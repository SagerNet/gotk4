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
		{T: externglib.Type(C.gtk_layout_child_get_type()), F: marshalLayoutChilder},
	})
}

// LayoutChilder describes LayoutChild's methods.
type LayoutChilder interface {
	gextras.Objector

	ChildWidget() *Widget
	LayoutManager() *LayoutManager
}

// LayoutChild: `GtkLayoutChild` is the base class for objects that are meant to
// hold layout properties.
//
// If a `GtkLayoutManager` has per-child properties, like their packing type, or
// the horizontal and vertical span, or the icon name, then the layout manager
// should use a `GtkLayoutChild` implementation to store those properties.
//
// A `GtkLayoutChild` instance is only ever valid while a widget is part of a
// layout.
type LayoutChild struct {
	*externglib.Object
}

var _ LayoutChilder = (*LayoutChild)(nil)

func wrapLayoutChilder(obj *externglib.Object) LayoutChilder {
	return &LayoutChild{
		Object: obj,
	}
}

func marshalLayoutChilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLayoutChilder(obj), nil
}

// ChildWidget retrieves the `GtkWidget` associated to the given @layout_child.
func (layoutChild *LayoutChild) ChildWidget() *Widget {
	var _arg0 *C.GtkLayoutChild // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(layoutChild.Native()))

	_cret = C.gtk_layout_child_get_child_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// LayoutManager retrieves the `GtkLayoutManager` instance that created the
// given @layout_child.
func (layoutChild *LayoutChild) LayoutManager() *LayoutManager {
	var _arg0 *C.GtkLayoutChild   // out
	var _cret *C.GtkLayoutManager // in

	_arg0 = (*C.GtkLayoutChild)(unsafe.Pointer(layoutChild.Native()))

	_cret = C.gtk_layout_child_get_layout_manager(_arg0)

	var _layoutManager *LayoutManager // out

	_layoutManager = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LayoutManager)

	return _layoutManager
}
