// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gint _gotk4_glib2_CompareDataFunc(gconstpointer, gconstpointer, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_custom_sorter_get_type()), F: marshalCustomSorterer},
	})
}

// CustomSorter: GtkCustomSorter is a GtkSorter implementation that sorts via a
// callback function.
type CustomSorter struct {
	Sorter
}

func WrapCustomSorter(obj *externglib.Object) *CustomSorter {
	return &CustomSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalCustomSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCustomSorter(obj), nil
}

// NewCustomSorter creates a new GtkSorter that works by calling sort_func to
// compare items.
//
// If sort_func is NULL, all items are considered equal.
func NewCustomSorter(sortFunc glib.CompareDataFunc) *CustomSorter {
	var _arg1 C.GCompareDataFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify
	var _cret *C.GtkCustomSorter // in

	if sortFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
		_arg2 = C.gpointer(gbox.Assign(sortFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_custom_sorter_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(sortFunc)

	var _customSorter *CustomSorter // out

	_customSorter = WrapCustomSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _customSorter
}

// SetSortFunc sets (or unsets) the function used for sorting items.
//
// If sort_func is NULL, all items are considered equal.
//
// If the sort func changes its sorting behavior, gtk_sorter_changed() needs to
// be called.
//
// If a previous function was set, its user_destroy will be called now.
func (self *CustomSorter) SetSortFunc(sortFunc glib.CompareDataFunc) {
	var _arg0 *C.GtkCustomSorter // out
	var _arg1 C.GCompareDataFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkCustomSorter)(unsafe.Pointer(self.Native()))
	if sortFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_glib2_CompareDataFunc)
		_arg2 = C.gpointer(gbox.Assign(sortFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_custom_sorter_set_sort_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sortFunc)
}
