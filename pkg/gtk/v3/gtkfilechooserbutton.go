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
		{T: externglib.Type(C.gtk_file_chooser_button_get_type()), F: marshalFileChooserButtonner},
	})
}

// FileChooserButtonnerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FileChooserButtonnerOverrider interface {
	FileSet()
}

// FileChooserButtonner describes FileChooserButton's methods.
type FileChooserButtonner interface {
	gextras.Objector

	FocusOnClick() bool
	Title() string
	WidthChars() int
	SetFocusOnClick(focusOnClick bool)
	SetTitle(title string)
	SetWidthChars(nChars int)
}

// FileChooserButton: the FileChooserButton is a widget that lets the user
// select a file. It implements the FileChooser interface. Visually, it is a
// file name with a button to bring up a FileChooserDialog. The user can then
// use that dialog to change the file associated with that button. This widget
// does not support setting the FileChooser:select-multiple property to true.
//
// Create a button to let the user select a file in /etc
//
//    {
//      GtkWidget *button;
//
//      button = gtk_file_chooser_button_new (_("Select a file"),
//                                            GTK_FILE_CHOOSER_ACTION_OPEN);
//      gtk_file_chooser_set_current_folder (GTK_FILE_CHOOSER (button),
//                                           "/etc");
//    }
//
// The FileChooserButton supports the FileChooserActions
// GTK_FILE_CHOOSER_ACTION_OPEN and GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
//
// > The FileChooserButton will ellipsize the label, and will thus > request
// little horizontal space. To give the button more space, > you should call
// gtk_widget_get_preferred_size(), > gtk_file_chooser_button_set_width_chars(),
// or pack the button in > such a way that other interface elements give space
// to the > widget.
//
//
// CSS nodes
//
// GtkFileChooserButton has a CSS node with name “filechooserbutton”, containing
// a subnode for the internal button with name “button” and style class “.file”.
type FileChooserButton struct {
	*externglib.Object
	Box
	Buildable
	FileChooser
	Orientable
}

var _ FileChooserButtonner = (*FileChooserButton)(nil)

func wrapFileChooserButtonner(obj *externglib.Object) FileChooserButtonner {
	return &FileChooserButton{
		Object: obj,
		Box: Box{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
		FileChooser: FileChooser{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalFileChooserButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFileChooserButtonner(obj), nil
}

// NewFileChooserButtonWithDialog creates a FileChooserButton widget which uses
// @dialog as its file-picking window.
//
// Note that @dialog must be a Dialog (or subclass) which implements the
// FileChooser interface and must not have GTK_DIALOG_DESTROY_WITH_PARENT set.
//
// Also note that the dialog needs to have its confirmative button added with
// response GTK_RESPONSE_ACCEPT or GTK_RESPONSE_OK in order for the button to
// take over the file selected in the dialog.
func NewFileChooserButtonWithDialog(dialog Dialogger) *FileChooserButton {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_file_chooser_button_new_with_dialog(_arg1)

	var _fileChooserButton *FileChooserButton // out

	_fileChooserButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*FileChooserButton)

	return _fileChooserButton
}

// FocusOnClick returns whether the button grabs focus when it is clicked with
// the mouse. See gtk_file_chooser_button_set_focus_on_click().
//
// Deprecated: Use gtk_widget_get_focus_on_click() instead.
func (button *FileChooserButton) FocusOnClick() bool {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title retrieves the title of the browse dialog used by @button. The returned
// value should not be modified or freed.
func (button *FileChooserButton) Title() string {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret *C.gchar                // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// WidthChars retrieves the width in characters of the @button widget’s entry
// and/or label.
func (button *FileChooserButton) WidthChars() int {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gint                  // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))

	_cret = C.gtk_file_chooser_button_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetFocusOnClick sets whether the button will grab focus when it is clicked
// with the mouse. Making mouse clicks not grab focus is useful in places like
// toolbars where you don’t want the keyboard focus removed from the main area
// of the application.
//
// Deprecated: Use gtk_widget_set_focus_on_click() instead.
func (button *FileChooserButton) SetFocusOnClick(focusOnClick bool) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_button_set_focus_on_click(_arg0, _arg1)
}

// SetTitle modifies the @title of the browse dialog used by @button.
func (button *FileChooserButton) SetTitle(title string) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_button_set_title(_arg0, _arg1)
}

// SetWidthChars sets the width (in characters) that @button will use to
// @n_chars.
func (button *FileChooserButton) SetWidthChars(nChars int) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gint                  // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(button.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_file_chooser_button_set_width_chars(_arg0, _arg1)
}
