// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFrame},
	})
}

// Frame: the frame widget is a bin that surrounds its child with a decorative
// frame and an optional label. If present, the label is drawn in a gap in the
// top side of the frame. The position of the label can be controlled with
// gtk_frame_set_label_align().
//
//
// GtkFrame as GtkBuildable
//
// The GtkFrame implementation of the GtkBuildable interface supports placing a
// child in the label position by specifying “label” as the “type” attribute of
// a <child> element. A normal content child can be specified without specifying
// a <child> type attribute.
//
// An example of a UI definition fragment with GtkFrame:
//
//    <object class="GtkFrame">
//      <child type="label">
//        <object class="GtkLabel" id="frame-label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame-content"/>
//      </child>
//    </object>
//
// CSS nodes
//
//    frame
//    ├── border[.flat]
//    ├── <label widget>
//    ╰── <child>
//
// GtkFrame has a main CSS node named “frame” and a subnode named “border”. The
// “border” node is used to draw the visible border. You can set the appearance
// of the border using CSS properties like “border-style” on the “border” node.
//
// The border node can be given the style class “.flat”, which is used by themes
// to disable drawing of the border. To do this from code, call
// gtk_frame_set_shadow_type() with GTK_SHADOW_NONE to add the “.flat” class or
// any other shadow type to remove it.
type Frame interface {
	Bin

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// Label: if the frame’s label widget is a Label, returns the text in the
	// label widget. (The frame will have a Label for the label widget if a
	// non-nil argument was passed to gtk_frame_new().)
	Label() string
	// LabelAlign retrieves the X and Y alignment of the frame’s label. See
	// gtk_frame_set_label_align().
	LabelAlign() (xalign float32, yalign float32)
	// LabelWidget retrieves the label widget for the frame. See
	// gtk_frame_set_label_widget().
	LabelWidget() Widget
	// ShadowType retrieves the shadow type of the frame. See
	// gtk_frame_set_shadow_type().
	ShadowType() ShadowType
	// SetLabel removes the current Frame:label-widget. If @label is not nil,
	// creates a new Label with that text and adds it as the Frame:label-widget.
	SetLabel(label string)
	// SetLabelAlign sets the alignment of the frame widget’s label. The default
	// values for a newly created frame are 0.0 and 0.5.
	SetLabelAlign(xalign float32, yalign float32)
	// SetLabelWidget sets the Frame:label-widget for the frame. This is the
	// widget that will appear embedded in the top edge of the frame as a title.
	SetLabelWidget(labelWidget Widget)
	// SetShadowType sets the Frame:shadow-type for @frame, i.e. whether it is
	// drawn without (GTK_SHADOW_NONE) or with (other values) a visible border.
	// Values other than GTK_SHADOW_NONE are treated identically by GtkFrame.
	// The chosen type is applied by removing or adding the .flat class to the
	// CSS node named border.
	SetShadowType(typ ShadowType)
}

// frame implements the Frame class.
type frame struct {
	Bin
}

// WrapFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapFrame(obj *externglib.Object) Frame {
	return frame{
		Bin: WrapBin(obj),
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFrame(obj), nil
}

// NewFrame creates a new Frame, with optional label @label. If @label is nil,
// the label is omitted.
func NewFrame(label string) Frame {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_frame_new(_arg1)

	var _frame Frame // out

	_frame = WrapFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

func (f frame) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(f))
}

func (f frame) Label() string {
	var _arg0 *C.GtkFrame // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (f frame) LabelAlign() (xalign float32, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // in
	var _arg2 C.gfloat    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	C.gtk_frame_get_label_align(_arg0, &_arg1, &_arg2)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
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

func (f frame) ShadowType() ShadowType {
	var _arg0 *C.GtkFrame     // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_frame_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

func (f frame) SetLabel(label string) {
	var _arg0 *C.GtkFrame // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_frame_set_label(_arg0, _arg1)
}

func (f frame) SetLabelAlign(xalign float32, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // out
	var _arg2 C.gfloat    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_frame_set_label_align(_arg0, _arg1, _arg2)
}

func (f frame) SetLabelWidget(labelWidget Widget) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_frame_set_label_widget(_arg0, _arg1)
}

func (f frame) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkFrame     // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_frame_set_shadow_type(_arg0, _arg1)
}
