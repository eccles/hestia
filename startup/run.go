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

package startup

import (
	"errors"
	"os"

	"github.com/eccles/hestia/logger"
)

type Runner func(string, Logger) error

var ErrLogLevelUndefined = errors.New("no loglevel specified")

func getLogLevel() string {
	value, ok := os.LookupEnv("LOGLEVEL")

	if !ok {
		panic(ErrLogLevelUndefined)
	}

	return value
}

// Run() is run from the main() method in order to isolate
// code that executes os.Exit(). This enables defers in the run
// function in the second argument.
func Run(serviceName string, run Runner) {
	var exitCode int

	logger.New(getLogLevel())

	log := logger.WithServiceName(serviceName)

	err := run(serviceName, log)

	if err != nil {
		log.Info("Error terminating", "err", err)
		exitCode = 1
	}

	log.Info("Shutting down")
	logger.Close()
	os.Exit(exitCode)
}
