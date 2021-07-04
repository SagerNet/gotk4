// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_font_map_get_type()), F: marshalFontMap},
	})
}

// FontMap: a `PangoFontMap` represents the set of fonts available for a
// particular rendering system.
//
// This is a virtual object with implementations being specific to particular
// rendering systems.
type FontMap interface {
	gextras.Objector

	// ChangedFontMap:
	ChangedFontMap()
	// CreateContextFontMap:
	CreateContextFontMap() Context
	// Family:
	Family(name string) FontFamily
	// Serial:
	Serial() uint
	// ListFamiliesFontMap:
	ListFamiliesFontMap() []FontFamily
	// LoadFontFontMap:
	LoadFontFontMap(context Context, desc *FontDescription) Font
	// LoadFontsetFontMap:
	LoadFontsetFontMap(context Context, desc *FontDescription, language *Language) Fontset
}

// fontMap implements the FontMap class.
type fontMap struct {
	gextras.Objector
}

// WrapFontMap wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{
		Objector: obj,
	}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontMap(obj), nil
}

func (f fontMap) ChangedFontMap() {
	var _arg0 *C.PangoFontMap // out

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	C.pango_font_map_changed(_arg0)
}

func (f fontMap) CreateContextFontMap() Context {
	var _arg0 *C.PangoFontMap // out
	var _cret *C.PangoContext // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	_cret = C.pango_font_map_create_context(_arg0)

	var _context Context // out

	_context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Context)

	return _context
}

func (f fontMap) Family(name string) FontFamily {
	var _arg0 *C.PangoFontMap    // out
	var _arg1 *C.char            // out
	var _cret *C.PangoFontFamily // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_font_map_get_family(_arg0, _arg1)

	var _fontFamily FontFamily // out

	_fontFamily = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FontFamily)

	return _fontFamily
}

func (f fontMap) Serial() uint {
	var _arg0 *C.PangoFontMap // out
	var _cret C.guint         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	_cret = C.pango_font_map_get_serial(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (f fontMap) ListFamiliesFontMap() []FontFamily {
	var _arg0 *C.PangoFontMap // out
	var _arg1 **C.PangoFontFamily
	var _arg2 C.int // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))

	C.pango_font_map_list_families(_arg0, &_arg1, &_arg2)

	var _families []FontFamily

	defer C.free(unsafe.Pointer(_arg1))
	{
		src := unsafe.Slice(_arg1, _arg2)
		_families = make([]FontFamily, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_families[i] = gextras.CastObject(externglib.Take(unsafe.Pointer(src[i]))).(FontFamily)
		}
	}

	return _families
}

func (f fontMap) LoadFontFontMap(context Context, desc *FontDescription) Font {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _cret *C.PangoFont            // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))

	_cret = C.pango_font_map_load_font(_arg0, _arg1, _arg2)

	var _font Font // out

	_font = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Font)

	return _font
}

func (f fontMap) LoadFontsetFontMap(context Context, desc *FontDescription, language *Language) Fontset {
	var _arg0 *C.PangoFontMap         // out
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.PangoFontDescription // out
	var _arg3 *C.PangoLanguage        // out
	var _cret *C.PangoFontset         // in

	_arg0 = (*C.PangoFontMap)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))
	_arg3 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	_cret = C.pango_font_map_load_fontset(_arg0, _arg1, _arg2, _arg3)

	var _fontset Fontset // out

	_fontset = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Fontset)

	return _fontset
}
