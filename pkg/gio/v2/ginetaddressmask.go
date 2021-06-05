// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_inet_address_mask_get_type()), F: marshalInetAddressMask},
	})
}

// InetAddressMask represents a range of IPv4 or IPv6 addresses described by a
// base address and a length indicating how many bits of the base address are
// relevant for matching purposes. These are often given in string form. Eg,
// "10.0.0.0/8", or "fe80::/10".
type InetAddressMask interface {
	gextras.Objector
	Initable

	// Equal tests if @mask and @mask2 are the same mask.
	Equal(mask2 InetAddressMask) bool
	// Address gets @mask's base address
	Address() InetAddress
	// Family gets the Family of @mask's address
	Family() SocketFamily
	// Length gets @mask's length
	Length() uint
	// Matches tests if @address falls within the range described by @mask.
	Matches(address InetAddress) bool
	// String converts @mask back to its corresponding string form.
	String() string
}

// inetAddressMask implements the InetAddressMask interface.
type inetAddressMask struct {
	gextras.Objector
	Initable
}

var _ InetAddressMask = (*inetAddressMask)(nil)

// WrapInetAddressMask wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetAddressMask(obj *externglib.Object) InetAddressMask {
	return InetAddressMask{
		Objector: obj,
		Initable: WrapInitable(obj),
	}
}

func marshalInetAddressMask(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetAddressMask(obj), nil
}

// NewInetAddressMask constructs a class InetAddressMask.
func NewInetAddressMask(addr InetAddress, length uint) (inetAddressMask InetAddressMask, err error) {
	var arg1 *C.GInetAddress
	var arg2 C.guint
	var errout *C.GError

	arg1 = (*C.GInetAddress)(unsafe.Pointer(addr.Native()))
	arg2 = C.guint(length)

	var cret C.GInetAddressMask
	var goret1 InetAddressMask
	var goerr error

	cret = C.g_inet_address_mask_new(addr, length, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(InetAddressMask)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// NewInetAddressMaskFromString constructs a class InetAddressMask.
func NewInetAddressMaskFromString(maskString string) (inetAddressMask InetAddressMask, err error) {
	var arg1 *C.gchar
	var errout *C.GError

	arg1 = (*C.gchar)(C.CString(maskString))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GInetAddressMask
	var goret1 InetAddressMask
	var goerr error

	cret = C.g_inet_address_mask_new_from_string(maskString, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(InetAddressMask)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Equal tests if @mask and @mask2 are the same mask.
func (m inetAddressMask) Equal(mask2 InetAddressMask) bool {
	var arg0 *C.GInetAddressMask
	var arg1 *C.GInetAddressMask

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GInetAddressMask)(unsafe.Pointer(mask2.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_inet_address_mask_equal(arg0, mask2)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Address gets @mask's base address
func (m inetAddressMask) Address() InetAddress {
	var arg0 *C.GInetAddressMask

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	var cret *C.GInetAddress
	var goret1 InetAddress

	cret = C.g_inet_address_mask_get_address(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(InetAddress)

	return goret1
}

// Family gets the Family of @mask's address
func (m inetAddressMask) Family() SocketFamily {
	var arg0 *C.GInetAddressMask

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	var cret C.GSocketFamily
	var goret1 SocketFamily

	cret = C.g_inet_address_mask_get_family(arg0)

	goret1 = SocketFamily(cret)

	return goret1
}

// Length gets @mask's length
func (m inetAddressMask) Length() uint {
	var arg0 *C.GInetAddressMask

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	var cret C.guint
	var goret1 uint

	cret = C.g_inet_address_mask_get_length(arg0)

	goret1 = C.guint(cret)

	return goret1
}

// Matches tests if @address falls within the range described by @mask.
func (m inetAddressMask) Matches(address InetAddress) bool {
	var arg0 *C.GInetAddressMask
	var arg1 *C.GInetAddress

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_inet_address_mask_matches(arg0, address)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// String converts @mask back to its corresponding string form.
func (m inetAddressMask) String() string {
	var arg0 *C.GInetAddressMask

	arg0 = (*C.GInetAddressMask)(unsafe.Pointer(m.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_inet_address_mask_to_string(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

type InetAddressMaskPrivate struct {
	native C.GInetAddressMaskPrivate
}

// WrapInetAddressMaskPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInetAddressMaskPrivate(ptr unsafe.Pointer) *InetAddressMaskPrivate {
	if ptr == nil {
		return nil
	}

	return (*InetAddressMaskPrivate)(ptr)
}

func marshalInetAddressMaskPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapInetAddressMaskPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *InetAddressMaskPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
