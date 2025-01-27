// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// GetCharset obtains the character set for the [current locale][setlocale]; you
// might use this character set as an argument to g_convert(), to convert from
// the current locale's encoding to some other encoding. (Frequently
// g_locale_to_utf8() and g_locale_from_utf8() are nice shortcuts, though.)
//
// On Windows the character set returned by this function is the so-called
// system default ANSI code-page. That is the character set used by the "narrow"
// versions of C library and Win32 functions that handle file names. It might be
// different from the character set used by the C library's current locale.
//
// On Linux, the character set is found by consulting nl_langinfo() if
// available. If not, the environment variables LC_ALL, LC_CTYPE, LANG and
// CHARSET are queried in order.
//
// The return value is TRUE if the locale's encoding is UTF-8, in that case you
// can perhaps avoid calling g_convert().
//
// The string returned in charset is not allocated, and should not be freed.
func GetCharset() (string, bool) {
	var _arg1 *C.char    // in
	var _cret C.gboolean // in

	_cret = C.g_get_charset(&_arg1)

	var _charset string // out
	var _ok bool        // out

	if _arg1 != nil {
		_charset = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	if _cret != 0 {
		_ok = true
	}

	return _charset, _ok
}

// GetCodeset gets the character set for the current locale.
func GetCodeset() string {
	var _cret *C.gchar // in

	_cret = C.g_get_codeset()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GetConsoleCharset obtains the character set used by the console attached to
// the process, which is suitable for printing output to the terminal.
//
// Usually this matches the result returned by g_get_charset(), but in
// environments where the locale's character set does not match the encoding of
// the console this function tries to guess a more suitable value instead.
//
// On Windows the character set returned by this function is the output code
// page used by the console associated with the calling process. If the codepage
// can't be determined (for example because there is no console attached) UTF-8
// is assumed.
//
// The return value is TRUE if the locale's encoding is UTF-8, in that case you
// can perhaps avoid calling g_convert().
//
// The string returned in charset is not allocated, and should not be freed.
func GetConsoleCharset() (string, bool) {
	var _arg1 *C.char    // in
	var _cret C.gboolean // in

	_cret = C.g_get_console_charset(&_arg1)

	var _charset string // out
	var _ok bool        // out

	if _arg1 != nil {
		_charset = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	if _cret != 0 {
		_ok = true
	}

	return _charset, _ok
}

// GetLanguageNames computes a list of applicable locale names, which can be
// used to e.g. construct locale-dependent filenames or search paths. The
// returned list is sorted from most desirable to least desirable and always
// contains the default locale "C".
//
// For example, if LANGUAGE=de:en_US, then the returned list is "de", "en_US",
// "en", "C".
//
// This function consults the environment variables LANGUAGE, LC_ALL,
// LC_MESSAGES and LANG to find the list of locales specified by the user.
func GetLanguageNames() []string {
	var _cret **C.gchar // in

	_cret = C.g_get_language_names()

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// GetLanguageNamesWithCategory computes a list of applicable locale names with
// a locale category name, which can be used to construct the fallback
// locale-dependent filenames or search paths. The returned list is sorted from
// most desirable to least desirable and always contains the default locale "C".
//
// This function consults the environment variables LANGUAGE, LC_ALL,
// category_name, and LANG to find the list of locales specified by the user.
//
// g_get_language_names() returns
// g_get_language_names_with_category("LC_MESSAGES").
func GetLanguageNamesWithCategory(categoryName string) []string {
	var _arg1 *C.gchar  // out
	var _cret **C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(categoryName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_get_language_names_with_category(_arg1)
	runtime.KeepAlive(categoryName)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// GetLocaleVariants returns a list of derived variants of locale, which can be
// used to e.g. construct locale-dependent filenames or search paths. The
// returned list is sorted from most desirable to least desirable. This function
// handles territory, charset and extra locale modifiers. See setlocale(3)
// (man:setlocale) for information about locales and their format.
//
// locale itself is guaranteed to be returned in the output.
//
// For example, if locale is fr_BE, then the returned list is fr_BE, fr. If
// locale is en_GB.UTF-8euro, then the returned list is en_GB.UTF-8euro,
// en_GB.UTF-8, en_GBeuro, en_GB, en.UTF-8euro, en.UTF-8, eneuro, en.
//
// If you need the list of variants for the current locale, use
// g_get_language_names().
func GetLocaleVariants(locale string) []string {
	var _arg1 *C.gchar  // out
	var _cret **C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(locale)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_get_locale_variants(_arg1)
	runtime.KeepAlive(locale)

	var _utf8s []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}
