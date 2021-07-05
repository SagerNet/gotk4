// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
//
// void gotk4_IconViewForeachFunc(GtkIconView*, GtkTreePath*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_view_drop_position_get_type()), F: marshalIconViewDropPosition},
		{T: externglib.Type(C.gtk_icon_view_get_type()), F: marshalIconView},
	})
}

// IconViewDropPosition: enum for determining where a dropped item goes.
type IconViewDropPosition int

const (
	// NoDrop: no drop possible
	IconViewDropPositionNoDrop IconViewDropPosition = iota
	// DropInto: dropped item replaces the item
	IconViewDropPositionDropInto
	// DropLeft: dropped item is inserted to the left
	IconViewDropPositionDropLeft
	// DropRight: dropped item is inserted to the right
	IconViewDropPositionDropRight
	// DropAbove: dropped item is inserted above
	IconViewDropPositionDropAbove
	// DropBelow: dropped item is inserted below
	IconViewDropPositionDropBelow
)

func marshalIconViewDropPosition(p uintptr) (interface{}, error) {
	return IconViewDropPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IconViewForeachFunc: function used by gtk_icon_view_selected_foreach() to map
// all selected rows.
//
// It will be called on every selected row in the view.
type IconViewForeachFunc func(iconView IconView, path *TreePath)

//export gotk4_IconViewForeachFunc
func gotk4_IconViewForeachFunc(arg0 *C.GtkIconView, arg1 *C.GtkTreePath, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var iconView IconView // out
	var path *TreePath    // out

	iconView = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(IconView)
	path = (*TreePath)(unsafe.Pointer(arg1))

	fn := v.(IconViewForeachFunc)
	fn(iconView, path)
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

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellLayout casts the class to the CellLayout interface.
	AsCellLayout() CellLayout
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsScrollable casts the class to the Scrollable interface.
	AsScrollable() Scrollable

	// EnableModelDragDest turns @icon_view into a drop destination for
	// automatic DND. Calling this method sets IconView:reorderable to false.
	EnableModelDragDest(formats *gdk.ContentFormats, actions gdk.DragAction)
	// EnableModelDragSource turns @icon_view into a drag source for automatic
	// DND. Calling this method sets IconView:reorderable to false.
	EnableModelDragSource(startButtonMask gdk.ModifierType, formats *gdk.ContentFormats, actions gdk.DragAction)
	// ActivateOnSingleClick gets the setting set by
	// gtk_icon_view_set_activate_on_single_click().
	ActivateOnSingleClick() bool
	// CellRect fills the bounding rectangle in widget coordinates for the cell
	// specified by @path and @cell. If @cell is nil the main cell area is used.
	//
	// This function is only valid if @icon_view is realized.
	CellRect(path *TreePath, cell CellRenderer) (gdk.Rectangle, bool)
	// ColumnSpacing returns the value of the ::column-spacing property.
	ColumnSpacing() int
	// Columns returns the value of the ::columns property.
	Columns() int
	// Cursor fills in @path and @cell with the current cursor path and cell. If
	// the cursor isn’t currently set, then *@path will be nil. If no cell
	// currently has focus, then *@cell will be nil.
	//
	// The returned TreePath must be freed with gtk_tree_path_free().
	Cursor() (*TreePath, CellRenderer, bool)
	// DestItemAtPos determines the destination item for a given position.
	DestItemAtPos(dragX int, dragY int) (*TreePath, IconViewDropPosition, bool)
	// DragDestItem gets information about the item that is highlighted for
	// feedback.
	DragDestItem() (*TreePath, IconViewDropPosition)
	// ItemAtPos gets the path and cell for the icon at the given position.
	ItemAtPos(x int, y int) (*TreePath, CellRenderer, bool)
	// ItemColumn gets the column in which the item @path is currently
	// displayed. Column numbers start at 0.
	ItemColumn(path *TreePath) int
	// ItemOrientation returns the value of the ::item-orientation property
	// which determines whether the labels are drawn beside the icons instead of
	// below.
	ItemOrientation() Orientation
	// ItemPadding returns the value of the ::item-padding property.
	ItemPadding() int
	// ItemRow gets the row in which the item @path is currently displayed. Row
	// numbers start at 0.
	ItemRow(path *TreePath) int
	// ItemWidth returns the value of the ::item-width property.
	ItemWidth() int
	// Margin returns the value of the ::margin property.
	Margin() int
	// MarkupColumn returns the column with markup text for @icon_view.
	MarkupColumn() int
	// Model returns the model the IconView is based on. Returns nil if the
	// model is unset.
	Model() TreeModel
	// PathAtPos gets the path for the icon at the given position.
	PathAtPos(x int, y int) *TreePath
	// PixbufColumn returns the column with pixbufs for @icon_view.
	PixbufColumn() int
	// Reorderable retrieves whether the user can reorder the list via
	// drag-and-drop. See gtk_icon_view_set_reorderable().
	Reorderable() bool
	// RowSpacing returns the value of the ::row-spacing property.
	RowSpacing() int
	// SelectionMode gets the selection mode of the @icon_view.
	SelectionMode() SelectionMode
	// Spacing returns the value of the ::spacing property.
	Spacing() int
	// TextColumn returns the column with text for @icon_view.
	TextColumn() int
	// TooltipColumn returns the column of @icon_view’s model which is being
	// used for displaying tooltips on @icon_view’s rows.
	TooltipColumn() int
	// TooltipContext: this function is supposed to be used in a
	// Widget::query-tooltip signal handler for IconView. The @x, @y and
	// @keyboard_tip values which are received in the signal handler, should be
	// passed to this function without modification.
	//
	// The return value indicates whether there is an icon view item at the
	// given coordinates (true) or not (false) for mouse tooltips. For keyboard
	// tooltips the item returned will be the cursor item. When true, then any
	// of @model, @path and @iter which have been provided will be set to point
	// to that row and the corresponding model.
	TooltipContext(x int, y int, keyboardTip bool) (TreeModel, *TreePath, TreeIter, bool)
	// VisibleRange sets @start_path and @end_path to be the first and last
	// visible path. Note that there may be invisible paths in between.
	//
	// Both paths should be freed with gtk_tree_path_free() after use.
	VisibleRange() (startPath *TreePath, endPath *TreePath, ok bool)
	// ItemActivated activates the item determined by @path.
	ItemActivated(path *TreePath)
	// PathIsSelected returns true if the icon pointed to by @path is currently
	// selected. If @path does not point to a valid location, false is returned.
	PathIsSelected(path *TreePath) bool
	// ScrollToPath moves the alignments of @icon_view to the position specified
	// by @path. @row_align determines where the row is placed, and @col_align
	// determines where @column is placed. Both are expected to be between 0.0
	// and 1.0. 0.0 means left/top alignment, 1.0 means right/bottom alignment,
	// 0.5 means center.
	//
	// If @use_align is false, then the alignment arguments are ignored, and the
	// tree does the minimum amount of work to scroll the item onto the screen.
	// This means that the item will be scrolled to the edge closest to its
	// current position. If the item is currently visible on the screen, nothing
	// is done.
	//
	// This function only works if the model is set, and @path is a valid row on
	// the model. If the model changes before the @icon_view is realized, the
	// centered path will be modified to reflect this change.
	ScrollToPath(path *TreePath, useAlign bool, rowAlign float32, colAlign float32)
	// SelectAll selects all the icons. @icon_view must has its selection mode
	// set to K_SELECTION_MULTIPLE.
	SelectAll()
	// SelectPath selects the row at @path.
	SelectPath(path *TreePath)
	// SelectedForeach calls a function for each selected icon. Note that the
	// model or selection cannot be modified from within this function.
	SelectedForeach(fn IconViewForeachFunc)
	// SetActivateOnSingleClick causes the IconView::item-activated signal to be
	// emitted on a single click instead of a double click.
	SetActivateOnSingleClick(single bool)
	// SetColumnSpacing sets the ::column-spacing property which specifies the
	// space which is inserted between the columns of the icon view.
	SetColumnSpacing(columnSpacing int)
	// SetColumns sets the ::columns property which determines in how many
	// columns the icons are arranged. If @columns is -1, the number of columns
	// will be chosen automatically to fill the available area.
	SetColumns(columns int)
	// SetCursor sets the current keyboard focus to be at @path, and selects it.
	// This is useful when you want to focus the user’s attention on a
	// particular item. If @cell is not nil, then focus is given to the cell
	// specified by it. Additionally, if @start_editing is true, then editing
	// should be started in the specified cell.
	//
	// This function is often followed by `gtk_widget_grab_focus (icon_view)` in
	// order to give keyboard focus to the widget. Please note that editing can
	// only happen when the widget is realized.
	SetCursor(path *TreePath, cell CellRenderer, startEditing bool)
	// SetDragDestItem sets the item that is highlighted for feedback.
	SetDragDestItem(path *TreePath, pos IconViewDropPosition)
	// SetItemOrientation sets the ::item-orientation property which determines
	// whether the labels are drawn beside the icons instead of below.
	SetItemOrientation(orientation Orientation)
	// SetItemPadding sets the IconView:item-padding property which specifies
	// the padding around each of the icon view’s items.
	SetItemPadding(itemPadding int)
	// SetItemWidth sets the ::item-width property which specifies the width to
	// use for each item. If it is set to -1, the icon view will automatically
	// determine a suitable item size.
	SetItemWidth(itemWidth int)
	// SetMargin sets the ::margin property which specifies the space which is
	// inserted at the top, bottom, left and right of the icon view.
	SetMargin(margin int)
	// SetMarkupColumn sets the column with markup information for @icon_view to
	// be @column. The markup column must be of type TYPE_STRING. If the markup
	// column is set to something, it overrides the text column set by
	// gtk_icon_view_set_text_column().
	SetMarkupColumn(column int)
	// SetModel sets the model for a IconView. If the @icon_view already has a
	// model set, it will remove it before setting the new model. If @model is
	// nil, then it will unset the old model.
	SetModel(model TreeModel)
	// SetPixbufColumn sets the column with pixbufs for @icon_view to be
	// @column. The pixbuf column must be of type K_TYPE_PIXBUF
	SetPixbufColumn(column int)
	// SetReorderable: this function is a convenience function to allow you to
	// reorder models that support the TreeDragSourceIface and the
	// TreeDragDestIface. Both TreeStore and ListStore support these. If
	// @reorderable is true, then the user can reorder the model by dragging and
	// dropping rows. The developer can listen to these changes by connecting to
	// the model's row_inserted and row_deleted signals. The reordering is
	// implemented by setting up the icon view as a drag source and destination.
	// Therefore, drag and drop can not be used in a reorderable view for any
	// other purpose.
	//
	// This function does not give you any degree of control over the order --
	// any reordering is allowed. If more control is needed, you should probably
	// handle drag and drop manually.
	SetReorderable(reorderable bool)
	// SetRowSpacing sets the ::row-spacing property which specifies the space
	// which is inserted between the rows of the icon view.
	SetRowSpacing(rowSpacing int)
	// SetSelectionMode sets the selection mode of the @icon_view.
	SetSelectionMode(mode SelectionMode)
	// SetSpacing sets the ::spacing property which specifies the space which is
	// inserted between the cells (i.e. the icon and the text) of an item.
	SetSpacing(spacing int)
	// SetTextColumn sets the column with text for @icon_view to be @column. The
	// text column must be of type TYPE_STRING.
	SetTextColumn(column int)
	// SetTooltipCell sets the tip area of @tooltip to the area which @cell
	// occupies in the item pointed to by @path. See also
	// gtk_tooltip_set_tip_area().
	//
	// See also gtk_icon_view_set_tooltip_column() for a simpler alternative.
	SetTooltipCell(tooltip Tooltip, path *TreePath, cell CellRenderer)
	// SetTooltipColumn: if you only plan to have simple (text-only) tooltips on
	// full items, you can use this function to have IconView handle these
	// automatically for you. @column should be set to the column in
	// @icon_view’s model containing the tooltip texts, or -1 to disable this
	// feature.
	//
	// When enabled, Widget:has-tooltip will be set to true and @icon_view will
	// connect a Widget::query-tooltip signal handler.
	//
	// Note that the signal handler sets the text with gtk_tooltip_set_markup(),
	// so &, <, etc have to be escaped in the text.
	SetTooltipColumn(column int)
	// SetTooltipItem sets the tip area of @tooltip to be the area covered by
	// the item at @path. See also gtk_icon_view_set_tooltip_column() for a
	// simpler alternative. See also gtk_tooltip_set_tip_area().
	SetTooltipItem(tooltip Tooltip, path *TreePath)
	// UnselectAll unselects all the icons.
	UnselectAll()
	// UnselectPath unselects the row at @path.
	UnselectPath(path *TreePath)
	// UnsetModelDragDest undoes the effect of
	// gtk_icon_view_enable_model_drag_dest(). Calling this method sets
	// IconView:reorderable to false.
	UnsetModelDragDest()
	// UnsetModelDragSource undoes the effect of
	// gtk_icon_view_enable_model_drag_source(). Calling this method sets
	// IconView:reorderable to false.
	UnsetModelDragSource()
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

// NewIconView creates a new IconView widget
func NewIconView() IconView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_icon_view_new()

	var _iconView IconView // out

	_iconView = WrapIconView(externglib.Take(unsafe.Pointer(_cret)))

	return _iconView
}

// NewIconViewWithArea creates a new IconView widget using the specified @area
// to layout cells inside the icons.
func NewIconViewWithArea(area CellArea) IconView {
	var _arg1 *C.GtkCellArea // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_icon_view_new_with_area(_arg1)

	var _iconView IconView // out

	_iconView = WrapIconView(externglib.Take(unsafe.Pointer(_cret)))

	return _iconView
}

// NewIconViewWithModel creates a new IconView widget with the model @model.
func NewIconViewWithModel(model TreeModel) IconView {
	var _arg1 *C.GtkTreeModel // out
	var _cret *C.GtkWidget    // in

	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_icon_view_new_with_model(_arg1)

	var _iconView IconView // out

	_iconView = WrapIconView(externglib.Take(unsafe.Pointer(_cret)))

	return _iconView
}

func (i iconView) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(i))
}

func (i iconView) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(i))
}

func (i iconView) AsCellLayout() CellLayout {
	return WrapCellLayout(gextras.InternObject(i))
}

func (i iconView) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(i))
}

func (i iconView) AsScrollable() Scrollable {
	return WrapScrollable(gextras.InternObject(i))
}

func (i iconView) EnableModelDragDest(formats *gdk.ContentFormats, actions gdk.DragAction) {
	var _arg0 *C.GtkIconView       // out
	var _arg1 *C.GdkContentFormats // out
	var _arg2 C.GdkDragAction      // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(formats))
	_arg2 = C.GdkDragAction(actions)

	C.gtk_icon_view_enable_model_drag_dest(_arg0, _arg1, _arg2)
}

func (i iconView) EnableModelDragSource(startButtonMask gdk.ModifierType, formats *gdk.ContentFormats, actions gdk.DragAction) {
	var _arg0 *C.GtkIconView       // out
	var _arg1 C.GdkModifierType    // out
	var _arg2 *C.GdkContentFormats // out
	var _arg3 C.GdkDragAction      // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GdkModifierType(startButtonMask)
	_arg2 = (*C.GdkContentFormats)(unsafe.Pointer(formats))
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
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))
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
	runtime.SetFinalizer(_path, func(v *TreePath) {
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
	runtime.SetFinalizer(_path, func(v *TreePath) {
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
	runtime.SetFinalizer(_path, func(v *TreePath) {
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
	runtime.SetFinalizer(_path, func(v *TreePath) {
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
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

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
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

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
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
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
	runtime.SetFinalizer(_path, func(v *TreePath) {
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
	runtime.SetFinalizer(_startPath, func(v *TreePath) {
		C.free(unsafe.Pointer(v))
	})
	_endPath = (*TreePath)(unsafe.Pointer(_arg2))
	runtime.SetFinalizer(_endPath, func(v *TreePath) {
		C.free(unsafe.Pointer(v))
	})
	if _cret != 0 {
		_ok = true
	}

	return _startPath, _endPath, _ok
}

func (i iconView) ItemActivated(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_icon_view_item_activated(_arg0, _arg1)
}

func (i iconView) PathIsSelected(path *TreePath) bool {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	_cret = C.gtk_icon_view_path_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i iconView) ScrollToPath(path *TreePath, useAlign bool, rowAlign float32, colAlign float32) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out
	var _arg2 C.gboolean     // out
	var _arg3 C.float        // out
	var _arg4 C.float        // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))
	if useAlign {
		_arg2 = C.TRUE
	}
	_arg3 = C.float(rowAlign)
	_arg4 = C.float(colAlign)

	C.gtk_icon_view_scroll_to_path(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (i iconView) SelectAll() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_select_all(_arg0)
}

func (i iconView) SelectPath(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_icon_view_select_path(_arg0, _arg1)
}

func (i iconView) SelectedForeach(fn IconViewForeachFunc) {
	var _arg0 *C.GtkIconView           // out
	var _arg1 C.GtkIconViewForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*[0]byte)(C.gotk4_IconViewForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_icon_view_selected_foreach(_arg0, _arg1, _arg2)
}

func (i iconView) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	if single {
		_arg1 = C.TRUE
	}

	C.gtk_icon_view_set_activate_on_single_click(_arg0, _arg1)
}

func (i iconView) SetColumnSpacing(columnSpacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(columnSpacing)

	C.gtk_icon_view_set_column_spacing(_arg0, _arg1)
}

func (i iconView) SetColumns(columns int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(columns)

	C.gtk_icon_view_set_columns(_arg0, _arg1)
}

func (i iconView) SetCursor(path *TreePath, cell CellRenderer, startEditing bool) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTreePath     // out
	var _arg2 *C.GtkCellRenderer // out
	var _arg3 C.gboolean         // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if startEditing {
		_arg3 = C.TRUE
	}

	C.gtk_icon_view_set_cursor(_arg0, _arg1, _arg2, _arg3)
}

func (i iconView) SetDragDestItem(path *TreePath, pos IconViewDropPosition) {
	var _arg0 *C.GtkIconView            // out
	var _arg1 *C.GtkTreePath            // out
	var _arg2 C.GtkIconViewDropPosition // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))
	_arg2 = C.GtkIconViewDropPosition(pos)

	C.gtk_icon_view_set_drag_dest_item(_arg0, _arg1, _arg2)
}

func (i iconView) SetItemOrientation(orientation Orientation) {
	var _arg0 *C.GtkIconView   // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_icon_view_set_item_orientation(_arg0, _arg1)
}

func (i iconView) SetItemPadding(itemPadding int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(itemPadding)

	C.gtk_icon_view_set_item_padding(_arg0, _arg1)
}

func (i iconView) SetItemWidth(itemWidth int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(itemWidth)

	C.gtk_icon_view_set_item_width(_arg0, _arg1)
}

func (i iconView) SetMargin(margin int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(margin)

	C.gtk_icon_view_set_margin(_arg0, _arg1)
}

func (i iconView) SetMarkupColumn(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_markup_column(_arg0, _arg1)
}

func (i iconView) SetModel(model TreeModel) {
	var _arg0 *C.GtkIconView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_icon_view_set_model(_arg0, _arg1)
}

func (i iconView) SetPixbufColumn(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_pixbuf_column(_arg0, _arg1)
}

func (i iconView) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_icon_view_set_reorderable(_arg0, _arg1)
}

func (i iconView) SetRowSpacing(rowSpacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(rowSpacing)

	C.gtk_icon_view_set_row_spacing(_arg0, _arg1)
}

func (i iconView) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.GtkSelectionMode(mode)

	C.gtk_icon_view_set_selection_mode(_arg0, _arg1)
}

func (i iconView) SetSpacing(spacing int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(spacing)

	C.gtk_icon_view_set_spacing(_arg0, _arg1)
}

func (i iconView) SetTextColumn(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_text_column(_arg0, _arg1)
}

func (i iconView) SetTooltipCell(tooltip Tooltip, path *TreePath, cell CellRenderer) {
	var _arg0 *C.GtkIconView     // out
	var _arg1 *C.GtkTooltip      // out
	var _arg2 *C.GtkTreePath     // out
	var _arg3 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(path))
	_arg3 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_icon_view_set_tooltip_cell(_arg0, _arg1, _arg2, _arg3)
}

func (i iconView) SetTooltipColumn(column int) {
	var _arg0 *C.GtkIconView // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = C.int(column)

	C.gtk_icon_view_set_tooltip_column(_arg0, _arg1)
}

func (i iconView) SetTooltipItem(tooltip Tooltip, path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTooltip  // out
	var _arg2 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg2 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_icon_view_set_tooltip_item(_arg0, _arg1, _arg2)
}

func (i iconView) UnselectAll() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unselect_all(_arg0)
}

func (i iconView) UnselectPath(path *TreePath) {
	var _arg0 *C.GtkIconView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_icon_view_unselect_path(_arg0, _arg1)
}

func (i iconView) UnsetModelDragDest() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unset_model_drag_dest(_arg0)
}

func (i iconView) UnsetModelDragSource() {
	var _arg0 *C.GtkIconView // out

	_arg0 = (*C.GtkIconView)(unsafe.Pointer(i.Native()))

	C.gtk_icon_view_unset_model_drag_source(_arg0)
}
