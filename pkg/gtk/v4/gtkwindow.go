// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_window_get_type()), F: marshalWindow},
	})
}

// Window: `GtkWindow` is a toplevel window which can contain other widgets.
//
// !An example GtkWindow (window.png)
//
// Windows normally have decorations that are under the control of the windowing
// system and allow the user to manipulate the window (resize it, move it, close
// it,...).
//
//
// GtkWindow as GtkBuildable
//
// The `GtkWindow` implementation of the [iface@Gtk.Buildable] interface
// supports setting a child as the titlebar by specifying “titlebar” as the
// “type” attribute of a <child> element.
//
//
// CSS nodes
//
// “` window.background [.csd / .solid-csd / .ssd] [.maximized / .fullscreen /
// .tiled] ├── <child> ╰── <titlebar child>.titlebar [.default-decoration] “`
//
// `GtkWindow` has a main CSS node with name window and style class .background.
//
// Style classes that are typically used with the main CSS node are .csd (when
// client-side decorations are in use), .solid-csd (for client-side decorations
// without invisible borders), .ssd (used by mutter when rendering server-side
// decorations). GtkWindow also represents window states with the following
// style classes on the main node: .maximized, .fullscreen, .tiled (when
// supported, also .tiled-top, .tiled-left, .tiled-right, .tiled-bottom).
//
// `GtkWindow` subclasses often add their own discriminating style classes, such
// as .dialog, .popup or .tooltip.
//
// Generally, some CSS properties don't make sense on the toplevel window node,
// such as margins or padding. When client-side decorations without invisible
// borders are in use (i.e. the .solid-csd style class is added to the main
// window node), the CSS border of the toplevel window is used for resize drags.
// In the .csd case, the shadow area outside of the window can be used to resize
// it.
//
// `GtkWindow` adds the .titlebar and .default-decoration style classes to the
// widget that is added as a titlebar child.
//
//
// Accessibility
//
// `GtkWindow` uses the GTK_ACCESSIBLE_ROLE_WINDOW role.
type Window interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsNative casts the class to the Native interface.
	AsNative() Native
	// AsRoot casts the class to the Root interface.
	AsRoot() Root
	// AsShortcutManager casts the class to the ShortcutManager interface.
	AsShortcutManager() ShortcutManager

	// Close requests that the window is closed.
	//
	// This is similar to what happens when a window manager close button is
	// clicked.
	//
	// This function can be used with close buttons in custom titlebars.
	Close()
	// Destroy: drop the internal reference GTK holds on toplevel windows.
	Destroy()
	// Fullscreen asks to place @window in the fullscreen state.
	//
	// Note that you shouldn’t assume the window is definitely fullscreen
	// afterward, because other entities (e.g. the user or window manager
	// unfullscreen it again, and not all window managers honor requests to
	// fullscreen windows.
	//
	// You can track the result of this operation via the
	// [property@Gdk.Toplevel:state] property, or by listening to notifications
	// of the [property@Gtk.Window:fullscreened] property.
	Fullscreen()
	// FullscreenOnMonitor asks to place @window in the fullscreen state on the
	// given @monitor.
	//
	// Note that you shouldn't assume the window is definitely fullscreen
	// afterward, or that the windowing system allows fullscreen windows on any
	// given monitor.
	//
	// You can track the result of this operation via the
	// [property@Gdk.Toplevel:state] property, or by listening to notifications
	// of the [property@Gtk.Window:fullscreened] property.
	FullscreenOnMonitor(monitor gdk.Monitor)
	// Application gets the `GtkApplication` associated with the window.
	Application() Application
	// Child gets the child widget of @window.
	Child() Widget
	// Decorated returns whether the window has been set to have decorations.
	Decorated() bool
	// DefaultSize gets the default size of the window.
	//
	// A value of 0 for the width or height indicates that a default size has
	// not been explicitly set for that dimension, so the “natural” size of the
	// window will be used.
	DefaultSize() (width int, height int)
	// DefaultWidget returns the default widget for @window.
	DefaultWidget() Widget
	// Deletable returns whether the window has been set to have a close button.
	Deletable() bool
	// DestroyWithParent returns whether the window will be destroyed with its
	// transient parent.
	DestroyWithParent() bool
	// Focus retrieves the current focused widget within the window.
	//
	// Note that this is the widget that would have the focus if the toplevel
	// window focused; if the toplevel window is not focused then
	// `gtk_widget_has_focus (widget)` will not be true for the widget.
	Focus() Widget
	// FocusVisible gets whether “focus rectangles” are supposed to be visible.
	FocusVisible() bool
	// Group returns the group for @window.
	//
	// If the window has no group, then the default group is returned.
	Group() WindowGroup
	// HandleMenubarAccel returns whether this window reacts to F10 key presses
	// by activating a menubar it contains.
	HandleMenubarAccel() bool
	// HideOnClose returns whether the window will be hidden when the close
	// button is clicked.
	HideOnClose() bool
	// IconName returns the name of the themed icon for the window.
	IconName() string
	// MnemonicsVisible gets whether mnemonics are supposed to be visible.
	MnemonicsVisible() bool
	// Modal returns whether the window is modal.
	Modal() bool
	// Resizable gets the value set by gtk_window_set_resizable().
	Resizable() bool
	// Title retrieves the title of the window.
	Title() string
	// Titlebar returns the custom titlebar that has been set with
	// gtk_window_set_titlebar().
	Titlebar() Widget
	// TransientFor fetches the transient parent for this window.
	TransientFor() Window
	// HasGroup returns whether @window has an explicit window group.
	HasGroup() bool
	// IsActive returns whether the window is part of the current active
	// toplevel.
	//
	// The active toplevel is the window receiving keystrokes.
	//
	// The return value is true if the window is active toplevel itself. You
	// might use this function if you wanted to draw a widget differently in an
	// active window from a widget in an inactive window.
	IsActive() bool
	// IsFullscreen retrieves the current fullscreen state of @window.
	//
	// Note that since fullscreening is ultimately handled by the window manager
	// and happens asynchronously to an application request, you shouldn’t
	// assume the return value of this function changing immediately (or at
	// all), as an effect of calling [method@Gtk.Window.fullscreen] or
	// [method@Gtk.Window.unfullscreen].
	//
	// If the window isn't yet mapped, the value returned will whether the
	// initial requested state is fullscreen.
	IsFullscreen() bool
	// IsMaximized retrieves the current maximized state of @window.
	//
	// Note that since maximization is ultimately handled by the window manager
	// and happens asynchronously to an application request, you shouldn’t
	// assume the return value of this function changing immediately (or at
	// all), as an effect of calling [method@Gtk.Window.maximize] or
	// [method@Gtk.Window.unmaximize].
	//
	// If the window isn't yet mapped, the value returned will whether the
	// initial requested state is maximized.
	IsMaximized() bool
	// Maximize asks to maximize @window, so that it fills the screen.
	//
	// Note that you shouldn’t assume the window is definitely maximized
	// afterward, because other entities (e.g. the user or window manager could
	// unmaximize it again, and not all window managers support maximization.
	//
	// It’s permitted to call this function before showing a window, in which
	// case the window will be maximized when it appears onscreen initially.
	//
	// You can track the result of this operation via the
	// [property@Gdk.Toplevel:state] property, or by listening to notifications
	// on the [property@Gtk.Window:maximized] property.
	Maximize()
	// Minimize asks to minimize the specified @window.
	//
	// Note that you shouldn’t assume the window is definitely minimized
	// afterward, because the windowing system might not support this
	// functionality; other entities (e.g. the user or the window manager could
	// unminimize it again, or there may not be a window manager in which case
	// minimization isn’t possible, etc.
	//
	// It’s permitted to call this function before showing a window, in which
	// case the window will be minimized before it ever appears onscreen.
	//
	// You can track result of this operation via the
	// [property@Gdk.Toplevel:state] property.
	Minimize()
	// Present presents a window to the user.
	//
	// This function should not be used as when it is called, it is too late to
	// gather a valid timestamp to allow focus stealing prevention to work
	// correctly.
	Present()
	// PresentWithTime presents a window to the user.
	//
	// This may mean raising the window in the stacking order, unminimizing it,
	// moving it to the current desktop, and/or giving it the keyboard focus,
	// possibly dependent on the user’s platform, window manager, and
	// preferences.
	//
	// If @window is hidden, this function calls [method@Gtk.Widget.show] as
	// well.
	//
	// This function should be used when the user tries to open a window that’s
	// already open. Say for example the preferences dialog is currently open,
	// and the user chooses Preferences from the menu a second time; use
	// [method@Gtk.Window.present] to move the already-open dialog where the
	// user can see it.
	//
	// Presents a window to the user in response to a user interaction. The
	// timestamp should be gathered when the window was requested to be shown
	// (when clicking a link for example), rather than once the window is ready
	// to be shown.
	PresentWithTime(timestamp uint32)
	// SetApplication sets or unsets the `GtkApplication` associated with the
	// window.
	//
	// The application will be kept alive for at least as long as it has any
	// windows associated with it (see g_application_hold() for a way to keep it
	// alive without windows).
	//
	// Normally, the connection between the application and the window will
	// remain until the window is destroyed, but you can explicitly remove it by
	// setting the @application to nil.
	//
	// This is equivalent to calling [method@Gtk.Application.remove_window]
	// and/or [method@Gtk.Application.add_window] on the old/new applications as
	// relevant.
	SetApplication(application Application)
	// SetChild sets the child widget of @window.
	SetChild(child Widget)
	// SetDecorated sets whether the window should be decorated.
	//
	// By default, windows are decorated with a title bar, resize controls, etc.
	// Some window managers allow GTK to disable these decorations, creating a
	// borderless window. If you set the decorated property to false using this
	// function, GTK will do its best to convince the window manager not to
	// decorate the window. Depending on the system, this function may not have
	// any effect when called on a window that is already visible, so you should
	// call it before calling [method@Gtk.Widget.show].
	//
	// On Windows, this function always works, since there’s no window manager
	// policy involved.
	SetDecorated(setting bool)
	// SetDefaultSize sets the default size of a window.
	//
	// If the window’s “natural” size (its size request) is larger than the
	// default, the default will be ignored.
	//
	// Unlike [method@Gtk.Widget.set_size_request], which sets a size request
	// for a widget and thus would keep users from shrinking the window, this
	// function only sets the initial size, just as if the user had resized the
	// window themselves. Users can still shrink the window again as they
	// normally would. Setting a default size of -1 means to use the “natural”
	// default size (the size request of the window).
	//
	// The default size of a window only affects the first time a window is
	// shown; if a window is hidden and re-shown, it will remember the size it
	// had prior to hiding, rather than using the default size.
	//
	// Windows can’t actually be 0x0 in size, they must be at least 1x1, but
	// passing 0 for @width and @height is OK, resulting in a 1x1 default size.
	//
	// If you use this function to reestablish a previously saved window size,
	// note that the appropriate size to save is the one returned by
	// [method@Gtk.Window.get_default_size]. Using the window allocation
	// directly will not work in all circumstances and can lead to growing or
	// shrinking windows.
	SetDefaultSize(width int, height int)
	// SetDefaultWidget sets the default widget.
	//
	// The default widget is the widget that is activated when the user presses
	// Enter in a dialog (for example).
	SetDefaultWidget(defaultWidget Widget)
	// SetDeletable sets whether the window should be deletable.
	//
	// By default, windows have a close button in the window frame. Some window
	// managers allow GTK to disable this button. If you set the deletable
	// property to false using this function, GTK will do its best to convince
	// the window manager not to show a close button. Depending on the system,
	// this function may not have any effect when called on a window that is
	// already visible, so you should call it before calling
	// [method@Gtk.Widget.show].
	//
	// On Windows, this function always works, since there’s no window manager
	// policy involved.
	SetDeletable(setting bool)
	// SetDestroyWithParent: if @setting is true, then destroying the transient
	// parent of @window will also destroy @window itself.
	//
	// This is useful for dialogs that shouldn’t persist beyond the lifetime of
	// the main window they are associated with, for example.
	SetDestroyWithParent(setting bool)
	// SetDisplay sets the `GdkDisplay` where the @window is displayed.
	//
	// If the window is already mapped, it will be unmapped, and then remapped
	// on the new display.
	SetDisplay(display gdk.Display)
	// SetFocus sets the focus widget.
	//
	// If @focus is not the current focus widget, and is focusable, sets it as
	// the focus widget for the window. If @focus is nil, unsets the focus
	// widget for this window. To set the focus to a particular widget in the
	// toplevel, it is usually more convenient to use
	// [method@Gtk.Widget.grab_focus] instead of this function.
	SetFocus(focus Widget)
	// SetFocusVisible sets whether “focus rectangles” are supposed to be
	// visible.
	SetFocusVisible(setting bool)
	// SetHandleMenubarAccel sets whether this window should react to F10 key
	// presses by activating a menubar it contains.
	SetHandleMenubarAccel(handleMenubarAccel bool)
	// SetHideOnClose: if @setting is true, then clicking the close button on
	// the window will not destroy it, but only hide it.
	SetHideOnClose(setting bool)
	// SetIconName sets the icon for the window from a named themed icon.
	//
	// See the docs for [class@Gtk.IconTheme] for more details. On some
	// platforms, the window icon is not used at all.
	//
	// Note that this has nothing to do with the WM_ICON_NAME property which is
	// mentioned in the ICCCM.
	SetIconName(name string)
	// SetMnemonicsVisible sets whether mnemonics are supposed to be visible.
	SetMnemonicsVisible(setting bool)
	// SetModal sets a window modal or non-modal.
	//
	// Modal windows prevent interaction with other windows in the same
	// application. To keep modal dialogs on top of main application windows,
	// use [method@Gtk.Window.set_transient_for] to make the dialog transient
	// for the parent; most window managers will then disallow lowering the
	// dialog below the parent.
	SetModal(modal bool)
	// SetResizable sets whether the user can resize a window.
	//
	// Windows are user resizable by default.
	SetResizable(resizable bool)
	// SetStartupID sets the startup notification ID.
	//
	// Startup notification identifiers are used by desktop environment to track
	// application startup, to provide user feedback and other features. This
	// function changes the corresponding property on the underlying
	// `GdkSurface`.
	//
	// Normally, startup identifier is managed automatically and you should only
	// use this function in special cases like transferring focus from other
	// processes. You should use this function before calling
	// [method@Gtk.Window.present] or any equivalent function generating a
	// window map event.
	//
	// This function is only useful on X11, not with other GTK targets.
	SetStartupID(startupId string)
	// SetTitle sets the title of the `GtkWindow`.
	//
	// The title of a window will be displayed in its title bar; on the X Window
	// System, the title bar is rendered by the window manager so exactly how
	// the title appears to users may vary according to a user’s exact
	// configuration. The title should help a user distinguish this window from
	// other windows they may have open. A good title might include the
	// application name and current document filename, for example.
	//
	// Passing nil does the same as setting the title to an empty string.
	SetTitle(title string)
	// SetTitlebar sets a custom titlebar for @window.
	//
	// A typical widget used here is [class@Gtk.HeaderBar], as it provides
	// various features expected of a titlebar while allowing the addition of
	// child widgets to it.
	//
	// If you set a custom titlebar, GTK will do its best to convince the window
	// manager not to put its own titlebar on the window. Depending on the
	// system, this function may not work for a window that is already visible,
	// so you set the titlebar before calling [method@Gtk.Widget.show].
	SetTitlebar(titlebar Widget)
	// SetTransientFor: dialog windows should be set transient for the main
	// application window they were spawned from. This allows window managers to
	// e.g. keep the dialog on top of the main window, or center the dialog over
	// the main window. [ctor@Gtk.Dialog.new_with_buttons] and other convenience
	// functions in GTK will sometimes call gtk_window_set_transient_for() on
	// your behalf.
	//
	// Passing nil for @parent unsets the current transient window.
	//
	// On Windows, this function puts the child window on top of the parent,
	// much as the window manager would have done on X.
	SetTransientFor(parent Window)
	// Unfullscreen asks to remove the fullscreen state for @window, and return
	// to its previous state.
	//
	// Note that you shouldn’t assume the window is definitely not fullscreen
	// afterward, because other entities (e.g. the user or window manager could
	// fullscreen it again, and not all window managers honor requests to
	// unfullscreen windows; normally the window will end up restored to its
	// normal state. Just don’t write code that crashes if not.
	//
	// You can track the result of this operation via the
	// [property@Gdk.Toplevel:state] property, or by listening to notifications
	// of the [property@Gtk.Window:fullscreened] property.
	Unfullscreen()
	// Unmaximize asks to unmaximize @window.
	//
	// Note that you shouldn’t assume the window is definitely unmaximized
	// afterward, because other entities (e.g. the user or window manager
	// maximize it again, and not all window managers honor requests to
	// unmaximize.
	//
	// You can track the result of this operation via the
	// [property@Gdk.Toplevel:state] property, or by listening to notifications
	// on the [property@Gtk.Window:maximized] property.
	Unmaximize()
	// Unminimize asks to unminimize the specified @window.
	//
	// Note that you shouldn’t assume the window is definitely unminimized
	// afterward, because the windowing system might not support this
	// functionality; other entities (e.g. the user or the window manager could
	// minimize it again, or there may not be a window manager in which case
	// minimization isn’t possible, etc.
	//
	// You can track result of this operation via the
	// [property@Gdk.Toplevel:state] property.
	Unminimize()
}

// window implements the Window class.
type window struct {
	Widget
}

// WrapWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapWindow(obj *externglib.Object) Window {
	return window{
		Widget: WrapWidget(obj),
	}
}

func marshalWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindow(obj), nil
}

// NewWindow creates a new `GtkWindow`.
//
// To get an undecorated window (no window borders), use
// [method@Gtk.Window.set_decorated].
//
// All top-level windows created by gtk_window_new() are stored in an internal
// top-level window list. This list can be obtained from
// [func@Gtk.Window.list_toplevels]. Due to GTK keeping a reference to the
// window internally, gtk_window_new() does not return a reference to the
// caller.
//
// To delete a `GtkWindow`, call [method@Gtk.Window.destroy].
func NewWindow() Window {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_window_new()

	var _window Window // out

	_window = WrapWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _window
}

func (w window) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(w))
}

func (w window) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(w))
}

func (w window) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(w))
}

func (w window) AsNative() Native {
	return WrapNative(gextras.InternObject(w))
}

func (w window) AsRoot() Root {
	return WrapRoot(gextras.InternObject(w))
}

func (w window) AsShortcutManager() ShortcutManager {
	return WrapShortcutManager(gextras.InternObject(w))
}

func (w window) Close() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_close(_arg0)
}

func (w window) Destroy() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_destroy(_arg0)
}

func (w window) Fullscreen() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_fullscreen(_arg0)
}

func (w window) FullscreenOnMonitor(monitor gdk.Monitor) {
	var _arg0 *C.GtkWindow  // out
	var _arg1 *C.GdkMonitor // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gtk_window_fullscreen_on_monitor(_arg0, _arg1)
}

func (w window) Application() Application {
	var _arg0 *C.GtkWindow      // out
	var _cret *C.GtkApplication // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_application(_arg0)

	var _application Application // out

	_application = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Application)

	return _application
}

func (w window) Child() Widget {
	var _arg0 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (w window) Decorated() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_decorated(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) DefaultSize() (width int, height int) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.int        // in
	var _arg2 C.int        // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_get_default_size(_arg0, &_arg1, &_arg2)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

func (w window) DefaultWidget() Widget {
	var _arg0 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_default_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (w window) Deletable() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_deletable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) DestroyWithParent() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_destroy_with_parent(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Focus() Widget {
	var _arg0 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_focus(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (w window) FocusVisible() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_focus_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Group() WindowGroup {
	var _arg0 *C.GtkWindow      // out
	var _cret *C.GtkWindowGroup // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_group(_arg0)

	var _windowGroup WindowGroup // out

	_windowGroup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(WindowGroup)

	return _windowGroup
}

func (w window) HandleMenubarAccel() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_handle_menubar_accel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) HideOnClose() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_hide_on_close(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) IconName() string {
	var _arg0 *C.GtkWindow // out
	var _cret *C.char      // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (w window) MnemonicsVisible() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_mnemonics_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Modal() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Resizable() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_resizable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Title() string {
	var _arg0 *C.GtkWindow // out
	var _cret *C.char      // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (w window) Titlebar() Widget {
	var _arg0 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_titlebar(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (w window) TransientFor() Window {
	var _arg0 *C.GtkWindow // out
	var _cret *C.GtkWindow // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_get_transient_for(_arg0)

	var _ret Window // out

	_ret = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _ret
}

func (w window) HasGroup() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_has_group(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) IsActive() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_is_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) IsFullscreen() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_is_fullscreen(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) IsMaximized() bool {
	var _arg0 *C.GtkWindow // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	_cret = C.gtk_window_is_maximized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (w window) Maximize() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_maximize(_arg0)
}

func (w window) Minimize() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_minimize(_arg0)
}

func (w window) Present() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_present(_arg0)
}

func (w window) PresentWithTime(timestamp uint32) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.guint32    // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = C.guint32(timestamp)

	C.gtk_window_present_with_time(_arg0, _arg1)
}

func (w window) SetApplication(application Application) {
	var _arg0 *C.GtkWindow      // out
	var _arg1 *C.GtkApplication // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	C.gtk_window_set_application(_arg0, _arg1)
}

func (w window) SetChild(child Widget) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_window_set_child(_arg0, _arg1)
}

func (w window) SetDecorated(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_decorated(_arg0, _arg1)
}

func (w window) SetDefaultSize(width int, height int) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	C.gtk_window_set_default_size(_arg0, _arg1, _arg2)
}

func (w window) SetDefaultWidget(defaultWidget Widget) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(defaultWidget.Native()))

	C.gtk_window_set_default_widget(_arg0, _arg1)
}

func (w window) SetDeletable(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_deletable(_arg0, _arg1)
}

func (w window) SetDestroyWithParent(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_destroy_with_parent(_arg0, _arg1)
}

func (w window) SetDisplay(display gdk.Display) {
	var _arg0 *C.GtkWindow  // out
	var _arg1 *C.GdkDisplay // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_window_set_display(_arg0, _arg1)
}

func (w window) SetFocus(focus Widget) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(focus.Native()))

	C.gtk_window_set_focus(_arg0, _arg1)
}

func (w window) SetFocusVisible(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_focus_visible(_arg0, _arg1)
}

func (w window) SetHandleMenubarAccel(handleMenubarAccel bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if handleMenubarAccel {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_handle_menubar_accel(_arg0, _arg1)
}

func (w window) SetHideOnClose(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_hide_on_close(_arg0, _arg1)
}

func (w window) SetIconName(name string) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_window_set_icon_name(_arg0, _arg1)
}

func (w window) SetMnemonicsVisible(setting bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_mnemonics_visible(_arg0, _arg1)
}

func (w window) SetModal(modal bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_modal(_arg0, _arg1)
}

func (w window) SetResizable(resizable bool) {
	var _arg0 *C.GtkWindow // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gtk_window_set_resizable(_arg0, _arg1)
}

func (w window) SetStartupID(startupId string) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_window_set_startup_id(_arg0, _arg1)
}

func (w window) SetTitle(title string) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_window_set_title(_arg0, _arg1)
}

func (w window) SetTitlebar(titlebar Widget) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(titlebar.Native()))

	C.gtk_window_set_titlebar(_arg0, _arg1)
}

func (w window) SetTransientFor(parent Window) {
	var _arg0 *C.GtkWindow // out
	var _arg1 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_window_set_transient_for(_arg0, _arg1)
}

func (w window) Unfullscreen() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_unfullscreen(_arg0)
}

func (w window) Unmaximize() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_unmaximize(_arg0)
}

func (w window) Unminimize() {
	var _arg0 *C.GtkWindow // out

	_arg0 = (*C.GtkWindow)(unsafe.Pointer(w.Native()))

	C.gtk_window_unminimize(_arg0)
}
