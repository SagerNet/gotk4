// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drag_cancel_reason_get_type()), F: marshalDragCancelReason},
		{T: externglib.Type(C.gdk_drag_protocol_get_type()), F: marshalDragProtocol},
		{T: externglib.Type(C.gdk_drag_action_get_type()), F: marshalDragAction},
	})
}

// DragCancelReason: used in DragContext to the reason of a cancelled DND
// operation.
type DragCancelReason int

const (
	// NoTarget: there is no suitable drop target.
	DragCancelReasonNoTarget DragCancelReason = 0
	// UserCancelled: drag cancelled by the user
	DragCancelReasonUserCancelled DragCancelReason = 1
	// error: unspecified error.
	DragCancelReasonError DragCancelReason = 2
)

func marshalDragCancelReason(p uintptr) (interface{}, error) {
	return DragCancelReason(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragProtocol: used in DragContext to indicate the protocol according to which
// DND is done.
type DragProtocol int

const (
	// none: no protocol.
	DragProtocolNone DragProtocol = 0
	// motif: the Motif DND protocol. No longer supported
	DragProtocolMotif DragProtocol = 1
	// xdnd: the Xdnd protocol.
	DragProtocolXdnd DragProtocol = 2
	// rootwin: an extension to the Xdnd protocol for unclaimed root window
	// drops.
	DragProtocolRootwin DragProtocol = 3
	// Win32Dropfiles: the simple WM_DROPFILES protocol.
	DragProtocolWin32Dropfiles DragProtocol = 4
	// ole2: the complex OLE2 DND protocol (not implemented).
	DragProtocolOle2 DragProtocol = 5
	// local: intra-application DND.
	DragProtocolLocal DragProtocol = 6
	// wayland: wayland DND protocol.
	DragProtocolWayland DragProtocol = 7
)

func marshalDragProtocol(p uintptr) (interface{}, error) {
	return DragProtocol(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragAction: used in DragContext to indicate what the destination should do
// with the dropped data.
type DragAction int

const (
	// DragActionDefault means nothing, and should not be used.
	DragActionDefault DragAction = 0b1
	// DragActionCopy: copy the data.
	DragActionCopy DragAction = 0b10
	// DragActionMove: move the data, i.e. first copy it, then delete it from
	// the source using the DELETE target of the X selection protocol.
	DragActionMove DragAction = 0b100
	// DragActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means.
	DragActionLink DragAction = 0b1000
	// DragActionPrivate: special action which tells the source that the
	// destination will do something that the source doesn’t understand.
	DragActionPrivate DragAction = 0b10000
	// DragActionAsk: ask the user what to do with the data.
	DragActionAsk DragAction = 0b100000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragAbort aborts a drag without dropping.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragAbort(context DragContext, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.guint32(time_)

	C.gdk_drag_abort(_arg1, _arg2)
}

// DragDrop drops on the current destination.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragDrop(context DragContext, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.guint32(time_)

	C.gdk_drag_drop(_arg1, _arg2)
}

// DragDropDone: inform GDK if the drop ended successfully. Passing false for
// @success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the @context.
//
// The DragContext will only take the first gdk_drag_drop_done() call as
// effective, if this function is called multiple times, all subsequent calls
// will be ignored.
func DragDropDone(context DragContext, success bool) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}

	C.gdk_drag_drop_done(_arg1, _arg2)
}

// DragDropSucceeded returns whether the dropped data has been successfully
// transferred. This function is intended to be used while handling a
// GDK_DROP_FINISHED event, its return value is meaningless at other times.
func DragDropSucceeded(context DragContext) bool {
	var _arg1 *C.GdkDragContext // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_drop_succeeded(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragFindWindowForScreen finds the destination window and DND protocol to use
// at the given pointer position.
//
// This function is called by the drag source to obtain the @dest_window and
// @protocol parameters for gdk_drag_motion().
func DragFindWindowForScreen(context DragContext, dragWindow Window, screen Screen, xRoot int, yRoot int) (Window, DragProtocol) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkWindow      // out
	var _arg3 *C.GdkScreen      // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out
	var _arg6 *C.GdkWindow      // in
	var _arg7 C.GdkDragProtocol // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(dragWindow.Native()))
	_arg3 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg4 = C.gint(xRoot)
	_arg5 = C.gint(yRoot)

	C.gdk_drag_find_window_for_screen(_arg1, _arg2, _arg3, _arg4, _arg5, &_arg6, &_arg7)

	var _destWindow Window     // out
	var _protocol DragProtocol // out

	_destWindow = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_arg6))).(Window)
	_protocol = DragProtocol(_arg7)

	return _destWindow, _protocol
}

// DragGetSelection returns the selection atom for the current source window.
func DragGetSelection(context DragContext) *Atom {
	var _arg1 *C.GdkDragContext // out
	var _cret C.GdkAtom         // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_drag_get_selection(_arg1)

	var _atom *Atom // out

	{
		var refTmpIn *C.GdkAtom
		var refTmpOut *Atom

		in0 := &_cret
		refTmpIn = in0

		refTmpOut = (*Atom)(unsafe.Pointer(refTmpIn))

		_atom = refTmpOut
	}

	return _atom
}

// DragMotion updates the drag context when the pointer moves or the set of
// actions changes.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragMotion(context DragContext, destWindow Window, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time_ uint32) bool {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkWindow      // out
	var _arg3 C.GdkDragProtocol // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out
	var _arg6 C.GdkDragAction   // out
	var _arg7 C.GdkDragAction   // out
	var _arg8 C.guint32         // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(destWindow.Native()))
	_arg3 = C.GdkDragProtocol(protocol)
	_arg4 = C.gint(xRoot)
	_arg5 = C.gint(yRoot)
	_arg6 = C.GdkDragAction(suggestedAction)
	_arg7 = C.GdkDragAction(possibleActions)
	_arg8 = C.guint32(time_)

	_cret = C.gdk_drag_motion(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragStatus selects one of the actions offered by the drag source.
//
// This function is called by the drag destination in response to
// gdk_drag_motion() called by the drag source.
func DragStatus(context DragContext, action DragAction, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.GdkDragAction   // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.GdkDragAction(action)
	_arg3 = C.guint32(time_)

	C.gdk_drag_status(_arg1, _arg2, _arg3)
}

// DropFinish ends the drag operation after a drop.
//
// This function is called by the drag destination.
func DropFinish(context DragContext, success bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}
	_arg3 = C.guint32(time_)

	C.gdk_drop_finish(_arg1, _arg2, _arg3)
}

// DropReply accepts or rejects a drop.
//
// This function is called by the drag destination in response to a drop
// initiated by the drag source.
func DropReply(context DragContext, accepted bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if accepted {
		_arg2 = C.TRUE
	}
	_arg3 = C.guint32(time_)

	C.gdk_drop_reply(_arg1, _arg2, _arg3)
}