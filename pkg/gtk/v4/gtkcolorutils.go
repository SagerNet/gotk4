// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gtk/gtk.h>
import "C"

// HSVToRGB converts a color from HSV space to RGB.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
func HSVToRGB(h float32, s float32, v float32) (r float32, g float32, b float32) {
	var _arg1 C.float // out
	var _arg2 C.float // out
	var _arg3 C.float // out
	var _arg4 C.float // in
	var _arg5 C.float // in
	var _arg6 C.float // in

	_arg1 = C.float(h)
	_arg2 = C.float(s)
	_arg3 = C.float(v)

	C.gtk_hsv_to_rgb(_arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6)

	var _r float32 // out
	var _g float32 // out
	var _b float32 // out

	_r = float32(_arg4)
	_g = float32(_arg5)
	_b = float32(_arg6)

	return _r, _g, _b
}

// RGBToHSV converts a color from RGB space to HSV.
//
// Input values must be in the [0.0, 1.0] range; output values will be in the
// same range.
func RGBToHSV(r float32, g float32, b float32) (h float32, s float32, v float32) {
	var _arg1 C.float // out
	var _arg2 C.float // out
	var _arg3 C.float // out
	var _arg4 C.float // in
	var _arg5 C.float // in
	var _arg6 C.float // in

	_arg1 = C.float(r)
	_arg2 = C.float(g)
	_arg3 = C.float(b)

	C.gtk_rgb_to_hsv(_arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6)

	var _h float32 // out
	var _s float32 // out
	var _v float32 // out

	_h = float32(_arg4)
	_s = float32(_arg5)
	_v = float32(_arg6)

	return _h, _s, _v
}
