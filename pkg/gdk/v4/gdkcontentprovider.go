// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_provider_get_type()), F: marshalContentProviderrer},
	})
}

// ContentProviderrerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ContentProviderrerOverrider interface {
	AttachClipboard(clipboard Clipboarder)
	// ContentChanged emits the ::content-changed signal.
	ContentChanged()
	DetachClipboard(clipboard Clipboarder)
	// Value gets the contents of @provider stored in @value.
	//
	// The @value will have been initialized to the `GType` the value should be
	// provided in. This given `GType` does not need to be listed in the formats
	// returned by [method@Gdk.ContentProvider.ref_formats]. However, if the
	// given `GType` is not supported, this operation can fail and
	// IO_ERROR_NOT_SUPPORTED will be reported.
	Value(value *externglib.Value) error
	// RefFormats gets the formats that the provider can provide its current
	// contents in.
	RefFormats() *ContentFormats
	// RefStorableFormats gets the formats that the provider suggests other
	// applications to store the data in.
	//
	// An example of such an application would be a clipboard manager.
	//
	// This can be assumed to be a subset of
	// [method@Gdk.ContentProvider.ref_formats].
	RefStorableFormats() *ContentFormats
}

// ContentProviderrer describes ContentProvider's methods.
type ContentProviderrer interface {
	gextras.Objector

	ContentChanged()
	Value(value *externglib.Value) error
	RefFormats() *ContentFormats
	RefStorableFormats() *ContentFormats
}

// ContentProvider: `GdkContentProvider` is used to provide content for the
// clipboard or for drag-and-drop operations in a number of formats.
//
// To create a `GdkContentProvider`, use
// [ctor@Gdk.ContentProvider.new_for_value] or
// [ctor@Gdk.ContentProvider.new_for_bytes].
//
// GDK knows how to handle common text and image formats out-of-the-box. See
// [class@Gdk.ContentSerializer] and [class@Gdk.ContentDeserializer] if you want
// to add support for application-specific data formats.
type ContentProvider struct {
	*externglib.Object
}

var _ ContentProviderrer = (*ContentProvider)(nil)

func wrapContentProviderrer(obj *externglib.Object) ContentProviderrer {
	return &ContentProvider{
		Object: obj,
	}
}

func marshalContentProviderrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapContentProviderrer(obj), nil
}

// NewContentProviderForValue: create a content provider that provides the given
// @value.
func NewContentProviderForValue(value *externglib.Value) *ContentProvider {
	var _arg1 *C.GValue             // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gdk_content_provider_new_for_value(_arg1)

	var _contentProvider *ContentProvider // out

	_contentProvider = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ContentProvider)

	return _contentProvider
}

// NewContentProviderUnion creates a content provider that represents all the
// given @providers.
//
// Whenever data needs to be written, the union provider will try the given
// @providers in the given order and the first one supporting a format will be
// chosen to provide it.
//
// This allows an easy way to support providing data in different formats. For
// example, an image may be provided by its file and by the image contents with
// a call such as “`c gdk_content_provider_new_union ((GdkContentProvider *[2])
// { gdk_content_provider_new_typed (G_TYPE_FILE, file),
// gdk_content_provider_new_typed (G_TYPE_TEXTURE, texture) }, 2); “`
func NewContentProviderUnion(providers []*ContentProvider) *ContentProvider {
	var _arg1 **C.GdkContentProvider
	var _arg2 C.gsize
	var _cret *C.GdkContentProvider // in

	_arg2 = C.gsize(len(providers))
	_arg1 = (**C.GdkContentProvider)(C.malloc(C.ulong(len(providers)) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice(_arg1, len(providers))
		for i := range providers {
			out[i] = (*C.GdkContentProvider)(unsafe.Pointer(providers[i].Native()))
		}
	}

	_cret = C.gdk_content_provider_new_union(_arg1, _arg2)

	var _contentProvider *ContentProvider // out

	_contentProvider = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ContentProvider)

	return _contentProvider
}

// ContentChanged emits the ::content-changed signal.
func (provider *ContentProvider) ContentChanged() {
	var _arg0 *C.GdkContentProvider // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	C.gdk_content_provider_content_changed(_arg0)
}

// Value gets the contents of @provider stored in @value.
//
// The @value will have been initialized to the `GType` the value should be
// provided in. This given `GType` does not need to be listed in the formats
// returned by [method@Gdk.ContentProvider.ref_formats]. However, if the given
// `GType` is not supported, this operation can fail and IO_ERROR_NOT_SUPPORTED
// will be reported.
func (provider *ContentProvider) Value(value *externglib.Value) error {
	var _arg0 *C.GdkContentProvider // out
	var _arg1 *C.GValue             // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gdk_content_provider_get_value(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RefFormats gets the formats that the provider can provide its current
// contents in.
func (provider *ContentProvider) RefFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gdk_content_provider_ref_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(unsafe.Pointer(v)))
	})

	return _contentFormats
}

// RefStorableFormats gets the formats that the provider suggests other
// applications to store the data in.
//
// An example of such an application would be a clipboard manager.
//
// This can be assumed to be a subset of
// [method@Gdk.ContentProvider.ref_formats].
func (provider *ContentProvider) RefStorableFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gdk_content_provider_ref_storable_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.gdk_content_formats_unref((*C.GdkContentFormats)(unsafe.Pointer(v)))
	})

	return _contentFormats
}
