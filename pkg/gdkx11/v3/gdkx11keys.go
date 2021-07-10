// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gdk_x11_keymap_get_type()), F: marshalX11Keymapper},
	})
}

// X11Keymapper describes X11Keymap's methods.
type X11Keymapper interface {
	gextras.Objector

	GroupForState(state uint) int
	KeyIsModifier(keycode uint) bool
}

type X11Keymap struct {
	gdk.Keymap
}

var _ X11Keymapper = (*X11Keymap)(nil)

func wrapX11Keymapper(obj *externglib.Object) X11Keymapper {
	return &X11Keymap{
		Keymap: gdk.Keymap{
			Object: obj,
		},
	}
}

func marshalX11Keymapper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapX11Keymapper(obj), nil
}

// GroupForState extracts the group from the state field sent in an X Key event.
// This is only needed for code processing raw X events, since EventKey directly
// includes an is_modifier field.
func (keymap *X11Keymap) GroupForState(state uint) int {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gint       // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(state)

	_cret = C.gdk_x11_keymap_get_group_for_state(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// KeyIsModifier determines whether a particular key code represents a key that
// is a modifier. That is, it’s a key that normally just affects the keyboard
// state and the behavior of other keys rather than producing a direct effect
// itself. This is only needed for code processing raw X events, since EventKey
// directly includes an is_modifier field.
func (keymap *X11Keymap) KeyIsModifier(keycode uint) bool {
	var _arg0 *C.GdkKeymap // out
	var _arg1 C.guint      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GdkKeymap)(unsafe.Pointer(keymap.Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_x11_keymap_key_is_modifier(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
