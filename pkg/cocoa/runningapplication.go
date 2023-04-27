package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
//#include "runningapplication.h"
import "C"
import "unsafe"

type RunningApplication struct {
	ptr unsafe.Pointer
}

func GetRunningApplication() *RunningApplication {
	app := new(RunningApplication)
	app.ptr = C.GetRunningApplications()
	return app
}

func (r *RunningApplication) GetBundleID() string {
	return C.GoString(C.GetBundleID(r.ptr))
}
