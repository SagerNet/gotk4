// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_arrow_placement_get_type()), F: marshalArrowPlacement},
		{T: externglib.Type(C.gtk_menu_get_type()), F: marshalMenuer},
	})
}

// ArrowPlacement: used to specify the placement of scroll arrows in scrolling
// menus.
type ArrowPlacement int

const (
	// ArrowsBoth: place one arrow on each end of the menu.
	ArrowsBoth ArrowPlacement = iota
	// ArrowsStart: place both arrows at the top of the menu.
	ArrowsStart
	// ArrowsEnd: place both arrows at the bottom of the menu.
	ArrowsEnd
)

func marshalArrowPlacement(p uintptr) (interface{}, error) {
	return ArrowPlacement(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ArrowPlacement.
func (a ArrowPlacement) String() string {
	switch a {
	case ArrowsBoth:
		return "Both"
	case ArrowsStart:
		return "Start"
	case ArrowsEnd:
		return "End"
	default:
		return fmt.Sprintf("ArrowPlacement(%d)", a)
	}
}

// Menu is a MenuShell that implements a drop down menu consisting of a list of
// MenuItem objects which can be navigated and activated by the user to perform
// application functions.
//
// A Menu is most commonly dropped down by activating a MenuItem in a MenuBar or
// popped up by activating a MenuItem in another Menu.
//
// A Menu can also be popped up by activating a ComboBox. Other composite
// widgets such as the Notebook can pop up a Menu as well.
//
// Applications can display a Menu as a popup menu by calling the
// gtk_menu_popup() function. The example below shows how an application can pop
// up a menu when the 3rd mouse button is pressed.
//
// Connecting the popup signal handler.
//
//    menu
//    ├── arrow.top
//    ├── <child>
//    ┊
//    ├── <child>
//    ╰── arrow.bottom
//
// The main CSS node of GtkMenu has name menu, and there are two subnodes with
// name arrow, for scrolling menu arrows. These subnodes get the .top and
// .bottom style classes.
type Menu struct {
	MenuShell
}

func WrapMenu(obj *externglib.Object) *Menu {
	return &Menu{
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
	}
}

func marshalMenuer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenu(obj), nil
}

// NewMenu creates a new Menu
func NewMenu() *Menu {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_new()

	var _menu *Menu // out

	_menu = WrapMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _menu
}

// NewMenuFromModel creates a Menu and populates it with menu items and submenus
// according to model.
//
// The created menu items are connected to actions found in the
// ApplicationWindow to which the menu belongs - typically by means of being
// attached to a widget (see gtk_menu_attach_to_widget()) that is contained
// within the ApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group() on the
// menu's attach widget or on any of its parent widgets.
func NewMenuFromModel(model gio.MenuModeller) *Menu {
	var _arg1 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_menu_new_from_model(_arg1)
	runtime.KeepAlive(model)

	var _menu *Menu // out

	_menu = WrapMenu(externglib.Take(unsafe.Pointer(_cret)))

	return _menu
}

// Attach adds a new MenuItem to a (table) menu. The number of “cells” that an
// item will occupy is specified by left_attach, right_attach, top_attach and
// bottom_attach. These each represent the leftmost, rightmost, uppermost and
// lower column and row numbers of the table. (Columns and rows are indexed from
// zero).
//
// Note that this function is not related to gtk_menu_detach().
func (menu *Menu) Attach(child Widgetter, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var _arg0 *C.GtkMenu   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.guint      // out
	var _arg3 C.guint      // out
	var _arg4 C.guint      // out
	var _arg5 C.guint      // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_menu_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(child)
	runtime.KeepAlive(leftAttach)
	runtime.KeepAlive(rightAttach)
	runtime.KeepAlive(topAttach)
	runtime.KeepAlive(bottomAttach)
}

// Detach detaches the menu from the widget to which it had been attached. This
// function will call the callback function, detacher, provided when the
// gtk_menu_attach_to_widget() function was called.
func (menu *Menu) Detach() {
	var _arg0 *C.GtkMenu // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_detach(_arg0)
	runtime.KeepAlive(menu)
}

// AccelGroup gets the AccelGroup which holds global accelerators for the menu.
// See gtk_menu_set_accel_group().
func (menu *Menu) AccelGroup() *AccelGroup {
	var _arg0 *C.GtkMenu       // out
	var _cret *C.GtkAccelGroup // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_accel_group(_arg0)
	runtime.KeepAlive(menu)

	var _accelGroup *AccelGroup // out

	_accelGroup = WrapAccelGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _accelGroup
}

// AccelPath retrieves the accelerator path set on the menu.
func (menu *Menu) AccelPath() string {
	var _arg0 *C.GtkMenu // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_accel_path(_arg0)
	runtime.KeepAlive(menu)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Active returns the selected menu item from the menu. This is used by the
// ComboBox.
func (menu *Menu) Active() Widgetter {
	var _arg0 *C.GtkMenu   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_active(_arg0)
	runtime.KeepAlive(menu)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// AttachWidget returns the Widget that the menu is attached to.
func (menu *Menu) AttachWidget() Widgetter {
	var _arg0 *C.GtkMenu   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_attach_widget(_arg0)
	runtime.KeepAlive(menu)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Monitor retrieves the number of the monitor on which to show the menu.
func (menu *Menu) Monitor() int {
	var _arg0 *C.GtkMenu // out
	var _cret C.gint     // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_monitor(_arg0)
	runtime.KeepAlive(menu)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ReserveToggleSize returns whether the menu reserves space for toggles and
// icons, regardless of their actual presence.
func (menu *Menu) ReserveToggleSize() bool {
	var _arg0 *C.GtkMenu // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_reserve_toggle_size(_arg0)
	runtime.KeepAlive(menu)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TearoffState returns whether the menu is torn off. See
// gtk_menu_set_tearoff_state().
//
// Deprecated: since version 3.10.
func (menu *Menu) TearoffState() bool {
	var _arg0 *C.GtkMenu // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_tearoff_state(_arg0)
	runtime.KeepAlive(menu)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title returns the title of the menu. See gtk_menu_set_title().
//
// Deprecated: since version 3.10.
func (menu *Menu) Title() string {
	var _arg0 *C.GtkMenu // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	_cret = C.gtk_menu_get_title(_arg0)
	runtime.KeepAlive(menu)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PlaceOnMonitor places menu on the given monitor.
func (menu *Menu) PlaceOnMonitor(monitor *gdk.Monitor) {
	var _arg0 *C.GtkMenu    // out
	var _arg1 *C.GdkMonitor // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gtk_menu_place_on_monitor(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(monitor)
}

// Popdown removes the menu from the screen.
func (menu *Menu) Popdown() {
	var _arg0 *C.GtkMenu // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_popdown(_arg0)
	runtime.KeepAlive(menu)
}

// ReorderChild moves child to a new position in the list of menu children.
func (menu *Menu) ReorderChild(child Widgetter, position int) {
	var _arg0 *C.GtkMenu   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_menu_reorder_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// Reposition repositions the menu according to its position function.
func (menu *Menu) Reposition() {
	var _arg0 *C.GtkMenu // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))

	C.gtk_menu_reposition(_arg0)
	runtime.KeepAlive(menu)
}

// SetAccelGroup: set the AccelGroup which holds global accelerators for the
// menu. This accelerator group needs to also be added to all windows that this
// menu is being used in with gtk_window_add_accel_group(), in order for those
// windows to support all the accelerators contained in this group.
func (menu *Menu) SetAccelGroup(accelGroup *AccelGroup) {
	var _arg0 *C.GtkMenu       // out
	var _arg1 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if accelGroup != nil {
		_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))
	}

	C.gtk_menu_set_accel_group(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(accelGroup)
}

// SetAccelPath sets an accelerator path for this menu from which accelerator
// paths for its immediate children, its menu items, can be constructed. The
// main purpose of this function is to spare the programmer the inconvenience of
// having to call gtk_menu_item_set_accel_path() on each menu item that should
// support runtime user changable accelerators. Instead, by just calling
// gtk_menu_set_accel_path() on their parent, each menu item of this menu, that
// contains a label describing its purpose, automatically gets an accel path
// assigned.
//
// For example, a menu containing menu items “New” and “Exit”, will, after
// gtk_menu_set_accel_path (menu, "<Gnumeric-Sheet>/File"); has been called,
// assign its items the accel paths: "<Gnumeric-Sheet>/File/New" and
// "<Gnumeric-Sheet>/File/Exit".
//
// Assigning accel paths to menu items then enables the user to change their
// accelerators at runtime. More details about accelerator paths and their
// default setups can be found at gtk_accel_map_add_entry().
//
// Note that accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
func (menu *Menu) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkMenu // out
	var _arg1 *C.gchar   // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if accelPath != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(accelPath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_menu_set_accel_path(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(accelPath)
}

// SetActive selects the specified menu item within the menu. This is used by
// the ComboBox and should not be used by anyone else.
func (menu *Menu) SetActive(index uint) {
	var _arg0 *C.GtkMenu // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = C.guint(index)

	C.gtk_menu_set_active(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(index)
}

// SetMonitor informs GTK+ on which monitor a menu should be popped up. See
// gdk_monitor_get_geometry().
//
// This function should be called from a MenuPositionFunc if the menu should not
// appear on the same monitor as the pointer. This information can’t be reliably
// inferred from the coordinates returned by a MenuPositionFunc, since, for very
// long menus, these coordinates may extend beyond the monitor boundaries or
// even the screen boundaries.
func (menu *Menu) SetMonitor(monitorNum int) {
	var _arg0 *C.GtkMenu // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	_arg1 = C.gint(monitorNum)

	C.gtk_menu_set_monitor(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(monitorNum)
}

// SetReserveToggleSize sets whether the menu should reserve space for drawing
// toggles or icons, regardless of their actual presence.
func (menu *Menu) SetReserveToggleSize(reserveToggleSize bool) {
	var _arg0 *C.GtkMenu // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if reserveToggleSize {
		_arg1 = C.TRUE
	}

	C.gtk_menu_set_reserve_toggle_size(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(reserveToggleSize)
}

// SetScreen sets the Screen on which the menu will be displayed.
func (menu *Menu) SetScreen(screen *gdk.Screen) {
	var _arg0 *C.GtkMenu   // out
	var _arg1 *C.GdkScreen // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if screen != nil {
		_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	}

	C.gtk_menu_set_screen(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(screen)
}

// SetTearoffState changes the tearoff state of the menu. A menu is normally
// displayed as drop down menu which persists as long as the menu is active. It
// can also be displayed as a tearoff menu which persists until it is closed or
// reattached.
//
// Deprecated: since version 3.10.
func (menu *Menu) SetTearoffState(tornOff bool) {
	var _arg0 *C.GtkMenu // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if tornOff {
		_arg1 = C.TRUE
	}

	C.gtk_menu_set_tearoff_state(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(tornOff)
}

// SetTitle sets the title string for the menu.
//
// The title is displayed when the menu is shown as a tearoff menu. If title is
// NULL, the menu will see if it is attached to a parent menu item, and if so it
// will try to use the same text as that menu item’s label.
//
// Deprecated: since version 3.10.
func (menu *Menu) SetTitle(title string) {
	var _arg0 *C.GtkMenu // out
	var _arg1 *C.gchar   // out

	_arg0 = (*C.GtkMenu)(unsafe.Pointer(menu.Native()))
	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_menu_set_title(_arg0, _arg1)
	runtime.KeepAlive(menu)
	runtime.KeepAlive(title)
}

// MenuGetForAttachWidget returns a list of the menus which are attached to this
// widget. This list is owned by GTK+ and must not be modified.
func MenuGetForAttachWidget(widget Widgetter) []Widgetter {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GList     // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_menu_get_for_attach_widget(_arg1)
	runtime.KeepAlive(widget)

	var _list []Widgetter // out

	_list = make([]Widgetter, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkWidget)(v)
		var dst Widgetter // out
		dst = (externglib.CastObject(externglib.Take(unsafe.Pointer(src)))).(Widgetter)
		_list = append(_list, dst)
	})

	return _list
}
