package widgetsAPI

//
// Generate Go code from protobuf definition file.
//go:generate ../../../scripts/protoc.sh pkg/apis/widgets/widgets.proto

// External version of Widget struct without the proto overhead.
type WidgetExternal struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

func WidgetToExternal(w *Widget) WidgetExternal {
	return WidgetExternal{
		Uuid: w.Uuid,
		Name: w.Name,
	}
}

func WidgetFromExternal(w *WidgetExternal) Widget {
	return Widget{
		Uuid: w.Uuid,
		Name: w.Name,
	}
}
