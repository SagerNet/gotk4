// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_cell_view_get_type()), F: marshalCellView},
	})
}

// CellView: a widget displaying a single row of a GtkTreeModel
//
// A CellView displays a single row of a TreeModel using a CellArea and
// CellAreaContext. A CellAreaContext can be provided to the CellView at
// construction time in order to keep the cellview in context of a group of cell
// views, this ensures that the renderers displayed will be properly aligned
// with each other (like the aligned cells in the menus of ComboBox).
//
// CellView is Orientable in order to decide in which orientation the underlying
// CellAreaContext should be allocated. Taking the ComboBox menu as an example,
// cellviews should be oriented horizontally if the menus are listed
// top-to-bottom and thus all share the same width but may have separate
// individual heights (left-to-right menus should be allocated vertically since
// they all share the same height but may have variable widths).
//
//
// CSS nodes
//
// GtkCellView has a single CSS node with name cellview.
type CellView interface {
	Widget
	CellLayout
	Orientable

	// DisplayedRow:
	DisplayedRow() *TreePath
	// DrawSensitive:
	DrawSensitive() bool
	// FitModel:
	FitModel() bool
	// Model:
	Model() TreeModel
	// SetDisplayedRowCellView:
	SetDisplayedRowCellView(path *TreePath)
	// SetDrawSensitiveCellView:
	SetDrawSensitiveCellView(drawSensitive bool)
	// SetFitModelCellView:
	SetFitModelCellView(fitModel bool)
	// SetModelCellView:
	SetModelCellView(model TreeModel)
}

// cellView implements the CellView class.
type cellView struct {
	Widget
}

// WrapCellView wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellView(obj *externglib.Object) CellView {
	return cellView{
		Widget: WrapWidget(obj),
	}
}

func marshalCellView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellView(obj), nil
}

// NewCellView:
func NewCellView() CellView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_cell_view_new()

	var _cellView CellView // out

	_cellView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellView)

	return _cellView
}

// NewCellViewWithContext:
func NewCellViewWithContext(area CellArea, context CellAreaContext) CellView {
	var _arg1 *C.GtkCellArea        // out
	var _arg2 *C.GtkCellAreaContext // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))
	_arg2 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_cell_view_new_with_context(_arg1, _arg2)

	var _cellView CellView // out

	_cellView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellView)

	return _cellView
}

// NewCellViewWithMarkup:
func NewCellViewWithMarkup(markup string) CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_markup(_arg1)

	var _cellView CellView // out

	_cellView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellView)

	return _cellView
}

// NewCellViewWithText:
func NewCellViewWithText(text string) CellView {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_text(_arg1)

	var _cellView CellView // out

	_cellView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellView)

	return _cellView
}

// NewCellViewWithTexture:
func NewCellViewWithTexture(texture gdk.Texture) CellView {
	var _arg1 *C.GdkTexture // out
	var _cret *C.GtkWidget  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gtk_cell_view_new_with_texture(_arg1)

	var _cellView CellView // out

	_cellView = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellView)

	return _cellView
}

func (c cellView) DisplayedRow() *TreePath {
	var _arg0 *C.GtkCellView // out
	var _cret *C.GtkTreePath // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_view_get_displayed_row(_arg0)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_treePath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

func (c cellView) DrawSensitive() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_view_get_draw_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c cellView) FitModel() bool {
	var _arg0 *C.GtkCellView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_view_get_fit_model(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c cellView) Model() TreeModel {
	var _arg0 *C.GtkCellView  // out
	var _cret *C.GtkTreeModel // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_view_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

func (c cellView) SetDisplayedRowCellView(path *TreePath) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_cell_view_set_displayed_row(_arg0, _arg1)
}

func (c cellView) SetDrawSensitiveCellView(drawSensitive bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	if drawSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_draw_sensitive(_arg0, _arg1)
}

func (c cellView) SetFitModelCellView(fitModel bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	if fitModel {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_fit_model(_arg0, _arg1)
}

func (c cellView) SetModelCellView(model TreeModel) {
	var _arg0 *C.GtkCellView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_cell_view_set_model(_arg0, _arg1)
}

func (s cellView) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s cellView) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s cellView) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s cellView) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s cellView) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s cellView) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s cellView) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b cellView) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (c cellView) AddAttribute(cell CellRenderer, attribute string, column int) {
	WrapCellLayout(gextras.InternObject(c)).AddAttribute(cell, attribute, column)
}

func (c cellView) Clear() {
	WrapCellLayout(gextras.InternObject(c)).Clear()
}

func (c cellView) ClearAttributes(cell CellRenderer) {
	WrapCellLayout(gextras.InternObject(c)).ClearAttributes(cell)
}

func (c cellView) Area() CellArea {
	return WrapCellLayout(gextras.InternObject(c)).Area()
}

func (c cellView) PackEnd(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackEnd(cell, expand)
}

func (c cellView) PackStart(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackStart(cell, expand)
}

func (c cellView) Reorder(cell CellRenderer, position int) {
	WrapCellLayout(gextras.InternObject(c)).Reorder(cell, position)
}

func (o cellView) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o cellView) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
