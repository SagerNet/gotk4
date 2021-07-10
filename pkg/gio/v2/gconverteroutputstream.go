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
		{T: externglib.Type(C.g_converter_output_stream_get_type()), F: marshalConverterOutputStreamer},
	})
}

// ConverterOutputStreamer describes ConverterOutputStream's methods.
type ConverterOutputStreamer interface {
	gextras.Objector

	Converter() *Converter
}

// ConverterOutputStream: converter output stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, OutputStream implements OutputStream.
type ConverterOutputStream struct {
	*externglib.Object
	FilterOutputStream
	PollableOutputStream
}

var _ ConverterOutputStreamer = (*ConverterOutputStream)(nil)

func wrapConverterOutputStreamer(obj *externglib.Object) ConverterOutputStreamer {
	return &ConverterOutputStream{
		Object: obj,
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
	}
}

func marshalConverterOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapConverterOutputStreamer(obj), nil
}

// NewConverterOutputStream creates a new converter output stream for the
// @base_stream.
func NewConverterOutputStream(baseStream OutputStreamer, converter Converterrer) *ConverterOutputStream {
	var _arg1 *C.GOutputStream // out
	var _arg2 *C.GConverter    // out
	var _cret *C.GOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))
	_arg2 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	_cret = C.g_converter_output_stream_new(_arg1, _arg2)

	var _converterOutputStream *ConverterOutputStream // out

	_converterOutputStream = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*ConverterOutputStream)

	return _converterOutputStream
}

// Converter gets the #GConverter that is used by @converter_stream.
func (converterStream *ConverterOutputStream) Converter() *Converter {
	var _arg0 *C.GConverterOutputStream // out
	var _cret *C.GConverter             // in

	_arg0 = (*C.GConverterOutputStream)(unsafe.Pointer(converterStream.Native()))

	_cret = C.g_converter_output_stream_get_converter(_arg0)

	var _converter *Converter // out

	_converter = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Converter)

	return _converter
}
