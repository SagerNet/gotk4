// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_color_button_get_type()), F: marshalColorButton},
	})
}

// ColorButton: the ColorButton is a button which displays the currently
// selected color and allows to open a color selection dialog to change the
// color. It is suitable widget for selecting a color in a preference dialog.
//
//
// CSS nodes
//
// GtkColorButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .color style class.
type ColorButton interface {
	Button
	ColorChooser

	// Alpha:
	Alpha() uint16
	// Color:
	Color() gdk.Color
	// Title:
	Title() string
	// GetUseAlpha:
	GetUseAlpha() bool
	// SetAlphaColorButton:
	SetAlphaColorButton(alpha uint16)
	// SetColorColorButton:
	SetColorColorButton(color *gdk.Color)
	// SetTitleColorButton:
	SetTitleColorButton(title string)
	// SetUseAlphaColorButton:
	SetUseAlphaColorButton(useAlpha bool)
}

// colorButton implements the ColorButton class.
type colorButton struct {
	Button
}

// WrapColorButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorButton(obj *externglib.Object) ColorButton {
	return colorButton{
		Button: WrapButton(obj),
	}
}

func marshalColorButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorButton(obj), nil
}

// NewColorButton:
func NewColorButton() ColorButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_color_button_new()

	var _colorButton ColorButton // out

	_colorButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorButton)

	return _colorButton
}

// NewColorButtonWithColor:
func NewColorButtonWithColor(color *gdk.Color) ColorButton {
	var _arg1 *C.GdkColor  // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	_cret = C.gtk_color_button_new_with_color(_arg1)

	var _colorButton ColorButton // out

	_colorButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorButton)

	return _colorButton
}

// NewColorButtonWithRGBA:
func NewColorButtonWithRGBA(rgba *gdk.RGBA) ColorButton {
	var _arg1 *C.GdkRGBA   // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba.Native()))

	_cret = C.gtk_color_button_new_with_rgba(_arg1)

	var _colorButton ColorButton // out

	_colorButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ColorButton)

	return _colorButton
}

func (b colorButton) Alpha() uint16 {
	var _arg0 *C.GtkColorButton // out
	var _cret C.guint16         // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_alpha(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (b colorButton) Color() gdk.Color {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.GdkColor        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	C.gtk_color_button_get_color(_arg0, &_arg1)

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

func (b colorButton) Title() string {
	var _arg0 *C.GtkColorButton // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b colorButton) GetUseAlpha() bool {
	var _arg0 *C.GtkColorButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_color_button_get_use_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b colorButton) SetAlphaColorButton(alpha uint16) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.guint16         // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint16(alpha)

	C.gtk_color_button_set_alpha(_arg0, _arg1)
}

func (b colorButton) SetColorColorButton(color *gdk.Color) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.GdkColor       // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(color.Native()))

	C.gtk_color_button_set_color(_arg0, _arg1)
}

func (b colorButton) SetTitleColorButton(title string) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_color_button_set_title(_arg0, _arg1)
}

func (b colorButton) SetUseAlphaColorButton(useAlpha bool) {
	var _arg0 *C.GtkColorButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkColorButton)(unsafe.Pointer(b.Native()))
	if useAlpha {
		_arg1 = C.TRUE
	}

	C.gtk_color_button_set_use_alpha(_arg0, _arg1)
}

func (b colorButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b colorButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b colorButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b colorButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b colorButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b colorButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b colorButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a colorButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a colorButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a colorButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a colorButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a colorButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b colorButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b colorButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b colorButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b colorButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b colorButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b colorButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b colorButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a colorButton) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a colorButton) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a colorButton) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a colorButton) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a colorButton) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a colorButton) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}

func (c colorButton) AddPalette(orientation Orientation, colorsPerLine int, colors []gdk.RGBA) {
	WrapColorChooser(gextras.InternObject(c)).AddPalette(orientation, colorsPerLine, colors)
}

func (c colorButton) RGBA() gdk.RGBA {
	return WrapColorChooser(gextras.InternObject(c)).RGBA()
}

func (c colorButton) UseAlpha() bool {
	return WrapColorChooser(gextras.InternObject(c)).UseAlpha()
}

func (c colorButton) SetRGBA(color *gdk.RGBA) {
	WrapColorChooser(gextras.InternObject(c)).SetRGBA(color)
}

func (c colorButton) SetUseAlpha(useAlpha bool) {
	WrapColorChooser(gextras.InternObject(c)).SetUseAlpha(useAlpha)
}
