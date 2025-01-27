// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
import "C"

// DBusActionGroupGet obtains a BusActionGroup for the action group which is
// exported at the given bus_name and object_path.
//
// The thread default main context is taken at the time of this call. All
// signals on the menu model (and any linked models) are reported with respect
// to this context. All calls on the returned menu model (and linked models)
// must also originate from this same context, with the thread default main
// context unchanged.
//
// This call is non-blocking. The returned action group may or may not already
// be filled in. The correct thing to do is connect the signals for the action
// group to monitor for changes and then to call g_action_group_list_actions()
// to get the initial list.
func DBusActionGroupGet(connection *DBusConnection, busName string, objectPath string) *DBusActionGroup {
	var _arg1 *C.GDBusConnection  // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _cret *C.GDBusActionGroup // in

	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	if busName != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(busName)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(objectPath)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_dbus_action_group_get(_arg1, _arg2, _arg3)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(busName)
	runtime.KeepAlive(objectPath)

	var _dBusActionGroup *DBusActionGroup // out

	_dBusActionGroup = WrapDBusActionGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dBusActionGroup
}
