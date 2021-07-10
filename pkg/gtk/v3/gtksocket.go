// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_get_type()), F: marshalSocketter},
	})
}

// SocketterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SocketterOverrider interface {
	PlugAdded()
	PlugRemoved() bool
}

// Socketter describes Socket's methods.
type Socketter interface {
	gextras.Objector

	PlugWindow() *gdk.Window
}

// Socket: together with Plug, Socket provides the ability to embed widgets from
// one process into another process in a fashion that is transparent to the
// user. One process creates a Socket widget and passes that widget’s window ID
// to the other process, which then creates a Plug with that window ID. Any
// widgets contained in the Plug then will appear inside the first application’s
// window.
//
// The socket’s window ID is obtained by using gtk_socket_get_id(). Before using
// this function, the socket must have been realized, and for hence, have been
// added to its parent.
//
// Obtaining the window ID of a socket.
//
//    GtkWidget *socket = gtk_socket_new ();
//    gtk_widget_show (socket);
//    gtk_container_add (GTK_CONTAINER (parent), socket);
//
//    // The following call is only necessary if one of
//    // the ancestors of the socket is not yet visible.
//    gtk_widget_realize (socket);
//    g_print ("The ID of the sockets window is %#x\n",
//             gtk_socket_get_id (socket));
//
// Note that if you pass the window ID of the socket to another process that
// will create a plug in the socket, you must make sure that the socket widget
// is not destroyed until that plug is created. Violating this rule will cause
// unpredictable consequences, the most likely consequence being that the plug
// will appear as a separate toplevel window. You can check if the plug has been
// created by using gtk_socket_get_plug_window(). If it returns a non-nil value,
// then the plug has been successfully created inside of the socket.
//
// When GTK+ is notified that the embedded window has been destroyed, then it
// will destroy the socket as well. You should always, therefore, be prepared
// for your sockets to be destroyed at any time when the main event loop is
// running. To prevent this from happening, you can connect to the
// Socket::plug-removed signal.
//
// The communication between a Socket and a Plug follows the XEmbed Protocol
// (http://www.freedesktop.org/Standards/xembed-spec). This protocol has also
// been implemented in other toolkits, e.g. Qt, allowing the same level of
// integration when embedding a Qt widget in GTK or vice versa.
//
// The Plug and Socket widgets are only available when GTK+ is compiled for the
// X11 platform and GDK_WINDOWING_X11 is defined. They can only be used on a
// X11Display. To use Plug and Socket, you need to include the `gtk/gtkx.h`
// header.
type Socket struct {
	*externglib.Object
	Container
	Buildable
}

var _ Socketter = (*Socket)(nil)

func wrapSocketter(obj *externglib.Object) Socketter {
	return &Socket{
		Object: obj,
		Container: Container{
			Object: obj,
			Widget: Widget{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalSocketter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketter(obj), nil
}

// NewSocket: create a new empty Socket.
func NewSocket() *Socket {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_socket_new()

	var _socket *Socket // out

	_socket = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Socket)

	return _socket
}

// PlugWindow retrieves the window of the plug. Use this to check if the plug
// has been created inside of the socket.
func (socket_ *Socket) PlugWindow() *gdk.Window {
	var _arg0 *C.GtkSocket // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkSocket)(unsafe.Pointer(socket_.Native()))

	_cret = C.gtk_socket_get_plug_window(_arg0)

	var _window *gdk.Window // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.Window)

	return _window
}
