// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_separator_menu_item_get_type()), F: marshalSeparatorMenuItemmer},
	})
}

// SeparatorMenuItem is a separator used to group items within a menu. It
// displays a horizontal line with a shadow to make it appear sunken into the
// interface.
//
//
// CSS nodes
//
// GtkSeparatorMenuItem has a single CSS node with name separator.
type SeparatorMenuItem struct {
	MenuItem
}

func WrapSeparatorMenuItem(obj *externglib.Object) *SeparatorMenuItem {
	return &SeparatorMenuItem{
		MenuItem: MenuItem{
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
		},
	}
}

func marshalSeparatorMenuItemmer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeparatorMenuItem(obj), nil
}

// NewSeparatorMenuItem creates a new SeparatorMenuItem.
func NewSeparatorMenuItem() *SeparatorMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_separator_menu_item_new()

	var _separatorMenuItem *SeparatorMenuItem // out

	_separatorMenuItem = WrapSeparatorMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _separatorMenuItem
}

func (*SeparatorMenuItem) privateSeparatorMenuItem() {}
