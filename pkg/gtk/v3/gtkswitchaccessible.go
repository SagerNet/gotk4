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
		{T: externglib.Type(C.gtk_switch_accessible_get_type()), F: marshalSwitchAccessibler},
	})
}

// SwitchAccessibler describes SwitchAccessible's methods.
type SwitchAccessibler interface {
	gextras.Objector

	privateSwitchAccessible()
}

type SwitchAccessible struct {
	WidgetAccessible
	atk.Action
}

var _ SwitchAccessibler = (*SwitchAccessible)(nil)

func wrapSwitchAccessibler(obj *externglib.Object) SwitchAccessibler {
	return &SwitchAccessible{
		WidgetAccessible: WidgetAccessible{
			Accessible: Accessible{
				Object: atk.Object{
					Object: obj,
				},
			},
		},
		Action: atk.Action{
			Object: obj,
		},
	}
}

func marshalSwitchAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSwitchAccessibler(obj), nil
}

func (*SwitchAccessible) privateSwitchAccessible() {}
