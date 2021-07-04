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
		{T: externglib.Type(C.gtk_color_selection_get_type()), F: marshalColorSelection},
	})
}

// ColorSelection:
type ColorSelection interface {
	Box

	// CurrentAlpha:
	CurrentAlpha() uint16
	// CurrentColor:
	CurrentColor() gdk.Color
	// CurrentRGBA:
	CurrentRGBA() gdk.RGBA
	// HasOpacityControl:
	HasOpacityControl() bool
	// HasPalette:
	HasPalette() bool
	// PreviousAlpha:
	PreviousAlpha() uint16
	// PreviousColor:
	PreviousColor() gdk.Color
	// PreviousRGBA:
	PreviousRGBA() gdk.RGBA
	// IsAdjustingColorSelection:
	IsAdjustingColorSelection() bool
	// SetCurrentAlphaColorSelection:
	SetCurrentAlphaColorSelection(alpha uint16)
	// SetCurrentColorColorSelection:
	SetCurrentColorColorSelection(color *gdk.Color)
	// SetCurrentRGBAColorSelection:
	SetCurrentRGBAColorSelection(rgba *gdk.RGBA)
	// SetHasOpacityControlColorSelection:
	SetHasOpacityControlColorSelection(hasOpacity bool)
	// SetHasPaletteColorSelection:
	SetHasPaletteColorSelection(hasPalette bool)
	// SetPreviousAlphaColorSelection:
	SetPreviousAlphaColorSelection(alpha uint16)
	// SetPreviousColorColorSelection:
	SetPreviousColorColorSelection(color *gdk.Color)
	// SetPreviousRGBAColorSelection:
	SetPreviousRGBAColorSelection(rgba *gdk.RGBA)
}

// colorSelection implements the ColorSelection class.
type colorSelection struct {
	Box
}

// WrapColorSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorSelection(obj *externglib.Object) ColorSelection {
	return colorSelection{
		Box: WrapBox(obj),
	}
}

func marshalColorSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorSelection(obj), nil
}

// NewColorSelection:
func NewColorSelection() ColorSelection {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_selection_new()

	var _colorSelection ColorSelection // out

	_colorSelection = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorSelection)

	return _colorSelection
}

func (c colorSelection) CurrentAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_get_current_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (c colorSelection) CurrentColor() gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	C.gtk_color_selection_get_current_color(_arg0, &_arg1)

	var _color gdk.Color // out

	{
		var refTmpIn *C.GdkColor
		var refTmpOut *gdk.Color

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Color)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (c colorSelection) CurrentRGBA() gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	C.gtk_color_selection_get_current_rgba(_arg0, &_arg1)

	var _rgba gdk.RGBA // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_rgba = *refTmpOut
	}

	return _rgba
}

func (c colorSelection) HasOpacityControl() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_get_has_opacity_control(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c colorSelection) HasPalette() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_get_has_palette(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c colorSelection) PreviousAlpha() uint16 {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.guint16            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_get_previous_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (c colorSelection) PreviousColor() gdk.Color {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkColor           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	C.gtk_color_selection_get_previous_color(_arg0, &_arg1)

	var _color gdk.Color // out

	{
		var refTmpIn *C.GdkColor
		var refTmpOut *gdk.Color

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Color)(unsafe.Pointer(refTmpIn))

		_color = *refTmpOut
	}

	return _color
}

func (c colorSelection) PreviousRGBA() gdk.RGBA {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.GdkRGBA            // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	C.gtk_color_selection_get_previous_rgba(_arg0, &_arg1)

	var _rgba gdk.RGBA // out

	{
		var refTmpIn *C.GdkRGBA
		var refTmpOut *gdk.RGBA

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.RGBA)(unsafe.Pointer(refTmpIn))

		_rgba = *refTmpOut
	}

	return _rgba
}

func (c colorSelection) IsAdjustingColorSelection() bool {
	var _arg0 *C.GtkColorSelection // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_color_selection_is_adjusting(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c colorSelection) SetCurrentAlphaColorSelection(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_current_alpha(_arg0, _arg1)
}

func (c colorSelection) SetCurrentColorColorSelection(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_selection_set_current_color(_arg0, _arg1)
}

func (c colorSelection) SetCurrentRGBAColorSelection(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gtk_color_selection_set_current_rgba(_arg0, _arg1)
}

func (c colorSelection) SetHasOpacityControlColorSelection(hasOpacity bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	if hasOpacity {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_opacity_control(_arg0, _arg1)
}

func (c colorSelection) SetHasPaletteColorSelection(hasPalette bool) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	if hasPalette {
		_arg1 = C.TRUE
	}

	C.gtk_color_selection_set_has_palette(_arg0, _arg1)
}

func (c colorSelection) SetPreviousAlphaColorSelection(alpha uint16) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 C.guint16            // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_previous_alpha(_arg0, _arg1)
}

func (c colorSelection) SetPreviousColorColorSelection(color *gdk.Color) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkColor          // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_selection_set_previous_color(_arg0, _arg1)
}

func (c colorSelection) SetPreviousRGBAColorSelection(rgba *gdk.RGBA) {
	var _arg0 *C.GtkColorSelection // out
	var _arg1 *C.GdkRGBA           // out

	_arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gtk_color_selection_set_previous_rgba(_arg0, _arg1)
}

func (b colorSelection) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b colorSelection) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b colorSelection) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b colorSelection) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b colorSelection) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b colorSelection) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b colorSelection) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o colorSelection) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o colorSelection) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
