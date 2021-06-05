// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_popup_get_type()), F: marshalPopup},
	})
}

// Popup: a Popup is a surface that is attached to another surface, called its
// Popup:parent, and is positioned relative to it.
//
// Popups are typically used to implement menus and similar popups. They can be
// modal, which is indicated by the Popup:autohide property.
type Popup interface {
	Surface

	// Autohide returns whether this popup is set to hide on outside clicks.
	Autohide() bool
	// Parent returns the parent surface of a popup.
	Parent() Surface
	// PositionX obtains the position of the popup relative to its parent.
	PositionX() int
	// PositionY obtains the position of the popup relative to its parent.
	PositionY() int
	// RectAnchor gets the current popup rectangle anchor.
	//
	// The value returned may change after calling gdk_popup_present(), or after
	// the Surface::layout signal is emitted.
	RectAnchor() Gravity
	// SurfaceAnchor gets the current popup surface anchor.
	//
	// The value returned may change after calling gdk_popup_present(), or after
	// the Surface::layout signal is emitted.
	SurfaceAnchor() Gravity
	// Present: present @popup after having processed the PopupLayout rules. If
	// the popup was previously now showing, it will be showed, otherwise it
	// will change position according to @layout.
	//
	// After calling this function, the result should be handled in response to
	// the Surface::layout signal being emitted. The resulting popup position
	// can be queried using gdk_popup_get_position_x(),
	// gdk_popup_get_position_y(), and the resulting size will be sent as
	// parameters in the layout signal. Use gdk_popup_get_rect_anchor() and
	// gdk_popup_get_surface_anchor() to get the resulting anchors.
	//
	// Presenting may fail, for example if the @popup is set to autohide and is
	// immediately hidden upon being presented. If presenting failed, the
	// Surface::layout signal will not me emitted.
	Present(width int, height int, layout *PopupLayout) bool
}

// popup implements the Popup interface.
type popup struct {
	Surface
}

var _ Popup = (*popup)(nil)

// WrapPopup wraps a GObject to a type that implements interface
// Popup. It is primarily used internally.
func WrapPopup(obj *externglib.Object) Popup {
	return Popup{
		Surface: WrapSurface(obj),
	}
}

func marshalPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopup(obj), nil
}

// Autohide returns whether this popup is set to hide on outside clicks.
func (p popup) Autohide() bool {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_popup_get_autohide(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Parent returns the parent surface of a popup.
func (p popup) Parent() Surface {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret *C.GdkSurface
	var ret1 Surface

	cret = C.gdk_popup_get_parent(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Surface)

	return ret1
}

// PositionX obtains the position of the popup relative to its parent.
func (p popup) PositionX() int {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_popup_get_position_x(arg0)

	ret1 = C.int(cret)

	return ret1
}

// PositionY obtains the position of the popup relative to its parent.
func (p popup) PositionY() int {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret C.int
	var ret1 int

	cret = C.gdk_popup_get_position_y(arg0)

	ret1 = C.int(cret)

	return ret1
}

// RectAnchor gets the current popup rectangle anchor.
//
// The value returned may change after calling gdk_popup_present(), or after
// the Surface::layout signal is emitted.
func (p popup) RectAnchor() Gravity {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret C.GdkGravity
	var ret1 Gravity

	cret = C.gdk_popup_get_rect_anchor(arg0)

	ret1 = Gravity(cret)

	return ret1
}

// SurfaceAnchor gets the current popup surface anchor.
//
// The value returned may change after calling gdk_popup_present(), or after
// the Surface::layout signal is emitted.
func (p popup) SurfaceAnchor() Gravity {
	var arg0 *C.GdkPopup

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))

	var cret C.GdkGravity
	var ret1 Gravity

	cret = C.gdk_popup_get_surface_anchor(arg0)

	ret1 = Gravity(cret)

	return ret1
}

// Present: present @popup after having processed the PopupLayout rules. If
// the popup was previously now showing, it will be showed, otherwise it
// will change position according to @layout.
//
// After calling this function, the result should be handled in response to
// the Surface::layout signal being emitted. The resulting popup position
// can be queried using gdk_popup_get_position_x(),
// gdk_popup_get_position_y(), and the resulting size will be sent as
// parameters in the layout signal. Use gdk_popup_get_rect_anchor() and
// gdk_popup_get_surface_anchor() to get the resulting anchors.
//
// Presenting may fail, for example if the @popup is set to autohide and is
// immediately hidden upon being presented. If presenting failed, the
// Surface::layout signal will not me emitted.
func (p popup) Present(width int, height int, layout *PopupLayout) bool {
	var arg0 *C.GdkPopup
	var arg1 C.int
	var arg2 C.int
	var arg3 *C.GdkPopupLayout

	arg0 = (*C.GdkPopup)(unsafe.Pointer(p.Native()))
	arg1 = C.int(width)
	arg2 = C.int(height)
	arg3 = (*C.GdkPopupLayout)(unsafe.Pointer(layout.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gdk_popup_present(arg0, width, height, layout)

	ret1 = C.bool(cret) != C.false

	return ret1
}