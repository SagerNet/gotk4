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
		{T: externglib.Type(C.g_mount_get_type()), F: marshalMount},
	})
}

// Mount: the #GMount interface represents user-visible mounts. Note, when
// porting from GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an @error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
type Mount interface {
	gextras.Objector

	// CanEject decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	CanEject() bool
	// CanUnmount decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	CanUnmount() bool
	// Eject decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectFinish decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	EjectFinish(result AsyncResult) error
	// EjectWithOperation decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// EjectWithOperationFinish decrements the shadow count on @mount. Usually
	// used by Monitor implementations when destroying a shadow mount for
	// @mount, see g_mount_is_shadowed() for more information. The caller will
	// need to emit the #GMount::changed signal on @mount manually.
	EjectWithOperationFinish(result AsyncResult) error
	// DefaultLocation decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	DefaultLocation() File
	// Drive decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Drive() Drive
	// Icon decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Icon() Icon
	// Name decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Name() string
	// Root decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Root() File
	// SortKey decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	SortKey() string
	// SymbolicIcon decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	SymbolicIcon() Icon
	// UUID decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	UUID() string
	// Volume decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Volume() Volume
	// GuessContentType decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	GuessContentType(forceRescan bool, cancellable Cancellable, callback AsyncReadyCallback)
	// GuessContentTypeFinish decrements the shadow count on @mount. Usually
	// used by Monitor implementations when destroying a shadow mount for
	// @mount, see g_mount_is_shadowed() for more information. The caller will
	// need to emit the #GMount::changed signal on @mount manually.
	GuessContentTypeFinish(result AsyncResult) ([]string, error)
	// GuessContentTypeSync decrements the shadow count on @mount. Usually used
	// by Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	GuessContentTypeSync(forceRescan bool, cancellable Cancellable) ([]string, error)
	// IsShadowed decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	IsShadowed() bool
	// Remount decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Remount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// RemountFinish decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	RemountFinish(result AsyncResult) error
	// Shadow decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Shadow()
	// Unmount decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Unmount(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountFinish decrements the shadow count on @mount. Usually used by
	// Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	UnmountFinish(result AsyncResult) error
	// UnmountWithOperation decrements the shadow count on @mount. Usually used
	// by Monitor implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	UnmountWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback)
	// UnmountWithOperationFinish decrements the shadow count on @mount. Usually
	// used by Monitor implementations when destroying a shadow mount for
	// @mount, see g_mount_is_shadowed() for more information. The caller will
	// need to emit the #GMount::changed signal on @mount manually.
	UnmountWithOperationFinish(result AsyncResult) error
	// Unshadow decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Unshadow()
}

// mount implements the Mount interface.
type mount struct {
	gextras.Objector
}

var _ Mount = (*mount)(nil)

// WrapMount wraps a GObject to a type that implements
// interface Mount. It is primarily used internally.
func WrapMount(obj *externglib.Object) Mount {
	return mount{
		Objector: obj,
	}
}

func marshalMount(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMount(obj), nil
}

func (m mount) CanEject() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_can_eject(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m mount) CanUnmount() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_can_unmount(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m mount) Eject(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_eject(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (m mount) EjectFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m mount) EjectWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_eject_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (m mount) EjectWithOperationFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_eject_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m mount) DefaultLocation() File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_default_location(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

func (m mount) Drive() Drive {
	var _arg0 *C.GMount // out
	var _cret *C.GDrive // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_drive(_arg0)

	var _drive Drive // out

	_drive = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Drive)

	return _drive
}

func (m mount) Icon() Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (m mount) Name() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (m mount) Root() File {
	var _arg0 *C.GMount // out
	var _cret *C.GFile  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_root(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

func (m mount) SortKey() string {
	var _arg0 *C.GMount // out
	var _cret *C.gchar  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_sort_key(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (m mount) SymbolicIcon() Icon {
	var _arg0 *C.GMount // out
	var _cret *C.GIcon  // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Icon)

	return _icon
}

func (m mount) UUID() string {
	var _arg0 *C.GMount // out
	var _cret *C.char   // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_uuid(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (m mount) Volume() Volume {
	var _arg0 *C.GMount  // out
	var _cret *C.GVolume // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_get_volume(_arg0)

	var _volume Volume // out

	_volume = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Volume)

	return _volume
}

func (m mount) GuessContentType(forceRescan bool, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.gboolean            // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	if forceRescan {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_guess_content_type(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (m mount) GuessContentTypeFinish(result AsyncResult) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_mount_guess_content_type_finish(_arg0, _arg1, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

func (m mount) GuessContentTypeSync(forceRescan bool, cancellable Cancellable) ([]string, error) {
	var _arg0 *C.GMount       // out
	var _arg1 C.gboolean      // out
	var _arg2 *C.GCancellable // out
	var _cret **C.gchar
	var _cerr *C.GError // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	if forceRescan {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_mount_guess_content_type_sync(_arg0, _arg1, _arg2, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

func (m mount) IsShadowed() bool {
	var _arg0 *C.GMount  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	_cret = C.g_mount_is_shadowed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m mount) Remount(flags MountMountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountMountFlags    // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountMountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_remount(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (m mount) RemountFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_remount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m mount) Shadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_shadow(_arg0)
}

func (m mount) Unmount(flags MountUnmountFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_mount_unmount(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (m mount) UnmountFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m mount) UnmountWithOperation(flags MountUnmountFlags, mountOperation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GMount             // out
	var _arg1 C.GMountUnmountFlags  // out
	var _arg2 *C.GMountOperation    // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = C.GMountUnmountFlags(flags)
	_arg2 = (*C.GMountOperation)(unsafe.Pointer(mountOperation.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_mount_unmount_with_operation(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (m mount) UnmountWithOperationFinish(result AsyncResult) error {
	var _arg0 *C.GMount       // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_mount_unmount_with_operation_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m mount) Unshadow() {
	var _arg0 *C.GMount // out

	_arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_unshadow(_arg0)
}