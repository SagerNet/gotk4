// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_radio_menu_item_get_type()), F: marshalRadioMenuItem},
	})
}

// RadioMenuItem: a radio menu item is a check menu item that belongs to a
// group. At each instant exactly one of the radio menu items from a group is
// selected.
//
// The group list does not need to be freed, as each RadioMenuItem will remove
// itself and its list item when it is destroyed.
//
// The correct way to create a group of radio menu items is approximatively
// this:
//
// How to create a group of radio menu items.
//
//    menuitem
//    ├── radio.left
//    ╰── <child>
//
// GtkRadioMenuItem has a main CSS node with name menuitem, and a subnode with
// name radio, which gets the .left or .right style class.
type RadioMenuItem interface {
	CheckMenuItem
	Actionable
	Activatable
	Buildable

	// Group returns the group to which the radio menu item belongs, as a #GList
	// of RadioMenuItem. The list belongs to GTK+ and should not be freed.
	Group() *glib.SList
	// JoinGroup joins a RadioMenuItem object to the group of another
	// RadioMenuItem object.
	//
	// This function should be used by language bindings to avoid the memory
	// manangement of the opaque List of gtk_radio_menu_item_get_group() and
	// gtk_radio_menu_item_set_group().
	//
	// A common way to set up a group of RadioMenuItem instances is:
	//
	//      GtkRadioMenuItem *last_item = NULL;
	//
	//      while ( ...more items to add... )
	//        {
	//          GtkRadioMenuItem *radio_item;
	//
	//          radio_item = gtk_radio_menu_item_new (...);
	//
	//          gtk_radio_menu_item_join_group (radio_item, last_item);
	//          last_item = radio_item;
	//        }
	JoinGroup(groupSource RadioMenuItem)
	// SetGroup sets the group of a radio menu item, or changes it.
	SetGroup(group *glib.SList)
}

// radioMenuItem implements the RadioMenuItem interface.
type radioMenuItem struct {
	CheckMenuItem
	Actionable
	Activatable
	Buildable
}

var _ RadioMenuItem = (*radioMenuItem)(nil)

// WrapRadioMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioMenuItem(obj *externglib.Object) RadioMenuItem {
	return RadioMenuItem{
		CheckMenuItem: WrapCheckMenuItem(obj),
		Actionable:    WrapActionable(obj),
		Activatable:   WrapActivatable(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioMenuItem(obj), nil
}

// NewRadioMenuItem constructs a class RadioMenuItem.
func NewRadioMenuItem(group *glib.SList) RadioMenuItem {
	var arg1 *C.GSList

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new(group)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// NewRadioMenuItemFromWidget constructs a class RadioMenuItem.
func NewRadioMenuItemFromWidget(group RadioMenuItem) RadioMenuItem {
	var arg1 *C.GtkRadioMenuItem

	arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new_from_widget(group)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// NewRadioMenuItemWithLabel constructs a class RadioMenuItem.
func NewRadioMenuItemWithLabel(group *glib.SList, label string) RadioMenuItem {
	var arg1 *C.GSList
	var arg2 *C.gchar

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new_with_label(group, label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// NewRadioMenuItemWithLabelFromWidget constructs a class RadioMenuItem.
func NewRadioMenuItemWithLabelFromWidget(group RadioMenuItem, label string) RadioMenuItem {
	var arg1 *C.GtkRadioMenuItem
	var arg2 *C.gchar

	arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new_with_label_from_widget(group, label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// NewRadioMenuItemWithMnemonic constructs a class RadioMenuItem.
func NewRadioMenuItemWithMnemonic(group *glib.SList, label string) RadioMenuItem {
	var arg1 *C.GSList
	var arg2 *C.gchar

	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new_with_mnemonic(group, label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// NewRadioMenuItemWithMnemonicFromWidget constructs a class RadioMenuItem.
func NewRadioMenuItemWithMnemonicFromWidget(group RadioMenuItem, label string) RadioMenuItem {
	var arg1 *C.GtkRadioMenuItem
	var arg2 *C.gchar

	arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.GtkRadioMenuItem
	var goret1 RadioMenuItem

	cret = C.gtk_radio_menu_item_new_with_mnemonic_from_widget(group, label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RadioMenuItem)

	return goret1
}

// Group returns the group to which the radio menu item belongs, as a #GList
// of RadioMenuItem. The list belongs to GTK+ and should not be freed.
func (r radioMenuItem) Group() *glib.SList {
	var arg0 *C.GtkRadioMenuItem

	arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(r.Native()))

	var cret *C.GSList
	var goret1 *glib.SList

	cret = C.gtk_radio_menu_item_get_group(arg0)

	goret1 = glib.WrapSList(unsafe.Pointer(cret))

	return goret1
}

// JoinGroup joins a RadioMenuItem object to the group of another
// RadioMenuItem object.
//
// This function should be used by language bindings to avoid the memory
// manangement of the opaque List of gtk_radio_menu_item_get_group() and
// gtk_radio_menu_item_set_group().
//
// A common way to set up a group of RadioMenuItem instances is:
//
//      GtkRadioMenuItem *last_item = NULL;
//
//      while ( ...more items to add... )
//        {
//          GtkRadioMenuItem *radio_item;
//
//          radio_item = gtk_radio_menu_item_new (...);
//
//          gtk_radio_menu_item_join_group (radio_item, last_item);
//          last_item = radio_item;
//        }
func (r radioMenuItem) JoinGroup(groupSource RadioMenuItem) {
	var arg0 *C.GtkRadioMenuItem
	var arg1 *C.GtkRadioMenuItem

	arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_menu_item_join_group(arg0, groupSource)
}

// SetGroup sets the group of a radio menu item, or changes it.
func (r radioMenuItem) SetGroup(group *glib.SList) {
	var arg0 *C.GtkRadioMenuItem
	var arg1 *C.GSList

	arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	C.gtk_radio_menu_item_set_group(arg0, group)
}

type RadioMenuItemPrivate struct {
	native C.GtkRadioMenuItemPrivate
}

// WrapRadioMenuItemPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRadioMenuItemPrivate(ptr unsafe.Pointer) *RadioMenuItemPrivate {
	if ptr == nil {
		return nil
	}

	return (*RadioMenuItemPrivate)(ptr)
}

func marshalRadioMenuItemPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRadioMenuItemPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RadioMenuItemPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}
