// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparator},
	})
}

// Separator: gtkSeparator is a horizontal or vertical separator widget,
// depending on the value of the Orientable:orientation property, used to group
// the widgets within a window. It displays a line with a shadow to make it
// appear sunken into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
type Separator interface {
	Widget
	Orientable
}

// separator implements the Separator class.
type separator struct {
	Widget
}

// WrapSeparator wraps a GObject to the right type. It is
// primarily used internally.
func WrapSeparator(obj *externglib.Object) Separator {
	return separator{
		Widget: WrapWidget(obj),
	}
}

func marshalSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeparator(obj), nil
}

// NewSeparator:
func NewSeparator(orientation Orientation) Separator {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_separator_new(_arg1)

	var _separator Separator // out

	_separator = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Separator)

	return _separator
}

func (b separator) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b separator) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b separator) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b separator) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b separator) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b separator) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b separator) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o separator) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o separator) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
