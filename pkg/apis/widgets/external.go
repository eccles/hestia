package widgetsapi

// External version of Widget struct without the proto overhead.
// This struct must be kept in synchronization with the proto
// message.
type WidgetExternal struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
}

// WidgetToExternal creates and external representation of the proto
// message without the proto members.
func WidgetToExternal(w *Widget) WidgetExternal {
	return WidgetExternal{
		Uuid: w.Uuid,
		Name: w.Name,
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
