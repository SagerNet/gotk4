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
		{T: externglib.Type(C.g_dbus_object_manager_client_get_type()), F: marshalDBusObjectManagerClient},
	})
}

// DBusObjectManagerClient is used to create, monitor and delete object proxies
// for remote objects exported by a BusObjectManagerServer (or any code
// implementing the org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface).
//
// Once an instance of this type has been created, you can connect to the
// BusObjectManager::object-added and BusObjectManager::object-removed signals
// and inspect the BusObjectProxy objects returned by
// g_dbus_object_manager_get_objects().
//
// If the name for a BusObjectManagerClient is not owned by anyone at object
// construction time, the default behavior is to request the message bus to
// launch an owner for the name. This behavior can be disabled using the
// G_DBUS_OBJECT_MANAGER_CLIENT_FLAGS_DO_NOT_AUTO_START flag. It's also worth
// noting that this only works if the name of interest is activatable in the
// first place. E.g. in some cases it is not possible to launch an owner for the
// requested name. In this case, BusObjectManagerClient object construction
// still succeeds but there will be no object proxies (e.g.
// g_dbus_object_manager_get_objects() returns the empty list) and the
// BusObjectManagerClient:name-owner property is nil.
//
// The owner of the requested name can come and go (for example consider a
// system service being restarted) – BusObjectManagerClient handles this case
// too; simply connect to the #GObject::notify signal to watch for changes on
// the BusObjectManagerClient:name-owner property. When the name owner vanishes,
// the behavior is that BusObjectManagerClient:name-owner is set to nil (this
// includes emission of the #GObject::notify signal) and then
// BusObjectManager::object-removed signals are synthesized for all currently
// existing object proxies. Since BusObjectManagerClient:name-owner is nil when
// this happens, you can use this information to disambiguate a synthesized
// signal from a genuine signal caused by object removal on the remote
// BusObjectManager. Similarly, when a new name owner appears,
// BusObjectManager::object-added signals are synthesized while
// BusObjectManagerClient:name-owner is still nil. Only when all object proxies
// have been added, the BusObjectManagerClient:name-owner is set to the new name
// owner (this includes emission of the #GObject::notify signal). Furthermore,
// you are guaranteed that BusObjectManagerClient:name-owner will alternate
// between a name owner (e.g. `:1.42`) and nil even in the case where the name
// of interest is atomically replaced
//
// Ultimately, BusObjectManagerClient is used to obtain BusProxy instances. All
// signals (including the org.freedesktop.DBus.Properties::PropertiesChanged
// signal) delivered to BusProxy instances are guaranteed to originate from the
// name owner. This guarantee along with the behavior described above, means
// that certain race conditions including the "half the proxy is from the old
// owner and the other half is from the new owner" problem cannot happen.
//
// To avoid having the application connect to signals on the returned
// BusObjectProxy and BusProxy objects, the BusObject::interface-added,
// BusObject::interface-removed, BusProxy::g-properties-changed and
// BusProxy::g-signal signals are also emitted on the BusObjectManagerClient
// instance managing these objects. The signals emitted are
// BusObjectManager::interface-added, BusObjectManager::interface-removed,
// BusObjectManagerClient::interface-proxy-properties-changed and
// BusObjectManagerClient::interface-proxy-signal.
//
// Note that all callbacks and signals are emitted in the [thread-default main
// context][g-main-context-push-thread-default] that the BusObjectManagerClient
// object was constructed in. Additionally, the BusObjectProxy and BusProxy
// objects originating from the BusObjectManagerClient object will be created in
// the same context and, consequently, will deliver signals in the same main
// loop.
type DBusObjectManagerClient interface {
	gextras.Objector
	AsyncInitable
	DBusObjectManager
	Initable

	// Connection gets the BusConnection used by @manager.
	Connection() DBusConnection
	// Flags gets the flags that @manager was constructed with.
	Flags() DBusObjectManagerClientFlags
	// Name gets the name that @manager is for, or nil if not a message bus
	// connection.
	Name() string
	// NameOwner: the unique name that owns the name that @manager is for or nil
	// if no-one currently owns that name. You can connect to the
	// #GObject::notify signal to track changes to the
	// BusObjectManagerClient:name-owner property.
	NameOwner() string
}

// dBusObjectManagerClient implements the DBusObjectManagerClient interface.
type dBusObjectManagerClient struct {
	gextras.Objector
	AsyncInitable
	DBusObjectManager
	Initable
}

var _ DBusObjectManagerClient = (*dBusObjectManagerClient)(nil)

// WrapDBusObjectManagerClient wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusObjectManagerClient(obj *externglib.Object) DBusObjectManagerClient {
	return DBusObjectManagerClient{
		Objector:          obj,
		AsyncInitable:     WrapAsyncInitable(obj),
		DBusObjectManager: WrapDBusObjectManager(obj),
		Initable:          WrapInitable(obj),
	}
}

func marshalDBusObjectManagerClient(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusObjectManagerClient(obj), nil
}

// NewDBusObjectManagerClientFinish constructs a class DBusObjectManagerClient.
func NewDBusObjectManagerClientFinish(res AsyncResult) (dBusObjectManagerClient DBusObjectManagerClient, err error) {
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var cret C.GDBusObjectManagerClient
	var goret1 DBusObjectManagerClient
	var goerr error

	cret = C.g_dbus_object_manager_client_new_finish(res, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObjectManagerClient)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// NewDBusObjectManagerClientForBusFinish constructs a class DBusObjectManagerClient.
func NewDBusObjectManagerClientForBusFinish(res AsyncResult) (dBusObjectManagerClient DBusObjectManagerClient, err error) {
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var cret C.GDBusObjectManagerClient
	var goret1 DBusObjectManagerClient
	var goerr error

	cret = C.g_dbus_object_manager_client_new_for_bus_finish(res, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusObjectManagerClient)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Connection gets the BusConnection used by @manager.
func (m dBusObjectManagerClient) Connection() DBusConnection {
	var arg0 *C.GDBusObjectManagerClient

	arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(m.Native()))

	var cret *C.GDBusConnection
	var goret1 DBusConnection

	cret = C.g_dbus_object_manager_client_get_connection(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(DBusConnection)

	return goret1
}

// Flags gets the flags that @manager was constructed with.
func (m dBusObjectManagerClient) Flags() DBusObjectManagerClientFlags {
	var arg0 *C.GDBusObjectManagerClient

	arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(m.Native()))

	var cret C.GDBusObjectManagerClientFlags
	var goret1 DBusObjectManagerClientFlags

	cret = C.g_dbus_object_manager_client_get_flags(arg0)

	goret1 = DBusObjectManagerClientFlags(cret)

	return goret1
}

// Name gets the name that @manager is for, or nil if not a message bus
// connection.
func (m dBusObjectManagerClient) Name() string {
	var arg0 *C.GDBusObjectManagerClient

	arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(m.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dbus_object_manager_client_get_name(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// NameOwner: the unique name that owns the name that @manager is for or nil
// if no-one currently owns that name. You can connect to the
// #GObject::notify signal to track changes to the
// BusObjectManagerClient:name-owner property.
func (m dBusObjectManagerClient) NameOwner() string {
	var arg0 *C.GDBusObjectManagerClient

	arg0 = (*C.GDBusObjectManagerClient)(unsafe.Pointer(m.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_dbus_object_manager_client_get_name_owner(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

type DBusObjectManagerClientPrivate struct {
	native C.GDBusObjectManagerClientPrivate
}

// WrapDBusObjectManagerClientPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusObjectManagerClientPrivate(ptr unsafe.Pointer) *DBusObjectManagerClientPrivate {
	if ptr == nil {
		return nil
	}

	return (*DBusObjectManagerClientPrivate)(ptr)
}

func marshalDBusObjectManagerClientPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusObjectManagerClientPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusObjectManagerClientPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
