// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_unix_output_stream_get_type()), F: marshalUnixOutputStream},
	})
}

// UnixOutputStream implements Stream for writing to a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixoutputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixOutputStream interface {
	FileDescriptorBased
	PollableOutputStream

	// CloseFd:
	CloseFd() bool
	// GetFd:
	GetFd() int
	// SetCloseFdUnixOutputStream:
	SetCloseFdUnixOutputStream(closeFd bool)
}

// unixOutputStream implements the UnixOutputStream class.
type unixOutputStream struct {
	OutputStream
}

// WrapUnixOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixOutputStream(obj *externglib.Object) UnixOutputStream {
	return unixOutputStream{
		OutputStream: WrapOutputStream(obj),
	}
}

func marshalUnixOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixOutputStream(obj), nil
}

// NewUnixOutputStream:
func NewUnixOutputStream(fd int, closeFd bool) UnixOutputStream {
	var _arg1 C.gint           // out
	var _arg2 C.gboolean       // out
	var _cret *C.GOutputStream // in

	_arg1 = C.gint(fd)
	if closeFd {
		_arg2 = C.TRUE
	}

	_cret = C.g_unix_output_stream_new(_arg1, _arg2)

	var _unixOutputStream UnixOutputStream // out

	_unixOutputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(UnixOutputStream)

	return _unixOutputStream
}

func (s unixOutputStream) CloseFd() bool {
	var _arg0 *C.GUnixOutputStream // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_unix_output_stream_get_close_fd(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s unixOutputStream) GetFd() int {
	var _arg0 *C.GUnixOutputStream // out
	var _cret C.gint               // in

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_unix_output_stream_get_fd(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s unixOutputStream) SetCloseFdUnixOutputStream(closeFd bool) {
	var _arg0 *C.GUnixOutputStream // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GUnixOutputStream)(unsafe.Pointer(s.Native()))
	if closeFd {
		_arg1 = C.TRUE
	}

	C.g_unix_output_stream_set_close_fd(_arg0, _arg1)
}

func (f unixOutputStream) Fd() int {
	return WrapFileDescriptorBased(gextras.InternObject(f)).Fd()
}

func (s unixOutputStream) CanPoll() bool {
	return WrapPollableOutputStream(gextras.InternObject(s)).CanPoll()
}

func (s unixOutputStream) CreateSource(cancellable Cancellable) *glib.Source {
	return WrapPollableOutputStream(gextras.InternObject(s)).CreateSource(cancellable)
}

func (s unixOutputStream) IsWritable() bool {
	return WrapPollableOutputStream(gextras.InternObject(s)).IsWritable()
}

func (s unixOutputStream) WriteNonblocking(buffer []byte, cancellable Cancellable) (int, error) {
	return WrapPollableOutputStream(gextras.InternObject(s)).WriteNonblocking(buffer, cancellable)
}

func (s unixOutputStream) WritevNonblocking(vectors []OutputVector, cancellable Cancellable) (uint, PollableReturn, error) {
	return WrapPollableOutputStream(gextras.InternObject(s)).WritevNonblocking(vectors, cancellable)
}
