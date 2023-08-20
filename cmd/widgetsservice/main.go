package main

import (
	"github.com/eccles/hestia/pkg/startup"
	widgetsservice "github.com/eccles/hestia/pkg/widgets/service"
)

func main() {
	startup.Run("widgets", widgetsservice.Run)
}
