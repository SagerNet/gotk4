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
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewport},
	})
}

// Viewport: the Viewport widget acts as an adaptor class, implementing
// scrollability for child widgets that lack their own scrolling capabilities.
// Use GtkViewport to scroll child widgets such as Grid, Box, and so on.
//
// If a widget has native scrolling abilities, such as TextView, TreeView or
// IconView, it can be added to a ScrolledWindow with gtk_container_add(). If a
// widget does not, you must first add the widget to a Viewport, then add the
// viewport to the scrolled window. gtk_container_add() does this automatically
// if a child that does not implement Scrollable is added to a ScrolledWindow,
// so you can ignore the presence of the viewport.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
type Viewport interface {
	Bin

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsScrollable casts the class to the Scrollable interface.
	AsScrollable() Scrollable

	// BinWindow gets the bin window of the Viewport.
	BinWindow() gdk.Window
	// HAdjustment returns the horizontal adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	HAdjustment() Adjustment
	// ShadowType gets the shadow type of the Viewport. See
	// gtk_viewport_set_shadow_type().
	ShadowType() ShadowType
	// VAdjustment returns the vertical adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	VAdjustment() Adjustment
	// ViewWindow gets the view window of the Viewport.
	ViewWindow() gdk.Window
	// SetHAdjustment sets the horizontal adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	SetHAdjustment(adjustment Adjustment)
	// SetShadowType sets the shadow type of the viewport.
	SetShadowType(typ ShadowType)
	// SetVAdjustment sets the vertical adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	SetVAdjustment(adjustment Adjustment)
}

// viewport implements the Viewport class.
type viewport struct {
	Bin
}

// WrapViewport wraps a GObject to the right type. It is
// primarily used internally.
func WrapViewport(obj *externglib.Object) Viewport {
	return viewport{
		Bin: WrapBin(obj),
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapViewport(obj), nil
}

// NewViewport creates a new Viewport with the given adjustments, or with
// default adjustments if none are given.
func NewViewport(hadjustment Adjustment, vadjustment Adjustment) Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport Viewport // out

	_viewport = WrapViewport(externglib.Take(unsafe.Pointer(_cret)))

	return _viewport
}

func (v viewport) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(v))
}

func (v viewport) AsScrollable() Scrollable {
	return WrapScrollable(gextras.InternObject(v))
}

func (v viewport) BinWindow() gdk.Window {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_bin_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (v viewport) HAdjustment() Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (v viewport) ShadowType() ShadowType {
	var _arg0 *C.GtkViewport  // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

func (v viewport) VAdjustment() Adjustment {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (v viewport) ViewWindow() gdk.Window {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_view_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (v viewport) SetHAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_hadjustment(_arg0, _arg1)
}

func (v viewport) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkViewport  // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_viewport_set_shadow_type(_arg0, _arg1)
}

func (v viewport) SetVAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_vadjustment(_arg0, _arg1)
}
