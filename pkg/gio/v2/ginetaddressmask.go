// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_inet_address_mask_get_type()), F: marshalInetAddressMask},
	})
}

// InetAddressMask represents a range of IPv4 or IPv6 addresses described by a
// base address and a length indicating how many bits of the base address are
// relevant for matching purposes. These are often given in string form. Eg,
// "10.0.0.0/8", or "fe80::/10".
type InetAddressMask interface {
	Initable

	EqualInetAddressMask(mask2 InetAddressMask) bool

	Address() InetAddress

	Family() SocketFamily

	Length() uint

	MatchesInetAddressMask(address InetAddress) bool

	String() string
}

// inetAddressMask implements the InetAddressMask class.
type inetAddressMask struct {
	gextras.Objector
}

// WrapInetAddressMask wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetAddressMask(obj *externglib.Object) InetAddressMask {
	return inetAddressMask{
		Objector: obj,
	}
}

func marshalInetAddressMask(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetAddressMask(obj), nil
}

func NewInetAddressMask(addr InetAddress, length uint) (InetAddressMask, error) {
	var _arg1 *C.GInetAddress     // out
	var _arg2 C.guint             // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GInetAddress)(unsafe.Pointer(addr.Native()))
	_arg2 = C.guint(length)

	_cret = C.g_inet_address_mask_new(_arg1, _arg2, &_cerr)

	var _inetAddressMask InetAddressMask // out
	var _goerr error                     // out

	_inetAddressMask = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InetAddressMask)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inetAddressMask, _goerr
}

func NewInetAddressMaskFromString(maskString string) (InetAddressMask, error) {
	var _arg1 *C.gchar            // out
	var _cret *C.GInetAddressMask // in
	var _cerr *C.GError           // in

	_arg1 = (*C.gchar)(C.CString(maskString))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_inet_address_mask_new_from_string(_arg1, &_cerr)

	var _inetAddressMask InetAddressMask // out
	var _goerr error                     // out

	_inetAddressMask = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InetAddressMask)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inetAddressMask, _goerr
}

func (m inetAddressMask) EqualInetAddressMask(mask2 InetAddressMask) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddressMask // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GInetAddressMask)(unsafe.Pointer(mask2.Native()))

	_cret = C.g_inet_address_mask_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m inetAddressMask) Address() InetAddress {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.GInetAddress     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	_cret = C.g_inet_address_mask_get_address(_arg0)

	var _inetAddress InetAddress // out

	_inetAddress = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(InetAddress)

	return _inetAddress
}

func (m inetAddressMask) Family() SocketFamily {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.GSocketFamily     // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	_cret = C.g_inet_address_mask_get_family(_arg0)

	var _socketFamily SocketFamily // out

	_socketFamily = SocketFamily(_cret)

	return _socketFamily
}

func (m inetAddressMask) Length() uint {
	var _arg0 *C.GInetAddressMask // out
	var _cret C.guint             // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	_cret = C.g_inet_address_mask_get_length(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (m inetAddressMask) MatchesInetAddressMask(address InetAddress) bool {
	var _arg0 *C.GInetAddressMask // out
	var _arg1 *C.GInetAddress     // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_inet_address_mask_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m inetAddressMask) String() string {
	var _arg0 *C.GInetAddressMask // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	_cret = C.g_inet_address_mask_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (i inetAddressMask) Init(cancellable Cancellable) error {
	return WrapInitable(gextras.InternObject(i)).Init(cancellable)
}