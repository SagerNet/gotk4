// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_get_type()), F: marshalPixbufSimpleAnim},
	})
}

// PixbufSimpleAnim: an opaque struct representing a simple animation.
type PixbufSimpleAnim interface {
	PixbufAnimation

	// AddFrame adds a new frame to @animation. The @pixbuf must have the
	// dimensions specified when the animation was constructed.
	AddFrame(pixbuf Pixbuf)
	// Loop gets whether @animation should loop indefinitely when it reaches the
	// end.
	Loop() bool
	// SetLoop sets whether @animation should loop indefinitely when it reaches
	// the end.
	SetLoop(loop bool)
}

// pixbufSimpleAnim implements the PixbufSimpleAnim interface.
type pixbufSimpleAnim struct {
	PixbufAnimation
}

var _ PixbufSimpleAnim = (*pixbufSimpleAnim)(nil)

// WrapPixbufSimpleAnim wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufSimpleAnim(obj *externglib.Object) PixbufSimpleAnim {
	return PixbufSimpleAnim{
		PixbufAnimation: WrapPixbufAnimation(obj),
	}
}

func marshalPixbufSimpleAnim(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufSimpleAnim(obj), nil
}

// NewPixbufSimpleAnim constructs a class PixbufSimpleAnim.
func NewPixbufSimpleAnim(width int, height int, rate float32) PixbufSimpleAnim {
	var arg1 C.gint
	var arg2 C.gint
	var arg3 C.gfloat

	arg1 = C.gint(width)
	arg2 = C.gint(height)
	arg3 = C.gfloat(rate)

	var cret C.GdkPixbufSimpleAnim
	var goret1 PixbufSimpleAnim

	cret = C.gdk_pixbuf_simple_anim_new(width, height, rate)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(PixbufSimpleAnim)

	return goret1
}

// AddFrame adds a new frame to @animation. The @pixbuf must have the
// dimensions specified when the animation was constructed.
func (a pixbufSimpleAnim) AddFrame(pixbuf Pixbuf) {
	var arg0 *C.GdkPixbufSimpleAnim
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gdk_pixbuf_simple_anim_add_frame(arg0, pixbuf)
}

// Loop gets whether @animation should loop indefinitely when it reaches the
// end.
func (a pixbufSimpleAnim) Loop() bool {
	var arg0 *C.GdkPixbufSimpleAnim

	arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gdk_pixbuf_simple_anim_get_loop(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetLoop sets whether @animation should loop indefinitely when it reaches
// the end.
func (a pixbufSimpleAnim) SetLoop(loop bool) {
	var arg0 *C.GdkPixbufSimpleAnim
	var arg1 C.gboolean

	arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))
	if loop {
		arg1 = C.gboolean(1)
	}

	C.gdk_pixbuf_simple_anim_set_loop(arg0, loop)
}
