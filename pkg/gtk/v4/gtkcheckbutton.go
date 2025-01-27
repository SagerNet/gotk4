// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_button_get_type()), F: marshalCheckButtonner},
	})
}

// CheckButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CheckButtonOverrider interface {
	Activate()
	Toggled()
}

// CheckButton: GtkCheckButton places a label next to an indicator.
//
// !Example GtkCheckButtons (check-button.png)
//
// A GtkCheckButton is created by calling either gtk.CheckButton.New or
// gtk.CheckButton.NewWithLabel.
//
// The state of a GtkCheckButton can be set specifically using
// gtk.CheckButton.SetActive(), and retrieved using gtk.CheckButton.GetActive().
//
//
// Inconsistent state
//
// In addition to "on" and "off", check buttons can be an "in between" state
// that is neither on nor off. This can be used e.g. when the user has selected
// a range of elements (such as some text or spreadsheet cells) that are
// affected by a check button, and the current values in that range are
// inconsistent.
//
// To set a GtkCheckButton to inconsistent state, use
// gtk.CheckButton.SetInconsistent().
//
//
// Grouping
//
// Check buttons can be grouped together, to form mutually exclusive groups -
// only one of the buttons can be toggled at a time, and toggling another one
// will switch the currently toggled one off.
//
// Grouped check buttons use a different indicator, and are commonly referred to
// as *radio buttons*.
//
// !Example GtkCheckButtons (radio-button.png)
//
// To add a GtkCheckButton to a group, use gtk.CheckButton.SetGroup().
//
// CSS nodes
//
//    checkbutton[.text-button]
//    ├── check
//    ╰── [label]
//
//
// A GtkCheckButton has a main node with name checkbutton. If the
// gtk.CheckButton:label property is set, it contains a label child. The
// indicator node is named check when no group is set, and radio if the
// checkbutton is grouped together with other checkbuttons.
//
//
// Accessibility
//
// GtkCheckButton uses the GTK_ACCESSIBLE_ROLE_CHECKBOX role.
type CheckButton struct {
	Widget

	Actionable
	*externglib.Object
}

func WrapCheckButton(obj *externglib.Object) *CheckButton {
	return &CheckButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Actionable: Actionable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalCheckButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCheckButton(obj), nil
}

// NewCheckButton creates a new GtkCheckButton.
func NewCheckButton() *CheckButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_check_button_new()

	var _checkButton *CheckButton // out

	_checkButton = WrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithLabel creates a new GtkCheckButton with the given text.
func NewCheckButtonWithLabel(label string) *CheckButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_check_button_new_with_label(_arg1)
	runtime.KeepAlive(label)

	var _checkButton *CheckButton // out

	_checkButton = WrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// NewCheckButtonWithMnemonic creates a new GtkCheckButton with the given text
// and a mnemonic.
func NewCheckButtonWithMnemonic(label string) *CheckButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_check_button_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _checkButton *CheckButton // out

	_checkButton = WrapCheckButton(externglib.Take(unsafe.Pointer(_cret)))

	return _checkButton
}

// Active returns whether the check button is active.
func (self *CheckButton) Active() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_check_button_get_active(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inconsistent returns whether the check button is in an inconsistent state.
func (checkButton *CheckButton) Inconsistent() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(checkButton.Native()))

	_cret = C.gtk_check_button_get_inconsistent(_arg0)
	runtime.KeepAlive(checkButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label returns the label of the check button.
func (self *CheckButton) Label() string {
	var _arg0 *C.GtkCheckButton // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_check_button_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline returns whether underlines in the label indicate mnemonics.
func (self *CheckButton) UseUnderline() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_check_button_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive changes the check buttons active state.
func (self *CheckButton) SetActive(setting bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetGroup adds self to the group of group.
//
// In a group of multiple check buttons, only one button can be active at a
// time. The behavior of a checkbutton in a group is also commonly known as a
// *radio button*.
//
// Setting the group of a check button also changes the css name of the
// indicator widget's CSS node to 'radio'.
//
// Setting up groups in a cycle leads to undefined behavior.
//
// Note that the same effect can be achieved via the gtk.Actionable API, by
// using the same action with parameter type and state type 's' for all buttons
// in the group, and giving each button its own target value.
func (self *CheckButton) SetGroup(group *CheckButton) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 *C.GtkCheckButton // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))
	if group != nil {
		_arg1 = (*C.GtkCheckButton)(unsafe.Pointer(group.Native()))
	}

	C.gtk_check_button_set_group(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(group)
}

// SetInconsistent sets the GtkCheckButton to inconsistent state.
//
// You shoud turn off the inconsistent state again if the user checks the check
// button. This has to be done manually.
func (checkButton *CheckButton) SetInconsistent(inconsistent bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(checkButton.Native()))
	if inconsistent {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_inconsistent(_arg0, _arg1)
	runtime.KeepAlive(checkButton)
	runtime.KeepAlive(inconsistent)
}

// SetLabel sets the text of self.
//
// If gtk.CheckButton:use-underline is TRUE, an underscore in label is
// interpreted as mnemonic indicator, see gtk.CheckButton.SetUseUnderline() for
// details on this behavior.
func (self *CheckButton) SetLabel(label string) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))
	if label != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_check_button_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(label)
}

// SetUseUnderline sets whether underlines in the label indicate mnemonics.
//
// If setting is TRUE, an underscore character in self's label indicates a
// mnemonic accelerator key. This behavior is similar to
// gtk.Label:use-underline.
func (self *CheckButton) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}
