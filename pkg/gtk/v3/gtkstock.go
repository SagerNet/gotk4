// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

type Stock = string

// TranslateFunc: the function used to translate messages in e.g. IconFactory
// and ActionGroup.
//
// Deprecated: since version 3.10.
type TranslateFunc func(path string, utf8 string)

//export gotk4_TranslateFunc
func gotk4_TranslateFunc(arg0 *C.gchar, arg1 C.gpointer) *C.gchar {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var path string // out

	path = C.GoString(arg0)

	fn := v.(TranslateFunc)
	utf8 := fn(path)

	var cret *C.gchar // out

	cret = (*C.gchar)(C.CString(utf8))

	return cret
}

// StockAdd registers each of the stock items in @items. If an item already
// exists with the same stock ID as one of the @items, the old item gets
// replaced. The stock items are copied, so GTK+ does not hold any pointer into
// @items and @items can be freed. Use gtk_stock_add_static() if @items is
// persistent and GTK+ need not copy the array.
//
// Deprecated: since version 3.10.
func StockAdd(items []StockItem) {
	var _arg1 *C.GtkStockItem
	var _arg2 C.guint

	_arg2 = C.guint(len(items))
	_arg1 = (*C.GtkStockItem)(unsafe.Pointer(&items[0]))

	C.gtk_stock_add(_arg1, _arg2)
}

// StockAddStatic: same as gtk_stock_add(), but doesn’t copy @items, so @items
// must persist until application exit.
//
// Deprecated: since version 3.10.
func StockAddStatic(items []StockItem) {
	var _arg1 *C.GtkStockItem
	var _arg2 C.guint

	_arg2 = C.guint(len(items))
	_arg1 = (*C.GtkStockItem)(unsafe.Pointer(&items[0]))

	C.gtk_stock_add_static(_arg1, _arg2)
}

// StockLookup fills @item with the registered values for @stock_id, returning
// true if @stock_id was known.
//
// Deprecated: since version 3.10.
func StockLookup(stockId string) (StockItem, bool) {
	var _arg1 *C.gchar       // out
	var _arg2 C.GtkStockItem // in
	var _cret C.gboolean     // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_stock_lookup(_arg1, &_arg2)

	var _item StockItem // out
	var _ok bool        // out

	{
		var refTmpIn *C.GtkStockItem
		var refTmpOut *StockItem

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*StockItem)(unsafe.Pointer(refTmpIn))

		_item = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _item, _ok
}

// StockItem: deprecated: since version 3.10.
type StockItem C.GtkStockItem

// WrapStockItem wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStockItem(ptr unsafe.Pointer) *StockItem {
	return (*StockItem)(ptr)
}

// Native returns the underlying C source pointer.
func (s *StockItem) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}

// Free frees a stock item allocated on the heap, such as one returned by
// gtk_stock_item_copy(). Also frees the fields inside the stock item, if they
// are not nil.
//
// Deprecated: since version 3.10.
func (i *StockItem) Free() {
	var _arg0 *C.GtkStockItem // out

	_arg0 = (*C.GtkStockItem)(unsafe.Pointer(i.Native()))

	C.gtk_stock_item_free(_arg0)
}