// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_zoom_get_type()), F: marshalGestureZoom},
	})
}

// GestureZoom is a Gesture implementation able to recognize pinch/zoom
// gestures, whenever the distance between both tracked sequences changes, the
// GestureZoom::scale-changed signal is emitted to report the scale factor.
type GestureZoom interface {
	Gesture

	// ScaleDelta: if @gesture is active, this function returns the zooming
	// difference since the gesture was recognized (hence the starting point is
	// considered 1:1). If @gesture is not active, 1 is returned.
	ScaleDelta() float64
}

// gestureZoom implements the GestureZoom interface.
type gestureZoom struct {
	Gesture
}

var _ GestureZoom = (*gestureZoom)(nil)

// WrapGestureZoom wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureZoom(obj *externglib.Object) GestureZoom {
	return GestureZoom{
		Gesture: WrapGesture(obj),
	}
}

func marshalGestureZoom(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureZoom(obj), nil
}

// NewGestureZoom constructs a class GestureZoom.
func NewGestureZoom() GestureZoom {
	var cret C.GtkGestureZoom
	var ret1 GestureZoom

	cret = C.gtk_gesture_zoom_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureZoom)

	return ret1
}

// ScaleDelta: if @gesture is active, this function returns the zooming
// difference since the gesture was recognized (hence the starting point is
// considered 1:1). If @gesture is not active, 1 is returned.
func (g gestureZoom) ScaleDelta() float64 {
	var arg0 *C.GtkGestureZoom

	arg0 = (*C.GtkGestureZoom)(unsafe.Pointer(g.Native()))

	var cret C.double
	var ret1 float64

	cret = C.gtk_gesture_zoom_get_scale_delta(arg0)

	ret1 = C.double(cret)

	return ret1
}