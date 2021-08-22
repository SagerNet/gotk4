// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_grid_get_type()), F: marshalGridder},
	})
}

// Grid is a container which arranges its child widgets in rows and columns,
// with arbitrary positions and horizontal/vertical spans.
//
// Children are added using gtk_grid_attach(). They can span multiple rows or
// columns. It is also possible to add a child next to an existing child, using
// gtk_grid_attach_next_to(). The behaviour of GtkGrid when several children
// occupy the same grid cell is undefined.
//
// GtkGrid can be used like a Box by just using gtk_container_add(), which will
// place children next to each other in the direction determined by the
// Orientable:orientation property. However, if all you want is a single row or
// column, then Box is the preferred widget.
//
//
// CSS nodes
//
// GtkGrid uses a single CSS node with name grid.
type Grid struct {
	Container

	Orientable
	*externglib.Object
}

func WrapGrid(obj *externglib.Object) *Grid {
	return &Grid{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalGridder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGrid(obj), nil
}

// NewGrid creates a new grid widget.
func NewGrid() *Grid {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_grid_new()

	var _grid *Grid // out

	_grid = WrapGrid(externglib.Take(unsafe.Pointer(_cret)))

	return _grid
}

// Attach adds a widget to the grid.
//
// The position of child is determined by left and top. The number of “cells”
// that child will occupy is determined by width and height.
func (grid *Grid) Attach(child Widgetter, left int, top int, width int, height int) {
	var _arg0 *C.GtkGrid   // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out
	var _arg4 C.gint       // out
	var _arg5 C.gint       // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(left)
	_arg3 = C.gint(top)
	_arg4 = C.gint(width)
	_arg5 = C.gint(height)

	C.gtk_grid_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(child)
	runtime.KeepAlive(left)
	runtime.KeepAlive(top)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// AttachNextTo adds a widget to the grid.
//
// The widget is placed next to sibling, on the side determined by side. When
// sibling is NULL, the widget is placed in row (for left or right placement) or
// column 0 (for top or bottom placement), at the end indicated by side.
//
// Attaching widgets labeled [1], [2], [3] with sibling == NULL and side ==
// GTK_POS_LEFT yields a layout of [3][2][1].
func (grid *Grid) AttachNextTo(child Widgetter, sibling Widgetter, side PositionType, width int, height int) {
	var _arg0 *C.GtkGrid        // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.GtkWidget      // out
	var _arg3 C.GtkPositionType // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if sibling != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	}
	_arg3 = C.GtkPositionType(side)
	_arg4 = C.gint(width)
	_arg5 = C.gint(height)

	C.gtk_grid_attach_next_to(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(child)
	runtime.KeepAlive(sibling)
	runtime.KeepAlive(side)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// BaselineRow returns which row defines the global baseline of grid.
func (grid *Grid) BaselineRow() int {
	var _arg0 *C.GtkGrid // out
	var _cret C.gint     // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_baseline_row(_arg0)
	runtime.KeepAlive(grid)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ChildAt gets the child of grid whose area covers the grid cell whose upper
// left corner is at left, top.
func (grid *Grid) ChildAt(left int, top int) Widgetter {
	var _arg0 *C.GtkGrid   // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(left)
	_arg2 = C.gint(top)

	_cret = C.gtk_grid_get_child_at(_arg0, _arg1, _arg2)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(left)
	runtime.KeepAlive(top)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// ColumnHomogeneous returns whether all columns of grid have the same width.
func (grid *Grid) ColumnHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_column_homogeneous(_arg0)
	runtime.KeepAlive(grid)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ColumnSpacing returns the amount of space between the columns of grid.
func (grid *Grid) ColumnSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_column_spacing(_arg0)
	runtime.KeepAlive(grid)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RowBaselinePosition returns the baseline position of row as set by
// gtk_grid_set_row_baseline_position() or the default value
// GTK_BASELINE_POSITION_CENTER.
func (grid *Grid) RowBaselinePosition(row int) BaselinePosition {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.gint                // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(row)

	_cret = C.gtk_grid_get_row_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(row)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// RowHomogeneous returns whether all rows of grid have the same height.
func (grid *Grid) RowHomogeneous() bool {
	var _arg0 *C.GtkGrid // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_row_homogeneous(_arg0)
	runtime.KeepAlive(grid)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing returns the amount of space between the rows of grid.
func (grid *Grid) RowSpacing() uint {
	var _arg0 *C.GtkGrid // out
	var _cret C.guint    // in

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))

	_cret = C.gtk_grid_get_row_spacing(_arg0)
	runtime.KeepAlive(grid)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// InsertColumn inserts a column at the specified position.
//
// Children which are attached at or to the right of this position are moved one
// column to the right. Children which span across this position are grown to
// span the new column.
func (grid *Grid) InsertColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_insert_column(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(position)
}

// InsertNextTo inserts a row or column at the specified position.
//
// The new row or column is placed next to sibling, on the side determined by
// side. If side is GTK_POS_TOP or GTK_POS_BOTTOM, a row is inserted. If side is
// GTK_POS_LEFT of GTK_POS_RIGHT, a column is inserted.
func (grid *Grid) InsertNextTo(sibling Widgetter, side PositionType) {
	var _arg0 *C.GtkGrid        // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 C.GtkPositionType // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))
	_arg2 = C.GtkPositionType(side)

	C.gtk_grid_insert_next_to(_arg0, _arg1, _arg2)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(sibling)
	runtime.KeepAlive(side)
}

// InsertRow inserts a row at the specified position.
//
// Children which are attached at or below this position are moved one row down.
// Children which span across this position are grown to span the new row.
func (grid *Grid) InsertRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_insert_row(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(position)
}

// RemoveColumn removes a column from the grid.
//
// Children that are placed in this column are removed, spanning children that
// overlap this column have their width reduced by one, and children after the
// column are moved to the left.
func (grid *Grid) RemoveColumn(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_remove_column(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(position)
}

// RemoveRow removes a row from the grid.
//
// Children that are placed in this row are removed, spanning children that
// overlap this row have their height reduced by one, and children below the row
// are moved up.
func (grid *Grid) RemoveRow(position int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(position)

	C.gtk_grid_remove_row(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(position)
}

// SetBaselineRow sets which row defines the global baseline for the entire
// grid. Each row in the grid can have its own local baseline, but only one of
// those is global, meaning it will be the baseline in the parent of the grid.
func (grid *Grid) SetBaselineRow(row int) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gint     // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(row)

	C.gtk_grid_set_baseline_row(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(row)
}

// SetColumnHomogeneous sets whether all columns of grid will have the same
// width.
func (grid *Grid) SetColumnHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_column_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetColumnSpacing sets the amount of space between columns of grid.
func (grid *Grid) SetColumnSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_column_spacing(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}

// SetRowBaselinePosition sets how the baseline should be positioned on row of
// the grid, in case that row is assigned more space than is requested.
func (grid *Grid) SetRowBaselinePosition(row int, pos BaselinePosition) {
	var _arg0 *C.GtkGrid            // out
	var _arg1 C.gint                // out
	var _arg2 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.gint(row)
	_arg2 = C.GtkBaselinePosition(pos)

	C.gtk_grid_set_row_baseline_position(_arg0, _arg1, _arg2)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(row)
	runtime.KeepAlive(pos)
}

// SetRowHomogeneous sets whether all rows of grid will have the same height.
func (grid *Grid) SetRowHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_grid_set_row_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(homogeneous)
}

// SetRowSpacing sets the amount of space between rows of grid.
func (grid *Grid) SetRowSpacing(spacing uint) {
	var _arg0 *C.GtkGrid // out
	var _arg1 C.guint    // out

	_arg0 = (*C.GtkGrid)(unsafe.Pointer(grid.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_grid_set_row_spacing(_arg0, _arg1)
	runtime.KeepAlive(grid)
	runtime.KeepAlive(spacing)
}
