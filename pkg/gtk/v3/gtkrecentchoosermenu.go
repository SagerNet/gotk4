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
		{T: externglib.Type(C.gtk_recent_chooser_menu_get_type()), F: marshalRecentChooserMenu},
	})
}

// RecentChooserMenu is a widget suitable for displaying recently used files
// inside a menu. It can be used to set a sub-menu of a MenuItem using
// gtk_menu_item_set_submenu(), or as the menu of a MenuToolButton.
//
// Note that RecentChooserMenu does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
// Note also that RecentChooserMenu does not support multiple filters, as it has
// no way to let the user choose between them as the RecentChooserWidget and
// RecentChooserDialog widgets do. Thus using gtk_recent_chooser_add_filter() on
// a RecentChooserMenu widget will yield the same effects as using
// gtk_recent_chooser_set_filter(), replacing any currently set filter with the
// supplied filter; gtk_recent_chooser_remove_filter() will remove any currently
// set RecentFilter object and will unset the current filter;
// gtk_recent_chooser_list_filters() will return a list containing a single
// RecentFilter object.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserMenu interface {
	Menu
	Activatable
	Buildable
	RecentChooser

	// ShowNumbers returns the value set by
	// gtk_recent_chooser_menu_set_show_numbers().
	ShowNumbers() bool
	// SetShowNumbers sets whether a number should be added to the items of
	// @menu. The numbers are shown to provide a unique character for a mnemonic
	// to be used inside ten menu item’s label. Only the first the items get a
	// number to avoid clashes.
	SetShowNumbers(showNumbers bool)
}

// recentChooserMenu implements the RecentChooserMenu interface.
type recentChooserMenu struct {
	Menu
	Activatable
	Buildable
	RecentChooser
}

var _ RecentChooserMenu = (*recentChooserMenu)(nil)

// WrapRecentChooserMenu wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentChooserMenu(obj *externglib.Object) RecentChooserMenu {
	return RecentChooserMenu{
		Menu:          WrapMenu(obj),
		Activatable:   WrapActivatable(obj),
		Buildable:     WrapBuildable(obj),
		RecentChooser: WrapRecentChooser(obj),
	}
}

func marshalRecentChooserMenu(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentChooserMenu(obj), nil
}

// NewRecentChooserMenu constructs a class RecentChooserMenu.
func NewRecentChooserMenu() RecentChooserMenu {
	var cret C.GtkRecentChooserMenu
	var ret1 RecentChooserMenu

	cret = C.gtk_recent_chooser_menu_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RecentChooserMenu)

	return ret1
}

// NewRecentChooserMenuForManager constructs a class RecentChooserMenu.
func NewRecentChooserMenuForManager(manager RecentManager) RecentChooserMenu {
	var arg1 *C.GtkRecentManager

	arg1 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	var cret C.GtkRecentChooserMenu
	var ret1 RecentChooserMenu

	cret = C.gtk_recent_chooser_menu_new_for_manager(manager)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(RecentChooserMenu)

	return ret1
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
func (m recentChooserMenu) ShowNumbers() bool {
	var arg0 *C.GtkRecentChooserMenu

	arg0 = (*C.GtkRecentChooserMenu)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_recent_chooser_menu_get_show_numbers(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetShowNumbers sets whether a number should be added to the items of
// @menu. The numbers are shown to provide a unique character for a mnemonic
// to be used inside ten menu item’s label. Only the first the items get a
// number to avoid clashes.
func (m recentChooserMenu) SetShowNumbers(showNumbers bool) {
	var arg0 *C.GtkRecentChooserMenu
	var arg1 C.gboolean

	arg0 = (*C.GtkRecentChooserMenu)(unsafe.Pointer(m.Native()))
	if showNumbers {
		arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_menu_set_show_numbers(arg0, showNumbers)
}

type RecentChooserMenuPrivate struct {
	native C.GtkRecentChooserMenuPrivate
}

// WrapRecentChooserMenuPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentChooserMenuPrivate(ptr unsafe.Pointer) *RecentChooserMenuPrivate {
	if ptr == nil {
		return nil
	}

	return (*RecentChooserMenuPrivate)(ptr)
}

func marshalRecentChooserMenuPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecentChooserMenuPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecentChooserMenuPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}