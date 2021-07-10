// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevicer},
	})
}

// WaylandDevicer describes WaylandDevice's methods.
type WaylandDevicer interface {
	gextras.Objector

	NodePath() string
}

// WaylandDevice: the Wayland implementation of `GdkDevice`.
//
// Beyond the regular [class@Gdk.Device] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_seat` with
// [method@GdkWayland.WaylandDevice.get_wl_seat], the `wl_keyboard` with
// [method@GdkWayland.WaylandDevice.get_wl_keyboard] and the `wl_pointer` with
// [method@GdkWayland.WaylandDevice.get_wl_pointer].
type WaylandDevice struct {
	gdk.Device
}

var _ WaylandDevicer = (*WaylandDevice)(nil)

func wrapWaylandDevicer(obj *externglib.Object) WaylandDevicer {
	return &WaylandDevice{
		Device: gdk.Device{
			Object: obj,
		},
	}
}

func marshalWaylandDevicer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWaylandDevicer(obj), nil
}

// NodePath returns the `/dev/input/event*` path of this device.
//
// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg. mouse,
// keyboard, touch,...), this function will return nil.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
func (device *WaylandDevice) NodePath() string {
	var _arg0 *C.GdkDevice // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_wayland_device_get_node_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
