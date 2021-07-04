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
		{T: externglib.Type(C.gtk_overlay_get_type()), F: marshalOverlay},
	})
}

// Overlay: `GtkOverlay` is a container which contains a single main child, on
// top of which it can place “overlay” widgets.
//
// !An example GtkOverlay (overlay.png)
//
// The position of each overlay widget is determined by its
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties.
// E.g. a widget with both alignments set to GTK_ALIGN_START will be placed at
// the top left corner of the `GtkOverlay` container, whereas an overlay with
// halign set to GTK_ALIGN_CENTER and valign set to GTK_ALIGN_END will be placed
// a the bottom edge of the `GtkOverlay`, horizontally centered. The position
// can be adjusted by setting the margin properties of the child to non-zero
// values.
//
// More complicated placement of overlays is possible by connecting to the
// [signal@Gtk.Overlay::get-child-position] signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The `GtkOverlay` implementation of the `GtkBuildable` interface supports
// placing a child as an overlay by specifying “overlay” as the “type” attribute
// of a `<child>` element.
//
//
// CSS nodes
//
// `GtkOverlay` has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay interface {
	Widget

	// AddOverlayOverlay:
	AddOverlayOverlay(widget Widget)
	// Child:
	Child() Widget
	// ClipOverlay:
	ClipOverlay(widget Widget) bool
	// MeasureOverlay:
	MeasureOverlay(widget Widget) bool
	// RemoveOverlayOverlay:
	RemoveOverlayOverlay(widget Widget)
	// SetChildOverlay:
	SetChildOverlay(child Widget)
	// SetClipOverlayOverlay:
	SetClipOverlayOverlay(widget Widget, clipOverlay bool)
	// SetMeasureOverlayOverlay:
	SetMeasureOverlayOverlay(widget Widget, measure bool)
}

// overlay implements the Overlay class.
type overlay struct {
	Widget
}

// WrapOverlay wraps a GObject to the right type. It is
// primarily used internally.
func WrapOverlay(obj *externglib.Object) Overlay {
	return overlay{
		Widget: WrapWidget(obj),
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOverlay(obj), nil
}

// NewOverlay:
func NewOverlay() Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay Overlay // out

	_overlay = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Overlay)

	return _overlay
}

func (o overlay) AddOverlayOverlay(widget Widget) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
}

func (o overlay) Child() Widget {
	var _arg0 *C.GtkOverlay // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_overlay_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (o overlay) ClipOverlay(widget Widget) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_clip_overlay(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o overlay) MeasureOverlay(widget Widget) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_measure_overlay(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o overlay) RemoveOverlayOverlay(widget Widget) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_remove_overlay(_arg0, _arg1)
}

func (o overlay) SetChildOverlay(child Widget) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_overlay_set_child(_arg0, _arg1)
}

func (o overlay) SetClipOverlayOverlay(widget Widget, clipOverlay bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if clipOverlay {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_clip_overlay(_arg0, _arg1, _arg2)
}

func (o overlay) SetMeasureOverlayOverlay(widget Widget, measure bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if measure {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_measure_overlay(_arg0, _arg1, _arg2)
}

func (s overlay) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s overlay) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s overlay) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s overlay) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s overlay) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s overlay) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s overlay) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b overlay) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
