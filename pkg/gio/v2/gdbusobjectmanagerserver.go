// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.g_dbus_object_manager_server_get_type()), F: marshalDBusObjectManagerServer},
	})
}

// DBusObjectManagerServer is used to export BusObject instances using the
// standardized org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface. For example, remote D-Bus clients can get all objects and
// properties in a single call. Additionally, any change in the object hierarchy
// is broadcast using signals. This means that D-Bus clients can keep caches up
// to date by only listening to D-Bus signals.
//
// The recommended path to export an object manager at is the path form of the
// well-known name of a D-Bus service, or below. For example, if a D-Bus service
// is available at the well-known name `net.example.ExampleService1`, the object
// manager should typically be exported at `/net/example/ExampleService1`, or
// below (to allow for multiple object managers in a service).
//
// It is supported, but not recommended, to export an object manager at the root
// path, `/`.
//
// See BusObjectManagerClient for the client-side code that is intended to be
// used with BusObjectManagerServer or any D-Bus object implementing the
// org.freedesktop.DBus.ObjectManager interface.
type DBusObjectManagerServer interface {
	DBusObjectManager

	ExportDBusObjectManagerServer(object DBusObjectSkeleton)

	ExportUniquelyDBusObjectManagerServer(object DBusObjectSkeleton)

	Connection() DBusConnection

	IsExportedDBusObjectManagerServer(object DBusObjectSkeleton) bool

	SetConnectionDBusObjectManagerServer(connection DBusConnection)

	UnexportDBusObjectManagerServer(objectPath string) bool
}

// dBusObjectManagerServer implements the DBusObjectManagerServer class.
type dBusObjectManagerServer struct {
	gextras.Objector
}

// WrapDBusObjectManagerServer wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectManagerServer(obj *externglib.Object) DBusObjectManagerServer {
	return dBusObjectManagerServer{
		Objector: obj,
	}
}

func marshalDBusObjectManagerServer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectManagerServer(obj), nil
}

func NewDBusObjectManagerServer(objectPath string) DBusObjectManagerServer {
	var _arg1 *C.gchar                    // out
	var _cret *C.GDBusObjectManagerServer // in

	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_manager_server_new(_arg1)

	var _dBusObjectManagerServer DBusObjectManagerServer // out

	_dBusObjectManagerServer = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DBusObjectManagerServer)

	return _dBusObjectManagerServer
}

func (m dBusObjectManagerServer) ExportDBusObjectManagerServer(object DBusObjectSkeleton) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export(_arg0, _arg1)
}

func (m dBusObjectManagerServer) ExportUniquelyDBusObjectManagerServer(object DBusObjectSkeleton) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export_uniquely(_arg0, _arg1)
}

func (m dBusObjectManagerServer) Connection() DBusConnection {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _cret *C.GDBusConnection          // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))

	_cret = C.g_dbus_object_manager_server_get_connection(_arg0)

	var _dBusConnection DBusConnection // out

	_dBusConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DBusConnection)

	return _dBusConnection
}

func (m dBusObjectManagerServer) IsExportedDBusObjectManagerServer(object DBusObjectSkeleton) bool {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusObjectSkeleton      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	_cret = C.g_dbus_object_manager_server_is_exported(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m dBusObjectManagerServer) SetConnectionDBusObjectManagerServer(connection DBusConnection) {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.GDBusConnection          // out

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	C.g_dbus_object_manager_server_set_connection(_arg0, _arg1)
}

func (m dBusObjectManagerServer) UnexportDBusObjectManagerServer(objectPath string) bool {
	var _arg0 *C.GDBusObjectManagerServer // out
	var _arg1 *C.gchar                    // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_manager_server_unexport(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m dBusObjectManagerServer) Interface(objectPath string, interfaceName string) DBusInterface {
	return WrapDBusObjectManager(gextras.InternObject(m)).Interface(objectPath, interfaceName)
}

func (m dBusObjectManagerServer) Object(objectPath string) DBusObject {
	return WrapDBusObjectManager(gextras.InternObject(m)).Object(objectPath)
}

func (m dBusObjectManagerServer) ObjectPath() string {
	return WrapDBusObjectManager(gextras.InternObject(m)).ObjectPath()
}