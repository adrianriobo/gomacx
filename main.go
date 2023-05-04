//go:build darwin
// +build darwin

package main

import (
	"fmt"

	"github.com/adrianriobo/gomacx/pkg/api/appkit"
)

// const (
// 	mainPane                 string = "Podman Desktop"
// 	checkEnableTelemetryName string = "Enable telemetry"
// 	buttonName               string = "Go to Podman Desktop"
// )

func main() {
	app := appkit.GetApp()
	fmt.Println(app.BundleIdentifier())
}
