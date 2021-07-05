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
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFrame},
	})
}

// Frame: `GtkFrame` is a widget that surrounds its child with a decorative
// frame and an optional label.
//
// !An example GtkFrame (frame.png)
//
// If present, the label is drawn inside the top edge of the frame. The
// horizontal position of the label can be controlled with
// [method@Gtk.Frame.set_label_align].
//
// `GtkFrame` clips its child. You can use this to add rounded corners to
// widgets, but be aware that it also cuts off shadows.
//
//
// GtkFrame as GtkBuildable
//
// The `GtkFrame` implementation of the `GtkBuildable` interface supports
// placing a child in the label position by specifying “label” as the “type”
// attribute of a <child> element. A normal content child can be specified
// without specifying a <child> type attribute.
//
// An example of a UI definition fragment with GtkFrame: “`xml <object
// class="GtkFrame"> <child type="label"> <object class="GtkLabel"
// id="frame_label"/> </child> <child> <object class="GtkEntry"
// id="frame_content"/> </child> </object> “`
//
//
// CSS nodes
//
// “` frame ├── <label widget> ╰── <child> “`
//
// `GtkFrame` has a main CSS node with name “frame”, which is used to draw the
// visible border. You can set the appearance of the border using CSS properties
// like “border-style” on this node.
type Frame interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget

	// Child gets the child widget of @frame.
	Child() Widget
	// Label returns the frame labels text.
	//
	// If the frame's label widget is not a `GtkLabel`, nil is returned.
	Label() string
	// LabelAlign retrieves the X alignment of the frame’s label.
	LabelAlign() float32
	// LabelWidget retrieves the label widget for the frame.
	LabelWidget() Widget
	// SetChild sets the child widget of @frame.
	SetChild(child Widget)
	// SetLabel creates a new `GtkLabel` with the @label and sets it as the
	// frame's label widget.
	SetLabel(label string)
	// SetLabelAlign sets the X alignment of the frame widget’s label.
	//
	// The default value for a newly created frame is 0.0.
	SetLabelAlign(xalign float32)
	// SetLabelWidget sets the label widget for the frame.
	//
	// This is the widget that will appear embedded in the top edge of the frame
	// as a title.
	SetLabelWidget(labelWidget Widget)
}

// frame implements the Frame class.
type frame struct {
	Widget
}

// WrapFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapFrame(obj *externglib.Object) Frame {
	return frame{
		Widget: WrapWidget(obj),
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFrame(obj), nil
}

// NewFrame creates a new `GtkFrame`, with optional label @label.
//
// If @label is nil, the label is omitted.
func NewFrame(label string) Frame {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_frame_new(_arg1)

	var _frame Frame // out

	_frame = WrapFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

func (f frame) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(f))
}

func (f frame) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(f))
}

func (f frame) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(f))
}

func (f frame) Child() Widget {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f frame) Label() string {
	var _arg0 *C.GtkFrame // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f frame) LabelAlign() float32 {
	var _arg0 *C.GtkFrame // out
	var _cret C.float     // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_label_align(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (f frame) LabelWidget() Widget {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_label_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (f frame) SetChild(child Widget) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_frame_set_child(_arg0, _arg1)
}

func (f frame) SetLabel(label string) {
	var _arg0 *C.GtkFrame // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_frame_set_label(_arg0, _arg1)
}

func (f frame) SetLabelAlign(xalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = C.float(xalign)

	C.gtk_frame_set_label_align(_arg0, _arg1)
}

func (f frame) SetLabelWidget(labelWidget Widget) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_frame_set_label_widget(_arg0, _arg1)
}
