// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_app_chooser_dialog_get_type()), F: marshalAppChooserDialogger},
	})
}

// AppChooserDialogger describes AppChooserDialog's methods.
type AppChooserDialogger interface {
	gextras.Objector

	Heading() string
	Widget() *Widget
	SetHeading(heading string)
}

// AppChooserDialog shows a AppChooserWidget inside a Dialog.
//
// Note that AppChooserDialog does not have any interesting methods of its own.
// Instead, you should get the embedded AppChooserWidget using
// gtk_app_chooser_dialog_get_widget() and call its methods if the generic
// AppChooser interface is not sufficient for your needs.
//
// To set the heading that is shown above the AppChooserWidget, use
// gtk_app_chooser_dialog_set_heading().
type AppChooserDialog struct {
	*externglib.Object
	Dialog
	AppChooser
	Buildable
}

var _ AppChooserDialogger = (*AppChooserDialog)(nil)

func wrapAppChooserDialogger(obj *externglib.Object) AppChooserDialogger {
	return &AppChooserDialog{
		Object: obj,
		Dialog: Dialog{
			Object: obj,
			Window: Window{
				Object: obj,
				Bin: Bin{
					Object: obj,
					Container: Container{
						Object: obj,
						Widget: Widget{
							Object: obj,
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		AppChooser: AppChooser{
			Object: obj,
			Widget: Widget{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalAppChooserDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserDialogger(obj), nil
}

// Heading returns the text to display at the top of the dialog.
func (self *AppChooserDialog) Heading() string {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Widget returns the AppChooserWidget of this dialog.
func (self *AppChooserDialog) Widget() *Widget {
	var _arg0 *C.GtkAppChooserDialog // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_dialog_get_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// SetHeading sets the text to display at the top of the dialog. If the heading
// is not set, the dialog displays a default text.
func (self *AppChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
}
