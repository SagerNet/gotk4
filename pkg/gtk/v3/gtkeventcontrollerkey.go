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
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_key_get_type()), F: marshalEventControllerKey},
	})
}

// EventControllerKey is an event controller meant for situations where you need
// access to key events.
//
// This object was added in 3.24.
type EventControllerKey interface {
	EventController

	Forward(widget Widget) bool

	Group() uint
	// ImContext gets the IM context of a key controller.
	ImContext() IMContext

	SetImContext(imContext IMContext)
}

// eventControllerKey implements the EventControllerKey interface.
type eventControllerKey struct {
	EventController
}

var _ EventControllerKey = (*eventControllerKey)(nil)

// WrapEventControllerKey wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerKey(obj *externglib.Object) EventControllerKey {
	return EventControllerKey{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerKey(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerKey(obj), nil
}

// NewEventControllerKey constructs a class EventControllerKey.
func NewEventControllerKey(widget Widget) EventControllerKey {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var cret C.GtkEventControllerKey
	var goret1 EventControllerKey

	cret = C.gtk_event_controller_key_new(widget)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(EventControllerKey)

	return goret1
}

func (c eventControllerKey) Forward(widget Widget) bool {
	var arg0 *C.GtkEventControllerKey
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_event_controller_key_forward(arg0, widget)

	goret1 = C.bool(cret) != C.false

	return goret1
}

func (c eventControllerKey) Group() uint {
	var arg0 *C.GtkEventControllerKey

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var cret C.guint
	var goret1 uint

	cret = C.gtk_event_controller_key_get_group(arg0)

	goret1 = C.guint(cret)

	return goret1
}

// ImContext gets the IM context of a key controller.
func (c eventControllerKey) ImContext() IMContext {
	var arg0 *C.GtkEventControllerKey

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))

	var cret *C.GtkIMContext
	var goret1 IMContext

	cret = C.gtk_event_controller_key_get_im_context(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(IMContext)

	return goret1
}

func (c eventControllerKey) SetImContext(imContext IMContext) {
	var arg0 *C.GtkEventControllerKey
	var arg1 *C.GtkIMContext

	arg0 = (*C.GtkEventControllerKey)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkIMContext)(unsafe.Pointer(imContext.Native()))

	C.gtk_event_controller_key_set_im_context(arg0, imContext)
}
