// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.g_simple_proxy_resolver_get_type()), F: marshalSimpleProxyResolver},
	})
}

// SimpleProxyResolver is a simple Resolver implementation that handles a single
// default proxy, multiple URI-scheme-specific proxies, and a list of hosts that
// proxies should not be used for.
//
// ProxyResolver is never the default proxy resolver, but it can be used as the
// base class for another proxy resolver implementation, or it can be created
// and used manually, such as with g_socket_client_set_proxy_resolver().
type SimpleProxyResolver interface {
	ProxyResolver

	// SetDefaultProxySimpleProxyResolver:
	SetDefaultProxySimpleProxyResolver(defaultProxy string)
	// SetURIProxySimpleProxyResolver:
	SetURIProxySimpleProxyResolver(uriScheme string, proxy string)
}

// simpleProxyResolver implements the SimpleProxyResolver class.
type simpleProxyResolver struct {
	gextras.Objector
}

// WrapSimpleProxyResolver wraps a GObject to the right type. It is
// primarily used internally.
func WrapSimpleProxyResolver(obj *externglib.Object) SimpleProxyResolver {
	return simpleProxyResolver{
		Objector: obj,
	}
}

func marshalSimpleProxyResolver(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSimpleProxyResolver(obj), nil
}

func (r simpleProxyResolver) SetDefaultProxySimpleProxyResolver(defaultProxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(defaultProxy))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_proxy_resolver_set_default_proxy(_arg0, _arg1)
}

func (r simpleProxyResolver) SetURIProxySimpleProxyResolver(uriScheme string, proxy string) {
	var _arg0 *C.GSimpleProxyResolver // out
	var _arg1 *C.gchar                // out
	var _arg2 *C.gchar                // out

	_arg0 = (*C.GSimpleProxyResolver)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.gchar)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(proxy))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_simple_proxy_resolver_set_uri_proxy(_arg0, _arg1, _arg2)
}

func (r simpleProxyResolver) IsSupported() bool {
	return WrapProxyResolver(gextras.InternObject(r)).IsSupported()
}

func (r simpleProxyResolver) Lookup(uri string, cancellable Cancellable) ([]string, error) {
	return WrapProxyResolver(gextras.InternObject(r)).Lookup(uri, cancellable)
}

func (r simpleProxyResolver) LookupFinish(result AsyncResult) ([]string, error) {
	return WrapProxyResolver(gextras.InternObject(r)).LookupFinish(result)
}
