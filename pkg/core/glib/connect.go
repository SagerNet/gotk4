package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/closure"
	"github.com/diamondburned/gotk4/pkg/core/intern"
)

// SignalHandle is the ID of a signal handler.
type SignalHandle uint

// Connect is a wrapper around g_signal_connect_closure(). f must be a function
// with at least one parameter matching the type it is connected to.
//
// It is optional to list the rest of the required types from GTK, as values
// that don't fit into the function parameter will simply be ignored; however,
// extraneous types will trigger a runtime panic. Arguments for f must be a
// matching Go equivalent type for the C callback, or an interface type which
// the value may be packed in. If the type is not suitable, a runtime panic will
// occur when the signal is emitted.
func (v *Object) Connect(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(false, detailedSignal, f)
}

// ConnectAfter is a wrapper around g_signal_connect_closure(). The difference
// between Connect and ConnectAfter is that the latter will be invoked after the
// default handler, not before. For more information, refer to Connect.
func (v *Object) ConnectAfter(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(true, detailedSignal, f)
}

func (v *Object) connectClosure(after bool, detailedSignal string, f interface{}) SignalHandle {
	fs := closure.NewFuncStack(f, 2)

	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	gclosure := closureNew(v, fs)
	c := C.g_signal_connect_closure(C.gpointer(v.Native()), (*C.gchar)(cstr), gclosure, gbool(after))

	runtime.KeepAlive(v)

	return SignalHandle(c)
}

// closureNew creates a new GClosure that's bound to the current object and adds
// its callback function to the internal registry. It's exported for visibility
// to other gotk3 packages and should not be used in a regular application.
func closureNew(v *Object, f interface{}) *C.GClosure {
	fs, ok := f.(*closure.FuncStack)
	if !ok {
		fs = closure.NewFuncStack(f, 2)
	}

	gclosure := C.g_closure_new_simple(C.sizeof_GClosure, nil)
	v.box.Closures.Register(unsafe.Pointer(gclosure), fs)

	C.g_closure_set_meta_marshal(gclosure, C.gpointer(v.Native()), (*[0]byte)(C.goMarshal))
	C.g_closure_add_finalize_notifier(gclosure, C.gpointer(v.Native()), (*[0]byte)(C.removeClosure))

	return gclosure
}

//export removeClosure
func removeClosure(obj *C.GObject, gclosure *C.GClosure) {
	intern.RemoveClosure(unsafe.Pointer(obj), unsafe.Pointer(gclosure))
}
