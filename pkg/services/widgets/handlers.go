package widgets

import (
	"context"

	empty "google.golang.org/protobuf/types/known/emptypb"

	widgetsAPI "github.com/eccles/hestia/pkg/apis/widgets"
)

func (s *Service) Create(
	ctx context.Context,
	r *widgetsAPI.CreateRequest,
) (*widgetsAPI.Widget, error) {
	return &widgetsAPI.Widget{
		Name: r.Name,
	}, nil
}

func (s *Service) FindByID(
	ctx context.Context,
	r *widgetsAPI.FindRequest,
) (*widgetsAPI.Widget, error) {
	return &widgetsAPI.Widget{
		Uuid: r.Uuid,
	}, nil
}

func (s *Service) List(
	ctx context.Context,
	r *widgetsAPI.ListRequest,
) (*widgetsAPI.ListResponse, error) {
	return &widgetsAPI.ListResponse{}, nil
}

func (s *Service) Update(
	ctx context.Context,
	r *widgetsAPI.UpdateRequest,
) (*widgetsAPI.Widget, error) {
	return &widgetsAPI.Widget{
		Uuid: r.Uuid,
	}, nil
}

func (s *Service) Delete(
	ctx context.Context,
	r *widgetsAPI.DeleteRequest,
) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
