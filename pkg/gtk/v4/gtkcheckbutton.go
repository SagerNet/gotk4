// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_button_get_type()), F: marshalCheckButton},
	})
}

// CheckButton: a `GtkCheckButton` places a label next to an indicator.
//
// !Example GtkCheckButtons (check-button.png)
//
// A `GtkCheckButton` is created by calling either [ctor@Gtk.CheckButton.new] or
// [ctor@Gtk.CheckButton.new_with_label].
//
// The state of a `GtkCheckButton` can be set specifically using
// [method@Gtk.CheckButton.set_active], and retrieved using
// [method@Gtk.CheckButton.get_active].
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
// To set a `GtkCheckButton` to inconsistent state, use
// [method@Gtk.CheckButton.set_inconsistent].
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
// To add a `GtkCheckButton` to a group, use [method@Gtk.CheckButton.set_group].
//
//
// CSS nodes
//
// “` checkbutton[.text-button] ├── check ╰── [label] “`
//
// A `GtkCheckButton` has a main node with name checkbutton. If the
// [property@Gtk.CheckButton:label] property is set, it contains a label child.
// The indicator node is named check when no group is set, and radio if the
// checkbutton is grouped together with other checkbuttons.
//
//
// Accessibility
//
// `GtkCheckButton` uses the GTK_ACCESSIBLE_ROLE_CHECKBOX role.
type CheckButton interface {
	Actionable

	// Active:
	Active() bool
	// Inconsistent:
	Inconsistent() bool
	// Label:
	Label() string
	// UseUnderline:
	UseUnderline() bool
	// SetActiveCheckButton:
	SetActiveCheckButton(setting bool)
	// SetGroupCheckButton:
	SetGroupCheckButton(group CheckButton)
	// SetInconsistentCheckButton:
	SetInconsistentCheckButton(inconsistent bool)
	// SetLabelCheckButton:
	SetLabelCheckButton(label string)
	// SetUseUnderlineCheckButton:
	SetUseUnderlineCheckButton(setting bool)
}

// checkButton implements the CheckButton class.
type checkButton struct {
	Widget
}

// WrapCheckButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapCheckButton(obj *externglib.Object) CheckButton {
	return checkButton{
		Widget: WrapWidget(obj),
	}
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCheckButton(obj), nil
}

// NewCheckButton:
func NewCheckButton() CheckButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_check_button_new()

	var _checkButton CheckButton // out

	_checkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CheckButton)

	return _checkButton
}

// NewCheckButtonWithLabel:
func NewCheckButtonWithLabel(label string) CheckButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_check_button_new_with_label(_arg1)

	var _checkButton CheckButton // out

	_checkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CheckButton)

	return _checkButton
}

// NewCheckButtonWithMnemonic:
func NewCheckButtonWithMnemonic(label string) CheckButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_check_button_new_with_mnemonic(_arg1)

	var _checkButton CheckButton // out

	_checkButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CheckButton)

	return _checkButton
}

func (s checkButton) Active() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_check_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c checkButton) Inconsistent() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_check_button_get_inconsistent(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s checkButton) Label() string {
	var _arg0 *C.GtkCheckButton // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_check_button_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s checkButton) UseUnderline() bool {
	var _arg0 *C.GtkCheckButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_check_button_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s checkButton) SetActiveCheckButton(setting bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_active(_arg0, _arg1)
}

func (s checkButton) SetGroupCheckButton(group CheckButton) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 *C.GtkCheckButton // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkCheckButton)(unsafe.Pointer(group.Native()))

	C.gtk_check_button_set_group(_arg0, _arg1)
}

func (c checkButton) SetInconsistentCheckButton(inconsistent bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(c.Native()))
	if inconsistent {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_inconsistent(_arg0, _arg1)
}

func (s checkButton) SetLabelCheckButton(label string) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_check_button_set_label(_arg0, _arg1)
}

func (s checkButton) SetUseUnderlineCheckButton(setting bool) {
	var _arg0 *C.GtkCheckButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_check_button_set_use_underline(_arg0, _arg1)
}

func (a checkButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a checkButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a checkButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a checkButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a checkButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (s checkButton) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s checkButton) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s checkButton) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s checkButton) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s checkButton) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s checkButton) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s checkButton) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b checkButton) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
