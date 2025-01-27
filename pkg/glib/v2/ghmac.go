// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// ComputeHMACForBytes computes the HMAC for a binary data. This is a
// convenience wrapper for g_hmac_new(), g_hmac_get_string() and g_hmac_unref().
//
// The hexadecimal string returned will be in lower case.
func ComputeHMACForBytes(digestType ChecksumType, key *Bytes, data *Bytes) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.GBytes       // out
	var _arg3 *C.GBytes       // out
	var _cret *C.gchar        // in

	_arg1 = C.GChecksumType(digestType)
	_arg2 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(key)))
	_arg3 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(data)))

	_cret = C.g_compute_hmac_for_bytes(_arg1, _arg2, _arg3)
	runtime.KeepAlive(digestType)
	runtime.KeepAlive(key)
	runtime.KeepAlive(data)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ComputeHMACForData computes the HMAC for a binary data of length. This is a
// convenience wrapper for g_hmac_new(), g_hmac_get_string() and g_hmac_unref().
//
// The hexadecimal string returned will be in lower case.
func ComputeHMACForData(digestType ChecksumType, key []byte, data []byte) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar       // out
	var _arg3 C.gsize
	var _arg4 *C.guchar // out
	var _arg5 C.gsize
	var _cret *C.gchar // in

	_arg1 = C.GChecksumType(digestType)
	_arg3 = (C.gsize)(len(key))
	if len(key) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&key[0]))
	}
	_arg5 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg4 = (*C.guchar)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_compute_hmac_for_data(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(digestType)
	runtime.KeepAlive(key)
	runtime.KeepAlive(data)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ComputeHMACForString computes the HMAC for a string.
//
// The hexadecimal string returned will be in lower case.
func ComputeHMACForString(digestType ChecksumType, key []byte, str string, length int) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar       // out
	var _arg3 C.gsize
	var _arg4 *C.gchar // out
	var _arg5 C.gssize // out
	var _cret *C.gchar // in

	_arg1 = C.GChecksumType(digestType)
	_arg3 = (C.gsize)(len(key))
	if len(key) > 0 {
		_arg2 = (*C.guchar)(unsafe.Pointer(&key[0]))
	}
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gssize(length)

	_cret = C.g_compute_hmac_for_string(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(digestType)
	runtime.KeepAlive(key)
	runtime.KeepAlive(str)
	runtime.KeepAlive(length)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
