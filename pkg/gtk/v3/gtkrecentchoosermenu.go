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
		{T: externglib.Type(C.gtk_recent_chooser_menu_get_type()), F: marshalRecentChooserMenuer},
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
type RecentChooserMenu struct {
	Menu

	Activatable
	RecentChooser
	*externglib.Object
}

func WrapRecentChooserMenu(obj *externglib.Object) *RecentChooserMenu {
	return &RecentChooserMenu{
		Menu: Menu{
			MenuShell: MenuShell{
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
			},
		},
		Activatable: Activatable{
			Object: obj,
		},
		RecentChooser: RecentChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalRecentChooserMenuer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentChooserMenu(obj), nil
}

// NewRecentChooserMenu creates a new RecentChooserMenu widget.
//
// This kind of widget shows the list of recently used resources as a menu, each
// item as a menu item. Each item inside the menu might have an icon,
// representing its MIME type, and a number, for mnemonic access.
//
// This widget implements the RecentChooser interface.
//
// This widget creates its own RecentManager object. See the
// gtk_recent_chooser_menu_new_for_manager() function to know how to create a
// RecentChooserMenu widget bound to another RecentManager object.
func NewRecentChooserMenu() *RecentChooserMenu {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_recent_chooser_menu_new()

	var _recentChooserMenu *RecentChooserMenu // out

	_recentChooserMenu = WrapRecentChooserMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserMenu
}

// NewRecentChooserMenuForManager creates a new RecentChooserMenu widget using
// manager as the underlying recently used resources manager.
//
// This is useful if you have implemented your own recent manager, or if you
// have a customized instance of a RecentManager object or if you wish to share
// a common RecentManager object among multiple RecentChooser widgets.
func NewRecentChooserMenuForManager(manager *RecentManager) *RecentChooserMenu {
	var _arg1 *C.GtkRecentManager // out
	var _cret *C.GtkWidget        // in

	_arg1 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_chooser_menu_new_for_manager(_arg1)
	runtime.KeepAlive(manager)

	var _recentChooserMenu *RecentChooserMenu // out

	_recentChooserMenu = WrapRecentChooserMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserMenu
}

// ShowNumbers returns the value set by
// gtk_recent_chooser_menu_set_show_numbers().
func (menu *RecentChooserMenu) ShowNumbers() bool {
	var _arg0 *C.GtkRecentChooserMenu // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkRecentChooserMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_recent_chooser_menu_get_show_numbers(_arg0)
	runtime.KeepAlive(menu)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetShowNumbers sets whether a number should be added to the items of menu.
// The numbers are shown to provide a unique character for a mnemonic to be used
// inside ten menu item’s label. Only the first the items get a number to avoid
// clashes.
func (menu *RecentChooserMenu) SetShowNumbers(showNumbers bool) {
	var _arg0 *C.GtkRecentChooserMenu // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkRecentChooserMenu)(unsafe.Pointer(menu.Native()))
	if showNumbers {
		_arg1 = C.TRUE
	}

	C.gtk_recent_chooser_menu_set_show_numbers(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(showNumbers)
}
