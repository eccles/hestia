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

package logger

// Creates global logger from which all loggers are derived.
import (
	"log/slog"
	"os"
)

const (
	serviceLogKey = "service"
)

var RootLogger *slog.Logger //nolint:gochecknoglobals // all loggers derive from this

func New(level string) {
	var options slog.HandlerOptions
	if level == "DEBUG" {
		options.AddSource = true
		options.Level = slog.LevelDebug
		RootLogger = slog.New(slog.NewTextHandler(os.Stderr, &options))
	} else {
		options.Level = slog.LevelInfo
		RootLogger = slog.New(slog.NewJSONHandler(os.Stdout, &options))
	}
}

// close is a dummy function at the moment - it should be deferred.
func Close() {
}

func WithServiceName(serviceName string) *slog.Logger {
	return RootLogger.With(serviceLogKey, serviceName)
}
