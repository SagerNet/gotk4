// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooserer},
	})
}

// AppChooser is an interface that can be implemented by widgets which allow the
// user to choose an application (typically for the purpose of opening a file).
// The main objects that implement this interface are AppChooserWidget,
// AppChooserDialog and AppChooserButton.
//
// Applications are represented by GIO Info objects here. GIO has a concept of
// recommended and fallback applications for a given content type. Recommended
// applications are those that claim to handle the content type itself, while
// fallback also includes applications that handle a more generic content type.
// GIO also knows the default and last-used application for a given content
// type. The AppChooserWidget provides detailed control over whether the shown
// list of applications should include default, recommended or fallback
// applications.
//
// To obtain the application that has been selected in a AppChooser, use
// gtk_app_chooser_get_app_info().
type AppChooser struct {
	Widget
}

// AppChooserer describes AppChooser's abstract methods.
type AppChooserer interface {
	externglib.Objector

	// AppInfo returns the currently selected application.
	AppInfo() gio.AppInfor
	// ContentType returns the current value of the AppChooser:content-type
	// property.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

var _ AppChooserer = (*AppChooser)(nil)

func WrapAppChooser(obj *externglib.Object) *AppChooser {
	return &AppChooser{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalAppChooserer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooser(obj), nil
}

// AppInfo returns the currently selected application.
func (self *AppChooser) AppInfo() gio.AppInfor {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.GAppInfo      // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_app_info(_arg0)
	runtime.KeepAlive(self)

	var _appInfo gio.AppInfor // out

	if _cret != nil {
		_appInfo = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.AppInfor)
	}

	return _appInfo
}

// ContentType returns the current value of the AppChooser:content-type
// property.
func (self *AppChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_content_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Refresh reloads the list of applications.
func (self *AppChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_refresh(_arg0)
	runtime.KeepAlive(self)
}
