// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_drop_target_async_get_type()), F: marshalDropTargetAsync},
	})
}

// DropTargetAsync: `GtkDropTargetAsync` is an event controller to receive
// Drag-and-Drop operations, asynchronously.
//
// It is the more complete but also more complex method of handling drop
// operations compared to [class@Gtk.DropTarget], and you should only use it if
// `GtkDropTarget` doesn't provide all the features you need.
//
// To use a `GtkDropTargetAsync` to receive drops on a widget, you create a
// `GtkDropTargetAsync` object, configure which data formats and actions you
// support, connect to its signals, and then attach it to the widget with
// [method@Gtk.Widget.add_controller].
//
// During a drag operation, the first signal that a `GtkDropTargetAsync` emits
// is [signal@Gtk.DropTargetAsync::accept], which is meant to determine whether
// the target is a possible drop site for the ongoing drop. The default handler
// for the ::accept signal accepts the drop if it finds a compatible data format
// and an action that is supported on both sides.
//
// If it is, and the widget becomes a target, you will receive a
// [signal@Gtk.DropTargetAsync::drag-enter] signal, followed by
// [signal@Gtk.DropTargetAsync::drag-motion] signals as the pointer moves,
// optionally a [signal@Gtk.DropTargetAsync::drop] signal when a drop happens,
// and finally a [signal@Gtk.DropTargetAsync::drag-leave] signal when the
// pointer moves off the widget.
//
// The ::drag-enter and ::drag-motion handler return a `GdkDragAction` to update
// the status of the ongoing operation. The ::drop handler should decide if it
// ultimately accepts the drop and if it does, it should initiate the data
// transfer and finish the operation by calling [method@Gdk.Drop.finish].
//
// Between the ::drag-enter and ::drag-leave signals the widget is a current
// drop target, and will receive the GTK_STATE_FLAG_DROP_ACTIVE state, which can
// be used by themes to style the widget as a drop target.
type DropTargetAsync interface {
	EventController

	// Actions:
	Actions() gdk.DragAction
	// Formats:
	Formats() *gdk.ContentFormats
	// RejectDropDropTargetAsync:
	RejectDropDropTargetAsync(drop gdk.Drop)
	// SetActionsDropTargetAsync:
	SetActionsDropTargetAsync(actions gdk.DragAction)
	// SetFormatsDropTargetAsync:
	SetFormatsDropTargetAsync(formats *gdk.ContentFormats)
}

// dropTargetAsync implements the DropTargetAsync class.
type dropTargetAsync struct {
	EventController
}

// WrapDropTargetAsync wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropTargetAsync(obj *externglib.Object) DropTargetAsync {
	return dropTargetAsync{
		EventController: WrapEventController(obj),
	}
}

func marshalDropTargetAsync(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropTargetAsync(obj), nil
}

// NewDropTargetAsync:
func NewDropTargetAsync(formats *gdk.ContentFormats, actions gdk.DragAction) DropTargetAsync {
	var _arg1 *C.GdkContentFormats  // out
	var _arg2 C.GdkDragAction       // out
	var _cret *C.GtkDropTargetAsync // in

	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(formats.Native()))
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gtk_drop_target_async_new(_arg1, _arg2)

	var _dropTargetAsync DropTargetAsync // out

	_dropTargetAsync = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(DropTargetAsync)

	return _dropTargetAsync
}

func (s dropTargetAsync) Actions() gdk.DragAction {
	var _arg0 *C.GtkDropTargetAsync // out
	var _cret C.GdkDragAction       // in

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_async_get_actions(_arg0)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

func (s dropTargetAsync) Formats() *gdk.ContentFormats {
	var _arg0 *C.GtkDropTargetAsync // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_target_async_get_formats(_arg0)

	var _contentFormats *gdk.ContentFormats // out

	_contentFormats = (*gdk.ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **gdk.ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

func (s dropTargetAsync) RejectDropDropTargetAsync(drop gdk.Drop) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 *C.GdkDrop            // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkDrop)(unsafe.Pointer(drop.Native()))

	C.gtk_drop_target_async_reject_drop(_arg0, _arg1)
}

func (s dropTargetAsync) SetActionsDropTargetAsync(actions gdk.DragAction) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 C.GdkDragAction       // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(s.Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drop_target_async_set_actions(_arg0, _arg1)
}

func (s dropTargetAsync) SetFormatsDropTargetAsync(formats *gdk.ContentFormats) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 *C.GdkContentFormats  // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(formats.Native()))

	C.gtk_drop_target_async_set_formats(_arg0, _arg1)
}
