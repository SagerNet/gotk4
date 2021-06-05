// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_frame_clock_get_type()), F: marshalFrameClock},
	})
}

// FrameClock: a FrameClock tells the application when to update and repaint a
// window. This may be synced to the vertical refresh rate of the monitor, for
// example. Even when the frame clock uses a simple timer rather than a
// hardware-based vertical sync, the frame clock helps because it ensures
// everything paints at the same time (reducing the total number of frames). The
// frame clock can also automatically stop painting when it knows the frames
// will not be visible, or scale back animation framerates.
//
// FrameClock is designed to be compatible with an OpenGL-based implementation
// or with mozRequestAnimationFrame in Firefox, for example.
//
// A frame clock is idle until someone requests a frame with
// gdk_frame_clock_request_phase(). At some later point that makes sense for the
// synchronization being implemented, the clock will process a frame and emit
// signals for each phase that has been requested. (See the signals of the
// FrameClock class for documentation of the phases.
// GDK_FRAME_CLOCK_PHASE_UPDATE and the FrameClock::update signal are most
// interesting for application writers, and are used to update the animations,
// using the frame time given by gdk_frame_clock_get_frame_time().
//
// The frame time is reported in microseconds and generally in the same
// timescale as g_get_monotonic_time(), however, it is not the same as
// g_get_monotonic_time(). The frame time does not advance during the time a
// frame is being painted, and outside of a frame, an attempt is made so that
// all calls to gdk_frame_clock_get_frame_time() that are called at a “similar”
// time get the same value. This means that if different animations are timed by
// looking at the difference in time between an initial value from
// gdk_frame_clock_get_frame_time() and the value inside the FrameClock::update
// signal of the clock, they will stay exactly synchronized.
type FrameClock interface {
	gextras.Objector

	// BeginUpdating starts updates for an animation. Until a matching call to
	// gdk_frame_clock_end_updating() is made, the frame clock will continually
	// request a new frame with the GDK_FRAME_CLOCK_PHASE_UPDATE phase. This
	// function may be called multiple times and frames will be requested until
	// gdk_frame_clock_end_updating() is called the same number of times.
	BeginUpdating()
	// EndUpdating stops updates for an animation. See the documentation for
	// gdk_frame_clock_begin_updating().
	EndUpdating()
	// CurrentTimings gets the frame timings for the current frame.
	CurrentTimings() *FrameTimings
	// FrameCounter: a FrameClock maintains a 64-bit counter that increments for
	// each frame drawn.
	FrameCounter() int64
	// FrameTime gets the time that should currently be used for animations.
	// Inside the processing of a frame, it’s the time used to compute the
	// animation position of everything in a frame. Outside of a frame, it's the
	// time of the conceptual “previous frame,” which may be either the actual
	// previous frame time, or if that’s too old, an updated time.
	FrameTime() int64
	// HistoryStart: FrameClock internally keeps a history of FrameTimings
	// objects for recent frames that can be retrieved with
	// gdk_frame_clock_get_timings(). The set of stored frames is the set from
	// the counter values given by gdk_frame_clock_get_history_start() and
	// gdk_frame_clock_get_frame_counter(), inclusive.
	HistoryStart() int64
	// RefreshInfo: using the frame history stored in the frame clock, finds the
	// last known presentation time and refresh interval, and assuming that
	// presentation times are separated by the refresh interval, predicts a
	// presentation time that is a multiple of the refresh interval after the
	// last presentation time, and later than @base_time.
	RefreshInfo(baseTime int64) (refreshIntervalReturn int64, presentationTimeReturn int64)
	// Timings retrieves a FrameTimings object holding timing information for
	// the current frame or a recent frame. The FrameTimings object may not yet
	// be complete: see gdk_frame_timings_get_complete().
	Timings(frameCounter int64) *FrameTimings
	// RequestPhase asks the frame clock to run a particular phase. The signal
	// corresponding the requested phase will be emitted the next time the frame
	// clock processes. Multiple calls to gdk_frame_clock_request_phase() will
	// be combined together and only one frame processed. If you are displaying
	// animated content and want to continually request the
	// GDK_FRAME_CLOCK_PHASE_UPDATE phase for a period of time, you should use
	// gdk_frame_clock_begin_updating() instead, since this allows GTK+ to
	// adjust system parameters to get maximally smooth animations.
	RequestPhase(phase FrameClockPhase)
}

// frameClock implements the FrameClock interface.
type frameClock struct {
	gextras.Objector
}

var _ FrameClock = (*frameClock)(nil)

// WrapFrameClock wraps a GObject to the right type. It is
// primarily used internally.
func WrapFrameClock(obj *externglib.Object) FrameClock {
	return FrameClock{
		Objector: obj,
	}
}

func marshalFrameClock(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFrameClock(obj), nil
}

// BeginUpdating starts updates for an animation. Until a matching call to
// gdk_frame_clock_end_updating() is made, the frame clock will continually
// request a new frame with the GDK_FRAME_CLOCK_PHASE_UPDATE phase. This
// function may be called multiple times and frames will be requested until
// gdk_frame_clock_end_updating() is called the same number of times.
func (f frameClock) BeginUpdating() {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	C.gdk_frame_clock_begin_updating(arg0)
}

// EndUpdating stops updates for an animation. See the documentation for
// gdk_frame_clock_begin_updating().
func (f frameClock) EndUpdating() {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	C.gdk_frame_clock_end_updating(arg0)
}

// CurrentTimings gets the frame timings for the current frame.
func (f frameClock) CurrentTimings() *FrameTimings {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	var cret *C.GdkFrameTimings
	var goret1 *FrameTimings

	cret = C.gdk_frame_clock_get_current_timings(arg0)

	goret1 = WrapFrameTimings(unsafe.Pointer(cret))

	return goret1
}

// FrameCounter: a FrameClock maintains a 64-bit counter that increments for
// each frame drawn.
func (f frameClock) FrameCounter() int64 {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	var cret C.gint64
	var goret1 int64

	cret = C.gdk_frame_clock_get_frame_counter(arg0)

	goret1 = C.gint64(cret)

	return goret1
}

// FrameTime gets the time that should currently be used for animations.
// Inside the processing of a frame, it’s the time used to compute the
// animation position of everything in a frame. Outside of a frame, it's the
// time of the conceptual “previous frame,” which may be either the actual
// previous frame time, or if that’s too old, an updated time.
func (f frameClock) FrameTime() int64 {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	var cret C.gint64
	var goret1 int64

	cret = C.gdk_frame_clock_get_frame_time(arg0)

	goret1 = C.gint64(cret)

	return goret1
}

// HistoryStart: FrameClock internally keeps a history of FrameTimings
// objects for recent frames that can be retrieved with
// gdk_frame_clock_get_timings(). The set of stored frames is the set from
// the counter values given by gdk_frame_clock_get_history_start() and
// gdk_frame_clock_get_frame_counter(), inclusive.
func (f frameClock) HistoryStart() int64 {
	var arg0 *C.GdkFrameClock

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))

	var cret C.gint64
	var goret1 int64

	cret = C.gdk_frame_clock_get_history_start(arg0)

	goret1 = C.gint64(cret)

	return goret1
}

// RefreshInfo: using the frame history stored in the frame clock, finds the
// last known presentation time and refresh interval, and assuming that
// presentation times are separated by the refresh interval, predicts a
// presentation time that is a multiple of the refresh interval after the
// last presentation time, and later than @base_time.
func (f frameClock) RefreshInfo(baseTime int64) (refreshIntervalReturn int64, presentationTimeReturn int64) {
	var arg0 *C.GdkFrameClock
	var arg1 C.gint64

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))
	arg1 = C.gint64(baseTime)

	var arg2 *C.gint64
	var ret2 int64
	var arg3 *C.gint64
	var ret3 int64

	C.gdk_frame_clock_get_refresh_info(arg0, baseTime, &arg2, &arg3)

	ret2 = *C.gint64(arg2)
	ret3 = *C.gint64(arg3)

	return ret2, ret3
}

// Timings retrieves a FrameTimings object holding timing information for
// the current frame or a recent frame. The FrameTimings object may not yet
// be complete: see gdk_frame_timings_get_complete().
func (f frameClock) Timings(frameCounter int64) *FrameTimings {
	var arg0 *C.GdkFrameClock
	var arg1 C.gint64

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))
	arg1 = C.gint64(frameCounter)

	var cret *C.GdkFrameTimings
	var goret1 *FrameTimings

	cret = C.gdk_frame_clock_get_timings(arg0, frameCounter)

	goret1 = WrapFrameTimings(unsafe.Pointer(cret))

	return goret1
}

// RequestPhase asks the frame clock to run a particular phase. The signal
// corresponding the requested phase will be emitted the next time the frame
// clock processes. Multiple calls to gdk_frame_clock_request_phase() will
// be combined together and only one frame processed. If you are displaying
// animated content and want to continually request the
// GDK_FRAME_CLOCK_PHASE_UPDATE phase for a period of time, you should use
// gdk_frame_clock_begin_updating() instead, since this allows GTK+ to
// adjust system parameters to get maximally smooth animations.
func (f frameClock) RequestPhase(phase FrameClockPhase) {
	var arg0 *C.GdkFrameClock
	var arg1 C.GdkFrameClockPhase

	arg0 = (*C.GdkFrameClock)(unsafe.Pointer(f.Native()))
	arg1 = (C.GdkFrameClockPhase)(phase)

	C.gdk_frame_clock_request_phase(arg0, phase)
}

type FrameClockPrivate struct {
	native C.GdkFrameClockPrivate
}

// WrapFrameClockPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFrameClockPrivate(ptr unsafe.Pointer) *FrameClockPrivate {
	if ptr == nil {
		return nil
	}

	return (*FrameClockPrivate)(ptr)
}

func marshalFrameClockPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFrameClockPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FrameClockPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}
