// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.g_vfs_get_type()), F: marshalVFS},
	})
}

// VfsFileLookupFunc: this function type is used by g_vfs_register_uri_scheme()
// to make it possible for a client to associate an URI scheme to a different
// #GFile implementation.
//
// The client should return a reference to the new file that has been created
// for @uri, or nil to continue with the default implementation.
type VFSFileLookupFunc func(vfs VFS, identifier string) File

//export gotk4_VFSFileLookupFunc
func gotk4_VFSFileLookupFunc(arg0 *C.GVfs, arg1 *C.char, arg2 C.gpointer) *C.GFile {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}
	fn := v.(VFSFileLookupFunc)

	ret := fn(vfs, identifier, userData)

	cret = (*C.GFile)(unsafe.Pointer(ret.Native()))

	return cret
}

// VFS: entry point for using GIO functionality.
type VFS interface {
	gextras.Objector

	// FileForPath gets a #GFile for @path.
	FileForPath(path string) File
	// FileForURI gets a #GFile for @uri.
	//
	// This operation never fails, but the returned object might not support any
	// I/O operation if the URI is malformed or if the URI scheme is not
	// supported.
	FileForURI(uri string) File
	// SupportedURISchemes gets a list of URI schemes supported by @vfs.
	SupportedURISchemes() []string
	// IsActive checks if the VFS is active.
	IsActive() bool
	// ParseName: this operation never fails, but the returned object might not
	// support any I/O operations if the @parse_name cannot be parsed by the
	// #GVfs module.
	ParseName(parseName string) File
	// RegisterURIScheme registers @uri_func and @parse_name_func as the #GFile
	// URI and parse name lookup functions for URIs with a scheme matching
	// @scheme. Note that @scheme is registered only within the running
	// application, as opposed to desktop-wide as it happens with GVfs backends.
	//
	// When a #GFile is requested with an URI containing @scheme (e.g. through
	// g_file_new_for_uri()), @uri_func will be called to allow a custom
	// constructor. The implementation of @uri_func should not be blocking, and
	// must not call g_vfs_register_uri_scheme() or
	// g_vfs_unregister_uri_scheme().
	//
	// When g_file_parse_name() is called with a parse name obtained from such
	// file, @parse_name_func will be called to allow the #GFile to be created
	// again. In that case, it's responsibility of @parse_name_func to make sure
	// the parse name matches what the custom #GFile implementation returned
	// when g_file_get_parse_name() was previously called. The implementation of
	// @parse_name_func should not be blocking, and must not call
	// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
	//
	// It's an error to call this function twice with the same scheme. To
	// unregister a custom URI scheme, use g_vfs_unregister_uri_scheme().
	RegisterURIScheme(scheme string, uriFunc VFSFileLookupFunc, parseNameFunc VFSFileLookupFunc) bool
	// UnregisterURIScheme unregisters the URI handler for @scheme previously
	// registered with g_vfs_register_uri_scheme().
	UnregisterURIScheme(scheme string) bool
}

// vfS implements the VFS interface.
type vfS struct {
	gextras.Objector
}

var _ VFS = (*vfS)(nil)

// WrapVFS wraps a GObject to the right type. It is
// primarily used internally.
func WrapVFS(obj *externglib.Object) VFS {
	return VFS{
		Objector: obj,
	}
}

func marshalVFS(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVFS(obj), nil
}

// FileForPath gets a #GFile for @path.
func (v vfS) FileForPath(path string) File {
	var arg0 *C.GVfs
	var arg1 *C.char

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GFile
	var goret1 File

	cret = C.g_vfs_get_file_for_path(arg0, path)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(File)

	return goret1
}

// FileForURI gets a #GFile for @uri.
//
// This operation never fails, but the returned object might not support any
// I/O operation if the URI is malformed or if the URI scheme is not
// supported.
func (v vfS) FileForURI(uri string) File {
	var arg0 *C.GVfs
	var arg1 *C.char

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GFile
	var goret1 File

	cret = C.g_vfs_get_file_for_uri(arg0, uri)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(File)

	return goret1
}

// SupportedURISchemes gets a list of URI schemes supported by @vfs.
func (v vfS) SupportedURISchemes() []string {
	var arg0 *C.GVfs

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))

	var cret **C.gchar
	var goret1 []string

	cret = C.g_vfs_get_supported_uri_schemes(arg0)

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		goret1 = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
			goret1[i] = C.GoString(src)
		}
	}

	return goret1
}

// IsActive checks if the VFS is active.
func (v vfS) IsActive() bool {
	var arg0 *C.GVfs

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_vfs_is_active(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// ParseName: this operation never fails, but the returned object might not
// support any I/O operations if the @parse_name cannot be parsed by the
// #GVfs module.
func (v vfS) ParseName(parseName string) File {
	var arg0 *C.GVfs
	var arg1 *C.char

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(parseName))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GFile
	var goret1 File

	cret = C.g_vfs_parse_name(arg0, parseName)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(File)

	return goret1
}

// RegisterURIScheme registers @uri_func and @parse_name_func as the #GFile
// URI and parse name lookup functions for URIs with a scheme matching
// @scheme. Note that @scheme is registered only within the running
// application, as opposed to desktop-wide as it happens with GVfs backends.
//
// When a #GFile is requested with an URI containing @scheme (e.g. through
// g_file_new_for_uri()), @uri_func will be called to allow a custom
// constructor. The implementation of @uri_func should not be blocking, and
// must not call g_vfs_register_uri_scheme() or
// g_vfs_unregister_uri_scheme().
//
// When g_file_parse_name() is called with a parse name obtained from such
// file, @parse_name_func will be called to allow the #GFile to be created
// again. In that case, it's responsibility of @parse_name_func to make sure
// the parse name matches what the custom #GFile implementation returned
// when g_file_get_parse_name() was previously called. The implementation of
// @parse_name_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// It's an error to call this function twice with the same scheme. To
// unregister a custom URI scheme, use g_vfs_unregister_uri_scheme().
func (v vfS) RegisterURIScheme(scheme string, uriFunc VFSFileLookupFunc, parseNameFunc VFSFileLookupFunc) bool {
	var arg0 *C.GVfs

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_vfs_register_uri_scheme(arg0, scheme, uriFunc, uriData, uriDestroy, parseNameFunc, parseNameData, parseNameDestroy)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// UnregisterURIScheme unregisters the URI handler for @scheme previously
// registered with g_vfs_register_uri_scheme().
func (v vfS) UnregisterURIScheme(scheme string) bool {
	var arg0 *C.GVfs
	var arg1 *C.char

	arg0 = (*C.GVfs)(unsafe.Pointer(v.Native()))
	arg1 = (*C.char)(C.CString(scheme))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_vfs_unregister_uri_scheme(arg0, scheme)

	goret1 = C.bool(cret) != C.false

	return goret1
}
