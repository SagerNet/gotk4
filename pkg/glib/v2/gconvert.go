// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/ptr"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <glib.h>
import "C"

// ConvertError: error codes returned by character set conversion routines.
type ConvertError int

const (
	// ConvertErrorNoConversion: conversion between the requested character sets
	// is not supported.
	ConvertErrorNoConversion ConvertError = 0
	// ConvertErrorIllegalSequence: invalid byte sequence in conversion input;
	// or the character sequence could not be represented in the target
	// character set.
	ConvertErrorIllegalSequence ConvertError = 1
	// ConvertErrorFailed: conversion failed for some reason.
	ConvertErrorFailed ConvertError = 2
	// ConvertErrorPartialInput: partial character sequence at end of input.
	ConvertErrorPartialInput ConvertError = 3
	// ConvertErrorBadURI: URI is invalid.
	ConvertErrorBadURI ConvertError = 4
	// ConvertErrorNotAbsolutePath: pathname is not an absolute path.
	ConvertErrorNotAbsolutePath ConvertError = 5
	// ConvertErrorNoMemory: no memory available. Since: 2.40
	ConvertErrorNoMemory ConvertError = 6
	// ConvertErrorEmbeddedNUL: an embedded NUL character is present in
	// conversion output where a NUL-terminated string is expected. Since: 2.56
	ConvertErrorEmbeddedNUL ConvertError = 7
)

// FilenameDisplayBasename returns the display basename for the particular
// filename, guaranteed to be valid UTF-8. The display name might not be
// identical to the filename, for instance there might be problems converting it
// to UTF-8, and some files can be translated in the display.
//
// If GLib cannot make sense of the encoding of @filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if @filename was in an invalid
// encoding.
//
// You must pass the whole absolute pathname to this functions so that
// translation of well known locations can be done.
//
// This function is preferred over g_filename_display_name() if you know the
// whole path, as it allows translation.
func FilenameDisplayBasename(filename string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_filename_display_basename(filename)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FilenameDisplayName converts a filename into a valid UTF-8 string. The
// conversion is not necessarily reversible, so you should keep the original
// around and use the return value of this function only for display purposes.
// Unlike g_filename_to_utf8(), the result is guaranteed to be non-nil even if
// the filename actually isn't in the GLib file name encoding.
//
// If GLib cannot make sense of the encoding of @filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if @filename was in an invalid
// encoding.
//
// If you know the whole pathname of the file you should use
// g_filename_display_basename(), since that allows location-based translation
// of filenames.
func FilenameDisplayName(filename string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var ret1 string

	cret = C.g_filename_display_name(filename)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// FilenameFromURI converts an escaped ASCII-encoded URI to a local filename in
// the encoding used for filenames.
func FilenameFromURI(uri string) (hostname string, filename string, err error) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.gchar
	var ret2 string
	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var ret3 string

	cret = C.g_filename_from_uri(uri, &arg2, &errout)

	*ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	goerr = gerror.Take(unsafe.Pointer(errout))
	ret3 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret2, goerr, ret3
}

// FilenameFromUTF8 converts a string from UTF-8 to the encoding GLib uses for
// filenames. Note that on Windows GLib uses UTF-8 for filenames; on other
// platforms, this function indirectly depends on the [current
// locale][setlocale].
//
// The input string shall not contain nul characters even if the @len argument
// is positive. A nul character found inside the string will result in error
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. If the filename encoding is not UTF-8 and
// the conversion output contains a nul character, the error
// G_CONVERT_ERROR_EMBEDDED_NUL is set and the function returns nil.
func FilenameFromUTF8(utf8String string, len int) (bytesRead uint, bytesWritten uint, filename string, err error) {
	var arg1 *C.gchar
	var arg2 C.gssize

	arg1 = (*C.gchar)(C.CString(utf8String))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	var arg3 C.gsize
	var ret3 uint
	var arg4 C.gsize
	var ret4 uint
	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var ret4 string

	cret = C.g_filename_from_utf8(utf8String, len, &arg3, &arg4, &errout)

	*ret3 = C.gsize(arg3)
	*ret4 = C.gsize(arg4)
	goerr = gerror.Take(unsafe.Pointer(errout))
	ret4 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret3, ret4, goerr, ret4
}

// FilenameToURI converts an absolute filename to an escaped ASCII-encoded URI,
// with the path component following Section 3.3. of RFC 2396.
func FilenameToURI(filename string, hostname string) (utf8 string, err error) {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(arg2))

	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var ret2 string

	cret = C.g_filename_to_uri(filename, hostname, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goerr, ret2
}

// FilenameToUTF8 converts a string which is in the encoding used by GLib for
// filenames into a UTF-8 string. Note that on Windows GLib uses UTF-8 for
// filenames; on other platforms, this function indirectly depends on the
// [current locale][setlocale].
//
// The input string shall not contain nul characters even if the @len argument
// is positive. A nul character found inside the string will result in error
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. If the source encoding is not UTF-8 and the
// conversion output contains a nul character, the error
// G_CONVERT_ERROR_EMBEDDED_NUL is set and the function returns nil. Use
// g_convert() to produce output that may contain embedded nul characters.
func FilenameToUTF8(opsysstring string, len int) (bytesRead uint, bytesWritten uint, utf8 string, err error) {
	var arg1 *C.gchar
	var arg2 C.gssize

	arg1 = (*C.gchar)(C.CString(opsysstring))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	var arg3 C.gsize
	var ret3 uint
	var arg4 C.gsize
	var ret4 uint
	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var ret4 string

	cret = C.g_filename_to_utf8(opsysstring, len, &arg3, &arg4, &errout)

	*ret3 = C.gsize(arg3)
	*ret4 = C.gsize(arg4)
	goerr = gerror.Take(unsafe.Pointer(errout))
	ret4 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret3, ret4, goerr, ret4
}

// GetFilenameCharsets determines the preferred character sets used for
// filenames. The first character set from the @charsets is the filename
// encoding, the subsequent character sets are used when trying to generate a
// displayable representation of a filename, see g_filename_display_name().
//
// On Unix, the character sets are determined by consulting the environment
// variables `G_FILENAME_ENCODING` and `G_BROKEN_FILENAMES`. On Windows, the
// character set used in the GLib API is always UTF-8 and said environment
// variables have no effect.
//
// `G_FILENAME_ENCODING` may be set to a comma-separated list of character set
// names. The special token "\@locale" is taken to mean the character set for
// the [current locale][setlocale]. If `G_FILENAME_ENCODING` is not set, but
// `G_BROKEN_FILENAMES` is, the character set of the current locale is taken as
// the filename encoding. If neither environment variable is set, UTF-8 is taken
// as the filename encoding, but the character set of the current locale is also
// put in the list of encodings.
//
// The returned @charsets belong to GLib and must not be freed.
//
// Note that on Unix, regardless of the locale character set or
// `G_FILENAME_ENCODING` value, the actual file names present on a system might
// be in any random encoding or just gibberish.
func GetFilenameCharsets() (filenameCharsets []string, ok bool) {
	var arg1 ***C.gchar
	var ret1 []string
	var cret C.gboolean
	var ret2 bool

	cret = C.g_get_filename_charsets(&arg1)

	{
		var length int
		for p := arg1; *p != 0; p = (***C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (**C.gchar)(ptr.Add(unsafe.Pointer(arg1), i))
			ret1[i] = C.GoString(src)
		}
	}
	ret2 = C.bool(cret) != C.false

	return ret1, ret2
}

// Iconv: same as the standard UNIX routine iconv(), but may be implemented via
// libiconv on UNIX flavors that lack a native implementation.
//
// GLib provides g_convert() and g_locale_to_utf8() which are likely more
// convenient than the raw iconv wrappers.
//
// Note that the behaviour of iconv() for characters which are valid in the
// input character set, but which have no representation in the output character
// set, is implementation defined. This function may return success (with a
// positive number of non-reversible conversions as replacement characters were
// used), or it may return -1 and set an error such as EILSEQ, in such a
// situation.
func Iconv(converter IConv, inbuf string, inbytesLeft uint, outbuf string, outbytesLeft uint) uint {
	var arg1 C.GIConv
	var arg2 **C.gchar
	var arg3 *C.gsize
	var arg4 **C.gchar
	var arg5 *C.gsize

	arg1 = (C.GIConv)(unsafe.Pointer(converter.Native()))
	arg2 = (**C.gchar)(C.CString(inbuf))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = *C.gsize(inbytesLeft)
	arg4 = (**C.gchar)(C.CString(outbuf))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = *C.gsize(outbytesLeft)

	var cret C.gsize
	var ret1 uint

	cret = C.g_iconv(converter, inbuf, inbytesLeft, outbuf, outbytesLeft)

	ret1 = C.gsize(cret)

	return ret1
}

// IconvOpen: same as the standard UNIX routine iconv_open(), but may be
// implemented via libiconv on UNIX flavors that lack a native implementation.
//
// GLib provides g_convert() and g_locale_to_utf8() which are likely more
// convenient than the raw iconv wrappers.
func IconvOpen(toCodeset string, fromCodeset string) IConv {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(toCodeset))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(fromCodeset))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GIConv
	var ret1 IConv

	cret = C.g_iconv_open(toCodeset, fromCodeset)

	ret1 = WrapIConv(unsafe.Pointer(cret))

	return ret1
}

// LocaleFromUTF8 converts a string from UTF-8 to the encoding used for strings
// by the C runtime (usually the same as that used by the operating system) in
// the [current locale][setlocale]. On Windows this means the system codepage.
//
// The input string shall not contain nul characters even if the @len argument
// is positive. A nul character found inside the string will result in error
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. Use g_convert() to convert input that may
// contain embedded nul characters.
func LocaleFromUTF8(utf8String string, len int) (bytesRead uint, bytesWritten uint, guint8s []byte, err error) {
	var arg1 *C.gchar
	var arg2 C.gssize

	arg1 = (*C.gchar)(C.CString(utf8String))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	var errout *C.GError
	var goerr error
	var cret *C.gchar
	var arg3 *C.gsize
	var ret4 []byte

	cret = C.g_locale_from_utf8(utf8String, len, &arg3, &arg4, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ptr.SetSlice(unsafe.Pointer(&ret4), unsafe.Pointer(cret), int(arg3))
	runtime.SetFinalizer(&ret4, func(v *[]byte) {
		C.free(ptr.Slice(unsafe.Pointer(v)))
	})

	return ret3, ret4, goerr, ret4
}

// URIListExtractUris splits an URI list conforming to the text/uri-list mime
// type defined in RFC 2483 into individual URIs, discarding any comments. The
// URIs are not validated.
func URIListExtractUris(uriList string) []string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(uriList))
	defer C.free(unsafe.Pointer(arg1))

	var cret **C.gchar
	var ret1 []string

	cret = C.g_uri_list_extract_uris(uriList)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		ret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			ret1[i] = C.GoString(src)
			defer C.free(unsafe.Pointer(src))
		}
	}

	return ret1
}

// IConv: the GIConv struct wraps an iconv() conversion descriptor. It contains
// private data and should only be accessed using the following functions.
type IConv struct {
	native C.GIConv
}

// WrapIConv wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIConv(ptr unsafe.Pointer) *IConv {
	if ptr == nil {
		return nil
	}

	return (*IConv)(ptr)
}

func marshalIConv(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapIConv(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *IConv) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// _: same as the standard UNIX routine iconv(), but may be implemented via
// libiconv on UNIX flavors that lack a native implementation.
//
// GLib provides g_convert() and g_locale_to_utf8() which are likely more
// convenient than the raw iconv wrappers.
//
// Note that the behaviour of iconv() for characters which are valid in the
// input character set, but which have no representation in the output character
// set, is implementation defined. This function may return success (with a
// positive number of non-reversible conversions as replacement characters were
// used), or it may return -1 and set an error such as EILSEQ, in such a
// situation.
func (c *IConv) _(inbuf string, inbytesLeft uint, outbuf string, outbytesLeft uint) uint {
	var arg0 C.GIConv
	var arg1 **C.gchar
	var arg2 *C.gsize
	var arg3 **C.gchar
	var arg4 *C.gsize

	arg0 = (C.GIConv)(unsafe.Pointer(c.Native()))
	arg1 = (**C.gchar)(C.CString(inbuf))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = *C.gsize(inbytesLeft)
	arg3 = (**C.gchar)(C.CString(outbuf))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = *C.gsize(outbytesLeft)

	var cret C.gsize
	var ret1 uint

	cret = C.g_iconv(arg0, inbuf, inbytesLeft, outbuf, outbytesLeft)

	ret1 = C.gsize(cret)

	return ret1
}

// Close: same as the standard UNIX routine iconv_close(), but may be
// implemented via libiconv on UNIX flavors that lack a native implementation.
// Should be called to clean up the conversion descriptor from g_iconv_open()
// when you are done converting things.
//
// GLib provides g_convert() and g_locale_to_utf8() which are likely more
// convenient than the raw iconv wrappers.
func (c *IConv) Close() int {
	var arg0 C.GIConv

	arg0 = (C.GIConv)(unsafe.Pointer(c.Native()))

	var cret C.gint
	var ret1 int

	cret = C.g_iconv_close(arg0)

	ret1 = C.gint(cret)

	return ret1
}