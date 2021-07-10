// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_menu_item_accessible_get_type()), F: marshalRadioMenuItemAccessibler},
	})
}

// RadioMenuItemAccessibler describes RadioMenuItemAccessible's methods.
type RadioMenuItemAccessibler interface {
	gextras.Objector

	privateRadioMenuItemAccessible()
}

type RadioMenuItemAccessible struct {
	CheckMenuItemAccessible
	atk.Action
}

var _ RadioMenuItemAccessibler = (*RadioMenuItemAccessible)(nil)

func wrapRadioMenuItemAccessibler(obj *externglib.Object) RadioMenuItemAccessibler {
	return &RadioMenuItemAccessible{
		CheckMenuItemAccessible: CheckMenuItemAccessible{
			MenuItemAccessible: MenuItemAccessible{
				ContainerAccessible: ContainerAccessible{
					WidgetAccessible: WidgetAccessible{
						Accessible: Accessible{
							Object: atk.Object{
								Object: obj,
							},
						},
					},
				},
				Action: atk.Action{
					Object: obj,
				},
			},
			Action: atk.Action{
				Object: obj,
			},
		},
		Action: atk.Action{
			Object: obj,
		},
	}
}

func marshalRadioMenuItemAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioMenuItemAccessibler(obj), nil
}

func (*RadioMenuItemAccessible) privateRadioMenuItemAccessible() {}
