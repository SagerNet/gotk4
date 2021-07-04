// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk.h>
import "C"

// RenderActivity renders an activity indicator (such as in Spinner). The state
// GTK_STATE_FLAG_CHECKED determines whether there is activity going on.
func RenderActivity(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_activity(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderArrow renders an arrow pointing to @angle.
//
// Typical arrow rendering at 0, 1⁄2 π;, π; and 3⁄2 π:
//
// ! (arrows.png)
func RenderArrow(context StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(angle)
	_arg4 = C.double(x)
	_arg5 = C.double(y)
	_arg6 = C.double(size)

	C.gtk_render_arrow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderBackground renders the background of an element.
//
// Typical background rendering, showing the effect of `background-image`,
// `border-width` and `border-radius`:
//
// ! (background.png)
func RenderBackground(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_background(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderCheck renders a checkmark (as in a CheckButton).
//
// The GTK_STATE_FLAG_CHECKED state determines whether the check is on or off,
// and GTK_STATE_FLAG_INCONSISTENT determines whether it should be marked as
// undefined.
//
// Typical checkmark rendering:
//
// ! (checks.png)
func RenderCheck(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_check(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderExpander renders an expander (as used in TreeView and Expander) in the
// area defined by @x, @y, @width, @height. The state GTK_STATE_FLAG_CHECKED
// determines whether the expander is collapsed or expanded.
//
// Typical expander rendering:
//
// ! (expanders.png)
func RenderExpander(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_expander(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderFocus renders a focus indicator on the rectangle determined by @x, @y,
// @width, @height.
//
// Typical focus rendering:
//
// ! (focus.png)
func RenderFocus(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_focus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderFrame renders a frame around the rectangle defined by @x, @y, @width,
// @height.
//
// Examples of frame rendering, showing the effect of `border-image`,
// `border-color`, `border-width`, `border-radius` and junctions:
//
// ! (frames.png)
func RenderFrame(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_frame(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderHandle renders a handle (as in Paned and Window’s resize grip), in the
// rectangle determined by @x, @y, @width, @height.
//
// Handles rendered for the paned and grip classes:
//
// ! (handles.png)
func RenderHandle(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_handle(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderIcon renders the icon in @texture at the specified @x and @y
// coordinates.
//
// This function will render the icon in @texture at exactly its size,
// regardless of scaling factors, which may not be appropriate when drawing on
// displays with high pixel densities.
func RenderIcon(context StyleContext, cr *cairo.Context, texture gdk.Texture, x float64, y float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 *C.GdkTexture      // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))
	_arg4 = C.double(x)
	_arg5 = C.double(y)

	C.gtk_render_icon(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderLayout renders @layout on the coordinates @x, @y
func RenderLayout(context StyleContext, cr *cairo.Context, x float64, y float64, layout pango.Layout) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 *C.PangoLayout     // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_render_layout(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderLine renders a line from (x0, y0) to (x1, y1).
func RenderLine(context StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x0)
	_arg4 = C.double(y0)
	_arg5 = C.double(x1)
	_arg6 = C.double(y1)

	C.gtk_render_line(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderOption renders an option mark (as in a radio button), the
// GTK_STATE_FLAG_CHECKED state will determine whether the option is on or off,
// and GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.
//
// Typical option mark rendering:
//
// ! (options.png)
func RenderOption(context StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.double           // out
	var _arg4 C.double           // out
	var _arg5 C.double           // out
	var _arg6 C.double           // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg3 = C.double(x)
	_arg4 = C.double(y)
	_arg5 = C.double(width)
	_arg6 = C.double(height)

	C.gtk_render_option(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}
