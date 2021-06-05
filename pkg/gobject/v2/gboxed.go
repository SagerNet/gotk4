// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
import "C"

// BoxedCopy: provide a copy of a boxed structure @src_boxed which is of type
// @boxed_type.
func BoxedCopy(boxedType externglib.Type, srcBoxed interface{}) interface{} {
	var arg1 C.GType
	var arg2 C.gpointer

	arg1 := C.GType(boxedType)
	arg2 = C.gpointer(srcBoxed)

	var cret C.gpointer
	var goret1 interface{}

	cret = C.g_boxed_copy(boxedType, srcBoxed)

	goret1 = C.gpointer(cret)

	return goret1
}

// BoxedFree: free the boxed structure @boxed which is of type @boxed_type.
func BoxedFree(boxedType externglib.Type, boxed interface{}) {
	var arg1 C.GType
	var arg2 C.gpointer

	arg1 := C.GType(boxedType)
	arg2 = C.gpointer(boxed)

	C.g_boxed_free(boxedType, boxed)
}
