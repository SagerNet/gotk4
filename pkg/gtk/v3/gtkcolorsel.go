// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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

type ColorSelection interface {
	Box
	Buildable
	Orientable

	// CurrentAlpha returns the current alpha value.
	CurrentAlpha() uint16
	// CurrentColor sets @color to be the current color in the GtkColorSelection
	// widget.
	CurrentColor() gdk.Color
	// CurrentRGBA sets @rgba to be the current color in the GtkColorSelection
	// widget.
	CurrentRGBA() gdk.RGBA
	// HasOpacityControl determines whether the colorsel has an opacity control.
	HasOpacityControl() bool
	// HasPalette determines whether the color selector has a color palette.
	HasPalette() bool
	// PreviousAlpha returns the previous alpha value.
	PreviousAlpha() uint16
	// PreviousColor fills @color in with the original color value.
	PreviousColor() gdk.Color
	// PreviousRGBA fills @rgba in with the original color value.
	PreviousRGBA() gdk.RGBA
	// IsAdjusting gets the current state of the @colorsel.
	IsAdjusting() bool
	// SetCurrentAlpha sets the current opacity to be @alpha.
	//
	// The first time this is called, it will also set the original opacity to
	// be @alpha too.
	SetCurrentAlpha(alpha uint16)
	// SetCurrentColor sets the current color to be @color.
	//
	// The first time this is called, it will also set the original color to be
	// @color too.
	SetCurrentColor(color *gdk.Color)
	// SetCurrentRGBA sets the current color to be @rgba.
	//
	// The first time this is called, it will also set the original color to be
	// @rgba too.
	SetCurrentRGBA(rgba *gdk.RGBA)
	// SetHasOpacityControl sets the @colorsel to use or not use opacity.
	SetHasOpacityControl(hasOpacity bool)
	// SetHasPalette shows and hides the palette based upon the value of
	// @has_palette.
	SetHasPalette(hasPalette bool)
	// SetPreviousAlpha sets the “previous” alpha to be @alpha.
	//
	// This function should be called with some hesitations, as it might seem
	// confusing to have that alpha change.
	SetPreviousAlpha(alpha uint16)
	// SetPreviousColor sets the “previous” color to be @color.
	//
	// This function should be called with some hesitations, as it might seem
	// confusing to have that color change. Calling
	// gtk_color_selection_set_current_color() will also set this color the
	// first time it is called.
	SetPreviousColor(color *gdk.Color)
	// SetPreviousRGBA sets the “previous” color to be @rgba.
	//
	// This function should be called with some hesitations, as it might seem
	// confusing to have that color change. Calling
	// gtk_color_selection_set_current_rgba() will also set this color the first
	// time it is called.
	SetPreviousRGBA(rgba *gdk.RGBA)
}

// colorSelection implements the ColorSelection interface.
type colorSelection struct {
	Box
	Buildable
	Orientable
}

var _ ColorSelection = (*colorSelection)(nil)

// WrapColorSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorSelection(obj *externglib.Object) ColorSelection {
	return ColorSelection{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalColorSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorSelection(obj), nil
}

// NewColorSelection constructs a class ColorSelection.
func NewColorSelection() ColorSelection {
	var cret C.GtkColorSelection
	var goret1 ColorSelection

	cret = C.gtk_color_selection_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorSelection)

	return goret1
}

// CurrentAlpha returns the current alpha value.
func (c colorSelection) CurrentAlpha() uint16 {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var cret C.guint16
	var goret1 uint16

	cret = C.gtk_color_selection_get_current_alpha(arg0)

	goret1 = C.guint16(cret)

	return goret1
}

// CurrentColor sets @color to be the current color in the GtkColorSelection
// widget.
func (c colorSelection) CurrentColor() gdk.Color {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkColor
	var ret1 *gdk.Color

	C.gtk_color_selection_get_current_color(arg0, &arg1)

	ret1 = gdk.WrapColor(unsafe.Pointer(arg1))

	return ret1
}

// CurrentRGBA sets @rgba to be the current color in the GtkColorSelection
// widget.
func (c colorSelection) CurrentRGBA() gdk.RGBA {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkRGBA
	var ret1 *gdk.RGBA

	C.gtk_color_selection_get_current_rgba(arg0, &arg1)

	ret1 = gdk.WrapRGBA(unsafe.Pointer(arg1))

	return ret1
}

// HasOpacityControl determines whether the colorsel has an opacity control.
func (c colorSelection) HasOpacityControl() bool {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_color_selection_get_has_opacity_control(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// HasPalette determines whether the color selector has a color palette.
func (c colorSelection) HasPalette() bool {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_color_selection_get_has_palette(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// PreviousAlpha returns the previous alpha value.
func (c colorSelection) PreviousAlpha() uint16 {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var cret C.guint16
	var goret1 uint16

	cret = C.gtk_color_selection_get_previous_alpha(arg0)

	goret1 = C.guint16(cret)

	return goret1
}

// PreviousColor fills @color in with the original color value.
func (c colorSelection) PreviousColor() gdk.Color {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkColor
	var ret1 *gdk.Color

	C.gtk_color_selection_get_previous_color(arg0, &arg1)

	ret1 = gdk.WrapColor(unsafe.Pointer(arg1))

	return ret1
}

// PreviousRGBA fills @rgba in with the original color value.
func (c colorSelection) PreviousRGBA() gdk.RGBA {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkRGBA
	var ret1 *gdk.RGBA

	C.gtk_color_selection_get_previous_rgba(arg0, &arg1)

	ret1 = gdk.WrapRGBA(unsafe.Pointer(arg1))

	return ret1
}

// IsAdjusting gets the current state of the @colorsel.
func (c colorSelection) IsAdjusting() bool {
	var arg0 *C.GtkColorSelection

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_color_selection_is_adjusting(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetCurrentAlpha sets the current opacity to be @alpha.
//
// The first time this is called, it will also set the original opacity to
// be @alpha too.
func (c colorSelection) SetCurrentAlpha(alpha uint16) {
	var arg0 *C.GtkColorSelection
	var arg1 C.guint16

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_current_alpha(arg0, alpha)
}

// SetCurrentColor sets the current color to be @color.
//
// The first time this is called, it will also set the original color to be
// @color too.
func (c colorSelection) SetCurrentColor(color *gdk.Color) {
	var arg0 *C.GtkColorSelection
	var arg1 *C.GdkColor

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_selection_set_current_color(arg0, color)
}

// SetCurrentRGBA sets the current color to be @rgba.
//
// The first time this is called, it will also set the original color to be
// @rgba too.
func (c colorSelection) SetCurrentRGBA(rgba *gdk.RGBA) {
	var arg0 *C.GtkColorSelection
	var arg1 *C.GdkRGBA

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gtk_color_selection_set_current_rgba(arg0, rgba)
}

// SetHasOpacityControl sets the @colorsel to use or not use opacity.
func (c colorSelection) SetHasOpacityControl(hasOpacity bool) {
	var arg0 *C.GtkColorSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	if hasOpacity {
		arg1 = C.gboolean(1)
	}

	C.gtk_color_selection_set_has_opacity_control(arg0, hasOpacity)
}

// SetHasPalette shows and hides the palette based upon the value of
// @has_palette.
func (c colorSelection) SetHasPalette(hasPalette bool) {
	var arg0 *C.GtkColorSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	if hasPalette {
		arg1 = C.gboolean(1)
	}

	C.gtk_color_selection_set_has_palette(arg0, hasPalette)
}

// SetPreviousAlpha sets the “previous” alpha to be @alpha.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that alpha change.
func (c colorSelection) SetPreviousAlpha(alpha uint16) {
	var arg0 *C.GtkColorSelection
	var arg1 C.guint16

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = C.guint16(alpha)

	C.gtk_color_selection_set_previous_alpha(arg0, alpha)
}

// SetPreviousColor sets the “previous” color to be @color.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_color() will also set this color the
// first time it is called.
func (c colorSelection) SetPreviousColor(color *gdk.Color) {
	var arg0 *C.GtkColorSelection
	var arg1 *C.GdkColor

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_selection_set_previous_color(arg0, color)
}

// SetPreviousRGBA sets the “previous” color to be @rgba.
//
// This function should be called with some hesitations, as it might seem
// confusing to have that color change. Calling
// gtk_color_selection_set_current_rgba() will also set this color the first
// time it is called.
func (c colorSelection) SetPreviousRGBA(rgba *gdk.RGBA) {
	var arg0 *C.GtkColorSelection
	var arg1 *C.GdkRGBA

	arg0 = (*C.GtkColorSelection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	C.gtk_color_selection_set_previous_rgba(arg0, rgba)
}

type ColorSelectionPrivate struct {
	native C.GtkColorSelectionPrivate
}

// WrapColorSelectionPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColorSelectionPrivate(ptr unsafe.Pointer) *ColorSelectionPrivate {
	if ptr == nil {
		return nil
	}

	return (*ColorSelectionPrivate)(ptr)
}

func marshalColorSelectionPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColorSelectionPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *ColorSelectionPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
