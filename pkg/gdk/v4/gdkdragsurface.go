// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drag_surface_get_type()), F: marshalDragSurfacer},
	})
}

// DragSurface is an interface for surfaces used during DND.
type DragSurface struct {
	Surface
}

// DragSurfacer describes DragSurface's abstract methods.
type DragSurfacer interface {
	externglib.Objector

	// Present drag_surface.
	Present(width int, height int) bool
}

var _ DragSurfacer = (*DragSurface)(nil)

func WrapDragSurface(obj *externglib.Object) *DragSurface {
	return &DragSurface{
		Surface: Surface{
			Object: obj,
		},
	}
}

func marshalDragSurfacer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDragSurface(obj), nil
}

// Present drag_surface.
func (dragSurface *DragSurface) Present(width int, height int) bool {
	var _arg0 *C.GdkDragSurface // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDragSurface)(unsafe.Pointer(dragSurface.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	_cret = C.gdk_drag_surface_present(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragSurface)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
