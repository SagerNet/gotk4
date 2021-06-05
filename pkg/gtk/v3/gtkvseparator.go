// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vseparator_get_type()), F: marshalVSeparator},
	})
}

// VSeparator: the VSeparator widget is a vertical separator, used to group the
// widgets within a window. It displays a vertical line with a shadow to make it
// appear sunken into the interface.
//
// GtkVSeparator has been deprecated, use Separator instead.
type VSeparator interface {
	Separator
	Buildable
	Orientable
}

// vSeparator implements the VSeparator interface.
type vSeparator struct {
	Separator
	Buildable
	Orientable
}

var _ VSeparator = (*vSeparator)(nil)

// WrapVSeparator wraps a GObject to the right type. It is
// primarily used internally.
func WrapVSeparator(obj *externglib.Object) VSeparator {
	return VSeparator{
		Separator:  WrapSeparator(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalVSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVSeparator(obj), nil
}

// NewVSeparator constructs a class VSeparator.
func NewVSeparator() VSeparator {
	var cret C.GtkVSeparator
	var goret1 VSeparator

	cret = C.gtk_vseparator_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(VSeparator)

	return goret1
}
