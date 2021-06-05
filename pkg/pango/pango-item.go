// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_item_get_type()), F: marshalItem},
	})
}

// Analysis: the `PangoAnalysis` structure stores information about the
// properties of a segment of text.
type Analysis struct {
	native C.PangoAnalysis
}

// WrapAnalysis wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAnalysis(ptr unsafe.Pointer) *Analysis {
	if ptr == nil {
		return nil
	}

	return (*Analysis)(ptr)
}

func marshalAnalysis(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAnalysis(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *Analysis) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// ShapeEngine gets the field inside the struct.
func (a *Analysis) ShapeEngine() interface{} {
	v = C.gpointer(a.native.shape_engine)
}

// LangEngine gets the field inside the struct.
func (a *Analysis) LangEngine() interface{} {
	v = C.gpointer(a.native.lang_engine)
}

// Font gets the field inside the struct.
func (a *Analysis) Font() Font {
	v = gextras.CastObject(externglib.Take(unsafe.Pointer(a.native.font.Native()))).(Font)
}

// Level gets the field inside the struct.
func (a *Analysis) Level() byte {
	v = C.guint8(a.native.level)
}

// Gravity gets the field inside the struct.
func (a *Analysis) Gravity() byte {
	v = C.guint8(a.native.gravity)
}

// Flags gets the field inside the struct.
func (a *Analysis) Flags() byte {
	v = C.guint8(a.native.flags)
}

// Script gets the field inside the struct.
func (a *Analysis) Script() byte {
	v = C.guint8(a.native.script)
}

// Language gets the field inside the struct.
func (a *Analysis) Language() *Language {
	v = WrapLanguage(unsafe.Pointer(a.native.language))
}

// ExtraAttrs gets the field inside the struct.
func (a *Analysis) ExtraAttrs() *glib.SList {
	v = glib.WrapSList(unsafe.Pointer(a.native.extra_attrs))
}

// Item: the `PangoItem` structure stores information about a segment of text.
//
// You typically obtain `PangoItems` by itemizing a piece of text with
// [func@itemize].
type Item struct {
	native C.PangoItem
}

// WrapItem wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapItem(ptr unsafe.Pointer) *Item {
	if ptr == nil {
		return nil
	}

	return (*Item)(ptr)
}

func marshalItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapItem(unsafe.Pointer(b)), nil
}

// NewItem constructs a struct Item.
func NewItem() *Item {
	var cret *C.PangoItem
	var goret1 *Item

	cret = C.pango_item_new()

	goret1 = WrapItem(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Item) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Native returns the underlying C source pointer.
func (i *Item) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// Offset gets the field inside the struct.
func (i *Item) Offset() int {
	v = C.gint(i.native.offset)
}

// Length gets the field inside the struct.
func (i *Item) Length() int {
	v = C.gint(i.native.length)
}

// NumChars gets the field inside the struct.
func (i *Item) NumChars() int {
	v = C.gint(i.native.num_chars)
}

// Analysis gets the field inside the struct.
func (i *Item) Analysis() Analysis {
	v = WrapAnalysis(unsafe.Pointer(i.native.analysis))
}

// ApplyAttrs: add attributes to a `PangoItem`.
//
// The idea is that you have attributes that don't affect itemization, such as
// font features, so you filter them out using [method@Pango.AttrList.filter],
// itemize your text, then reapply the attributes to the resulting items using
// this function.
//
// The @iter should be positioned before the range of the item, and will be
// advanced past it. This function is meant to be called in a loop over the
// items resulting from itemization, while passing the iter to each call.
func (i *Item) ApplyAttrs(iter *AttrIterator) {
	var arg0 *C.PangoItem
	var arg1 *C.PangoAttrIterator

	arg0 = (*C.PangoItem)(unsafe.Pointer(i.Native()))
	arg1 = (*C.PangoAttrIterator)(unsafe.Pointer(iter.Native()))

	C.pango_item_apply_attrs(arg0, iter)
}

// Copy: copy an existing `PangoItem` structure.
func (i *Item) Copy() *Item {
	var arg0 *C.PangoItem

	arg0 = (*C.PangoItem)(unsafe.Pointer(i.Native()))

	var cret *C.PangoItem
	var goret1 *Item

	cret = C.pango_item_copy(arg0)

	goret1 = WrapItem(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Item) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Free: free a `PangoItem` and all associated memory.
func (i *Item) Free() {
	var arg0 *C.PangoItem

	arg0 = (*C.PangoItem)(unsafe.Pointer(i.Native()))

	C.pango_item_free(arg0)
}

// Split modifies @orig to cover only the text after @split_index, and returns a
// new item that covers the text before @split_index that used to be in @orig.
//
// You can think of @split_index as the length of the returned item.
// @split_index may not be 0, and it may not be greater than or equal to the
// length of @orig (that is, there must be at least one byte assigned to each
// item, you can't create a zero-length item). @split_offset is the length of
// the first item in chars, and must be provided because the text used to
// generate the item isn't available, so `pango_item_split()` can't count the
// char length of the split items itself.
func (o *Item) Split(splitIndex int, splitOffset int) *Item {
	var arg0 *C.PangoItem
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.PangoItem)(unsafe.Pointer(o.Native()))
	arg1 = C.int(splitIndex)
	arg2 = C.int(splitOffset)

	var cret *C.PangoItem
	var goret1 *Item

	cret = C.pango_item_split(arg0, splitIndex, splitOffset)

	goret1 = WrapItem(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Item) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}
