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
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialog},
	})
}

// ColorChooserDialog: the ColorChooserDialog widget is a dialog for choosing a
// color. It implements the ColorChooser interface.
type ColorChooserDialog interface {
	Dialog
	Buildable
	ColorChooser
}

// colorChooserDialog implements the ColorChooserDialog interface.
type colorChooserDialog struct {
	Dialog
	Buildable
	ColorChooser
}

var _ ColorChooserDialog = (*colorChooserDialog)(nil)

// WrapColorChooserDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorChooserDialog(obj *externglib.Object) ColorChooserDialog {
	return ColorChooserDialog{
		Dialog:       WrapDialog(obj),
		Buildable:    WrapBuildable(obj),
		ColorChooser: WrapColorChooser(obj),
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorChooserDialog(obj), nil
}

// NewColorChooserDialog constructs a class ColorChooserDialog.
func NewColorChooserDialog(title string, parent Window) ColorChooserDialog {
	var arg1 *C.gchar
	var arg2 *C.GtkWindow

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	var cret C.GtkColorChooserDialog
	var goret1 ColorChooserDialog

	cret = C.gtk_color_chooser_dialog_new(title, parent)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorChooserDialog)

	return goret1
}

type ColorChooserDialogPrivate struct {
	native C.GtkColorChooserDialogPrivate
}

// WrapColorChooserDialogPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColorChooserDialogPrivate(ptr unsafe.Pointer) *ColorChooserDialogPrivate {
	if ptr == nil {
		return nil
	}

	return (*ColorChooserDialogPrivate)(ptr)
}

func marshalColorChooserDialogPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColorChooserDialogPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *ColorChooserDialogPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}
