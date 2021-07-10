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
		{T: externglib.Type(C.pango_language_get_type()), F: marshalLanguage},
	})
}

// Language: the `PangoLanguage` structure is used to represent a language.
//
// `PangoLanguage` pointers can be efficiently copied and compared with each
// other.
type Language struct {
	native C.PangoLanguage
}

func marshalLanguage(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Language)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *Language) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// SampleString: get a string that is representative of the characters needed to
// render a particular language.
//
// The sample text may be a pangram, but is not necessarily. It is chosen to be
// demonstrative of normal text in the language, as well as exposing font
// feature requirements unique to the language. It is suitable for use as sample
// text in a font selection dialog.
//
// If @language is nil, the default language as found by
// [type_func@Pango.Language.get_default] is used.
//
// If Pango does not have a sample string for @language, the classic "The quick
// brown fox..." is returned. This can be detected by comparing the returned
// pointer value to that returned for (non-existent) language code "xx". That
// is, compare to:
//
// “` pango_language_get_sample_string (pango_language_from_string ("xx")) “`
func (language *Language) SampleString() string {
	var _arg0 *C.PangoLanguage // out
	var _cret *C.char          // in

	_arg0 = (*C.PangoLanguage)(unsafe.Pointer(language))

	_cret = C.pango_language_get_sample_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Matches checks if a language tag matches one of the elements in a list of
// language ranges.
//
// A language tag is considered to match a range in the list if the range is
// '*', the range is exactly the tag, or the range is a prefix of the tag, and
// the character after it in the tag is '-'.
func (language *Language) Matches(rangeList string) bool {
	var _arg0 *C.PangoLanguage // out
	var _arg1 *C.char          // out
	var _cret C.gboolean       // in

	_arg0 = (*C.PangoLanguage)(unsafe.Pointer(language))
	_arg1 = (*C.char)(C.CString(rangeList))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.pango_language_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String gets the RFC-3066 format string representing the given language tag.
func (language *Language) String() string {
	var _arg0 *C.PangoLanguage // out
	var _cret *C.char          // in

	_arg0 = (*C.PangoLanguage)(unsafe.Pointer(language))

	_cret = C.pango_language_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
