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
		{T: externglib.Type(C.gtk_cell_renderer_text_get_type()), F: marshalCellRendererText},
	})
}

// CellRendererText renders a given text in its cell, using the font, color and
// style information provided by its properties. The text will be ellipsized if
// it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText interface {
	CellRenderer

	// SetFixedHeightFromFont sets the height of a renderer to explicitly be
	// determined by the “font” and “y_pad” property set on it. Further changes
	// in these properties do not affect the height, so they must be accompanied
	// by a subsequent call to this function. Using this function is unflexible,
	// and should really only be used if calculating the size of a cell is too
	// slow (ie, a massive number of cells displayed). If @number_of_rows is -1,
	// then the fixed height is unset, and the height is determined by the
	// properties again.
	SetFixedHeightFromFont(numberOfRows int)
}

// cellRendererText implements the CellRendererText class.
type cellRendererText struct {
	CellRenderer
}

// WrapCellRendererText wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererText(obj *externglib.Object) CellRendererText {
	return cellRendererText{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererText(obj), nil
}

// NewCellRendererText creates a new CellRendererText. Adjust how text is drawn
// using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView
func NewCellRendererText() CellRendererText {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_text_new()

	var _cellRendererText CellRendererText // out

	_cellRendererText = WrapCellRendererText(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererText
}

func (r cellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText // out
	var _arg1 C.gint                 // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
}
