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
		{T: externglib.Type(C.gtk_widget_accessible_get_type()), F: marshalWidgetAccessibler},
	})
}

type WidgetAccessible struct {
	Accessible

	atk.Component
}

func WrapWidgetAccessible(obj *externglib.Object) *WidgetAccessible {
	return &WidgetAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Component: atk.Component{
			Object: obj,
		},
	}
}

func marshalWidgetAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWidgetAccessible(obj), nil
}

func (*WidgetAccessible) privateWidgetAccessible() {}
