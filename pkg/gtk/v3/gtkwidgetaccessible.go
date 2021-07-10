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
		{T: externglib.Type(C.gtk_widget_accessible_get_type()), F: marshalWidgetAccessibler},
	})
}

// WidgetAccessibler describes WidgetAccessible's methods.
type WidgetAccessibler interface {
	gextras.Objector

	privateWidgetAccessible()
}

type WidgetAccessible struct {
	Accessible
}

var _ WidgetAccessibler = (*WidgetAccessible)(nil)

func wrapWidgetAccessibler(obj *externglib.Object) WidgetAccessibler {
	return &WidgetAccessible{
		Accessible: Accessible{
			Object: atk.Object{
				Object: obj,
			},
		},
	}
}

func marshalWidgetAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidgetAccessibler(obj), nil
}

func (*WidgetAccessible) privateWidgetAccessible() {}
