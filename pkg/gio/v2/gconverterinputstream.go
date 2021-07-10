// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_converter_input_stream_get_type()), F: marshalConverterInputStreamer},
	})
}

// ConverterInputStreamer describes ConverterInputStream's methods.
type ConverterInputStreamer interface {
	gextras.Objector

	Converter() *Converter
}

// ConverterInputStream: converter input stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, InputStream implements InputStream.
type ConverterInputStream struct {
	*externglib.Object
	FilterInputStream
	PollableInputStream
}

var _ ConverterInputStreamer = (*ConverterInputStream)(nil)

func wrapConverterInputStreamer(obj *externglib.Object) ConverterInputStreamer {
	return &ConverterInputStream{
		Object: obj,
		FilterInputStream: FilterInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
		PollableInputStream: PollableInputStream{
			InputStream: InputStream{
				Object: obj,
			},
		},
	}
}

func marshalConverterInputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConverterInputStreamer(obj), nil
}

// NewConverterInputStream creates a new converter input stream for the
// @base_stream.
func NewConverterInputStream(baseStream InputStreamer, converter Converterrer) *ConverterInputStream {
	var _arg1 *C.GInputStream // out
	var _arg2 *C.GConverter   // out
	var _cret *C.GInputStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_converter_input_stream_new(_arg1, _arg2)

	var _converterInputStream *ConverterInputStream // out

	_converterInputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ConverterInputStream)

	return _converterInputStream
}

// Converter gets the #GConverter that is used by @converter_stream.
func (converterStream *ConverterInputStream) Converter() *Converter {
	var _arg0 *C.GConverterInputStream // out
	var _cret *C.GConverter            // in

	_arg0 = (*C.GConverterInputStream)(unsafe.Pointer(converterStream.Native()))

	_cret = C.g_converter_input_stream_get_converter(_arg0)

	var _converter *Converter // out

	_converter = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Converter)

	return _converter
}
