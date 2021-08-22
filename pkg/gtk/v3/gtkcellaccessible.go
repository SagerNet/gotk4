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
		{T: externglib.Type(C.gtk_cell_accessible_get_type()), F: marshalCellAccessibler},
	})
}

// CellAccessibleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellAccessibleOverrider interface {
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	Accessible

	atk.Action
	atk.Component
	atk.TableCell
	*externglib.Object
}

func WrapCellAccessible(obj *externglib.Object) *CellAccessible {
	return &CellAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
		TableCell: atk.TableCell{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalCellAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAccessible(obj), nil
}

func (*CellAccessible) privateCellAccessible() {}
