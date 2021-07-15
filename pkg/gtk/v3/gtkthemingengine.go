// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/gotk3/gotk3/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_theming_engine_get_type()), F: marshalThemingEnginer},
	})
}

// ThemingEngineOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ThemingEngineOverrider interface {
	RenderActivity(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderArrow(cr *cairo.Context, angle float64, x float64, y float64, size float64)
	RenderBackground(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderCheck(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderExpander(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderExtension(cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType)
	RenderFocus(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderFrame(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderFrameGap(cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64)
	RenderHandle(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x float64, y float64)
	RenderIconSurface(cr *cairo.Context, surface *cairo.Surface, x float64, y float64)
	RenderLayout(cr *cairo.Context, x float64, y float64, layout *pango.Layout)
	RenderLine(cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64)
	RenderOption(cr *cairo.Context, x float64, y float64, width float64, height float64)
	RenderSlider(cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation)
}

// ThemingEnginer describes ThemingEngine's methods.
type ThemingEnginer interface {
	// BackgroundColor gets the background color for a given state.
	BackgroundColor(state StateFlags) gdk.RGBA
	// Border gets the border for a given state as a Border.
	Border(state StateFlags) Border
	// BorderColor gets the border color for a given state.
	BorderColor(state StateFlags) gdk.RGBA
	// Color gets the foreground color for a given state.
	Color(state StateFlags) gdk.RGBA
	// Direction returns the widget direction used for rendering.
	Direction() TextDirection
	// Font returns the font description for a given state.
	Font(state StateFlags) *pango.FontDescription
	// JunctionSides returns the widget direction used for rendering.
	JunctionSides() JunctionSides
	// Margin gets the margin for a given state as a Border.
	Margin(state StateFlags) Border
	// Padding gets the padding for a given state as a Border.
	Padding(state StateFlags) Border
	// Path returns the widget path used for style matching.
	Path() *WidgetPath
	// Property gets a property value as retrieved from the style settings that
	// apply to the currently rendered element.
	Property(property string, state StateFlags) externglib.Value
	// Screen returns the Screen to which engine currently rendering to.
	Screen() *gdk.Screen
	// State returns the state used when rendering.
	State() StateFlags
	// StyleProperty gets the value for a widget style property.
	StyleProperty(propertyName string) externglib.Value
	// HasClass returns TRUE if the currently rendered contents have defined the
	// given class name.
	HasClass(styleClass string) bool
	// HasRegion returns TRUE if the currently rendered contents have the region
	// defined.
	HasRegion(styleRegion string) (RegionFlags, bool)
	// LookupColor looks up and resolves a color name in the current style’s
	// color map.
	LookupColor(colorName string) (gdk.RGBA, bool)
	// StateIsRunning returns TRUE if there is a transition animation running
	// for the current region (see gtk_style_context_push_animatable_region()).
	StateIsRunning(state StateType) (float64, bool)
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine struct {
	*externglib.Object
}

var (
	_ ThemingEnginer  = (*ThemingEngine)(nil)
	_ gextras.Nativer = (*ThemingEngine)(nil)
)

func wrapThemingEngine(obj *externglib.Object) *ThemingEngine {
	return &ThemingEngine{
		Object: obj,
	}
}

func marshalThemingEnginer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapThemingEngine(obj), nil
}

// BackgroundColor gets the background color for a given state.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) BackgroundColor(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _color gdk.RGBA

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_background_color(_arg0, _arg1, (*C.GdkRGBA)(unsafe.Pointer(&_color)))

	return _color
}

// Border gets the border for a given state as a Border.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Border(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _border Border

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border(_arg0, _arg1, (*C.GtkBorder)(unsafe.Pointer(&_border)))

	return _border
}

// BorderColor gets the border color for a given state.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) BorderColor(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _color gdk.RGBA

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border_color(_arg0, _arg1, (*C.GdkRGBA)(unsafe.Pointer(&_color)))

	return _color
}

// Color gets the foreground color for a given state.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Color(state StateFlags) gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _color gdk.RGBA

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_color(_arg0, _arg1, (*C.GdkRGBA)(unsafe.Pointer(&_color)))

	return _color
}

// Direction returns the widget direction used for rendering.
//
// Deprecated: Use gtk_theming_engine_get_state() and check for
// K_STATE_FLAG_DIR_LTR and K_STATE_FLAG_DIR_RTL instead.
func (engine *ThemingEngine) Direction() TextDirection {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkTextDirection  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_direction(_arg0)

	var _textDirection TextDirection // out

	_textDirection = TextDirection(_cret)

	return _textDirection
}

// Font returns the font description for a given state.
//
// Deprecated: Use gtk_theming_engine_get().
func (engine *ThemingEngine) Font(state StateFlags) *pango.FontDescription {
	var _arg0 *C.GtkThemingEngine     // out
	var _arg1 C.GtkStateFlags         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	_cret = C.gtk_theming_engine_get_font(_arg0, _arg1)

	var _fontDescription *pango.FontDescription // out

	_fontDescription = (*pango.FontDescription)(unsafe.Pointer(_cret))

	return _fontDescription
}

// JunctionSides returns the widget direction used for rendering.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) JunctionSides() JunctionSides {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkJunctionSides  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_junction_sides(_arg0)

	var _junctionSides JunctionSides // out

	_junctionSides = JunctionSides(_cret)

	return _junctionSides
}

// Margin gets the margin for a given state as a Border.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Margin(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _margin Border

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_margin(_arg0, _arg1, (*C.GtkBorder)(unsafe.Pointer(&_margin)))

	return _margin
}

// Padding gets the padding for a given state as a Border.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Padding(state StateFlags) Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _padding Border

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_padding(_arg0, _arg1, (*C.GtkBorder)(unsafe.Pointer(&_padding)))

	return _padding
}

// Path returns the widget path used for style matching.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Path() *WidgetPath {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GtkWidgetPath    // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_path(_arg0)

	var _widgetPath *WidgetPath // out

	_widgetPath = (*WidgetPath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_widgetPath, func(v *WidgetPath) {
		C.gtk_widget_path_unref((*C.GtkWidgetPath)(unsafe.Pointer(v)))
	})

	return _widgetPath
}

// Property gets a property value as retrieved from the style settings that
// apply to the currently rendered element.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Property(property string, state StateFlags) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkStateFlags     // out
	var _arg3 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(property)))
	_arg2 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_property(_arg0, _arg1, _arg2, &_arg3)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))
	runtime.SetFinalizer(_value, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(unsafe.Pointer(v.GValue)))
	})

	return _value
}

// Screen returns the Screen to which engine currently rendering to.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) Screen() *gdk.Screen {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GdkScreen        // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_screen(_arg0)

	var _screen *gdk.Screen // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_screen = &gdk.Screen{
			Object: obj,
		}
	}

	return _screen
}

// State returns the state used when rendering.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) State() StateFlags {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkStateFlags     // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_state(_arg0)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// StyleProperty gets the value for a widget style property.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) StyleProperty(propertyName string) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))

	C.gtk_theming_engine_get_style_property(_arg0, _arg1, &_arg2)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg2)))

	return _value
}

// HasClass returns TRUE if the currently rendered contents have defined the
// given class name.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) HasClass(styleClass string) bool {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleClass)))

	_cret = C.gtk_theming_engine_has_class(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasRegion returns TRUE if the currently rendered contents have the region
// defined. If flags_return is not NULL, it is set to the flags affecting the
// region.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) HasRegion(styleRegion string) (RegionFlags, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkRegionFlags    // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleRegion)))

	_cret = C.gtk_theming_engine_has_region(_arg0, _arg1, &_arg2)

	var _flags RegionFlags // out
	var _ok bool           // out

	_flags = RegionFlags(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _flags, _ok
}

// LookupColor looks up and resolves a color name in the current style’s color
// map.
//
// Deprecated: since version 3.14.
func (engine *ThemingEngine) LookupColor(colorName string) (gdk.RGBA, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _color gdk.RGBA
	var _cret C.gboolean // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(colorName)))

	_cret = C.gtk_theming_engine_lookup_color(_arg0, _arg1, (*C.GdkRGBA)(unsafe.Pointer(&_color)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// StateIsRunning returns TRUE if there is a transition animation running for
// the current region (see gtk_style_context_push_animatable_region()).
//
// If progress is not NULL, the animation progress will be returned there, 0.0
// means the state is closest to being FALSE, while 1.0 means it’s closest to
// being TRUE. This means transition animations will run from 0 to 1 when state
// is being set to TRUE and from 1 to 0 when it’s being set to FALSE.
//
// Deprecated: Always returns FALSE.
func (engine *ThemingEngine) StateIsRunning(state StateType) (float64, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateType      // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateType(state)

	_cret = C.gtk_theming_engine_state_is_running(_arg0, _arg1, &_arg2)

	var _progress float64 // out
	var _ok bool          // out

	_progress = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _progress, _ok
}

// ThemingEngineLoad loads and initializes a theming engine module from the
// standard directories.
//
// Deprecated: since version 3.14.
func ThemingEngineLoad(name string) *ThemingEngine {
	var _arg1 *C.gchar            // out
	var _cret *C.GtkThemingEngine // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_theming_engine_load(_arg1)

	var _themingEngine *ThemingEngine // out

	_themingEngine = wrapThemingEngine(externglib.Take(unsafe.Pointer(_cret)))

	return _themingEngine
}
