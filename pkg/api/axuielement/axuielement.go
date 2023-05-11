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
	value    string
	role     string
	Parent   *AXUIElementRef
	children []*AXUIElementRef
}

func (e *AXUIElementRef) GetID() string {
	return e.id
}

func (e *AXUIElementRef) GetRef() core.Ref {
	return e.ref
}

func GetAXUIElementRef(ref core.Ref, parent *AXUIElementRef) (*AXUIElementRef, error) {
	r := AXUIElementRef{}
	r.ref = ref
	hasChildren := hasChildren(ref)
	if hasChildren {
		r.children = getChildren(ref, &r)
	}
	id, err := getID(ref)
	if err != nil && len(r.children) == 0 {
		return nil, fmt.Errorf("elto has no id and no children")
	}
	r.id = id
	r.role = getRole(ref)
	if parent != nil {
		r.Parent = parent
	}
	// showActions(ref)
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

func (e *AXUIElementRef) FindElementByID(id string) (*AXUIElementRef, error) {
	fmt.Println("id", e.id)
	if len(e.id) > 0 && e.id == id {
		return e, nil
	}
	for _, child := range e.children {
		if element, err := child.FindElementByID(id); err == nil {
			return element, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func (e *AXUIElementRef) ShowElements() {
	// elements with empty id are parents holding children with ids or with children below
	if len(e.children) > 0 {
		fmt.Printf("Parent type %s with id %s and ref %v\n", e.role, e.id, e.ref)
		for _, child := range e.children {
			child.ShowElements()
		}
	} else {
		fmt.Printf("Element %s id %s \n", e.role, e.id)
	}
}

func (e *AXUIElementRef) Press() {
	// println(e.ref)
	// println(getRole(e.ref))
	C.Press(C.CFTypeRef(e.ref))
}

func hasChildren(ref core.Ref) bool {
	has, err := strconv.ParseBool(C.GoString(C.HasChildren(C.CFTypeRef(ref))))
	if err != nil {
		return false
	}
	return has
}

func getChildren(ref core.Ref, parent *AXUIElementRef) []*AXUIElementRef {
	var children []*AXUIElementRef
	childrenASCFArray := C.GetChildren(C.CFTypeRef(ref))
	count := C.CFArrayGetCount(childrenASCFArray)
	for i := 0; i < int(count); i++ {
		if child, err := GetAXUIElementRef(core.Ref(C.GetChild(childrenASCFArray, C.CFIndex(i))), parent); err == nil {
			children = append(children, child)
		}
	}
	return children
}

// TODO id should be only title or description
func getID(ref core.Ref) (string, error) {
	id := C.GoString(C.GetTitle(C.CFTypeRef(ref)))
	if len(id) == 0 {
		id = C.GoString(C.GetValue(C.CFTypeRef(ref)))
	}
	if len(id) == 0 {
		id = C.GoString(C.GetDescription(C.CFTypeRef(ref)))
	}
	if len(id) == 0 {
		return "", fmt.Errorf("object has no id")
	}
	return strings.TrimSpace(id), nil
}

func getRole(ref core.Ref) string {
	return C.GoString(C.GetRole(C.CFTypeRef(ref)))
}

func showActions(ref core.Ref) {
	C.ShowActions(C.CFTypeRef(ref))
}
