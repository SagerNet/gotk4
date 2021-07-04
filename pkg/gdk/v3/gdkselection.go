// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
import "C"

// SelectionConvert retrieves the contents of a selection in a given form.
func SelectionConvert(requestor Window, selection *Atom, target *Atom, time_ uint32) {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.GdkAtom    // out
	var _arg4 C.guint32    // out

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	_arg4 = C.guint32(time_)

	C.gdk_selection_convert(_arg1, _arg2, _arg3, _arg4)
}

// SelectionOwnerGet determines the owner of the given selection.
func SelectionOwnerGet(selection *Atom) Window {
	var _arg1 C.GdkAtom    // out
	var _cret *C.GdkWindow // in

	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}

	_cret = C.gdk_selection_owner_get(_arg1)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

// SelectionOwnerGetForDisplay: determine the owner of the given selection.
//
// Note that the return value may be owned by a different process if a foreign
// window was previously created for that window, but a new foreign window will
// never be created by this call.
func SelectionOwnerGetForDisplay(display Display, selection *Atom) Window {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.GdkAtom     // out
	var _cret *C.GdkWindow  // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}

	_cret = C.gdk_selection_owner_get_for_display(_arg1, _arg2)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

// SelectionOwnerSet sets the owner of the given selection.
func SelectionOwnerSet(owner Window, selection *Atom, time_ uint32, sendEvent bool) bool {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.guint32    // out
	var _arg4 C.gboolean   // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(owner.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	_arg3 = C.guint32(time_)
	if sendEvent {
		_arg4 = C.TRUE
	}

	_cret = C.gdk_selection_owner_set(_arg1, _arg2, _arg3, _arg4)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionOwnerSetForDisplay sets the Window @owner as the current owner of
// the selection @selection.
func SelectionOwnerSetForDisplay(display Display, owner Window, selection *Atom, time_ uint32, sendEvent bool) bool {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.GdkWindow  // out
	var _arg3 C.GdkAtom     // out
	var _arg4 C.guint32     // out
	var _arg5 C.gboolean    // out
	var _cret C.gboolean    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(owner.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	_arg4 = C.guint32(time_)
	if sendEvent {
		_arg5 = C.TRUE
	}

	_cret = C.gdk_selection_owner_set_for_display(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionSendNotify sends a response to SelectionRequest event.
func SelectionSendNotify(requestor Window, selection *Atom, target *Atom, property *Atom, time_ uint32) {
	var _arg1 *C.GdkWindow // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.GdkAtom    // out
	var _arg4 C.GdkAtom    // out
	var _arg5 C.guint32    // out

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = property

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg4 = *refTmpOut
	}
	_arg5 = C.guint32(time_)

	C.gdk_selection_send_notify(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// SelectionSendNotifyForDisplay: send a response to SelectionRequest event.
func SelectionSendNotifyForDisplay(display Display, requestor Window, selection *Atom, target *Atom, property *Atom, time_ uint32) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.GdkWindow  // out
	var _arg3 C.GdkAtom     // out
	var _arg4 C.GdkAtom     // out
	var _arg5 C.GdkAtom     // out
	var _arg6 C.guint32     // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg4 = *refTmpOut
	}
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = property

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg5 = *refTmpOut
	}
	_arg6 = C.guint32(time_)

	C.gdk_selection_send_notify_for_display(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}
