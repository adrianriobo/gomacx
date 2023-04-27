package main

import (
	"fmt"
	"github.com/adrianriobo/gomacx/pkg/cocoa"
)

const (
	mainPane                 string = "Podman Desktop"
	checkEnableTelemetryName string = "Enable telemetry"
	buttonName               string = "Go to Podman Desktop"
)

func main() {
	a := cocoa.GetRunningApplication()
	b := a.GetBundleID()
	fmt.Println(b)
}
