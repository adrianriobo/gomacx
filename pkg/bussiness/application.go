package bussiness

import (
	"fmt"

	"github.com/adrianriobo/gomacx/pkg/api/appkit"
	"github.com/adrianriobo/gomacx/pkg/api/applescript"
)

const (
	MenuDashboard            = "Dashboard"
	MenuSettings             = "Settings"
	MenuExtensions           = "Extensions"
	MenuPreferences          = "Preferences"
	MenuPreferencesTelemetry = "Telemetry"

	ButtonInstall = "Install"
	ButtonYes     = "Yes"

	installerBundleID       = "com.apple.installer"
	installerPodmanTitle    = "Install Podman"
	installerButtonContinue = "Continue"
	installerButtonAgree    = "Agree"
	installerButtonInstall  = "Install"
	installerButtonClose    = "Close"
)

type PodmanDesktopApp struct {
	app *appkit.NSRunningApplication
}

func Initialize(appPath string) (*PodmanDesktopApp, error) {
	app, err := appkit.GetApp(appPath)
	if err != nil {
		return nil, err
	}
	return &PodmanDesktopApp{
		app: app}, nil
}

func (a *PodmanDesktopApp) DisableTelemetry() error {
	if err := a.app.Click(MenuSettings); err != nil {
		return fmt.Errorf("Error disabling telemetry", err)
	}
	if err := a.app.Click(MenuPreferences); err != nil {
		return fmt.Errorf("Error disabling telemetry", err)
	}
	if err := a.app.Click(MenuPreferencesTelemetry); err != nil {
		return fmt.Errorf("Error disabling telemetry", err)
	}
	if err := a.app.SetCheck("Enabled", "Disabled"); err != nil {
		return fmt.Errorf("Error disabling telemetry", err)
	}
	return nil
}

// This function is expected to be run at the very beggining
// when podman is the only extension to be installed
func (a *PodmanDesktopApp) InstallPodman(userPassword string) error {
	if err := a.app.Click(MenuDashboard); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	if err := a.app.Click(ButtonInstall); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	if err := a.app.Click(ButtonYes); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Installer is open on a different app, we need to pick it
	iApp, err := appkit.GetAppByBundleAndWindow(installerBundleID, installerPodmanTitle)
	if err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Click introducction
	if err := iApp.Click(installerButtonContinue); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Click License
	if err := iApp.Click(installerButtonContinue); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Click Agree License
	if err := iApp.Click(installerButtonAgree); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Click Install
	if err := iApp.Click(installerButtonInstall); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	// Need to type the user password
	applescript.Keystroke(userPassword)
	if err := iApp.LoadFocusedWindow(); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	if err := iApp.Click(installerButtonClose); err != nil {
		return fmt.Errorf("Error installing podman", err)
	}
	return nil
}

// test ghcr.io/crc-org/crc-extension:latest
func (a *PodmanDesktopApp) InstallExtension(extensionsOCI string) error {
	if err := a.app.Click(MenuSettings); err != nil {
		return fmt.Errorf("Error installing extension", err)
	}
	if err := a.app.Click(MenuExtensions); err != nil {
		return fmt.Errorf("Error installing extension", err)
	}
	if err := a.app.InstallExtension(extensionsOCI); err != nil {
		return fmt.Errorf("Error installing extension", err)
	}
	return nil
}
