// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.gtk_cell_area_context_get_type()), F: marshalCellAreaContexter},
	})
}

// CellAreaContextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellAreaContextOverrider interface {
	// Allocate allocates a width and/or a height for all rows which are to be
	// rendered with context.
	//
	// Usually allocation is performed only horizontally or sometimes vertically
	// since a group of rows are usually rendered side by side vertically or
	// horizontally and share either the same width or the same height.
	// Sometimes they are allocated in both horizontal and vertical orientations
	// producing a homogeneous effect of the rows. This is generally the case
	// for TreeView when TreeView:fixed-height-mode is enabled.
	//
	// Since 3.0
	Allocate(width int, height int)
	// PreferredHeightForWidth gets the accumulative preferred height for width
	// for all rows which have been requested for the same said width with this
	// context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	PreferredHeightForWidth(width int) (minimumHeight int, naturalHeight int)
	// PreferredWidthForHeight gets the accumulative preferred width for height
	// for all rows which have been requested for the same said height with this
	// context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	PreferredWidthForHeight(height int) (minimumWidth int, naturalWidth int)
	// Reset resets any previously cached request and allocation data.
	//
	// When underlying TreeModel data changes its important to reset the context
	// if the content size is allowed to shrink. If the content size is only
	// allowed to grow (this is usually an option for views rendering large data
	// stores as a measure of optimization), then only the row that changed or
	// was inserted needs to be (re)requested with
	// gtk_cell_area_get_preferred_width().
	//
	// When the new overall size of the context requires that the allocated size
	// changes (or whenever this allocation changes at all), the variable row
	// sizes need to be re-requested for every row.
	//
	// For instance, if the rows are displayed all with the same width from top
	// to bottom then a change in the allocated width necessitates a
	// recalculation of all the displayed row heights using
	// gtk_cell_area_get_preferred_height_for_width().
	//
	// Since 3.0
	Reset()
}

// CellAreaContext object is created by a given CellArea implementation via its
// CellAreaClass.create_context() virtual method and is used to store cell sizes
// and alignments for a series of TreeModel rows that are requested and rendered
// in the same context.
//
// CellLayout widgets can create any number of contexts in which to request and
// render groups of data rows. However, it’s important that the same context
// which was used to request sizes for a given TreeModel row also be used for
// the same row when calling other CellArea APIs such as gtk_cell_area_render()
// and gtk_cell_area_event().
type CellAreaContext struct {
	*externglib.Object
}

func WrapCellAreaContext(obj *externglib.Object) *CellAreaContext {
	return &CellAreaContext{
		Object: obj,
	}
}

func marshalCellAreaContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAreaContext(obj), nil
}

// Allocate allocates a width and/or a height for all rows which are to be
// rendered with context.
//
// Usually allocation is performed only horizontally or sometimes vertically
// since a group of rows are usually rendered side by side vertically or
// horizontally and share either the same width or the same height. Sometimes
// they are allocated in both horizontal and vertical orientations producing a
// homogeneous effect of the rows. This is generally the case for TreeView when
// TreeView:fixed-height-mode is enabled.
//
// Since 3.0
func (context *CellAreaContext) Allocate(width int, height int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // out

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	C.gtk_cell_area_context_allocate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Allocation fetches the current allocation size for context.
//
// If the context was not allocated in width or height, or if the context was
// recently reset with gtk_cell_area_context_reset(), the returned value will be
// -1.
func (context *CellAreaContext) Allocation() (width int, height int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // in
	var _arg2 C.gint                // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	C.gtk_cell_area_context_get_allocation(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _width int  // out
	var _height int // out

	_width = int(_arg1)
	_height = int(_arg2)

	return _width, _height
}

// Area fetches the CellArea this context was created by.
//
// This is generally unneeded by layouting widgets; however, it is important for
// the context implementation itself to fetch information about the area it is
// being used for.
//
// For instance at CellAreaContextClass.allocate() time it’s important to know
// details about any cell spacing that the CellArea is configured with in order
// to compute a proper allocation.
func (context *CellAreaContext) Area() CellAreaer {
	var _arg0 *C.GtkCellAreaContext // out
	var _cret *C.GtkCellArea        // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_cell_area_context_get_area(_arg0)
	runtime.KeepAlive(context)

	var _cellArea CellAreaer // out

	_cellArea = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(CellAreaer)

	return _cellArea
}

// PreferredHeight gets the accumulative preferred height for all rows which
// have been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are 0.
func (context *CellAreaContext) PreferredHeight() (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // in
	var _arg2 C.gint                // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	C.gtk_cell_area_context_get_preferred_height(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg1)
	_naturalHeight = int(_arg2)

	return _minimumHeight, _naturalHeight
}

// PreferredHeightForWidth gets the accumulative preferred height for width for
// all rows which have been requested for the same said width with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are -1.
func (context *CellAreaContext) PreferredHeightForWidth(width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // in
	var _arg3 C.gint                // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(width)

	C.gtk_cell_area_context_get_preferred_height_for_width(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(width)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = int(_arg2)
	_naturalHeight = int(_arg3)

	return _minimumHeight, _naturalHeight
}

// PreferredWidth gets the accumulative preferred width for all rows which have
// been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are 0.
func (context *CellAreaContext) PreferredWidth() (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // in
	var _arg2 C.gint                // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	C.gtk_cell_area_context_get_preferred_width(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg1)
	_naturalWidth = int(_arg2)

	return _minimumWidth, _naturalWidth
}

// PreferredWidthForHeight gets the accumulative preferred width for height for
// all rows which have been requested for the same said height with this
// context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are -1.
func (context *CellAreaContext) PreferredWidthForHeight(height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // in
	var _arg3 C.gint                // in

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(height)

	C.gtk_cell_area_context_get_preferred_width_for_height(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(height)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = int(_arg2)
	_naturalWidth = int(_arg3)

	return _minimumWidth, _naturalWidth
}

// PushPreferredHeight causes the minimum and/or natural height to grow if the
// new proposed sizes exceed the current minimum and natural height.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested height
// over a series of gtk_cell_area_get_preferred_height() requests.
func (context *CellAreaContext) PushPreferredHeight(minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // out

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(minimumHeight)
	_arg2 = C.gint(naturalHeight)

	C.gtk_cell_area_context_push_preferred_height(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumHeight)
	runtime.KeepAlive(naturalHeight)
}

// PushPreferredWidth causes the minimum and/or natural width to grow if the new
// proposed sizes exceed the current minimum and natural width.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested width
// over a series of gtk_cell_area_get_preferred_width() requests.
func (context *CellAreaContext) PushPreferredWidth(minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellAreaContext // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // out

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg1 = C.gint(minimumWidth)
	_arg2 = C.gint(naturalWidth)

	C.gtk_cell_area_context_push_preferred_width(_arg0, _arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumWidth)
	runtime.KeepAlive(naturalWidth)
}

// Reset resets any previously cached request and allocation data.
//
// When underlying TreeModel data changes its important to reset the context if
// the content size is allowed to shrink. If the content size is only allowed to
// grow (this is usually an option for views rendering large data stores as a
// measure of optimization), then only the row that changed or was inserted
// needs to be (re)requested with gtk_cell_area_get_preferred_width().
//
// When the new overall size of the context requires that the allocated size
// changes (or whenever this allocation changes at all), the variable row sizes
// need to be re-requested for every row.
//
// For instance, if the rows are displayed all with the same width from top to
// bottom then a change in the allocated width necessitates a recalculation of
// all the displayed row heights using
// gtk_cell_area_get_preferred_height_for_width().
//
// Since 3.0
func (context *CellAreaContext) Reset() {
	var _arg0 *C.GtkCellAreaContext // out

	_arg0 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))

	C.gtk_cell_area_context_reset(_arg0)
	runtime.KeepAlive(context)
}
