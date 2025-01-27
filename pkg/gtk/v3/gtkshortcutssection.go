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
		{T: externglib.Type(C.gtk_shortcuts_section_get_type()), F: marshalShortcutsSectioner},
	})
}

// ShortcutsSection collects all the keyboard shortcuts and gestures for a major
// application mode. If your application needs multiple sections, you should
// give each section a unique ShortcutsSection:section-name and a
// ShortcutsSection:title that can be shown in the section selector of the
// GtkShortcutsWindow.
//
// The ShortcutsSection:max-height property can be used to influence how the
// groups in the section are distributed over pages and columns.
//
// This widget is only meant to be used with ShortcutsWindow.
type ShortcutsSection struct {
	Box
}

func WrapShortcutsSection(obj *externglib.Object) *ShortcutsSection {
	return &ShortcutsSection{
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
	}
}

func marshalShortcutsSectioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsSection(obj), nil
}

func (*ShortcutsSection) privateShortcutsSection() {}
