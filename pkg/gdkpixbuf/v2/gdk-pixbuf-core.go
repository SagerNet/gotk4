// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

// PixbufSaveFunc specifies the type of the function passed to
// gdk_pixbuf_save_to_callback(). It is called once for each block of bytes that
// is "written" by gdk_pixbuf_save_to_callback(). If successful it should return
// true. If an error occurs it should set @error and return false, in which case
// gdk_pixbuf_save_to_callback() will fail with the same error.
type PixbufSaveFunc func(buf []byte) (err *glib.Error, ok bool)

//export gotk4_PixbufSaveFunc
func gotk4_PixbufSaveFunc(arg0 *C.gchar, arg1 C.gsize, arg2 **C.GError, arg3 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}
	fn := v.(PixbufSaveFunc)

	error, ret := fn(buf, count, data)

	*arg2 = (**C.GError)(unsafe.Pointer(error.Native()))
	if ret {
		cret = C.gboolean(1)
	}

	return cret
}
