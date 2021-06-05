// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_drop_controller_motion_get_type()), F: marshalDropControllerMotion},
	})
}

// DropControllerMotion is an event controller meant for tracking the pointer
// hovering over a widget during a drag and drop operation.
//
// It is modeled after EventControllerMotion so if you have used that, this
// should feel really familiar.
//
// The drop controller is not able to accept drops, use DropTarget for that
// purpose.
type DropControllerMotion interface {
	EventController

	// ContainsPointer returns the value of the
	// GtkDropControllerMotion:contains-pointer property.
	ContainsPointer() bool
	// Drop returns the value of the GtkDropControllerMotion:drop property.
	Drop() gdk.Drop
	// IsPointer returns the value of the GtkDropControllerMotion:is-pointer
	// property.
	IsPointer() bool
}

// dropControllerMotion implements the DropControllerMotion interface.
type dropControllerMotion struct {
	EventController
}

var _ DropControllerMotion = (*dropControllerMotion)(nil)

// WrapDropControllerMotion wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropControllerMotion(obj *externglib.Object) DropControllerMotion {
	return DropControllerMotion{
		EventController: WrapEventController(obj),
	}
}

func marshalDropControllerMotion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropControllerMotion(obj), nil
}

// NewDropControllerMotion constructs a class DropControllerMotion.
func NewDropControllerMotion() DropControllerMotion {
	var cret C.GtkDropControllerMotion
	var goret1 DropControllerMotion

	cret = C.gtk_drop_controller_motion_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DropControllerMotion)

	return goret1
}

// ContainsPointer returns the value of the
// GtkDropControllerMotion:contains-pointer property.
func (s dropControllerMotion) ContainsPointer() bool {
	var arg0 *C.GtkDropControllerMotion

	arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_drop_controller_motion_contains_pointer(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Drop returns the value of the GtkDropControllerMotion:drop property.
func (s dropControllerMotion) Drop() gdk.Drop {
	var arg0 *C.GtkDropControllerMotion

	arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(s.Native()))

	var cret *C.GdkDrop
	var goret1 gdk.Drop

	cret = C.gtk_drop_controller_motion_get_drop(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Drop)

	return goret1
}

// IsPointer returns the value of the GtkDropControllerMotion:is-pointer
// property.
func (s dropControllerMotion) IsPointer() bool {
	var arg0 *C.GtkDropControllerMotion

	arg0 = (*C.GtkDropControllerMotion)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_drop_controller_motion_is_pointer(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}
