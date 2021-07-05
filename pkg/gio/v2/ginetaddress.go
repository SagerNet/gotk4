// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_inet_address_get_type()), F: marshalInetAddress},
	})
}

// InetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to look up
// the Address for a hostname. Use g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a Address.
//
// To actually connect to a remote host, you will need a SocketAddress (which
// includes a Address as well as a port number).
type InetAddress interface {
	gextras.Objector

	// Equal checks if two Address instances are equal, e.g. the same address.
	Equal(otherAddress InetAddress) bool
	// Family gets @address's family
	Family() SocketFamily
	// IsAny tests whether @address is the "any" address for its family.
	IsAny() bool
	// IsLinkLocal tests whether @address is a link-local address (that is, if
	// it identifies a host on a local network that is not connected to the
	// Internet).
	IsLinkLocal() bool
	// IsLoopback tests whether @address is the loopback address for its family.
	IsLoopback() bool
	// IsMcGlobal tests whether @address is a global multicast address.
	IsMcGlobal() bool
	// IsMcLinkLocal tests whether @address is a link-local multicast address.
	IsMcLinkLocal() bool
	// IsMcNodeLocal tests whether @address is a node-local multicast address.
	IsMcNodeLocal() bool
	// IsMcOrgLocal tests whether @address is an organization-local multicast
	// address.
	IsMcOrgLocal() bool
	// IsMcSiteLocal tests whether @address is a site-local multicast address.
	IsMcSiteLocal() bool
	// IsMulticast tests whether @address is a multicast address.
	IsMulticast() bool
	// IsSiteLocal tests whether @address is a site-local address such as
	// 10.0.0.1 (that is, the address identifies a host on a local network that
	// can not be reached directly from the Internet, but which may have
	// outgoing Internet connectivity via a NAT or firewall).
	IsSiteLocal() bool
	// NativeSize gets the size of the native raw binary address for @address.
	// This is the size of the data that you get from g_inet_address_to_bytes().
	NativeSize() uint
	// String converts @address to string form.
	String() string
}

// inetAddress implements the InetAddress class.
type inetAddress struct {
	gextras.Objector
}

// WrapInetAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetAddress(obj *externglib.Object) InetAddress {
	return inetAddress{
		Objector: obj,
	}
}

func marshalInetAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetAddress(obj), nil
}

// NewInetAddressAny creates a Address for the "any" address (unassigned/"don't
// care") for @family.
func NewInetAddressAny(family SocketFamily) InetAddress {
	var _arg1 C.GSocketFamily // out
	var _cret *C.GInetAddress // in

	_arg1 = C.GSocketFamily(family)

	_cret = C.g_inet_address_new_any(_arg1)

	var _inetAddress InetAddress // out

	_inetAddress = WrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetAddress
}

// NewInetAddressFromString parses @string as an IP address and creates a new
// Address.
func NewInetAddressFromString(_string string) InetAddress {
	var _arg1 *C.gchar        // out
	var _cret *C.GInetAddress // in

	_arg1 = (*C.gchar)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_inet_address_new_from_string(_arg1)

	var _inetAddress InetAddress // out

	_inetAddress = WrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetAddress
}

// NewInetAddressLoopback creates a Address for the loopback address for
// @family.
func NewInetAddressLoopback(family SocketFamily) InetAddress {
	var _arg1 C.GSocketFamily // out
	var _cret *C.GInetAddress // in

	_arg1 = C.GSocketFamily(family)

	_cret = C.g_inet_address_new_loopback(_arg1)

	var _inetAddress InetAddress // out

	_inetAddress = WrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _inetAddress
}

func (a inetAddress) Equal(otherAddress InetAddress) bool {
	var _arg0 *C.GInetAddress // out
	var _arg1 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(otherAddress.Native()))

	_cret = C.g_inet_address_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) Family() SocketFamily {
	var _arg0 *C.GInetAddress // out
	var _cret C.GSocketFamily // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_family(_arg0)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

func (a inetAddress) IsAny() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_any(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsLinkLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_link_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsLoopback() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_loopback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMcGlobal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_mc_global(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMcLinkLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_mc_link_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMcNodeLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_mc_node_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMcOrgLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_mc_org_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMcSiteLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_mc_site_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsMulticast() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_multicast(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) IsSiteLocal() bool {
	var _arg0 *C.GInetAddress // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_is_site_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a inetAddress) NativeSize() uint {
	var _arg0 *C.GInetAddress // out
	var _cret C.gsize         // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_get_native_size(_arg0)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

func (a inetAddress) String() string {
	var _arg0 *C.GInetAddress // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GInetAddress)(unsafe.Pointer(a.Native()))

	_cret = C.g_inet_address_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
