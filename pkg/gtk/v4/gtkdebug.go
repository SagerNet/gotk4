// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// GetDebugFlags returns the GTK debug flags that are currently active.
//
// This function is intended for GTK modules that want to adjust their debug
// output based on GTK debug flags.
func GetDebugFlags() DebugFlags {
	var _cret C.GtkDebugFlags // in

	_cret = C.gtk_get_debug_flags()

	var _debugFlags DebugFlags // out

	_debugFlags = DebugFlags(_cret)

	return _debugFlags
}

// SetDebugFlags sets the GTK debug flags.
func SetDebugFlags(flags DebugFlags) {
	var _arg1 C.GtkDebugFlags // out

	_arg1 = C.GtkDebugFlags(flags)

	C.gtk_set_debug_flags(_arg1)
	runtime.KeepAlive(flags)
}
