// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
	})
}

// X11RegisterStandardEventType registers interest in receiving extension events
// with type codes between @event_base and `event_base + n_events - 1`. The
// registered events must have the window field in the same place as core X
// events (this is not the case for e.g. XKB extension events).
//
// If an event type is registered, events of this type will go through global
// and window-specific filters (see gdk_window_add_filter()). Unregistered
// events will only go through global filters. GDK may register the events of
// some X extensions on its own.
//
// This function should only be needed in unusual circumstances, e.g. when
// filtering XInput extension events on the root window.
func X11RegisterStandardEventType(display X11Display, eventBase int, nEvents int) {
	var arg1 *C.GdkDisplay
	var arg2 C.gint
	var arg3 C.gint

	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	arg2 = C.gint(eventBase)
	arg3 = C.gint(nEvents)

	C.gdk_x11_register_standard_event_type(display, eventBase, nEvents)
}

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientID string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(smClientID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_set_sm_client_id(smClientID)
}

type X11Display interface {
	gdk.Display

	// ErrorTrapPop pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	//
	// See gdk_error_trap_pop() for the all-displays-at-once equivalent.
	ErrorTrapPop() int
	// ErrorTrapPopIgnored pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	//
	// See gdk_error_trap_pop_ignored() for the all-displays-at-once equivalent.
	ErrorTrapPopIgnored()
	// ErrorTrapPush begins a range of X requests on @display for which X error
	// events will be ignored. Unignored errors (when no trap is pushed) will
	// abort the application. Use gdk_x11_display_error_trap_pop() or
	// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
	// function.
	//
	// See also gdk_error_trap_push() to push a trap on all displays.
	ErrorTrapPush()
	// StartupNotificationID gets the startup notification ID for a display.
	StartupNotificationID() string
	// UserTime returns the timestamp of the last user interaction on @display.
	// The timestamp is taken from events caused by user interaction such as key
	// presses or pointer movements. See gdk_x11_window_set_user_time().
	UserTime() uint32
	// Grab: call XGrabServer() on @display. To ungrab the display again, use
	// gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	Grab()
	// SetCursorTheme sets the cursor theme from which the images for cursor
	// should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new(), gdk_cursor_new_for_display() and
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_pixbuf() will have to
	// be handled by the application (GTK+ applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorTheme(theme string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// If the ID contains the string "_TIME" then the portion following that
	// string is taken to be the X11 timestamp of the event that triggered the
	// application to be launched and the GDK current event time is set
	// accordingly.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_notify_startup_complete()).
	SetStartupNotificationID(startupID string)
	// SetWindowScale forces a specific window scale for all windows on this
	// display, instead of using the default or user configured scale. This is
	// can be used to disable scaling support by setting @scale to 1, or to
	// programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetWindowScale(scale int)
	// StringToCompoundText: convert a string from the encoding of the current
	// locale into a form suitable for storing in a window property.
	StringToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, length int, gint int)
	// TextPropertyToTextList: convert a text string from the encoding as it is
	// stored in a property into an array of strings in the encoding of the
	// current locale. (The elements of the array represent the nul-separated
	// elements of the original text string.)
	TextPropertyToTextList(encoding gdk.Atom, format int, text byte, length int, list string) int
	// Ungrab: ungrab @display after it has been grabbed with
	// gdk_x11_display_grab().
	Ungrab()
	// UTF8ToCompoundText converts from UTF-8 to compound text.
	UTF8ToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, length int, ok bool)
}

// x11Display implements the X11Display interface.
type x11Display struct {
	gdk.Display
}

var _ X11Display = (*x11Display)(nil)

// WrapX11Display wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return X11Display{
		gdk.Display: gdk.WrapDisplay(obj),
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

// ErrorTrapPop pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
// always block until the error is known to have occurred or not occurred,
// so the error code can be returned.
//
// If you don’t need to use the return value,
// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
//
// See gdk_error_trap_pop() for the all-displays-at-once equivalent.
func (d x11Display) ErrorTrapPop() int {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gdk_x11_display_error_trap_pop(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// ErrorTrapPopIgnored pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Does not block to see if an error
// occurred; merely records the range of requests to ignore errors for, and
// ignores those errors if they arrive asynchronously.
//
// See gdk_error_trap_pop_ignored() for the all-displays-at-once equivalent.
func (d x11Display) ErrorTrapPopIgnored() {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(arg0)
}

// ErrorTrapPush begins a range of X requests on @display for which X error
// events will be ignored. Unignored errors (when no trap is pushed) will
// abort the application. Use gdk_x11_display_error_trap_pop() or
// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
// function.
//
// See also gdk_error_trap_push() to push a trap on all displays.
func (d x11Display) ErrorTrapPush() {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_push(arg0)
}

// StartupNotificationID gets the startup notification ID for a display.
func (d x11Display) StartupNotificationID() string {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.gdk_x11_display_get_startup_notification_id(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// UserTime returns the timestamp of the last user interaction on @display.
// The timestamp is taken from events caused by user interaction such as key
// presses or pointer movements. See gdk_x11_window_set_user_time().
func (d x11Display) UserTime() uint32 {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var cret C.guint32
	var goret1 uint32

	cret = C.gdk_x11_display_get_user_time(arg0)

	goret1 = C.guint32(cret)

	return goret1
}

// Grab: call XGrabServer() on @display. To ungrab the display again, use
// gdk_x11_display_ungrab().
//
// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
func (d x11Display) Grab() {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(arg0)
}

// SetCursorTheme sets the cursor theme from which the images for cursor
// should be taken.
//
// If the windowing system supports it, existing cursors created with
// gdk_cursor_new(), gdk_cursor_new_for_display() and
// gdk_cursor_new_from_name() are updated to reflect the theme change.
// Custom cursors constructed with gdk_cursor_new_from_pixbuf() will have to
// be handled by the application (GTK+ applications can learn about cursor
// theme changes by listening for change notification for the corresponding
// Setting).
func (d x11Display) SetCursorTheme(theme string, size int) {
	var arg0 *C.GdkDisplay
	var arg1 *C.gchar
	var arg2 C.gint

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.gchar)(C.CString(theme))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(size)

	C.gdk_x11_display_set_cursor_theme(arg0, theme, size)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID
// environment variable, but in some cases (such as the application not
// being launched using exec()) it can come from other sources.
//
// If the ID contains the string "_TIME" then the portion following that
// string is taken to be the X11 timestamp of the event that triggered the
// application to be launched and the GDK current event time is set
// accordingly.
//
// The startup ID is also what is used to signal that the startup is
// complete (for example, when opening a window or when calling
// gdk_notify_startup_complete()).
func (d x11Display) SetStartupNotificationID(startupID string) {
	var arg0 *C.GdkDisplay
	var arg1 *C.gchar

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.gchar)(C.CString(startupID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_display_set_startup_notification_id(arg0, startupID)
}

// SetWindowScale forces a specific window scale for all windows on this
// display, instead of using the default or user configured scale. This is
// can be used to disable scaling support by setting @scale to 1, or to
// programmatically set the window scale.
//
// Once the scale is set by this call it will not change in response to
// later user configuration changes.
func (d x11Display) SetWindowScale(scale int) {
	var arg0 *C.GdkDisplay
	var arg1 C.gint

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = C.gint(scale)

	C.gdk_x11_display_set_window_scale(arg0, scale)
}

// StringToCompoundText: convert a string from the encoding of the current
// locale into a form suitable for storing in a window property.
func (d x11Display) StringToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, length int, gint int) {
	var arg0 *C.GdkDisplay
	var arg1 *C.gchar

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gint
	var goret5 int

	cret = C.gdk_x11_display_string_to_compound_text(arg0, str, &arg2, &arg3, &arg4, &arg5)

	goret5 = C.gint(cret)

	return ret2, ret3, ret4, ret5, goret5
}

// TextPropertyToTextList: convert a text string from the encoding as it is
// stored in a property into an array of strings in the encoding of the
// current locale. (The elements of the array represent the nul-separated
// elements of the original text string.)
func (d x11Display) TextPropertyToTextList(encoding gdk.Atom, format int, text byte, length int, list string) int {
	var arg0 *C.GdkDisplay
	var arg1 C.GdkAtom
	var arg2 C.gint
	var arg3 *C.guchar
	var arg4 C.gint
	var arg5 ***C.gchar

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (C.GdkAtom)(unsafe.Pointer(encoding.Native()))
	arg2 = C.gint(format)
	arg3 = *C.guchar(text)
	arg4 = C.gint(length)
	arg5 = (***C.gchar)(C.CString(list))
	defer C.free(unsafe.Pointer(arg5))

	var cret C.gint
	var goret1 int

	cret = C.gdk_x11_display_text_property_to_text_list(arg0, encoding, format, text, length, list)

	goret1 = C.gint(cret)

	return goret1
}

// Ungrab: ungrab @display after it has been grabbed with
// gdk_x11_display_grab().
func (d x11Display) Ungrab() {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(arg0)
}

// UTF8ToCompoundText converts from UTF-8 to compound text.
func (d x11Display) UTF8ToCompoundText(str string) (encoding gdk.Atom, format int, ctext []byte, length int, ok bool) {
	var arg0 *C.GdkDisplay
	var arg1 *C.gchar

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret5 bool

	cret = C.gdk_x11_display_utf8_to_compound_text(arg0, str, &arg2, &arg3, &arg4, &arg5)

	goret5 = C.bool(cret) != C.false

	return ret2, ret3, ret4, ret5, goret5
}
