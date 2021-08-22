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
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBoxer},
	})
}

// Box widget arranges child widgets into a single row or column, depending upon
// the value of its Orientable:orientation property. Within the other dimension,
// all children are allocated the same size. Of course, the Widget:halign and
// Widget:valign properties can be used on the children to influence their
// allocation.
//
// GtkBox uses a notion of packing. Packing refers to adding widgets with
// reference to a particular position in a Container. For a GtkBox, there are
// two reference positions: the start and the end of the box. For a vertical
// Box, the start is defined as the top of the box and the end is defined as the
// bottom. For a horizontal Box the start is defined as the left side and the
// end is defined as the right side.
//
// Use repeated calls to gtk_box_pack_start() to pack widgets into a GtkBox from
// start to end. Use gtk_box_pack_end() to add widgets from end to start. You
// may intersperse these calls and add widgets from both ends of the same
// GtkBox.
//
// Because GtkBox is a Container, you may also use gtk_container_add() to insert
// widgets into the box, and they will be packed with the default values for
// expand and fill child properties. Use gtk_container_remove() to remove
// widgets from the GtkBox.
//
// Use gtk_box_set_homogeneous() to specify whether or not all children of the
// GtkBox are forced to get the same amount of space.
//
// Use gtk_box_set_spacing() to determine how much space will be minimally
// placed between all children in the GtkBox. Note that spacing is added between
// the children, while padding added by gtk_box_pack_start() or
// gtk_box_pack_end() is added on either side of the widget it belongs to.
//
// Use gtk_box_reorder_child() to move a GtkBox child to a different place in
// the box.
//
// Use gtk_box_set_child_packing() to reset the expand, fill and padding child
// properties. Use gtk_box_query_child_packing() to query these fields.
//
//
// CSS nodes
//
// GtkBox uses a single CSS node with name box.
//
// In horizontal orientation, the nodes of the children are always arranged from
// left to right. So :first-child will always select the leftmost child,
// regardless of text direction.
type Box struct {
	Container

	Orientable
	*externglib.Object
}

func WrapBox(obj *externglib.Object) *Box {
	return &Box{
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
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalBoxer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBox(obj), nil
}

// NewBox creates a new Box.
func NewBox(orientation Orientation, spacing int) *Box {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.gint           // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.gint(spacing)

	_cret = C.gtk_box_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(spacing)

	var _box *Box // out

	_box = WrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// BaselinePosition gets the value set by gtk_box_set_baseline_position().
func (box *Box) BaselinePosition() BaselinePosition {
	var _arg0 *C.GtkBox             // out
	var _cret C.GtkBaselinePosition // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_baseline_position(_arg0)
	runtime.KeepAlive(box)

	var _baselinePosition BaselinePosition // out

	_baselinePosition = BaselinePosition(_cret)

	return _baselinePosition
}

// CenterWidget retrieves the center widget of the box.
func (box *Box) CenterWidget() Widgetter {
	var _arg0 *C.GtkBox    // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_center_widget(_arg0)
	runtime.KeepAlive(box)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Homogeneous returns whether the box is homogeneous (all children are the same
// size). See gtk_box_set_homogeneous().
func (box *Box) Homogeneous() bool {
	var _arg0 *C.GtkBox  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_homogeneous(_arg0)
	runtime.KeepAlive(box)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Spacing gets the value set by gtk_box_set_spacing().
func (box *Box) Spacing() int {
	var _arg0 *C.GtkBox // out
	var _cret C.gint    // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))

	_cret = C.gtk_box_get_spacing(_arg0)
	runtime.KeepAlive(box)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PackEnd adds child to box, packed with reference to the end of box. The child
// is packed after (away from end of) any other child packed with reference to
// the end of box.
func (box *Box) PackEnd(child Widgetter, expand bool, fill bool, padding uint) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out
	var _arg4 C.guint      // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)

	C.gtk_box_pack_end(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(fill)
	runtime.KeepAlive(padding)
}

// PackStart adds child to box, packed with reference to the start of box. The
// child is packed after any other child packed with reference to the start of
// box.
func (box *Box) PackStart(child Widgetter, expand bool, fill bool, padding uint) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gboolean   // out
	var _arg3 C.gboolean   // out
	var _arg4 C.guint      // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)

	C.gtk_box_pack_start(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(fill)
	runtime.KeepAlive(padding)
}

// QueryChildPacking obtains information about how child is packed into box.
func (box *Box) QueryChildPacking(child Widgetter) (expand bool, fill bool, padding uint, packType PackType) {
	var _arg0 *C.GtkBox     // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // in
	var _arg3 C.gboolean    // in
	var _arg4 C.guint       // in
	var _arg5 C.GtkPackType // in

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_query_child_packing(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)

	var _expand bool       // out
	var _fill bool         // out
	var _padding uint      // out
	var _packType PackType // out

	if _arg2 != 0 {
		_expand = true
	}
	if _arg3 != 0 {
		_fill = true
	}
	_padding = uint(_arg4)
	_packType = PackType(_arg5)

	return _expand, _fill, _padding, _packType
}

// ReorderChild moves child to a new position in the list of box children. The
// list contains widgets packed K_PACK_START as well as widgets packed
// K_PACK_END, in the order that these widgets were added to box.
//
// A widget’s position in the box children list determines where the widget is
// packed into box. A child widget at some position in the list will be packed
// just after all other widgets of the same packing type that appear earlier in
// the list.
func (box *Box) ReorderChild(child Widgetter, position int) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_box_reorder_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// SetBaselinePosition sets the baseline position of a box. This affects only
// horizontal boxes with at least one baseline aligned child. If there is more
// vertical space available than requested, and the baseline is not allocated by
// the parent then position is used to allocate the baseline wrt the extra space
// available.
func (box *Box) SetBaselinePosition(position BaselinePosition) {
	var _arg0 *C.GtkBox             // out
	var _arg1 C.GtkBaselinePosition // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.GtkBaselinePosition(position)

	C.gtk_box_set_baseline_position(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(position)
}

// SetCenterWidget sets a center widget; that is a child widget that will be
// centered with respect to the full width of the box, even if the children at
// either side take up different amounts of space.
func (box *Box) SetCenterWidget(widget Widgetter) {
	var _arg0 *C.GtkBox    // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_box_set_center_widget(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(widget)
}

// SetChildPacking sets the way child is packed into box.
func (box *Box) SetChildPacking(child Widgetter, expand bool, fill bool, padding uint, packType PackType) {
	var _arg0 *C.GtkBox     // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out
	var _arg3 C.gboolean    // out
	var _arg4 C.guint       // out
	var _arg5 C.GtkPackType // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if expand {
		_arg2 = C.TRUE
	}
	if fill {
		_arg3 = C.TRUE
	}
	_arg4 = C.guint(padding)
	_arg5 = C.GtkPackType(packType)

	C.gtk_box_set_child_packing(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(box)
	runtime.KeepAlive(child)
	runtime.KeepAlive(expand)
	runtime.KeepAlive(fill)
	runtime.KeepAlive(padding)
	runtime.KeepAlive(packType)
}

// SetHomogeneous sets the Box:homogeneous property of box, controlling whether
// or not all children of box are given equal space in the box.
func (box *Box) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkBox  // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_box_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(homogeneous)
}

// SetSpacing sets the Box:spacing property of box, which is the number of
// pixels to place between children of box.
func (box *Box) SetSpacing(spacing int) {
	var _arg0 *C.GtkBox // out
	var _arg1 C.gint    // out

	_arg0 = (*C.GtkBox)(unsafe.Pointer(box.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_box_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(box)
	runtime.KeepAlive(spacing)
}
