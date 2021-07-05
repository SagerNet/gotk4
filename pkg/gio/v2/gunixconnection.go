// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
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

	// ReceiveCredentials receives credentials from the sending end of the
	// connection. The sending end has to call
	// g_unix_connection_send_credentials() (or similar) for this to work.
	//
	// As well as reading the credentials this also reads (and discards) a
	// single byte from the stream, as this is required for credentials passing
	// to work on some implementations.
	//
	// This method can be expected to be available on the following platforms:
	//
	// - Linux since GLib 2.26 - FreeBSD since GLib 2.26 - GNU/kFreeBSD since
	// GLib 2.36 - Solaris, Illumos and OpenSolaris since GLib 2.40 - GNU/Hurd
	// since GLib 2.40
	//
	// Other ways to exchange credentials with a foreign peer includes the
	// CredentialsMessage type and g_socket_get_credentials() function.
	ReceiveCredentials(cancellable Cancellable) (Credentials, error)
	// ReceiveCredentialsAsync: asynchronously receive credentials.
	//
	// For more details, see g_unix_connection_receive_credentials() which is
	// the synchronous version of this call.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_unix_connection_receive_credentials_finish() to get the result of
	// the operation.
	ReceiveCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// ReceiveCredentialsFinish finishes an asynchronous receive credentials
	// operation started with g_unix_connection_receive_credentials_async().
	ReceiveCredentialsFinish(result AsyncResult) (Credentials, error)
	// ReceiveFd receives a file descriptor from the sending end of the
	// connection. The sending end has to call g_unix_connection_send_fd() for
	// this to work.
	//
	// As well as reading the fd this also reads a single byte from the stream,
	// as this is required for fd passing to work on some implementations.
	ReceiveFd(cancellable Cancellable) (int, error)
	// SendCredentials passes the credentials of the current user the receiving
	// side of the connection. The receiving end has to call
	// g_unix_connection_receive_credentials() (or similar) to accept the
	// credentials.
	//
	// As well as sending the credentials this also writes a single NUL byte to
	// the stream, as this is required for credentials passing to work on some
	// implementations.
	//
	// This method can be expected to be available on the following platforms:
	//
	// - Linux since GLib 2.26 - FreeBSD since GLib 2.26 - GNU/kFreeBSD since
	// GLib 2.36 - Solaris, Illumos and OpenSolaris since GLib 2.40 - GNU/Hurd
	// since GLib 2.40
	//
	// Other ways to exchange credentials with a foreign peer includes the
	// CredentialsMessage type and g_socket_get_credentials() function.
	SendCredentials(cancellable Cancellable) error
	// SendCredentialsAsync: asynchronously send credentials.
	//
	// For more details, see g_unix_connection_send_credentials() which is the
	// synchronous version of this call.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_unix_connection_send_credentials_finish() to get the result of the
	// operation.
	SendCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// SendCredentialsFinish finishes an asynchronous send credentials operation
	// started with g_unix_connection_send_credentials_async().
	SendCredentialsFinish(result AsyncResult) error
	// SendFd passes a file descriptor to the receiving side of the connection.
	// The receiving end has to call g_unix_connection_receive_fd() to accept
	// the file descriptor.
	//
	// As well as sending the fd this also writes a single byte to the stream,
	// as this is required for fd passing to work on some implementations.
	SendFd(fd int, cancellable Cancellable) error
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

func (c unixConnection) ReceiveCredentials(cancellable Cancellable) (Credentials, error) {
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

func (c unixConnection) ReceiveCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GUnixConnection    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_unix_connection_receive_credentials_async(_arg0, _arg1, _arg2, _arg3)
}

func (c unixConnection) ReceiveCredentialsFinish(result AsyncResult) (Credentials, error) {
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

func (c unixConnection) ReceiveFd(cancellable Cancellable) (int, error) {
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

func (c unixConnection) SendCredentials(cancellable Cancellable) error {
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

func (c unixConnection) SendCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GUnixConnection    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GUnixConnection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_unix_connection_send_credentials_async(_arg0, _arg1, _arg2, _arg3)
}

func (c unixConnection) SendCredentialsFinish(result AsyncResult) error {
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

func (c unixConnection) SendFd(fd int, cancellable Cancellable) error {
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
