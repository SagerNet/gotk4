// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tooltip_get_type()), F: marshalTooltipper},
	})
}

// Tooltip: GtkTooltip is an object representing a widget tooltip.
//
// Basic tooltips can be realized simply by using gtk.Widget.SetTooltipText() or
// gtk.Widget.SetTooltipMarkup() without any explicit tooltip object.
//
// When you need a tooltip with a little more fancy contents, like adding an
// image, or you want the tooltip to have different contents per GtkTreeView row
// or cell, you will have to do a little more work:
//
// - Set the gtk.Widget:has-tooltip property to TRUE. This will make GTK monitor
// the widget for motion and related events which are needed to determine when
// and where to show a tooltip.
//
// - Connect to the gtk.Widget::query-tooltip signal. This signal will be
// emitted when a tooltip is supposed to be shown. One of the arguments passed
// to the signal handler is a GtkTooltip object. This is the object that we are
// about to display as a tooltip, and can be manipulated in your callback using
// functions like gtk.Tooltip.SetIcon(). There are functions for setting the
// tooltip’s markup, setting an image from a named icon, or even putting in a
// custom widget.
//
// - Return TRUE from your ::query-tooltip handler. This causes the tooltip to
// be show. If you return FALSE, it will not be shown.
type Tooltip struct {
	*externglib.Object
}

func WrapTooltip(obj *externglib.Object) *Tooltip {
	return &Tooltip{
		Object: obj,
	}
}

func marshalTooltipper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTooltip(obj), nil
}

// SetCustom replaces the widget packed into the tooltip with custom_widget.
// custom_widget does not get destroyed when the tooltip goes away. By default a
// box with a Image and Label is embedded in the tooltip, which can be
// configured using gtk_tooltip_set_markup() and gtk_tooltip_set_icon().
func (tooltip *Tooltip) SetCustom(customWidget Widgetter) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if customWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(customWidget.Native()))
	}

	C.gtk_tooltip_set_custom(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(customWidget)
}

// SetIcon sets the icon of the tooltip (which is in front of the text) to be
// paintable. If paintable is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIcon(paintable gdk.Paintabler) {
	var _arg0 *C.GtkTooltip   // out
	var _arg1 *C.GdkPaintable // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if paintable != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))
	}

	C.gtk_tooltip_set_icon(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(paintable)
}

// SetIconFromGIcon sets the icon of the tooltip (which is in front of the text)
// to be the icon indicated by gicon with the size indicated by size. If gicon
// is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIconFromGIcon(gicon gio.Iconner) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GIcon      // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if gicon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(gicon.Native()))
	}

	C.gtk_tooltip_set_icon_from_gicon(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(gicon)
}

// SetIconFromIconName sets the icon of the tooltip (which is in front of the
// text) to be the icon indicated by icon_name with the size indicated by size.
// If icon_name is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIconFromIconName(iconName string) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_tooltip_set_icon_from_icon_name(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(iconName)
}

// SetMarkup sets the text of the tooltip to be markup.
//
// The string must be marked up with Pango markup. If markup is NULL, the label
// will be hidden.
func (tooltip *Tooltip) SetMarkup(markup string) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if markup != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(markup)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_tooltip_set_markup(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(markup)
}

// SetText sets the text of the tooltip to be text.
//
// If text is NULL, the label will be hidden. See also gtk.Tooltip.SetMarkup().
func (tooltip *Tooltip) SetText(text string) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	if text != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_tooltip_set_text(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(text)
}

// SetTipArea sets the area of the widget, where the contents of this tooltip
// apply, to be rect (in widget coordinates). This is especially useful for
// properly setting tooltips on TreeView rows and cells, IconViews, etc.
//
// For setting tooltips on TreeView, please refer to the convenience functions
// for this: gtk_tree_view_set_tooltip_row() and
// gtk_tree_view_set_tooltip_cell().
func (tooltip *Tooltip) SetTipArea(rect *gdk.Rectangle) {
	var _arg0 *C.GtkTooltip   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect)))

	C.gtk_tooltip_set_tip_area(_arg0, _arg1)
	runtime.KeepAlive(tooltip)
	runtime.KeepAlive(rect)
}
