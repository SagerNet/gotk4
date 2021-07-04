// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_volume_monitor_get_type()), F: marshalVolumeMonitor},
	})
}

// VolumeMonitor is for listing the user interesting devices and volumes on the
// computer. In other words, what a file selector or file manager would show in
// a sidebar.
//
// Monitor is not [thread-default-context
// aware][g-main-context-push-thread-default], and so should not be used other
// than from the main thread, with no thread-default-context active.
//
// In order to receive updates about volumes and mounts monitored through GVFS,
// a main loop must be running.
type VolumeMonitor interface {
	gextras.Objector

	// MountForUUID:
	MountForUUID(uuid string) Mount
	// VolumeForUUID:
	VolumeForUUID(uuid string) Volume
}

// volumeMonitor implements the VolumeMonitor class.
type volumeMonitor struct {
	gextras.Objector
}

// WrapVolumeMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapVolumeMonitor(obj *externglib.Object) VolumeMonitor {
	return volumeMonitor{
		Objector: obj,
	}
}

func marshalVolumeMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVolumeMonitor(obj), nil
}

func (v volumeMonitor) MountForUUID(uuid string) Mount {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GMount         // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_mount_for_uuid(_arg0, _arg1)

	var _mount Mount // out

	_mount = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Mount)

	return _mount
}

func (v volumeMonitor) VolumeForUUID(uuid string) Volume {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GVolume        // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_volume_for_uuid(_arg0, _arg1)

	var _volume Volume // out

	_volume = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Volume)

	return _volume
}
