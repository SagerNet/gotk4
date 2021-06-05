// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_stack_get_type()), F: marshalStack},
	})
}

// Stack: the GtkStack widget is a container which only shows one of its
// children at a time. In contrast to GtkNotebook, GtkStack does not provide a
// means for users to change the visible child. Instead, the StackSwitcher
// widget can be used with GtkStack to provide this functionality.
//
// Transitions between pages can be animated as slides or fades. This can be
// controlled with gtk_stack_set_transition_type(). These animations respect the
// Settings:gtk-enable-animations setting.
//
// The GtkStack widget was added in GTK+ 3.10.
//
//
// CSS nodes
//
// GtkStack has a single CSS node named stack.
type Stack interface {
	Container
	Buildable

	// AddNamed adds a child to @stack. The child is identified by the @name.
	AddNamed(child Widget, name string)
	// AddTitled adds a child to @stack. The child is identified by the @name.
	// The @title will be used by StackSwitcher to represent @child in a tab
	// bar, so it should be short.
	AddTitled(child Widget, name string, title string)
	// ChildByName finds the child of the Stack with the name given as the
	// argument. Returns nil if there is no child with this name.
	ChildByName(name string) Widget
	// Hhomogeneous gets whether @stack is horizontally homogeneous. See
	// gtk_stack_set_hhomogeneous().
	Hhomogeneous() bool
	// Homogeneous gets whether @stack is homogeneous. See
	// gtk_stack_set_homogeneous().
	Homogeneous() bool
	// InterpolateSize returns wether the Stack is set up to interpolate between
	// the sizes of children on page switch.
	InterpolateSize() bool
	// TransitionDuration returns the amount of time (in milliseconds) that
	// transitions between pages in @stack will take.
	TransitionDuration() uint
	// TransitionRunning returns whether the @stack is currently in a transition
	// from one page to another.
	TransitionRunning() bool
	// TransitionType gets the type of animation that will be used for
	// transitions between pages in @stack.
	TransitionType() StackTransitionType
	// Vhomogeneous gets whether @stack is vertically homogeneous. See
	// gtk_stack_set_vhomogeneous().
	Vhomogeneous() bool
	// VisibleChild gets the currently visible child of @stack, or nil if there
	// are no visible children.
	VisibleChild() Widget
	// VisibleChildName returns the name of the currently visible child of
	// @stack, or nil if there is no visible child.
	VisibleChildName() string
	// SetHhomogeneous sets the Stack to be horizontally homogeneous or not. If
	// it is homogeneous, the Stack will request the same width for all its
	// children. If it isn't, the stack may change width when a different child
	// becomes visible.
	SetHhomogeneous(hhomogeneous bool)
	// SetHomogeneous sets the Stack to be homogeneous or not. If it is
	// homogeneous, the Stack will request the same size for all its children.
	// If it isn't, the stack may change size when a different child becomes
	// visible.
	//
	// Since 3.16, homogeneity can be controlled separately for horizontal and
	// vertical size, with the Stack:hhomogeneous and Stack:vhomogeneous.
	SetHomogeneous(homogeneous bool)
	// SetInterpolateSize sets whether or not @stack will interpolate its size
	// when changing the visible child. If the Stack:interpolate-size property
	// is set to true, @stack will interpolate its size between the current one
	// and the one it'll take after changing the visible child, according to the
	// set transition duration.
	SetInterpolateSize(interpolateSize bool)
	// SetTransitionDuration sets the duration that transitions between pages in
	// @stack will take.
	SetTransitionDuration(duration uint)
	// SetTransitionType sets the type of animation that will be used for
	// transitions between pages in @stack. Available types include various
	// kinds of fades and slides.
	//
	// The transition type can be changed without problems at runtime, so it is
	// possible to change the animation based on the page that is about to
	// become current.
	SetTransitionType(transition StackTransitionType)
	// SetVhomogeneous sets the Stack to be vertically homogeneous or not. If it
	// is homogeneous, the Stack will request the same height for all its
	// children. If it isn't, the stack may change height when a different child
	// becomes visible.
	SetVhomogeneous(vhomogeneous bool)
	// SetVisibleChild makes @child the visible child of @stack.
	//
	// If @child is different from the currently visible child, the transition
	// between the two will be animated with the current transition type of
	// @stack.
	//
	// Note that the @child widget has to be visible itself (see
	// gtk_widget_show()) in order to become the visible child of @stack.
	SetVisibleChild(child Widget)
	// SetVisibleChildFull makes the child with the given name visible.
	//
	// Note that the child widget has to be visible itself (see
	// gtk_widget_show()) in order to become the visible child of @stack.
	SetVisibleChildFull(name string, transition StackTransitionType)
	// SetVisibleChildName makes the child with the given name visible.
	//
	// If @child is different from the currently visible child, the transition
	// between the two will be animated with the current transition type of
	// @stack.
	//
	// Note that the child widget has to be visible itself (see
	// gtk_widget_show()) in order to become the visible child of @stack.
	SetVisibleChildName(name string)
}

// stack implements the Stack interface.
type stack struct {
	Container
	Buildable
}

var _ Stack = (*stack)(nil)

// WrapStack wraps a GObject to the right type. It is
// primarily used internally.
func WrapStack(obj *externglib.Object) Stack {
	return Stack{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalStack(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStack(obj), nil
}

// NewStack constructs a class Stack.
func NewStack() Stack {
	var cret C.GtkStack
	var ret1 Stack

	cret = C.gtk_stack_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Stack)

	return ret1
}

// AddNamed adds a child to @stack. The child is identified by the @name.
func (s stack) AddNamed(child Widget, name string) {
	var arg0 *C.GtkStack
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_stack_add_named(arg0, child, name)
}

// AddTitled adds a child to @stack. The child is identified by the @name.
// The @title will be used by StackSwitcher to represent @child in a tab
// bar, so it should be short.
func (s stack) AddTitled(child Widget, name string, title string) {
	var arg0 *C.GtkStack
	var arg1 *C.GtkWidget
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_stack_add_titled(arg0, child, name, title)
}

// ChildByName finds the child of the Stack with the name given as the
// argument. Returns nil if there is no child with this name.
func (s stack) ChildByName(name string) Widget {
	var arg0 *C.GtkStack
	var arg1 *C.gchar

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_stack_get_child_by_name(arg0, name)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Hhomogeneous gets whether @stack is horizontally homogeneous. See
// gtk_stack_set_hhomogeneous().
func (s stack) Hhomogeneous() bool {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_stack_get_hhomogeneous(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Homogeneous gets whether @stack is homogeneous. See
// gtk_stack_set_homogeneous().
func (s stack) Homogeneous() bool {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_stack_get_homogeneous(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// InterpolateSize returns wether the Stack is set up to interpolate between
// the sizes of children on page switch.
func (s stack) InterpolateSize() bool {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_stack_get_interpolate_size(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in @stack will take.
func (s stack) TransitionDuration() uint {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gtk_stack_get_transition_duration(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// TransitionRunning returns whether the @stack is currently in a transition
// from one page to another.
func (s stack) TransitionRunning() bool {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_stack_get_transition_running(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// TransitionType gets the type of animation that will be used for
// transitions between pages in @stack.
func (s stack) TransitionType() StackTransitionType {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.GtkStackTransitionType
	var ret1 StackTransitionType

	cret = C.gtk_stack_get_transition_type(arg0)

	ret1 = StackTransitionType(cret)

	return ret1
}

// Vhomogeneous gets whether @stack is vertically homogeneous. See
// gtk_stack_set_vhomogeneous().
func (s stack) Vhomogeneous() bool {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_stack_get_vhomogeneous(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// VisibleChild gets the currently visible child of @stack, or nil if there
// are no visible children.
func (s stack) VisibleChild() Widget {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_stack_get_visible_child(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// VisibleChildName returns the name of the currently visible child of
// @stack, or nil if there is no visible child.
func (s stack) VisibleChildName() string {
	var arg0 *C.GtkStack

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_stack_get_visible_child_name(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// SetHhomogeneous sets the Stack to be horizontally homogeneous or not. If
// it is homogeneous, the Stack will request the same width for all its
// children. If it isn't, the stack may change width when a different child
// becomes visible.
func (s stack) SetHhomogeneous(hhomogeneous bool) {
	var arg0 *C.GtkStack
	var arg1 C.gboolean

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if hhomogeneous {
		arg1 = C.gboolean(1)
	}

	C.gtk_stack_set_hhomogeneous(arg0, hhomogeneous)
}

// SetHomogeneous sets the Stack to be homogeneous or not. If it is
// homogeneous, the Stack will request the same size for all its children.
// If it isn't, the stack may change size when a different child becomes
// visible.
//
// Since 3.16, homogeneity can be controlled separately for horizontal and
// vertical size, with the Stack:hhomogeneous and Stack:vhomogeneous.
func (s stack) SetHomogeneous(homogeneous bool) {
	var arg0 *C.GtkStack
	var arg1 C.gboolean

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if homogeneous {
		arg1 = C.gboolean(1)
	}

	C.gtk_stack_set_homogeneous(arg0, homogeneous)
}

// SetInterpolateSize sets whether or not @stack will interpolate its size
// when changing the visible child. If the Stack:interpolate-size property
// is set to true, @stack will interpolate its size between the current one
// and the one it'll take after changing the visible child, according to the
// set transition duration.
func (s stack) SetInterpolateSize(interpolateSize bool) {
	var arg0 *C.GtkStack
	var arg1 C.gboolean

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if interpolateSize {
		arg1 = C.gboolean(1)
	}

	C.gtk_stack_set_interpolate_size(arg0, interpolateSize)
}

// SetTransitionDuration sets the duration that transitions between pages in
// @stack will take.
func (s stack) SetTransitionDuration(duration uint) {
	var arg0 *C.GtkStack
	var arg1 C.guint

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(duration)

	C.gtk_stack_set_transition_duration(arg0, duration)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between pages in @stack. Available types include various
// kinds of fades and slides.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the page that is about to
// become current.
func (s stack) SetTransitionType(transition StackTransitionType) {
	var arg0 *C.GtkStack
	var arg1 C.GtkStackTransitionType

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type(arg0, transition)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not. If it
// is homogeneous, the Stack will request the same height for all its
// children. If it isn't, the stack may change height when a different child
// becomes visible.
func (s stack) SetVhomogeneous(vhomogeneous bool) {
	var arg0 *C.GtkStack
	var arg1 C.gboolean

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	if vhomogeneous {
		arg1 = C.gboolean(1)
	}

	C.gtk_stack_set_vhomogeneous(arg0, vhomogeneous)
}

// SetVisibleChild makes @child the visible child of @stack.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of
// @stack.
//
// Note that the @child widget has to be visible itself (see
// gtk_widget_show()) in order to become the visible child of @stack.
func (s stack) SetVisibleChild(child Widget) {
	var arg0 *C.GtkStack
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_stack_set_visible_child(arg0, child)
}

// SetVisibleChildFull makes the child with the given name visible.
//
// Note that the child widget has to be visible itself (see
// gtk_widget_show()) in order to become the visible child of @stack.
func (s stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	var arg0 *C.GtkStack
	var arg1 *C.gchar
	var arg2 C.GtkStackTransitionType

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full(arg0, name, transition)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If @child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of
// @stack.
//
// Note that the child widget has to be visible itself (see
// gtk_widget_show()) in order to become the visible child of @stack.
func (s stack) SetVisibleChildName(name string) {
	var arg0 *C.GtkStack
	var arg1 *C.gchar

	arg0 = (*C.GtkStack)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_stack_set_visible_child_name(arg0, name)
}