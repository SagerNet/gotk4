// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_layout_get_type()), F: marshalLayout},
	})
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
type Layout interface {
	Container
	Scrollable

	// BinWindow:
	BinWindow() gdk.Window
	// GetHAdjustment:
	GetHAdjustment() Adjustment
	// Size:
	Size() (width uint, height uint)
	// GetVAdjustment:
	GetVAdjustment() Adjustment
	// MoveLayout:
	MoveLayout(childWidget Widget, x int, y int)
	// PutLayout:
	PutLayout(childWidget Widget, x int, y int)
	// SetHAdjustmentLayout:
	SetHAdjustmentLayout(adjustment Adjustment)
	// SetSizeLayout:
	SetSizeLayout(width uint, height uint)
	// SetVAdjustmentLayout:
	SetVAdjustmentLayout(adjustment Adjustment)
}

// layout implements the Layout class.
type layout struct {
	Container
}

// WrapLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapLayout(obj *externglib.Object) Layout {
	return layout{
		Container: WrapContainer(obj),
	}
}

func marshalLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLayout(obj), nil
}

// NewLayout:
func NewLayout(hadjustment Adjustment, vadjustment Adjustment) Layout {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_layout_new(_arg1, _arg2)

	var _layout Layout // out

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Layout)

	return _layout
}

func (l layout) BinWindow() gdk.Window {
	var _arg0 *C.GtkLayout // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_layout_get_bin_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (l layout) GetHAdjustment() Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_layout_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (l layout) Size() (width uint, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // in
	var _arg2 C.guint      // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	C.gtk_layout_get_size(_arg0, &_arg1, &_arg2)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

func (l layout) GetVAdjustment() Adjustment {
	var _arg0 *C.GtkLayout     // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_layout_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (l layout) MoveLayout(childWidget Widget, x int, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_move(_arg0, _arg1, _arg2, _arg3)
}

func (l layout) PutLayout(childWidget Widget, x int, y int) {
	var _arg0 *C.GtkLayout // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(childWidget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_layout_put(_arg0, _arg1, _arg2, _arg3)
}

func (l layout) SetHAdjustmentLayout(adjustment Adjustment) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_hadjustment(_arg0, _arg1)
}

func (l layout) SetSizeLayout(width uint, height uint) {
	var _arg0 *C.GtkLayout // out
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint(width)
	_arg2 = C.guint(height)

	C.gtk_layout_set_size(_arg0, _arg1, _arg2)
}

func (l layout) SetVAdjustmentLayout(adjustment Adjustment) {
	var _arg0 *C.GtkLayout     // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_layout_set_vadjustment(_arg0, _arg1)
}

func (b layout) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b layout) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b layout) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b layout) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b layout) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b layout) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b layout) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (s layout) Border() (Border, bool) {
	return WrapScrollable(gextras.InternObject(s)).Border()
}

func (s layout) HAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).HAdjustment()
}

func (s layout) HScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).HScrollPolicy()
}

func (s layout) VAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).VAdjustment()
}

func (s layout) VScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).VScrollPolicy()
}

func (s layout) SetHAdjustment(hadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetHAdjustment(hadjustment)
}

func (s layout) SetHScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetHScrollPolicy(policy)
}

func (s layout) SetVAdjustment(vadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetVAdjustment(vadjustment)
}

func (s layout) SetVScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetVScrollPolicy(policy)
}
