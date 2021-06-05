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
		{T: externglib.Type(C.gtk_scrolled_window_accessible_get_type()), F: marshalScrolledWindowAccessible},
	})
}

type ScrolledWindowAccessible interface {
	ContainerAccessible
}

// scrolledWindowAccessible implements the ScrolledWindowAccessible interface.
type scrolledWindowAccessible struct {
	ContainerAccessible
}

var _ ScrolledWindowAccessible = (*scrolledWindowAccessible)(nil)

// WrapScrolledWindowAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrolledWindowAccessible(obj *externglib.Object) ScrolledWindowAccessible {
	return ScrolledWindowAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalScrolledWindowAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrolledWindowAccessible(obj), nil
}

type ScrolledWindowAccessiblePrivate struct {
	native C.GtkScrolledWindowAccessiblePrivate
}

// WrapScrolledWindowAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScrolledWindowAccessiblePrivate(ptr unsafe.Pointer) *ScrolledWindowAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*ScrolledWindowAccessiblePrivate)(ptr)
}

func marshalScrolledWindowAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScrolledWindowAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *ScrolledWindowAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
