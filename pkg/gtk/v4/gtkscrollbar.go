// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_scrollbar_get_type()), F: marshalScrollbarrer},
	})
}

// Scrollbar: GtkScrollbar widget is a horizontal or vertical scrollbar.
//
// !An example GtkScrollbar (scrollbar.png)
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by gtk.Scrollbar.New. See [class.Gtk.Adjustment] for more details.
// The gtk.Adjustment:value field sets the position of the thumb and must be
// between gtk.Adjustment:lower and gtk.Adjustment:upper -
// gtk.Adjustment:page-size. The gtk.Adjustment:page-size represents the size of
// the visible scrollable area.
//
// The fields gtk.Adjustment:step-increment and gtk.Adjustment:page-increment
// fields are added to or subtracted from the gtk.Adjustment:value when the user
// asks to move by a step (using e.g. the cursor arrow keys) or by a page (using
// e.g. the Page Down/Up keys).
//
// CSS nodes
//
//    scrollbar
//    ╰── range[.fine-tune]
//        ╰── trough
//            ╰── slider
//
//
// GtkScrollbar has a main CSS node with name scrollbar and a subnode for its
// contents. The main node gets the .horizontal or .vertical style classes
// applied, depending on the scrollbar's orientation.
//
// The range node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// Other style classes that may be added to scrollbars inside gtk.ScrolledWindow
// include the positional classes (.left, .right, .top, .bottom) and style
// classes related to overlay scrolling (.overlay-indicator, .dragging,
// .hovering).
//
//
// Accessibility
//
// GtkScrollbar uses the GTK_ACCESSIBLE_ROLE_SCROLLBAR role.
type Scrollbar struct {
	Widget

	Orientable
	*externglib.Object
}

func wrapScrollbar(obj *externglib.Object) *Scrollbar {
	return &Scrollbar{
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
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalScrollbarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapScrollbar(obj), nil
}

// NewScrollbar creates a new scrollbar with the given orientation.
func NewScrollbar(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	if adjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(adjustment)

	var _scrollbar *Scrollbar // out

	_scrollbar = wrapScrollbar(externglib.Take(unsafe.Pointer(_cret)))

	return _scrollbar
}

// Adjustment returns the scrollbar's adjustment.
func (self *Scrollbar) Adjustment() *Adjustment {
	var _arg0 *C.GtkScrollbar  // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_scrollbar_get_adjustment(_arg0)
	runtime.KeepAlive(self)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// SetAdjustment makes the scrollbar use the given adjustment.
func (self *Scrollbar) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkScrollbar  // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollbar)(unsafe.Pointer(self.Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}

	C.gtk_scrollbar_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(adjustment)
}
