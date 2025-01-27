// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_accel_label_get_type()), F: marshalAccelLabeller},
	})
}

// AccelLabel widget is a subclass of Label that also displays an accelerator
// key on the right of the label text, e.g. “Ctrl+S”. It is commonly used in
// menus to show the keyboard short-cuts for commands.
//
// The accelerator key to display is typically not set explicitly (although it
// can be, with gtk_accel_label_set_accel()). Instead, the AccelLabel displays
// the accelerators which have been added to a particular widget. This widget is
// set by calling gtk_accel_label_set_accel_widget().
//
// For example, a MenuItem widget may have an accelerator added to emit the
// “activate” signal when the “Ctrl+S” key combination is pressed. A AccelLabel
// is created and added to the MenuItem, and gtk_accel_label_set_accel_widget()
// is called with the MenuItem as the second argument. The AccelLabel will now
// display “Ctrl+S” after its label.
//
// Note that creating a MenuItem with gtk_menu_item_new_with_label() (or one of
// the similar functions for CheckMenuItem and RadioMenuItem) automatically adds
// a AccelLabel to the MenuItem and calls gtk_accel_label_set_accel_widget() to
// set it up for you.
//
// A AccelLabel will only display accelerators which have GTK_ACCEL_VISIBLE set
// (see AccelFlags). A AccelLabel can display multiple accelerators and even
// signal names, though it is almost always used to display just one accelerator
// key.
//
// Creating a simple menu item with an accelerator key.
//
//    label
//    ╰── accelerator
//
// Like Label, GtkAccelLabel has a main CSS node with the name label. It adds a
// subnode with name accelerator.
type AccelLabel struct {
	Label
}

func WrapAccelLabel(obj *externglib.Object) *AccelLabel {
	return &AccelLabel{
		Label: Label{
			Misc: Misc{
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
			},
		},
	}
}

func marshalAccelLabeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAccelLabel(obj), nil
}

// NewAccelLabel creates a new AccelLabel.
func NewAccelLabel(_string string) *AccelLabel {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_accel_label_new(_arg1)
	runtime.KeepAlive(_string)

	var _accelLabel *AccelLabel // out

	_accelLabel = WrapAccelLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _accelLabel
}

// Accel gets the keyval and modifier mask set with gtk_accel_label_set_accel().
func (accelLabel *AccelLabel) Accel() (uint, gdk.ModifierType) {
	var _arg0 *C.GtkAccelLabel  // out
	var _arg1 C.guint           // in
	var _arg2 C.GdkModifierType // in

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))

	C.gtk_accel_label_get_accel(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(accelLabel)

	var _acceleratorKey uint              // out
	var _acceleratorMods gdk.ModifierType // out

	_acceleratorKey = uint(_arg1)
	_acceleratorMods = gdk.ModifierType(_arg2)

	return _acceleratorKey, _acceleratorMods
}

// AccelWidget fetches the widget monitored by this accelerator label. See
// gtk_accel_label_set_accel_widget().
func (accelLabel *AccelLabel) AccelWidget() Widgetter {
	var _arg0 *C.GtkAccelLabel // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))

	_cret = C.gtk_accel_label_get_accel_widget(_arg0)
	runtime.KeepAlive(accelLabel)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// AccelWidth returns the width needed to display the accelerator key(s). This
// is used by menus to align all of the MenuItem widgets, and shouldn't be
// needed by applications.
func (accelLabel *AccelLabel) AccelWidth() uint {
	var _arg0 *C.GtkAccelLabel // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))

	_cret = C.gtk_accel_label_get_accel_width(_arg0)
	runtime.KeepAlive(accelLabel)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Refetch recreates the string representing the accelerator keys. This should
// not be needed since the string is automatically updated whenever accelerators
// are added or removed from the associated widget.
func (accelLabel *AccelLabel) Refetch() bool {
	var _arg0 *C.GtkAccelLabel // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))

	_cret = C.gtk_accel_label_refetch(_arg0)
	runtime.KeepAlive(accelLabel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAccel: manually sets a keyval and modifier mask as the accelerator
// rendered by accel_label.
//
// If a keyval and modifier are explicitly set then these values are used
// regardless of any associated accel closure or widget.
//
// Providing an accelerator_key of 0 removes the manual setting.
func (accelLabel *AccelLabel) SetAccel(acceleratorKey uint, acceleratorMods gdk.ModifierType) {
	var _arg0 *C.GtkAccelLabel  // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))
	_arg1 = C.guint(acceleratorKey)
	_arg2 = C.GdkModifierType(acceleratorMods)

	C.gtk_accel_label_set_accel(_arg0, _arg1, _arg2)
	runtime.KeepAlive(accelLabel)
	runtime.KeepAlive(acceleratorKey)
	runtime.KeepAlive(acceleratorMods)
}

// SetAccelWidget sets the widget to be monitored by this accelerator label.
// Passing NULL for accel_widget will dissociate accel_label from its current
// widget, if any.
func (accelLabel *AccelLabel) SetAccelWidget(accelWidget Widgetter) {
	var _arg0 *C.GtkAccelLabel // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel.Native()))
	if accelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(accelWidget.Native()))
	}

	C.gtk_accel_label_set_accel_widget(_arg0, _arg1)
	runtime.KeepAlive(accelLabel)
	runtime.KeepAlive(accelWidget)
}
