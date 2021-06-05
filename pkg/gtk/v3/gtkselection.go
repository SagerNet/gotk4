// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_target_entry_get_type()), F: marshalTargetEntry},
		{T: externglib.Type(C.gtk_target_list_get_type()), F: marshalTargetList},
	})
}

// SelectionAddTarget appends a specified target to the list of supported
// targets for a given widget and selection.
func SelectionAddTarget(widget Widget, selection gdk.Atom, target gdk.Atom, info uint) {
	var arg1 *C.GtkWidget
	var arg2 C.GdkAtom
	var arg3 C.GdkAtom
	var arg4 C.guint

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	arg3 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	arg4 = C.guint(info)

	C.gtk_selection_add_target(widget, selection, target, info)
}

// SelectionAddTargets prepends a table of targets to the list of supported
// targets for a given widget and selection.
func SelectionAddTargets(widget Widget, selection gdk.Atom, targets []TargetEntry) {

	C.gtk_selection_add_targets(widget, selection, targets, ntargets)
}

// SelectionClearTargets: remove all targets registered for the given selection
// for the widget.
func SelectionClearTargets(widget Widget, selection gdk.Atom) {
	var arg1 *C.GtkWidget
	var arg2 C.GdkAtom

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))

	C.gtk_selection_clear_targets(widget, selection)
}

// SelectionConvert requests the contents of a selection. When received, a
// “selection-received” signal will be generated.
func SelectionConvert(widget Widget, selection gdk.Atom, target gdk.Atom, time_ uint32) bool {
	var arg1 *C.GtkWidget
	var arg2 C.GdkAtom
	var arg3 C.GdkAtom
	var arg4 C.guint32

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	arg3 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	arg4 = C.guint32(time_)

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_selection_convert(widget, selection, target, time_)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SelectionOwnerSet claims ownership of a given selection for a particular
// widget, or, if @widget is nil, release ownership of the selection.
func SelectionOwnerSet(widget Widget, selection gdk.Atom, time_ uint32) bool {
	var arg1 *C.GtkWidget
	var arg2 C.GdkAtom
	var arg3 C.guint32

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	arg3 = C.guint32(time_)

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_selection_owner_set(widget, selection, time_)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SelectionOwnerSetForDisplay: claim ownership of a given selection for a
// particular widget, or, if @widget is nil, release ownership of the selection.
func SelectionOwnerSetForDisplay(display gdk.Display, widget Widget, selection gdk.Atom, time_ uint32) bool {
	var arg1 *C.GdkDisplay
	var arg2 *C.GtkWidget
	var arg3 C.GdkAtom
	var arg4 C.guint32

	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg3 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	arg4 = C.guint32(time_)

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_selection_owner_set_for_display(display, widget, selection, time_)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SelectionRemoveAll removes all handlers and unsets ownership of all
// selections for a widget. Called when widget is being destroyed. This function
// will not generally be called by applications.
func SelectionRemoveAll(widget Widget) {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_selection_remove_all(widget)
}

// TargetTableFree: this function frees a target table as returned by
// gtk_target_table_new_from_list()
func TargetTableFree(targets []TargetEntry) {

	C.gtk_target_table_free(targets, nTargets)
}

// TargetTableNewFromList: this function creates an TargetEntry array that
// contains the same targets as the passed list. The returned table is newly
// allocated and should be freed using gtk_target_table_free() when no longer
// needed.
func TargetTableNewFromList(list *TargetList) (nTargets int, targetEntrys []TargetEntry) {
	var arg1 *C.GtkTargetList

	arg1 = (*C.GtkTargetList)(unsafe.Pointer(list.Native()))

	var cret *C.GtkTargetEntry
	var arg2 *C.gint
	var goret2 []TargetEntry

	cret = C.gtk_target_table_new_from_list(list, &arg2)

	goret2 = make([]TargetEntry, arg2)
	for i := 0; i < uintptr(arg2); i++ {
		src := (C.GtkTargetEntry)(ptr.Add(unsafe.Pointer(cret), i))
		goret2[i] = WrapTargetEntry(unsafe.Pointer(src))
		runtime.SetFinalizer(goret2[i], func(v TargetEntry) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret2, goret2
}

// TargetsIncludeText determines if any of the targets in @targets can be used
// to provide text.
func TargetsIncludeText(targets []gdk.Atom) bool {

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_targets_include_text(targets, nTargets)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// TargetsIncludeURI determines if any of the targets in @targets can be used to
// provide an uri list.
func TargetsIncludeURI(targets []gdk.Atom) bool {

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_targets_include_uri(targets, nTargets)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// TargetEntry: a TargetEntry represents a single type of data than can be
// supplied for by a widget for a selection or for supplied or received during
// drag-and-drop.
type TargetEntry struct {
	native C.GtkTargetEntry
}

// WrapTargetEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetEntry(ptr unsafe.Pointer) *TargetEntry {
	if ptr == nil {
		return nil
	}

	return (*TargetEntry)(ptr)
}

func marshalTargetEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTargetEntry(unsafe.Pointer(b)), nil
}

// NewTargetEntry constructs a struct TargetEntry.
func NewTargetEntry(target string, flags uint, info uint) *TargetEntry {
	var arg1 *C.gchar
	var arg2 C.guint
	var arg3 C.guint

	arg1 = (*C.gchar)(C.CString(target))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.guint(flags)
	arg3 = C.guint(info)

	var cret *C.GtkTargetEntry
	var goret1 *TargetEntry

	cret = C.gtk_target_entry_new(target, flags, info)

	goret1 = WrapTargetEntry(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *TargetEntry) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Native returns the underlying C source pointer.
func (t *TargetEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Target gets the field inside the struct.
func (t *TargetEntry) Target() string {
	v = C.GoString(t.native.target)
}

// Flags gets the field inside the struct.
func (t *TargetEntry) Flags() uint {
	v = C.guint(t.native.flags)
}

// Info gets the field inside the struct.
func (t *TargetEntry) Info() uint {
	v = C.guint(t.native.info)
}

// Copy makes a copy of a TargetEntry and its data.
func (d *TargetEntry) Copy() *TargetEntry {
	var arg0 *C.GtkTargetEntry

	arg0 = (*C.GtkTargetEntry)(unsafe.Pointer(d.Native()))

	var cret *C.GtkTargetEntry
	var goret1 *TargetEntry

	cret = C.gtk_target_entry_copy(arg0)

	goret1 = WrapTargetEntry(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *TargetEntry) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Free frees a TargetEntry returned from gtk_target_entry_new() or
// gtk_target_entry_copy().
func (d *TargetEntry) Free() {
	var arg0 *C.GtkTargetEntry

	arg0 = (*C.GtkTargetEntry)(unsafe.Pointer(d.Native()))

	C.gtk_target_entry_free(arg0)
}

// TargetList: a TargetList-struct is a reference counted list of TargetPair and
// should be treated as opaque.
type TargetList struct {
	native C.GtkTargetList
}

// WrapTargetList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetList(ptr unsafe.Pointer) *TargetList {
	if ptr == nil {
		return nil
	}

	return (*TargetList)(ptr)
}

func marshalTargetList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTargetList(unsafe.Pointer(b)), nil
}

// NewTargetList constructs a struct TargetList.
func NewTargetList(targets []TargetEntry) *TargetList {

	var cret *C.GtkTargetList
	var goret1 *TargetList

	cret = C.gtk_target_list_new(targets, ntargets)

	goret1 = WrapTargetList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *TargetList) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Native returns the underlying C source pointer.
func (t *TargetList) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Add appends another target to a TargetList.
func (l *TargetList) Add(target gdk.Atom, flags uint, info uint) {
	var arg0 *C.GtkTargetList
	var arg1 C.GdkAtom
	var arg2 C.guint
	var arg3 C.guint

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	arg2 = C.guint(flags)
	arg3 = C.guint(info)

	C.gtk_target_list_add(arg0, target, flags, info)
}

// AddImageTargets appends the image targets supported by SelectionData to the
// target list. All targets are added with the same @info.
func (l *TargetList) AddImageTargets(info uint, writable bool) {
	var arg0 *C.GtkTargetList
	var arg1 C.guint
	var arg2 C.gboolean

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = C.guint(info)
	if writable {
		arg2 = C.gboolean(1)
	}

	C.gtk_target_list_add_image_targets(arg0, info, writable)
}

// AddRichTextTargets appends the rich text targets registered with
// gtk_text_buffer_register_serialize_format() or
// gtk_text_buffer_register_deserialize_format() to the target list. All targets
// are added with the same @info.
func (l *TargetList) AddRichTextTargets(info uint, deserializable bool, buffer TextBuffer) {
	var arg0 *C.GtkTargetList
	var arg1 C.guint
	var arg2 C.gboolean
	var arg3 *C.GtkTextBuffer

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = C.guint(info)
	if deserializable {
		arg2 = C.gboolean(1)
	}
	arg3 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_target_list_add_rich_text_targets(arg0, info, deserializable, buffer)
}

// AddTable prepends a table of TargetEntry to a target list.
func (l *TargetList) AddTable(targets []TargetEntry) {
	var arg0 *C.GtkTargetList

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))

	C.gtk_target_list_add_table(arg0, targets, ntargets)
}

// AddTextTargets appends the text targets supported by SelectionData to the
// target list. All targets are added with the same @info.
func (l *TargetList) AddTextTargets(info uint) {
	var arg0 *C.GtkTargetList
	var arg1 C.guint

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = C.guint(info)

	C.gtk_target_list_add_text_targets(arg0, info)
}

// AddURITargets appends the URI targets supported by SelectionData to the
// target list. All targets are added with the same @info.
func (l *TargetList) AddURITargets(info uint) {
	var arg0 *C.GtkTargetList
	var arg1 C.guint

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = C.guint(info)

	C.gtk_target_list_add_uri_targets(arg0, info)
}

// Find looks up a given target in a TargetList.
func (l *TargetList) Find(target gdk.Atom) (info uint, ok bool) {
	var arg0 *C.GtkTargetList
	var arg1 C.GdkAtom

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = (C.GdkAtom)(unsafe.Pointer(target.Native()))

	var arg2 *C.guint
	var ret2 uint
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_target_list_find(arg0, target, &arg2)

	ret2 = *C.guint(arg2)
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

// Ref increases the reference count of a TargetList by one.
func (l *TargetList) Ref() *TargetList {
	var arg0 *C.GtkTargetList

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))

	var cret *C.GtkTargetList
	var goret1 *TargetList

	cret = C.gtk_target_list_ref(arg0)

	goret1 = WrapTargetList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *TargetList) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Remove removes a target from a target list.
func (l *TargetList) Remove(target gdk.Atom) {
	var arg0 *C.GtkTargetList
	var arg1 C.GdkAtom

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))
	arg1 = (C.GdkAtom)(unsafe.Pointer(target.Native()))

	C.gtk_target_list_remove(arg0, target)
}

// Unref decreases the reference count of a TargetList by one. If the resulting
// reference count is zero, frees the list.
func (l *TargetList) Unref() {
	var arg0 *C.GtkTargetList

	arg0 = (*C.GtkTargetList)(unsafe.Pointer(l.Native()))

	C.gtk_target_list_unref(arg0)
}

// TargetPair: a TargetPair is used to represent the same information as a table
// of TargetEntry, but in an efficient form.
type TargetPair struct {
	native C.GtkTargetPair
}

// WrapTargetPair wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTargetPair(ptr unsafe.Pointer) *TargetPair {
	if ptr == nil {
		return nil
	}

	return (*TargetPair)(ptr)
}

func marshalTargetPair(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTargetPair(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TargetPair) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Target gets the field inside the struct.
func (t *TargetPair) Target() gdk.Atom {
	v = gdk.WrapAtom(unsafe.Pointer(t.native.target))
}

// Flags gets the field inside the struct.
func (t *TargetPair) Flags() uint {
	v = C.guint(t.native.flags)
}

// Info gets the field inside the struct.
func (t *TargetPair) Info() uint {
	v = C.guint(t.native.info)
}
