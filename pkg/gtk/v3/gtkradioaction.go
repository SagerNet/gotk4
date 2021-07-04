// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_radio_action_get_type()), F: marshalRadioAction},
	})
}

// RadioAction: a RadioAction is similar to RadioMenuItem. A number of radio
// actions can be linked together so that only one may be active at any one
// time.
type RadioAction interface {
	ToggleAction

	CurrentValue() int

	JoinGroupRadioAction(groupSource RadioAction)

	SetCurrentValueRadioAction(currentValue int)
}

// radioAction implements the RadioAction class.
type radioAction struct {
	ToggleAction
}

// WrapRadioAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioAction(obj *externglib.Object) RadioAction {
	return radioAction{
		ToggleAction: WrapToggleAction(obj),
	}
}

func marshalRadioAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioAction(obj), nil
}

func NewRadioAction(name string, label string, tooltip string, stockId string, value int) RadioAction {
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.gchar          // out
	var _arg4 *C.gchar          // out
	var _arg5 C.gint            // out
	var _cret *C.GtkRadioAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gint(value)

	_cret = C.gtk_radio_action_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _radioAction RadioAction // out

	_radioAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RadioAction)

	return _radioAction
}

func (a radioAction) CurrentValue() int {
	var _arg0 *C.GtkRadioAction // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_radio_action_get_current_value(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a radioAction) JoinGroupRadioAction(groupSource RadioAction) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 *C.GtkRadioAction // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkRadioAction)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_action_join_group(_arg0, _arg1)
}

func (a radioAction) SetCurrentValueRadioAction(currentValue int) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(currentValue)

	C.gtk_radio_action_set_current_value(_arg0, _arg1)
}

func (b radioAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioAction) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b radioAction) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b radioAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}