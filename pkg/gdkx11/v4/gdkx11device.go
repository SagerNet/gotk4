// Code generated by girgen. DO NOT EDIT.

package gdkx11

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceGetID returns the device ID as seen by XInput2.
func X11DeviceGetID(device X11DeviceXI2) int {
	var arg1 *C.GdkDevice

	arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	var cret C.int
	var goret1 int

	cret = C.gdk_x11_device_get_id(device)

	goret1 = C.int(cret)

	return goret1
}
