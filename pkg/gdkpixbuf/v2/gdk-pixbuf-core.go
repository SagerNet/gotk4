// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"context"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_colorspace_get_type()), F: marshalColorspace},
		{T: externglib.Type(C.gdk_pixbuf_alpha_mode_get_type()), F: marshalPixbufAlphaMode},
		{T: externglib.Type(C.gdk_pixbuf_error_get_type()), F: marshalPixbufError},
	})
}

// Colorspace: this enumeration defines the color spaces that are supported by
// the gdk-pixbuf library.
//
// Currently only RGB is supported.
type Colorspace int

const (
	// ColorspaceRGB indicates a red/green/blue additive color space.
	ColorspaceRGB Colorspace = iota
)

func marshalColorspace(p uintptr) (interface{}, error) {
	return Colorspace(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for Colorspace.
func (c Colorspace) String() string {
	switch c {
	case ColorspaceRGB:
		return "RGB"
	default:
		return fmt.Sprintf("Colorspace(%d)", c)
	}
}

// PixbufAlphaMode: control the alpha channel for drawables.
//
// These values can be passed to gdk_pixbuf_xlib_render_to_drawable_alpha() in
// gdk-pixbuf-xlib to control how the alpha channel of an image should be
// handled.
//
// This function can create a bilevel clipping mask (black and white) and use it
// while painting the image.
//
// In the future, when the X Window System gets an alpha channel extension, it
// will be possible to do full alpha compositing onto arbitrary drawables. For
// now both cases fall back to a bilevel clipping mask.
//
// Deprecated: There is no user of GdkPixbufAlphaMode in GdkPixbuf, and the Xlib
// utility functions have been split out to their own library, gdk-pixbuf-xlib.
type PixbufAlphaMode int

const (
	// PixbufAlphaBilevel clipping mask (black and white) will be created and
	// used to draw the image. Pixels below 0.5 opacity will be considered fully
	// transparent, and all others will be considered fully opaque.
	PixbufAlphaBilevel PixbufAlphaMode = iota
	// PixbufAlphaFull: for now falls back to K_PIXBUF_ALPHA_BILEVEL. In the
	// future it will do full alpha compositing.
	PixbufAlphaFull
)

func marshalPixbufAlphaMode(p uintptr) (interface{}, error) {
	return PixbufAlphaMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PixbufAlphaMode.
func (p PixbufAlphaMode) String() string {
	switch p {
	case PixbufAlphaBilevel:
		return "Bilevel"
	case PixbufAlphaFull:
		return "Full"
	default:
		return fmt.Sprintf("PixbufAlphaMode(%d)", p)
	}
}

// PixbufError: error code in the GDK_PIXBUF_ERROR domain.
//
// Many gdk-pixbuf operations can cause errors in this domain, or in the
// G_FILE_ERROR domain.
type PixbufError int

const (
	// PixbufErrorCorruptImage: image file was broken somehow.
	PixbufErrorCorruptImage PixbufError = iota
	// PixbufErrorInsufficientMemory: not enough memory.
	PixbufErrorInsufficientMemory
	// PixbufErrorBadOption: bad option was passed to a pixbuf save module.
	PixbufErrorBadOption
	// PixbufErrorUnknownType: unknown image type.
	PixbufErrorUnknownType
	// PixbufErrorUnsupportedOperation: don't know how to perform the given
	// operation on the type of image at hand.
	PixbufErrorUnsupportedOperation
	// PixbufErrorFailed: generic failure code, something went wrong.
	PixbufErrorFailed
	// PixbufErrorIncompleteAnimation: only part of the animation was loaded.
	PixbufErrorIncompleteAnimation
)

func marshalPixbufError(p uintptr) (interface{}, error) {
	return PixbufError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PixbufError.
func (p PixbufError) String() string {
	switch p {
	case PixbufErrorCorruptImage:
		return "CorruptImage"
	case PixbufErrorInsufficientMemory:
		return "InsufficientMemory"
	case PixbufErrorBadOption:
		return "BadOption"
	case PixbufErrorUnknownType:
		return "UnknownType"
	case PixbufErrorUnsupportedOperation:
		return "UnsupportedOperation"
	case PixbufErrorFailed:
		return "Failed"
	case PixbufErrorIncompleteAnimation:
		return "IncompleteAnimation"
	default:
		return fmt.Sprintf("PixbufError(%d)", p)
	}
}

// PixbufSaveFunc: save functions used by gdkpixbuf.Pixbuf.SaveToCallback().
//
// This function is called once for each block of bytes that is "written" by
// gdk_pixbuf_save_to_callback().
//
// If successful it should return TRUE; if an error occurs it should set error
// and return FALSE, in which case gdk_pixbuf_save_to_callback() will fail with
// the same error.
type PixbufSaveFunc func(buf []byte) (err error, ok bool)

//export _gotk4_gdkpixbuf2_PixbufSaveFunc
func _gotk4_gdkpixbuf2_PixbufSaveFunc(arg0 *C.gchar, arg1 C.gsize, arg2 **C.GError, arg3 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var buf []byte // out

	defer C.free(unsafe.Pointer(arg0))
	buf = make([]byte, arg1)
	copy(buf, unsafe.Slice((*byte)(unsafe.Pointer(arg0)), arg1))

	fn := v.(PixbufSaveFunc)
	err, ok := fn(buf)

	*arg2 = (*C.GError)(gerror.New(err))
	if ok {
		cret = C.TRUE
	}

	return cret
}

// PixbufCalculateRowstride calculates the rowstride that an image created with
// those values would have.
//
// This function is useful for front-ends and backends that want to check image
// values without needing to create a GdkPixbuf.
func PixbufCalculateRowstride(colorspace Colorspace, hasAlpha bool, bitsPerSample int, width int, height int) int {
	var _arg1 C.GdkColorspace // out
	var _arg2 C.gboolean      // out
	var _arg3 C.int           // out
	var _arg4 C.int           // out
	var _arg5 C.int           // out
	var _cret C.gint          // in

	_arg1 = C.GdkColorspace(colorspace)
	if hasAlpha {
		_arg2 = C.TRUE
	}
	_arg3 = C.int(bitsPerSample)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	_cret = C.gdk_pixbuf_calculate_rowstride(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NewPixbufFromStreamAsync creates a new pixbuf by asynchronously loading an
// image from an input stream.
//
// For more details see gdk_pixbuf_new_from_stream(), which is the synchronous
// version of this function.
//
// When the operation is finished, callback will be called in the main thread.
// You can then call gdk_pixbuf_new_from_stream_finish() to get the result of
// the operation.
func NewPixbufFromStreamAsync(ctx context.Context, stream gio.InputStreamer, callback gio.AsyncReadyCallback) {
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GInputStream       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GInputStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))
	_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg4 = C.gpointer(gbox.AssignOnce(callback))

	C.gdk_pixbuf_new_from_stream_async(_arg1, _arg2, _arg3, _arg4)
}

// NewPixbufFromStreamAtScaleAsync creates a new pixbuf by asynchronously
// loading an image from an input stream.
//
// For more details see gdk_pixbuf_new_from_stream_at_scale(), which is the
// synchronous version of this function.
//
// When the operation is finished, callback will be called in the main thread.
// You can then call gdk_pixbuf_new_from_stream_finish() to get the result of
// the operation.
func NewPixbufFromStreamAtScaleAsync(ctx context.Context, stream gio.InputStreamer, width int, height int, preserveAspectRatio bool, callback gio.AsyncReadyCallback) {
	var _arg5 *C.GCancellable       // out
	var _arg1 *C.GInputStream       // out
	var _arg2 C.gint                // out
	var _arg3 C.gint                // out
	var _arg4 C.gboolean            // out
	var _arg6 C.GAsyncReadyCallback // out
	var _arg7 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GInputStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))
	_arg2 = C.gint(width)
	_arg3 = C.gint(height)
	if preserveAspectRatio {
		_arg4 = C.TRUE
	}
	_arg6 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
	_arg7 = C.gpointer(gbox.AssignOnce(callback))

	C.gdk_pixbuf_new_from_stream_at_scale_async(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// PixbufSaveToStreamFinish finishes an asynchronous pixbuf save operation
// started with gdk_pixbuf_save_to_stream_async().
func PixbufSaveToStreamFinish(asyncResult gio.AsyncResulter) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer((asyncResult).(gextras.Nativer).Native()))

	C.gdk_pixbuf_save_to_stream_finish(_arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
