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
		{T: externglib.Type(C.graphene_size_get_type()), F: marshalSize},
	})
}

// Size: a size.
type Size C.graphene_size_t

// WrapSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSize(ptr unsafe.Pointer) *Size {
	return (*Size)(ptr)
}

func marshalSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Size)(unsafe.Pointer(b)), nil
}

// NewSizeAlloc constructs a struct Size.
func NewSizeAlloc() *Size {
	var _cret *C.graphene_size_t // in

	_cret = C.graphene_size_alloc()

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_size, func(v **Size) {
		C.free(unsafe.Pointer(v))
	})

	return _size
}

// Native returns the underlying C source pointer.
func (s *Size) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}

// Equal scales the components of a #graphene_size_t using the given @factor.
func (s *Size) Equal(b *Size) bool {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret C._Bool            // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(b.Native()))

	_cret = C.graphene_size_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Free scales the components of a #graphene_size_t using the given @factor.
func (s *Size) Free() {
	var _arg0 *C.graphene_size_t // out

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))

	C.graphene_size_free(_arg0)
}

// Init scales the components of a #graphene_size_t using the given @factor.
func (s *Size) Init(width float32, height float32) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _arg2 C.float            // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(width)
	_arg2 = C.float(height)

	_cret = C.graphene_size_init(_arg0, _arg1, _arg2)

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))

	return _size
}

// InitFromSize scales the components of a #graphene_size_t using the given
// @factor.
func (s *Size) InitFromSize(src *Size) *Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _cret *C.graphene_size_t // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(src.Native()))

	_cret = C.graphene_size_init_from_size(_arg0, _arg1)

	var _size *Size // out

	_size = (*Size)(unsafe.Pointer(_cret))

	return _size
}

// Interpolate scales the components of a #graphene_size_t using the given
// @factor.
func (s *Size) Interpolate(b *Size, factor float64) Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 *C.graphene_size_t // out
	var _arg2 C.double           // out
	var _arg3 C.graphene_size_t  // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.graphene_size_t)(unsafe.Pointer(b.Native()))
	_arg2 = C.double(factor)

	C.graphene_size_interpolate(_arg0, _arg1, _arg2, &_arg3)

	var _res Size // out

	{
		var refTmpIn *C.graphene_size_t
		var refTmpOut *Size

		in0 := &_arg3
		refTmpIn = in0

		refTmpOut = (*Size)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}

// Scale scales the components of a #graphene_size_t using the given @factor.
func (s *Size) Scale(factor float32) Size {
	var _arg0 *C.graphene_size_t // out
	var _arg1 C.float            // out
	var _arg2 C.graphene_size_t  // in

	_arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(factor)

	C.graphene_size_scale(_arg0, _arg1, &_arg2)

	var _res Size // out

	{
		var refTmpIn *C.graphene_size_t
		var refTmpOut *Size

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*Size)(unsafe.Pointer(refTmpIn))

		_res = *refTmpOut
	}

	return _res
}
