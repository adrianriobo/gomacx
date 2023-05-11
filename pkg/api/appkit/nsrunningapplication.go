//go:build darwin
// +build darwin

package appkit

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "nsrunningapplication.h"
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/adrianriobo/gomacx/pkg/api/axuielement"
	"github.com/adrianriobo/gomacx/pkg/core"
)

var (
	defaultDelay      = 3 * time.Second
	clickableElements = []string{"AXButton", "AXStaticText", "AXLink"}
)

// https://developer.apple.com/documentation/appkit/nsrunningapplication?language=objc
type NSRunningApplication struct {
	ref              unsafe.Pointer
	bundleIdentifier string
	axRef            core.Ref
	focusedWindow    *axuielement.AXUIElementRef
}

func ShowAllApplications() {
	C.ShowAllApplications()
}

// TODO windowTitle could be localized, how to handle that??
func GetAppByBundleAndWindow(bundleID, windowTitle string) (*NSRunningApplication, error) {
	cBundleID := C.CString(bundleID)
	defer C.free(unsafe.Pointer(cBundleID))
	cWindowTitle := C.CString(windowTitle)
	defer C.free(unsafe.Pointer(cWindowTitle))
	appRef := C.FindRunningApplication(cBundleID, cWindowTitle)
	if appRef == nil {
		return nil, fmt.Errorf("not found any app with bundle %s and window title %s", bundleID, windowTitle)
	}
	app := NSRunningApplication{
		ref: appRef}
	app.createAX()
	if err := app.loadFocusedWindow(); err != nil {
		fmt.Println(err)
	}
	return &app, nil
}

func GetApp() *NSRunningApplication {
	app := NSRunningApplication{
		ref: C.FrontmostApplication()}
	app.createAX()
	if err := app.loadFocusedWindow(); err != nil {
		fmt.Println(err)
	}
	return &app
}

func (r *NSRunningApplication) ShowElements() {
	r.loadFocusedWindow()
	r.focusedWindow.ShowElements()
}

// TODO something similar for role Checkbox
// ...Element AXCheckBox id Help improve Podman Desktop
func (r *NSRunningApplication) Click(id string) error {
	clickable, err := r.getClickable(id)
	if err != nil {
		return err
	}
	clickable.Press()
	time.Sleep(defaultDelay)
	return r.loadFocusedWindow()
}

func (r *NSRunningApplication) BundleIdentifier() string {
	if len(r.bundleIdentifier) == 0 {
		r.bundleIdentifier = C.GoString(C.BundleIdentifier(r.ref))
	}
	return r.bundleIdentifier
}

func (r *NSRunningApplication) createAX() {
	r.axRef = core.Ref(C.CreateApplicationAXRef(r.ref))
}

func (r *NSRunningApplication) loadFocusedWindow() (err error) {
	// Get the ax ui ref for the focused window
	fwAXRef := C.GetAXFocusedWindow(C.CFTypeRef(r.axRef))
	// Greate hierachy of elements
	r.focusedWindow, err = axuielement.GetAXUIElementRef(core.Ref(fwAXRef), nil)
	return
}

func (r *NSRunningApplication) getClickable(id string) (*axuielement.AXUIElementRef, error) {
	for _, ct := range clickableElements {
		clickable, err := r.focusedWindow.FindElementByName(ct, id)
		if err == nil {
			return clickable, nil
		}
	}
	return nil, fmt.Errorf("not found any clickable element with id %s", id)
}

func (r *NSRunningApplication) SetCheck(id, current, desired string) error {
	element, err := r.focusedWindow.FindElementByID(id)
	if err != nil {
		return fmt.Errorf("Can not find %s among the AX managed objects", id)
	}
	fmt.Println("got the holder", element.Parent.GetRef())
	check, err := element.Parent.FindElementByID(current)
	if err != nil {
		return fmt.Errorf("Can not find element with value %s within %s ", current, id)
	}
	check.Press()
	// Check after press we get the expected value
	// delay for refresh
	time.Sleep(defaultDelay)
	_, err = element.Parent.FindElementByID(desired)
	if err != nil {
		return fmt.Errorf("Set on %s has not set the expected %s state", id, desired)
	}
	return nil
}

// func (r *NSRunningApplication) GetGroupHolding(id string) (*axuielement.AXUIElementRef, error) {
// 	return r.focusedWindow.GetGroupHolding(id)
// }

// func (r *NSRunningApplication) EnableDisable(elementName string) error {
// 	_, err := r.focusedWindow.GetEnableDisableElement(elementName)
// 	if err != nil {
// 		fmt.Println("not found disabled enabled")
// 		return err
// 	}
// 	fmt.Println("found disabled enabled")
// 	return nil
// }
