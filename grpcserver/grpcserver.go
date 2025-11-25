package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_otrace "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/eccles/hestia/grpchealth"
)

// so we dont have to import grpc when using this package.
type grpcServer = grpc.Server
type grpcUnaryServerInterceptor = grpc.UnaryServerInterceptor

type RegisterServer func(*grpcServer)

func defaultRegisterServer(g *grpcServer) {}

type GRPCServer struct {
	name          string
	log           Logger
	listenStr     string
	health        bool
	healthService *grpchealth.HealthCheckingService
	interceptors  []grpcUnaryServerInterceptor
	register      RegisterServer
	server        *grpcServer
	reflection    bool
}

type GRPCServerOption func(*GRPCServer)

func WithAppendedInterceptor(i grpcUnaryServerInterceptor) GRPCServerOption {
	return func(g *GRPCServer) {
		g.interceptors = append(g.interceptors, i)
	}
}

func WithPrependedInterceptor(i grpcUnaryServerInterceptor) GRPCServerOption {
	return func(g *GRPCServer) {
		g.interceptors = append([]grpcUnaryServerInterceptor{i}, g.interceptors...)
	}
}

func WithRegisterServer(r RegisterServer) GRPCServerOption {
	return func(g *GRPCServer) {
		g.register = r
	}
}

func WithoutHealth() GRPCServerOption {
	return func(g *GRPCServer) {
		g.health = false
	}
}

func WithReflection(r bool) GRPCServerOption {
	return func(g *GRPCServer) {
		g.reflection = r
	}
}

func tracingFilter(ctx context.Context, fullMethodName string) bool {
	return fullMethodName != grpcHealth.Health_Check_FullMethodName
}

// New creates a new GRPCServer that is bound to a specific GRPC API. This object complies with
// the standard Listener service and can be managed by the startup.Listeners object.
func New(log Logger, name string, port string, opts ...GRPCServerOption) GRPCServer {
	g := GRPCServer{
		name:      strings.ToLower(name),
		listenStr: fmt.Sprintf(":%s", port),
		register:  defaultRegisterServer,
		interceptors: []grpc.UnaryServerInterceptor{
			grpc_otrace.UnaryServerInterceptor(grpc_otrace.WithFilterFunc(tracingFilter)),
			grpc_validator.UnaryServerInterceptor(),
		},
		health: true,
	}
	for _, opt := range opts {
		opt(&g)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(g.interceptors...),
		),
	)

	g.register(server)

	if g.health {
		healthService := grpchealth.New(log)
		g.healthService = &healthService
		grpcHealth.RegisterHealthServer(server, g.healthService)
	}

	if g.reflection {
		reflection.Register(server)
	}

	g.server = server
	g.log = log.With("grpcserver", g.String())
	return g
}

func (g *GRPCServer) String() string {
	// No logging in this method please.
	return fmt.Sprintf("%s%s", g.name, g.listenStr)
}

func (g *GRPCServer) Listen() error {
	listen, err := net.Listen("tcp", g.listenStr)
	if err != nil {
		return fmt.Errorf("failed to listen %s: %w", g, err)
	}

	if g.healthService != nil {
		g.healthService.Ready() // readiness
	}

	g.log.Info("Listen")
	err = g.server.Serve(listen)
	if err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("failed to serve %s: %w", g, err)
	}
	return nil
}

func (g *GRPCServer) Shutdown(_ context.Context) error {
	g.log.Info("Shutdown")
	if g.healthService != nil {
		g.healthService.NotReady() // readiness
		g.healthService.Dead()     // liveness
	}
	g.server.GracefulStop()
	return nil
}
