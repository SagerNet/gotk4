// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk/gdk.h>
import "C"

// TestRenderSync retrieves a pixel from @window to force the windowing system
// to carry out any pending rendering commands.
//
// This function is intended to be used to synchronize with rendering pipelines,
// to benchmark windowing system rendering operations.
func TestRenderSync(window Window) {
	var arg1 *C.GdkWindow

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gdk_test_render_sync(window)
}

// TestSimulateButton: this function is intended to be used in GTK+ test
// programs. It will warp the mouse pointer to the given (@x,@y) coordinates
// within @window and simulate a button press or release event. Because the
// mouse pointer needs to be warped to the target location, use of this function
// outside of test programs that run in their own virtual windowing system (e.g.
// Xvfb) is not recommended.
//
// Also, gdk_test_simulate_button() is a fairly low level function, for most
// testing purposes, gtk_test_widget_click() is the right function to call which
// will generate a button press event followed by its accompanying button
// release event.
func TestSimulateButton(window Window, x int, y int, button uint, modifiers ModifierType, buttonPressrelease EventType) bool {
	var arg1 *C.GdkWindow
	var arg2 C.gint
	var arg3 C.gint
	var arg4 C.guint
	var arg5 C.GdkModifierType
	var arg6 C.GdkEventType

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)
	arg4 = C.guint(button)
	arg5 = (C.GdkModifierType)(modifiers)
	arg6 = (C.GdkEventType)(buttonPressrelease)

	var cret C.gboolean
	var goret1 bool

	cret = C.gdk_test_simulate_button(window, x, y, button, modifiers, buttonPressrelease)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// TestSimulateKey: this function is intended to be used in GTK+ test programs.
// If (@x,@y) are > (-1,-1), it will warp the mouse pointer to the given (@x,@y)
// coordinates within @window and simulate a key press or release event.
//
// When the mouse pointer is warped to the target location, use of this function
// outside of test programs that run in their own virtual windowing system (e.g.
// Xvfb) is not recommended. If (@x,@y) are passed as (-1,-1), the mouse pointer
// will not be warped and @window origin will be used as mouse pointer location
// for the event.
//
// Also, gdk_test_simulate_key() is a fairly low level function, for most
// testing purposes, gtk_test_widget_send_key() is the right function to call
// which will generate a key press event followed by its accompanying key
// release event.
func TestSimulateKey(window Window, x int, y int, keyval uint, modifiers ModifierType, keyPressrelease EventType) bool {
	var arg1 *C.GdkWindow
	var arg2 C.gint
	var arg3 C.gint
	var arg4 C.guint
	var arg5 C.GdkModifierType
	var arg6 C.GdkEventType

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)
	arg4 = C.guint(keyval)
	arg5 = (C.GdkModifierType)(modifiers)
	arg6 = (C.GdkEventType)(keyPressrelease)

	var cret C.gboolean
	var goret1 bool

	cret = C.gdk_test_simulate_key(window, x, y, keyval, modifiers, keyPressrelease)

	goret1 = C.bool(cret) != C.false

	return goret1
}
