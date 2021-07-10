// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_entry_icon_position_get_type()), F: marshalEntryIconPosition},
		{T: externglib.Type(C.gtk_entry_get_type()), F: marshalyier},
	})
}

// EntryIconPosition specifies the side of the entry at which an icon is placed.
type EntryIconPosition int

const (
	// Primary: at the beginning of the entry (depending on the text direction).
	EntryIconPositionPrimary EntryIconPosition = iota
	// Secondary: at the end of the entry (depending on the text direction).
	EntryIconPositionSecondary
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	return EntryIconPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// yierOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type yierOverrider interface {
	Activate()
	Backspace()
	CopyClipboard()
	CutClipboard()
	FrameSize(x *int, y *int, width *int, height *int)
	TextAreaSize(x *int, y *int, width *int, height *int)
	InsertAtCursor(str string)
	InsertEmoji()
	PasteClipboard()
	PopulatePopup(popup Widgetter)
	ToggleOverwrite()
}

// yier describes Entry's methods.
type yier interface {
	gextras.Objector

	ActivatesDefault() bool
	Alignment() float32
	Attributes() *pango.AttrList
	Buffer() *EntryBuffer
	Completion() *EntryCompletion
	CurrentIconDragSource() int
	CursorHAdjustment() *Adjustment
	HasFrame() bool
	IconAtPos(x int, y int) int
	InnerBorder() *Border
	InputHints() InputHints
	InputPurpose() InputPurpose
	InvisibleChar() uint32
	Layout() *pango.Layout
	LayoutOffsets() (x int, y int)
	MaxLength() int
	MaxWidthChars() int
	OverwriteMode() bool
	PlaceholderText() string
	ProgressFraction() float64
	ProgressPulseStep() float64
	Tabs() *pango.TabArray
	Text() string
	TextArea() gdk.Rectangle
	TextLength() uint16
	Visibility() bool
	WidthChars() int
	GrabFocusWithoutSelecting()
	ImContextFilterKeypress(event *gdk.EventKey) bool
	LayoutIndexToTextIndex(layoutIndex int) int
	ProgressPulse()
	ResetImContext()
	SetActivatesDefault(setting bool)
	SetAlignment(xalign float32)
	SetAttributes(attrs *pango.AttrList)
	SetBuffer(buffer EntryBufferrer)
	SetCompletion(completion EntryCompletioner)
	SetCursorHAdjustment(adjustment Adjustmenter)
	SetHasFrame(setting bool)
	SetInnerBorder(border *Border)
	SetInvisibleChar(ch uint32)
	SetMaxLength(max int)
	SetMaxWidthChars(nChars int)
	SetOverwriteMode(overwrite bool)
	SetPlaceholderText(text string)
	SetProgressFraction(fraction float64)
	SetProgressPulseStep(fraction float64)
	SetTabs(tabs *pango.TabArray)
	SetText(text string)
	SetVisibility(visible bool)
	SetWidthChars(nChars int)
	TextIndexToLayoutIndex(textIndex int) int
	UnsetInvisibleChar()
}

// Entry: the Entry widget is a single line text entry widget. A fairly large
// set of key bindings are supported by default. If the entered text is longer
// than the allocation of the widget, the widget will scroll so that the cursor
// position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using gtk_entry_set_visibility(). In this mode,
// entered text is displayed using a “invisible” character. By default, GTK+
// picks the best invisible character that is available in the current font, but
// it can be changed with gtk_entry_set_invisible_char(). Since 2.16, GTK+
// displays a warning when Caps Lock or input methods might interfere with
// entering text in a password entry. The warning can be turned off with the
// Entry:caps-lock-warning property.
//
// Since 2.16, GtkEntry has the ability to display progress or activity
// information behind the text. To make an entry display such information, use
// gtk_entry_set_progress_fraction() or gtk_entry_set_progress_pulse_step().
//
// Additionally, GtkEntry can show icons at either side of the entry. These
// icons can be activatable by clicking, can be set up as drag source and can
// have tooltips. To add an icon, use gtk_entry_set_icon_from_gicon() or one of
// the various other functions that set an icon from a stock id, an icon name or
// a pixbuf. To trigger an action when the user clicks an icon, connect to the
// Entry::icon-press signal. To allow DND operations from an icon, use
// gtk_entry_set_icon_drag_source(). To set a tooltip on an icon, use
// gtk_entry_set_icon_tooltip_text() or the corresponding function for markup.
//
// Note that functionality or information that is only available by clicking on
// an icon in an entry may not be accessible at all to users which are not able
// to use a mouse or other pointing device. It is therefore recommended that any
// such functionality should also be available by other means, e.g. via the
// context menu of the entry.
//
// CSS nodes
//
//    entry[.read-only][.flat][.warning][.error]
//    ├── image.left
//    ├── image.right
//    ├── undershoot.left
//    ├── undershoot.right
//    ├── [selection]
//    ├── [progress[.pulse]]
//    ╰── [window.popup]
//
// GtkEntry has a main node with the name entry. Depending on the properties of
// the entry, the style classes .read-only and .flat may appear. The style
// classes .warning and .error may also be used with entries.
//
// When the entry shows icons, it adds subnodes with the name image and the
// style class .left or .right, depending on where the icon appears.
//
// When the entry has a selection, it adds a subnode with the name selection.
//
// When the entry shows progress, it adds a subnode with the name progress. The
// node has the style class .pulse when the shown progress is pulsing.
//
// The CSS node for a context menu is added as a subnode below entry as well.
//
// The undershoot nodes are used to draw the underflow indication when content
// is scrolled out of view. These nodes get the .left and .right style classes
// added depending on where the indication is drawn.
//
// When touch is used and touch selection handles are shown, they are using CSS
// nodes with name cursor-handle. They get the .top or .bottom style class
// depending on where they are shown in relation to the selection. If there is
// just a single handle for the text cursor, it gets the style class
// .insertion-cursor.
type Entry struct {
	*externglib.Object
	Widget
	Buildable
	CellEditable
	Editable
}

var _ yier = (*Entry)(nil)

func wrapyier(obj *externglib.Object) yier {
	return &Entry{
		Object: obj,
		Widget: Widget{
			Object: obj,
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
		},
		Buildable: Buildable{
			Object: obj,
		},
		CellEditable: CellEditable{
			Object: obj,
			Widget: Widget{
				Object: obj,
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
		Editable: Editable{
			Object: obj,
		},
	}
}

func marshalyier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapyier(obj), nil
}

// NewEntry creates a new entry.
func NewEntry() *Entry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_entry_new()

	var _entry *Entry // out

	_entry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Entry)

	return _entry
}

// NewEntryWithBuffer creates a new entry with the specified text buffer.
func NewEntryWithBuffer(buffer EntryBufferrer) *Entry {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_new_with_buffer(_arg1)

	var _entry *Entry // out

	_entry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Entry)

	return _entry
}

// ActivatesDefault retrieves the value set by
// gtk_entry_set_activates_default().
func (entry *Entry) ActivatesDefault() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Alignment gets the value set by gtk_entry_set_alignment().
func (entry *Entry) Alignment() float32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Attributes gets the attribute list that was set on the entry using
// gtk_entry_set_attributes(), if any.
func (entry *Entry) Attributes() *pango.AttrList {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_attributes(_arg0)

	var _attrList *pango.AttrList // out

	_attrList = (*pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)
	runtime.SetFinalizer(_attrList, func(v *pango.AttrList) {
		C.pango_attr_list_unref((*C.PangoAttrList)(unsafe.Pointer(v)))
	})

	return _attrList
}

// Buffer: get the EntryBuffer object which holds the text for this widget.
func (entry *Entry) Buffer() *EntryBuffer {
	var _arg0 *C.GtkEntry       // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_buffer(_arg0)

	var _entryBuffer *EntryBuffer // out

	_entryBuffer = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryBuffer)

	return _entryBuffer
}

// Completion returns the auxiliary completion object currently in use by
// @entry.
func (entry *Entry) Completion() *EntryCompletion {
	var _arg0 *C.GtkEntry           // out
	var _cret *C.GtkEntryCompletion // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_completion(_arg0)

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*EntryCompletion)

	return _entryCompletion
}

// CurrentIconDragSource returns the index of the icon which is the source of
// the current DND operation, or -1.
//
// This function is meant to be used in a Widget::drag-data-get callback.
func (entry *Entry) CurrentIconDragSource() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_current_icon_drag_source(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// CursorHAdjustment retrieves the horizontal cursor adjustment for the entry.
// See gtk_entry_set_cursor_hadjustment().
func (entry *Entry) CursorHAdjustment() *Adjustment {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_cursor_hadjustment(_arg0)

	var _adjustment *Adjustment // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Adjustment)

	return _adjustment
}

// HasFrame gets the value set by gtk_entry_set_has_frame().
func (entry *Entry) HasFrame() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconAtPos finds the icon at the given position and return its index. The
// position’s coordinates are relative to the @entry’s top left corner. If @x,
// @y doesn’t lie inside an icon, -1 is returned. This function is intended for
// use in a Widget::query-tooltip signal handler.
func (entry *Entry) IconAtPos(x int, y int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _arg2 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_entry_get_icon_at_pos(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InnerBorder: this function returns the entry’s Entry:inner-border property.
// See gtk_entry_set_inner_border() for more information.
//
// Deprecated: Use the standard border and padding CSS properties (through
// objects like StyleContext and CssProvider); the value returned by this
// function is ignored by Entry.
func (entry *Entry) InnerBorder() *Border {
	var _arg0 *C.GtkEntry  // out
	var _cret *C.GtkBorder // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_inner_border(_arg0)

	var _border *Border // out

	_border = (*Border)(unsafe.Pointer(_cret))

	return _border
}

// InputHints gets the value of the Entry:input-hints property.
func (entry *Entry) InputHints() InputHints {
	var _arg0 *C.GtkEntry     // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = (InputHints)(_cret)

	return _inputHints
}

// InputPurpose gets the value of the Entry:input-purpose property.
func (entry *Entry) InputPurpose() InputPurpose {
	var _arg0 *C.GtkEntry       // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = (InputPurpose)(_cret)

	return _inputPurpose
}

// InvisibleChar retrieves the character displayed in place of the real
// characters for entries with visibility set to false. See
// gtk_entry_set_invisible_char().
func (entry *Entry) InvisibleChar() uint32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gunichar  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

// Layout gets the Layout used to display the entry. The layout is useful to
// e.g. convert text positions to pixel positions, in combination with
// gtk_entry_get_layout_offsets(). The returned layout is owned by the entry and
// must not be modified or freed by the caller.
//
// Keep in mind that the layout text may contain a preedit string, so
// gtk_entry_layout_index_to_text_index() and
// gtk_entry_text_index_to_layout_index() are needed to convert byte indices in
// the layout to byte indices in the entry contents.
func (entry *Entry) Layout() *pango.Layout {
	var _arg0 *C.GtkEntry    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_layout(_arg0)

	var _layout *pango.Layout // out

	_layout = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*pango.Layout)

	return _layout
}

// LayoutOffsets obtains the position of the Layout used to render text in the
// entry, in widget coordinates. Useful if you want to line up the text in an
// entry with some other text, e.g. when using the entry to implement editable
// cells in a sheet widget.
//
// Also useful to convert mouse events into coordinates inside the Layout, e.g.
// to take some action if some part of the entry text is clicked.
//
// Note that as the user scrolls around in the entry the offsets will change;
// you’ll need to connect to the “notify::scroll-offset” signal to track this.
// Remember when using the Layout functions you need to convert to and from
// pixels using PANGO_PIXELS() or NGO_SCALE.
//
// Keep in mind that the layout text may contain a preedit string, so
// gtk_entry_layout_index_to_text_index() and
// gtk_entry_text_index_to_layout_index() are needed to convert byte indices in
// the layout to byte indices in the entry contents.
func (entry *Entry) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// MaxLength retrieves the maximum allowed length of the text in @entry. See
// gtk_entry_set_max_length().
//
// This is equivalent to getting @entry's EntryBuffer and calling
// gtk_entry_buffer_get_max_length() on it.
func (entry *Entry) MaxLength() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxWidthChars retrieves the desired maximum width of @entry, in characters.
// See gtk_entry_set_max_width_chars().
func (entry *Entry) MaxWidthChars() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_max_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// OverwriteMode gets the value set by gtk_entry_set_overwrite_mode().
func (entry *Entry) OverwriteMode() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PlaceholderText retrieves the text that will be displayed when @entry is
// empty and unfocused
func (entry *Entry) PlaceholderText() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ProgressFraction returns the current fraction of the task that’s been
// completed. See gtk_entry_set_progress_fraction().
func (entry *Entry) ProgressFraction() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_progress_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ProgressPulseStep retrieves the pulse step set with
// gtk_entry_set_progress_pulse_step().
func (entry *Entry) ProgressPulseStep() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_progress_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Tabs gets the tabstops that were set on the entry using gtk_entry_set_tabs(),
// if any.
func (entry *Entry) Tabs() *pango.TabArray {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_tabs(_arg0)

	var _tabArray *pango.TabArray // out

	_tabArray = (*pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

// Text retrieves the contents of the entry widget. See also
// gtk_editable_get_chars().
//
// This is equivalent to getting @entry's EntryBuffer and calling
// gtk_entry_buffer_get_text() on it.
func (entry *Entry) Text() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// TextArea gets the area where the entry’s text is drawn. This function is
// useful when drawing something to the entry in a draw callback.
//
// If the entry is not realized, @text_area is filled with zeros.
//
// See also gtk_entry_get_icon_area().
func (entry *Entry) TextArea() gdk.Rectangle {
	var _arg0 *C.GtkEntry    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_get_text_area(_arg0, &_arg1)

	var _textArea gdk.Rectangle // out

	_textArea = *(*gdk.Rectangle)(unsafe.Pointer((&_arg1)))

	return _textArea
}

// TextLength retrieves the current length of the text in @entry.
//
// This is equivalent to getting @entry's EntryBuffer and calling
// gtk_entry_buffer_get_length() on it.
func (entry *Entry) TextLength() uint16 {
	var _arg0 *C.GtkEntry // out
	var _cret C.guint16   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Visibility retrieves whether the text in @entry is visible. See
// gtk_entry_set_visibility().
func (entry *Entry) Visibility() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WidthChars gets the value set by gtk_entry_set_width_chars().
func (entry *Entry) WidthChars() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_entry_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GrabFocusWithoutSelecting causes @entry to have keyboard focus.
//
// It behaves like gtk_widget_grab_focus(), except that it doesn't select the
// contents of the entry. You only want to call this on some special entries
// which the user usually doesn't want to replace all text in, such as
// search-as-you-type entries.
func (entry *Entry) GrabFocusWithoutSelecting() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_grab_focus_without_selecting(_arg0)
}

// ImContextFilterKeypress: allow the Entry input method to internally handle
// key press and release events. If this function returns true, then no further
// processing should be done for this key event. See
// gtk_im_context_filter_keypress().
//
// Note that you are expected to call this function from your handler when
// overriding key event handling. This is needed in the case when you need to
// insert your own key handling between the input method and the default key
// event handling of the Entry. See gtk_text_view_reset_im_context() for an
// example of use.
func (entry *Entry) ImContextFilterKeypress(event *gdk.EventKey) bool {
	var _arg0 *C.GtkEntry    // out
	var _arg1 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GdkEventKey)(unsafe.Pointer(event))

	_cret = C.gtk_entry_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LayoutIndexToTextIndex converts from a position in the entry’s Layout
// (returned by gtk_entry_get_layout()) to a position in the entry contents
// (returned by gtk_entry_get_text()).
func (entry *Entry) LayoutIndexToTextIndex(layoutIndex int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(layoutIndex)

	_cret = C.gtk_entry_layout_index_to_text_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ProgressPulse indicates that some progress is made, but you don’t know how
// much. Causes the entry’s progress indicator to enter “activity mode,” where a
// block bounces back and forth. Each call to gtk_entry_progress_pulse() causes
// the block to move by a little bit (the amount of movement per pulse is
// determined by gtk_entry_set_progress_pulse_step()).
func (entry *Entry) ProgressPulse() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_progress_pulse(_arg0)
}

// ResetImContext: reset the input method context of the entry if needed.
//
// This can be necessary in the case where modifying the buffer would confuse
// on-going input method behavior.
func (entry *Entry) ResetImContext() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_reset_im_context(_arg0)
}

// SetActivatesDefault: if @setting is true, pressing Enter in the @entry will
// activate the default widget for the window containing the entry. This usually
// means that the dialog box containing the entry will be closed, since the
// default widget is usually one of the dialog buttons.
//
// (For experts: if @setting is true, the entry calls
// gtk_window_activate_default() on the window containing the entry, in the
// default handler for the Entry::activate signal.)
func (entry *Entry) SetActivatesDefault(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_activates_default(_arg0, _arg1)
}

// SetAlignment sets the alignment for the contents of the entry. This controls
// the horizontal positioning of the contents when the displayed text is shorter
// than the width of the entry.
func (entry *Entry) SetAlignment(xalign float32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_entry_set_alignment(_arg0, _arg1)
}

// SetAttributes sets a AttrList; the attributes in the list are applied to the
// entry text.
func (entry *Entry) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_entry_set_attributes(_arg0, _arg1)
}

// SetBuffer: set the EntryBuffer object which holds the text for this widget.
func (entry *Entry) SetBuffer(buffer EntryBufferrer) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_entry_set_buffer(_arg0, _arg1)
}

// SetCompletion sets @completion to be the auxiliary completion object to use
// with @entry. All further configuration of the completion mechanism is done on
// @completion using the EntryCompletion API. Completion is disabled if
// @completion is set to nil.
func (entry *Entry) SetCompletion(completion EntryCompletioner) {
	var _arg0 *C.GtkEntry           // out
	var _arg1 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_set_completion(_arg0, _arg1)
}

// SetCursorHAdjustment hooks up an adjustment to the cursor position in an
// entry, so that when the cursor is moved, the adjustment is scrolled to show
// that position. See gtk_scrolled_window_get_hadjustment() for a typical way of
// obtaining the adjustment.
//
// The adjustment has to be in pixel units and in the same coordinate system as
// the entry.
func (entry *Entry) SetCursorHAdjustment(adjustment Adjustmenter) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_entry_set_cursor_hadjustment(_arg0, _arg1)
}

// SetHasFrame sets whether the entry has a beveled frame around it.
func (entry *Entry) SetHasFrame(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_has_frame(_arg0, _arg1)
}

// SetInnerBorder sets entry’s inner-border property to @border, or clears it if
// nil is passed. The inner-border is the area around the entry’s text, but
// inside its frame.
//
// If set, this property overrides the inner-border style property. Overriding
// the style-provided border is useful when you want to do in-place editing of
// some text in a canvas or list widget, where pixel-exact positioning of the
// entry is important.
//
// Deprecated: Use the standard border and padding CSS properties (through
// objects like StyleContext and CssProvider); the value set with this function
// is ignored by Entry.
func (entry *Entry) SetInnerBorder(border *Border) {
	var _arg0 *C.GtkEntry  // out
	var _arg1 *C.GtkBorder // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkBorder)(unsafe.Pointer(border))

	C.gtk_entry_set_inner_border(_arg0, _arg1)
}

// SetInvisibleChar sets the character to use in place of the actual text when
// gtk_entry_set_visibility() has been called to set text visibility to false.
// i.e. this is the character used in “password mode” to show the user how many
// characters have been typed. By default, GTK+ picks the best invisible char
// available in the current font. If you set the invisible char to 0, then the
// user will get no feedback at all; there will be no text on the screen as they
// type.
func (entry *Entry) SetInvisibleChar(ch uint32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gunichar  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_entry_set_invisible_char(_arg0, _arg1)
}

// SetMaxLength sets the maximum allowed length of the contents of the widget.
// If the current contents are longer than the given length, then they will be
// truncated to fit.
//
// This is equivalent to getting @entry's EntryBuffer and calling
// gtk_entry_buffer_set_max_length() on it. ]|
func (entry *Entry) SetMaxLength(max int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(max)

	C.gtk_entry_set_max_length(_arg0, _arg1)
}

// SetMaxWidthChars sets the desired maximum width in characters of @entry.
func (entry *Entry) SetMaxWidthChars(nChars int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_entry_set_max_width_chars(_arg0, _arg1)
}

// SetOverwriteMode sets whether the text is overwritten when typing in the
// Entry.
func (entry *Entry) SetOverwriteMode(overwrite bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_overwrite_mode(_arg0, _arg1)
}

// SetPlaceholderText sets text to be displayed in @entry when it is empty and
// unfocused. This can be used to give a visual hint of the expected contents of
// the Entry.
//
// Note that since the placeholder text gets removed when the entry received
// focus, using this feature is a bit problematic if the entry is given the
// initial focus in a window. Sometimes this can be worked around by delaying
// the initial focus setting until the first key event arrives.
func (entry *Entry) SetPlaceholderText(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_placeholder_text(_arg0, _arg1)
}

// SetProgressFraction causes the entry’s progress indicator to “fill in” the
// given fraction of the bar. The fraction should be between 0.0 and 1.0,
// inclusive.
func (entry *Entry) SetProgressFraction(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_entry_set_progress_fraction(_arg0, _arg1)
}

// SetProgressPulseStep sets the fraction of total entry width to move the
// progress bouncing block for each call to gtk_entry_progress_pulse().
func (entry *Entry) SetProgressPulseStep(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_entry_set_progress_pulse_step(_arg0, _arg1)
}

// SetTabs sets a TabArray; the tabstops in the array are applied to the entry
// text.
func (entry *Entry) SetTabs(tabs *pango.TabArray) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.gtk_entry_set_tabs(_arg0, _arg1)
}

// SetText sets the text in the widget to the given value, replacing the current
// contents.
//
// See gtk_entry_buffer_set_text().
func (entry *Entry) SetText(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_text(_arg0, _arg1)
}

// SetVisibility sets whether the contents of the entry are visible or not. When
// visibility is set to false, characters are displayed as the invisible char,
// and will also appear that way when the text in the entry widget is copied
// elsewhere.
//
// By default, GTK+ picks the best invisible character available in the current
// font, but it can be changed with gtk_entry_set_invisible_char().
//
// Note that you probably want to set Entry:input-purpose to
// GTK_INPUT_PURPOSE_PASSWORD or GTK_INPUT_PURPOSE_PIN to inform input methods
// about the purpose of this entry, in addition to setting visibility to false.
func (entry *Entry) SetVisibility(visible bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_visibility(_arg0, _arg1)
}

// SetWidthChars changes the size request of the entry to be about the right
// size for @n_chars characters. Note that it changes the size request, the size
// can still be affected by how you pack the widget into containers. If @n_chars
// is -1, the size reverts to the default entry size.
func (entry *Entry) SetWidthChars(nChars int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_entry_set_width_chars(_arg0, _arg1)
}

// TextIndexToLayoutIndex converts from a position in the entry contents
// (returned by gtk_entry_get_text()) to a position in the entry’s Layout
// (returned by gtk_entry_get_layout(), with text retrieved via
// pango_layout_get_text()).
func (entry *Entry) TextIndexToLayoutIndex(textIndex int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = C.gint(textIndex)

	_cret = C.gtk_entry_text_index_to_layout_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UnsetInvisibleChar unsets the invisible char previously set with
// gtk_entry_set_invisible_char(). So that the default invisible char is used
// again.
func (entry *Entry) UnsetInvisibleChar() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))

	C.gtk_entry_unset_invisible_char(_arg0)
}
