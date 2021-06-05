// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_view_accessible_get_type()), F: marshalTreeViewAccessible},
	})
}

type TreeViewAccessible interface {
	ContainerAccessible
	CellAccessibleParent
}

// treeViewAccessible implements the TreeViewAccessible interface.
type treeViewAccessible struct {
	ContainerAccessible
	CellAccessibleParent
}

var _ TreeViewAccessible = (*treeViewAccessible)(nil)

// WrapTreeViewAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewAccessible(obj *externglib.Object) TreeViewAccessible {
	return TreeViewAccessible{
		ContainerAccessible:  WrapContainerAccessible(obj),
		CellAccessibleParent: WrapCellAccessibleParent(obj),
	}
}

func marshalTreeViewAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewAccessible(obj), nil
}

type TreeViewAccessiblePrivate struct {
	native C.GtkTreeViewAccessiblePrivate
}

// WrapTreeViewAccessiblePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTreeViewAccessiblePrivate(ptr unsafe.Pointer) *TreeViewAccessiblePrivate {
	if ptr == nil {
		return nil
	}

	return (*TreeViewAccessiblePrivate)(ptr)
}

func marshalTreeViewAccessiblePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTreeViewAccessiblePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TreeViewAccessiblePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
