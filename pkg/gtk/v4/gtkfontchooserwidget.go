// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidget},
	})
}

// FontChooserWidget: the `GtkFontChooserWidget` widget lets the user select a
// font.
//
// It is used in the `GtkFontChooserDialog` widget to provide a dialog for
// selecting fonts.
//
// To set the font which is initially selected, use
// [method@Gtk.FontChooser.set_font] or [method@Gtk.FontChooser.set_font_desc].
//
// To get the selected font use [method@Gtk.FontChooser.get_font] or
// [method@Gtk.FontChooser.get_font_desc].
//
// To change the text which is shown in the preview area, use
// [method@Gtk.FontChooser.set_preview_text].
//
//
// CSS nodes
//
// `GtkFontChooserWidget` has a single CSS node with name fontchooser.
type FontChooserWidget interface {
	Widget
	FontChooser
}

// fontChooserWidget implements the FontChooserWidget class.
type fontChooserWidget struct {
	Widget
}

// WrapFontChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontChooserWidget(obj *externglib.Object) FontChooserWidget {
	return fontChooserWidget{
		Widget: WrapWidget(obj),
	}
}

func marshalFontChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooserWidget(obj), nil
}

func NewFontChooserWidget() FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget FontChooserWidget // out

	_fontChooserWidget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontChooserWidget)

	return _fontChooserWidget
}

func (s fontChooserWidget) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s fontChooserWidget) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s fontChooserWidget) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s fontChooserWidget) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s fontChooserWidget) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s fontChooserWidget) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s fontChooserWidget) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b fontChooserWidget) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (f fontChooserWidget) Font() string {
	return WrapFontChooser(gextras.InternObject(f)).Font()
}

func (f fontChooserWidget) FontDesc() *pango.FontDescription {
	return WrapFontChooser(gextras.InternObject(f)).FontDesc()
}

func (f fontChooserWidget) FontFace() pango.FontFace {
	return WrapFontChooser(gextras.InternObject(f)).FontFace()
}

func (f fontChooserWidget) FontFamily() pango.FontFamily {
	return WrapFontChooser(gextras.InternObject(f)).FontFamily()
}

func (f fontChooserWidget) FontFeatures() string {
	return WrapFontChooser(gextras.InternObject(f)).FontFeatures()
}

func (f fontChooserWidget) FontMap() pango.FontMap {
	return WrapFontChooser(gextras.InternObject(f)).FontMap()
}

func (f fontChooserWidget) FontSize() int {
	return WrapFontChooser(gextras.InternObject(f)).FontSize()
}

func (f fontChooserWidget) Language() string {
	return WrapFontChooser(gextras.InternObject(f)).Language()
}

func (f fontChooserWidget) Level() FontChooserLevel {
	return WrapFontChooser(gextras.InternObject(f)).Level()
}

func (f fontChooserWidget) PreviewText() string {
	return WrapFontChooser(gextras.InternObject(f)).PreviewText()
}

func (f fontChooserWidget) ShowPreviewEntry() bool {
	return WrapFontChooser(gextras.InternObject(f)).ShowPreviewEntry()
}

func (f fontChooserWidget) SetFont(fontname string) {
	WrapFontChooser(gextras.InternObject(f)).SetFont(fontname)
}

func (f fontChooserWidget) SetFontDesc(fontDesc *pango.FontDescription) {
	WrapFontChooser(gextras.InternObject(f)).SetFontDesc(fontDesc)
}

func (f fontChooserWidget) SetFontMap(fontmap pango.FontMap) {
	WrapFontChooser(gextras.InternObject(f)).SetFontMap(fontmap)
}

func (f fontChooserWidget) SetLanguage(language string) {
	WrapFontChooser(gextras.InternObject(f)).SetLanguage(language)
}

func (f fontChooserWidget) SetLevel(level FontChooserLevel) {
	WrapFontChooser(gextras.InternObject(f)).SetLevel(level)
}

func (f fontChooserWidget) SetPreviewText(text string) {
	WrapFontChooser(gextras.InternObject(f)).SetPreviewText(text)
}

func (f fontChooserWidget) SetShowPreviewEntry(showPreviewEntry bool) {
	WrapFontChooser(gextras.InternObject(f)).SetShowPreviewEntry(showPreviewEntry)
}