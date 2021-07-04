// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_inet_socket_address_get_type()), F: marshalInetSocketAddress},
	})
}

// InetSocketAddress: an IPv4 or IPv6 socket address; that is, the combination
// of a Address and a port number.
type InetSocketAddress interface {
	SocketAddress

	// Address:
	Address() InetAddress
	// Flowinfo:
	Flowinfo() uint32
	// Port:
	Port() uint16
	// ScopeID:
	ScopeID() uint32
}

// inetSocketAddress implements the InetSocketAddress class.
type inetSocketAddress struct {
	SocketAddress
}

// WrapInetSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetSocketAddress(obj *externglib.Object) InetSocketAddress {
	return inetSocketAddress{
		SocketAddress: WrapSocketAddress(obj),
	}
}

func marshalInetSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetSocketAddress(obj), nil
}

// NewInetSocketAddress:
func NewInetSocketAddress(address InetAddress, port uint16) InetSocketAddress {
	var _arg1 *C.GInetAddress   // out
	var _arg2 C.guint16         // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	_arg2 = C.guint16(port)

	_cret = C.g_inet_socket_address_new(_arg1, _arg2)

	var _inetSocketAddress InetSocketAddress // out

	_inetSocketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InetSocketAddress)

	return _inetSocketAddress
}

// NewInetSocketAddressFromString:
func NewInetSocketAddressFromString(address string, port uint) InetSocketAddress {
	var _arg1 *C.char           // out
	var _arg2 C.guint           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (*C.char)(C.CString(address))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(port)

	_cret = C.g_inet_socket_address_new_from_string(_arg1, _arg2)

	var _inetSocketAddress InetSocketAddress // out

	_inetSocketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InetSocketAddress)

	return _inetSocketAddress
}

func (a inetSocketAddress) Address() InetAddress {
	var _arg0 *C.GInetSocketAddress // out
	var _cret *C.GInetAddress       // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_socket_address_get_address(_arg0)

	var _inetAddress InetAddress // out

	_inetAddress = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(InetAddress)

	return _inetAddress
}

func (a inetSocketAddress) Flowinfo() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_socket_address_get_flowinfo(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (a inetSocketAddress) Port() uint16 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint16             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_socket_address_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (a inetSocketAddress) ScopeID() uint32 {
	var _arg0 *C.GInetSocketAddress // out
	var _cret C.guint32             // in

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_socket_address_get_scope_id(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (c inetSocketAddress) Enumerate() SocketAddressEnumerator {
	return WrapSocketConnectable(gextras.InternObject(c)).Enumerate()
}

func (c inetSocketAddress) ProxyEnumerate() SocketAddressEnumerator {
	return WrapSocketConnectable(gextras.InternObject(c)).ProxyEnumerate()
}

func (c inetSocketAddress) String() string {
	return WrapSocketConnectable(gextras.InternObject(c)).String()
}
