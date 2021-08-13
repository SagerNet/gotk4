// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_unix_output_stream_get_type()), F: marshalUnixOutputStreamer},
	})
}

// UnixOutputStream implements Stream for writing to a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that <gio/gunixoutputstream.h> belongs to the UNIX-specific GIO
// interfaces, thus you have to use the gio-unix-2.0.pc pkg-config file when
// using it.
type UnixOutputStream struct {
	OutputStream

	FileDescriptorBased
	PollableOutputStream
	*externglib.Object
}

func wrapUnixOutputStream(obj *externglib.Object) *UnixOutputStream {
	return &UnixOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
		FileDescriptorBased: FileDescriptorBased{
			Object: obj,
		},
		PollableOutputStream: PollableOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalUnixOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapUnixOutputStream(obj), nil
}

// NewUnixOutputStream creates a new OutputStream for the given fd.
//
// If close_fd, is TRUE, the file descriptor will be closed when the output
// stream is destroyed.
func NewUnixOutputStream(fd int, closeFd bool) *UnixOutputStream {
	var _arg1 C.gint           // out
	var _arg2 C.gboolean       // out
	var _cret *C.GOutputStream // in

	_arg1 = C.gint(fd)
	if closeFd {
		_arg2 = C.TRUE
	}

	_cret = C.g_unix_output_stream_new(_arg1, _arg2)
	runtime.KeepAlive(fd)
	runtime.KeepAlive(closeFd)

	var _unixOutputStream *UnixOutputStream // out

	_unixOutputStream = wrapUnixOutputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _unixOutputStream
}

// CloseFd returns whether the file descriptor of stream will be closed when the
// stream is closed.
func (stream *UnixOutputStream) CloseFd() bool {
	var _arg0 *C.GUnixOutputStream // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_unix_output_stream_get_close_fd(_arg0)
	runtime.KeepAlive(stream)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Fd: return the UNIX file descriptor that the stream writes to.
func (stream *UnixOutputStream) Fd() int {
	var _arg0 *C.GUnixOutputStream // out
	var _cret C.gint               // in

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_unix_output_stream_get_fd(_arg0)
	runtime.KeepAlive(stream)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetCloseFd sets whether the file descriptor of stream shall be closed when
// the stream is closed.
func (stream *UnixOutputStream) SetCloseFd(closeFd bool) {
	var _arg0 *C.GUnixOutputStream // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(stream.Native()))
	if closeFd {
		_arg1 = C.TRUE
	}

	C.g_unix_output_stream_set_close_fd(_arg0, _arg1)
	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeFd)
}
