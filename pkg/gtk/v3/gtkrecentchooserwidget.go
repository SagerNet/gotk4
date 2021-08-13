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
		{T: externglib.Type(C.gtk_recent_chooser_widget_get_type()), F: marshalRecentChooserWidgetter},
	})
}

// RecentChooserWidget is a widget suitable for selecting recently used files.
// It is the main building block of a RecentChooserDialog. Most applications
// will only need to use the latter; you can use RecentChooserWidget as part of
// a larger window if you have special needs.
//
// Note that RecentChooserWidget does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserWidget struct {
	Box

	RecentChooser
	*externglib.Object
}

func wrapRecentChooserWidget(obj *externglib.Object) *RecentChooserWidget {
	return &RecentChooserWidget{
		Box: Box{
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
		},
		RecentChooser: RecentChooser{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalRecentChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRecentChooserWidget(obj), nil
}

// NewRecentChooserWidget creates a new RecentChooserWidget object. This is an
// embeddable widget used to access the recently used resources list.
func NewRecentChooserWidget() *RecentChooserWidget {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_recent_chooser_widget_new()

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}

// NewRecentChooserWidgetForManager creates a new RecentChooserWidget with a
// specified recent manager.
//
// This is useful if you have implemented your own recent manager, or if you
// have a customized instance of a RecentManager object.
func NewRecentChooserWidgetForManager(manager *RecentManager) *RecentChooserWidget {
	var _arg1 *C.GtkRecentManager // out
	var _cret *C.GtkWidget        // in

	_arg1 = (*C.GtkRecentManager)(unsafe.Pointer(manager.Native()))

	_cret = C.gtk_recent_chooser_widget_new_for_manager(_arg1)
	runtime.KeepAlive(manager)

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}

func (*RecentChooserWidget) privateRecentChooserWidget() {}
