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
		{T: externglib.Type(C.gtk_column_view_get_type()), F: marshalColumnView},
	})
}

// ColumnView: `GtkColumnView` presents a large dynamic list of items using
// multiple columns with headers.
//
// `GtkColumnView` uses the factories of its columns to generate a cell widget
// for each column, for each visible item and displays them together as the row
// for this item.
//
// The [property@Gtk.ColumnView:show-row-separators] and
// [propertyGtk.ColumnView:show-column-separators] properties offer a simple way
// to display separators between the rows or columns.
//
// `GtkColumnView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// [property@Gtk.ColumnView:enable-rubberband].
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the `GtkSorter` returned by
// [method@Gtk.ColumnView.get_sorter] must be attached to a sort model for the
// data that the view is showing, and the columns must have sorters attached to
// them by calling [method@Gtk.ColumnViewColumn.set_sorter]. The initial sort
// order can be set with [method@Gtk.ColumnView.sort_by_column].
//
// The column view also supports interactive resizing and reordering of columns,
// via Drag-and-Drop of the column headers. This can be enabled or disabled with
// the [property@Gtk.ColumnView:reorderable] and
// [property@Gtk.ColumnViewColumn:resizable] properties.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
//
// CSS nodes
//
// “`
// columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
// ├── header │ ├── <column header> ┊ ┊ │ ╰── <column header> │ ├── listview │ ┊
// ╰── [rubberband] “`
//
// `GtkColumnView` uses a single CSS node named columnview. It may carry the
// .column-separators style class, when
// [property@Gtk.ColumnView:show-column-separators] property is set. Header
// widgets appear below a node with name header. The rows are contained in a
// `GtkListView` widget, so there is a listview node with the same structure as
// for a standalone `GtkListView` widget. If
// [property@Gtk.ColumnView:show-row-separators] is set, it will be passed on to
// the list view, causing its CSS node to carry the .separators style class. For
// rubberband selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select the style of
// list presentation (section-list-widget.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// `GtkColumnView` uses the GTK_ACCESSIBLE_ROLE_TREE_GRID role, header title
// widgets are using the GTK_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets
// are using the GTK_ACCESSIBLE_ROLE_ROW role, and individual cells are using
// the GTK_ACCESSIBLE_ROLE_GRID_CELL role
type ColumnView interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsScrollable casts the class to the Scrollable interface.
	AsScrollable() Scrollable

	// AppendColumn appends the @column to the end of the columns in @self.
	AppendColumn(column ColumnViewColumn)
	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband() bool
	// Reorderable returns whether columns are reorderable.
	Reorderable() bool
	// ShowColumnSeparators returns whether the list should show separators
	// between columns.
	ShowColumnSeparators() bool
	// ShowRowSeparators returns whether the list should show separators between
	// rows.
	ShowRowSeparators() bool
	// SingleClickActivate returns whether rows will be activated on single
	// click and selected on hover.
	SingleClickActivate() bool
	// Sorter returns a special sorter that reflects the users sorting choices
	// in the column view.
	//
	// To allow users to customizable sorting by clicking on column headers,
	// this sorter needs to be set on the sort model underneath the model that
	// is displayed by the view.
	//
	// See [method@Gtk.ColumnViewColumn.set_sorter] for setting up per-column
	// sorting.
	//
	// Here is an example: “`c gtk_column_view_column_set_sorter (column,
	// sorter); gtk_column_view_append_column (view, column); sorter =
	// g_object_ref (gtk_column_view_get_sorter (view))); model =
	// gtk_sort_list_model_new (store, sorter); selection = gtk_no_selection_new
	// (model); gtk_column_view_set_model (view, selection); “`
	Sorter() Sorter
	// InsertColumn inserts a column at the given position in the columns of
	// @self.
	//
	// If @column is already a column of @self, it will be repositioned.
	InsertColumn(position uint, column ColumnViewColumn)
	// RemoveColumn removes the @column from the list of columns of @self.
	RemoveColumn(column ColumnViewColumn)
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(enableRubberband bool)
	// SetReorderable sets whether columns should be reorderable by dragging.
	SetReorderable(reorderable bool)
	// SetShowColumnSeparators sets whether the list should show separators
	// between columns.
	SetShowColumnSeparators(showColumnSeparators bool)
	// SetShowRowSeparators sets whether the list should show separators between
	// rows.
	SetShowRowSeparators(showRowSeparators bool)
	// SetSingleClickActivate sets whether rows should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(singleClickActivate bool)
	// SortByColumn sets the sorting of the view.
	//
	// This function should be used to set up the initial sorting. At runtime,
	// users can change the sorting of a column view by clicking on the list
	// headers.
	//
	// This call only has an effect if the sorter returned by
	// [method@Gtk.ColumnView.get_sorter] is set on a sort model, and
	// [method@Gtk.ColumnViewColumn.set_sorter] has been called on @column to
	// associate a sorter with the column.
	//
	// If @column is nil, the view will be unsorted.
	SortByColumn(column ColumnViewColumn, direction SortType)
}

// columnView implements the ColumnView class.
type columnView struct {
	Widget
}

// WrapColumnView wraps a GObject to the right type. It is
// primarily used internally.
func WrapColumnView(obj *externglib.Object) ColumnView {
	return columnView{
		Widget: WrapWidget(obj),
	}
}

func marshalColumnView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColumnView(obj), nil
}

func (c columnView) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(c))
}

func (c columnView) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(c))
}

func (c columnView) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(c))
}

func (c columnView) AsScrollable() Scrollable {
	return WrapScrollable(gextras.InternObject(c))
}

func (s columnView) AppendColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_append_column(_arg0, _arg1)
}

func (s columnView) EnableRubberband() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_enable_rubberband(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s columnView) Reorderable() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s columnView) ShowColumnSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_show_column_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s columnView) ShowRowSeparators() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_show_row_separators(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s columnView) SingleClickActivate() bool {
	var _arg0 *C.GtkColumnView // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_single_click_activate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s columnView) Sorter() Sorter {
	var _arg0 *C.GtkColumnView // out
	var _cret *C.GtkSorter     // in

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_column_view_get_sorter(_arg0)

	var _sorter Sorter // out

	_sorter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Sorter)

	return _sorter
}

func (s columnView) InsertColumn(position uint, column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 C.guint                // out
	var _arg2 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)
	_arg2 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_insert_column(_arg0, _arg1, _arg2)
}

func (s columnView) RemoveColumn(column ColumnViewColumn) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_remove_column(_arg0, _arg1)
}

func (s columnView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_enable_rubberband(_arg0, _arg1)
}

func (s columnView) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_reorderable(_arg0, _arg1)
}

func (s columnView) SetShowColumnSeparators(showColumnSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showColumnSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_column_separators(_arg0, _arg1)
}

func (s columnView) SetShowRowSeparators(showRowSeparators bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showRowSeparators {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_show_row_separators(_arg0, _arg1)
}

func (s columnView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkColumnView // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_column_view_set_single_click_activate(_arg0, _arg1)
}

func (s columnView) SortByColumn(column ColumnViewColumn, direction SortType) {
	var _arg0 *C.GtkColumnView       // out
	var _arg1 *C.GtkColumnViewColumn // out
	var _arg2 C.GtkSortType          // out

	_arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))
	_arg2 = C.GtkSortType(direction)

	C.gtk_column_view_sort_by_column(_arg0, _arg1, _arg2)
}
