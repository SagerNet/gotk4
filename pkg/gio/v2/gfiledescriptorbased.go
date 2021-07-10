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
		{T: externglib.Type(C.g_file_descriptor_based_get_type()), F: marshalFileDescriptorBasedder},
	})
}

// FileDescriptorBasedderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileDescriptorBasedderOverrider interface {
	// Fd gets the underlying file descriptor.
	Fd() int
}

// FileDescriptorBasedder describes FileDescriptorBased's methods.
type FileDescriptorBasedder interface {
	gextras.Objector

	Fd() int
}

// FileDescriptorBased is implemented by streams (implementations of Stream or
// Stream) that are based on file descriptors.
//
// Note that `<gio/gfiledescriptorbased.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type FileDescriptorBased struct {
	*externglib.Object
}

var _ FileDescriptorBasedder = (*FileDescriptorBased)(nil)

func wrapFileDescriptorBasedder(obj *externglib.Object) FileDescriptorBasedder {
	return &FileDescriptorBased{
		Object: obj,
	}
}

func marshalFileDescriptorBasedder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileDescriptorBasedder(obj), nil
}

// Fd gets the underlying file descriptor.
func (fdBased *FileDescriptorBased) Fd() int {
	var _arg0 *C.GFileDescriptorBased // out
	var _cret C.int                   // in

	_arg0 = (*C.GFileDescriptorBased)(unsafe.Pointer(fdBased.Native()))

	_cret = C.g_file_descriptor_based_get_fd(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
