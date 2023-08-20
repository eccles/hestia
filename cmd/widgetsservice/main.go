package main

import (
	"github.com/eccles/hestia/pkg/startup"
	widgetsService "github.com/eccles/hestia/pkg/widgets/service"
)

func main() {
	startup.Run("widgets", widgetsService.Run)
}
