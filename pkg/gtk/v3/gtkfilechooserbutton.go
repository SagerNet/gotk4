// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_file_chooser_button_get_type()), F: marshalFileChooserButton},
	})
}

// FileChooserButton: the FileChooserButton is a widget that lets the user
// select a file. It implements the FileChooser interface. Visually, it is a
// file name with a button to bring up a FileChooserDialog. The user can then
// use that dialog to change the file associated with that button. This widget
// does not support setting the FileChooser:select-multiple property to true.
//
// Create a button to let the user select a file in /etc
//
//    {
//      GtkWidget *button;
//
//      button = gtk_file_chooser_button_new (_("Select a file"),
//                                            GTK_FILE_CHOOSER_ACTION_OPEN);
//      gtk_file_chooser_set_current_folder (GTK_FILE_CHOOSER (button),
//                                           "/etc");
//    }
//
// The FileChooserButton supports the FileChooserActions
// GTK_FILE_CHOOSER_ACTION_OPEN and GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER.
//
// > The FileChooserButton will ellipsize the label, and will thus > request
// little horizontal space. To give the button more space, > you should call
// gtk_widget_get_preferred_size(), > gtk_file_chooser_button_set_width_chars(),
// or pack the button in > such a way that other interface elements give space
// to the > widget.
//
//
// CSS nodes
//
// GtkFileChooserButton has a CSS node with name “filechooserbutton”, containing
// a subnode for the internal button with name “button” and style class “.file”.
type FileChooserButton interface {
	Box
	FileChooser

	// FocusOnClick:
	FocusOnClick() bool
	// Title:
	Title() string
	// WidthChars:
	WidthChars() int
	// SetFocusOnClickFileChooserButton:
	SetFocusOnClickFileChooserButton(focusOnClick bool)
	// SetTitleFileChooserButton:
	SetTitleFileChooserButton(title string)
	// SetWidthCharsFileChooserButton:
	SetWidthCharsFileChooserButton(nChars int)
}

// fileChooserButton implements the FileChooserButton class.
type fileChooserButton struct {
	Box
}

// WrapFileChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserButton(obj *externglib.Object) FileChooserButton {
	return fileChooserButton{
		Box: WrapBox(obj),
	}
}

func marshalFileChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserButton(obj), nil
}

// NewFileChooserButton:
func NewFileChooserButton(title string, action FileChooserAction) FileChooserButton {
	var _arg1 *C.gchar               // out
	var _arg2 C.GtkFileChooserAction // out
	var _cret *C.GtkWidget           // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkFileChooserAction(action)

	_cret = C.gtk_file_chooser_button_new(_arg1, _arg2)

	var _fileChooserButton FileChooserButton // out

	_fileChooserButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FileChooserButton)

	return _fileChooserButton
}

// NewFileChooserButtonWithDialog:
func NewFileChooserButtonWithDialog(dialog Dialog) FileChooserButton {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_file_chooser_button_new_with_dialog(_arg1)

	var _fileChooserButton FileChooserButton // out

	_fileChooserButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(FileChooserButton)

	return _fileChooserButton
}

func (b fileChooserButton) FocusOnClick() bool {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_file_chooser_button_get_focus_on_click(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (b fileChooserButton) Title() string {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret *C.gchar                // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_file_chooser_button_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b fileChooserButton) WidthChars() int {
	var _arg0 *C.GtkFileChooserButton // out
	var _cret C.gint                  // in

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_file_chooser_button_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (b fileChooserButton) SetFocusOnClickFileChooserButton(focusOnClick bool) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	if focusOnClick {
		_arg1 = C.TRUE
	}

	C.gtk_file_chooser_button_set_focus_on_click(_arg0, _arg1)
}

func (b fileChooserButton) SetTitleFileChooserButton(title string) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_file_chooser_button_set_title(_arg0, _arg1)
}

func (b fileChooserButton) SetWidthCharsFileChooserButton(nChars int) {
	var _arg0 *C.GtkFileChooserButton // out
	var _arg1 C.gint                  // out

	_arg0 = (*C.GtkFileChooserButton)(unsafe.Pointer(b.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_file_chooser_button_set_width_chars(_arg0, _arg1)
}

func (b fileChooserButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b fileChooserButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b fileChooserButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b fileChooserButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b fileChooserButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b fileChooserButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b fileChooserButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o fileChooserButton) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o fileChooserButton) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}

func (c fileChooserButton) AddChoice(id string, label string, options []string, optionLabels []string) {
	WrapFileChooser(gextras.InternObject(c)).AddChoice(id, label, options, optionLabels)
}

func (c fileChooserButton) AddFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).AddFilter(filter)
}

func (c fileChooserButton) AddShortcutFolder(folder string) error {
	return WrapFileChooser(gextras.InternObject(c)).AddShortcutFolder(folder)
}

func (c fileChooserButton) AddShortcutFolderURI(uri string) error {
	return WrapFileChooser(gextras.InternObject(c)).AddShortcutFolderURI(uri)
}

func (c fileChooserButton) Action() FileChooserAction {
	return WrapFileChooser(gextras.InternObject(c)).Action()
}

func (c fileChooserButton) Choice(id string) string {
	return WrapFileChooser(gextras.InternObject(c)).Choice(id)
}

func (c fileChooserButton) CreateFolders() bool {
	return WrapFileChooser(gextras.InternObject(c)).CreateFolders()
}

func (c fileChooserButton) CurrentFolder() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentFolder()
}

func (c fileChooserButton) CurrentFolderFile() gio.File {
	return WrapFileChooser(gextras.InternObject(c)).CurrentFolderFile()
}

func (c fileChooserButton) CurrentFolderURI() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentFolderURI()
}

func (c fileChooserButton) CurrentName() string {
	return WrapFileChooser(gextras.InternObject(c)).CurrentName()
}

func (c fileChooserButton) DoOverwriteConfirmation() bool {
	return WrapFileChooser(gextras.InternObject(c)).DoOverwriteConfirmation()
}

func (c fileChooserButton) ExtraWidget() Widget {
	return WrapFileChooser(gextras.InternObject(c)).ExtraWidget()
}

func (c fileChooserButton) File() gio.File {
	return WrapFileChooser(gextras.InternObject(c)).File()
}

func (c fileChooserButton) Filename() string {
	return WrapFileChooser(gextras.InternObject(c)).Filename()
}

func (c fileChooserButton) Filter() FileFilter {
	return WrapFileChooser(gextras.InternObject(c)).Filter()
}

func (c fileChooserButton) LocalOnly() bool {
	return WrapFileChooser(gextras.InternObject(c)).LocalOnly()
}

func (c fileChooserButton) PreviewFile() gio.File {
	return WrapFileChooser(gextras.InternObject(c)).PreviewFile()
}

func (c fileChooserButton) PreviewFilename() string {
	return WrapFileChooser(gextras.InternObject(c)).PreviewFilename()
}

func (c fileChooserButton) PreviewURI() string {
	return WrapFileChooser(gextras.InternObject(c)).PreviewURI()
}

func (c fileChooserButton) PreviewWidget() Widget {
	return WrapFileChooser(gextras.InternObject(c)).PreviewWidget()
}

func (c fileChooserButton) PreviewWidgetActive() bool {
	return WrapFileChooser(gextras.InternObject(c)).PreviewWidgetActive()
}

func (c fileChooserButton) SelectMultiple() bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectMultiple()
}

func (c fileChooserButton) ShowHidden() bool {
	return WrapFileChooser(gextras.InternObject(c)).ShowHidden()
}

func (c fileChooserButton) URI() string {
	return WrapFileChooser(gextras.InternObject(c)).URI()
}

func (c fileChooserButton) UsePreviewLabel() bool {
	return WrapFileChooser(gextras.InternObject(c)).UsePreviewLabel()
}

func (c fileChooserButton) RemoveChoice(id string) {
	WrapFileChooser(gextras.InternObject(c)).RemoveChoice(id)
}

func (c fileChooserButton) RemoveFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).RemoveFilter(filter)
}

func (c fileChooserButton) RemoveShortcutFolder(folder string) error {
	return WrapFileChooser(gextras.InternObject(c)).RemoveShortcutFolder(folder)
}

func (c fileChooserButton) RemoveShortcutFolderURI(uri string) error {
	return WrapFileChooser(gextras.InternObject(c)).RemoveShortcutFolderURI(uri)
}

func (c fileChooserButton) SelectAll() {
	WrapFileChooser(gextras.InternObject(c)).SelectAll()
}

func (c fileChooserButton) SelectFile(file gio.File) error {
	return WrapFileChooser(gextras.InternObject(c)).SelectFile(file)
}

func (c fileChooserButton) SelectFilename(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectFilename(filename)
}

func (c fileChooserButton) SelectURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SelectURI(uri)
}

func (c fileChooserButton) SetAction(action FileChooserAction) {
	WrapFileChooser(gextras.InternObject(c)).SetAction(action)
}

func (c fileChooserButton) SetChoice(id string, option string) {
	WrapFileChooser(gextras.InternObject(c)).SetChoice(id, option)
}

func (c fileChooserButton) SetCreateFolders(createFolders bool) {
	WrapFileChooser(gextras.InternObject(c)).SetCreateFolders(createFolders)
}

func (c fileChooserButton) SetCurrentFolder(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetCurrentFolder(filename)
}

func (c fileChooserButton) SetCurrentFolderFile(file gio.File) error {
	return WrapFileChooser(gextras.InternObject(c)).SetCurrentFolderFile(file)
}

func (c fileChooserButton) SetCurrentFolderURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetCurrentFolderURI(uri)
}

func (c fileChooserButton) SetCurrentName(name string) {
	WrapFileChooser(gextras.InternObject(c)).SetCurrentName(name)
}

func (c fileChooserButton) SetDoOverwriteConfirmation(doOverwriteConfirmation bool) {
	WrapFileChooser(gextras.InternObject(c)).SetDoOverwriteConfirmation(doOverwriteConfirmation)
}

func (c fileChooserButton) SetExtraWidget(extraWidget Widget) {
	WrapFileChooser(gextras.InternObject(c)).SetExtraWidget(extraWidget)
}

func (c fileChooserButton) SetFile(file gio.File) error {
	return WrapFileChooser(gextras.InternObject(c)).SetFile(file)
}

func (c fileChooserButton) SetFilename(filename string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetFilename(filename)
}

func (c fileChooserButton) SetFilter(filter FileFilter) {
	WrapFileChooser(gextras.InternObject(c)).SetFilter(filter)
}

func (c fileChooserButton) SetLocalOnly(localOnly bool) {
	WrapFileChooser(gextras.InternObject(c)).SetLocalOnly(localOnly)
}

func (c fileChooserButton) SetPreviewWidget(previewWidget Widget) {
	WrapFileChooser(gextras.InternObject(c)).SetPreviewWidget(previewWidget)
}

func (c fileChooserButton) SetPreviewWidgetActive(active bool) {
	WrapFileChooser(gextras.InternObject(c)).SetPreviewWidgetActive(active)
}

func (c fileChooserButton) SetSelectMultiple(selectMultiple bool) {
	WrapFileChooser(gextras.InternObject(c)).SetSelectMultiple(selectMultiple)
}

func (c fileChooserButton) SetShowHidden(showHidden bool) {
	WrapFileChooser(gextras.InternObject(c)).SetShowHidden(showHidden)
}

func (c fileChooserButton) SetURI(uri string) bool {
	return WrapFileChooser(gextras.InternObject(c)).SetURI(uri)
}

func (c fileChooserButton) SetUsePreviewLabel(useLabel bool) {
	WrapFileChooser(gextras.InternObject(c)).SetUsePreviewLabel(useLabel)
}

func (c fileChooserButton) UnselectAll() {
	WrapFileChooser(gextras.InternObject(c)).UnselectAll()
}

func (c fileChooserButton) UnselectFile(file gio.File) {
	WrapFileChooser(gextras.InternObject(c)).UnselectFile(file)
}

func (c fileChooserButton) UnselectFilename(filename string) {
	WrapFileChooser(gextras.InternObject(c)).UnselectFilename(filename)
}

func (c fileChooserButton) UnselectURI(uri string) {
	WrapFileChooser(gextras.InternObject(c)).UnselectURI(uri)
}
