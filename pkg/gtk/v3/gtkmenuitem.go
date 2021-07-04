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
		{T: externglib.Type(C.gtk_menu_item_get_type()), F: marshalMenuItem},
	})
}

// MenuItem: the MenuItem widget and the derived widgets are the only valid
// children for menus. Their function is to correctly handle highlighting,
// alignment, events and submenus.
//
// As a GtkMenuItem derives from Bin it can hold any valid child widget,
// although only a few are really useful.
//
// By default, a GtkMenuItem sets a AccelLabel as its child. GtkMenuItem has
// direct functions to set the label and its mnemonic. For more advanced label
// settings, you can fetch the child widget from the GtkBin.
//
// An example for setting markup and accelerator on a MenuItem:
//
//    menuitem
//    ├── <child>
//    ╰── [arrow.right]
//
// GtkMenuItem has a single CSS node with name menuitem. If the menuitem has a
// submenu, it gets another CSS node with name arrow, which has the .left or
// .right style class.
type MenuItem interface {
	Bin
	Actionable
	Activatable

	// ActivateMenuItem:
	ActivateMenuItem()
	// DeselectMenuItem:
	DeselectMenuItem()
	// AccelPath:
	AccelPath() string
	// Label:
	Label() string
	// ReserveIndicator:
	ReserveIndicator() bool
	// RightJustified:
	RightJustified() bool
	// Submenu:
	Submenu() Widget
	// UseUnderline:
	UseUnderline() bool
	// SelectMenuItem:
	SelectMenuItem()
	// SetAccelPathMenuItem:
	SetAccelPathMenuItem(accelPath string)
	// SetLabelMenuItem:
	SetLabelMenuItem(label string)
	// SetReserveIndicatorMenuItem:
	SetReserveIndicatorMenuItem(reserve bool)
	// SetRightJustifiedMenuItem:
	SetRightJustifiedMenuItem(rightJustified bool)
	// SetSubmenuMenuItem:
	SetSubmenuMenuItem(submenu Menu)
	// SetUseUnderlineMenuItem:
	SetUseUnderlineMenuItem(setting bool)
	// ToggleSizeAllocateMenuItem:
	ToggleSizeAllocateMenuItem(allocation int)
}

// menuItem implements the MenuItem class.
type menuItem struct {
	Bin
}

// WrapMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuItem(obj *externglib.Object) MenuItem {
	return menuItem{
		Bin: WrapBin(obj),
	}
}

func marshalMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuItem(obj), nil
}

// NewMenuItem:
func NewMenuItem() MenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_item_new()

	var _menuItem MenuItem // out

	_menuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuItem)

	return _menuItem
}

// NewMenuItemWithLabel:
func NewMenuItemWithLabel(label string) MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_label(_arg1)

	var _menuItem MenuItem // out

	_menuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuItem)

	return _menuItem
}

// NewMenuItemWithMnemonic:
func NewMenuItemWithMnemonic(label string) MenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_item_new_with_mnemonic(_arg1)

	var _menuItem MenuItem // out

	_menuItem = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MenuItem)

	return _menuItem
}

func (m menuItem) ActivateMenuItem() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	C.gtk_menu_item_activate(_arg0)
}

func (m menuItem) DeselectMenuItem() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	C.gtk_menu_item_deselect(_arg0)
}

func (m menuItem) AccelPath() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_accel_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (m menuItem) Label() string {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (m menuItem) ReserveIndicator() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_reserve_indicator(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m menuItem) RightJustified() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_right_justified(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m menuItem) Submenu() Widget {
	var _arg0 *C.GtkMenuItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_submenu(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (m menuItem) UseUnderline() bool {
	var _arg0 *C.GtkMenuItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	_cret = C.gtk_menu_item_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m menuItem) SelectMenuItem() {
	var _arg0 *C.GtkMenuItem // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))

	C.gtk_menu_item_select(_arg0)
}

func (m menuItem) SetAccelPathMenuItem(accelPath string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_item_set_accel_path(_arg0, _arg1)
}

func (m menuItem) SetLabelMenuItem(label string) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_item_set_label(_arg0, _arg1)
}

func (m menuItem) SetReserveIndicatorMenuItem(reserve bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	if reserve {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_reserve_indicator(_arg0, _arg1)
}

func (m menuItem) SetRightJustifiedMenuItem(rightJustified bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	if rightJustified {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_right_justified(_arg0, _arg1)
}

func (m menuItem) SetSubmenuMenuItem(submenu Menu) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(submenu.Native()))

	C.gtk_menu_item_set_submenu(_arg0, _arg1)
}

func (m menuItem) SetUseUnderlineMenuItem(setting bool) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_menu_item_set_use_underline(_arg0, _arg1)
}

func (m menuItem) ToggleSizeAllocateMenuItem(allocation int) {
	var _arg0 *C.GtkMenuItem // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkMenuItem)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(allocation)

	C.gtk_menu_item_toggle_size_allocate(_arg0, _arg1)
}

func (b menuItem) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b menuItem) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b menuItem) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b menuItem) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b menuItem) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b menuItem) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b menuItem) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a menuItem) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a menuItem) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a menuItem) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a menuItem) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a menuItem) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b menuItem) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b menuItem) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b menuItem) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b menuItem) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b menuItem) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b menuItem) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b menuItem) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a menuItem) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a menuItem) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a menuItem) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a menuItem) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a menuItem) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a menuItem) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}
