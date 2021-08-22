// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevicer},
	})
}

// WaylandDevice: wayland implementation of GdkDevice.
//
// Beyond the regular gdk.Device API, the Wayland implementation provides access
// to Wayland objects such as the wl_seat with
// gdkwayland.WaylandDevice.GetWlSeat(), the wl_keyboard with
// gdkwayland.WaylandDevice.GetWlKeyboard() and the wl_pointer with
// gdkwayland.WaylandDevice.GetWlPointer().
type WaylandDevice struct {
	gdk.Device
}

func WrapWaylandDevice(obj *externglib.Object) *WaylandDevice {
	return &WaylandDevice{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalWaylandDevicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDevice(obj), nil
}

// NodePath returns the /dev/input/event* path of this device.
//
// For GdkDevices that possibly coalesce multiple hardware devices (eg. mouse,
// keyboard, touch,...), this function will return NULL.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
func (device *WaylandDevice) NodePath() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_wayland_device_get_node_path(_arg0)
	runtime.KeepAlive(device)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
