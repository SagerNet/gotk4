// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// together with the @user_data parameter passed to
// g_hash_table_foreach_remove(). It should return true if the key/value pair
// should be removed from the Table.
type HRFunc func(key interface{}, value interface{}, userData interface{}) (ok bool)

//export gotk4_HRFunc
func gotk4_HRFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key interface{}      // out
	var value interface{}    // out
	var userData interface{} // out

	key = box.Get(uintptr(arg0))
	value = box.Get(uintptr(arg1))
	userData = box.Get(uintptr(arg2))

	fn := v.(HRFunc)
	ok := fn(key, value, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// DirectEqual compares two #gpointer arguments and returns true if they are
// equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using opaque pointers compared by pointer value as keys in a
// Table.
//
// This equality function is also appropriate for keys that are integers stored
// in pointers, such as `GINT_TO_POINTER (n)`.
func DirectEqual(v1 interface{}, v2 interface{}) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(box.Assign(v1))
	_arg2 = (C.gconstpointer)(box.Assign(v2))

	_cret = C.g_direct_equal(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DirectHash converts a gpointer to a hash value. It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using opaque pointers
// compared by pointer value as keys in a Table.
//
// This hash function is also appropriate for keys that are integers stored in
// pointers, such as `GINT_TO_POINTER (n)`.
func DirectHash(v interface{}) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(box.Assign(v))

	_cret = C.g_direct_hash(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// DoubleEqual compares the two #gdouble values being pointed to and returns
// true if they are equal. It can be passed to g_hash_table_new() as the
// @key_equal_func parameter, when using non-nil pointers to doubles as keys in
// a Table.
func DoubleEqual(v1 interface{}, v2 interface{}) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(box.Assign(v1))
	_arg2 = (C.gconstpointer)(box.Assign(v2))

	_cret = C.g_double_equal(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DoubleHash converts a pointer to a #gdouble to a hash value. It can be passed
// to g_hash_table_new() as the @hash_func parameter, It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using non-nil pointers
// to doubles as keys in a Table.
func DoubleHash(v interface{}) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(box.Assign(v))

	_cret = C.g_double_hash(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Int64Equal compares the two #gint64 values being pointed to and returns true
// if they are equal. It can be passed to g_hash_table_new() as the
// @key_equal_func parameter, when using non-nil pointers to 64-bit integers as
// keys in a Table.
func Int64Equal(v1 interface{}, v2 interface{}) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(box.Assign(v1))
	_arg2 = (C.gconstpointer)(box.Assign(v2))

	_cret = C.g_int64_equal(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Int64Hash converts a pointer to a #gint64 to a hash value.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter, when
// using non-nil pointers to 64-bit integer values as keys in a Table.
func Int64Hash(v interface{}) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(box.Assign(v))

	_cret = C.g_int64_hash(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// IntEqual compares the two #gint values being pointed to and returns true if
// they are equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-nil pointers to integers as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form `GINT_TO_POINTER (n)`, use
// g_direct_equal() instead.
func IntEqual(v1 interface{}, v2 interface{}) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(box.Assign(v1))
	_arg2 = (C.gconstpointer)(box.Assign(v2))

	_cret = C.g_int_equal(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IntHash converts a pointer to a #gint to a hash value. It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using non-nil pointers
// to integer values as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form `GINT_TO_POINTER (n)`, use
// g_direct_hash() instead.
func IntHash(v interface{}) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(box.Assign(v))

	_cret = C.g_int_hash(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StrEqual compares two strings for byte-by-byte equality and returns true if
// they are equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-nil strings as keys in a Table.
//
// This function is typically used for hash table comparisons, but can be used
// for general purpose comparisons of non-nil strings. For a nil-safe string
// comparison function, see g_strcmp0().
func StrEqual(v1 interface{}, v2 interface{}) bool {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(box.Assign(v1))
	_arg2 = (C.gconstpointer)(box.Assign(v2))

	_cret = C.g_str_equal(_arg1, _arg2)

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
// starts at 5381 and for each byte 'c' in the string, is updated: `hash = hash
// * 33 + c`. This function uses the signed value of each byte.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter, when
// using non-nil strings as keys in a Table.
//
// Note that this function may not be a perfect fit for all use cases. For
// example, it produces some hash collisions with strings as short as 2.
func StrHash(v interface{}) uint {
	var _arg1 C.gconstpointer // out
	var _cret C.guint         // in

	_arg1 = (C.gconstpointer)(box.Assign(v))

	_cret = C.g_str_hash(_arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// HashTable: the Table struct is an opaque data structure to represent a [Hash
// Table][glib-Hash-Tables]. It should only be accessed via the following
// functions.
type HashTable struct {
	native C.GHashTable
}

func marshalHashTable(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*HashTable)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (h *HashTable) Native() unsafe.Pointer {
	return unsafe.Pointer(&h.native)
}

// HashTableIter structure represents an iterator that can be used to iterate
// over the elements of a Table. GHashTableIter structures are typically
// allocated on the stack and then initialized with g_hash_table_iter_init().
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
type HashTableIter struct {
	native C.GHashTableIter
}

// Native returns the underlying C source pointer.
func (h *HashTableIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&h.native)
}

// Init initializes a key/value pair iterator and associates it with
// @hash_table. Modifying the hash table after calling this function invalidates
// the returned iterator.
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
func (iter *HashTableIter) Init(hashTable *HashTable) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 *C.GHashTable     // out

	_arg0 = (*C.GHashTableIter)(unsafe.Pointer(iter))
	_arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable))

	C.g_hash_table_iter_init(_arg0, _arg1)
}

// Next advances @iter and retrieves the key and/or value that are now pointed
// to as a result of this advancement. If false is returned, @key and @value are
// not set, and the iterator becomes invalid.
func (iter *HashTableIter) Next() (key interface{}, value interface{}, ok bool) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 C.gpointer        // in
	var _arg2 C.gpointer        // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GHashTableIter)(unsafe.Pointer(iter))

	_cret = C.g_hash_table_iter_next(_arg0, &_arg1, &_arg2)

	var _key interface{}   // out
	var _value interface{} // out
	var _ok bool           // out

	_key = box.Get(uintptr(_arg1))
	_value = box.Get(uintptr(_arg2))
	if _cret != 0 {
		_ok = true
	}

	return _key, _value, _ok
}

// Remove removes the key/value pair currently pointed to by the iterator from
// its associated Table. Can only be called after g_hash_table_iter_next()
// returned true, and cannot be called more than once for the same key/value
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

	_arg0 = (*C.GHashTableIter)(unsafe.Pointer(iter))

	C.g_hash_table_iter_remove(_arg0)
}

// Replace replaces the value currently pointed to by the iterator from its
// associated Table. Can only be called after g_hash_table_iter_next() returned
// true.
//
// If you supplied a @value_destroy_func when creating the Table, the old value
// is freed using that function.
func (iter *HashTableIter) Replace(value interface{}) {
	var _arg0 *C.GHashTableIter // out
	var _arg1 C.gpointer        // out

	_arg0 = (*C.GHashTableIter)(unsafe.Pointer(iter))
	_arg1 = (C.gpointer)(box.Assign(value))

	C.g_hash_table_iter_replace(_arg0, _arg1)
}

// Steal removes the key/value pair currently pointed to by the iterator from
// its associated Table, without calling the key and value destroy functions.
// Can only be called after g_hash_table_iter_next() returned true, and cannot
// be called more than once for the same key/value pair.
func (iter *HashTableIter) Steal() {
	var _arg0 *C.GHashTableIter // out

	_arg0 = (*C.GHashTableIter)(unsafe.Pointer(iter))

	C.g_hash_table_iter_steal(_arg0)
}
