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
		{T: externglib.Type(C.gtk_gesture_long_press_get_type()), F: marshalGestureLongPresser},
	})
}

// GestureLongPress is a Gesture implementation able to recognize long presses,
// triggering the GestureLongPress::pressed after the timeout is exceeded.
//
// If the touchpoint is lifted before the timeout passes, or if it drifts too
// far of the initial press point, the GestureLongPress::cancelled signal will
// be emitted.
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

// NewGestureLongPress returns a newly created Gesture that recognizes long
// presses.
func NewGestureLongPress(widget Widgetter) *GestureLongPress {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_long_press_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureLongPress *GestureLongPress // out

	_gestureLongPress = WrapGestureLongPress(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureLongPress
}

func (*GestureLongPress) privateGestureLongPress() {}
