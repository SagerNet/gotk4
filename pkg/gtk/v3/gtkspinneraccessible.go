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
		{T: externglib.Type(C.gtk_spinner_accessible_get_type()), F: marshalSpinnerAccessible},
	})
}

type SpinnerAccessible interface {
	WidgetAccessible
}

// spinnerAccessible implements the SpinnerAccessible interface.
type spinnerAccessible struct {
	WidgetAccessible
}

var _ SpinnerAccessible = (*spinnerAccessible)(nil)

// WrapSpinnerAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapSpinnerAccessible(obj *externglib.Object) SpinnerAccessible {
	return SpinnerAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalSpinnerAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSpinnerAccessible(obj), nil
}

type SpinnerAccessiblePrivate struct {
	native C.GtkSpinnerAccessiblePrivate
}

// WrapSpinnerAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSpinnerAccessiblePrivate(ptr unsafe.Pointer) *SpinnerAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*SpinnerAccessiblePrivate)(ptr)
}

func marshalSpinnerAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSpinnerAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SpinnerAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}