// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_cursor_get_type()), F: marshalX11Cursorrer},
	})
}

// X11Cursorrer describes X11Cursor's methods.
type X11Cursorrer interface {
	gextras.Objector

	privateX11Cursor()
}

type X11Cursor struct {
	gdk.Cursor
}

var _ X11Cursorrer = (*X11Cursor)(nil)

func wrapX11Cursorrer(obj *externglib.Object) X11Cursorrer {
	return &X11Cursor{
		Cursor: gdk.Cursor{
			Object: obj,
		},
	}
}

func marshalX11Cursorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Cursorrer(obj), nil
}

func (*X11Cursor) privateX11Cursor() {}
