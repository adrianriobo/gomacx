//go:build darwin
// +build darwin

package appkit

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "nsrunningapplication.h"
import "C"
import (
	"unsafe"

	"github.com/adrianriobo/gomacx/pkg/api/axuielement"
	"github.com/adrianriobo/gomacx/pkg/core"
)

// https://developer.apple.com/documentation/appkit/nsrunningapplication?language=objc
type NSRunningApplication struct {
	ref              unsafe.Pointer
	bundleIdentifier string
	axRef            core.Ref
	focusedWindow    *axuielement.AXUIElementRef
}

func GetApp() *NSRunningApplication {
	app := NSRunningApplication{
		ref: C.FrontmostApplication()}
	app.createAX()
	app.loadFocusedWindow()
	return &app
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

func (r *NSRunningApplication) loadFocusedWindow() {
	// Get the ax ui ref for the focused window
	fwAXRef := C.GetAXFocusedWindow(C.CFTypeRef(r.axRef))
	// Greate hierachy of elements
	r.focusedWindow = axuielement.GetAXUIElementRef(core.Ref(fwAXRef))
}
