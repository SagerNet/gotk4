// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
	Area() (rect gdk.Rectangle, ok bool)
	// SetArea: if @rect is non-nil, the press area will be checked to be
	// confined within the rectangle, otherwise the button count will be reset
	// so the press is seen as being the first one. If @rect is nil, the area
	// will be reset to an unrestricted state.
	//
	// Note: The rectangle is only used to determine whether any non-first click
	// falls within the expected area. This is not akin to an input shape.
	SetArea(rect *gdk.Rectangle)
}

// gestureMultiPress implements the GestureMultiPress interface.
type gestureMultiPress struct {
	GestureSingle
}

var _ GestureMultiPress = (*gestureMultiPress)(nil)

// WrapGestureMultiPress wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureMultiPress(obj *externglib.Object) GestureMultiPress {
	return GestureMultiPress{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureMultiPress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureMultiPress(obj), nil
}

// NewGestureMultiPress constructs a class GestureMultiPress.
func NewGestureMultiPress(widget Widget) GestureMultiPress {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var cret C.GtkGestureMultiPress
	var goret1 GestureMultiPress

	cret = C.gtk_gesture_multi_press_new(widget)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureMultiPress)

	return goret1
}

// Area: if an area was set through gtk_gesture_multi_press_set_area(), this
// function will return true and fill in @rect with the press area. See
// gtk_gesture_multi_press_set_area() for more details on what the press
// area represents.
func (g gestureMultiPress) Area() (rect gdk.Rectangle, ok bool) {
	var arg0 *C.GtkGestureMultiPress

	arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))

	var arg1 *C.GdkRectangle
	var ret1 *gdk.Rectangle
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_gesture_multi_press_get_area(arg0, &arg1)

	ret1 = gdk.WrapRectangle(unsafe.Pointer(arg1))
	goret2 = C.bool(cret) != C.false

	return ret1, goret2
}

// SetArea: if @rect is non-nil, the press area will be checked to be
// confined within the rectangle, otherwise the button count will be reset
// so the press is seen as being the first one. If @rect is nil, the area
// will be reset to an unrestricted state.
//
// Note: The rectangle is only used to determine whether any non-first click
// falls within the expected area. This is not akin to an input shape.
func (g gestureMultiPress) SetArea(rect *gdk.Rectangle) {
	var arg0 *C.GtkGestureMultiPress
	var arg1 *C.GdkRectangle

	arg0 = (*C.GtkGestureMultiPress)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_gesture_multi_press_set_area(arg0, rect)
}
