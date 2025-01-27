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
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparatorrer},
	})
}

// Separator is a horizontal or vertical separator widget, depending on the
// value of the Orientable:orientation property, used to group the widgets
// within a window. It displays a line with a shadow to make it appear sunken
// into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
type Separator struct {
	Widget

	Orientable
	*externglib.Object
}

func WrapSeparator(obj *externglib.Object) *Separator {
	return &Separator{
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
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSeparatorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeparator(obj), nil
}

// NewSeparator creates a new Separator with the given orientation.
func NewSeparator(orientation Orientation) *Separator {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_separator_new(_arg1)
	runtime.KeepAlive(orientation)

	var _separator *Separator // out

	_separator = WrapSeparator(externglib.Take(unsafe.Pointer(_cret)))

	return _separator
}

func (*Separator) privateSeparator() {}
