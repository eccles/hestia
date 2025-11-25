package startup

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrNilListener      = errors.New("nil Listener")
	ErrNilListenerValue = errors.New("nil Listener value")
)

// based on gist found at
// https://gist.github.com/pteich/c0bb58b0b7c8af7cc6a689dd0d3d26ef?permalink_comment_id=4053701

// Listener is an interface that describes any kind of listener - HTTP Server, GRPC Server
// or servicebus receiver.
type Listener interface {
	Listen() error
	Shutdown(context.Context) error
}

// Listeners contains all servers that comply with the service.
type Listeners struct {
	name      string
	log       Logger
	listeners []Listener
}

type ListenersOption func(*Listeners)

// WithListeners add multiple listeners. Nil listeners will cause
// an error to be returned.
func WithListeners(listeners ...Listener) ListenersOption {
	return func(l *Listeners) {
		l.listeners = append(l.listeners, listeners...)
	}
}

func NewListeners(log Logger, name string, opts ...ListenersOption) Listeners {
	l := Listeners{name: strings.ToLower(name)}
	for _, opt := range opts {
		opt(&l)
	}
	l.log = log.With("listener", l.String())
	return l
}

func (l *Listeners) String() string {
	return l.name
}

func (l *Listeners) Listen() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, errCtx := errgroup.WithContext(ctx)

	for i := range l.listeners {
		h := l.listeners[i]
		if h == nil {
			return ErrNilListener
		}
		if reflect.ValueOf(h).IsNil() {
			return ErrNilListenerValue
		}
		l.log.Debug("Start %d %s", i, h)
		g.Go(func() error {
			err := h.Listen()
			if err != nil {
				return err
			}
			return nil
		})
	}

	g.Go(func() error {
		<-errCtx.Done()
		l.log.Info("Cancel from signal")
		return l.Shutdown()
	})

	err := g.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (l *Listeners) Shutdown() error {
	var err error
	for _, h := range l.listeners {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			e := h.Shutdown(ctx)
			if e != nil {
				if err != nil {
					err = fmt.Errorf("cannot shutdown %s: %w: %w", h, err, e)
				} else {
					err = fmt.Errorf("cannot shutdown %s: %w", h, e)
				}
			}
		}()
	}
	return err
}
