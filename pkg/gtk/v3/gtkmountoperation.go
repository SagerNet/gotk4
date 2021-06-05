// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation interface {
	gio.MountOperation

	// Parent gets the transient parent used by the MountOperation
	Parent() Window
	// Screen gets the screen on which windows of the MountOperation will be
	// shown.
	Screen() gdk.Screen
	// IsShowing returns whether the MountOperation is currently displaying a
	// window.
	IsShowing() bool
	// SetParent sets the transient parent for windows shown by the
	// MountOperation.
	SetParent(parent Window)
	// SetScreen sets the screen to show windows of the MountOperation on.
	SetScreen(screen gdk.Screen)
}

// mountOperation implements the MountOperation interface.
type mountOperation struct {
	gio.MountOperation
}

var _ MountOperation = (*mountOperation)(nil)

// WrapMountOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return MountOperation{
		gio.MountOperation: gio.WrapMountOperation(obj),
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

// NewMountOperation constructs a class MountOperation.
func NewMountOperation(parent Window) MountOperation {
	var arg1 *C.GtkWindow

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	var cret C.GtkMountOperation
	var goret1 MountOperation

	cret = C.gtk_mount_operation_new(parent)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MountOperation)

	return goret1
}

// Parent gets the transient parent used by the MountOperation
func (o mountOperation) Parent() Window {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret *C.GtkWindow
	var goret1 Window

	cret = C.gtk_mount_operation_get_parent(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Window)

	return goret1
}

// Screen gets the screen on which windows of the MountOperation will be
// shown.
func (o mountOperation) Screen() gdk.Screen {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret *C.GdkScreen
	var goret1 gdk.Screen

	cret = C.gtk_mount_operation_get_screen(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Screen)

	return goret1
}

// IsShowing returns whether the MountOperation is currently displaying a
// window.
func (o mountOperation) IsShowing() bool {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_mount_operation_is_showing(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetParent sets the transient parent for windows shown by the
// MountOperation.
func (o mountOperation) SetParent(parent Window) {
	var arg0 *C.GtkMountOperation
	var arg1 *C.GtkWindow

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(arg0, parent)
}

// SetScreen sets the screen to show windows of the MountOperation on.
func (o mountOperation) SetScreen(screen gdk.Screen) {
	var arg0 *C.GtkMountOperation
	var arg1 *C.GdkScreen

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_mount_operation_set_screen(arg0, screen)
}

type MountOperationPrivate struct {
	native C.GtkMountOperationPrivate
}

// WrapMountOperationPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMountOperationPrivate(ptr unsafe.Pointer) *MountOperationPrivate {
	if ptr == nil {
		return nil
	}

	return (*MountOperationPrivate)(ptr)
}

func marshalMountOperationPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMountOperationPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (m *MountOperationPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}
