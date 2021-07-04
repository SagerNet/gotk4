// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_editable_get_type()), F: marshalCellEditable},
	})
}

// CellEditable: the CellEditable interface must be implemented for widgets to
// be usable to edit the contents of a TreeView cell. It provides a way to
// specify how temporary widgets should be configured for editing, get the new
// value, etc.
type CellEditable interface {
	Widget

	// EditingDone begins editing on a @cell_editable.
	//
	// The CellRenderer for the cell creates and returns a CellEditable from
	// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
	//
	// gtk_cell_editable_start_editing() can then set up @cell_editable suitably
	// for editing a cell, e.g. making the Esc key emit
	// CellEditable::editing-done.
	//
	// Note that the @cell_editable is created on-demand for the current edit;
	// its lifetime is temporary and does not persist across other edits and/or
	// cells.
	EditingDone()
	// RemoveWidget begins editing on a @cell_editable.
	//
	// The CellRenderer for the cell creates and returns a CellEditable from
	// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
	//
	// gtk_cell_editable_start_editing() can then set up @cell_editable suitably
	// for editing a cell, e.g. making the Esc key emit
	// CellEditable::editing-done.
	//
	// Note that the @cell_editable is created on-demand for the current edit;
	// its lifetime is temporary and does not persist across other edits and/or
	// cells.
	RemoveWidget()
}

// cellEditable implements the CellEditable interface.
type cellEditable struct {
	Widget
}

var _ CellEditable = (*cellEditable)(nil)

// WrapCellEditable wraps a GObject to a type that implements
// interface CellEditable. It is primarily used internally.
func WrapCellEditable(obj *externglib.Object) CellEditable {
	return cellEditable{
		Widget: WrapWidget(obj),
	}
}

func marshalCellEditable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellEditable(obj), nil
}

func (c cellEditable) EditingDone() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(c.Native()))

	C.gtk_cell_editable_editing_done(_arg0)
}

func (c cellEditable) RemoveWidget() {
	var _arg0 *C.GtkCellEditable // out

	_arg0 = (*C.GtkCellEditable)(unsafe.Pointer(c.Native()))

	C.gtk_cell_editable_remove_widget(_arg0)
}