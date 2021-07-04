// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_fullscreen_mode_get_type()), F: marshalFullscreenMode},
		{T: externglib.Type(C.gdk_surface_edge_get_type()), F: marshalSurfaceEdge},
		{T: externglib.Type(C.gdk_toplevel_state_get_type()), F: marshalToplevelState},
		{T: externglib.Type(C.gdk_toplevel_get_type()), F: marshalToplevel},
	})
}

// FullscreenMode indicates which monitor a surface should span over when in
// fullscreen mode.
type FullscreenMode int

const (
	// CurrentMonitor: fullscreen on current monitor only.
	FullscreenModeCurrentMonitor FullscreenMode = 0
	// AllMonitors: span across all monitors when fullscreen.
	FullscreenModeAllMonitors FullscreenMode = 1
)

func marshalFullscreenMode(p uintptr) (interface{}, error) {
	return FullscreenMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SurfaceEdge determines a surface edge or corner.
type SurfaceEdge int

const (
	// NorthWest: the top left corner.
	SurfaceEdgeNorthWest SurfaceEdge = 0
	// north: the top edge.
	SurfaceEdgeNorth SurfaceEdge = 1
	// NorthEast: the top right corner.
	SurfaceEdgeNorthEast SurfaceEdge = 2
	// west: the left edge.
	SurfaceEdgeWest SurfaceEdge = 3
	// east: the right edge.
	SurfaceEdgeEast SurfaceEdge = 4
	// SouthWest: the lower left corner.
	SurfaceEdgeSouthWest SurfaceEdge = 5
	// south: the lower edge.
	SurfaceEdgeSouth SurfaceEdge = 6
	// SouthEast: the lower right corner.
	SurfaceEdgeSouthEast SurfaceEdge = 7
)

func marshalSurfaceEdge(p uintptr) (interface{}, error) {
	return SurfaceEdge(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ToplevelState specifies the state of a toplevel surface.
//
// On platforms that support information about individual edges, the
// GDK_TOPLEVEL_STATE_TILED state will be set whenever any of the individual
// tiled states is set. On platforms that lack that support, the tiled state
// will give an indication of tiledness without any of the per-edge states being
// set.
type ToplevelState int

const (
	// ToplevelStateMinimized: the surface is minimized
	ToplevelStateMinimized ToplevelState = 0b1
	// ToplevelStateMaximized: the surface is maximized
	ToplevelStateMaximized ToplevelState = 0b10
	// ToplevelStateSticky: the surface is sticky
	ToplevelStateSticky ToplevelState = 0b100
	// ToplevelStateFullscreen: the surface is maximized without decorations
	ToplevelStateFullscreen ToplevelState = 0b1000
	// ToplevelStateAbove: the surface is kept above other surfaces
	ToplevelStateAbove ToplevelState = 0b10000
	// ToplevelStateBelow: the surface is kept below other surfaces
	ToplevelStateBelow ToplevelState = 0b100000
	// ToplevelStateFocused: the surface is presented as focused (with active
	// decorations)
	ToplevelStateFocused ToplevelState = 0b1000000
	// ToplevelStateTiled: the surface is in a tiled state
	ToplevelStateTiled ToplevelState = 0b10000000
	// ToplevelStateTopTiled: whether the top edge is tiled
	ToplevelStateTopTiled ToplevelState = 0b100000000
	// ToplevelStateTopResizable: whether the top edge is resizable
	ToplevelStateTopResizable ToplevelState = 0b1000000000
	// ToplevelStateRightTiled: whether the right edge is tiled
	ToplevelStateRightTiled ToplevelState = 0b10000000000
	// ToplevelStateRightResizable: whether the right edge is resizable
	ToplevelStateRightResizable ToplevelState = 0b100000000000
	// ToplevelStateBottomTiled: whether the bottom edge is tiled
	ToplevelStateBottomTiled ToplevelState = 0b1000000000000
	// ToplevelStateBottomResizable: whether the bottom edge is resizable
	ToplevelStateBottomResizable ToplevelState = 0b10000000000000
	// ToplevelStateLeftTiled: whether the left edge is tiled
	ToplevelStateLeftTiled ToplevelState = 0b100000000000000
	// ToplevelStateLeftResizable: whether the left edge is resizable
	ToplevelStateLeftResizable ToplevelState = 0b1000000000000000
)

func marshalToplevelState(p uintptr) (interface{}, error) {
	return ToplevelState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Toplevel: a `GdkToplevel` is a freestanding toplevel surface.
//
// The `GdkToplevel` interface provides useful APIs for interacting with the
// windowing system, such as controlling maximization and size of the surface,
// setting icons and transient parents for dialogs.
type Toplevel interface {
	Surface

	// BeginMove returns whether the desktop environment supports tiled window
	// states.
	BeginMove(device Device, button int, x float64, y float64, timestamp uint32)
	// BeginResize returns whether the desktop environment supports tiled window
	// states.
	BeginResize(edge SurfaceEdge, device Device, button int, x float64, y float64, timestamp uint32)
	// Focus returns whether the desktop environment supports tiled window
	// states.
	Focus(timestamp uint32)
	// State returns whether the desktop environment supports tiled window
	// states.
	State() ToplevelState
	// InhibitSystemShortcuts returns whether the desktop environment supports
	// tiled window states.
	InhibitSystemShortcuts(event Event)
	// Lower returns whether the desktop environment supports tiled window
	// states.
	Lower() bool
	// Minimize returns whether the desktop environment supports tiled window
	// states.
	Minimize() bool
	// Present returns whether the desktop environment supports tiled window
	// states.
	Present(layout *ToplevelLayout)
	// RestoreSystemShortcuts returns whether the desktop environment supports
	// tiled window states.
	RestoreSystemShortcuts()
	// SetDecorated returns whether the desktop environment supports tiled
	// window states.
	SetDecorated(decorated bool)
	// SetDeletable returns whether the desktop environment supports tiled
	// window states.
	SetDeletable(deletable bool)
	// SetModal returns whether the desktop environment supports tiled window
	// states.
	SetModal(modal bool)
	// SetStartupID returns whether the desktop environment supports tiled
	// window states.
	SetStartupID(startupId string)
	// SetTitle returns whether the desktop environment supports tiled window
	// states.
	SetTitle(title string)
	// SetTransientFor returns whether the desktop environment supports tiled
	// window states.
	SetTransientFor(parent Surface)
	// ShowWindowMenu returns whether the desktop environment supports tiled
	// window states.
	ShowWindowMenu(event Event) bool
	// SupportsEdgeConstraints returns whether the desktop environment supports
	// tiled window states.
	SupportsEdgeConstraints() bool
}

// toplevel implements the Toplevel interface.
type toplevel struct {
	Surface
}

var _ Toplevel = (*toplevel)(nil)

// WrapToplevel wraps a GObject to a type that implements
// interface Toplevel. It is primarily used internally.
func WrapToplevel(obj *externglib.Object) Toplevel {
	return toplevel{
		Surface: WrapSurface(obj),
	}
}

func marshalToplevel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToplevel(obj), nil
}

func (t toplevel) BeginMove(device Device, button int, x float64, y float64, timestamp uint32) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkDevice   // out
	var _arg2 C.int          // out
	var _arg3 C.double       // out
	var _arg4 C.double       // out
	var _arg5 C.guint32      // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg2 = C.int(button)
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.guint32(timestamp)

	C.gdk_toplevel_begin_move(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (t toplevel) BeginResize(edge SurfaceEdge, device Device, button int, x float64, y float64, timestamp uint32) {
	var _arg0 *C.GdkToplevel   // out
	var _arg1 C.GdkSurfaceEdge // out
	var _arg2 *C.GdkDevice     // out
	var _arg3 C.int            // out
	var _arg4 C.double         // out
	var _arg5 C.double         // out
	var _arg6 C.guint32        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = C.GdkSurfaceEdge(edge)
	_arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg3 = C.int(button)
	_arg4 = C.double(x)
	_arg5 = C.double(y)
	_arg6 = C.guint32(timestamp)

	C.gdk_toplevel_begin_resize(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (t toplevel) Focus(timestamp uint32) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.guint32      // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint32(timestamp)

	C.gdk_toplevel_focus(_arg0, _arg1)
}

func (t toplevel) State() ToplevelState {
	var _arg0 *C.GdkToplevel     // out
	var _cret C.GdkToplevelState // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_toplevel_get_state(_arg0)

	var _toplevelState ToplevelState // out

	_toplevelState = ToplevelState(_cret)

	return _toplevelState
}

func (t toplevel) InhibitSystemShortcuts(event Event) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkEvent    // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	C.gdk_toplevel_inhibit_system_shortcuts(_arg0, _arg1)
}

func (t toplevel) Lower() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_toplevel_lower(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toplevel) Minimize() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_toplevel_minimize(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toplevel) Present(layout *ToplevelLayout) {
	var _arg0 *C.GdkToplevel       // out
	var _arg1 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkToplevelLayout)(unsafe.Pointer(layout.Native()))

	C.gdk_toplevel_present(_arg0, _arg1)
}

func (t toplevel) RestoreSystemShortcuts() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	C.gdk_toplevel_restore_system_shortcuts(_arg0)
}

func (t toplevel) SetDecorated(decorated bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	if decorated {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_decorated(_arg0, _arg1)
}

func (t toplevel) SetDeletable(deletable bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	if deletable {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_deletable(_arg0, _arg1)
}

func (t toplevel) SetModal(modal bool) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_set_modal(_arg0, _arg1)
}

func (t toplevel) SetStartupID(startupId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_toplevel_set_startup_id(_arg0, _arg1)
}

func (t toplevel) SetTitle(title string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_toplevel_set_title(_arg0, _arg1)
}

func (t toplevel) SetTransientFor(parent Surface) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkSurface  // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(parent.Native()))

	C.gdk_toplevel_set_transient_for(_arg0, _arg1)
}

func (t toplevel) ShowWindowMenu(event Event) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.GdkEvent    // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	_cret = C.gdk_toplevel_show_window_menu(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toplevel) SupportsEdgeConstraints() bool {
	var _arg0 *C.GdkToplevel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	_cret = C.gdk_toplevel_supports_edge_constraints(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}