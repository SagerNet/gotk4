// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_scrollable_get_type()), F: marshalScrollabler},
	})
}

// ScrollableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScrollableOverrider interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	//
	// An example for this would be treeview headers. GTK can use this
	// information to display overlaid graphics, like the overshoot indication,
	// at the right position.
	Border() (Border, bool)
}

// Scrollable: GtkScrollable is an interface for widgets with native scrolling
// ability.
//
// To implement this interface you should override the
// gtk.Scrollable:hadjustment and gtk.Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ gtk.Adjustment:lower,
// gtk.Adjustment:upper, gtk.Adjustment:step-increment,
// gtk.Adjustment:page-increment and gtk.Adjustment:page-size properties and
// connect to the gtk.Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its GtkWidgetClass.size_allocate()
// function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the gtk.Adjustment::value-changed signal,
// the scrollable widget should scroll its contents.
type Scrollable struct {
	*externglib.Object
}

// Scrollabler describes Scrollable's abstract methods.
type Scrollabler interface {
	externglib.Objector

	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable.
	Border() (Border, bool)
	// HAdjustment retrieves the GtkAdjustment used for horizontal scrolling.
	HAdjustment() *Adjustment
	// HScrollPolicy gets the horizontal GtkScrollablePolicy.
	HScrollPolicy() ScrollablePolicy
	// VAdjustment retrieves the GtkAdjustment used for vertical scrolling.
	VAdjustment() *Adjustment
	// VScrollPolicy gets the vertical GtkScrollablePolicy.
	VScrollPolicy() ScrollablePolicy
	// SetHAdjustment sets the horizontal adjustment of the GtkScrollable.
	SetHAdjustment(hadjustment *Adjustment)
	// SetHScrollPolicy sets the GtkScrollablePolicy.
	SetHScrollPolicy(policy ScrollablePolicy)
	// SetVAdjustment sets the vertical adjustment of the GtkScrollable.
	SetVAdjustment(vadjustment *Adjustment)
	// SetVScrollPolicy sets the GtkScrollablePolicy.
	SetVScrollPolicy(policy ScrollablePolicy)
}

var _ Scrollabler = (*Scrollable)(nil)

func wrapScrollable(obj *externglib.Object) *Scrollable {
	return &Scrollable{
		Object: obj,
	}
}

func marshalScrollabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScrollable(obj), nil
}

// Border returns the size of a non-scrolling border around the outside of the
// scrollable.
//
// An example for this would be treeview headers. GTK can use this information
// to display overlaid graphics, like the overshoot indication, at the right
// position.
func (scrollable *Scrollable) Border() (Border, bool) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 C.GtkBorder      // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_border(_arg0, &_arg1)
	runtime.KeepAlive(scrollable)

	var _border Border // out
	var _ok bool       // out

	_border = *(*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _border, _ok
}

// HAdjustment retrieves the GtkAdjustment used for horizontal scrolling.
func (scrollable *Scrollable) HAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_hadjustment(_arg0)
	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// HScrollPolicy gets the horizontal GtkScrollablePolicy.
func (scrollable *Scrollable) HScrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_hscroll_policy(_arg0)
	runtime.KeepAlive(scrollable)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// VAdjustment retrieves the GtkAdjustment used for vertical scrolling.
func (scrollable *Scrollable) VAdjustment() *Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_vadjustment(_arg0)
	runtime.KeepAlive(scrollable)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// VScrollPolicy gets the vertical GtkScrollablePolicy.
func (scrollable *Scrollable) VScrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))

	_cret = C.gtk_scrollable_get_vscroll_policy(_arg0)
	runtime.KeepAlive(scrollable)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

// SetHAdjustment sets the horizontal adjustment of the GtkScrollable.
func (scrollable *Scrollable) SetHAdjustment(hadjustment *Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	if hadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	}

	C.gtk_scrollable_set_hadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(hadjustment)
}

// SetHScrollPolicy sets the GtkScrollablePolicy.
//
// The policy determines whether horizontal scrolling should start below the
// minimum width or below the natural width.
func (scrollable *Scrollable) SetHScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_hscroll_policy(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(policy)
}

// SetVAdjustment sets the vertical adjustment of the GtkScrollable.
func (scrollable *Scrollable) SetVAdjustment(vadjustment *Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	if vadjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))
	}

	C.gtk_scrollable_set_vadjustment(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(vadjustment)
}

// SetVScrollPolicy sets the GtkScrollablePolicy.
//
// The policy determines whether vertical scrolling should start below the
// minimum height or below the natural height.
func (scrollable *Scrollable) SetVScrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(scrollable.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_vscroll_policy(_arg0, _arg1)
	runtime.KeepAlive(scrollable)
	runtime.KeepAlive(policy)
}
