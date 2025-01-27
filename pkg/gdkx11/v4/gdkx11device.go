// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceGetID returns the device ID as seen by XInput2.
func X11DeviceGetID(device *X11DeviceXI2) int {
	var _arg1 *C.GdkDevice // out
	var _cret C.int        // in

	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_x11_device_get_id(_arg1)
	runtime.KeepAlive(device)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
