// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_credentials_get_type()), F: marshalCredentialser},
	})
}

// Credentials type is a reference-counted wrapper for native credentials. This
// information is typically used for identifying, authenticating and authorizing
// other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a struct ucred - see the unix(7) man
// page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a struct xucred. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a struct cmsgcred. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a struct unpcbid. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a struct sockpeercred. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a ucred_t. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials struct {
	*externglib.Object
}

func WrapCredentials(obj *externglib.Object) *Credentials {
	return &Credentials{
		Object: obj,
	}
}

func marshalCredentialser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCredentials(obj), nil
}

// NewCredentials creates a new #GCredentials object with credentials matching
// the the current process.
func NewCredentials() *Credentials {
	var _cret *C.GCredentials // in

	_cret = C.g_credentials_new()

	var _credentials *Credentials // out

	_credentials = WrapCredentials(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _credentials
}

// UnixPid tries to get the UNIX process identifier from credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX process
// ID (for example this is the case for G_CREDENTIALS_TYPE_APPLE_XUCRED).
func (credentials *Credentials) UnixPid() (int, error) {
	var _arg0 *C.GCredentials // out
	var _cret C.pid_t         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))

	_cret = C.g_credentials_get_unix_pid(_arg0, &_cerr)
	runtime.KeepAlive(credentials)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _gint, _goerr
}

// UnixUser tries to get the UNIX user identifier from credentials. This method
// is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX user.
func (credentials *Credentials) UnixUser() (uint, error) {
	var _arg0 *C.GCredentials // out
	var _cret C.uid_t         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))

	_cret = C.g_credentials_get_unix_user(_arg0, &_cerr)
	runtime.KeepAlive(credentials)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// IsSameUser checks if credentials and other_credentials is the same user.
//
// This operation can fail if #GCredentials is not supported on the the OS.
func (credentials *Credentials) IsSameUser(otherCredentials *Credentials) error {
	var _arg0 *C.GCredentials // out
	var _arg1 *C.GCredentials // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))
	_arg1 = (*C.GCredentials)(unsafe.Pointer(otherCredentials.Native()))

	C.g_credentials_is_same_user(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(credentials)
	runtime.KeepAlive(otherCredentials)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetNative copies the native credentials of type native_type from native into
// credentials.
//
// It is a programming error (which will cause a warning to be logged) to use
// this method if there is no #GCredentials support for the OS or if native_type
// isn't supported by the OS.
func (credentials *Credentials) SetNative(nativeType CredentialsType, native cgo.Handle) {
	var _arg0 *C.GCredentials    // out
	var _arg1 C.GCredentialsType // out
	var _arg2 C.gpointer         // out

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))
	_arg1 = C.GCredentialsType(nativeType)
	_arg2 = (C.gpointer)(unsafe.Pointer(native))

	C.g_credentials_set_native(_arg0, _arg1, _arg2)
	runtime.KeepAlive(credentials)
	runtime.KeepAlive(nativeType)
	runtime.KeepAlive(native)
}

// SetUnixUser tries to set the UNIX user identifier on credentials. This method
// is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if the
// native credentials type does not contain information about the UNIX user. It
// can also fail if the OS does not allow the use of "spoofed" credentials.
func (credentials *Credentials) SetUnixUser(uid uint) error {
	var _arg0 *C.GCredentials // out
	var _arg1 C.uid_t         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))
	_arg1 = C.uid_t(uid)

	C.g_credentials_set_unix_user(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(credentials)
	runtime.KeepAlive(uid)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// String creates a human-readable textual representation of credentials that
// can be used in logging and debug messages. The format of the returned string
// may change in future GLib release.
func (credentials *Credentials) String() string {
	var _arg0 *C.GCredentials // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))

	_cret = C.g_credentials_to_string(_arg0)
	runtime.KeepAlive(credentials)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
