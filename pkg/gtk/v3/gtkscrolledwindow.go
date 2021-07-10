// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_scrolled_window_get_type()), F: marshalScrolledWindowwer},
	})
}

// CornerType specifies which corner a child widget should be placed in when
// packed into a ScrolledWindow. This is effectively the opposite of where the
// scroll bars are placed.
type CornerType int

const (
	// TopLeft: place the scrollbars on the right and bottom of the widget
	// (default behaviour).
	CornerTypeTopLeft CornerType = iota
	// BottomLeft: place the scrollbars on the top and right of the widget.
	CornerTypeBottomLeft
	// TopRight: place the scrollbars on the left and bottom of the widget.
	CornerTypeTopRight
	// BottomRight: place the scrollbars on the top and left of the widget.
	CornerTypeBottomRight
)

func marshalCornerType(p uintptr) (interface{}, error) {
	return CornerType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PolicyType determines how the size should be computed to achieve the one of
// the visibility mode for the scrollbars.
type PolicyType int

const (
	// Always: the scrollbar is always visible. The view size is independent of
	// the content.
	PolicyTypeAlways PolicyType = iota
	// Automatic: the scrollbar will appear and disappear as necessary. For
	// example, when all of a TreeView can not be seen.
	PolicyTypeAutomatic
	// Never: the scrollbar should never appear. In this mode the content
	// determines the size.
	PolicyTypeNever
	// External: don't show a scrollbar, but don't force the size to follow the
	// content. This can be used e.g. to make multiple scrolled windows share a
	// scrollbar. Since: 3.16
	PolicyTypeExternal
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	return PolicyType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrolledWindowwer describes ScrolledWindow's methods.
type ScrolledWindowwer interface {
	gextras.Objector

	AddWithViewport(child Widgetter)
	CaptureButtonPress() bool
	HAdjustment() *Adjustment
	Hscrollbar() *Widget
	KineticScrolling() bool
	MaxContentHeight() int
	MaxContentWidth() int
	MinContentHeight() int
	MinContentWidth() int
	OverlayScrolling() bool
	Placement() CornerType
	Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType)
	PropagateNaturalHeight() bool
	PropagateNaturalWidth() bool
	ShadowType() ShadowType
	VAdjustment() *Adjustment
	Vscrollbar() *Widget
	SetCaptureButtonPress(captureButtonPress bool)
	SetHAdjustment(hadjustment Adjustmenter)
	SetKineticScrolling(kineticScrolling bool)
	SetMaxContentHeight(height int)
	SetMaxContentWidth(width int)
	SetMinContentHeight(height int)
	SetMinContentWidth(width int)
	SetOverlayScrolling(overlayScrolling bool)
	SetPropagateNaturalHeight(propagate bool)
	SetPropagateNaturalWidth(propagate bool)
	SetVAdjustment(vadjustment Adjustmenter)
	UnsetPlacement()
}

// ScrolledWindow is a container that accepts a single child widget, makes that
// child scrollable using either internally added scrollbars or externally
// associated adjustments, and optionally draws a frame around the child.
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
type ScrolledWindow struct {
	*externglib.Object
	Bin
	Buildable
}

var _ ScrolledWindowwer = (*ScrolledWindow)(nil)

func wrapScrolledWindowwer(obj *externglib.Object) ScrolledWindowwer {
	return &ScrolledWindow{
		Object: obj,
		Bin: Bin{
			Object: obj,
			Container: Container{
				Object: obj,
				Widget: Widget{
					Object: obj,
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalScrolledWindowwer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScrolledWindowwer(obj), nil
}

// NewScrolledWindow creates a new scrolled window.
//
// The two arguments are the scrolled window’s adjustments; these will be shared
// with the scrollbars and the child widget to keep the bars in sync with the
// child. Usually you want to pass nil for the adjustments, which will cause the
// scrolled window to create them for you.
func NewScrolledWindow(hadjustment Adjustmenter, vadjustment Adjustmenter) *ScrolledWindow {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_scrolled_window_new(_arg1, _arg2)

	var _scrolledWindow *ScrolledWindow // out

	_scrolledWindow = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ScrolledWindow)

	return _scrolledWindow
}

// AddWithViewport: used to add children without native scrolling capabilities.
// This is simply a convenience function; it is equivalent to adding the
// unscrollable child to a viewport, then adding the viewport to the scrolled
// window. If a child has native scrolling, use gtk_container_add() instead of
// this function.
//
// The viewport scrolls the child by moving its Window, and takes the size of
// the child to be the size of its toplevel Window. This will be very wrong for
// most widgets that support native scrolling; for example, if you add a widget
// such as TreeView with a viewport, the whole widget will scroll, including the
// column headings. Thus, widgets with native scrolling support should not be
// used with the Viewport proxy.
//
// A widget supports scrolling natively if it implements the Scrollable
// interface.
//
// Deprecated: gtk_container_add() will automatically add a Viewport if the
// child doesn’t implement Scrollable.
func (scrolledWindow *ScrolledWindow) AddWithViewport(child Widgetter) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_scrolled_window_add_with_viewport(_arg0, _arg1)
}

// CaptureButtonPress: return whether button presses are captured during kinetic
// scrolling. See gtk_scrolled_window_set_capture_button_press().
func (scrolledWindow *ScrolledWindow) CaptureButtonPress() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_capture_button_press(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HAdjustment returns the horizontal scrollbar’s adjustment, used to connect
// the horizontal scrollbar to the child widget’s horizontal scroll
// functionality.
func (scrolledWindow *ScrolledWindow) HAdjustment() *Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_hadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// Hscrollbar returns the horizontal scrollbar of @scrolled_window.
func (scrolledWindow *ScrolledWindow) Hscrollbar() *Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_hscrollbar(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// KineticScrolling returns the specified kinetic scrolling behavior.
func (scrolledWindow *ScrolledWindow) KineticScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_kinetic_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxContentHeight returns the maximum content height set.
func (scrolledWindow *ScrolledWindow) MaxContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxContentWidth returns the maximum content width set.
func (scrolledWindow *ScrolledWindow) MaxContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MinContentHeight gets the minimal content height of @scrolled_window, or -1
// if not set.
func (scrolledWindow *ScrolledWindow) MinContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_height(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MinContentWidth gets the minimum content width of @scrolled_window, or -1 if
// not set.
func (scrolledWindow *ScrolledWindow) MinContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_width(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverlayScrolling returns whether overlay scrolling is enabled for this
// scrolled window.
func (scrolledWindow *ScrolledWindow) OverlayScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_overlay_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Placement gets the placement of the contents with respect to the scrollbars
// for the scrolled window. See gtk_scrolled_window_set_placement().
func (scrolledWindow *ScrolledWindow) Placement() CornerType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkCornerType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_placement(_arg0)

	var _cornerType CornerType // out

	_cornerType = (CornerType)(_cret)

	return _cornerType
}

// Policy retrieves the current policy values for the horizontal and vertical
// scrollbars. See gtk_scrolled_window_set_policy().
func (scrolledWindow *ScrolledWindow) Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // in
	var _arg2 C.GtkPolicyType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	C.gtk_scrolled_window_get_policy(_arg0, &_arg1, &_arg2)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	_hscrollbarPolicy = (PolicyType)(_arg1)
	_vscrollbarPolicy = (PolicyType)(_arg2)

	return _hscrollbarPolicy, _vscrollbarPolicy
}

// PropagateNaturalHeight reports whether the natural height of the child will
// be calculated and propagated through the scrolled window’s requested natural
// height.
func (scrolledWindow *ScrolledWindow) PropagateNaturalHeight() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_height(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PropagateNaturalWidth reports whether the natural width of the child will be
// calculated and propagated through the scrolled window’s requested natural
// width.
func (scrolledWindow *ScrolledWindow) PropagateNaturalWidth() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShadowType gets the shadow type of the scrolled window. See
// gtk_scrolled_window_set_shadow_type().
func (scrolledWindow *ScrolledWindow) ShadowType() ShadowType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkShadowType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = (ShadowType)(_cret)

	return _shadowType
}

// VAdjustment returns the vertical scrollbar’s adjustment, used to connect the
// vertical scrollbar to the child widget’s vertical scroll functionality.
func (scrolledWindow *ScrolledWindow) VAdjustment() *Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_vadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// Vscrollbar returns the vertical scrollbar of @scrolled_window.
func (scrolledWindow *ScrolledWindow) Vscrollbar() *Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	_cret = C.gtk_scrolled_window_get_vscrollbar(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// SetCaptureButtonPress changes the behaviour of @scrolled_window with regard
// to the initial event that possibly starts kinetic scrolling. When
// @capture_button_press is set to true, the event is captured by the scrolled
// window, and then later replayed if it is meant to go to the child widget.
//
// This should be enabled if any child widgets perform non-reversible actions on
// Widget::button-press-event. If they don't, and handle additionally handle
// Widget::grab-broken-event, it might be better to set @capture_button_press to
// false.
//
// This setting only has an effect if kinetic scrolling is enabled.
func (scrolledWindow *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if captureButtonPress {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_capture_button_press(_arg0, _arg1)
}

// SetHAdjustment sets the Adjustment for the horizontal scrollbar.
func (scrolledWindow *ScrolledWindow) SetHAdjustment(hadjustment Adjustmenter) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))

	C.gtk_scrolled_window_set_hadjustment(_arg0, _arg1)
}

// SetKineticScrolling turns kinetic scrolling on or off. Kinetic scrolling only
// applies to devices with source GDK_SOURCE_TOUCHSCREEN.
func (scrolledWindow *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if kineticScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_kinetic_scrolling(_arg0, _arg1)
}

// SetMaxContentHeight sets the maximum height that @scrolled_window should keep
// visible. The @scrolled_window will grow up to this height before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than ScrolledWindow:min-content-height.
func (scrolledWindow *ScrolledWindow) SetMaxContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.gint(height)

	C.gtk_scrolled_window_set_max_content_height(_arg0, _arg1)
}

// SetMaxContentWidth sets the maximum width that @scrolled_window should keep
// visible. The @scrolled_window will grow up to this width before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content width to a value smaller
// than ScrolledWindow:min-content-width.
func (scrolledWindow *ScrolledWindow) SetMaxContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.gint(width)

	C.gtk_scrolled_window_set_max_content_width(_arg0, _arg1)
}

// SetMinContentHeight sets the minimum height that @scrolled_window should keep
// visible. Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content height to a value
// greater than ScrolledWindow:max-content-height.
func (scrolledWindow *ScrolledWindow) SetMinContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.gint(height)

	C.gtk_scrolled_window_set_min_content_height(_arg0, _arg1)
}

// SetMinContentWidth sets the minimum width that @scrolled_window should keep
// visible. Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content width to a value greater
// than ScrolledWindow:max-content-width.
func (scrolledWindow *ScrolledWindow) SetMinContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = C.gint(width)

	C.gtk_scrolled_window_set_min_content_width(_arg0, _arg1)
}

// SetOverlayScrolling enables or disables overlay scrolling for this scrolled
// window.
func (scrolledWindow *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if overlayScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_overlay_scrolling(_arg0, _arg1)
}

// SetPropagateNaturalHeight sets whether the natural height of the child should
// be calculated and propagated through the scrolled window’s requested natural
// height.
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_height(_arg0, _arg1)
}

// SetPropagateNaturalWidth sets whether the natural width of the child should
// be calculated and propagated through the scrolled window’s requested natural
// width.
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_width(_arg0, _arg1)
}

// SetVAdjustment sets the Adjustment for the vertical scrollbar.
func (scrolledWindow *ScrolledWindow) SetVAdjustment(vadjustment Adjustmenter) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	C.gtk_scrolled_window_set_vadjustment(_arg0, _arg1)
}

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars for the scrolled window. If no window placement is set for a
// scrolled window, it defaults to GTK_CORNER_TOP_LEFT.
//
// See also gtk_scrolled_window_set_placement() and
// gtk_scrolled_window_get_placement().
func (scrolledWindow *ScrolledWindow) UnsetPlacement() {
	var _arg0 *C.GtkScrolledWindow // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow.Native()))

	C.gtk_scrolled_window_unset_placement(_arg0)
}
