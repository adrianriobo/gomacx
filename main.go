//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"

	"github.com/adrianriobo/gomacx/pkg/api/appkit"
)

// const (
// 	mainPane                 string = "Podman Desktop"
// 	checkEnableTelemetryName string = "Enable telemetry"
// 	buttonName               string = "Go to Podman Desktop"
// )

func main() {
	// TODO check why we need to get the app like this to see the installer as nsrunning app
	pdApp := appkit.GetApp()
	// pdApp.ShowElements()
	if err := pdApp.Click("Settings"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// pdApp.ShowElements()
	if err := pdApp.Click("Preferences"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// pdApp.ShowElements()
	if err := pdApp.Click("Telemetry"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	pdApp.ShowElements()
	// if err := pdApp.SetCheck("Telemetry", "Enabled", "Disabled"); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	// pdApp, err := appkit.GetAppByBundleAndWindow("io.podmandesktop.PodmanDesktop", "Podman Desktop")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(pdApp.BundleIdentifier())

	// pdApp.Click("Settings")
	// pdApp.ShowElements()
	// pdApp.EnableDisable("Telemetry")

}

// func clickInstall(pdApp *appkit.NSRunningApplication) {
// 	pdApp.Click("Install")
// 	pdApp.ShowElements()
// 	pdApp.Click("Yes")

// 	piApp, err := appkit.GetAppByBundleAndWindow("com.apple.installer", "Install Podman")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	piApp.ShowElements()
// }
