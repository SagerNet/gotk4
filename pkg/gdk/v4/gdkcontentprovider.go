// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_provider_get_type()), F: marshalContentProviderer},
	})
}

// ContentProviderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ContentProviderOverrider interface {
	AttachClipboard(clipboard *Clipboard)
	// ContentChanged emits the ::content-changed signal.
	ContentChanged()
	DetachClipboard(clipboard *Clipboard)
	// Value gets the contents of provider stored in value.
	//
	// The value will have been initialized to the GType the value should be
	// provided in. This given GType does not need to be listed in the formats
	// returned by gdk.ContentProvider.RefFormats(). However, if the given GType
	// is not supported, this operation can fail and IO_ERROR_NOT_SUPPORTED will
	// be reported.
	Value(value *externglib.Value) error
	// RefFormats gets the formats that the provider can provide its current
	// contents in.
	RefFormats() *ContentFormats
	// RefStorableFormats gets the formats that the provider suggests other
	// applications to store the data in.
	//
	// An example of such an application would be a clipboard manager.
	//
	// This can be assumed to be a subset of gdk.ContentProvider.RefFormats().
	RefStorableFormats() *ContentFormats
	// WriteMIMETypeAsync: asynchronously writes the contents of provider to
	// stream in the given mime_type.
	//
	// When the operation is finished callback will be called. You must then
	// call gdk.ContentProvider.WriteMIMETypeFinish() to get the result of the
	// operation.
	//
	// The given mime type does not need to be listed in the formats returned by
	// gdk.ContentProvider.RefFormats(). However, if the given GType is not
	// supported, IO_ERROR_NOT_SUPPORTED will be reported.
	//
	// The given stream will not be closed.
	WriteMIMETypeAsync(ctx context.Context, mimeType string, stream gio.OutputStreamer, ioPriority int, callback gio.AsyncReadyCallback)
	// WriteMIMETypeFinish finishes an asynchronous write operation.
	//
	// See gdk.ContentProvider.WriteMIMETypeAsync().
	WriteMIMETypeFinish(result gio.AsyncResulter) error
}

// ContentProvider: GdkContentProvider is used to provide content for the
// clipboard or for drag-and-drop operations in a number of formats.
//
// To create a GdkContentProvider, use gdk.ContentProvider.NewForValue or
// gdk.ContentProvider.NewForBytes.
//
// GDK knows how to handle common text and image formats out-of-the-box. See
// gdk.ContentSerializer and gdk.ContentDeserializer if you want to add support
// for application-specific data formats.
type ContentProvider struct {
	*externglib.Object
}

func WrapContentProvider(obj *externglib.Object) *ContentProvider {
	return &ContentProvider{
		Object: obj,
	}
}

func marshalContentProviderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContentProvider(obj), nil
}

// NewContentProviderForBytes: create a content provider that provides the given
// bytes as data for the given mime_type.
func NewContentProviderForBytes(mimeType string, bytes *glib.Bytes) *ContentProvider {
	var _arg1 *C.char               // out
	var _arg2 *C.GBytes             // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))

	_cret = C.gdk_content_provider_new_for_bytes(_arg1, _arg2)
	runtime.KeepAlive(mimeType)
	runtime.KeepAlive(bytes)

	var _contentProvider *ContentProvider // out

	_contentProvider = WrapContentProvider(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contentProvider
}

// NewContentProviderForValue: create a content provider that provides the given
// value.
func NewContentProviderForValue(value *externglib.Value) *ContentProvider {
	var _arg1 *C.GValue             // out
	var _cret *C.GdkContentProvider // in

	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	_cret = C.gdk_content_provider_new_for_value(_arg1)
	runtime.KeepAlive(value)

	var _contentProvider *ContentProvider // out

	_contentProvider = WrapContentProvider(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contentProvider
}

// NewContentProviderUnion creates a content provider that represents all the
// given providers.
//
// Whenever data needs to be written, the union provider will try the given
// providers in the given order and the first one supporting a format will be
// chosen to provide it.
//
// This allows an easy way to support providing data in different formats. For
// example, an image may be provided by its file and by the image contents with
// a call such as
//
//    gdk_content_provider_new_union ((GdkContentProvider *[2]) {
//                                      gdk_content_provider_new_typed (G_TYPE_FILE, file),
//                                      gdk_content_provider_new_typed (G_TYPE_TEXTURE, texture)
//                                    }, 2);
func NewContentProviderUnion(providers []*ContentProvider) *ContentProvider {
	var _arg1 **C.GdkContentProvider // out
	var _arg2 C.gsize
	var _cret *C.GdkContentProvider // in

	if providers != nil {
		_arg2 = (C.gsize)(len(providers))
		_arg1 = (**C.GdkContentProvider)(C.malloc(C.ulong(len(providers)) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice((**C.GdkContentProvider)(_arg1), len(providers))
			for i := range providers {
				out[i] = (*C.GdkContentProvider)(unsafe.Pointer(providers[i].Native()))
				C.g_object_ref(C.gpointer(providers[i].Native()))
			}
		}
	}

	_cret = C.gdk_content_provider_new_union(_arg1, _arg2)
	runtime.KeepAlive(providers)

	var _contentProvider *ContentProvider // out

	_contentProvider = WrapContentProvider(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _contentProvider
}

// ContentChanged emits the ::content-changed signal.
func (provider *ContentProvider) ContentChanged() {
	var _arg0 *C.GdkContentProvider // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	C.gdk_content_provider_content_changed(_arg0)
	runtime.KeepAlive(provider)
}

// Value gets the contents of provider stored in value.
//
// The value will have been initialized to the GType the value should be
// provided in. This given GType does not need to be listed in the formats
// returned by gdk.ContentProvider.RefFormats(). However, if the given GType is
// not supported, this operation can fail and IO_ERROR_NOT_SUPPORTED will be
// reported.
func (provider *ContentProvider) Value(value *externglib.Value) error {
	var _arg0 *C.GdkContentProvider // out
	var _arg1 *C.GValue             // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gdk_content_provider_get_value(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(value)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// RefFormats gets the formats that the provider can provide its current
// contents in.
func (provider *ContentProvider) RefFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gdk_content_provider_ref_formats(_arg0)
	runtime.KeepAlive(provider)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// RefStorableFormats gets the formats that the provider suggests other
// applications to store the data in.
//
// An example of such an application would be a clipboard manager.
//
// This can be assumed to be a subset of gdk.ContentProvider.RefFormats().
func (provider *ContentProvider) RefStorableFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gdk_content_provider_ref_storable_formats(_arg0)
	runtime.KeepAlive(provider)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// WriteMIMETypeAsync: asynchronously writes the contents of provider to stream
// in the given mime_type.
//
// When the operation is finished callback will be called. You must then call
// gdk.ContentProvider.WriteMIMETypeFinish() to get the result of the operation.
//
// The given mime type does not need to be listed in the formats returned by
// gdk.ContentProvider.RefFormats(). However, if the given GType is not
// supported, IO_ERROR_NOT_SUPPORTED will be reported.
//
// The given stream will not be closed.
func (provider *ContentProvider) WriteMIMETypeAsync(ctx context.Context, mimeType string, stream gio.OutputStreamer, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkContentProvider // out
	var _arg4 *C.GCancellable       // out
	var _arg1 *C.char               // out
	var _arg2 *C.GOutputStream      // out
	var _arg3 C.int                 // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GOutputStream)(unsafe.Pointer(stream.Native()))
	_arg3 = C.int(ioPriority)
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_content_provider_write_mime_type_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(mimeType)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// WriteMIMETypeFinish finishes an asynchronous write operation.
//
// See gdk.ContentProvider.WriteMIMETypeAsync().
func (provider *ContentProvider) WriteMIMETypeFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GdkContentProvider // out
	var _arg1 *C.GAsyncResult       // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gdk_content_provider_write_mime_type_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(provider)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
