// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_renderer_text_get_type()), F: marshalCellRendererText},
	})
}

// CellRendererText renders text in a cell
//
// A CellRendererText renders a given text in its cell, using the font, color
// and style information provided by its properties. The text will be ellipsized
// if it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText interface {
	CellRenderer

	SetFixedHeightFromFontCellRendererText(numberOfRows int)
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

func NewCellRendererText() CellRendererText {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_text_new()

	var _cellRendererText CellRendererText // out

	_cellRendererText = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererText)

	return _cellRendererText
}

func (r cellRendererText) SetFixedHeightFromFontCellRendererText(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(r.Native()))
	_arg1 = C.int(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
}