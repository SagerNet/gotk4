// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// IMContextInfo: bookkeeping information about a loadable input method.
type IMContextInfo C.GtkIMContextInfo

// WrapIMContextInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIMContextInfo(ptr unsafe.Pointer) *IMContextInfo {
	return (*IMContextInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IMContextInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}