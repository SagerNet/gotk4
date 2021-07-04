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
		{T: externglib.Type(C.gtk_hpaned_get_type()), F: marshalHPaned},
	})
}

// HPaned: the HPaned widget is a container widget with two children arranged
// horizontally. The division between the two panes is adjustable by the user by
// dragging a handle. See Paned for details.
//
// GtkHPaned has been deprecated, use Paned instead.
type HPaned interface {
	Paned
}

// hPaned implements the HPaned class.
type hPaned struct {
	Paned
}

// WrapHPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapHPaned(obj *externglib.Object) HPaned {
	return hPaned{
		Paned: WrapPaned(obj),
	}
}

func marshalHPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHPaned(obj), nil
}

// NewHPaned:
func NewHPaned() HPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hpaned_new()

	var _hPaned HPaned // out

	_hPaned = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(HPaned)

	return _hPaned
}

func (b hPaned) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b hPaned) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b hPaned) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b hPaned) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b hPaned) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b hPaned) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b hPaned) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o hPaned) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o hPaned) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
