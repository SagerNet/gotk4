// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

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
		{T: externglib.Type(C.pango_bidi_type_get_type()), F: marshalBidiType},
	})
}

// BidiType: `PangoBidiType` represents the bidirectional character type of a
// Unicode character as specified by the <ulink
// url="http://www.unicode.org/reports/tr9/">Unicode bidirectional
// algorithm</ulink>.
//
// Deprecated: since version 1.44.
type BidiType int

const (
	// l: left-to-Right
	BidiTypeL BidiType = 0
	// lre: left-to-Right Embedding
	BidiTypeLre BidiType = 1
	// lro: left-to-Right Override
	BidiTypeLro BidiType = 2
	// r: right-to-Left
	BidiTypeR BidiType = 3
	// al: right-to-Left Arabic
	BidiTypeAl BidiType = 4
	// rle: right-to-Left Embedding
	BidiTypeRle BidiType = 5
	// rlo: right-to-Left Override
	BidiTypeRlo BidiType = 6
	// pdf: pop Directional Format
	BidiTypePDF BidiType = 7
	// en: european Number
	BidiTypeEn BidiType = 8
	// es: european Number Separator
	BidiTypeES BidiType = 9
	// et: european Number Terminator
	BidiTypeEt BidiType = 10
	// an: arabic Number
	BidiTypeAn BidiType = 11
	// cs: common Number Separator
	BidiTypeCs BidiType = 12
	// nsm: nonspacing Mark
	BidiTypeNsm BidiType = 13
	// bn: boundary Neutral
	BidiTypeBn BidiType = 14
	// b: paragraph Separator
	BidiTypeB BidiType = 15
	// s: segment Separator
	BidiTypeS BidiType = 16
	// ws: whitespace
	BidiTypeWs BidiType = 17
	// on: other Neutrals
	BidiTypeOn BidiType = 18
)

func marshalBidiType(p uintptr) (interface{}, error) {
	return BidiType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FindBaseDir searches a string the first character that has a strong
// direction, according to the Unicode bidirectional algorithm.
func FindBaseDir(text string, length int) Direction {
	var _arg1 *C.gchar         // out
	var _arg2 C.gint           // out
	var _cret C.PangoDirection // in

	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(length)

	_cret = C.pango_find_base_dir(_arg1, _arg2)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}

// GetMirrorChar returns the mirrored character of a Unicode character.
//
// Mirror characters are determined by the Unicode mirrored property.
//
// Use g_unichar_get_mirror_char() instead; the docs for that function provide
// full details.
func GetMirrorChar(ch uint32, mirroredCh *uint32) bool {
	var _arg1 C.gunichar  // out
	var _arg2 *C.gunichar // out
	var _cret C.gboolean  // in

	_arg1 = C.gunichar(ch)
	_arg2 = (*C.gunichar)(unsafe.Pointer(mirroredCh))

	_cret = C.pango_get_mirror_char(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnicharDirection determines the inherent direction of a character.
//
// The inherent direction is either PANGO_DIRECTION_LTR, PANGO_DIRECTION_RTL, or
// PANGO_DIRECTION_NEUTRAL.
//
// This function is useful to categorize characters into left-to-right letters,
// right-to-left letters, and everything else. If full Unicode bidirectional
// type of a character is needed, [type_func@Pango.BidiType.for_unichar] can be
// used instead.
func UnicharDirection(ch uint32) Direction {
	var _arg1 C.gunichar       // out
	var _cret C.PangoDirection // in

	_arg1 = C.gunichar(ch)

	_cret = C.pango_unichar_direction(_arg1)

	var _direction Direction // out

	_direction = Direction(_cret)

	return _direction
}
