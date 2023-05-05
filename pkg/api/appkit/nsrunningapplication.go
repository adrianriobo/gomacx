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
	defaultDelay = 10 * time.Second
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
	r.focusedWindow.ShowElements()
}

func (r *NSRunningApplication) Click(buttonName string) error {
	button, err := r.focusedWindow.FindElementByName("AXButton", buttonName)
	if err != nil {
		return err
	}
	button.Click()
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
	r.focusedWindow, err = axuielement.GetAXUIElementRef(core.Ref(fwAXRef))
	return
}
