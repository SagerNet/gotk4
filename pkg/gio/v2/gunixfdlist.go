// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.g_unix_fd_list_get_type()), F: marshalUnixFDList},
	})
}

// UnixFDList: a FDList contains a list of file descriptors. It owns the file
// descriptors that it contains, closing them when finalized.
//
// It may be wrapped in a FDMessage and sent over a #GSocket in the
// G_SOCKET_FAMILY_UNIX family by using g_socket_send_message() and received
// using g_socket_receive_message().
//
// Note that `<gio/gunixfdlist.h>` belongs to the UNIX-specific GIO interfaces,
// thus you have to use the `gio-unix-2.0.pc` pkg-config file when using it.
type UnixFDList interface {
	gextras.Objector

	// AppendUnixFDList:
	AppendUnixFDList(fd int) (int, error)
	// GetUnixFDList:
	GetUnixFDList(index_ int) (int, error)
	// Length:
	Length() int
}

// unixFDList implements the UnixFDList class.
type unixFDList struct {
	gextras.Objector
}

// WrapUnixFDList wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixFDList(obj *externglib.Object) UnixFDList {
	return unixFDList{
		Objector: obj,
	}
}

func marshalUnixFDList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixFDList(obj), nil
}

// NewUnixFDList:
func NewUnixFDList() UnixFDList {
	var _cret *C.GUnixFDList // in

	_cret = C.g_unix_fd_list_new()

	var _unixFDList UnixFDList // out

	_unixFDList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(UnixFDList)

	return _unixFDList
}

// NewUnixFDListFromArray:
func NewUnixFDListFromArray(fds []int) UnixFDList {
	var _arg1 *C.gint
	var _arg2 C.gint
	var _cret *C.GUnixFDList // in

	_arg2 = C.gint(len(fds))
	_arg1 = (*C.gint)(unsafe.Pointer(&fds[0]))

	_cret = C.g_unix_fd_list_new_from_array(_arg1, _arg2)

	var _unixFDList UnixFDList // out

	_unixFDList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(UnixFDList)

	return _unixFDList
}

func (l unixFDList) AppendUnixFDList(fd int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(fd)

	_cret = C.g_unix_fd_list_append(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

func (l unixFDList) GetUnixFDList(index_ int) (int, error) {
	var _arg0 *C.GUnixFDList // out
	var _arg1 C.gint         // out
	var _cret C.gint         // in
	var _cerr *C.GError      // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(index_)

	_cret = C.g_unix_fd_list_get(_arg0, _arg1, &_cerr)

	var _gint int    // out
	var _goerr error // out

	_gint = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gint, _goerr
}

func (l unixFDList) Length() int {
	var _arg0 *C.GUnixFDList // out
	var _cret C.gint         // in

	_arg0 = (*C.GUnixFDList)(unsafe.Pointer(l.Native()))

	_cret = C.g_unix_fd_list_get_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
