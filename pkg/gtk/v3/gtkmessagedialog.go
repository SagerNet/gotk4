// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buttons_type_get_type()), F: marshalButtonsType},
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialog},
	})
}

// ButtonsType: prebuilt sets of buttons for the dialog. If none of these
// choices are appropriate, simply use GTK_BUTTONS_NONE then call
// gtk_dialog_add_buttons().
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType int

const (
	// none: no buttons at all
	ButtonsTypeNone ButtonsType = 0
	// ok: an OK button
	ButtonsTypeOk ButtonsType = 1
	// close: a Close button
	ButtonsTypeClose ButtonsType = 2
	// cancel: a Cancel button
	ButtonsTypeCancel ButtonsType = 3
	// YesNo yes and No buttons
	ButtonsTypeYesNo ButtonsType = 4
	// OkCancel: OK and Cancel buttons
	ButtonsTypeOkCancel ButtonsType = 5
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MessageDialog presents a dialog with some message text. It’s simply a
// convenience widget; you could construct the equivalent of MessageDialog from
// Dialog without too much effort, but MessageDialog saves typing.
//
// One difference from Dialog is that MessageDialog sets the
// Window:skip-taskbar-hint property to true, so that the dialog is hidden from
// the taskbar by default.
//
// The easiest way to do a modal message dialog is to use gtk_dialog_run(),
// though you can also pass in the GTK_DIALOG_MODAL flag, gtk_dialog_run()
// automatically makes the dialog modal and waits for the user to respond to it.
// gtk_dialog_run() returns when any dialog button is clicked.
//
// An example for using a modal dialog:
//
//     GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_message_dialog_new (parent_window,
//                                      flags,
//                                      GTK_MESSAGE_ERROR,
//                                      GTK_BUTTONS_CLOSE,
//                                      "Error reading “s”: s",
//                                      filename,
//                                      g_strerror (errno));
//
//     // Destroy the dialog when the user responds to it
//     // (e.g. clicks a button)
//
//     g_signal_connect_swapped (dialog, "response",
//                               G_CALLBACK (gtk_widget_destroy),
//                               dialog);
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog interface {
	Dialog

	// Image:
	Image() Widget
	// MessageArea:
	MessageArea() Widget
	// SetImageMessageDialog:
	SetImageMessageDialog(image Widget)
	// SetMarkupMessageDialog:
	SetMarkupMessageDialog(str string)
}

// messageDialog implements the MessageDialog class.
type messageDialog struct {
	Dialog
}

// WrapMessageDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapMessageDialog(obj *externglib.Object) MessageDialog {
	return messageDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalMessageDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMessageDialog(obj), nil
}

func (d messageDialog) Image() Widget {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_message_dialog_get_image(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (m messageDialog) MessageArea() Widget {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_message_dialog_get_message_area(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (d messageDialog) SetImageMessageDialog(image Widget) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(image.Native()))

	C.gtk_message_dialog_set_image(_arg0, _arg1)
}

func (m messageDialog) SetMarkupMessageDialog(str string) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_message_dialog_set_markup(_arg0, _arg1)
}

func (b messageDialog) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b messageDialog) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b messageDialog) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b messageDialog) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b messageDialog) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b messageDialog) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b messageDialog) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
