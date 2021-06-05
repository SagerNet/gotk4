// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_permission_get_type()), F: marshalPermission},
	})
}

// Permission: a #GPermission represents the status of the caller's permission
// to perform a certain action.
//
// You can query if the action is currently allowed and if it is possible to
// acquire the permission so that the action will be allowed in the future.
//
// There is also an API to actually acquire the permission and one to release
// it.
//
// As an example, a #GPermission might represent the ability for the user to
// write to a #GSettings object. This #GPermission object could then be used to
// decide if it is appropriate to show a "Click here to unlock" button in a
// dialog and to provide the mechanism to invoke when that button is clicked.
type Permission interface {
	gextras.Objector

	// Acquire attempts to acquire the permission represented by @permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. A simple example is that a
	// dialog may appear asking the user to enter their password.
	//
	// You should check with g_permission_get_can_acquire() before calling this
	// function.
	//
	// If the permission is acquired then true is returned. Otherwise, false is
	// returned and @error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_acquire_async() for the
	// non-blocking version.
	Acquire(cancellable Cancellable) error
	// AcquireAsync attempts to acquire the permission represented by
	// @permission.
	//
	// This is the first half of the asynchronous version of
	// g_permission_acquire().
	AcquireAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// AcquireFinish collects the result of attempting to acquire the permission
	// represented by @permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_acquire().
	AcquireFinish(result AsyncResult) error
	// Allowed gets the value of the 'allowed' property. This property is true
	// if the caller currently has permission to perform the action that
	// @permission represents the permission to perform.
	Allowed() bool
	// CanAcquire gets the value of the 'can-acquire' property. This property is
	// true if it is generally possible to acquire the permission by calling
	// g_permission_acquire().
	CanAcquire() bool
	// CanRelease gets the value of the 'can-release' property. This property is
	// true if it is generally possible to release the permission by calling
	// g_permission_release().
	CanRelease() bool
	// ImplUpdate: this function is called by the #GPermission implementation to
	// update the properties of the permission. You should never call this
	// function except from a #GPermission implementation.
	//
	// GObject notify signals are generated, as appropriate.
	ImplUpdate(allowed bool, canAcquire bool, canRelease bool)
	// Release attempts to release the permission represented by @permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. In most cases the permission
	// will be dropped immediately without further action.
	//
	// You should check with g_permission_get_can_release() before calling this
	// function.
	//
	// If the permission is released then true is returned. Otherwise, false is
	// returned and @error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_release_async() for the
	// non-blocking version.
	Release(cancellable Cancellable) error
	// ReleaseAsync attempts to release the permission represented by
	// @permission.
	//
	// This is the first half of the asynchronous version of
	// g_permission_release().
	ReleaseAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// ReleaseFinish collects the result of attempting to release the permission
	// represented by @permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_release().
	ReleaseFinish(result AsyncResult) error
}

// permission implements the Permission interface.
type permission struct {
	gextras.Objector
}

var _ Permission = (*permission)(nil)

// WrapPermission wraps a GObject to the right type. It is
// primarily used internally.
func WrapPermission(obj *externglib.Object) Permission {
	return Permission{
		Objector: obj,
	}
}

func marshalPermission(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPermission(obj), nil
}

// Acquire attempts to acquire the permission represented by @permission.
//
// The precise method by which this happens depends on the permission and
// the underlying authentication mechanism. A simple example is that a
// dialog may appear asking the user to enter their password.
//
// You should check with g_permission_get_can_acquire() before calling this
// function.
//
// If the permission is acquired then true is returned. Otherwise, false is
// returned and @error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_acquire_async() for the
// non-blocking version.
func (p permission) Acquire(cancellable Cancellable) error {
	var arg0 *C.GPermission
	var arg1 *C.GCancellable

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var goerr error

	C.g_permission_acquire(arg0, cancellable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// AcquireAsync attempts to acquire the permission represented by
// @permission.
//
// This is the first half of the asynchronous version of
// g_permission_acquire().
func (p permission) AcquireAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GPermission

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	C.g_permission_acquire_async(arg0, cancellable, callback, userData)
}

// AcquireFinish collects the result of attempting to acquire the permission
// represented by @permission.
//
// This is the second half of the asynchronous version of
// g_permission_acquire().
func (p permission) AcquireFinish(result AsyncResult) error {
	var arg0 *C.GPermission
	var arg1 *C.GAsyncResult

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var goerr error

	C.g_permission_acquire_finish(arg0, result, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// Allowed gets the value of the 'allowed' property. This property is true
// if the caller currently has permission to perform the action that
// @permission represents the permission to perform.
func (p permission) Allowed() bool {
	var arg0 *C.GPermission

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_permission_get_allowed(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// CanAcquire gets the value of the 'can-acquire' property. This property is
// true if it is generally possible to acquire the permission by calling
// g_permission_acquire().
func (p permission) CanAcquire() bool {
	var arg0 *C.GPermission

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_permission_get_can_acquire(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// CanRelease gets the value of the 'can-release' property. This property is
// true if it is generally possible to release the permission by calling
// g_permission_release().
func (p permission) CanRelease() bool {
	var arg0 *C.GPermission

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_permission_get_can_release(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ImplUpdate: this function is called by the #GPermission implementation to
// update the properties of the permission. You should never call this
// function except from a #GPermission implementation.
//
// GObject notify signals are generated, as appropriate.
func (p permission) ImplUpdate(allowed bool, canAcquire bool, canRelease bool) {
	var arg0 *C.GPermission
	var arg1 C.gboolean
	var arg2 C.gboolean
	var arg3 C.gboolean

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	if allowed {
		arg1 = C.gboolean(1)
	}
	if canAcquire {
		arg2 = C.gboolean(1)
	}
	if canRelease {
		arg3 = C.gboolean(1)
	}

	C.g_permission_impl_update(arg0, allowed, canAcquire, canRelease)
}

// Release attempts to release the permission represented by @permission.
//
// The precise method by which this happens depends on the permission and
// the underlying authentication mechanism. In most cases the permission
// will be dropped immediately without further action.
//
// You should check with g_permission_get_can_release() before calling this
// function.
//
// If the permission is released then true is returned. Otherwise, false is
// returned and @error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_release_async() for the
// non-blocking version.
func (p permission) Release(cancellable Cancellable) error {
	var arg0 *C.GPermission
	var arg1 *C.GCancellable

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var goerr error

	C.g_permission_release(arg0, cancellable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// ReleaseAsync attempts to release the permission represented by
// @permission.
//
// This is the first half of the asynchronous version of
// g_permission_release().
func (p permission) ReleaseAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GPermission

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))

	C.g_permission_release_async(arg0, cancellable, callback, userData)
}

// ReleaseFinish collects the result of attempting to release the permission
// represented by @permission.
//
// This is the second half of the asynchronous version of
// g_permission_release().
func (p permission) ReleaseFinish(result AsyncResult) error {
	var arg0 *C.GPermission
	var arg1 *C.GAsyncResult

	arg0 = (*C.GPermission)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var goerr error

	C.g_permission_release_finish(arg0, result, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

type PermissionPrivate struct {
	native C.GPermissionPrivate
}

// WrapPermissionPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPermissionPrivate(ptr unsafe.Pointer) *PermissionPrivate {
	if ptr == nil {
		return nil
	}

	return (*PermissionPrivate)(ptr)
}

func marshalPermissionPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPermissionPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PermissionPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}