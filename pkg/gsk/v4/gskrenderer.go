// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_renderer_get_type()), F: marshalRendererer},
	})
}

// Rendererer describes Renderer's methods.
type Rendererer interface {
	// Surface retrieves the GdkSurface set using gsk_enderer_realize().
	Surface() *gdk.Surface
	// IsRealized checks whether the renderer is realized or not.
	IsRealized() bool
	// Realize creates the resources needed by the renderer to render the scene
	// graph.
	Realize(surface gdk.Surfacer) error
	// Render renders the scene graph, described by a tree of GskRenderNode
	// instances, ensuring that the given region gets redrawn.
	Render(root RenderNoder, region *cairo.Region)
	// RenderTexture renders the scene graph, described by a tree of
	// GskRenderNode instances, to a GdkTexture.
	RenderTexture(root RenderNoder, viewport *graphene.Rect) *gdk.Texture
	// Unrealize releases all the resources created by gsk_renderer_realize().
	Unrealize()
}

// Renderer: GskRenderer is a class that renders a scene graph defined via a
// tree of gsk.RenderNode instances.
//
// Typically you will use a GskRenderer instance to repeatedly call
// gsk.Renderer.Render() to update the contents of its associated gdk.Surface.
//
// It is necessary to realize a GskRenderer instance using
// gsk.Renderer.Realize() before calling gsk.Renderer.Render(), in order to
// create the appropriate windowing system resources needed to render the scene.
type Renderer struct {
	*externglib.Object
}

var (
	_ Rendererer      = (*Renderer)(nil)
	_ gextras.Nativer = (*Renderer)(nil)
)

func wrapRenderer(obj *externglib.Object) *Renderer {
	return &Renderer{
		Object: obj,
	}
}

func marshalRendererer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRenderer(obj), nil
}

// NewRendererForSurface creates an appropriate GskRenderer instance for the
// given surface.
//
// If the GSK_RENDERER environment variable is set, GSK will try that renderer
// first, before trying the backend-specific default. The ultimate fallback is
// the cairo renderer.
//
// The renderer will be realized before it is returned.
func NewRendererForSurface(surface gdk.Surfacer) *Renderer {
	var _arg1 *C.GdkSurface  // out
	var _cret *C.GskRenderer // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer((surface).(gextras.Nativer).Native()))

	_cret = C.gsk_renderer_new_for_surface(_arg1)

	var _renderer *Renderer // out

	_renderer = wrapRenderer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _renderer
}

// Surface retrieves the GdkSurface set using gsk_enderer_realize().
//
// If the renderer has not been realized yet, NULL will be returned.
func (renderer *Renderer) Surface() *gdk.Surface {
	var _arg0 *C.GskRenderer // out
	var _cret *C.GdkSurface  // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gsk_renderer_get_surface(_arg0)

	var _surface *gdk.Surface // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_surface = &gdk.Surface{
			Object: obj,
		}
	}

	return _surface
}

// IsRealized checks whether the renderer is realized or not.
func (renderer *Renderer) IsRealized() bool {
	var _arg0 *C.GskRenderer // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	_cret = C.gsk_renderer_is_realized(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Realize creates the resources needed by the renderer to render the scene
// graph.
func (renderer *Renderer) Realize(surface gdk.Surfacer) error {
	var _arg0 *C.GskRenderer // out
	var _arg1 *C.GdkSurface  // out
	var _cerr *C.GError      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer((surface).(gextras.Nativer).Native()))

	C.gsk_renderer_realize(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Render renders the scene graph, described by a tree of GskRenderNode
// instances, ensuring that the given region gets redrawn.
//
// Renderers must ensure that changes of the contents given by the root node as
// well as the area given by region are redrawn. They are however free to not
// redraw any pixel outside of region if they can guarantee that it didn't
// change.
//
// The renderer will acquire a reference on the GskRenderNode tree while the
// rendering is in progress.
func (renderer *Renderer) Render(root RenderNoder, region *cairo.Region) {
	var _arg0 *C.GskRenderer    // out
	var _arg1 *C.GskRenderNode  // out
	var _arg2 *C.cairo_region_t // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer((root).(gextras.Nativer).Native()))
	_arg2 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gsk_renderer_render(_arg0, _arg1, _arg2)
}

// RenderTexture renders the scene graph, described by a tree of GskRenderNode
// instances, to a GdkTexture.
//
// The renderer will acquire a reference on the GskRenderNode tree while the
// rendering is in progress.
//
// If you want to apply any transformations to root, you should put it into a
// transform node and pass that node instead.
func (renderer *Renderer) RenderTexture(root RenderNoder, viewport *graphene.Rect) *gdk.Texture {
	var _arg0 *C.GskRenderer     // out
	var _arg1 *C.GskRenderNode   // out
	var _arg2 *C.graphene_rect_t // out
	var _cret *C.GdkTexture      // in

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer((root).(gextras.Nativer).Native()))
	_arg2 = (*C.graphene_rect_t)(unsafe.Pointer(viewport))

	_cret = C.gsk_renderer_render_texture(_arg0, _arg1, _arg2)

	var _texture *gdk.Texture // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_texture = &gdk.Texture{
			Object: obj,
			Paintable: gdk.Paintable{
				Object: obj,
			},
		}
	}

	return _texture
}

// Unrealize releases all the resources created by gsk_renderer_realize().
func (renderer *Renderer) Unrealize() {
	var _arg0 *C.GskRenderer // out

	_arg0 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	C.gsk_renderer_unrealize(_arg0)
}
