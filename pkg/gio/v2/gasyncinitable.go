// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_async_initable_get_type()), F: marshalAsyncInitabler},
	})
}

// AsyncInitablerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AsyncInitablerOverrider interface {
	// InitAsync starts asynchronous initialization of the object implementing
	// the interface. This must be done before any real use of the object after
	// initial construction. If the object also implements #GInitable you can
	// optionally call g_initable_init() instead.
	//
	// This method is intended for language bindings. If writing in C,
	// g_async_initable_new_async() should typically be used instead.
	//
	// When the initialization is finished, @callback will be called. You can
	// then call g_async_initable_init_finish() to get the result of the
	// initialization.
	//
	// Implementations may also support cancellation. If @cancellable is not
	// nil, then initialization can be cancelled by triggering the cancellable
	// object from another thread. If the operation was cancelled, the error
	// G_IO_ERROR_CANCELLED will be returned. If @cancellable is not nil, and
	// the object doesn't support cancellable initialization, the error
	// G_IO_ERROR_NOT_SUPPORTED will be returned.
	//
	// As with #GInitable, if the object is not initialized, or initialization
	// returns with an error, then all operations on the object except
	// g_object_ref() and g_object_unref() are considered to be invalid, and
	// have undefined behaviour. They will often fail with g_critical() or
	// g_warning(), but this must not be relied on.
	//
	// Callers should not assume that a class which implements Initable can be
	// initialized multiple times; for more information, see g_initable_init().
	// If a class explicitly supports being initialized multiple times,
	// implementation requires yielding all subsequent calls to init_async() on
	// the results of the first call.
	//
	// For classes that also support the #GInitable interface, the default
	// implementation of this method will run the g_initable_init() function in
	// a thread, so if you want to support asynchronous initialization via
	// threads, just implement the Initable interface without overriding any
	// interface methods.
	InitAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	// InitFinish finishes asynchronous initialization and returns the result.
	// See g_async_initable_init_async().
	InitFinish(res AsyncResulter) error
}

// AsyncInitabler describes AsyncInitable's methods.
type AsyncInitabler interface {
	gextras.Objector

	InitAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback)
	InitFinish(res AsyncResulter) error
	NewFinish(res AsyncResulter) (*externglib.Object, error)
}

// AsyncInitable: this is the asynchronous version of #GInitable; it behaves the
// same in all ways except that initialization is asynchronous. For more details
// see the descriptions on #GInitable.
//
// A class may implement both the #GInitable and Initable interfaces.
//
// Users of objects implementing this are not intended to use the interface
// method directly; instead it will be used automatically in various ways. For C
// applications you generally just call g_async_initable_new_async() directly,
// or indirectly via a foo_thing_new_async() wrapper. This will call
// g_async_initable_init_async() under the cover, calling back with nil and a
// set GError on failure.
//
// A typical implementation might look something like this:
//
//    enum {
//       NOT_INITIALIZED,
//       INITIALIZING,
//       INITIALIZED
//    };
//
//    static void
//    _foo_ready_cb (Foo *self)
//    {
//      GList *l;
//
//      self->priv->state = INITIALIZED;
//
//      for (l = self->priv->init_results; l != NULL; l = l->next)
//        {
//          GTask *task = l->data;
//
//          if (self->priv->success)
//            g_task_return_boolean (task, TRUE);
//          else
//            g_task_return_new_error (task, ...);
//          g_object_unref (task);
//        }
//
//      g_list_free (self->priv->init_results);
//      self->priv->init_results = NULL;
//    }
//
//    static void
//    foo_init_async (GAsyncInitable       *initable,
//                    int                   io_priority,
//                    GCancellable         *cancellable,
//                    GAsyncReadyCallback   callback,
//                    gpointer              user_data)
//    {
//      Foo *self = FOO (initable);
//      GTask *task;
//
//      task = g_task_new (initable, cancellable, callback, user_data);
//      g_task_set_name (task, G_STRFUNC);
//
//      switch (self->priv->state)
//        {
//          case NOT_INITIALIZED:
//            _foo_get_ready (self);
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            self->priv->state = INITIALIZING;
//            break;
//          case INITIALIZING:
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            break;
//          case INITIALIZED:
//            if (!self->priv->success)
//              g_task_return_new_error (task, ...);
//            else
//              g_task_return_boolean (task, TRUE);
//            g_object_unref (task);
//            break;
//        }
//    }
//
//    static gboolean
//    foo_init_finish (GAsyncInitable       *initable,
//                     GAsyncResult         *result,
//                     GError              **error)
//    {
//      g_return_val_if_fail (g_task_is_valid (result, initable), FALSE);
//
//      return g_task_propagate_boolean (G_TASK (result), error);
//    }
//
//    static void
//    foo_async_initable_iface_init (gpointer g_iface,
//                                   gpointer data)
//    {
//      GAsyncInitableIface *iface = g_iface;
//
//      iface->init_async = foo_init_async;
//      iface->init_finish = foo_init_finish;
//    }
type AsyncInitable struct {
	*externglib.Object
}

var _ AsyncInitabler = (*AsyncInitable)(nil)

func wrapAsyncInitabler(obj *externglib.Object) AsyncInitabler {
	return &AsyncInitable{
		Object: obj,
	}
}

func marshalAsyncInitabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAsyncInitabler(obj), nil
}

// InitAsync starts asynchronous initialization of the object implementing the
// interface. This must be done before any real use of the object after initial
// construction. If the object also implements #GInitable you can optionally
// call g_initable_init() instead.
//
// This method is intended for language bindings. If writing in C,
// g_async_initable_new_async() should typically be used instead.
//
// When the initialization is finished, @callback will be called. You can then
// call g_async_initable_init_finish() to get the result of the initialization.
//
// Implementations may also support cancellation. If @cancellable is not nil,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// G_IO_ERROR_CANCELLED will be returned. If @cancellable is not nil, and the
// object doesn't support cancellable initialization, the error
// G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// As with #GInitable, if the object is not initialized, or initialization
// returns with an error, then all operations on the object except
// g_object_ref() and g_object_unref() are considered to be invalid, and have
// undefined behaviour. They will often fail with g_critical() or g_warning(),
// but this must not be relied on.
//
// Callers should not assume that a class which implements Initable can be
// initialized multiple times; for more information, see g_initable_init(). If a
// class explicitly supports being initialized multiple times, implementation
// requires yielding all subsequent calls to init_async() on the results of the
// first call.
//
// For classes that also support the #GInitable interface, the default
// implementation of this method will run the g_initable_init() function in a
// thread, so if you want to support asynchronous initialization via threads,
// just implement the Initable interface without overriding any interface
// methods.
func (initable *AsyncInitable) InitAsync(ioPriority int, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GAsyncInitable     // out
	var _arg1 C.int                 // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(initable.Native()))
	_arg1 = C.int(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_async_initable_init_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// InitFinish finishes asynchronous initialization and returns the result. See
// g_async_initable_init_async().
func (initable *AsyncInitable) InitFinish(res AsyncResulter) error {
	var _arg0 *C.GAsyncInitable // out
	var _arg1 *C.GAsyncResult   // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(initable.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	C.g_async_initable_init_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// NewFinish finishes the async construction for the various
// g_async_initable_new calls, returning the created object or nil on error.
func (initable *AsyncInitable) NewFinish(res AsyncResulter) (*externglib.Object, error) {
	var _arg0 *C.GAsyncInitable // out
	var _arg1 *C.GAsyncResult   // out
	var _cret *C.GObject        // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GAsyncInitable)(unsafe.Pointer(initable.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	_cret = C.g_async_initable_new_finish(_arg0, _arg1, &_cerr)

	var _object *externglib.Object // out
	var _goerr error               // out

	_object = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*externglib.Object)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _object, _goerr
}
