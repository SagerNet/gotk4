// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButtonner},
	})
}

// ToggleButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToggleButtonOverrider interface {
	// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There
	// is no good reason for an application ever to call this function.
	Toggled()
}

// ToggleButton is a Button which will remain “pressed-in” when clicked.
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either gtk_toggle_button_new() or
// gtk_toggle_button_new_with_label(). If using the former, it is advisable to
// pack a widget, (such as a Label and/or a Image), into the toggle button’s
// container. (See Button for more information).
//
// The state of a ToggleButton can be set specifically using
// gtk_toggle_button_set_active(), and retrieved using
// gtk_toggle_button_get_active().
//
// To simply switch the state of a toggle button, use
// gtk_toggle_button_toggled().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .toggle style class.
//
// Creating two ToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data) {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void) {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      // Makes this toggle button invisible
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle1),
//                                  TRUE);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle2),
//                                  FALSE);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle2);
//
//      gtk_container_add (GTK_CONTAINER (window), box);
//      gtk_widget_show_all (window);
//    }
type ToggleButton struct {
	Button
}

func WrapToggleButton(obj *externglib.Object) *ToggleButton {
	return &ToggleButton{
		Button: Button{
			Bin: Bin{
				Container: Container{
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
			Actionable: Actionable{
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
			Activatable: Activatable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalToggleButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleButton(obj), nil
}

// NewToggleButton creates a new toggle button. A widget should be packed into
// the button, as in gtk_button_new().
func NewToggleButton() *ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton *ToggleButton // out

	_toggleButton = WrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithLabel creates a new toggle button with a text label.
func NewToggleButtonWithLabel(label string) *ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = WrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithMnemonic creates a new ToggleButton containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the button.
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = WrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// Active queries a ToggleButton and returns its current state. Returns TRUE if
// the toggle button is pressed in and FALSE if it is raised.
func (toggleButton *ToggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inconsistent gets the value set by gtk_toggle_button_set_inconsistent().
func (toggleButton *ToggleButton) Inconsistent() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	_cret = C.gtk_toggle_button_get_inconsistent(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Mode retrieves whether the button is displayed as a separate indicator and
// label. See gtk_toggle_button_set_mode().
func (toggleButton *ToggleButton) Mode() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	_cret = C.gtk_toggle_button_get_mode(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle button. Set to TRUE if you want the
// GtkToggleButton to be “pressed in”, and FALSE to raise it. This action causes
// the ToggleButton::toggled signal and the Button::clicked signal to be
// emitted.
func (toggleButton *ToggleButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(isActive)
}

// SetInconsistent: if the user has selected a range of elements (such as some
// text or spreadsheet cells) that are affected by a toggle button, and the
// current values in that range are inconsistent, you may want to display the
// toggle in an “in between” state. This function turns on “in between” display.
// Normally you would turn off the inconsistent state again if the user toggles
// the toggle button. This has to be done manually,
// gtk_toggle_button_set_inconsistent() only affects visual appearance, it
// doesn’t affect the semantics of the button.
func (toggleButton *ToggleButton) SetInconsistent(setting bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_inconsistent(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(setting)
}

// SetMode sets whether the button is displayed as a separate indicator and
// label. You can call this function on a checkbutton or a radiobutton with
// draw_indicator = FALSE to make the button look like a normal button.
//
// This can be used to create linked strip of buttons that work like a
// StackSwitcher.
//
// This function only affects instances of classes like CheckButton and
// RadioButton that derive from ToggleButton, not instances of ToggleButton
// itself.
func (toggleButton *ToggleButton) SetMode(drawIndicator bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if drawIndicator {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_mode(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(drawIndicator)
}

// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There is
// no good reason for an application ever to call this function.
func (toggleButton *ToggleButton) Toggled() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	C.gtk_toggle_button_toggled(_arg0)
	runtime.KeepAlive(toggleButton)
}
