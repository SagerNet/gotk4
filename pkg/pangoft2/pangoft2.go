// Code generated by girgen. DO NOT EDIT.

package pangoft2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoft2 pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangoft2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_ft2_font_map_get_type()), F: marshalFontMap},
	})
}

// FontGetCoverage gets the Coverage for a `PangoFT2Font`. Use
// pango_font_get_coverage() instead.
func FontGetCoverage(font pango.Font, language *pango.Language) pango.Coverage {
	var arg1 *C.PangoFont
	var arg2 *C.PangoLanguage

	arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	arg2 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	var cret *C.PangoCoverage
	var goret1 pango.Coverage

	cret = C.pango_ft2_font_get_coverage(font, language)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(pango.Coverage)

	return goret1
}

// GetContext retrieves a `PangoContext` for the default PangoFT2 fontmap (see
// pango_ft2_font_map_for_display()) and sets the resolution for the default
// fontmap to @dpi_x by @dpi_y.
func GetContext(dpiX float64, dpiY float64) pango.Context {
	var arg1 C.double
	var arg2 C.double

	arg1 = C.double(dpiX)
	arg2 = C.double(dpiY)

	var cret *C.PangoContext
	var goret1 pango.Context

	cret = C.pango_ft2_get_context(dpiX, dpiY)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(pango.Context)

	return goret1
}

// ShutdownDisplay: free the global fontmap. (See
// pango_ft2_font_map_for_display()) Use of the global PangoFT2 fontmap is
// deprecated.
func ShutdownDisplay() {
	C.pango_ft2_shutdown_display()
}

// FontMap: the `PangoFT2FontMap` is the `PangoFontMap` implementation for
// FreeType fonts.
type FontMap interface {
	pangofc.FontMap

	// CreateContext: create a `PangoContext` for the given fontmap.
	CreateContext() pango.Context
	// SetDefaultSubstitute sets a function that will be called to do final
	// configuration substitution on a `FcPattern` before it is used to load the
	// font.
	//
	// This function can be used to do things like set hinting and antialiasing
	// options.
	SetDefaultSubstitute(fn SubstituteFunc)
	// SetResolution sets the horizontal and vertical resolutions for the
	// fontmap.
	SetResolution(dpiX float64, dpiY float64)
	// SubstituteChanged: call this function any time the results of the default
	// substitution function set with
	// pango_ft2_font_map_set_default_substitute() change.
	//
	// That is, if your substitution function will return different results for
	// the same input pattern, you must call this function.
	SubstituteChanged()
}

// fontMap implements the FontMap interface.
type fontMap struct {
	pangofc.FontMap
}

var _ FontMap = (*fontMap)(nil)

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return FontMap{
		pangofc.FontMap: pangofc.WrapFontMap(obj),
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

// NewFontMap constructs a class FontMap.
func NewFontMap() FontMap {
	var cret C.PangoFT2FontMap
	var goret1 FontMap

	cret = C.pango_ft2_font_map_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(FontMap)

	return goret1
}

// CreateContext: create a `PangoContext` for the given fontmap.
func (f fontMap) CreateContext() pango.Context {
	var arg0 *C.PangoFT2FontMap

	arg0 = (*C.PangoFT2FontMap)(unsafe.Pointer(f.Native()))

	var cret *C.PangoContext
	var goret1 pango.Context

	cret = C.pango_ft2_font_map_create_context(arg0)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(pango.Context)

	return goret1
}

// SetDefaultSubstitute sets a function that will be called to do final
// configuration substitution on a `FcPattern` before it is used to load the
// font.
//
// This function can be used to do things like set hinting and antialiasing
// options.
func (f fontMap) SetDefaultSubstitute(fn SubstituteFunc) {
	var arg0 *C.PangoFT2FontMap

	arg0 = (*C.PangoFT2FontMap)(unsafe.Pointer(f.Native()))

	C.pango_ft2_font_map_set_default_substitute(arg0, fn, data, notify)
}

// SetResolution sets the horizontal and vertical resolutions for the
// fontmap.
func (f fontMap) SetResolution(dpiX float64, dpiY float64) {
	var arg0 *C.PangoFT2FontMap
	var arg1 C.double
	var arg2 C.double

	arg0 = (*C.PangoFT2FontMap)(unsafe.Pointer(f.Native()))
	arg1 = C.double(dpiX)
	arg2 = C.double(dpiY)

	C.pango_ft2_font_map_set_resolution(arg0, dpiX, dpiY)
}

// SubstituteChanged: call this function any time the results of the default
// substitution function set with
// pango_ft2_font_map_set_default_substitute() change.
//
// That is, if your substitution function will return different results for
// the same input pattern, you must call this function.
func (f fontMap) SubstituteChanged() {
	var arg0 *C.PangoFT2FontMap

	arg0 = (*C.PangoFT2FontMap)(unsafe.Pointer(f.Native()))

	C.pango_ft2_font_map_substitute_changed(arg0)
}
