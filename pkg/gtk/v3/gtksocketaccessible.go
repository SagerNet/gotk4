// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_socket_accessible_get_type()), F: marshalSocketAccessibler},
	})
}

// SocketAccessibler describes SocketAccessible's methods.
type SocketAccessibler interface {
	gextras.Objector

	Embed(path string)
}

type SocketAccessible struct {
	ContainerAccessible
}

var _ SocketAccessibler = (*SocketAccessible)(nil)

func wrapSocketAccessibler(obj *externglib.Object) SocketAccessibler {
	return &SocketAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					Object: atk.Object{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalSocketAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSocketAccessibler(obj), nil
}

func (socket *SocketAccessible) Embed(path string) {
	var _arg0 *C.GtkSocketAccessible // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkSocketAccessible)(unsafe.Pointer(socket.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_socket_accessible_embed(_arg0, _arg1)
}
