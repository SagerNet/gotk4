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
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKeyer},
	})
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey struct {
	EventController
}

func WrapEventControllerKey(obj *externglib.Object) *EventControllerKey {
	return &EventControllerKey{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalEventControllerKeyer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

func NewEventControllerKey(widget Widgetter) *EventControllerKey {
	var _arg1 *C.GtkWidget          // out
	var _cret *C.GtkEventController // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_new(_arg1)
	runtime.KeepAlive(widget)

	var _eventControllerKey *EventControllerKey // out

	_eventControllerKey = WrapEventControllerKey(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _eventControllerKey
}

func (controller *EventControllerKey) Forward(widget Widgetter) bool {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkWidget             // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_event_controller_key_forward(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (controller *EventControllerKey) Group() uint {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret C.guint                  // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_group(_arg0)
	runtime.KeepAlive(controller)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ImContext gets the IM context of a key controller.
func (controller *EventControllerKey) ImContext() IMContexter {
	var _arg0 *C.GtkEventControllerKey // out
	var _cret *C.GtkIMContext          // in

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))

	_cret = C.gtk_event_controller_key_get_im_context(_arg0)
	runtime.KeepAlive(controller)

	var _imContext IMContexter // out

	_imContext = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(IMContexter)

	return _imContext
}

func (controller *EventControllerKey) SetImContext(imContext IMContexter) {
	var _arg0 *C.GtkEventControllerKey // out
	var _arg1 *C.GtkIMContext          // out

	_arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(controller.Native()))
	_arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(_arg0, _arg1)
	runtime.KeepAlive(controller)
	runtime.KeepAlive(imContext)
}
