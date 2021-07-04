// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// BindingsActivate: find a key binding matching @keyval and @modifiers and
// activate the binding on @object.
func BindingsActivate(object gextras.Objector, keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GObject        // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	_cret = C.gtk_bindings_activate(_arg1, _arg2, _arg3)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingsActivateEvent looks up key bindings for @object to find one matching
// @event, and if one was found, activate it.
func BindingsActivateEvent(object gextras.Objector, event *gdk.EventKey) bool {
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.GdkEventKey)(unsafe.Pointer(event.Native()))

	_cret = C.gtk_bindings_activate_event(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingArg: a BindingArg holds the data associated with an argument for a key
// binding signal emission as stored in BindingSignal.
type BindingArg C.GtkBindingArg

// WrapBindingArg wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingArg(ptr unsafe.Pointer) *BindingArg {
	return (*BindingArg)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingArg) Native() unsafe.Pointer {
	return unsafe.Pointer(b)
}

// BindingEntry: each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
type BindingEntry C.GtkBindingEntry

// WrapBindingEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingEntry(ptr unsafe.Pointer) *BindingEntry {
	return (*BindingEntry)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(b)
}

// BindingSignal: a GtkBindingSignal stores the necessary information to
// activate a widget in response to a key press via a signal emission.
type BindingSignal C.GtkBindingSignal

// WrapBindingSignal wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBindingSignal(ptr unsafe.Pointer) *BindingSignal {
	return (*BindingSignal)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BindingSignal) Native() unsafe.Pointer {
	return unsafe.Pointer(b)
}