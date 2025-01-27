// Code generated by girgen. DO NOT EDIT.

package gio

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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

// IOSchedulerCancelAllJobs cancels all cancellable I/O jobs.
//
// A job is cancellable if a #GCancellable was passed into
// g_io_scheduler_push_job().
//
// Deprecated: You should never call this function, since you don't know how
// other libraries in your program might be making use of gioscheduler.
func IOSchedulerCancelAllJobs() {
	C.g_io_scheduler_cancel_all_jobs()
}
