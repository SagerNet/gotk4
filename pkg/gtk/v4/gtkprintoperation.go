// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_print_error_get_type()), F: marshalPrintError},
		{T: externglib.Type(C.gtk_print_operation_action_get_type()), F: marshalPrintOperationAction},
		{T: externglib.Type(C.gtk_print_operation_result_get_type()), F: marshalPrintOperationResult},
		{T: externglib.Type(C.gtk_print_status_get_type()), F: marshalPrintStatus},
		{T: externglib.Type(C.gtk_print_operation_get_type()), F: marshalPrintOperation},
	})
}

// PrintError: error codes that identify various errors that can occur while
// using the GTK printing support.
type PrintError int

const (
	// general: an unspecified error occurred.
	PrintErrorGeneral PrintError = 0
	// InternalError: an internal error occurred.
	PrintErrorInternalError PrintError = 1
	// nomem: a memory allocation failed.
	PrintErrorNOMEM PrintError = 2
	// InvalidFile: an error occurred while loading a page setup or paper size
	// from a key file.
	PrintErrorInvalidFile PrintError = 3
)

func marshalPrintError(p uintptr) (interface{}, error) {
	return PrintError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintOperationAction determines what action the print operation should
// perform.
//
// A parameter of this typs is passed to [method@Gtk.PrintOperation.run].
type PrintOperationAction int

const (
	// PrintDialog: show the print dialog.
	PrintOperationActionPrintDialog PrintOperationAction = 0
	// print: start to print without showing the print dialog, based on the
	// current print settings.
	PrintOperationActionPrint PrintOperationAction = 1
	// preview: show the print preview.
	PrintOperationActionPreview PrintOperationAction = 2
	// export: export to a file. This requires the export-filename property to
	// be set.
	PrintOperationActionExport PrintOperationAction = 3
)

func marshalPrintOperationAction(p uintptr) (interface{}, error) {
	return PrintOperationAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintOperationResult: the result of a print operation.
//
// A value of this type is returned by [method@Gtk.PrintOperation.run].
type PrintOperationResult int

const (
	// error: an error has occurred.
	PrintOperationResultError PrintOperationResult = 0
	// apply: the print settings should be stored.
	PrintOperationResultApply PrintOperationResult = 1
	// cancel: the print operation has been canceled, the print settings should
	// not be stored.
	PrintOperationResultCancel PrintOperationResult = 2
	// InProgress: the print operation is not complete yet. This value will only
	// be returned when running asynchronously.
	PrintOperationResultInProgress PrintOperationResult = 3
)

func marshalPrintOperationResult(p uintptr) (interface{}, error) {
	return PrintOperationResult(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintStatus: the status gives a rough indication of the completion of a
// running print operation.
type PrintStatus int

const (
	// initial: the printing has not started yet; this status is set initially,
	// and while the print dialog is shown.
	PrintStatusInitial PrintStatus = 0
	// preparing: this status is set while the begin-print signal is emitted and
	// during pagination.
	PrintStatusPreparing PrintStatus = 1
	// GeneratingData: this status is set while the pages are being rendered.
	PrintStatusGeneratingData PrintStatus = 2
	// SendingData: the print job is being sent off to the printer.
	PrintStatusSendingData PrintStatus = 3
	// pending: the print job has been sent to the printer, but is not printed
	// for some reason, e.g. the printer may be stopped.
	PrintStatusPending PrintStatus = 4
	// PendingIssue: some problem has occurred during printing, e.g. a paper
	// jam.
	PrintStatusPendingIssue PrintStatus = 5
	// printing: the printer is processing the print job.
	PrintStatusPrinting PrintStatus = 6
	// finished: the printing has been completed successfully.
	PrintStatusFinished PrintStatus = 7
	// FinishedAborted: the printing has been aborted.
	PrintStatusFinishedAborted PrintStatus = 8
)

func marshalPrintStatus(p uintptr) (interface{}, error) {
	return PrintStatus(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintRunPageSetupDialog runs a page setup dialog, letting the user modify the
// values from @page_setup. If the user cancels the dialog, the returned
// PageSetup is identical to the passed in @page_setup, otherwise it contains
// the modifications done in the dialog.
//
// Note that this function may use a recursive mainloop to show the page setup
// dialog. See gtk_print_run_page_setup_dialog_async() if this is a problem.
func PrintRunPageSetupDialog(parent Window, pageSetup PageSetup, settings PrintSettings) PageSetup {
	var _arg1 *C.GtkWindow        // out
	var _arg2 *C.GtkPageSetup     // out
	var _arg3 *C.GtkPrintSettings // out
	var _cret *C.GtkPageSetup     // in

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_run_page_setup_dialog(_arg1, _arg2, _arg3)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

// PrintOperation: `GtkPrintOperation` is the high-level, portable printing API.
//
// It looks a bit different than other GTK dialogs such as the `GtkFileChooser`,
// since some platforms don’t expose enough infrastructure to implement a good
// print dialog. On such platforms, `GtkPrintOperation` uses the native print
// dialog. On platforms which do not provide a native print dialog, GTK uses its
// own, see [class@Gtk.PrintUnixDialog].
//
// The typical way to use the high-level printing API is to create a
// `GtkPrintOperation` object with [ctor@Gtk.PrintOperation.new] when the user
// selects to print. Then you set some properties on it, e.g. the page size, any
// [class@Gtk.PrintSettings] from previous print operations, the number of
// pages, the current page, etc.
//
// Then you start the print operation by calling
// [method@Gtk.PrintOperation.run]. It will then show a dialog, let the user
// select a printer and options. When the user finished the dialog, various
// signals will be emitted on the `GtkPrintOperation`, the main one being
// [signal@Gtk.PrintOperation::draw-page], which you are supposed to handle and
// render the page on the provided [class@Gtk.PrintContext] using Cairo.
//
//
// The high-level printing API
//
// “`c static GtkPrintSettings *settings = NULL;
//
// static void do_print (void) { GtkPrintOperation *print;
// GtkPrintOperationResult res;
//
//    print = gtk_print_operation_new ();
//
//    if (settings != NULL)
//      gtk_print_operation_set_print_settings (print, settings);
//
//    g_signal_connect (print, "begin_print", G_CALLBACK (begin_print), NULL);
//    g_signal_connect (print, "draw_page", G_CALLBACK (draw_page), NULL);
//
//    res = gtk_print_operation_run (print, GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG,
//                                   GTK_WINDOW (main_window), NULL);
//
//    if (res == GTK_PRINT_OPERATION_RESULT_APPLY)
//      {
//        if (settings != NULL)
//          g_object_unref (settings);
//        settings = g_object_ref (gtk_print_operation_get_print_settings (print));
//      }
//
//    g_object_unref (print);
//
// } “`
//
// By default `GtkPrintOperation` uses an external application to do print
// preview. To implement a custom print preview, an application must connect to
// the preview signal. The functions
// [method@Gtk.PrintOperationPreview.render_page],
// [method@Gtk.PrintOperationPreview.end_preview] and
// [method@Gtk.PrintOperationPreview.is_selected] are useful when implementing a
// print preview.
type PrintOperation interface {
	PrintOperationPreview

	// CancelPrintOperation:
	CancelPrintOperation()
	// DrawPageFinishPrintOperation:
	DrawPageFinishPrintOperation()
	// DefaultPageSetup:
	DefaultPageSetup() PageSetup
	// EmbedPageSetup:
	EmbedPageSetup() bool
	// Error:
	Error() error
	// HasSelection:
	HasSelection() bool
	// NPagesToPrint:
	NPagesToPrint() int
	// PrintSettings:
	PrintSettings() PrintSettings
	// Status:
	Status() PrintStatus
	// StatusString:
	StatusString() string
	// SupportSelection:
	SupportSelection() bool
	// IsFinishedPrintOperation:
	IsFinishedPrintOperation() bool
	// RunPrintOperation:
	RunPrintOperation(action PrintOperationAction, parent Window) (PrintOperationResult, error)
	// SetAllowAsyncPrintOperation:
	SetAllowAsyncPrintOperation(allowAsync bool)
	// SetCurrentPagePrintOperation:
	SetCurrentPagePrintOperation(currentPage int)
	// SetCustomTabLabelPrintOperation:
	SetCustomTabLabelPrintOperation(label string)
	// SetDefaultPageSetupPrintOperation:
	SetDefaultPageSetupPrintOperation(defaultPageSetup PageSetup)
	// SetDeferDrawingPrintOperation:
	SetDeferDrawingPrintOperation()
	// SetEmbedPageSetupPrintOperation:
	SetEmbedPageSetupPrintOperation(embed bool)
	// SetExportFilenamePrintOperation:
	SetExportFilenamePrintOperation(filename string)
	// SetHasSelectionPrintOperation:
	SetHasSelectionPrintOperation(hasSelection bool)
	// SetJobNamePrintOperation:
	SetJobNamePrintOperation(jobName string)
	// SetNPagesPrintOperation:
	SetNPagesPrintOperation(nPages int)
	// SetPrintSettingsPrintOperation:
	SetPrintSettingsPrintOperation(printSettings PrintSettings)
	// SetShowProgressPrintOperation:
	SetShowProgressPrintOperation(showProgress bool)
	// SetSupportSelectionPrintOperation:
	SetSupportSelectionPrintOperation(supportSelection bool)
	// SetTrackPrintStatusPrintOperation:
	SetTrackPrintStatusPrintOperation(trackStatus bool)
	// SetUnitPrintOperation:
	SetUnitPrintOperation(unit Unit)
	// SetUseFullPagePrintOperation:
	SetUseFullPagePrintOperation(fullPage bool)
}

// printOperation implements the PrintOperation class.
type printOperation struct {
	gextras.Objector
}

// WrapPrintOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrintOperation(obj *externglib.Object) PrintOperation {
	return printOperation{
		Objector: obj,
	}
}

func marshalPrintOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintOperation(obj), nil
}

// NewPrintOperation:
func NewPrintOperation() PrintOperation {
	var _cret *C.GtkPrintOperation // in

	_cret = C.gtk_print_operation_new()

	var _printOperation PrintOperation // out

	_printOperation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(PrintOperation)

	return _printOperation
}

func (o printOperation) CancelPrintOperation() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_cancel(_arg0)
}

func (o printOperation) DrawPageFinishPrintOperation() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_draw_page_finish(_arg0)
}

func (o printOperation) DefaultPageSetup() PageSetup {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPageSetup      // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_default_page_setup(_arg0)

	var _pageSetup PageSetup // out

	_pageSetup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PageSetup)

	return _pageSetup
}

func (o printOperation) EmbedPageSetup() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_embed_page_setup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o printOperation) Error() error {
	var _arg0 *C.GtkPrintOperation // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_get_error(_arg0, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (o printOperation) HasSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_has_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o printOperation) NPagesToPrint() int {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.int                // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_n_pages_to_print(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (o printOperation) PrintSettings() PrintSettings {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPrintSettings  // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_print_settings(_arg0)

	var _printSettings PrintSettings // out

	_printSettings = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(PrintSettings)

	return _printSettings
}

func (o printOperation) Status() PrintStatus {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.GtkPrintStatus     // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_status(_arg0)

	var _printStatus PrintStatus // out

	_printStatus = PrintStatus(_cret)

	return _printStatus
}

func (o printOperation) StatusString() string {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_status_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (o printOperation) SupportSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_get_support_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o printOperation) IsFinishedPrintOperation() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_print_operation_is_finished(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o printOperation) RunPrintOperation(action PrintOperationAction, parent Window) (PrintOperationResult, error) {
	var _arg0 *C.GtkPrintOperation      // out
	var _arg1 C.GtkPrintOperationAction // out
	var _arg2 *C.GtkWindow              // out
	var _cret C.GtkPrintOperationResult // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.GtkPrintOperationAction(action)
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_print_operation_run(_arg0, _arg1, _arg2, &_cerr)

	var _printOperationResult PrintOperationResult // out
	var _goerr error                               // out

	_printOperationResult = PrintOperationResult(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _printOperationResult, _goerr
}

func (o printOperation) SetAllowAsyncPrintOperation(allowAsync bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if allowAsync {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_allow_async(_arg0, _arg1)
}

func (o printOperation) SetCurrentPagePrintOperation(currentPage int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.int(currentPage)

	C.gtk_print_operation_set_current_page(_arg0, _arg1)
}

func (o printOperation) SetCustomTabLabelPrintOperation(label string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_custom_tab_label(_arg0, _arg1)
}

func (o printOperation) SetDefaultPageSetupPrintOperation(defaultPageSetup PageSetup) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPageSetup      // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(defaultPageSetup.Native()))

	C.gtk_print_operation_set_default_page_setup(_arg0, _arg1)
}

func (o printOperation) SetDeferDrawingPrintOperation() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_set_defer_drawing(_arg0)
}

func (o printOperation) SetEmbedPageSetupPrintOperation(embed bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if embed {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_embed_page_setup(_arg0, _arg1)
}

func (o printOperation) SetExportFilenamePrintOperation(filename string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_export_filename(_arg0, _arg1)
}

func (o printOperation) SetHasSelectionPrintOperation(hasSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if hasSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_has_selection(_arg0, _arg1)
}

func (o printOperation) SetJobNamePrintOperation(jobName string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(jobName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_job_name(_arg0, _arg1)
}

func (o printOperation) SetNPagesPrintOperation(nPages int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.int                // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.int(nPages)

	C.gtk_print_operation_set_n_pages(_arg0, _arg1)
}

func (o printOperation) SetPrintSettingsPrintOperation(printSettings PrintSettings) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPrintSettings  // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(printSettings.Native()))

	C.gtk_print_operation_set_print_settings(_arg0, _arg1)
}

func (o printOperation) SetShowProgressPrintOperation(showProgress bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if showProgress {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_show_progress(_arg0, _arg1)
}

func (o printOperation) SetSupportSelectionPrintOperation(supportSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if supportSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_support_selection(_arg0, _arg1)
}

func (o printOperation) SetTrackPrintStatusPrintOperation(trackStatus bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if trackStatus {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_track_print_status(_arg0, _arg1)
}

func (o printOperation) SetUnitPrintOperation(unit Unit) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.GtkUnit            // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.GtkUnit(unit)

	C.gtk_print_operation_set_unit(_arg0, _arg1)
}

func (o printOperation) SetUseFullPagePrintOperation(fullPage bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if fullPage {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_use_full_page(_arg0, _arg1)
}

func (p printOperation) EndPreview() {
	WrapPrintOperationPreview(gextras.InternObject(p)).EndPreview()
}

func (p printOperation) IsSelected(pageNr int) bool {
	return WrapPrintOperationPreview(gextras.InternObject(p)).IsSelected(pageNr)
}

func (p printOperation) RenderPage(pageNr int) {
	WrapPrintOperationPreview(gextras.InternObject(p)).RenderPage(pageNr)
}
