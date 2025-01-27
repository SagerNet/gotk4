// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// IMContextInfo: bookkeeping information about a loadable input method.
//
// An instance of this type is always passed by reference.
type IMContextInfo struct {
	*imContextInfo
}

// imContextInfo is the struct that's finalized.
type imContextInfo struct {
	native *C.GtkIMContextInfo
}

// ContextID: unique identification string of the input method.
func (i *IMContextInfo) ContextID() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(i.native.context_id)))
	return v
}

// ContextName: human-readable name of the input method.
func (i *IMContextInfo) ContextName() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(i.native.context_name)))
	return v
}

// Domain: translation domain to be used with dgettext()
func (i *IMContextInfo) Domain() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(i.native.domain)))
	return v
}

// DomainDirname: name of locale directory for use with bindtextdomain()
func (i *IMContextInfo) DomainDirname() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(i.native.domain_dirname)))
	return v
}

// DefaultLocales: colon-separated list of locales where this input method
// should be the default. The asterisk “*” sets the default for all locales.
func (i *IMContextInfo) DefaultLocales() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(i.native.default_locales)))
	return v
}
