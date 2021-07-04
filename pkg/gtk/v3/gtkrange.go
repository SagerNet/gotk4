// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRange},
	})
}

// Range is the common base class for widgets which visualize an adjustment, e.g
// Scale or Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, Range
// provides properties and methods for influencing the sensitivity of the
// “steppers”. It also provides properties and methods for setting a “fill
// level” on range widgets. See gtk_range_set_fill_level().
type Range interface {
	Widget
	Orientable

	Adjustment() Adjustment

	FillLevel() float64

	Flippable() bool

	Inverted() bool

	LowerStepperSensitivity() SensitivityType

	MinSliderSize() int

	RangeRect() gdk.Rectangle

	RestrictToFillLevel() bool

	RoundDigits() int

	ShowFillLevel() bool

	SliderRange() (sliderStart int, sliderEnd int)

	SliderSizeFixed() bool

	UpperStepperSensitivity() SensitivityType

	Value() float64

	SetAdjustmentRange(adjustment Adjustment)

	SetFillLevelRange(fillLevel float64)

	SetFlippableRange(flippable bool)

	SetIncrementsRange(step float64, page float64)

	SetInvertedRange(setting bool)

	SetLowerStepperSensitivityRange(sensitivity SensitivityType)

	SetMinSliderSizeRange(minSize int)

	SetRangeRange(min float64, max float64)

	SetRestrictToFillLevelRange(restrictToFillLevel bool)

	SetRoundDigitsRange(roundDigits int)

	SetShowFillLevelRange(showFillLevel bool)

	SetSliderSizeFixedRange(sizeFixed bool)

	SetUpperStepperSensitivityRange(sensitivity SensitivityType)

	SetValueRange(value float64)
}

// _range implements the Range class.
type _range struct {
	Widget
}

// WrapRange wraps a GObject to the right type. It is
// primarily used internally.
func WrapRange(obj *externglib.Object) Range {
	return _range{
		Widget: WrapWidget(obj),
	}
}

func marshalRange(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRange(obj), nil
}

func (r _range) Adjustment() Adjustment {
	var _arg0 *C.GtkRange      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_adjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (r _range) FillLevel() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (r _range) Flippable() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_flippable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) Inverted() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) LowerStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_lower_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

func (r _range) MinSliderSize() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_min_slider_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (r _range) RangeRect() gdk.Rectangle {
	var _arg0 *C.GtkRange    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	C.gtk_range_get_range_rect(_arg0, &_arg1)

	var _rangeRect gdk.Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_rangeRect = *refTmpOut
	}

	return _rangeRect
}

func (r _range) RestrictToFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_restrict_to_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_round_digits(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (r _range) ShowFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_show_fill_level(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

func (r _range) SliderSizeFixed() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_slider_size_fixed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (r _range) UpperStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_upper_stepper_sensitivity(_arg0)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

func (r _range) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))

	_cret = C.gtk_range_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (r _range) SetAdjustmentRange(adjustment Adjustment) {
	var _arg0 *C.GtkRange      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_range_set_adjustment(_arg0, _arg1)
}

func (r _range) SetFillLevelRange(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
}

func (r _range) SetFlippableRange(flippable bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if flippable {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(_arg0, _arg1)
}

func (r _range) SetIncrementsRange(step float64, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
}

func (r _range) SetInvertedRange(setting bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(_arg0, _arg1)
}

func (r _range) SetLowerStepperSensitivityRange(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_lower_stepper_sensitivity(_arg0, _arg1)
}

func (r _range) SetMinSliderSizeRange(minSize int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint(minSize)

	C.gtk_range_set_min_slider_size(_arg0, _arg1)
}

func (r _range) SetRangeRange(min float64, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
}

func (r _range) SetRestrictToFillLevelRange(restrictToFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if restrictToFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(_arg0, _arg1)
}

func (r _range) SetRoundDigitsRange(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
}

func (r _range) SetShowFillLevelRange(showFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if showFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(_arg0, _arg1)
}

func (r _range) SetSliderSizeFixedRange(sizeFixed bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	if sizeFixed {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(_arg0, _arg1)
}

func (r _range) SetUpperStepperSensitivityRange(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_upper_stepper_sensitivity(_arg0, _arg1)
}

func (r _range) SetValueRange(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_range_set_value(_arg0, _arg1)
}

func (b _range) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b _range) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b _range) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b _range) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b _range) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b _range) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b _range) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b _range) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b _range) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b _range) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o _range) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o _range) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}