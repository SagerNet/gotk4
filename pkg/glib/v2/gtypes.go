// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// CompareDataFunc specifies the type of a comparison function used to compare
// two values. The function should return a negative integer if the first value
// comes before the second, 0 if they are equal, or a positive integer if the
// first value comes after the second.
type CompareDataFunc func(a cgo.Handle, b cgo.Handle) (gint int)

//export _gotk4_glib2_CompareDataFunc
func _gotk4_glib2_CompareDataFunc(arg0 C.gconstpointer, arg1 C.gconstpointer, arg2 C.gpointer) (cret C.gint) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var a cgo.Handle // out
	var b cgo.Handle // out

	a = (cgo.Handle)(unsafe.Pointer(arg0))
	b = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(CompareDataFunc)
	gint := fn(a, b)

	cret = C.gint(gint)

	return cret
}

// Func specifies the type of functions passed to g_list_foreach() and
// g_slist_foreach().
type Func func(data cgo.Handle)

//export _gotk4_glib2_Func
func _gotk4_glib2_Func(arg0 C.gpointer, arg1 C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var data cgo.Handle // out

	data = (cgo.Handle)(unsafe.Pointer(arg0))

	fn := v.(Func)
	fn(data)
}

// HFunc specifies the type of the function passed to g_hash_table_foreach(). It
// is called with each key/value pair, together with the user_data parameter
// which is passed to g_hash_table_foreach().
type HFunc func(key cgo.Handle, value cgo.Handle)

//export _gotk4_glib2_HFunc
func _gotk4_glib2_HFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key cgo.Handle   // out
	var value cgo.Handle // out

	key = (cgo.Handle)(unsafe.Pointer(arg0))
	value = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(HFunc)
	fn(key, value)
}

// TimeVal represents a precise time, with seconds and microseconds. Similar to
// the struct timeval returned by the gettimeofday() UNIX system call.
//
// GLib is attempting to unify around the use of 64-bit integers to represent
// microsecond-precision time. As such, this type will be removed from a future
// version of GLib. A consequence of using glong for tv_sec is that on 32-bit
// systems GTimeVal is subject to the year 2038 problem.
//
// Deprecated: Use Time or #guint64 instead.
type TimeVal struct {
	native C.GTimeVal
}

// Native returns the underlying C source pointer.
func (t *TimeVal) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TvSec: seconds
func (t *TimeVal) TvSec() int32 {
	var v int32 // out
	v = int32(t.native.tv_sec)
	return v
}

// TvUsec: microseconds
func (t *TimeVal) TvUsec() int32 {
	var v int32 // out
	v = int32(t.native.tv_usec)
	return v
}

// Add adds the given number of microseconds to time_. microseconds can also be
// negative to decrease the value of time_.
//
// Deprecated: Val is not year-2038-safe. Use guint64 for representing
// microseconds since the epoch, or use Time.
func (time_ *TimeVal) Add(microseconds int32) {
	var _arg0 *C.GTimeVal // out
	var _arg1 C.glong     // out

	_arg0 = (*C.GTimeVal)(unsafe.Pointer(time_))
	_arg1 = C.glong(microseconds)

	C.g_time_val_add(_arg0, _arg1)
}

// ToISO8601 converts time_ into an RFC 3339 encoded string, relative to the
// Coordinated Universal Time (UTC). This is one of the many formats allowed by
// ISO 8601.
//
// ISO 8601 allows a large number of date/time formats, with or without
// punctuation and optional elements. The format returned by this function is a
// complete date and time, with optional punctuation included, the UTC time zone
// represented as "Z", and the tv_usec part included if and only if it is
// nonzero, i.e. either "YYYY-MM-DDTHH:MM:SSZ" or "YYYY-MM-DDTHH:MM:SS.fffffZ".
//
// This corresponds to the Internet date/time format defined by RFC 3339
// (https://www.ietf.org/rfc/rfc3339.txt), and to either of the two most-precise
// formats defined by the W3C Note Date and Time Formats
// (http://www.w3.org/TR/NOTE-datetime-19980827). Both of these documents are
// profiles of ISO 8601.
//
// Use g_date_time_format() or g_strdup_printf() if a different variation of ISO
// 8601 format is required.
//
// If time_ represents a date which is too large to fit into a struct tm, NULL
// will be returned. This is platform dependent. Note also that since GTimeVal
// stores the number of seconds as a glong, on 32-bit systems it is subject to
// the year 2038 problem. Accordingly, since GLib 2.62, this function has been
// deprecated. Equivalent functionality is available using:
//
//    GDateTime *dt = g_date_time_new_from_unix_utc (time_val);
//    iso8601_string = g_date_time_format_iso8601 (dt);
//    g_date_time_unref (dt);
//
// The return value of g_time_val_to_iso8601() has been nullable since GLib
// 2.54; before then, GLib would crash under the same conditions.
//
// Deprecated: Val is not year-2038-safe. Use g_date_time_format_iso8601(dt)
// instead.
func (time_ *TimeVal) ToISO8601() string {
	var _arg0 *C.GTimeVal // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GTimeVal)(unsafe.Pointer(time_))

	_cret = C.g_time_val_to_iso8601(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
