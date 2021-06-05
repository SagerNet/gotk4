// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// ShellError: error codes returned by shell functions.
type ShellError int

const (
	// ShellErrorBadQuoting: mismatched or otherwise mangled quoting.
	ShellErrorBadQuoting ShellError = 0
	// ShellErrorEmptyString: string to be parsed was empty.
	ShellErrorEmptyString ShellError = 1
	// ShellErrorFailed: some other error.
	ShellErrorFailed ShellError = 2
)

// ShellQuote quotes a string so that the shell (/bin/sh) will interpret the
// quoted string to mean @unquoted_string. If you pass a filename to the shell,
// for example, you should first quote it with this function. The return value
// must be freed with g_free(). The quoting style used is undefined (single or
// double quotes may be used).
func ShellQuote(unquotedString string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(unquotedString))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret1 string

	cret = C.g_shell_quote(unquotedString)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// ShellUnquote unquotes a string as the shell (/bin/sh) would. Only handles
// quotes; if a string contains file globs, arithmetic operators, variables,
// backticks, redirections, or other special-to-the-shell features, the result
// will be different from the result a real shell would produce (the variables,
// backticks, etc. will be passed through literally instead of being expanded).
// This function is guaranteed to succeed if applied to the result of
// g_shell_quote(). If it fails, it returns nil and sets the error. The
// @quoted_string need not actually contain quoted or escaped text;
// g_shell_unquote() simply goes through the string and unquotes/unescapes
// anything that the shell would. Both single and double quotes are handled, as
// are escapes including escaped newlines. The return value must be freed with
// g_free(). Possible errors are in the SHELL_ERROR domain.
//
// Shell quoting rules are a bit strange. Single quotes preserve the literal
// string exactly. escape sequences are not allowed; not even \' - if you want a
// ' in the quoted text, you have to do something like 'foo'\”bar'. Double
// quotes allow $, `, ", \, and newline to be escaped with backslash. Otherwise
// double quotes preserve things literally.
func ShellUnquote(quotedString string) (filename string, err error) {
	var arg1 *C.gchar
	var errout *C.GError

	arg1 = (*C.gchar)(C.CString(quotedString))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret1 string
	var goerr error

	cret = C.g_shell_unquote(quotedString, &errout)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}
