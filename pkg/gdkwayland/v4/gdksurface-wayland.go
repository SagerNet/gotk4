// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_popup_get_type()), F: marshalWaylandPopup},
		{T: externglib.Type(C.gdk_wayland_surface_get_type()), F: marshalWaylandSurface},
		{T: externglib.Type(C.gdk_wayland_toplevel_get_type()), F: marshalWaylandToplevel},
	})
}

// WaylandPopup: the Wayland implementation of `GdkPopup`.
type WaylandPopup interface {
	WaylandSurface
}

// waylandPopup implements the WaylandPopup class.
type waylandPopup struct {
	WaylandSurface
}

// WrapWaylandPopup wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandPopup(obj *externglib.Object) WaylandPopup {
	return waylandPopup{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandPopup(obj), nil
}

// WaylandSurface: the Wayland implementation of `GdkSurface`.
//
// Beyond the [class@Gdk.Surface] API, the Wayland implementation offers access
// to the Wayland `wl_surface` object with
// [method@GdkWayland.WaylandSurface.get_wl_surface].
type WaylandSurface interface {
	gdk.Surface
}

// waylandSurface implements the WaylandSurface class.
type waylandSurface struct {
	gdk.Surface
}

// WrapWaylandSurface wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandSurface(obj *externglib.Object) WaylandSurface {
	return waylandSurface{
		Surface: gdk.WrapSurface(obj),
	}
}

func marshalWaylandSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandSurface(obj), nil
}

// WaylandToplevel: the Wayland implementation of `GdkToplevel`.
//
// Beyond the [interface@Gdk.Toplevel] API, the Wayland implementation has API
// to set up cross-process parent-child relationships between surfaces with
// [method@GdkWayland.WaylandToplevel.export_handle] and
// [method@GdkWayland.WaylandToplevel.set_transient_for_exported].
type WaylandToplevel interface {
	WaylandSurface

	// SetApplicationID sets the application id on a `GdkToplevel`.
	SetApplicationID(applicationId string)
	// SetTransientForExported marks @toplevel as transient for the surface to
	// which the given @parent_handle_str refers.
	//
	// Typically, the handle will originate from a
	// [method@GdkWayland.WaylandToplevel.export_handle] call in another
	// process.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	SetTransientForExported(parentHandleStr string) bool
	// UnexportHandle destroys the handle that was obtained with
	// gdk_wayland_toplevel_export_handle().
	//
	// It is an error to call this function on a surface that does not have a
	// handle.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	UnexportHandle()
}

// waylandToplevel implements the WaylandToplevel class.
type waylandToplevel struct {
	WaylandSurface
}

// WrapWaylandToplevel wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandToplevel(obj *externglib.Object) WaylandToplevel {
	return waylandToplevel{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandToplevel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandToplevel(obj), nil
}

func (t waylandToplevel) SetApplicationID(applicationId string) {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_wayland_toplevel_set_application_id(_arg0, _arg1)
}

func (t waylandToplevel) SetTransientForExported(parentHandleStr string) bool {
	var _arg0 *C.GdkToplevel // out
	var _arg1 *C.char        // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.char)(C.CString(parentHandleStr))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_wayland_toplevel_set_transient_for_exported(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t waylandToplevel) UnexportHandle() {
	var _arg0 *C.GdkToplevel // out

	_arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	C.gdk_wayland_toplevel_unexport_handle(_arg0)
}
