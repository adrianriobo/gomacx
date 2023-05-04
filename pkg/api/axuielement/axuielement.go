//go:build darwin
// +build darwin

package axuielement

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "axuielement.h"
import "C"
import (
	"fmt"

	"github.com/adrianriobo/gomacx/pkg/core"
)

type AXUIElementRef struct {
	ref      core.Ref
	title    string
	role     string
	value    string
	children []*AXUIElementRef
}

func GetAXUIElementRef(ref core.Ref) *AXUIElementRef {
	r := AXUIElementRef{}
	has := C.GoString(C.HasAXUIElementChildren(C.CFTypeRef(ref)))
	fmt.Println(has)
	if has == "true" {
		childrenASCFArray := C.GetAXUIElementChildren(C.CFTypeRef(ref))
		fmt.Println("before counting children")
		count := C.CFArrayGetCount(childrenASCFArray)
		fmt.Println("after counting children")
		r.children = make([]*AXUIElementRef, count)
		for i := 0; i < int(count); i++ {
			r.children[i] = GetAXUIElementRef(core.Ref(C.GetChild(childrenASCFArray, C.CFIndex(i))))
		}
	}
	r.role = C.GoString(C.GetRoleAttribute(C.CFTypeRef(ref)))
	r.title = C.GoString(C.GetTitleAttribute(C.CFTypeRef(ref)))
	r.value = C.GoString(C.GetValueAttribute(C.CFTypeRef(ref)))
	return &r
}
