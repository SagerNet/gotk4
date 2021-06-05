// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// BitsetIterInitAt initializes @iter to point to @target. If @target is not
// found, finds the next value after it. If no value >= @target exists in @set,
// this function returns false.
func BitsetIterInitAt(set *Bitset, target uint) (iter BitsetIter, value uint, ok bool) {
	var arg2 *C.GtkBitset
	var arg3 C.guint

	arg2 = (*C.GtkBitset)(unsafe.Pointer(set.Native()))
	arg3 = C.guint(target)

	var arg1 C.GtkBitsetIter
	var ret1 *BitsetIter
	var arg4 C.guint
	var ret4 uint
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_bitset_iter_init_at(&arg1, set, target, &arg4)

	*ret1 = WrapBitsetIter(unsafe.Pointer(arg1))
	*ret4 = C.guint(arg4)
	ret3 = C.bool(cret) != C.false

	return ret1, ret4, ret3
}

// BitsetIterInitFirst initializes an iterator for @set and points it to the
// first value in @set. If @set is empty, false is returned and @value is set to
// G_MAXUINT.
func BitsetIterInitFirst(set *Bitset) (iter BitsetIter, value uint, ok bool) {
	var arg2 *C.GtkBitset

	arg2 = (*C.GtkBitset)(unsafe.Pointer(set.Native()))

	var arg1 C.GtkBitsetIter
	var ret1 *BitsetIter
	var arg3 C.guint
	var ret3 uint
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_bitset_iter_init_first(&arg1, set, &arg3)

	*ret1 = WrapBitsetIter(unsafe.Pointer(arg1))
	*ret3 = C.guint(arg3)
	ret3 = C.bool(cret) != C.false

	return ret1, ret3, ret3
}

// BitsetIterInitLast initializes an iterator for @set and points it to the last
// value in @set. If @set is empty, false is returned.
func BitsetIterInitLast(set *Bitset) (iter BitsetIter, value uint, ok bool) {
	var arg2 *C.GtkBitset

	arg2 = (*C.GtkBitset)(unsafe.Pointer(set.Native()))

	var arg1 C.GtkBitsetIter
	var ret1 *BitsetIter
	var arg3 C.guint
	var ret3 uint
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_bitset_iter_init_last(&arg1, set, &arg3)

	*ret1 = WrapBitsetIter(unsafe.Pointer(arg1))
	*ret3 = C.guint(arg3)
	ret3 = C.bool(cret) != C.false

	return ret1, ret3, ret3
}

// BitsetIter: an opaque, stack-allocated struct for iterating over the elements
// of a Bitset. Before a GtkBitsetIter can be used, it needs to be initialized
// with gtk_bitset_iter_init_first(), gtk_bitset_iter_init_last() or
// gtk_bitset_iter_init_at().
type BitsetIter struct {
	native C.GtkBitsetIter
}

// WrapBitsetIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBitsetIter(ptr unsafe.Pointer) *BitsetIter {
	if ptr == nil {
		return nil
	}

	return (*BitsetIter)(ptr)
}

func marshalBitsetIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBitsetIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BitsetIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Value gets the current value that @iter points to.
//
// If @iter is not valid and gtk_bitset_iter_is_valid() returns false, this
// function returns 0.
func (i *BitsetIter) Value() uint {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gtk_bitset_iter_get_value(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// IsValid checks if @iter points to a valid value.
func (i *BitsetIter) IsValid() bool {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_bitset_iter_is_valid(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Next moves @iter to the next value in the set. If it was already pointing to
// the last value in the set, false is returned and @iter is invalidated.
func (i *BitsetIter) Next() (value uint, ok bool) {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var arg1 C.guint
	var ret1 uint
	var cret C.gboolean
	var ret2 bool

	cret = C.gtk_bitset_iter_next(arg0, &arg1)

	*ret1 = C.guint(arg1)
	ret2 = C.bool(cret) != C.false

	return ret1, ret2
}

// Previous moves @iter to the previous value in the set. If it was already
// pointing to the first value in the set, false is returned and @iter is
// invalidated.
func (i *BitsetIter) Previous() (value uint, ok bool) {
	var arg0 *C.GtkBitsetIter

	arg0 = (*C.GtkBitsetIter)(unsafe.Pointer(i.Native()))

	var arg1 C.guint
	var ret1 uint
	var cret C.gboolean
	var ret2 bool

	cret = C.gtk_bitset_iter_previous(arg0, &arg1)

	*ret1 = C.guint(arg1)
	ret2 = C.bool(cret) != C.false

	return ret1, ret2
}