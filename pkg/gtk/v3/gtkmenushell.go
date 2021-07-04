// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_menu_shell_get_type()), F: marshalMenuShell},
	})
}

// MenuShell: a MenuShell is the abstract base class used to derive the Menu and
// MenuBar subclasses.
//
// A MenuShell is a container of MenuItem objects arranged in a list which can
// be navigated, selected, and activated by the user to perform application
// functions. A MenuItem can have a submenu associated with it, allowing for
// nested hierarchical menus.
//
//
// Terminology
//
// A menu item can be “selected”, this means that it is displayed in the
// prelight state, and if it has a submenu, that submenu will be popped up.
//
// A menu is “active” when it is visible onscreen and the user is selecting from
// it. A menubar is not active until the user clicks on one of its menuitems.
// When a menu is active, passing the mouse over a submenu will pop it up.
//
// There is also is a concept of the current menu and a current menu item. The
// current menu item is the selected menu item that is furthest down in the
// hierarchy. (Every active menu shell does not necessarily contain a selected
// menu item, but if it does, then the parent menu shell must also contain a
// selected menu item.) The current menu is the menu that contains the current
// menu item. It will always have a GTK grab and receive all key presses.
type MenuShell interface {
	Container

	// ActivateItemMenuShell:
	ActivateItemMenuShell(menuItem Widget, forceDeactivate bool)
	// AppendMenuShell:
	AppendMenuShell(child MenuItem)
	// BindModelMenuShell:
	BindModelMenuShell(model gio.MenuModel, actionNamespace string, withSeparators bool)
	// CancelMenuShell:
	CancelMenuShell()
	// DeactivateMenuShell:
	DeactivateMenuShell()
	// DeselectMenuShell:
	DeselectMenuShell()
	// ParentShell:
	ParentShell() Widget
	// SelectedItem:
	SelectedItem() Widget
	// TakeFocus:
	TakeFocus() bool
	// InsertMenuShell:
	InsertMenuShell(child Widget, position int)
	// PrependMenuShell:
	PrependMenuShell(child Widget)
	// SelectFirstMenuShell:
	SelectFirstMenuShell(searchSensitive bool)
	// SelectItemMenuShell:
	SelectItemMenuShell(menuItem Widget)
	// SetTakeFocusMenuShell:
	SetTakeFocusMenuShell(takeFocus bool)
}

// menuShell implements the MenuShell class.
type menuShell struct {
	Container
}

// WrapMenuShell wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuShell(obj *externglib.Object) MenuShell {
	return menuShell{
		Container: WrapContainer(obj),
	}
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuShell(obj), nil
}

func (m menuShell) ActivateItemMenuShell(menuItem Widget, forceDeactivate bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))
	if forceDeactivate {
		_arg2 = C.TRUE
	}

	C.gtk_menu_shell_activate_item(_arg0, _arg1, _arg2)
}

func (m menuShell) AppendMenuShell(child MenuItem) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_append(_arg0, _arg1)
}

func (m menuShell) BindModelMenuShell(model gio.MenuModel, actionNamespace string, withSeparators bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GMenuModel   // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.gchar)(C.CString(actionNamespace))
	defer C.free(unsafe.Pointer(_arg2))
	if withSeparators {
		_arg3 = C.TRUE
	}

	C.gtk_menu_shell_bind_model(_arg0, _arg1, _arg2, _arg3)
}

func (m menuShell) CancelMenuShell() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_cancel(_arg0)
}

func (m menuShell) DeactivateMenuShell() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_deactivate(_arg0)
}

func (m menuShell) DeselectMenuShell() {
	var _arg0 *C.GtkMenuShell // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_deselect(_arg0)
}

func (m menuShell) ParentShell() Widget {
	var _arg0 *C.GtkMenuShell // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_shell_get_parent_shell(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (m menuShell) SelectedItem() Widget {
	var _arg0 *C.GtkMenuShell // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_shell_get_selected_item(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (m menuShell) TakeFocus() bool {
	var _arg0 *C.GtkMenuShell // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_shell_get_take_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m menuShell) InsertMenuShell(child Widget, position int) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_menu_shell_insert(_arg0, _arg1, _arg2)
}

func (m menuShell) PrependMenuShell(child Widget) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_prepend(_arg0, _arg1)
}

func (m menuShell) SelectFirstMenuShell(searchSensitive bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	if searchSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_menu_shell_select_first(_arg0, _arg1)
}

func (m menuShell) SelectItemMenuShell(menuItem Widget) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_shell_select_item(_arg0, _arg1)
}

func (m menuShell) SetTakeFocusMenuShell(takeFocus bool) {
	var _arg0 *C.GtkMenuShell // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	if takeFocus {
		_arg1 = C.TRUE
	}

	C.gtk_menu_shell_set_take_focus(_arg0, _arg1)
}

func (b menuShell) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b menuShell) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b menuShell) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b menuShell) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b menuShell) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b menuShell) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b menuShell) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
