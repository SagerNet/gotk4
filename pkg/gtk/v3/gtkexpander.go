// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpander},
	})
}

// Expander: a Expander allows the user to hide or show its child by clicking on
// an expander triangle similar to the triangles used in a TreeView.
//
// Normally you use an expander as you would use any other descendant of Bin;
// you create the child widget and use gtk_container_add() to add it to the
// expander. When the expander is toggled, it will take care of showing and
// hiding the child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a Expander but do not add a child to it.
// The expander widget has an Expander:expanded property which can be used to
// monitor its expansion state. You should watch this property with a signal
// connection as follows:
//
//    expander
//    ├── title
//    │   ├── arrow
//    │   ╰── <label widget>
//    ╰── <child>
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
type Expander interface {
	Bin
	Buildable

	// Expanded queries a Expander and returns its current state. Returns true
	// if the child widget is revealed.
	//
	// See gtk_expander_set_expanded().
	Expanded() bool
	// Label fetches the text from a label widget including any embedded
	// underlines indicating mnemonics and Pango markup, as set by
	// gtk_expander_set_label(). If the label text has not been set the return
	// value will be nil. This will be the case if you create an empty button
	// with gtk_button_new() to use as a container.
	//
	// Note that this function behaved differently in versions prior to 2.14 and
	// used to return the label text stripped of embedded underlines indicating
	// mnemonics and Pango markup. This problem can be avoided by fetching the
	// label text directly from the label widget.
	Label() string
	// LabelFill returns whether the label widget will fill all available
	// horizontal space allocated to @expander.
	LabelFill() bool
	// LabelWidget retrieves the label widget for the frame. See
	// gtk_expander_set_label_widget().
	LabelWidget() Widget
	// ResizeToplevel returns whether the expander will resize the toplevel
	// widget containing the expander upon resizing and collpasing.
	ResizeToplevel() bool
	// Spacing gets the value set by gtk_expander_set_spacing().
	Spacing() int
	// UseMarkup returns whether the label’s text is interpreted as marked up
	// with the [Pango text markup language][PangoMarkupFormat]. See
	// gtk_expander_set_use_markup().
	UseMarkup() bool
	// UseUnderline returns whether an embedded underline in the expander label
	// indicates a mnemonic. See gtk_expander_set_use_underline().
	UseUnderline() bool
	// SetExpanded sets the state of the expander. Set to true, if you want the
	// child widget to be revealed, and false if you want the child widget to be
	// hidden.
	SetExpanded(expanded bool)
	// SetLabel sets the text of the label of the expander to @label.
	//
	// This will also clear any previously set labels.
	SetLabel(label string)
	// SetLabelFill sets whether the label widget should fill all available
	// horizontal space allocated to @expander.
	//
	// Note that this function has no effect since 3.20.
	SetLabelFill(labelFill bool)
	// SetLabelWidget: set the label widget for the expander. This is the widget
	// that will appear embedded alongside the expander arrow.
	SetLabelWidget(labelWidget Widget)
	// SetResizeToplevel sets whether the expander will resize the toplevel
	// widget containing the expander upon resizing and collpasing.
	SetResizeToplevel(resizeToplevel bool)
	// SetSpacing sets the spacing field of @expander, which is the number of
	// pixels to place between expander and the child.
	SetSpacing(spacing int)
	// SetUseMarkup sets whether the text of the label contains markup in
	// [Pango’s text markup language][PangoMarkupFormat]. See
	// gtk_label_set_markup().
	SetUseMarkup(useMarkup bool)
	// SetUseUnderline: if true, an underline in the text of the expander label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(useUnderline bool)
}

// expander implements the Expander interface.
type expander struct {
	Bin
	Buildable
}

var _ Expander = (*expander)(nil)

// WrapExpander wraps a GObject to the right type. It is
// primarily used internally.
func WrapExpander(obj *externglib.Object) Expander {
	return Expander{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalExpander(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapExpander(obj), nil
}

// NewExpander constructs a class Expander.
func NewExpander(label string) Expander {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkExpander
	var goret1 Expander

	cret = C.gtk_expander_new(label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Expander)

	return goret1
}

// NewExpanderWithMnemonic constructs a class Expander.
func NewExpanderWithMnemonic(label string) Expander {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkExpander
	var goret1 Expander

	cret = C.gtk_expander_new_with_mnemonic(label)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Expander)

	return goret1
}

// Expanded queries a Expander and returns its current state. Returns true
// if the child widget is revealed.
//
// See gtk_expander_set_expanded().
func (e expander) Expanded() bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_expander_get_expanded(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Label fetches the text from a label widget including any embedded
// underlines indicating mnemonics and Pango markup, as set by
// gtk_expander_set_label(). If the label text has not been set the return
// value will be nil. This will be the case if you create an empty button
// with gtk_button_new() to use as a container.
//
// Note that this function behaved differently in versions prior to 2.14 and
// used to return the label text stripped of embedded underlines indicating
// mnemonics and Pango markup. This problem can be avoided by fetching the
// label text directly from the label widget.
func (e expander) Label() string {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret *C.gchar
	var goret1 string

	cret = C.gtk_expander_get_label(arg0)

	goret1 = C.GoString(cret)

	return goret1
}

// LabelFill returns whether the label widget will fill all available
// horizontal space allocated to @expander.
func (e expander) LabelFill() bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_expander_get_label_fill(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
func (e expander) LabelWidget() Widget {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret *C.GtkWidget
	var goret1 Widget

	cret = C.gtk_expander_get_label_widget(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret1
}

// ResizeToplevel returns whether the expander will resize the toplevel
// widget containing the expander upon resizing and collpasing.
func (e expander) ResizeToplevel() bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_expander_get_resize_toplevel(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Spacing gets the value set by gtk_expander_set_spacing().
func (e expander) Spacing() int {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gtk_expander_get_spacing(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// UseMarkup returns whether the label’s text is interpreted as marked up
// with the [Pango text markup language][PangoMarkupFormat]. See
// gtk_expander_set_use_markup().
func (e expander) UseMarkup() bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_expander_get_use_markup(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// UseUnderline returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
func (e expander) UseUnderline() bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_expander_get_use_underline(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetExpanded sets the state of the expander. Set to true, if you want the
// child widget to be revealed, and false if you want the child widget to be
// hidden.
func (e expander) SetExpanded(expanded bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if expanded {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_expanded(arg0, expanded)
}

// SetLabel sets the text of the label of the expander to @label.
//
// This will also clear any previously set labels.
func (e expander) SetLabel(label string) {
	var arg0 *C.GtkExpander
	var arg1 *C.gchar

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_expander_set_label(arg0, label)
}

// SetLabelFill sets whether the label widget should fill all available
// horizontal space allocated to @expander.
//
// Note that this function has no effect since 3.20.
func (e expander) SetLabelFill(labelFill bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if labelFill {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_label_fill(arg0, labelFill)
}

// SetLabelWidget: set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
func (e expander) SetLabelWidget(labelWidget Widget) {
	var arg0 *C.GtkExpander
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_expander_set_label_widget(arg0, labelWidget)
}

// SetResizeToplevel sets whether the expander will resize the toplevel
// widget containing the expander upon resizing and collpasing.
func (e expander) SetResizeToplevel(resizeToplevel bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if resizeToplevel {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_resize_toplevel(arg0, resizeToplevel)
}

// SetSpacing sets the spacing field of @expander, which is the number of
// pixels to place between expander and the child.
func (e expander) SetSpacing(spacing int) {
	var arg0 *C.GtkExpander
	var arg1 C.gint

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(spacing)

	C.gtk_expander_set_spacing(arg0, spacing)
}

// SetUseMarkup sets whether the text of the label contains markup in
// [Pango’s text markup language][PangoMarkupFormat]. See
// gtk_label_set_markup().
func (e expander) SetUseMarkup(useMarkup bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useMarkup {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_use_markup(arg0, useMarkup)
}

// SetUseUnderline: if true, an underline in the text of the expander label
// indicates the next character should be used for the mnemonic accelerator
// key.
func (e expander) SetUseUnderline(useUnderline bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useUnderline {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_use_underline(arg0, useUnderline)
}

type ExpanderPrivate struct {
	native C.GtkExpanderPrivate
}

// WrapExpanderPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapExpanderPrivate(ptr unsafe.Pointer) *ExpanderPrivate {
	if ptr == nil {
		return nil
	}

	return (*ExpanderPrivate)(ptr)
}

func marshalExpanderPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapExpanderPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *ExpanderPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}
