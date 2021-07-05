// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_bookmark_list_get_type()), F: marshalBookmarkList},
	})
}

// BookmarkList: `GtkBookmarkList` is a list model that wraps `GBookmarkFile`.
//
// It presents a `GListModel` and fills it asynchronously with the `GFileInfo`s
// returned from that function.
//
// The `GFileInfo`s in the list have some attributes in the recent namespace
// added: `recent::private` (boolean) and `recent:applications` (stringv).
type BookmarkList interface {
	gextras.Objector

	// Attributes gets the attributes queried on the children.
	Attributes() string
	// Filename returns the filename of the bookmark file that this list is
	// loading.
	Filename() string
	// IOPriority gets the IO priority to use while loading file.
	IOPriority() int
	// IsLoading returns true if the files are currently being loaded.
	//
	// Files will be added to @self from time to time while loading is going on.
	// The order in which are added is undefined and may change in between runs.
	IsLoading() bool
	// SetAttributes sets the @attributes to be enumerated and starts the
	// enumeration.
	//
	// If @attributes is nil, no attributes will be queried, but a list of Infos
	// will still be created.
	SetAttributes(attributes string)
	// SetIOPriority sets the IO priority to use while loading files.
	//
	// The default IO priority is G_PRIORITY_DEFAULT.
	SetIOPriority(ioPriority int)
}

// bookmarkList implements the BookmarkList class.
type bookmarkList struct {
	gextras.Objector
}

// WrapBookmarkList wraps a GObject to the right type. It is
// primarily used internally.
func WrapBookmarkList(obj *externglib.Object) BookmarkList {
	return bookmarkList{
		Objector: obj,
	}
}

func marshalBookmarkList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBookmarkList(obj), nil
}

// NewBookmarkList creates a new `GtkBookmarkList` with the given @attributes.
func NewBookmarkList(filename string, attributes string) BookmarkList {
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out
	var _cret *C.GtkBookmarkList // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_bookmark_list_new(_arg1, _arg2)

	var _bookmarkList BookmarkList // out

	_bookmarkList = WrapBookmarkList(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _bookmarkList
}

func (s bookmarkList) Attributes() string {
	var _arg0 *C.GtkBookmarkList // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_bookmark_list_get_attributes(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s bookmarkList) Filename() string {
	var _arg0 *C.GtkBookmarkList // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_bookmark_list_get_filename(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s bookmarkList) IOPriority() int {
	var _arg0 *C.GtkBookmarkList // out
	var _cret C.int              // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_bookmark_list_get_io_priority(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s bookmarkList) IsLoading() bool {
	var _arg0 *C.GtkBookmarkList // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_bookmark_list_is_loading(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s bookmarkList) SetAttributes(attributes string) {
	var _arg0 *C.GtkBookmarkList // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_bookmark_list_set_attributes(_arg0, _arg1)
}

func (s bookmarkList) SetIOPriority(ioPriority int) {
	var _arg0 *C.GtkBookmarkList // out
	var _arg1 C.int              // out

	_arg0 = (*C.GtkBookmarkList)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(ioPriority)

	C.gtk_bookmark_list_set_io_priority(_arg0, _arg1)
}
