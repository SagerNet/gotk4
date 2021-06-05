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
		{T: externglib.Type(C.graphene_point_get_type()), F: marshalPoint},
	})
}

// PointZero returns a point fixed at (0, 0).
func PointZero() *Point {
	var cret *C.graphene_point_t
	var goret1 *Point

	cret = C.graphene_point_zero()

	goret1 = WrapPoint(unsafe.Pointer(cret))

	return goret1
}

// Point: a point with two coordinates.
type Point struct {
	native C.graphene_point_t
}

// WrapPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint(ptr unsafe.Pointer) *Point {
	if ptr == nil {
		return nil
	}

	return (*Point)(ptr)
}

func marshalPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPoint(unsafe.Pointer(b)), nil
}

// NewPointAlloc constructs a struct Point.
func NewPointAlloc() *Point {
	var cret *C.graphene_point_t
	var goret1 *Point

	cret = C.graphene_point_alloc()

	goret1 = WrapPoint(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *Point) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// Native returns the underlying C source pointer.
func (p *Point) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// X gets the field inside the struct.
func (p *Point) X() float32 {
	v = C.float(p.native.x)
}

// Y gets the field inside the struct.
func (p *Point) Y() float32 {
	v = C.float(p.native.y)
}

// Distance computes the distance between @a and @b.
func (a *Point) Distance(b *Point) (dX float32, dY float32, gfloat float32) {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))

	var arg2 *C.float
	var ret2 float32
	var arg3 *C.float
	var ret3 float32
	var cret C.float
	var goret3 float32

	cret = C.graphene_point_distance(arg0, b, &arg2, &arg3)

	ret2 = *C.float(arg2)
	ret3 = *C.float(arg3)
	goret3 = C.float(cret)

	return ret2, ret3, goret3
}

// Equal checks if the two points @a and @b point to the same coordinates.
//
// This function accounts for floating point fluctuations; if you want to
// control the fuzziness of the match, you can use graphene_point_near()
// instead.
func (a *Point) Equal(b *Point) bool {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var goret1 bool

	cret = C.graphene_point_equal(arg0, b)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Free frees the resources allocated by graphene_point_alloc().
func (p *Point) Free() {
	var arg0 *C.graphene_point_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	C.graphene_point_free(arg0)
}

// Init initializes @p to the given @x and @y coordinates.
//
// It's safe to call this function multiple times.
func (p *Point) Init(x float32, y float32) *Point {
	var arg0 *C.graphene_point_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	arg1 = C.float(x)
	arg2 = C.float(y)

	var cret *C.graphene_point_t
	var goret1 *Point

	cret = C.graphene_point_init(arg0, x, y)

	goret1 = WrapPoint(unsafe.Pointer(cret))

	return goret1
}

// InitFromPoint initializes @p with the same coordinates of @src.
func (p *Point) InitFromPoint(src *Point) *Point {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_point_t
	var goret1 *Point

	cret = C.graphene_point_init_from_point(arg0, src)

	goret1 = WrapPoint(unsafe.Pointer(cret))

	return goret1
}

// InitFromVec2 initializes @p with the coordinates inside the given
// #graphene_vec2_t.
func (p *Point) InitFromVec2(src *Vec2) *Point {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_vec2_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_point_t
	var goret1 *Point

	cret = C.graphene_point_init_from_vec2(arg0, src)

	goret1 = WrapPoint(unsafe.Pointer(cret))

	return goret1
}

// Interpolate: linearly interpolates the coordinates of @a and @b using the
// given @factor.
func (a *Point) Interpolate(b *Point, factor float64) Point {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_point_t
	var arg2 C.double

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))
	arg2 = C.double(factor)

	var arg3 *C.graphene_point_t
	var ret3 *Point

	C.graphene_point_interpolate(arg0, b, factor, &arg3)

	ret3 = WrapPoint(unsafe.Pointer(arg3))

	return ret3
}

// Near checks whether the two points @a and @b are within the threshold of
// @epsilon.
func (a *Point) Near(b *Point, epsilon float32) bool {
	var arg0 *C.graphene_point_t
	var arg1 *C.graphene_point_t
	var arg2 C.float

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))
	arg2 = C.float(epsilon)

	var cret C._Bool
	var goret1 bool

	cret = C.graphene_point_near(arg0, b, epsilon)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// ToVec2 stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) ToVec2() Vec2 {
	var arg0 *C.graphene_point_t

	arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var arg1 *C.graphene_vec2_t
	var ret1 *Vec2

	C.graphene_point_to_vec2(arg0, &arg1)

	ret1 = WrapVec2(unsafe.Pointer(arg1))

	return ret1
}
