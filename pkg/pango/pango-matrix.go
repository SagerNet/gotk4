// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_matrix_get_type()), F: marshalMatrix},
	})
}

// Matrix: PangoMatrix specifies a transformation between user-space and device
// coordinates.
//
// The transformation is given by
//
//    x_device = x_user * matrix->xx + y_user * matrix->xy + matrix->x0;
//    y_device = x_user * matrix->yx + y_user * matrix->yy + matrix->y0;
//
//
// An instance of this type is always passed by reference.
type Matrix struct {
	*matrix
}

// matrix is the struct that's finalized.
type matrix struct {
	native *C.PangoMatrix
}

func marshalMatrix(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Matrix{&matrix{(*C.PangoMatrix)(unsafe.Pointer(b))}}, nil
}

// XX: 1st component of the transformation matrix
func (m *Matrix) XX() float64 {
	var v float64 // out
	v = float64(m.native.xx)
	return v
}

// XY: 2nd component of the transformation matrix
func (m *Matrix) XY() float64 {
	var v float64 // out
	v = float64(m.native.xy)
	return v
}

// YX: 3rd component of the transformation matrix
func (m *Matrix) YX() float64 {
	var v float64 // out
	v = float64(m.native.yx)
	return v
}

// YY: 4th component of the transformation matrix
func (m *Matrix) YY() float64 {
	var v float64 // out
	v = float64(m.native.yy)
	return v
}

// X0: x translation
func (m *Matrix) X0() float64 {
	var v float64 // out
	v = float64(m.native.x0)
	return v
}

// Y0: y translation
func (m *Matrix) Y0() float64 {
	var v float64 // out
	v = float64(m.native.y0)
	return v
}

// Concat changes the transformation represented by matrix to be the
// transformation given by first applying transformation given by new_matrix
// then applying the original transformation.
func (matrix *Matrix) Concat(newMatrix *Matrix) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 *C.PangoMatrix // out

	_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	_arg1 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(newMatrix)))

	C.pango_matrix_concat(_arg0, _arg1)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(newMatrix)
}

// Copy copies a PangoMatrix.
func (matrix *Matrix) Copy() *Matrix {
	var _arg0 *C.PangoMatrix // out
	var _cret *C.PangoMatrix // in

	if matrix != nil {
		_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	_cret = C.pango_matrix_copy(_arg0)
	runtime.KeepAlive(matrix)

	var _ret *Matrix // out

	if _cret != nil {
		_ret = (*Matrix)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_ret)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_matrix_free((*C.PangoMatrix)(intern.C))
			},
		)
	}

	return _ret
}

// FontScaleFactor returns the scale factor of a matrix on the height of the
// font.
//
// That is, the scale factor in the direction perpendicular to the vector that
// the X coordinate is mapped to. If the scale in the X coordinate is needed as
// well, use pango.Matrix.GetFontScaleFactors().
func (matrix *Matrix) FontScaleFactor() float64 {
	var _arg0 *C.PangoMatrix // out
	var _cret C.double       // in

	if matrix != nil {
		_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	_cret = C.pango_matrix_get_font_scale_factor(_arg0)
	runtime.KeepAlive(matrix)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// FontScaleFactors calculates the scale factor of a matrix on the width and
// height of the font.
//
// That is, xscale is the scale factor in the direction of the X coordinate, and
// yscale is the scale factor in the direction perpendicular to the vector that
// the X coordinate is mapped to.
//
// Note that output numbers will always be non-negative.
func (matrix *Matrix) FontScaleFactors() (xscale float64, yscale float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // in
	var _arg2 C.double       // in

	if matrix != nil {
		_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	}

	C.pango_matrix_get_font_scale_factors(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(matrix)

	var _xscale float64 // out
	var _yscale float64 // out

	_xscale = float64(_arg1)
	_yscale = float64(_arg2)

	return _xscale, _yscale
}

// Rotate changes the transformation represented by matrix to be the
// transformation given by first rotating by degrees degrees counter-clockwise
// then applying the original transformation.
func (matrix *Matrix) Rotate(degrees float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out

	_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	_arg1 = C.double(degrees)

	C.pango_matrix_rotate(_arg0, _arg1)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(degrees)
}

// Scale changes the transformation represented by matrix to be the
// transformation given by first scaling by sx in the X direction and sy in the
// Y direction then applying the original transformation.
func (matrix *Matrix) Scale(scaleX float64, scaleY float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out
	var _arg2 C.double       // out

	_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	_arg1 = C.double(scaleX)
	_arg2 = C.double(scaleY)

	C.pango_matrix_scale(_arg0, _arg1, _arg2)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(scaleX)
	runtime.KeepAlive(scaleY)
}

// Translate changes the transformation represented by matrix to be the
// transformation given by first translating by (tx, ty) then applying the
// original transformation.
func (matrix *Matrix) Translate(tx float64, ty float64) {
	var _arg0 *C.PangoMatrix // out
	var _arg1 C.double       // out
	var _arg2 C.double       // out

	_arg0 = (*C.PangoMatrix)(gextras.StructNative(unsafe.Pointer(matrix)))
	_arg1 = C.double(tx)
	_arg2 = C.double(ty)

	C.pango_matrix_translate(_arg0, _arg1, _arg2)
	runtime.KeepAlive(matrix)
	runtime.KeepAlive(tx)
	runtime.KeepAlive(ty)
}
