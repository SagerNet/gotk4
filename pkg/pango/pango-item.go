// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_item_get_type()), F: marshalItem},
	})
}

// ANALYSIS_FLAG_CENTERED_BASELINE: whether the segment should be shifted to
// center around the baseline.
//
// This is mainly used in vertical writing directions.
const ANALYSIS_FLAG_CENTERED_BASELINE = 1

// ANALYSIS_FLAG_IS_ELLIPSIS: whether this run holds ellipsized text.
const ANALYSIS_FLAG_IS_ELLIPSIS = 2

// ANALYSIS_FLAG_NEED_HYPHEN: whether to add a hyphen at the end of the run
// during shaping.
const ANALYSIS_FLAG_NEED_HYPHEN = 4

// Analysis: PangoAnalysis structure stores information about the properties of
// a segment of text.
//
// An instance of this type is always passed by reference.
type Analysis struct {
	*analysis
}

// analysis is the struct that's finalized.
type analysis struct {
	native *C.PangoAnalysis
}

// ShapeEngine: unused
func (a *Analysis) ShapeEngine() cgo.Handle {
	var v cgo.Handle // out
	v = (cgo.Handle)(unsafe.Pointer(a.native.shape_engine))
	return v
}

// LangEngine: unused
func (a *Analysis) LangEngine() cgo.Handle {
	var v cgo.Handle // out
	v = (cgo.Handle)(unsafe.Pointer(a.native.lang_engine))
	return v
}

// Font: font for this segment.
func (a *Analysis) Font() Fonter {
	var v Fonter // out
	v = (externglib.CastObject(externglib.Take(unsafe.Pointer(a.native.font)))).(Fonter)
	return v
}

// Level: bidirectional level for this segment.
func (a *Analysis) Level() byte {
	var v byte // out
	v = byte(a.native.level)
	return v
}

// Gravity: glyph orientation for this segment (A PangoGravity).
func (a *Analysis) Gravity() byte {
	var v byte // out
	v = byte(a.native.gravity)
	return v
}

// Flags: boolean flags for this segment (Since: 1.16).
func (a *Analysis) Flags() byte {
	var v byte // out
	v = byte(a.native.flags)
	return v
}

// Script: detected script for this segment (A PangoScript) (Since: 1.18).
func (a *Analysis) Script() byte {
	var v byte // out
	v = byte(a.native.script)
	return v
}

// Language: detected language for this segment.
func (a *Analysis) Language() *Language {
	var v *Language // out
	v = (*Language)(gextras.NewStructNative(unsafe.Pointer(a.native.language)))
	return v
}

// Item: PangoItem structure stores information about a segment of text.
//
// You typically obtain PangoItems by itemizing a piece of text with itemize.
//
// An instance of this type is always passed by reference.
type Item struct {
	*item
}

// item is the struct that's finalized.
type item struct {
	native *C.PangoItem
}

func marshalItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Item{&item{(*C.PangoItem)(unsafe.Pointer(b))}}, nil
}

// NewItem constructs a struct Item.
func NewItem() *Item {
	var _cret *C.PangoItem // in

	_cret = C.pango_item_new()

	var _item *Item // out

	_item = (*Item)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_item)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_item_free((*C.PangoItem)(intern.C))
		},
	)

	return _item
}

// Offset: byte offset of the start of this item in text.
func (i *Item) Offset() int {
	var v int // out
	v = int(i.native.offset)
	return v
}

// Length: length of this item in bytes.
func (i *Item) Length() int {
	var v int // out
	v = int(i.native.length)
	return v
}

// NumChars: number of Unicode characters in the item.
func (i *Item) NumChars() int {
	var v int // out
	v = int(i.native.num_chars)
	return v
}

// Analysis analysis results for the item.
func (i *Item) Analysis() Analysis {
	var v Analysis // out
	v = *(*Analysis)(gextras.NewStructNative(unsafe.Pointer((&i.native.analysis))))
	return v
}

// ApplyAttrs: add attributes to a PangoItem.
//
// The idea is that you have attributes that don't affect itemization, such as
// font features, so you filter them out using pango.AttrList.Filter(), itemize
// your text, then reapply the attributes to the resulting items using this
// function.
//
// The iter should be positioned before the range of the item, and will be
// advanced past it. This function is meant to be called in a loop over the
// items resulting from itemization, while passing the iter to each call.
func (item *Item) ApplyAttrs(iter *AttrIterator) {
	var _arg0 *C.PangoItem         // out
	var _arg1 *C.PangoAttrIterator // out

	_arg0 = (*C.PangoItem)(gextras.StructNative(unsafe.Pointer(item)))
	_arg1 = (*C.PangoAttrIterator)(gextras.StructNative(unsafe.Pointer(iter)))

	C.pango_item_apply_attrs(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(iter)
}

// Copy an existing PangoItem structure.
func (item *Item) Copy() *Item {
	var _arg0 *C.PangoItem // out
	var _cret *C.PangoItem // in

	if item != nil {
		_arg0 = (*C.PangoItem)(gextras.StructNative(unsafe.Pointer(item)))
	}

	_cret = C.pango_item_copy(_arg0)
	runtime.KeepAlive(item)

	var _ret *Item // out

	if _cret != nil {
		_ret = (*Item)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_ret)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_item_free((*C.PangoItem)(intern.C))
			},
		)
	}

	return _ret
}

// Split modifies orig to cover only the text after split_index, and returns a
// new item that covers the text before split_index that used to be in orig.
//
// You can think of split_index as the length of the returned item. split_index
// may not be 0, and it may not be greater than or equal to the length of orig
// (that is, there must be at least one byte assigned to each item, you can't
// create a zero-length item). split_offset is the length of the first item in
// chars, and must be provided because the text used to generate the item isn't
// available, so pango_item_split() can't count the char length of the split
// items itself.
func (orig *Item) Split(splitIndex int, splitOffset int) *Item {
	var _arg0 *C.PangoItem // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _cret *C.PangoItem // in

	_arg0 = (*C.PangoItem)(gextras.StructNative(unsafe.Pointer(orig)))
	_arg1 = C.int(splitIndex)
	_arg2 = C.int(splitOffset)

	_cret = C.pango_item_split(_arg0, _arg1, _arg2)
	runtime.KeepAlive(orig)
	runtime.KeepAlive(splitIndex)
	runtime.KeepAlive(splitOffset)

	var _item *Item // out

	_item = (*Item)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_item)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_item_free((*C.PangoItem)(intern.C))
		},
	)

	return _item
}
