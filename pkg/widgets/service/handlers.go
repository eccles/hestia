package widgetsservice

import (
	"context"

	empty "google.golang.org/protobuf/types/known/emptypb"

	widgetsapi "github.com/eccles/hestia/pkg/apis/widgets"
)

func (s *Service) Create(
	ctx context.Context,
	r *widgetsapi.CreateRequest,
) (*widgetsapi.Widget, error) {
	return &widgetsapi.Widget{
		Name: r.Name,
	}, nil
}

func (s *Service) FindByID(
	ctx context.Context,
	r *widgetsapi.FindRequest,
) (*widgetsapi.Widget, error) {
	return &widgetsapi.Widget{
		Uuid: r.Uuid,
	}, nil
}

func (s *Service) List(
	ctx context.Context,
	r *widgetsapi.ListRequest,
) (*widgetsapi.ListResponse, error) {
	return &widgetsapi.ListResponse{}, nil
}

func (s *Service) Update(
	ctx context.Context,
	r *widgetsapi.UpdateRequest,
) (*widgetsapi.Widget, error) {
	return &widgetsapi.Widget{
		Uuid: r.Uuid,
	}, nil
}

func (s *Service) Delete(
	ctx context.Context,
	r *widgetsapi.DeleteRequest,
) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
