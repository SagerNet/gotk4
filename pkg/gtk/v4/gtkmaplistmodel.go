// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gpointer _gotk4_gtk4_MapListModelMapFunc(gpointer, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_map_list_model_get_type()), F: marshalMapListModeller},
	})
}

// MapListModelMapFunc: user function that is called to map an item of the
// original model to an item expected by the map model.
//
// The returned items must conform to the item type of the model they are used
// with.
type MapListModelMapFunc func(item *externglib.Object) (object *externglib.Object)

//export _gotk4_gtk4_MapListModelMapFunc
func _gotk4_gtk4_MapListModelMapFunc(arg0 C.gpointer, arg1 C.gpointer) (cret C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item *externglib.Object // out

	item = externglib.AssumeOwnership(unsafe.Pointer(arg0))

	fn := v.(MapListModelMapFunc)
	object := fn(item)

	cret = C.gpointer(unsafe.Pointer(object.Native()))

	return cret
}

// MapListModel: GtkMapListModel maps the items in a list model to different
// items.
//
// GtkMapListModel uses a gtk.MapListModelMapFunc.
//
// Example: Create a list of GtkEventControllers
//
//    static gpointer
//    map_to_controllers (gpointer widget,
//                        gpointer data)
//    {
//      gpointer result = gtk_widget_observe_controllers (widget);
//      g_object_unref (widget);
//      return result;
//    }
//
//    widgets = gtk_widget_observe_children (widget);
//
//    controllers = gtk_map_list_model_new (G_TYPE_LIST_MODEL,
//                                          widgets,
//                                          map_to_controllers,
//                                          NULL, NULL);
//
//    model = gtk_flatten_list_model_new (GTK_TYPE_EVENT_CONTROLLER,
//                                        controllers);
//
//
// GtkMapListModel will attempt to discard the mapped objects as soon as they
// are no longer needed and recreate them if necessary.
type MapListModel struct {
	*externglib.Object

	gio.ListModel
}

func WrapMapListModel(obj *externglib.Object) *MapListModel {
	return &MapListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalMapListModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMapListModel(obj), nil
}

// NewMapListModel creates a new GtkMapListModel for the given arguments.
func NewMapListModel(model gio.ListModeller, mapFunc MapListModelMapFunc) *MapListModel {
	var _arg1 *C.GListModel            // out
	var _arg2 C.GtkMapListModelMapFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _cret *C.GtkMapListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}
	if mapFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk4_MapListModelMapFunc)
		_arg3 = C.gpointer(gbox.Assign(mapFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gtk_map_list_model_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(model)
	runtime.KeepAlive(mapFunc)

	var _mapListModel *MapListModel // out

	_mapListModel = WrapMapListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mapListModel
}

// Model gets the model that is currently being mapped or NULL if none.
func (self *MapListModel) Model() gio.ListModeller {
	var _arg0 *C.GtkMapListModel // out
	var _cret *C.GListModel      // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_map_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	if _cret != nil {
		_listModel = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)
	}

	return _listModel
}

// HasMap checks if a map function is currently set on self.
func (self *MapListModel) HasMap() bool {
	var _arg0 *C.GtkMapListModel // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_map_list_model_has_map(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMapFunc sets the function used to map items.
//
// The function will be called whenever an item needs to be mapped and must
// return the item to use for the given input item.
//
// Note that GtkMapListModel may call this function multiple times on the same
// item, because it may delete items it doesn't need anymore.
//
// GTK makes no effort to ensure that map_func conforms to the item type of
// self. It assumes that the caller knows what they are doing and the map
// function returns items of the appropriate type.
func (self *MapListModel) SetMapFunc(mapFunc MapListModelMapFunc) {
	var _arg0 *C.GtkMapListModel       // out
	var _arg1 C.GtkMapListModelMapFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))
	if mapFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_MapListModelMapFunc)
		_arg2 = C.gpointer(gbox.Assign(mapFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_map_list_model_set_map_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mapFunc)
}

// SetModel sets the model to be mapped.
//
// GTK makes no effort to ensure that model conforms to the item type expected
// by the map function. It assumes that the caller knows what they are doing and
// have set up an appropriate map function.
func (self *MapListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkMapListModel // out
	var _arg1 *C.GListModel      // out

	_arg0 = (*C.GtkMapListModel)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_map_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}
