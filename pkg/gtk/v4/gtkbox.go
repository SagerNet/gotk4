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
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBoxxer},
	})
}

// Boxxer describes Box's methods.
type Boxxer interface {
	gextras.Objector

	Append(child Widgetter)
	BaselinePosition() BaselinePosition
	Homogeneous() bool
	Spacing() int
	InsertChildAfter(child Widgetter, sibling Widgetter)
	Prepend(child Widgetter)
	Remove(child Widgetter)
	ReorderChildAfter(child Widgetter, sibling Widgetter)
	SetHomogeneous(homogeneous bool)
	SetSpacing(spacing int)
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
type Box struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Boxxer = (*Box)(nil)

func wrapBoxxer(obj *externglib.Object) Boxxer {
	return &Box{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBoxxer(obj), nil
}

// Append adds @child as the last child to @box.
func (box *Box) Append(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_append(_arg0, _arg1)
}

// BaselinePosition gets the value set by gtk_box_set_baseline_position().
func (box *Box) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBox             // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(_cret)

	return _baselinePosition
}

// Homogeneous returns whether the box is homogeneous (all children are the same
// size).
func (box *Box) Homogeneous() bool {
	var _arg0 *C.GtkBox  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_box_set_spacing().
func (box *Box) Spacing() int {
	var _arg0 *C.GtkBox // out
	var _cret C.int     // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertChildAfter inserts @child in the position after @sibling in the list of
// @box children.
//
// If @sibling is nil, insert @child at the first position.
func (box *Box) InsertChildAfter(child Widgetter, sibling Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_insert_child_after(_arg0, _arg1, _arg2)
}

// Prepend adds @child as the first child to @box.
func (box *Box) Prepend(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_prepend(_arg0, _arg1)
}

// Remove removes a child widget from @box.
//
// The child must have been added before with [method@Gtk.Box.append],
// [method@Gtk.Box.prepend], or [method@Gtk.Box.insert_child_after].
func (box *Box) Remove(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_remove(_arg0, _arg1)
}

// ReorderChildAfter moves @child to the position after @sibling in the list of
// @box children.
//
// If @sibling is nil, move @child to the first position.
func (box *Box) ReorderChildAfter(child Widgetter, sibling Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_reorder_child_after(_arg0, _arg1, _arg2)
}

// SetHomogeneous sets whether or not all children of @box are given equal space
// in the box.
func (box *Box) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBox  // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_set_homogeneous(_arg0, _arg1)
}

// SetSpacing sets the number of pixels to place between children of @box.
func (box *Box) SetSpacing(spacing int) {
	var _arg0 *C.GtkBox // out
	var _arg1 C.int     // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(spacing)

	C.gtk_box_set_spacing(_arg0, _arg1)
}
