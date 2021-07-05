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
		{T: externglib.Type(C.gtk_grid_view_get_type()), F: marshalGridView},
	})
}

// GridView: `GtkGridView` presents a large dynamic grid of items.
//
// `GtkGridView` uses its factory to generate one child widget for each visible
// item and shows them in a grid. The orientation of the grid view determines if
// the grid reflows vertically or horizontally.
//
// `GtkGridView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// [property@Gtk.GridView:enable-rubberband].
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
//
// CSS nodes
//
// “` gridview ├── child │ ├── child │ ┊ ╰── [rubberband] “`
//
// `GtkGridView` uses a single CSS node with name gridview. Each child uses a
// single CSS node with name child. For rubberband selection, a subnode with
// name rubberband is used.
//
//
// Accessibility
//
// `GtkGridView` uses the GTK_ACCESSIBLE_ROLE_GRID role, and the items use the
// GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type GridView interface {
	ListBase

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable
	// AsScrollable casts the class to the Scrollable interface.
	AsScrollable() Scrollable

	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband() bool
	// Factory gets the factory that's currently used to populate list items.
	Factory() ListItemFactory
	// MaxColumns gets the maximum number of columns that the grid will use.
	MaxColumns() uint
	// MinColumns gets the minimum number of columns that the grid will use.
	MinColumns() uint
	// SingleClickActivate returns whether items will be activated on single
	// click and selected on hover.
	SingleClickActivate() bool
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(enableRubberband bool)
	// SetFactory sets the `GtkListItemFactory` to use for populating list
	// items.
	SetFactory(factory ListItemFactory)
	// SetMaxColumns sets the maximum number of columns to use.
	//
	// This number must be at least 1.
	//
	// If @max_columns is smaller than the minimum set via
	// [method@Gtk.GridView.set_min_columns], that value is used instead.
	SetMaxColumns(maxColumns uint)
	// SetMinColumns sets the minimum number of columns to use.
	//
	// This number must be at least 1.
	//
	// If @min_columns is smaller than the minimum set via
	// [method@Gtk.GridView.set_max_columns], that value is ignored.
	SetMinColumns(minColumns uint)
	// SetSingleClickActivate sets whether items should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(singleClickActivate bool)
}

// gridView implements the GridView class.
type gridView struct {
	ListBase
}

// WrapGridView wraps a GObject to the right type. It is
// primarily used internally.
func WrapGridView(obj *externglib.Object) GridView {
	return gridView{
		ListBase: WrapListBase(obj),
	}
}

func marshalGridView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGridView(obj), nil
}

func (g gridView) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(g))
}

func (g gridView) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(g))
}

func (g gridView) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(g))
}

func (g gridView) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(g))
}

func (g gridView) AsScrollable() Scrollable {
	return WrapScrollable(gextras.InternObject(g))
}

func (s gridView) EnableRubberband() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_grid_view_get_enable_rubberband(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s gridView) Factory() ListItemFactory {
	var _arg0 *C.GtkGridView        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_grid_view_get_factory(_arg0)

	var _listItemFactory ListItemFactory // out

	_listItemFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListItemFactory)

	return _listItemFactory
}

func (s gridView) MaxColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_grid_view_get_max_columns(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s gridView) MinColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_grid_view_get_min_columns(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s gridView) SingleClickActivate() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_grid_view_get_single_click_activate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s gridView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_enable_rubberband(_arg0, _arg1)
}

func (s gridView) SetFactory(factory ListItemFactory) {
	var _arg0 *C.GtkGridView        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_grid_view_set_factory(_arg0, _arg1)
}

func (s gridView) SetMaxColumns(maxColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(maxColumns)

	C.gtk_grid_view_set_max_columns(_arg0, _arg1)
}

func (s gridView) SetMinColumns(minColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(minColumns)

	C.gtk_grid_view_set_min_columns(_arg0, _arg1)
}

func (s gridView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_single_click_activate(_arg0, _arg1)
}
