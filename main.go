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
	pdApp := appkit.GetApp()
	fmt.Println(pdApp.BundleIdentifier())
	pdApp.ShowElements()
	pdApp.Click("Install")
	pdApp.ShowElements()
	pdApp.Click("Yes")
	// This will run a installer for podman so we need to pick it
	// piApp := appkit.GetApp()
	// fmt.Println(piApp.BundleIdentifier())
	// piApp.ShowElements()
	// appkit.ShowAllApplications()
	piApp := appkit.GetAppByBundleAndWindow("com.apple.installer", "Install Podman")
	piApp.ShowElements()
}
