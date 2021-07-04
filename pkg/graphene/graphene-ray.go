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
		{T: externglib.Type(C.graphene_ray_get_type()), F: marshalRay},
	})
}

// RayIntersectionKind: the type of intersection.
type RayIntersectionKind int

const (
	// none: no intersection
	RayIntersectionKindNone RayIntersectionKind = 0
	// enter: the ray is entering the intersected object
	RayIntersectionKindEnter RayIntersectionKind = 1
	// leave: the ray is leaving the intersected object
	RayIntersectionKindLeave RayIntersectionKind = 2
)

// Ray: a ray emitted from an origin in a given direction.
//
// The contents of the `graphene_ray_t` structure are private, and should not be
// modified directly.
type Ray C.graphene_ray_t

// WrapRay wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRay(ptr unsafe.Pointer) *Ray {
	return (*Ray)(ptr)
}

func marshalRay(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Ray)(unsafe.Pointer(b)), nil
}

// NewRayAlloc constructs a struct Ray.
func NewRayAlloc() *Ray {
	var _cret *C.graphene_ray_t // in

	_cret = C.graphene_ray_alloc()

	var _ray *Ray // out

	_ray = (*Ray)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_ray, func(v **Ray) {
		C.free(unsafe.Pointer(v))
	})

	return _ray
}

// Native returns the underlying C source pointer.
func (r *Ray) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}

// Equal checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) Equal(b *Ray) bool {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_ray_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_ray_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_ray_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) Free() {
	var _arg0 *C.graphene_ray_t // out

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	C.graphene_ray_free(_arg0)
}

// ClosestPointToPoint checks whether the given #graphene_ray_t @r intersects
// the given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) ClosestPointToPoint(p *Point3D) Point3D {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 C.graphene_point3d_t  // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	C.graphene_ray_get_closest_point_to_point(_arg0, _arg1, &_arg2)

	var _res Point3D // out

	{
		var refTmpIn *C.graphene_point3d_t
		var refTmpOut *Point3D

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Point3D)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Direction checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) Direction() Vec3 {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 C.graphene_vec3_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	C.graphene_ray_get_direction(_arg0, &_arg1)

	var _direction Vec3 // out

	{
		var refTmpIn *C.graphene_vec3_t
		var refTmpOut *Vec3

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Vec3)(unsafe.Pointer(refTmpIn))

		_direction = *refTmpOut
	}

	return _direction
}

// DistanceToPlane checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) DistanceToPlane(p *Plane) float32 {
	var _arg0 *C.graphene_ray_t   // out
	var _arg1 *C.graphene_plane_t // out
	var _cret C.float             // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_plane_t)(unsafe.Pointer(p.Native()))

	_cret = C.graphene_ray_get_distance_to_plane(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// DistanceToPoint checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) DistanceToPoint(p *Point3D) float32 {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _cret C.float               // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(p.Native()))

	_cret = C.graphene_ray_get_distance_to_point(_arg0, _arg1)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Origin checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) Origin() Point3D {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))

	C.graphene_ray_get_origin(_arg0, &_arg1)

	var _origin Point3D // out

	{
		var refTmpIn *C.graphene_point3d_t
		var refTmpOut *Point3D

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Point3D)(unsafe.Pointer(refTmpIn))

		_origin = *refTmpOut
	}

	return _origin
}

// PositionAt checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) PositionAt(t float32) Point3D {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 C.float              // out
	var _arg2 C.graphene_point3d_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = C.float(t)

	C.graphene_ray_get_position_at(_arg0, _arg1, &_arg2)

	var _position Point3D // out

	{
		var refTmpIn *C.graphene_point3d_t
		var refTmpOut *Point3D

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Point3D)(unsafe.Pointer(refTmpIn))

		_position = *refTmpOut
	}

	return _position
}

// Init checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) Init(origin *Point3D, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t     // out
	var _arg1 *C.graphene_point3d_t // out
	var _arg2 *C.graphene_vec3_t    // out
	var _cret *C.graphene_ray_t     // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(origin.Native()))
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	_cret = C.graphene_ray_init(_arg0, _arg1, _arg2)

	var _ray *Ray // out

	_ray = (*Ray)(unsafe.Pointer(_cret))

	return _ray
}

// InitFromRay checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) InitFromRay(src *Ray) *Ray {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_ray_t // out
	var _cret *C.graphene_ray_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_ray_t)(unsafe.Pointer(src.Native()))

	_cret = C.graphene_ray_init_from_ray(_arg0, _arg1)

	var _ray *Ray // out

	_ray = (*Ray)(unsafe.Pointer(_cret))

	return _ray
}

// InitFromVec3 checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) InitFromVec3(origin *Vec3, direction *Vec3) *Ray {
	var _arg0 *C.graphene_ray_t  // out
	var _arg1 *C.graphene_vec3_t // out
	var _arg2 *C.graphene_vec3_t // out
	var _cret *C.graphene_ray_t  // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_vec3_t)(unsafe.Pointer(origin.Native()))
	_arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(direction.Native()))

	_cret = C.graphene_ray_init_from_vec3(_arg0, _arg1, _arg2)

	var _ray *Ray // out

	_ray = (*Ray)(unsafe.Pointer(_cret))

	return _ray
}

// IntersectBox checks whether the given #graphene_ray_t @r intersects the given
// #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectBox(b *Box) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_box_t                  // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_ray_intersect_box(_arg0, _arg1, &_arg2)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectSphere checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectSphere(s *Sphere) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_sphere_t               // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	_cret = C.graphene_ray_intersect_sphere(_arg0, _arg1, &_arg2)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectTriangle checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectTriangle(t *Triangle) (float32, RayIntersectionKind) {
	var _arg0 *C.graphene_ray_t                  // out
	var _arg1 *C.graphene_triangle_t             // out
	var _arg2 C.float                            // in
	var _cret C.graphene_ray_intersection_kind_t // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	_cret = C.graphene_ray_intersect_triangle(_arg0, _arg1, &_arg2)

	var _tOut float32                            // out
	var _rayIntersectionKind RayIntersectionKind // out

	_tOut = float32(_arg2)
	_rayIntersectionKind = RayIntersectionKind(_cret)

	return _tOut, _rayIntersectionKind
}

// IntersectsBox checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectsBox(b *Box) bool {
	var _arg0 *C.graphene_ray_t // out
	var _arg1 *C.graphene_box_t // out
	var _cret C._Bool           // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_box_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_ray_intersects_box(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsSphere checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectsSphere(s *Sphere) bool {
	var _arg0 *C.graphene_ray_t    // out
	var _arg1 *C.graphene_sphere_t // out
	var _cret C._Bool              // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(s.Native()))

	_cret = C.graphene_ray_intersects_sphere(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IntersectsTriangle checks whether the given #graphene_ray_t @r intersects the
// given #graphene_triangle_t @b.
//
// See also: graphene_ray_intersect_triangle()
func (r *Ray) IntersectsTriangle(t *Triangle) bool {
	var _arg0 *C.graphene_ray_t      // out
	var _arg1 *C.graphene_triangle_t // out
	var _cret C._Bool                // in

	_arg0 = (*C.graphene_ray_t)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.graphene_triangle_t)(unsafe.Pointer(t.Native()))

	_cret = C.graphene_ray_intersects_triangle(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
