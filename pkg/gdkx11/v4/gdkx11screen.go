// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screen},
	})
}

// X11Screen:
type X11Screen interface {
	gextras.Objector

	// CurrentDesktop:
	CurrentDesktop() uint32
	// NumberOfDesktops:
	NumberOfDesktops() uint32
	// ScreenNumber:
	ScreenNumber() int
	// WindowManagerName:
	WindowManagerName() string
	// SupportsNetWmHintX11Screen:
	SupportsNetWmHintX11Screen(propertyName string) bool
}

// x11Screen implements the X11Screen class.
type x11Screen struct {
	gextras.Objector
}

// WrapX11Screen wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Screen(obj *externglib.Object) X11Screen {
	return x11Screen{
		Objector: obj,
	}
}

func marshalX11Screen(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Screen(obj), nil
}

func (s x11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (s x11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (s x11Screen) ScreenNumber() int {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.int           // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s x11Screen) WindowManagerName() string {
	var _arg0 *C.GdkX11Screen // out
	var _cret *C.char         // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s x11Screen) SupportsNetWmHintX11Screen(propertyName string) bool {
	var _arg0 *C.GdkX11Screen // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
