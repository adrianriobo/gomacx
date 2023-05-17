//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"os"

	"github.com/adrianriobo/gomacx/pkg/bussiness"
)

func main() {
	podmanDesktopApp, err := bussiness.Initialize("/Users/crcqe/Desktop/Podman Desktop.app")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = podmanDesktopApp.DisableTelemetry(); err != nil {
		fmt.Println(err)
	}
	if err = podmanDesktopApp.InstallPodman("userPassword"); err != nil {
		fmt.Println(err)
	}
}
