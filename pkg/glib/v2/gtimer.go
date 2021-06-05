// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// TimeValFromISO8601 converts a string containing an ISO 8601 encoded date and
// time to a Val and puts it into @time_.
//
// @iso_date must include year, month, day, hours, minutes, and seconds. It can
// optionally include fractions of a second and a time zone indicator. (In the
// absence of any time zone indication, the timestamp is assumed to be in local
// time.)
//
// Any leading or trailing space in @iso_date is ignored.
//
// This function was deprecated, along with Val itself, in GLib 2.62. Equivalent
// functionality is available using code like:
//
//    GDateTime *dt = g_date_time_new_from_iso8601 (iso8601_string, NULL);
//    gint64 time_val = g_date_time_to_unix (dt);
//    g_date_time_unref (dt);
func TimeValFromISO8601(isoDate string) (time_ TimeVal, ok bool) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(isoDate))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.GTimeVal
	var ret2 *TimeVal
	var cret C.gboolean
	var goret2 bool

	cret = C.g_time_val_from_iso8601(isoDate, &arg2)

	ret2 = WrapTimeVal(unsafe.Pointer(arg2))
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

// Usleep pauses the current thread for the given number of microseconds.
//
// There are 1 million microseconds per second (represented by the USEC_PER_SEC
// macro). g_usleep() may have limited precision, depending on hardware and
// operating system; don't rely on the exact length of the sleep.
func Usleep(microseconds uint32) {
	var arg1 C.gulong

	arg1 = C.gulong(microseconds)

	C.g_usleep(microseconds)
}

// Timer: opaque datatype that records a start time.
type Timer struct {
	native C.GTimer
}

// WrapTimer wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTimer(ptr unsafe.Pointer) *Timer {
	if ptr == nil {
		return nil
	}

	return (*Timer)(ptr)
}

func marshalTimer(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTimer(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *Timer) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Continue resumes a timer that has previously been stopped with
// g_timer_stop(). g_timer_stop() must be called before using this function.
func (t *Timer) Continue() {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	C.g_timer_continue(arg0)
}

// Destroy destroys a timer, freeing associated resources.
func (t *Timer) Destroy() {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	C.g_timer_destroy(arg0)
}

// Elapsed: if @timer has been started but not stopped, obtains the time since
// the timer was started. If @timer has been stopped, obtains the elapsed time
// between the time it was started and the time it was stopped. The return value
// is the number of seconds elapsed, including any fractional part. The
// @microseconds out parameter is essentially useless.
func (t *Timer) Elapsed(microseconds uint32) float64 {
	var arg0 *C.GTimer
	var arg1 *C.gulong

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))
	arg1 = *C.gulong(microseconds)

	var cret C.gdouble
	var goret1 float64

	cret = C.g_timer_elapsed(arg0, microseconds)

	goret1 = C.gdouble(cret)

	return goret1
}

// IsActive exposes whether the timer is currently active.
func (t *Timer) IsActive() bool {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_timer_is_active(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Reset: this function is useless; it's fine to call g_timer_start() on an
// already-started timer to reset the start time, so g_timer_reset() serves no
// purpose.
func (t *Timer) Reset() {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	C.g_timer_reset(arg0)
}

// Start marks a start time, so that future calls to g_timer_elapsed() will
// report the time since g_timer_start() was called. g_timer_new() automatically
// marks the start time, so no need to call g_timer_start() immediately after
// creating the timer.
func (t *Timer) Start() {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	C.g_timer_start(arg0)
}

// Stop marks an end time, so calls to g_timer_elapsed() will return the
// difference between this end time and the start time.
func (t *Timer) Stop() {
	var arg0 *C.GTimer

	arg0 = (*C.GTimer)(unsafe.Pointer(t.Native()))

	C.g_timer_stop(arg0)
}
