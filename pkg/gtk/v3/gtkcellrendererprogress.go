// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_progress_get_type()), F: marshalCellRendererProgress},
	})
}

// CellRendererProgress renders a numeric value as a progress par in a cell.
// Additionally, it can display a text on top of the progress bar.
//
// The CellRendererProgress cell renderer was added in GTK+ 2.6.
type CellRendererProgress interface {
	CellRenderer
	Orientable
}

// cellRendererProgress implements the CellRendererProgress interface.
type cellRendererProgress struct {
	CellRenderer
	Orientable
}

var _ CellRendererProgress = (*cellRendererProgress)(nil)

// WrapCellRendererProgress wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellRendererProgress(obj *externglib.Object) CellRendererProgress {
	return CellRendererProgress{
		CellRenderer: WrapCellRenderer(obj),
		Orientable:   WrapOrientable(obj),
	}
}

func marshalCellRendererProgress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellRendererProgress(obj), nil
}

// NewCellRendererProgress constructs a class CellRendererProgress.
func NewCellRendererProgress() CellRendererProgress {
	var cret C.GtkCellRendererProgress
	var ret1 CellRendererProgress

	cret = C.gtk_cell_renderer_progress_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CellRendererProgress)

	return ret1
}

type CellRendererProgressPrivate struct {
	native C.GtkCellRendererProgressPrivate
}

// WrapCellRendererProgressPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCellRendererProgressPrivate(ptr unsafe.Pointer) *CellRendererProgressPrivate {
	if ptr == nil {
		return nil
	}

	return (*CellRendererProgressPrivate)(ptr)
}

func marshalCellRendererProgressPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCellRendererProgressPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CellRendererProgressPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}