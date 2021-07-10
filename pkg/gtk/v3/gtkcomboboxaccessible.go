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
		{T: externglib.Type(C.gtk_combo_box_accessible_get_type()), F: marshalComboBoxAccessibler},
	})
}

// ComboBoxAccessibler describes ComboBoxAccessible's methods.
type ComboBoxAccessibler interface {
	gextras.Objector

	privateComboBoxAccessible()
}

type ComboBoxAccessible struct {
	ContainerAccessible
	atk.Action
}

var _ ComboBoxAccessibler = (*ComboBoxAccessible)(nil)

func wrapComboBoxAccessibler(obj *externglib.Object) ComboBoxAccessibler {
	return &ComboBoxAccessible{
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
	}
}

func marshalComboBoxAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBoxAccessibler(obj), nil
}

func (*ComboBoxAccessible) privateComboBoxAccessible() {}
