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
	"net"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	widgetsapi "github.com/eccles/hestia/apis/widgets"
)

// GRPCService represents an incoming connection i.e a server that handles GRPC messages
// sent from a client.
type GRPCService struct {
	Server *grpc.Server
	Port   string
}

func (g *GRPCService) Stop() {
	g.Server.GracefulStop()
}

// StartGRPCService initialises the GRPC service structs generated in thee qpis
// package and starts the GRPC service.
func (s *Service) StartGRPCService() error {
	s.Log.Info("Start GRPCService")

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			validator.UnaryServerInterceptor(),
		)),
	)

	widgetsapi.RegisterWidgetsServer(grpcServer, s) // s is of type widgetsapi.WidgetsServer
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":"+s.GRPC.Port)

	if err != nil {
		return fmt.Errorf("listen ':%s' failure: %w", s.GRPC.Port, err)
	}

	s.GRPC.Server = grpcServer

	g := new(errgroup.Group)

	g.Go(func() error {
		err = s.GRPC.Server.Serve(listen)
		if err != nil {
			s.Log.Info("Failed to start")
			return err
		}

		return nil
	})

	return g.Wait()
}
