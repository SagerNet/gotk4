// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_action_group_get_type()), F: marshalActionGroup},
	})
}

// ActionGroup actions are organised into groups. An action group is essentially
// a map from names to Action objects.
//
// All actions that would make sense to use in a particular context should be in
// a single group. Multiple action groups may be used for a particular user
// interface. In fact, it is expected that most nontrivial applications will
// make use of multiple groups. For example, in an application that can edit
// multiple documents, one group holding global actions (e.g. quit, about, new),
// and one group per document holding actions that act on that document (eg.
// save, cut/copy/paste, etc). Each window’s menus would be constructed from a
// combination of two action groups.
//
//
// Accelerators
//
// Accelerators are handled by the GTK+ accelerator map. All actions are
// assigned an accelerator path (which normally has the form
// `<Actions>/group-name/action-name`) and a shortcut is associated with this
// accelerator path. All menuitems and toolitems take on this accelerator path.
// The GTK+ accelerator map code makes sure that the correct shortcut is
// displayed next to the menu item.
//
//
// GtkActionGroup as GtkBuildable
//
// The ActionGroup implementation of the Buildable interface accepts Action
// objects as <child> elements in UI definitions.
//
// Note that it is probably more common to define actions and action groups in
// the code, since they are directly related to what the code can do.
//
// The GtkActionGroup implementation of the GtkBuildable interface supports a
// custom <accelerator> element, which has attributes named “key“ and
// “modifiers“ and allows to specify accelerators. This is similar to the
// <accelerator> element of Widget, the main difference is that it doesn’t allow
// you to specify a signal.
//
// A Dialog UI definition fragment. ##
//
//    <object class="GtkActionGroup" id="actiongroup">
//      <child>
//          <object class="GtkAction" id="About">
//              <property name="name">About</property>
//              <property name="stock_id">gtk-about</property>
//              <signal handler="about_activate" name="activate"/>
//          </object>
//          <accelerator key="F1" modifiers="GDK_CONTROL_MASK | GDK_SHIFT_MASK"/>
//      </child>
//    </object>
type ActionGroup interface {
	Buildable

	// AddActionActionGroup:
	AddActionActionGroup(action Action)
	// AddActionWithAccelActionGroup:
	AddActionWithAccelActionGroup(action Action, accelerator string)
	// AccelGroup:
	AccelGroup() AccelGroup
	// Action:
	Action(actionName string) Action
	// GetName:
	GetName() string
	// Sensitive:
	Sensitive() bool
	// Visible:
	Visible() bool
	// RemoveActionActionGroup:
	RemoveActionActionGroup(action Action)
	// SetAccelGroupActionGroup:
	SetAccelGroupActionGroup(accelGroup AccelGroup)
	// SetSensitiveActionGroup:
	SetSensitiveActionGroup(sensitive bool)
	// SetTranslationDomainActionGroup:
	SetTranslationDomainActionGroup(domain string)
	// SetVisibleActionGroup:
	SetVisibleActionGroup(visible bool)
	// TranslateStringActionGroup:
	TranslateStringActionGroup(_string string) string
}

// actionGroup implements the ActionGroup class.
type actionGroup struct {
	gextras.Objector
}

// WrapActionGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapActionGroup(obj *externglib.Object) ActionGroup {
	return actionGroup{
		Objector: obj,
	}
}

func marshalActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActionGroup(obj), nil
}

// NewActionGroup:
func NewActionGroup(name string) ActionGroup {
	var _arg1 *C.gchar          // out
	var _cret *C.GtkActionGroup // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_new(_arg1)

	var _actionGroup ActionGroup // out

	_actionGroup = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(ActionGroup)

	return _actionGroup
}

func (a actionGroup) AddActionActionGroup(action Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_add_action(_arg0, _arg1)
}

func (a actionGroup) AddActionWithAccelActionGroup(action Action, accelerator string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out
	var _arg2 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg2 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_action_group_add_action_with_accel(_arg0, _arg1, _arg2)
}

func (a actionGroup) AccelGroup() AccelGroup {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.GtkAccelGroup  // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_group_get_accel_group(_arg0)

	var _accelGroup AccelGroup // out

	_accelGroup = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(AccelGroup)

	return _accelGroup
}

func (a actionGroup) Action(actionName string) Action {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.GtkAction      // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_get_action(_arg0, _arg1)

	var _action Action // out

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (a actionGroup) GetName() string {
	var _arg0 *C.GtkActionGroup // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_group_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a actionGroup) Sensitive() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_group_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a actionGroup) Visible() bool {
	var _arg0 *C.GtkActionGroup // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_group_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a actionGroup) RemoveActionActionGroup(action Action) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_group_remove_action(_arg0, _arg1)
}

func (a actionGroup) SetAccelGroupActionGroup(accelGroup AccelGroup) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.GtkAccelGroup  // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_action_group_set_accel_group(_arg0, _arg1)
}

func (a actionGroup) SetSensitiveActionGroup(sensitive bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_sensitive(_arg0, _arg1)
}

func (a actionGroup) SetTranslationDomainActionGroup(domain string) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_group_set_translation_domain(_arg0, _arg1)
}

func (a actionGroup) SetVisibleActionGroup(visible bool) {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_group_set_visible(_arg0, _arg1)
}

func (a actionGroup) TranslateStringActionGroup(_string string) string {
	var _arg0 *C.GtkActionGroup // out
	var _arg1 *C.gchar          // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkActionGroup)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(_string))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_action_group_translate_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (b actionGroup) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b actionGroup) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b actionGroup) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b actionGroup) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b actionGroup) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b actionGroup) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b actionGroup) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

// RadioActionEntry structs are used with gtk_action_group_add_radio_actions()
// to construct groups of radio actions.
//
// Deprecated: since version 3.10.
type RadioActionEntry C.GtkRadioActionEntry

// WrapRadioActionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRadioActionEntry(ptr unsafe.Pointer) *RadioActionEntry {
	return (*RadioActionEntry)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RadioActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}
