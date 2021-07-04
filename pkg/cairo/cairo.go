// Code generated by girgen. DO NOT EDIT.

package cairo

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: cairo-gobject gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <cairo-gobject.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.cairo_gobject_status_get_type()), F: marshalStatus},
		{T: externglib.Type(C.cairo_gobject_content_get_type()), F: marshalContent},
		{T: externglib.Type(C.cairo_gobject_operator_get_type()), F: marshalOperator},
		{T: externglib.Type(C.cairo_gobject_antialias_get_type()), F: marshalAntialias},
		{T: externglib.Type(C.cairo_gobject_fill_rule_get_type()), F: marshalFillRule},
		{T: externglib.Type(C.cairo_gobject_line_cap_get_type()), F: marshalLineCap},
		{T: externglib.Type(C.cairo_gobject_line_join_get_type()), F: marshalLineJoin},
		{T: externglib.Type(C.cairo_gobject_text_cluster_flags_get_type()), F: marshalTextClusterFlags},
		{T: externglib.Type(C.cairo_gobject_font_slant_get_type()), F: marshalFontSlant},
		{T: externglib.Type(C.cairo_gobject_font_weight_get_type()), F: marshalFontWeight},
		{T: externglib.Type(C.cairo_gobject_subpixel_order_get_type()), F: marshalSubpixelOrder},
		{T: externglib.Type(C.cairo_gobject_hint_style_get_type()), F: marshalHintStyle},
		{T: externglib.Type(C.cairo_gobject_hint_metrics_get_type()), F: marshalHintMetrics},
		{T: externglib.Type(C.cairo_gobject_font_type_get_type()), F: marshalFontType},
		{T: externglib.Type(C.cairo_gobject_path_data_type_get_type()), F: marshalPathDataType},
		{T: externglib.Type(C.cairo_gobject_device_type_get_type()), F: marshalDeviceType},
		{T: externglib.Type(C.cairo_gobject_surface_type_get_type()), F: marshalSurfaceType},
		{T: externglib.Type(C.cairo_gobject_format_get_type()), F: marshalFormat},
		{T: externglib.Type(C.cairo_gobject_pattern_type_get_type()), F: marshalPatternType},
		{T: externglib.Type(C.cairo_gobject_extend_get_type()), F: marshalExtend},
		{T: externglib.Type(C.cairo_gobject_filter_get_type()), F: marshalFilter},
		{T: externglib.Type(C.cairo_gobject_region_overlap_get_type()), F: marshalRegionOverlap},
		{T: externglib.Type(C.cairo_gobject_context_get_type()), F: marshalContext},
		{T: externglib.Type(C.cairo_gobject_device_get_type()), F: marshalDevice},
		{T: externglib.Type(C.cairo_gobject_surface_get_type()), F: marshalSurface},
		{T: externglib.Type(C.cairo_gobject_pattern_get_type()), F: marshalPattern},
		{T: externglib.Type(C.cairo_gobject_region_get_type()), F: marshalRegion},
		{T: externglib.Type(C.cairo_gobject_font_options_get_type()), F: marshalFontOptions},
		{T: externglib.Type(C.cairo_gobject_font_face_get_type()), F: marshalFontFace},
		{T: externglib.Type(C.cairo_gobject_scaled_font_get_type()), F: marshalScaledFont},
		{T: externglib.Type(C.cairo_gobject_rectangle_get_type()), F: marshalRectangle},
		{T: externglib.Type(C.cairo_gobject_rectangle_int_get_type()), F: marshalRectangleInt},
	})
}

// Status:
type Status int

const (
	StatusSuccess                 Status = 0
	StatusNoMemory                Status = 1
	StatusInvalidRestore          Status = 2
	StatusInvalidPopGroup         Status = 3
	StatusNoCurrentPoint          Status = 4
	StatusInvalidMatrix           Status = 5
	StatusInvalidStatus           Status = 6
	StatusNullPointer             Status = 7
	StatusInvalidString           Status = 8
	StatusInvalidPathData         Status = 9
	StatusReadError               Status = 10
	StatusWriteError              Status = 11
	StatusSurfaceFinished         Status = 12
	StatusSurfaceTypeMismatch     Status = 13
	StatusPatternTypeMismatch     Status = 14
	StatusInvalidContent          Status = 15
	StatusInvalidFormat           Status = 16
	StatusInvalidVisual           Status = 17
	StatusFileNotFound            Status = 18
	StatusInvalidDash             Status = 19
	StatusInvalidDscComment       Status = 20
	StatusInvalidIndex            Status = 21
	StatusClipNotRepresentable    Status = 22
	StatusTempFileError           Status = 23
	StatusInvalidStride           Status = 24
	StatusFontTypeMismatch        Status = 25
	StatusUserFontImmutable       Status = 26
	StatusUserFontError           Status = 27
	StatusNegativeCount           Status = 28
	StatusInvalidClusters         Status = 29
	StatusInvalidSlant            Status = 30
	StatusInvalidWeight           Status = 31
	StatusInvalidSize             Status = 32
	StatusUserFontNotImplemented  Status = 33
	StatusDeviceTypeMismatch      Status = 34
	StatusDeviceError             Status = 35
	StatusInvalidMeshConstruction Status = 36
	StatusDeviceFinished          Status = 37
	StatusJbig2GlobalMissing      Status = 38
)

func marshalStatus(p uintptr) (interface{}, error) {
	return Status(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Content:
type Content int

const (
	ContentColor      Content = 4096
	ContentAlpha      Content = 8192
	ContentColorAlpha Content = 12288
)

func marshalContent(p uintptr) (interface{}, error) {
	return Content(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Operator:
type Operator int

const (
	OperatorClear         Operator = 0
	OperatorSource        Operator = 1
	OperatorOver          Operator = 2
	OperatorIn            Operator = 3
	OperatorOut           Operator = 4
	OperatorAtop          Operator = 5
	OperatorDest          Operator = 6
	OperatorDestOver      Operator = 7
	OperatorDestIn        Operator = 8
	OperatorDestOut       Operator = 9
	OperatorDestAtop      Operator = 10
	OperatorXor           Operator = 11
	OperatorAdd           Operator = 12
	OperatorSaturate      Operator = 13
	OperatorMultiply      Operator = 14
	OperatorScreen        Operator = 15
	OperatorOverlay       Operator = 16
	OperatorDarken        Operator = 17
	OperatorLighten       Operator = 18
	OperatorColorDodge    Operator = 19
	OperatorColorBurn     Operator = 20
	OperatorHardLight     Operator = 21
	OperatorSoftLight     Operator = 22
	OperatorDifference    Operator = 23
	OperatorExclusion     Operator = 24
	OperatorHSLHue        Operator = 25
	OperatorHSLSaturation Operator = 26
	OperatorHSLColor      Operator = 27
	OperatorHSLLuminosity Operator = 28
)

func marshalOperator(p uintptr) (interface{}, error) {
	return Operator(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Antialias:
type Antialias int

const (
	AntialiasDefault  Antialias = 0
	AntialiasNone     Antialias = 1
	AntialiasGray     Antialias = 2
	AntialiasSubpixel Antialias = 3
	AntialiasFast     Antialias = 4
	AntialiasGood     Antialias = 5
	AntialiasBest     Antialias = 6
)

func marshalAntialias(p uintptr) (interface{}, error) {
	return Antialias(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FillRule:
type FillRule int

const (
	FillRuleWinding FillRule = 0
	FillRuleEvenOdd FillRule = 1
)

func marshalFillRule(p uintptr) (interface{}, error) {
	return FillRule(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// LineCap:
type LineCap int

const (
	LineCapButt   LineCap = 0
	LineCapRound  LineCap = 1
	LineCapSquare LineCap = 2
)

func marshalLineCap(p uintptr) (interface{}, error) {
	return LineCap(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// LineJoin:
type LineJoin int

const (
	LineJoinMiter LineJoin = 0
	LineJoinRound LineJoin = 1
	LineJoinBevel LineJoin = 2
)

func marshalLineJoin(p uintptr) (interface{}, error) {
	return LineJoin(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextClusterFlags:
type TextClusterFlags int

const (
	TextClusterFlagsBackward TextClusterFlags = 1
)

func marshalTextClusterFlags(p uintptr) (interface{}, error) {
	return TextClusterFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontSlant:
type FontSlant int

const (
	FontSlantNormal  FontSlant = 0
	FontSlantItalic  FontSlant = 1
	FontSlantOblique FontSlant = 2
)

func marshalFontSlant(p uintptr) (interface{}, error) {
	return FontSlant(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontWeight:
type FontWeight int

const (
	FontWeightNormal FontWeight = 0
	FontWeightBold   FontWeight = 1
)

func marshalFontWeight(p uintptr) (interface{}, error) {
	return FontWeight(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SubpixelOrder:
type SubpixelOrder int

const (
	SubpixelOrderDefault SubpixelOrder = 0
	SubpixelOrderRGB     SubpixelOrder = 1
	SubpixelOrderBGR     SubpixelOrder = 2
	SubpixelOrderVRGB    SubpixelOrder = 3
	SubpixelOrderVBGR    SubpixelOrder = 4
)

func marshalSubpixelOrder(p uintptr) (interface{}, error) {
	return SubpixelOrder(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// HintStyle:
type HintStyle int

const (
	HintStyleDefault HintStyle = 0
	HintStyleNone    HintStyle = 1
	HintStyleSlight  HintStyle = 2
	HintStyleMedium  HintStyle = 3
	HintStyleFull    HintStyle = 4
)

func marshalHintStyle(p uintptr) (interface{}, error) {
	return HintStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// HintMetrics:
type HintMetrics int

const (
	HintMetricsDefault HintMetrics = 0
	HintMetricsOff     HintMetrics = 1
	HintMetricsOn      HintMetrics = 2
)

func marshalHintMetrics(p uintptr) (interface{}, error) {
	return HintMetrics(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontType:
type FontType int

const (
	FontTypeToy    FontType = 0
	FontTypeFT     FontType = 1
	FontTypeWin32  FontType = 2
	FontTypeQuartz FontType = 3
	FontTypeUser   FontType = 4
)

func marshalFontType(p uintptr) (interface{}, error) {
	return FontType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PathDataType:
type PathDataType int

const (
	PathDataTypeMoveTo    PathDataType = 0
	PathDataTypeLineTo    PathDataType = 1
	PathDataTypeCurveTo   PathDataType = 2
	PathDataTypeClosePath PathDataType = 3
)

func marshalPathDataType(p uintptr) (interface{}, error) {
	return PathDataType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DeviceType:
type DeviceType int

const (
	DeviceTypeDRM     DeviceType = 0
	DeviceTypeGL      DeviceType = 1
	DeviceTypeScript  DeviceType = 2
	DeviceTypeXcb     DeviceType = 3
	DeviceTypeXlib    DeviceType = 4
	DeviceTypeXML     DeviceType = 5
	DeviceTypeCOGL    DeviceType = 6
	DeviceTypeWin32   DeviceType = 7
	DeviceTypeInvalid DeviceType = -1
)

func marshalDeviceType(p uintptr) (interface{}, error) {
	return DeviceType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SurfaceType:
type SurfaceType int

const (
	SurfaceTypeImage         SurfaceType = 0
	SurfaceTypePDF           SurfaceType = 1
	SurfaceTypePS            SurfaceType = 2
	SurfaceTypeXlib          SurfaceType = 3
	SurfaceTypeXcb           SurfaceType = 4
	SurfaceTypeGlitz         SurfaceType = 5
	SurfaceTypeQuartz        SurfaceType = 6
	SurfaceTypeWin32         SurfaceType = 7
	SurfaceTypeBeOS          SurfaceType = 8
	SurfaceTypeDirectFB      SurfaceType = 9
	SurfaceTypeSvg           SurfaceType = 10
	SurfaceTypeOS2           SurfaceType = 11
	SurfaceTypeWin32Printing SurfaceType = 12
	SurfaceTypeQuartzImage   SurfaceType = 13
	SurfaceTypeScript        SurfaceType = 14
	SurfaceTypeQt            SurfaceType = 15
	SurfaceTypeRecording     SurfaceType = 16
	SurfaceTypeVg            SurfaceType = 17
	SurfaceTypeGL            SurfaceType = 18
	SurfaceTypeDRM           SurfaceType = 19
	SurfaceTypeTee           SurfaceType = 20
	SurfaceTypeXML           SurfaceType = 21
	SurfaceTypeSkia          SurfaceType = 22
	SurfaceTypeSubsurface    SurfaceType = 23
	SurfaceTypeCOGL          SurfaceType = 24
)

func marshalSurfaceType(p uintptr) (interface{}, error) {
	return SurfaceType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Format:
type Format int

const (
	FormatInvalid  Format = -1
	FormatARGB32   Format = 0
	FormatRGB24    Format = 1
	FormatA8       Format = 2
	FormatA1       Format = 3
	FormatRGB16565 Format = 4
	FormatRGB30    Format = 5
)

func marshalFormat(p uintptr) (interface{}, error) {
	return Format(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PatternType:
type PatternType int

const (
	PatternTypeSolid        PatternType = 0
	PatternTypeSurface      PatternType = 1
	PatternTypeLinear       PatternType = 2
	PatternTypeRadial       PatternType = 3
	PatternTypeMesh         PatternType = 4
	PatternTypeRasterSource PatternType = 5
)

func marshalPatternType(p uintptr) (interface{}, error) {
	return PatternType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Extend:
type Extend int

const (
	ExtendNone    Extend = 0
	ExtendRepeat  Extend = 1
	ExtendReflect Extend = 2
	ExtendPad     Extend = 3
)

func marshalExtend(p uintptr) (interface{}, error) {
	return Extend(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Filter:
type Filter int

const (
	FilterFast     Filter = 0
	FilterGood     Filter = 1
	FilterBest     Filter = 2
	FilterNearest  Filter = 3
	FilterBilinear Filter = 4
	FilterGaussian Filter = 5
)

func marshalFilter(p uintptr) (interface{}, error) {
	return Filter(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RegionOverlap:
type RegionOverlap int

const (
	RegionOverlapIn   RegionOverlap = 0
	RegionOverlapOut  RegionOverlap = 1
	RegionOverlapPart RegionOverlap = 2
)

func marshalRegionOverlap(p uintptr) (interface{}, error) {
	return RegionOverlap(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Context:
type Context C.cairo_t

// WrapContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapContext(ptr unsafe.Pointer) *Context {
	return (*Context)(ptr)
}

func marshalContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Context)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Context) Native() unsafe.Pointer {
	return unsafe.Pointer(c)
}

// Device:
type Device C.cairo_device_t

// WrapDevice wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDevice(ptr unsafe.Pointer) *Device {
	return (*Device)(ptr)
}

func marshalDevice(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Device)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *Device) Native() unsafe.Pointer {
	return unsafe.Pointer(d)
}

// Surface:
type Surface C.cairo_surface_t

// WrapSurface wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSurface(ptr unsafe.Pointer) *Surface {
	return (*Surface)(ptr)
}

func marshalSurface(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Surface)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *Surface) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}

// Matrix:
type Matrix C.cairo_matrix_t

// WrapMatrix wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMatrix(ptr unsafe.Pointer) *Matrix {
	return (*Matrix)(ptr)
}

// Native returns the underlying C source pointer.
func (m *Matrix) Native() unsafe.Pointer {
	return unsafe.Pointer(m)
}

// Pattern:
type Pattern C.cairo_pattern_t

// WrapPattern wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPattern(ptr unsafe.Pointer) *Pattern {
	return (*Pattern)(ptr)
}

func marshalPattern(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Pattern)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *Pattern) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// Region:
type Region C.cairo_region_t

// WrapRegion wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRegion(ptr unsafe.Pointer) *Region {
	return (*Region)(ptr)
}

func marshalRegion(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Region)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Region) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}

// FontOptions:
type FontOptions C.cairo_font_options_t

// WrapFontOptions wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontOptions(ptr unsafe.Pointer) *FontOptions {
	return (*FontOptions)(ptr)
}

func marshalFontOptions(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*FontOptions)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontOptions) Native() unsafe.Pointer {
	return unsafe.Pointer(f)
}

// FontFace:
type FontFace C.cairo_font_face_t

// WrapFontFace wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontFace(ptr unsafe.Pointer) *FontFace {
	return (*FontFace)(ptr)
}

func marshalFontFace(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*FontFace)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontFace) Native() unsafe.Pointer {
	return unsafe.Pointer(f)
}

// ScaledFont:
type ScaledFont C.cairo_scaled_font_t

// WrapScaledFont wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScaledFont(ptr unsafe.Pointer) *ScaledFont {
	return (*ScaledFont)(ptr)
}

func marshalScaledFont(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ScaledFont)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *ScaledFont) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}

// Path:
type Path C.cairo_path_t

// WrapPath wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPath(ptr unsafe.Pointer) *Path {
	return (*Path)(ptr)
}

// Native returns the underlying C source pointer.
func (p *Path) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// Rectangle:
type Rectangle C.cairo_rectangle_t

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Rectangle)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}

// RectangleInt:
type RectangleInt C.cairo_rectangle_int_t

// WrapRectangleInt wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangleInt(ptr unsafe.Pointer) *RectangleInt {
	return (*RectangleInt)(ptr)
}

func marshalRectangleInt(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*RectangleInt)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RectangleInt) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}
