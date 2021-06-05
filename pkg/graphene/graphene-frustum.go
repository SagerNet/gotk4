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
		{T: externglib.Type(C.graphene_frustum_get_type()), F: marshalFrustum},
	})
}

// Frustum: a 3D volume delimited by 2D clip planes.
//
// The contents of the `graphene_frustum_t` are private, and should not be
// modified directly.
type Frustum struct {
	native C.graphene_frustum_t
}

// WrapFrustum wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFrustum(ptr unsafe.Pointer) *Frustum {
	if ptr == nil {
		return nil
	}

	return (*Frustum)(ptr)
}

func marshalFrustum(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFrustum(unsafe.Pointer(b)), nil
}

// NewFrustumAlloc constructs a struct Frustum.
func NewFrustumAlloc() *Frustum {
	var cret *C.graphene_frustum_t
	var ret1 *Frustum

	cret = C.graphene_frustum_alloc()

	ret1 = WrapFrustum(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Frustum) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (f *Frustum) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// ContainsPoint checks whether a point is inside the volume defined by the
// given #graphene_frustum_t.
func (f *Frustum) ContainsPoint(point *Point3D) bool {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_point3d_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_point3d_t)(unsafe.Pointer(point.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_frustum_contains_point(arg0, point)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Equal checks whether the two given #graphene_frustum_t are equal.
func (a *Frustum) Equal(b *Frustum) bool {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_frustum_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_frustum_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_frustum_equal(arg0, b)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Free frees the resources allocated by graphene_frustum_alloc().
func (f *Frustum) Free() {
	var arg0 *C.graphene_frustum_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))

	C.graphene_frustum_free(arg0)
}

// Planes retrieves the planes that define the given #graphene_frustum_t.
func (f *Frustum) Planes() [6]Plane {
	var arg0 *C.graphene_frustum_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))

	var arg1 [6]C.graphene_plane_t
	var ret1 [6]Plane

	C.graphene_frustum_get_planes(arg0, &arg1)

	{
		tmp := *(*[6]Plane)(unsafe.Pointer(&arg1))
		for i := 0; i < 6; i++ {
			src := tmp[i]
			ret1[i] = WrapPlane(unsafe.Pointer(src))
		}
	}

	return ret1
}

// Init initializes the given #graphene_frustum_t using the provided clipping
// planes.
func (f *Frustum) Init(p0 *Plane, p1 *Plane, p2 *Plane, p3 *Plane, p4 *Plane, p5 *Plane) *Frustum {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_plane_t
	var arg2 *C.graphene_plane_t
	var arg3 *C.graphene_plane_t
	var arg4 *C.graphene_plane_t
	var arg5 *C.graphene_plane_t
	var arg6 *C.graphene_plane_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_plane_t)(unsafe.Pointer(p0.Native()))
	arg2 = (*C.graphene_plane_t)(unsafe.Pointer(p1.Native()))
	arg3 = (*C.graphene_plane_t)(unsafe.Pointer(p2.Native()))
	arg4 = (*C.graphene_plane_t)(unsafe.Pointer(p3.Native()))
	arg5 = (*C.graphene_plane_t)(unsafe.Pointer(p4.Native()))
	arg6 = (*C.graphene_plane_t)(unsafe.Pointer(p5.Native()))

	var cret *C.graphene_frustum_t
	var ret1 *Frustum

	cret = C.graphene_frustum_init(arg0, p0, p1, p2, p3, p4, p5)

	ret1 = WrapFrustum(unsafe.Pointer(cret))

	return ret1
}

// InitFromFrustum initializes the given #graphene_frustum_t using the clipping
// planes of another #graphene_frustum_t.
func (f *Frustum) InitFromFrustum(src *Frustum) *Frustum {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_frustum_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_frustum_t)(unsafe.Pointer(src.Native()))

	var cret *C.graphene_frustum_t
	var ret1 *Frustum

	cret = C.graphene_frustum_init_from_frustum(arg0, src)

	ret1 = WrapFrustum(unsafe.Pointer(cret))

	return ret1
}

// InitFromMatrix initializes a #graphene_frustum_t using the given @matrix.
func (f *Frustum) InitFromMatrix(matrix *Matrix) *Frustum {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_matrix_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_matrix_t)(unsafe.Pointer(matrix.Native()))

	var cret *C.graphene_frustum_t
	var ret1 *Frustum

	cret = C.graphene_frustum_init_from_matrix(arg0, matrix)

	ret1 = WrapFrustum(unsafe.Pointer(cret))

	return ret1
}

// IntersectsBox checks whether the given @box intersects a plane of a
// #graphene_frustum_t.
func (f *Frustum) IntersectsBox(box *Box) bool {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_box_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_box_t)(unsafe.Pointer(box.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_frustum_intersects_box(arg0, box)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IntersectsSphere checks whether the given @sphere intersects a plane of a
// #graphene_frustum_t.
func (f *Frustum) IntersectsSphere(sphere *Sphere) bool {
	var arg0 *C.graphene_frustum_t
	var arg1 *C.graphene_sphere_t

	arg0 = (*C.graphene_frustum_t)(unsafe.Pointer(f.Native()))
	arg1 = (*C.graphene_sphere_t)(unsafe.Pointer(sphere.Native()))

	var cret C._Bool
	var ret1 bool

	cret = C.graphene_frustum_intersects_sphere(arg0, sphere)

	ret1 = C.bool(cret) != C.false

	return ret1
}