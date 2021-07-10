// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalyier},
	})
}

// yier describes PasswordEntry's methods.
type yier interface {
	gextras.Objector

	ExtraMenu() *gio.MenuModel
	ShowPeekIcon() bool
	SetExtraMenu(model gio.MenuModeller)
	SetShowPeekIcon(showPeekIcon bool)
}

// PasswordEntry: `GtkPasswordEntry` is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, `GtkPasswordEntry` will also place the text in
// a non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// `GtkPasswordEntry` provides only minimal API and should be used with the
// [iface@Gtk.Editable] API.
//
//
// CSS Nodes
//
// “` entry.password ╰── text ├── image.caps-lock-indicator ┊ “`
//
// `GtkPasswordEntry` has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// `GtkPasswordEntry` uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Editable
}

var _ yier = (*PasswordEntry)(nil)

func wrapyier(obj *externglib.Object) yier {
	return &PasswordEntry{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Accessible: Accessible{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		ConstraintTarget: ConstraintTarget{
			Object: obj,
		},
		Editable: Editable{
			Object: obj,
			Widget: Widget{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalyier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapyier(obj), nil
}

// NewPasswordEntry creates a `GtkPasswordEntry`.
func NewPasswordEntry() *PasswordEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_password_entry_new()

	var _passwordEntry *PasswordEntry // out

	_passwordEntry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*PasswordEntry)

	return _passwordEntry
}

// ExtraMenu gets the menu model set with gtk_password_entry_set_extra_menu().
func (entry *PasswordEntry) ExtraMenu() *gio.MenuModel {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret *C.GMenuModel       // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_extra_menu(_arg0)

	var _menuModel *gio.MenuModel // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.MenuModel)

	return _menuModel
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
func (entry *PasswordEntry) ShowPeekIcon() bool {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_show_peek_icon(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// @entry.
func (entry *PasswordEntry) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 *C.GMenuModel       // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_password_entry_set_extra_menu(_arg0, _arg1)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to reveal
// the contents.
//
// Setting this to false also hides the text again.
func (entry *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	if showPeekIcon {
		_arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(_arg0, _arg1)
}
