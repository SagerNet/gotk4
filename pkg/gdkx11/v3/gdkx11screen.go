// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screen},
	})
}

// X11GetDefaultScreen gets the default GTK+ screen number.
func X11GetDefaultScreen() int {
	var _cret C.gint // in

	_cret = C.gdk_x11_get_default_screen()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// X11Screen:
type X11Screen interface {
	gdk.Screen

	// CurrentDesktop:
	CurrentDesktop() uint32
	// NumberOfDesktops:
	NumberOfDesktops() uint32
	// ScreenNumber:
	ScreenNumber() int
	// WindowManagerName:
	WindowManagerName() string
	// SupportsNetWmHintX11Screen:
	SupportsNetWmHintX11Screen(property *gdk.Atom) bool
}

// x11Screen implements the X11Screen class.
type x11Screen struct {
	gdk.Screen
}

// WrapX11Screen wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Screen(obj *externglib.Object) X11Screen {
	return x11Screen{
		Screen: gdk.WrapScreen(obj),
	}
}

func marshalX11Screen(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Screen(obj), nil
}

func (s x11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (s x11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkScreen // out
	var _cret C.guint32    // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

func (s x11Screen) ScreenNumber() int {
	var _arg0 *C.GdkScreen // out
	var _cret C.int        // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s x11Screen) WindowManagerName() string {
	var _arg0 *C.GdkScreen // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s x11Screen) SupportsNetWmHintX11Screen(property *gdk.Atom) bool {
	var _arg0 *C.GdkScreen // out
	var _arg1 C.GdkAtom    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkScreen)(unsafe.Pointer(s.Native()))
	{
		var refTmpIn *gdk.Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = property

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
