// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_size_group_get_type()), F: marshalSizeGroup},
	})
}

// SizeGroup provides a mechanism for grouping a number of widgets together so
// they all request the same amount of space. This is typically useful when you
// want a column of widgets to have the same size, but you can’t use a Grid
// widget.
//
// In detail, the size requested for each widget in a SizeGroup is the maximum
// of the sizes that would have been requested for each widget in the size group
// if they were not in the size group. The mode of the size group (see
// gtk_size_group_set_mode()) determines whether this applies to the horizontal
// size, the vertical size, or both sizes.
//
// Note that size groups only affect the amount of space requested, not the size
// that the widgets finally receive. If you want the widgets in a SizeGroup to
// actually be the same size, you need to pack them in such a way that they get
// the size they request and not more. For example, if you are packing your
// widgets into a table, you would not include the GTK_FILL flag.
//
// SizeGroup objects are referenced by each widget in the size group, so once
// you have added all widgets to a SizeGroup, you can drop the initial reference
// to the size group with g_object_unref(). If the widgets in the size group are
// subsequently destroyed, then they will be removed from the size group and
// drop their references on the size group; when all widgets have been removed,
// the size group will be freed.
//
// Widgets can be part of multiple size groups; GTK+ will compute the horizontal
// size of a widget from the horizontal requisition of all widgets that can be
// reached from the widget by a chain of size groups of type
// GTK_SIZE_GROUP_HORIZONTAL or GTK_SIZE_GROUP_BOTH, and the vertical size from
// the vertical requisition of all widgets that can be reached from the widget
// by a chain of size groups of type GTK_SIZE_GROUP_VERTICAL or
// GTK_SIZE_GROUP_BOTH.
//
// Note that only non-contextual sizes of every widget are ever consulted by
// size groups (since size groups have no knowledge of what size a widget will
// be allocated in one dimension, it cannot derive how much height a widget will
// receive for a given width). When grouping widgets that trade height for width
// in mode GTK_SIZE_GROUP_VERTICAL or GTK_SIZE_GROUP_BOTH: the height for the
// minimum width will be the requested height for all widgets in the group. The
// same is of course true when horizontally grouping width for height widgets.
//
// Widgets that trade height-for-width should set a reasonably large minimum
// width by way of Label:width-chars for instance. Widgets with static sizes as
// well as widgets that grow (such as ellipsizing text) need no such
// considerations.
//
//
// GtkSizeGroup as GtkBuildable
//
// Size groups can be specified in a UI definition by placing an <object>
// element with `class="GtkSizeGroup"` somewhere in the UI definition. The
// widgets that belong to the size group are specified by a <widgets> element
// that may contain multiple <widget> elements, one for each member of the size
// group. The ”name” attribute gives the id of the widget.
//
// An example of a UI definition fragment with GtkSizeGroup:
//
//    <object class="GtkSizeGroup">
//      <property name="mode">GTK_SIZE_GROUP_HORIZONTAL</property>
//      <widgets>
//        <widget name="radio1"/>
//        <widget name="radio2"/>
//      </widgets>
//    </object>
type SizeGroup interface {
	gextras.Objector
	Buildable

	// AddWidget adds a widget to a SizeGroup. In the future, the requisition of
	// the widget will be determined as the maximum of its requisition and the
	// requisition of the other widgets in the size group. Whether this applies
	// horizontally, vertically, or in both directions depends on the mode of
	// the size group. See gtk_size_group_set_mode().
	//
	// When the widget is destroyed or no longer referenced elsewhere, it will
	// be removed from the size group.
	AddWidget(widget Widget)
	// IgnoreHidden returns if invisible widgets are ignored when calculating
	// the size.
	IgnoreHidden() bool
	// Mode gets the current mode of the size group. See
	// gtk_size_group_set_mode().
	Mode() SizeGroupMode
	// Widgets returns the list of widgets associated with @size_group.
	Widgets() *glib.SList
	// RemoveWidget removes a widget from a SizeGroup.
	RemoveWidget(widget Widget)
	// SetIgnoreHidden sets whether unmapped widgets should be ignored when
	// calculating the size.
	SetIgnoreHidden(ignoreHidden bool)
	// SetMode sets the SizeGroupMode of the size group. The mode of the size
	// group determines whether the widgets in the size group should all have
	// the same horizontal requisition (GTK_SIZE_GROUP_HORIZONTAL) all have the
	// same vertical requisition (GTK_SIZE_GROUP_VERTICAL), or should all have
	// the same requisition in both directions (GTK_SIZE_GROUP_BOTH).
	SetMode(mode SizeGroupMode)
}

// sizeGroup implements the SizeGroup interface.
type sizeGroup struct {
	gextras.Objector
	Buildable
}

var _ SizeGroup = (*sizeGroup)(nil)

// WrapSizeGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapSizeGroup(obj *externglib.Object) SizeGroup {
	return SizeGroup{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalSizeGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSizeGroup(obj), nil
}

// NewSizeGroup constructs a class SizeGroup.
func NewSizeGroup(mode SizeGroupMode) SizeGroup {
	var arg1 C.GtkSizeGroupMode

	arg1 = (C.GtkSizeGroupMode)(mode)

	var cret C.GtkSizeGroup
	var goret1 SizeGroup

	cret = C.gtk_size_group_new(mode)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SizeGroup)

	return goret1
}

// AddWidget adds a widget to a SizeGroup. In the future, the requisition of
// the widget will be determined as the maximum of its requisition and the
// requisition of the other widgets in the size group. Whether this applies
// horizontally, vertically, or in both directions depends on the mode of
// the size group. See gtk_size_group_set_mode().
//
// When the widget is destroyed or no longer referenced elsewhere, it will
// be removed from the size group.
func (s sizeGroup) AddWidget(widget Widget) {
	var arg0 *C.GtkSizeGroup
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_size_group_add_widget(arg0, widget)
}

// IgnoreHidden returns if invisible widgets are ignored when calculating
// the size.
func (s sizeGroup) IgnoreHidden() bool {
	var arg0 *C.GtkSizeGroup

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_size_group_get_ignore_hidden(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Mode gets the current mode of the size group. See
// gtk_size_group_set_mode().
func (s sizeGroup) Mode() SizeGroupMode {
	var arg0 *C.GtkSizeGroup

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))

	var cret C.GtkSizeGroupMode
	var goret1 SizeGroupMode

	cret = C.gtk_size_group_get_mode(arg0)

	goret1 = SizeGroupMode(cret)

	return goret1
}

// Widgets returns the list of widgets associated with @size_group.
func (s sizeGroup) Widgets() *glib.SList {
	var arg0 *C.GtkSizeGroup

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))

	var cret *C.GSList
	var goret1 *glib.SList

	cret = C.gtk_size_group_get_widgets(arg0)

	goret1 = glib.WrapSList(unsafe.Pointer(cret))

	return goret1
}

// RemoveWidget removes a widget from a SizeGroup.
func (s sizeGroup) RemoveWidget(widget Widget) {
	var arg0 *C.GtkSizeGroup
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_size_group_remove_widget(arg0, widget)
}

// SetIgnoreHidden sets whether unmapped widgets should be ignored when
// calculating the size.
func (s sizeGroup) SetIgnoreHidden(ignoreHidden bool) {
	var arg0 *C.GtkSizeGroup
	var arg1 C.gboolean

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	if ignoreHidden {
		arg1 = C.gboolean(1)
	}

	C.gtk_size_group_set_ignore_hidden(arg0, ignoreHidden)
}

// SetMode sets the SizeGroupMode of the size group. The mode of the size
// group determines whether the widgets in the size group should all have
// the same horizontal requisition (GTK_SIZE_GROUP_HORIZONTAL) all have the
// same vertical requisition (GTK_SIZE_GROUP_VERTICAL), or should all have
// the same requisition in both directions (GTK_SIZE_GROUP_BOTH).
func (s sizeGroup) SetMode(mode SizeGroupMode) {
	var arg0 *C.GtkSizeGroup
	var arg1 C.GtkSizeGroupMode

	arg0 = (*C.GtkSizeGroup)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkSizeGroupMode)(mode)

	C.gtk_size_group_set_mode(arg0, mode)
}

type SizeGroupPrivate struct {
	native C.GtkSizeGroupPrivate
}

// WrapSizeGroupPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSizeGroupPrivate(ptr unsafe.Pointer) *SizeGroupPrivate {
	if ptr == nil {
		return nil
	}

	return (*SizeGroupPrivate)(ptr)
}

func marshalSizeGroupPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSizeGroupPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SizeGroupPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
