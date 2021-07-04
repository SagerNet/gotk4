// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_corner_type_get_type()), F: marshalCornerType},
		{T: externglib.Type(C.gtk_policy_type_get_type()), F: marshalPolicyType},
		{T: externglib.Type(C.gtk_scrolled_window_get_type()), F: marshalScrolledWindow},
	})
}

// CornerType specifies which corner a child widget should be placed in when
// packed into a ScrolledWindow. This is effectively the opposite of where the
// scroll bars are placed.
type CornerType int

const (
	// TopLeft: place the scrollbars on the right and bottom of the widget
	// (default behaviour).
	CornerTypeTopLeft CornerType = 0
	// BottomLeft: place the scrollbars on the top and right of the widget.
	CornerTypeBottomLeft CornerType = 1
	// TopRight: place the scrollbars on the left and bottom of the widget.
	CornerTypeTopRight CornerType = 2
	// BottomRight: place the scrollbars on the top and left of the widget.
	CornerTypeBottomRight CornerType = 3
)

func marshalCornerType(p uintptr) (interface{}, error) {
	return CornerType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PolicyType determines how the size should be computed to achieve the one of
// the visibility mode for the scrollbars.
type PolicyType int

const (
	// always: the scrollbar is always visible. The view size is independent of
	// the content.
	PolicyTypeAlways PolicyType = 0
	// automatic: the scrollbar will appear and disappear as necessary. For
	// example, when all of a TreeView can not be seen.
	PolicyTypeAutomatic PolicyType = 1
	// never: the scrollbar should never appear. In this mode the content
	// determines the size.
	PolicyTypeNever PolicyType = 2
	// external: don't show a scrollbar, but don't force the size to follow the
	// content. This can be used e.g. to make multiple scrolled windows share a
	// scrollbar. Since: 3.16
	PolicyTypeExternal PolicyType = 3
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	return PolicyType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrolledWindow: gtkScrolledWindow is a container that accepts a single child
// widget, makes that child scrollable using either internally added scrollbars
// or externally associated adjustments, and optionally draws a frame around the
// child.
//
// Widgets with native scrolling support, i.e. those whose classes implement the
// Scrollable interface, are added directly. For other types of widget, the
// class Viewport acts as an adaptor, giving scrollability to other widgets.
// GtkScrolledWindow’s implementation of gtk_container_add() intelligently
// accounts for whether or not the added child is a Scrollable. If it isn’t,
// ScrolledWindow wraps the child in a Viewport and adds that for you.
// Therefore, you can just add any child widget and not worry about the details.
//
// If gtk_container_add() has added a Viewport for you, you can remove both your
// added child widget from the Viewport, and the Viewport from the
// GtkScrolledWindow, like this:
//
//    GtkWidget *scrolled_window = gtk_scrolled_window_new (NULL, NULL);
//    GtkWidget *child_widget = gtk_button_new ();
//
//    // GtkButton is not a GtkScrollable, so GtkScrolledWindow will automatically
//    // add a GtkViewport.
//    gtk_container_add (GTK_CONTAINER (scrolled_window),
//                       child_widget);
//
//    // Either of these will result in child_widget being unparented:
//    gtk_container_remove (GTK_CONTAINER (scrolled_window),
//                          child_widget);
//    // or
//    gtk_container_remove (GTK_CONTAINER (scrolled_window),
//                          gtk_bin_get_child (GTK_BIN (scrolled_window)));
//
// Unless ScrolledWindow:policy is GTK_POLICY_NEVER or GTK_POLICY_EXTERNAL,
// GtkScrolledWindow adds internal Scrollbar widgets around its child. The
// scroll position of the child, and if applicable the scrollbars, is controlled
// by the ScrolledWindow:hadjustment and ScrolledWindow:vadjustment that are
// associated with the GtkScrolledWindow. See the docs on Scrollbar for the
// details, but note that the “step_increment” and “page_increment” fields are
// only effective if the policy causes scrollbars to be present.
//
// If a GtkScrolledWindow doesn’t behave quite as you would like, or doesn’t
// have exactly the right layout, it’s very possible to set up your own
// scrolling with Scrollbar and for example a Grid.
//
//
// Touch support
//
// GtkScrolledWindow has built-in support for touch devices. When a touchscreen
// is used, swiping will move the scrolled window, and will expose 'kinetic'
// behavior. This can be turned off with the ScrolledWindow:kinetic-scrolling
// property if it is undesired.
//
// GtkScrolledWindow also displays visual 'overshoot' indication when the
// content is pulled beyond the end, and this situation can be captured with the
// ScrolledWindow::edge-overshot signal.
//
// If no mouse device is present, the scrollbars will overlayed as narrow,
// auto-hiding indicators over the content. If traditional scrollbars are
// desired although no mouse is present, this behaviour can be turned off with
// the ScrolledWindow:overlay-scrolling property.
//
//
// CSS nodes
//
// GtkScrolledWindow has a main CSS node with name scrolledwindow.
//
// It uses subnodes with names overshoot and undershoot to draw the overflow and
// underflow indications. These nodes get the .left, .right, .top or .bottom
// style class added depending on where the indication is drawn.
//
// GtkScrolledWindow also sets the positional style classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering) on its scrollbars.
//
// If both scrollbars are visible, the area where they meet is drawn with a
// subnode named junction.
type ScrolledWindow interface {
	Bin

	// AddWithViewportScrolledWindow:
	AddWithViewportScrolledWindow(child Widget)
	// CaptureButtonPress:
	CaptureButtonPress() bool
	// HAdjustment:
	HAdjustment() Adjustment
	// HScrollbar:
	HScrollbar() Widget
	// KineticScrolling:
	KineticScrolling() bool
	// MaxContentHeight:
	MaxContentHeight() int
	// MaxContentWidth:
	MaxContentWidth() int
	// MinContentHeight:
	MinContentHeight() int
	// MinContentWidth:
	MinContentWidth() int
	// OverlayScrolling:
	OverlayScrolling() bool
	// Placement:
	Placement() CornerType
	// Policy:
	Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType)
	// PropagateNaturalHeight:
	PropagateNaturalHeight() bool
	// PropagateNaturalWidth:
	PropagateNaturalWidth() bool
	// ShadowType:
	ShadowType() ShadowType
	// VAdjustment:
	VAdjustment() Adjustment
	// VScrollbar:
	VScrollbar() Widget
	// SetCaptureButtonPressScrolledWindow:
	SetCaptureButtonPressScrolledWindow(captureButtonPress bool)
	// SetHAdjustmentScrolledWindow:
	SetHAdjustmentScrolledWindow(hadjustment Adjustment)
	// SetKineticScrollingScrolledWindow:
	SetKineticScrollingScrolledWindow(kineticScrolling bool)
	// SetMaxContentHeightScrolledWindow:
	SetMaxContentHeightScrolledWindow(height int)
	// SetMaxContentWidthScrolledWindow:
	SetMaxContentWidthScrolledWindow(width int)
	// SetMinContentHeightScrolledWindow:
	SetMinContentHeightScrolledWindow(height int)
	// SetMinContentWidthScrolledWindow:
	SetMinContentWidthScrolledWindow(width int)
	// SetOverlayScrollingScrolledWindow:
	SetOverlayScrollingScrolledWindow(overlayScrolling bool)
	// SetPlacementScrolledWindow:
	SetPlacementScrolledWindow(windowPlacement CornerType)
	// SetPolicyScrolledWindow:
	SetPolicyScrolledWindow(hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType)
	// SetPropagateNaturalHeightScrolledWindow:
	SetPropagateNaturalHeightScrolledWindow(propagate bool)
	// SetPropagateNaturalWidthScrolledWindow:
	SetPropagateNaturalWidthScrolledWindow(propagate bool)
	// SetShadowTypeScrolledWindow:
	SetShadowTypeScrolledWindow(typ ShadowType)
	// SetVAdjustmentScrolledWindow:
	SetVAdjustmentScrolledWindow(vadjustment Adjustment)
	// UnsetPlacementScrolledWindow:
	UnsetPlacementScrolledWindow()
}

// scrolledWindow implements the ScrolledWindow class.
type scrolledWindow struct {
	Bin
}

// WrapScrolledWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrolledWindow(obj *externglib.Object) ScrolledWindow {
	return scrolledWindow{
		Bin: WrapBin(obj),
	}
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrolledWindow(obj), nil
}

// NewScrolledWindow:
func NewScrolledWindow(hadjustment Adjustment, vadjustment Adjustment) ScrolledWindow {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_scrolled_window_new(_arg1, _arg2)

	var _scrolledWindow ScrolledWindow // out

	_scrolledWindow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ScrolledWindow)

	return _scrolledWindow
}

func (s scrolledWindow) AddWithViewportScrolledWindow(child Widget) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_scrolled_window_add_with_viewport(_arg0, _arg1)
}

func (s scrolledWindow) CaptureButtonPress() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_capture_button_press(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s scrolledWindow) HAdjustment() Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (s scrolledWindow) HScrollbar() Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_hscrollbar(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s scrolledWindow) KineticScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_kinetic_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s scrolledWindow) MaxContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s scrolledWindow) MaxContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s scrolledWindow) MinContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s scrolledWindow) MinContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s scrolledWindow) OverlayScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_overlay_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s scrolledWindow) Placement() CornerType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkCornerType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_placement(_arg0)

	var _cornerType CornerType // out

	_cornerType = CornerType(_cret)

	return _cornerType
}

func (s scrolledWindow) Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // in
	var _arg2 C.GtkPolicyType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	C.gtk_scrolled_window_get_policy(_arg0, &_arg1, &_arg2)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	_hscrollbarPolicy = PolicyType(_arg1)
	_vscrollbarPolicy = PolicyType(_arg2)

	return _hscrollbarPolicy, _vscrollbarPolicy
}

func (s scrolledWindow) PropagateNaturalHeight() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_height(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s scrolledWindow) PropagateNaturalWidth() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s scrolledWindow) ShadowType() ShadowType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkShadowType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

func (s scrolledWindow) VAdjustment() Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (s scrolledWindow) VScrollbar() Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_vscrollbar(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s scrolledWindow) SetCaptureButtonPressScrolledWindow(captureButtonPress bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if captureButtonPress {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_capture_button_press(_arg0, _arg1)
}

func (s scrolledWindow) SetHAdjustmentScrolledWindow(hadjustment Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))

	C.gtk_scrolled_window_set_hadjustment(_arg0, _arg1)
}

func (s scrolledWindow) SetKineticScrollingScrolledWindow(kineticScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if kineticScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_kinetic_scrolling(_arg0, _arg1)
}

func (s scrolledWindow) SetMaxContentHeightScrolledWindow(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(height)

	C.gtk_scrolled_window_set_max_content_height(_arg0, _arg1)
}

func (s scrolledWindow) SetMaxContentWidthScrolledWindow(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(width)

	C.gtk_scrolled_window_set_max_content_width(_arg0, _arg1)
}

func (s scrolledWindow) SetMinContentHeightScrolledWindow(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(height)

	C.gtk_scrolled_window_set_min_content_height(_arg0, _arg1)
}

func (s scrolledWindow) SetMinContentWidthScrolledWindow(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.gint(width)

	C.gtk_scrolled_window_set_min_content_width(_arg0, _arg1)
}

func (s scrolledWindow) SetOverlayScrollingScrolledWindow(overlayScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if overlayScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_overlay_scrolling(_arg0, _arg1)
}

func (s scrolledWindow) SetPlacementScrolledWindow(windowPlacement CornerType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkCornerType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkCornerType(windowPlacement)

	C.gtk_scrolled_window_set_placement(_arg0, _arg1)
}

func (s scrolledWindow) SetPolicyScrolledWindow(hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // out
	var _arg2 C.GtkPolicyType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPolicyType(hscrollbarPolicy)
	_arg2 = C.GtkPolicyType(vscrollbarPolicy)

	C.gtk_scrolled_window_set_policy(_arg0, _arg1, _arg2)
}

func (s scrolledWindow) SetPropagateNaturalHeightScrolledWindow(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_height(_arg0, _arg1)
}

func (s scrolledWindow) SetPropagateNaturalWidthScrolledWindow(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_width(_arg0, _arg1)
}

func (s scrolledWindow) SetShadowTypeScrolledWindow(typ ShadowType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkShadowType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_scrolled_window_set_shadow_type(_arg0, _arg1)
}

func (s scrolledWindow) SetVAdjustmentScrolledWindow(vadjustment Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	C.gtk_scrolled_window_set_vadjustment(_arg0, _arg1)
}

func (s scrolledWindow) UnsetPlacementScrolledWindow() {
	var _arg0 *C.GtkScrolledWindow // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	C.gtk_scrolled_window_unset_placement(_arg0)
}

func (b scrolledWindow) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b scrolledWindow) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b scrolledWindow) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b scrolledWindow) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b scrolledWindow) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b scrolledWindow) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b scrolledWindow) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
