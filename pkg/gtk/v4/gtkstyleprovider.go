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
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProvider},
	})
}

// StyleProvider: `GtkStyleProvider` is an interface for style information used
// by `GtkStyleContext`.
//
// See [method@Gtk.StyleContext.add_provider] and
// [func@Gtk.StyleContext.add_provider_for_display] for adding
// `GtkStyleProviders`.
//
// GTK uses the `GtkStyleProvider` implementation for CSS in
// [iface@Gtk.CssProvider].
type StyleProvider interface {
	gextras.Objector
}

// styleProvider implements the StyleProvider interface.
type styleProvider struct {
	gextras.Objector
}

var _ StyleProvider = (*styleProvider)(nil)

// WrapStyleProvider wraps a GObject to a type that implements
// interface StyleProvider. It is primarily used internally.
func WrapStyleProvider(obj *externglib.Object) StyleProvider {
	return styleProvider{
		Objector: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProvider(obj), nil
}
