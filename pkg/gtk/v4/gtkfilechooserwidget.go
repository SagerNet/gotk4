// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_widget_get_type()), F: marshalFileChooserWidgetter},
	})
}

// FileChooserWidgetter describes FileChooserWidget's methods.
type FileChooserWidgetter interface {
	gextras.Objector

	privateFileChooserWidget()
}

// FileChooserWidget: `GtkFileChooserWidget` is a widget for choosing files.
//
// It exposes the [iface@Gtk.FileChooser] interface, and you should use the
// methods of this interface to interact with the widget.
//
//
// CSS nodes
//
// `GtkFileChooserWidget` has a single CSS node with name filechooser.
type FileChooserWidget struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
	FileChooser
}

var _ FileChooserWidgetter = (*FileChooserWidget)(nil)

func wrapFileChooserWidgetter(obj *externglib.Object) FileChooserWidgetter {
	return &FileChooserWidget{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		FileChooser: FileChooser{
			Object: obj,
		},
	}
}

func marshalFileChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooserWidgetter(obj), nil
}

func (*FileChooserWidget) privateFileChooserWidget() {}
