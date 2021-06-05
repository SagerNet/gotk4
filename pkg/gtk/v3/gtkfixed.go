// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_fixed_get_type()), F: marshalFixed},
	})
}

// Fixed: the Fixed widget is a container which can place child widgets at fixed
// positions and with fixed sizes, given in pixels. Fixed performs no automatic
// layout management.
//
// For most applications, you should not use this container! It keeps you from
// having to learn about the other GTK+ containers, but it results in broken
// applications. With Fixed, the following things will result in truncated text,
// overlapping widgets, and other display bugs:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, Fixed does not pay attention to text direction and thus may
// produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK+ will order containers
// appropriately for the text direction, e.g. to put labels to the right of the
// thing they label when using an RTL language, but it can’t do that with Fixed.
// So if you need to reorder widgets depending on the text direction, you would
// need to manually detect it and adjust child positions accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove GUI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application, and
// prefer the simplicity of Fixed, by all means use the widget. But you should
// be aware of the tradeoffs.
//
// See also Layout, which shares the ability to perform fixed positioning of
// child widgets and additionally adds custom drawing and scrollability.
type Fixed interface {
	Container
	Buildable

	// Move moves a child of a Fixed container to the given position.
	Move(widget Widget, x int, y int)
	// Put adds a widget to a Fixed container at the given position.
	Put(widget Widget, x int, y int)
}

// fixed implements the Fixed interface.
type fixed struct {
	Container
	Buildable
}

var _ Fixed = (*fixed)(nil)

// WrapFixed wraps a GObject to the right type. It is
// primarily used internally.
func WrapFixed(obj *externglib.Object) Fixed {
	return Fixed{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalFixed(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFixed(obj), nil
}

// NewFixed constructs a class Fixed.
func NewFixed() Fixed {
	var cret C.GtkFixed
	var goret1 Fixed

	cret = C.gtk_fixed_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Fixed)

	return goret1
}

// Move moves a child of a Fixed container to the given position.
func (f fixed) Move(widget Widget, x int, y int) {
	var arg0 *C.GtkFixed
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 C.gint

	arg0 = (*C.GtkFixed)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)

	C.gtk_fixed_move(arg0, widget, x, y)
}

// Put adds a widget to a Fixed container at the given position.
func (f fixed) Put(widget Widget, x int, y int) {
	var arg0 *C.GtkFixed
	var arg1 *C.GtkWidget
	var arg2 C.gint
	var arg3 C.gint

	arg0 = (*C.GtkFixed)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.gint(x)
	arg3 = C.gint(y)

	C.gtk_fixed_put(arg0, widget, x, y)
}

type FixedChild struct {
	native C.GtkFixedChild
}

// WrapFixedChild wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFixedChild(ptr unsafe.Pointer) *FixedChild {
	if ptr == nil {
		return nil
	}

	return (*FixedChild)(ptr)
}

func marshalFixedChild(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFixedChild(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FixedChild) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// Widget gets the field inside the struct.
func (f *FixedChild) Widget() Widget {
	v = gextras.CastObject(externglib.Take(unsafe.Pointer(f.native.widget.Native()))).(Widget)
}

// X gets the field inside the struct.
func (f *FixedChild) X() int {
	v = C.gint(f.native.x)
}

// Y gets the field inside the struct.
func (f *FixedChild) Y() int {
	v = C.gint(f.native.y)
}

type FixedPrivate struct {
	native C.GtkFixedPrivate
}

// WrapFixedPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFixedPrivate(ptr unsafe.Pointer) *FixedPrivate {
	if ptr == nil {
		return nil
	}

	return (*FixedPrivate)(ptr)
}

func marshalFixedPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFixedPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FixedPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}
