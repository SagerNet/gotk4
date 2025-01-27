// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_action_bar_get_type()), F: marshalActionBarrer},
	})
}

// ActionBar is designed to present contextual actions. It is expected to be
// displayed below the content and expand horizontally to fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// GtkActionBar has a single CSS node with name actionbar.
type ActionBar struct {
	Bin
}

func WrapActionBar(obj *externglib.Object) *ActionBar {
	return &ActionBar{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
	}
}

func marshalActionBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActionBar(obj), nil
}

// NewActionBar creates a new ActionBar widget.
func NewActionBar() *ActionBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_action_bar_new()

	var _actionBar *ActionBar // out

	_actionBar = WrapActionBar(externglib.Take(unsafe.Pointer(_cret)))

	return _actionBar
}

// CenterWidget retrieves the center bar widget of the bar.
func (actionBar *ActionBar) CenterWidget() Widgetter {
	var _arg0 *C.GtkActionBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(actionBar.Native()))

	_cret = C.gtk_action_bar_get_center_widget(_arg0)
	runtime.KeepAlive(actionBar)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// PackEnd adds child to action_bar, packed with reference to the end of the
// action_bar.
func (actionBar *ActionBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(actionBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// PackStart adds child to action_bar, packed with reference to the start of the
// action_bar.
func (actionBar *ActionBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(actionBar.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(child)
}

// SetCenterWidget sets the center widget for the ActionBar.
func (actionBar *ActionBar) SetCenterWidget(centerWidget Widgetter) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(actionBar.Native()))
	if centerWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(centerWidget.Native()))
	}

	C.gtk_action_bar_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(actionBar)
	runtime.KeepAlive(centerWidget)
}
