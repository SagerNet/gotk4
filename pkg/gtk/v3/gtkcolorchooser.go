// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_color_chooser_get_type()), F: marshalColorChooser},
	})
}

// ColorChooserOverrider contains methods that are overridable. This
// interface is a subset of the interface ColorChooser.
type ColorChooserOverrider interface {
	ColorActivated(color *gdk.RGBA)
	// RGBA gets the currently-selected color.
	RGBA() gdk.RGBA
	// SetRGBA sets the color.
	SetRGBA(color *gdk.RGBA)
}

// ColorChooser is an interface that is implemented by widgets for choosing
// colors. Depending on the situation, colors may be allowed to have alpha
// (translucency).
//
// In GTK+, the main widgets that implement this interface are
// ColorChooserWidget, ColorChooserDialog and ColorButton.
type ColorChooser interface {
	gextras.Objector
	ColorChooserOverrider

	// UseAlpha returns whether the color chooser shows the alpha channel.
	UseAlpha() bool
	// SetUseAlpha sets whether or not the color chooser should use the alpha
	// channel.
	SetUseAlpha(useAlpha bool)
}

// colorChooser implements the ColorChooser interface.
type colorChooser struct {
	gextras.Objector
}

var _ ColorChooser = (*colorChooser)(nil)

// WrapColorChooser wraps a GObject to a type that implements interface
// ColorChooser. It is primarily used internally.
func WrapColorChooser(obj *externglib.Object) ColorChooser {
	return ColorChooser{
		Objector: obj,
	}
}

func marshalColorChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorChooser(obj), nil
}

// RGBA gets the currently-selected color.
func (c colorChooser) RGBA() gdk.RGBA {
	var arg0 *C.GtkColorChooser

	arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkRGBA
	var ret1 *gdk.RGBA

	C.gtk_color_chooser_get_rgba(arg0, &arg1)

	ret1 = gdk.WrapRGBA(unsafe.Pointer(arg1))

	return ret1
}

// UseAlpha returns whether the color chooser shows the alpha channel.
func (c colorChooser) UseAlpha() bool {
	var arg0 *C.GtkColorChooser

	arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_color_chooser_get_use_alpha(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetRGBA sets the color.
func (c colorChooser) SetRGBA(color *gdk.RGBA) {
	var arg0 *C.GtkColorChooser
	var arg1 *C.GdkRGBA

	arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))

	C.gtk_color_chooser_set_rgba(arg0, color)
}

// SetUseAlpha sets whether or not the color chooser should use the alpha
// channel.
func (c colorChooser) SetUseAlpha(useAlpha bool) {
	var arg0 *C.GtkColorChooser
	var arg1 C.gboolean

	arg0 = (*C.GtkColorChooser)(unsafe.Pointer(c.Native()))
	if useAlpha {
		arg1 = C.gboolean(1)
	}

	C.gtk_color_chooser_set_use_alpha(arg0, useAlpha)
}
