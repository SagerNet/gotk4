// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_license_get_type()), F: marshalLicense},
		{T: externglib.Type(C.gtk_about_dialog_get_type()), F: marshalAboutDialog},
	})
}

// License: the type of license for an application.
//
// This enumeration can be expanded at later date.
type License int

const (
	// unknown: no license specified
	LicenseUnknown License = iota
	// custom: license text is going to be specified by the developer
	LicenseCustom
	// GPL20: the GNU General Public License, version 2.0 or later
	LicenseGPL20
	// GPL30: the GNU General Public License, version 3.0 or later
	LicenseGPL30
	// LGPL21: the GNU Lesser General Public License, version 2.1 or later
	LicenseLGPL21
	// LGPL30: the GNU Lesser General Public License, version 3.0 or later
	LicenseLGPL30
	// bsd: the BSD standard license
	LicenseBSD
	// MITX11: the MIT/X11 standard license
	LicenseMITX11
	// artistic: the Artistic License, version 2.0
	LicenseArtistic
	// GPL20Only: the GNU General Public License, version 2.0 only. Since 3.12.
	LicenseGPL20Only
	// GPL30Only: the GNU General Public License, version 3.0 only. Since 3.12.
	LicenseGPL30Only
	// LGPL21Only: the GNU Lesser General Public License, version 2.1 only.
	// Since 3.12.
	LicenseLGPL21Only
	// LGPL30Only: the GNU Lesser General Public License, version 3.0 only.
	// Since 3.12.
	LicenseLGPL30Only
	// AGPL30: the GNU Affero General Public License, version 3.0 or later.
	// Since: 3.22.
	LicenseAGPL30
	// AGPL30Only: the GNU Affero General Public License, version 3.0 only.
	// Since: 3.22.27.
	LicenseAGPL30Only
	// BSD3: the 3-clause BSD licence. Since: 3.24.20.
	LicenseBSD3
	// Apache20: the Apache License, version 2.0. Since: 3.24.20.
	LicenseApache20
	// MPL20: the Mozilla Public License, version 2.0. Since: 3.24.20.
	LicenseMPL20
)

func marshalLicense(p uintptr) (interface{}, error) {
	return License(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AboutDialog: the GtkAboutDialog offers a simple way to display information
// about a program like its logo, name, copyright, website and license. It is
// also possible to give credits to the authors, documenters, translators and
// artists who have worked on the program. An about dialog is typically opened
// when the user selects the `About` option from the `Help` menu. All parts of
// the dialog are optional.
//
// About dialogs often contain links and email addresses. GtkAboutDialog
// displays these as clickable links. By default, it calls
// gtk_show_uri_on_window() when a user clicks one. The behaviour can be
// overridden with the AboutDialog::activate-link signal.
//
// To specify a person with an email address, use a string like "Edgar Allan Poe
// <edgar\@poe.com>". To specify a website with a title, use a string like "GTK+
// team http://www.gtk.org".
//
// To make constructing a GtkAboutDialog as convenient as possible, you can use
// the function gtk_show_about_dialog() which constructs and shows a dialog and
// keeps it around so that it can be shown again.
//
// Note that GTK+ sets a default title of `_("About s")` on the dialog window
// (where \s is replaced by the name of the application, but in order to ensure
// proper translation of the title, applications should set the title property
// explicitly when constructing a GtkAboutDialog, as shown in the following
// example:
//
//    GdkPixbuf *example_logo = gdk_pixbuf_new_from_file ("./logo.png", NULL);
//    gtk_show_about_dialog (NULL,
//                           "program-name", "ExampleCode",
//                           "logo", example_logo,
//                           "title", _("About ExampleCode"),
//                           NULL);
//
// It is also possible to show a AboutDialog like any other Dialog, e.g. using
// gtk_dialog_run(). In this case, you might need to know that the “Close”
// button returns the K_RESPONSE_CANCEL response id.
type AboutDialog interface {
	Dialog

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// AddCreditSection creates a new section in the Credits page.
	AddCreditSection(sectionName string, people []string)
	// Artists returns the string which are displayed in the artists tab of the
	// secondary credits dialog.
	Artists() []string
	// Authors returns the string which are displayed in the authors tab of the
	// secondary credits dialog.
	Authors() []string
	// Comments returns the comments string.
	Comments() string
	// Copyright returns the copyright string.
	Copyright() string
	// Documenters returns the string which are displayed in the documenters tab
	// of the secondary credits dialog.
	Documenters() []string
	// License returns the license information.
	License() string
	// LicenseType retrieves the license set using
	// gtk_about_dialog_set_license_type()
	LicenseType() License
	// Logo returns the pixbuf displayed as logo in the about dialog.
	Logo() gdkpixbuf.Pixbuf
	// LogoIconName returns the icon name displayed as logo in the about dialog.
	LogoIconName() string
	// ProgramName returns the program name displayed in the about dialog.
	ProgramName() string
	// TranslatorCredits returns the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	TranslatorCredits() string
	// Version returns the version string.
	Version() string
	// Website returns the website URL.
	Website() string
	// WebsiteLabel returns the label used for the website link.
	WebsiteLabel() string
	// WrapLicense returns whether the license text in @about is automatically
	// wrapped.
	WrapLicense() bool
	// SetArtists sets the strings which are displayed in the artists tab of the
	// secondary credits dialog.
	SetArtists(artists []string)
	// SetAuthors sets the strings which are displayed in the authors tab of the
	// secondary credits dialog.
	SetAuthors(authors []string)
	// SetComments sets the comments string to display in the about dialog. This
	// should be a short string of one or two lines.
	SetComments(comments string)
	// SetCopyright sets the copyright string to display in the about dialog.
	// This should be a short string of one or two lines.
	SetCopyright(copyright string)
	// SetDocumenters sets the strings which are displayed in the documenters
	// tab of the secondary credits dialog.
	SetDocumenters(documenters []string)
	// SetLicense sets the license information to be displayed in the secondary
	// license dialog. If @license is nil, the license button is hidden.
	SetLicense(license string)
	// SetLicenseType sets the license of the application showing the @about
	// dialog from a list of known licenses.
	//
	// This function overrides the license set using
	// gtk_about_dialog_set_license().
	SetLicenseType(licenseType License)
	// SetLogo sets the pixbuf to be displayed as logo in the about dialog. If
	// it is nil, the default window icon set with gtk_window_set_default_icon()
	// will be used.
	SetLogo(logo gdkpixbuf.Pixbuf)
	// SetLogoIconName sets the pixbuf to be displayed as logo in the about
	// dialog. If it is nil, the default window icon set with
	// gtk_window_set_default_icon() will be used.
	SetLogoIconName(iconName string)
	// SetProgramName sets the name to display in the about dialog. If this is
	// not set, it defaults to g_get_application_name().
	SetProgramName(name string)
	// SetTranslatorCredits sets the translator credits string which is
	// displayed in the translators tab of the secondary credits dialog.
	//
	// The intended use for this string is to display the translator of the
	// language which is currently used in the user interface. Using gettext(),
	// a simple way to achieve that is to mark the string for translation:
	//
	//    GtkWidget *about = gtk_about_dialog_new ();
	//    gtk_about_dialog_set_translator_credits (GTK_ABOUT_DIALOG (about),
	//                                             _("translator-credits"));
	//
	// It is a good idea to use the customary msgid “translator-credits” for
	// this purpose, since translators will already know the purpose of that
	// msgid, and since AboutDialog will detect if “translator-credits” is
	// untranslated and hide the tab.
	SetTranslatorCredits(translatorCredits string)
	// SetVersion sets the version string to display in the about dialog.
	SetVersion(version string)
	// SetWebsite sets the URL to use for the website link.
	SetWebsite(website string)
	// SetWebsiteLabel sets the label to be used for the website link.
	SetWebsiteLabel(websiteLabel string)
	// SetWrapLicense sets whether the license text in @about is automatically
	// wrapped.
	SetWrapLicense(wrapLicense bool)
}

// aboutDialog implements the AboutDialog class.
type aboutDialog struct {
	Dialog
}

// WrapAboutDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapAboutDialog(obj *externglib.Object) AboutDialog {
	return aboutDialog{
		Dialog: WrapDialog(obj),
	}
}

func marshalAboutDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAboutDialog(obj), nil
}

// NewAboutDialog creates a new AboutDialog.
func NewAboutDialog() AboutDialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_about_dialog_new()

	var _aboutDialog AboutDialog // out

	_aboutDialog = WrapAboutDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _aboutDialog
}

func (a aboutDialog) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(a))
}

func (a aboutDialog) AddCreditSection(sectionName string, people []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out
	var _arg2 **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(sectionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.gchar)(C.malloc(C.ulong(len(people)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(people))
		for i := range people {
			out[i] = (*C.gchar)(C.CString(people[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_add_credit_section(_arg0, _arg1, _arg2)
}

func (a aboutDialog) Artists() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_artists(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

func (a aboutDialog) Authors() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_authors(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

func (a aboutDialog) Comments() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_comments(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) Copyright() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_copyright(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) Documenters() []string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_documenters(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

func (a aboutDialog) License() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_license(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) LicenseType() License {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.GtkLicense      // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_license_type(_arg0)

	var _license License // out

	_license = License(_cret)

	return _license
}

func (a aboutDialog) Logo() gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.GdkPixbuf      // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_logo(_arg0)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (a aboutDialog) LogoIconName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_logo_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) ProgramName() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_program_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) TranslatorCredits() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_translator_credits(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) Version() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_version(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) Website() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_website(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) WebsiteLabel() string {
	var _arg0 *C.GtkAboutDialog // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_website_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a aboutDialog) WrapLicense() bool {
	var _arg0 *C.GtkAboutDialog // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_about_dialog_get_wrap_license(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a aboutDialog) SetArtists(artists []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (**C.gchar)(C.malloc(C.ulong(len(artists)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(artists))
		for i := range artists {
			out[i] = (*C.gchar)(C.CString(artists[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_artists(_arg0, _arg1)
}

func (a aboutDialog) SetAuthors(authors []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (**C.gchar)(C.malloc(C.ulong(len(authors)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(authors))
		for i := range authors {
			out[i] = (*C.gchar)(C.CString(authors[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_authors(_arg0, _arg1)
}

func (a aboutDialog) SetComments(comments string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(comments))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_comments(_arg0, _arg1)
}

func (a aboutDialog) SetCopyright(copyright string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(copyright))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_copyright(_arg0, _arg1)
}

func (a aboutDialog) SetDocumenters(documenters []string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 **C.gchar

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (**C.gchar)(C.malloc(C.ulong(len(documenters)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(documenters))
		for i := range documenters {
			out[i] = (*C.gchar)(C.CString(documenters[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_about_dialog_set_documenters(_arg0, _arg1)
}

func (a aboutDialog) SetLicense(license string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(license))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_license(_arg0, _arg1)
}

func (a aboutDialog) SetLicenseType(licenseType License) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.GtkLicense      // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = C.GtkLicense(licenseType)

	C.gtk_about_dialog_set_license_type(_arg0, _arg1)
}

func (a aboutDialog) SetLogo(logo gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.GdkPixbuf      // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(logo.Native()))

	C.gtk_about_dialog_set_logo(_arg0, _arg1)
}

func (a aboutDialog) SetLogoIconName(iconName string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_logo_icon_name(_arg0, _arg1)
}

func (a aboutDialog) SetProgramName(name string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_program_name(_arg0, _arg1)
}

func (a aboutDialog) SetTranslatorCredits(translatorCredits string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(translatorCredits))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_translator_credits(_arg0, _arg1)
}

func (a aboutDialog) SetVersion(version string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(version))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_version(_arg0, _arg1)
}

func (a aboutDialog) SetWebsite(website string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(website))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_website(_arg0, _arg1)
}

func (a aboutDialog) SetWebsiteLabel(websiteLabel string) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(websiteLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_about_dialog_set_website_label(_arg0, _arg1)
}

func (a aboutDialog) SetWrapLicense(wrapLicense bool) {
	var _arg0 *C.GtkAboutDialog // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAboutDialog)(unsafe.Pointer(a.Native()))
	if wrapLicense {
		_arg1 = C.TRUE
	}

	C.gtk_about_dialog_set_wrap_license(_arg0, _arg1)
}
