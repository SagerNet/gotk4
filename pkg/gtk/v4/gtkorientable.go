// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_orientable_get_type()), F: marshalOrientable},
	})
}

// Orientable: the Orientable interface is implemented by all widgets that can
// be oriented horizontally or vertically. Orientable is more flexible in that
// it allows the orientation to be changed at runtime, allowing the widgets to
// “flip”.
type Orientable interface {
	gextras.Objector

	// Orientation retrieves the orientation of the @orientable.
	Orientation() Orientation
	// SetOrientation sets the orientation of the @orientable.
	SetOrientation(orientation Orientation)
}

// orientable implements the Orientable interface.
type orientable struct {
	gextras.Objector
}

var _ Orientable = (*orientable)(nil)

// WrapOrientable wraps a GObject to a type that implements interface
// Orientable. It is primarily used internally.
func WrapOrientable(obj *externglib.Object) Orientable {
	return Orientable{
		Objector: obj,
	}
}

func marshalOrientable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOrientable(obj), nil
}

// Orientation retrieves the orientation of the @orientable.
func (o orientable) Orientation() Orientation {
	var arg0 *C.GtkOrientable

	arg0 = (*C.GtkOrientable)(unsafe.Pointer(o.Native()))

	var cret C.GtkOrientation
	var goret1 Orientation

	cret = C.gtk_orientable_get_orientation(arg0)

	goret1 = Orientation(cret)

	return goret1
}

// SetOrientation sets the orientation of the @orientable.
func (o orientable) SetOrientation(orientation Orientation) {
	var arg0 *C.GtkOrientable
	var arg1 C.GtkOrientation

	arg0 = (*C.GtkOrientable)(unsafe.Pointer(o.Native()))
	arg1 = (C.GtkOrientation)(orientation)

	C.gtk_orientable_set_orientation(arg0, orientation)
}
