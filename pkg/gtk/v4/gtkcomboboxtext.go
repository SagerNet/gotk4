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
		{T: externglib.Type(C.gtk_combo_box_text_get_type()), F: marshalComboBoxTexter},
	})
}

// ComboBoxText: GtkComboBoxText is a simple variant of GtkComboBox for
// text-only use cases.
//
// !An example GtkComboBoxText (combo-box-text.png)
//
// GtkComboBoxText hides the model-view complexity of GtkComboBox.
//
// To create a GtkComboBoxText, use gtk.ComboBoxText.New or
// gtk.ComboBoxText.NewWithEntry.
//
// You can add items to a GtkComboBoxText with gtk.ComboBoxText.AppendText(),
// gtk.ComboBoxText.InsertText() or gtk.ComboBoxText.PrependText() and remove
// options with gtk.ComboBoxText.Remove().
//
// If the GtkComboBoxText contains an entry (via the gtk.ComboBox:has-entry
// property), its contents can be retrieved using
// gtk.ComboBoxText.GetActiveText().
//
// You should not call gtk.ComboBox.SetModel() or attempt to pack more cells
// into this combo box via its gtk.CellLayout interface.
//
//
// GtkComboBoxText as GtkBuildable
//
// The GtkComboBoxText implementation of the GtkBuildable interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element can specify the “id”
// corresponding to the appended text and also supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying GtkComboBoxText items:
//
//    <object class="GtkComboBoxText">
//      <items>
//        <item translatable="yes" id="factory">Factory</item>
//        <item translatable="yes" id="home">Home</item>
//        <item translatable="yes" id="subway">Subway</item>
//      </items>
//    </object>
//
//
// CSS nodes
//
//    combobox
//    ╰── box.linked
//        ├── entry.combo
//        ├── button.combo
//        ╰── window.popup
//
//
// GtkComboBoxText has a single CSS node with name combobox. It adds the style
// class .combo to the main CSS nodes of its entry and button children, and the
// .linked class to the node of its internal box.
type ComboBoxText struct {
	ComboBox
}

func WrapComboBoxText(obj *externglib.Object) *ComboBoxText {
	return &ComboBoxText{
		ComboBox: ComboBox{
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
			CellEditable: CellEditable{
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
			},
			CellLayout: CellLayout{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalComboBoxTexter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapComboBoxText(obj), nil
}

// NewComboBoxText creates a new GtkComboBoxText.
func NewComboBoxText() *ComboBoxText {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new()

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = WrapComboBoxText(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBoxText
}

// NewComboBoxTextWithEntry creates a new GtkComboBoxText with an entry.
func NewComboBoxTextWithEntry() *ComboBoxText {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new_with_entry()

	var _comboBoxText *ComboBoxText // out

	_comboBoxText = WrapComboBoxText(externglib.Take(unsafe.Pointer(_cret)))

	return _comboBoxText
}

// Append appends text to the list of strings stored in combo_box.
//
// If id is non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk.ComboBoxText.Insert() with a position of -1.
func (comboBox *ComboBoxText) Append(id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	if id != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_append(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// AppendText appends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk.ComboBoxText.InsertText() with a position of
// -1.
func (comboBox *ComboBoxText) AppendText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_append_text(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// ActiveText returns the currently active string in combo_box.
//
// If no row is currently selected, NULL is returned. If combo_box contains an
// entry, this function will return its contents (which will not necessarily be
// an item from the list).
func (comboBox *ComboBoxText) ActiveText() string {
	var _arg0 *C.GtkComboBoxText // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))

	_cret = C.gtk_combo_box_text_get_active_text(_arg0)
	runtime.KeepAlive(comboBox)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Insert inserts text at position in the list of strings stored in combo_box.
//
// If id is non-NULL then it is used as the ID of the row. See
// gtk.ComboBox:id-column.
//
// If position is negative then text is appended.
func (comboBox *ComboBoxText) Insert(position int, id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out
	var _arg2 *C.char            // out
	var _arg3 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(position)
	if id != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_combo_box_text_insert(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// InsertText inserts text at position in the list of strings stored in
// combo_box.
//
// If position is negative then text is appended.
//
// This is the same as calling gtk.ComboBoxText.Insert() with a NULL ID string.
func (comboBox *ComboBoxText) InsertText(position int, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(position)
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_insert_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
	runtime.KeepAlive(text)
}

// Prepend prepends text to the list of strings stored in combo_box.
//
// If id is non-NULL then it is used as the ID of the row.
//
// This is the same as calling gtk.ComboBoxText.Insert() with a position of 0.
func (comboBox *ComboBoxText) Prepend(id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	if id != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(id)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_prepend(_arg0, _arg1, _arg2)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(id)
	runtime.KeepAlive(text)
}

// PrependText prepends text to the list of strings stored in combo_box.
//
// This is the same as calling gtk.ComboBoxText.InsertText() with a position of
// 0.
func (comboBox *ComboBoxText) PrependText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_prepend_text(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(text)
}

// Remove removes the string at position from combo_box.
func (comboBox *ComboBoxText) Remove(position int) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))
	_arg1 = C.int(position)

	C.gtk_combo_box_text_remove(_arg0, _arg1)
	runtime.KeepAlive(comboBox)
	runtime.KeepAlive(position)
}

// RemoveAll removes all the text entries from the combo box.
func (comboBox *ComboBoxText) RemoveAll() {
	var _arg0 *C.GtkComboBoxText // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(comboBox.Native()))

	C.gtk_combo_box_text_remove_all(_arg0)
	runtime.KeepAlive(comboBox)
}
