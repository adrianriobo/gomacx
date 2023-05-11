package bussiness

import (
	"fmt"

	"github.com/adrianriobo/gomacx/pkg/api/appkit"
)

type Application struct {
	ref *appkit.NSRunningApplication
}

func Get() *Application {
	return &Application{
		ref: appkit.GetApp()}
}

func (a *Application) Click(buttonName string) error {
	return fmt.Errorf("not implemented yet")
}
