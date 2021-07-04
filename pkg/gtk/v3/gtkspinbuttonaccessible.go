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
		{T: externglib.Type(C.gtk_spin_button_accessible_get_type()), F: marshalSpinButtonAccessible},
	})
}

type SpinButtonAccessible interface {
	EntryAccessible
}

// spinButtonAccessible implements the SpinButtonAccessible class.
type spinButtonAccessible struct {
	EntryAccessible
}

// WrapSpinButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapSpinButtonAccessible(obj *externglib.Object) SpinButtonAccessible {
	return spinButtonAccessible{
		EntryAccessible: WrapEntryAccessible(obj),
	}
}

func marshalSpinButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSpinButtonAccessible(obj), nil
}

func (a spinButtonAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a spinButtonAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a spinButtonAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a spinButtonAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a spinButtonAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a spinButtonAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a spinButtonAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}