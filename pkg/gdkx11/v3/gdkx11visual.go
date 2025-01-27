// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_visual_get_type()), F: marshalX11Visualer},
	})
}

type X11Visual struct {
	gdk.Visual
}

func WrapX11Visual(obj *externglib.Object) *X11Visual {
	return &X11Visual{
		Visual: gdk.Visual{
			Object: obj,
		},
	}
}

func marshalX11Visualer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Visual(obj), nil
}

func (*X11Visual) privateX11Visual() {}
