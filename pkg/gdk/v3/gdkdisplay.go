// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplay},
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
type Display interface {
	gextras.Objector

	BeepDisplay()

	CloseDisplay()

	DeviceIsGrabbedDisplay(device Device) bool

	FlushDisplay()

	AppLaunchContext() AppLaunchContext

	DefaultCursorSize() uint

	DefaultGroup() Window

	DefaultScreen() Screen

	DefaultSeat() Seat

	DeviceManager() DeviceManager

	MaximalCursorSize() (width uint, height uint)

	Monitor(monitorNum int) Monitor

	MonitorAtPoint(x int, y int) Monitor

	MonitorAtWindow(window Window) Monitor

	NMonitors() int

	NScreens() int

	Name() string

	Pointer() (screen Screen, x int, y int, mask ModifierType)

	PrimaryMonitor() Monitor

	Screen(screenNum int) Screen

	WindowAtPointer() (winX int, winY int, window Window)

	HasPendingDisplay() bool

	IsClosedDisplay() bool

	KeyboardUngrabDisplay(time_ uint32)

	NotifyStartupCompleteDisplay(startupId string)

	PointerIsGrabbedDisplay() bool

	PointerUngrabDisplay(time_ uint32)

	RequestSelectionNotificationDisplay(selection *Atom) bool

	SetDoubleClickDistanceDisplay(distance uint)

	SetDoubleClickTimeDisplay(msec uint)

	StoreClipboardDisplay(clipboardWindow Window, time_ uint32, targets []*Atom)

	SupportsClipboardPersistenceDisplay() bool

	SupportsCompositeDisplay() bool

	SupportsCursorAlphaDisplay() bool

	SupportsCursorColorDisplay() bool

	SupportsInputShapesDisplay() bool

	SupportsSelectionNotificationDisplay() bool

	SupportsShapesDisplay() bool

	SyncDisplay()

	WarpPointerDisplay(screen Screen, x int, y int)
}

// display implements the Display class.
type display struct {
	gextras.Objector
}

// WrapDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapDisplay(obj *externglib.Object) Display {
	return display{
		Objector: obj,
	}
}

func marshalDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDisplay(obj), nil
}

func (d display) BeepDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_beep(_arg0)
}

func (d display) CloseDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_close(_arg0)
}

func (d display) DeviceIsGrabbedDisplay(device Device) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) FlushDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_flush(_arg0)
}

func (d display) AppLaunchContext() AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)

	var _appLaunchContext AppLaunchContext // out

	_appLaunchContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(AppLaunchContext)

	return _appLaunchContext
}

func (d display) DefaultCursorSize() uint {
	var _arg0 *C.GdkDisplay // out
	var _cret C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_cursor_size(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (d display) DefaultGroup() Window {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_group(_arg0)

	var _window Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _window
}

func (d display) DefaultScreen() Screen {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_screen(_arg0)

	var _screen Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Screen)

	return _screen
}

func (d display) DefaultSeat() Seat {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Seat)

	return _seat
}

func (d display) DeviceManager() DeviceManager {
	var _arg0 *C.GdkDisplay       // out
	var _cret *C.GdkDeviceManager // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_device_manager(_arg0)

	var _deviceManager DeviceManager // out

	_deviceManager = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DeviceManager)

	return _deviceManager
}

func (d display) MaximalCursorSize() (width uint, height uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // in
	var _arg2 C.guint       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_get_maximal_cursor_size(_arg0, &_arg1, &_arg2)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

func (d display) Monitor(monitorNum int) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(monitorNum)

	_cret = C.gdk_display_get_monitor(_arg0, _arg1)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

func (d display) MonitorAtPoint(x int, y int) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gdk_display_get_monitor_at_point(_arg0, _arg1, _arg2)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

func (d display) MonitorAtWindow(window Window) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkWindow  // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	_cret = C.gdk_display_get_monitor_at_window(_arg0, _arg1)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

func (d display) NMonitors() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.int         // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_n_monitors(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d display) NScreens() int {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gint        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_n_screens(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d display) Pointer() (screen Screen, x int, y int, mask ModifierType) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 *C.GdkScreen      // in
	var _arg2 C.gint            // in
	var _arg3 C.gint            // in
	var _arg4 C.GdkModifierType // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_get_pointer(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _screen Screen     // out
	var _x int             // out
	var _y int             // out
	var _mask ModifierType // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg1))).(Screen)
	_x = int(_arg2)
	_y = int(_arg3)
	_mask = ModifierType(_arg4)

	return _screen, _x, _y, _mask
}

func (d display) PrimaryMonitor() Monitor {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_primary_monitor(_arg0)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

func (d display) Screen(screenNum int) Screen {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // out
	var _cret *C.GdkScreen  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.gint(screenNum)

	_cret = C.gdk_display_get_screen(_arg0, _arg1)

	var _screen Screen // out

	_screen = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Screen)

	return _screen
}

func (d display) WindowAtPointer() (winX int, winY int, window Window) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.gint        // in
	var _arg2 C.gint        // in
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_window_at_pointer(_arg0, &_arg1, &_arg2)

	var _winX int      // out
	var _winY int      // out
	var _window Window // out

	_winX = int(_arg1)
	_winY = int(_arg2)
	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Window)

	return _winX, _winY, _window
}

func (d display) HasPendingDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_has_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) IsClosedDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) KeyboardUngrabDisplay(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_display_keyboard_ungrab(_arg0, _arg1)
}

func (d display) NotifyStartupCompleteDisplay(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
}

func (d display) PointerIsGrabbedDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_pointer_is_grabbed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) PointerUngrabDisplay(time_ uint32) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint32     // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint32(time_)

	C.gdk_display_pointer_ungrab(_arg0, _arg1)
}

func (d display) RequestSelectionNotificationDisplay(selection *Atom) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.GdkAtom     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	{
		var refTmpIn *Atom
		var refTmpOut *C.GdkAtom

		refTmpIn = selection

		refTmpOut = (*C.GdkAtom)(unsafe.Pointer(refTmpIn.Native()))

		_arg1 = *refTmpOut
	}

	_cret = C.gdk_display_request_selection_notification(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SetDoubleClickDistanceDisplay(distance uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(distance)

	C.gdk_display_set_double_click_distance(_arg0, _arg1)
}

func (d display) SetDoubleClickTimeDisplay(msec uint) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(msec)

	C.gdk_display_set_double_click_time(_arg0, _arg1)
}

func (d display) StoreClipboardDisplay(clipboardWindow Window, time_ uint32, targets []*Atom) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkWindow  // out
	var _arg2 C.guint32     // out
	var _arg3 *C.GdkAtom
	var _arg4 C.gint

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(clipboardWindow.Native()))
	_arg2 = C.guint32(time_)
	_arg4 = C.gint(len(targets))
	_arg3 = (*C.GdkAtom)(unsafe.Pointer(&targets[0]))

	C.gdk_display_store_clipboard(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (d display) SupportsClipboardPersistenceDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_clipboard_persistence(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsCompositeDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_composite(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsCursorAlphaDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_cursor_alpha(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsCursorColorDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_cursor_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsInputShapesDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsSelectionNotificationDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_selection_notification(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SupportsShapesDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SyncDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_sync(_arg0)
}

func (d display) WarpPointerDisplay(screen Screen, x int, y int) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkScreen  // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gdk_display_warp_pointer(_arg0, _arg1, _arg2, _arg3)
}