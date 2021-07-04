// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_recent_action_get_type()), F: marshalRecentAction},
	})
}

// RecentAction: a RecentAction represents a list of recently used files, which
// can be shown by widgets such as RecentChooserDialog or RecentChooserMenu.
//
// To construct a submenu showing recently used files, use a RecentAction as the
// action for a <menuitem>. To construct a menu toolbutton showing the recently
// used files in the popup menu, use a RecentAction as the action for a
// <toolitem> element.
type RecentAction interface {
	Action
	RecentChooser

	// ShowNumbers:
	ShowNumbers() bool
	// SetShowNumbersRecentAction:
	SetShowNumbersRecentAction(showNumbers bool)
}

// recentAction implements the RecentAction class.
type recentAction struct {
	Action
}

// WrapRecentAction wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentAction(obj *externglib.Object) RecentAction {
	return recentAction{
		Action: WrapAction(obj),
	}
}

func marshalRecentAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentAction(obj), nil
}

// NewRecentAction:
func NewRecentAction(name string, label string, tooltip string, stockId string) RecentAction {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_recent_action_new(_arg1, _arg2, _arg3, _arg4)

	var _recentAction RecentAction // out

	_recentAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RecentAction)

	return _recentAction
}

// NewRecentActionForManager:
func NewRecentActionForManager(name string, label string, tooltip string, stockId string, manager RecentManager) RecentAction {
	var _arg1 *C.gchar            // out
	var _arg2 *C.gchar            // out
	var _arg3 *C.gchar            // out
	var _arg4 *C.gchar            // out
	var _arg5 *C.GtkRecentManager // out
	var _cret *C.GtkAction        // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_action_new_for_manager(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _recentAction RecentAction // out

	_recentAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RecentAction)

	return _recentAction
}

func (a recentAction) ShowNumbers() bool {
	var _arg0 *C.GtkRecentAction // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_recent_action_get_show_numbers(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a recentAction) SetShowNumbersRecentAction(showNumbers bool) {
	var _arg0 *C.GtkRecentAction // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkRecentAction)(unsafe.Pointer(a.Native()))
	if showNumbers {
		_arg1 = C.TRUE
	}

	C.gtk_recent_action_set_show_numbers(_arg0, _arg1)
}

func (b recentAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b recentAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b recentAction) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b recentAction) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b recentAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b recentAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b recentAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c recentAction) AddFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).AddFilter(filter)
}

func (c recentAction) CurrentItem() *RecentInfo {
	return WrapRecentChooser(gextras.InternObject(c)).CurrentItem()
}

func (c recentAction) CurrentURI() string {
	return WrapRecentChooser(gextras.InternObject(c)).CurrentURI()
}

func (c recentAction) Filter() RecentFilter {
	return WrapRecentChooser(gextras.InternObject(c)).Filter()
}

func (c recentAction) Limit() int {
	return WrapRecentChooser(gextras.InternObject(c)).Limit()
}

func (c recentAction) LocalOnly() bool {
	return WrapRecentChooser(gextras.InternObject(c)).LocalOnly()
}

func (c recentAction) SelectMultiple() bool {
	return WrapRecentChooser(gextras.InternObject(c)).SelectMultiple()
}

func (c recentAction) ShowIcons() bool {
	return WrapRecentChooser(gextras.InternObject(c)).ShowIcons()
}

func (c recentAction) ShowNotFound() bool {
	return WrapRecentChooser(gextras.InternObject(c)).ShowNotFound()
}

func (c recentAction) ShowPrivate() bool {
	return WrapRecentChooser(gextras.InternObject(c)).ShowPrivate()
}

func (c recentAction) ShowTips() bool {
	return WrapRecentChooser(gextras.InternObject(c)).ShowTips()
}

func (c recentAction) SortType() RecentSortType {
	return WrapRecentChooser(gextras.InternObject(c)).SortType()
}

func (c recentAction) RemoveFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).RemoveFilter(filter)
}

func (c recentAction) SelectAll() {
	WrapRecentChooser(gextras.InternObject(c)).SelectAll()
}

func (c recentAction) SelectURI(uri string) error {
	return WrapRecentChooser(gextras.InternObject(c)).SelectURI(uri)
}

func (c recentAction) SetCurrentURI(uri string) error {
	return WrapRecentChooser(gextras.InternObject(c)).SetCurrentURI(uri)
}

func (c recentAction) SetFilter(filter RecentFilter) {
	WrapRecentChooser(gextras.InternObject(c)).SetFilter(filter)
}

func (c recentAction) SetLimit(limit int) {
	WrapRecentChooser(gextras.InternObject(c)).SetLimit(limit)
}

func (c recentAction) SetLocalOnly(localOnly bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetLocalOnly(localOnly)
}

func (c recentAction) SetSelectMultiple(selectMultiple bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetSelectMultiple(selectMultiple)
}

func (c recentAction) SetShowIcons(showIcons bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowIcons(showIcons)
}

func (c recentAction) SetShowNotFound(showNotFound bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowNotFound(showNotFound)
}

func (c recentAction) SetShowPrivate(showPrivate bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowPrivate(showPrivate)
}

func (c recentAction) SetShowTips(showTips bool) {
	WrapRecentChooser(gextras.InternObject(c)).SetShowTips(showTips)
}

func (c recentAction) SetSortType(sortType RecentSortType) {
	WrapRecentChooser(gextras.InternObject(c)).SetSortType(sortType)
}

func (c recentAction) UnselectAll() {
	WrapRecentChooser(gextras.InternObject(c)).UnselectAll()
}

func (c recentAction) UnselectURI(uri string) {
	WrapRecentChooser(gextras.InternObject(c)).UnselectURI(uri)
}
