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
		{T: externglib.Type(C.g_unix_connection_get_type()), F: marshalUnixConnection},
	})
}

// UnixConnection: this is the subclass of Connection that is created for UNIX
// domain sockets.
//
// It contains functions to do some of the UNIX socket specific functionality
// like passing file descriptors.
//
// Note that `<gio/gunixconnection.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixConnection interface {
	SocketConnection

	// ReceiveCredentialsUnixConnection:
	ReceiveCredentialsUnixConnection(cancellable Cancellable) (Credentials, error)
	// ReceiveCredentialsFinishUnixConnection:
	ReceiveCredentialsFinishUnixConnection(result AsyncResult) (Credentials, error)
	// ReceiveFdUnixConnection:
	ReceiveFdUnixConnection(cancellable Cancellable) (int, error)
	// SendCredentialsUnixConnection:
	SendCredentialsUnixConnection(cancellable Cancellable) error
	// SendCredentialsFinishUnixConnection:
	SendCredentialsFinishUnixConnection(result AsyncResult) error
	// SendFdUnixConnection:
	SendFdUnixConnection(fd int, cancellable Cancellable) error
}

// unixConnection implements the UnixConnection class.
type unixConnection struct {
	SocketConnection
}

// WrapUnixConnection wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixConnection(obj *externglib.Object) UnixConnection {
	return unixConnection{
		SocketConnection: WrapSocketConnection(obj),
	}
}

func marshalUnixConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixConnection(obj), nil
}

func (c unixConnection) ReceiveCredentialsUnixConnection(cancellable Cancellable) (Credentials, error) {
	var _arg0 *C.GUnixConnection // out
	var _arg1 *C.GCancellable    // out
	var _cret *C.GCredentials    // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_unix_connection_receive_credentials(_arg0, _arg1, &_cerr)

	var _credentials Credentials // out
	var _goerr error             // out

	_credentials = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Credentials)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _credentials, _goerr
}

func (c unixConnection) ReceiveCredentialsFinishUnixConnection(result AsyncResult) (Credentials, error) {
	var _arg0 *C.GUnixConnection // out
	var _arg1 *C.GAsyncResult    // out
	var _cret *C.GCredentials    // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_unix_connection_receive_credentials_finish(_arg0, _arg1, &_cerr)

	var _credentials Credentials // out
	var _goerr error             // out

	_credentials = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Credentials)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _credentials, _goerr
}

func (c unixConnection) ReceiveFdUnixConnection(cancellable Cancellable) (int, error) {
	var _arg0 *C.GUnixConnection // out
	var _arg1 *C.GCancellable    // out
	var _cret C.gint             // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_unix_connection_receive_fd(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

func (c unixConnection) SendCredentialsUnixConnection(cancellable Cancellable) error {
	var _arg0 *C.GUnixConnection // out
	var _arg1 *C.GCancellable    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_unix_connection_send_credentials(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (c unixConnection) SendCredentialsFinishUnixConnection(result AsyncResult) error {
	var _arg0 *C.GUnixConnection // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_unix_connection_send_credentials_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (c unixConnection) SendFdUnixConnection(fd int, cancellable Cancellable) error {
	var _arg0 *C.GUnixConnection // out
	var _arg1 C.gint             // out
	var _arg2 *C.GCancellable    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(fd)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_unix_connection_send_fd(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
