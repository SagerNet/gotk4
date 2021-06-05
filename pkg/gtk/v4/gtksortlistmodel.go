// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_sort_list_model_get_type()), F: marshalSortListModel},
	})
}

// SortListModel is a list model that takes a list model and sorts its elements
// according to a Sorter.
//
// The model can be set up to do incremental sorting, so that sorting long lists
// doesn't block the UI. See gtk_sort_list_model_set_incremental() for details.
//
// SortListModel is a generic model and because of that it cannot take advantage
// of any external knowledge when sorting. If you run into performance issues
// with SortListModel, it is strongly recommended that you write your own
// sorting list model.
type SortListModel interface {
	gextras.Objector
	gio.ListModel

	// Incremental returns whether incremental sorting was enabled via
	// gtk_sort_list_model_set_incremental().
	Incremental() bool
	// Model gets the model currently sorted or nil if none.
	Model() gio.ListModel
	// Pending estimates progress of an ongoing sorting operation
	//
	// The estimate is the number of items that would still need to be sorted to
	// finish the sorting operation if this was a linear algorithm. So this
	// number is not related to how many items are already correctly sorted.
	//
	// If you want to estimate the progress, you can use code like this:
	//
	//    pending = gtk_sort_list_model_get_pending (self);
	//    model = gtk_sort_list_model_get_model (self);
	//    progress = 1.0 - pending / (double) MAX (1, g_list_model_get_n_items (model));
	//
	// If no sort operation is ongoing - in particular when
	// SortListModel:incremental is false - this function returns 0.
	Pending() uint
	// Sorter gets the sorter that is used to sort @self.
	Sorter() Sorter
	// SetIncremental sets the sort model to do an incremental sort.
	//
	// When incremental sorting is enabled, the sortlistmodel will not do a
	// complete sort immediately, but will instead queue an idle handler that
	// incrementally sorts the items towards their correct position. This of
	// course means that items do not instantly appear in the right place. It
	// also means that the total sorting time is a lot slower.
	//
	// When your filter blocks the UI while sorting, you might consider turning
	// this on. Depending on your model and sorters, this may become interesting
	// around 10,000 to 100,000 items.
	//
	// By default, incremental sorting is disabled.
	//
	// See gtk_sort_list_model_get_pending() for progress information about an
	// ongoing incremental sorting operation.
	SetIncremental(incremental bool)
	// SetModel sets the model to be sorted. The @model's item type must conform
	// to the item type of @self.
	SetModel(model gio.ListModel)
	// SetSorter sets a new sorter on @self.
	SetSorter(sorter Sorter)
}

// sortListModel implements the SortListModel interface.
type sortListModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SortListModel = (*sortListModel)(nil)

// WrapSortListModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSortListModel(obj *externglib.Object) SortListModel {
	return SortListModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSortListModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSortListModel(obj), nil
}

// NewSortListModel constructs a class SortListModel.
func NewSortListModel(model gio.ListModel, sorter Sorter) SortListModel {
	var arg1 *C.GListModel
	var arg2 *C.GtkSorter

	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	var cret C.GtkSortListModel
	var goret1 SortListModel

	cret = C.gtk_sort_list_model_new(model, sorter)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SortListModel)

	return goret1
}

// Incremental returns whether incremental sorting was enabled via
// gtk_sort_list_model_set_incremental().
func (s sortListModel) Incremental() bool {
	var arg0 *C.GtkSortListModel

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_sort_list_model_get_incremental(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Model gets the model currently sorted or nil if none.
func (s sortListModel) Model() gio.ListModel {
	var arg0 *C.GtkSortListModel

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var cret *C.GListModel
	var goret1 gio.ListModel

	cret = C.gtk_sort_list_model_get_model(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gio.ListModel)

	return goret1
}

// Pending estimates progress of an ongoing sorting operation
//
// The estimate is the number of items that would still need to be sorted to
// finish the sorting operation if this was a linear algorithm. So this
// number is not related to how many items are already correctly sorted.
//
// If you want to estimate the progress, you can use code like this:
//
//    pending = gtk_sort_list_model_get_pending (self);
//    model = gtk_sort_list_model_get_model (self);
//    progress = 1.0 - pending / (double) MAX (1, g_list_model_get_n_items (model));
//
// If no sort operation is ongoing - in particular when
// SortListModel:incremental is false - this function returns 0.
func (s sortListModel) Pending() uint {
	var arg0 *C.GtkSortListModel

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var cret C.guint
	var goret1 uint

	cret = C.gtk_sort_list_model_get_pending(arg0)

	goret1 = C.guint(cret)

	return goret1
}

// Sorter gets the sorter that is used to sort @self.
func (s sortListModel) Sorter() Sorter {
	var arg0 *C.GtkSortListModel

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))

	var cret *C.GtkSorter
	var goret1 Sorter

	cret = C.gtk_sort_list_model_get_sorter(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Sorter)

	return goret1
}

// SetIncremental sets the sort model to do an incremental sort.
//
// When incremental sorting is enabled, the sortlistmodel will not do a
// complete sort immediately, but will instead queue an idle handler that
// incrementally sorts the items towards their correct position. This of
// course means that items do not instantly appear in the right place. It
// also means that the total sorting time is a lot slower.
//
// When your filter blocks the UI while sorting, you might consider turning
// this on. Depending on your model and sorters, this may become interesting
// around 10,000 to 100,000 items.
//
// By default, incremental sorting is disabled.
//
// See gtk_sort_list_model_get_pending() for progress information about an
// ongoing incremental sorting operation.
func (s sortListModel) SetIncremental(incremental bool) {
	var arg0 *C.GtkSortListModel
	var arg1 C.gboolean

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	if incremental {
		arg1 = C.gboolean(1)
	}

	C.gtk_sort_list_model_set_incremental(arg0, incremental)
}

// SetModel sets the model to be sorted. The @model's item type must conform
// to the item type of @self.
func (s sortListModel) SetModel(model gio.ListModel) {
	var arg0 *C.GtkSortListModel
	var arg1 *C.GListModel

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_sort_list_model_set_model(arg0, model)
}

// SetSorter sets a new sorter on @self.
func (s sortListModel) SetSorter(sorter Sorter) {
	var arg0 *C.GtkSortListModel
	var arg1 *C.GtkSorter

	arg0 = (*C.GtkSortListModel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_sort_list_model_set_sorter(arg0, sorter)
}
