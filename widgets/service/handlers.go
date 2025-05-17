// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package widgetsservice

import (
	"context"

	empty "google.golang.org/protobuf/types/known/emptypb"

	widgetsapi "github.com/eccles/hestia/apis/widgets"
)

func (s *Service) Create(
	ctx context.Context,
	r *widgetsapi.CreateRequest,
) (*widgetsapi.Widget, error) {

	return &widgetsapi.Widget{
		Name: r.GetName(),
	}, nil
}

func (s *Service) FindByID(
	ctx context.Context,
	r *widgetsapi.FindRequest,
) (*widgetsapi.Widget, error) {

	return &widgetsapi.Widget{
		Uuid: r.GetUuid(),
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
		Uuid: r.GetUuid(),
	}, nil
}

func (s *Service) Delete(
	ctx context.Context,
	_ *widgetsapi.DeleteRequest,
) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}
