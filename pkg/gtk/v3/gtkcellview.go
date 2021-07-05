// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_cell_view_get_type()), F: marshalCellView},
	})
}

// CellView displays a single row of a TreeModel using a CellArea and
// CellAreaContext. A CellAreaContext can be provided to the CellView at
// construction time in order to keep the cellview in context of a group of cell
// views, this ensures that the renderers displayed will be properly aligned
// with eachother (like the aligned cells in the menus of ComboBox).
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

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsCellLayout casts the class to the CellLayout interface.
	AsCellLayout() CellLayout
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// DisplayedRow returns a TreePath referring to the currently displayed row.
	// If no row is currently displayed, nil is returned.
	DisplayedRow() *TreePath
	// DrawSensitive gets whether @cell_view is configured to draw all of its
	// cells in a sensitive state.
	DrawSensitive() bool
	// FitModel gets whether @cell_view is configured to request space to fit
	// the entire TreeModel.
	FitModel() bool
	// Model returns the model for @cell_view. If no model is used nil is
	// returned.
	Model() TreeModel
	// SizeOfRow sets @requisition to the size needed by @cell_view to display
	// the model row pointed to by @path.
	//
	// Deprecated: since version 3.0.
	SizeOfRow(path *TreePath) (Requisition, bool)
	// SetBackgroundColor sets the background color of @view.
	//
	// Deprecated: since version 3.4.
	SetBackgroundColor(color *gdk.Color)
	// SetBackgroundRGBA sets the background color of @cell_view.
	SetBackgroundRGBA(rgba *gdk.RGBA)
	// SetDisplayedRow sets the row of the model that is currently displayed by
	// the CellView. If the path is unset, then the contents of the cellview
	// “stick” at their last value; this is not normally a desired result, but
	// may be a needed intermediate state if say, the model for the CellView
	// becomes temporarily empty.
	SetDisplayedRow(path *TreePath)
	// SetDrawSensitive sets whether @cell_view should draw all of its cells in
	// a sensitive state, this is used by ComboBox menus to ensure that rows
	// with insensitive cells that contain children appear sensitive in the
	// parent menu item.
	SetDrawSensitive(drawSensitive bool)
	// SetFitModel sets whether @cell_view should request space to fit the
	// entire TreeModel.
	//
	// This is used by ComboBox to ensure that the cell view displayed on the
	// combo box’s button always gets enough space and does not resize when
	// selection changes.
	SetFitModel(fitModel bool)
	// SetModel sets the model for @cell_view. If @cell_view already has a model
	// set, it will remove it before setting the new model. If @model is nil,
	// then it will unset the old model.
	SetModel(model TreeModel)
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

// NewCellView creates a new CellView widget.
func NewCellView() CellView {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_cell_view_new()

	var _cellView CellView // out

	_cellView = WrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithContext creates a new CellView widget with a specific CellArea
// to layout cells and a specific CellAreaContext.
//
// Specifying the same context for a handfull of cells lets the underlying area
// synchronize the geometry for those cells, in this way alignments with
// cellviews for other rows are possible.
func NewCellViewWithContext(area CellArea, context CellAreaContext) CellView {
	var _arg1 *C.GtkCellArea        // out
	var _arg2 *C.GtkCellAreaContext // out
	var _cret *C.GtkWidget          // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))
	_arg2 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_cell_view_new_with_context(_arg1, _arg2)

	var _cellView CellView // out

	_cellView = WrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithMarkup creates a new CellView widget, adds a CellRendererText
// to it, and makes it show @markup. The text can be marked up with the [Pango
// text markup language][PangoMarkupFormat].
func NewCellViewWithMarkup(markup string) CellView {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_markup(_arg1)

	var _cellView CellView // out

	_cellView = WrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithPixbuf creates a new CellView widget, adds a
// CellRendererPixbuf to it, and makes it show @pixbuf.
func NewCellViewWithPixbuf(pixbuf gdkpixbuf.Pixbuf) CellView {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_cell_view_new_with_pixbuf(_arg1)

	var _cellView CellView // out

	_cellView = WrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

// NewCellViewWithText creates a new CellView widget, adds a CellRendererText to
// it, and makes it show @text.
func NewCellViewWithText(text string) CellView {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_cell_view_new_with_text(_arg1)

	var _cellView CellView // out

	_cellView = WrapCellView(externglib.Take(unsafe.Pointer(_cret)))

	return _cellView
}

func (c cellView) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(c))
}

func (c cellView) AsCellLayout() CellLayout {
	return WrapCellLayout(gextras.InternObject(c))
}

func (c cellView) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(c))
}

func (c cellView) DisplayedRow() *TreePath {
	var _arg0 *C.GtkCellView // out
	var _cret *C.GtkTreePath // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_view_get_displayed_row(_arg0)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_treePath, func(v *TreePath) {
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

func (c cellView) SizeOfRow(path *TreePath) (Requisition, bool) {
	var _arg0 *C.GtkCellView   // out
	var _arg1 *C.GtkTreePath   // out
	var _arg2 C.GtkRequisition // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	_cret = C.gtk_cell_view_get_size_of_row(_arg0, _arg1, &_arg2)

	var _requisition Requisition // out
	var _ok bool                 // out

	{
		var refTmpIn *C.GtkRequisition
		var refTmpOut *Requisition

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Requisition)(unsafe.Pointer(refTmpIn))

		_requisition = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _requisition, _ok
}

func (c cellView) SetBackgroundColor(color *gdk.Color) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GdkColor    // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_cell_view_set_background_color(_arg0, _arg1)
}

func (c cellView) SetBackgroundRGBA(rgba *gdk.RGBA) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GdkRGBA     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(rgba))

	C.gtk_cell_view_set_background_rgba(_arg0, _arg1)
}

func (c cellView) SetDisplayedRow(path *TreePath) {
	var _arg0 *C.GtkCellView // out
	var _arg1 *C.GtkTreePath // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_cell_view_set_displayed_row(_arg0, _arg1)
}

func (c cellView) SetDrawSensitive(drawSensitive bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	if drawSensitive {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_draw_sensitive(_arg0, _arg1)
}

func (c cellView) SetFitModel(fitModel bool) {
	var _arg0 *C.GtkCellView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	if fitModel {
		_arg1 = C.TRUE
	}

	C.gtk_cell_view_set_fit_model(_arg0, _arg1)
}

func (c cellView) SetModel(model TreeModel) {
	var _arg0 *C.GtkCellView  // out
	var _arg1 *C.GtkTreeModel // out

	_arg0 = (*C.GtkCellView)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_cell_view_set_model(_arg0, _arg1)
}
