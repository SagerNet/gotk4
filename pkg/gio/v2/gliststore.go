// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_list_store_get_type()), F: marshalListStore},
	})
}

// ListStore is a simple implementation of Model that stores all items in
// memory.
//
// It provides insertions, deletions, and lookups in logarithmic time with a
// fast path for the common case of iterating the list linearly.
type ListStore interface {
	gextras.Objector
	ListModel

	// Append appends @item to @store. @item must be of type Store:item-type.
	//
	// This function takes a ref on @item.
	//
	// Use g_list_store_splice() to append multiple items at the same time
	// efficiently.
	Append(item gextras.Objector)
	// Find looks up the given @item in the list store by looping over the items
	// until the first occurrence of @item. If @item was not found, then
	// @position will not be set, and this method will return false.
	//
	// If you need to compare the two items with a custom comparison function,
	// use g_list_store_find_with_equal_func() with a custom Func instead.
	Find(item gextras.Objector) (position uint, ok bool)
	// Insert inserts @item into @store at @position. @item must be of type
	// Store:item-type or derived from it. @position must be smaller than the
	// length of the list, or equal to it to append.
	//
	// This function takes a ref on @item.
	//
	// Use g_list_store_splice() to insert multiple items at the same time
	// efficiently.
	Insert(position uint, item gextras.Objector)
	// InsertSorted inserts @item into @store at a position to be determined by
	// the @compare_func.
	//
	// The list must already be sorted before calling this function or the
	// result is undefined. Usually you would approach this by only ever
	// inserting items by way of this function.
	//
	// This function takes a ref on @item.
	InsertSorted(item gextras.Objector, compareFunc glib.CompareDataFunc) uint
	// Remove removes the item from @store that is at @position. @position must
	// be smaller than the current length of the list.
	//
	// Use g_list_store_splice() to remove multiple items at the same time
	// efficiently.
	Remove(position uint)
	// RemoveAll removes all items from @store.
	RemoveAll()
	// Sort: sort the items in @store according to @compare_func.
	Sort(compareFunc glib.CompareDataFunc)
	// Splice changes @store by removing @n_removals items and adding
	// @n_additions items to it. @additions must contain @n_additions items of
	// type Store:item-type. nil is not permitted.
	//
	// This function is more efficient than g_list_store_insert() and
	// g_list_store_remove(), because it only emits Model::items-changed once
	// for the change.
	//
	// This function takes a ref on each item in @additions.
	//
	// The parameters @position and @n_removals must be correct (ie: @position +
	// @n_removals must be less than or equal to the length of the list at the
	// time this function is called).
	Splice(position uint, nRemovals uint, additions []gextras.Objector)
}

// listStore implements the ListStore interface.
type listStore struct {
	gextras.Objector
	ListModel
}

var _ ListStore = (*listStore)(nil)

// WrapListStore wraps a GObject to the right type. It is
// primarily used internally.
func WrapListStore(obj *externglib.Object) ListStore {
	return ListStore{
		Objector:  obj,
		ListModel: WrapListModel(obj),
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListStore(obj), nil
}

// NewListStore constructs a class ListStore.
func NewListStore(itemType externglib.Type) ListStore {
	var arg1 C.GType

	arg1 := C.GType(itemType)

	var cret C.GListStore
	var ret1 ListStore

	cret = C.g_list_store_new(itemType)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ListStore)

	return ret1
}

// Append appends @item to @store. @item must be of type Store:item-type.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to append multiple items at the same time
// efficiently.
func (s listStore) Append(item gextras.Objector) {
	var arg0 *C.GListStore
	var arg1 C.gpointer

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(item.Native()))

	C.g_list_store_append(arg0, item)
}

// Find looks up the given @item in the list store by looping over the items
// until the first occurrence of @item. If @item was not found, then
// @position will not be set, and this method will return false.
//
// If you need to compare the two items with a custom comparison function,
// use g_list_store_find_with_equal_func() with a custom Func instead.
func (s listStore) Find(item gextras.Objector) (position uint, ok bool) {
	var arg0 *C.GListStore
	var arg1 C.gpointer

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(item.Native()))

	var arg2 C.guint
	var ret2 uint
	var cret C.gboolean
	var ret2 bool

	cret = C.g_list_store_find(arg0, item, &arg2)

	*ret2 = C.guint(arg2)
	ret2 = C.bool(cret) != C.false

	return ret2, ret2
}

// Insert inserts @item into @store at @position. @item must be of type
// Store:item-type or derived from it. @position must be smaller than the
// length of the list, or equal to it to append.
//
// This function takes a ref on @item.
//
// Use g_list_store_splice() to insert multiple items at the same time
// efficiently.
func (s listStore) Insert(position uint, item gextras.Objector) {
	var arg0 *C.GListStore
	var arg1 C.guint
	var arg2 C.gpointer

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)
	arg2 = (*C.GObject)(unsafe.Pointer(item.Native()))

	C.g_list_store_insert(arg0, position, item)
}

// InsertSorted inserts @item into @store at a position to be determined by
// the @compare_func.
//
// The list must already be sorted before calling this function or the
// result is undefined. Usually you would approach this by only ever
// inserting items by way of this function.
//
// This function takes a ref on @item.
func (s listStore) InsertSorted(item gextras.Objector, compareFunc glib.CompareDataFunc) uint {
	var arg0 *C.GListStore

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.g_list_store_insert_sorted(arg0, item, compareFunc, userData)

	ret1 = C.guint(cret)

	return ret1
}

// Remove removes the item from @store that is at @position. @position must
// be smaller than the current length of the list.
//
// Use g_list_store_splice() to remove multiple items at the same time
// efficiently.
func (s listStore) Remove(position uint) {
	var arg0 *C.GListStore
	var arg1 C.guint

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)

	C.g_list_store_remove(arg0, position)
}

// RemoveAll removes all items from @store.
func (s listStore) RemoveAll() {
	var arg0 *C.GListStore

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))

	C.g_list_store_remove_all(arg0)
}

// Sort: sort the items in @store according to @compare_func.
func (s listStore) Sort(compareFunc glib.CompareDataFunc) {
	var arg0 *C.GListStore

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))

	C.g_list_store_sort(arg0, compareFunc, userData)
}

// Splice changes @store by removing @n_removals items and adding
// @n_additions items to it. @additions must contain @n_additions items of
// type Store:item-type. nil is not permitted.
//
// This function is more efficient than g_list_store_insert() and
// g_list_store_remove(), because it only emits Model::items-changed once
// for the change.
//
// This function takes a ref on each item in @additions.
//
// The parameters @position and @n_removals must be correct (ie: @position +
// @n_removals must be less than or equal to the length of the list at the
// time this function is called).
func (s listStore) Splice(position uint, nRemovals uint, additions []gextras.Objector) {
	var arg0 *C.GListStore

	arg0 = (*C.GListStore)(unsafe.Pointer(s.Native()))

	C.g_list_store_splice(arg0, position, nRemovals, additions, nAdditions)
}