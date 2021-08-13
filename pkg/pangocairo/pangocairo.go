// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"runtime"
	_ "runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: pangocairo pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pangocairo.h>
// extern void callbackDelete(gpointer);
// void _gotk4_pangocairo1_ShapeRendererFunc(cairo_t*, PangoAttrShape*, gboolean, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_cairo_font_get_type()), F: marshalFonter},
		{T: externglib.Type(C.pango_cairo_font_map_get_type()), F: marshalFontMapper},
	})
}

// ShapeRendererFunc: function type for rendering attributes of type
// PANGO_ATTR_SHAPE with Pango's Cairo renderer.
type ShapeRendererFunc func(cr *cairo.Context, attr *pango.AttrShape, doPath bool)

//export _gotk4_pangocairo1_ShapeRendererFunc
func _gotk4_pangocairo1_ShapeRendererFunc(arg0 *C.cairo_t, arg1 *C.PangoAttrShape, arg2 C.gboolean, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var cr *cairo.Context     // out
	var attr *pango.AttrShape // out
	var doPath bool           // out

	cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg0)))
	runtime.SetFinalizer(cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})
	attr = (*pango.AttrShape)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(attr, func(v *pango.AttrShape) {
		C.free(gextras.StructNative(unsafe.Pointer(v)))
	})
	if arg2 != 0 {
		doPath = true
	}

	fn := v.(ShapeRendererFunc)
	fn(cr, attr, doPath)
}

// ContextGetFontOptions retrieves any font rendering options previously set
// with pangocairo.ContextSetFontOptions().
//
// This function does not report options that are derived from the target
// surface by update_context.
func ContextGetFontOptions(context *pango.Context) *cairo.FontOptions {
	var _arg1 *C.PangoContext         // out
	var _cret *C.cairo_font_options_t // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_cairo_context_get_font_options(_arg1)
	runtime.KeepAlive(context)

	var _fontOptions *cairo.FontOptions // out

	if _cret != nil {
		_fontOptions = (*cairo.FontOptions)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _fontOptions
}

// ContextGetResolution gets the resolution for the context. See
// pangocairo.ContextSetResolution()
func ContextGetResolution(context *pango.Context) float64 {
	var _arg1 *C.PangoContext // out
	var _cret C.double        // in

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	_cret = C.pango_cairo_context_get_resolution(_arg1)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ContextSetFontOptions sets the font options used when rendering text with
// this context.
//
// These options override any options that update_context derives from the
// target surface.
func ContextSetFontOptions(context *pango.Context, options *cairo.FontOptions) {
	var _arg1 *C.PangoContext         // out
	var _arg2 *C.cairo_font_options_t // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if options != nil {
		_arg2 = (*C.cairo_font_options_t)(gextras.StructNative(unsafe.Pointer(options)))
	}

	C.pango_cairo_context_set_font_options(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(options)
}

// ContextSetResolution sets the resolution for the context.
//
// This is a scale factor between points specified in a PangoFontDescription and
// Cairo units. The default value is 96, meaning that a 10 point font will be 13
// units high. (10 * 96. / 72. = 13.3).
func ContextSetResolution(context *pango.Context, dpi float64) {
	var _arg1 *C.PangoContext // out
	var _arg2 C.double        // out

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.double(dpi)

	C.pango_cairo_context_set_resolution(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(dpi)
}

// ContextSetShapeRenderer sets callback function for context to use for
// rendering attributes of type PANGO_ATTR_SHAPE.
//
// See PangoCairoShapeRendererFunc for details.
func ContextSetShapeRenderer(context *pango.Context, fn ShapeRendererFunc) {
	var _arg1 *C.PangoContext               // out
	var _arg2 C.PangoCairoShapeRendererFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	if fn != nil {
		_arg2 = (*[0]byte)(C._gotk4_pangocairo1_ShapeRendererFunc)
		_arg3 = C.gpointer(gbox.Assign(fn))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.pango_cairo_context_set_shape_renderer(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(fn)
}

// CreateContext creates a context object set up to match the current
// transformation and target surface of the Cairo context.
//
// This context can then be used to create a layout using pango.Layout.New.
//
// This function is a convenience function that creates a context using the
// default font map, then updates it to cr. If you just need to create a layout
// for use with cr and do not need to access PangoContext directly, you can use
// create_layout instead.
func CreateContext(cr *cairo.Context) *pango.Context {
	var _arg1 *C.cairo_t      // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.pango_cairo_create_context(_arg1)
	runtime.KeepAlive(cr)

	var _context *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}

// CreateLayout creates a layout object set up to match the current
// transformation and target surface of the Cairo context.
//
// This layout can then be used for text measurement with functions like
// pango.Layout.GetSize() or drawing with functions like show_layout. If you
// change the transformation or target surface for cr, you need to call
// update_layout.
//
// This function is the most convenient way to use Cairo with Pango, however it
// is slightly inefficient since it creates a separate PangoContext object for
// each layout. This might matter in an application that was laying out large
// amounts of text.
func CreateLayout(cr *cairo.Context) *pango.Layout {
	var _arg1 *C.cairo_t     // out
	var _cret *C.PangoLayout // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.pango_cairo_create_layout(_arg1)
	runtime.KeepAlive(cr)

	var _layout *pango.Layout // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// ErrorUnderlinePath: add a squiggly line to the current path in the specified
// cairo context that approximately covers the given rectangle in the style of
// an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_error_underline_path(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// GlyphStringPath adds the glyphs in glyphs to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be at the
// current point of the cairo context.
func GlyphStringPath(cr *cairo.Context, font pango.Fonter, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg3 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_cairo_glyph_string_path(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyphs)
}

// LayoutLinePath adds the text in PangoLayoutLine to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be at the current
// point of the cairo context.
func LayoutLinePath(cr *cairo.Context, line *pango.LayoutLine) {
	var _arg1 *C.cairo_t         // out
	var _arg2 *C.PangoLayoutLine // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayoutLine)(gextras.StructNative(unsafe.Pointer(line)))

	C.pango_cairo_layout_line_path(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(line)
}

// LayoutPath adds the text in a PangoLayout to the current path in the
// specified cairo context.
//
// The top-left corner of the PangoLayout will be at the current point of the
// cairo context.
func LayoutPath(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_layout_path(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// ShowErrorUnderline: draw a squiggly line in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used to
// indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.cairo_t // out
	var _arg2 C.double   // out
	var _arg3 C.double   // out
	var _arg4 C.double   // out
	var _arg5 C.double   // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(x)
	_arg3 = C.double(y)
	_arg4 = C.double(width)
	_arg5 = C.double(height)

	C.pango_cairo_show_error_underline(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// ShowGlyphItem draws the glyphs in glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// show_glyph_string.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.char           // out
	var _arg3 *C.PangoGlyphItem // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))

	C.pango_cairo_show_glyph_item(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(text)
	runtime.KeepAlive(glyphItem)
}

// ShowGlyphString draws the glyphs in glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
func ShowGlyphString(cr *cairo.Context, font pango.Fonter, glyphs *pango.GlyphString) {
	var _arg1 *C.cairo_t          // out
	var _arg2 *C.PangoFont        // out
	var _arg3 *C.PangoGlyphString // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoFont)(unsafe.Pointer(font.Native()))
	_arg3 = (*C.PangoGlyphString)(gextras.StructNative(unsafe.Pointer(glyphs)))

	C.pango_cairo_show_glyph_string(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(font)
	runtime.KeepAlive(glyphs)
}

// ShowLayout draws a PangoLayout in the specified cairo context.
//
// The top-left corner of the PangoLayout will be drawn at the current point of
// the cairo context.
func ShowLayout(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_show_layout(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// ShowLayoutLine draws a PangoLayoutLine in the specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be drawn at the
// current point of the cairo context.
func ShowLayoutLine(cr *cairo.Context, line *pango.LayoutLine) {
	var _arg1 *C.cairo_t         // out
	var _arg2 *C.PangoLayoutLine // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayoutLine)(gextras.StructNative(unsafe.Pointer(line)))

	C.pango_cairo_show_layout_line(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(line)
}

// UpdateContext updates a PangoContext previously created for use with Cairo to
// match the current transformation and target surface of a Cairo context.
//
// If any layouts have been created for the context, it's necessary to call
// pango.Layout.ContextChanged() on those layouts.
func UpdateContext(cr *cairo.Context, context *pango.Context) {
	var _arg1 *C.cairo_t      // out
	var _arg2 *C.PangoContext // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoContext)(unsafe.Pointer(context.Native()))

	C.pango_cairo_update_context(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(context)
}

// UpdateLayout updates the private PangoContext of a PangoLayout created with
// create_layout to match the current transformation and target surface of a
// Cairo context.
func UpdateLayout(cr *cairo.Context, layout *pango.Layout) {
	var _arg1 *C.cairo_t     // out
	var _arg2 *C.PangoLayout // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.pango_cairo_update_layout(_arg1, _arg2)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(layout)
}

// Font: PangoCairoFont is an interface exported by fonts for use with Cairo.
//
// The actual type of the font will depend on the particular font technology
// Cairo was compiled to use.
type Font struct {
	pango.Font
}

// Fonter describes Font's abstract methods.
type Fonter interface {
	externglib.Objector

	privateFont()
}

var _ Fonter = (*Font)(nil)

func wrapFont(obj *externglib.Object) *Font {
	return &Font{
		Font: pango.Font{
			Object: obj,
		},
	}
}

func marshalFonter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFont(obj), nil
}

func (*Font) privateFont() {}

// FontMap: PangoCairoFontMap is an interface exported by font maps for use with
// Cairo.
//
// The actual type of the font map will depend on the particular font technology
// Cairo was compiled to use.
type FontMap struct {
	pango.FontMap
}

// FontMapper describes FontMap's abstract methods.
type FontMapper interface {
	externglib.Objector

	// Resolution gets the resolution for the fontmap.
	Resolution() float64
	// SetDefault sets a default PangoCairoFontMap to use with Cairo.
	SetDefault()
	// SetResolution sets the resolution for the fontmap.
	SetResolution(dpi float64)
}

var _ FontMapper = (*FontMap)(nil)

func wrapFontMap(obj *externglib.Object) *FontMap {
	return &FontMap{
		FontMap: pango.FontMap{
			Object: obj,
		},
	}
}

func marshalFontMapper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFontMap(obj), nil
}

// Resolution gets the resolution for the fontmap.
//
// See pangocairo.FontMap.SetResolution().
func (fontmap *FontMap) Resolution() float64 {
	var _arg0 *C.PangoCairoFontMap // out
	var _cret C.double             // in

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(fontmap.Native()))

	_cret = C.pango_cairo_font_map_get_resolution(_arg0)
	runtime.KeepAlive(fontmap)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetDefault sets a default PangoCairoFontMap to use with Cairo.
//
// This can be used to change the Cairo font backend that the default fontmap
// uses for example. The old default font map is unreffed and the new font map
// referenced.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. This
// function only changes the default fontmap for the current thread. Default
// fontmaps of existing threads are not changed. Default fontmaps of any new
// threads will still be created using pangocairo.FontMap.New.
//
// A value of NULL for fontmap will cause the current default font map to be
// released and a new default font map to be created on demand, using
// pangocairo.FontMap.New.
func (fontmap *FontMap) SetDefault() {
	var _arg0 *C.PangoCairoFontMap // out

	if fontmap != nil {
		_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(fontmap.Native()))
	}

	C.pango_cairo_font_map_set_default(_arg0)
	runtime.KeepAlive(fontmap)
}

// SetResolution sets the resolution for the fontmap.
//
// This is a scale factor between points specified in a PangoFontDescription and
// Cairo units. The default value is 96, meaning that a 10 point font will be 13
// units high. (10 * 96. / 72. = 13.3).
func (fontmap *FontMap) SetResolution(dpi float64) {
	var _arg0 *C.PangoCairoFontMap // out
	var _arg1 C.double             // out

	_arg0 = (*C.PangoCairoFontMap)(unsafe.Pointer(fontmap.Native()))
	_arg1 = C.double(dpi)

	C.pango_cairo_font_map_set_resolution(_arg0, _arg1)
	runtime.KeepAlive(fontmap)
	runtime.KeepAlive(dpi)
}

// FontMapGetDefault gets a default PangoCairoFontMap to use with Cairo.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; you generally should only use the
// PangoFontMap and PangoCairoFontMap interfaces on the returned object.
//
// The default Cairo fontmap can be changed by using
// pangocairo.FontMap.SetDefault(). This can be used to change the Cairo font
// backend that the default fontmap uses for example.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. Each thread
// gets its own default fontmap. In this way, PangoCairo can be used safely from
// multiple threads.
func FontMapGetDefault() pango.FontMapper {
	var _cret *C.PangoFontMap // in

	_cret = C.pango_cairo_font_map_get_default()

	var _fontMap pango.FontMapper // out

	_fontMap = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(pango.FontMapper)

	return _fontMap
}

// NewFontMap creates a new PangoCairoFontMap object.
//
// A fontmap is used to cache information about available fonts, and holds
// certain global parameters such as the resolution. In most cases, you can use
// `funcPangoCairo.font_map_get_default] instead.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; You generally should only use the
// PangoFontMap and PangoCairoFontMap interfaces on the returned object.
//
// You can override the type of backend returned by using an environment
// variable PANGOCAIRO_BACKEND. Supported types, based on your build, are fc
// (fontconfig), win32, and coretext. If requested type is not available, NULL
// is returned. Ie. this is only useful for testing, when at least two backends
// are compiled in.
func NewFontMap() pango.FontMapper {
	var _cret *C.PangoFontMap // in

	_cret = C.pango_cairo_font_map_new()

	var _fontMap pango.FontMapper // out

	_fontMap = (externglib.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(pango.FontMapper)

	return _fontMap
}
