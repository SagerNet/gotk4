// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_accessible_get_type()), F: marshalSocketAccessibler},
	})
}

type SocketAccessible struct {
	ContainerAccessible
}

func WrapSocketAccessible(obj *externglib.Object) *SocketAccessible {
	return &SocketAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
	}
}

func marshalSocketAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketAccessible(obj), nil
}

func (socket *SocketAccessible) Embed(path string) {
	var _arg0 *C.GtkSocketAccessible // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(socket.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_socket_accessible_embed(_arg0, _arg1)
	runtime.KeepAlive(socket)
	runtime.KeepAlive(path)
}
