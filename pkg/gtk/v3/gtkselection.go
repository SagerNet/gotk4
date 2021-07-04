// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_target_flags_get_type()), F: marshalTargetFlags},
		{T: externglib.Type(C.gtk_target_entry_get_type()), F: marshalTargetEntry},
		{T: externglib.Type(C.gtk_target_list_get_type()), F: marshalTargetList},
	})
}

// TargetFlags: the TargetFlags enumeration is used to specify constraints on a
// TargetEntry.
type TargetFlags int

const (
	// TargetFlagsSameApp: if this is set, the target will only be selected for
	// drags within a single application.
	TargetFlagsSameApp TargetFlags = 0b1
	// TargetFlagsSameWidget: if this is set, the target will only be selected
	// for drags within a single widget.
	TargetFlagsSameWidget TargetFlags = 0b10
	// TargetFlagsOtherApp: if this is set, the target will not be selected for
	// drags within a single application.
	TargetFlagsOtherApp TargetFlags = 0b100
	// TargetFlagsOtherWidget: if this is set, the target will not be selected
	// for drags withing a single widget.
	TargetFlagsOtherWidget TargetFlags = 0b1000
)

func marshalTargetFlags(p uintptr) (interface{}, error) {
	return TargetFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SelectionAddTarget appends a specified target to the list of supported
// targets for a given widget and selection.
func SelectionAddTarget(widget Widget, selection *gdk.Atom, target *gdk.Atom, info uint) {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.GdkAtom    // out
	var _arg4 C.guint      // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	_arg4 = C.guint(info)

	C.gtk_selection_add_target(_arg1, _arg2, _arg3, _arg4)
}

// SelectionAddTargets prepends a table of targets to the list of supported
// targets for a given widget and selection.
func SelectionAddTargets(widget Widget, selection *gdk.Atom, targets []TargetEntry) {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GdkAtom    // out
	var _arg3 *C.GtkTargetEntry
	var _arg4 C.guint

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	_arg4 = C.guint(len(targets))
	_arg3 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	C.gtk_selection_add_targets(_arg1, _arg2, _arg3, _arg4)
}

// SelectionClearTargets: remove all targets registered for the given selection
// for the widget.
func SelectionClearTargets(widget Widget, selection *gdk.Atom) {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GdkAtom    // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}

	C.gtk_selection_clear_targets(_arg1, _arg2)
}

// SelectionConvert requests the contents of a selection. When received, a
// “selection-received” signal will be generated.
func SelectionConvert(widget Widget, selection *gdk.Atom, target *gdk.Atom, time_ uint32) bool {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.GdkAtom    // out
	var _arg4 C.guint32    // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	_arg4 = C.guint32(time_)

	_cret = C.gtk_selection_convert(_arg1, _arg2, _arg3, _arg4)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionOwnerSet claims ownership of a given selection for a particular
// widget, or, if @widget is nil, release ownership of the selection.
func SelectionOwnerSet(widget Widget, selection *gdk.Atom, time_ uint32) bool {
	var _arg1 *C.GtkWidget // out
	var _arg2 C.GdkAtom    // out
	var _arg3 C.guint32    // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg2 = *refTmpOut
	}
	_arg3 = C.guint32(time_)

	_cret = C.gtk_selection_owner_set(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionOwnerSetForDisplay: claim ownership of a given selection for a
// particular widget, or, if @widget is nil, release ownership of the selection.
func SelectionOwnerSetForDisplay(display gdk.Display, widget Widget, selection *gdk.Atom, time_ uint32) bool {
	var _arg1 *C.GdkDisplay // out
	var _arg2 *C.GtkWidget  // out
	var _arg3 C.GdkAtom     // out
	var _arg4 C.guint32     // out
	var _cret C.gboolean    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg3 = *refTmpOut
	}
	_arg4 = C.guint32(time_)

	_cret = C.gtk_selection_owner_set_for_display(_arg1, _arg2, _arg3, _arg4)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionRemoveAll removes all handlers and unsets ownership of all
// selections for a widget. Called when widget is being destroyed. This function
// will not generally be called by applications.
func SelectionRemoveAll(widget Widget) {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_selection_remove_all(_arg1)
}

// TargetTableFree: this function frees a target table as returned by
// gtk_target_table_new_from_list()
func TargetTableFree(targets []TargetEntry) {
	var _arg1 *C.GtkTargetEntry
	var _arg2 C.gint

	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	C.gtk_target_table_free(_arg1, _arg2)
}

// TargetsIncludeImage determines if any of the targets in @targets can be used
// to provide a Pixbuf.
func TargetsIncludeImage(targets []*gdk.Atom, writable bool) bool {
	var _arg1 *C.GdkAtom
	var _arg2 C.gint
	var _arg3 C.gboolean // out
	var _cret C.gboolean // in

	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))
	if writable {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_targets_include_image(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TargetsIncludeRichText determines if any of the targets in @targets can be
// used to provide rich text.
func TargetsIncludeRichText(targets []*gdk.Atom, buffer TextBuffer) bool {
	var _arg1 *C.GdkAtom
	var _arg2 C.gint
	var _arg3 *C.GtkTextBuffer // out
	var _cret C.gboolean       // in

	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))
	_arg3 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_targets_include_rich_text(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TargetsIncludeText determines if any of the targets in @targets can be used
// to provide text.
func TargetsIncludeText(targets []*gdk.Atom) bool {
	var _arg1 *C.GdkAtom
	var _arg2 C.gint
	var _cret C.gboolean // in

	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))

	_cret = C.gtk_targets_include_text(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TargetsIncludeURI determines if any of the targets in @targets can be used to
// provide an uri list.
func TargetsIncludeURI(targets []*gdk.Atom) bool {
	var _arg1 *C.GdkAtom
	var _arg2 C.gint
	var _cret C.gboolean // in

	_arg2 = C.gint(len(targets))
	_arg1 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))

	_cret = C.gtk_targets_include_uri(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TargetEntry: a TargetEntry represents a single type of data than can be
// supplied for by a widget for a selection or for supplied or received during
// drag-and-drop.
type TargetEntry C.GtkTargetEntry

// WrapTargetEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetEntry(ptr unsafe.Pointer) *TargetEntry {
	return (*TargetEntry)(ptr)
}

func marshalTargetEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TargetEntry)(unsafe.Pointer(b)), nil
}

// NewTargetEntry constructs a struct TargetEntry.
func NewTargetEntry(target string, flags uint, info uint) *TargetEntry {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _cret *C.GtkTargetEntry // in

	_arg1 = (*C.gchar)(C.CString(target))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(flags)
	_arg3 = C.guint(info)

	_cret = C.gtk_target_entry_new(_arg1, _arg2, _arg3)

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_targetEntry, func(v **TargetEntry) {
		C.free(unsafe.Pointer(v))
	})

	return _targetEntry
}

// Native returns the underlying C source pointer.
func (t *TargetEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// Copy frees a TargetEntry returned from gtk_target_entry_new() or
// gtk_target_entry_copy().
func (d *TargetEntry) Copy() *TargetEntry {
	var _arg0 *C.GtkTargetEntry // out
	var _cret *C.GtkTargetEntry // in

	_arg0 = (*C.GtkTargetEntry)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_target_entry_copy(_arg0)

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_targetEntry, func(v **TargetEntry) {
		C.free(unsafe.Pointer(v))
	})

	return _targetEntry
}

// Free frees a TargetEntry returned from gtk_target_entry_new() or
// gtk_target_entry_copy().
func (d *TargetEntry) Free() {
	var _arg0 *C.GtkTargetEntry // out

	_arg0 = (*C.GtkTargetEntry)(unsafe.Pointer(d.Native()))

	C.gtk_target_entry_free(_arg0)
}

// TargetList: a TargetList-struct is a reference counted list of TargetPair and
// should be treated as opaque.
type TargetList C.GtkTargetList

// WrapTargetList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetList(ptr unsafe.Pointer) *TargetList {
	return (*TargetList)(ptr)
}

func marshalTargetList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TargetList)(unsafe.Pointer(b)), nil
}

// NewTargetList constructs a struct TargetList.
func NewTargetList(targets []TargetEntry) *TargetList {
	var _arg1 *C.GtkTargetEntry
	var _arg2 C.guint
	var _cret *C.GtkTargetList // in

	_arg2 = C.guint(len(targets))
	_arg1 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	_cret = C.gtk_target_list_new(_arg1, _arg2)

	var _targetList *TargetList // out

	_targetList = (*TargetList)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_targetList, func(v **TargetList) {
		C.free(unsafe.Pointer(v))
	})

	return _targetList
}

// Native returns the underlying C source pointer.
func (t *TargetList) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// Add decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Add(target *gdk.Atom, flags uint, info uint) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.GdkAtom        // out
	var _arg2 C.guint          // out
	var _arg3 C.guint          // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}
	_arg2 = C.guint(flags)
	_arg3 = C.guint(info)

	C.gtk_target_list_add(_arg0, _arg1, _arg2, _arg3)
}

// AddImageTargets decreases the reference count of a TargetList by one. If the
// resulting reference count is zero, frees the list.
func (l *TargetList) AddImageTargets(info uint, writable bool) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out
	var _arg2 C.gboolean       // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(info)
	if writable {
		_arg2 = C.TRUE
	}

	C.gtk_target_list_add_image_targets(_arg0, _arg1, _arg2)
}

// AddRichTextTargets decreases the reference count of a TargetList by one. If
// the resulting reference count is zero, frees the list.
func (l *TargetList) AddRichTextTargets(info uint, deserializable bool, buffer TextBuffer) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out
	var _arg2 C.gboolean       // out
	var _arg3 *C.GtkTextBuffer // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(info)
	if deserializable {
		_arg2 = C.TRUE
	}
	_arg3 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_target_list_add_rich_text_targets(_arg0, _arg1, _arg2, _arg3)
}

// AddTable decreases the reference count of a TargetList by one. If the
// resulting reference count is zero, frees the list.
func (l *TargetList) AddTable(targets []TargetEntry) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 *C.GtkTargetEntry
	var _arg2 C.guint

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	_arg2 = C.guint(len(targets))
	_arg1 = (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	C.gtk_target_list_add_table(_arg0, _arg1, _arg2)
}

// AddTextTargets decreases the reference count of a TargetList by one. If the
// resulting reference count is zero, frees the list.
func (l *TargetList) AddTextTargets(info uint) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(info)

	C.gtk_target_list_add_text_targets(_arg0, _arg1)
}

// AddURITargets decreases the reference count of a TargetList by one. If the
// resulting reference count is zero, frees the list.
func (l *TargetList) AddURITargets(info uint) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(info)

	C.gtk_target_list_add_uri_targets(_arg0, _arg1)
}

// Find decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Find(target *gdk.Atom) (uint, bool) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.GdkAtom        // out
	var _arg2 C.guint          // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}

	_cret = C.gtk_target_list_find(_arg0, _arg1, &_arg2)

	var _info uint // out
	var _ok bool   // out

	_info = uint(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _info, _ok
}

// Ref decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Ref() *TargetList {
	var _arg0 *C.GtkTargetList // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_target_list_ref(_arg0)

	var _targetList *TargetList // out

	_targetList = (*TargetList)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_targetList, func(v **TargetList) {
		C.free(unsafe.Pointer(v))
	})

	return _targetList
}

// Remove decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Remove(target *gdk.Atom) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.GdkAtom        // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = target

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}

	C.gtk_target_list_remove(_arg0, _arg1)
}

// Unref decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Unref() {
	var _arg0 *C.GtkTargetList // out

	_arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))

	C.gtk_target_list_unref(_arg0)
}

// TargetPair: a TargetPair is used to represent the same information as a table
// of TargetEntry, but in an efficient form.
type TargetPair C.GtkTargetPair

// WrapTargetPair wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetPair(ptr unsafe.Pointer) *TargetPair {
	return (*TargetPair)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TargetPair) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}