// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

func CSSParserWarningQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_css_parser_warning_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}