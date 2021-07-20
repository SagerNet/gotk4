// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_snapshot_get_type()), F: marshalSnapshotter},
	})
}

// Snapshot: GtkSnapshot assists in creating GskRenderNodes for widgets.
//
// It functions in a similar way to a cairo context, and maintains a stack of
// render nodes and their associated transformations.
//
// The node at the top of the stack is the the one that gtk_snapshot_append_…
// functions operate on. Use the gtk_snapshot_push_… functions and
// gtk_snapshot_pop() to change the current node.
//
// The typical way to obtain a GtkSnapshot object is as an argument to the
// GtkWidgetClass.snapshot() vfunc. If you need to create your own GtkSnapshot,
// use gtk.Snapshot.New.
type Snapshot struct {
	gdk.Snapshot
}

var _ gextras.Nativer = (*Snapshot)(nil)

func wrapSnapshot(obj *externglib.Object) *Snapshot {
	return &Snapshot{
		Snapshot: gdk.Snapshot{
			Object: obj,
		},
	}
}

func marshalSnapshotter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSnapshot(obj), nil
}

// NewSnapshot creates a new GtkSnapshot.
func NewSnapshot() *Snapshot {
	var _cret *C.GtkSnapshot // in

	_cret = C.gtk_snapshot_new()

	var _snapshot *Snapshot // out

	_snapshot = wrapSnapshot(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _snapshot
}

// AppendBorder appends a stroked border rectangle inside the given outline.
//
// The four sides of the border can have different widths and colors.
func (snapshot *Snapshot) AppendBorder(outline *gsk.RoundedRect, borderWidth [4]float32, borderColor [4]gdk.RGBA) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.float          // out
	var _arg3 *C.GdkRGBA        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskRoundedRect)(gextras.StructNative(unsafe.Pointer(outline)))
	_arg2 = (*C.float)(unsafe.Pointer(&borderWidth))
	_arg3 = (*C.GdkRGBA)(unsafe.Pointer(&borderColor))

	C.gtk_snapshot_append_border(_arg0, _arg1, _arg2, _arg3)
}

// AppendCairo creates a new GskCairoNode and appends it to the current render
// node of snapshot, without changing the current node.
func (snapshot *Snapshot) AppendCairo(bounds *graphene.Rect) *cairo.Context {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))

	_cret = C.gtk_snapshot_append_cairo(_arg0, _arg1)

	var _context *cairo.Context // out

	_context = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	C.cairo_reference(_cret)
	runtime.SetFinalizer(_context, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _context
}

// AppendColor creates a new render node drawing the color into the given bounds
// and appends it to the current render node of snapshot.
//
// You should try to avoid calling this function if color is transparent.
func (snapshot *Snapshot) AppendColor(color *gdk.RGBA, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkRGBA         // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))
	_arg2 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))

	C.gtk_snapshot_append_color(_arg0, _arg1, _arg2)
}

// AppendConicGradient appends a conic gradient node with the given stops to
// snapshot.
func (snapshot *Snapshot) AppendConicGradient(bounds *graphene.Rect, center *graphene.Point, rotation float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 *C.GskColorStop     // out
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(center)))
	_arg3 = C.float(rotation)
	_arg5 = (C.gsize)(len(stops))
	if len(stops) > 0 {
		_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))
	}

	C.gtk_snapshot_append_conic_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// AppendInsetShadow appends an inset shadow into the box given by outline.
func (snapshot *Snapshot) AppendInsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskRoundedRect)(gextras.StructNative(unsafe.Pointer(outline)))
	_arg2 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_inset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (snapshot *Snapshot) AppendLayout(layout *pango.Layout, color *gdk.RGBA) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.PangoLayout // out
	var _arg2 *C.GdkRGBA     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg2 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))

	C.gtk_snapshot_append_layout(_arg0, _arg1, _arg2)
}

// AppendLinearGradient appends a linear gradient node with the given stops to
// snapshot.
func (snapshot *Snapshot) AppendLinearGradient(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.GskColorStop     // out
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(startPoint)))
	_arg3 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(endPoint)))
	_arg5 = (C.gsize)(len(stops))
	if len(stops) > 0 {
		_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))
	}

	C.gtk_snapshot_append_linear_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// AppendNode appends node to the current render node of snapshot, without
// changing the current node.
//
// If snapshot does not have a current node yet, node will become the initial
// node.
func (snapshot *Snapshot) AppendNode(node gsk.RenderNoder) {
	var _arg0 *C.GtkSnapshot   // out
	var _arg1 *C.GskRenderNode // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskRenderNode)(unsafe.Pointer((node).(gextras.Nativer).Native()))

	C.gtk_snapshot_append_node(_arg0, _arg1)
}

// AppendOutsetShadow appends an outset shadow node around the box given by
// outline.
func (snapshot *Snapshot) AppendOutsetShadow(outline *gsk.RoundedRect, color *gdk.RGBA, dx float32, dy float32, spread float32, blurRadius float32) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out
	var _arg2 *C.GdkRGBA        // out
	var _arg3 C.float           // out
	var _arg4 C.float           // out
	var _arg5 C.float           // out
	var _arg6 C.float           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskRoundedRect)(gextras.StructNative(unsafe.Pointer(outline)))
	_arg2 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(color)))
	_arg3 = C.float(dx)
	_arg4 = C.float(dy)
	_arg5 = C.float(spread)
	_arg6 = C.float(blurRadius)

	C.gtk_snapshot_append_outset_shadow(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// AppendRadialGradient appends a radial gradient node with the given stops to
// snapshot.
func (snapshot *Snapshot) AppendRadialGradient(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 C.float             // out
	var _arg5 C.float             // out
	var _arg6 C.float             // out
	var _arg7 *C.GskColorStop     // out
	var _arg8 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(center)))
	_arg3 = C.float(hradius)
	_arg4 = C.float(vradius)
	_arg5 = C.float(start)
	_arg6 = C.float(end)
	_arg8 = (C.gsize)(len(stops))
	if len(stops) > 0 {
		_arg7 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))
	}

	C.gtk_snapshot_append_radial_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// AppendRepeatingLinearGradient appends a repeating linear gradient node with
// the given stops to snapshot.
func (snapshot *Snapshot) AppendRepeatingLinearGradient(bounds *graphene.Rect, startPoint *graphene.Point, endPoint *graphene.Point, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 *C.graphene_point_t // out
	var _arg4 *C.GskColorStop     // out
	var _arg5 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(startPoint)))
	_arg3 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(endPoint)))
	_arg5 = (C.gsize)(len(stops))
	if len(stops) > 0 {
		_arg4 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))
	}

	C.gtk_snapshot_append_repeating_linear_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// AppendRepeatingRadialGradient appends a repeating radial gradient node with
// the given stops to snapshot.
func (snapshot *Snapshot) AppendRepeatingRadialGradient(bounds *graphene.Rect, center *graphene.Point, hradius float32, vradius float32, start float32, end float32, stops []gsk.ColorStop) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_rect_t  // out
	var _arg2 *C.graphene_point_t // out
	var _arg3 C.float             // out
	var _arg4 C.float             // out
	var _arg5 C.float             // out
	var _arg6 C.float             // out
	var _arg7 *C.GskColorStop     // out
	var _arg8 C.gsize

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(center)))
	_arg3 = C.float(hradius)
	_arg4 = C.float(vradius)
	_arg5 = C.float(start)
	_arg6 = C.float(end)
	_arg8 = (C.gsize)(len(stops))
	if len(stops) > 0 {
		_arg7 = (*C.GskColorStop)(unsafe.Pointer(&stops[0]))
	}

	C.gtk_snapshot_append_repeating_radial_gradient(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8)
}

// AppendTexture creates a new render node drawing the texture into the given
// bounds and appends it to the current render node of snapshot.
func (snapshot *Snapshot) AppendTexture(texture gdk.Texturer, bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GdkTexture      // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GdkTexture)(unsafe.Pointer((texture).(gextras.Nativer).Native()))
	_arg2 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))

	C.gtk_snapshot_append_texture(_arg0, _arg1, _arg2)
}

// GLShaderPopTexture removes the top element from the stack of render nodes and
// adds it to the nearest GskGLShaderNode below it.
//
// This must be called the same number of times as the number of textures is
// needed for the shader in gtk.Snapshot.PushGLShader().
func (snapshot *Snapshot) GLShaderPopTexture() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))

	C.gtk_snapshot_gl_shader_pop_texture(_arg0)
}

// Perspective applies a perspective projection transform.
//
// See gsk.Transform.Perspective() for a discussion on the details.
func (snapshot *Snapshot) Perspective(depth float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.float(depth)

	C.gtk_snapshot_perspective(_arg0, _arg1)
}

// Pop removes the top element from the stack of render nodes, and appends it to
// the node underneath it.
func (snapshot *Snapshot) Pop() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))

	C.gtk_snapshot_pop(_arg0)
}

// PushBlend blends together two images with the given blend mode.
//
// Until the first call to gtk.Snapshot.Pop(), the bottom image for the blend
// operation will be recorded. After that call, the top image to be blended will
// be recorded until the second call to gtk.Snapshot.Pop().
//
// Calling this function requires two subsequent calls to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushBlend(blendMode gsk.BlendMode) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.GskBlendMode // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.GskBlendMode(blendMode)

	C.gtk_snapshot_push_blend(_arg0, _arg1)
}

// PushBlur blurs an image.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushBlur(radius float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.double(radius)

	C.gtk_snapshot_push_blur(_arg0, _arg1)
}

// PushClip clips an image to a rectangle.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushClip(bounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))

	C.gtk_snapshot_push_clip(_arg0, _arg1)
}

// PushColorMatrix modifies the colors of an image by applying an affine
// transformation in RGB space.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushColorMatrix(colorMatrix *graphene.Matrix, colorOffset *graphene.Vec4) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out
	var _arg2 *C.graphene_vec4_t   // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(colorMatrix)))
	_arg2 = (*C.graphene_vec4_t)(gextras.StructNative(unsafe.Pointer(colorOffset)))

	C.gtk_snapshot_push_color_matrix(_arg0, _arg1, _arg2)
}

// PushCrossFade snapshots a cross-fade operation between two images with the
// given progress.
//
// Until the first call to gtk.Snapshot.Pop(), the start image will be snapshot.
// After that call, the end image will be recorded until the second call to
// gtk.Snapshot.Pop().
//
// Calling this function requires two subsequent calls to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushCrossFade(progress float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.double(progress)

	C.gtk_snapshot_push_cross_fade(_arg0, _arg1)
}

// PushOpacity modifies the opacity of an image.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushOpacity(opacity float64) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.double(opacity)

	C.gtk_snapshot_push_opacity(_arg0, _arg1)
}

// PushRepeat creates a node that repeats the child node.
//
// The child is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushRepeat(bounds *graphene.Rect, childBounds *graphene.Rect) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_rect_t // out
	var _arg2 *C.graphene_rect_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(bounds)))
	_arg2 = (*C.graphene_rect_t)(gextras.StructNative(unsafe.Pointer(childBounds)))

	C.gtk_snapshot_push_repeat(_arg0, _arg1, _arg2)
}

// PushRoundedClip clips an image to a rounded rectangle.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushRoundedClip(bounds *gsk.RoundedRect) {
	var _arg0 *C.GtkSnapshot    // out
	var _arg1 *C.GskRoundedRect // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskRoundedRect)(gextras.StructNative(unsafe.Pointer(bounds)))

	C.gtk_snapshot_push_rounded_clip(_arg0, _arg1)
}

// PushShadow applies a shadow to an image.
//
// The image is recorded until the next call to gtk.Snapshot.Pop().
func (snapshot *Snapshot) PushShadow(shadow *gsk.Shadow, nShadows uint) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 *C.GskShadow   // out
	var _arg2 C.gsize        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskShadow)(gextras.StructNative(unsafe.Pointer(shadow)))
	_arg2 = C.gsize(nShadows)

	C.gtk_snapshot_push_shadow(_arg0, _arg1, _arg2)
}

// RenderBackground creates a render node for the CSS background according to
// context, and appends it to the current node of snapshot, without changing the
// current node.
func (snapshot *Snapshot) RenderBackground(context *StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_background(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderFocus creates a render node for the focus outline according to context,
// and appends it to the current node of snapshot, without changing the current
// node.
func (snapshot *Snapshot) RenderFocus(context *StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_focus(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderFrame creates a render node for the CSS border according to context,
// and appends it to the current node of snapshot, without changing the current
// node.
func (snapshot *Snapshot) RenderFrame(context *StyleContext, x float64, y float64, width float64, height float64) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.gtk_snapshot_render_frame(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderInsertionCursor draws a text caret using snapshot at the specified
// index of layout.
func (snapshot *Snapshot) RenderInsertionCursor(context *StyleContext, x float64, y float64, layout *pango.Layout, index int, direction pango.Direction) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 *C.PangoLayout     // out
	var _arg5 C.int              // out
	var _arg6 C.PangoDirection   // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	_arg5 = C.int(index)
	_arg6 = C.PangoDirection(direction)

	C.gtk_snapshot_render_insertion_cursor(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderLayout creates a render node for rendering layout according to the
// style information in context, and appends it to the current node of snapshot,
// without changing the current node.
func (snapshot *Snapshot) RenderLayout(context *StyleContext, x float64, y float64, layout *pango.Layout) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out
	var _arg4 *C.PangoLayout     // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_snapshot_render_layout(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Restore restores snapshot to the state saved by a preceding call to
// gtk_snapshot_save() and removes that state from the stack of saved states.
func (snapshot *Snapshot) Restore() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))

	C.gtk_snapshot_restore(_arg0)
}

// Rotate rotates @snapshot's coordinate system by angle degrees in 2D space -
// or in 3D speak, rotates around the Z axis.
//
// To rotate around other axes, use gsk.Transform.Rotate3D().
func (snapshot *Snapshot) Rotate(angle float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.float(angle)

	C.gtk_snapshot_rotate(_arg0, _arg1)
}

// Rotate3D rotates snapshot's coordinate system by angle degrees around axis.
//
// For a rotation in 2D space, use gsk.Transform.Rotate().
func (snapshot *Snapshot) Rotate3D(angle float32, axis *graphene.Vec3) {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 C.float            // out
	var _arg2 *C.graphene_vec3_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.float(angle)
	_arg2 = (*C.graphene_vec3_t)(gextras.StructNative(unsafe.Pointer(axis)))

	C.gtk_snapshot_rotate_3d(_arg0, _arg1, _arg2)
}

// Save makes a copy of the current state of snapshot and saves it on an
// internal stack.
//
// When gtk.Snapshot.Restore() is called, snapshot will be restored to the saved
// state. Multiple calls to gtk_snapshot_save() and gtk_snapshot_restore() can
// be nested; each call to gtk_snapshot_restore() restores the state from the
// matching paired gtk_snapshot_save().
//
// It is necessary to clear all saved states with corresponding calls to
// gtk_snapshot_restore().
func (snapshot *Snapshot) Save() {
	var _arg0 *C.GtkSnapshot // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))

	C.gtk_snapshot_save(_arg0)
}

// Scale scales snapshot's coordinate system in 2-dimensional space by the given
// factors.
//
// Use gtk.Snapshot.Scale3D() to scale in all 3 dimensions.
func (snapshot *Snapshot) Scale(factorX float32, factorY float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)

	C.gtk_snapshot_scale(_arg0, _arg1, _arg2)
}

// Scale3D scales snapshot's coordinate system by the given factors.
func (snapshot *Snapshot) Scale3D(factorX float32, factorY float32, factorZ float32) {
	var _arg0 *C.GtkSnapshot // out
	var _arg1 C.float        // out
	var _arg2 C.float        // out
	var _arg3 C.float        // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = C.float(factorX)
	_arg2 = C.float(factorY)
	_arg3 = C.float(factorZ)

	C.gtk_snapshot_scale_3d(_arg0, _arg1, _arg2, _arg3)
}

// ToNode returns the render node that was constructed by snapshot.
//
// After calling this function, it is no longer possible to add more nodes to
// snapshot. The only function that should be called after this is
// g_object_unref().
func (snapshot *Snapshot) ToNode() gsk.RenderNoder {
	var _arg0 *C.GtkSnapshot   // out
	var _cret *C.GskRenderNode // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))

	_cret = C.gtk_snapshot_to_node(_arg0)

	var _renderNode gsk.RenderNoder // out

	_renderNode = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gsk.RenderNoder)

	return _renderNode
}

// ToPaintable returns a paintable encapsulating the render node that was
// constructed by snapshot.
//
// After calling this function, it is no longer possible to add more nodes to
// snapshot. The only function that should be called after this is
// g_object_unref().
func (snapshot *Snapshot) ToPaintable(size *graphene.Size) gdk.Paintabler {
	var _arg0 *C.GtkSnapshot     // out
	var _arg1 *C.graphene_size_t // out
	var _cret *C.GdkPaintable    // in

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_size_t)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_snapshot_to_paintable(_arg0, _arg1)

	var _paintable gdk.Paintabler // out

	_paintable = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gdk.Paintabler)

	return _paintable
}

// Transform transforms snapshot's coordinate system with the given transform.
func (snapshot *Snapshot) Transform(transform *gsk.Transform) {
	var _arg0 *C.GtkSnapshot  // out
	var _arg1 *C.GskTransform // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.GskTransform)(gextras.StructNative(unsafe.Pointer(transform)))

	C.gtk_snapshot_transform(_arg0, _arg1)
}

// TransformMatrix transforms snapshot's coordinate system with the given
// matrix.
func (snapshot *Snapshot) TransformMatrix(matrix *graphene.Matrix) {
	var _arg0 *C.GtkSnapshot       // out
	var _arg1 *C.graphene_matrix_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_matrix_t)(gextras.StructNative(unsafe.Pointer(matrix)))

	C.gtk_snapshot_transform_matrix(_arg0, _arg1)
}

// Translate translates snapshot's coordinate system by point in 2-dimensional
// space.
func (snapshot *Snapshot) Translate(point *graphene.Point) {
	var _arg0 *C.GtkSnapshot      // out
	var _arg1 *C.graphene_point_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_point_t)(gextras.StructNative(unsafe.Pointer(point)))

	C.gtk_snapshot_translate(_arg0, _arg1)
}

// Translate3D translates snapshot's coordinate system by point.
func (snapshot *Snapshot) Translate3D(point *graphene.Point3D) {
	var _arg0 *C.GtkSnapshot        // out
	var _arg1 *C.graphene_point3d_t // out

	_arg0 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg1 = (*C.graphene_point3d_t)(gextras.StructNative(unsafe.Pointer(point)))

	C.gtk_snapshot_translate_3d(_arg0, _arg1)
}
