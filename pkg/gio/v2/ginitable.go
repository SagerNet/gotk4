// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_initable_get_type()), F: marshalInitabler},
	})
}

// InitablerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type InitablerOverrider interface {
	// Init initializes the object implementing the interface.
	//
	// This method is intended for language bindings. If writing in C,
	// g_initable_new() should typically be used instead.
	//
	// The object must be initialized before any real use after initial
	// construction, either with this function or g_async_initable_init_async().
	//
	// Implementations may also support cancellation. If @cancellable is not
	// nil, then initialization can be cancelled by triggering the cancellable
	// object from another thread. If the operation was cancelled, the error
	// G_IO_ERROR_CANCELLED will be returned. If @cancellable is not nil and the
	// object doesn't support cancellable initialization the error
	// G_IO_ERROR_NOT_SUPPORTED will be returned.
	//
	// If the object is not initialized, or initialization returns with an
	// error, then all operations on the object except g_object_ref() and
	// g_object_unref() are considered to be invalid, and have undefined
	// behaviour. See the [introduction][ginitable] for more details.
	//
	// Callers should not assume that a class which implements #GInitable can be
	// initialized multiple times, unless the class explicitly documents itself
	// as supporting this. Generally, a class’ implementation of init() can
	// assume (and assert) that it will only be called once. Previously, this
	// documentation recommended all #GInitable implementations should be
	// idempotent; that recommendation was relaxed in GLib 2.54.
	//
	// If a class explicitly supports being initialized multiple times, it is
	// recommended that the method is idempotent: multiple calls with the same
	// arguments should return the same results. Only the first call initializes
	// the object; further calls return the result of the first call.
	//
	// One reason why a class might need to support idempotent initialization is
	// if it is designed to be used via the singleton pattern, with a
	// Class.constructor that sometimes returns an existing instance. In this
	// pattern, a caller would expect to be able to call g_initable_init() on
	// the result of g_object_new(), regardless of whether it is in fact a new
	// instance.
	Init(cancellable Cancellabler) error
}

// Initabler describes Initable's methods.
type Initabler interface {
	gextras.Objector

	Init(cancellable Cancellabler) error
}

// Initable is implemented by objects that can fail during initialization. If an
// object implements this interface then it must be initialized as the first
// thing after construction, either via g_initable_init() or
// g_async_initable_init_async() (the latter is only available if it also
// implements Initable).
//
// If the object is not initialized, or initialization returns with an error,
// then all operations on the object except g_object_ref() and g_object_unref()
// are considered to be invalid, and have undefined behaviour. They will often
// fail with g_critical() or g_warning(), but this must not be relied on.
//
// Users of objects implementing this are not intended to use the interface
// method directly, instead it will be used automatically in various ways. For C
// applications you generally just call g_initable_new() directly, or indirectly
// via a foo_thing_new() wrapper. This will call g_initable_init() under the
// cover, returning nil and setting a #GError on failure (at which point the
// instance is unreferenced).
//
// For bindings in languages where the native constructor supports exceptions
// the binding could check for objects implementing GInitable during normal
// construction and automatically initialize them, throwing an exception on
// failure.
type Initable struct {
	*externglib.Object
}

var _ Initabler = (*Initable)(nil)

func wrapInitabler(obj *externglib.Object) Initabler {
	return &Initable{
		Object: obj,
	}
}

func marshalInitabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInitabler(obj), nil
}

// Init initializes the object implementing the interface.
//
// This method is intended for language bindings. If writing in C,
// g_initable_new() should typically be used instead.
//
// The object must be initialized before any real use after initial
// construction, either with this function or g_async_initable_init_async().
//
// Implementations may also support cancellation. If @cancellable is not nil,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// G_IO_ERROR_CANCELLED will be returned. If @cancellable is not nil and the
// object doesn't support cancellable initialization the error
// G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// If the object is not initialized, or initialization returns with an error,
// then all operations on the object except g_object_ref() and g_object_unref()
// are considered to be invalid, and have undefined behaviour. See the
// [introduction][ginitable] for more details.
//
// Callers should not assume that a class which implements #GInitable can be
// initialized multiple times, unless the class explicitly documents itself as
// supporting this. Generally, a class’ implementation of init() can assume (and
// assert) that it will only be called once. Previously, this documentation
// recommended all #GInitable implementations should be idempotent; that
// recommendation was relaxed in GLib 2.54.
//
// If a class explicitly supports being initialized multiple times, it is
// recommended that the method is idempotent: multiple calls with the same
// arguments should return the same results. Only the first call initializes the
// object; further calls return the result of the first call.
//
// One reason why a class might need to support idempotent initialization is if
// it is designed to be used via the singleton pattern, with a Class.constructor
// that sometimes returns an existing instance. In this pattern, a caller would
// expect to be able to call g_initable_init() on the result of g_object_new(),
// regardless of whether it is in fact a new instance.
func (initable *Initable) Init(cancellable Cancellabler) error {
	var _arg0 *C.GInitable    // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GInitable)(unsafe.Pointer(initable.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_initable_init(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
