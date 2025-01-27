// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_widget_get_type()), F: marshalFontChooserWidgetter},
	})
}

// FontChooserWidget: GtkFontChooserWidget widget lets the user select a font.
//
// It is used in the GtkFontChooserDialog widget to provide a dialog for
// selecting fonts.
//
// To set the font which is initially selected, use gtk.FontChooser.SetFont() or
// gtk.FontChooser.SetFontDesc().
//
// To get the selected font use gtk.FontChooser.GetFont() or
// gtk.FontChooser.GetFontDesc().
//
// To change the text which is shown in the preview area, use
// gtk.FontChooser.SetPreviewText().
//
//
// CSS nodes
//
// GtkFontChooserWidget has a single CSS node with name fontchooser.
type FontChooserWidget struct {
	Widget

	FontChooser
	*externglib.Object
}

func WrapFontChooserWidget(obj *externglib.Object) *FontChooserWidget {
	return &FontChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		FontChooser: FontChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalFontChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooserWidget(obj), nil
}

// NewFontChooserWidget creates a new GtkFontChooserWidget.
func NewFontChooserWidget() *FontChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_font_chooser_widget_new()

	var _fontChooserWidget *FontChooserWidget // out

	_fontChooserWidget = WrapFontChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _fontChooserWidget
}

func (*FontChooserWidget) privateFontChooserWidget() {}
