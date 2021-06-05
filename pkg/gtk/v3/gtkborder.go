// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_border_get_type()), F: marshalBorder},
	})
}

// Border: a struct that specifies a border around a rectangular area that can
// be of different width on each side.
type Border struct {
	native C.GtkBorder
}

// WrapBorder wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBorder(ptr unsafe.Pointer) *Border {
	if ptr == nil {
		return nil
	}

	return (*Border)(ptr)
}

func marshalBorder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBorder(unsafe.Pointer(b)), nil
}

// NewBorder constructs a struct Border.
func NewBorder() *Border {
	var cret *C.GtkBorder
	var goret1 *Border

	cret = C.gtk_border_new()

	goret1 = WrapBorder(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Border) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Native returns the underlying C source pointer.
func (b *Border) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Left gets the field inside the struct.
func (b *Border) Left() int16 {
	v = C.gint16(b.native.left)
}

// Right gets the field inside the struct.
func (b *Border) Right() int16 {
	v = C.gint16(b.native.right)
}

// Top gets the field inside the struct.
func (b *Border) Top() int16 {
	v = C.gint16(b.native.top)
}

// Bottom gets the field inside the struct.
func (b *Border) Bottom() int16 {
	v = C.gint16(b.native.bottom)
}

// Copy copies a Border-struct.
func (b *Border) Copy() *Border {
	var arg0 *C.GtkBorder

	arg0 = (*C.GtkBorder)(unsafe.Pointer(b.Native()))

	var cret *C.GtkBorder
	var goret1 *Border

	cret = C.gtk_border_copy(arg0)

	goret1 = WrapBorder(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Border) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Free frees a Border-struct.
func (b *Border) Free() {
	var arg0 *C.GtkBorder

	arg0 = (*C.GtkBorder)(unsafe.Pointer(b.Native()))

	C.gtk_border_free(arg0)
}
