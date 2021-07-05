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
		{T: externglib.Type(C.g_tls_certificate_get_type()), F: marshalTLSCertificate},
	})
}

// TLSCertificate: certificate used for TLS authentication and encryption. This
// can represent either a certificate only (eg, the certificate received by a
// client from a server), or the combination of a certificate and a private key
// (which is needed when acting as a ServerConnection).
type TLSCertificate interface {
	gextras.Objector

	// Issuer gets the Certificate representing @cert's issuer, if known
	Issuer() TLSCertificate
	// IsSame: check if two Certificate objects represent the same certificate.
	// The raw DER byte data of the two certificates are checked for equality.
	// This has the effect that two certificates may compare equal even if their
	// Certificate:issuer, Certificate:private-key, or
	// Certificate:private-key-pem properties differ.
	IsSame(certTwo TLSCertificate) bool
	// Verify: this verifies @cert and returns a set of CertificateFlags
	// indicating any problems found with it. This can be used to verify a
	// certificate outside the context of making a connection, or to check a
	// certificate against a CA that is not part of the system CA database.
	//
	// If @identity is not nil, @cert's name(s) will be compared against it, and
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value if it does
	// not match. If @identity is nil, that bit will never be set in the return
	// value.
	//
	// If @trusted_ca is not nil, then @cert (or one of the certificates in its
	// chain) must be signed by it, or else G_TLS_CERTIFICATE_UNKNOWN_CA will be
	// set in the return value. If @trusted_ca is nil, that bit will never be
	// set in the return value.
	//
	// (All other CertificateFlags values will always be set or unset as
	// appropriate.)
	Verify(identity SocketConnectable, trustedCa TLSCertificate) TLSCertificateFlags
}

// tlsCertificate implements the TLSCertificate class.
type tlsCertificate struct {
	gextras.Objector
}

// WrapTLSCertificate wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSCertificate(obj *externglib.Object) TLSCertificate {
	return tlsCertificate{
		Objector: obj,
	}
}

func marshalTLSCertificate(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSCertificate(obj), nil
}

// NewTLSCertificateFromFile creates a Certificate from the PEM-encoded data in
// @file. The returned certificate will be the first certificate found in @file.
// As of GLib 2.44, if @file contains more certificates it will try to load a
// certificate chain. All certificates will be verified in the order found
// (top-level certificate should be the last one in the file) and the
// Certificate:issuer property of each certificate will be set accordingly if
// the verification succeeds. If any certificate in the chain cannot be
// verified, the first certificate in the file will still be returned.
//
// If @file cannot be read or parsed, the function will return nil and set
// @error. Otherwise, this behaves like g_tls_certificate_new_from_pem().
func NewTLSCertificateFromFile(file string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(file))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_tls_certificate_new_from_file(_arg1, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = WrapTLSCertificate(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromFiles creates a Certificate from the PEM-encoded data in
// @cert_file and @key_file. The returned certificate will be the first
// certificate found in @cert_file. As of GLib 2.44, if @cert_file contains more
// certificates it will try to load a certificate chain. All certificates will
// be verified in the order found (top-level certificate should be the last one
// in the file) and the Certificate:issuer property of each certificate will be
// set accordingly if the verification succeeds. If any certificate in the chain
// cannot be verified, the first certificate in the file will still be returned.
//
// If either file cannot be read or parsed, the function will return nil and set
// @error. Otherwise, this behaves like g_tls_certificate_new_from_pem().
func NewTLSCertificateFromFiles(certFile string, keyFile string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(certFile))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(keyFile))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_certificate_new_from_files(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = WrapTLSCertificate(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromPem creates a Certificate from the PEM-encoded data in
// @data. If @data includes both a certificate and a private key, then the
// returned certificate will include the private key data as well. (See the
// Certificate:private-key-pem property for information about supported
// formats.)
//
// The returned certificate will be the first certificate found in @data. As of
// GLib 2.44, if @data contains more certificates it will try to load a
// certificate chain. All certificates will be verified in the order found
// (top-level certificate should be the last one in the file) and the
// Certificate:issuer property of each certificate will be set accordingly if
// the verification succeeds. If any certificate in the chain cannot be
// verified, the first certificate in the file will still be returned.
func NewTLSCertificateFromPem(data string, length int) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 C.gssize           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(data))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.g_tls_certificate_new_from_pem(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = WrapTLSCertificate(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromPkcs11Uris creates a Certificate from a PKCS \#11 URI.
//
// An example @pkcs11_uri would be
// `pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01`
//
// Where the token’s layout is:
//
// “` Object 0: URL:
// pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01;object=private20key;type=private
// Type: Private key (RSA-2048) ID: 01
//
// Object 1: URL:
// pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01;object=Certificate20for20Authentication;type=cert
// Type: X.509 Certificate (RSA-2048) ID: 01 “`
//
// In this case the certificate and private key would both be detected and used
// as expected. @pkcs_uri may also just reference an X.509 certificate object
// and then optionally @private_key_pkcs11_uri allows using a private key
// exposed under a different URI.
//
// Note that the private key is not accessed until usage and may fail or require
// a PIN later.
func NewTLSCertificateFromPkcs11Uris(pkcs11Uri string, privateKeyPkcs11Uri string) (TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(C.CString(pkcs11Uri))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(privateKeyPkcs11Uri))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_certificate_new_from_pkcs11_uris(_arg1, _arg2, &_cerr)

	var _tlsCertificate TLSCertificate // out
	var _goerr error                   // out

	_tlsCertificate = WrapTLSCertificate(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _tlsCertificate, _goerr
}

func (c tlsCertificate) Issuer() TLSCertificate {
	var _arg0 *C.GTlsCertificate // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))

	_cret = C.g_tls_certificate_get_issuer(_arg0)

	var _tlsCertificate TLSCertificate // out

	_tlsCertificate = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TLSCertificate)

	return _tlsCertificate
}

func (c tlsCertificate) IsSame(certTwo TLSCertificate) bool {
	var _arg0 *C.GTlsCertificate // out
	var _arg1 *C.GTlsCertificate // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certTwo.Native()))

	_cret = C.g_tls_certificate_is_same(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c tlsCertificate) Verify(identity SocketConnectable, trustedCa TLSCertificate) TLSCertificateFlags {
	var _arg0 *C.GTlsCertificate     // out
	var _arg1 *C.GSocketConnectable  // out
	var _arg2 *C.GTlsCertificate     // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))
	_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(trustedCa.Native()))

	_cret = C.g_tls_certificate_verify(_arg0, _arg1, _arg2)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}
