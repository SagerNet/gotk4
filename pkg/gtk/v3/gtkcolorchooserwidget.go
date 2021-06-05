// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_widget_get_type()), F: marshalColorChooserWidget},
	})
}

// ColorChooserWidget: the ColorChooserWidget widget lets the user select a
// color. By default, the chooser presents a predefined palette of colors, plus
// a small number of settable custom colors. It is also possible to select a
// different color with the single-color editor. To enter the single-color
// editing mode, use the context menu of any color of the palette, or use the
// '+' button to add a new custom color.
//
// The chooser automatically remembers the last selection, as well as custom
// colors.
//
// To change the initially selected color, use gtk_color_chooser_set_rgba(). To
// get the selected color use gtk_color_chooser_get_rgba().
//
// The ColorChooserWidget is used in the ColorChooserDialog to provide a dialog
// for selecting colors.
//
//
// CSS names
//
// GtkColorChooserWidget has a single CSS node with name colorchooser.
type ColorChooserWidget interface {
	Box
	Buildable
	ColorChooser
	Orientable
}

// colorChooserWidget implements the ColorChooserWidget interface.
type colorChooserWidget struct {
	Box
	Buildable
	ColorChooser
	Orientable
}

var _ ColorChooserWidget = (*colorChooserWidget)(nil)

// WrapColorChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorChooserWidget(obj *externglib.Object) ColorChooserWidget {
	return ColorChooserWidget{
		Box:          WrapBox(obj),
		Buildable:    WrapBuildable(obj),
		ColorChooser: WrapColorChooser(obj),
		Orientable:   WrapOrientable(obj),
	}
}

func marshalColorChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorChooserWidget(obj), nil
}

// NewColorChooserWidget constructs a class ColorChooserWidget.
func NewColorChooserWidget() ColorChooserWidget {
	var cret C.GtkColorChooserWidget
	var goret1 ColorChooserWidget

	cret = C.gtk_color_chooser_widget_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorChooserWidget)

	return goret1
}

type ColorChooserWidgetPrivate struct {
	native C.GtkColorChooserWidgetPrivate
}

// WrapColorChooserWidgetPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColorChooserWidgetPrivate(ptr unsafe.Pointer) *ColorChooserWidgetPrivate {
	if ptr == nil {
		return nil
	}

	return (*ColorChooserWidgetPrivate)(ptr)
}

func marshalColorChooserWidgetPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColorChooserWidgetPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *ColorChooserWidgetPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
