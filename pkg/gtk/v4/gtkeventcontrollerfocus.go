// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_focus_get_type()), F: marshalEventControllerFocus},
	})
}

// EventControllerFocus is an event controller meant for situations where you
// need to know where the focus is.
type EventControllerFocus interface {
	EventController

	// ContainsFocus returns the value of the
	// GtkEventControllerFocus:contains-focus property.
	ContainsFocus() bool
	// IsFocus returns the value of the GtkEventControllerFocus:is-focus
	// property.
	IsFocus() bool
}

// eventControllerFocus implements the EventControllerFocus interface.
type eventControllerFocus struct {
	EventController
}

var _ EventControllerFocus = (*eventControllerFocus)(nil)

// WrapEventControllerFocus wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerFocus(obj *externglib.Object) EventControllerFocus {
	return EventControllerFocus{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerFocus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerFocus(obj), nil
}

// NewEventControllerFocus constructs a class EventControllerFocus.
func NewEventControllerFocus() EventControllerFocus {
	var cret C.GtkEventControllerFocus
	var goret1 EventControllerFocus

	cret = C.gtk_event_controller_focus_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(EventControllerFocus)

	return goret1
}

// ContainsFocus returns the value of the
// GtkEventControllerFocus:contains-focus property.
func (s eventControllerFocus) ContainsFocus() bool {
	var arg0 *C.GtkEventControllerFocus

	arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_event_controller_focus_contains_focus(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// IsFocus returns the value of the GtkEventControllerFocus:is-focus
// property.
func (s eventControllerFocus) IsFocus() bool {
	var arg0 *C.GtkEventControllerFocus

	arg0 = (*C.GtkEventControllerFocus)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_event_controller_focus_is_focus(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}
