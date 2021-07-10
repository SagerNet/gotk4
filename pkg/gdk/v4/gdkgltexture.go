// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_texture_get_type()), F: marshalGLTexturer},
	})
}

// GLTexturer describes GLTexture's methods.
type GLTexturer interface {
	gextras.Objector

	Release()
}

// GLTexture: gdkTexture representing a GL texture object.
type GLTexture struct {
	*externglib.Object
	Texture
	Paintable
}

var _ GLTexturer = (*GLTexture)(nil)

func wrapGLTexturer(obj *externglib.Object) GLTexturer {
	return &GLTexture{
		Object: obj,
		Texture: Texture{
			Object: obj,
			Paintable: Paintable{
				Object: obj,
			},
		},
		Paintable: Paintable{
			Object: obj,
		},
	}
}

func marshalGLTexturer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGLTexturer(obj), nil
}

// Release releases the GL resources held by a `GdkGLTexture`.
//
// The texture contents are still available via the
// [method@Gdk.Texture.download] function, after this function has been called.
func (self *GLTexture) Release() {
	var _arg0 *C.GdkGLTexture // out

	_arg0 = (*C.GdkGLTexture)(unsafe.Pointer(self.Native()))

	C.gdk_gl_texture_release(_arg0)
}
