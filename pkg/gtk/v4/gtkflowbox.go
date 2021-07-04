// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_FlowBoxForeachFunc(GtkFlowBox*, GtkFlowBoxChild*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_get_type()), F: marshalFlowBox},
		{T: externglib.Type(C.gtk_flow_box_child_get_type()), F: marshalFlowBoxChild},
	})
}

// FlowBoxCreateWidgetFunc: called for flow boxes that are bound to a
// `GListModel`.
//
// This function is called for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func(item gextras.Objector, widget Widget)

//export gotk4_FlowBoxCreateWidgetFunc
func gotk4_FlowBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item gextras.Objector // out

	item = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(gextras.Objector)

	fn := v.(FlowBoxCreateWidgetFunc)
	widget := fn(item)

	var cret *C.GtkWidget // out

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// FlowBoxFilterFunc: a function that will be called whenever a child changes or
// is added.
//
// It lets you control if the child should be visible or not.
type FlowBoxFilterFunc func(child FlowBoxChild, ok bool)

//export gotk4_FlowBoxFilterFunc
func gotk4_FlowBoxFilterFunc(arg0 *C.GtkFlowBoxChild, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var child FlowBoxChild // out

	child = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(FlowBoxChild)

	fn := v.(FlowBoxFilterFunc)
	ok := fn(child)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FlowBoxForeachFunc: a function used by gtk_flow_box_selected_foreach().
//
// It will be called on every selected child of the @box.
type FlowBoxForeachFunc func(box FlowBox, child FlowBoxChild)

//export gotk4_FlowBoxForeachFunc
func gotk4_FlowBoxForeachFunc(arg0 *C.GtkFlowBox, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box FlowBox        // out
	var child FlowBoxChild // out

	box = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(FlowBox)
	child = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(FlowBoxChild)

	fn := v.(FlowBoxForeachFunc)
	fn(box, child)
}

// FlowBoxSortFunc: a function to compare two children to determine which should
// come first.
type FlowBoxSortFunc func(child1 FlowBoxChild, child2 FlowBoxChild, gint int)

//export gotk4_FlowBoxSortFunc
func gotk4_FlowBoxSortFunc(arg0 *C.GtkFlowBoxChild, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) C.int {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var child1 FlowBoxChild // out
	var child2 FlowBoxChild // out

	child1 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(FlowBoxChild)
	child2 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(FlowBoxChild)

	fn := v.(FlowBoxSortFunc)
	gint := fn(child1, child2)

	var cret C.int // out

	cret = C.int(gint)

	return cret
}

// FlowBox: a `GtkFlowBox` puts child widgets in reflowing grid.
//
// For instance, with the horizontal orientation, the widgets will be arranged
// from left to right, starting a new row under the previous row when necessary.
// Reducing the width in this case will require more rows, so a larger height
// will be requested.
//
// Likewise, with the vertical orientation, the widgets will be arranged from
// top to bottom, starting a new column to the right when necessary. Reducing
// the height will require more columns, so a larger width will be requested.
//
// The size request of a `GtkFlowBox` alone may not be what you expect; if you
// need to be able to shrink it along both axes and dynamically reflow its
// children, you may have to wrap it in a `GtkScrolledWindow` to enable that.
//
// The children of a `GtkFlowBox` can be dynamically sorted and filtered.
//
// Although a `GtkFlowBox` must have only `GtkFlowBoxChild` children, you can
// add any kind of widget to it via [method@Gtk.FlowBox.insert], and a
// `GtkFlowBoxChild` widget will automatically be inserted between the box and
// the widget.
//
// Also see [class@Gtk.ListBox].
//
//
// CSS nodes
//
// “` flowbox ├── flowboxchild │ ╰── <child> ├── flowboxchild │ ╰── <child> ┊
// ╰── [rubberband] “`
//
// `GtkFlowBox` uses a single CSS node with name flowbox. `GtkFlowBoxChild` uses
// a single CSS node with name flowboxchild. For rubberband selection, a subnode
// with name rubberband is used.
//
//
// Accessibility
//
// `GtkFlowBox` uses the GTK_ACCESSIBLE_ROLE_GRID role, and `GtkFlowBoxChild`
// uses the GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type FlowBox interface {
	Widget
	Orientable

	ActivateOnSingleClick() bool

	ChildAtIndex(idx int) FlowBoxChild

	ChildAtPos(x int, y int) FlowBoxChild

	ColumnSpacing() uint

	Homogeneous() bool

	MaxChildrenPerLine() uint

	MinChildrenPerLine() uint

	RowSpacing() uint

	SelectionMode() SelectionMode

	InsertFlowBox(widget Widget, position int)

	InvalidateFilterFlowBox()

	InvalidateSortFlowBox()

	RemoveFlowBox(widget Widget)

	SelectAllFlowBox()

	SelectChildFlowBox(child FlowBoxChild)

	SelectedForeachFlowBox(fn FlowBoxForeachFunc)

	SetActivateOnSingleClickFlowBox(single bool)

	SetColumnSpacingFlowBox(spacing uint)

	SetHAdjustmentFlowBox(adjustment Adjustment)

	SetHomogeneousFlowBox(homogeneous bool)

	SetMaxChildrenPerLineFlowBox(nChildren uint)

	SetMinChildrenPerLineFlowBox(nChildren uint)

	SetRowSpacingFlowBox(spacing uint)

	SetSelectionModeFlowBox(mode SelectionMode)

	SetVAdjustmentFlowBox(adjustment Adjustment)

	UnselectAllFlowBox()

	UnselectChildFlowBox(child FlowBoxChild)
}

// flowBox implements the FlowBox class.
type flowBox struct {
	Widget
}

// WrapFlowBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBox(obj *externglib.Object) FlowBox {
	return flowBox{
		Widget: WrapWidget(obj),
	}
}

func marshalFlowBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBox(obj), nil
}

func NewFlowBox() FlowBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_flow_box_new()

	var _flowBox FlowBox // out

	_flowBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FlowBox)

	return _flowBox
}

func (b flowBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b flowBox) ChildAtIndex(idx int) FlowBoxChild {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.int              // out
	var _cret *C.GtkFlowBoxChild // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(idx)

	_cret = C.gtk_flow_box_get_child_at_index(_arg0, _arg1)

	var _flowBoxChild FlowBoxChild // out

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FlowBoxChild)

	return _flowBoxChild
}

func (b flowBox) ChildAtPos(x int, y int) FlowBoxChild {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out
	var _cret *C.GtkFlowBoxChild // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gtk_flow_box_get_child_at_pos(_arg0, _arg1, _arg2)

	var _flowBoxChild FlowBoxChild // out

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FlowBoxChild)

	return _flowBoxChild
}

func (b flowBox) ColumnSpacing() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_column_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b flowBox) Homogeneous() bool {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b flowBox) MaxChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_max_children_per_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b flowBox) MinChildrenPerLine() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_min_children_per_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b flowBox) RowSpacing() uint {
	var _arg0 *C.GtkFlowBox // out
	var _cret C.guint       // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (b flowBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkFlowBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_flow_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

func (b flowBox) InsertFlowBox(widget Widget, position int) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.int(position)

	C.gtk_flow_box_insert(_arg0, _arg1, _arg2)
}

func (b flowBox) InvalidateFilterFlowBox() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_invalidate_filter(_arg0)
}

func (b flowBox) InvalidateSortFlowBox() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_invalidate_sort(_arg0)
}

func (b flowBox) RemoveFlowBox(widget Widget) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_flow_box_remove(_arg0, _arg1)
}

func (b flowBox) SelectAllFlowBox() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_select_all(_arg0)
}

func (b flowBox) SelectChildFlowBox(child FlowBoxChild) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_select_child(_arg0, _arg1)
}

func (b flowBox) SelectedForeachFlowBox(fn FlowBoxForeachFunc) {
	var _arg0 *C.GtkFlowBox           // out
	var _arg1 C.GtkFlowBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*[0]byte)(C.gotk4_FlowBoxForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_flow_box_selected_foreach(_arg0, _arg1, _arg2)
}

func (b flowBox) SetActivateOnSingleClickFlowBox(single bool) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_flow_box_set_activate_on_single_click(_arg0, _arg1)
}

func (b flowBox) SetColumnSpacingFlowBox(spacing uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_column_spacing(_arg0, _arg1)
}

func (b flowBox) SetHAdjustmentFlowBox(adjustment Adjustment) {
	var _arg0 *C.GtkFlowBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_hadjustment(_arg0, _arg1)
}

func (b flowBox) SetHomogeneousFlowBox(homogeneous bool) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_flow_box_set_homogeneous(_arg0, _arg1)
}

func (b flowBox) SetMaxChildrenPerLineFlowBox(nChildren uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_max_children_per_line(_arg0, _arg1)
}

func (b flowBox) SetMinChildrenPerLineFlowBox(nChildren uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(nChildren)

	C.gtk_flow_box_set_min_children_per_line(_arg0, _arg1)
}

func (b flowBox) SetRowSpacingFlowBox(spacing uint) {
	var _arg0 *C.GtkFlowBox // out
	var _arg1 C.guint       // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_flow_box_set_row_spacing(_arg0, _arg1)
}

func (b flowBox) SetSelectionModeFlowBox(mode SelectionMode) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_flow_box_set_selection_mode(_arg0, _arg1)
}

func (b flowBox) SetVAdjustmentFlowBox(adjustment Adjustment) {
	var _arg0 *C.GtkFlowBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_flow_box_set_vadjustment(_arg0, _arg1)
}

func (b flowBox) UnselectAllFlowBox() {
	var _arg0 *C.GtkFlowBox // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))

	C.gtk_flow_box_unselect_all(_arg0)
}

func (b flowBox) UnselectChildFlowBox(child FlowBoxChild) {
	var _arg0 *C.GtkFlowBox      // out
	var _arg1 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkFlowBoxChild)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_unselect_child(_arg0, _arg1)
}

func (s flowBox) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s flowBox) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s flowBox) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s flowBox) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s flowBox) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s flowBox) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s flowBox) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b flowBox) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o flowBox) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o flowBox) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}

// FlowBoxChild: `GtkFlowBoxChild` is the kind of widget that can be added to a
// `GtkFlowBox`.
type FlowBoxChild interface {
	Widget

	ChangedFlowBoxChild()

	Child() Widget

	Index() int

	IsSelectedFlowBoxChild() bool

	SetChildFlowBoxChild(child Widget)
}

// flowBoxChild implements the FlowBoxChild class.
type flowBoxChild struct {
	Widget
}

// WrapFlowBoxChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBoxChild(obj *externglib.Object) FlowBoxChild {
	return flowBoxChild{
		Widget: WrapWidget(obj),
	}
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBoxChild(obj), nil
}

func NewFlowBoxChild() FlowBoxChild {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_flow_box_child_new()

	var _flowBoxChild FlowBoxChild // out

	_flowBoxChild = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FlowBoxChild)

	return _flowBoxChild
}

func (c flowBoxChild) ChangedFlowBoxChild() {
	var _arg0 *C.GtkFlowBoxChild // out

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	C.gtk_flow_box_child_changed(_arg0)
}

func (s flowBoxChild) Child() Widget {
	var _arg0 *C.GtkFlowBoxChild // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_flow_box_child_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (c flowBoxChild) Index() int {
	var _arg0 *C.GtkFlowBoxChild // out
	var _cret C.int              // in

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_flow_box_child_get_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c flowBoxChild) IsSelectedFlowBoxChild() bool {
	var _arg0 *C.GtkFlowBoxChild // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_flow_box_child_is_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s flowBoxChild) SetChildFlowBoxChild(child Widget) {
	var _arg0 *C.GtkFlowBoxChild // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_child_set_child(_arg0, _arg1)
}

func (s flowBoxChild) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s flowBoxChild) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s flowBoxChild) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s flowBoxChild) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s flowBoxChild) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s flowBoxChild) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s flowBoxChild) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b flowBoxChild) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}