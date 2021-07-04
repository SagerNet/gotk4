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
		{T: externglib.Type(C.gtk_cell_renderer_pixbuf_get_type()), F: marshalCellRendererPixbuf},
	})
}

// CellRendererPixbuf: a CellRendererPixbuf can be used to render an image in a
// cell. It allows to render either a given Pixbuf (set via the
// CellRendererPixbuf:pixbuf property) or a named icon (set via the
// CellRendererPixbuf:icon-name property).
//
// To support the tree view, CellRendererPixbuf also supports rendering two
// alternative pixbufs, when the CellRenderer:is-expander property is true. If
// the CellRenderer:is-expanded property is true and the
// CellRendererPixbuf:pixbuf-expander-open property is set to a pixbuf, it
// renders that pixbuf, if the CellRenderer:is-expanded property is false and
// the CellRendererPixbuf:pixbuf-expander-closed property is set to a pixbuf, it
// renders that one.
type CellRendererPixbuf interface {
	CellRenderer
}

// cellRendererPixbuf implements the CellRendererPixbuf class.
type cellRendererPixbuf struct {
	CellRenderer
}

// WrapCellRendererPixbuf wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererPixbuf(obj *externglib.Object) CellRendererPixbuf {
	return cellRendererPixbuf{
		CellRenderer: WrapCellRenderer(obj),
	}
}

func marshalCellRendererPixbuf(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererPixbuf(obj), nil
}

// NewCellRendererPixbuf:
func NewCellRendererPixbuf() CellRendererPixbuf {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_pixbuf_new()

	var _cellRendererPixbuf CellRendererPixbuf // out

	_cellRendererPixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellRendererPixbuf)

	return _cellRendererPixbuf
}
