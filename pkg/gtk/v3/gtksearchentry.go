// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_search_entry_get_type()), F: marshalSearchEntry},
	})
}

// SearchEntry is a subclass of Entry that has been tailored for use as a search
// entry.
//
// It will show an inactive symbolic “find” icon when the search entry is empty,
// and a symbolic “clear” icon when there is text. Clicking on the “clear” icon
// will empty the search entry.
//
// Note that the search/clear icon is shown using a secondary icon, and thus
// does not work if you are using the secondary icon position for some other
// purpose.
//
// To make filtering appear more reactive, it is a good idea to not react to
// every change in the entry text immediately, but only after a short delay. To
// support this, SearchEntry emits the SearchEntry::search-changed signal which
// can be used instead of the Editable::changed signal.
//
// The SearchEntry::previous-match, SearchEntry::next-match and
// SearchEntry::stop-search signals can be used to implement moving between
// search results and ending the search.
//
// Often, GtkSearchEntry will be fed events by means of being placed inside a
// SearchBar. If that is not the case, you can use
// gtk_search_entry_handle_event() to pass events.
type SearchEntry interface {
	Entry
	Buildable
	CellEditable
	Editable
}

// searchEntry implements the SearchEntry interface.
type searchEntry struct {
	Entry
	Buildable
	CellEditable
	Editable
}

var _ SearchEntry = (*searchEntry)(nil)

// WrapSearchEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapSearchEntry(obj *externglib.Object) SearchEntry {
	return SearchEntry{
		Entry:        WrapEntry(obj),
		Buildable:    WrapBuildable(obj),
		CellEditable: WrapCellEditable(obj),
		Editable:     WrapEditable(obj),
	}
}

func marshalSearchEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSearchEntry(obj), nil
}

// NewSearchEntry constructs a class SearchEntry.
func NewSearchEntry() SearchEntry {
	var cret C.GtkSearchEntry
	var goret1 SearchEntry

	cret = C.gtk_search_entry_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(SearchEntry)

	return goret1
}
