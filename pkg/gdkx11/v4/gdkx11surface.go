// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_surface_get_type()), F: marshalX11Surface},
	})
}

// X11GetServerTime: routine to get the current X server time stamp.
func X11GetServerTime(surface X11Surface) uint32 {
	var arg1 *C.GdkSurface

	arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	var cret C.guint32
	var ret1 uint32

	cret = C.gdk_x11_get_server_time(surface)

	ret1 = C.guint32(cret)

	return ret1
}

type X11Surface interface {
	gdk.Surface

	// Desktop gets the number of the workspace @surface is on.
	Desktop() uint32
	// Group returns the group this surface belongs to.
	Group() gdk.Surface
	// MoveToCurrentDesktop moves the surface to the correct workspace when
	// running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification. Will not do
	// anything if the surface is already on all workspaces.
	MoveToCurrentDesktop()
	// MoveToDesktop moves the surface to the given workspace when running unde
	// a window manager that supports multiple workspaces, as described in the
	// Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	MoveToDesktop(desktop uint32)
	// SetFrameSyncEnabled: this function can be used to disable frame
	// synchronization for a surface. Normally frame synchronziation will be
	// enabled or disabled based on whether the system has a compositor that
	// supports frame synchronization, but if the surface is not directly
	// managed by the window manager, then frame synchronziation may need to be
	// disabled. This is the case for a surface embedded via the XEMBED
	// protocol.
	SetFrameSyncEnabled(frameSyncEnabled bool)
	// SetGroup sets the group leader of @surface to be @leader. See the ICCCM
	// for details.
	SetGroup(leader gdk.Surface)
	// SetSkipPagerHint sets a hint on @surface that pagers should not display
	// it. See the EWMH for details.
	SetSkipPagerHint(skipsPager bool)
	// SetSkipTaskbarHint sets a hint on @surface that taskbars should not
	// display it. See the EWMH for details.
	SetSkipTaskbarHint(skipsTaskbar bool)
	// SetThemeVariant: GTK applications can request a dark theme variant. In
	// order to make other applications - namely window managers using GTK for
	// themeing - aware of this choice, GTK uses this function to export the
	// requested theme variant as _GTK_THEME_VARIANT property on toplevel
	// surfaces.
	//
	// Note that this property is automatically updated by GTK, so this function
	// should only be used by applications which do not use GTK to create
	// toplevel surfaces.
	SetThemeVariant(variant string)
	// SetUrgencyHint sets a hint on @surface that it needs user attention. See
	// the ICCCM for details.
	SetUrgencyHint(urgent bool)
	// SetUserTime: the application can use this call to update the
	// _NET_WM_USER_TIME property on a toplevel surface. This property stores an
	// Xserver time which represents the time of the last user input event
	// received for this surface. This property may be used by the window
	// manager to alter the focus, stacking, and/or placement behavior of
	// surfaces when they are mapped depending on whether the new surface was
	// created by a user action or is a "pop-up" surface activated by a timer or
	// some other event.
	//
	// Note that this property is automatically updated by GDK, so this function
	// should only be used by applications which handle input events bypassing
	// GDK.
	SetUserTime(timestamp uint32)
	// SetUTF8Property: this function modifies or removes an arbitrary X11
	// window property of type UTF8_STRING. If the given @surface is not a
	// toplevel surface, it is ignored.
	SetUTF8Property(name string, value string)
}

// x11Surface implements the X11Surface interface.
type x11Surface struct {
	gdk.Surface
}

var _ X11Surface = (*x11Surface)(nil)

// WrapX11Surface wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Surface(obj *externglib.Object) X11Surface {
	return X11Surface{
		gdk.Surface: gdk.WrapSurface(obj),
	}
}

func marshalX11Surface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Surface(obj), nil
}

// Desktop gets the number of the workspace @surface is on.
func (s x11Surface) Desktop() uint32 {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret C.guint32
	var ret1 uint32

	cret = C.gdk_x11_surface_get_desktop(arg0)

	ret1 = C.guint32(cret)

	return ret1
}

// Group returns the group this surface belongs to.
func (s x11Surface) Group() gdk.Surface {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var cret *C.GdkSurface
	var ret1 gdk.Surface

	cret = C.gdk_x11_surface_get_group(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Surface)

	return ret1
}

// MoveToCurrentDesktop moves the surface to the correct workspace when
// running under a window manager that supports multiple workspaces, as
// described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification. Will not do
// anything if the surface is already on all workspaces.
func (s x11Surface) MoveToCurrentDesktop() {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_x11_surface_move_to_current_desktop(arg0)
}

// MoveToDesktop moves the surface to the given workspace when running unde
// a window manager that supports multiple workspaces, as described in the
// Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification.
func (s x11Surface) MoveToDesktop(desktop uint32) {
	var arg0 *C.GdkSurface
	var arg1 C.guint32

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = C.guint32(desktop)

	C.gdk_x11_surface_move_to_desktop(arg0, desktop)
}

// SetFrameSyncEnabled: this function can be used to disable frame
// synchronization for a surface. Normally frame synchronziation will be
// enabled or disabled based on whether the system has a compositor that
// supports frame synchronization, but if the surface is not directly
// managed by the window manager, then frame synchronziation may need to be
// disabled. This is the case for a surface embedded via the XEMBED
// protocol.
func (s x11Surface) SetFrameSyncEnabled(frameSyncEnabled bool) {
	var arg0 *C.GdkSurface
	var arg1 C.gboolean

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if frameSyncEnabled {
		arg1 = C.gboolean(1)
	}

	C.gdk_x11_surface_set_frame_sync_enabled(arg0, frameSyncEnabled)
}

// SetGroup sets the group leader of @surface to be @leader. See the ICCCM
// for details.
func (s x11Surface) SetGroup(leader gdk.Surface) {
	var arg0 *C.GdkSurface
	var arg1 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkSurface)(unsafe.Pointer(leader.Native()))

	C.gdk_x11_surface_set_group(arg0, leader)
}

// SetSkipPagerHint sets a hint on @surface that pagers should not display
// it. See the EWMH for details.
func (s x11Surface) SetSkipPagerHint(skipsPager bool) {
	var arg0 *C.GdkSurface
	var arg1 C.gboolean

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if skipsPager {
		arg1 = C.gboolean(1)
	}

	C.gdk_x11_surface_set_skip_pager_hint(arg0, skipsPager)
}

// SetSkipTaskbarHint sets a hint on @surface that taskbars should not
// display it. See the EWMH for details.
func (s x11Surface) SetSkipTaskbarHint(skipsTaskbar bool) {
	var arg0 *C.GdkSurface
	var arg1 C.gboolean

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if skipsTaskbar {
		arg1 = C.gboolean(1)
	}

	C.gdk_x11_surface_set_skip_taskbar_hint(arg0, skipsTaskbar)
}

// SetThemeVariant: GTK applications can request a dark theme variant. In
// order to make other applications - namely window managers using GTK for
// themeing - aware of this choice, GTK uses this function to export the
// requested theme variant as _GTK_THEME_VARIANT property on toplevel
// surfaces.
//
// Note that this property is automatically updated by GTK, so this function
// should only be used by applications which do not use GTK to create
// toplevel surfaces.
func (s x11Surface) SetThemeVariant(variant string) {
	var arg0 *C.GdkSurface
	var arg1 *C.char

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(variant))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_surface_set_theme_variant(arg0, variant)
}

// SetUrgencyHint sets a hint on @surface that it needs user attention. See
// the ICCCM for details.
func (s x11Surface) SetUrgencyHint(urgent bool) {
	var arg0 *C.GdkSurface
	var arg1 C.gboolean

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	if urgent {
		arg1 = C.gboolean(1)
	}

	C.gdk_x11_surface_set_urgency_hint(arg0, urgent)
}

// SetUserTime: the application can use this call to update the
// _NET_WM_USER_TIME property on a toplevel surface. This property stores an
// Xserver time which represents the time of the last user input event
// received for this surface. This property may be used by the window
// manager to alter the focus, stacking, and/or placement behavior of
// surfaces when they are mapped depending on whether the new surface was
// created by a user action or is a "pop-up" surface activated by a timer or
// some other event.
//
// Note that this property is automatically updated by GDK, so this function
// should only be used by applications which handle input events bypassing
// GDK.
func (s x11Surface) SetUserTime(timestamp uint32) {
	var arg0 *C.GdkSurface
	var arg1 C.guint32

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = C.guint32(timestamp)

	C.gdk_x11_surface_set_user_time(arg0, timestamp)
}

// SetUTF8Property: this function modifies or removes an arbitrary X11
// window property of type UTF8_STRING. If the given @surface is not a
// toplevel surface, it is ignored.
func (s x11Surface) SetUTF8Property(name string, value string) {
	var arg0 *C.GdkSurface
	var arg1 *C.char
	var arg2 *C.char

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(arg2))

	C.gdk_x11_surface_set_utf8_property(arg0, name, value)
}