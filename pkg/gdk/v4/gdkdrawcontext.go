// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_draw_context_get_type()), F: marshalDrawContexter},
	})
}

// DrawContexter describes DrawContext's methods.
type DrawContexter interface {
	// BeginFrame indicates that you are beginning the process of redrawing
	// region on the context's surface.
	BeginFrame(region *cairo.Region)
	// EndFrame ends a drawing operation started with
	// gdk_draw_context_begin_frame().
	EndFrame()
	// Display retrieves the GdkDisplay the context is created for
	Display() *Display
	// FrameRegion retrieves the region that is currently being repainted.
	FrameRegion() *cairo.Region
	// Surface retrieves the surface that context is bound to.
	Surface() *Surface
	// IsInFrame returns TRUE if context is in the process of drawing to its
	// surface.
	IsInFrame() bool
}

// DrawContext: base class for objects implementing different rendering methods.
//
// GdkDrawContext is the base object used by contexts implementing different
// rendering methods, such as gdk.CairoContext or gdk.GLContext. It provides
// shared functionality between those contexts.
//
// You will always interact with one of those subclasses.
//
// A GdkDrawContext is always associated with a single toplevel surface.
type DrawContext struct {
	*externglib.Object
}

var (
	_ DrawContexter   = (*DrawContext)(nil)
	_ gextras.Nativer = (*DrawContext)(nil)
)

func wrapDrawContext(obj *externglib.Object) *DrawContext {
	return &DrawContext{
		Object: obj,
	}
}

func marshalDrawContexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDrawContext(obj), nil
}

// BeginFrame indicates that you are beginning the process of redrawing region
// on the context's surface.
//
// Calling this function begins a drawing operation using context on the surface
// that context was created from. The actual requirements and guarantees for the
// drawing operation vary for different implementations of drawing, so a
// gdk.CairoContext and a gdk.GLContext need to be treated differently.
//
// A call to this function is a requirement for drawing and must be followed by
// a call to gdk.DrawContext.EndFrame(), which will complete the drawing
// operation and ensure the contents become visible on screen.
//
// Note that the region passed to this function is the minimum region that needs
// to be drawn and depending on implementation, windowing system and hardware in
// use, it might be necessary to draw a larger region. Drawing implementation
// must use [methodGdk.DrawContext.get_frame_region() to query the region that
// must be drawn.
//
// When using GTK, the widget system automatically places calls to
// gdk_draw_context_begin_frame() and gdk_draw_context_end_frame() via the use
// of gsk.Renderers, so application code does not need to call these functions
// explicitly.
func (context *DrawContext) BeginFrame(region *cairo.Region) {
	var _arg0 *C.GdkDrawContext // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_draw_context_begin_frame(_arg0, _arg1)
}

// EndFrame ends a drawing operation started with
// gdk_draw_context_begin_frame().
//
// This makes the drawing available on screen. See gdk.DrawContext.BeginFrame()
// for more details about drawing.
//
// When using a gdk.GLContext, this function may call glFlush() implicitly
// before returning; it is not recommended to call glFlush() explicitly before
// calling this function.
func (context *DrawContext) EndFrame() {
	var _arg0 *C.GdkDrawContext // out

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))

	C.gdk_draw_context_end_frame(_arg0)
}

// Display retrieves the GdkDisplay the context is created for
func (context *DrawContext) Display() *Display {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.GdkDisplay     // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_draw_context_get_display(_arg0)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// FrameRegion retrieves the region that is currently being repainted.
//
// After a call to gdk.DrawContext.BeginFrame() this function will return a
// union of the region passed to that function and the area of the surface that
// the context determined needs to be repainted.
//
// If context is not in between calls to gdk.DrawContext.BeginFrame() and
// gdk.DrawContext.EndFrame(), NULL will be returned.
func (context *DrawContext) FrameRegion() *cairo.Region {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.cairo_region_t // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_draw_context_get_frame_region(_arg0)

	var _region *cairo.Region // out

	{
		v := &struct{ p unsafe.Pointer }{unsafe.Pointer(_cret)}
		_region = (*cairo.Region)(unsafe.Pointer(v))
	}
	runtime.SetFinalizer(_region, func(v *cairo.Region) {
		C.cairo_region_destroy((*C.cairo_region_t)(unsafe.Pointer(v.Native())))
	})

	return _region
}

// Surface retrieves the surface that context is bound to.
func (context *DrawContext) Surface() *Surface {
	var _arg0 *C.GdkDrawContext // out
	var _cret *C.GdkSurface     // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_draw_context_get_surface(_arg0)

	var _surface *Surface // out

	_surface = wrapSurface(externglib.Take(unsafe.Pointer(_cret)))

	return _surface
}

// IsInFrame returns TRUE if context is in the process of drawing to its
// surface.
//
// This is the case between calls to gdk.DrawContext.BeginFrame() and
// gdk.DrawContext.EndFrame(). In this situation, drawing commands may be
// effecting the contents of the context's surface.
func (context *DrawContext) IsInFrame() bool {
	var _arg0 *C.GdkDrawContext // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDrawContext)(unsafe.Pointer(context.Native()))

	_cret = C.gdk_draw_context_is_in_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
