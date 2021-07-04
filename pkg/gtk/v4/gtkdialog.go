// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
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
		{T: externglib.Type(C.gtk_response_type_get_type()), F: marshalResponseType},
		{T: externglib.Type(C.gtk_dialog_flags_get_type()), F: marshalDialogFlags},
		{T: externglib.Type(C.gtk_dialog_get_type()), F: marshalDialog},
	})
}

// ResponseType: predefined values for use as response ids in
// gtk_dialog_add_button().
//
// All predefined values are negative; GTK leaves values of 0 or greater for
// application-defined response ids.
type ResponseType int

const (
	// none: returned if an action widget has no response id, or if the dialog
	// gets programmatically hidden or destroyed
	ResponseTypeNone ResponseType = -1
	// reject: generic response id, not used by GTK dialogs
	ResponseTypeReject ResponseType = -2
	// accept: generic response id, not used by GTK dialogs
	ResponseTypeAccept ResponseType = -3
	// DeleteEvent: returned if the dialog is deleted
	ResponseTypeDeleteEvent ResponseType = -4
	// ok: returned by OK buttons in GTK dialogs
	ResponseTypeOk ResponseType = -5
	// cancel: returned by Cancel buttons in GTK dialogs
	ResponseTypeCancel ResponseType = -6
	// close: returned by Close buttons in GTK dialogs
	ResponseTypeClose ResponseType = -7
	// yes: returned by Yes buttons in GTK dialogs
	ResponseTypeYes ResponseType = -8
	// no: returned by No buttons in GTK dialogs
	ResponseTypeNo ResponseType = -9
	// apply: returned by Apply buttons in GTK dialogs
	ResponseTypeApply ResponseType = -10
	// help: returned by Help buttons in GTK dialogs
	ResponseTypeHelp ResponseType = -11
)

func marshalResponseType(p uintptr) (interface{}, error) {
	return ResponseType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DialogFlags flags used to influence dialog construction.
type DialogFlags int

const (
	// DialogFlagsModal: make the constructed dialog modal
	DialogFlagsModal DialogFlags = 0b1
	// DialogFlagsDestroyWithParent: destroy the dialog when its parent is
	// destroyed
	DialogFlagsDestroyWithParent DialogFlags = 0b10
	// DialogFlagsUseHeaderBar: create dialog with actions in header bar instead
	// of action area
	DialogFlagsUseHeaderBar DialogFlags = 0b100
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	return DialogFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Dialog dialogs are a convenient way to prompt the user for a small amount of
// input.
//
// !An example GtkDialog (dialog.png)
//
// Typical uses are to display a message, ask a question, or anything else that
// does not require extensive effort on the user’s part.
//
// The main area of a `GtkDialog` is called the "content area", and is yours to
// populate with widgets such a `GtkLabel` or `GtkEntry`, to present your
// information, questions, or tasks to the user.
//
// In addition, dialogs allow you to add "action widgets". Most commonly, action
// widgets are buttons. Depending on the platform, action widgets may be
// presented in the header bar at the top of the window, or at the bottom of the
// window. To add action widgets, create your `GtkDialog` using
// [ctor@Gtk.Dialog.new_with_buttons], or use [method@Gtk.Dialog.add_button],
// [method@Gtk.Dialog.add_buttons], or [method@Gtk.Dialog.add_action_widget].
//
// `GtkDialogs` uses some heuristics to decide whether to add a close button to
// the window decorations. If any of the action buttons use the response ID
// GTK_RESPONSE_CLOSE or GTK_RESPONSE_CANCEL, the close button is omitted.
//
// Clicking a button that was added as an action widget will emit the
// [signal@Gtk.Dialog::response] signal with a response ID that you specified.
// GTK will never assign a meaning to positive response IDs; these are entirely
// user-defined. But for convenience, you can use the response IDs in the
// [enum@Gtk.ResponseType] enumeration (these all have values less than zero).
// If a dialog receives a delete event, the [signal@Gtk.Dialog::response] signal
// will be emitted with the GTK_RESPONSE_DELETE_EVENT response ID.
//
// Dialogs are created with a call to [ctor@Gtk.Dialog.new] or
// [ctor@Gtk.Dialog.new_with_buttons]. The latter is recommended; it allows you
// to set the dialog title, some convenient flags, and add buttons.
//
// A “modal” dialog (that is, one which freezes the rest of the application from
// user input), can be created by calling [method@Gtk.Window.set_modal] on the
// dialog. When using [ctor@Gtk.Dialog.new_with_buttons], you can also pass the
// GTK_DIALOG_MODAL flag to make a dialog modal.
//
// For the simple dialog in the following example, a [class@Gtk.MessageDialog]
// would save some effort. But you’d need to create the dialog contents manually
// if you had more than a simple message in the dialog.
//
// An example for simple `GtkDialog` usage: “`c // Function to open a dialog box
// with a message void quick_message (GtkWindow *parent, char *message) {
// GtkWidget *dialog, *label, *content_area; GtkDialogFlags flags;
//
//    // Create the widgets
//    flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//    dialog = gtk_dialog_new_with_buttons ("Message",
//                                          parent,
//                                          flags,
//                                          _("_OK"),
//                                          GTK_RESPONSE_NONE,
//                                          NULL);
//    content_area = gtk_dialog_get_content_area (GTK_DIALOG (dialog));
//    label = gtk_label_new (message);
//
//    // Ensure that the dialog box is destroyed when the user responds
//
//    g_signal_connect_swapped (dialog,
//                              "response",
//                              G_CALLBACK (gtk_window_destroy),
//                              dialog);
//
//    // Add the label, and show everything we’ve added
//
//    gtk_box_append (GTK_BOX (content_area), label);
//    gtk_widget_show (dialog);
//
// } “`
//
//
// GtkDialog as GtkBuildable
//
// The `GtkDialog` implementation of the `GtkBuildable` interface exposes the
// @content_area as an internal child with the name “content_area”.
//
// `GtkDialog` supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs @action_area). To mark a response as
// default, set the “default“ attribute of the <action-widget> element to true.
//
// `GtkDialog` supports adding action widgets by specifying “action“ as the
// “type“ attribute of a <child> element. The widget will be added either to the
// action area or the headerbar of the dialog, depending on the “use-header-bar“
// property. The response id has to be associated with the action widget using
// the <action-widgets> element.
//
// An example of a Dialog UI definition fragment: “`xml <object
// class="GtkDialog" id="dialog1"> <child type="action"> <object
// class="GtkButton" id="button_cancel"/> </child> <child type="action"> <object
// class="GtkButton" id="button_ok"> </object> </child> <action-widgets>
// <action-widget response="cancel">button_cancel</action-widget> <action-widget
// response="ok" default="true">button_ok</action-widget> </action-widgets>
// </object> “`
//
//
// Accessibility
//
// `GtkDialog` uses the GTK_ACCESSIBLE_ROLE_DIALOG role.
type Dialog interface {
	Window

	AddActionWidgetDialog(child Widget, responseId int)

	AddButtonDialog(buttonText string, responseId int) Widget

	ContentArea() Box

	HeaderBar() HeaderBar

	ResponseForWidget(widget Widget) int

	WidgetForResponse(responseId int) Widget

	ResponseDialog(responseId int)

	SetDefaultResponseDialog(responseId int)

	SetResponseSensitiveDialog(responseId int, setting bool)
}

// dialog implements the Dialog class.
type dialog struct {
	Window
}

// WrapDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapDialog(obj *externglib.Object) Dialog {
	return dialog{
		Window: WrapWindow(obj),
	}
}

func marshalDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDialog(obj), nil
}

func NewDialog() Dialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_new()

	var _dialog Dialog // out

	_dialog = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Dialog)

	return _dialog
}

func (d dialog) AddActionWidgetDialog(child Widget, responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(responseId)

	C.gtk_dialog_add_action_widget(_arg0, _arg1, _arg2)
}

func (d dialog) AddButtonDialog(buttonText string, responseId int) Widget {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(buttonText))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(responseId)

	_cret = C.gtk_dialog_add_button(_arg0, _arg1, _arg2)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (d dialog) ContentArea() Box {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_dialog_get_content_area(_arg0)

	var _box Box // out

	_box = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Box)

	return _box
}

func (d dialog) HeaderBar() HeaderBar {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))

	_cret = C.gtk_dialog_get_header_bar(_arg0)

	var _headerBar HeaderBar // out

	_headerBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(HeaderBar)

	return _headerBar
}

func (d dialog) ResponseForWidget(widget Widget) int {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _cret C.int        // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_dialog_get_response_for_widget(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d dialog) WidgetForResponse(responseId int) Widget {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(responseId)

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (d dialog) ResponseDialog(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_response(_arg0, _arg1)
}

func (d dialog) SetDefaultResponseDialog(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_set_default_response(_arg0, _arg1)
}

func (d dialog) SetResponseSensitiveDialog(responseId int, setting bool) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_dialog_set_response_sensitive(_arg0, _arg1, _arg2)
}

func (s dialog) Display() gdk.Display {
	return WrapRoot(gextras.InternObject(s)).Display()
}

func (s dialog) Focus() Widget {
	return WrapRoot(gextras.InternObject(s)).Focus()
}

func (s dialog) SetFocus(focus Widget) {
	WrapRoot(gextras.InternObject(s)).SetFocus(focus)
}

func (s dialog) Renderer() gsk.Renderer {
	return WrapNative(gextras.InternObject(s)).Renderer()
}

func (s dialog) Surface() gdk.Surface {
	return WrapNative(gextras.InternObject(s)).Surface()
}

func (s dialog) SurfaceTransform() (x float64, y float64) {
	return WrapNative(gextras.InternObject(s)).SurfaceTransform()
}

func (s dialog) Realize() {
	WrapNative(gextras.InternObject(s)).Realize()
}

func (s dialog) Unrealize() {
	WrapNative(gextras.InternObject(s)).Unrealize()
}

func (s dialog) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s dialog) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s dialog) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s dialog) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s dialog) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s dialog) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s dialog) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b dialog) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}