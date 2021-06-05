// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_formats_builder_get_type()), F: marshalContentFormatsBuilder},
	})
}

// InternMIMEType canonicalizes the given mime type and interns the result.
//
// If @string is not a valid mime type, nil is returned instead. See RFC 2048
// for the syntax if mime types.
func InternMIMEType(string string) string {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.char
	var ret1 string

	cret = C.gdk_intern_mime_type(string)

	ret1 = C.GoString(cret)

	return ret1
}

// ContentFormatsBuilder: a ContentFormatsBuilder struct is an opaque struct. It
// is meant to not be kept around and only be used to create new ContentFormats
// objects.
type ContentFormatsBuilder struct {
	native C.GdkContentFormatsBuilder
}

// WrapContentFormatsBuilder wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapContentFormatsBuilder(ptr unsafe.Pointer) *ContentFormatsBuilder {
	if ptr == nil {
		return nil
	}

	return (*ContentFormatsBuilder)(ptr)
}

func marshalContentFormatsBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapContentFormatsBuilder(unsafe.Pointer(b)), nil
}

// NewContentFormatsBuilder constructs a struct ContentFormatsBuilder.
func NewContentFormatsBuilder() *ContentFormatsBuilder {
	var cret *C.GdkContentFormatsBuilder
	var ret1 *ContentFormatsBuilder

	cret = C.gdk_content_formats_builder_new()

	ret1 = WrapContentFormatsBuilder(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ContentFormatsBuilder) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (c *ContentFormatsBuilder) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// AddFormats appends all formats from @formats to @builder, skipping those that
// already exist.
func (b *ContentFormatsBuilder) AddFormats(formats *ContentFormats) {
	var arg0 *C.GdkContentFormatsBuilder
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(formats.Native()))

	C.gdk_content_formats_builder_add_formats(arg0, formats)
}

// AddGType appends @gtype to @builder if it has not already been added.
func (b *ContentFormatsBuilder) AddGType(typ externglib.Type) {
	var arg0 *C.GdkContentFormatsBuilder
	var arg1 C.GType

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))
	arg1 := C.GType(typ)

	C.gdk_content_formats_builder_add_gtype(arg0, typ)
}

// AddMIMEType appends @mime_type to @builder if it has not already been added.
func (b *ContentFormatsBuilder) AddMIMEType(mimeType string) {
	var arg0 *C.GdkContentFormatsBuilder
	var arg1 *C.char

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_content_formats_builder_add_mime_type(arg0, mimeType)
}

// FreeToFormats creates a new ContentFormats from the current state of the
// given @builder, and frees the @builder instance.
func (b *ContentFormatsBuilder) FreeToFormats() *ContentFormats {
	var arg0 *C.GdkContentFormatsBuilder

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GdkContentFormats
	var ret1 *ContentFormats

	cret = C.gdk_content_formats_builder_free_to_formats(arg0)

	ret1 = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Ref acquires a reference on the given @builder.
//
// This function is intended primarily for bindings. ContentFormatsBuilder
// objects should not be kept around.
func (b *ContentFormatsBuilder) Ref() *ContentFormatsBuilder {
	var arg0 *C.GdkContentFormatsBuilder

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GdkContentFormatsBuilder
	var ret1 *ContentFormatsBuilder

	cret = C.gdk_content_formats_builder_ref(arg0)

	ret1 = WrapContentFormatsBuilder(unsafe.Pointer(cret))

	return ret1
}

// ToFormats creates a new ContentFormats from the given @builder.
//
// The given ContentFormatsBuilder is reset once this function returns; you
// cannot call this function multiple times on the same @builder instance.
//
// This function is intended primarily for bindings. C code should use
// gdk_content_formats_builder_free_to_formats().
func (b *ContentFormatsBuilder) ToFormats() *ContentFormats {
	var arg0 *C.GdkContentFormatsBuilder

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GdkContentFormats
	var ret1 *ContentFormats

	cret = C.gdk_content_formats_builder_to_formats(arg0)

	ret1 = WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Unref releases a reference on the given @builder.
func (b *ContentFormatsBuilder) Unref() {
	var arg0 *C.GdkContentFormatsBuilder

	arg0 = (*C.GdkContentFormatsBuilder)(unsafe.Pointer(b.Native()))

	C.gdk_content_formats_builder_unref(arg0)
}