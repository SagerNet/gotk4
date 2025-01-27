// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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

// DBusEscapeObjectPath: this is a language binding friendly version of
// g_dbus_escape_object_path_bytestring().
func DBusEscapeObjectPath(s string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_escape_object_path(_arg1)
	runtime.KeepAlive(s)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusEscapeObjectPathBytestring escapes bytes for use in a D-Bus object path
// component. bytes is an array of zero or more nonzero bytes in an unspecified
// encoding, followed by a single zero byte.
//
// The escaping method consists of replacing all non-alphanumeric characters
// (see g_ascii_isalnum()) with their hexadecimal value preceded by an
// underscore (_). For example: foo.bar.baz will become foo_2ebar_2ebaz.
//
// This method is appropriate to use when the input is nearly a valid object
// path component but is not when your input is far from being a valid object
// path component. Other escaping algorithms are also valid to use with D-Bus
// object paths.
//
// This can be reversed with g_dbus_unescape_object_path().
func DBusEscapeObjectPathBytestring(bytes []byte) string {
	var _arg1 *C.guint8 // out
	var _cret *C.gchar  // in

	{
		var zero byte
		bytes = append(bytes, zero)
		_arg1 = (*C.guint8)(unsafe.Pointer(&bytes[0]))
	}

	_cret = C.g_dbus_escape_object_path_bytestring(_arg1)
	runtime.KeepAlive(bytes)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusGenerateGUID: generate a D-Bus GUID that can be used with e.g.
// g_dbus_connection_new().
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusGenerateGUID() string {
	var _cret *C.gchar // in

	_cret = C.g_dbus_generate_guid()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusGValueToGVariant converts a #GValue to a #GVariant of the type indicated
// by the type parameter.
//
// The conversion is using the following rules:
//
// - TYPE_STRING: 's', 'o', 'g' or 'ay'
//
// - TYPE_STRV: 'as', 'ao' or 'aay'
//
// - TYPE_BOOLEAN: 'b'
//
// - TYPE_UCHAR: 'y'
//
// - TYPE_INT: 'i', 'n'
//
// - TYPE_UINT: 'u', 'q'
//
// - TYPE_INT64 'x'
//
// - TYPE_UINT64: 't'
//
// - TYPE_DOUBLE: 'd'
//
// - TYPE_VARIANT: Any Type
//
// This can fail if e.g. gvalue is of type TYPE_STRING and type is
// ['i'][G-VARIANT-TYPE-INT32:CAPS]. It will also fail for any #GType (including
// e.g. TYPE_OBJECT and TYPE_BOXED derived-types) not in the table above.
//
// Note that if gvalue is of type TYPE_VARIANT and its value is NULL, the empty
// #GVariant instance (never NULL) for type is returned (e.g. 0 for scalar
// types, the empty string for string types, '/' for object path types, the
// empty array for any array type and so on).
//
// See the g_dbus_gvariant_to_gvalue() function for how to convert a #GVariant
// to a #GValue.
func DBusGValueToGVariant(gvalue *externglib.Value, typ *glib.VariantType) *glib.Variant {
	var _arg1 *C.GValue       // out
	var _arg2 *C.GVariantType // out
	var _cret *C.GVariant     // in

	_arg1 = (*C.GValue)(unsafe.Pointer(gvalue.Native()))
	_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(typ)))

	_cret = C.g_dbus_gvalue_to_gvariant(_arg1, _arg2)
	runtime.KeepAlive(gvalue)
	runtime.KeepAlive(typ)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// DBusGVariantToGValue converts a #GVariant to a #GValue. If value is floating,
// it is consumed.
//
// The rules specified in the g_dbus_gvalue_to_gvariant() function are used -
// this function is essentially its reverse form. So, a #GVariant containing any
// basic or string array type will be converted to a #GValue containing a basic
// value or string array. Any other #GVariant (handle, variant, tuple, dict
// entry) will be converted to a #GValue containing that #GVariant.
//
// The conversion never fails - a valid #GValue is always returned in
// out_gvalue.
func DBusGVariantToGValue(value *glib.Variant) externglib.Value {
	var _arg1 *C.GVariant // out
	var _arg2 C.GValue    // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_dbus_gvariant_to_gvalue(_arg1, &_arg2)
	runtime.KeepAlive(value)

	var _outGvalue externglib.Value // out

	_outGvalue = *externglib.ValueFromNative(unsafe.Pointer((&_arg2)))

	return _outGvalue
}

// DBusIsGUID checks if string is a D-Bus GUID.
//
// See the D-Bus specification regarding what strings are valid D-Bus GUID (for
// example, D-Bus GUIDs are not RFC-4122 compliant).
func DBusIsGUID(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_guid(_arg1)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsInterfaceName checks if string is a valid D-Bus interface name.
func DBusIsInterfaceName(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_interface_name(_arg1)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsMemberName checks if string is a valid D-Bus member (e.g. signal or
// method) name.
func DBusIsMemberName(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_member_name(_arg1)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsName checks if string is a valid D-Bus bus name (either unique or
// well-known).
func DBusIsName(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_name(_arg1)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusIsUniqueName checks if string is a valid D-Bus unique bus name.
func DBusIsUniqueName(_string string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_is_unique_name(_arg1)
	runtime.KeepAlive(_string)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusUnescapeObjectPath unescapes an string that was previously escaped with
// g_dbus_escape_object_path(). If the string is in a format that could not have
// been returned by g_dbus_escape_object_path(), this function returns NULL.
//
// Encoding alphanumeric characters which do not need to be encoded is not
// allowed (e.g _63 is not valid, the string should contain c instead).
func DBusUnescapeObjectPath(s string) []byte {
	var _arg1 *C.gchar  // out
	var _cret *C.guint8 // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(s)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_unescape_object_path(_arg1)
	runtime.KeepAlive(s)

	var _guint8s []byte // out

	if _cret != nil {
		{
			var i int
			var z C.guint8
			for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_guint8s = make([]byte, i)
			for i := range src {
				_guint8s[i] = byte(src[i])
			}
		}
	}

	return _guint8s
}
