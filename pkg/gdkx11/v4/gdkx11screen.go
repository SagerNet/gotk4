// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screener},
	})
}

// X11Screener describes X11Screen's methods.
type X11Screener interface {
	gextras.Objector

	CurrentDesktop() uint32
	NumberOfDesktops() uint32
	ScreenNumber() int
	WindowManagerName() string
	SupportsNetWmHint(propertyName string) bool
}

type X11Screen struct {
	*externglib.Object
}

var _ X11Screener = (*X11Screen)(nil)

func wrapX11Screener(obj *externglib.Object) X11Screener {
	return &X11Screen{
		Object: obj,
	}
}

func marshalX11Screener(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Screener(obj), nil
}

// CurrentDesktop returns the current workspace for @screen when running under a
// window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
func (screen *X11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// NumberOfDesktops returns the number of workspaces for @screen when running
// under a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
func (screen *X11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.guint32       // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// ScreenNumber returns the index of a X11Screen.
func (screen *X11Screen) ScreenNumber() int {
	var _arg0 *C.GdkX11Screen // out
	var _cret C.int           // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WindowManagerName returns the name of the window manager for @screen.
func (screen *X11Screen) WindowManagerName() string {
	var _arg0 *C.GdkX11Screen // out
	var _cret *C.char         // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SupportsNetWmHint: this function is specific to the X11 backend of GDK, and
// indicates whether the window manager supports a certain hint from the
// Extended Window Manager Hints (http://www.freedesktop.org/Standards/wm-spec)
// specification.
//
// When using this function, keep in mind that the window manager can change
// over time; so you shouldn’t use this function in a way that impacts
// persistent application state. A common bug is that your application can start
// up before the window manager does when the user logs in, and before the
// window manager starts gdk_x11_screen_supports_net_wm_hint() will return false
// for every property. You can monitor the window_manager_changed signal on
// X11Screen to detect a window manager change.
func (screen *X11Screen) SupportsNetWmHint(propertyName string) bool {
	var _arg0 *C.GdkX11Screen // out
	var _arg1 *C.char         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(screen.Native()))
	_arg1 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
