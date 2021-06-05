// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_button_accessible_get_type()), F: marshalButtonAccessible},
	})
}

type ButtonAccessible interface {
	ContainerAccessible
}

// buttonAccessible implements the ButtonAccessible interface.
type buttonAccessible struct {
	ContainerAccessible
}

var _ ButtonAccessible = (*buttonAccessible)(nil)

// WrapButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapButtonAccessible(obj *externglib.Object) ButtonAccessible {
	return ButtonAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButtonAccessible(obj), nil
}

type ButtonAccessiblePrivate struct {
	native C.GtkButtonAccessiblePrivate
}

// WrapButtonAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapButtonAccessiblePrivate(ptr unsafe.Pointer) *ButtonAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*ButtonAccessiblePrivate)(ptr)
}

func marshalButtonAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapButtonAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *ButtonAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}
