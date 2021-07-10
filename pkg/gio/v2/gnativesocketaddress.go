// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.g_native_socket_address_get_type()), F: marshalNativeSocketAddresser},
	})
}

// NativeSocketAddresser describes NativeSocketAddress's methods.
type NativeSocketAddresser interface {
	gextras.Objector

	privateNativeSocketAddress()
}

// NativeSocketAddress: socket address of some unknown native type.
type NativeSocketAddress struct {
	*externglib.Object
	SocketAddress
	SocketConnectable
}

var _ NativeSocketAddresser = (*NativeSocketAddress)(nil)

func wrapNativeSocketAddresser(obj *externglib.Object) NativeSocketAddresser {
	return &NativeSocketAddress{
		Object: obj,
		SocketAddress: SocketAddress{
			Object: obj,
			SocketConnectable: SocketConnectable{
				Object: obj,
			},
		},
		SocketConnectable: SocketConnectable{
			Object: obj,
		},
	}
}

func marshalNativeSocketAddresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNativeSocketAddresser(obj), nil
}

// NewNativeSocketAddress creates a new SocketAddress for @native and @len.
func NewNativeSocketAddress(native interface{}, len uint) *NativeSocketAddress {
	var _arg1 C.gpointer        // out
	var _arg2 C.gsize           // out
	var _cret *C.GSocketAddress // in

	_arg1 = (C.gpointer)(box.Assign(native))
	_arg2 = C.gsize(len)

	_cret = C.g_native_socket_address_new(_arg1, _arg2)

	var _nativeSocketAddress *NativeSocketAddress // out

	_nativeSocketAddress = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NativeSocketAddress)

	return _nativeSocketAddress
}

func (*NativeSocketAddress) privateNativeSocketAddress() {}
