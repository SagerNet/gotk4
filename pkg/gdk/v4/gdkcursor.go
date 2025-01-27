// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_cursor_get_type()), F: marshalCursorrer},
	})
}

// Cursor: GdkCursor is used to create and destroy cursors.
//
// Cursors are immutable objects, so once you created them, there is no way to
// modify them later. You should create a new cursor when you want to change
// something about it.
//
// Cursors by themselves are not very interesting: they must be bound to a
// window for users to see them. This is done with gdk.Surface.SetCursor() or
// gdk.Surface.SetDeviceCursor(). Applications will typically use higher-level
// GTK functions such as gtk.Widget.SetCursor()` instead.
//
// Cursors are not bound to a given gdk.Display, so they can be shared. However,
// the appearance of cursors may vary when used on different platforms.
//
//
// Named and texture cursors
//
// There are multiple ways to create cursors. The platform's own cursors can be
// created with gdk.Cursor.NewFromName. That function lists the commonly
// available names that are shared with the CSS specification. Other names may
// be available, depending on the platform in use. On some platforms, what
// images are used for named cursors may be influenced by the cursor theme.
//
// Another option to create a cursor is to use gdk.Cursor.NewFromTexture and
// provide an image to use for the cursor.
//
// To ease work with unsupported cursors, a fallback cursor can be provided. If
// a gdk.Surface cannot use a cursor because of the reasons mentioned above, it
// will try the fallback cursor. Fallback cursors can themselves have fallback
// cursors again, so it is possible to provide a chain of progressively easier
// to support cursors. If none of the provided cursors can be supported, the
// default cursor will be the ultimate fallback.
type Cursor struct {
	*externglib.Object
}

func WrapCursor(obj *externglib.Object) *Cursor {
	return &Cursor{
		Object: obj,
	}
}

func marshalCursorrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCursor(obj), nil
}

// NewCursorFromName creates a new cursor by looking up name in the current
// cursor theme.
//
// A recommended set of cursor names that will work across different platforms
// can be found in the CSS specification:
//
// | | | | | | --- | --- | ---- | --- | | "none" | ! (default_cursor.png)
// "default" | ! (help_cursor.png) "help" | ! (pointer_cursor.png) "pointer" | |
// ! (context_menu_cursor.png) "context-menu" | ! (progress_cursor.png)
// "progress" | ! (wait_cursor.png) "wait" | ! (cell_cursor.png) "cell" | | !
// (crosshair_cursor.png) "crosshair" | ! (text_cursor.png) "text" | !
// (vertical_text_cursor.png) "vertical-text" | ! (alias_cursor.png) "alias" | |
// ! (copy_cursor.png) "copy" | ! (no_drop_cursor.png) "no-drop" | !
// (move_cursor.png) "move" | ! (not_allowed_cursor.png) "not-allowed" | | !
// (grab_cursor.png) "grab" | ! (grabbing_cursor.png) "grabbing" | !
// (all_scroll_cursor.png) "all-scroll" | ! (col_resize_cursor.png) "col-resize"
// | | ! (row_resize_cursor.png) "row-resize" | ! (n_resize_cursor.png)
// "n-resize" | ! (e_resize_cursor.png) "e-resize" | ! (s_resize_cursor.png)
// "s-resize" | | ! (w_resize_cursor.png) "w-resize" | ! (ne_resize_cursor.png)
// "ne-resize" | ! (nw_resize_cursor.png) "nw-resize" | ! (sw_resize_cursor.png)
// "sw-resize" | | ! (se_resize_cursor.png) "se-resize" | !
// (ew_resize_cursor.png) "ew-resize" | ! (ns_resize_cursor.png) "ns-resize" | !
// (nesw_resize_cursor.png) "nesw-resize" | | ! (nwse_resize_cursor.png)
// "nwse-resize" | ! (zoom_in_cursor.png) "zoom-in" | ! (zoom_out_cursor.png)
// "zoom-out" | |
func NewCursorFromName(name string, fallback *Cursor) *Cursor {
	var _arg1 *C.char      // out
	var _arg2 *C.GdkCursor // out
	var _cret *C.GdkCursor // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if fallback != nil {
		_arg2 = (*C.GdkCursor)(unsafe.Pointer(fallback.Native()))
	}

	_cret = C.gdk_cursor_new_from_name(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(fallback)

	var _cursor *Cursor // out

	if _cret != nil {
		_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _cursor
}

// NewCursorFromTexture creates a new cursor from a GdkTexture.
func NewCursorFromTexture(texture Texturer, hotspotX int, hotspotY int, fallback *Cursor) *Cursor {
	var _arg1 *C.GdkTexture // out
	var _arg2 C.int         // out
	var _arg3 C.int         // out
	var _arg4 *C.GdkCursor  // out
	var _cret *C.GdkCursor  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))
	_arg2 = C.int(hotspotX)
	_arg3 = C.int(hotspotY)
	if fallback != nil {
		_arg4 = (*C.GdkCursor)(unsafe.Pointer(fallback.Native()))
	}

	_cret = C.gdk_cursor_new_from_texture(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(texture)
	runtime.KeepAlive(hotspotX)
	runtime.KeepAlive(hotspotY)
	runtime.KeepAlive(fallback)

	var _cursor *Cursor // out

	_cursor = WrapCursor(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cursor
}

// Fallback returns the fallback for this cursor.
//
// The fallback will be used if this cursor is not available on a given
// GdkDisplay. For named cursors, this can happen when using nonstandard names
// or when using an incomplete cursor theme. For textured cursors, this can
// happen when the texture is too large or when the GdkDisplay it is used on
// does not support textured cursors.
func (cursor *Cursor) Fallback() *Cursor {
	var _arg0 *C.GdkCursor // out
	var _cret *C.GdkCursor // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_fallback(_arg0)
	runtime.KeepAlive(cursor)

	var _ret *Cursor // out

	if _cret != nil {
		_ret = WrapCursor(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _ret
}

// HotspotX returns the horizontal offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function will
// only return the hotspot position for cursors created with
// gdk.Cursor.NewFromTexture.
func (cursor *Cursor) HotspotX() int {
	var _arg0 *C.GdkCursor // out
	var _cret C.int        // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_hotspot_x(_arg0)
	runtime.KeepAlive(cursor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HotspotY returns the vertical offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function will
// only return the hotspot position for cursors created with
// gdk.Cursor.NewFromTexture.
func (cursor *Cursor) HotspotY() int {
	var _arg0 *C.GdkCursor // out
	var _cret C.int        // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_hotspot_y(_arg0)
	runtime.KeepAlive(cursor)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Name returns the name of the cursor.
//
// If the cursor is not a named cursor, NULL will be returned.
func (cursor *Cursor) Name() string {
	var _arg0 *C.GdkCursor // out
	var _cret *C.char      // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_name(_arg0)
	runtime.KeepAlive(cursor)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Texture returns the texture for the cursor.
//
// If the cursor is a named cursor, NULL will be returned.
func (cursor *Cursor) Texture() Texturer {
	var _arg0 *C.GdkCursor  // out
	var _cret *C.GdkTexture // in

	_arg0 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	_cret = C.gdk_cursor_get_texture(_arg0)
	runtime.KeepAlive(cursor)

	var _texture Texturer // out

	if _cret != nil {
		_texture = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Texturer)
	}

	return _texture
}
