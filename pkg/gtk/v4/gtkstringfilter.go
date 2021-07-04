// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_string_filter_match_mode_get_type()), F: marshalStringFilterMatchMode},
		{T: externglib.Type(C.gtk_string_filter_get_type()), F: marshalStringFilter},
	})
}

// StringFilterMatchMode specifies how search strings are matched inside text.
type StringFilterMatchMode int

const (
	// exact: the search string and text must match exactly.
	StringFilterMatchModeExact StringFilterMatchMode = 0
	// substring: the search string must be contained as a substring inside the
	// text.
	StringFilterMatchModeSubstring StringFilterMatchMode = 1
	// prefix: the text must begin with the search string.
	StringFilterMatchModePrefix StringFilterMatchMode = 2
)

func marshalStringFilterMatchMode(p uintptr) (interface{}, error) {
	return StringFilterMatchMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StringFilter: `GtkStringFilter` determines whether to include items by
// comparing strings to a fixed search term.
//
// The strings are obtained from the items by evaluating a `GtkExpression` set
// with [method@Gtk.StringFilter.set_expression], and they are compared against
// a search term set with [method@Gtk.StringFilter.set_search].
//
// `GtkStringFilter` has several different modes of comparison - it can match
// the whole string, just a prefix, or any substring. Use
// [method@Gtk.StringFilter.set_match_mode] choose a mode.
//
// It is also possible to make case-insensitive comparisons, with
// [method@Gtk.StringFilter.set_ignore_case].
type StringFilter interface {
	Filter

	// Expression:
	Expression() Expression
	// IgnoreCase:
	IgnoreCase() bool
	// MatchMode:
	MatchMode() StringFilterMatchMode
	// Search:
	Search() string
	// SetExpressionStringFilter:
	SetExpressionStringFilter(expression Expression)
	// SetIgnoreCaseStringFilter:
	SetIgnoreCaseStringFilter(ignoreCase bool)
	// SetMatchModeStringFilter:
	SetMatchModeStringFilter(mode StringFilterMatchMode)
	// SetSearchStringFilter:
	SetSearchStringFilter(search string)
}

// stringFilter implements the StringFilter class.
type stringFilter struct {
	Filter
}

// WrapStringFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapStringFilter(obj *externglib.Object) StringFilter {
	return stringFilter{
		Filter: WrapFilter(obj),
	}
}

func marshalStringFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStringFilter(obj), nil
}

// NewStringFilter:
func NewStringFilter(expression Expression) StringFilter {
	var _arg1 *C.GtkExpression   // out
	var _cret *C.GtkStringFilter // in

	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_string_filter_new(_arg1)

	var _stringFilter StringFilter // out

	_stringFilter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(StringFilter)

	return _stringFilter
}

func (s stringFilter) Expression() Expression {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.GtkExpression   // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_string_filter_get_expression(_arg0)

	var _expression Expression // out

	_expression = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expression)

	return _expression
}

func (s stringFilter) IgnoreCase() bool {
	var _arg0 *C.GtkStringFilter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_string_filter_get_ignore_case(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stringFilter) MatchMode() StringFilterMatchMode {
	var _arg0 *C.GtkStringFilter         // out
	var _cret C.GtkStringFilterMatchMode // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_string_filter_get_match_mode(_arg0)

	var _stringFilterMatchMode StringFilterMatchMode // out

	_stringFilterMatchMode = StringFilterMatchMode(_cret)

	return _stringFilterMatchMode
}

func (s stringFilter) Search() string {
	var _arg0 *C.GtkStringFilter // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_string_filter_get_search(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s stringFilter) SetExpressionStringFilter(expression Expression) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.GtkExpression   // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_string_filter_set_expression(_arg0, _arg1)
}

func (s stringFilter) SetIgnoreCaseStringFilter(ignoreCase bool) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))
	if ignoreCase {
		_arg1 = C.TRUE
	}

	C.gtk_string_filter_set_ignore_case(_arg0, _arg1)
}

func (s stringFilter) SetMatchModeStringFilter(mode StringFilterMatchMode) {
	var _arg0 *C.GtkStringFilter         // out
	var _arg1 C.GtkStringFilterMatchMode // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkStringFilterMatchMode(mode)

	C.gtk_string_filter_set_match_mode(_arg0, _arg1)
}

func (s stringFilter) SetSearchStringFilter(search string) {
	var _arg0 *C.GtkStringFilter // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkStringFilter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(search))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_string_filter_set_search(_arg0, _arg1)
}
