// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_swipe_get_type()), F: marshalGestureSwiper},
	})
}

// GestureSwipe is a Gesture implementation able to recognize swipes, after a
// press/move/.../move/release sequence happens, the GestureSwipe::swipe signal
// will be emitted, providing the velocity and directionality of the sequence at
// the time it was lifted.
//
// If the velocity is desired in intermediate points,
// gtk_gesture_swipe_get_velocity() can be called on eg. a Gesture::update
// handler.
//
// All velocities are reported in pixels/sec units.
type GestureSwipe struct {
	GestureSingle
}

func WrapGestureSwipe(obj *externglib.Object) *GestureSwipe {
	return &GestureSwipe{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureSwiper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureSwipe(obj), nil
}

// NewGestureSwipe returns a newly created Gesture that recognizes swipes.
func NewGestureSwipe(widget Widgetter) *GestureSwipe {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_swipe_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureSwipe *GestureSwipe // out

	_gestureSwipe = WrapGestureSwipe(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureSwipe
}

// Velocity: if the gesture is recognized, this function returns TRUE and fill
// in velocity_x and velocity_y with the recorded velocity, as per the last
// event(s) processed.
func (gesture *GestureSwipe) Velocity() (velocityX float64, velocityY float64, ok bool) {
	var _arg0 *C.GtkGestureSwipe // out
	var _arg1 C.gdouble          // in
	var _arg2 C.gdouble          // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkGestureSwipe)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_swipe_get_velocity(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _velocityX float64 // out
	var _velocityY float64 // out
	var _ok bool           // out

	_velocityX = float64(_arg1)
	_velocityY = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _velocityX, _velocityY, _ok
}
