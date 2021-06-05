// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_pollable_output_stream_get_type()), F: marshalPollableOutputStream},
	})
}

// PollableOutputStreamOverrider contains methods that are overridable. This
// interface is a subset of the interface PollableOutputStream.
type PollableOutputStreamOverrider interface {
	// CanPoll checks if @stream is actually pollable. Some classes may
	// implement OutputStream but have only certain instances of that class be
	// pollable. If this method returns false, then the behavior of other
	// OutputStream methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when @stream can be
	// written, or @cancellable is triggered or an error occurs. The callback on
	// the source is of the SourceFunc type.
	//
	// As with g_pollable_output_stream_is_writable(), it is possible that the
	// stream may not actually be writable even after the source triggers, so
	// you should use g_pollable_output_stream_write_nonblocking() rather than
	// g_output_stream_write() from the callback.
	CreateSource(cancellable Cancellable) *glib.Source
	// IsWritable checks if @stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns true would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	IsWritable() bool
	// WriteNonblocking attempts to write up to @count bytes from @buffer to
	// @stream, as with g_output_stream_write(). If @stream is not currently
	// writable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_output_stream_create_source() to create a #GSource
	// that will be triggered when @stream is writable.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	//
	// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
	// transports like D/TLS require that you re-send the same @buffer and
	// @count in the next write call.
	WriteNonblocking(buffer []byte) (gssize int, err error)
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
type PollableOutputStream interface {
	OutputStream
	PollableOutputStreamOverrider
}

// pollableOutputStream implements the PollableOutputStream interface.
type pollableOutputStream struct {
	OutputStream
}

var _ PollableOutputStream = (*pollableOutputStream)(nil)

// WrapPollableOutputStream wraps a GObject to a type that implements interface
// PollableOutputStream. It is primarily used internally.
func WrapPollableOutputStream(obj *externglib.Object) PollableOutputStream {
	return PollableOutputStream{
		OutputStream: WrapOutputStream(obj),
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPollableOutputStream(obj), nil
}

// CanPoll checks if @stream is actually pollable. Some classes may
// implement OutputStream but have only certain instances of that class be
// pollable. If this method returns false, then the behavior of other
// OutputStream methods is undefined.
//
// For any given stream, the value returned by this method is constant; a
// stream cannot switch from pollable to non-pollable or vice versa.
func (s pollableOutputStream) CanPoll() bool {
	var arg0 *C.GPollableOutputStream

	arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_pollable_output_stream_can_poll(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// CreateSource creates a #GSource that triggers when @stream can be
// written, or @cancellable is triggered or an error occurs. The callback on
// the source is of the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so
// you should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
func (s pollableOutputStream) CreateSource(cancellable Cancellable) *glib.Source {
	var arg0 *C.GPollableOutputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GSource
	var goret1 *glib.Source

	cret = C.g_pollable_output_stream_create_source(arg0, cancellable)

	goret1 = glib.WrapSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.Source) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// IsWritable checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100%
// reliably, and it is possible that a call to g_output_stream_write() after
// this returns true would still block. To guarantee non-blocking behavior,
// you should always use g_pollable_output_stream_write_nonblocking(), which
// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (s pollableOutputStream) IsWritable() bool {
	var arg0 *C.GPollableOutputStream

	arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_pollable_output_stream_is_writable(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}
