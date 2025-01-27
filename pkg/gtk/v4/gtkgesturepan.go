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
		{T: externglib.Type(C.gtk_gesture_pan_get_type()), F: marshalGesturePanner},
	})
}

// GesturePan: GtkGesturePan is a GtkGesture for pan gestures.
//
// These are drags that are locked to happen along one axis. The axis that a
// GtkGesturePan handles is defined at construct time, and can be changed
// through gtk.GesturePan.SetOrientation().
//
// When the gesture starts to be recognized, GtkGesturePan will attempt to
// determine as early as possible whether the sequence is moving in the expected
// direction, and denying the sequence if this does not happen.
//
// Once a panning gesture along the expected axis is recognized, the
// gtk.GesturePan::pan signal will be emitted as input events are received,
// containing the offset in the given axis.
type GesturePan struct {
	GestureDrag
}

func WrapGesturePan(obj *externglib.Object) *GesturePan {
	return &GesturePan{
		GestureDrag: GestureDrag{
			GestureSingle: GestureSingle{
				Gesture: Gesture{
					EventController: EventController{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalGesturePanner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGesturePan(obj), nil
}

// NewGesturePan returns a newly created GtkGesture that recognizes pan
// gestures.
func NewGesturePan(orientation Orientation) *GesturePan {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkGesture    // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_gesture_pan_new(_arg1)
	runtime.KeepAlive(orientation)

	var _gesturePan *GesturePan // out

	_gesturePan = WrapGesturePan(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gesturePan
}

// Orientation returns the orientation of the pan gestures that this gesture
// expects.
func (gesture *GesturePan) Orientation() Orientation {
	var _arg0 *C.GtkGesturePan // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_pan_get_orientation(_arg0)
	runtime.KeepAlive(gesture)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation to be expected on pan gestures.
func (gesture *GesturePan) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkGesturePan // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkGesturePan)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_gesture_pan_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(orientation)
}
