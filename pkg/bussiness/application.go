package bussiness

import "fmt"

type Application struct {
	ExecPath string
}

func (a *Application) Click(buttonName string) error {
	return fmt.Errorf("not implemented yet")
}
