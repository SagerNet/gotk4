// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_file_filter_get_type()), F: marshalFileFilter},
	})
}

// FileFilter: `GtkFileFilter` filters files by name or mime type.
//
// `GtkFileFilter` can be used to restrict the files being shown in a
// `GtkFileChooser`. Files can be filtered based on their name (with
// [method@Gtk.FileFilter.add_pattern]) or on their mime type (with
// [method@Gtk.FileFilter.add_mime_type]).
//
// Filtering by mime types handles aliasing and subclassing of mime types; e.g.
// a filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that `GtkFileFilter`
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, file filters are used by adding them to a `GtkFileChooser` (see
// [method@Gtk.FileChooser.add_filter]), but it is also possible to manually use
// a file filter on any [class@Gtk.FilterListModel] containing `GFileInfo`
// objects.
//
//
// GtkFileFilter as GtkBuildable
//
// The `GtkFileFilter` implementation of the `GtkBuildable` interface supports
// adding rules using the <mime-types> and <patterns> elements and listing the
// rules within. Specifying a <mime-type> or <pattern> has the same effect as as
// calling [method@Gtk.FileFilter.add_mime_type] or
// [method@Gtk.FileFilter.add_pattern].
//
// An example of a UI definition fragment specifying `GtkFileFilter` rules:
// “`xml <object class="GtkFileFilter"> <property name="name"
// translatable="yes">Text and Images</property> <mime-types>
// <mime-type>text/plain</mime-type> <mime-type>image/ *</mime-type>
// </mime-types> <patterns> <pattern>*.txt</pattern> <pattern>*.png</pattern>
// </patterns> </object> “`
type FileFilter interface {
	Filter
	Buildable

	// AddMIMETypeFileFilter:
	AddMIMETypeFileFilter(mimeType string)
	// AddPatternFileFilter:
	AddPatternFileFilter(pattern string)
	// AddPixbufFormatsFileFilter:
	AddPixbufFormatsFileFilter()
	// Attributes:
	Attributes() []string
	// Name:
	Name() string
	// SetNameFileFilter:
	SetNameFileFilter(name string)
	// ToGVariantFileFilter:
	ToGVariantFileFilter() *glib.Variant
}

// fileFilter implements the FileFilter class.
type fileFilter struct {
	Filter
}

// WrapFileFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileFilter(obj *externglib.Object) FileFilter {
	return fileFilter{
		Filter: WrapFilter(obj),
	}
}

func marshalFileFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileFilter(obj), nil
}

// NewFileFilter:
func NewFileFilter() FileFilter {
	var _cret *C.GtkFileFilter // in

	_cret = C.gtk_file_filter_new()

	var _fileFilter FileFilter // out

	_fileFilter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileFilter)

	return _fileFilter
}

// NewFileFilterFromGVariant:
func NewFileFilterFromGVariant(variant *glib.Variant) FileFilter {
	var _arg1 *C.GVariant      // out
	var _cret *C.GtkFileFilter // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	_cret = C.gtk_file_filter_new_from_gvariant(_arg1)

	var _fileFilter FileFilter // out

	_fileFilter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileFilter)

	return _fileFilter
}

func (f fileFilter) AddMIMETypeFileFilter(mimeType string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_mime_type(_arg0, _arg1)
}

func (f fileFilter) AddPatternFileFilter(pattern string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(pattern))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_add_pattern(_arg0, _arg1)
}

func (f fileFilter) AddPixbufFormatsFileFilter() {
	var _arg0 *C.GtkFileFilter // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	C.gtk_file_filter_add_pixbuf_formats(_arg0)
}

func (f fileFilter) Attributes() []string {
	var _arg0 *C.GtkFileFilter // out
	var _cret **C.char

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_get_attributes(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

func (f fileFilter) Name() string {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f fileFilter) SetNameFileFilter(name string) {
	var _arg0 *C.GtkFileFilter // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_filter_set_name(_arg0, _arg1)
}

func (f fileFilter) ToGVariantFileFilter() *glib.Variant {
	var _arg0 *C.GtkFileFilter // out
	var _cret *C.GVariant      // in

	_arg0 = (*C.GtkFileFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_file_filter_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))

	return _variant
}

func (b fileFilter) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
