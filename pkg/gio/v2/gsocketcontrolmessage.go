// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_socket_control_message_get_type()), F: marshalSocketControlMessager},
	})
}

// SocketControlMessageOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketControlMessageOverrider interface {
	// Level returns the "level" (i.e. the originating protocol) of the control
	// message. This is often SOL_SOCKET.
	Level() int
	// Size returns the space required for the control message, not including
	// headers or alignment.
	Size() uint
	Type() int
	// Serialize converts the data in the message to bytes placed in the
	// message.
	//
	// data is guaranteed to have enough space to fit the size returned by
	// g_socket_control_message_get_size() on this object.
	Serialize(data cgo.Handle)
}

// SocketControlMessage is a special-purpose utility message that can be sent to
// or received from a #GSocket. These types of messages are often called
// "ancillary data".
//
// The message can represent some sort of special instruction to or information
// from the socket or can represent a special kind of transfer to the peer (for
// example, sending a file descriptor over a UNIX socket).
//
// These messages are sent with g_socket_send_message() and received with
// g_socket_receive_message().
//
// To extend the set of control message that can be sent, subclass this class
// and override the get_size, get_level, get_type and serialize methods.
//
// To extend the set of control messages that can be received, subclass this
// class and implement the deserialize method. Also, make sure your class is
// registered with the GType typesystem before calling
// g_socket_receive_message() to read such a message.
type SocketControlMessage struct {
	*externglib.Object
}

// SocketControlMessager describes SocketControlMessage's abstract methods.
type SocketControlMessager interface {
	externglib.Objector

	// Level returns the "level" (i.e.
	Level() int
	// MsgType returns the protocol specific type of the control message.
	MsgType() int
	// Size returns the space required for the control message, not including
	// headers or alignment.
	Size() uint
	// Serialize converts the data in the message to bytes placed in the
	// message.
	Serialize(data cgo.Handle)
}

var _ SocketControlMessager = (*SocketControlMessage)(nil)

func WrapSocketControlMessage(obj *externglib.Object) *SocketControlMessage {
	return &SocketControlMessage{
		Object: obj,
	}
}

func marshalSocketControlMessager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketControlMessage(obj), nil
}

// Level returns the "level" (i.e. the originating protocol) of the control
// message. This is often SOL_SOCKET.
func (message *SocketControlMessage) Level() int {
	var _arg0 *C.GSocketControlMessage // out
	var _cret C.int                    // in

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(message.Native()))

	_cret = C.g_socket_control_message_get_level(_arg0)
	runtime.KeepAlive(message)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MsgType returns the protocol specific type of the control message. For
// instance, for UNIX fd passing this would be SCM_RIGHTS.
func (message *SocketControlMessage) MsgType() int {
	var _arg0 *C.GSocketControlMessage // out
	var _cret C.int                    // in

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(message.Native()))

	_cret = C.g_socket_control_message_get_msg_type(_arg0)
	runtime.KeepAlive(message)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Size returns the space required for the control message, not including
// headers or alignment.
func (message *SocketControlMessage) Size() uint {
	var _arg0 *C.GSocketControlMessage // out
	var _cret C.gsize                  // in

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(message.Native()))

	_cret = C.g_socket_control_message_get_size(_arg0)
	runtime.KeepAlive(message)

	var _gsize uint // out

	_gsize = uint(_cret)

	return _gsize
}

// Serialize converts the data in the message to bytes placed in the message.
//
// data is guaranteed to have enough space to fit the size returned by
// g_socket_control_message_get_size() on this object.
func (message *SocketControlMessage) Serialize(data cgo.Handle) {
	var _arg0 *C.GSocketControlMessage // out
	var _arg1 C.gpointer               // out

	_arg0 = (*C.GSocketControlMessage)(unsafe.Pointer(message.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(data))

	C.g_socket_control_message_serialize(_arg0, _arg1)
	runtime.KeepAlive(message)
	runtime.KeepAlive(data)
}

// SocketControlMessageDeserialize tries to deserialize a socket control message
// of a given level and type. This will ask all known (to GType) subclasses of
// ControlMessage if they can understand this kind of message and if so
// deserialize it into a ControlMessage.
//
// If there is no implementation for this kind of control message, NULL will be
// returned.
func SocketControlMessageDeserialize(level int, typ int, data []byte) SocketControlMessager {
	var _arg1 C.int      // out
	var _arg2 C.int      // out
	var _arg4 C.gpointer // out
	var _arg3 C.gsize
	var _cret *C.GSocketControlMessage // in

	_arg1 = C.int(level)
	_arg2 = C.int(typ)
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg4 = (C.gpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.g_socket_control_message_deserialize(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(level)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(data)

	var _socketControlMessage SocketControlMessager // out

	_socketControlMessage = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(SocketControlMessager)

	return _socketControlMessage
}
