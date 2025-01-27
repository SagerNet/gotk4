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
		{T: externglib.Type(C.gtk_hbutton_box_get_type()), F: marshalHButtonBoxer},
	})
}

type HButtonBox struct {
	ButtonBox
}

func WrapHButtonBox(obj *externglib.Object) *HButtonBox {
	return &HButtonBox{
		ButtonBox: ButtonBox{
			Box: Box{
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
				Orientable: Orientable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalHButtonBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHButtonBox(obj), nil
}

// NewHButtonBox creates a new horizontal button box.
//
// Deprecated: Use gtk_button_box_new() with GTK_ORIENTATION_HORIZONTAL instead.
func NewHButtonBox() *HButtonBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hbutton_box_new()

	var _hButtonBox *HButtonBox // out

	_hButtonBox = WrapHButtonBox(externglib.Take(unsafe.Pointer(_cret)))

	return _hButtonBox
}

func (*HButtonBox) privateHButtonBox() {}
