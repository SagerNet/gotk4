// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

type SliceConfig int

const (
	SliceConfigAlwaysMalloc SliceConfig = 1

	SliceConfigBypassMagazines SliceConfig = 2

	SliceConfigWorkingSetMsecs SliceConfig = 3

	SliceConfigColorIncrement SliceConfig = 4

	SliceConfigChunkSizes SliceConfig = 5

	SliceConfigContentionCounter SliceConfig = 6
)

// SliceAlloc allocates a block of memory from the slice allocator. The block
// address handed out can be expected to be aligned to at least 1 * sizeof
// (void*), though in general slices are 2 * sizeof (void*) bytes aligned, if a
// malloc() fallback implementation is used instead, the alignment may be
// reduced in a libc dependent fashion. Note that the underlying slice
// allocation mechanism can be changed with the
// [`G_SLICE=always-malloc`][G_SLICE] environment variable.
func SliceAlloc(blockSize uint) interface{} {
	var arg1 C.gsize

	arg1 = C.gsize(blockSize)

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_slice_alloc(blockSize)

	ret1 = C.gpointer(cret)

	return ret1
}

// SliceAlloc0 allocates a block of memory via g_slice_alloc() and initializes
// the returned memory to 0. Note that the underlying slice allocation mechanism
// can be changed with the [`G_SLICE=always-malloc`][G_SLICE] environment
// variable.
func SliceAlloc0(blockSize uint) interface{} {
	var arg1 C.gsize

	arg1 = C.gsize(blockSize)

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_slice_alloc0(blockSize)

	ret1 = C.gpointer(cret)

	return ret1
}

// SliceCopy allocates a block of memory from the slice allocator and copies
// @block_size bytes into it from @mem_block.
//
// @mem_block must be non-nil if @block_size is non-zero.
func SliceCopy(blockSize uint, memBlock interface{}) interface{} {
	var arg1 C.gsize
	var arg2 C.gpointer

	arg1 = C.gsize(blockSize)
	arg2 = C.gpointer(memBlock)

	var cret C.gpointer
	var ret1 interface{}

	cret = C.g_slice_copy(blockSize, memBlock)

	ret1 = C.gpointer(cret)

	return ret1
}

// SliceFree1 frees a block of memory.
//
// The memory must have been allocated via g_slice_alloc() or g_slice_alloc0()
// and the @block_size has to match the size specified upon allocation. Note
// that the exact release behaviour can be changed with the
// [`G_DEBUG=gc-friendly`][G_DEBUG] environment variable, also see
// [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_block is nil, this function does nothing.
func SliceFree1(blockSize uint, memBlock interface{}) {
	var arg1 C.gsize
	var arg2 C.gpointer

	arg1 = C.gsize(blockSize)
	arg2 = C.gpointer(memBlock)

	C.g_slice_free1(blockSize, memBlock)
}

// SliceFreeChainWithOffset frees a linked list of memory blocks of structure
// type @type.
//
// The memory blocks must be equal-sized, allocated via g_slice_alloc() or
// g_slice_alloc0() and linked together by a @next pointer (similar to List).
// The offset of the @next field in each block is passed as third argument. Note
// that the exact release behaviour can be changed with the
// [`G_DEBUG=gc-friendly`][G_DEBUG] environment variable, also see
// [`G_SLICE`][G_SLICE] for related debugging options.
//
// If @mem_chain is nil, this function does nothing.
func SliceFreeChainWithOffset(blockSize uint, memChain interface{}, nextOffset uint) {
	var arg1 C.gsize
	var arg2 C.gpointer
	var arg3 C.gsize

	arg1 = C.gsize(blockSize)
	arg2 = C.gpointer(memChain)
	arg3 = C.gsize(nextOffset)

	C.g_slice_free_chain_with_offset(blockSize, memChain, nextOffset)
}

func SliceGetConfig(ckey SliceConfig) int64 {
	var arg1 C.GSliceConfig

	arg1 = (C.GSliceConfig)(ckey)

	var cret C.gint64
	var ret1 int64

	cret = C.g_slice_get_config(ckey)

	ret1 = C.gint64(cret)

	return ret1
}

func SliceGetConfigState(ckey SliceConfig, address int64, nValues uint) int64 {
	var arg1 C.GSliceConfig
	var arg2 C.gint64
	var arg3 *C.guint

	arg1 = (C.GSliceConfig)(ckey)
	arg2 = C.gint64(address)
	arg3 = *C.guint(nValues)

	var cret *C.gint64
	var ret1 int64

	cret = C.g_slice_get_config_state(ckey, address, nValues)

	ret1 = *C.gint64(cret)

	return ret1
}

func SliceSetConfig(ckey SliceConfig, value int64) {
	var arg1 C.GSliceConfig
	var arg2 C.gint64

	arg1 = (C.GSliceConfig)(ckey)
	arg2 = C.gint64(value)

	C.g_slice_set_config(ckey, value)
}