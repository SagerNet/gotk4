// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_dbus_annotation_info_get_type()), F: marshalDBusAnnotationInfo},
		{T: externglib.Type(C.g_dbus_arg_info_get_type()), F: marshalDBusArgInfo},
		{T: externglib.Type(C.g_dbus_interface_info_get_type()), F: marshalDBusInterfaceInfo},
		{T: externglib.Type(C.g_dbus_method_info_get_type()), F: marshalDBusMethodInfo},
		{T: externglib.Type(C.g_dbus_node_info_get_type()), F: marshalDBusNodeInfo},
		{T: externglib.Type(C.g_dbus_property_info_get_type()), F: marshalDBusPropertyInfo},
		{T: externglib.Type(C.g_dbus_signal_info_get_type()), F: marshalDBusSignalInfo},
	})
}

// DBusAnnotationInfo: information about an annotation.
type DBusAnnotationInfo struct {
	native C.GDBusAnnotationInfo
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusAnnotationInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusAnnotationInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusAnnotationInfo) ref() *DBusAnnotationInfo {
	var _arg0 *C.GDBusAnnotationInfo // out
	var _cret *C.GDBusAnnotationInfo // in

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_annotation_info_ref(_arg0)

	var _dBusAnnotationInfo *DBusAnnotationInfo // out

	_dBusAnnotationInfo = (*DBusAnnotationInfo)(unsafe.Pointer(_cret))
	C.g_dbus_annotation_info_ref(_cret)
	runtime.SetFinalizer(_dBusAnnotationInfo, func(v *DBusAnnotationInfo) {
		C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(unsafe.Pointer(v)))
	})

	return _dBusAnnotationInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusAnnotationInfo) unref() {
	var _arg0 *C.GDBusAnnotationInfo // out

	_arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(info))

	C.g_dbus_annotation_info_unref(_arg0)
}

// DBusArgInfo: information about an argument for a method or a signal.
type DBusArgInfo struct {
	native C.GDBusArgInfo
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusArgInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusArgInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusArgInfo) ref() *DBusArgInfo {
	var _arg0 *C.GDBusArgInfo // out
	var _cret *C.GDBusArgInfo // in

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_arg_info_ref(_arg0)

	var _dBusArgInfo *DBusArgInfo // out

	_dBusArgInfo = (*DBusArgInfo)(unsafe.Pointer(_cret))
	C.g_dbus_arg_info_ref(_cret)
	runtime.SetFinalizer(_dBusArgInfo, func(v *DBusArgInfo) {
		C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(unsafe.Pointer(v)))
	})

	return _dBusArgInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusArgInfo) unref() {
	var _arg0 *C.GDBusArgInfo // out

	_arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(info))

	C.g_dbus_arg_info_unref(_arg0)
}

// DBusInterfaceInfo: information about a D-Bus interface.
type DBusInterfaceInfo struct {
	native C.GDBusInterfaceInfo
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusInterfaceInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with @info, the existing cache is used and
// its use count is increased.
//
// Note that @info cannot be modified until
// g_dbus_interface_info_cache_release() is called.
func (info *DBusInterfaceInfo) CacheBuild() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_interface_info_cache_build(_arg0)
}

// CacheRelease decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (info *DBusInterfaceInfo) CacheRelease() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_interface_info_cache_release(_arg0)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (info *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusMethodInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_method(_arg0, _arg1)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	C.g_dbus_method_info_ref(_cret)
	runtime.SetFinalizer(_dBusMethodInfo, func(v *DBusMethodInfo) {
		C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
	})

	return _dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (info *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusPropertyInfo  // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_property(_arg0, _arg1)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	C.g_dbus_property_info_ref(_cret)
	runtime.SetFinalizer(_dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
	})

	return _dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (info *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusSignalInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_signal(_arg0, _arg1)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	C.g_dbus_signal_info_ref(_cret)
	runtime.SetFinalizer(_dBusSignalInfo, func(v *DBusSignalInfo) {
		C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
	})

	return _dBusSignalInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusInterfaceInfo) ref() *DBusInterfaceInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_interface_info_ref(_arg0)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusInterfaceInfo) unref() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info))

	C.g_dbus_interface_info_unref(_arg0)
}

// DBusMethodInfo: information about a method on an D-Bus interface.
type DBusMethodInfo struct {
	native C.GDBusMethodInfo
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusMethodInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusMethodInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusMethodInfo) ref() *DBusMethodInfo {
	var _arg0 *C.GDBusMethodInfo // out
	var _cret *C.GDBusMethodInfo // in

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_method_info_ref(_arg0)

	var _dBusMethodInfo *DBusMethodInfo // out

	_dBusMethodInfo = (*DBusMethodInfo)(unsafe.Pointer(_cret))
	C.g_dbus_method_info_ref(_cret)
	runtime.SetFinalizer(_dBusMethodInfo, func(v *DBusMethodInfo) {
		C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(unsafe.Pointer(v)))
	})

	return _dBusMethodInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusMethodInfo) unref() {
	var _arg0 *C.GDBusMethodInfo // out

	_arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(info))

	C.g_dbus_method_info_unref(_arg0)
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
type DBusNodeInfo struct {
	native C.GDBusNodeInfo
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusNodeInfo)(unsafe.Pointer(b)), nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (*DBusNodeInfo, error) {
	var _arg1 *C.gchar         // out
	var _cret *C.GDBusNodeInfo // in
	var _cerr *C.GError        // in

	_arg1 = (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_new_for_xml(_arg1, &_cerr)

	var _dBusNodeInfo *DBusNodeInfo // out
	var _goerr error                // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	C.g_dbus_node_info_ref(_cret)
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _dBusNodeInfo, _goerr
}

// Native returns the underlying C source pointer.
func (d *DBusNodeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
func (info *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var _arg0 *C.GDBusNodeInfo      // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(info))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_lookup_interface(_arg0, _arg1)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	_dBusInterfaceInfo = (*DBusInterfaceInfo)(unsafe.Pointer(_cret))
	C.g_dbus_interface_info_ref(_cret)
	runtime.SetFinalizer(_dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(unsafe.Pointer(v)))
	})

	return _dBusInterfaceInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusNodeInfo) ref() *DBusNodeInfo {
	var _arg0 *C.GDBusNodeInfo // out
	var _cret *C.GDBusNodeInfo // in

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_node_info_ref(_arg0)

	var _dBusNodeInfo *DBusNodeInfo // out

	_dBusNodeInfo = (*DBusNodeInfo)(unsafe.Pointer(_cret))
	C.g_dbus_node_info_ref(_cret)
	runtime.SetFinalizer(_dBusNodeInfo, func(v *DBusNodeInfo) {
		C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(unsafe.Pointer(v)))
	})

	return _dBusNodeInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusNodeInfo) unref() {
	var _arg0 *C.GDBusNodeInfo // out

	_arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(info))

	C.g_dbus_node_info_unref(_arg0)
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
type DBusPropertyInfo struct {
	native C.GDBusPropertyInfo
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusPropertyInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusPropertyInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusPropertyInfo) ref() *DBusPropertyInfo {
	var _arg0 *C.GDBusPropertyInfo // out
	var _cret *C.GDBusPropertyInfo // in

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_property_info_ref(_arg0)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	_dBusPropertyInfo = (*DBusPropertyInfo)(unsafe.Pointer(_cret))
	C.g_dbus_property_info_ref(_cret)
	runtime.SetFinalizer(_dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(unsafe.Pointer(v)))
	})

	return _dBusPropertyInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusPropertyInfo) unref() {
	var _arg0 *C.GDBusPropertyInfo // out

	_arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(info))

	C.g_dbus_property_info_unref(_arg0)
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
type DBusSignalInfo struct {
	native C.GDBusSignalInfo
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*DBusSignalInfo)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusSignalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (info *DBusSignalInfo) ref() *DBusSignalInfo {
	var _arg0 *C.GDBusSignalInfo // out
	var _cret *C.GDBusSignalInfo // in

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(info))

	_cret = C.g_dbus_signal_info_ref(_arg0)

	var _dBusSignalInfo *DBusSignalInfo // out

	_dBusSignalInfo = (*DBusSignalInfo)(unsafe.Pointer(_cret))
	C.g_dbus_signal_info_ref(_cret)
	runtime.SetFinalizer(_dBusSignalInfo, func(v *DBusSignalInfo) {
		C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(unsafe.Pointer(v)))
	})

	return _dBusSignalInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (info *DBusSignalInfo) unref() {
	var _arg0 *C.GDBusSignalInfo // out

	_arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(info))

	C.g_dbus_signal_info_unref(_arg0)
}
