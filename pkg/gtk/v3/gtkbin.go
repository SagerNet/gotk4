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
		{T: externglib.Type(C.gtk_bin_get_type()), F: marshalBinner},
	})
}

// Bin widget is a container with just one child. It is not very useful itself,
// but it is useful for deriving subclasses, since it provides common code
// needed for handling a single child widget.
//
// Many GTK+ widgets are subclasses of Bin, including Window, Button, Frame,
// HandleBox or ScrolledWindow.
type Bin struct {
	Container
}

// Binner describes Bin's abstract methods.
type Binner interface {
	externglib.Objector

	// Child gets the child of the Bin, or NULL if the bin contains no child
	// widget.
	Child() Widgetter
}

var _ Binner = (*Bin)(nil)

func WrapBin(obj *externglib.Object) *Bin {
	return &Bin{
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
	}
}

func marshalBinner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBin(obj), nil
}

// Child gets the child of the Bin, or NULL if the bin contains no child widget.
// The returned widget does not have a reference added, so you do not need to
// unref it.
func (bin *Bin) Child() Widgetter {
	var _arg0 *C.GtkBin    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkBin)(unsafe.Pointer(bin.Native()))

	_cret = C.gtk_bin_get_child(_arg0)
	runtime.KeepAlive(bin)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}
