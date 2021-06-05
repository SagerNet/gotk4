// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_tls_database_get_type()), F: marshalTLSDatabase},
	})
}

// TLSDatabase is used to look up certificates and other information from a
// certificate or key store. It is an abstract base class which TLS library
// specific subtypes override.
//
// A Database may be accessed from multiple threads by the TLS backend. All
// implementations are required to be fully thread-safe.
//
// Most common client applications will not directly interact with Database. It
// is used internally by Connection.
type TLSDatabase interface {
	gextras.Objector

	// CreateCertificateHandle: create a handle string for the certificate. The
	// database will only be able to create a handle for certificates that
	// originate from the database. In cases where the database cannot create a
	// handle for a certificate, nil will be returned.
	//
	// This handle should be stable across various instances of the application,
	// and between applications. If a certificate is modified in the database,
	// then it is not guaranteed that this handle will continue to point to it.
	CreateCertificateHandle(certificate TLSCertificate) string
	// LookupCertificateForHandle: look up a certificate by its handle.
	//
	// The handle should have been created by calling
	// g_tls_database_create_certificate_handle() on a Database object of the
	// same TLS backend. The handle is designed to remain valid across
	// instantiations of the database.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_for_handle_async() to perform the
	// lookup operation asynchronously.
	LookupCertificateForHandle(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (tlsCertificate TLSCertificate, err error)
	// LookupCertificateForHandleAsync: asynchronously look up a certificate by
	// its handle in the database. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	LookupCertificateForHandleAsync(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
	// certificate by its handle. See
	// g_tls_database_lookup_certificate_for_handle() for more information.
	//
	// If the handle is no longer valid, or does not point to a certificate in
	// this database, then nil will be returned.
	LookupCertificateForHandleFinish(result AsyncResult) (tlsCertificate TLSCertificate, err error)
	// LookupCertificateIssuer: look up the issuer of @certificate in the
	// database.
	//
	// The Certificate:issuer property of @certificate is not modified, and the
	// two certificates are not hooked into a chain.
	//
	// This function can block, use
	// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
	// operation asynchronously.
	LookupCertificateIssuer(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (tlsCertificate TLSCertificate, err error)
	// LookupCertificateIssuerAsync: asynchronously look up the issuer of
	// @certificate in the database. See
	// g_tls_database_lookup_certificate_issuer() for more information.
	LookupCertificateIssuerAsync(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
	// operation. See g_tls_database_lookup_certificate_issuer() for more
	// information.
	LookupCertificateIssuerFinish(result AsyncResult) (tlsCertificate TLSCertificate, err error)
	// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
	// the database.
	//
	// This function can block, use
	// g_tls_database_lookup_certificates_issued_by_async() to perform the
	// lookup operation asynchronously.
	LookupCertificatesIssuedBy(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (list *glib.List, err error)
	// LookupCertificatesIssuedByAsync: asynchronously look up certificates
	// issued by this issuer in the database. See
	// g_tls_database_lookup_certificates_issued_by() for more information.
	//
	// The database may choose to hold a reference to the issuer byte array for
	// the duration of of this asynchronous operation. The byte array should not
	// be modified during this time.
	LookupCertificatesIssuedByAsync(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
	// certificates. See g_tls_database_lookup_certificates_issued_by() for more
	// information.
	LookupCertificatesIssuedByFinish(result AsyncResult) (list *glib.List, err error)
	// VerifyChain determines the validity of a certificate chain after looking
	// up and adding any missing certificates to the chain.
	//
	// @chain is a chain of Certificate objects each pointing to the next
	// certificate in the chain by its Certificate:issuer property. The chain
	// may initially consist of one or more certificates. After the verification
	// process is complete, @chain may be modified by adding missing
	// certificates, or removing extra certificates. If a certificate anchor was
	// found, then it is added to the @chain.
	//
	// @purpose describes the purpose (or usage) for which the certificate is
	// being used. Typically @purpose will be set to
	// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
	// is being used to authenticate a server (and we are acting as the client).
	//
	// The @identity is used to ensure the server certificate is valid for the
	// expected peer identity. If the identity does not match the certificate,
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
	// @identity is nil, that bit will never be set in the return value. The
	// peer identity may also be used to check for pinned certificates (trust
	// exceptions) in the database. These may override the normal verification
	// process on a host-by-host basis.
	//
	// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be
	// used.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	//
	// This function can block, use g_tls_database_verify_chain_async() to
	// perform the verification operation asynchronously.
	VerifyChain(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (tlsCertificateFlags TLSCertificateFlags, err error)
	// VerifyChainAsync: asynchronously determines the validity of a certificate
	// chain after looking up and adding any missing certificates to the chain.
	// See g_tls_database_verify_chain() for more information.
	VerifyChainAsync(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable, callback AsyncReadyCallback)
	// VerifyChainFinish: finish an asynchronous verify chain operation. See
	// g_tls_database_verify_chain() for more information.
	//
	// If @chain is found to be valid, then the return value will be 0. If
	// @chain is found to be invalid, then the return value will indicate the
	// problems found. If the function is unable to determine whether @chain is
	// valid or not (eg, because @cancellable is triggered before it completes)
	// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
	// will be set accordingly. @error is not set when @chain is successfully
	// analyzed but found to be invalid.
	VerifyChainFinish(result AsyncResult) (tlsCertificateFlags TLSCertificateFlags, err error)
}

// tlsDatabase implements the TLSDatabase interface.
type tlsDatabase struct {
	gextras.Objector
}

var _ TLSDatabase = (*tlsDatabase)(nil)

// WrapTLSDatabase wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSDatabase(obj *externglib.Object) TLSDatabase {
	return TLSDatabase{
		Objector: obj,
	}
}

func marshalTLSDatabase(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSDatabase(obj), nil
}

// CreateCertificateHandle: create a handle string for the certificate. The
// database will only be able to create a handle for certificates that
// originate from the database. In cases where the database cannot create a
// handle for a certificate, nil will be returned.
//
// This handle should be stable across various instances of the application,
// and between applications. If a certificate is modified in the database,
// then it is not guaranteed that this handle will continue to point to it.
func (s tlsDatabase) CreateCertificateHandle(certificate TLSCertificate) string {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GTlsCertificate

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.g_tls_database_create_certificate_handle(arg0, certificate)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}

// LookupCertificateForHandle: look up a certificate by its handle.
//
// The handle should have been created by calling
// g_tls_database_create_certificate_handle() on a Database object of the
// same TLS backend. The handle is designed to remain valid across
// instantiations of the database.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then nil will be returned.
//
// This function can block, use
// g_tls_database_lookup_certificate_for_handle_async() to perform the
// lookup operation asynchronously.
func (s tlsDatabase) LookupCertificateForHandle(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (tlsCertificate TLSCertificate, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.gchar
	var arg2 *C.GTlsInteraction
	var arg3 C.GTlsDatabaseLookupFlags
	var arg4 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(handle))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GTlsCertificate
	var goret1 TLSCertificate
	var goerr error

	cret = C.g_tls_database_lookup_certificate_for_handle(arg0, handle, interaction, flags, cancellable, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TLSCertificate)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// LookupCertificateForHandleAsync: asynchronously look up a certificate by
// its handle in the database. See
// g_tls_database_lookup_certificate_for_handle() for more information.
func (s tlsDatabase) LookupCertificateForHandleAsync(handle string, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GTlsDatabase

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificate_for_handle_async(arg0, handle, interaction, flags, cancellable, callback, userData)
}

// LookupCertificateForHandleFinish: finish an asynchronous lookup of a
// certificate by its handle. See
// g_tls_database_lookup_certificate_for_handle() for more information.
//
// If the handle is no longer valid, or does not point to a certificate in
// this database, then nil will be returned.
func (s tlsDatabase) LookupCertificateForHandleFinish(result AsyncResult) (tlsCertificate TLSCertificate, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret *C.GTlsCertificate
	var goret1 TLSCertificate
	var goerr error

	cret = C.g_tls_database_lookup_certificate_for_handle_finish(arg0, result, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TLSCertificate)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// LookupCertificateIssuer: look up the issuer of @certificate in the
// database.
//
// The Certificate:issuer property of @certificate is not modified, and the
// two certificates are not hooked into a chain.
//
// This function can block, use
// g_tls_database_lookup_certificate_issuer_async() to perform the lookup
// operation asynchronously.
func (s tlsDatabase) LookupCertificateIssuer(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (tlsCertificate TLSCertificate, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GTlsCertificate
	var arg2 *C.GTlsInteraction
	var arg3 C.GTlsDatabaseLookupFlags
	var arg4 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GTlsCertificate
	var goret1 TLSCertificate
	var goerr error

	cret = C.g_tls_database_lookup_certificate_issuer(arg0, certificate, interaction, flags, cancellable, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TLSCertificate)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// LookupCertificateIssuerAsync: asynchronously look up the issuer of
// @certificate in the database. See
// g_tls_database_lookup_certificate_issuer() for more information.
func (s tlsDatabase) LookupCertificateIssuerAsync(certificate TLSCertificate, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GTlsDatabase

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificate_issuer_async(arg0, certificate, interaction, flags, cancellable, callback, userData)
}

// LookupCertificateIssuerFinish: finish an asynchronous lookup issuer
// operation. See g_tls_database_lookup_certificate_issuer() for more
// information.
func (s tlsDatabase) LookupCertificateIssuerFinish(result AsyncResult) (tlsCertificate TLSCertificate, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret *C.GTlsCertificate
	var goret1 TLSCertificate
	var goerr error

	cret = C.g_tls_database_lookup_certificate_issuer_finish(arg0, result, &errout)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TLSCertificate)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// LookupCertificatesIssuedBy: look up certificates issued by this issuer in
// the database.
//
// This function can block, use
// g_tls_database_lookup_certificates_issued_by_async() to perform the
// lookup operation asynchronously.
func (s tlsDatabase) LookupCertificatesIssuedBy(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable) (list *glib.List, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GByteArray
	var arg2 *C.GTlsInteraction
	var arg3 C.GTlsDatabaseLookupFlags
	var arg4 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = C.malloc(len(issuerRawDn) * (C.sizeof_guint8 + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []C.guint8
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(issuerRawDn)))

		for i := range issuerRawDn {
			out[i] = C.guint8(issuerRawDn[i])
		}
	}
	arg2 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	arg3 = (C.GTlsDatabaseLookupFlags)(flags)
	arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GList
	var goret1 *glib.List
	var goerr error

	cret = C.g_tls_database_lookup_certificates_issued_by(arg0, issuerRawDn, interaction, flags, cancellable, &errout)

	goret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// LookupCertificatesIssuedByAsync: asynchronously look up certificates
// issued by this issuer in the database. See
// g_tls_database_lookup_certificates_issued_by() for more information.
//
// The database may choose to hold a reference to the issuer byte array for
// the duration of of this asynchronous operation. The byte array should not
// be modified during this time.
func (s tlsDatabase) LookupCertificatesIssuedByAsync(issuerRawDn []byte, interaction TLSInteraction, flags TLSDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GTlsDatabase

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_lookup_certificates_issued_by_async(arg0, issuerRawDn, interaction, flags, cancellable, callback, userData)
}

// LookupCertificatesIssuedByFinish: finish an asynchronous lookup of
// certificates. See g_tls_database_lookup_certificates_issued_by() for more
// information.
func (s tlsDatabase) LookupCertificatesIssuedByFinish(result AsyncResult) (list *glib.List, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret *C.GList
	var goret1 *glib.List
	var goerr error

	cret = C.g_tls_database_lookup_certificates_issued_by_finish(arg0, result, &errout)

	goret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// VerifyChain determines the validity of a certificate chain after looking
// up and adding any missing certificates to the chain.
//
// @chain is a chain of Certificate objects each pointing to the next
// certificate in the chain by its Certificate:issuer property. The chain
// may initially consist of one or more certificates. After the verification
// process is complete, @chain may be modified by adding missing
// certificates, or removing extra certificates. If a certificate anchor was
// found, then it is added to the @chain.
//
// @purpose describes the purpose (or usage) for which the certificate is
// being used. Typically @purpose will be set to
// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER which means that the certificate
// is being used to authenticate a server (and we are acting as the client).
//
// The @identity is used to ensure the server certificate is valid for the
// expected peer identity. If the identity does not match the certificate,
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value. If
// @identity is nil, that bit will never be set in the return value. The
// peer identity may also be used to check for pinned certificates (trust
// exceptions) in the database. These may override the normal verification
// process on a host-by-host basis.
//
// Currently there are no @flags, and G_TLS_DATABASE_VERIFY_NONE should be
// used.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate the
// problems found. If the function is unable to determine whether @chain is
// valid or not (eg, because @cancellable is triggered before it completes)
// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
// will be set accordingly. @error is not set when @chain is successfully
// analyzed but found to be invalid.
//
// This function can block, use g_tls_database_verify_chain_async() to
// perform the verification operation asynchronously.
func (s tlsDatabase) VerifyChain(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable) (tlsCertificateFlags TLSCertificateFlags, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GTlsCertificate
	var arg2 *C.gchar
	var arg3 *C.GSocketConnectable
	var arg4 *C.GTlsInteraction
	var arg5 C.GTlsDatabaseVerifyFlags
	var arg6 *C.GCancellable
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(chain.Native()))
	arg2 = (*C.gchar)(C.CString(purpose))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	arg4 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))
	arg5 = (C.GTlsDatabaseVerifyFlags)(flags)
	arg6 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.GTlsCertificateFlags
	var goret1 TLSCertificateFlags
	var goerr error

	cret = C.g_tls_database_verify_chain(arg0, chain, purpose, identity, interaction, flags, cancellable, &errout)

	goret1 = TLSCertificateFlags(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// VerifyChainAsync: asynchronously determines the validity of a certificate
// chain after looking up and adding any missing certificates to the chain.
// See g_tls_database_verify_chain() for more information.
func (s tlsDatabase) VerifyChainAsync(chain TLSCertificate, purpose string, identity SocketConnectable, interaction TLSInteraction, flags TLSDatabaseVerifyFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	var arg0 *C.GTlsDatabase

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))

	C.g_tls_database_verify_chain_async(arg0, chain, purpose, identity, interaction, flags, cancellable, callback, userData)
}

// VerifyChainFinish: finish an asynchronous verify chain operation. See
// g_tls_database_verify_chain() for more information.
//
// If @chain is found to be valid, then the return value will be 0. If
// @chain is found to be invalid, then the return value will indicate the
// problems found. If the function is unable to determine whether @chain is
// valid or not (eg, because @cancellable is triggered before it completes)
// then the return value will be G_TLS_CERTIFICATE_GENERIC_ERROR and @error
// will be set accordingly. @error is not set when @chain is successfully
// analyzed but found to be invalid.
func (s tlsDatabase) VerifyChainFinish(result AsyncResult) (tlsCertificateFlags TLSCertificateFlags, err error) {
	var arg0 *C.GTlsDatabase
	var arg1 *C.GAsyncResult
	var errout *C.GError

	arg0 = (*C.GTlsDatabase)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.GTlsCertificateFlags
	var goret1 TLSCertificateFlags
	var goerr error

	cret = C.g_tls_database_verify_chain_finish(arg0, result, &errout)

	goret1 = TLSCertificateFlags(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

type TLSDatabasePrivate struct {
	native C.GTlsDatabasePrivate
}

// WrapTLSDatabasePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTLSDatabasePrivate(ptr unsafe.Pointer) *TLSDatabasePrivate {
	if ptr == nil {
		return nil
	}

	return (*TLSDatabasePrivate)(ptr)
}

func marshalTLSDatabasePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTLSDatabasePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TLSDatabasePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
