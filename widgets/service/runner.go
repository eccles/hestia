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

// Widgets - a GRPC Service that implements the widgetsapi.WidgetsServer interface.
// This file (runner.go must be defined manually) .
//
// Additional files to be created manually are handlers.go and connetions.go.
// All other fioes are generated using 'go generate' commands

import (
	"fmt"
	"os"
)

func port() string {
	val, ok := os.LookupEnv("GRPC_SERVICE_PORT")

	if !ok {
		val = ""
	}

	return val
}

// Run initialises the Service struct and executes its Run method.
// The Service struct specifies the grpc server code and any other interfaces
// to external services defined in connect.go.
func Run(serviceName string, log Logger) error {
	var err error

	s := Service{
		Log: log,
		GRPC: GRPCService{
			Port: port(),
		},
	}

	err = s.Run()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
