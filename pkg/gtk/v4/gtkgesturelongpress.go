// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_long_press_get_type()), F: marshalGestureLongPresser},
	})
}

// GestureLongPress: GtkGestureLongPress is a GtkGesture for long presses.
//
// This gesture is also known as “Press and Hold”.
//
// When the timeout is exceeded, the gesture is triggering the
// gtk.GestureLongPress::pressed signal.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the gtk.GestureLongPress::cancelled signal
// will be emitted.
//
// How long the timeout is before the ::pressed signal gets emitted is
// determined by the gtk.Settings:gtk-long-press-time setting. It can be
// modified by the gtk.GestureLongPress:delay-factor property.
type GestureLongPress struct {
	GestureSingle
}

func WrapGestureLongPress(obj *externglib.Object) *GestureLongPress {
	return &GestureLongPress{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureLongPresser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureLongPress(obj), nil
}

// NewGestureLongPress returns a newly created GtkGesture that recognizes long
// presses.
func NewGestureLongPress() *GestureLongPress {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_long_press_new()

	var _gestureLongPress *GestureLongPress // out

	_gestureLongPress = WrapGestureLongPress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureLongPress
}

// DelayFactor returns the delay factor.
func (gesture *GestureLongPress) DelayFactor() float64 {
	var _arg0 *C.GtkGestureLongPress // out
	var _cret C.double               // in

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_long_press_get_delay_factor(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDelayFactor applies the given delay factor.
//
// The default long press time will be multiplied by this value. Valid values
// are in the range [0.5..2.0].
func (gesture *GestureLongPress) SetDelayFactor(delayFactor float64) {
	var _arg0 *C.GtkGestureLongPress // out
	var _arg1 C.double               // out

	_arg0 = (*C.GtkGestureLongPress)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.double(delayFactor)

	C.gtk_gesture_long_press_set_delay_factor(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(delayFactor)
}
