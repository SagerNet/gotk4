// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_legacy_get_type()), F: marshalyier},
	})
}

// yier describes EventControllerLegacy's methods.
type yier interface {
	gextras.Objector

	privateEventControllerLegacy()
}

// EventControllerLegacy: `GtkEventControllerLegacy` is an event controller that
// provides raw access to the event stream.
//
// It should only be used as a last resort if none of the other event
// controllers or gestures do the job.
type EventControllerLegacy struct {
	EventController
}

var _ yier = (*EventControllerLegacy)(nil)

func wrapyier(obj *externglib.Object) yier {
	return &EventControllerLegacy{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalyier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapyier(obj), nil
}

// NewEventControllerLegacy creates a new legacy event controller.
func NewEventControllerLegacy() *EventControllerLegacy {
	var _cret *C.GtkEventController // in

	_cret = C.gtk_event_controller_legacy_new()

	var _eventControllerLegacy *EventControllerLegacy // out

	_eventControllerLegacy = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*EventControllerLegacy)

	return _eventControllerLegacy
}

func (*EventControllerLegacy) privateEventControllerLegacy() {}
