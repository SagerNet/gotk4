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
		{T: externglib.Type(C.gtk_event_controller_get_type()), F: marshalEventControllerer},
	})
}

// EventController is a base, low-level implementation for event controllers.
// Those react to a series of Events, and possibly trigger actions as a
// consequence of those.
type EventController struct {
	*externglib.Object
}

// EventControllerer describes EventController's abstract methods.
type EventControllerer interface {
	externglib.Objector

	// PropagationPhase gets the propagation phase at which controller handles
	// events.
	PropagationPhase() PropagationPhase
	// Widget returns the Widget this controller relates to.
	Widget() Widgetter
	// Reset resets the controller to a clean state.
	Reset()
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	SetPropagationPhase(phase PropagationPhase)
}

var _ EventControllerer = (*EventController)(nil)

func WrapEventController(obj *externglib.Object) *EventController {
	return &EventController{
		Object: obj,
	}
}

func marshalEventControllerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventController(obj), nil
}

// PropagationPhase gets the propagation phase at which controller handles
// events.
func (controller *EventController) PropagationPhase() PropagationPhase {
	var _arg0 *C.GtkEventController // out
	var _cret C.GtkPropagationPhase // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_propagation_phase(_arg0)
	runtime.KeepAlive(controller)

	var _propagationPhase PropagationPhase // out

	_propagationPhase = PropagationPhase(_cret)

	return _propagationPhase
}

// Widget returns the Widget this controller relates to.
func (controller *EventController) Widget() Widgetter {
	var _arg0 *C.GtkEventController // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_get_widget(_arg0)
	runtime.KeepAlive(controller)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Reset resets the controller to a clean state. Every interaction the
// controller did through EventController::handle-event will be dropped at this
// point.
func (controller *EventController) Reset() {
	var _arg0 *C.GtkEventController // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))

	C.gtk_event_controller_reset(_arg0)
	runtime.KeepAlive(controller)
}

// SetPropagationPhase sets the propagation phase at which a controller handles
// events.
//
// If phase is GTK_PHASE_NONE, no automatic event handling will be performed,
// but other additional gesture maintenance will. In that phase, the events can
// be managed by calling gtk_event_controller_handle_event().
func (controller *EventController) SetPropagationPhase(phase PropagationPhase) {
	var _arg0 *C.GtkEventController // out
	var _arg1 C.GtkPropagationPhase // out

	_arg0 = (*C.GtkEventController)(unsafe.Pointer(controller.Native()))
	_arg1 = C.GtkPropagationPhase(phase)

	C.gtk_event_controller_set_propagation_phase(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(phase)
}
