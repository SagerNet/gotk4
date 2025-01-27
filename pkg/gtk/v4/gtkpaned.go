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
		{T: externglib.Type(C.gtk_paned_get_type()), F: marshalPanedder},
	})
}

// Paned: GtkPaned has two panes, arranged either horizontally or vertically.
//
// !An example GtkPaned (panes.png)
//
// The division between the two panes is adjustable by the user by dragging a
// handle.
//
// Child widgets are added to the panes of the widget with
// gtk.Paned.SetStartChild() and gtk.Paned.SetEndChild(). The division between
// the two children is set by default from the size requests of the children,
// but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a gtk.Frame so that the gutter appears as a ridge. No separator is drawn if
// one of the children is missing.
//
// Each child has two options that can be set, resize and shrink. If resize is
// true, then when the GtkPaned is resized, that child will expand or shrink
// along with the paned widget. If shrink is true, then that child can be made
// smaller than its requisition by the user. Setting shrink to FALSE allows the
// application to set a minimum size. If resize is false for both children, then
// this is treated as if resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling gtk.Paned.SetPosition().
//
// CSS nodes
//
//    paned
//    ├── <child>
//    ├── separator[.wide]
//    ╰── <child>
//
//
// GtkPaned has a main CSS node with name paned, and a subnode for the separator
// with name separator. The subnode gets a .wide style class when the paned is
// supposed to be wide.
//
// In horizontal orientation, the nodes are arranged based on the text
// direction, so in left-to-right mode, :first-child will select the leftmost
// child, while it will select the rightmost child in RTL layouts.
//
// Creating a paned widget with minimum sizes.
//
//    GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
//    GtkWidget *frame1 = gtk_frame_new (NULL);
//    GtkWidget *frame2 = gtk_frame_new (NULL);
//
//    gtk_widget_set_size_request (hpaned, 200, -1);
//
//    gtk_paned_set_start_child (GTK_PANED (hpaned), frame1);
//    gtk_paned_set_start_child_resize (GTK_PANED (hpaned), TRUE);
//    gtk_paned_set_start_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame1, 50, -1);
//
//    gtk_paned_set_end_child (GTK_PANED (hpaned), frame2);
//    gtk_paned_set_end_child_resize (GTK_PANED (hpaned), FALSE);
//    gtk_paned_set_end_child_shrink (GTK_PANED (hpaned), FALSE);
//    gtk_widget_set_size_request (frame2, 50, -1);
type Paned struct {
	Widget

	Orientable
	*externglib.Object
}

func WrapPaned(obj *externglib.Object) *Paned {
	return &Paned{
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

func marshalPanedder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPaned(obj), nil
}

// NewPaned creates a new GtkPaned widget.
func NewPaned(orientation Orientation) *Paned {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_paned_new(_arg1)
	runtime.KeepAlive(orientation)

	var _paned *Paned // out

	_paned = WrapPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _paned
}

// EndChild retrieves the end child of the given GtkPaned.
//
// See also: GtkPaned:end-child
func (paned *Paned) EndChild() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Position obtains the position of the divider between the two panes.
func (paned *Paned) Position() int {
	var _arg0 *C.GtkPaned // out
	var _cret C.int       // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_position(_arg0)
	runtime.KeepAlive(paned)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResizeEndChild returns whether the end child can be resized.
func (paned *Paned) ResizeEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_resize_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ResizeStartChild returns whether the start child can be resized.
func (paned *Paned) ResizeStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_resize_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShrinkEndChild returns whether the end child can be shrunk.
func (paned *Paned) ShrinkEndChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_shrink_end_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShrinkStartChild returns whether the start child can be shrunk.
func (paned *Paned) ShrinkStartChild() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_shrink_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartChild retrieves the start child of the given GtkPaned.
//
// See also: GtkPaned:start-child
func (paned *Paned) StartChild() Widgetter {
	var _arg0 *C.GtkPaned  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_start_child(_arg0)
	runtime.KeepAlive(paned)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// WideHandle gets whether the separator should be wide.
func (paned *Paned) WideHandle() bool {
	var _arg0 *C.GtkPaned // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))

	_cret = C.gtk_paned_get_wide_handle(_arg0)
	runtime.KeepAlive(paned)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEndChild sets the end child of paned to child.
func (paned *Paned) SetEndChild(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// SetPosition sets the position of the divider between the two panes.
func (paned *Paned) SetPosition(position int) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = C.int(position)

	C.gtk_paned_set_position(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(position)
}

// SetResizeEndChild sets the GtkPaned:resize-end-child property
func (paned *Paned) SetResizeEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetResizeStartChild sets the GtkPaned:resize-start-child property
func (paned *Paned) SetResizeStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_resize_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetShrinkEndChild sets the GtkPaned:shrink-end-child property
func (paned *Paned) SetShrinkEndChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_end_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetShrinkStartChild sets the GtkPaned:shrink-start-child property
func (paned *Paned) SetShrinkStartChild(resize bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if resize {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_shrink_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(resize)
}

// SetStartChild sets the start child of paned to child.
func (paned *Paned) SetStartChild(child Widgetter) {
	var _arg0 *C.GtkPaned  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_set_start_child(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(child)
}

// SetWideHandle sets whether the separator should be wide.
func (paned *Paned) SetWideHandle(wide bool) {
	var _arg0 *C.GtkPaned // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkPaned)(unsafe.Pointer(paned.Native()))
	if wide {
		_arg1 = C.TRUE
	}

	C.gtk_paned_set_wide_handle(_arg0, _arg1)
	runtime.KeepAlive(paned)
	runtime.KeepAlive(wide)
}
