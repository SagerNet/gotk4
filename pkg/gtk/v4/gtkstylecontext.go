// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_style_context_get_type()), F: marshalStyleContext},
	})
}

// StyleContext is an object that stores styling information affecting a widget.
//
// In order to construct the final style information, StyleContext queries
// information from all attached StyleProviders. Style providers can be either
// attached explicitly to the context through gtk_style_context_add_provider(),
// or to the display through gtk_style_context_add_provider_for_display(). The
// resulting style is a combination of all providers’ information in priority
// order.
//
// For GTK widgets, any StyleContext returned by gtk_widget_get_style_context()
// will already have a Display and RTL/LTR information set. The style context
// will also be updated automatically if any of these settings change on the
// widget.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
// GTK defines macros for a number of style classes.
//
//
// Custom styling in UI libraries and applications
//
// If you are developing a library with custom Widgets that render differently
// than standard components, you may need to add a StyleProvider yourself with
// the GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority, either a CssProvider or a
// custom object implementing the StyleProvider interface. This way themes may
// still attempt to style your UI elements in a different way if needed so.
//
// If you are using custom styling on an applications, you probably want then to
// make your style information prevail to the theme’s, so you must use a
// StyleProvider with the GTK_STYLE_PROVIDER_PRIORITY_APPLICATION priority, keep
// in mind that the user settings in `XDG_CONFIG_HOME/gtk-4.0/gtk.css` will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext interface {
	gextras.Objector

	// AddClass adds a style class to @context, so later uses of the style
	// context will make use of this new class for styling.
	//
	// In the CSS file format, a Entry defining a “search” class, would be
	// matched by:
	//
	// |[ <!-- language="CSS" --> entry.search { ... } ]|
	//
	// While any widget defining a “search” class would be matched by: |[ <!--
	// language="CSS" --> .search { ... } ]|
	AddClass(className string)
	// AddProvider adds a style provider to @context, to be used in style
	// construction. Note that a style provider added by this function only
	// affects the style of the widget to which @context belongs. If you want to
	// affect the style of all widgets, use
	// gtk_style_context_add_provider_for_display().
	//
	// Note: If both priorities are the same, a StyleProvider added through this
	// function takes precedence over another added through
	// gtk_style_context_add_provider_for_display().
	AddProvider(provider StyleProvider, priority uint)
	// Border gets the border for a given state as a Border.
	Border() Border
	// Color gets the foreground color for a given state.
	Color() gdk.RGBA
	// Display returns the Display to which @context is attached.
	Display() gdk.Display
	// Margin gets the margin for a given state as a Border.
	Margin() Border
	// Padding gets the padding for a given state as a Border.
	Padding() Border
	// Scale returns the scale used for assets.
	Scale() int
	// State returns the state used for style matching.
	//
	// This method should only be used to retrieve the StateFlags to pass to
	// StyleContext methods, like gtk_style_context_get_padding(). If you need
	// to retrieve the current state of a Widget, use
	// gtk_widget_get_state_flags().
	State() StateFlags
	// HasClass returns true if @context currently has defined the given class
	// name.
	HasClass(className string) bool
	// LookupColor looks up and resolves a color name in the @context color map.
	LookupColor(colorName string) (color gdk.RGBA, ok bool)
	// RemoveClass removes @class_name from @context.
	RemoveClass(className string)
	// RemoveProvider removes @provider from the style providers list in
	// @context.
	RemoveProvider(provider StyleProvider)
	// Restore restores @context state to a previous stage. See
	// gtk_style_context_save().
	Restore()
	// Save saves the @context state, so temporary modifications done through
	// gtk_style_context_add_class(), gtk_style_context_remove_class(),
	// gtk_style_context_set_state(), etc. can quickly be reverted in one go
	// through gtk_style_context_restore().
	//
	// The matching call to gtk_style_context_restore() must be done before GTK
	// returns to the main loop.
	Save()
	// SetDisplay attaches @context to the given display.
	//
	// The display is used to add style information from “global” style
	// providers, such as the display's Settings instance.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), you do not need to call this yourself.
	SetDisplay(display gdk.Display)
	// SetScale sets the scale to use when getting image assets for the style.
	SetScale(scale int)
	// SetState sets the state to be used for style matching.
	SetState(flags StateFlags)
	// String converts the style context into a string representation.
	//
	// The string representation always includes information about the name,
	// state, id, visibility and style classes of the CSS node that is backing
	// @context. Depending on the flags, more information may be included.
	//
	// This function is intended for testing and debugging of the CSS
	// implementation in GTK. There are no guarantees about the format of the
	// returned string, it may change.
	String(flags StyleContextPrintFlags) string
}

// styleContext implements the StyleContext interface.
type styleContext struct {
	gextras.Objector
}

var _ StyleContext = (*styleContext)(nil)

// WrapStyleContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapStyleContext(obj *externglib.Object) StyleContext {
	return StyleContext{
		Objector: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleContext(obj), nil
}

// AddClass adds a style class to @context, so later uses of the style
// context will make use of this new class for styling.
//
// In the CSS file format, a Entry defining a “search” class, would be
// matched by:
//
// |[ <!-- language="CSS" --> entry.search { ... } ]|
//
// While any widget defining a “search” class would be matched by: |[ <!--
// language="CSS" --> .search { ... } ]|
func (c styleContext) AddClass(className string) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_style_context_add_class(arg0, className)
}

// AddProvider adds a style provider to @context, to be used in style
// construction. Note that a style provider added by this function only
// affects the style of the widget to which @context belongs. If you want to
// affect the style of all widgets, use
// gtk_style_context_add_provider_for_display().
//
// Note: If both priorities are the same, a StyleProvider added through this
// function takes precedence over another added through
// gtk_style_context_add_provider_for_display().
func (c styleContext) AddProvider(provider StyleProvider, priority uint) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GtkStyleProvider
	var arg2 C.guint

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))
	arg2 = C.guint(priority)

	C.gtk_style_context_add_provider(arg0, provider, priority)
}

// Border gets the border for a given state as a Border.
func (c styleContext) Border() Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.GtkBorder
	var ret1 *Border

	C.gtk_style_context_get_border(arg0, &arg1)

	ret1 = WrapBorder(unsafe.Pointer(arg1))

	return ret1
}

// Color gets the foreground color for a given state.
func (c styleContext) Color() gdk.RGBA {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.GdkRGBA
	var ret1 *gdk.RGBA

	C.gtk_style_context_get_color(arg0, &arg1)

	ret1 = gdk.WrapRGBA(unsafe.Pointer(arg1))

	return ret1
}

// Display returns the Display to which @context is attached.
func (c styleContext) Display() gdk.Display {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var cret *C.GdkDisplay
	var goret1 gdk.Display

	cret = C.gtk_style_context_get_display(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Display)

	return goret1
}

// Margin gets the margin for a given state as a Border.
func (c styleContext) Margin() Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.GtkBorder
	var ret1 *Border

	C.gtk_style_context_get_margin(arg0, &arg1)

	ret1 = WrapBorder(unsafe.Pointer(arg1))

	return ret1
}

// Padding gets the padding for a given state as a Border.
func (c styleContext) Padding() Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 *C.GtkBorder
	var ret1 *Border

	C.gtk_style_context_get_padding(arg0, &arg1)

	ret1 = WrapBorder(unsafe.Pointer(arg1))

	return ret1
}

// Scale returns the scale used for assets.
func (c styleContext) Scale() int {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var cret C.int
	var goret1 int

	cret = C.gtk_style_context_get_scale(arg0)

	goret1 = C.int(cret)

	return goret1
}

// State returns the state used for style matching.
//
// This method should only be used to retrieve the StateFlags to pass to
// StyleContext methods, like gtk_style_context_get_padding(). If you need
// to retrieve the current state of a Widget, use
// gtk_widget_get_state_flags().
func (c styleContext) State() StateFlags {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var cret C.GtkStateFlags
	var goret1 StateFlags

	cret = C.gtk_style_context_get_state(arg0)

	goret1 = StateFlags(cret)

	return goret1
}

// HasClass returns true if @context currently has defined the given class
// name.
func (c styleContext) HasClass(className string) bool {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_style_context_has_class(arg0, className)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// LookupColor looks up and resolves a color name in the @context color map.
func (c styleContext) LookupColor(colorName string) (color gdk.RGBA, ok bool) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(colorName))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 *C.GdkRGBA
	var ret2 *gdk.RGBA
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_style_context_lookup_color(arg0, colorName, &arg2)

	ret2 = gdk.WrapRGBA(unsafe.Pointer(arg2))
	goret2 = C.bool(cret) != C.false

	return ret2, goret2
}

// RemoveClass removes @class_name from @context.
func (c styleContext) RemoveClass(className string) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_style_context_remove_class(arg0, className)
}

// RemoveProvider removes @provider from the style providers list in
// @context.
func (c styleContext) RemoveProvider(provider StyleProvider) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GtkStyleProvider

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))

	C.gtk_style_context_remove_provider(arg0, provider)
}

// Restore restores @context state to a previous stage. See
// gtk_style_context_save().
func (c styleContext) Restore() {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_restore(arg0)
}

// Save saves the @context state, so temporary modifications done through
// gtk_style_context_add_class(), gtk_style_context_remove_class(),
// gtk_style_context_set_state(), etc. can quickly be reverted in one go
// through gtk_style_context_restore().
//
// The matching call to gtk_style_context_restore() must be done before GTK
// returns to the main loop.
func (c styleContext) Save() {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_save(arg0)
}

// SetDisplay attaches @context to the given display.
//
// The display is used to add style information from “global” style
// providers, such as the display's Settings instance.
//
// If you are using a StyleContext returned from
// gtk_widget_get_style_context(), you do not need to call this yourself.
func (c styleContext) SetDisplay(display gdk.Display) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GdkDisplay

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_style_context_set_display(arg0, display)
}

// SetScale sets the scale to use when getting image assets for the style.
func (c styleContext) SetScale(scale int) {
	var arg0 *C.GtkStyleContext
	var arg1 C.int

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(scale)

	C.gtk_style_context_set_scale(arg0, scale)
}

// SetState sets the state to be used for style matching.
func (c styleContext) SetState(flags StateFlags) {
	var arg0 *C.GtkStyleContext
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkStateFlags)(flags)

	C.gtk_style_context_set_state(arg0, flags)
}

// String converts the style context into a string representation.
//
// The string representation always includes information about the name,
// state, id, visibility and style classes of the CSS node that is backing
// @context. Depending on the flags, more information may be included.
//
// This function is intended for testing and debugging of the CSS
// implementation in GTK. There are no guarantees about the format of the
// returned string, it may change.
func (c styleContext) String(flags StyleContextPrintFlags) string {
	var arg0 *C.GtkStyleContext
	var arg1 C.GtkStyleContextPrintFlags

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkStyleContextPrintFlags)(flags)

	var cret *C.char
	var goret1 string

	cret = C.gtk_style_context_to_string(arg0, flags)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return goret1
}
