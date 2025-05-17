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

package widgetsapi

// External version of Widget struct without the proto overhead.
// This struct must be kept in synchronisation with the proto
// message.
type WidgetExternal struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// WidgetToExternal creates and external representation of the proto
// message without the proto members.
func WidgetToExternal(w *Widget) WidgetExternal {
	return WidgetExternal{
		Uuid: w.GetUuid(),
		Name: w.GetName(),
	}
}

// WidgetFromExternal creates a prto message from a vanilla copy
// of the widgets record.
func WidgetFromExternal(w *WidgetExternal) Widget {
	return Widget{
		Uuid: w.Uuid,
		Name: w.Name,
	}
}
