// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_iter_get_type()), F: marshalPixbufSimpleAnimIter},
	})
}

// PixbufNonAnim:
type PixbufNonAnim interface {
	PixbufAnimation
}

// pixbufNonAnim implements the PixbufNonAnim class.
type pixbufNonAnim struct {
	PixbufAnimation
}

// WrapPixbufNonAnim wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufNonAnim(obj *externglib.Object) PixbufNonAnim {
	return pixbufNonAnim{
		PixbufAnimation: WrapPixbufAnimation(obj),
	}
}

func marshalPixbufNonAnim(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufNonAnim(obj), nil
}

// NewPixbufNonAnim:
func NewPixbufNonAnim(pixbuf Pixbuf) PixbufNonAnim {
	var _arg1 *C.GdkPixbuf          // out
	var _cret *C.GdkPixbufAnimation // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gdk_pixbuf_non_anim_new(_arg1)

	var _pixbufNonAnim PixbufNonAnim // out

	_pixbufNonAnim = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PixbufNonAnim)

	return _pixbufNonAnim
}

// PixbufSimpleAnimIter:
type PixbufSimpleAnimIter interface {
	PixbufAnimationIter
}

// pixbufSimpleAnimIter implements the PixbufSimpleAnimIter class.
type pixbufSimpleAnimIter struct {
	PixbufAnimationIter
}

// WrapPixbufSimpleAnimIter wraps a GObject to the right type. It is
// primarily used internally.
func WrapPixbufSimpleAnimIter(obj *externglib.Object) PixbufSimpleAnimIter {
	return pixbufSimpleAnimIter{
		PixbufAnimationIter: WrapPixbufAnimationIter(obj),
	}
}

func marshalPixbufSimpleAnimIter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPixbufSimpleAnimIter(obj), nil
}
