// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_error_get_type()), F: marshalError},
	})
}

// Error: the `GError` structure contains information about an error that has
// occurred.
type Error C.GError

// WrapError wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapError(ptr unsafe.Pointer) *Error {
	return (*Error)(ptr)
}

func marshalError(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Error)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *Error) Native() unsafe.Pointer {
	return unsafe.Pointer(e)
}

// Copy returns true if @error matches @domain and @code, false otherwise. In
// particular, when @error is nil, false will be returned.
//
// If @domain contains a `FAILED` (or otherwise generic) error code, you should
// generally not check for it explicitly, but should instead treat any
// not-explicitly-recognized error code as being equivalent to the `FAILED`
// code. This way, if the domain is extended in the future to provide a more
// specific error code for a certain case, your code will still work.
func (e *Error) Copy() error {
	var _arg0 *C.GError // out
	var _cret *C.GError // in

	_arg0 = (*C.GError)(gerror.New(e))
	defer C.g_error_free(_arg0)

	_cret = C.g_error_copy(_arg0)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

// Free returns true if @error matches @domain and @code, false otherwise. In
// particular, when @error is nil, false will be returned.
//
// If @domain contains a `FAILED` (or otherwise generic) error code, you should
// generally not check for it explicitly, but should instead treat any
// not-explicitly-recognized error code as being equivalent to the `FAILED`
// code. This way, if the domain is extended in the future to provide a more
// specific error code for a certain case, your code will still work.
func (e *Error) Free() {
	var _arg0 *C.GError // out

	_arg0 = (*C.GError)(gerror.New(e))
	defer C.g_error_free(_arg0)

	C.g_error_free(_arg0)
}