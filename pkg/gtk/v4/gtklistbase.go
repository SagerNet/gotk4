// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_list_base_get_type()), F: marshalListBase},
	})
}

// ListBase: `GtkListBase` is the abstract base class for GTK's list widgets.
type ListBase interface {
	Widget
	Orientable
	Scrollable
}

// listBase implements the ListBase class.
type listBase struct {
	Widget
}

// WrapListBase wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBase(obj *externglib.Object) ListBase {
	return listBase{
		Widget: WrapWidget(obj),
	}
}

func marshalListBase(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBase(obj), nil
}

func (s listBase) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s listBase) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s listBase) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s listBase) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s listBase) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s listBase) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s listBase) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b listBase) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o listBase) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o listBase) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}

func (s listBase) Border() (Border, bool) {
	return WrapScrollable(gextras.InternObject(s)).Border()
}

func (s listBase) HAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).HAdjustment()
}

func (s listBase) HScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).HScrollPolicy()
}

func (s listBase) VAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).VAdjustment()
}

func (s listBase) VScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).VScrollPolicy()
}

func (s listBase) SetHAdjustment(hadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetHAdjustment(hadjustment)
}

func (s listBase) SetHScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetHScrollPolicy(policy)
}

func (s listBase) SetVAdjustment(vadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetVAdjustment(vadjustment)
}

func (s listBase) SetVScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetVScrollPolicy(policy)
}