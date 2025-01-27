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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_backend_get_type()), F: marshalTLSBackender},
	})
}

// TLS_BACKEND_EXTENSION_POINT_NAME: extension point for TLS functionality via
// Backend. See [Extending GIO][extending-gio].
const TLS_BACKEND_EXTENSION_POINT_NAME = "gio-tls-backend"

// TLSBackendOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSBackendOverrider interface {
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() TLSDatabaser
	// SupportsDTLS checks if DTLS is supported. DTLS support may not be
	// available even if TLS support is available, and vice-versa.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

// TLSBackend: TLS (Transport Layer Security, aka SSL) and DTLS backend.
type TLSBackend struct {
	*externglib.Object
}

// TLSBackender describes TLSBackend's abstract methods.
type TLSBackender interface {
	externglib.Objector

	// CertificateType gets the #GType of backend's Certificate implementation.
	CertificateType() externglib.Type
	// ClientConnectionType gets the #GType of backend's ClientConnection
	// implementation.
	ClientConnectionType() externglib.Type
	// DefaultDatabase gets the default Database used to verify TLS connections.
	DefaultDatabase() TLSDatabaser
	// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
	// implementation.
	DTLSClientConnectionType() externglib.Type
	// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
	// implementation.
	DTLSServerConnectionType() externglib.Type
	// FileDatabaseType gets the #GType of backend's FileDatabase
	// implementation.
	FileDatabaseType() externglib.Type
	// ServerConnectionType gets the #GType of backend's ServerConnection
	// implementation.
	ServerConnectionType() externglib.Type
	// SetDefaultDatabase: set the default Database used to verify TLS
	// connections Any subsequent call to g_tls_backend_get_default_database()
	// will return the database set in this call.
	SetDefaultDatabase(database TLSDatabaser)
	// SupportsDTLS checks if DTLS is supported.
	SupportsDTLS() bool
	// SupportsTLS checks if TLS is supported; if this returns FALSE for the
	// default Backend, it means no "real" TLS backend is available.
	SupportsTLS() bool
}

var _ TLSBackender = (*TLSBackend)(nil)

func WrapTLSBackend(obj *externglib.Object) *TLSBackend {
	return &TLSBackend{
		Object: obj,
	}
}

func marshalTLSBackender(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSBackend(obj), nil
}

// CertificateType gets the #GType of backend's Certificate implementation.
func (backend *TLSBackend) CertificateType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_certificate_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ClientConnectionType gets the #GType of backend's ClientConnection
// implementation.
func (backend *TLSBackend) ClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DefaultDatabase gets the default Database used to verify TLS connections.
func (backend *TLSBackend) DefaultDatabase() TLSDatabaser {
	var _arg0 *C.GTlsBackend  // out
	var _cret *C.GTlsDatabase // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_default_database(_arg0)
	runtime.KeepAlive(backend)

	var _tlsDatabase TLSDatabaser // out

	_tlsDatabase = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(TLSDatabaser)

	return _tlsDatabase
}

// DTLSClientConnectionType gets the #GType of backend’s ClientConnection
// implementation.
func (backend *TLSBackend) DTLSClientConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_dtls_client_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// DTLSServerConnectionType gets the #GType of backend’s ServerConnection
// implementation.
func (backend *TLSBackend) DTLSServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_dtls_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// FileDatabaseType gets the #GType of backend's FileDatabase implementation.
func (backend *TLSBackend) FileDatabaseType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_file_database_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// ServerConnectionType gets the #GType of backend's ServerConnection
// implementation.
func (backend *TLSBackend) ServerConnectionType() externglib.Type {
	var _arg0 *C.GTlsBackend // out
	var _cret C.GType        // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_get_server_connection_type(_arg0)
	runtime.KeepAlive(backend)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetDefaultDatabase: set the default Database used to verify TLS connections
//
// Any subsequent call to g_tls_backend_get_default_database() will return the
// database set in this call. Existing databases and connections are not
// modified.
//
// Setting a NULL default database will reset to using the system default
// database as if g_tls_backend_set_default_database() had never been called.
func (backend *TLSBackend) SetDefaultDatabase(database TLSDatabaser) {
	var _arg0 *C.GTlsBackend  // out
	var _arg1 *C.GTlsDatabase // out

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))
	if database != nil {
		_arg1 = (*C.GTlsDatabase)(unsafe.Pointer(database.Native()))
	}

	C.g_tls_backend_set_default_database(_arg0, _arg1)
	runtime.KeepAlive(backend)
	runtime.KeepAlive(database)
}

// SupportsDTLS checks if DTLS is supported. DTLS support may not be available
// even if TLS support is available, and vice-versa.
func (backend *TLSBackend) SupportsDTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_supports_dtls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsTLS checks if TLS is supported; if this returns FALSE for the default
// Backend, it means no "real" TLS backend is available.
func (backend *TLSBackend) SupportsTLS() bool {
	var _arg0 *C.GTlsBackend // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GTlsBackend)(unsafe.Pointer(backend.Native()))

	_cret = C.g_tls_backend_supports_tls(_arg0)
	runtime.KeepAlive(backend)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TLSBackendGetDefault gets the default Backend for the system.
func TLSBackendGetDefault() TLSBackender {
	var _cret *C.GTlsBackend // in

	_cret = C.g_tls_backend_get_default()

	var _tlsBackend TLSBackender // out

	_tlsBackend = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(TLSBackender)

	return _tlsBackend
}
