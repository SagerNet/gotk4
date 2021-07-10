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
		{T: externglib.Type(C.g_tls_password_get_type()), F: marshalTLSPassworder},
	})
}

// TLSPassworderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSPassworderOverrider interface {
	DefaultWarning() string
	// Value: get the password value. If @length is not nil then it will be
	// filled in with the length of the password value. (Note that the password
	// value is not nul-terminated, so you can only pass nil for @length in
	// contexts where you know the password will have a certain fixed length.)
	Value(length *uint) *byte
}

// TLSPassworder describes TLSPassword's methods.
type TLSPassworder interface {
	gextras.Objector

	Description() string
	Flags() TLSPasswordFlags
	Value(length *uint) *byte
	Warning() string
	SetDescription(description string)
	SetValue(value []byte)
	SetWarning(warning string)
}

// TLSPassword holds a password used in TLS.
type TLSPassword struct {
	*externglib.Object
}

var _ TLSPassworder = (*TLSPassword)(nil)

func wrapTLSPassworder(obj *externglib.Object) TLSPassworder {
	return &TLSPassword{
		Object: obj,
	}
}

func marshalTLSPassworder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSPassworder(obj), nil
}

// Description: get a description string about what the password will be used
// for.
func (password *TLSPassword) Description() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Flags: get flags about the password.
func (password *TLSPassword) Flags() TLSPasswordFlags {
	var _arg0 *C.GTlsPassword     // out
	var _cret C.GTlsPasswordFlags // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_flags(_arg0)

	var _tlsPasswordFlags TLSPasswordFlags // out

	_tlsPasswordFlags = (TLSPasswordFlags)(_cret)

	return _tlsPasswordFlags
}

// Value: get the password value. If @length is not nil then it will be filled
// in with the length of the password value. (Note that the password value is
// not nul-terminated, so you can only pass nil for @length in contexts where
// you know the password will have a certain fixed length.)
func (password *TLSPassword) Value(length *uint) *byte {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gsize        // out
	var _cret *C.guchar       // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = (*C.gsize)(unsafe.Pointer(length))

	_cret = C.g_tls_password_get_value(_arg0, _arg1)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

// Warning: get a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
func (password *TLSPassword) Warning() string {
	var _arg0 *C.GTlsPassword // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))

	_cret = C.g_tls_password_get_warning(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetDescription: set a description string about what the password will be used
// for.
func (password *TLSPassword) SetDescription(description string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_description(_arg0, _arg1)
}

// SetValue: set the value for this password. The @value will be copied by the
// password object.
//
// Specify the @length, for a non-nul-terminated password. Pass -1 as @length if
// using a nul-terminated password, and @length will be calculated
// automatically. (Note that the terminating nul is not considered part of the
// password in this case.)
func (password *TLSPassword) SetValue(value []byte) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.guchar
	var _arg2 C.gssize

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg2 = C.gssize(len(value))
	_arg1 = (*C.guchar)(unsafe.Pointer(&value[0]))

	C.g_tls_password_set_value(_arg0, _arg1, _arg2)
}

// SetWarning: set a user readable translated warning. Usually this warning is a
// representation of the password flags returned from
// g_tls_password_get_flags().
func (password *TLSPassword) SetWarning(warning string) {
	var _arg0 *C.GTlsPassword // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg1 = (*C.gchar)(C.CString(warning))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_tls_password_set_warning(_arg0, _arg1)
}
