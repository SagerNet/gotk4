// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"context"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_serializer_get_type()), F: marshalContentSerializerer},
	})
}

// ContentSerializeAsync: serialize content and write it to the given output
// stream, asynchronously.
//
// The default I/O priority is G_PRIORITY_DEFAULT (i.e. 0), and lower numbers
// indicate a higher priority.
//
// When the operation is finished, callback will be called. You must then call
// content_serialize_finish to get the result of the operation.
func ContentSerializeAsync(ctx context.Context, stream gio.OutputStreamer, mimeType string, value *externglib.Value, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg5 *C.GCancellable       // out
	var _arg1 *C.GOutputStream      // out
	var _arg2 *C.char               // out
	var _arg3 *C.GValue             // out
	var _arg4 C.int                 // out
	var _arg6 C.GAsyncReadyCallback // out
	var _arg7 C.gpointer

	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GOutputStream)(unsafe.Pointer(stream.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))
	_arg4 = C.int(ioPriority)
	if callback != nil {
		_arg6 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg7 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_content_serialize_async(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(mimeType)
	runtime.KeepAlive(value)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ContentSerializeFinish finishes a content serialization operation.
func ContentSerializeFinish(result gio.AsyncResulter) error {
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gdk_content_serialize_finish(_arg1, &_cerr)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ContentSerializer: GdkContentSerializer is used to serialize content for
// inter-application data transfers.
//
// The GdkContentSerializer transforms an object that is identified by a GType
// into a serialized form (i.e. a byte stream) that is identified by a mime
// type.
//
// GTK provides serializers and deserializers for common data types such as
// text, colors, images or file lists. To register your own serialization
// functions, use content_register_serializer.
//
// Also see gdk.ContentDeserializer.
type ContentSerializer struct {
	*externglib.Object

	gio.AsyncResult
}

func WrapContentSerializer(obj *externglib.Object) *ContentSerializer {
	return &ContentSerializer{
		Object: obj,
		AsyncResult: gio.AsyncResult{
			Object: obj,
		},
	}
}

func marshalContentSerializerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContentSerializer(obj), nil
}

// Cancellable gets the cancellable for the current operation.
//
// This is the GCancellable that was passed to [content_serialize_async].
func (serializer *ContentSerializer) Cancellable() *gio.Cancellable {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GCancellable         // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_cancellable(_arg0)
	runtime.KeepAlive(serializer)

	var _cancellable *gio.Cancellable // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_cancellable = &gio.Cancellable{
			Object: obj,
		}
	}

	return _cancellable
}

// GType gets the GType to of the object to serialize.
func (serializer *ContentSerializer) GType() externglib.Type {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.GType                 // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_gtype(_arg0)
	runtime.KeepAlive(serializer)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// MIMEType gets the mime type to serialize to.
func (serializer *ContentSerializer) MIMEType() string {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.char                 // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_mime_type(_arg0)
	runtime.KeepAlive(serializer)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// OutputStream gets the output stream for the current operation.
//
// This is the stream that was passed to content_serialize_async.
func (serializer *ContentSerializer) OutputStream() gio.OutputStreamer {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GOutputStream        // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_output_stream(_arg0)
	runtime.KeepAlive(serializer)

	var _outputStream gio.OutputStreamer // out

	_outputStream = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.OutputStreamer)

	return _outputStream
}

// Priority gets the I/O priority for the current operation.
//
// This is the priority that was passed to content_serialize_async.
func (serializer *ContentSerializer) Priority() int {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.int                   // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_priority(_arg0)
	runtime.KeepAlive(serializer)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TaskData gets the data that was associated with the current operation.
//
// See gdk.ContentSerializer.SetTaskData().
func (serializer *ContentSerializer) TaskData() cgo.Handle {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.gpointer              // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_task_data(_arg0)
	runtime.KeepAlive(serializer)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// UserData gets the user data that was passed when the serializer was
// registered.
func (serializer *ContentSerializer) UserData() cgo.Handle {
	var _arg0 *C.GdkContentSerializer // out
	var _cret C.gpointer              // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_user_data(_arg0)
	runtime.KeepAlive(serializer)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// Value gets the GValue to read the object to serialize from.
func (serializer *ContentSerializer) Value() *externglib.Value {
	var _arg0 *C.GdkContentSerializer // out
	var _cret *C.GValue               // in

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	_cret = C.gdk_content_serializer_get_value(_arg0)
	runtime.KeepAlive(serializer)

	var _value *externglib.Value // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _value
}

// ReturnError: indicate that the serialization has ended with an error.
//
// This function consumes error.
func (serializer *ContentSerializer) ReturnError(err error) {
	var _arg0 *C.GdkContentSerializer // out
	var _arg1 *C.GError               // out

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))
	_arg1 = (*C.GError)(gerror.New(err))

	C.gdk_content_serializer_return_error(_arg0, _arg1)
	runtime.KeepAlive(serializer)
	runtime.KeepAlive(err)
}

// ReturnSuccess: indicate that the serialization has been successfully
// completed.
func (serializer *ContentSerializer) ReturnSuccess() {
	var _arg0 *C.GdkContentSerializer // out

	_arg0 = (*C.GdkContentSerializer)(unsafe.Pointer(serializer.Native()))

	C.gdk_content_serializer_return_success(_arg0)
	runtime.KeepAlive(serializer)
}
