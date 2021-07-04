// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_tearoff_menu_item_get_type()), F: marshalTearoffMenuItem},
	})
}

// TearoffMenuItem: a TearoffMenuItem is a special MenuItem which is used to
// tear off and reattach its menu.
//
// When its menu is shown normally, the TearoffMenuItem is drawn as a dotted
// line indicating that the menu can be torn off. Activating it causes its menu
// to be torn off and displayed in its own window as a tearoff menu.
//
// When its menu is shown as a tearoff menu, the TearoffMenuItem is drawn as a
// dotted line which has a left pointing arrow graphic indicating that the
// tearoff menu can be reattached. Activating it will erase the tearoff menu
// window.
//
// > TearoffMenuItem is deprecated and should not be used in newly > written
// code. Menus are not meant to be torn around.
type TearoffMenuItem interface {
	MenuItem
}

// tearoffMenuItem implements the TearoffMenuItem class.
type tearoffMenuItem struct {
	MenuItem
}

// WrapTearoffMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapTearoffMenuItem(obj *externglib.Object) TearoffMenuItem {
	return tearoffMenuItem{
		MenuItem: WrapMenuItem(obj),
	}
}

func marshalTearoffMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTearoffMenuItem(obj), nil
}

// NewTearoffMenuItem:
func NewTearoffMenuItem() TearoffMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tearoff_menu_item_new()

	var _tearoffMenuItem TearoffMenuItem // out

	_tearoffMenuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TearoffMenuItem)

	return _tearoffMenuItem
}

func (b tearoffMenuItem) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b tearoffMenuItem) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b tearoffMenuItem) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b tearoffMenuItem) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b tearoffMenuItem) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b tearoffMenuItem) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b tearoffMenuItem) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a tearoffMenuItem) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a tearoffMenuItem) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a tearoffMenuItem) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a tearoffMenuItem) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a tearoffMenuItem) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b tearoffMenuItem) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b tearoffMenuItem) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b tearoffMenuItem) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b tearoffMenuItem) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b tearoffMenuItem) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b tearoffMenuItem) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b tearoffMenuItem) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a tearoffMenuItem) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a tearoffMenuItem) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a tearoffMenuItem) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a tearoffMenuItem) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a tearoffMenuItem) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a tearoffMenuItem) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}
