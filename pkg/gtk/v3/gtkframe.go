// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFramer},
	})
}

// FrameOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FrameOverrider interface {
	ComputeChildAllocation(allocation *Allocation)
}

// Frame: frame widget is a bin that surrounds its child with a decorative frame
// and an optional label. If present, the label is drawn in a gap in the top
// side of the frame. The position of the label can be controlled with
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
type Frame struct {
	Bin
}

func wrapFrame(obj *externglib.Object) *Frame {
	return &Frame{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
	}
}

func marshalFramer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFrame(obj), nil
}

// NewFrame creates a new Frame, with optional label label. If label is NULL,
// the label is omitted.
func NewFrame(label string) *Frame {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_frame_new(_arg1)
	runtime.KeepAlive(label)

	var _frame *Frame // out

	_frame = wrapFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

// Label: if the frame’s label widget is a Label, returns the text in the label
// widget. (The frame will have a Label for the label widget if a non-NULL
// argument was passed to gtk_frame_new().)
func (frame *Frame) Label() string {
	var _arg0 *C.GtkFrame // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_label(_arg0)
	runtime.KeepAlive(frame)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelAlign retrieves the X and Y alignment of the frame’s label. See
// gtk_frame_set_label_align().
func (frame *Frame) LabelAlign() (xalign float32, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // in
	var _arg2 C.gfloat    // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	C.gtk_frame_get_label_align(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(frame)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_frame_set_label_widget().
func (frame *Frame) LabelWidget() Widgetter {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_label_widget(_arg0)
	runtime.KeepAlive(frame)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// ShadowType retrieves the shadow type of the frame. See
// gtk_frame_set_shadow_type().
func (frame *Frame) ShadowType() ShadowType {
	var _arg0 *C.GtkFrame     // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_shadow_type(_arg0)
	runtime.KeepAlive(frame)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// SetLabel removes the current Frame:label-widget. If label is not NULL,
// creates a new Label with that text and adds it as the Frame:label-widget.
func (frame *Frame) SetLabel(label string) {
	var _arg0 *C.GtkFrame // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_frame_set_label(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(label)
}

// SetLabelAlign sets the alignment of the frame widget’s label. The default
// values for a newly created frame are 0.0 and 0.5.
func (frame *Frame) SetLabelAlign(xalign float32, yalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.gfloat    // out
	var _arg2 C.gfloat    // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_frame_set_label_align(_arg0, _arg1, _arg2)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
}

// SetLabelWidget sets the Frame:label-widget for the frame. This is the widget
// that will appear embedded in the top edge of the frame as a title.
func (frame *Frame) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	if labelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))
	}

	C.gtk_frame_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(labelWidget)
}

// SetShadowType sets the Frame:shadow-type for frame, i.e. whether it is drawn
// without (GTK_SHADOW_NONE) or with (other values) a visible border. Values
// other than GTK_SHADOW_NONE are treated identically by GtkFrame. The chosen
// type is applied by removing or adding the .flat class to the CSS node named
// border.
func (frame *Frame) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkFrame     // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_frame_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(typ)
}
