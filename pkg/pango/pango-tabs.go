// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_tab_align_get_type()), F: marshalTabAlign},
		{T: externglib.Type(C.pango_tab_array_get_type()), F: marshalTabArray},
	})
}

// TabAlign: `PangoTabAlign` specifies where a tab stop appears relative to the
// text.
type TabAlign int

const (
	// left: the tab stop appears to the left of the text.
	TabAlignLeft TabAlign = 0
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	return TabAlign(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TabArray: a `PangoTabArray` contains an array of tab stops.
//
// `PangoTabArray` can be used to set tab stops in a `PangoLayout`. Each tab
// stop has an alignment and a position.
type TabArray C.PangoTabArray

// WrapTabArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTabArray(ptr unsafe.Pointer) *TabArray {
	return (*TabArray)(ptr)
}

func marshalTabArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*TabArray)(unsafe.Pointer(b)), nil
}

// NewTabArray constructs a struct TabArray.
func NewTabArray(initialSize int, positionsInPixels bool) *TabArray {
	var _arg1 C.gint           // out
	var _arg2 C.gboolean       // out
	var _cret *C.PangoTabArray // in

	_arg1 = C.gint(initialSize)
	if positionsInPixels {
		_arg2 = C.TRUE
	}

	_cret = C.pango_tab_array_new(_arg1, _arg2)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_tabArray, func(v **TabArray) {
		C.free(unsafe.Pointer(v))
	})

	return _tabArray
}

// Native returns the underlying C source pointer.
func (t *TabArray) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// Copy sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) Copy() *TabArray {
	var _arg0 *C.PangoTabArray // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(s.Native()))

	_cret = C.pango_tab_array_copy(_arg0)

	var _tabArray *TabArray // out

	_tabArray = (*TabArray)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_tabArray, func(v **TabArray) {
		C.free(unsafe.Pointer(v))
	})

	return _tabArray
}

// Free sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) Free() {
	var _arg0 *C.PangoTabArray // out

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))

	C.pango_tab_array_free(_arg0)
}

// PositionsInPixels sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) PositionsInPixels() bool {
	var _arg0 *C.PangoTabArray // out
	var _cret C.gboolean       // in

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))

	_cret = C.pango_tab_array_get_positions_in_pixels(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) Size() int {
	var _arg0 *C.PangoTabArray // out
	var _cret C.gint           // in

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))

	_cret = C.pango_tab_array_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Tab sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) Tab(tabIndex int) (TabAlign, int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out
	var _arg2 C.PangoTabAlign  // in
	var _arg3 C.gint           // in

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(tabIndex)

	C.pango_tab_array_get_tab(_arg0, _arg1, &_arg2, &_arg3)

	var _alignment TabAlign // out
	var _location int       // out

	_alignment = TabAlign(_arg2)
	_location = int(_arg3)

	return _alignment, _location
}

// Resize sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) Resize(newSize int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(newSize)

	C.pango_tab_array_resize(_arg0, _arg1)
}

// SetTab sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) SetTab(tabIndex int, alignment TabAlign, location int) {
	var _arg0 *C.PangoTabArray // out
	var _arg1 C.gint           // out
	var _arg2 C.PangoTabAlign  // out
	var _arg3 C.gint           // out

	_arg0 = (*C.PangoTabArray)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(tabIndex)
	_arg2 = C.PangoTabAlign(alignment)
	_arg3 = C.gint(location)

	C.pango_tab_array_set_tab(_arg0, _arg1, _arg2, _arg3)
}
