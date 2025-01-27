// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_cell_renderer_pixbuf_get_type()), F: marshalCellRendererPixbuffer},
	})
}

// CellRendererPixbuf renders a pixbuf in a cell
//
// A CellRendererPixbuf can be used to render an image in a cell. It allows to
// render either a given Pixbuf (set via the CellRendererPixbuf:pixbuf property)
// or a named icon (set via the CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is TRUE. If
// the CellRenderer:is-expanded property is TRUE and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is FALSE and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf struct {
	CellRenderer
}

func WrapCellRendererPixbuf(obj *externglib.Object) *CellRendererPixbuf {
	return &CellRendererPixbuf{
		CellRenderer: CellRenderer{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererPixbuffer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererPixbuf(obj), nil
}

// NewCellRendererPixbuf creates a new CellRendererPixbuf. Adjust rendering
// parameters using object properties. Object properties can be set globally
// (with g_object_set()). Also, with TreeViewColumn, you can bind a property to
// a value in a TreeModel. For example, you can bind the “pixbuf” property on
// the cell renderer to a pixbuf value in the model, thus rendering a different
// image in each row of the TreeView.
func NewCellRendererPixbuf() *CellRendererPixbuf {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_pixbuf_new()

	var _cellRendererPixbuf *CellRendererPixbuf // out

	_cellRendererPixbuf = WrapCellRendererPixbuf(externglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererPixbuf
}

func (*CellRendererPixbuf) privateCellRendererPixbuf() {}
