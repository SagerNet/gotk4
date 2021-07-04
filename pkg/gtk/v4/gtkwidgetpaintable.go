// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_widget_paintable_get_type()), F: marshalWidgetPaintable},
	})
}

// WidgetPaintable: `GtkWidgetPaintable` is a `GdkPaintable` that displays the
// contents of a widget.
//
// `GtkWidgetPaintable` will also take care of the widget not being in a state
// where it can be drawn (like when it isn't shown) and just draw nothing or
// where it does not have a size (like when it is hidden) and report no size in
// that case.
//
// Of course, `GtkWidgetPaintable` allows you to monitor widgets for size
// changes by emitting the [signal@Gdk.Paintable::invalidate-size] signal
// whenever the size of the widget changes as well as for visual changes by
// emitting the [signal@Gdk.Paintable::invalidate-contents] signal whenever the
// widget changes.
//
// You can use a `GtkWidgetPaintable` everywhere a `GdkPaintable` is allowed,
// including using it on a `GtkPicture` (or one of its parents) that it was set
// on itself via gtk_picture_set_paintable(). The paintable will take care of
// recursion when this happens. If you do this however, ensure that the
// [property@Gtk.Picture:can-shrink] property is set to true or you might end up
// with an infinitely growing widget.
type WidgetPaintable interface {
	gdk.Paintable

	// Widget:
	Widget() Widget
	// SetWidgetWidgetPaintable:
	SetWidgetWidgetPaintable(widget Widget)
}

// widgetPaintable implements the WidgetPaintable class.
type widgetPaintable struct {
	gextras.Objector
}

// WrapWidgetPaintable wraps a GObject to the right type. It is
// primarily used internally.
func WrapWidgetPaintable(obj *externglib.Object) WidgetPaintable {
	return widgetPaintable{
		Objector: obj,
	}
}

func marshalWidgetPaintable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWidgetPaintable(obj), nil
}

// NewWidgetPaintable:
func NewWidgetPaintable(widget Widget) WidgetPaintable {
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPaintable // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_widget_paintable_new(_arg1)

	var _widgetPaintable WidgetPaintable // out

	_widgetPaintable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(WidgetPaintable)

	return _widgetPaintable
}

func (s widgetPaintable) Widget() Widget {
	var _arg0 *C.GtkWidgetPaintable // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_widget_paintable_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s widgetPaintable) SetWidgetWidgetPaintable(widget Widget) {
	var _arg0 *C.GtkWidgetPaintable // out
	var _arg1 *C.GtkWidget          // out

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_widget_paintable_set_widget(_arg0, _arg1)
}

func (p widgetPaintable) ComputeConcreteSize(specifiedWidth float64, specifiedHeight float64, defaultWidth float64, defaultHeight float64) (concreteWidth float64, concreteHeight float64) {
	return gdk.WrapPaintable(gextras.InternObject(p)).ComputeConcreteSize(specifiedWidth, specifiedHeight, defaultWidth, defaultHeight)
}

func (p widgetPaintable) CurrentImage() gdk.Paintable {
	return gdk.WrapPaintable(gextras.InternObject(p)).CurrentImage()
}

func (p widgetPaintable) Flags() gdk.PaintableFlags {
	return gdk.WrapPaintable(gextras.InternObject(p)).Flags()
}

func (p widgetPaintable) IntrinsicAspectRatio() float64 {
	return gdk.WrapPaintable(gextras.InternObject(p)).IntrinsicAspectRatio()
}

func (p widgetPaintable) IntrinsicHeight() int {
	return gdk.WrapPaintable(gextras.InternObject(p)).IntrinsicHeight()
}

func (p widgetPaintable) IntrinsicWidth() int {
	return gdk.WrapPaintable(gextras.InternObject(p)).IntrinsicWidth()
}

func (p widgetPaintable) InvalidateContents() {
	gdk.WrapPaintable(gextras.InternObject(p)).InvalidateContents()
}

func (p widgetPaintable) InvalidateSize() {
	gdk.WrapPaintable(gextras.InternObject(p)).InvalidateSize()
}

func (p widgetPaintable) Snapshot(snapshot gdk.Snapshot, width float64, height float64) {
	gdk.WrapPaintable(gextras.InternObject(p)).Snapshot(snapshot, width, height)
}
