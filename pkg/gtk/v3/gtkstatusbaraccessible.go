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
		{T: externglib.Type(C.gtk_statusbar_accessible_get_type()), F: marshalStatusbarAccessible},
	})
}

type StatusbarAccessible interface {
	ContainerAccessible
}

// statusbarAccessible implements the StatusbarAccessible interface.
type statusbarAccessible struct {
	ContainerAccessible
}

var _ StatusbarAccessible = (*statusbarAccessible)(nil)

// WrapStatusbarAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapStatusbarAccessible(obj *externglib.Object) StatusbarAccessible {
	return StatusbarAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalStatusbarAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusbarAccessible(obj), nil
}

type StatusbarAccessiblePrivate struct {
	native C.GtkStatusbarAccessiblePrivate
}

// WrapStatusbarAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStatusbarAccessiblePrivate(ptr unsafe.Pointer) *StatusbarAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*StatusbarAccessiblePrivate)(ptr)
}

func marshalStatusbarAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStatusbarAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StatusbarAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
