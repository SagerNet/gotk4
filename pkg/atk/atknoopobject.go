// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_no_op_object_get_type()), F: marshalNoOpObjector},
	})
}

// NoOpObject is an AtkObject which purports to implement all ATK interfaces. It
// is the type of AtkObject which is created if an accessible object is
// requested for an object type for which no factory type is specified.
type NoOpObject struct {
	ObjectClass

	Action
	Component
	Document
	EditableText
	Hypertext
	Image
	Selection
	Table
	TableCell
	Text
	Value
	Window
	*externglib.Object
}

func WrapNoOpObject(obj *externglib.Object) *NoOpObject {
	return &NoOpObject{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Action: Action{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
		Document: Document{
			Object: obj,
		},
		EditableText: EditableText{
			Object: obj,
		},
		Hypertext: Hypertext{
			Object: obj,
		},
		Image: Image{
			Object: obj,
		},
		Selection: Selection{
			Object: obj,
		},
		Table: Table{
			Object: obj,
		},
		TableCell: TableCell{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
		Text: Text{
			Object: obj,
		},
		Value: Value{
			Object: obj,
		},
		Window: Window{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalNoOpObjector(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNoOpObject(obj), nil
}

// NewNoOpObject provides a default (non-functioning stub) Object. Application
// maintainers should not use this method.
func NewNoOpObject(obj *externglib.Object) *NoOpObject {
	var _arg1 *C.GObject   // out
	var _cret *C.AtkObject // in

	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_no_op_object_new(_arg1)
	runtime.KeepAlive(obj)

	var _noOpObject *NoOpObject // out

	_noOpObject = WrapNoOpObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noOpObject
}

func (*NoOpObject) privateNoOpObject() {}
