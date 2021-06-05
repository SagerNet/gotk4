// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_simple_get_type()), F: marshalIMContextSimple},
	})
}

// IMContextSimple: gtkIMContextSimple is a simple input method context
// supporting table-based input methods. It has a built-in table of compose
// sequences that is derived from the X11 Compose files.
//
// GtkIMContextSimple reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-4.0/Compose, ~/.XCompose,
// /usr/share/X11/locale/$locale/Compose (for locales that have a nontrivial
// Compose file). The syntax of these files is described in the Compose(5)
// manual page.
//
//
// Unicode characters
//
// GtkIMContextSimple also supports numeric entry of Unicode characters by
// typing Ctrl-Shift-u, followed by a hexadecimal Unicode codepoint. For
// example, Ctrl-Shift-u 1 2 3 Enter yields U+0123 LATIN SMALL LETTER G WITH
// CEDILLA, i.e. ģ.
type IMContextSimple interface {
	IMContext

	// AddComposeFile adds an additional table from the X11 compose file.
	AddComposeFile(composeFile string)
}

// imContextSimple implements the IMContextSimple interface.
type imContextSimple struct {
	IMContext
}

var _ IMContextSimple = (*imContextSimple)(nil)

// WrapIMContextSimple wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMContextSimple(obj *externglib.Object) IMContextSimple {
	return IMContextSimple{
		IMContext: WrapIMContext(obj),
	}
}

func marshalIMContextSimple(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMContextSimple(obj), nil
}

// NewIMContextSimple constructs a class IMContextSimple.
func NewIMContextSimple() IMContextSimple {
	var cret C.GtkIMContextSimple
	var ret1 IMContextSimple

	cret = C.gtk_im_context_simple_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(IMContextSimple)

	return ret1
}

// AddComposeFile adds an additional table from the X11 compose file.
func (c imContextSimple) AddComposeFile(composeFile string) {
	var arg0 *C.GtkIMContextSimple
	var arg1 *C.char

	arg0 = (*C.GtkIMContextSimple)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(composeFile))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_im_context_simple_add_compose_file(arg0, composeFile)
}

type IMContextSimplePrivate struct {
	native C.GtkIMContextSimplePrivate
}

// WrapIMContextSimplePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIMContextSimplePrivate(ptr unsafe.Pointer) *IMContextSimplePrivate {
	if ptr == nil {
		return nil
	}

	return (*IMContextSimplePrivate)(ptr)
}

func marshalIMContextSimplePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapIMContextSimplePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *IMContextSimplePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}