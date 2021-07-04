// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
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
		{T: externglib.Type(C.g_socket_address_enumerator_get_type()), F: marshalSocketAddressEnumerator},
	})
}

// SocketAddressEnumerator is an enumerator type for Address instances. It is
// returned by enumeration functions such as g_socket_connectable_enumerate(),
// which returns a AddressEnumerator to list each Address which could be used to
// connect to that Connectable.
//
// Enumeration is typically a blocking operation, so the asynchronous methods
// g_socket_address_enumerator_next_async() and
// g_socket_address_enumerator_next_finish() should be used where possible.
//
// Each AddressEnumerator can only be enumerated once. Once
// g_socket_address_enumerator_next() has returned nil, further enumeration with
// that AddressEnumerator is not possible, and it can be unreffed.
type SocketAddressEnumerator interface {
	gextras.Objector

	// NextSocketAddressEnumerator:
	NextSocketAddressEnumerator(cancellable Cancellable) (SocketAddress, error)
	// NextFinishSocketAddressEnumerator:
	NextFinishSocketAddressEnumerator(result AsyncResult) (SocketAddress, error)
}

// socketAddressEnumerator implements the SocketAddressEnumerator class.
type socketAddressEnumerator struct {
	gextras.Objector
}

// WrapSocketAddressEnumerator wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketAddressEnumerator(obj *externglib.Object) SocketAddressEnumerator {
	return socketAddressEnumerator{
		Objector: obj,
	}
}

func marshalSocketAddressEnumerator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAddressEnumerator(obj), nil
}

func (e socketAddressEnumerator) NextSocketAddressEnumerator(cancellable Cancellable) (SocketAddress, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GCancellable             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_socket_address_enumerator_next(_arg0, _arg1, &_cerr)

	var _socketAddress SocketAddress // out
	var _goerr error                 // out

	_socketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(SocketAddress)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}

func (e socketAddressEnumerator) NextFinishSocketAddressEnumerator(result AsyncResult) (SocketAddress, error) {
	var _arg0 *C.GSocketAddressEnumerator // out
	var _arg1 *C.GAsyncResult             // out
	var _cret *C.GSocketAddress           // in
	var _cerr *C.GError                   // in

	_arg0 = (*C.GSocketAddressEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_socket_address_enumerator_next_finish(_arg0, _arg1, &_cerr)

	var _socketAddress SocketAddress // out
	var _goerr error                 // out

	_socketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(SocketAddress)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _socketAddress, _goerr
}
