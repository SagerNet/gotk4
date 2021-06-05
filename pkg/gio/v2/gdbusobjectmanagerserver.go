// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
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
	gextras.Objector
	DBusObjectManager

	// Export exports @object on @manager.
	//
	// If there is already a BusObject exported at the object path, then the old
	// object is removed.
	//
	// The object path for @object must be in the hierarchy rooted by the object
	// path for @manager.
	//
	// Note that @manager will take a reference on @object for as long as it is
	// exported.
	Export(object DBusObjectSkeleton)
	// ExportUniquely: like g_dbus_object_manager_server_export() but appends a
	// string of the form _N (with N being a natural number) to @object's object
	// path if an object with the given path already exists. As such, the
	// BusObjectProxy:g-object-path property of @object may be modified.
	ExportUniquely(object DBusObjectSkeleton)
	// Connection gets the BusConnection used by @manager.
	Connection() DBusConnection
	// IsExported returns whether @object is currently exported on @manager.
	IsExported(object DBusObjectSkeleton) bool
	// SetConnection exports all objects managed by @manager on @connection. If
	// @connection is nil, stops exporting objects.
	SetConnection(connection DBusConnection)
	// Unexport: if @manager has an object at @path, removes the object.
	// Otherwise does nothing.
	//
	// Note that @object_path must be in the hierarchy rooted by the object path
	// for @manager.
	Unexport(objectPath string) bool
}

// dBusObjectManagerServer implements the DBusObjectManagerServer interface.
type dBusObjectManagerServer struct {
	gextras.Objector
	DBusObjectManager
}

var _ DBusObjectManagerServer = (*dBusObjectManagerServer)(nil)

// WrapDBusObjectManagerServer wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectManagerServer(obj *externglib.Object) DBusObjectManagerServer {
	return DBusObjectManagerServer{
		Objector:          obj,
		DBusObjectManager: WrapDBusObjectManager(obj),
	}
}

func marshalDBusObjectManagerServer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectManagerServer(obj), nil
}

// NewDBusObjectManagerServer constructs a class DBusObjectManagerServer.
func NewDBusObjectManagerServer(objectPath string) DBusObjectManagerServer {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GDBusObjectManagerServer
	var goret1 DBusObjectManagerServer

	cret = C.g_dbus_object_manager_server_new(objectPath)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObjectManagerServer)

	return goret1
}

// Export exports @object on @manager.
//
// If there is already a BusObject exported at the object path, then the old
// object is removed.
//
// The object path for @object must be in the hierarchy rooted by the object
// path for @manager.
//
// Note that @manager will take a reference on @object for as long as it is
// exported.
func (m dBusObjectManagerServer) Export(object DBusObjectSkeleton) {
	var arg0 *C.GDBusObjectManagerServer
	var arg1 *C.GDBusObjectSkeleton

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export(arg0, object)
}

// ExportUniquely: like g_dbus_object_manager_server_export() but appends a
// string of the form _N (with N being a natural number) to @object's object
// path if an object with the given path already exists. As such, the
// BusObjectProxy:g-object-path property of @object may be modified.
func (m dBusObjectManagerServer) ExportUniquely(object DBusObjectSkeleton) {
	var arg0 *C.GDBusObjectManagerServer
	var arg1 *C.GDBusObjectSkeleton

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	C.g_dbus_object_manager_server_export_uniquely(arg0, object)
}

// Connection gets the BusConnection used by @manager.
func (m dBusObjectManagerServer) Connection() DBusConnection {
	var arg0 *C.GDBusObjectManagerServer

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))

	var cret *C.GDBusConnection
	var goret1 DBusConnection

	cret = C.g_dbus_object_manager_server_get_connection(arg0)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusConnection)

	return goret1
}

// IsExported returns whether @object is currently exported on @manager.
func (m dBusObjectManagerServer) IsExported(object DBusObjectSkeleton) bool {
	var arg0 *C.GDBusObjectManagerServer
	var arg1 *C.GDBusObjectSkeleton

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GDBusObjectSkeleton)(unsafe.Pointer(object.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_dbus_object_manager_server_is_exported(arg0, object)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetConnection exports all objects managed by @manager on @connection. If
// @connection is nil, stops exporting objects.
func (m dBusObjectManagerServer) SetConnection(connection DBusConnection) {
	var arg0 *C.GDBusObjectManagerServer
	var arg1 *C.GDBusConnection

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))

	C.g_dbus_object_manager_server_set_connection(arg0, connection)
}

// Unexport: if @manager has an object at @path, removes the object.
// Otherwise does nothing.
//
// Note that @object_path must be in the hierarchy rooted by the object path
// for @manager.
func (m dBusObjectManagerServer) Unexport(objectPath string) bool {
	var arg0 *C.GDBusObjectManagerServer
	var arg1 *C.gchar

	arg0 = (*C.GDBusObjectManagerServer)(unsafe.Pointer(m.Native()))
	arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_dbus_object_manager_server_unexport(arg0, objectPath)

	goret1 = C.bool(cret) != C.false

	return goret1
}

type DBusObjectManagerServerPrivate struct {
	native C.GDBusObjectManagerServerPrivate
}

// WrapDBusObjectManagerServerPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusObjectManagerServerPrivate(ptr unsafe.Pointer) *DBusObjectManagerServerPrivate {
	if ptr == nil {
		return nil
	}

	return (*DBusObjectManagerServerPrivate)(ptr)
}

func marshalDBusObjectManagerServerPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusObjectManagerServerPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusObjectManagerServerPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
