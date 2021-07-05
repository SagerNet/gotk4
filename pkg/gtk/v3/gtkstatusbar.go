// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_statusbar_get_type()), F: marshalStatusbar},
	})
}

// Statusbar is usually placed along the bottom of an application's main Window.
// It may provide a regular commentary of the application's status (as is
// usually the case in a web browser, for example), or may be used to simply
// output a message when the status changes, (when an upload is complete in an
// FTP client, for example).
//
// Status bars in GTK+ maintain a stack of messages. The message at the top of
// the each bar’s stack is the one that will currently be displayed.
//
// Any messages added to a statusbar’s stack must specify a context id that is
// used to uniquely identify the source of a message. This context id can be
// generated by gtk_statusbar_get_context_id(), given a message and the
// statusbar that it will be added to. Note that messages are stored in a stack,
// and when choosing which message to display, the stack structure is adhered
// to, regardless of the context identifier of a message.
//
// One could say that a statusbar maintains one stack of messages for display
// purposes, but allows multiple message producers to maintain sub-stacks of the
// messages they produced (via context ids).
//
// Status bars are created using gtk_statusbar_new().
//
// Messages are added to the bar’s stack with gtk_statusbar_push().
//
// The message at the top of the stack can be removed using gtk_statusbar_pop().
// A message can be removed from anywhere in the stack if its message id was
// recorded at the time it was added. This is done using gtk_statusbar_remove().
//
//
// CSS node
//
// GtkStatusbar has a single CSS node with name statusbar.
type Statusbar interface {
	Box

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// ContextID returns a new context identifier, given a description of the
	// actual context. Note that the description is not shown in the UI.
	ContextID(contextDescription string) uint
	// MessageArea retrieves the box containing the label widget.
	MessageArea() Box
	// Pop removes the first message in the Statusbar’s stack with the given
	// context id.
	//
	// Note that this may not change the displayed message, if the message at
	// the top of the stack has a different context id.
	Pop(contextId uint)
	// Push pushes a new message onto a statusbar’s stack.
	Push(contextId uint, text string) uint
	// Remove forces the removal of a message from a statusbar’s stack. The
	// exact @context_id and @message_id must be specified.
	Remove(contextId uint, messageId uint)
	// RemoveAll forces the removal of all messages from a statusbar's stack
	// with the exact @context_id.
	RemoveAll(contextId uint)
}

// statusbar implements the Statusbar class.
type statusbar struct {
	Box
}

// WrapStatusbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapStatusbar(obj *externglib.Object) Statusbar {
	return statusbar{
		Box: WrapBox(obj),
	}
}

func marshalStatusbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusbar(obj), nil
}

// NewStatusbar creates a new Statusbar ready for messages.
func NewStatusbar() Statusbar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_statusbar_new()

	var _statusbar Statusbar // out

	_statusbar = WrapStatusbar(externglib.Take(unsafe.Pointer(_cret)))

	return _statusbar
}

func (s statusbar) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(s))
}

func (s statusbar) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(s))
}

func (s statusbar) ContextID(contextDescription string) uint {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 *C.gchar        // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(contextDescription))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_statusbar_get_context_id(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s statusbar) MessageArea() Box {
	var _arg0 *C.GtkStatusbar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_statusbar_get_message_area(_arg0)

	var _box Box // out

	_box = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Box)

	return _box
}

func (s statusbar) Pop(contextId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(contextId)

	C.gtk_statusbar_pop(_arg0, _arg1)
}

func (s statusbar) Push(contextId uint, text string) uint {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out
	var _arg2 *C.gchar        // out
	var _cret C.guint         // in

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(contextId)
	_arg2 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_statusbar_push(_arg0, _arg1, _arg2)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s statusbar) Remove(contextId uint, messageId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out
	var _arg2 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(contextId)
	_arg2 = C.guint(messageId)

	C.gtk_statusbar_remove(_arg0, _arg1, _arg2)
}

func (s statusbar) RemoveAll(contextId uint) {
	var _arg0 *C.GtkStatusbar // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GtkStatusbar)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(contextId)

	C.gtk_statusbar_remove_all(_arg0, _arg1)
}
