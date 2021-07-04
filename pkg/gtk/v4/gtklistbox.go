// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBox},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRow},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a
// `GListModel` with gtk_list_box_bind_model() for each item that gets added to
// the model.
type ListBoxCreateWidgetFunc func(item gextras.Objector, widget Widget)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item gextras.Objector // out

	item = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(gextras.Objector)

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item)

	var cret *C.GtkWidget // out

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return cret
}

// ListBoxFilterFunc: will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row ListBoxRow, ok bool)

//export gotk4_ListBoxFilterFunc
func gotk4_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var row ListBoxRow // out

	row = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(ListBoxRow)

	fn := v.(ListBoxFilterFunc)
	ok := fn(row)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ListBoxForeachFunc: a function used by gtk_list_box_selected_foreach().
//
// It will be called on every selected child of the @box.
type ListBoxForeachFunc func(box ListBox, row ListBoxRow)

//export gotk4_ListBoxForeachFunc
func gotk4_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var box ListBox    // out
	var row ListBoxRow // out

	box = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(ListBox)
	row = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(ListBoxRow)

	fn := v.(ListBoxForeachFunc)
	fn(box, row)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 ListBoxRow, row2 ListBoxRow, gint int)

//export gotk4_ListBoxSortFunc
func gotk4_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) C.int {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row1 ListBoxRow // out
	var row2 ListBoxRow // out

	row1 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(ListBoxRow)
	row2 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(ListBoxRow)

	fn := v.(ListBoxSortFunc)
	gint := fn(row1, row2)

	var cret C.int // out

	cret = C.int(gint)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever @row changes or which row is before @row
// changes this is called, which lets you update the header on @row.
//
// You may remove or set a new one via [method@Gtk.ListBoxRow.set_header] or
// just change the state of the current header widget.
type ListBoxUpdateHeaderFunc func(row ListBoxRow, before ListBoxRow)

//export gotk4_ListBoxUpdateHeaderFunc
func gotk4_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var row ListBoxRow    // out
	var before ListBoxRow // out

	row = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(ListBoxRow)
	before = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(ListBoxRow)

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before)
}

// ListBox: `GtkListBox` is a vertical list.
//
// A `GtkListBox` only contains `GtkListBoxRow` children. These rows can by
// dynamically sorted and filtered, and headers can be added dynamically
// depending on the row content. It also allows keyboard and mouse navigation
// and selection like a typical list.
//
// Using `GtkListBox` is often an alternative to `GtkTreeView`, especially when
// the list contents has a more complicated layout than what is allowed by a
// `GtkCellRenderer`, or when the contents is interactive (i.e. has a button in
// it).
//
// Although a `GtkListBox` must have only `GtkListBoxRow` children, you can add
// any kind of widget to it via [method@Gtk.ListBox.prepend],
// [method@Gtk.ListBox.append] and [method@Gtk.ListBox.insert] and a
// `GtkListBoxRow` widget will automatically be inserted between the list and
// the widget.
//
// `GtkListBoxRows` can be marked as activatable or selectable. If a row is
// activatable, [signal@Gtk.ListBox::row-activated] will be emitted for it when
// the user tries to activate it. If it is selectable, the row will be marked as
// selected when the user tries to select it.
//
//
// GtkListBox as GtkBuildable
//
// The `GtkListBox` implementation of the `GtkBuildable` interface supports
// setting a child as the placeholder by specifying “placeholder” as the “type”
// attribute of a <child> element. See [method@Gtk.ListBox.set_placeholder] for
// info.
//
// CSS nodes
//
//    list[.separators][.rich-list][.navigation-sidebar]
//    ╰── row[.activatable]
//
// `GtkListBox` uses a single CSS node named list. It may carry the .separators
// style class, when the [property@Gtk.ListBox:show-separators] property is set.
// Each `GtkListBoxRow` uses a single CSS node named row. The row nodes get the
// .activatable style class added when appropriate.
//
// The main list node may also carry style classes to select the style of list
// presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkListBox` uses the GTK_ACCESSIBLE_ROLE_LIST role and `GtkListBoxRow` uses
// the GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListBox interface {
	Widget

	AppendListBox(child Widget)

	DragHighlightRowListBox(row ListBoxRow)

	DragUnhighlightRowListBox()

	ActivateOnSingleClick() bool

	Adjustment() Adjustment

	RowAtIndex(index_ int) ListBoxRow

	RowAtY(y int) ListBoxRow

	SelectedRow() ListBoxRow

	SelectionMode() SelectionMode

	ShowSeparators() bool

	InsertListBox(child Widget, position int)

	InvalidateFilterListBox()

	InvalidateHeadersListBox()

	InvalidateSortListBox()

	PrependListBox(child Widget)

	RemoveListBox(child Widget)

	SelectAllListBox()

	SelectRowListBox(row ListBoxRow)

	SelectedForeachListBox(fn ListBoxForeachFunc)

	SetActivateOnSingleClickListBox(single bool)

	SetAdjustmentListBox(adjustment Adjustment)

	SetPlaceholderListBox(placeholder Widget)

	SetSelectionModeListBox(mode SelectionMode)

	SetShowSeparatorsListBox(showSeparators bool)

	UnselectAllListBox()

	UnselectRowListBox(row ListBoxRow)
}

// listBox implements the ListBox class.
type listBox struct {
	Widget
}

// WrapListBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBox(obj *externglib.Object) ListBox {
	return listBox{
		Widget: WrapWidget(obj),
	}
}

func marshalListBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBox(obj), nil
}

func NewListBox() ListBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_new()

	var _listBox ListBox // out

	_listBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListBox)

	return _listBox
}

func (b listBox) AppendListBox(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_append(_arg0, _arg1)
}

func (b listBox) DragHighlightRowListBox(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
}

func (b listBox) DragUnhighlightRowListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
}

func (b listBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b listBox) Adjustment() Adjustment {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_list_box_get_adjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (b listBox) RowAtIndex(index_ int) ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(index_)

	_cret = C.gtk_list_box_get_row_at_index(_arg0, _arg1)

	var _listBoxRow ListBoxRow // out

	_listBoxRow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListBoxRow)

	return _listBoxRow
}

func (b listBox) RowAtY(y int) ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _arg1 C.int            // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.int(y)

	_cret = C.gtk_list_box_get_row_at_y(_arg0, _arg1)

	var _listBoxRow ListBoxRow // out

	_listBoxRow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListBoxRow)

	return _listBoxRow
}

func (b listBox) SelectedRow() ListBoxRow {
	var _arg0 *C.GtkListBox    // out
	var _cret *C.GtkListBoxRow // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_list_box_get_selected_row(_arg0)

	var _listBoxRow ListBoxRow // out

	_listBoxRow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListBoxRow)

	return _listBoxRow
}

func (b listBox) SelectionMode() SelectionMode {
	var _arg0 *C.GtkListBox      // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_list_box_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

func (b listBox) ShowSeparators() bool {
	var _arg0 *C.GtkListBox // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_list_box_get_show_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b listBox) InsertListBox(child Widget, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
}

func (b listBox) InvalidateFilterListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
}

func (b listBox) InvalidateHeadersListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
}

func (b listBox) InvalidateSortListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
}

func (b listBox) PrependListBox(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
}

func (b listBox) RemoveListBox(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_remove(_arg0, _arg1)
}

func (b listBox) SelectAllListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_select_all(_arg0)
}

func (b listBox) SelectRowListBox(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_select_row(_arg0, _arg1)
}

func (b listBox) SelectedForeachListBox(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ListBoxForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
}

func (b listBox) SetActivateOnSingleClickListBox(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
}

func (b listBox) SetAdjustmentListBox(adjustment Adjustment) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
}

func (b listBox) SetPlaceholderListBox(placeholder Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(placeholder.Native()))

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
}

func (b listBox) SetSelectionModeListBox(mode SelectionMode) {
	var _arg0 *C.GtkListBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_list_box_set_selection_mode(_arg0, _arg1)
}

func (b listBox) SetShowSeparatorsListBox(showSeparators bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	if showSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_set_show_separators(_arg0, _arg1)
}

func (b listBox) UnselectAllListBox() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_unselect_all(_arg0)
}

func (b listBox) UnselectRowListBox(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
}

func (s listBox) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s listBox) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s listBox) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s listBox) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s listBox) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s listBox) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s listBox) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b listBox) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

// ListBoxRow: `GtkListBoxRow` is the kind of widget that can be added to a
// `GtkListBox`.
type ListBoxRow interface {
	Actionable

	ChangedListBoxRow()

	Activatable() bool

	Child() Widget

	Header() Widget

	Index() int

	Selectable() bool

	IsSelectedListBoxRow() bool

	SetActivatableListBoxRow(activatable bool)

	SetChildListBoxRow(child Widget)

	SetHeaderListBoxRow(header Widget)

	SetSelectableListBoxRow(selectable bool)
}

// listBoxRow implements the ListBoxRow class.
type listBoxRow struct {
	Widget
}

// WrapListBoxRow wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBoxRow(obj *externglib.Object) ListBoxRow {
	return listBoxRow{
		Widget: WrapWidget(obj),
	}
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBoxRow(obj), nil
}

func NewListBoxRow() ListBoxRow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_list_box_row_new()

	var _listBoxRow ListBoxRow // out

	_listBoxRow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListBoxRow)

	return _listBoxRow
}

func (r listBoxRow) ChangedListBoxRow() {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

func (r listBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r listBoxRow) Child() Widget {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (r listBoxRow) Header() Widget {
	var _arg0 *C.GtkListBoxRow // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_get_header(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (r listBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.int            // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (r listBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r listBoxRow) IsSelectedListBoxRow() bool {
	var _arg0 *C.GtkListBoxRow // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r listBoxRow) SetActivatableListBoxRow(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

func (r listBoxRow) SetChildListBoxRow(child Widget) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_row_set_child(_arg0, _arg1)
}

func (r listBoxRow) SetHeaderListBoxRow(header Widget) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

func (r listBoxRow) SetSelectableListBoxRow(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if selectable {
		_arg1 = C.TRUE
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}

func (a listBoxRow) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a listBoxRow) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a listBoxRow) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a listBoxRow) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a listBoxRow) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (s listBoxRow) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s listBoxRow) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s listBoxRow) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s listBoxRow) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s listBoxRow) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s listBoxRow) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s listBoxRow) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b listBoxRow) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}