// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// Popover: gtkPopover is a bubble-like context window, primarily meant to
// provide context-dependent information or options. Popovers are attached to a
// widget, passed at construction time on gtk_popover_new(), or updated
// afterwards through gtk_popover_set_relative_to(), by default they will point
// to the whole widget area, although this behavior can be changed through
// gtk_popover_set_pointing_to().
//
// The position of a popover relative to the widget it is attached to can also
// be changed through gtk_popover_set_position().
//
// By default, Popover performs a GTK+ grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Esc key being
// pressed). If no such modal behavior is desired on a popover,
// gtk_popover_set_modal() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. To facilitate this, it supports
// being populated from a Model, using gtk_popover_new_from_model(). In addition
// to all the regular menu model features, this function supports rendering
// sections in the model in a more compact form, as a row of icon buttons
// instead of menu items.
//
// To use this rendering, set the ”display-hint” attribute of the section to
// ”horizontal-buttons” and set the icons of your items with the ”verb-icon”
// attribute.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
// GtkPopover has a single css node called popover. It always gets the
// .background style class and it gets the .menu style class if it is menu-like
// (e.g. PopoverMenu or created using gtk_popover_new_from_model().
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in Entry or TextView get style classes like .touch-selection or .magnifier to
// differentiate from plain popovers.
type Popover interface {
	Bin

	// BindModelPopover:
	BindModelPopover(model gio.MenuModel, actionNamespace string)
	// ConstrainTo:
	ConstrainTo() PopoverConstraint
	// DefaultWidget:
	DefaultWidget() Widget
	// Modal:
	Modal() bool
	// PointingTo:
	PointingTo() (gdk.Rectangle, bool)
	// Position:
	Position() PositionType
	// RelativeTo:
	RelativeTo() Widget
	// TransitionsEnabled:
	TransitionsEnabled() bool
	// PopdownPopover:
	PopdownPopover()
	// PopupPopover:
	PopupPopover()
	// SetConstrainToPopover:
	SetConstrainToPopover(constraint PopoverConstraint)
	// SetDefaultWidgetPopover:
	SetDefaultWidgetPopover(widget Widget)
	// SetModalPopover:
	SetModalPopover(modal bool)
	// SetPointingToPopover:
	SetPointingToPopover(rect *gdk.Rectangle)
	// SetPositionPopover:
	SetPositionPopover(position PositionType)
	// SetRelativeToPopover:
	SetRelativeToPopover(relativeTo Widget)
	// SetTransitionsEnabledPopover:
	SetTransitionsEnabledPopover(transitionsEnabled bool)
}

// popover implements the Popover class.
type popover struct {
	Bin
}

// WrapPopover wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopover(obj *externglib.Object) Popover {
	return popover{
		Bin: WrapBin(obj),
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopover(obj), nil
}

// NewPopover:
func NewPopover(relativeTo Widget) Popover {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	_cret = C.gtk_popover_new(_arg1)

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Popover)

	return _popover
}

// NewPopoverFromModel:
func NewPopoverFromModel(relativeTo Widget, model gio.MenuModel) Popover {
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_popover_new_from_model(_arg1, _arg2)

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Popover)

	return _popover
}

func (p popover) BindModelPopover(model gio.MenuModel, actionNamespace string) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GMenuModel // out
	var _arg2 *C.gchar      // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.gchar)(C.CString(actionNamespace))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_popover_bind_model(_arg0, _arg1, _arg2)
}

func (p popover) ConstrainTo() PopoverConstraint {
	var _arg0 *C.GtkPopover          // out
	var _cret C.GtkPopoverConstraint // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_constrain_to(_arg0)

	var _popoverConstraint PopoverConstraint // out

	_popoverConstraint = PopoverConstraint(_cret)

	return _popoverConstraint
}

func (p popover) DefaultWidget() Widget {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_default_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (p popover) Modal() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_modal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) PointingTo() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkPopover  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_pointing_to(_arg0, &_arg1)

	var _rect gdk.Rectangle // out
	var _ok bool            // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rect = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

func (p popover) Position() PositionType {
	var _arg0 *C.GtkPopover     // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_position(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

func (p popover) RelativeTo() Widget {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_relative_to(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (p popover) TransitionsEnabled() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_popover_get_transitions_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p popover) PopdownPopover() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(_arg0)
}

func (p popover) PopupPopover() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(_arg0)
}

func (p popover) SetConstrainToPopover(constraint PopoverConstraint) {
	var _arg0 *C.GtkPopover          // out
	var _arg1 C.GtkPopoverConstraint // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkPopoverConstraint(constraint)

	C.gtk_popover_set_constrain_to(_arg0, _arg1)
}

func (p popover) SetDefaultWidgetPopover(widget Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(_arg0, _arg1)
}

func (p popover) SetModalPopover(modal bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_modal(_arg0, _arg1)
}

func (p popover) SetPointingToPopover(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
}

func (p popover) SetPositionPopover(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_popover_set_position(_arg0, _arg1)
}

func (p popover) SetRelativeToPopover(relativeTo Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	C.gtk_popover_set_relative_to(_arg0, _arg1)
}

func (p popover) SetTransitionsEnabledPopover(transitionsEnabled bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if transitionsEnabled {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_transitions_enabled(_arg0, _arg1)
}

func (b popover) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b popover) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b popover) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b popover) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b popover) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b popover) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b popover) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
