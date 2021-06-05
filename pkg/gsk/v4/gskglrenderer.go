// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_gl_renderer_get_type()), F: marshalGLRenderer},
	})
}

type GLRenderer interface {
	Renderer
}

// glRenderer implements the GLRenderer interface.
type glRenderer struct {
	Renderer
}

var _ GLRenderer = (*glRenderer)(nil)

// WrapGLRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLRenderer(obj *externglib.Object) GLRenderer {
	return GLRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalGLRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLRenderer(obj), nil
}

// NewGLRenderer constructs a class GLRenderer.
func NewGLRenderer() GLRenderer {
	var cret C.GskGLRenderer
	var goret1 GLRenderer

	cret = C.gsk_gl_renderer_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GLRenderer)

	return goret1
}
