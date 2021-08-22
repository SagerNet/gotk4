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
		{T: externglib.Type(C.gtk_toggle_tool_button_get_type()), F: marshalToggleToolButtonner},
	})
}

// ToggleToolButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToggleToolButtonOverrider interface {
	Toggled()
}

// ToggleToolButton is a ToolItem that contains a toggle button.
//
// Use gtk_toggle_tool_button_new() to create a new GtkToggleToolButton.
//
//
// CSS nodes
//
// GtkToggleToolButton has a single CSS node with name togglebutton.
type ToggleToolButton struct {
	ToolButton
}

func WrapToggleToolButton(obj *externglib.Object) *ToggleToolButton {
	return &ToggleToolButton{
		ToolButton: ToolButton{
			ToolItem: ToolItem{
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
				Activatable: Activatable{
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
			Object: obj,
		},
	}
}

func marshalToggleToolButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleToolButton(obj), nil
}

// NewToggleToolButton returns a new ToggleToolButton
func NewToggleToolButton() *ToggleToolButton {
	var _cret *C.GtkToolItem // in

	_cret = C.gtk_toggle_tool_button_new()

	var _toggleToolButton *ToggleToolButton // out

	_toggleToolButton = WrapToggleToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleToolButton
}

// NewToggleToolButtonFromStock creates a new ToggleToolButton containing the
// image and text from a stock item. Some stock ids have preprocessor macros
// like K_STOCK_OK and K_STOCK_APPLY.
//
// It is an error if stock_id is not a name of a stock item.
//
// Deprecated: Use gtk_toggle_tool_button_new() instead.
func NewToggleToolButtonFromStock(stockId string) *ToggleToolButton {
	var _arg1 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_tool_button_new_from_stock(_arg1)
	runtime.KeepAlive(stockId)

	var _toggleToolButton *ToggleToolButton // out

	_toggleToolButton = WrapToggleToolButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleToolButton
}

// Active queries a ToggleToolButton and returns its current state. Returns TRUE
// if the toggle button is pressed in and FALSE if it is raised.
func (button *ToggleToolButton) Active() bool {
	var _arg0 *C.GtkToggleToolButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_toggle_tool_button_get_active(_arg0)
	runtime.KeepAlive(button)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle tool button. Set to TRUE if you want
// the GtkToggleButton to be “pressed in”, and FALSE to raise it. This action
// causes the toggled signal to be emitted.
func (button *ToggleToolButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleToolButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkToggleToolButton)(unsafe.Pointer(button.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_tool_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(isActive)
}
