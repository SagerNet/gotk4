// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_pollfd_get_type()), F: marshalPollFD},
	})
}

// Poll polls @fds, as with the poll() system call, but portably. (On systems
// that don't have poll(), it is emulated using select().) This is used
// internally by Context, but it can be called directly if you need to block
// until a file descriptor is ready, but don't want to run the full main loop.
//
// Each element of @fds is a FD describing a single file descriptor to poll. The
// @fd field indicates the file descriptor, and the @events field indicates the
// events to poll for. On return, the @revents fields will be filled with the
// events that actually occurred.
//
// On POSIX systems, the file descriptors in @fds can be any sort of file
// descriptor, but the situation is much more complicated on Windows. If you
// need to use g_poll() in code that has to run on Windows, the easiest solution
// is to construct all of your FDs with g_io_channel_win32_make_pollfd().
func Poll(fds *PollFD, nfds uint, timeout int) int {
	var arg1 *C.GPollFD
	var arg2 C.guint
	var arg3 C.gint

	arg1 = (*C.GPollFD)(unsafe.Pointer(fds.Native()))
	arg2 = C.guint(nfds)
	arg3 = C.gint(timeout)

	var cret C.gint
	var goret1 int

	cret = C.g_poll(fds, nfds, timeout)

	goret1 = C.gint(cret)

	return goret1
}

// PollFD represents a file descriptor, which events to poll for, and which
// events occurred.
type PollFD struct {
	native C.GPollFD
}

// WrapPollFD wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPollFD(ptr unsafe.Pointer) *PollFD {
	if ptr == nil {
		return nil
	}

	return (*PollFD)(ptr)
}

func marshalPollFD(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPollFD(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PollFD) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Fd gets the field inside the struct.
func (p *PollFD) Fd() int {
	v = C.gint(p.native.fd)
}

// Events gets the field inside the struct.
func (p *PollFD) Events() uint16 {
	v = C.gushort(p.native.events)
}

// Revents gets the field inside the struct.
func (p *PollFD) Revents() uint16 {
	v = C.gushort(p.native.revents)
}
