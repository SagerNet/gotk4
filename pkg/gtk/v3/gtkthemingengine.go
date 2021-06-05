// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_theming_engine_get_type()), F: marshalThemingEngine},
	})
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine interface {
	gextras.Objector

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
	// Screen returns the Screen to which @engine currently rendering to.
	Screen() gdk.Screen
	// State returns the state used when rendering.
	State() StateFlags
	// StyleProperty gets the value for a widget style property.
	StyleProperty(propertyName string) externglib.Value
	// HasClass returns true if the currently rendered contents have defined the
	// given class name.
	HasClass(styleClass string) bool
	// HasRegion returns true if the currently rendered contents have the region
	// defined. If @flags_return is not nil, it is set to the flags affecting
	// the region.
	HasRegion(styleRegion string) (flags RegionFlags, ok bool)
	// LookupColor looks up and resolves a color name in the current style’s
	// color map.
	LookupColor(colorName string) (color gdk.RGBA, ok bool)
	// StateIsRunning returns true if there is a transition animation running
	// for the current region (see gtk_style_context_push_animatable_region()).
	//
	// If @progress is not nil, the animation progress will be returned there,
	// 0.0 means the state is closest to being false, while 1.0 means it’s
	// closest to being true. This means transition animations will run from 0
	// to 1 when @state is being set to true and from 1 to 0 when it’s being set
	// to false.
	StateIsRunning(state StateType) (progress float64, ok bool)
}

// themingEngine implements the ThemingEngine interface.
type themingEngine struct {
	gextras.Objector
}

var _ ThemingEngine = (*themingEngine)(nil)

// WrapThemingEngine wraps a GObject to the right type. It is
// primarily used internally.
func WrapThemingEngine(obj *externglib.Object) ThemingEngine {
	return ThemingEngine{
		Objector: obj,
	}
}

func marshalThemingEngine(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapThemingEngine(obj), nil
}

// BackgroundColor gets the background color for a given state.
func (e themingEngine) BackgroundColor(state StateFlags) gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GdkRGBA
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_background_color(arg0, state, &arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Border gets the border for a given state as a Border.
func (e themingEngine) Border(state StateFlags) Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GtkBorder
	var ret2 *Border

	C.gtk_theming_engine_get_border(arg0, state, &arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// BorderColor gets the border color for a given state.
func (e themingEngine) BorderColor(state StateFlags) gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GdkRGBA
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_border_color(arg0, state, &arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Color gets the foreground color for a given state.
func (e themingEngine) Color(state StateFlags) gdk.RGBA {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GdkRGBA
	var ret2 *gdk.RGBA

	C.gtk_theming_engine_get_color(arg0, state, &arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))

	return ret2
}

// Direction returns the widget direction used for rendering.
func (e themingEngine) Direction() TextDirection {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkTextDirection
	var goret1 TextDirection

	cret = C.gtk_theming_engine_get_direction(arg0)

	goret1 = TextDirection(cret)

	return goret1
}

// Font returns the font description for a given state.
func (e themingEngine) Font(state StateFlags) *pango.FontDescription {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var cret *C.PangoFontDescription
	var goret1 *pango.FontDescription

	cret = C.gtk_theming_engine_get_font(arg0, state)

	goret1 = pango.WrapFontDescription(unsafe.Pointer(cret))

	return goret1
}

// JunctionSides returns the widget direction used for rendering.
func (e themingEngine) JunctionSides() JunctionSides {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkJunctionSides
	var goret1 JunctionSides

	cret = C.gtk_theming_engine_get_junction_sides(arg0)

	goret1 = JunctionSides(cret)

	return goret1
}

// Margin gets the margin for a given state as a Border.
func (e themingEngine) Margin(state StateFlags) Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GtkBorder
	var ret2 *Border

	C.gtk_theming_engine_get_margin(arg0, state, &arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// Padding gets the padding for a given state as a Border.
func (e themingEngine) Padding(state StateFlags) Border {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateFlags)(state)

	var arg2 *C.GtkBorder
	var ret2 *Border

	C.gtk_theming_engine_get_padding(arg0, state, &arg2)

	ret2 = WrapBorder(unsafe.Pointer(arg2))

	return ret2
}

// Path returns the widget path used for style matching.
func (e themingEngine) Path() *WidgetPath {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret *C.GtkWidgetPath
	var goret1 *WidgetPath

	cret = C.gtk_theming_engine_get_path(arg0)

	goret1 = WrapWidgetPath(unsafe.Pointer(cret))

	return goret1
}

// Property gets a property value as retrieved from the style settings that
// apply to the currently rendered element.
func (e themingEngine) Property(property string, state StateFlags) externglib.Value {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar
	var arg2 C.GtkStateFlags

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GtkStateFlags)(state)

	var arg3 *C.GValue
	var ret3 *externglib.Value

	C.gtk_theming_engine_get_property(arg0, property, state, &arg3)

	ret3 = externglib.ValueFromNative(unsafe.Pointer(arg3))
	runtime.SetFinalizer(ret3, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(v.GValue))
	})

	return ret3
}

// Screen returns the Screen to which @engine currently rendering to.
func (e themingEngine) Screen() gdk.Screen {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret *C.GdkScreen
	var goret1 gdk.Screen

	cret = C.gtk_theming_engine_get_screen(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Screen)

	return goret1
}

// State returns the state used when rendering.
func (e themingEngine) State() StateFlags {
	var arg0 *C.GtkThemingEngine

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))

	var cret C.GtkStateFlags
	var goret1 StateFlags

	cret = C.gtk_theming_engine_get_state(arg0)

	goret1 = StateFlags(cret)

	return goret1
}

// StyleProperty gets the value for a widget style property.
func (e themingEngine) StyleProperty(propertyName string) externglib.Value {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.GValue
	var ret2 *externglib.Value

	C.gtk_theming_engine_get_style_property(arg0, propertyName, &arg2)

	ret2 = externglib.ValueFromNative(unsafe.Pointer(arg2))

	return ret2
}

// HasClass returns true if the currently rendered contents have defined the
// given class name.
func (e themingEngine) HasClass(styleClass string) bool {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(styleClass))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_theming_engine_has_class(arg0, styleClass)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// HasRegion returns true if the currently rendered contents have the region
// defined. If @flags_return is not nil, it is set to the flags affecting
// the region.
func (e themingEngine) HasRegion(styleRegion string) (flags RegionFlags, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(styleRegion))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.GtkRegionFlags
	var ret2 *RegionFlags
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_theming_engine_has_region(arg0, styleRegion, &arg2)

	ret2 = *RegionFlags(arg2)
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

// LookupColor looks up and resolves a color name in the current style’s
// color map.
func (e themingEngine) LookupColor(colorName string) (color gdk.RGBA, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 *C.gchar

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.GdkRGBA
	var ret2 *gdk.RGBA
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_theming_engine_lookup_color(arg0, colorName, &arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

// StateIsRunning returns true if there is a transition animation running
// for the current region (see gtk_style_context_push_animatable_region()).
//
// If @progress is not nil, the animation progress will be returned there,
// 0.0 means the state is closest to being false, while 1.0 means it’s
// closest to being true. This means transition animations will run from 0
// to 1 when @state is being set to true and from 1 to 0 when it’s being set
// to false.
func (e themingEngine) StateIsRunning(state StateType) (progress float64, ok bool) {
	var arg0 *C.GtkThemingEngine
	var arg1 C.GtkStateType

	arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(e.Native()))
	arg1 = (C.GtkStateType)(state)

	var arg2 *C.gdouble
	var ret2 float64
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_theming_engine_state_is_running(arg0, state, &arg2)

	ret2 = *C.gdouble(arg2)
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

type ThemingEnginePrivate struct {
	native C.GtkThemingEnginePrivate
}

// WrapThemingEnginePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapThemingEnginePrivate(ptr unsafe.Pointer) *ThemingEnginePrivate {
	if ptr == nil {
		return nil
	}

	return (*ThemingEnginePrivate)(ptr)
}

func marshalThemingEnginePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapThemingEnginePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ThemingEnginePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
