package main

import (
	_ "embed"
)

//go:generate ../../scripts/version.sh cmd/widgetsservice
//go:embed version.txt
var version string
