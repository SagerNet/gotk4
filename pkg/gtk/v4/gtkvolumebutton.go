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
		{T: externglib.Type(C.gtk_volume_button_get_type()), F: marshalVolumeButton},
	})
}

// VolumeButton: `GtkVolumeButton` is a `GtkScaleButton` subclass tailored for
// volume control.
//
// !An example GtkVolumeButton (volumebutton.png)
type VolumeButton interface {
	ScaleButton
}

// volumeButton implements the VolumeButton class.
type volumeButton struct {
	ScaleButton
}

// WrapVolumeButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapVolumeButton(obj *externglib.Object) VolumeButton {
	return volumeButton{
		ScaleButton: WrapScaleButton(obj),
	}
}

func marshalVolumeButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVolumeButton(obj), nil
}

func NewVolumeButton() VolumeButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_volume_button_new()

	var _volumeButton VolumeButton // out

	_volumeButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(VolumeButton)

	return _volumeButton
}

func (s volumeButton) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s volumeButton) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s volumeButton) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s volumeButton) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s volumeButton) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s volumeButton) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s volumeButton) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b volumeButton) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o volumeButton) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o volumeButton) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}