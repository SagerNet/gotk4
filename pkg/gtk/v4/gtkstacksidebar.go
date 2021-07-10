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
		{T: externglib.Type(C.gtk_stack_sidebar_get_type()), F: marshalStackSidebarrer},
	})
}

// StackSidebarrer describes StackSidebar's methods.
type StackSidebarrer interface {
	gextras.Objector

	Stack() *Stack
	SetStack(stack Stacker)
}

// StackSidebar: `GtkStackSidebar` uses a sidebar to switch between `GtkStack`
// pages.
//
// In order to use a `GtkStackSidebar`, you simply use a `GtkStack` to organize
// your UI flow, and add the sidebar to your sidebar area. You can use
// [method@Gtk.StackSidebar.set_stack] to connect the `GtkStackSidebar` to the
// `GtkStack`.
//
//
// CSS nodes
//
// `GtkStackSidebar` has a single CSS node with name stacksidebar and style
// class .sidebar.
//
// When circumstances require it, `GtkStackSidebar` adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSidebar struct {
	*externglib.Object
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ StackSidebarrer = (*StackSidebar)(nil)

func wrapStackSidebarrer(obj *externglib.Object) StackSidebarrer {
	return &StackSidebar{
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
	}
}

func marshalStackSidebarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStackSidebarrer(obj), nil
}

// NewStackSidebar creates a new `GtkStackSidebar`.
func NewStackSidebar() *StackSidebar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_sidebar_new()

	var _stackSidebar *StackSidebar // out

	_stackSidebar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*StackSidebar)

	return _stackSidebar
}

// Stack retrieves the stack.
func (self *StackSidebar) Stack() *Stack {
	var _arg0 *C.GtkStackSidebar // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_stack_sidebar_get_stack(_arg0)

	var _stack *Stack // out

	_stack = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Stack)

	return _stack
}

// SetStack: set the `GtkStack` associated with this `GtkStackSidebar`.
//
// The sidebar widget will automatically update according to the order and items
// within the given `GtkStack`.
func (self *StackSidebar) SetStack(stack Stacker) {
	var _arg0 *C.GtkStackSidebar // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.GtkStackSidebar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_sidebar_set_stack(_arg0, _arg1)
}
