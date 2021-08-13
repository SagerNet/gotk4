// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_time_zone_get_type()), F: marshalTimeZone},
	})
}

// TimeType disambiguates a given time in two ways.
//
// First, specifies if the given time is in universal or local time.
//
// Second, if the time is in local time, specifies if it is local standard time
// or local daylight time. This is important for the case where the same local
// time occurs twice (during daylight savings time transitions, for example).
type TimeType int

const (
	// TimeTypeStandard: time is in local standard time
	TimeTypeStandard TimeType = iota
	// TimeTypeDaylight: time is in local daylight time
	TimeTypeDaylight
	// TimeTypeUniversal: time is in UTC
	TimeTypeUniversal
)

// String returns the name in string for TimeType.
func (t TimeType) String() string {
	switch t {
	case TimeTypeStandard:
		return "Standard"
	case TimeTypeDaylight:
		return "Daylight"
	case TimeTypeUniversal:
		return "Universal"
	default:
		return fmt.Sprintf("TimeType(%d)", t)
	}
}

// TimeZone is an opaque structure whose members cannot be accessed directly.
type TimeZone struct {
	nocopy gextras.NoCopy
	native *C.GTimeZone
}

func marshalTimeZone(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &TimeZone{native: (*C.GTimeZone)(unsafe.Pointer(b))}, nil
}

// NewTimeZone constructs a struct TimeZone.
func NewTimeZone(identifier string) *TimeZone {
	var _arg1 *C.gchar     // out
	var _cret *C.GTimeZone // in

	if identifier != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(identifier)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_time_zone_new(_arg1)
	runtime.KeepAlive(identifier)

	var _timeZone *TimeZone // out

	_timeZone = (*TimeZone)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_timeZone, func(v *TimeZone) {
		C.g_time_zone_unref((*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _timeZone
}

// NewTimeZoneIdentifier constructs a struct TimeZone.
func NewTimeZoneIdentifier(identifier string) *TimeZone {
	var _arg1 *C.gchar     // out
	var _cret *C.GTimeZone // in

	if identifier != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(identifier)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_time_zone_new_identifier(_arg1)
	runtime.KeepAlive(identifier)

	var _timeZone *TimeZone // out

	if _cret != nil {
		_timeZone = (*TimeZone)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_timeZone, func(v *TimeZone) {
			C.g_time_zone_unref((*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _timeZone
}

// NewTimeZoneLocal constructs a struct TimeZone.
func NewTimeZoneLocal() *TimeZone {
	var _cret *C.GTimeZone // in

	_cret = C.g_time_zone_new_local()

	var _timeZone *TimeZone // out

	_timeZone = (*TimeZone)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_timeZone, func(v *TimeZone) {
		C.g_time_zone_unref((*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _timeZone
}

// NewTimeZoneOffset constructs a struct TimeZone.
func NewTimeZoneOffset(seconds int32) *TimeZone {
	var _arg1 C.gint32     // out
	var _cret *C.GTimeZone // in

	_arg1 = C.gint32(seconds)

	_cret = C.g_time_zone_new_offset(_arg1)
	runtime.KeepAlive(seconds)

	var _timeZone *TimeZone // out

	_timeZone = (*TimeZone)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_timeZone, func(v *TimeZone) {
		C.g_time_zone_unref((*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _timeZone
}

// NewTimeZoneUtc constructs a struct TimeZone.
func NewTimeZoneUtc() *TimeZone {
	var _cret *C.GTimeZone // in

	_cret = C.g_time_zone_new_utc()

	var _timeZone *TimeZone // out

	_timeZone = (*TimeZone)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_timeZone, func(v *TimeZone) {
		C.g_time_zone_unref((*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _timeZone
}

// AdjustTime finds an interval within tz that corresponds to the given time_,
// possibly adjusting time_ if required to fit into an interval. The meaning of
// time_ depends on type.
//
// This function is similar to g_time_zone_find_interval(), with the difference
// that it always succeeds (by making the adjustments described below).
//
// In any of the cases where g_time_zone_find_interval() succeeds then this
// function returns the same value, without modifying time_.
//
// This function may, however, modify time_ in order to deal with non-existent
// times. If the non-existent local time_ of 02:30 were requested on March 14th
// 2010 in Toronto then this function would adjust time_ to be 03:00 and return
// the interval containing the adjusted time.
func (tz *TimeZone) AdjustTime(typ TimeType, time_ *int64) int {
	var _arg0 *C.GTimeZone // out
	var _arg1 C.GTimeType  // out
	var _arg2 *C.gint64    // out
	var _cret C.gint       // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))
	_arg1 = C.GTimeType(typ)
	_arg2 = (*C.gint64)(unsafe.Pointer(time_))

	_cret = C.g_time_zone_adjust_time(_arg0, _arg1, _arg2)
	runtime.KeepAlive(tz)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(time_)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// FindInterval finds an interval within tz that corresponds to the given time_.
// The meaning of time_ depends on type.
//
// If type is G_TIME_TYPE_UNIVERSAL then this function will always succeed
// (since universal time is monotonic and continuous).
//
// Otherwise time_ is treated as local time. The distinction between
// G_TIME_TYPE_STANDARD and G_TIME_TYPE_DAYLIGHT is ignored except in the case
// that the given time_ is ambiguous. In Toronto, for example, 01:30 on November
// 7th 2010 occurred twice (once inside of daylight savings time and the next,
// an hour later, outside of daylight savings time). In this case, the different
// value of type would result in a different interval being returned.
//
// It is still possible for this function to fail. In Toronto, for example,
// 02:00 on March 14th 2010 does not exist (due to the leap forward to begin
// daylight savings time). -1 is returned in that case.
func (tz *TimeZone) FindInterval(typ TimeType, time_ int64) int {
	var _arg0 *C.GTimeZone // out
	var _arg1 C.GTimeType  // out
	var _arg2 C.gint64     // out
	var _cret C.gint       // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))
	_arg1 = C.GTimeType(typ)
	_arg2 = C.gint64(time_)

	_cret = C.g_time_zone_find_interval(_arg0, _arg1, _arg2)
	runtime.KeepAlive(tz)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(time_)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Abbreviation determines the time zone abbreviation to be used during a
// particular interval of time in the time zone tz.
//
// For example, in Toronto this is currently "EST" during the winter months and
// "EDT" during the summer months when daylight savings time is in effect.
func (tz *TimeZone) Abbreviation(interval int) string {
	var _arg0 *C.GTimeZone // out
	var _arg1 C.gint       // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))
	_arg1 = C.gint(interval)

	_cret = C.g_time_zone_get_abbreviation(_arg0, _arg1)
	runtime.KeepAlive(tz)
	runtime.KeepAlive(interval)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Identifier: get the identifier of this Zone, as passed to g_time_zone_new().
// If the identifier passed at construction time was not recognised, UTC will be
// returned. If it was NULL, the identifier of the local timezone at
// construction time will be returned.
//
// The identifier will be returned in the same format as provided at
// construction time: if provided as a time offset, that will be returned by
// this function.
func (tz *TimeZone) Identifier() string {
	var _arg0 *C.GTimeZone // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))

	_cret = C.g_time_zone_get_identifier(_arg0)
	runtime.KeepAlive(tz)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Offset determines the offset to UTC in effect during a particular interval of
// time in the time zone tz.
//
// The offset is the number of seconds that you add to UTC time to arrive at
// local time for tz (ie: negative numbers for time zones west of GMT, positive
// numbers for east).
func (tz *TimeZone) Offset(interval int) int32 {
	var _arg0 *C.GTimeZone // out
	var _arg1 C.gint       // out
	var _cret C.gint32     // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))
	_arg1 = C.gint(interval)

	_cret = C.g_time_zone_get_offset(_arg0, _arg1)
	runtime.KeepAlive(tz)
	runtime.KeepAlive(interval)

	var _gint32 int32 // out

	_gint32 = int32(_cret)

	return _gint32
}

// IsDst determines if daylight savings time is in effect during a particular
// interval of time in the time zone tz.
func (tz *TimeZone) IsDst(interval int) bool {
	var _arg0 *C.GTimeZone // out
	var _arg1 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GTimeZone)(gextras.StructNative(unsafe.Pointer(tz)))
	_arg1 = C.gint(interval)

	_cret = C.g_time_zone_is_dst(_arg0, _arg1)
	runtime.KeepAlive(tz)
	runtime.KeepAlive(interval)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
