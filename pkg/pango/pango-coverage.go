// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_coverage_level_get_type()), F: marshalCoverageLevel},
		{T: externglib.Type(C.pango_coverage_get_type()), F: marshalCoverager},
	})
}

// CoverageLevel: PangoCoverageLevel is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel int

const (
	// CoverageNone: character is not representable with the font.
	CoverageNone CoverageLevel = iota
	// CoverageFallback: character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageFallback
	// CoverageApproximate: character is represented as basically the correct
	// graphical form, but with a stylistic variant inappropriate for the
	// current script.
	CoverageApproximate
	// CoverageExact: character is represented as the correct graphical form.
	CoverageExact
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CoverageLevel.
func (c CoverageLevel) String() string {
	switch c {
	case CoverageNone:
		return "None"
	case CoverageFallback:
		return "Fallback"
	case CoverageApproximate:
		return "Approximate"
	case CoverageExact:
		return "Exact"
	default:
		return fmt.Sprintf("CoverageLevel(%d)", c)
	}
}

// Coverage structure is a map from Unicode characters to CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage struct {
	*externglib.Object
}

var _ gextras.Nativer = (*Coverage)(nil)

func wrapCoverage(obj *externglib.Object) *Coverage {
	return &Coverage{
		Object: obj,
	}
}

func marshalCoverager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCoverage(obj), nil
}

// NewCoverage: create a new PangoCoverage
func NewCoverage() *Coverage {
	var _cret *C.PangoCoverage // in

	_cret = C.pango_coverage_new()

	var _coverage *Coverage // out

	_coverage = wrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _coverage
}

// Copy an existing PangoCoverage.
func (coverage *Coverage) Copy() *Coverage {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))

	_cret = C.pango_coverage_copy(_arg0)

	var _ret *Coverage // out

	_ret = wrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// Get: determine whether a particular index is covered by coverage.
func (coverage *Coverage) Get(index_ int) CoverageLevel {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _cret C.PangoCoverageLevel // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))
	_arg1 = C.int(index_)

	_cret = C.pango_coverage_get(_arg0, _arg1)

	var _coverageLevel CoverageLevel // out

	_coverageLevel = CoverageLevel(_cret)

	return _coverageLevel
}

// Max: set the coverage for each index in coverage to be the max (better) value
// of the current coverage for the index and the coverage for the corresponding
// index in other.
//
// Deprecated: This function does nothing.
func (coverage *Coverage) Max(other *Coverage) {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))
	_arg1 = (*C.PangoCoverage)(unsafe.Pointer(other.Native()))

	C.pango_coverage_max(_arg0, _arg1)
}

// Ref: increase the reference count on the PangoCoverage by one.
func (coverage *Coverage) ref() *Coverage {
	var _arg0 *C.PangoCoverage // out
	var _cret *C.PangoCoverage // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))

	_cret = C.pango_coverage_ref(_arg0)

	var _ret *Coverage // out

	_ret = wrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _ret
}

// Set: modify a particular index within coverage
func (coverage *Coverage) Set(index_ int, level CoverageLevel) {
	var _arg0 *C.PangoCoverage     // out
	var _arg1 C.int                // out
	var _arg2 C.PangoCoverageLevel // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))
	_arg1 = C.int(index_)
	_arg2 = C.PangoCoverageLevel(level)

	C.pango_coverage_set(_arg0, _arg1, _arg2)
}

// ToBytes: convert a PangoCoverage structure into a flat binary format.
//
// Deprecated: This returns NULL.
func (coverage *Coverage) ToBytes() []byte {
	var _arg0 *C.PangoCoverage // out
	var _arg1 *C.guchar        // in
	var _arg2 C.int            // in

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))

	C.pango_coverage_to_bytes(_arg0, &_arg1, &_arg2)

	var _bytes []byte // out

	defer C.free(unsafe.Pointer(_arg1))
	_bytes = make([]byte, _arg2)
	copy(_bytes, unsafe.Slice((*byte)(unsafe.Pointer(_arg1)), _arg2))

	return _bytes
}

// Unref: decrease the reference count on the PangoCoverage by one.
//
// If the result is zero, free the coverage and all associated memory.
func (coverage *Coverage) unref() {
	var _arg0 *C.PangoCoverage // out

	_arg0 = (*C.PangoCoverage)(unsafe.Pointer(coverage.Native()))

	C.pango_coverage_unref(_arg0)
}

// CoverageFromBytes: convert data generated from pango_coverage_to_bytes() back
// to a PangoCoverage.
//
// Deprecated: This returns NULL.
func CoverageFromBytes(bytes []byte) *Coverage {
	var _arg1 *C.guchar // out
	var _arg2 C.int
	var _cret *C.PangoCoverage // in

	_arg2 = (C.int)(len(bytes))
	if len(bytes) > 0 {
		_arg1 = (*C.guchar)(unsafe.Pointer(&bytes[0]))
	}

	_cret = C.pango_coverage_from_bytes(_arg1, _arg2)

	var _coverage *Coverage // out

	_coverage = wrapCoverage(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _coverage
}
