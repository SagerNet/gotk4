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
		{T: externglib.Type(C.gtk_overlay_get_type()), F: marshalOverlayyer},
	})
}

// Overlayyer describes Overlay's methods.
type Overlayyer interface {
	gextras.Objector

	AddOverlay(widget Widgetter)
	Child() *Widget
	ClipOverlay(widget Widgetter) bool
	MeasureOverlay(widget Widgetter) bool
	RemoveOverlay(widget Widgetter)
	SetChild(child Widgetter)
	SetClipOverlay(widget Widgetter, clipOverlay bool)
	SetMeasureOverlay(widget Widgetter, measure bool)
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
type Overlay struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Overlayyer = (*Overlay)(nil)

func wrapOverlayyer(obj *externglib.Object) Overlayyer {
	return &Overlay{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
	}
}

func marshalOverlayyer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOverlayyer(obj), nil
}

// NewOverlay creates a new `GtkOverlay`.
func NewOverlay() *Overlay {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay *Overlay // out

	_overlay = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Overlay)

	return _overlay
}

// AddOverlay adds @widget to @overlay.
//
// The widget will be stacked on top of the main widget added with
// [method@Gtk.Overlay.set_child].
//
// The position at which @widget is placed is determined from its
// [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign] properties.
func (overlay *Overlay) AddOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
}

// Child gets the child widget of @overlay.
func (overlay *Overlay) Child() *Widget {
	var _arg0 *C.GtkOverlay // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))

	_cret = C.gtk_overlay_get_child(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// ClipOverlay gets whether @widget should be clipped within the parent.
func (overlay *Overlay) ClipOverlay(widget Widgetter) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_clip_overlay(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MeasureOverlay gets whether @widget's size is included in the measurement of
// @overlay.
func (overlay *Overlay) MeasureOverlay(widget Widgetter) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_measure_overlay(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RemoveOverlay removes an overlay that was added with
// gtk_overlay_add_overlay().
func (overlay *Overlay) RemoveOverlay(widget Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_remove_overlay(_arg0, _arg1)
}

// SetChild sets the child widget of @overlay.
func (overlay *Overlay) SetChild(child Widgetter) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_overlay_set_child(_arg0, _arg1)
}

// SetClipOverlay sets whether @widget should be clipped within the parent.
func (overlay *Overlay) SetClipOverlay(widget Widgetter, clipOverlay bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if clipOverlay {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_clip_overlay(_arg0, _arg1, _arg2)
}

// SetMeasureOverlay sets whether @widget is included in the measured size of
// @overlay.
//
// The overlay will request the size of the largest child that has this property
// set to true. Children who are not included may be drawn outside of @overlay's
// allocation if they are too large.
func (overlay *Overlay) SetMeasureOverlay(widget Widgetter, measure bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(overlay.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if measure {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_measure_overlay(_arg0, _arg1, _arg2)
}
