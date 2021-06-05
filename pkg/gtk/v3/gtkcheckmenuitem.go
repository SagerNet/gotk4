// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_menu_item_get_type()), F: marshalCheckMenuItem},
	})
}

// CheckMenuItem: a CheckMenuItem is a menu item that maintains the state of a
// boolean value in addition to a MenuItem usual role in activating application
// code.
//
// A check box indicating the state of the boolean value is displayed at the
// left side of the MenuItem. Activating the MenuItem toggles the value.
//
// CSS nodes
//
//    menuitem
//    ├── check.left
//    ╰── <child>
//
// GtkCheckMenuItem has a main CSS node with name menuitem, and a subnode with
// name check, which gets the .left or .right style class.
type CheckMenuItem interface {
	MenuItem
	Actionable
	Activatable
	Buildable

	// Active returns whether the check menu item is active. See
	// gtk_check_menu_item_set_active ().
	Active() bool
	// DrawAsRadio returns whether @check_menu_item looks like a RadioMenuItem
	DrawAsRadio() bool
	// Inconsistent retrieves the value set by
	// gtk_check_menu_item_set_inconsistent().
	Inconsistent() bool
	// SetActive sets the active state of the menu item’s check box.
	SetActive(isActive bool)
	// SetDrawAsRadio sets whether @check_menu_item is drawn like a
	// RadioMenuItem
	SetDrawAsRadio(drawAsRadio bool)
	// SetInconsistent: if the user has selected a range of elements (such as
	// some text or spreadsheet cells) that are affected by a boolean setting,
	// and the current values in that range are inconsistent, you may want to
	// display the check in an “in between” state. This function turns on “in
	// between” display. Normally you would turn off the inconsistent state
	// again if the user explicitly selects a setting. This has to be done
	// manually, gtk_check_menu_item_set_inconsistent() only affects visual
	// appearance, it doesn’t affect the semantics of the widget.
	SetInconsistent(setting bool)
	// Toggled emits the CheckMenuItem::toggled signal.
	Toggled()
}

// checkMenuItem implements the CheckMenuItem interface.
type checkMenuItem struct {
	MenuItem
	Actionable
	Activatable
	Buildable
}

var _ CheckMenuItem = (*checkMenuItem)(nil)

// WrapCheckMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapCheckMenuItem(obj *externglib.Object) CheckMenuItem {
	return CheckMenuItem{
		MenuItem:    WrapMenuItem(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalCheckMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCheckMenuItem(obj), nil
}

// NewCheckMenuItem constructs a class CheckMenuItem.
func NewCheckMenuItem() CheckMenuItem {
	var cret C.GtkCheckMenuItem
	var ret1 CheckMenuItem

	cret = C.gtk_check_menu_item_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CheckMenuItem)

	return ret1
}

// NewCheckMenuItemWithLabel constructs a class CheckMenuItem.
func NewCheckMenuItemWithLabel(label string) CheckMenuItem {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkCheckMenuItem
	var ret1 CheckMenuItem

	cret = C.gtk_check_menu_item_new_with_label(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CheckMenuItem)

	return ret1
}

// NewCheckMenuItemWithMnemonic constructs a class CheckMenuItem.
func NewCheckMenuItemWithMnemonic(label string) CheckMenuItem {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkCheckMenuItem
	var ret1 CheckMenuItem

	cret = C.gtk_check_menu_item_new_with_mnemonic(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CheckMenuItem)

	return ret1
}

// Active returns whether the check menu item is active. See
// gtk_check_menu_item_set_active ().
func (c checkMenuItem) Active() bool {
	var arg0 *C.GtkCheckMenuItem

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_check_menu_item_get_active(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// DrawAsRadio returns whether @check_menu_item looks like a RadioMenuItem
func (c checkMenuItem) DrawAsRadio() bool {
	var arg0 *C.GtkCheckMenuItem

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_check_menu_item_get_draw_as_radio(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Inconsistent retrieves the value set by
// gtk_check_menu_item_set_inconsistent().
func (c checkMenuItem) Inconsistent() bool {
	var arg0 *C.GtkCheckMenuItem

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_check_menu_item_get_inconsistent(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetActive sets the active state of the menu item’s check box.
func (c checkMenuItem) SetActive(isActive bool) {
	var arg0 *C.GtkCheckMenuItem
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))
	if isActive {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_menu_item_set_active(arg0, isActive)
}

// SetDrawAsRadio sets whether @check_menu_item is drawn like a
// RadioMenuItem
func (c checkMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	var arg0 *C.GtkCheckMenuItem
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))
	if drawAsRadio {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_menu_item_set_draw_as_radio(arg0, drawAsRadio)
}

// SetInconsistent: if the user has selected a range of elements (such as
// some text or spreadsheet cells) that are affected by a boolean setting,
// and the current values in that range are inconsistent, you may want to
// display the check in an “in between” state. This function turns on “in
// between” display. Normally you would turn off the inconsistent state
// again if the user explicitly selects a setting. This has to be done
// manually, gtk_check_menu_item_set_inconsistent() only affects visual
// appearance, it doesn’t affect the semantics of the widget.
func (c checkMenuItem) SetInconsistent(setting bool) {
	var arg0 *C.GtkCheckMenuItem
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_menu_item_set_inconsistent(arg0, setting)
}

// Toggled emits the CheckMenuItem::toggled signal.
func (c checkMenuItem) Toggled() {
	var arg0 *C.GtkCheckMenuItem

	arg0 = (*C.GtkCheckMenuItem)(unsafe.Pointer(c.Native()))

	C.gtk_check_menu_item_toggled(arg0)
}

type CheckMenuItemPrivate struct {
	native C.GtkCheckMenuItemPrivate
}

// WrapCheckMenuItemPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCheckMenuItemPrivate(ptr unsafe.Pointer) *CheckMenuItemPrivate {
	if ptr == nil {
		return nil
	}

	return (*CheckMenuItemPrivate)(ptr)
}

func marshalCheckMenuItemPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCheckMenuItemPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CheckMenuItemPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}