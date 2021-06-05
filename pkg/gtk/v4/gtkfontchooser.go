// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_get_type()), F: marshalFontChooser},
	})
}

// FontFilterFunc: the type of function that is used for deciding what fonts get
// shown in a FontChooser. See gtk_font_chooser_set_filter_func().
type FontFilterFunc func(family pango.FontFamily, face pango.FontFace) bool

//export gotk4_FontFilterFunc
func gotk4_FontFilterFunc(arg0 *C.PangoFontFamily, arg1 *C.PangoFontFace, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}
	fn := v.(FontFilterFunc)

	ret := fn(family, face, data)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// FontChooserOverrider contains methods that are overridable. This
// interface is a subset of the interface FontChooser.
type FontChooserOverrider interface {
	FontActivated(fontname string)
	// FontFace gets the FontFace representing the selected font group details
	// (i.e. family, slant, weight, width, etc).
	//
	// If the selected font is not installed, returns nil.
	FontFace() pango.FontFace
	// FontFamily gets the FontFamily representing the selected font family.
	// Font families are a collection of font faces.
	//
	// If the selected font is not installed, returns nil.
	FontFamily() pango.FontFamily
	// FontMap gets the custom font map of this font chooser widget, or nil if
	// it does not have one.
	FontMap() pango.FontMap
	// FontSize: the selected font size.
	FontSize() int
	// SetFilterFunc adds a filter function that decides which fonts to display
	// in the font chooser.
	SetFilterFunc(filter FontFilterFunc)
	// SetFontMap sets a custom font map to use for this font chooser widget. A
	// custom font map can be used to present application-specific fonts instead
	// of or in addition to the normal system fonts.
	//
	//    FcConfig *config;
	//    PangoFontMap *fontmap;
	//
	//    config = FcInitLoadConfigAndFonts ();
	//    FcConfigAppFontAddFile (config, my_app_font_file);
	//
	//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
	//
	// Note that other GTK widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	//    context = gtk_widget_get_pango_context (label);
	//    pango_context_set_font_map (context, fontmap);
	SetFontMap(fontmap pango.FontMap)
}

// FontChooser is an interface that can be implemented by widgets displaying the
// list of fonts. In GTK, the main objects that implement this interface are
// FontChooserWidget, FontChooserDialog and FontButton.
type FontChooser interface {
	gextras.Objector
	FontChooserOverrider

	// Font gets the currently-selected font name.
	//
	// Note that this can be a different string than what you set with
	// gtk_font_chooser_set_font(), as the font chooser widget may normalize
	// font names and thus return a string with a different structure. For
	// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”.
	//
	// Use pango_font_description_equal() if you want to compare two font
	// descriptions.
	Font() string
	// FontDesc gets the currently-selected font.
	//
	// Note that this can be a different string than what you set with
	// gtk_font_chooser_set_font(), as the font chooser widget may normalize
	// font names and thus return a string with a different structure. For
	// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”.
	//
	// Use pango_font_description_equal() if you want to compare two font
	// descriptions.
	FontDesc() *pango.FontDescription
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// Language gets the language that is used for font features.
	Language() string
	// Level returns the current level of granularity for selecting fonts.
	Level() FontChooserLevel
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from @font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetLevel sets the desired level of granularity for selecting fonts.
	SetLevel(level FontChooserLevel)
	// SetPreviewText sets the text displayed in the preview area. The @text is
	// used to show how the selected font looks.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)
}

// fontChooser implements the FontChooser interface.
type fontChooser struct {
	gextras.Objector
}

var _ FontChooser = (*fontChooser)(nil)

// WrapFontChooser wraps a GObject to a type that implements interface
// FontChooser. It is primarily used internally.
func WrapFontChooser(obj *externglib.Object) FontChooser {
	return FontChooser{
		Objector: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooser(obj), nil
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize
// font names and thus return a string with a different structure. For
// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
// Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
func (f fontChooser) Font() string {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_font_chooser_get_font(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// FontDesc gets the currently-selected font.
//
// Note that this can be a different string than what you set with
// gtk_font_chooser_set_font(), as the font chooser widget may normalize
// font names and thus return a string with a different structure. For
// example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
// Bold Italic 12”.
//
// Use pango_font_description_equal() if you want to compare two font
// descriptions.
func (f fontChooser) FontDesc() *pango.FontDescription {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontDescription
	var goret1 *pango.FontDescription

	cret = C.gtk_font_chooser_get_font_desc(arg0)

	goret1 = pango.WrapFontDescription(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *pango.FontDescription) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// FontFace gets the FontFace representing the selected font group details
// (i.e. family, slant, weight, width, etc).
//
// If the selected font is not installed, returns nil.
func (f fontChooser) FontFace() pango.FontFace {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontFace
	var goret1 pango.FontFace

	cret = C.gtk_font_chooser_get_font_face(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.FontFace)

	return goret1
}

// FontFamily gets the FontFamily representing the selected font family.
// Font families are a collection of font faces.
//
// If the selected font is not installed, returns nil.
func (f fontChooser) FontFamily() pango.FontFamily {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontFamily
	var goret1 pango.FontFamily

	cret = C.gtk_font_chooser_get_font_family(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.FontFamily)

	return goret1
}

// FontFeatures gets the currently-selected font features.
func (f fontChooser) FontFeatures() string {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_font_chooser_get_font_features(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// FontMap gets the custom font map of this font chooser widget, or nil if
// it does not have one.
func (f fontChooser) FontMap() pango.FontMap {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.PangoFontMap
	var goret1 pango.FontMap

	cret = C.gtk_font_chooser_get_font_map(arg0)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(pango.FontMap)

	return goret1
}

// FontSize: the selected font size.
func (f fontChooser) FontSize() int {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret C.int
	var goret1 int

	cret = C.gtk_font_chooser_get_font_size(arg0)

	goret1 = C.int(cret)

	return goret1
}

// Language gets the language that is used for font features.
func (f fontChooser) Language() string {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_font_chooser_get_language(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// Level returns the current level of granularity for selecting fonts.
func (f fontChooser) Level() FontChooserLevel {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret C.GtkFontChooserLevel
	var goret1 FontChooserLevel

	cret = C.gtk_font_chooser_get_level(arg0)

	goret1 = FontChooserLevel(cret)

	return goret1
}

// PreviewText gets the text displayed in the preview area.
func (f fontChooser) PreviewText() string {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret *C.char
	var goret1 string

	cret = C.gtk_font_chooser_get_preview_text(arg0)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
func (f fontChooser) ShowPreviewEntry() bool {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_font_chooser_get_show_preview_entry(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetFilterFunc adds a filter function that decides which fonts to display
// in the font chooser.
func (f fontChooser) SetFilterFunc(filter FontFilterFunc) {
	var arg0 *C.GtkFontChooser

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	C.gtk_font_chooser_set_filter_func(arg0, filter, userData, destroy)
}

// SetFont sets the currently-selected font.
func (f fontChooser) SetFont(fontname string) {
	var arg0 *C.GtkFontChooser
	var arg1 *C.char

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(fontname))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_font_chooser_set_font(arg0, fontname)
}

// SetFontDesc sets the currently-selected font from @font_desc.
func (f fontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var arg0 *C.GtkFontChooser
	var arg1 *C.PangoFontDescription

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(fontDesc.Native()))

	C.gtk_font_chooser_set_font_desc(arg0, fontDesc)
}

// SetFontMap sets a custom font map to use for this font chooser widget. A
// custom font map can be used to present application-specific fonts instead
// of or in addition to the normal system fonts.
//
//    FcConfig *config;
//    PangoFontMap *fontmap;
//
//    config = FcInitLoadConfigAndFonts ();
//    FcConfigAppFontAddFile (config, my_app_font_file);
//
//    fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
//    pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
//    gtk_font_chooser_set_font_map (font_chooser, fontmap);
//
// Note that other GTK widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
//    context = gtk_widget_get_pango_context (label);
//    pango_context_set_font_map (context, fontmap);
func (f fontChooser) SetFontMap(fontmap pango.FontMap) {
	var arg0 *C.GtkFontChooser
	var arg1 *C.PangoFontMap

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	C.gtk_font_chooser_set_font_map(arg0, fontmap)
}

// SetLanguage sets the language to use for font features.
func (f fontChooser) SetLanguage(language string) {
	var arg0 *C.GtkFontChooser
	var arg1 *C.char

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(language))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_font_chooser_set_language(arg0, language)
}

// SetLevel sets the desired level of granularity for selecting fonts.
func (f fontChooser) SetLevel(level FontChooserLevel) {
	var arg0 *C.GtkFontChooser
	var arg1 C.GtkFontChooserLevel

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (C.GtkFontChooserLevel)(level)

	C.gtk_font_chooser_set_level(arg0, level)
}

// SetPreviewText sets the text displayed in the preview area. The @text is
// used to show how the selected font looks.
func (f fontChooser) SetPreviewText(text string) {
	var arg0 *C.GtkFontChooser
	var arg1 *C.char

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_font_chooser_set_preview_text(arg0, text)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
func (f fontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var arg0 *C.GtkFontChooser
	var arg1 C.gboolean

	arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	if showPreviewEntry {
		arg1 = C.gboolean(1)
	}

	C.gtk_font_chooser_set_show_preview_entry(arg0, showPreviewEntry)
}
