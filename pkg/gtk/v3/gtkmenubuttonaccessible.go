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
		{T: externglib.Type(C.gtk_menu_button_accessible_get_type()), F: marshalMenuButtonAccessible},
	})
}

type MenuButtonAccessible interface {
	ToggleButtonAccessible
}

// menuButtonAccessible implements the MenuButtonAccessible interface.
type menuButtonAccessible struct {
	ToggleButtonAccessible
}

var _ MenuButtonAccessible = (*menuButtonAccessible)(nil)

// WrapMenuButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuButtonAccessible(obj *externglib.Object) MenuButtonAccessible {
	return MenuButtonAccessible{
		ToggleButtonAccessible: WrapToggleButtonAccessible(obj),
	}
}

func marshalMenuButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuButtonAccessible(obj), nil
}

type MenuButtonAccessiblePrivate struct {
	native C.GtkMenuButtonAccessiblePrivate
}

// WrapMenuButtonAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMenuButtonAccessiblePrivate(ptr unsafe.Pointer) *MenuButtonAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*MenuButtonAccessiblePrivate)(ptr)
}

func marshalMenuButtonAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMenuButtonAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MenuButtonAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}
