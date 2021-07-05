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
		{T: externglib.Type(C.gtk_application_inhibit_flags_get_type()), F: marshalApplicationInhibitFlags},
		{T: externglib.Type(C.gtk_application_get_type()), F: marshalApplication},
	})
}

// ApplicationInhibitFlags types of user actions that may be blocked by
// `GtkApplication`.
//
// See [method@Gtk.Application.inhibit].
type ApplicationInhibitFlags int

const (
	// ApplicationInhibitFlagsLogout: inhibit ending the user session by logging
	// out or by shutting down the computer
	ApplicationInhibitFlagsLogout ApplicationInhibitFlags = 0b1
	// ApplicationInhibitFlagsSwitch: inhibit user switching
	ApplicationInhibitFlagsSwitch ApplicationInhibitFlags = 0b10
	// ApplicationInhibitFlagsSuspend: inhibit suspending the session or
	// computer
	ApplicationInhibitFlagsSuspend ApplicationInhibitFlags = 0b100
	// ApplicationInhibitFlagsIdle: inhibit the session being marked as idle
	// (and possibly locked)
	ApplicationInhibitFlagsIdle ApplicationInhibitFlags = 0b1000
)

func marshalApplicationInhibitFlags(p uintptr) (interface{}, error) {
	return ApplicationInhibitFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Application: `GtkApplication` is a high-level API for writing applications.
//
// It supports many aspects of writing a GTK application in a convenient
// fashion, without enforcing a one-size-fits-all model.
//
// Currently, `GtkApplication` handles GTK initialization, application
// uniqueness, session management, provides some basic scriptability and desktop
// shell integration by exporting actions and menus and manages a list of
// toplevel windows whose life-cycle is automatically tied to the life-cycle of
// your application.
//
// While `GtkApplication` works fine with plain [class@Gtk.Window]s, it is
// recommended to use it together with [class@Gtk.ApplicationWindow].
//
//
// Automatic resources
//
// `GtkApplication` will automatically load menus from the `GtkBuilder` resource
// located at "gtk/menus.ui", relative to the application's resource base path
// (see `g_application_set_resource_base_path()`). The menu with the ID
// "menubar" is taken as the application's menubar. Additional menus (most
// interesting submenus) can be named and accessed via
// [method@Gtk.Application.get_menu_by_id] which allows for dynamic population
// of a part of the menu structure.
//
// It is also possible to provide the menubar manually using
// [method@Gtk.Application.set_menubar].
//
// `GtkApplication` will also automatically setup an icon search path for the
// default icon theme by appending "icons" to the resource base path. This
// allows your application to easily store its icons as resources. See
// [method@Gtk.IconTheme.add_resource_path] for more information.
//
// If there is a resource located at "gtk/help-overlay.ui" which defines a
// [class@Gtk.ShortcutsWindow] with ID "help_overlay" then `GtkApplication`
// associates an instance of this shortcuts window with each
// [class@Gtk.ApplicationWindow] and sets up the keyboard accelerator
// <kbd>Control</kbd>+<kbd>?</kbd> to open it. To create a menu item that
// displays the shortcuts window, associate the item with the action
// `win.show-help-overlay`.
//
//
// A simple application
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/bp/bloatpad.c) is
// available in the GTK source code repository
//
// `GtkApplication` optionally registers with a session manager of the users
// session (if you set the [property@Gtk.Application:register-session] property)
// and offers various functionality related to the session life-cycle.
//
// An application can block various ways to end the session with the
// [method@Gtk.Application.inhibit] function. Typical use cases for this kind of
// inhibiting are long-running, uninterruptible operations, such as burning a CD
// or performing a disk backup. The session manager may not honor the inhibitor,
// but it can be expected to inform the user about the negative consequences of
// ending the session while inhibitors are present.
//
//
// See Also
//
// HowDoI: Using GtkApplication (https://wiki.gnome.org/HowDoI/GtkApplication),
// Getting Started with GTK: Basics (getting_started.html#basics)
type Application interface {
	gio.Application

	// AddWindow adds a window to `application`.
	//
	// This call can only happen after the `application` has started; typically,
	// you should add new application windows in response to the emission of the
	// `GApplication::activate` signal.
	//
	// This call is equivalent to setting the [property@Gtk.Window:application]
	// property of `window` to `application`.
	//
	// Normally, the connection between the application and the window will
	// remain until the window is destroyed, but you can explicitly remove it
	// with [method@Gtk.Application.remove_window].
	//
	// GTK will keep the `application` running as long as it has any windows.
	AddWindow(window Window)
	// AccelsForAction gets the accelerators that are currently associated with
	// the given action.
	AccelsForAction(detailedActionName string) []string
	// ActionsForAccel returns the list of actions (possibly empty) that `accel`
	// maps to.
	//
	// Each item in the list is a detailed action name in the usual form.
	//
	// This might be useful to discover if an accel already exists in order to
	// prevent installation of a conflicting accelerator (from an accelerator
	// editor or a plugin system, for example). Note that having more than one
	// action per accelerator may not be a bad thing and might make sense in
	// cases where the actions never appear in the same context.
	//
	// In case there are no actions for a given accelerator, an empty array is
	// returned. `NULL` is never returned.
	//
	// It is a programmer error to pass an invalid accelerator string.
	//
	// If you are unsure, check it with [func@Gtk.accelerator_parse] first.
	ActionsForAccel(accel string) []string
	// ActiveWindow gets the “active” window for the application.
	//
	// The active window is the one that was most recently focused (within the
	// application). This window may not have the focus at the moment if another
	// application has it — this is just the most recently-focused window within
	// this application.
	ActiveWindow() Window
	// MenuByID gets a menu from automatically loaded resources.
	//
	// See the section on Automatic resources
	// (class.Application.html#automatic-resources) for more information.
	MenuByID(id string) gio.Menu
	// Menubar returns the menu model that has been set with
	// [method@Gtk.Application.set_menubar].
	Menubar() gio.MenuModel
	// WindowByID returns the [class@Gtk.ApplicationWindow] with the given ID.
	//
	// The ID of a `GtkApplicationWindow` can be retrieved with
	// [method@Gtk.ApplicationWindow.get_id].
	WindowByID(id uint) Window
	// Inhibit: inform the session manager that certain types of actions should
	// be inhibited.
	//
	// This is not guaranteed to work on all platforms and for all types of
	// actions.
	//
	// Applications should invoke this method when they begin an operation that
	// should not be interrupted, such as creating a CD or DVD. The types of
	// actions that may be blocked are specified by the `flags` parameter. When
	// the application completes the operation it should call
	// [method@Gtk.Application.uninhibit] to remove the inhibitor. Note that an
	// application can have multiple inhibitors, and all of them must be
	// individually removed. Inhibitors are also cleared when the application
	// exits.
	//
	// Applications should not expect that they will always be able to block the
	// action. In most cases, users will be given the option to force the action
	// to take place.
	//
	// The `reason` message should be short and to the point.
	//
	// If `window` is given, the session manager may point the user to this
	// window to find out more about why the action is inhibited.
	Inhibit(window Window, flags ApplicationInhibitFlags, reason string) uint
	// ListActionDescriptions lists the detailed action names which have
	// associated accelerators.
	//
	// See [method@Gtk.Application.set_accels_for_action].
	ListActionDescriptions() []string
	// RemoveWindow: remove a window from `application`.
	//
	// If `window` belongs to `application` then this call is equivalent to
	// setting the [property@Gtk.Window:application] property of `window` to
	// `NULL`.
	//
	// The application may stop running as a result of a call to this function,
	// if `window` was the last window of the `application`.
	RemoveWindow(window Window)
	// SetAccelsForAction sets zero or more keyboard accelerators that will
	// trigger the given action.
	//
	// The first item in `accels` will be the primary accelerator, which may be
	// displayed in the UI.
	//
	// To remove all accelerators for an action, use an empty, zero-terminated
	// array for `accels`.
	//
	// For the `detailed_action_name`, see `g_action_parse_detailed_name()` and
	// `g_action_print_detailed_name()`.
	SetAccelsForAction(detailedActionName string, accels []string)
	// SetMenubar sets or unsets the menubar for windows of `application`.
	//
	// This is a menubar in the traditional sense.
	//
	// This can only be done in the primary instance of the application, after
	// it has been registered. `GApplication::startup` is a good place to call
	// this.
	//
	// Depending on the desktop environment, this may appear at the top of each
	// window, or at the top of the screen. In some environments, if both the
	// application menu and the menubar are set, the application menu will be
	// presented as if it were the first item of the menubar. Other environments
	// treat the two as completely separate — for example, the application menu
	// may be rendered by the desktop shell while the menubar (if set) remains
	// in each individual window.
	//
	// Use the base `GActionMap` interface to add actions, to respond to the
	// user selecting these menu items.
	SetMenubar(menubar gio.MenuModel)
	// Uninhibit removes an inhibitor that has been previously established.
	//
	// See [method@Gtk.Application.inhibit].
	//
	// Inhibitors are also cleared when the application exits.
	Uninhibit(cookie uint)
}

// application implements the Application class.
type application struct {
	gio.Application
}

// WrapApplication wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplication(obj *externglib.Object) Application {
	return application{
		Application: gio.WrapApplication(obj),
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplication(obj), nil
}

// NewApplication creates a new `GtkApplication` instance.
//
// When using `GtkApplication`, it is not necessary to call [func@Gtk.init]
// manually. It is called as soon as the application gets registered as the
// primary instance.
//
// Concretely, [func@Gtk.init] is called in the default handler for the
// `GApplication::startup` signal. Therefore, `GtkApplication` subclasses should
// always chain up in their `GApplication::startup` handler before using any GTK
// API.
//
// Note that commandline arguments are not passed to [func@Gtk.init].
//
// If `application_id` is not nil, then it must be valid. See
// `g_application_id_is_valid()`.
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled.
func NewApplication(applicationId string, flags gio.ApplicationFlags) Application {
	var _arg1 *C.char             // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.GtkApplication   // in

	_arg1 = (*C.char)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.gtk_application_new(_arg1, _arg2)

	var _application Application // out

	_application = WrapApplication(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _application
}

func (a application) AddWindow(window Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_add_window(_arg0, _arg1)
}

func (a application) AccelsForAction(detailedActionName string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.char           // out
	var _cret **C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_accels_for_action(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) ActionsForAccel(accel string) []string {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.char           // out
	var _cret **C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(accel))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_actions_for_accel(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) ActiveWindow() Window {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_get_active_window(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (a application) MenuByID(id string) gio.Menu {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.char           // out
	var _cret *C.GMenu          // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_application_get_menu_by_id(_arg0, _arg1)

	var _menu gio.Menu // out

	_menu = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.Menu)

	return _menu
}

func (a application) Menubar() gio.MenuModel {
	var _arg0 *C.GtkApplication // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_get_menubar(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (a application) WindowByID(id uint) Window {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out
	var _cret *C.GtkWindow      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(id)

	_cret = C.gtk_application_get_window_by_id(_arg0, _arg1)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (a application) Inhibit(window Window, flags ApplicationInhibitFlags, reason string) uint {
	var _arg0 *C.GtkApplication            // out
	var _arg1 *C.GtkWindow                 // out
	var _arg2 C.GtkApplicationInhibitFlags // out
	var _arg3 *C.char                      // out
	var _cret C.guint                      // in

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.GtkApplicationInhibitFlags(flags)
	_arg3 = (*C.char)(C.CString(reason))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_application_inhibit(_arg0, _arg1, _arg2, _arg3)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (a application) ListActionDescriptions() []string {
	var _arg0 *C.GtkApplication // out
	var _cret **C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_application_list_action_descriptions(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (a application) RemoveWindow(window Window) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_application_remove_window(_arg0, _arg1)
}

func (a application) SetAccelsForAction(detailedActionName string, accels []string) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.char           // out
	var _arg2 **C.char

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(detailedActionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(accels)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(accels))
		for i := range accels {
			out[i] = (*C.char)(C.CString(accels[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_application_set_accels_for_action(_arg0, _arg1, _arg2)
}

func (a application) SetMenubar(menubar gio.MenuModel) {
	var _arg0 *C.GtkApplication // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menubar.Native()))

	C.gtk_application_set_menubar(_arg0, _arg1)
}

func (a application) Uninhibit(cookie uint) {
	var _arg0 *C.GtkApplication // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(cookie)

	C.gtk_application_uninhibit(_arg0, _arg1)
}
