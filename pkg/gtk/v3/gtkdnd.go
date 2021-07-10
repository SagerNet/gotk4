// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// DragCancel cancels an ongoing drag operation on the source side.
//
// If you want to be able to cancel a drag operation in this way, you need to
// keep a pointer to the drag context, either from an explicit call to
// gtk_drag_begin_with_coordinates(), or by connecting to Widget::drag-begin.
//
// If @context does not refer to an ongoing drag operation, this function does
// nothing.
//
// If a drag is cancelled in this way, the @result argument of
// Widget::drag-failed is set to @GTK_DRAG_RESULT_ERROR.
func DragCancel(context gdk.DragContexter) {
	var _arg1 *C.GdkDragContext // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_cancel(_arg1)
}

// DragFinish informs the drag source that the drop is finished, and that the
// data of the drag will no longer be required.
func DragFinish(context gdk.DragContexter, success bool, del bool, time_ uint32) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 C.gboolean        // out
	var _arg3 C.gboolean        // out
	var _arg4 C.guint32         // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		_arg2 = C.TRUE
	}
	if del {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint32(time_)

	C.gtk_drag_finish(_arg1, _arg2, _arg3, _arg4)
}

// DragGetSourceWidget determines the source widget for a drag.
func DragGetSourceWidget(context gdk.DragContexter) *Widget {
	var _arg1 *C.GdkDragContext // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_drag_get_source_widget(_arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// DragSetIconDefault sets the icon for a particular drag to the default icon.
func DragSetIconDefault(context gdk.DragContexter) {
	var _arg1 *C.GdkDragContext // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_set_icon_default(_arg1)
}

// DragSetIconName sets the icon for a given drag from a named themed icon. See
// the docs for IconTheme for more details. Note that the size of the icon
// depends on the icon theme (the icon is loaded at the symbolic size
// K_ICON_SIZE_DND), thus @hot_x and @hot_y have to be used with care.
func DragSetIconName(context gdk.DragContexter, iconName string, hotX int, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.gchar          // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_name(_arg1, _arg2, _arg3, _arg4)
}

// DragSetIconPixbuf sets @pixbuf as the icon for a given drag.
func DragSetIconPixbuf(context gdk.DragContexter, pixbuf gdkpixbuf.Pixbuffer, hotX int, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkPixbuf      // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_pixbuf(_arg1, _arg2, _arg3, _arg4)
}

// DragSetIconStock sets the icon for a given drag from a stock ID.
//
// Deprecated: Use gtk_drag_set_icon_name() instead.
func DragSetIconStock(context gdk.DragContexter, stockId string, hotX int, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.gchar          // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_stock(_arg1, _arg2, _arg3, _arg4)
}

// DragSetIconSurface sets @surface as the icon for a given drag. GTK+ retains
// references for the arguments, and will release them when they are no longer
// needed.
//
// To position the surface relative to the mouse, use
// cairo_surface_set_device_offset() on @surface. The mouse cursor will be
// positioned at the (0,0) coordinate of the surface.
func DragSetIconSurface(context gdk.DragContexter, surface *cairo.Surface) {
	var _arg1 *C.GdkDragContext  // out
	var _arg2 *C.cairo_surface_t // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface))

	C.gtk_drag_set_icon_surface(_arg1, _arg2)
}

// DragSetIconWidget changes the icon for drag operation to a given widget. GTK+
// will not destroy the widget, so if you don’t want it to persist, you should
// connect to the “drag-end” signal and destroy it yourself.
func DragSetIconWidget(context gdk.DragContexter, widget Widgetter, hotX int, hotY int) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GtkWidget      // out
	var _arg3 C.gint            // out
	var _arg4 C.gint            // out

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.gint(hotX)
	_arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_widget(_arg1, _arg2, _arg3, _arg4)
}
