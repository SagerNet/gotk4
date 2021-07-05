// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBox},
	})
}

// Box: the `GtkBox` widget arranges child widgets into a single row or column.
//
// !An example GtkBox (box.png)
//
// Whether it is a row or column depends on the value of its
// [property@Gtk.Orientable:orientation] property. Within the other dimension,
// all children are allocated the same size. Of course, the
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties can
// be used on the children to influence their allocation.
//
// Use repeated calls to [method@Gtk.Box.append] to pack widgets into a `GtkBox`
// from start to end. Use [method@Gtk.Box.remove] to remove widgets from the
// `GtkBox`. [method@Gtk.Box.insert_child_after] can be used to add a child at a
// particular position.
//
// Use [method@Gtk.Box.set_homogeneous] to specify whether or not all children
// of the `GtkBox` are forced to get the same amount of space.
//
// Use [method@Gtk.Box.set_spacing] to determine how much space will be
// minimally placed between all children in the `GtkBox`. Note that spacing is
// added *between* the children.
//
// Use [method@Gtk.Box.reorder_child_after] to move a child to a different place
// in the box.
//
//
// CSS nodes
//
// `GtkBox` uses a single CSS node with name box.
//
//
// Accessibility
//
// `GtkBox` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Box interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// Append adds @child as the last child to @box.
	Append(child Widget)
	// BaselinePosition gets the value set by gtk_box_set_baseline_position().
	BaselinePosition() BaselinePosition
	// Homogeneous returns whether the box is homogeneous (all children are the
	// same size).
	Homogeneous() bool
	// Spacing gets the value set by gtk_box_set_spacing().
	Spacing() int
	// InsertChildAfter inserts @child in the position after @sibling in the
	// list of @box children.
	//
	// If @sibling is nil, insert @child at the first position.
	InsertChildAfter(child Widget, sibling Widget)
	// Prepend adds @child as the first child to @box.
	Prepend(child Widget)
	// Remove removes a child widget from @box.
	//
	// The child must have been added before with [method@Gtk.Box.append],
	// [method@Gtk.Box.prepend], or [method@Gtk.Box.insert_child_after].
	Remove(child Widget)
	// ReorderChildAfter moves @child to the position after @sibling in the list
	// of @box children.
	//
	// If @sibling is nil, move @child to the first position.
	ReorderChildAfter(child Widget, sibling Widget)
	// SetBaselinePosition sets the baseline position of a box.
	//
	// This affects only horizontal boxes with at least one baseline aligned
	// child. If there is more vertical space available than requested, and the
	// baseline is not allocated by the parent then @position is used to
	// allocate the baseline with respect to the extra space available.
	SetBaselinePosition(position BaselinePosition)
	// SetHomogeneous sets whether or not all children of @box are given equal
	// space in the box.
	SetHomogeneous(homogeneous bool)
	// SetSpacing sets the number of pixels to place between children of @box.
	SetSpacing(spacing int)
}

// box implements the Box class.
type box struct {
	Widget
}

// WrapBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapBox(obj *externglib.Object) Box {
	return box{
		Widget: WrapWidget(obj),
	}
}

func marshalBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBox(obj), nil
}

// NewBox creates a new `GtkBox`.
func NewBox(orientation Orientation, spacing int) Box {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.int            // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.int(spacing)

	_cret = C.gtk_box_new(_arg1, _arg2)

	var _box Box // out

	_box = WrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

func (b box) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(b))
}

func (b box) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(b))
}

func (b box) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(b))
}

func (b box) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(b))
}

func (b box) Append(child Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_append(_arg0, _arg1)
}

func (b box) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBox             // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (b box) Homogeneous() bool {
	var _arg0 *C.GtkBox  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b box) Spacing() int {
	var _arg0 *C.GtkBox // out
	var _cret C.int     // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_box_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (b box) InsertChildAfter(child Widget, sibling Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_insert_child_after(_arg0, _arg1, _arg2)
}

func (b box) Prepend(child Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_prepend(_arg0, _arg1)
}

func (b box) Remove(child Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_remove(_arg0, _arg1)
}

func (b box) ReorderChildAfter(child Widget, sibling Widget) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_reorder_child_after(_arg0, _arg1, _arg2)
}

func (b box) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBox             // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_set_baseline_position(_arg0, _arg1)
}

func (b box) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBox  // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_set_homogeneous(_arg0, _arg1)
}

func (b box) SetSpacing(spacing int) {
	var _arg0 *C.GtkBox // out
	var _arg1 C.int     // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(spacing)

	C.gtk_box_set_spacing(_arg0, _arg1)
}
