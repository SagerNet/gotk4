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
		{T: externglib.Type(C.gtk_accel_map_get_type()), F: marshalAccelMapper},
	})
}

// AccelMapper describes AccelMap's methods.
type AccelMapper interface {
	gextras.Objector

	privateAccelMap()
}

// AccelMap: accelerator maps are used to define runtime configurable
// accelerators. Functions for manipulating them are are usually used by higher
// level convenience mechanisms like UIManager and are thus considered
// “low-level”. You’ll want to use them if you’re manually creating menus that
// should have user-configurable accelerators.
//
// An accelerator is uniquely defined by: - accelerator path - accelerator key -
// accelerator modifiers
//
// The accelerator path must consist of
// “<WINDOWTYPE>/Category1/Category2/.../Action”, where WINDOWTYPE should be a
// unique application-specific identifier that corresponds to the kind of window
// the accelerator is being used in, e.g. “Gimp-Image”, “Abiword-Document” or
// “Gnumeric-Settings”. The “Category1/.../Action” portion is most appropriately
// chosen by the action the accelerator triggers, i.e. for accelerators on menu
// items, choose the item’s menu path, e.g. “File/Save As”, “Image/View/Zoom” or
// “Edit/Select All”. So a full valid accelerator path may look like:
// “<Gimp-Toolbox>/File/Dialogs/Tool Options...”.
//
// All accelerators are stored inside one global AccelMap that can be obtained
// using gtk_accel_map_get(). See [Monitoring changes][monitoring-changes] for
// additional details.
//
//
// Manipulating accelerators
//
// New accelerators can be added using gtk_accel_map_add_entry(). To search for
// specific accelerator, use gtk_accel_map_lookup_entry(). Modifications of
// existing accelerators should be done using gtk_accel_map_change_entry().
//
// In order to avoid having some accelerators changed, they can be locked using
// gtk_accel_map_lock_path(). Unlocking is done using
// gtk_accel_map_unlock_path().
//
//
// Saving and loading accelerator maps
//
// Accelerator maps can be saved to and loaded from some external resource. For
// simple saving and loading from file, gtk_accel_map_save() and
// gtk_accel_map_load() are provided. Saving and loading can also be done by
// providing file descriptor to gtk_accel_map_save_fd() and
// gtk_accel_map_load_fd().
//
//
// Monitoring changes
//
// AccelMap object is only useful for monitoring changes of accelerators. By
// connecting to AccelMap::changed signal, one can monitor changes of all
// accelerators. It is also possible to monitor only single accelerator path by
// using it as a detail of the AccelMap::changed signal.
type AccelMap struct {
	*externglib.Object
}

var _ AccelMapper = (*AccelMap)(nil)

func wrapAccelMapper(obj *externglib.Object) AccelMapper {
	return &AccelMap{
		Object: obj,
	}
}

func marshalAccelMapper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAccelMapper(obj), nil
}

func (*AccelMap) privateAccelMap() {}
