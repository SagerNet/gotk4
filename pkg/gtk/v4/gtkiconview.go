// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_icon_view_drop_position_get_type()), F: marshalIconViewDropPosition},
		{T: externglib.Type(C.gtk_icon_view_get_type()), F: marshalIconView},
	})
}

// IconViewDropPosition: an enum for determining where a dropped item goes.
type IconViewDropPosition int

const (
	// NoDrop: no drop possible
	IconViewDropPositionNoDrop IconViewDropPosition = 0
	// DropInto: dropped item replaces the item
	IconViewDropPositionDropInto IconViewDropPosition = 1
	// DropLeft: dropped item is inserted to the left
	IconViewDropPositionDropLeft IconViewDropPosition = 2
	// DropRight: dropped item is inserted to the right
	IconViewDropPositionDropRight IconViewDropPosition = 3
	// DropAbove: dropped item is inserted above
	IconViewDropPositionDropAbove IconViewDropPosition = 4
	// DropBelow: dropped item is inserted below
	IconViewDropPositionDropBelow IconViewDropPosition = 5
)

func marshalIconViewDropPosition(p uintptr) (interface{}, error) {
	return IconViewDropPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IconView: `GtkIconView` is a widget which displays data in a grid of icons.
//
// `GtkIconView` provides an alternative view on a `GtkTreeModel`. It displays
// the model as a grid of icons with labels. Like [class@Gtk.TreeView], it
// allows to select one or multiple items (depending on the selection mode, see
// [method@Gtk.IconView.set_selection_mode]). In addition to selection with the
// arrow keys, `GtkIconView` supports rubberband selection, which is controlled
// by dragging the pointer.
//
// Note that if the tree model is backed by an actual tree store (as opposed to
// a flat list where the mapping to icons is obvious), IconView will only
// display the first level of the tree and ignore the tree’s branches.
//
//
// CSS nodes
//
// “` iconview.view ╰── [rubberband] “`
//
// `GtkIconView` has a single CSS node with name iconview and style class .view.
// For rubberband selection, a subnode with name rubberband is used.
type IconView interface {
	Widget
	CellLayout
	Scrollable

	// CreateDragIconIconView:
	CreateDragIconIconView(path *TreePath) gdk.Paintable
	// EnableModelDragDestIconView:
	EnableModelDragDestIconView(formats *gdk.ContentFormats, actions gdk.DragAction)
	// EnableModelDragSourceIconView:
	EnableModelDragSourceIconView(startButtonMask gdk.ModifierType, formats *gdk.ContentFormats, actions gdk.DragAction)
	// ActivateOnSingleClick:
	ActivateOnSingleClick() bool
	// CellRect:
	CellRect(path *TreePath, cell CellRenderer) (gdk.Rectangle, bool)
	// ColumnSpacing:
	ColumnSpacing() int
	// Columns:
	Columns() int
	// Cursor:
	Cursor() (*TreePath, CellRenderer, bool)
	// DestItemAtPos:
	DestItemAtPos(dragX int, dragY int) (*TreePath, IconViewDropPosition, bool)
	// DragDestItem:
	DragDestItem() (*TreePath, IconViewDropPosition)
	// ItemAtPos:
	ItemAtPos(x int, y int) (*TreePath, CellRenderer, bool)
	// ItemColumn:
	ItemColumn(path *TreePath) int
	// ItemOrientation:
	ItemOrientation() Orientation
	// ItemPadding:
	ItemPadding() int
	// ItemRow:
	ItemRow(path *TreePath) int
	// ItemWidth:
	ItemWidth() int
	// Margin:
	Margin() int
	// MarkupColumn:
	MarkupColumn() int
	// Model:
	Model() TreeModel
	// PathAtPos:
	PathAtPos(x int, y int) *TreePath
	// PixbufColumn:
	PixbufColumn() int
	// Reorderable:
	Reorderable() bool
	// RowSpacing:
	RowSpacing() int
	// SelectionMode:
	SelectionMode() SelectionMode
	// Spacing:
	Spacing() int
	// TextColumn:
	TextColumn() int
	// TooltipColumn:
	TooltipColumn() int
	// TooltipContext:
	TooltipContext(x int, y int, keyboardTip bool) (TreeModel, *TreePath, TreeIter, bool)
	// VisibleRange:
	VisibleRange() (startPath *TreePath, endPath *TreePath, ok bool)
	// ItemActivatedIconView:
	ItemActivatedIconView(path *TreePath)
	// PathIsSelectedIconView:
	PathIsSelectedIconView(path *TreePath) bool
	// ScrollToPathIconView:
	ScrollToPathIconView(path *TreePath, useAlign bool, rowAlign float32, colAlign float32)
	// SelectAllIconView:
	SelectAllIconView()
	// SelectPathIconView:
	SelectPathIconView(path *TreePath)
	// SetActivateOnSingleClickIconView:
	SetActivateOnSingleClickIconView(single bool)
	// SetColumnSpacingIconView:
	SetColumnSpacingIconView(columnSpacing int)
	// SetColumnsIconView:
	SetColumnsIconView(columns int)
	// SetCursorIconView:
	SetCursorIconView(path *TreePath, cell CellRenderer, startEditing bool)
	// SetDragDestItemIconView:
	SetDragDestItemIconView(path *TreePath, pos IconViewDropPosition)
	// SetItemOrientationIconView:
	SetItemOrientationIconView(orientation Orientation)
	// SetItemPaddingIconView:
	SetItemPaddingIconView(itemPadding int)
	// SetItemWidthIconView:
	SetItemWidthIconView(itemWidth int)
	// SetMarginIconView:
	SetMarginIconView(margin int)
	// SetMarkupColumnIconView:
	SetMarkupColumnIconView(column int)
	// SetModelIconView:
	SetModelIconView(model TreeModel)
	// SetPixbufColumnIconView:
	SetPixbufColumnIconView(column int)
	// SetReorderableIconView:
	SetReorderableIconView(reorderable bool)
	// SetRowSpacingIconView:
	SetRowSpacingIconView(rowSpacing int)
	// SetSelectionModeIconView:
	SetSelectionModeIconView(mode SelectionMode)
	// SetSpacingIconView:
	SetSpacingIconView(spacing int)
	// SetTextColumnIconView:
	SetTextColumnIconView(column int)
	// SetTooltipCellIconView:
	SetTooltipCellIconView(tooltip Tooltip, path *TreePath, cell CellRenderer)
	// SetTooltipColumnIconView:
	SetTooltipColumnIconView(column int)
	// SetTooltipItemIconView:
	SetTooltipItemIconView(tooltip Tooltip, path *TreePath)
	// UnselectAllIconView:
	UnselectAllIconView()
	// UnselectPathIconView:
	UnselectPathIconView(path *TreePath)
	// UnsetModelDragDestIconView:
	UnsetModelDragDestIconView()
	// UnsetModelDragSourceIconView:
	UnsetModelDragSourceIconView()
}

// iconView implements the IconView class.
type iconView struct {
	Widget
}

// WrapIconView wraps a GObject to the right type. It is
// primarily used internally.
func WrapIconView(obj *externglib.Object) IconView {
	return iconView{
		Widget: WrapWidget(obj),
	}
}

func marshalIconView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIconView(obj), nil
}

// NewIconView:
func NewIconView() IconView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_icon_view_new()

	var _iconView IconView // out

	_iconView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IconView)

	return _iconView
}

// NewIconViewWithArea:
func NewIconViewWithArea(area CellArea) IconView {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_icon_view_new_with_area(_arg1)

	var _iconView IconView // out

	_iconView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IconView)

	return _iconView
}

// NewIconViewWithModel:
func NewIconViewWithModel(model TreeModel) IconView {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_icon_view_new_with_model(_arg1)

	var _iconView IconView // out

	_iconView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(IconView)

	return _iconView
}

func (i iconView) CreateDragIconIconView(path *TreePath) gdk.Paintable {
	var _arg0 *C.GtkIconView  // out
	var _arg1 *C.GtkTreePath  // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_icon_view_create_drag_icon(_arg0, _arg1)

	var _paintable gdk.Paintable // out

	_paintable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gdk.Paintable)

	return _paintable
}

func (i iconView) EnableModelDragDestIconView(formats *gdk.ContentFormats, actions gdk.DragAction) {
	var _arg0 *C.GtkIconView       // out
	var _arg1 *C.GdkContentFormats // out
	var _arg2 C.GdkDragAction      // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(formats.Native()))
	_arg2 = C.GdkDragAction(actions)

	C.gtk_icon_view_enable_model_drag_dest(_arg0, _arg1, _arg2)
}

func (i iconView) EnableModelDragSourceIconView(startButtonMask gdk.ModifierType, formats *gdk.ContentFormats, actions gdk.DragAction) {
	var _arg0 *C.GtkIconView       // out
	var _arg1 C.GdkModifierType    // out
	var _arg2 *C.GdkContentFormats // out
	var _arg3 C.GdkDragAction      // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GdkModifierType(startButtonMask)
	_arg2 = (*C.GdkContentFormats)(unsafe.Pointer(formats.Native()))
	_arg3 = C.GdkDragAction(actions)

	C.gtk_icon_view_enable_model_drag_source(_arg0, _arg1, _arg2, _arg3)
}

func (i iconView) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkIconView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i iconView) CellRect(path *TreePath, cell CellRenderer) (gdk.Rectangle, bool) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GtkCellRenderer // out
	var _arg3 C.GdkRectangle     // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	_cret = C.gtk_icon_view_get_cell_rect(_arg0, _arg1, _arg2, &_arg3)

	var _rect gdk.Rectangle // out
	var _ok bool            // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rect = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

func (i iconView) ColumnSpacing() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_column_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) Columns() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_columns(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) Cursor() (*TreePath, CellRenderer, bool) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTreePath     // in
	var _arg2 *C.GtkCellRenderer // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_cursor(_arg0, &_arg1, &_arg2)

	var _path *TreePath    // out
	var _cell CellRenderer // out
	var _ok bool           // out

	_path = (*TreePath)(unsafe.Pointer(_arg1))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_cell = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg2))).(CellRenderer)
	if _cret != 0 {
		_ok = true
	}

	return _path, _cell, _ok
}

func (i iconView) DestItemAtPos(dragX int, dragY int) (*TreePath, IconViewDropPosition, bool) {
	var _arg0 *C.GtkIconView            // out
	var _arg1 C.int                     // out
	var _arg2 C.int                     // out
	var _arg3 *C.GtkTreePath            // in
	var _arg4 C.GtkIconViewDropPosition // in
	var _cret C.gboolean                // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(dragX)
	_arg2 = C.int(dragY)

	_cret = C.gtk_icon_view_get_dest_item_at_pos(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _path *TreePath           // out
	var _pos IconViewDropPosition // out
	var _ok bool                  // out

	_path = (*TreePath)(unsafe.Pointer(_arg3))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_pos = IconViewDropPosition(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _path, _pos, _ok
}

func (i iconView) DragDestItem() (*TreePath, IconViewDropPosition) {
	var _arg0 *C.GtkIconView            // out
	var _arg1 *C.GtkTreePath            // in
	var _arg2 C.GtkIconViewDropPosition // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_get_drag_dest_item(_arg0, &_arg1, &_arg2)

	var _path *TreePath           // out
	var _pos IconViewDropPosition // out

	_path = (*TreePath)(unsafe.Pointer(_arg1))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_pos = IconViewDropPosition(_arg2)

	return _path, _pos
}

func (i iconView) ItemAtPos(x int, y int) (*TreePath, CellRenderer, bool) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 C.int              // out
	var _arg2 C.int              // out
	var _arg3 *C.GtkTreePath     // in
	var _arg4 *C.GtkCellRenderer // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gtk_icon_view_get_item_at_pos(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _path *TreePath    // out
	var _cell CellRenderer // out
	var _ok bool           // out

	_path = (*TreePath)(unsafe.Pointer(_arg3))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_cell = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg4))).(CellRenderer)
	if _cret != 0 {
		_ok = true
	}

	return _path, _cell, _ok
}

func (i iconView) ItemColumn(path *TreePath) int {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_icon_view_get_item_column(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) ItemOrientation() Orientation {
	var _arg0 *C.GtkIconView   // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_item_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

func (i iconView) ItemPadding() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_item_padding(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) ItemRow(path *TreePath) int {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_icon_view_get_item_row(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) ItemWidth() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_item_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) Margin() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_margin(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) MarkupColumn() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_markup_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) Model() TreeModel {
	var _arg0 *C.GtkIconView  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

func (i iconView) PathAtPos(x int, y int) *TreePath {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out
	var _arg2 C.int          // out
	var _cret *C.GtkTreePath // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gtk_icon_view_get_path_at_pos(_arg0, _arg1, _arg2)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_treePath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

func (i iconView) PixbufColumn() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_pixbuf_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) Reorderable() bool {
	var _arg0 *C.GtkIconView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i iconView) RowSpacing() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_row_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) SelectionMode() SelectionMode {
	var _arg0 *C.GtkIconView     // out
	var _cret C.GtkSelectionMode // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_selection_mode(_arg0)

	var _selectionMode SelectionMode // out

	_selectionMode = SelectionMode(_cret)

	return _selectionMode
}

func (i iconView) Spacing() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_spacing(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) TextColumn() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) TooltipColumn() int {
	var _arg0 *C.GtkIconView // out
	var _cret C.int          // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_tooltip_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (i iconView) TooltipContext(x int, y int, keyboardTip bool) (TreeModel, *TreePath, TreeIter, bool) {
	var _arg0 *C.GtkIconView  // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _arg3 C.gboolean      // out
	var _arg4 *C.GtkTreeModel // in
	var _arg5 *C.GtkTreePath  // in
	var _arg6 C.GtkTreeIter   // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)
	if keyboardTip {
		_arg3 = C.TRUE
	}

	_cret = C.gtk_icon_view_get_tooltip_context(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6)

	var _model TreeModel // out
	var _path *TreePath  // out
	var _iter TreeIter   // out
	var _ok bool         // out

	_model = gextras.CastObject(externglib.Take(unsafe.Pointer(_arg4))).(TreeModel)
	_path = (*TreePath)(unsafe.Pointer(_arg5))
	runtime.SetFinalizer(&_path, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg6
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_iter = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _model, _path, _iter, _ok
}

func (i iconView) VisibleRange() (startPath *TreePath, endPath *TreePath, ok bool) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // in
	var _arg2 *C.GtkTreePath // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_icon_view_get_visible_range(_arg0, &_arg1, &_arg2)

	var _startPath *TreePath // out
	var _endPath *TreePath   // out
	var _ok bool             // out

	_startPath = (*TreePath)(unsafe.Pointer(_arg1))
	runtime.SetFinalizer(&_startPath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_endPath = (*TreePath)(unsafe.Pointer(_arg2))
	runtime.SetFinalizer(&_endPath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})
	if _cret != 0 {
		_ok = true
	}

	return _startPath, _endPath, _ok
}

func (i iconView) ItemActivatedIconView(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_icon_view_item_activated(_arg0, _arg1)
}

func (i iconView) PathIsSelectedIconView(path *TreePath) bool {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	_cret = C.gtk_icon_view_path_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i iconView) ScrollToPathIconView(path *TreePath, useAlign bool, rowAlign float32, colAlign float32) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _arg2 C.gboolean     // out
	var _arg3 C.float        // out
	var _arg4 C.float        // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	if useAlign {
		_arg2 = C.TRUE
	}
	_arg3 = C.float(rowAlign)
	_arg4 = C.float(colAlign)

	C.gtk_icon_view_scroll_to_path(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (i iconView) SelectAllIconView() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_select_all(_arg0)
}

func (i iconView) SelectPathIconView(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_icon_view_select_path(_arg0, _arg1)
}

func (i iconView) SetActivateOnSingleClickIconView(single bool) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_icon_view_set_activate_on_single_click(_arg0, _arg1)
}

func (i iconView) SetColumnSpacingIconView(columnSpacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(columnSpacing)

	C.gtk_icon_view_set_column_spacing(_arg0, _arg1)
}

func (i iconView) SetColumnsIconView(columns int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(columns)

	C.gtk_icon_view_set_columns(_arg0, _arg1)
}

func (i iconView) SetCursorIconView(path *TreePath, cell CellRenderer, startEditing bool) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GtkCellRenderer // out
	var _arg3 C.gboolean         // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if startEditing {
		_arg3 = C.TRUE
	}

	C.gtk_icon_view_set_cursor(_arg0, _arg1, _arg2, _arg3)
}

func (i iconView) SetDragDestItemIconView(path *TreePath, pos IconViewDropPosition) {
	var _arg0 *C.GtkIconView            // out
	var _arg1 *C.GtkTreePath            // out
	var _arg2 C.GtkIconViewDropPosition // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	_arg2 = C.GtkIconViewDropPosition(pos)

	C.gtk_icon_view_set_drag_dest_item(_arg0, _arg1, _arg2)
}

func (i iconView) SetItemOrientationIconView(orientation Orientation) {
	var _arg0 *C.GtkIconView   // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_icon_view_set_item_orientation(_arg0, _arg1)
}

func (i iconView) SetItemPaddingIconView(itemPadding int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(itemPadding)

	C.gtk_icon_view_set_item_padding(_arg0, _arg1)
}

func (i iconView) SetItemWidthIconView(itemWidth int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(itemWidth)

	C.gtk_icon_view_set_item_width(_arg0, _arg1)
}

func (i iconView) SetMarginIconView(margin int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(margin)

	C.gtk_icon_view_set_margin(_arg0, _arg1)
}

func (i iconView) SetMarkupColumnIconView(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_markup_column(_arg0, _arg1)
}

func (i iconView) SetModelIconView(model TreeModel) {
	var _arg0 *C.GtkIconView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_icon_view_set_model(_arg0, _arg1)
}

func (i iconView) SetPixbufColumnIconView(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_pixbuf_column(_arg0, _arg1)
}

func (i iconView) SetReorderableIconView(reorderable bool) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_icon_view_set_reorderable(_arg0, _arg1)
}

func (i iconView) SetRowSpacingIconView(rowSpacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(rowSpacing)

	C.gtk_icon_view_set_row_spacing(_arg0, _arg1)
}

func (i iconView) SetSelectionModeIconView(mode SelectionMode) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_icon_view_set_selection_mode(_arg0, _arg1)
}

func (i iconView) SetSpacingIconView(spacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(spacing)

	C.gtk_icon_view_set_spacing(_arg0, _arg1)
}

func (i iconView) SetTextColumnIconView(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_text_column(_arg0, _arg1)
}

func (i iconView) SetTooltipCellIconView(tooltip Tooltip, path *TreePath, cell CellRenderer) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTooltip      // out
	var _arg2 *C.GtkTreePath     // out
	var _arg3 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	_arg3 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_icon_view_set_tooltip_cell(_arg0, _arg1, _arg2, _arg3)
}

func (i iconView) SetTooltipColumnIconView(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_tooltip_column(_arg0, _arg1)
}

func (i iconView) SetTooltipItemIconView(tooltip Tooltip, path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTooltip  // out
	var _arg2 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_icon_view_set_tooltip_item(_arg0, _arg1, _arg2)
}

func (i iconView) UnselectAllIconView() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unselect_all(_arg0)
}

func (i iconView) UnselectPathIconView(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_icon_view_unselect_path(_arg0, _arg1)
}

func (i iconView) UnsetModelDragDestIconView() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unset_model_drag_dest(_arg0)
}

func (i iconView) UnsetModelDragSourceIconView() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unset_model_drag_source(_arg0)
}

func (s iconView) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s iconView) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s iconView) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s iconView) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s iconView) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s iconView) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s iconView) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b iconView) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (c iconView) AddAttribute(cell CellRenderer, attribute string, column int) {
	WrapCellLayout(gextras.InternObject(c)).AddAttribute(cell, attribute, column)
}

func (c iconView) Clear() {
	WrapCellLayout(gextras.InternObject(c)).Clear()
}

func (c iconView) ClearAttributes(cell CellRenderer) {
	WrapCellLayout(gextras.InternObject(c)).ClearAttributes(cell)
}

func (c iconView) Area() CellArea {
	return WrapCellLayout(gextras.InternObject(c)).Area()
}

func (c iconView) PackEnd(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackEnd(cell, expand)
}

func (c iconView) PackStart(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackStart(cell, expand)
}

func (c iconView) Reorder(cell CellRenderer, position int) {
	WrapCellLayout(gextras.InternObject(c)).Reorder(cell, position)
}

func (s iconView) Border() (Border, bool) {
	return WrapScrollable(gextras.InternObject(s)).Border()
}

func (s iconView) HAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).HAdjustment()
}

func (s iconView) HScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).HScrollPolicy()
}

func (s iconView) VAdjustment() Adjustment {
	return WrapScrollable(gextras.InternObject(s)).VAdjustment()
}

func (s iconView) VScrollPolicy() ScrollablePolicy {
	return WrapScrollable(gextras.InternObject(s)).VScrollPolicy()
}

func (s iconView) SetHAdjustment(hadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetHAdjustment(hadjustment)
}

func (s iconView) SetHScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetHScrollPolicy(policy)
}

func (s iconView) SetVAdjustment(vadjustment Adjustment) {
	WrapScrollable(gextras.InternObject(s)).SetVAdjustment(vadjustment)
}

func (s iconView) SetVScrollPolicy(policy ScrollablePolicy) {
	WrapScrollable(gextras.InternObject(s)).SetVScrollPolicy(policy)
}
