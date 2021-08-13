// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_action_map_get_type()), F: marshalActionMapper},
	})
}

// ActionMapOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionMapOverrider interface {
	// AddAction adds an action to the action_map.
	//
	// If the action map already contains an action with the same name as action
	// then the old action is dropped from the action map.
	//
	// The action map takes its own reference on action.
	AddAction(action Actioner)
	// LookupAction looks up the action with the name action_name in action_map.
	//
	// If no such action exists, returns NULL.
	LookupAction(actionName string) Actioner
	// RemoveAction removes the named action from the action map.
	//
	// If no action of this name is in the map then nothing happens.
	RemoveAction(actionName string)
}

// ActionMap interface is implemented by Group implementations that operate by
// containing a number of named #GAction instances, such as ActionGroup.
//
// One useful application of this interface is to map the names of actions from
// various action groups to unique, prefixed names (e.g. by prepending "app." or
// "win."). This is the motivation for the 'Map' part of the interface name.
type ActionMap struct {
	*externglib.Object
}

// ActionMapper describes ActionMap's abstract methods.
type ActionMapper interface {
	externglib.Objector

	// AddAction adds an action to the action_map.
	AddAction(action Actioner)
	// AddActionEntries: convenience function for creating multiple Action
	// instances and adding them to a Map.
	AddActionEntries(entries []ActionEntry, userData cgo.Handle)
	// LookupAction looks up the action with the name action_name in action_map.
	LookupAction(actionName string) Actioner
	// RemoveAction removes the named action from the action map.
	RemoveAction(actionName string)
}

var _ ActionMapper = (*ActionMap)(nil)

func wrapActionMap(obj *externglib.Object) *ActionMap {
	return &ActionMap{
		Object: obj,
	}
}

func marshalActionMapper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapActionMap(obj), nil
}

// AddAction adds an action to the action_map.
//
// If the action map already contains an action with the same name as action
// then the old action is dropped from the action map.
//
// The action map takes its own reference on action.
func (actionMap *ActionMap) AddAction(action Actioner) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.GAction    // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(actionMap.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_action_map_add_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(action)
}

// AddActionEntries: convenience function for creating multiple Action instances
// and adding them to a Map.
//
// Each action is constructed as per one Entry.
//
//    static void
//    activate_quit (GSimpleAction *simple,
//                   GVariant      *parameter,
//                   gpointer       user_data)
//    {
//      exit (0);
//    }
//
//    static void
//    activate_print_string (GSimpleAction *simple,
//                           GVariant      *parameter,
//                           gpointer       user_data)
//    {
//      g_print ("s\n", g_variant_get_string (parameter, NULL));
//    }
//
//    static GActionGroup *
//    create_action_group (void)
//    {
//      const GActionEntry entries[] = {
//        { "quit",         activate_quit              },
//        { "print-string", activate_print_string, "s" }
//      };
//      GSimpleActionGroup *group;
//
//      group = g_simple_action_group_new ();
//      g_action_map_add_action_entries (G_ACTION_MAP (group), entries, G_N_ELEMENTS (entries), NULL);
//
//      return G_ACTION_GROUP (group);
//    }
func (actionMap *ActionMap) AddActionEntries(entries []ActionEntry, userData cgo.Handle) {
	var _arg0 *C.GActionMap   // out
	var _arg1 *C.GActionEntry // out
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(actionMap.Native()))
	_arg2 = (C.gint)(len(entries))
	if len(entries) > 0 {
		_arg1 = (*C.GActionEntry)(unsafe.Pointer(&entries[0]))
	}
	_arg3 = (C.gpointer)(unsafe.Pointer(userData))

	C.g_action_map_add_action_entries(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(entries)
	runtime.KeepAlive(userData)
}

// LookupAction looks up the action with the name action_name in action_map.
//
// If no such action exists, returns NULL.
func (actionMap *ActionMap) LookupAction(actionName string) Actioner {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out
	var _cret *C.GAction    // in

	_arg0 = (*C.GActionMap)(unsafe.Pointer(actionMap.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_action_map_lookup_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)

	var _action Actioner // out

	if _cret != nil {
		_action = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Actioner)
	}

	return _action
}

// RemoveAction removes the named action from the action map.
//
// If no action of this name is in the map then nothing happens.
func (actionMap *ActionMap) RemoveAction(actionName string) {
	var _arg0 *C.GActionMap // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GActionMap)(unsafe.Pointer(actionMap.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(actionName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_action_map_remove_action(_arg0, _arg1)
	runtime.KeepAlive(actionMap)
	runtime.KeepAlive(actionName)
}

// ActionEntry: this struct defines a single action. It is for use with
// g_action_map_add_action_entries().
//
// The order of the items in the structure are intended to reflect frequency of
// use. It is permissible to use an incomplete initialiser in order to leave
// some of the later values as NULL. All values after name are optional.
// Additional optional fields may be added in the future.
//
// See g_action_map_add_action_entries() for an example.
type ActionEntry struct {
	nocopy gextras.NoCopy
	native *C.GActionEntry
}

// Name: name of the action
func (a *ActionEntry) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.name)))
	return v
}

// ParameterType: type of the parameter that must be passed to the activate
// function for this action, given as a single GVariant type string (or NULL for
// no parameter)
func (a *ActionEntry) ParameterType() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.parameter_type)))
	return v
}

// State: initial state for this action, given in [GVariant text
// format][gvariant-text]. The state is parsed with no extra type information,
// so type tags must be added to the string if they are necessary. Stateless
// actions should give NULL here.
func (a *ActionEntry) State() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(a.native.state)))
	return v
}
