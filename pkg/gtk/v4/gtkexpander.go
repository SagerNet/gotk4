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
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpander},
	})
}

// Expander: `GtkExpander` allows the user to reveal its child by clicking on an
// expander triangle.
//
// !An example GtkExpander (expander.png)
//
// This is similar to the triangles used in a `GtkTreeView`.
//
// Normally you use an expander as you would use a frame; you create the child
// widget and use [method@Gtk.Expander.set_child] to add it to the expander.
// When the expander is toggled, it will take care of showing and hiding the
// child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a `GtkExpander` but do not add a child
// to it. The expander widget has an [property@Gtk.Expander:expanded[ property
// which can be used to monitor its expansion state. You should watch this
// property with a signal connection as follows:
//
// “`c static void expander_callback (GObject *object, GParamSpec *param_spec,
// gpointer user_data) { GtkExpander *expander;
//
//    expander = GTK_EXPANDER (object);
//
//    if (gtk_expander_get_expanded (expander))
//      {
//        // Show or create widgets
//      }
//    else
//      {
//        // Hide or destroy widgets
//      }
//
// }
//
// static void create_expander (void) { GtkWidget *expander =
// gtk_expander_new_with_mnemonic ("_More Options"); g_signal_connect (expander,
// "notify::expanded", G_CALLBACK (expander_callback), NULL);
//
//    // ...
//
// } “`
//
//
// GtkExpander as GtkBuildable
//
// The `GtkExpander` implementation of the `GtkBuildable` interface supports
// placing a child in the label position by specifying “label” as the “type”
// attribute of a <child> element. A normal content child can be specified
// without specifying a <child> type attribute.
//
// An example of a UI definition fragment with GtkExpander:
//
// “`xml <object class="GtkExpander"> <child type="label"> <object
// class="GtkLabel" id="expander-label"/> </child> <child> <object
// class="GtkEntry" id="expander-content"/> </child> </object> “`
//
//
// CSS nodes
//
// “` expander ╰── box ├── title │ ├── arrow │ ╰── <label widget> ╰── <child> “`
//
// `GtkExpander` has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
//
//
// Accessibility
//
// `GtkExpander` uses the GTK_ACCESSIBLE_ROLE_BUTTON role.
type Expander interface {
	Widget

	Child() Widget

	Expanded() bool

	Label() string

	LabelWidget() Widget

	ResizeToplevel() bool

	UseMarkup() bool

	UseUnderline() bool

	SetChildExpander(child Widget)

	SetExpandedExpander(expanded bool)

	SetLabelExpander(label string)

	SetLabelWidgetExpander(labelWidget Widget)

	SetResizeToplevelExpander(resizeToplevel bool)

	SetUseMarkupExpander(useMarkup bool)

	SetUseUnderlineExpander(useUnderline bool)
}

// expander implements the Expander class.
type expander struct {
	Widget
}

// WrapExpander wraps a GObject to the right type. It is
// primarily used internally.
func WrapExpander(obj *externglib.Object) Expander {
	return expander{
		Widget: WrapWidget(obj),
	}
}

func marshalExpander(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapExpander(obj), nil
}

func NewExpander(label string) Expander {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new(_arg1)

	var _expander Expander // out

	_expander = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expander)

	return _expander
}

func NewExpanderWithMnemonic(label string) Expander {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_expander_new_with_mnemonic(_arg1)

	var _expander Expander // out

	_expander = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expander)

	return _expander
}

func (e expander) Child() Widget {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (e expander) Expanded() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_expanded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) Label() string {
	var _arg0 *C.GtkExpander // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e expander) LabelWidget() Widget {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_label_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (e expander) ResizeToplevel() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_resize_toplevel(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) UseMarkup() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_use_markup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) UseUnderline() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_expander_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e expander) SetChildExpander(child Widget) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_expander_set_child(_arg0, _arg1)
}

func (e expander) SetExpandedExpander(expanded bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_expanded(_arg0, _arg1)
}

func (e expander) SetLabelExpander(label string) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_expander_set_label(_arg0, _arg1)
}

func (e expander) SetLabelWidgetExpander(labelWidget Widget) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_expander_set_label_widget(_arg0, _arg1)
}

func (e expander) SetResizeToplevelExpander(resizeToplevel bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if resizeToplevel {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_resize_toplevel(_arg0, _arg1)
}

func (e expander) SetUseMarkupExpander(useMarkup bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_markup(_arg0, _arg1)
}

func (e expander) SetUseUnderlineExpander(useUnderline bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_underline(_arg0, _arg1)
}

func (s expander) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s expander) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s expander) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s expander) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s expander) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s expander) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s expander) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b expander) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}