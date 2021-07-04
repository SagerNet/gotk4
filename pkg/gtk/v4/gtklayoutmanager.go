// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_layout_manager_get_type()), F: marshalLayoutManager},
	})
}

// LayoutManager: layout managers are delegate classes that handle the preferred
// size and the allocation of a widget.
//
// You typically subclass `GtkLayoutManager` if you want to implement a layout
// policy for the children of a widget, or if you want to determine the size of
// a widget depending on its contents.
//
// Each `GtkWidget` can only have a `GtkLayoutManager` instance associated to it
// at any given time; it is possible, though, to replace the layout manager
// instance using [method@Gtk.Widget.set_layout_manager].
//
//
// Layout properties
//
// A layout manager can expose properties for controlling the layout of each
// child, by creating an object type derived from [class@Gtk.LayoutChild] and
// installing the properties on it as normal `GObject` properties.
//
// Each `GtkLayoutChild` instance storing the layout properties for a specific
// child is created through the [method@Gtk.LayoutManager.get_layout_child]
// method; a `GtkLayoutManager` controls the creation of its `GtkLayoutChild`
// instances by overriding the GtkLayoutManagerClass.create_layout_child()
// virtual function. The typical implementation should look like:
//
// “`c static GtkLayoutChild * create_layout_child (GtkLayoutManager *manager,
// GtkWidget *container, GtkWidget *child) { return g_object_new
// (your_layout_child_get_type (), "layout-manager", manager, "child-widget",
// child, NULL); } “`
//
// The [property@Gtk.LayoutChild:layout-manager] and
// [property@Gtk.LayoutChild:child-widget] properties on the newly created
// `GtkLayoutChild` instance are mandatory. The `GtkLayoutManager` will cache
// the newly created `GtkLayoutChild` instance until the widget is removed from
// its parent, or the parent removes the layout manager.
//
// Each `GtkLayoutManager` instance creating a `GtkLayoutChild` should use
// [method@Gtk.LayoutManager.get_layout_child] every time it needs to query the
// layout properties; each `GtkLayoutChild` instance should call
// [method@Gtk.LayoutManager.layout_changed] every time a property is updated,
// in order to queue a new size measuring and allocation.
type LayoutManager interface {
	gextras.Objector

	// AllocateLayoutManager:
	AllocateLayoutManager(widget Widget, width int, height int, baseline int)
	// LayoutChild:
	LayoutChild(child Widget) LayoutChild
	// RequestMode:
	RequestMode() SizeRequestMode
	// Widget:
	Widget() Widget
	// LayoutChangedLayoutManager:
	LayoutChangedLayoutManager()
	// MeasureLayoutManager:
	MeasureLayoutManager(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int)
}

// layoutManager implements the LayoutManager class.
type layoutManager struct {
	gextras.Objector
}

// WrapLayoutManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapLayoutManager(obj *externglib.Object) LayoutManager {
	return layoutManager{
		Objector: obj,
	}
}

func marshalLayoutManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLayoutManager(obj), nil
}

func (m layoutManager) AllocateLayoutManager(widget Widget, width int, height int, baseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.int               // out
	var _arg3 C.int               // out
	var _arg4 C.int               // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(width)
	_arg3 = C.int(height)
	_arg4 = C.int(baseline)

	C.gtk_layout_manager_allocate(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (m layoutManager) LayoutChild(child Widget) LayoutChild {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.GtkLayoutChild   // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_layout_manager_get_layout_child(_arg0, _arg1)

	var _layoutChild LayoutChild // out

	_layoutChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(LayoutChild)

	return _layoutChild
}

func (m layoutManager) RequestMode() SizeRequestMode {
	var _arg0 *C.GtkLayoutManager  // out
	var _cret C.GtkSizeRequestMode // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_layout_manager_get_request_mode(_arg0)

	var _sizeRequestMode SizeRequestMode // out

	_sizeRequestMode = SizeRequestMode(_cret)

	return _sizeRequestMode
}

func (m layoutManager) Widget() Widget {
	var _arg0 *C.GtkLayoutManager // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_layout_manager_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (m layoutManager) LayoutChangedLayoutManager() {
	var _arg0 *C.GtkLayoutManager // out

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	C.gtk_layout_manager_layout_changed(_arg0)
}

func (m layoutManager) MeasureLayoutManager(widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int) {
	var _arg0 *C.GtkLayoutManager // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 C.GtkOrientation    // out
	var _arg3 C.int               // out
	var _arg4 C.int               // in
	var _arg5 C.int               // in
	var _arg6 C.int               // in
	var _arg7 C.int               // in

	_arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkOrientation(orientation)
	_arg3 = C.int(forSize)

	C.gtk_layout_manager_measure(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)

	var _minimum int         // out
	var _natural int         // out
	var _minimumBaseline int // out
	var _naturalBaseline int // out

	_minimum = int(_arg4)
	_natural = int(_arg5)
	_minimumBaseline = int(_arg6)
	_naturalBaseline = int(_arg7)

	return _minimum, _natural, _minimumBaseline, _naturalBaseline
}
