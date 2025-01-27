// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// BindingsActivate: find a key binding matching keyval and modifiers and
// activate the binding on object.
func BindingsActivate(object *externglib.Object, keyval uint, modifiers gdk.ModifierType) bool {
	var _arg1 *C.GObject        // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _cret C.gboolean        // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	_cret = C.gtk_bindings_activate(_arg1, _arg2, _arg3)
	runtime.KeepAlive(object)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingsActivateEvent looks up key bindings for object to find one matching
// event, and if one was found, activate it.
func BindingsActivateEvent(object *externglib.Object, event *gdk.EventKey) bool {
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.GdkEventKey)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gtk_bindings_activate_event(_arg1, _arg2)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// BindingArg holds the data associated with an argument for a key binding
// signal emission as stored in BindingSignal.
//
// An instance of this type is always passed by reference.
type BindingArg struct {
	*bindingArg
}

// bindingArg is the struct that's finalized.
type bindingArg struct {
	native *C.GtkBindingArg
}

// ArgType: implementation detail
func (b *BindingArg) ArgType() externglib.Type {
	var v externglib.Type // out
	v = externglib.Type(b.native.arg_type)
	return v
}

// BindingEntry: each key binding element of a binding sets binding list is
// represented by a GtkBindingEntry.
//
// An instance of this type is always passed by reference.
type BindingEntry struct {
	*bindingEntry
}

// bindingEntry is the struct that's finalized.
type bindingEntry struct {
	native *C.GtkBindingEntry
}

// Keyval: key value to match
func (b *BindingEntry) Keyval() uint {
	var v uint // out
	v = uint(b.native.keyval)
	return v
}

// Modifiers: key modifiers to match
func (b *BindingEntry) Modifiers() gdk.ModifierType {
	var v gdk.ModifierType // out
	v = gdk.ModifierType(b.native.modifiers)
	return v
}

// BindingSet: binding set this entry belongs to
func (b *BindingEntry) BindingSet() *BindingSet {
	var v *BindingSet // out
	v = (*BindingSet)(gextras.NewStructNative(unsafe.Pointer(b.native.binding_set)))
	return v
}

// SetNext: linked list of entries maintained by binding set
func (b *BindingEntry) SetNext() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.set_next)))
	return v
}

// HashNext: implementation detail
func (b *BindingEntry) HashNext() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.hash_next)))
	return v
}

// Signals: action signals of this entry
func (b *BindingEntry) Signals() *BindingSignal {
	var v *BindingSignal // out
	v = (*BindingSignal)(gextras.NewStructNative(unsafe.Pointer(b.native.signals)))
	return v
}

// BindingEntryAddSignalFromString parses a signal description from signal_desc
// and incorporates it into binding_set.
//
// Signal descriptions may either bind a key combination to one or more signals:
//
//    bind "key" {
//      "signalname" (param, ...)
//      ...
//    }
//
// Or they may also unbind a key combination:
//
//    unbind "key"
//
// Key combinations must be in a format that can be parsed by
// gtk_accelerator_parse().
func BindingEntryAddSignalFromString(bindingSet *BindingSet, signalDesc string) glib.TokenType {
	var _arg1 *C.GtkBindingSet // out
	var _arg2 *C.gchar         // out
	var _cret C.GTokenType     // in

	_arg1 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(signalDesc)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_binding_entry_add_signal_from_string(_arg1, _arg2)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(signalDesc)

	var _tokenType glib.TokenType // out

	_tokenType = glib.TokenType(_cret)

	return _tokenType
}

// BindingEntryAddSignall: override or install a new key binding for keyval with
// modifiers on binding_set.
func BindingEntryAddSignall(bindingSet *BindingSet, keyval uint, modifiers gdk.ModifierType, signalName string, bindingArgs []BindingArg) {
	var _arg1 *C.GtkBindingSet  // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out
	var _arg4 *C.gchar          // out
	var _arg5 *C.GSList         // out

	_arg1 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(signalName)))
	defer C.free(unsafe.Pointer(_arg4))
	for i := len(bindingArgs) - 1; i >= 0; i-- {
		src := bindingArgs[i]
		var dst *C.GtkBindingArg // out
		dst = (*C.GtkBindingArg)(gextras.StructNative(unsafe.Pointer((&src))))
		_arg5 = C.g_slist_prepend(_arg5, C.gpointer(unsafe.Pointer(dst)))
	}
	defer C.g_slist_free(_arg5)

	C.gtk_binding_entry_add_signall(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
	runtime.KeepAlive(signalName)
	runtime.KeepAlive(bindingArgs)
}

// BindingEntryRemove: remove a binding previously installed via
// gtk_binding_entry_add_signal() on binding_set.
func BindingEntryRemove(bindingSet *BindingSet, keyval uint, modifiers gdk.ModifierType) {
	var _arg1 *C.GtkBindingSet  // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	C.gtk_binding_entry_remove(_arg1, _arg2, _arg3)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
}

// BindingEntrySkip: install a binding on binding_set which causes key lookups
// to be aborted, to prevent bindings from lower priority sets to be activated.
func BindingEntrySkip(bindingSet *BindingSet, keyval uint, modifiers gdk.ModifierType) {
	var _arg1 *C.GtkBindingSet  // out
	var _arg2 C.guint           // out
	var _arg3 C.GdkModifierType // out

	_arg1 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg2 = C.guint(keyval)
	_arg3 = C.GdkModifierType(modifiers)

	C.gtk_binding_entry_skip(_arg1, _arg2, _arg3)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
}

// BindingSet: binding set maintains a list of activatable key bindings. A
// single binding set can match multiple types of widgets. Similar to style
// contexts, can be matched by any information contained in a widgets
// WidgetPath. When a binding within a set is matched upon activation, an action
// signal is emitted on the target widget to carry out the actual activation.
//
// An instance of this type is always passed by reference.
type BindingSet struct {
	*bindingSet
}

// bindingSet is the struct that's finalized.
type bindingSet struct {
	native *C.GtkBindingSet
}

// SetName: unique name of this binding set
func (b *BindingSet) SetName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(b.native.set_name)))
	return v
}

// Priority: unused
func (b *BindingSet) Priority() int {
	var v int // out
	v = int(b.native.priority)
	return v
}

// Entries: key binding entries in this binding set
func (b *BindingSet) Entries() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.entries)))
	return v
}

// Current: implementation detail
func (b *BindingSet) Current() *BindingEntry {
	var v *BindingEntry // out
	v = (*BindingEntry)(gextras.NewStructNative(unsafe.Pointer(b.native.current)))
	return v
}

// Activate: find a key binding matching keyval and modifiers within binding_set
// and activate the binding on object.
func (bindingSet *BindingSet) Activate(keyval uint, modifiers gdk.ModifierType, object *externglib.Object) bool {
	var _arg0 *C.GtkBindingSet  // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 *C.GObject        // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg1 = C.guint(keyval)
	_arg2 = C.GdkModifierType(modifiers)
	_arg3 = (*C.GObject)(unsafe.Pointer(object.Native()))

	_cret = C.gtk_binding_set_activate(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(keyval)
	runtime.KeepAlive(modifiers)
	runtime.KeepAlive(object)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddPath: this function was used internally by the GtkRC parsing mechanism to
// assign match patterns to BindingSet structures.
//
// In GTK+ 3, these match patterns are unused.
//
// Deprecated: since version 3.0.
func (bindingSet *BindingSet) AddPath(pathType PathType, pathPattern string, priority PathPriorityType) {
	var _arg0 *C.GtkBindingSet      // out
	var _arg1 C.GtkPathType         // out
	var _arg2 *C.gchar              // out
	var _arg3 C.GtkPathPriorityType // out

	_arg0 = (*C.GtkBindingSet)(gextras.StructNative(unsafe.Pointer(bindingSet)))
	_arg1 = C.GtkPathType(pathType)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(pathPattern)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.GtkPathPriorityType(priority)

	C.gtk_binding_set_add_path(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(bindingSet)
	runtime.KeepAlive(pathType)
	runtime.KeepAlive(pathPattern)
	runtime.KeepAlive(priority)
}

// BindingSetFind: find a binding set by its globally unique name.
//
// The set_name can either be a name used for gtk_binding_set_new() or the type
// name of a class used in gtk_binding_set_by_class().
func BindingSetFind(setName string) *BindingSet {
	var _arg1 *C.gchar         // out
	var _cret *C.GtkBindingSet // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(setName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_binding_set_find(_arg1)
	runtime.KeepAlive(setName)

	var _bindingSet *BindingSet // out

	if _cret != nil {
		_bindingSet = (*BindingSet)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _bindingSet
}

// BindingSignal stores the necessary information to activate a widget in
// response to a key press via a signal emission.
//
// An instance of this type is always passed by reference.
type BindingSignal struct {
	*bindingSignal
}

// bindingSignal is the struct that's finalized.
type bindingSignal struct {
	native *C.GtkBindingSignal
}
