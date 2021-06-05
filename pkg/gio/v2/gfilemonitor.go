// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
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
		{T: externglib.Type(C.g_file_monitor_get_type()), F: marshalFileMonitor},
	})
}

// FileMonitor monitors a file or directory for changes.
//
// To obtain a Monitor for a file or directory, use g_file_monitor(),
// g_file_monitor_file(), or g_file_monitor_directory().
//
// To get informed about changes to the file or directory you are monitoring,
// connect to the Monitor::changed signal. The signal will be emitted in the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in (though if the global default main
// context is blocked, this may cause notifications to be blocked even if the
// thread-default context is still running).
type FileMonitor interface {
	gextras.Objector

	// Cancel cancels a file monitor.
	Cancel() bool
	// EmitEvent emits the Monitor::changed signal if a change has taken place.
	// Should be called from file monitor implementations only.
	//
	// Implementations are responsible to call this method from the
	// [thread-default main context][g-main-context-push-thread-default] of the
	// thread that the monitor was created in.
	EmitEvent(child File, otherFile File, eventType FileMonitorEvent)
	// IsCancelled returns whether the monitor is canceled.
	IsCancelled() bool
	// SetRateLimit sets the rate limit to which the @monitor will report
	// consecutive change events to the same file.
	SetRateLimit(limitMsecs int)
}

// fileMonitor implements the FileMonitor interface.
type fileMonitor struct {
	gextras.Objector
}

var _ FileMonitor = (*fileMonitor)(nil)

// WrapFileMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileMonitor(obj *externglib.Object) FileMonitor {
	return FileMonitor{
		Objector: obj,
	}
}

func marshalFileMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileMonitor(obj), nil
}

// Cancel cancels a file monitor.
func (m fileMonitor) Cancel() bool {
	var arg0 *C.GFileMonitor

	arg0 = (*C.GFileMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_file_monitor_cancel(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// EmitEvent emits the Monitor::changed signal if a change has taken place.
// Should be called from file monitor implementations only.
//
// Implementations are responsible to call this method from the
// [thread-default main context][g-main-context-push-thread-default] of the
// thread that the monitor was created in.
func (m fileMonitor) EmitEvent(child File, otherFile File, eventType FileMonitorEvent) {
	var arg0 *C.GFileMonitor
	var arg1 *C.GFile
	var arg2 *C.GFile
	var arg3 C.GFileMonitorEvent

	arg0 = (*C.GFileMonitor)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GFile)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GFile)(unsafe.Pointer(otherFile.Native()))
	arg3 = (C.GFileMonitorEvent)(eventType)

	C.g_file_monitor_emit_event(arg0, child, otherFile, eventType)
}

// IsCancelled returns whether the monitor is canceled.
func (m fileMonitor) IsCancelled() bool {
	var arg0 *C.GFileMonitor

	arg0 = (*C.GFileMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.g_file_monitor_is_cancelled(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetRateLimit sets the rate limit to which the @monitor will report
// consecutive change events to the same file.
func (m fileMonitor) SetRateLimit(limitMsecs int) {
	var arg0 *C.GFileMonitor
	var arg1 C.gint

	arg0 = (*C.GFileMonitor)(unsafe.Pointer(m.Native()))
	arg1 = C.gint(limitMsecs)

	C.g_file_monitor_set_rate_limit(arg0, limitMsecs)
}

type FileMonitorPrivate struct {
	native C.GFileMonitorPrivate
}

// WrapFileMonitorPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileMonitorPrivate(ptr unsafe.Pointer) *FileMonitorPrivate {
	if ptr == nil {
		return nil
	}

	return (*FileMonitorPrivate)(ptr)
}

func marshalFileMonitorPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFileMonitorPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FileMonitorPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}