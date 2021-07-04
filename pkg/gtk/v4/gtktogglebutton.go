// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButton},
	})
}

// ToggleButton: a `GtkToggleButton` is a button which remains “pressed-in” when
// clicked.
//
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either [ctor@Gtk.ToggleButton.new] or
// [ctor@Gtk.ToggleButton.new_with_label]. If using the former, it is advisable
// to pack a widget, (such as a `GtkLabel` and/or a `GtkImage`), into the toggle
// button’s container. (See [class@Gtk.Button] for more information).
//
// The state of a `GtkToggleButton` can be set specifically using
// [method@Gtk.ToggleButton.set_active], and retrieved using
// [method@Gtk.ToggleButton.get_active].
//
// To simply switch the state of a toggle button, use
// [method@Gtk.ToggleButton.toggled].
//
//
// Grouping
//
// Toggle buttons can be grouped together, to form mutually exclusive groups -
// only one of the buttons can be toggled at a time, and toggling another one
// will switch the currently toggled one off.
//
// To add a `GtkToggleButton` to a group, use
// [method@Gtk.ToggleButton.set_group].
//
//
// CSS nodes
//
// `GtkToggleButton` has a single CSS node with name button. To differentiate it
// from a plain `GtkButton`, it gets the .toggle style class.
//
// Creating two `GtkToggleButton` widgets.
//
// “`c static void output_state (GtkToggleButton *source, gpointer user_data) {
// printf ("Active: d\n", gtk_toggle_button_get_active (source)); }
//
// void make_toggles (void) { GtkWidget *window, *toggle1, *toggle2; GtkWidget
// *box; const char *text;
//
//    window = gtk_window_new ();
//    box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//    text = "Hi, I’m a toggle button.";
//    toggle1 = gtk_toggle_button_new_with_label (text);
//
//    g_signal_connect (toggle1, "toggled",
//                      G_CALLBACK (output_state),
//                      NULL);
//    gtk_box_append (GTK_BOX (box), toggle1);
//
//    text = "Hi, I’m a toggle button.";
//    toggle2 = gtk_toggle_button_new_with_label (text);
//    g_signal_connect (toggle2, "toggled",
//                      G_CALLBACK (output_state),
//                      NULL);
//    gtk_box_append (GTK_BOX (box), toggle2);
//
//    gtk_window_set_child (GTK_WINDOW (window), box);
//    gtk_widget_show (window);
//
// } “`
type ToggleButton interface {
	Button

	Active() bool

	SetActiveToggleButton(isActive bool)

	SetGroupToggleButton(group ToggleButton)

	ToggledToggleButton()
}

// toggleButton implements the ToggleButton class.
type toggleButton struct {
	Button
}

// WrapToggleButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleButton(obj *externglib.Object) ToggleButton {
	return toggleButton{
		Button: WrapButton(obj),
	}
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleButton(obj), nil
}

func NewToggleButton() ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func NewToggleButtonWithLabel(label string) ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func NewToggleButtonWithMnemonic(label string) ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func (t toggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toggleButton) SetActiveToggleButton(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
}

func (t toggleButton) SetGroupToggleButton(group ToggleButton) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkToggleButton)(unsafe.Pointer(group.Native()))

	C.gtk_toggle_button_set_group(_arg0, _arg1)
}

func (t toggleButton) ToggledToggleButton() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	C.gtk_toggle_button_toggled(_arg0)
}

func (a toggleButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a toggleButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a toggleButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a toggleButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a toggleButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (s toggleButton) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s toggleButton) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s toggleButton) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s toggleButton) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s toggleButton) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s toggleButton) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s toggleButton) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b toggleButton) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}