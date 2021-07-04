// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_im_context_get_type()), F: marshalIMContext},
	})
}

// IMContext defines the interface for GTK+ input methods. An input method is
// used by GTK+ text input widgets like Entry to map from key events to Unicode
// character strings.
//
// The default input method can be set programmatically via the
// Settings:gtk-im-module GtkSettings property. Alternatively, you may set the
// GTK_IM_MODULE environment variable as documented in [Running GTK+
// Applications][gtk-running].
//
// The Entry Entry:im-module and TextView TextView:im-module properties may also
// be used to set input methods for specific widget instances. For instance, a
// certain entry widget might be expected to contain certain characters which
// would be easier to input with a certain input method.
//
// An input method may consume multiple key events in sequence and finally
// output the composed result. This is called preediting, and an input method
// may provide feedback about this process by displaying the intermediate
// composition states as preedit text. For instance, the default GTK+ input
// method implements the input of arbitrary Unicode code points by holding down
// the Control and Shift keys and then typing “U” followed by the hexadecimal
// digits of the code point. When releasing the Control and Shift keys,
// preediting ends and the character is inserted as text. Ctrl+Shift+u20AC for
// example results in the € sign.
//
// Additional input methods can be made available for use by GTK+ widgets as
// loadable modules. An input method module is a small shared library which
// implements a subclass of IMContext or IMContextSimple and exports these four
// functions:
//
//    GtkIMContext * im_module_create(const gchar *context_id);
//
// This function should return a pointer to a newly created instance of the
// IMContext subclass identified by @context_id. The context ID is the same as
// specified in the IMContextInfo array returned by im_module_list().
//
// After a new loadable input method module has been installed on the system,
// the configuration file `gtk.immodules` needs to be regenerated by
// [gtk-query-immodules-3.0][gtk-query-immodules-3.0], in order for the new
// input method to become available to GTK+ applications.
type IMContext interface {
	gextras.Objector

	DeleteSurroundingIMContext(offset int, nChars int) bool

	FilterKeypressIMContext(event *gdk.EventKey) bool

	FocusInIMContext()

	FocusOutIMContext()

	PreeditString() (string, *pango.AttrList, int)

	Surrounding() (string, int, bool)

	ResetIMContext()

	SetClientWindowIMContext(window gdk.Window)

	SetCursorLocationIMContext(area *gdk.Rectangle)

	SetSurroundingIMContext(text string, len int, cursorIndex int)

	SetUsePreeditIMContext(usePreedit bool)
}

// imContext implements the IMContext class.
type imContext struct {
	gextras.Objector
}

// WrapIMContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMContext(obj *externglib.Object) IMContext {
	return imContext{
		Objector: obj,
	}
}

func marshalIMContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMContext(obj), nil
}

func (c imContext) DeleteSurroundingIMContext(offset int, nChars int) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gint          // out
	var _arg2 C.gint          // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(offset)
	_arg2 = C.gint(nChars)

	_cret = C.gtk_im_context_delete_surrounding(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c imContext) FilterKeypressIMContext(event *gdk.EventKey) bool {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkEventKey  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkEventKey)(unsafe.Pointer(event.Native()))

	_cret = C.gtk_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c imContext) FocusInIMContext() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_in(_arg0)
}

func (c imContext) FocusOutIMContext() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_focus_out(_arg0)
}

func (c imContext) PreeditString() (string, *pango.AttrList, int) {
	var _arg0 *C.GtkIMContext  // out
	var _arg1 *C.gchar         // in
	var _arg2 *C.PangoAttrList // in
	var _arg3 C.gint           // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_get_preedit_string(_arg0, &_arg1, &_arg2, &_arg3)

	var _str string            // out
	var _attrs *pango.AttrList // out
	var _cursorPos int         // out

	_str = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_attrs = (*pango.AttrList)(unsafe.Pointer(_arg2))
	runtime.SetFinalizer(&_attrs, func(v **pango.AttrList) {
		C.free(unsafe.Pointer(v))
	})
	_cursorPos = int(_arg3)

	return _str, _attrs, _cursorPos
}

func (c imContext) Surrounding() (string, int, bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.gchar        // in
	var _arg2 C.gint          // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_im_context_get_surrounding(_arg0, &_arg1, &_arg2)

	var _text string     // out
	var _cursorIndex int // out
	var _ok bool         // out

	_text = C.GoString(_arg1)
	defer C.free(unsafe.Pointer(_arg1))
	_cursorIndex = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _text, _cursorIndex, _ok
}

func (c imContext) ResetIMContext() {
	var _arg0 *C.GtkIMContext // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))

	C.gtk_im_context_reset(_arg0)
}

func (c imContext) SetClientWindowIMContext(window gdk.Window) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkWindow    // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_im_context_set_client_window(_arg0, _arg1)
}

func (c imContext) SetCursorLocationIMContext(area *gdk.Rectangle) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(area.Native()))

	C.gtk_im_context_set_cursor_location(_arg0, _arg1)
}

func (c imContext) SetSurroundingIMContext(text string, len int, cursorIndex int) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gint          // out
	var _arg3 C.gint          // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(len)
	_arg3 = C.gint(cursorIndex)

	C.gtk_im_context_set_surrounding(_arg0, _arg1, _arg2, _arg3)
}

func (c imContext) SetUsePreeditIMContext(usePreedit bool) {
	var _arg0 *C.GtkIMContext // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkIMContext)(unsafe.Pointer(c.Native()))
	if usePreedit {
		_arg1 = C.TRUE
	}

	C.gtk_im_context_set_use_preedit(_arg0, _arg1)
}