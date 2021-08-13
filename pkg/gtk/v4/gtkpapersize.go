// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_paper_size_get_type()), F: marshalPaperSize},
	})
}

// PAPER_NAME_A3: name for the A3 paper size.
const PAPER_NAME_A3 = "iso_a3"

// PAPER_NAME_A4: name for the A4 paper size.
const PAPER_NAME_A4 = "iso_a4"

// PAPER_NAME_A5: name for the A5 paper size.
const PAPER_NAME_A5 = "iso_a5"

// PAPER_NAME_B5: name for the B5 paper size.
const PAPER_NAME_B5 = "iso_b5"

// PAPER_NAME_EXECUTIVE: name for the Executive paper size.
const PAPER_NAME_EXECUTIVE = "na_executive"

// PAPER_NAME_LEGAL: name for the Legal paper size.
const PAPER_NAME_LEGAL = "na_legal"

// PAPER_NAME_LETTER: name for the Letter paper size.
const PAPER_NAME_LETTER = "na_letter"

// PaperSize: GtkPaperSize handles paper sizes.
//
// It uses the standard called PWG 5101.1-2002 PWG: Standard for Media
// Standardized Names (http://www.pwg.org/standards.html) to name the paper
// sizes (and to get the data for the page sizes). In addition to standard paper
// sizes, GtkPaperSize allows to construct custom paper sizes with arbitrary
// dimensions.
//
// The GtkPaperSize object stores not only the dimensions (width and height) of
// a paper size and its name, it also provides default print margins.
type PaperSize struct {
	nocopy gextras.NoCopy
	native *C.GtkPaperSize
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &PaperSize{native: (*C.GtkPaperSize)(unsafe.Pointer(b))}, nil
}

// NewPaperSize constructs a struct PaperSize.
func NewPaperSize(name string) *PaperSize {
	var _arg1 *C.char         // out
	var _cret *C.GtkPaperSize // in

	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_paper_size_new(_arg1)
	runtime.KeepAlive(name)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// NewPaperSizeCustom constructs a struct PaperSize.
func NewPaperSizeCustom(name string, displayName string, width float64, height float64, unit Unit) *PaperSize {
	var _arg1 *C.char         // out
	var _arg2 *C.char         // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out
	var _arg5 C.GtkUnit       // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(displayName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.double(width)
	_arg4 = C.double(height)
	_arg5 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_new_custom(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(name)
	runtime.KeepAlive(displayName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(unit)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// NewPaperSizeFromGVariant constructs a struct PaperSize.
func NewPaperSizeFromGVariant(variant *glib.Variant) *PaperSize {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	_cret = C.gtk_paper_size_new_from_gvariant(_arg1)
	runtime.KeepAlive(variant)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// NewPaperSizeFromIPP constructs a struct PaperSize.
func NewPaperSizeFromIPP(ippName string, width float64, height float64) *PaperSize {
	var _arg1 *C.char         // out
	var _arg2 C.double        // out
	var _arg3 C.double        // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(ippName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(width)
	_arg3 = C.double(height)

	_cret = C.gtk_paper_size_new_from_ipp(_arg1, _arg2, _arg3)
	runtime.KeepAlive(ippName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// NewPaperSizeFromKeyFile constructs a struct PaperSize.
func NewPaperSizeFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PaperSize, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.char         // out
	var _cret *C.GtkPaperSize // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_paper_size_new_from_key_file(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _paperSize *PaperSize // out
	var _goerr error          // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _paperSize, _goerr
}

// NewPaperSizeFromPPD constructs a struct PaperSize.
func NewPaperSizeFromPPD(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	var _arg1 *C.char         // out
	var _arg2 *C.char         // out
	var _arg3 C.double        // out
	var _arg4 C.double        // out
	var _cret *C.GtkPaperSize // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(ppdName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(ppdDisplayName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.double(width)
	_arg4 = C.double(height)

	_cret = C.gtk_paper_size_new_from_ppd(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(ppdName)
	runtime.KeepAlive(ppdDisplayName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// Copy copies an existing GtkPaperSize.
func (other *PaperSize) Copy() *PaperSize {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gtk_paper_size_copy(_arg0)
	runtime.KeepAlive(other)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// DefaultBottomMargin gets the default bottom margin for the GtkPaperSize.
func (size *PaperSize) DefaultBottomMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_bottom_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultLeftMargin gets the default left margin for the GtkPaperSize.
func (size *PaperSize) DefaultLeftMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_left_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultRightMargin gets the default right margin for the GtkPaperSize.
func (size *PaperSize) DefaultRightMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_right_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DefaultTopMargin gets the default top margin for the GtkPaperSize.
func (size *PaperSize) DefaultTopMargin(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_default_top_margin(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DisplayName gets the human-readable name of the GtkPaperSize.
func (size *PaperSize) DisplayName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_display_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Height gets the paper height of the GtkPaperSize, in units of unit.
func (size *PaperSize) Height(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_height(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Name gets the name of the GtkPaperSize.
func (size *PaperSize) Name() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PPDName gets the PPD name of the GtkPaperSize, which may be NULL.
func (size *PaperSize) PPDName() string {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_get_ppd_name(_arg0)
	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Width gets the paper width of the GtkPaperSize, in units of unit.
func (size *PaperSize) Width(unit Unit) float64 {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.GtkUnit       // out
	var _cret C.double        // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_paper_size_get_width(_arg0, _arg1)
	runtime.KeepAlive(size)
	runtime.KeepAlive(unit)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// IsCustom returns TRUE if size is not a standard paper size.
func (size *PaperSize) IsCustom() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_is_custom(_arg0)
	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqual compares two GtkPaperSize objects.
func (size1 *PaperSize) IsEqual(size2 *PaperSize) bool {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size1)))
	_arg1 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size2)))

	_cret = C.gtk_paper_size_is_equal(_arg0, _arg1)
	runtime.KeepAlive(size1)
	runtime.KeepAlive(size2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsIPP returns TRUE if size is an IPP standard paper size.
func (size *PaperSize) IsIPP() bool {
	var _arg0 *C.GtkPaperSize // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))

	_cret = C.gtk_paper_size_is_ipp(_arg0)
	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSize changes the dimensions of a size to width x height.
func (size *PaperSize) SetSize(width float64, height float64, unit Unit) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 C.double        // out
	var _arg2 C.double        // out
	var _arg3 C.GtkUnit       // out

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = C.double(width)
	_arg2 = C.double(height)
	_arg3 = C.GtkUnit(unit)

	C.gtk_paper_size_set_size(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(size)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(unit)
}

// ToGVariant: serialize a paper size to an a{sv} variant.
func (paperSize *PaperSize) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPaperSize // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(paperSize)))

	_cret = C.gtk_paper_size_to_gvariant(_arg0)
	runtime.KeepAlive(paperSize)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variant
}

// ToKeyFile: this function adds the paper size from size to key_file.
func (size *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPaperSize // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.char         // out

	_arg0 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_paper_size_to_key_file(_arg0, _arg1, _arg2)
	runtime.KeepAlive(size)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)
}

// PaperSizeGetDefault returns the name of the default paper size, which depends
// on the current locale.
func PaperSizeGetDefault() string {
	var _cret *C.char // in

	_cret = C.gtk_paper_size_get_default()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PaperSizeGetPaperSizes creates a list of known paper sizes.
func PaperSizeGetPaperSizes(includeCustom bool) []*PaperSize {
	var _arg1 C.gboolean // out
	var _cret *C.GList   // in

	if includeCustom {
		_arg1 = C.TRUE
	}

	_cret = C.gtk_paper_size_get_paper_sizes(_arg1)
	runtime.KeepAlive(includeCustom)

	var _list []*PaperSize // out

	_list = make([]*PaperSize, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkPaperSize)(v)
		var dst *PaperSize // out
		dst = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(dst, func(v *PaperSize) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
		})
		_list = append(_list, dst)
	})

	return _list
}
