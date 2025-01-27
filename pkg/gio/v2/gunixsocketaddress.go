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
		{T: externglib.Type(C.g_unix_socket_address_get_type()), F: marshalUnixSocketAddresser},
	})
}

// UnixSocketAddress: support for UNIX-domain (also known as local) sockets.
//
// UNIX domain sockets are generally visible in the filesystem. However, some
// systems support abstract socket names which are not visible in the filesystem
// and not affected by the filesystem permissions, visibility, etc. Currently
// this is only supported under Linux. If you attempt to use abstract sockets on
// other systems, function calls may return G_IO_ERROR_NOT_SUPPORTED errors. You
// can use g_unix_socket_address_abstract_names_supported() to see if abstract
// names are supported.
//
// Note that <gio/gunixsocketaddress.h> belongs to the UNIX-specific GIO
// interfaces, thus you have to use the gio-unix-2.0.pc pkg-config file when
// using it.
type UnixSocketAddress struct {
	SocketAddress
}

func WrapUnixSocketAddress(obj *externglib.Object) *UnixSocketAddress {
	return &UnixSocketAddress{
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
	}
}

func marshalUnixSocketAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixSocketAddress(obj), nil
}

// NewUnixSocketAddress creates a new SocketAddress for path.
//
// To create abstract socket addresses, on systems that support that, use
// g_unix_socket_address_new_abstract().
func NewUnixSocketAddress(path string) *UnixSocketAddress {
	var _arg1 *C.gchar          // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_unix_socket_address_new(_arg1)
	runtime.KeepAlive(path)

	var _unixSocketAddress *UnixSocketAddress // out

	_unixSocketAddress = WrapUnixSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixSocketAddress
}

// NewUnixSocketAddressAbstract creates a new
// G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED SocketAddress for path.
//
// Deprecated: Use g_unix_socket_address_new_with_type().
func NewUnixSocketAddressAbstract(path []byte) *UnixSocketAddress {
	var _arg1 *C.gchar // out
	var _arg2 C.gint
	var _cret *C.GSocketAddress // in

	_arg2 = (C.gint)(len(path))
	if len(path) > 0 {
		_arg1 = (*C.gchar)(unsafe.Pointer(&path[0]))
	}

	_cret = C.g_unix_socket_address_new_abstract(_arg1, _arg2)
	runtime.KeepAlive(path)

	var _unixSocketAddress *UnixSocketAddress // out

	_unixSocketAddress = WrapUnixSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixSocketAddress
}

// NewUnixSocketAddressWithType creates a new SocketAddress of type type with
// name path.
//
// If type is G_UNIX_SOCKET_ADDRESS_PATH, this is equivalent to calling
// g_unix_socket_address_new().
//
// If type is G_UNIX_SOCKET_ADDRESS_ANONYMOUS, path and path_len will be
// ignored.
//
// If path_type is G_UNIX_SOCKET_ADDRESS_ABSTRACT, then path_len bytes of path
// will be copied to the socket's path, and only those bytes will be considered
// part of the name. (If path_len is -1, then path is assumed to be
// NUL-terminated.) For example, if path was "test", then calling
// g_socket_address_get_native_size() on the returned socket would return 7 (2
// bytes of overhead, 1 byte for the abstract-socket indicator byte, and 4 bytes
// for the name "test").
//
// If path_type is G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED, then path_len bytes of
// path will be copied to the socket's path, the rest of the path will be padded
// with 0 bytes, and the entire zero-padded buffer will be considered the name.
// (As above, if path_len is -1, then path is assumed to be NUL-terminated.) In
// this case, g_socket_address_get_native_size() will always return the full
// size of a struct sockaddr_un, although g_unix_socket_address_get_path_len()
// will still return just the length of path.
//
// G_UNIX_SOCKET_ADDRESS_ABSTRACT is preferred over
// G_UNIX_SOCKET_ADDRESS_ABSTRACT_PADDED for new programs. Of course, when
// connecting to a server created by another process, you must use the
// appropriate type corresponding to how that process created its listening
// socket.
func NewUnixSocketAddressWithType(path []byte, typ UnixSocketAddressType) *UnixSocketAddress {
	var _arg1 *C.gchar // out
	var _arg2 C.gint
	var _arg3 C.GUnixSocketAddressType // out
	var _cret *C.GSocketAddress        // in

	_arg2 = (C.gint)(len(path))
	if len(path) > 0 {
		_arg1 = (*C.gchar)(unsafe.Pointer(&path[0]))
	}
	_arg3 = C.GUnixSocketAddressType(typ)

	_cret = C.g_unix_socket_address_new_with_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(path)
	runtime.KeepAlive(typ)

	var _unixSocketAddress *UnixSocketAddress // out

	_unixSocketAddress = WrapUnixSocketAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixSocketAddress
}

// AddressType gets address's type.
func (address *UnixSocketAddress) AddressType() UnixSocketAddressType {
	var _arg0 *C.GUnixSocketAddress    // out
	var _cret C.GUnixSocketAddressType // in

	_arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_unix_socket_address_get_address_type(_arg0)
	runtime.KeepAlive(address)

	var _unixSocketAddressType UnixSocketAddressType // out

	_unixSocketAddressType = UnixSocketAddressType(_cret)

	return _unixSocketAddressType
}

// IsAbstract tests if address is abstract.
//
// Deprecated: Use g_unix_socket_address_get_address_type().
func (address *UnixSocketAddress) IsAbstract() bool {
	var _arg0 *C.GUnixSocketAddress // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_unix_socket_address_get_is_abstract(_arg0)
	runtime.KeepAlive(address)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Path gets address's path, or for abstract sockets the "name".
//
// Guaranteed to be zero-terminated, but an abstract socket may contain embedded
// zeros, and thus you should use g_unix_socket_address_get_path_len() to get
// the true length of this string.
func (address *UnixSocketAddress) Path() string {
	var _arg0 *C.GUnixSocketAddress // out
	var _cret *C.char               // in

	_arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_unix_socket_address_get_path(_arg0)
	runtime.KeepAlive(address)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PathLen gets the length of address's path.
//
// For details, see g_unix_socket_address_get_path().
func (address *UnixSocketAddress) PathLen() uint {
	var _arg0 *C.GUnixSocketAddress // out
	var _cret C.gsize               // in

	_arg0 = (*C.GUnixSocketAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_unix_socket_address_get_path_len(_arg0)
	runtime.KeepAlive(address)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// UnixSocketAddressAbstractNamesSupported checks if abstract UNIX domain socket
// names are supported.
func UnixSocketAddressAbstractNamesSupported() bool {
	var _cret C.gboolean // in

	_cret = C.g_unix_socket_address_abstract_names_supported()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
