// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib.h>
import "C"

// HostnameIsASCIIEncoded tests if @hostname contains segments with an
// ASCII-compatible encoding of an Internationalized Domain Name. If this
// returns true, you should decode the hostname with g_hostname_to_unicode()
// before displaying it to the user.
//
// Note that a hostname might contain a mix of encoded and unencoded segments,
// and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return true for a name.
func HostnameIsASCIIEncoded(hostname string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_hostname_is_ascii_encoded(hostname)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// HostnameIsIpAddress tests if @hostname is the string form of an IPv4 or IPv6
// address. (Eg, "192.168.0.1".)
//
// Since 2.66, IPv6 addresses with a zone-id are accepted (RFC6874).
func HostnameIsIpAddress(hostname string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_hostname_is_ip_address(hostname)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// HostnameIsNonASCII tests if @hostname contains Unicode characters. If this
// returns true, you need to encode the hostname with g_hostname_to_ascii()
// before using it in non-IDN-aware contexts.
//
// Note that a hostname might contain a mix of encoded and unencoded segments,
// and so it is possible for g_hostname_is_non_ascii() and
// g_hostname_is_ascii_encoded() to both return true for a name.
func HostnameIsNonASCII(hostname string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_hostname_is_non_ascii(hostname)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// HostnameToASCII converts @hostname to its canonical ASCII form; an ASCII-only
// string containing no uppercase letters and not ending with a trailing dot.
func HostnameToASCII(hostname string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_hostname_to_ascii(hostname)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// HostnameToUnicode converts @hostname to its canonical presentation form; a
// UTF-8 string in Unicode normalization form C, containing no uppercase
// letters, no forbidden characters, and no ASCII-encoded segments, and not
// ending with a trailing dot.
//
// Of course if @hostname is not an internationalized hostname, then the
// canonical presentation form will be entirely ASCII.
func HostnameToUnicode(hostname string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_hostname_to_unicode(hostname)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}