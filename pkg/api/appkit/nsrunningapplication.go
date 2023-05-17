//go:build darwin
// +build darwin

package appkit

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "nsrunningapplication.h"
import "C"
import (
	"fmt"
	"os/exec"
	"time"
	"unsafe"

	"github.com/adrianriobo/gomacx/pkg/api/axuielement"
	"github.com/adrianriobo/gomacx/pkg/core"
)

var (
	defaultDelay      = 5 * time.Second
	longDelay         = 15 * time.Second
	clickableElements = []string{"AXButton", "AXStaticText", "AXLink"}
	checkableElements = []string{"AXCheckBox"}
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
	if err := app.LoadFocusedWindow(); err != nil {
		fmt.Println(err)
	}
	return &app, nil
}

func GetApp(applicationPath string) (*NSRunningApplication, error) {
	if err := openApplication(applicationPath); err != nil {
		return nil, fmt.Errorf("error getting app %s", applicationPath, err)
	}
	time.Sleep(defaultDelay)
	// Get the fromtmost application
	app := NSRunningApplication{
		ref: C.FrontmostApplication()}
	// create an ax element to access the app
	app.createAX()
	time.Sleep(defaultDelay)
	// Load all AX elements for the app
	if err := app.LoadFocusedWindow(); err != nil {
		return nil, fmt.Errorf("error getting app %s", applicationPath, err)
	}
	return &app, nil
}

func GetFrontmostApp() (*NSRunningApplication, error) {
	// Get the fromtmost application
	app := NSRunningApplication{
		ref: C.FrontmostApplication()}
	// create an ax element to access the app
	app.createAX()
	time.Sleep(defaultDelay)
	// Load all AX elements for the app
	if err := app.LoadFocusedWindow(); err != nil {
		return nil, fmt.Errorf("error getting frontmost app", err)
	}
	return &app, nil
}

func openApplication(path string) error {
	cmd := exec.Command("open", path)
	return cmd.Start()
}

func (r *NSRunningApplication) ShowElements() {
	r.LoadFocusedWindow()
	r.focusedWindow.ShowElements()
}

// TODO something similar for role Checkbox
// ...Element AXCheckBox id Help improve Podman Desktop
func (r *NSRunningApplication) Click(id string) error {
	clickable, err := r.getElementbyRoleAndID(id, clickableElements)
	if err != nil {
		return err
	}
	return r.pressElement(clickable)
}

func (r *NSRunningApplication) Check(id string) error {
	clickable, err := r.getElementbyRoleAndID(id, checkableElements)
	if err != nil {
		return err
	}
	return r.pressElement(clickable)
}

func (r *NSRunningApplication) pressElement(element *axuielement.AXUIElementRef) error {
	element.Press()
	time.Sleep(longDelay)
	return r.LoadFocusedWindow()
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

func (r *NSRunningApplication) LoadFocusedWindow() (err error) {
	// Get the ax ui ref for the focused window
	fwAXRef := C.GetAXFocusedWindow(C.CFTypeRef(r.axRef))
	// Greate hierachy of elements
	r.focusedWindow, err = axuielement.GetAXUIElementRef(core.Ref(fwAXRef), nil)
	return
}

func (r *NSRunningApplication) getElementbyRoleAndID(id string, elementTypes []string) (*axuielement.AXUIElementRef, error) {
	for _, ct := range elementTypes {
		clickable, err := r.focusedWindow.FindElementByRoleAndID(ct, id)
		if err == nil {
			return clickable, nil
		}
	}
	return nil, fmt.Errorf("not found any clickable element with id %s", id)
}

// check is a label (AXStaticText) and checkbox (AXCheckBox)
func (r *NSRunningApplication) SetCheck(current, desired string) error {
	check, err := r.focusedWindow.FindElementByRole("AXCheckBox")
	if err != nil {
		return fmt.Errorf("Can not find any AXCheckBox element")
	}
	_, err = check.Parent.FindElementByID(current)
	if err != nil {
		return fmt.Errorf("Can not find element with value %s ", current)
	}
	check.Press()
	time.Sleep(defaultDelay)
	err = r.LoadFocusedWindow()
	if err != nil {
		return fmt.Errorf("error re loading the focused window", err)
	}
	check, err = r.focusedWindow.FindElementByRole("AXCheckBox")
	if err != nil {
		return fmt.Errorf("Can not find any AXCheckBox element")
	}
	_, err = check.Parent.FindElementByID(desired)
	if err != nil {
		return fmt.Errorf("Error checking the desired state %s ", desired, err)
	}
	return nil
}

func (r *NSRunningApplication) InstallExtension(oci string) error {
	button, err := r.focusedWindow.FindElementByRoleAndID("AXButton", "Install extension from the OCI image")
	if err != nil {
		return fmt.Errorf("Can not find button for install new extension from OCI")
	}
	// fmt.Println("el button es ", button.GetID(), button.Role())
	// fmt.Println("y el parent", button.Parent.GetID(), button.Parent.Role(), button.Parent.GetRef(), len(button.Parent.Children()))
	// button.Parent.ShowElements()
	textField, err := button.Parent.FindElementByRole("AXTextField")
	if err != nil {
		return fmt.Errorf("Can not find text field for install new extension from OCI")
	}
	textField.SetValue(oci)
	time.Sleep(defaultDelay)
	button.Press()
	time.Sleep(longDelay)
	return r.LoadFocusedWindow()
}

// This function is used to click on install button offered after installing the
// openshift local extension on systems without any previous openshift local installation
func (r *NSRunningApplication) InstallOpenshiftLocal() error {
	extensionHeaders, err := r.focusedWindow.FindElementsByRoleAndID("AXStaticText", "OpenShift Local")
	// fmt.Println("number of extension headers", len(extensionHeaders))
	if err != nil {
		return fmt.Errorf("Can not find section for install for Openshift Local")
	}
	var installButton *axuielement.AXUIElementRef
	for _, extensionHeader := range extensionHeaders {
		// fmt.Println("header role", extensionHeader.Role(), len(extensionHeader.Children()))
		// installButton, err = extensionHeader.Parent.Parent.FindElementByRoleAndID("AXButton", "Install")
		// if err == nil {
		// 	break
		// }
		extensionHeader.Parent.Parent.ShowElements()
	}
	if installButton == nil {
		return fmt.Errorf("Can not find button for installing the Openshift Local extension")
	}
	installButton.Press()
	time.Sleep(defaultDelay)
	return r.LoadFocusedWindow()
}
