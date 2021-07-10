// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
//
// gboolean gotk4_FontsetForeachFunc(PangoFontset*, PangoFont*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_fontset_get_type()), F: marshalFontsetter},
		{T: externglib.Type(C.pango_fontset_simple_get_type()), F: marshalFontsetSimpler},
	})
}

// FontsetForeachFunc: callback used by pango_fontset_foreach() when enumerating
// fonts in a fontset.
type FontsetForeachFunc func(fontset *Fontset, font *Font, userData interface{}) (ok bool)

//export gotk4_FontsetForeachFunc
func gotk4_FontsetForeachFunc(arg0 *C.PangoFontset, arg1 *C.PangoFont, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var fontset *Fontset     // out
	var font *Font           // out
	var userData interface{} // out

	fontset = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*Fontset)
	font = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*Font)
	userData = box.Get(uintptr(arg2))

	fn := v.(FontsetForeachFunc)
	ok := fn(fontset, font, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FontsetterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FontsetterOverrider interface {
	// Foreach iterates through all the fonts in a fontset, calling @func for
	// each one.
	//
	// If @func returns true, that stops the iteration.
	Foreach(fn FontsetForeachFunc)
	// Font returns the font in the fontset that contains the best glyph for a
	// Unicode character.
	Font(wc uint) *Font
	Language() *Language
	// Metrics: get overall metric information for the fonts in the fontset.
	Metrics() *FontMetrics
}

// Fontsetter describes Fontset's methods.
type Fontsetter interface {
	gextras.Objector

	Foreach(fn FontsetForeachFunc)
	Font(wc uint) *Font
	Metrics() *FontMetrics
}

// Fontset: `PangoFontset` represents a set of `PangoFont` to use when rendering
// text.
//
// A `PAngoFontset` is the result of resolving a `PangoFontDescription` against
// a particular `PangoContext`. It has operations for finding the component font
// for a particular Unicode character, and for finding a composite set of
// metrics for the entire fontset.
type Fontset struct {
	*externglib.Object
}

var _ Fontsetter = (*Fontset)(nil)

func wrapFontsetter(obj *externglib.Object) Fontsetter {
	return &Fontset{
		Object: obj,
	}
}

func marshalFontsetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontsetter(obj), nil
}

// Foreach iterates through all the fonts in a fontset, calling @func for each
// one.
//
// If @func returns true, that stops the iteration.
func (fontset *Fontset) Foreach(fn FontsetForeachFunc) {
	var _arg0 *C.PangoFontset           // out
	var _arg1 C.PangoFontsetForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*[0]byte)(C.gotk4_FontsetForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.pango_fontset_foreach(_arg0, _arg1, _arg2)
}

// Font returns the font in the fontset that contains the best glyph for a
// Unicode character.
func (fontset *Fontset) Font(wc uint) *Font {
	var _arg0 *C.PangoFontset // out
	var _arg1 C.guint         // out
	var _cret *C.PangoFont    // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))
	_arg1 = C.guint(wc)

	_cret = C.pango_fontset_get_font(_arg0, _arg1)

	var _font *Font // out

	_font = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Font)

	return _font
}

// Metrics: get overall metric information for the fonts in the fontset.
func (fontset *Fontset) Metrics() *FontMetrics {
	var _arg0 *C.PangoFontset     // out
	var _cret *C.PangoFontMetrics // in

	_arg0 = (*C.PangoFontset)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_get_metrics(_arg0)

	var _fontMetrics *FontMetrics // out

	_fontMetrics = (*FontMetrics)(unsafe.Pointer(_cret))
	C.pango_font_metrics_ref(_cret)
	runtime.SetFinalizer(_fontMetrics, func(v *FontMetrics) {
		C.pango_font_metrics_unref((*C.PangoFontMetrics)(unsafe.Pointer(v)))
	})

	return _fontMetrics
}

// FontsetSimpler describes FontsetSimple's methods.
type FontsetSimpler interface {
	gextras.Objector

	Append(font Fonter)
	Size() int
}

// FontsetSimple: `PangoFontsetSimple` is a implementation of the abstract
// `PangoFontset` base class as an array of fonts.
//
// When creating a `PangoFontsetSimple`, you have to provide the array of fonts
// that make up the fontset.
type FontsetSimple struct {
	Fontset
}

var _ FontsetSimpler = (*FontsetSimple)(nil)

func wrapFontsetSimpler(obj *externglib.Object) FontsetSimpler {
	return &FontsetSimple{
		Fontset: Fontset{
			Object: obj,
		},
	}
}

func marshalFontsetSimpler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontsetSimpler(obj), nil
}

// NewFontsetSimple creates a new `PangoFontsetSimple` for the given language.
func NewFontsetSimple(language *Language) *FontsetSimple {
	var _arg1 *C.PangoLanguage      // out
	var _cret *C.PangoFontsetSimple // in

	_arg1 = (*C.PangoLanguage)(unsafe.Pointer(language))

	_cret = C.pango_fontset_simple_new(_arg1)

	var _fontsetSimple *FontsetSimple // out

	_fontsetSimple = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*FontsetSimple)

	return _fontsetSimple
}

// Append adds a font to the fontset.
func (fontset *FontsetSimple) Append(font Fonter) {
	var _arg0 *C.PangoFontsetSimple // out
	var _arg1 *C.PangoFont          // out

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))
	_arg1 = (*C.PangoFont)(unsafe.Pointer(font.Native()))

	C.pango_fontset_simple_append(_arg0, _arg1)
}

// Size returns the number of fonts in the fontset.
func (fontset *FontsetSimple) Size() int {
	var _arg0 *C.PangoFontsetSimple // out
	var _cret C.int                 // in

	_arg0 = (*C.PangoFontsetSimple)(unsafe.Pointer(fontset.Native()))

	_cret = C.pango_fontset_simple_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
