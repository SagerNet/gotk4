// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_layout_get_type()), F: marshalLayouter},
	})
}

// Layouter describes Layout's methods.
type Layouter interface {
	gextras.Objector

	BinWindow() *gdk.Window
	HAdjustment() *Adjustment
	Size() (width uint, height uint)
	VAdjustment() *Adjustment
	Move(childWidget Widgetter, x int, y int)
	Put(childWidget Widgetter, x int, y int)
	SetHAdjustment(adjustment Adjustmenter)
	SetSize(width uint, height uint)
	SetVAdjustment(adjustment Adjustmenter)
}

// Layout is similar to DrawingArea in that it’s a “blank slate” and doesn’t do
// anything except paint a blank background by default. It’s different in that
// it supports scrolling natively due to implementing Scrollable, and can
// contain child widgets since it’s a Container.
//
// If you just want to draw, a DrawingArea is a better choice since it has lower
// overhead. If you just need to position child widgets at specific points, then
// Fixed provides that functionality on its own.
//
// When handling expose events on a Layout, you must draw to the Window returned
// by gtk_layout_get_bin_window(), rather than to the one returned by
// gtk_widget_get_window() as you would for a DrawingArea.
type Layout struct {
	*externglib.Object
	Container
	Buildable
	Scrollable
}

var _ Layouter = (*Layout)(nil)

func wrapLayouter(obj *externglib.Object) Layouter {
	return &Layout{
		Object: obj,
		Container: Container{
			Object: obj,
			Widget: Widget{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
	}
}

func marshalLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLayouter(obj), nil
}

// NewLayout creates a new Layout. Unless you have a specific adjustment you’d
// like the layout to use for scrolling, pass nil for @hadjustment and
// @vadjustment.
func NewLayout(hadjustment Adjustmenter, vadjustment Adjustmenter) *Layout {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_layout_new(_arg1, _arg2)

	var _layout *Layout // out

	_layout = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Layout)

	return _layout
}

// BinWindow: retrieve the bin window of the layout used for drawing operations.
func (layout *Layout) BinWindow() *gdk.Window {
	var _arg0 *C.GtkLayout // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))

	_cret = C.gtk_layout_get_bin_window(_arg0)

	var _window *gdk.Window // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Window)

	return _window
}

// HAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the horizontal scrollbar and
// @layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
func (layout *Layout) HAdjustment() *Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))

	_cret = C.gtk_layout_get_hadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// Size gets the size that has been set on the layout, and that determines the
// total extents of the layout’s scrollbar area. See gtk_layout_set_size ().
func (layout *Layout) Size() (width uint, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // in
	var _arg2 C.guint      // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_layout_get_size(_arg0, &_arg1, &_arg2)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

// VAdjustment: this function should only be called after the layout has been
// placed in a ScrolledWindow or otherwise configured for scrolling. It returns
// the Adjustment used for communication between the vertical scrollbar and
// @layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
func (layout *Layout) VAdjustment() *Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))

	_cret = C.gtk_layout_get_vadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// Move moves a current child of @layout to a new position.
func (layout *Layout) Move(childWidget Widgetter, x int, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_move(_arg0, _arg1, _arg2, _arg3)
}

// Put adds @child_widget to @layout, at position (@x,@y). @layout becomes the
// new parent container of @child_widget.
func (layout *Layout) Put(childWidget Widgetter, x int, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_put(_arg0, _arg1, _arg2, _arg3)
}

// SetHAdjustment sets the horizontal scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_hadjustment().
func (layout *Layout) SetHAdjustment(adjustment Adjustmenter) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_hadjustment(_arg0, _arg1)
}

// SetSize sets the size of the scrollable area of the layout.
func (layout *Layout) SetSize(width uint, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = C.guint(width)
	_arg2 = C.guint(height)

	C.gtk_layout_set_size(_arg0, _arg1, _arg2)
}

// SetVAdjustment sets the vertical scroll adjustment for the layout.
//
// See ScrolledWindow, Scrollbar, Adjustment for details.
//
// Deprecated: Use gtk_scrollable_set_vadjustment().
func (layout *Layout) SetVAdjustment(adjustment Adjustmenter) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(layout.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_vadjustment(_arg0, _arg1)
}
