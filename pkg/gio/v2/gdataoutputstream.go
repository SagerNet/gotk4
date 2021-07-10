// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_data_output_stream_get_type()), F: marshalDataOutputStreamer},
	})
}

// DataOutputStreamer describes DataOutputStream's methods.
type DataOutputStreamer interface {
	gextras.Objector

	ByteOrder() DataStreamByteOrder
	PutByte(data byte, cancellable Cancellabler) error
	PutInt16(data int16, cancellable Cancellabler) error
	PutInt32(data int32, cancellable Cancellabler) error
	PutInt64(data int64, cancellable Cancellabler) error
	PutString(str string, cancellable Cancellabler) error
	PutUint16(data uint16, cancellable Cancellabler) error
	PutUint32(data uint32, cancellable Cancellabler) error
	PutUint64(data uint64, cancellable Cancellabler) error
}

// DataOutputStream: data output stream implements Stream and includes functions
// for writing data directly to an output stream.
type DataOutputStream struct {
	FilterOutputStream
	Seekable
}

var _ DataOutputStreamer = (*DataOutputStream)(nil)

func wrapDataOutputStreamer(obj *externglib.Object) DataOutputStreamer {
	return &DataOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalDataOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDataOutputStreamer(obj), nil
}

// NewDataOutputStream creates a new data output stream for @base_stream.
func NewDataOutputStream(baseStream OutputStreamer) *DataOutputStream {
	var _arg1 *C.GOutputStream     // out
	var _cret *C.GDataOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))

	_cret = C.g_data_output_stream_new(_arg1)

	var _dataOutputStream *DataOutputStream // out

	_dataOutputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DataOutputStream)

	return _dataOutputStream
}

// ByteOrder gets the byte order for the stream.
func (stream *DataOutputStream) ByteOrder() DataStreamByteOrder {
	var _arg0 *C.GDataOutputStream   // out
	var _cret C.GDataStreamByteOrder // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_data_output_stream_get_byte_order(_arg0)

	var _dataStreamByteOrder DataStreamByteOrder // out

	_dataStreamByteOrder = (DataStreamByteOrder)(_cret)

	return _dataStreamByteOrder
}

// PutByte puts a byte into the output stream.
func (stream *DataOutputStream) PutByte(data byte, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.guchar             // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.guchar(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_byte(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt16 puts a signed 16-bit integer into the output stream.
func (stream *DataOutputStream) PutInt16(data int16, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.gint16             // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gint16(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_int16(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt32 puts a signed 32-bit integer into the output stream.
func (stream *DataOutputStream) PutInt32(data int32, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.gint32             // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gint32(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_int32(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutInt64 puts a signed 64-bit integer into the stream.
func (stream *DataOutputStream) PutInt64(data int64, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.gint64             // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.gint64(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_int64(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutString puts a string into the output stream.
func (stream *DataOutputStream) PutString(str string, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 *C.char              // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_string(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint16 puts an unsigned 16-bit integer into the output stream.
func (stream *DataOutputStream) PutUint16(data uint16, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.guint16            // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.guint16(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_uint16(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint32 puts an unsigned 32-bit integer into the stream.
func (stream *DataOutputStream) PutUint32(data uint32, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.guint32            // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.guint32(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_uint32(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PutUint64 puts an unsigned 64-bit integer into the stream.
func (stream *DataOutputStream) PutUint64(data uint64, cancellable Cancellabler) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg1 C.guint64            // out
	var _arg2 *C.GCancellable      // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.guint64(data)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_data_output_stream_put_uint64(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
