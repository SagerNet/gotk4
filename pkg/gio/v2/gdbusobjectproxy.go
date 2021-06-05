// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_dbus_object_proxy_get_type()), F: marshalDBusObjectProxy},
	})
}

// DBusObjectProxy: a BusObjectProxy is an object used to represent a remote
// object with one or more D-Bus interfaces. Normally, you don't instantiate a
// BusObjectProxy yourself - typically BusObjectManagerClient is used to obtain
// it.
type DBusObjectProxy interface {
	gextras.Objector
	DBusObject

	// Connection gets the connection that @proxy is for.
	Connection() DBusConnection
}

// dBusObjectProxy implements the DBusObjectProxy interface.
type dBusObjectProxy struct {
	gextras.Objector
	DBusObject
}

var _ DBusObjectProxy = (*dBusObjectProxy)(nil)

// WrapDBusObjectProxy wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectProxy(obj *externglib.Object) DBusObjectProxy {
	return DBusObjectProxy{
		Objector:   obj,
		DBusObject: WrapDBusObject(obj),
	}
}

func marshalDBusObjectProxy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectProxy(obj), nil
}

// NewDBusObjectProxy constructs a class DBusObjectProxy.
func NewDBusObjectProxy(connection DBusConnection, objectPath string) DBusObjectProxy {
	var arg1 *C.GDBusConnection
	var arg2 *C.gchar

	arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	arg2 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GDBusObjectProxy
	var ret1 DBusObjectProxy

	cret = C.g_dbus_object_proxy_new(connection, objectPath)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObjectProxy)

	return ret1
}

// Connection gets the connection that @proxy is for.
func (p dBusObjectProxy) Connection() DBusConnection {
	var arg0 *C.GDBusObjectProxy

	arg0 = (*C.GDBusObjectProxy)(unsafe.Pointer(p.Native()))

	var cret *C.GDBusConnection
	var ret1 DBusConnection

	cret = C.g_dbus_object_proxy_get_connection(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(DBusConnection)

	return ret1
}

type DBusObjectProxyPrivate struct {
	native C.GDBusObjectProxyPrivate
}

// WrapDBusObjectProxyPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusObjectProxyPrivate(ptr unsafe.Pointer) *DBusObjectProxyPrivate {
	if ptr == nil {
		return nil
	}

	return (*DBusObjectProxyPrivate)(ptr)
}

func marshalDBusObjectProxyPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusObjectProxyPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusObjectProxyPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}