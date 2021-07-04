// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_section_type_get_type()), F: marshalCSSSectionType},
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSectionType: the different types of sections indicate parts of a CSS
// document as parsed by GTK’s CSS parser. They are oriented towards the CSS
// Grammar (http://www.w3.org/TR/CSS21/grammar.html), but may contain
// extensions.
//
// More types might be added in the future as the parser incorporates more
// features.
type CSSSectionType int

const (
	// document: the section describes a complete document. This section time is
	// the only one where gtk_css_section_get_parent() might return nil.
	CSSSectionTypeDocument CSSSectionType = 0
	// import: the section defines an import rule.
	CSSSectionTypeImport CSSSectionType = 1
	// ColorDefinition: the section defines a color. This is a GTK extension to
	// CSS.
	CSSSectionTypeColorDefinition CSSSectionType = 2
	// BindingSet: the section defines a binding set. This is a GTK extension to
	// CSS.
	CSSSectionTypeBindingSet CSSSectionType = 3
	// ruleset: the section defines a CSS ruleset.
	CSSSectionTypeRuleset CSSSectionType = 4
	// selector: the section defines a CSS selector.
	CSSSectionTypeSelector CSSSectionType = 5
	// declaration: the section defines the declaration of a CSS variable.
	CSSSectionTypeDeclaration CSSSectionType = 6
	// value: the section defines the value of a CSS declaration.
	CSSSectionTypeValue CSSSectionType = 7
	// keyframes: the section defines keyframes. See [CSS
	// Animations](http://dev.w3.org/csswg/css3-animations/#keyframes) for
	// details. Since 3.6
	CSSSectionTypeKeyframes CSSSectionType = 8
)

func marshalCSSSectionType(p uintptr) (interface{}, error) {
	return CSSSectionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CSSSection defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the containing
// region.
type CSSSection C.GtkCssSection

// WrapCSSSection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCSSSection(ptr unsafe.Pointer) *CSSSection {
	return (*CSSSection)(ptr)
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*CSSSection)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(c)
}

// EndLine decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) EndLine() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_end_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EndPosition decrements the reference count on @section, freeing the structure
// if the reference count reaches 0.
func (s *CSSSection) EndPosition() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_end_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// File decrements the reference count on @section, freeing the structure if the
// reference count reaches 0.
func (s *CSSSection) File() gio.File {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GFile         // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_file(_arg0)

	var _file gio.File // out

	_file = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.File)

	return _file
}

// Parent decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Parent() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_parent(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))

	return _cssSection
}

// SectionType decrements the reference count on @section, freeing the structure
// if the reference count reaches 0.
func (s *CSSSection) SectionType() CSSSectionType {
	var _arg0 *C.GtkCssSection    // out
	var _cret C.GtkCssSectionType // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_section_type(_arg0)

	var _cssSectionType CSSSectionType // out

	_cssSectionType = CSSSectionType(_cret)

	return _cssSectionType
}

// StartLine decrements the reference count on @section, freeing the structure
// if the reference count reaches 0.
func (s *CSSSection) StartLine() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_start_line(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StartPosition decrements the reference count on @section, freeing the
// structure if the reference count reaches 0.
func (s *CSSSection) StartPosition() uint {
	var _arg0 *C.GtkCssSection // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_get_start_position(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Ref decrements the reference count on @section, freeing the structure if the
// reference count reaches 0.
func (s *CSSSection) Ref() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_css_section_ref(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_cssSection, func(v **CSSSection) {
		C.free(unsafe.Pointer(v))
	})

	return _cssSection
}

// Unref decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Unref() {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	C.gtk_css_section_unref(_arg0)
}
