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
		{T: externglib.Type(C.gtk_switch_get_type()), F: marshalSwitcher},
	})
}

// SwitchOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SwitchOverrider interface {
	Activate()
	StateSet(state bool) bool
}

// Switch is a widget that has two states: on or off. The user can control which
// state should be active by clicking the empty area, or by dragging the handle.
//
// GtkSwitch can also handle situations where the underlying state changes with
// a delay. See Switch::state-set for details.
//
// CSS nodes
//
//    switch
//    ╰── slider
//
// GtkSwitch has two css nodes, the main node with the name switch and a subnode
// named slider. Neither of them is using any style classes.
type Switch struct {
	Widget

	Actionable
	Activatable
	*externglib.Object
}

func WrapSwitch(obj *externglib.Object) *Switch {
	return &Switch{
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
		Actionable: Actionable{
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
		Activatable: Activatable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSwitcher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSwitch(obj), nil
}

// NewSwitch creates a new Switch widget.
func NewSwitch() *Switch {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_switch_new()

	var __switch *Switch // out

	__switch = WrapSwitch(externglib.Take(unsafe.Pointer(_cret)))

	return __switch
}

// Active gets whether the Switch is in its “on” or “off” state.
func (sw *Switch) Active() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(sw.Native()))

	_cret = C.gtk_switch_get_active(_arg0)
	runtime.KeepAlive(sw)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// State gets the underlying state of the Switch.
func (sw *Switch) State() bool {
	var _arg0 *C.GtkSwitch // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(sw.Native()))

	_cret = C.gtk_switch_get_state(_arg0)
	runtime.KeepAlive(sw)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the state of sw to the desired one.
func (sw *Switch) SetActive(isActive bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(sw.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_active(_arg0, _arg1)
	runtime.KeepAlive(sw)
	runtime.KeepAlive(isActive)
}

// SetState sets the underlying state of the Switch.
//
// Normally, this is the same as Switch:active, unless the switch is set up for
// delayed state changes. This function is typically called from a
// Switch::state-set signal handler.
//
// See Switch::state-set for details.
func (sw *Switch) SetState(state bool) {
	var _arg0 *C.GtkSwitch // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkSwitch)(unsafe.Pointer(sw.Native()))
	if state {
		_arg1 = C.TRUE
	}

	C.gtk_switch_set_state(_arg0, _arg1)
	runtime.KeepAlive(sw)
	runtime.KeepAlive(state)
}
