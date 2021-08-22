// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplayer},
	})
}

// Display objects purpose are two fold:
//
// - To manage and provide information about input devices (pointers and
// keyboards)
//
// - To manage and provide information about the available Screens
//
// GdkDisplay objects are the GDK representation of an X Display, which can be
// described as a workstation consisting of a keyboard, a pointing device (such
// as a mouse) and one or more screens. It is used to open and keep track of
// various GdkScreen objects currently instantiated by the application. It is
// also used to access the keyboard(s) and mouse pointer(s) of the display.
//
// Most of the input device handling has been factored out into the separate
// DeviceManager object. Every display has a device manager, which you can
// obtain using gdk_display_get_device_manager().
type Display struct {
	*externglib.Object
}

func WrapDisplay(obj *externglib.Object) *Display {
	return &Display{
		Object: obj,
	}
}

func marshalDisplayer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDisplay(obj), nil
}

// Beep emits a short beep on display
func (display *Display) Beep() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_beep(_arg0)
	runtime.KeepAlive(display)
}

// Close closes the connection to the windowing system for the given display,
// and cleans up associated resources.
func (display *Display) Close() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_close(_arg0)
	runtime.KeepAlive(display)
}

// DeviceIsGrabbed returns TRUE if there is an ongoing grab on device for
// display.
func (display *Display) DeviceIsGrabbed(device Devicer) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(device)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Flush flushes any requests queued for the windowing system; this happens
// automatically when the main loop blocks waiting for new events, but if your
// application is drawing without returning control to the main loop, you may
// need to call this function explicitly. A common case where this function
// needs to be called is when an application is executing drawing commands from
// a thread other than the thread where the main loop is running.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *Display) Flush() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_flush(_arg0)
	runtime.KeepAlive(display)
}

// AppLaunchContext returns a AppLaunchContext suitable for launching
// applications on the given display.
func (display *Display) AppLaunchContext() *AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)
	runtime.KeepAlive(display)

	var _appLaunchContext *AppLaunchContext // out

	_appLaunchContext = WrapAppLaunchContext(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _appLaunchContext
}

// DefaultCursorSize returns the default size to use for cursors on display.
func (display *Display) DefaultCursorSize() uint {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_cursor_size(_arg0)
	runtime.KeepAlive(display)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DefaultGroup returns the default group leader window for all toplevel windows
// on display. This window is implicitly created by GDK. See
// gdk_window_set_group().
func (display *Display) DefaultGroup() Windower {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_group(_arg0)
	runtime.KeepAlive(display)

	var _window Windower // out

	_window = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Windower)

	return _window
}

// DefaultScreen: get the default Screen for display.
func (display *Display) DefaultScreen() *Screen {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_screen(_arg0)
	runtime.KeepAlive(display)

	var _screen *Screen // out

	_screen = WrapScreen(externglib.Take(unsafe.Pointer(_cret)))

	return _screen
}

// DefaultSeat returns the default Seat for this display.
func (display *Display) DefaultSeat() Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)
	runtime.KeepAlive(display)

	var _seat Seater // out

	_seat = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Seater)

	return _seat
}

// DeviceManager returns the DeviceManager associated to display.
//
// Deprecated: Use gdk_display_get_default_seat() and Seat operations.
func (display *Display) DeviceManager() DeviceManagerer {
	var _arg0 *C.GdkDisplay       // out
	var _cret *C.GdkDeviceManager // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_device_manager(_arg0)
	runtime.KeepAlive(display)

	var _deviceManager DeviceManagerer // out

	if _cret != nil {
		_deviceManager = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(DeviceManagerer)
	}

	return _deviceManager
}

// MaximalCursorSize gets the maximal size to use for cursors on display.
func (display *Display) MaximalCursorSize() (width uint, height uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // in
	var _arg2 C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_get_maximal_cursor_size(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(display)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

// Monitor gets a monitor associated with this display.
func (display *Display) Monitor(monitorNum int) *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.int(monitorNum)

	_cret = C.gdk_display_get_monitor(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(monitorNum)

	var _monitor *Monitor // out

	if _cret != nil {
		_monitor = WrapMonitor(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _monitor
}

// MonitorAtPoint gets the monitor in which the point (x, y) is located, or a
// nearby monitor if the point is not in any monitor.
func (display *Display) MonitorAtPoint(x int, y int) *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gdk_display_get_monitor_at_point(_arg0, _arg1, _arg2)
	runtime.KeepAlive(display)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _monitor *Monitor // out

	_monitor = WrapMonitor(externglib.Take(unsafe.Pointer(_cret)))

	return _monitor
}

// MonitorAtWindow gets the monitor in which the largest area of window resides,
// or a monitor close to window if it is outside of all monitors.
func (display *Display) MonitorAtWindow(window Windower) *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkWindow  // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gdk_display_get_monitor_at_window(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(window)

	var _monitor *Monitor // out

	_monitor = WrapMonitor(externglib.Take(unsafe.Pointer(_cret)))

	return _monitor
}

// NMonitors gets the number of monitors that belong to display.
//
// The returned number is valid until the next emission of the
// Display::monitor-added or Display::monitor-removed signal.
func (display *Display) NMonitors() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_n_monitors(_arg0)
	runtime.KeepAlive(display)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NScreens gets the number of screen managed by the display.
//
// Deprecated: The number of screens is always 1.
func (display *Display) NScreens() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_n_screens(_arg0)
	runtime.KeepAlive(display)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name gets the name of the display.
func (display *Display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_name(_arg0)
	runtime.KeepAlive(display)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Pointer gets the current location of the pointer and the current modifier
// mask for a given display.
//
// Deprecated: Use gdk_device_get_position() instead.
func (display *Display) Pointer() (screen *Screen, x int, y int, mask ModifierType) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 *C.GdkScreen      // in
	var _arg2 C.gint            // in
	var _arg3 C.gint            // in
	var _arg4 C.GdkModifierType // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_get_pointer(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(display)

	var _screen *Screen    // out
	var _x int             // out
	var _y int             // out
	var _mask ModifierType // out

	if _arg1 != nil {
		_screen = WrapScreen(externglib.Take(unsafe.Pointer(_arg1)))
	}
	_x = int(_arg2)
	_y = int(_arg3)
	_mask = ModifierType(_arg4)

	return _screen, _x, _y, _mask
}

// PrimaryMonitor gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop” lives.
// While normal application windows typically allow the window manager to place
// the windows, specialized desktop applications such as panels should place
// themselves on the primary monitor.
func (display *Display) PrimaryMonitor() *Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_primary_monitor(_arg0)
	runtime.KeepAlive(display)

	var _monitor *Monitor // out

	if _cret != nil {
		_monitor = WrapMonitor(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _monitor
}

// Screen returns a screen object for one of the screens of the display.
//
// Deprecated: There is only one screen; use gdk_display_get_default_screen() to
// get it.
func (display *Display) Screen(screenNum int) *Screen {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.gint(screenNum)

	_cret = C.gdk_display_get_screen(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(screenNum)

	var _screen *Screen // out

	_screen = WrapScreen(externglib.Take(unsafe.Pointer(_cret)))

	return _screen
}

// WindowAtPointer obtains the window underneath the mouse pointer, returning
// the location of the pointer in that window in win_x, win_y for screen.
// Returns NULL if the window under the mouse pointer is not known to GDK (for
// example, belongs to another application).
//
// Deprecated: Use gdk_device_get_window_at_position() instead.
func (display *Display) WindowAtPointer() (winX int, winY int, window Windower) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // in
	var _arg2 C.gint        // in
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_get_window_at_pointer(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(display)

	var _winX int        // out
	var _winY int        // out
	var _window Windower // out

	_winX = int(_arg1)
	_winY = int(_arg2)
	if _cret != nil {
		_window = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Windower)
	}

	return _winX, _winY, _window
}

// HasPending returns whether the display has events that are waiting to be
// processed.
func (display *Display) HasPending() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_has_pending(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsClosed finds out if the display has been closed.
func (display *Display) IsClosed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_is_closed(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// KeyboardUngrab: release any keyboard grab
//
// Deprecated: Use gdk_device_ungrab(), together with gdk_device_grab() instead.
func (display *Display) KeyboardUngrab(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_display_keyboard_ungrab(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(time_)
}

// ListDevices returns the list of available input devices attached to display.
// The list is statically allocated and should not be freed.
//
// Deprecated: Use gdk_device_manager_list_devices() instead.
func (display *Display) ListDevices() []Devicer {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GList      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_list_devices(_arg0)
	runtime.KeepAlive(display)

	var _list []Devicer // out

	_list = make([]Devicer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GdkDevice)(v)
		var dst Devicer // out
		dst = (externglib.CastObject(externglib.Take(unsafe.Pointer(src)))).(Devicer)
		_list = append(_list, dst)
	})

	return _list
}

// ListSeats returns the list of seats known to display.
func (display *Display) ListSeats() []Seater {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GList      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_list_seats(_arg0)
	runtime.KeepAlive(display)

	var _list []Seater // out

	_list = make([]Seater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkSeat)(v)
		var dst Seater // out
		dst = (externglib.CastObject(externglib.Take(unsafe.Pointer(src)))).(Seater)
		_list = append(_list, dst)
	})

	return _list
}

// NotifyStartupComplete indicates to the GUI environment that the application
// has finished loading, using a given identifier.
//
// GTK+ will call this function automatically for Window with custom
// startup-notification identifier unless
// gtk_window_set_auto_startup_notification() is called to disable that feature.
func (display *Display) NotifyStartupComplete(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(startupId)
}

// PointerIsGrabbed: test if the pointer is grabbed.
//
// Deprecated: Use gdk_display_device_is_grabbed() instead.
func (display *Display) PointerIsGrabbed() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_pointer_is_grabbed(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PointerUngrab: release any pointer grab.
//
// Deprecated: Use gdk_device_ungrab(), together with gdk_device_grab() instead.
func (display *Display) PointerUngrab(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_display_pointer_ungrab(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(time_)
}

// SetDoubleClickDistance sets the double click distance (two clicks within this
// distance count as a double click and result in a K_2BUTTON_PRESS event). See
// also gdk_display_set_double_click_time(). Applications should not set this,
// it is a global user-configured setting.
func (display *Display) SetDoubleClickDistance(distance uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(distance)

	C.gdk_display_set_double_click_distance(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(distance)
}

// SetDoubleClickTime sets the double click time (two clicks within this time
// interval count as a double click and result in a K_2BUTTON_PRESS event).
// Applications should not set this, it is a global user-configured setting.
func (display *Display) SetDoubleClickTime(msec uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = C.guint(msec)

	C.gdk_display_set_double_click_time(_arg0, _arg1)
	runtime.KeepAlive(display)
	runtime.KeepAlive(msec)
}

// SupportsClipboardPersistence returns whether the speicifed display supports
// clipboard persistance; i.e. if it’s possible to store the clipboard data
// after an application has quit. On X11 this checks if a clipboard daemon is
// running.
func (display *Display) SupportsClipboardPersistence() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_clipboard_persistence(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsComposite returns TRUE if gdk_window_set_composited() can be used to
// redirect drawing on the window using compositing.
//
// Currently this only works on X11 with XComposite and XDamage extensions
// available.
//
// Deprecated: Compositing is an outdated technology that only ever worked on
// X11.
func (display *Display) SupportsComposite() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_composite(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsCursorAlpha returns TRUE if cursors can use an 8bit alpha channel on
// display. Otherwise, cursors are restricted to bilevel alpha (i.e. a mask).
func (display *Display) SupportsCursorAlpha() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_cursor_alpha(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsCursorColor returns TRUE if multicolored cursors are supported on
// display. Otherwise, cursors have only a forground and a background color.
func (display *Display) SupportsCursorColor() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_cursor_color(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsInputShapes returns TRUE if gdk_window_input_shape_combine_mask() can
// be used to modify the input shape of windows on display.
func (display *Display) SupportsInputShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsSelectionNotification returns whether EventOwnerChange events will be
// sent when the owner of a selection changes.
func (display *Display) SupportsSelectionNotification() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_selection_notification(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SupportsShapes returns TRUE if gdk_window_shape_combine_mask() can be used to
// create shaped windows on display.
func (display *Display) SupportsShapes() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_display_supports_shapes(_arg0)
	runtime.KeepAlive(display)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sync flushes any requests queued for the windowing system and waits until all
// requests have been handled. This is often used for making sure that the
// display is synchronized with the current state of the program. Calling
// gdk_display_sync() before gdk_error_trap_pop() makes sure that any errors
// generated from earlier requests are handled before the error trap is removed.
//
// This is most useful for X11. On windowing systems where requests are handled
// synchronously, this function will do nothing.
func (display *Display) Sync() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gdk_display_sync(_arg0)
	runtime.KeepAlive(display)
}

// WarpPointer warps the pointer of display to the point x,y on the screen
// screen, unless the pointer is confined to a window by a grab, in which case
// it will be moved as far as allowed by the grab. Warping the pointer creates
// events as if the user had moved the mouse instantaneously to the destination.
//
// Note that the pointer should normally be under the control of the user. This
// function was added to cover some rare use cases like keyboard navigation
// support for the color picker in the ColorSelectionDialog.
//
// Deprecated: Use gdk_device_warp() instead.
func (display *Display) WarpPointer(screen *Screen, x int, y int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkScreen  // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_display_warp_pointer(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(display)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// DisplayGetDefault gets the default Display. This is a convenience function
// for: gdk_display_manager_get_default_display (gdk_display_manager_get ()).
func DisplayGetDefault() *Display {
	var _cret *C.GdkDisplay // in

	_cret = C.gdk_display_get_default()

	var _display *Display // out

	if _cret != nil {
		_display = WrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// DisplayOpen opens a display.
func DisplayOpen(displayName string) *Display {
	var _arg1 *C.gchar      // out
	var _cret *C.GdkDisplay // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(displayName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_display_open(_arg1)
	runtime.KeepAlive(displayName)

	var _display *Display // out

	if _cret != nil {
		_display = WrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// DisplayOpenDefaultLibgtkOnly opens the default display specified by command
// line arguments or environment variables, sets it as the default display, and
// returns it. gdk_parse_args() must have been called first. If the default
// display has previously been set, simply returns that. An internal function
// that should not be used by applications.
//
// Deprecated: This symbol was never meant to be used outside of GTK+.
func DisplayOpenDefaultLibgtkOnly() *Display {
	var _cret *C.GdkDisplay // in

	_cret = C.gdk_display_open_default_libgtk_only()

	var _display *Display // out

	if _cret != nil {
		_display = WrapDisplay(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}
