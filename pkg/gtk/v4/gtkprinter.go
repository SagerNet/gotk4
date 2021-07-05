// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_capabilities_get_type()), F: marshalPrintCapabilities},
		{T: externglib.Type(C.gtk_printer_get_type()), F: marshalPrinter},
	})
}

// PrintCapabilities specifies which features the print dialog should offer.
//
// If neither GTK_PRINT_CAPABILITY_GENERATE_PDF nor
// GTK_PRINT_CAPABILITY_GENERATE_PS is specified, GTK assumes that all formats
// are supported.
type PrintCapabilities int

const (
	// PrintCapabilitiesPageSet: print dialog will offer printing even/odd
	// pages.
	PrintCapabilitiesPageSet PrintCapabilities = 0b1
	// PrintCapabilitiesCopies: print dialog will allow to print multiple
	// copies.
	PrintCapabilitiesCopies PrintCapabilities = 0b10
	// PrintCapabilitiesCollate: print dialog will allow to collate multiple
	// copies.
	PrintCapabilitiesCollate PrintCapabilities = 0b100
	// PrintCapabilitiesReverse: print dialog will allow to print pages in
	// reverse order.
	PrintCapabilitiesReverse PrintCapabilities = 0b1000
	// PrintCapabilitiesScale: print dialog will allow to scale the output.
	PrintCapabilitiesScale PrintCapabilities = 0b10000
	// PrintCapabilitiesGeneratePDF: the program will send the document to the
	// printer in PDF format
	PrintCapabilitiesGeneratePDF PrintCapabilities = 0b100000
	// PrintCapabilitiesGeneratePS: the program will send the document to the
	// printer in Postscript format
	PrintCapabilitiesGeneratePS PrintCapabilities = 0b1000000
	// PrintCapabilitiesPreview: print dialog will offer a preview
	PrintCapabilitiesPreview PrintCapabilities = 0b10000000
	// PrintCapabilitiesNumberUp: print dialog will offer printing multiple
	// pages per sheet
	PrintCapabilitiesNumberUp PrintCapabilities = 0b100000000
	// PrintCapabilitiesNumberUpLayout: print dialog will allow to rearrange
	// pages when printing multiple pages per sheet
	PrintCapabilitiesNumberUpLayout PrintCapabilities = 0b1000000000
)

func marshalPrintCapabilities(p uintptr) (interface{}, error) {
	return PrintCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrinterFunc: the type of function passed to gtk_enumerate_printers().
//
// Note that you need to ref @printer, if you want to keep a reference to it
// after the function has returned.
type PrinterFunc func(printer Printer) (ok bool)

//export gotk4_PrinterFunc
func gotk4_PrinterFunc(arg0 *C.GtkPrinter, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var printer Printer // out

	printer = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Printer)

	fn := v.(PrinterFunc)
	ok := fn(printer)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// Printer: `GtkPrinter` object represents a printer.
//
// You only need to deal directly with printers if you use the non-portable
// [class@Gtk.PrintUnixDialog] API.
//
// A `GtkPrinter` allows to get status information about the printer, such as
// its description, its location, the number of queued jobs, etc. Most
// importantly, a `GtkPrinter` object can be used to create a
// [class@Gtk.PrintJob] object, which lets you print to the printer.
type Printer interface {
	gextras.Objector

	// AcceptsPDF returns whether the printer accepts input in PDF format.
	AcceptsPDF() bool
	// AcceptsPS returns whether the printer accepts input in PostScript format.
	AcceptsPS() bool
	// Compare compares two printers.
	Compare(b Printer) int
	// Capabilities returns the printer’s capabilities.
	//
	// This is useful when you’re using `GtkPrintUnixDialog`’s
	// manual-capabilities setting and need to know which settings the printer
	// can handle and which you must handle yourself.
	//
	// This will return 0 unless the printer’s details are available, see
	// [method@Gtk.Printer.has_details] and
	// [method@Gtk.Printer.request_details].
	Capabilities() PrintCapabilities
	// DefaultPageSize returns default page size of @printer.
	DefaultPageSize() PageSetup
	// Description gets the description of the printer.
	Description() string
	// HardMargins: retrieve the hard margins of @printer.
	//
	// These are the margins that define the area at the borders of the paper
	// that the printer cannot print to.
	//
	// Note: This will not succeed unless the printer’s details are available,
	// see [method@Gtk.Printer.has_details] and
	// [method@Gtk.Printer.request_details].
	HardMargins() (top float64, bottom float64, left float64, right float64, ok bool)
	// HardMarginsForPaperSize: retrieve the hard margins of @printer for
	// @paper_size.
	//
	// These are the margins that define the area at the borders of the paper
	// that the printer cannot print to.
	//
	// Note: This will not succeed unless the printer’s details are available,
	// see [method@Gtk.Printer.has_details] and
	// [method@Gtk.Printer.request_details].
	HardMarginsForPaperSize(paperSize *PaperSize) (top float64, bottom float64, left float64, right float64, ok bool)
	// IconName gets the name of the icon to use for the printer.
	IconName() string
	// JobCount gets the number of jobs currently queued on the printer.
	JobCount() int
	// Location returns a description of the location of the printer.
	Location() string
	// Name returns the name of the printer.
	Name() string
	// StateMessage returns the state message describing the current state of
	// the printer.
	StateMessage() string
	// HasDetails returns whether the printer details are available.
	HasDetails() bool
	// IsAcceptingJobs returns whether the printer is accepting jobs
	IsAcceptingJobs() bool
	// IsActive returns whether the printer is currently active (i.e. accepts
	// new jobs).
	IsActive() bool
	// IsDefault returns whether the printer is the default printer.
	IsDefault() bool
	// IsPaused returns whether the printer is currently paused.
	//
	// A paused printer still accepts jobs, but it is not printing them.
	IsPaused() bool
	// IsVirtual returns whether the printer is virtual (i.e. does not represent
	// actual printer hardware, but something like a CUPS class).
	IsVirtual() bool
	// RequestDetails requests the printer details.
	//
	// When the details are available, the
	// [signal@Gtk.Printer::details-acquired] signal will be emitted on
	// @printer.
	RequestDetails()
}

// printer implements the Printer class.
type printer struct {
	gextras.Objector
}

// WrapPrinter wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrinter(obj *externglib.Object) Printer {
	return printer{
		Objector: obj,
	}
}

func marshalPrinter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrinter(obj), nil
}

func (p printer) AcceptsPDF() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_accepts_pdf(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) AcceptsPS() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_accepts_ps(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a printer) Compare(b Printer) int {
	var _arg0 *C.GtkPrinter // out
	var _arg1 *C.GtkPrinter // out
	var _cret C.int         // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkPrinter)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_printer_compare(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p printer) Capabilities() PrintCapabilities {
	var _arg0 *C.GtkPrinter          // out
	var _cret C.GtkPrintCapabilities // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_capabilities(_arg0)

	var _printCapabilities PrintCapabilities // out

	_printCapabilities = PrintCapabilities(_cret)

	return _printCapabilities
}

func (p printer) DefaultPageSize() PageSetup {
	var _arg0 *C.GtkPrinter   // out
	var _cret *C.GtkPageSetup // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_default_page_size(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

func (p printer) Description() string {
	var _arg0 *C.GtkPrinter // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p printer) HardMargins() (top float64, bottom float64, left float64, right float64, ok bool) {
	var _arg0 *C.GtkPrinter // out
	var _arg1 C.double      // in
	var _arg2 C.double      // in
	var _arg3 C.double      // in
	var _arg4 C.double      // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_hard_margins(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _top float64    // out
	var _bottom float64 // out
	var _left float64   // out
	var _right float64  // out
	var _ok bool        // out

	_top = float64(_arg1)
	_bottom = float64(_arg2)
	_left = float64(_arg3)
	_right = float64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _top, _bottom, _left, _right, _ok
}

func (p printer) HardMarginsForPaperSize(paperSize *PaperSize) (top float64, bottom float64, left float64, right float64, ok bool) {
	var _arg0 *C.GtkPrinter   // out
	var _arg1 *C.GtkPaperSize // out
	var _arg2 C.double        // in
	var _arg3 C.double        // in
	var _arg4 C.double        // in
	var _arg5 C.double        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(paperSize))

	_cret = C.gtk_printer_get_hard_margins_for_paper_size(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _top float64    // out
	var _bottom float64 // out
	var _left float64   // out
	var _right float64  // out
	var _ok bool        // out

	_top = float64(_arg2)
	_bottom = float64(_arg3)
	_left = float64(_arg4)
	_right = float64(_arg5)
	if _cret != 0 {
		_ok = true
	}

	return _top, _bottom, _left, _right, _ok
}

func (p printer) IconName() string {
	var _arg0 *C.GtkPrinter // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p printer) JobCount() int {
	var _arg0 *C.GtkPrinter // out
	var _cret C.int         // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_job_count(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p printer) Location() string {
	var _arg0 *C.GtkPrinter // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_location(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p printer) Name() string {
	var _arg0 *C.GtkPrinter // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p printer) StateMessage() string {
	var _arg0 *C.GtkPrinter // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_get_state_message(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p printer) HasDetails() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_has_details(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) IsAcceptingJobs() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_is_accepting_jobs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) IsActive() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_is_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) IsDefault() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_is_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) IsPaused() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_is_paused(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) IsVirtual() bool {
	var _arg0 *C.GtkPrinter // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_printer_is_virtual(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printer) RequestDetails() {
	var _arg0 *C.GtkPrinter // out

	_arg0 = (*C.GtkPrinter)(unsafe.Pointer(p.Native()))

	C.gtk_printer_request_details(_arg0)
}
