// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
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
		{T: externglib.Type(C.g_network_monitor_get_type()), F: marshalNetworkMonitor},
	})
}

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
type NetworkMonitor interface {
	Initable

	// CanReach checks if the network is metered. See Monitor:network-metered
	// for more details.
	CanReach(connectable SocketConnectable, cancellable Cancellable) error
	// CanReachFinish checks if the network is metered. See
	// Monitor:network-metered for more details.
	CanReachFinish(result AsyncResult) error
	// Connectivity checks if the network is metered. See
	// Monitor:network-metered for more details.
	Connectivity() NetworkConnectivity
	// NetworkAvailable checks if the network is metered. See
	// Monitor:network-metered for more details.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered. See
	// Monitor:network-metered for more details.
	NetworkMetered() bool
}

// networkMonitor implements the NetworkMonitor interface.
type networkMonitor struct {
	Initable
}

var _ NetworkMonitor = (*networkMonitor)(nil)

// WrapNetworkMonitor wraps a GObject to a type that implements
// interface NetworkMonitor. It is primarily used internally.
func WrapNetworkMonitor(obj *externglib.Object) NetworkMonitor {
	return networkMonitor{
		Initable: WrapInitable(obj),
	}
}

func marshalNetworkMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkMonitor(obj), nil
}

func (m networkMonitor) CanReach(connectable SocketConnectable, cancellable Cancellable) error {
	var _arg0 *C.GNetworkMonitor    // out
	var _arg1 *C.GSocketConnectable // out
	var _arg2 *C.GCancellable       // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_network_monitor_can_reach(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m networkMonitor) CanReachFinish(result AsyncResult) error {
	var _arg0 *C.GNetworkMonitor // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr *C.GError          // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_network_monitor_can_reach_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (m networkMonitor) Connectivity() NetworkConnectivity {
	var _arg0 *C.GNetworkMonitor     // out
	var _cret C.GNetworkConnectivity // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	_cret = C.g_network_monitor_get_connectivity(_arg0)

	var _networkConnectivity NetworkConnectivity // out

	_networkConnectivity = NetworkConnectivity(_cret)

	return _networkConnectivity
}

func (m networkMonitor) NetworkAvailable() bool {
	var _arg0 *C.GNetworkMonitor // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	_cret = C.g_network_monitor_get_network_available(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (m networkMonitor) NetworkMetered() bool {
	var _arg0 *C.GNetworkMonitor // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	_cret = C.g_network_monitor_get_network_metered(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
