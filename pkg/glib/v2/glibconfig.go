// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// Pid: a type which is used to hold a process identification.
//
// On UNIX, processes are identified by a process id (an integer), while Windows
// uses process handles (which are pointers).
//
// GPid is used in GLib only for descendant processes spawned with the g_spawn
// functions.
type Pid int
