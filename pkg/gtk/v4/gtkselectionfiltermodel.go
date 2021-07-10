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
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModeller},
	})
}

// SelectionFilterModeller describes SelectionFilterModel's methods.
type SelectionFilterModeller interface {
	gextras.Objector

	privateSelectionFilterModel()
}

// SelectionFilterModel: `GtkSelectionFilterModel` is a list model that presents
// the selection from a `GtkSelectionModel`.
type SelectionFilterModel struct {
	*externglib.Object
}

var _ SelectionFilterModeller = (*SelectionFilterModel)(nil)

func wrapSelectionFilterModeller(obj *externglib.Object) SelectionFilterModeller {
	return &SelectionFilterModel{
		Object: obj,
	}
}

func marshalSelectionFilterModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSelectionFilterModeller(obj), nil
}

func (*SelectionFilterModel) privateSelectionFilterModel() {}
