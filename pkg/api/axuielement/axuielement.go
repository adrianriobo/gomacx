//go:build darwin
// +build darwin

package axuielement

// #cgo darwin CFLAGS:-mmacosx-version-min=11 -x objective-c
// #cgo darwin LDFLAGS: -lobjc -framework Foundation -framework Cocoa
//#include "axuielement.h"
import "C"
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adrianriobo/gomacx/pkg/core"
)

type AXUIElementRef struct {
	ref      core.Ref
	id       string
	role     string
	children []*AXUIElementRef
}

func GetAXUIElementRef(ref core.Ref) (*AXUIElementRef, error) {
	r := AXUIElementRef{}
	r.ref = ref
	hasChildren := hasChildren(ref)
	if hasChildren {
		r.children = getChildren(ref)
	}
	id, err := getID(ref)
	if err != nil && len(r.children) == 0 {
		return nil, fmt.Errorf("elto has no id and no children")
	}
	r.id = id
	r.role = getRole(ref)
	return &r, nil
}

func (e *AXUIElementRef) FindElementByName(role, id string) (*AXUIElementRef, error) {
	if e.role == role && e.id == id {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByName(role, id); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (e *AXUIElementRef) ShowElements() {
	// elements with empty id are parents holding children with ids or with children below
	if len(e.id) > 0 {
		fmt.Printf("Element type %s with id %s and ref %v\n", e.role, e.id, e.ref)
	}
	for _, child := range e.children {
		child.ShowElements()
	}
}

func (e *AXUIElementRef) Click() {
	println(e.ref)
	println(getRole(e.ref))
	C.ClickButton(C.CFTypeRef(e.ref))
}

func hasChildren(ref core.Ref) bool {
	has, err := strconv.ParseBool(C.GoString(C.HasAXUIElementChildren(C.CFTypeRef(ref))))
	if err != nil {
		return false
	}
	return has
}

func getChildren(ref core.Ref) []*AXUIElementRef {
	var children []*AXUIElementRef
	childrenASCFArray := C.GetAXUIElementChildren(C.CFTypeRef(ref))
	count := C.CFArrayGetCount(childrenASCFArray)
	for i := 0; i < int(count); i++ {
		if child, err := GetAXUIElementRef(core.Ref(C.GetChild(childrenASCFArray, C.CFIndex(i)))); err == nil {
			children = append(children, child)
		}
	}
	return children
}

func getID(ref core.Ref) (string, error) {
	id := C.GoString(C.GetTitleAttribute(C.CFTypeRef(ref)))
	if len(id) == 0 {
		id = C.GoString(C.GetValueAttribute(C.CFTypeRef(ref)))
	}
	if len(id) == 0 {
		return "", fmt.Errorf("object has no id")
	}
	return strings.TrimSpace(id), nil
}

func getRole(ref core.Ref) string {
	return C.GoString(C.GetRoleAttribute(C.CFTypeRef(ref)))
}
