// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_provider_error_get_type()), F: marshalCSSProviderError},
		{T: externglib.Type(C.gtk_css_provider_get_type()), F: marshalCSSProviderer},
	})
}

// CSSProviderError: error codes for GTK_CSS_PROVIDER_ERROR.
type CSSProviderError int

const (
	// CSSProviderErrorFailed: failed.
	CSSProviderErrorFailed CSSProviderError = iota
	// CSSProviderErrorSyntax: syntax error.
	CSSProviderErrorSyntax
	// CSSProviderErrorImport: import error.
	CSSProviderErrorImport
	// CSSProviderErrorName: name error.
	CSSProviderErrorName
	// CSSProviderErrorDeprecated: deprecation error.
	CSSProviderErrorDeprecated
	// CSSProviderErrorUnknownValue: unknown value.
	CSSProviderErrorUnknownValue
)

func marshalCSSProviderError(p uintptr) (interface{}, error) {
	return CSSProviderError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for CSSProviderError.
func (c CSSProviderError) String() string {
	switch c {
	case CSSProviderErrorFailed:
		return "Failed"
	case CSSProviderErrorSyntax:
		return "Syntax"
	case CSSProviderErrorImport:
		return "Import"
	case CSSProviderErrorName:
		return "Name"
	case CSSProviderErrorDeprecated:
		return "Deprecated"
	case CSSProviderErrorUnknownValue:
		return "UnknownValue"
	default:
		return fmt.Sprintf("CSSProviderError(%d)", c)
	}
}

// CSSProviderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CSSProviderOverrider interface {
	ParsingError(section *CSSSection, err error)
}

// CSSProvider is an object implementing the StyleProvider interface. It is able
// to parse [CSS-like][css-overview] input in order to style widgets.
//
// An application can make GTK+ parse a specific CSS style sheet by calling
// gtk_css_provider_load_from_file() or gtk_css_provider_load_from_resource()
// and adding the provider with gtk_style_context_add_provider() or
// gtk_style_context_add_provider_for_screen().
//
// In addition, certain files will be read when GTK+ is initialized. First, the
// file $XDG_CONFIG_HOME/gtk-3.0/gtk.css is loaded if it exists. Then, GTK+
// loads the first existing file among
// XDG_DATA_HOME/themes/THEME/gtk-VERSION/gtk.css,
// $HOME/.themes/THEME/gtk-VERSION/gtk.css,
// $XDG_DATA_DIRS/themes/THEME/gtk-VERSION/gtk.css and
// DATADIR/share/themes/THEME/gtk-VERSION/gtk.css, where THEME is the name of
// the current theme (see the Settings:gtk-theme-name setting), DATADIR is the
// prefix configured when GTK+ was compiled (unless overridden by the
// GTK_DATA_PREFIX environment variable), and VERSION is the GTK+ version
// number. If no file is found for the current version, GTK+ tries older
// versions all the way back to 3.0.
//
// In the same way, GTK+ tries to load a gtk-keys.css file for the current key
// theme, as defined by Settings:gtk-key-theme-name.
type CSSProvider struct {
	*externglib.Object

	StyleProvider
}

func WrapCSSProvider(obj *externglib.Object) *CSSProvider {
	return &CSSProvider{
		Object: obj,
		StyleProvider: StyleProvider{
			Object: obj,
		},
	}
}

func marshalCSSProviderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCSSProvider(obj), nil
}

// NewCSSProvider returns a newly created CssProvider.
func NewCSSProvider() *CSSProvider {
	var _cret *C.GtkCssProvider // in

	_cret = C.gtk_css_provider_new()

	var _cssProvider *CSSProvider // out

	_cssProvider = WrapCSSProvider(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _cssProvider
}

// LoadFromData loads data into css_provider, and by doing so clears any
// previously loaded information.
func (cssProvider *CSSProvider) LoadFromData(data []byte) error {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar          // out
	var _arg2 C.gssize
	var _cerr *C.GError // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg2 = (C.gssize)(len(data))
	if len(data) > 0 {
		_arg1 = (*C.gchar)(unsafe.Pointer(&data[0]))
	}

	C.gtk_css_provider_load_from_data(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(data)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromFile loads the data contained in file into css_provider, making it
// clear any previously loaded information.
func (cssProvider *CSSProvider) LoadFromFile(file gio.Filer) error {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.GFile          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_css_provider_load_from_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(file)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromPath loads the data contained in path into css_provider, making it
// clear any previously loaded information.
func (cssProvider *CSSProvider) LoadFromPath(path string) error {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar          // out
	var _cerr *C.GError         // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_path(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(path)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadFromResource loads the data contained in the resource at resource_path
// into the CssProvider, clearing any previously loaded information.
//
// To track errors while loading CSS, connect to the CssProvider::parsing-error
// signal.
func (cssProvider *CSSProvider) LoadFromResource(resourcePath string) {
	var _arg0 *C.GtkCssProvider // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(cssProvider.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_css_provider_load_from_resource(_arg0, _arg1)
	runtime.KeepAlive(cssProvider)
	runtime.KeepAlive(resourcePath)
}

// String converts the provider into a string representation in CSS format.
//
// Using gtk_css_provider_load_from_data() with the return value from this
// function on a new provider created with gtk_css_provider_new() will basically
// create a duplicate of this provider.
func (provider *CSSProvider) String() string {
	var _arg0 *C.GtkCssProvider // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkCssProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gtk_css_provider_to_string(_arg0)
	runtime.KeepAlive(provider)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// CSSProviderGetDefault returns the provider containing the style settings used
// as a fallback for all widgets.
//
// Deprecated: Use gtk_css_provider_new() instead.
func CSSProviderGetDefault() *CSSProvider {
	var _cret *C.GtkCssProvider // in

	_cret = C.gtk_css_provider_get_default()

	var _cssProvider *CSSProvider // out

	_cssProvider = WrapCSSProvider(externglib.Take(unsafe.Pointer(_cret)))

	return _cssProvider
}

// CSSProviderGetNamed loads a theme from the usual theme paths
func CSSProviderGetNamed(name string, variant string) *CSSProvider {
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _cret *C.GtkCssProvider // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if variant != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variant)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_css_provider_get_named(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(variant)

	var _cssProvider *CSSProvider // out

	_cssProvider = WrapCSSProvider(externglib.Take(unsafe.Pointer(_cret)))

	return _cssProvider
}
