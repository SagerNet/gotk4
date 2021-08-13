// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager *X11DeviceManagerCore, deviceId int) *X11DeviceCore {
	var _arg1 *C.GdkDeviceManager // out
	var _arg2 C.gint              // out
	var _cret *C.GdkDevice        // in

	_arg1 = (*C.GdkDeviceManager)(unsafe.Pointer(deviceManager.Native()))
	_arg2 = C.gint(deviceId)

	_cret = C.gdk_x11_device_manager_lookup(_arg1, _arg2)
	runtime.KeepAlive(deviceManager)
	runtime.KeepAlive(deviceId)

	var _x11DeviceCore *X11DeviceCore // out

	if _cret != nil {
		_x11DeviceCore = wrapX11DeviceCore(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _x11DeviceCore
}
