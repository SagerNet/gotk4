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
		{T: externglib.Type(C.gtk_lock_button_accessible_get_type()), F: marshalLockButtonAccessible},
	})
}

type LockButtonAccessible interface {
	ButtonAccessible
}

// lockButtonAccessible implements the LockButtonAccessible class.
type lockButtonAccessible struct {
	ButtonAccessible
}

// WrapLockButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapLockButtonAccessible(obj *externglib.Object) LockButtonAccessible {
	return lockButtonAccessible{
		ButtonAccessible: WrapButtonAccessible(obj),
	}
}

func marshalLockButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLockButtonAccessible(obj), nil
}

func (a lockButtonAccessible) DoAction(i int) bool {
	return atk.WrapAction(gextras.InternObject(a)).DoAction(i)
}

func (a lockButtonAccessible) Description(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Description(i)
}

func (a lockButtonAccessible) Keybinding(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Keybinding(i)
}

func (a lockButtonAccessible) LocalizedName(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).LocalizedName(i)
}

func (a lockButtonAccessible) NActions() int {
	return atk.WrapAction(gextras.InternObject(a)).NActions()
}

func (a lockButtonAccessible) Name(i int) string {
	return atk.WrapAction(gextras.InternObject(a)).Name(i)
}

func (a lockButtonAccessible) SetDescription(i int, desc string) bool {
	return atk.WrapAction(gextras.InternObject(a)).SetDescription(i, desc)
}

func (i lockButtonAccessible) ImageDescription() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageDescription()
}

func (i lockButtonAccessible) ImageLocale() string {
	return atk.WrapImage(gextras.InternObject(i)).ImageLocale()
}

func (i lockButtonAccessible) ImagePosition(coordType atk.CoordType) (x int, y int) {
	return atk.WrapImage(gextras.InternObject(i)).ImagePosition(coordType)
}

func (i lockButtonAccessible) ImageSize() (width int, height int) {
	return atk.WrapImage(gextras.InternObject(i)).ImageSize()
}

func (i lockButtonAccessible) SetImageDescription(description string) bool {
	return atk.WrapImage(gextras.InternObject(i)).SetImageDescription(description)
}