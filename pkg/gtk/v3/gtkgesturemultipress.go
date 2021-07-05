// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_multi_press_get_type()), F: marshalGestureMultiPress},
	})
}

// GestureMultiPress is a Gesture implementation able to recognize multiple
// clicks on a nearby zone, which can be listened for through the
// GestureMultiPress::pressed signal. Whenever time or distance between clicks
// exceed the GTK+ defaults, GestureMultiPress::stopped is emitted, and the
// click counter is reset.
//
// Callers may also restrict the area that is considered valid for a >1
// touch/button press through gtk_gesture_multi_press_set_area(), so any click
// happening outside that area is considered to be a first click of its own.
type GestureMultiPress interface {
	GestureSingle

	// Area: if an area was set through gtk_gesture_multi_press_set_area(), this
	// function will return true and fill in @rect with the press area. See
	// gtk_gesture_multi_press_set_area() for more details on what the press
	// area represents.
	Area() (gdk.Rectangle, bool)
	// SetArea: if @rect is non-nil, the press area will be checked to be
	// confined within the rectangle, otherwise the button count will be reset
	// so the press is seen as being the first one. If @rect is nil, the area
	// will be reset to an unrestricted state.
	//
	// Note: The rectangle is only used to determine whether any non-first click
	// falls within the expected area. This is not akin to an input shape.
	SetArea(rect *gdk.Rectangle)
}

// gestureMultiPress implements the GestureMultiPress class.
type gestureMultiPress struct {
	GestureSingle
}

// WrapGestureMultiPress wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureMultiPress(obj *externglib.Object) GestureMultiPress {
	return gestureMultiPress{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureMultiPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureMultiPress(obj), nil
}

// NewGestureMultiPress returns a newly created Gesture that recognizes single
// and multiple presses.
func NewGestureMultiPress(widget Widget) GestureMultiPress {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_multi_press_new(_arg1)

	var _gestureMultiPress GestureMultiPress // out

	_gestureMultiPress = WrapGestureMultiPress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureMultiPress
}

func (g gestureMultiPress) Area() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkGestureMultiPress // out
	var _arg1 C.GdkRectangle          // in
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_multi_press_get_area(_arg0, &_arg1)

	var _rect gdk.Rectangle // out
	var _ok bool            // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rect = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

func (g gestureMultiPress) SetArea(rect *gdk.Rectangle) {
	var _arg0 *C.GtkGestureMultiPress // out
	var _arg1 *C.GdkRectangle         // out

	_arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_gesture_multi_press_set_area(_arg0, _arg1)
}
