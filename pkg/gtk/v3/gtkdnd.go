// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
func DragCancel(context gdk.DragContext) {
	var arg1 *C.GdkDragContext

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_cancel(context)
}

// DragFinish informs the drag source that the drop is finished, and that the
// data of the drag will no longer be required.
func DragFinish(context gdk.DragContext, success bool, del bool, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.gboolean
	var arg3 C.gboolean
	var arg4 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		arg2 = C.gboolean(1)
	}
	if del {
		arg3 = C.gboolean(1)
	}
	arg4 = C.guint32(time_)

	C.gtk_drag_finish(context, success, del, time_)
}

// DragGetSourceWidget determines the source widget for a drag.
func DragGetSourceWidget(context gdk.DragContext) Widget {
	var arg1 *C.GdkDragContext

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	var cret *C.GtkWidget
	var goret1 Widget

	cret = C.gtk_drag_get_source_widget(context)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret1
}

// DragSetIconDefault sets the icon for a particular drag to the default icon.
func DragSetIconDefault(context gdk.DragContext) {
	var arg1 *C.GdkDragContext

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	C.gtk_drag_set_icon_default(context)
}

// DragSetIconGIcon sets the icon for a given drag from the given @icon. See the
// documentation for gtk_drag_set_icon_name() for more details about using icons
// in drag and drop.
func DragSetIconGIcon(context gdk.DragContext, icon gio.Icon, hotX int, hotY int) {
	var arg1 *C.GdkDragContext
	var arg2 *C.GIcon
	var arg3 C.gint
	var arg4 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	arg3 = C.gint(hotX)
	arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_gicon(context, icon, hotX, hotY)
}

// DragSetIconName sets the icon for a given drag from a named themed icon. See
// the docs for IconTheme for more details. Note that the size of the icon
// depends on the icon theme (the icon is loaded at the symbolic size
// K_ICON_SIZE_DND), thus @hot_x and @hot_y have to be used with care.
func DragSetIconName(context gdk.DragContext, iconName string, hotX int, hotY int) {
	var arg1 *C.GdkDragContext
	var arg2 *C.gchar
	var arg3 C.gint
	var arg4 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gint(hotX)
	arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_name(context, iconName, hotX, hotY)
}

// DragSetIconPixbuf sets @pixbuf as the icon for a given drag.
func DragSetIconPixbuf(context gdk.DragContext, pixbuf gdkpixbuf.Pixbuf, hotX int, hotY int) {
	var arg1 *C.GdkDragContext
	var arg2 *C.GdkPixbuf
	var arg3 C.gint
	var arg4 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	arg3 = C.gint(hotX)
	arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_pixbuf(context, pixbuf, hotX, hotY)
}

// DragSetIconStock sets the icon for a given drag from a stock ID.
func DragSetIconStock(context gdk.DragContext, stockID string, hotX int, hotY int) {
	var arg1 *C.GdkDragContext
	var arg2 *C.gchar
	var arg3 C.gint
	var arg4 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gint(hotX)
	arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_stock(context, stockID, hotX, hotY)
}

// DragSetIconSurface sets @surface as the icon for a given drag. GTK+ retains
// references for the arguments, and will release them when they are no longer
// needed.
//
// To position the surface relative to the mouse, use
// cairo_surface_set_device_offset() on @surface. The mouse cursor will be
// positioned at the (0,0) coordinate of the surface.
func DragSetIconSurface(context gdk.DragContext, surface *cairo.Surface) {
	var arg1 *C.GdkDragContext
	var arg2 *C.cairo_surface_t

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	C.gtk_drag_set_icon_surface(context, surface)
}

// DragSetIconWidget changes the icon for drag operation to a given widget. GTK+
// will not destroy the widget, so if you don’t want it to persist, you should
// connect to the “drag-end” signal and destroy it yourself.
func DragSetIconWidget(context gdk.DragContext, widget Widget, hotX int, hotY int) {
	var arg1 *C.GdkDragContext
	var arg2 *C.GtkWidget
	var arg3 C.gint
	var arg4 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg3 = C.gint(hotX)
	arg4 = C.gint(hotY)

	C.gtk_drag_set_icon_widget(context, widget, hotX, hotY)
}
