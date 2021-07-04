// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_axis_use_get_type()), F: marshalAxisUse},
		{T: externglib.Type(C.gdk_gl_error_get_type()), F: marshalGLError},
		{T: externglib.Type(C.gdk_gravity_get_type()), F: marshalGravity},
		{T: externglib.Type(C.gdk_vulkan_error_get_type()), F: marshalVulkanError},
		{T: externglib.Type(C.gdk_axis_flags_get_type()), F: marshalAxisFlags},
		{T: externglib.Type(C.gdk_drag_action_get_type()), F: marshalDragAction},
		{T: externglib.Type(C.gdk_modifier_type_get_type()), F: marshalModifierType},
		{T: externglib.Type(C.gdk_content_formats_get_type()), F: marshalContentFormats},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// AxisUse defines how device axes are interpreted by GTK.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse int

const (
	// ignore: the axis is ignored.
	AxisUseIgnore AxisUse = 0
	// x: the axis is used as the x axis.
	AxisUseX AxisUse = 1
	// y: the axis is used as the y axis.
	AxisUseY AxisUse = 2
	// DeltaX: the axis is used as the scroll x delta
	AxisUseDeltaX AxisUse = 3
	// DeltaY: the axis is used as the scroll y delta
	AxisUseDeltaY AxisUse = 4
	// pressure: the axis is used for pressure information.
	AxisUsePressure AxisUse = 5
	// xtilt: the axis is used for x tilt information.
	AxisUseXtilt AxisUse = 6
	// ytilt: the axis is used for y tilt information.
	AxisUseYtilt AxisUse = 7
	// wheel: the axis is used for wheel information.
	AxisUseWheel AxisUse = 8
	// distance: the axis is used for pen/tablet distance information
	AxisUseDistance AxisUse = 9
	// rotation: the axis is used for pen rotation information
	AxisUseRotation AxisUse = 10
	// slider: the axis is used for pen slider information
	AxisUseSlider AxisUse = 11
	// last: a constant equal to the numerically highest axis value.
	AxisUseLast AxisUse = 12
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GLError: error enumeration for `GdkGLContext`.
type GLError int

const (
	// NotAvailable: openGL support is not available
	GLErrorNotAvailable GLError = 0
	// UnsupportedFormat: the requested visual format is not supported
	GLErrorUnsupportedFormat GLError = 1
	// UnsupportedProfile: the requested profile is not supported
	GLErrorUnsupportedProfile GLError = 2
	// CompilationFailed: the shader compilation failed
	GLErrorCompilationFailed GLError = 3
	// LinkFailed: the shader linking failed
	GLErrorLinkFailed GLError = 4
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Gravity defines the reference point of a surface and is used in PopupLayout.
type Gravity int

const (
	// NorthWest: the reference point is at the top left corner.
	GravityNorthWest Gravity = 1
	// north: the reference point is in the middle of the top edge.
	GravityNorth Gravity = 2
	// NorthEast: the reference point is at the top right corner.
	GravityNorthEast Gravity = 3
	// west: the reference point is at the middle of the left edge.
	GravityWest Gravity = 4
	// center: the reference point is at the center of the surface.
	GravityCenter Gravity = 5
	// east: the reference point is at the middle of the right edge.
	GravityEast Gravity = 6
	// SouthWest: the reference point is at the lower left corner.
	GravitySouthWest Gravity = 7
	// south: the reference point is at the middle of the lower edge.
	GravitySouth Gravity = 8
	// SouthEast: the reference point is at the lower right corner.
	GravitySouthEast Gravity = 9
	// static: the reference point is at the top left corner of the surface
	// itself, ignoring window manager decorations.
	GravityStatic Gravity = 10
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// VulkanError: error enumeration for VulkanContext.
type VulkanError int

const (
	// unsupported: vulkan is not supported on this backend or has not been
	// compiled in.
	VulkanErrorUnsupported VulkanError = 0
	// NotAvailable: vulkan support is not available on this Surface
	VulkanErrorNotAvailable VulkanError = 1
)

func marshalVulkanError(p uintptr) (interface{}, error) {
	return VulkanError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags int

const (
	// AxisFlagsX: x axis is present
	AxisFlagsX AxisFlags = 0b10
	// AxisFlagsY: y axis is present
	AxisFlagsY AxisFlags = 0b100
	// AxisFlagsDeltaX: scroll X delta axis is present
	AxisFlagsDeltaX AxisFlags = 0b1000
	// AxisFlagsDeltaY: scroll Y delta axis is present
	AxisFlagsDeltaY AxisFlags = 0b10000
	// AxisFlagsPressure: pressure axis is present
	AxisFlagsPressure AxisFlags = 0b100000
	// AxisFlagsXtilt: x tilt axis is present
	AxisFlagsXtilt AxisFlags = 0b1000000
	// AxisFlagsYtilt: y tilt axis is present
	AxisFlagsYtilt AxisFlags = 0b10000000
	// AxisFlagsWheel: wheel axis is present
	AxisFlagsWheel AxisFlags = 0b100000000
	// AxisFlagsDistance: distance axis is present
	AxisFlagsDistance AxisFlags = 0b1000000000
	// AxisFlagsRotation z-axis rotation is present
	AxisFlagsRotation AxisFlags = 0b10000000000
	// AxisFlagsSlider: slider axis is present
	AxisFlagsSlider AxisFlags = 0b100000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragAction: used in `GdkDrop` and `GdkDrag` to indicate the actions that the
// destination can and should do with the dropped data.
type DragAction int

const (
	// DragActionCopy: copy the data.
	DragActionCopy DragAction = 0b1
	// DragActionMove: move the data, i.e. first copy it, then delete it from
	// the source using the DELETE target of the X selection protocol.
	DragActionMove DragAction = 0b10
	// DragActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means, and is not supported on
	// all platforms.
	DragActionLink DragAction = 0b100
	// DragActionAsk: ask the user what to do with the data.
	DragActionAsk DragAction = 0b1000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ModifierType flags to indicate the state of modifier keys and mouse buttons
// in events.
//
// Typical modifier keys are Shift, Control, Meta, Super, Hyper, Alt, Compose,
// Apple, CapsLock or ShiftLock.
//
// Note that GDK may add internal values to events which include values outside
// of this enumeration. Your code should preserve and ignore them. You can use
// GDK_MODIFIER_MASK to remove all private values.
type ModifierType int

const (
	// ModifierTypeShiftMask: the Shift key.
	ModifierTypeShiftMask ModifierType = 0b1
	// ModifierTypeLockMask: a Lock key (depending on the modifier mapping of
	// the X server this may either be CapsLock or ShiftLock).
	ModifierTypeLockMask ModifierType = 0b10
	// ModifierTypeControlMask: the Control key.
	ModifierTypeControlMask ModifierType = 0b100
	// ModifierTypeAltMask: the fourth modifier key (it depends on the modifier
	// mapping of the X server which key is interpreted as this modifier, but
	// normally it is the Alt key).
	ModifierTypeAltMask ModifierType = 0b1000
	// ModifierTypeButton1Mask: the first mouse button.
	ModifierTypeButton1Mask ModifierType = 0b100000000
	// ModifierTypeButton2Mask: the second mouse button.
	ModifierTypeButton2Mask ModifierType = 0b1000000000
	// ModifierTypeButton3Mask: the third mouse button.
	ModifierTypeButton3Mask ModifierType = 0b10000000000
	// ModifierTypeButton4Mask: the fourth mouse button.
	ModifierTypeButton4Mask ModifierType = 0b100000000000
	// ModifierTypeButton5Mask: the fifth mouse button.
	ModifierTypeButton5Mask ModifierType = 0b1000000000000
	// ModifierTypeSuperMask: the Super modifier
	ModifierTypeSuperMask ModifierType = 0b100000000000000000000000000
	// ModifierTypeHyperMask: the Hyper modifier
	ModifierTypeHyperMask ModifierType = 0b1000000000000000000000000000
	// ModifierTypeMetaMask: the Meta modifier
	ModifierTypeMetaMask ModifierType = 0b10000000000000000000000000000
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ContentFormats: the `GdkContentFormats` structure is used to advertise and
// negotiate the format of content.
//
// You will encounter `GdkContentFormats` when interacting with objects
// controlling operations that pass data between different widgets, window or
// application, like [class@Gdk.Drag], [class@Gdk.Drop], [class@Gdk.Clipboard]
// or [class@Gdk.ContentProvider].
//
// GDK supports content in 2 forms: `GType` and mime type. Using `GTypes` is
// meant only for in-process content transfers. Mime types are meant to be used
// for data passing both in-process and out-of-process. The details of how data
// is passed is described in the documentation of the actual implementations. To
// transform between the two forms, [class@Gdk.ContentSerializer] and
// [class@Gdk.ContentDeserializer] are used.
//
// A `GdkContentFormats` describes a set of possible formats content can be
// exchanged in. It is assumed that this set is ordered. `GTypes` are more
// important than mime types. Order between different `GTypes` or mime types is
// the order they were added in, most important first. Functions that care about
// order, such as [method@Gdk.ContentFormats.union], will describe in their
// documentation how they interpret that order, though in general the order of
// the first argument is considered the primary order of the result, followed by
// the order of further arguments.
//
// For debugging purposes, the function [method@Gdk.ContentFormats.to_string]
// exists. It will print a comma-separated list of formats from most important
// to least important.
//
// `GdkContentFormats` is an immutable struct. After creation, you cannot change
// the types it represents. Instead, new `GdkContentFormats` have to be created.
// The [struct@Gdk.ContentFormatsBuilder]` structure is meant to help in this
// endeavor.
type ContentFormats C.GdkContentFormats

// WrapContentFormats wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapContentFormats(ptr unsafe.Pointer) *ContentFormats {
	return (*ContentFormats)(ptr)
}

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ContentFormats)(unsafe.Pointer(b)), nil
}

// NewContentFormats constructs a struct ContentFormats.
func NewContentFormats(mimeTypes []string) *ContentFormats {
	var _arg1 **C.char
	var _arg2 C.guint
	var _cret *C.GdkContentFormats // in

	_arg2 = C.guint(len(mimeTypes))
	_arg1 = (**C.char)(C.malloc(C.ulong(len(mimeTypes)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(mimeTypes))
		for i := range mimeTypes {
			out[i] = (*C.char)(C.CString(mimeTypes[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.gdk_content_formats_new(_arg1, _arg2)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// NewContentFormatsForGType constructs a struct ContentFormats.
func NewContentFormatsForGType(typ externglib.Type) *ContentFormats {
	var _arg1 C.GType              // out
	var _cret *C.GdkContentFormats // in

	_arg1 = (C.GType)(typ)

	_cret = C.gdk_content_formats_new_for_gtype(_arg1)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// Native returns the underlying C source pointer.
func (c *ContentFormats) Native() unsafe.Pointer {
	return unsafe.Pointer(c)
}

// ContainGType decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) ContainGType(typ externglib.Type) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 C.GType              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (C.GType)(typ)

	_cret = C.gdk_content_formats_contain_gtype(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainMIMEType decreases the reference count of a `GdkContentFormats` by
// one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) ContainMIMEType(mimeType string) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.char              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_content_formats_contain_mime_type(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Match decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) Match(second *ContentFormats) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	_cret = C.gdk_content_formats_match(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchGType decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) MatchGType(second *ContentFormats) externglib.Type {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.GType              // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	_cret = C.gdk_content_formats_match_gtype(_arg0, _arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// MatchMIMEType decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) MatchMIMEType(second *ContentFormats) string {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	_cret = C.gdk_content_formats_match_mime_type(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Ref decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) Ref() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_ref(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// String decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) String() string {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Union decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) Union(second *ContentFormats) *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	_cret = C.gdk_content_formats_union(_arg0, _arg1)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// UnionDeserializeGTypes decreases the reference count of a `GdkContentFormats`
// by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) UnionDeserializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_union_deserialize_gtypes(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// UnionDeserializeMIMETypes decreases the reference count of a
// `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) UnionDeserializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_union_deserialize_mime_types(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// UnionSerializeGTypes decreases the reference count of a `GdkContentFormats`
// by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) UnionSerializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_union_serialize_gtypes(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// UnionSerializeMIMETypes decreases the reference count of a
// `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) UnionSerializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_content_formats_union_serialize_mime_types(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_contentFormats, func(v **ContentFormats) {
		C.free(unsafe.Pointer(v))
	})

	return _contentFormats
}

// Unref decreases the reference count of a `GdkContentFormats` by one.
//
// If the resulting reference count is zero, frees the formats.
func (f *ContentFormats) Unref() {
	var _arg0 *C.GdkContentFormats // out

	_arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_unref(_arg0)
}

// KeymapKey: a `GdkKeymapKey` is a hardware key that can be mapped to a keyval.
type KeymapKey C.GdkKeymapKey

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	return (*KeymapKey)(ptr)
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(k)
}

// Rectangle: a `GdkRectangle` data type for representing rectangles.
//
// `GdkRectangle` is identical to `cairo_rectangle_t`. Together with Cairo’s
// `cairo_region_t` data type, these are the central types for representing sets
// of pixels.
//
// The intersection of two rectangles can be computed with
// [method@Gdk.Rectangle.intersect]; to find the union of two rectangles use
// [method@Gdk.Rectangle.union].
//
// The `cairo_region_t` type provided by Cairo is usually used for managing
// non-rectangular clipping of graphical operations.
//
// The Graphene library has a number of other data types for regions and volumes
// in 2D and 3D.
type Rectangle C.GdkRectangle

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

// ContainsPoint calculates the union of two rectangles.
//
// The union of rectangles @src1 and @src2 is the smallest rectangle which
// includes both @src1 and @src2 within it. It is allowed for @dest to be the
// same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) ContainsPoint(x int, y int) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gdk_rectangle_contains_point(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Equal calculates the union of two rectangles.
//
// The union of rectangles @src1 and @src2 is the smallest rectangle which
// includes both @src1 and @src2 within it. It is allowed for @dest to be the
// same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Equal(rect2 *Rectangle) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect2.Native()))

	_cret = C.gdk_rectangle_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Intersect calculates the union of two rectangles.
//
// The union of rectangles @src1 and @src2 is the smallest rectangle which
// includes both @src1 and @src2 within it. It is allowed for @dest to be the
// same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Intersect(src2 *Rectangle) (Rectangle, bool) {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	_cret = C.gdk_rectangle_intersect(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out
	var _ok bool        // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *Rectangle

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Rectangle)(unsafe.Pointer(refTmpIn))

		_dest = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _dest, _ok
}

// Union calculates the union of two rectangles.
//
// The union of rectangles @src1 and @src2 is the smallest rectangle which
// includes both @src1 and @src2 within it. It is allowed for @dest to be the
// same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Union(src2 *Rectangle) Rectangle {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in

	_arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	C.gdk_rectangle_union(_arg0, _arg1, &_arg2)

	var _dest Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *Rectangle

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Rectangle)(unsafe.Pointer(refTmpIn))

		_dest = *refTmpOut
	}

	return _dest
}