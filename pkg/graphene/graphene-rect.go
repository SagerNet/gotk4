// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_rect_get_type()), F: marshalRect},
	})
}

// RectAlloc allocates a new #graphene_rect_t.
//
// The contents of the returned rectangle are undefined.
func RectAlloc() *Rect {
	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_alloc()

	ret1 = WrapRect(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Rect) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// RectZero returns a degenerate rectangle with origin fixed at (0, 0) and a
// size of 0, 0.
func RectZero() *Rect {
	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_zero()

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// Rect: the location and size of a rectangle region.
//
// The width and height of a #graphene_rect_t can be negative; for instance, a
// #graphene_rect_t with an origin of [ 0, 0 ] and a size of [ 10, 10 ] is
// equivalent to a #graphene_rect_t with an origin of [ 10, 10 ] and a size of [
// -10, -10 ].
//
// Application code can normalize rectangles using graphene_rect_normalize();
// this function will ensure that the width and height of a rectangle are
// positive values. All functions taking a #graphene_rect_t as an argument will
// internally operate on a normalized copy; all functions returning a
// #graphene_rect_t will always return a normalized rectangle.
type Rect struct {
	native C.graphene_rect_t
}

// WrapRect wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRect(ptr unsafe.Pointer) *Rect {
	if ptr == nil {
		return nil
	}

	return (*Rect)(ptr)
}

func marshalRect(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRect(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rect) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Origin gets the field inside the struct.
func (r *Rect) Origin() Point {
	v = WrapPoint(unsafe.Pointer(r.native.origin))
}

// Size gets the field inside the struct.
func (r *Rect) Size() Size {
	v = WrapSize(unsafe.Pointer(r.native.size))
}

// ContainsPoint checks whether a #graphene_rect_t contains the given
// coordinates.
func (r *Rect) ContainsPoint(p *Point) bool {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_rect_contains_point(arg0, p)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ContainsRect checks whether a #graphene_rect_t fully contains the given
// rectangle.
func (a *Rect) ContainsRect(b *Rect) bool {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_rect_contains_rect(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Equal checks whether the two given rectangle are equal.
func (a *Rect) Equal(b *Rect) bool {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_rect_equal(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Expand expands a #graphene_rect_t to contain the given #graphene_point_t.
func (r *Rect) Expand(p *Point) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var arg2 C.graphene_rect_t
	var ret2 *Rect

	C.graphene_rect_expand(arg0, p, &arg2)

	*ret2 = WrapRect(unsafe.Pointer(arg2))

	return ret2
}

// Free frees the resources allocated by graphene_rect_alloc().
func (r *Rect) Free() {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	C.graphene_rect_free(arg0)
}

// Area: compute the area of given normalized rectangle.
func (r *Rect) Area() float32 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_rect_get_area(arg0)

	ret1 = C.float(cret)

	return ret1
}

// BottomLeft retrieves the coordinates of the bottom-left corner of the given
// rectangle.
func (r *Rect) BottomLeft() Point {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_point_t
	var ret1 *Point

	C.graphene_rect_get_bottom_left(arg0, &arg1)

	*ret1 = WrapPoint(unsafe.Pointer(arg1))

	return ret1
}

// BottomRight retrieves the coordinates of the bottom-right corner of the given
// rectangle.
func (r *Rect) BottomRight() Point {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_point_t
	var ret1 *Point

	C.graphene_rect_get_bottom_right(arg0, &arg1)

	*ret1 = WrapPoint(unsafe.Pointer(arg1))

	return ret1
}

// Center retrieves the coordinates of the center of the given rectangle.
func (r *Rect) Center() Point {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_point_t
	var ret1 *Point

	C.graphene_rect_get_center(arg0, &arg1)

	*ret1 = WrapPoint(unsafe.Pointer(arg1))

	return ret1
}

// Height retrieves the normalized height of the given rectangle.
func (r *Rect) Height() float32 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_rect_get_height(arg0)

	ret1 = C.float(cret)

	return ret1
}

// TopLeft retrieves the coordinates of the top-left corner of the given
// rectangle.
func (r *Rect) TopLeft() Point {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_point_t
	var ret1 *Point

	C.graphene_rect_get_top_left(arg0, &arg1)

	*ret1 = WrapPoint(unsafe.Pointer(arg1))

	return ret1
}

// TopRight retrieves the coordinates of the top-right corner of the given
// rectangle.
func (r *Rect) TopRight() Point {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_point_t
	var ret1 *Point

	C.graphene_rect_get_top_right(arg0, &arg1)

	*ret1 = WrapPoint(unsafe.Pointer(arg1))

	return ret1
}

// Vertices computes the four vertices of a #graphene_rect_t.
func (r *Rect) Vertices() [4]Vec2 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 [4]C.graphene_vec2_t
	var ret1 [4]Vec2

	C.graphene_rect_get_vertices(arg0, &arg1)

	{
		tmp := *(*[4]Vec2)(unsafe.Pointer(&arg1))
		for i := 0; i < 4; i++ {
			src := tmp[i]
			ret1[i] = WrapVec2(unsafe.Pointer(src))
		}
	}

	return ret1
}

// Width retrieves the normalized width of the given rectangle.
func (r *Rect) Width() float32 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_rect_get_width(arg0)

	ret1 = C.float(cret)

	return ret1
}

// X retrieves the normalized X coordinate of the origin of the given rectangle.
func (r *Rect) X() float32 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_rect_get_x(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Y retrieves the normalized Y coordinate of the origin of the given rectangle.
func (r *Rect) Y() float32 {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret C.float
	var ret1 float32

	cret = C.graphene_rect_get_y(arg0)

	ret1 = C.float(cret)

	return ret1
}

// Init initializes the given #graphene_rect_t with the given values.
//
// This function will implicitly normalize the #graphene_rect_t before
// returning.
func (r *Rect) Init(x float32, y float32, width float32, height float32) *Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float
	var arg3 C.float
	var arg4 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)
	arg3 = C.float(width)
	arg4 = C.float(height)

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_init(arg0, x, y, width, height)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// InitFromRect initializes @r using the given @src rectangle.
//
// This function will implicitly normalize the #graphene_rect_t before
// returning.
func (r *Rect) InitFromRect(src *Rect) *Rect {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_init_from_rect(arg0, src)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// Inset changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive @d_x or @d_y values; to create a
// larger, encompassing rectangle, use negative @d_x or @d_y values.
//
// The origin of the rectangle is offset by @d_x and @d_y, while the size is
// adjusted by `(2 * @d_x, 2 * @d_y)`. If @d_x and @d_y are positive values, the
// size of the rectangle is decreased; if @d_x and @d_y are negative values, the
// size of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or height
// then the size will be set to zero.
func (r *Rect) Inset(dX float32, dY float32) *Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(dX)
	arg2 = C.float(dY)

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_inset(arg0, dX, dY)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// InsetR changes the given rectangle to be smaller, or larger depending on the
// given inset parameters.
//
// To create an inset rectangle, use positive @d_x or @d_y values; to create a
// larger, encompassing rectangle, use negative @d_x or @d_y values.
//
// The origin of the rectangle is offset by @d_x and @d_y, while the size is
// adjusted by `(2 * @d_x, 2 * @d_y)`. If @d_x and @d_y are positive values, the
// size of the rectangle is decreased; if @d_x and @d_y are negative values, the
// size of the rectangle is increased.
//
// If the size of the resulting inset rectangle has a negative width or height
// then the size will be set to zero.
func (r *Rect) InsetR(dX float32, dY float32) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(dX)
	arg2 = C.float(dY)

	var arg3 C.graphene_rect_t
	var ret3 *Rect

	C.graphene_rect_inset_r(arg0, dX, dY, &arg3)

	*ret3 = WrapRect(unsafe.Pointer(arg3))

	return ret3
}

// Interpolate: linearly interpolates the origin and size of the two given
// rectangles.
func (a *Rect) Interpolate(b *Rect, factor float64) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t
	var arg2 C.double

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(b.Native()))
	arg2 = C.double(factor)

	var arg3 C.graphene_rect_t
	var ret3 *Rect

	C.graphene_rect_interpolate(arg0, b, factor, &arg3)

	*ret3 = WrapRect(unsafe.Pointer(arg3))

	return ret3
}

// Intersection computes the intersection of the two given rectangles.
//
// ! (rectangle-intersection.png)
//
// The intersection in the image above is the blue outline.
//
// If the two rectangles do not intersect, @res will contain a degenerate
// rectangle with origin in (0, 0) and a size of 0.
func (a *Rect) Intersection(b *Rect) (res Rect, ok bool) {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_rect_t
	var ret2 *Rect
	var cret C._Bool
	var ret2 bool

	cret = C.graphene_rect_intersection(arg0, b, &arg2)

	*ret2 = WrapRect(unsafe.Pointer(arg2))
	ret2 = C.bool(cret) != C.false

	return ret2, ret2
}

// Normalize normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of positive
// values, and that the origin is the top-left corner of the rectangle.
func (r *Rect) Normalize() *Rect {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_normalize(arg0)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// NormalizeR normalizes the passed rectangle.
//
// This function ensures that the size of the rectangle is made of positive
// values, and that the origin is in the top-left corner of the rectangle.
func (r *Rect) NormalizeR() Rect {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_rect_t
	var ret1 *Rect

	C.graphene_rect_normalize_r(arg0, &arg1)

	*ret1 = WrapRect(unsafe.Pointer(arg1))

	return ret1
}

// Offset offsets the origin by @d_x and @d_y.
//
// The size of the rectangle is unchanged.
func (r *Rect) Offset(dX float32, dY float32) *Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(dX)
	arg2 = C.float(dY)

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_offset(arg0, dX, dY)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// OffsetR offsets the origin of the given rectangle by @d_x and @d_y.
//
// The size of the rectangle is left unchanged.
func (r *Rect) OffsetR(dX float32, dY float32) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(dX)
	arg2 = C.float(dY)

	var arg3 C.graphene_rect_t
	var ret3 *Rect

	C.graphene_rect_offset_r(arg0, dX, dY, &arg3)

	*ret3 = WrapRect(unsafe.Pointer(arg3))

	return ret3
}

// Round rounds the origin and size of the given rectangle to their nearest
// integer values; the rounding is guaranteed to be large enough to have an area
// bigger or equal to the original rectangle, but might not fully contain its
// extents. Use graphene_rect_round_extents() in case you need to round to a
// rectangle that covers fully the original one.
//
// This function is the equivalent of calling `floor` on the coordinates of the
// origin, and `ceil` on the size.
func (r *Rect) Round() Rect {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_rect_t
	var ret1 *Rect

	C.graphene_rect_round(arg0, &arg1)

	*ret1 = WrapRect(unsafe.Pointer(arg1))

	return ret1
}

// RoundExtents rounds the origin of the given rectangle to its nearest integer
// value and and recompute the size so that the rectangle is large enough to
// contain all the conrners of the original rectangle.
//
// This function is the equivalent of calling `floor` on the coordinates of the
// origin, and recomputing the size calling `ceil` on the bottom-right
// coordinates.
//
// If you want to be sure that the rounded rectangle completely covers the area
// that was covered by the original rectangle — i.e. you want to cover the area
// including all its corners — this function will make sure that the size is
// recomputed taking into account the ceiling of the coordinates of the
// bottom-right corner. If the difference between the original coordinates and
// the coordinates of the rounded rectangle is greater than the difference
// between the original size and and the rounded size, then the move of the
// origin would not be compensated by a move in the anti-origin, leaving the
// corners of the original rectangle outside the rounded one.
func (r *Rect) RoundExtents() Rect {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var arg1 C.graphene_rect_t
	var ret1 *Rect

	C.graphene_rect_round_extents(arg0, &arg1)

	*ret1 = WrapRect(unsafe.Pointer(arg1))

	return ret1
}

// RoundToPixel rounds the origin and the size of the given rectangle to their
// nearest integer values; the rounding is guaranteed to be large enough to
// contain the original rectangle.
func (r *Rect) RoundToPixel() *Rect {
	var arg0 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	var cret *C.graphene_rect_t
	var ret1 *Rect

	cret = C.graphene_rect_round_to_pixel(arg0)

	ret1 = WrapRect(unsafe.Pointer(cret))

	return ret1
}

// Scale scales the size and origin of a rectangle horizontaly by @s_h, and
// vertically by @s_v. The result @res is normalized.
func (r *Rect) Scale(sH float32, sV float32) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))
	arg1 = C.float(sH)
	arg2 = C.float(sV)

	var arg3 C.graphene_rect_t
	var ret3 *Rect

	C.graphene_rect_scale(arg0, sH, sV, &arg3)

	*ret3 = WrapRect(unsafe.Pointer(arg3))

	return ret3
}

// Union computes the union of the two given rectangles.
//
// ! (rectangle-union.png)
//
// The union in the image above is the blue outline.
func (a *Rect) Union(b *Rect) Rect {
	var arg0 *C.graphene_rect_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_rect_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(b.Native()))

	var arg2 C.graphene_rect_t
	var ret2 *Rect

	C.graphene_rect_union(arg0, b, &arg2)

	*ret2 = WrapRect(unsafe.Pointer(arg2))

	return ret2
}