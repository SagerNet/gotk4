// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBoxer},
	})
}

// Box: GtkBox widget arranges child widgets into a single row or column.
//
// !An example GtkBox (box.png)
//
// Whether it is a row or column depends on the value of its
// gtk.Orientable:orientation property. Within the other dimension, all children
// are allocated the same size. Of course, the gtk.Widget:halign and
// gtk.Widget:valign properties can be used on the children to influence their
// allocation.
//
// Use repeated calls to gtk.Box.Append() to pack widgets into a GtkBox from
// start to end. Use gtk.Box.Remove() to remove widgets from the GtkBox.
// gtk.Box.InsertChildAfter() can be used to add a child at a particular
// position.
//
// Use gtk.Box.SetHomogeneous() to specify whether or not all children of the
// GtkBox are forced to get the same amount of space.
//
// Use gtk.Box.SetSpacing() to determine how much space will be minimally placed
// between all children in the GtkBox. Note that spacing is added *between* the
// children.
//
// Use gtk.Box.ReorderChildAfter() to move a child to a different place in the
// box.
//
//
// CSS nodes
//
// GtkBox uses a single CSS node with name box.
//
//
// Accessibility
//
// GtkBox uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Box struct {
	Widget

	Orientable
	*externglib.Object
}

func WrapBox(obj *externglib.Object) *Box {
	return &Box{
		Widget: Widget{
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
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBox(obj), nil
}

// NewBox creates a new GtkBox.
func NewBox(orientation Orientation, spacing int) *Box {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.int            // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.int(spacing)

	_cret = C.gtk_box_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(spacing)

	var _box *Box // out

	_box = WrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// Append adds child as the last child to box.
func (box *Box) Append(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_append(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// BaselinePosition gets the value set by gtk_box_set_baseline_position().
func (box *Box) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBox             // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_baseline_position(_arg0)
	runtime.KeepAlive(box)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// Homogeneous returns whether the box is homogeneous (all children are the same
// size).
func (box *Box) Homogeneous() bool {
	var _arg0 *C.GtkBox  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_homogeneous(_arg0)
	runtime.KeepAlive(box)

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
	runtime.KeepAlive(box)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertChildAfter inserts child in the position after sibling in the list of
// box children.
//
// If sibling is NULL, insert child at the first position.
func (box *Box) InsertChildAfter(child Widgetter, sibling Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	}

	C.gtk_box_insert_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// Prepend adds child as the first child to box.
func (box *Box) Prepend(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_prepend(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// Remove removes a child widget from box.
//
// The child must have been added before with gtk.Box.Append(),
// gtk.Box.Prepend(), or gtk.Box.InsertChildAfter().
func (box *Box) Remove(child Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_remove(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// box children.
//
// If sibling is NULL, move child to the first position.
func (box *Box) ReorderChildAfter(child Widgetter, sibling Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	}

	C.gtk_box_reorder_child_after(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
}

// SetBaselinePosition sets the baseline position of a box.
//
// This affects only horizontal boxes with at least one baseline aligned child.
// If there is more vertical space available than requested, and the baseline is
// not allocated by the parent then position is used to allocate the baseline
// with respect to the extra space available.
func (box *Box) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBox             // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_set_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(position)
}

// SetHomogeneous sets whether or not all children of box are given equal space
// in the box.
func (box *Box) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBox  // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(homogeneous)
}

// SetSpacing sets the number of pixels to place between children of box.
func (box *Box) SetSpacing(spacing int) {
	var _arg0 *C.GtkBox // out
	var _arg1 C.int     // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.int(spacing)

	C.gtk_box_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
