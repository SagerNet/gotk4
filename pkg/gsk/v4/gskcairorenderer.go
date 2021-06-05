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
		{T: externglib.Type(C.gsk_cairo_renderer_get_type()), F: marshalCairoRenderer},
	})
}

type CairoRenderer interface {
	Renderer
}

// cairoRenderer implements the CairoRenderer interface.
type cairoRenderer struct {
	Renderer
}

var _ CairoRenderer = (*cairoRenderer)(nil)

// WrapCairoRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapCairoRenderer(obj *externglib.Object) CairoRenderer {
	return CairoRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalCairoRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCairoRenderer(obj), nil
}

// NewCairoRenderer constructs a class CairoRenderer.
func NewCairoRenderer() CairoRenderer {
	var cret C.GskCairoRenderer
	var goret1 CairoRenderer

	cret = C.gsk_cairo_renderer_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(CairoRenderer)

	return goret1
}
