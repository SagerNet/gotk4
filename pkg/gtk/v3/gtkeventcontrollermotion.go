// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_event_controller_motion_get_type()), F: marshalEventControllerMotion},
	})
}

// EventControllerMotion is an event controller meant for situations where you
// need to track the position of the pointer.
//
// This object was added in 3.24.
type EventControllerMotion interface {
	EventController
}

// eventControllerMotion implements the EventControllerMotion class.
type eventControllerMotion struct {
	EventController
}

// WrapEventControllerMotion wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerMotion(obj *externglib.Object) EventControllerMotion {
	return eventControllerMotion{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerMotion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerMotion(obj), nil
}

// NewEventControllerMotion:
func NewEventControllerMotion(widget Widget) EventControllerMotion {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_motion_new(_arg1)

	var _eventControllerMotion EventControllerMotion // out

	_eventControllerMotion = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EventControllerMotion)

	return _eventControllerMotion
}
