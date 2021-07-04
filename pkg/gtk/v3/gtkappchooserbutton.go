// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButton},
	})
}

// AppChooserButton: the AppChooserButton is a widget that lets the user select
// an application. It implements the AppChooser interface.
//
// Initially, a AppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// AppChooserButton:show-default-item is true, the default application.
//
// The list of applications shown in a AppChooserButton includes the recommended
// applications for the given content type. When
// AppChooserButton:show-default-item is set, the default application is also
// included. To let the user chooser other applications, you can set the
// AppChooserButton:show-dialog-item property, which allows to open a full
// AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk_app_chooser_button_append_custom_item(). These items cause the
// AppChooserButton::custom-item-activated signal to be emitted when they are
// selected.
//
// To track changes in the selected application, use the ComboBox::changed
// signal.
type AppChooserButton interface {
	ComboBox
	AppChooser

	// AppendCustomItemAppChooserButton:
	AppendCustomItemAppChooserButton(name string, label string, icon gio.Icon)
	// AppendSeparatorAppChooserButton:
	AppendSeparatorAppChooserButton()
	// Heading:
	Heading() string
	// ShowDefaultItem:
	ShowDefaultItem() bool
	// ShowDialogItem:
	ShowDialogItem() bool
	// SetActiveCustomItemAppChooserButton:
	SetActiveCustomItemAppChooserButton(name string)
	// SetHeadingAppChooserButton:
	SetHeadingAppChooserButton(heading string)
	// SetShowDefaultItemAppChooserButton:
	SetShowDefaultItemAppChooserButton(setting bool)
	// SetShowDialogItemAppChooserButton:
	SetShowDialogItemAppChooserButton(setting bool)
}

// appChooserButton implements the AppChooserButton class.
type appChooserButton struct {
	ComboBox
}

// WrapAppChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserButton(obj *externglib.Object) AppChooserButton {
	return appChooserButton{
		ComboBox: WrapComboBox(obj),
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserButton(obj), nil
}

// NewAppChooserButton:
func NewAppChooserButton(contentType string) AppChooserButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_button_new(_arg1)

	var _appChooserButton AppChooserButton // out

	_appChooserButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(AppChooserButton)

	return _appChooserButton
}

func (s appChooserButton) AppendCustomItemAppChooserButton(name string, label string, icon gio.Icon) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.gchar               // out
	var _arg2 *C.gchar               // out
	var _arg3 *C.GIcon               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_app_chooser_button_append_custom_item(_arg0, _arg1, _arg2, _arg3)
}

func (s appChooserButton) AppendSeparatorAppChooserButton() {
	var _arg0 *C.GtkAppChooserButton // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_append_separator(_arg0)
}

func (s appChooserButton) Heading() string {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s appChooserButton) ShowDefaultItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_default_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserButton) ShowDialogItem() bool {
	var _arg0 *C.GtkAppChooserButton // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_button_get_show_dialog_item(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s appChooserButton) SetActiveCustomItemAppChooserButton(name string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_active_custom_item(_arg0, _arg1)
}

func (s appChooserButton) SetHeadingAppChooserButton(heading string) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_button_set_heading(_arg0, _arg1)
}

func (s appChooserButton) SetShowDefaultItemAppChooserButton(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_default_item(_arg0, _arg1)
}

func (s appChooserButton) SetShowDialogItemAppChooserButton(setting bool) {
	var _arg0 *C.GtkAppChooserButton // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_button_set_show_dialog_item(_arg0, _arg1)
}

func (b appChooserButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b appChooserButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b appChooserButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b appChooserButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b appChooserButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b appChooserButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b appChooserButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c appChooserButton) EditingDone() {
	WrapCellEditable(gextras.InternObject(c)).EditingDone()
}

func (c appChooserButton) RemoveWidget() {
	WrapCellEditable(gextras.InternObject(c)).RemoveWidget()
}

func (b appChooserButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b appChooserButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b appChooserButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b appChooserButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b appChooserButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b appChooserButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b appChooserButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (c appChooserButton) AddAttribute(cell CellRenderer, attribute string, column int) {
	WrapCellLayout(gextras.InternObject(c)).AddAttribute(cell, attribute, column)
}

func (c appChooserButton) Clear() {
	WrapCellLayout(gextras.InternObject(c)).Clear()
}

func (c appChooserButton) ClearAttributes(cell CellRenderer) {
	WrapCellLayout(gextras.InternObject(c)).ClearAttributes(cell)
}

func (c appChooserButton) Area() CellArea {
	return WrapCellLayout(gextras.InternObject(c)).Area()
}

func (c appChooserButton) PackEnd(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackEnd(cell, expand)
}

func (c appChooserButton) PackStart(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackStart(cell, expand)
}

func (c appChooserButton) Reorder(cell CellRenderer, position int) {
	WrapCellLayout(gextras.InternObject(c)).Reorder(cell, position)
}

func (s appChooserButton) AppInfo() gio.AppInfo {
	return WrapAppChooser(gextras.InternObject(s)).AppInfo()
}

func (s appChooserButton) ContentType() string {
	return WrapAppChooser(gextras.InternObject(s)).ContentType()
}

func (s appChooserButton) Refresh() {
	WrapAppChooser(gextras.InternObject(s)).Refresh()
}

func (b appChooserButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b appChooserButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b appChooserButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b appChooserButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b appChooserButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b appChooserButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b appChooserButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
