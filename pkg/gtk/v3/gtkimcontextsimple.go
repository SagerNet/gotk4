// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_context_simple_get_type()), F: marshalIMContextSimpler},
	})
}

// MAX_COMPOSE_LEN: maximum length of sequences in compose tables.
const MAX_COMPOSE_LEN = 7

// IMContextSimple is a simple input method context supporting table-based input
// methods. It has a built-in table of compose sequences that is derived from
// the X11 Compose files.
//
// GtkIMContextSimple reads additional compose sequences from the first of the
// following files that is found: ~/.config/gtk-3.0/Compose, ~/.XCompose,
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
type IMContextSimple struct {
	IMContext
}

func WrapIMContextSimple(obj *externglib.Object) *IMContextSimple {
	return &IMContextSimple{
		IMContext: IMContext{
			Object: obj,
		},
	}
}

func marshalIMContextSimpler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMContextSimple(obj), nil
}

// NewIMContextSimple creates a new IMContextSimple.
func NewIMContextSimple() *IMContextSimple {
	var _cret *C.GtkIMContext // in

	_cret = C.gtk_im_context_simple_new()

	var _imContextSimple *IMContextSimple // out

	_imContextSimple = WrapIMContextSimple(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _imContextSimple
}

// AddComposeFile adds an additional table from the X11 compose file.
func (contextSimple *IMContextSimple) AddComposeFile(composeFile string) {
	var _arg0 *C.GtkIMContextSimple // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GtkIMContextSimple)(unsafe.Pointer(contextSimple.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(composeFile)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_im_context_simple_add_compose_file(_arg0, _arg1)
	runtime.KeepAlive(contextSimple)
	runtime.KeepAlive(composeFile)
}
