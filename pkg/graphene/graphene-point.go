// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_point_get_type()), F: marshalPoint},
	})
}

// Point: a point with two coordinates.
type Point C.graphene_point_t

// WrapPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint(ptr unsafe.Pointer) *Point {
	return (*Point)(ptr)
}

func marshalPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Point)(unsafe.Pointer(b)), nil
}

// NewPointAlloc constructs a struct Point.
func NewPointAlloc() *Point {
	var _cret *C.graphene_point_t // in

	_cret = C.graphene_point_alloc()

	var _point *Point // out

	_point = (*Point)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_point, func(v **Point) {
		C.free(unsafe.Pointer(v))
	})

	return _point
}

// Native returns the underlying C source pointer.
func (p *Point) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// Distance stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Distance(b *Point) (dX float32, dY float32, gfloat float32) {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.float             // in
	var _arg3 C.float             // in
	var _cret C.float             // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_point_distance(_arg0, _arg1, &_arg2, &_arg3)

	var _dX float32     // out
	var _dY float32     // out
	var _gfloat float32 // out

	_dX = float32(_arg2)
	_dY = float32(_arg3)
	_gfloat = float32(_cret)

	return _dX, _dY, _gfloat
}

// Equal stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Equal(b *Point) bool {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_point_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Free() {
	var _arg0 *C.graphene_point_t // out

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	C.graphene_point_free(_arg0)
}

// Init stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Init(x float32, y float32) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 C.float             // out
	var _arg2 C.float             // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	_arg1 = C.float(x)
	_arg2 = C.float(y)

	_cret = C.graphene_point_init(_arg0, _arg1, _arg2)

	var _point *Point // out

	_point = (*Point)(unsafe.Pointer(_cret))

	return _point
}

// InitFromPoint stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) InitFromPoint(src *Point) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(src.Native()))

	_cret = C.graphene_point_init_from_point(_arg0, _arg1)

	var _point *Point // out

	_point = (*Point)(unsafe.Pointer(_cret))

	return _point
}

// InitFromVec2 stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) InitFromVec2(src *Vec2) *Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_vec2_t  // out
	var _cret *C.graphene_point_t // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.graphene_vec2_t)(unsafe.Pointer(src.Native()))

	_cret = C.graphene_point_init_from_vec2(_arg0, _arg1)

	var _point *Point // out

	_point = (*Point)(unsafe.Pointer(_cret))

	return _point
}

// Interpolate stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Interpolate(b *Point, factor float64) Point {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.double            // out
	var _arg3 C.graphene_point_t  // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.double(factor)

	C.graphene_point_interpolate(_arg0, _arg1, _arg2, &_arg3)

	var _res Point // out

	{
		var refTmpIn *C.graphene_point_t
		var refTmpOut *Point

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = (*Point)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Near stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) Near(b *Point, epsilon float32) bool {
	var _arg0 *C.graphene_point_t // out
	var _arg1 *C.graphene_point_t // out
	var _arg2 C.float             // out
	var _cret C._Bool             // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_point_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.float(epsilon)

	_cret = C.graphene_point_near(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ToVec2 stores the coordinates of the given #graphene_point_t into a
// #graphene_vec2_t.
func (p *Point) ToVec2() Vec2 {
	var _arg0 *C.graphene_point_t // out
	var _arg1 C.graphene_vec2_t   // in

	_arg0 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	C.graphene_point_to_vec2(_arg0, &_arg1)

	var _v Vec2 // out

	{
		var refTmpIn *C.graphene_vec2_t
		var refTmpOut *Vec2

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Vec2)(unsafe.Pointer(refTmpIn))

		_v = *refTmpOut
	}

	return _v
}
