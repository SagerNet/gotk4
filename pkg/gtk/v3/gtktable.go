// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_attach_options_get_type()), F: marshalAttachOptions},
		{T: externglib.Type(C.gtk_table_get_type()), F: marshalTable},
	})
}

// AttachOptions denotes the expansion properties that a widget will have when
// it (or its parent) is resized.
type AttachOptions int

const (
	// AttachOptionsExpand: the widget should expand to take up any extra space
	// in its container that has been allocated.
	AttachOptionsExpand AttachOptions = 0b1
	// AttachOptionsShrink: the widget should shrink as and when possible.
	AttachOptionsShrink AttachOptions = 0b10
	// AttachOptionsFill: the widget should fill the space allocated to it.
	AttachOptionsFill AttachOptions = 0b100
)

func marshalAttachOptions(p uintptr) (interface{}, error) {
	return AttachOptions(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Table: the Table functions allow the programmer to arrange widgets in rows
// and columns, making it easy to align many widgets next to each other,
// horizontally and vertically.
//
// Tables are created with a call to gtk_table_new(), the size of which can
// later be changed with gtk_table_resize().
//
// Widgets can be added to a table using gtk_table_attach() or the more
// convenient (but slightly less flexible) gtk_table_attach_defaults().
//
// To alter the space next to a specific row, use gtk_table_set_row_spacing(),
// and for a column, gtk_table_set_col_spacing(). The gaps between all rows or
// columns can be changed by calling gtk_table_set_row_spacings() or
// gtk_table_set_col_spacings() respectively. Note that spacing is added between
// the children, while padding added by gtk_table_attach() is added on either
// side of the widget it belongs to.
//
// gtk_table_set_homogeneous(), can be used to set whether all cells in the
// table will resize themselves to the size of the largest widget in the table.
//
// > Table has been deprecated. Use Grid instead. It provides the same >
// capabilities as GtkTable for arranging widgets in a rectangular grid, but >
// does support height-for-width geometry management.
type Table interface {
	Container

	// AttachTable:
	AttachTable(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint)
	// AttachDefaultsTable:
	AttachDefaultsTable(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint)
	// ColSpacing:
	ColSpacing(column uint) uint
	// DefaultColSpacing:
	DefaultColSpacing() uint
	// DefaultRowSpacing:
	DefaultRowSpacing() uint
	// Homogeneous:
	Homogeneous() bool
	// RowSpacing:
	RowSpacing(row uint) uint
	// Size:
	Size() (rows uint, columns uint)
	// ResizeTable:
	ResizeTable(rows uint, columns uint)
	// SetColSpacingTable:
	SetColSpacingTable(column uint, spacing uint)
	// SetColSpacingsTable:
	SetColSpacingsTable(spacing uint)
	// SetHomogeneousTable:
	SetHomogeneousTable(homogeneous bool)
	// SetRowSpacingTable:
	SetRowSpacingTable(row uint, spacing uint)
	// SetRowSpacingsTable:
	SetRowSpacingsTable(spacing uint)
}

// table implements the Table class.
type table struct {
	Container
}

// WrapTable wraps a GObject to the right type. It is
// primarily used internally.
func WrapTable(obj *externglib.Object) Table {
	return table{
		Container: WrapContainer(obj),
	}
}

func marshalTable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTable(obj), nil
}

// NewTable:
func NewTable(rows uint, columns uint, homogeneous bool) Table {
	var _arg1 C.guint      // out
	var _arg2 C.guint      // out
	var _arg3 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)
	if homogeneous {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_table_new(_arg1, _arg2, _arg3)

	var _table Table // out

	_table = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Table)

	return _table
}

func (t table) AttachTable(child Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	var _arg0 *C.GtkTable        // out
	var _arg1 *C.GtkWidget       // out
	var _arg2 C.guint            // out
	var _arg3 C.guint            // out
	var _arg4 C.guint            // out
	var _arg5 C.guint            // out
	var _arg6 C.GtkAttachOptions // out
	var _arg7 C.GtkAttachOptions // out
	var _arg8 C.guint            // out
	var _arg9 C.guint            // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)
	_arg6 = C.GtkAttachOptions(xoptions)
	_arg7 = C.GtkAttachOptions(yoptions)
	_arg8 = C.guint(xpadding)
	_arg9 = C.guint(ypadding)

	C.gtk_table_attach(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

func (t table) AttachDefaultsTable(widget Widget, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	var _arg0 *C.GtkTable  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.guint      // out
	var _arg3 C.guint      // out
	var _arg4 C.guint      // out
	var _arg5 C.guint      // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.guint(leftAttach)
	_arg3 = C.guint(rightAttach)
	_arg4 = C.guint(topAttach)
	_arg5 = C.guint(bottomAttach)

	C.gtk_table_attach_defaults(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (t table) ColSpacing(column uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(column)

	_cret = C.gtk_table_get_col_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (t table) DefaultColSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_default_col_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (t table) DefaultRowSpacing() uint {
	var _arg0 *C.GtkTable // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_default_row_spacing(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (t table) Homogeneous() bool {
	var _arg0 *C.GtkTable // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_table_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t table) RowSpacing(row uint) uint {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(row)

	_cret = C.gtk_table_get_row_spacing(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (t table) Size() (rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // in
	var _arg2 C.guint     // in

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))

	C.gtk_table_get_size(_arg0, &_arg1, &_arg2)

	var _rows uint    // out
	var _columns uint // out

	_rows = uint(_arg1)
	_columns = uint(_arg2)

	return _rows, _columns
}

func (t table) ResizeTable(rows uint, columns uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(rows)
	_arg2 = C.guint(columns)

	C.gtk_table_resize(_arg0, _arg1, _arg2)
}

func (t table) SetColSpacingTable(column uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(column)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_col_spacing(_arg0, _arg1, _arg2)
}

func (t table) SetColSpacingsTable(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_col_spacings(_arg0, _arg1)
}

func (t table) SetHomogeneousTable(homogeneous bool) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_table_set_homogeneous(_arg0, _arg1)
}

func (t table) SetRowSpacingTable(row uint, spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out
	var _arg2 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(row)
	_arg2 = C.guint(spacing)

	C.gtk_table_set_row_spacing(_arg0, _arg1, _arg2)
}

func (t table) SetRowSpacingsTable(spacing uint) {
	var _arg0 *C.GtkTable // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkTable)(unsafe.Pointer(t.Native()))
	_arg1 = C.guint(spacing)

	C.gtk_table_set_row_spacings(_arg0, _arg1)
}

func (b table) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b table) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b table) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b table) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b table) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b table) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b table) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

// TableChild:
type TableChild C.GtkTableChild

// WrapTableChild wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableChild(ptr unsafe.Pointer) *TableChild {
	return (*TableChild)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TableChild) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// TableRowCol:
type TableRowCol C.GtkTableRowCol

// WrapTableRowCol wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTableRowCol(ptr unsafe.Pointer) *TableRowCol {
	return (*TableRowCol)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TableRowCol) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}
