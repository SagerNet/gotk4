// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_grid_get_type()), F: marshalGrid},
	})
}

// Grid: `GtkGrid` is a container which arranges its child widgets in rows and
// columns.
//
// !An example GtkGrid (grid.png)
//
// It supports arbitrary positions and horizontal/vertical spans.
//
// Children are added using [method@Gtk.Grid.attach]. They can span multiple
// rows or columns. It is also possible to add a child next to an existing
// child, using [method@Gtk.Grid.attach_next_to]. To remove a child from the
// grid, use [method@Gtk.Grid.remove].
//
// The behaviour of `GtkGrid` when several children occupy the same grid cell is
// undefined.
//
//
// GtkGrid as GtkBuildable
//
// Every child in a `GtkGrid` has access to a custom [iface@Gtk.Buildable]
// element, called ´<layout>´. It can by used to specify a position in the grid
// and optionally spans. All properties that can be used in the ´<layout>´
// element are implemented by [class@Gtk.GridLayoutChild].
//
// It is implemented by `GtkWidget` using [class@Gtk.LayoutManager].
//
// To showcase it, here is a simple example:
//
// “`xml <object class="GtkGrid" id="my_grid"> <child> <object class="GtkButton"
// id="button1"> <property name="label">Button 1</property> <layout> <property
// name="column">0</property> <property name="row">0</property> </layout>
// </object> </child> <child> <object class="GtkButton" id="button2"> <property
// name="label">Button 2</property> <layout> <property
// name="column">1</property> <property name="row">0</property> </layout>
// </object> </child> <child> <object class="GtkButton" id="button3"> <property
// name="label">Button 3</property> <layout> <property
// name="column">2</property> <property name="row">0</property> <property
// name="row-span">2</property> </layout> </object> </child> <child> <object
// class="GtkButton" id="button4"> <property name="label">Button 4</property>
// <layout> <property name="column">0</property> <property
// name="row">1</property> <property name="column-span">2</property> </layout>
// </object> </child> </object> “`
//
// It organizes the first two buttons side-by-side in one cell each. The third
// button is in the last column but spans across two rows. This is defined by
// the ´row-span´ property. The last button is located in the second row and
// spans across two columns, which is defined by the ´column-span´ property.
//
//
// CSS nodes
//
// `GtkGrid` uses a single CSS node with name `grid`.
//
//
// Accessibility
//
// `GtkGrid` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Grid interface {
	Widget
	Orientable

	// AttachGrid:
	AttachGrid(child Widget, column int, row int, width int, height int)
	// AttachNextToGrid:
	AttachNextToGrid(child Widget, sibling Widget, side PositionType, width int, height int)
	// BaselineRow:
	BaselineRow() int
	// ChildAt:
	ChildAt(column int, row int) Widget
	// ColumnHomogeneous:
	ColumnHomogeneous() bool
	// ColumnSpacing:
	ColumnSpacing() uint
	// RowBaselinePosition:
	RowBaselinePosition(row int) BaselinePosition
	// RowHomogeneous:
	RowHomogeneous() bool
	// RowSpacing:
	RowSpacing() uint
	// InsertColumnGrid:
	InsertColumnGrid(position int)
	// InsertNextToGrid:
	InsertNextToGrid(sibling Widget, side PositionType)
	// InsertRowGrid:
	InsertRowGrid(position int)
	// QueryChildGrid:
	QueryChildGrid(child Widget) (column int, row int, width int, height int)
	// RemoveGrid:
	RemoveGrid(child Widget)
	// RemoveColumnGrid:
	RemoveColumnGrid(position int)
	// RemoveRowGrid:
	RemoveRowGrid(position int)
	// SetBaselineRowGrid:
	SetBaselineRowGrid(row int)
	// SetColumnHomogeneousGrid:
	SetColumnHomogeneousGrid(homogeneous bool)
	// SetColumnSpacingGrid:
	SetColumnSpacingGrid(spacing uint)
	// SetRowBaselinePositionGrid:
	SetRowBaselinePositionGrid(row int, pos BaselinePosition)
	// SetRowHomogeneousGrid:
	SetRowHomogeneousGrid(homogeneous bool)
	// SetRowSpacingGrid:
	SetRowSpacingGrid(spacing uint)
}

// grid implements the Grid class.
type grid struct {
	Widget
}

// WrapGrid wraps a GObject to the right type. It is
// primarily used internally.
func WrapGrid(obj *externglib.Object) Grid {
	return grid{
		Widget: WrapWidget(obj),
	}
}

func marshalGrid(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGrid(obj), nil
}

// NewGrid:
func NewGrid() Grid {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_grid_new()

	var _grid Grid // out

	_grid = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Grid)

	return _grid
}

func (g grid) AttachGrid(child Widget, column int, row int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // out
	var _arg3 C.int        // out
	var _arg4 C.int        // out
	var _arg5 C.int        // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(column)
	_arg3 = C.int(row)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.gtk_grid_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (g grid) AttachNextToGrid(child Widget, sibling Widget, side PositionType, width int, height int) {
	var _arg0 *C.GtkGrid        // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.GtkWidget      // out
	var _arg3 C.GtkPositionType // out
	var _arg4 C.int             // out
	var _arg5 C.int             // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	_arg3 = C.GtkPositionType(side)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	C.gtk_grid_attach_next_to(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (g grid) BaselineRow() int {
	var _arg0 *C.GtkGrid // out
	var _cret C.int      // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_grid_get_baseline_row(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (g grid) ChildAt(column int, row int) Widget {
	var _arg0 *C.GtkGrid   // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(column)
	_arg2 = C.int(row)

	_cret = C.gtk_grid_get_child_at(_arg0, _arg1, _arg2)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (g grid) ColumnHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_grid_get_column_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (g grid) ColumnSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_grid_get_column_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (g grid) RowBaselinePosition(row int) BaselinePosition {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.int                 // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(row)

	_cret = C.gtk_grid_get_row_baseline_position(_arg0, _arg1)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

func (g grid) RowHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_grid_get_row_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (g grid) RowSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_grid_get_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (g grid) InsertColumnGrid(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_insert_column(_arg0, _arg1)
}

func (g grid) InsertNextToGrid(sibling Widget, side PositionType) {
	var _arg0 *C.GtkGrid        // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.GtkPositionType // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	_arg2 = C.GtkPositionType(side)

	C.gtk_grid_insert_next_to(_arg0, _arg1, _arg2)
}

func (g grid) InsertRowGrid(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_insert_row(_arg0, _arg1)
}

func (g grid) QueryChildGrid(child Widget) (column int, row int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // in
	var _arg3 C.int        // in
	var _arg4 C.int        // in
	var _arg5 C.int        // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_grid_query_child(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _column int // out
	var _row int    // out
	var _width int  // out
	var _height int // out

	_column = int(_arg2)
	_row = int(_arg3)
	_width = int(_arg4)
	_height = int(_arg5)

	return _column, _row, _width, _height
}

func (g grid) RemoveGrid(child Widget) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_grid_remove(_arg0, _arg1)
}

func (g grid) RemoveColumnGrid(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_remove_column(_arg0, _arg1)
}

func (g grid) RemoveRowGrid(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(position)

	C.gtk_grid_remove_row(_arg0, _arg1)
}

func (g grid) SetBaselineRowGrid(row int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.int      // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(row)

	C.gtk_grid_set_baseline_row(_arg0, _arg1)
}

func (g grid) SetColumnHomogeneousGrid(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_column_homogeneous(_arg0, _arg1)
}

func (g grid) SetColumnSpacingGrid(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_column_spacing(_arg0, _arg1)
}

func (g grid) SetRowBaselinePositionGrid(row int, pos BaselinePosition) {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.int                 // out
	var _arg2 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.int(row)
	_arg2 = C.GtkBaselinePosition(pos)

	C.gtk_grid_set_row_baseline_position(_arg0, _arg1, _arg2)
}

func (g grid) SetRowHomogeneousGrid(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_row_homogeneous(_arg0, _arg1)
}

func (g grid) SetRowSpacingGrid(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_row_spacing(_arg0, _arg1)
}

func (s grid) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s grid) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s grid) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s grid) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s grid) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s grid) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s grid) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b grid) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (o grid) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o grid) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
