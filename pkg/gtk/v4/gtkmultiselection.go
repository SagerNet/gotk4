// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_multi_selection_get_type()), F: marshalMultiSelection},
	})
}

// MultiSelection: `GtkMultiSelection` is a `GtkSelectionModel` that allows
// selecting multiple elements.
type MultiSelection interface {
	SelectionModel

	// Model:
	Model() gio.ListModel
	// SetModelMultiSelection:
	SetModelMultiSelection(model gio.ListModel)
}

// multiSelection implements the MultiSelection class.
type multiSelection struct {
	gextras.Objector
}

// WrapMultiSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapMultiSelection(obj *externglib.Object) MultiSelection {
	return multiSelection{
		Objector: obj,
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiSelection(obj), nil
}

// NewMultiSelection:
func NewMultiSelection(model gio.ListModel) MultiSelection {
	var _arg1 *C.GListModel        // out
	var _cret *C.GtkMultiSelection // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_multi_selection_new(_arg1)

	var _multiSelection MultiSelection // out

	_multiSelection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MultiSelection)

	return _multiSelection
}

func (s multiSelection) Model() gio.ListModel {
	var _arg0 *C.GtkMultiSelection // out
	var _cret *C.GListModel        // in

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_multi_selection_get_model(_arg0)

	var _listModel gio.ListModel // out

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.ListModel)

	return _listModel
}

func (s multiSelection) SetModelMultiSelection(model gio.ListModel) {
	var _arg0 *C.GtkMultiSelection // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_multi_selection_set_model(_arg0, _arg1)
}

func (m multiSelection) Selection() *Bitset {
	return WrapSelectionModel(gextras.InternObject(m)).Selection()
}

func (m multiSelection) SelectionInRange(position uint, nItems uint) *Bitset {
	return WrapSelectionModel(gextras.InternObject(m)).SelectionInRange(position, nItems)
}

func (m multiSelection) IsSelected(position uint) bool {
	return WrapSelectionModel(gextras.InternObject(m)).IsSelected(position)
}

func (m multiSelection) SelectAll() bool {
	return WrapSelectionModel(gextras.InternObject(m)).SelectAll()
}

func (m multiSelection) SelectItem(position uint, unselectRest bool) bool {
	return WrapSelectionModel(gextras.InternObject(m)).SelectItem(position, unselectRest)
}

func (m multiSelection) SelectRange(position uint, nItems uint, unselectRest bool) bool {
	return WrapSelectionModel(gextras.InternObject(m)).SelectRange(position, nItems, unselectRest)
}

func (m multiSelection) SelectionChanged(position uint, nItems uint) {
	WrapSelectionModel(gextras.InternObject(m)).SelectionChanged(position, nItems)
}

func (m multiSelection) SetSelection(selected *Bitset, mask *Bitset) bool {
	return WrapSelectionModel(gextras.InternObject(m)).SetSelection(selected, mask)
}

func (m multiSelection) UnselectAll() bool {
	return WrapSelectionModel(gextras.InternObject(m)).UnselectAll()
}

func (m multiSelection) UnselectItem(position uint) bool {
	return WrapSelectionModel(gextras.InternObject(m)).UnselectItem(position)
}

func (m multiSelection) UnselectRange(position uint, nItems uint) bool {
	return WrapSelectionModel(gextras.InternObject(m)).UnselectRange(position, nItems)
}

func (l multiSelection) ItemType() externglib.Type {
	return gio.WrapListModel(gextras.InternObject(l)).ItemType()
}

func (l multiSelection) NItems() uint {
	return gio.WrapListModel(gextras.InternObject(l)).NItems()
}

func (l multiSelection) Object(position uint) gextras.Objector {
	return gio.WrapListModel(gextras.InternObject(l)).Object(position)
}

func (l multiSelection) ItemsChanged(position uint, removed uint, added uint) {
	gio.WrapListModel(gextras.InternObject(l)).ItemsChanged(position, removed, added)
}
