// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_controls_get_type()), F: marshalWindowControlser},
	})
}

// WindowControls: GtkWindowControls shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// GtkWindowControls only displays start or end side of the controls (see
// gtk.WindowControls:side), so it's intended to be always used in pair with
// another GtkWindowControls for the opposite side, for example:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">start</property>
//        </object>
//      </child>
//
//      ...
//
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">end</property>
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    windowcontrols
//    ├── [image.icon]
//    ├── [button.minimize]
//    ├── [button.maximize]
//    ╰── [button.close]
//
//
// A GtkWindowControls' CSS node is called windowcontrols. It contains subnodes
// corresponding to each title button. Which of the title buttons exist and
// where they are placed exactly depends on the desktop environment and
// gtk.WindowControls:decoration-layout value.
//
// When gtk.WindowControls:empty is TRUE, it gets the .empty style class.
//
//
// Accessibility
//
// GtkWindowControls uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls struct {
	Widget
}

func WrapWindowControls(obj *externglib.Object) *WindowControls {
	return &WindowControls{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalWindowControlser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindowControls(obj), nil
}

// NewWindowControls creates a new GtkWindowControls.
func NewWindowControls(side PackType) *WindowControls {
	var _arg1 C.GtkPackType // out
	var _cret *C.GtkWidget  // in

	_arg1 = C.GtkPackType(side)

	_cret = C.gtk_window_controls_new(_arg1)
	runtime.KeepAlive(side)

	var _windowControls *WindowControls // out

	_windowControls = WrapWindowControls(externglib.Take(unsafe.Pointer(_cret)))

	return _windowControls
}

// DecorationLayout gets the decoration layout of this GtkWindowControls.
func (self *WindowControls) DecorationLayout() string {
	var _arg0 *C.GtkWindowControls // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_window_controls_get_decoration_layout(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Empty gets whether the widget has any window buttons.
func (self *WindowControls) Empty() bool {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_window_controls_get_empty(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Side gets the side to which this GtkWindowControls instance belongs.
func (self *WindowControls) Side() PackType {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.GtkPackType        // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_window_controls_get_side(_arg0)
	runtime.KeepAlive(self)

	var _packType PackType // out

	_packType = PackType(_cret)

	return _packType
}

// SetDecorationLayout sets the decoration layout for the title buttons.
//
// This overrides the gtk.Settings:gtk-decoration-layout setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close and icon (the window
// icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon on the left, and
// minimize, maximize and close buttons on the right.
//
// If gtk.WindowControls:side value is GTK_PACK_START, self will display the
// part before the colon, otherwise after that.
func (self *WindowControls) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(self.Native()))
	if layout != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_window_controls_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(layout)
}

// SetSide determines which part of decoration layout the GtkWindowControls
// uses.
//
// See gtk.WindowControls:decoration-layout.
func (self *WindowControls) SetSide(side PackType) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 C.GtkPackType        // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkPackType(side)

	C.gtk_window_controls_set_side(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(side)
}
