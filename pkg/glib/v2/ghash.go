// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_hash_table_get_type()), F: marshalHashTable},
	})
}

// HRFunc specifies the type of the function passed to
// g_hash_table_foreach_remove(). It is called with each key/value pair,
// together with the user_data parameter passed to
// g_hash_table_foreach_remove(). It should return TRUE if the key/value pair
// should be removed from the Table.
type HRFunc func(key cgo.Handle, value cgo.Handle) (ok bool)

//export _gotk4_glib2_HRFunc
func _gotk4_glib2_HRFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key cgo.Handle   // out
	var value cgo.Handle // out

	key = (cgo.Handle)(unsafe.Pointer(arg0))
	value = (cgo.Handle)(unsafe.Pointer(arg1))

	fn := v.(HRFunc)
	ok := fn(key, value)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// DirectEqual compares two #gpointer arguments and returns TRUE if they are
// equal. It can be passed to g_hash_table_new() as the key_equal_func
// parameter, when using opaque pointers compared by pointer value as keys in a
// Table.
//
// This equality function is also appropriate for keys that are integers stored
// in pointers, such as GINT_TO_POINTER (n).
func DirectEqual(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.g_direct_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DirectHash converts a gpointer to a hash value. It can be passed to
// g_hash_table_new() as the hash_func parameter, when using opaque pointers
// compared by pointer value as keys in a Table.
//
// This hash function is also appropriate for keys that are integers stored in
// pointers, such as GINT_TO_POINTER (n).
func DirectHash(v cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v))

	_cret = C.g_direct_hash(_arg1)
	runtime.KeepAlive(v)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DoubleEqual compares the two #gdouble values being pointed to and returns
// TRUE if they are equal. It can be passed to g_hash_table_new() as the
// key_equal_func parameter, when using non-NULL pointers to doubles as keys in
// a Table.
func DoubleEqual(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.g_double_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DoubleHash converts a pointer to a #gdouble to a hash value. It can be passed
// to g_hash_table_new() as the hash_func parameter, It can be passed to
// g_hash_table_new() as the hash_func parameter, when using non-NULL pointers
// to doubles as keys in a Table.
func DoubleHash(v cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v))

	_cret = C.g_double_hash(_arg1)
	runtime.KeepAlive(v)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Int64Equal compares the two #gint64 values being pointed to and returns TRUE
// if they are equal. It can be passed to g_hash_table_new() as the
// key_equal_func parameter, when using non-NULL pointers to 64-bit integers as
// keys in a Table.
func Int64Equal(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.g_int64_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Int64Hash converts a pointer to a #gint64 to a hash value.
//
// It can be passed to g_hash_table_new() as the hash_func parameter, when using
// non-NULL pointers to 64-bit integer values as keys in a Table.
func Int64Hash(v cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v))

	_cret = C.g_int64_hash(_arg1)
	runtime.KeepAlive(v)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IntEqual compares the two #gint values being pointed to and returns TRUE if
// they are equal. It can be passed to g_hash_table_new() as the key_equal_func
// parameter, when using non-NULL pointers to integers as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form GINT_TO_POINTER (n), use
// g_direct_equal() instead.
func IntEqual(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.g_int_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IntHash converts a pointer to a #gint to a hash value. It can be passed to
// g_hash_table_new() as the hash_func parameter, when using non-NULL pointers
// to integer values as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form GINT_TO_POINTER (n), use
// g_direct_hash() instead.
func IntHash(v cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v))

	_cret = C.g_int_hash(_arg1)
	runtime.KeepAlive(v)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StrEqual compares two strings for byte-by-byte equality and returns TRUE if
// they are equal. It can be passed to g_hash_table_new() as the key_equal_func
// parameter, when using non-NULL strings as keys in a Table.
//
// This function is typically used for hash table comparisons, but can be used
// for general purpose comparisons of non-NULL strings. For a NULL-safe string
// comparison function, see g_strcmp0().
func StrEqual(v1 cgo.Handle, v2 cgo.Handle) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v1))
	_arg2 = (C.gconstpointer)(unsafe.Pointer(v2))

	_cret = C.g_str_equal(_arg1, _arg2)
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StrHash converts a string to a hash value.
//
// This function implements the widely used "djb" hash apparently posted by
// Daniel Bernstein to comp.lang.c some time ago. The 32 bit unsigned hash value
// starts at 5381 and for each byte 'c' in the string, is updated: hash = hash *
// 33 + c. This function uses the signed value of each byte.
//
// It can be passed to g_hash_table_new() as the hash_func parameter, when using
// non-NULL strings as keys in a Table.
//
// Note that this function may not be a perfect fit for all use cases. For
// example, it produces some hash collisions with strings as short as 2.
func StrHash(v cgo.Handle) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(v))

	_cret = C.g_str_hash(_arg1)
	runtime.KeepAlive(v)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// HashTable struct is an opaque data structure to represent a [Hash
// Table][glib-Hash-Tables]. It should only be accessed via the following
// functions.
//
// An instance of this type is always passed by reference.
type HashTable struct {
	*hashTable
}

// hashTable is the struct that's finalized.
type hashTable struct {
	native *C.GHashTable
}

func marshalHashTable(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &HashTable{&hashTable{(*C.GHashTable)(unsafe.Pointer(b))}}, nil
}

// HashTableAdd: this is a convenience function for using a Table as a set. It
// is equivalent to calling g_hash_table_replace() with key as both the key and
// the value.
//
// In particular, this means that if key already exists in the hash table, then
// the old copy of key in the hash table is freed and key replaces it in the
// table.
//
// When a hash table only ever contains keys that have themselves as the
// corresponding value it is able to be stored more efficiently. See the
// discussion in the section description.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableAdd(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle) bool {
	var _arg1 *C.GHashTable // out
	var _arg2 C.gpointer    // out
	var _cret C.gboolean    // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gpointer)(unsafe.Pointer(key))

	_cret = C.g_hash_table_add(_arg1, _arg2)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableContains checks if key is in hash_table.
func HashTableContains(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle) bool {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_hash_table_contains(_arg1, _arg2)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableDestroy destroys all keys and values in the Table and decrements its
// reference count by 1. If keys and/or values are dynamically allocated, you
// should either free them first or create the Table with destroy notifiers
// using g_hash_table_new_full(). In the latter case the destroy functions you
// supplied will be called on all keys and values during the destruction phase.
func HashTableDestroy(hashTable map[cgo.Handle]cgo.Handle) {
	var _arg1 *C.GHashTable // out

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	C.g_hash_table_destroy(_arg1)
	runtime.KeepAlive(hashTable)
}

// HashTableInsert inserts a new key and value into a Table.
//
// If the key already exists in the Table its current value is replaced with the
// new value. If you supplied a value_destroy_func when creating the Table, the
// old value is freed using that function. If you supplied a key_destroy_func
// when creating the Table, the passed key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableInsert(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle, value cgo.Handle) bool {
	var _arg1 *C.GHashTable // out
	var _arg2 C.gpointer    // out
	var _arg3 C.gpointer    // out
	var _cret C.gboolean    // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gpointer)(unsafe.Pointer(key))
	_arg3 = (C.gpointer)(unsafe.Pointer(value))

	_cret = C.g_hash_table_insert(_arg1, _arg2, _arg3)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableLookup looks up a key in a Table. Note that this function cannot
// distinguish between a key that is not present and one which is present and
// has the value NULL. If you need this distinction, use
// g_hash_table_lookup_extended().
func HashTableLookup(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle) cgo.Handle {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _cret C.gpointer      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_hash_table_lookup(_arg1, _arg2)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// HashTableLookupExtended looks up a key in the Table, returning the original
// key and the associated value and a #gboolean which is TRUE if the key was
// found. This is useful if you need to free the memory allocated for the
// original key, for example before calling g_hash_table_remove().
//
// You can actually pass NULL for lookup_key to test whether the NULL key
// exists, provided the hash and equal functions of hash_table are NULL-safe.
func HashTableLookupExtended(hashTable map[cgo.Handle]cgo.Handle, lookupKey cgo.Handle) (origKey cgo.Handle, value cgo.Handle, ok bool) {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _arg3 C.gpointer      // in
	var _arg4 C.gpointer      // in
	var _cret C.gboolean      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(lookupKey))

	_cret = C.g_hash_table_lookup_extended(_arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(lookupKey)

	var _origKey cgo.Handle // out
	var _value cgo.Handle   // out
	var _ok bool            // out

	_origKey = (cgo.Handle)(unsafe.Pointer(_arg3))
	_value = (cgo.Handle)(unsafe.Pointer(_arg4))
	if _cret != 0 {
		_ok = true
	}

	return _origKey, _value, _ok
}

// HashTableRemove removes a key and its associated value from a Table.
//
// If the Table was created using g_hash_table_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself.
func HashTableRemove(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle) bool {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_hash_table_remove(_arg1, _arg2)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableRemoveAll removes all keys and their associated values from a Table.
//
// If the Table was created using g_hash_table_new_full(), the keys and values
// are freed using the supplied destroy functions, otherwise you have to make
// sure that any dynamically allocated values are freed yourself.
func HashTableRemoveAll(hashTable map[cgo.Handle]cgo.Handle) {
	var _arg1 *C.GHashTable // out

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	C.g_hash_table_remove_all(_arg1)
	runtime.KeepAlive(hashTable)
}

// HashTableReplace inserts a new key and value into a Table similar to
// g_hash_table_insert(). The difference is that if the key already exists in
// the Table, it gets replaced by the new key. If you supplied a
// value_destroy_func when creating the Table, the old value is freed using that
// function. If you supplied a key_destroy_func when creating the Table, the old
// key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableReplace(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle, value cgo.Handle) bool {
	var _arg1 *C.GHashTable // out
	var _arg2 C.gpointer    // out
	var _arg3 C.gpointer    // out
	var _cret C.gboolean    // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gpointer)(unsafe.Pointer(key))
	_arg3 = (C.gpointer)(unsafe.Pointer(value))

	_cret = C.g_hash_table_replace(_arg1, _arg2, _arg3)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableSize returns the number of elements contained in the Table.
func HashTableSize(hashTable map[cgo.Handle]cgo.Handle) uint {
	var _arg1 *C.GHashTable // out
	var _cret C.guint       // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	_cret = C.g_hash_table_size(_arg1)
	runtime.KeepAlive(hashTable)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// HashTableSteal removes a key and its associated value from a Table without
// calling the key and value destroy functions.
func HashTableSteal(hashTable map[cgo.Handle]cgo.Handle, key cgo.Handle) bool {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(key))

	_cret = C.g_hash_table_steal(_arg1, _arg2)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(key)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HashTableStealAll removes all keys and their associated values from a Table
// without calling the key and value destroy functions.
func HashTableStealAll(hashTable map[cgo.Handle]cgo.Handle) {
	var _arg1 *C.GHashTable // out

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	C.g_hash_table_steal_all(_arg1)
	runtime.KeepAlive(hashTable)
}

// HashTableStealExtended looks up a key in the Table, stealing the original key
// and the associated value and returning TRUE if the key was found. If the key
// was not found, FALSE is returned.
//
// If found, the stolen key and value are removed from the hash table without
// calling the key and value destroy functions, and ownership is transferred to
// the caller of this method; as with g_hash_table_steal().
//
// You can pass NULL for lookup_key, provided the hash and equal functions of
// hash_table are NULL-safe.
func HashTableStealExtended(hashTable map[cgo.Handle]cgo.Handle, lookupKey cgo.Handle) (stolenKey cgo.Handle, stolenValue cgo.Handle, ok bool) {
	var _arg1 *C.GHashTable   // out
	var _arg2 C.gconstpointer // out
	var _arg3 C.gpointer      // in
	var _arg4 C.gpointer      // in
	var _cret C.gboolean      // in

	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)
	_arg2 = (C.gconstpointer)(unsafe.Pointer(lookupKey))

	_cret = C.g_hash_table_steal_extended(_arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(hashTable)
	runtime.KeepAlive(lookupKey)

	var _stolenKey cgo.Handle   // out
	var _stolenValue cgo.Handle // out
	var _ok bool                // out

	_stolenKey = (cgo.Handle)(unsafe.Pointer(_arg3))
	_stolenValue = (cgo.Handle)(unsafe.Pointer(_arg4))
	if _cret != 0 {
		_ok = true
	}

	return _stolenKey, _stolenValue, _ok
}

// HashTableIter structure represents an iterator that can be used to iterate
// over the elements of a Table. GHashTableIter structures are typically
// allocated on the stack and then initialized with g_hash_table_iter_init().
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
//
// An instance of this type is always passed by reference.
type HashTableIter struct {
	*hashTableIter
}

// hashTableIter is the struct that's finalized.
type hashTableIter struct {
	native *C.GHashTableIter
}

// Init initializes a key/value pair iterator and associates it with hash_table.
// Modifying the hash table after calling this function invalidates the returned
// iterator.
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
//
//    GHashTableIter iter;
//    gpointer key, value;
//
//    g_hash_table_iter_init (&iter, hash_table);
//    while (g_hash_table_iter_next (&iter, &key, &value))
//      {
//        // do something with key and value
//      }
func (iter *HashTableIter) Init(hashTable map[cgo.Handle]cgo.Handle) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 *C.GHashTable     // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg1 = C.g_hash_table_new_full(nil, nil, (*[0]byte)(C.free), (*[0]byte)(C.free))
	for ksrc, vsrc := range hashTable {
		var kdst *C.gpointer // out
		var vdst *C.gpointer // out
		kdst = (*C.gpointer)(unsafe.Pointer(ksrc))
		vdst = (*C.gpointer)(unsafe.Pointer(vsrc))
		C.g_hash_table_insert(_arg1, C.gpointer(unsafe.Pointer(kdst)), C.gpointer(unsafe.Pointer(vdst)))
	}
	defer C.g_hash_table_unref(_arg1)

	C.g_hash_table_iter_init(_arg0, _arg1)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(hashTable)
}

// Next advances iter and retrieves the key and/or value that are now pointed to
// as a result of this advancement. If FALSE is returned, key and value are not
// set, and the iterator becomes invalid.
func (iter *HashTableIter) Next() (key cgo.Handle, value cgo.Handle, ok bool) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 C.gpointer        // in
	var _arg2 C.gpointer        // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.g_hash_table_iter_next(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(iter)

	var _key cgo.Handle   // out
	var _value cgo.Handle // out
	var _ok bool          // out

	_key = (cgo.Handle)(unsafe.Pointer(_arg1))
	_value = (cgo.Handle)(unsafe.Pointer(_arg2))
	if _cret != 0 {
		_ok = true
	}

	return _key, _value, _ok
}

// Remove removes the key/value pair currently pointed to by the iterator from
// its associated Table. Can only be called after g_hash_table_iter_next()
// returned TRUE, and cannot be called more than once for the same key/value
// pair.
//
// If the Table was created using g_hash_table_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself.
//
// It is safe to continue iterating the Table afterward:
//
//    while (g_hash_table_iter_next (&iter, &key, &value))
//      {
//        if (condition)
//          g_hash_table_iter_remove (&iter);
//      }
func (iter *HashTableIter) Remove() {
	var _arg0 *C.GHashTableIter // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.g_hash_table_iter_remove(_arg0)
	runtime.KeepAlive(iter)
}

// Replace replaces the value currently pointed to by the iterator from its
// associated Table. Can only be called after g_hash_table_iter_next() returned
// TRUE.
//
// If you supplied a value_destroy_func when creating the Table, the old value
// is freed using that function.
func (iter *HashTableIter) Replace(value cgo.Handle) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 C.gpointer        // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg1 = (C.gpointer)(unsafe.Pointer(value))

	C.g_hash_table_iter_replace(_arg0, _arg1)
	runtime.KeepAlive(iter)
	runtime.KeepAlive(value)
}

// Steal removes the key/value pair currently pointed to by the iterator from
// its associated Table, without calling the key and value destroy functions.
// Can only be called after g_hash_table_iter_next() returned TRUE, and cannot
// be called more than once for the same key/value pair.
func (iter *HashTableIter) Steal() {
	var _arg0 *C.GHashTableIter // out

	_arg0 = (*C.GHashTableIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.g_hash_table_iter_steal(_arg0)
	runtime.KeepAlive(iter)
}
