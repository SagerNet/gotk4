// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_data_output_stream_get_type()), F: marshalDataOutputStream},
	})
}

// DataOutputStream: data output stream implements Stream and includes functions
// for writing data directly to an output stream.
type DataOutputStream interface {
	FilterOutputStream
	Seekable

	// ByteOrder gets the byte order for the stream.
	ByteOrder() DataStreamByteOrder
	// PutByte puts a byte into the output stream.
	PutByte(data byte, cancellable Cancellable) error
	// PutInt16 puts a signed 16-bit integer into the output stream.
	PutInt16(data int16, cancellable Cancellable) error
	// PutInt32 puts a signed 32-bit integer into the output stream.
	PutInt32(data int32, cancellable Cancellable) error
	// PutInt64 puts a signed 64-bit integer into the stream.
	PutInt64(data int64, cancellable Cancellable) error
	// PutString puts a string into the output stream.
	PutString(str string, cancellable Cancellable) error
	// PutUint16 puts an unsigned 16-bit integer into the output stream.
	PutUint16(data uint16, cancellable Cancellable) error
	// PutUint32 puts an unsigned 32-bit integer into the stream.
	PutUint32(data uint32, cancellable Cancellable) error
	// PutUint64 puts an unsigned 64-bit integer into the stream.
	PutUint64(data uint64, cancellable Cancellable) error
	// SetByteOrder sets the byte order of the data output stream to @order.
	SetByteOrder(order DataStreamByteOrder)
}

// dataOutputStream implements the DataOutputStream interface.
type dataOutputStream struct {
	FilterOutputStream
	Seekable
}

var _ DataOutputStream = (*dataOutputStream)(nil)

// WrapDataOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapDataOutputStream(obj *externglib.Object) DataOutputStream {
	return DataOutputStream{
		FilterOutputStream: WrapFilterOutputStream(obj),
		Seekable:           WrapSeekable(obj),
	}
}

func marshalDataOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDataOutputStream(obj), nil
}

// NewDataOutputStream constructs a class DataOutputStream.
func NewDataOutputStream(baseStream OutputStream) DataOutputStream {
	var arg1 *C.GOutputStream

	arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))

	var cret C.GDataOutputStream
	var goret1 DataOutputStream

	cret = C.g_data_output_stream_new(baseStream)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DataOutputStream)

	return goret1
}

// ByteOrder gets the byte order for the stream.
func (s dataOutputStream) ByteOrder() DataStreamByteOrder {
	var arg0 *C.GDataOutputStream

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.GDataStreamByteOrder
	var goret1 DataStreamByteOrder

	cret = C.g_data_output_stream_get_byte_order(arg0)

	goret1 = DataStreamByteOrder(cret)

	return goret1
}

// PutByte puts a byte into the output stream.
func (s dataOutputStream) PutByte(data byte, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.guchar
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.guchar(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_byte(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutInt16 puts a signed 16-bit integer into the output stream.
func (s dataOutputStream) PutInt16(data int16, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.gint16
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gint16(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_int16(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutInt32 puts a signed 32-bit integer into the output stream.
func (s dataOutputStream) PutInt32(data int32, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.gint32
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gint32(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_int32(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutInt64 puts a signed 64-bit integer into the stream.
func (s dataOutputStream) PutInt64(data int64, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.gint64
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.gint64(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_int64(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutString puts a string into the output stream.
func (s dataOutputStream) PutString(str string, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 *C.char
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_string(arg0, str, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutUint16 puts an unsigned 16-bit integer into the output stream.
func (s dataOutputStream) PutUint16(data uint16, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.guint16
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.guint16(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_uint16(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutUint32 puts an unsigned 32-bit integer into the stream.
func (s dataOutputStream) PutUint32(data uint32, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.guint32
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.guint32(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_uint32(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// PutUint64 puts an unsigned 64-bit integer into the stream.
func (s dataOutputStream) PutUint64(data uint64, cancellable Cancellable) error {
	var arg0 *C.GDataOutputStream
	var arg1 C.guint64
	var arg2 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = C.guint64(data)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var goerr error

	C.g_data_output_stream_put_uint64(arg0, data, cancellable, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// SetByteOrder sets the byte order of the data output stream to @order.
func (s dataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	var arg0 *C.GDataOutputStream
	var arg1 C.GDataStreamByteOrder

	arg0 = (*C.GDataOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (C.GDataStreamByteOrder)(order)

	C.g_data_output_stream_set_byte_order(arg0, order)
}

type DataOutputStreamPrivate struct {
	native C.GDataOutputStreamPrivate
}

// WrapDataOutputStreamPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDataOutputStreamPrivate(ptr unsafe.Pointer) *DataOutputStreamPrivate {
	if ptr == nil {
		return nil
	}

	return (*DataOutputStreamPrivate)(ptr)
}

func marshalDataOutputStreamPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDataOutputStreamPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DataOutputStreamPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
