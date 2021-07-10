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
		{T: externglib.Type(C.g_tls_interaction_get_type()), F: marshalTLSInteractioner},
	})
}

// TLSInteractionerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type TLSInteractionerOverrider interface {
	// AskPassword: run synchronous interaction to ask the user for a password.
	// In general, g_tls_interaction_invoke_ask_password() should be used
	// instead of this function.
	//
	// Derived subclasses usually implement a password prompt, although they may
	// also choose to provide a password from elsewhere. The @password value
	// will be filled in and then @callback will be called. Alternatively the
	// user may abort this password request, which will usually abort the TLS
	// connection.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code. Certain implementations may
	// not support immediate cancellation.
	AskPassword(password TLSPassworder, cancellable Cancellabler) (TLSInteractionResult, error)
	// AskPasswordAsync: run asynchronous interaction to ask the user for a
	// password. In general, g_tls_interaction_invoke_ask_password() should be
	// used instead of this function.
	//
	// Derived subclasses usually implement a password prompt, although they may
	// also choose to provide a password from elsewhere. The @password value
	// will be filled in and then @callback will be called. Alternatively the
	// user may abort this password request, which will usually abort the TLS
	// connection.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code. Certain implementations may
	// not support immediate cancellation.
	//
	// Certain implementations may not support immediate cancellation.
	AskPasswordAsync(password TLSPassworder, cancellable Cancellabler, callback AsyncReadyCallback)
	// AskPasswordFinish: complete an ask password user interaction request.
	// This should be once the g_tls_interaction_ask_password_async() completion
	// callback is called.
	//
	// If G_TLS_INTERACTION_HANDLED is returned, then the Password passed to
	// g_tls_interaction_ask_password() will have its password filled in.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code.
	AskPasswordFinish(result AsyncResulter) (TLSInteractionResult, error)
	// RequestCertificateFinish: complete a request certificate user interaction
	// request. This should be once the
	// g_tls_interaction_request_certificate_async() completion callback is
	// called.
	//
	// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
	// g_tls_interaction_request_certificate_async() will have had its
	// Connection:certificate filled in.
	//
	// If the interaction is cancelled by the cancellation object, or by the
	// user then G_TLS_INTERACTION_FAILED will be returned with an error that
	// contains a G_IO_ERROR_CANCELLED error code.
	RequestCertificateFinish(result AsyncResulter) (TLSInteractionResult, error)
}

// TLSInteractioner describes TLSInteraction's methods.
type TLSInteractioner interface {
	gextras.Objector

	AskPassword(password TLSPassworder, cancellable Cancellabler) (TLSInteractionResult, error)
	AskPasswordAsync(password TLSPassworder, cancellable Cancellabler, callback AsyncReadyCallback)
	AskPasswordFinish(result AsyncResulter) (TLSInteractionResult, error)
	InvokeAskPassword(password TLSPassworder, cancellable Cancellabler) (TLSInteractionResult, error)
	RequestCertificateFinish(result AsyncResulter) (TLSInteractionResult, error)
}

// TLSInteraction provides a mechanism for the TLS connection and database code
// to interact with the user. It can be used to ask the user for passwords.
//
// To use a Interaction with a TLS connection use
// g_tls_connection_set_interaction().
//
// Callers should instantiate a derived class that implements the various
// interaction methods to show the required dialogs.
//
// Callers should use the 'invoke' functions like
// g_tls_interaction_invoke_ask_password() to run interaction methods. These
// functions make sure that the interaction is invoked in the main loop and not
// in the current thread, if the current thread is not running the main loop.
//
// Derived classes can choose to implement whichever interactions methods they'd
// like to support by overriding those virtual methods in their class
// initialization function. Any interactions not implemented will return
// G_TLS_INTERACTION_UNHANDLED. If a derived class implements an async method,
// it must also implement the corresponding finish method.
type TLSInteraction struct {
	*externglib.Object
}

var _ TLSInteractioner = (*TLSInteraction)(nil)

func wrapTLSInteractioner(obj *externglib.Object) TLSInteractioner {
	return &TLSInteraction{
		Object: obj,
	}
}

func marshalTLSInteractioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSInteractioner(obj), nil
}

// AskPassword: run synchronous interaction to ask the user for a password. In
// general, g_tls_interaction_invoke_ask_password() should be used instead of
// this function.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The @password value will be
// filled in and then @callback will be called. Alternatively the user may abort
// this password request, which will usually abort the TLS connection.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
func (interaction *TLSInteraction) AskPassword(password TLSPassworder, cancellable Cancellabler) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GTlsPassword         // out
	var _arg2 *C.GCancellable         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_tls_interaction_ask_password(_arg0, _arg1, _arg2, &_cerr)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = (TLSInteractionResult)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsInteractionResult, _goerr
}

// AskPasswordAsync: run asynchronous interaction to ask the user for a
// password. In general, g_tls_interaction_invoke_ask_password() should be used
// instead of this function.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The @password value will be
// filled in and then @callback will be called. Alternatively the user may abort
// this password request, which will usually abort the TLS connection.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
//
// Certain implementations may not support immediate cancellation.
func (interaction *TLSInteraction) AskPasswordAsync(password TLSPassworder, cancellable Cancellabler, callback AsyncReadyCallback) {
	var _arg0 *C.GTlsInteraction    // out
	var _arg1 *C.GTlsPassword       // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_tls_interaction_ask_password_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// AskPasswordFinish: complete an ask password user interaction request. This
// should be once the g_tls_interaction_ask_password_async() completion callback
// is called.
//
// If G_TLS_INTERACTION_HANDLED is returned, then the Password passed to
// g_tls_interaction_ask_password() will have its password filled in.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code.
func (interaction *TLSInteraction) AskPasswordFinish(result AsyncResulter) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_interaction_ask_password_finish(_arg0, _arg1, &_cerr)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = (TLSInteractionResult)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsInteractionResult, _goerr
}

// InvokeAskPassword: invoke the interaction to ask the user for a password. It
// invokes this interaction in the main loop, specifically the Context returned
// by g_main_context_get_thread_default() when the interaction is created. This
// is called by called by Connection or Database to ask the user for a password.
//
// Derived subclasses usually implement a password prompt, although they may
// also choose to provide a password from elsewhere. The @password value will be
// filled in and then @callback will be called. Alternatively the user may abort
// this password request, which will usually abort the TLS connection.
//
// The implementation can either be a synchronous (eg: modal dialog) or an
// asynchronous one (eg: modeless dialog). This function will take care of
// calling which ever one correctly.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code. Certain implementations may not support
// immediate cancellation.
func (interaction *TLSInteraction) InvokeAskPassword(password TLSPassworder, cancellable Cancellabler) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GTlsPassword         // out
	var _arg2 *C.GCancellable         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg1 = (*C.GTlsPassword)(unsafe.Pointer(password.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_tls_interaction_invoke_ask_password(_arg0, _arg1, _arg2, &_cerr)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = (TLSInteractionResult)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsInteractionResult, _goerr
}

// RequestCertificateFinish: complete a request certificate user interaction
// request. This should be once the
// g_tls_interaction_request_certificate_async() completion callback is called.
//
// If G_TLS_INTERACTION_HANDLED is returned, then the Connection passed to
// g_tls_interaction_request_certificate_async() will have had its
// Connection:certificate filled in.
//
// If the interaction is cancelled by the cancellation object, or by the user
// then G_TLS_INTERACTION_FAILED will be returned with an error that contains a
// G_IO_ERROR_CANCELLED error code.
func (interaction *TLSInteraction) RequestCertificateFinish(result AsyncResulter) (TLSInteractionResult, error) {
	var _arg0 *C.GTlsInteraction      // out
	var _arg1 *C.GAsyncResult         // out
	var _cret C.GTlsInteractionResult // in
	var _cerr *C.GError               // in

	_arg0 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_tls_interaction_request_certificate_finish(_arg0, _arg1, &_cerr)

	var _tlsInteractionResult TLSInteractionResult // out
	var _goerr error                               // out

	_tlsInteractionResult = (TLSInteractionResult)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsInteractionResult, _goerr
}
