// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
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
		{T: externglib.Type(C.pango_glyph_item_get_type()), F: marshalGlyphItem},
		{T: externglib.Type(C.pango_glyph_item_iter_get_type()), F: marshalGlyphItemIter},
	})
}

// GlyphItem: `PangoGlyphItem` is a pair of a `PangoItem` and the glyphs
// resulting from shaping the items text.
//
// As an example of the usage of `PangoGlyphItem`, the results of shaping text
// with `PangoLayout` is a list of `PangoLayoutLine`, each of which contains a
// list of `PangoGlyphItem`.
type GlyphItem struct {
	native C.PangoGlyphItem
}

func marshalGlyphItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*GlyphItem)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *GlyphItem) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// Copy: make a deep copy of an existing `PangoGlyphItem` structure.
func (orig *GlyphItem) Copy() *GlyphItem {
	var _arg0 *C.PangoGlyphItem // out
	var _cret *C.PangoGlyphItem // in

	_arg0 = (*C.PangoGlyphItem)(unsafe.Pointer(orig))

	_cret = C.pango_glyph_item_copy(_arg0)

	var _glyphItem *GlyphItem // out

	_glyphItem = (*GlyphItem)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_glyphItem, func(v *GlyphItem) {
		C.pango_glyph_item_free((*C.PangoGlyphItem)(unsafe.Pointer(v)))
	})

	return _glyphItem
}

// Free frees a `PangoGlyphItem` and resources to which it points.
func (glyphItem *GlyphItem) free() {
	var _arg0 *C.PangoGlyphItem // out

	_arg0 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem))

	C.pango_glyph_item_free(_arg0)
}

// Split modifies @orig to cover only the text after @split_index, and returns a
// new item that covers the text before @split_index that used to be in @orig.
//
// You can think of @split_index as the length of the returned item.
// @split_index may not be 0, and it may not be greater than or equal to the
// length of @orig (that is, there must be at least one byte assigned to each
// item, you can't create a zero-length item).
//
// This function is similar in function to pango_item_split() (and uses it
// internally.)
func (orig *GlyphItem) Split(text string, splitIndex int) *GlyphItem {
	var _arg0 *C.PangoGlyphItem // out
	var _arg1 *C.char           // out
	var _arg2 C.int             // out
	var _cret *C.PangoGlyphItem // in

	_arg0 = (*C.PangoGlyphItem)(unsafe.Pointer(orig))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(splitIndex)

	_cret = C.pango_glyph_item_split(_arg0, _arg1, _arg2)

	var _glyphItem *GlyphItem // out

	_glyphItem = (*GlyphItem)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_glyphItem, func(v *GlyphItem) {
		C.pango_glyph_item_free((*C.PangoGlyphItem)(unsafe.Pointer(v)))
	})

	return _glyphItem
}

// GlyphItemIter: `PangoGlyphItemIter` is an iterator over the clusters in a
// `PangoGlyphItem`.
//
// The *forward direction* of the iterator is the logical direction of text.
// That is, with increasing @start_index and @start_char values. If @glyph_item
// is right-to-left (that is, if `glyph_item->item->analysis.level` is odd),
// then @start_glyph decreases as the iterator moves forward. Moreover, in
// right-to-left cases, @start_glyph is greater than @end_glyph.
//
// An iterator should be initialized using either
// pango_glyph_item_iter_init_start() or pango_glyph_item_iter_init_end(), for
// forward and backward iteration respectively, and walked over using any
// desired mixture of pango_glyph_item_iter_next_cluster() and
// pango_glyph_item_iter_prev_cluster().
//
// A common idiom for doing a forward iteration over the clusters is:
//
// “` PangoGlyphItemIter cluster_iter; gboolean have_cluster;
//
// for (have_cluster = pango_glyph_item_iter_init_start (&cluster_iter,
// glyph_item, text); have_cluster; have_cluster =
// pango_glyph_item_iter_next_cluster (&cluster_iter)) { ... } “`
//
// Note that @text is the start of the text for layout, which is then indexed by
// `glyph_item->item->offset` to get to the text of @glyph_item. The
// @start_index and @end_index values can directly index into @text. The
// @start_glyph, @end_glyph, @start_char, and @end_char values however are
// zero-based for the @glyph_item. For each cluster, the item pointed at by the
// start variables is included in the cluster while the one pointed at by end
// variables is not.
//
// None of the members of a `PangoGlyphItemIter` should be modified manually.
type GlyphItemIter struct {
	native C.PangoGlyphItemIter
}

func marshalGlyphItemIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*GlyphItemIter)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (g *GlyphItemIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&g.native)
}

// Copy: make a shallow copy of an existing `PangoGlyphItemIter` structure.
func (orig *GlyphItemIter) Copy() *GlyphItemIter {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret *C.PangoGlyphItemIter // in

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(orig))

	_cret = C.pango_glyph_item_iter_copy(_arg0)

	var _glyphItemIter *GlyphItemIter // out

	_glyphItemIter = (*GlyphItemIter)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_glyphItemIter, func(v *GlyphItemIter) {
		C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(unsafe.Pointer(v)))
	})

	return _glyphItemIter
}

// Free frees a `PangoGlyphItem`Iter.
func (iter *GlyphItemIter) free() {
	var _arg0 *C.PangoGlyphItemIter // out

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(iter))

	C.pango_glyph_item_iter_free(_arg0)
}

// InitEnd initializes a `PangoGlyphItemIter` structure to point to the last
// cluster in a glyph item.
//
// See `PangoGlyphItemIter` for details of cluster orders.
func (iter *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _arg1 *C.PangoGlyphItem     // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(iter))
	_arg1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.pango_glyph_item_iter_init_end(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InitStart initializes a `PangoGlyphItemIter` structure to point to the first
// cluster in a glyph item.
//
// See `PangoGlyphItemIter` for details of cluster orders.
func (iter *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _arg1 *C.PangoGlyphItem     // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(iter))
	_arg1 = (*C.PangoGlyphItem)(unsafe.Pointer(glyphItem))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.pango_glyph_item_iter_init_start(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NextCluster advances the iterator to the next cluster in the glyph item.
//
// See `PangoGlyphItemIter` for details of cluster orders.
func (iter *GlyphItemIter) NextCluster() bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(iter))

	_cret = C.pango_glyph_item_iter_next_cluster(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PrevCluster moves the iterator to the preceding cluster in the glyph item.
// See `PangoGlyphItemIter` for details of cluster orders.
func (iter *GlyphItemIter) PrevCluster() bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(unsafe.Pointer(iter))

	_cret = C.pango_glyph_item_iter_prev_cluster(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
