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
	"fmt"

	widgetsapi "github.com/eccles/hestia/apis/widgets"
)

// implements the widgetsapi.WidgetsServer interfacw.
type Service struct {
	widgetsapi.UnimplementedWidgetsServer

	// An interface as we may want to mock it out in tests.
	Log Logger

	// A concrete implementation of a GRPC service.
	GRPC GRPCService
}

func (s *Service) Run() error {
	err := s.StartGRPCService()

	if err != nil {
		return fmt.Errorf("grpcservice start failure: %w", err)
	}

	defer s.GRPC.Stop()

	return s.Connect()
}
