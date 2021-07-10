// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_center_box_get_type()), F: marshalCenterBoxxer},
	})
}

// CenterBoxxer describes CenterBox's methods.
type CenterBoxxer interface {
	gextras.Objector

	BaselinePosition() BaselinePosition
	CenterWidget() *Widget
	EndWidget() *Widget
	StartWidget() *Widget
	SetCenterWidget(child Widgetter)
	SetEndWidget(child Widgetter)
	SetStartWidget(child Widgetter)
}

// CenterBox: `GtkCenterBox` arranges three children in a row, keeping the
// middle child centered as well as possible.
//
// !An example GtkCenterBox (centerbox.png)
//
// To add children to `GtkCenterBox`, use
// [method@Gtk.CenterBox.set_start_widget],
// [method@Gtk.CenterBox.set_center_widget] and
// [method@Gtk.CenterBox.set_end_widget].
//
// The sizing and positioning of children can be influenced with the align and
// expand properties of the children.
//
//
// GtkCenterBox as GtkBuildable
//
// The `GtkCenterBox` implementation of the `GtkBuildable` interface supports
// placing children in the 3 positions by specifying “start”, “center” or “end”
// as the “type” attribute of a <child> element.
//
//
// CSS nodes
//
// `GtkCenterBox` uses a single CSS node with the name “box”,
//
// The first child of the `GtkCenterBox` will be allocated depending on the text
// direction, i.e. in left-to-right layouts it will be allocated on the left and
// in right-to-left layouts on the right.
//
// In vertical orientation, the nodes of the children are arranged from top to
// bottom.
//
//
// Accessibility
//
// `GtkCenterBox` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type CenterBox struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ CenterBoxxer = (*CenterBox)(nil)

func wrapCenterBoxxer(obj *externglib.Object) CenterBoxxer {
	return &CenterBox{
		Object: obj,
		Widget: Widget{
			Object: obj,
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
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalCenterBoxxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCenterBoxxer(obj), nil
}

// NewCenterBox creates a new `GtkCenterBox`.
func NewCenterBox() *CenterBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_center_box_new()

	var _centerBox *CenterBox // out

	_centerBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CenterBox)

	return _centerBox
}

// BaselinePosition gets the value set by
// gtk_center_box_set_baseline_position().
func (self *CenterBox) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkCenterBox       // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_baseline_position(_arg0)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = (BaselinePosition)(_cret)

	return _baselinePosition
}

// CenterWidget gets the center widget, or nil if there is none.
func (self *CenterBox) CenterWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_center_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// EndWidget gets the end widget, or nil if there is none.
func (self *CenterBox) EndWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_end_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// StartWidget gets the start widget, or nil if there is none.
func (self *CenterBox) StartWidget() *Widget {
	var _arg0 *C.GtkCenterBox // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_center_box_get_start_widget(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// SetCenterWidget sets the center widget.
//
// To remove the existing center widget, pas nil.
func (self *CenterBox) SetCenterWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_center_widget(_arg0, _arg1)
}

// SetEndWidget sets the end widget.
//
// To remove the existing end widget, pass nil.
func (self *CenterBox) SetEndWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_end_widget(_arg0, _arg1)
}

// SetStartWidget sets the start widget.
//
// To remove the existing start widget, pass nil.
func (self *CenterBox) SetStartWidget(child Widgetter) {
	var _arg0 *C.GtkCenterBox // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkCenterBox)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_center_box_set_start_widget(_arg0, _arg1)
}
