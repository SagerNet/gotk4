// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gtk/gtk.h>
import "C"

// TreeCellDataFunc: a function to set the properties of a cell instead of just
// using the straight mapping between the cell and the model. This is useful for
// customizing the cell renderer. For example, a function might get an integer
// from the @tree_model, and render it to the “text” attribute of “cell” by
// converting it to its written equivalent. This is set by calling
// gtk_tree_view_column_set_cell_data_func()
type TreeCellDataFunc func(treeColumn TreeViewColumn, cell CellRenderer, treeModel TreeModel, iter *TreeIter)

//export gotk4_TreeCellDataFunc
func gotk4_TreeCellDataFunc(arg0 *C.GtkTreeViewColumn, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}
	fn := v.(TreeCellDataFunc)

	fn(treeColumn, cell, treeModel, iter, data)
}
