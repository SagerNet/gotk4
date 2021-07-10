// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_page_setup_get_type()), F: marshalPageSetupper},
	})
}

// PageSetupper describes PageSetup's methods.
type PageSetupper interface {
	gextras.Objector

	Copy() *PageSetup
	Orientation() PageOrientation
	PaperSize() *PaperSize
	LoadFile(fileName string) error
	LoadKeyFile(keyFile *glib.KeyFile, groupName string) error
	SetPaperSize(size *PaperSize)
	SetPaperSizeAndDefaultMargins(size *PaperSize)
	ToFile(fileName string) error
	ToGVariant() *glib.Variant
	ToKeyFile(keyFile *glib.KeyFile, groupName string)
}

// PageSetup object stores the page size, orientation and margins. The idea is
// that you can get one of these from the page setup dialog and then pass it to
// the PrintOperation when printing. The benefit of splitting this out of the
// PrintSettings is that these affect the actual layout of the page, and thus
// need to be set long before user prints.
//
//
// Margins
//
// The margins specified in this object are the “print margins”, i.e. the parts
// of the page that the printer cannot print on. These are different from the
// layout margins that a word processor uses; they are typically used to
// determine the minimal size for the layout margins.
//
// To obtain a PageSetup use gtk_page_setup_new() to get the defaults, or use
// gtk_print_run_page_setup_dialog() to show the page setup dialog and receive
// the resulting page setup.
//
// A page setup dialog
//
//    static GtkPrintSettings *settings = NULL;
//    static GtkPageSetup *page_setup = NULL;
//
//    static void
//    do_page_setup (void)
//    {
//      GtkPageSetup *new_page_setup;
//
//      if (settings == NULL)
//        settings = gtk_print_settings_new ();
//
//      new_page_setup = gtk_print_run_page_setup_dialog (GTK_WINDOW (main_window),
//                                                        page_setup, settings);
//
//      if (page_setup)
//        g_object_unref (page_setup);
//
//      page_setup = new_page_setup;
//    }
//
// Printing support was added in GTK+ 2.10.
type PageSetup struct {
	*externglib.Object
}

var _ PageSetupper = (*PageSetup)(nil)

func wrapPageSetupper(obj *externglib.Object) PageSetupper {
	return &PageSetup{
		Object: obj,
	}
}

func marshalPageSetupper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPageSetupper(obj), nil
}

// NewPageSetup creates a new PageSetup.
func NewPageSetup() *PageSetup {
	var _cret *C.GtkPageSetup // in

	_cret = C.gtk_page_setup_new()

	var _pageSetup *PageSetup // out

	_pageSetup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PageSetup)

	return _pageSetup
}

// NewPageSetupFromFile reads the page setup from the file @file_name. Returns a
// new PageSetup object with the restored page setup, or nil if an error
// occurred. See gtk_page_setup_to_file().
func NewPageSetupFromFile(fileName string) (*PageSetup, error) {
	var _arg1 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_page_setup_new_from_file(_arg1, &_cerr)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PageSetup)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pageSetup, _goerr
}

// NewPageSetupFromGVariant: desrialize a page setup from an a{sv} variant in
// the format produced by gtk_page_setup_to_gvariant().
func NewPageSetupFromGVariant(variant *glib.Variant) *PageSetup {
	var _arg1 *C.GVariant     // out
	var _cret *C.GtkPageSetup // in

	_arg1 = (*C.GVariant)(unsafe.Pointer(variant))

	_cret = C.gtk_page_setup_new_from_gvariant(_arg1)

	var _pageSetup *PageSetup // out

	_pageSetup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PageSetup)

	return _pageSetup
}

// NewPageSetupFromKeyFile reads the page setup from the group @group_name in
// the key file @key_file. Returns a new PageSetup object with the restored page
// setup, or nil if an error occurred.
func NewPageSetupFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cret *C.GtkPageSetup // in
	var _cerr *C.GError       // in

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_page_setup_new_from_key_file(_arg1, _arg2, &_cerr)

	var _pageSetup *PageSetup // out
	var _goerr error          // out

	_pageSetup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PageSetup)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _pageSetup, _goerr
}

// Copy copies a PageSetup.
func (other *PageSetup) Copy() *PageSetup {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPageSetup // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(other.Native()))

	_cret = C.gtk_page_setup_copy(_arg0)

	var _pageSetup *PageSetup // out

	_pageSetup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PageSetup)

	return _pageSetup
}

// Orientation gets the page orientation of the PageSetup.
func (setup *PageSetup) Orientation() PageOrientation {
	var _arg0 *C.GtkPageSetup      // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))

	_cret = C.gtk_page_setup_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = (PageOrientation)(_cret)

	return _pageOrientation
}

// PaperSize gets the paper size of the PageSetup.
func (setup *PageSetup) PaperSize() *PaperSize {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GtkPaperSize // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))

	_cret = C.gtk_page_setup_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(unsafe.Pointer(_cret))

	return _paperSize
}

// LoadFile reads the page setup from the file @file_name. See
// gtk_page_setup_to_file().
func (setup *PageSetup) LoadFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// LoadKeyFile reads the page setup from the group @group_name in the key file
// @key_file.
func (setup *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_page_setup_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetPaperSize sets the paper size of the PageSetup without changing the
// margins. See gtk_page_setup_set_paper_size_and_default_margins().
func (setup *PageSetup) SetPaperSize(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size))

	C.gtk_page_setup_set_paper_size(_arg0, _arg1)
}

// SetPaperSizeAndDefaultMargins sets the paper size of the PageSetup and
// modifies the margins according to the new paper size.
func (setup *PageSetup) SetPaperSizeAndDefaultMargins(size *PaperSize) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GtkPaperSize // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.GtkPaperSize)(unsafe.Pointer(size))

	C.gtk_page_setup_set_paper_size_and_default_margins(_arg0, _arg1)
}

// ToFile: this function saves the information from @setup to @file_name.
func (setup *PageSetup) ToFile(fileName string) error {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.char         // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.char)(C.CString(fileName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_page_setup_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ToGVariant: serialize page setup to an a{sv} variant.
func (setup *PageSetup) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPageSetup // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))

	_cret = C.gtk_page_setup_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

// ToKeyFile: this function adds the page setup from @setup to @key_file.
func (setup *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPageSetup // out
	var _arg1 *C.GKeyFile     // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkPageSetup)(unsafe.Pointer(setup.Native()))
	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile))
	_arg2 = (*C.gchar)(C.CString(groupName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_page_setup_to_key_file(_arg0, _arg1, _arg2)
}
