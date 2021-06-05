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
		{T: externglib.Type(C.gtk_plug_accessible_get_type()), F: marshalPlugAccessible},
	})
}

type PlugAccessible interface {
	WindowAccessible

	ID() string
}

// plugAccessible implements the PlugAccessible interface.
type plugAccessible struct {
	WindowAccessible
}

var _ PlugAccessible = (*plugAccessible)(nil)

// WrapPlugAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapPlugAccessible(obj *externglib.Object) PlugAccessible {
	return PlugAccessible{
		WindowAccessible: WrapWindowAccessible(obj),
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlugAccessible(obj), nil
}

func (p plugAccessible) ID() string {
	var arg0 *C.GtkPlugAccessible

	arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(p.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.gtk_plug_accessible_get_id(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

type PlugAccessiblePrivate struct {
	native C.GtkPlugAccessiblePrivate
}

// WrapPlugAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPlugAccessiblePrivate(ptr unsafe.Pointer) *PlugAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*PlugAccessiblePrivate)(ptr)
}

func marshalPlugAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPlugAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PlugAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
