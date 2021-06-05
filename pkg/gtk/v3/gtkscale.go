// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_scale_get_type()), F: marshalScale},
	})
}

// Scale: a GtkScale is a slider control used to select a numeric value. To use
// it, you’ll probably want to investigate the methods on its base class, Range,
// in addition to the methods for GtkScale itself. To set the value of a scale,
// you would normally use gtk_range_set_value(). To detect changes to the value,
// you would normally use the Range::value-changed signal.
//
// Note that using the same upper and lower bounds for the Scale (through the
// Range methods) will hide the slider itself. This is useful for applications
// that want to show an undeterminate value on the scale, without changing the
// layout of the application (such as movie or music players).
//
//
// GtkScale as GtkBuildable
//
// GtkScale supports a custom <marks> element, which can contain multiple <mark>
// elements. The “value” and “position” attributes have the same meaning as
// gtk_scale_add_mark() parameters of the same name. If the element is not
// empty, its content is taken as the markup to show at the mark. It can be
// translated with the usual ”translatable” and “context” attributes.
//
// CSS nodes
//
//    scale[.fine-tune][.marks-before][.marks-after]
//    ├── marks.top
//    │   ├── mark
//    │   ┊    ├── [label]
//    │   ┊    ╰── indicator
//    ┊   ┊
//    │   ╰── mark
//    ├── [value]
//    ├── contents
//    │   ╰── trough
//    │       ├── slider
//    │       ├── [highlight]
//    │       ╰── [fill]
//    ╰── marks.bottom
//        ├── mark
//        ┊    ├── indicator
//        ┊    ╰── [label]
//        ╰── mark
//
// GtkScale has a main CSS node with name scale and a subnode for its contents,
// with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see gtk_scale_set_has_origin()), there is a
// subnode with name highlight below the trough node that is used for rendering
// the highlighted part of the trough.
//
// If the scale is showing a fill level (see gtk_range_set_show_fill_level()),
// there is a subnode with name fill below the trough node that is used for
// rendering the filled in part of the trough.
//
// If marks are present, there is a marks subnode before or after the contents
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see Scale:draw-value), there is subnode
// with name value.
type Scale interface {
	Range
	Buildable
	Orientable

	// AddMark adds a mark at @value.
	//
	// A mark is indicated visually by drawing a tick mark next to the scale,
	// and GTK+ makes it easy for the user to position the scale exactly at the
	// marks value.
	//
	// If @markup is not nil, text is shown next to the tick mark.
	//
	// To remove marks from a scale, use gtk_scale_clear_marks().
	AddMark(value float64, position PositionType, markup string)
	// ClearMarks removes any marks that have been added with
	// gtk_scale_add_mark().
	ClearMarks()
	// Digits gets the number of decimal places that are displayed in the value.
	Digits() int
	// DrawValue returns whether the current value is displayed as a string next
	// to the slider.
	DrawValue() bool
	// HasOrigin returns whether the scale has an origin.
	HasOrigin() bool
	// Layout gets the Layout used to display the scale. The returned object is
	// owned by the scale so does not need to be freed by the caller.
	Layout() pango.Layout
	// LayoutOffsets obtains the coordinates where the scale will draw the
	// Layout representing the text in the scale. Remember when using the Layout
	// function you need to convert to and from pixels using PANGO_PIXELS() or
	// NGO_SCALE.
	//
	// If the Scale:draw-value property is false, the return values are
	// undefined.
	LayoutOffsets() (x int, y int)
	// ValuePos gets the position in which the current value is displayed.
	ValuePos() PositionType
	// SetDigits sets the number of decimal places that are displayed in the
	// value. Also causes the value of the adjustment to be rounded to this
	// number of digits, so the retrieved value matches the displayed one, if
	// Scale:draw-value is true when the value changes. If you want to enforce
	// rounding the value when Scale:draw-value is false, you can set
	// Range:round-digits instead.
	//
	// Note that rounding to a small number of digits can interfere with the
	// smooth autoscrolling that is built into Scale. As an alternative, you can
	// use the Scale::format-value signal to format the displayed value
	// yourself.
	SetDigits(digits int)
	// SetDrawValue specifies whether the current value is displayed as a string
	// next to the slider.
	SetDrawValue(drawValue bool)
	// SetHasOrigin: if Scale:has-origin is set to true (the default), the scale
	// will highlight the part of the trough between the origin (bottom or left
	// side) and the current value.
	SetHasOrigin(hasOrigin bool)
	// SetValuePos sets the position in which the current value is displayed.
	SetValuePos(pos PositionType)
}

// scale implements the Scale interface.
type scale struct {
	Range
	Buildable
	Orientable
}

var _ Scale = (*scale)(nil)

// WrapScale wraps a GObject to the right type. It is
// primarily used internally.
func WrapScale(obj *externglib.Object) Scale {
	return Scale{
		Range:      WrapRange(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalScale(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScale(obj), nil
}

// NewScale constructs a class Scale.
func NewScale(orientation Orientation, adjustment Adjustment) Scale {
	var arg1 C.GtkOrientation
	var arg2 *C.GtkAdjustment

	arg1 = (C.GtkOrientation)(orientation)
	arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	var cret C.GtkScale
	var goret1 Scale

	cret = C.gtk_scale_new(orientation, adjustment)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Scale)

	return goret1
}

// NewScaleWithRange constructs a class Scale.
func NewScaleWithRange(orientation Orientation, min float64, max float64, step float64) Scale {
	var arg1 C.GtkOrientation
	var arg2 C.gdouble
	var arg3 C.gdouble
	var arg4 C.gdouble

	arg1 = (C.GtkOrientation)(orientation)
	arg2 = C.gdouble(min)
	arg3 = C.gdouble(max)
	arg4 = C.gdouble(step)

	var cret C.GtkScale
	var goret1 Scale

	cret = C.gtk_scale_new_with_range(orientation, min, max, step)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Scale)

	return goret1
}

// AddMark adds a mark at @value.
//
// A mark is indicated visually by drawing a tick mark next to the scale,
// and GTK+ makes it easy for the user to position the scale exactly at the
// marks value.
//
// If @markup is not nil, text is shown next to the tick mark.
//
// To remove marks from a scale, use gtk_scale_clear_marks().
func (s scale) AddMark(value float64, position PositionType, markup string) {
	var arg0 *C.GtkScale
	var arg1 C.gdouble
	var arg2 C.GtkPositionType
	var arg3 *C.gchar

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))
	arg1 = C.gdouble(value)
	arg2 = (C.GtkPositionType)(position)
	arg3 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_scale_add_mark(arg0, value, position, markup)
}

// ClearMarks removes any marks that have been added with
// gtk_scale_add_mark().
func (s scale) ClearMarks() {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	C.gtk_scale_clear_marks(arg0)
}

// Digits gets the number of decimal places that are displayed in the value.
func (s scale) Digits() int {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gtk_scale_get_digits(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// DrawValue returns whether the current value is displayed as a string next
// to the slider.
func (s scale) DrawValue() bool {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_scale_get_draw_value(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// HasOrigin returns whether the scale has an origin.
func (s scale) HasOrigin() bool {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_scale_get_has_origin(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Layout gets the Layout used to display the scale. The returned object is
// owned by the scale so does not need to be freed by the caller.
func (s scale) Layout() pango.Layout {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var cret *C.PangoLayout
	var goret1 pango.Layout

	cret = C.gtk_scale_get_layout(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.Layout)

	return goret1
}

// LayoutOffsets obtains the coordinates where the scale will draw the
// Layout representing the text in the scale. Remember when using the Layout
// function you need to convert to and from pixels using PANGO_PIXELS() or
// NGO_SCALE.
//
// If the Scale:draw-value property is false, the return values are
// undefined.
func (s scale) LayoutOffsets() (x int, y int) {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var arg1 *C.gint
	var ret1 int
	var arg2 *C.gint
	var ret2 int

	C.gtk_scale_get_layout_offsets(arg0, &arg1, &arg2)

	ret1 = *C.gint(arg1)
	ret2 = *C.gint(arg2)

	return ret1, ret2
}

// ValuePos gets the position in which the current value is displayed.
func (s scale) ValuePos() PositionType {
	var arg0 *C.GtkScale

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))

	var cret C.GtkPositionType
	var goret1 PositionType

	cret = C.gtk_scale_get_value_pos(arg0)

	goret1 = PositionType(cret)

	return goret1
}

// SetDigits sets the number of decimal places that are displayed in the
// value. Also causes the value of the adjustment to be rounded to this
// number of digits, so the retrieved value matches the displayed one, if
// Scale:draw-value is true when the value changes. If you want to enforce
// rounding the value when Scale:draw-value is false, you can set
// Range:round-digits instead.
//
// Note that rounding to a small number of digits can interfere with the
// smooth autoscrolling that is built into Scale. As an alternative, you can
// use the Scale::format-value signal to format the displayed value
// yourself.
func (s scale) SetDigits(digits int) {
	var arg0 *C.GtkScale
	var arg1 C.gint

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))
	arg1 = C.gint(digits)

	C.gtk_scale_set_digits(arg0, digits)
}

// SetDrawValue specifies whether the current value is displayed as a string
// next to the slider.
func (s scale) SetDrawValue(drawValue bool) {
	var arg0 *C.GtkScale
	var arg1 C.gboolean

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))
	if drawValue {
		arg1 = C.gboolean(1)
	}

	C.gtk_scale_set_draw_value(arg0, drawValue)
}

// SetHasOrigin: if Scale:has-origin is set to true (the default), the scale
// will highlight the part of the trough between the origin (bottom or left
// side) and the current value.
func (s scale) SetHasOrigin(hasOrigin bool) {
	var arg0 *C.GtkScale
	var arg1 C.gboolean

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))
	if hasOrigin {
		arg1 = C.gboolean(1)
	}

	C.gtk_scale_set_has_origin(arg0, hasOrigin)
}

// SetValuePos sets the position in which the current value is displayed.
func (s scale) SetValuePos(pos PositionType) {
	var arg0 *C.GtkScale
	var arg1 C.GtkPositionType

	arg0 = (*C.GtkScale)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkPositionType)(pos)

	C.gtk_scale_set_value_pos(arg0, pos)
}

type ScalePrivate struct {
	native C.GtkScalePrivate
}

// WrapScalePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScalePrivate(ptr unsafe.Pointer) *ScalePrivate {
	if ptr == nil {
		return nil
	}

	return (*ScalePrivate)(ptr)
}

func marshalScalePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScalePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *ScalePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
