// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_font_selection_get_type()), F: marshalFontSelection},
		{T: externglib.Type(C.gtk_font_selection_dialog_get_type()), F: marshalFontSelectionDialog},
	})
}

// FontSelection:
type FontSelection interface {
	Box

	// Face:
	Face() pango.FontFace
	// FaceList:
	FaceList() Widget
	// Family:
	Family() pango.FontFamily
	// FamilyList:
	FamilyList() Widget
	// FontName:
	FontName() string
	// PreviewEntry:
	PreviewEntry() Widget
	// PreviewText:
	PreviewText() string
	// Size:
	Size() int
	// SizeEntry:
	SizeEntry() Widget
	// SizeList:
	SizeList() Widget
	// SetFontNameFontSelection:
	SetFontNameFontSelection(fontname string) bool
	// SetPreviewTextFontSelection:
	SetPreviewTextFontSelection(text string)
}

// fontSelection implements the FontSelection class.
type fontSelection struct {
	Box
}

// WrapFontSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontSelection(obj *externglib.Object) FontSelection {
	return fontSelection{
		Box: WrapBox(obj),
	}
}

func marshalFontSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontSelection(obj), nil
}

// NewFontSelection:
func NewFontSelection() FontSelection {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_selection_new()

	var _fontSelection FontSelection // out

	_fontSelection = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontSelection)

	return _fontSelection
}

func (f fontSelection) Face() pango.FontFace {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.PangoFontFace    // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_face(_arg0)

	var _fontFace pango.FontFace // out

	_fontFace = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.FontFace)

	return _fontFace
}

func (f fontSelection) FaceList() Widget {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_face_list(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelection) Family() pango.FontFamily {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.PangoFontFamily  // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_family(_arg0)

	var _fontFamily pango.FontFamily // out

	_fontFamily = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.FontFamily)

	return _fontFamily
}

func (f fontSelection) FamilyList() Widget {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_family_list(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelection) FontName() string {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_font_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontSelection) PreviewEntry() Widget {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_preview_entry(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelection) PreviewText() string {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f fontSelection) Size() int {
	var _arg0 *C.GtkFontSelection // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (f fontSelection) SizeEntry() Widget {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_size_entry(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelection) SizeList() Widget {
	var _arg0 *C.GtkFontSelection // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_get_size_list(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelection) SetFontNameFontSelection(fontname string) bool {
	var _arg0 *C.GtkFontSelection // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_set_font_name(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontSelection) SetPreviewTextFontSelection(text string) {
	var _arg0 *C.GtkFontSelection // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkFontSelection)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_selection_set_preview_text(_arg0, _arg1)
}

func (b fontSelection) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b fontSelection) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b fontSelection) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b fontSelection) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b fontSelection) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b fontSelection) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b fontSelection) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o fontSelection) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o fontSelection) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}

// FontSelectionDialog:
type FontSelectionDialog interface {
	Dialog

	// CancelButton:
	CancelButton() Widget
	// FontName:
	FontName() string
	// FontSelection:
	FontSelection() Widget
	// OkButton:
	OkButton() Widget
	// PreviewText:
	PreviewText() string
	// SetFontNameFontSelectionDialog:
	SetFontNameFontSelectionDialog(fontname string) bool
	// SetPreviewTextFontSelectionDialog:
	SetPreviewTextFontSelectionDialog(text string)
}

// fontSelectionDialog implements the FontSelectionDialog class.
type fontSelectionDialog struct {
	Dialog
}

// WrapFontSelectionDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontSelectionDialog(obj *externglib.Object) FontSelectionDialog {
	return fontSelectionDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalFontSelectionDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontSelectionDialog(obj), nil
}

// NewFontSelectionDialog:
func NewFontSelectionDialog(title string) FontSelectionDialog {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_dialog_new(_arg1)

	var _fontSelectionDialog FontSelectionDialog // out

	_fontSelectionDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontSelectionDialog)

	return _fontSelectionDialog
}

func (f fontSelectionDialog) CancelButton() Widget {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_dialog_get_cancel_button(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelectionDialog) FontName() string {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_dialog_get_font_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (f fontSelectionDialog) FontSelection() Widget {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_dialog_get_font_selection(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelectionDialog) OkButton() Widget {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.GtkWidget              // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_dialog_get_ok_button(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f fontSelectionDialog) PreviewText() string {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _cret *C.gchar                  // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_font_selection_dialog_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f fontSelectionDialog) SetFontNameFontSelectionDialog(fontname string) bool {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _arg1 *C.gchar                  // out
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_font_selection_dialog_set_font_name(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (f fontSelectionDialog) SetPreviewTextFontSelectionDialog(text string) {
	var _arg0 *C.GtkFontSelectionDialog // out
	var _arg1 *C.gchar                  // out

	_arg0 = (*C.GtkFontSelectionDialog)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_selection_dialog_set_preview_text(_arg0, _arg1)
}

func (b fontSelectionDialog) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b fontSelectionDialog) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b fontSelectionDialog) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b fontSelectionDialog) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b fontSelectionDialog) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b fontSelectionDialog) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b fontSelectionDialog) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
