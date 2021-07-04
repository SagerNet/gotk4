// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_plug_get_type()), F: marshalPlug},
	})
}

// Plug: together with Socket, Plug provides the ability to embed widgets from
// one process into another process in a fashion that is transparent to the
// user. One process creates a Socket widget and passes the ID of that widget’s
// window to the other process, which then creates a Plug with that window ID.
// Any widgets contained in the Plug then will appear inside the first
// application’s window.
//
// The communication between a Socket and a Plug follows the XEmbed Protocol
// (http://www.freedesktop.org/Standards/xembed-spec). This protocol has also
// been implemented in other toolkits, e.g. Qt, allowing the same level of
// integration when embedding a Qt widget in GTK+ or vice versa.
//
// The Plug and Socket widgets are only available when GTK+ is compiled for the
// X11 platform and GDK_WINDOWING_X11 is defined. They can only be used on a
// X11Display. To use Plug and Socket, you need to include the `gtk/gtkx.h`
// header.
type Plug interface {
	Window

	// Embedded:
	Embedded() bool
	// SocketWindow:
	SocketWindow() gdk.Window
}

// plug implements the Plug class.
type plug struct {
	Window
}

// WrapPlug wraps a GObject to the right type. It is
// primarily used internally.
func WrapPlug(obj *externglib.Object) Plug {
	return plug{
		Window: WrapWindow(obj),
	}
}

func marshalPlug(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlug(obj), nil
}

func (p plug) Embedded() bool {
	var _arg0 *C.GtkPlug // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_plug_get_embedded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p plug) SocketWindow() gdk.Window {
	var _arg0 *C.GtkPlug   // out
	var _cret *C.GdkWindow // in

	_arg0 = (*C.GtkPlug)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_plug_get_socket_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.Window)

	return _window
}

func (b plug) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b plug) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b plug) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b plug) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b plug) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b plug) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b plug) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
