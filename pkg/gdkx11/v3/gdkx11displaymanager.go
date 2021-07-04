// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

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
		{T: externglib.Type(C.gdk_x11_display_manager_get_type()), F: marshalX11DisplayManager},
	})
}

// X11DisplayManager:
type X11DisplayManager interface {
	gdk.DisplayManager
}

// x11DisplayManager implements the X11DisplayManager class.
type x11DisplayManager struct {
	gdk.DisplayManager
}

// WrapX11DisplayManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DisplayManager(obj *externglib.Object) X11DisplayManager {
	return x11DisplayManager{
		DisplayManager: gdk.WrapDisplayManager(obj),
	}
}

func marshalX11DisplayManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DisplayManager(obj), nil
}
