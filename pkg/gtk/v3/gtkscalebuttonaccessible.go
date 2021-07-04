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
		{T: externglib.Type(C.gtk_scale_button_accessible_get_type()), F: marshalScaleButtonAccessible},
	})
}

type ScaleButtonAccessible interface {
	ButtonAccessible
}

// scaleButtonAccessible implements the ScaleButtonAccessible class.
type scaleButtonAccessible struct {
	ButtonAccessible
}

// WrapScaleButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapScaleButtonAccessible(obj *externglib.Object) ScaleButtonAccessible {
	return scaleButtonAccessible{
		ButtonAccessible: WrapButtonAccessible(obj),
	}
}

func marshalScaleButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScaleButtonAccessible(obj), nil
}

func (a scaleButtonAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a scaleButtonAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a scaleButtonAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a scaleButtonAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a scaleButtonAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a scaleButtonAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a scaleButtonAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (i scaleButtonAccessible) ImageDescription() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageDescription()
}

func (i scaleButtonAccessible) ImageLocale() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageLocale()
}

func (i scaleButtonAccessible) ImagePosition(coordType atk.CoordType) (x int, y int) {
	return atk.WrapImage(gextras.InternObject(i)).ImagePosition(coordType)
}

func (i scaleButtonAccessible) ImageSize() (width int, height int) {
	return atk.WrapImage(gextras.InternObject(i)).ImageSize()
}

func (i scaleButtonAccessible) SetImageDescription(description string) bool {
	return atk.WrapImage(gextras.InternObject(i)).SetImageDescription(description)
}