// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_application_get_type()), F: marshalApplication},
	})
}

// Application: a #GApplication is the foundation of an application. It wraps
// some low-level platform-specific services and is intended to act as the
// foundation for higher-level application classes such as Application or
// Application. In general, you should not use this class outside of a higher
// level framework.
//
// GApplication provides convenient life cycle management by maintaining a "use
// count" for the primary application instance. The use count can be changed
// using g_application_hold() and g_application_release(). If it drops to zero,
// the application exits. Higher-level classes such as Application employ the
// use count to ensure that the application stays alive as long as it has any
// opened windows.
//
// Another feature that GApplication (optionally) provides is process
// uniqueness. Applications can make use of this functionality by providing a
// unique application ID. If given, only one application with this ID can be
// running at a time per session. The session concept is platform-dependent, but
// corresponds roughly to a graphical desktop login. When your application is
// launched again, its arguments are passed through platform communication to
// the already running program. The already running instance of the program is
// called the "primary instance"; for non-unique applications this is always the
// current instance. On Linux, the D-Bus session bus is used for communication.
//
// The use of #GApplication differs from some other commonly-used uniqueness
// libraries (such as libunique) in important ways. The application is not
// expected to manually register itself and check if it is the primary instance.
// Instead, the main() function of a #GApplication should do very little more
// than instantiating the application instance, possibly connecting signal
// handlers, then calling g_application_run(). All checks for uniqueness are
// done internally. If the application is the primary instance then the startup
// signal is emitted and the mainloop runs. If the application is not the
// primary instance then a signal is sent to the primary instance and
// g_application_run() promptly returns. See the code examples below.
//
// If used, the expected form of an application identifier is the same as that
// of of a D-Bus well-known bus name
// (https://dbus.freedesktop.org/doc/dbus-specification.html#message-protocol-names-bus).
// Examples include: `com.example.MyApp`,
// `org.example.internal_apps.Calculator`, `org._7_zip.Archiver`. For details on
// valid application identifiers, see g_application_id_is_valid().
//
// On Linux, the application identifier is claimed as a well-known bus name on
// the user's session bus. This means that the uniqueness of your application is
// scoped to the current session. It also means that your application may
// provide additional services (through registration of other object paths) at
// that bus name. The registration of these object paths should be done with the
// shared GDBus session bus. Note that due to the internal architecture of
// GDBus, method calls can be dispatched at any time (even if a main loop is not
// running). For this reason, you must ensure that any object paths that you
// wish to register are registered before #GApplication attempts to acquire the
// bus name of your application (which happens in g_application_register()).
// Unfortunately, this means that you cannot use g_application_get_is_remote()
// to decide if you want to register object paths.
//
// GApplication also implements the Group and Map interfaces and lets you easily
// export actions by adding them with g_action_map_add_action(). When invoking
// an action by calling g_action_group_activate_action() on the application, it
// is always invoked in the primary instance. The actions are also exported on
// the session bus, and GIO provides the BusActionGroup wrapper to conveniently
// access them remotely. GIO provides a BusMenuModel wrapper for remote access
// to exported Models.
//
// There is a number of different entry points into a GApplication:
//
// - via 'Activate' (i.e. just starting the application)
//
// - via 'Open' (i.e. opening some files)
//
// - by handling a command-line
//
// - via activating an action
//
// The #GApplication::startup signal lets you handle the application
// initialization for all of these in a single place.
//
// Regardless of which of these entry points is used to start the application,
// GApplication passes some ‘platform data’ from the launching instance to the
// primary instance, in the form of a #GVariant dictionary mapping strings to
// variants. To use platform data, override the @before_emit or @after_emit
// virtual functions in your #GApplication subclass. When dealing with
// CommandLine objects, the platform data is directly available via
// g_application_command_line_get_cwd(),
// g_application_command_line_get_environ() and
// g_application_command_line_get_platform_data().
//
// As the name indicates, the platform data may vary depending on the operating
// system, but it always includes the current directory (key "cwd"), and
// optionally the environment (ie the set of environment variables and their
// values) of the calling process (key "environ"). The environment is only added
// to the platform data if the G_APPLICATION_SEND_ENVIRONMENT flag is set.
// #GApplication subclasses can add their own platform data by overriding the
// @add_platform_data virtual function. For instance, Application adds startup
// notification data in this way.
//
// To parse commandline arguments you may handle the #GApplication::command-line
// signal or override the local_command_line() vfunc, to parse them in either
// the primary instance or the local instance, respectively.
//
// For an example of opening files with a GApplication, see
// gapplication-example-open.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-open.c).
//
// For an example of using actions with GApplication, see
// gapplication-example-actions.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-actions.c).
//
// For an example of using extra D-Bus hooks with GApplication, see
// gapplication-example-dbushooks.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-dbushooks.c).
type Application interface {
	ActionGroup
	ActionMap

	// ActivateApplication:
	ActivateApplication()
	// AddMainOptionApplication:
	AddMainOptionApplication(longName string, shortName byte, flags glib.OptionFlags, arg glib.OptionArg, description string, argDescription string)
	// AddMainOptionEntriesApplication:
	AddMainOptionEntriesApplication(entries []glib.OptionEntry)
	// AddOptionGroupApplication:
	AddOptionGroupApplication(group *glib.OptionGroup)
	// BindBusyPropertyApplication:
	BindBusyPropertyApplication(object gextras.Objector, property string)
	// ApplicationID:
	ApplicationID() string
	// DBusConnection:
	DBusConnection() DBusConnection
	// DBusObjectPath:
	DBusObjectPath() string
	// Flags:
	Flags() ApplicationFlags
	// InactivityTimeout:
	InactivityTimeout() uint
	// IsBusy:
	IsBusy() bool
	// IsRegistered:
	IsRegistered() bool
	// IsRemote:
	IsRemote() bool
	// ResourceBasePath:
	ResourceBasePath() string
	// HoldApplication:
	HoldApplication()
	// MarkBusyApplication:
	MarkBusyApplication()
	// OpenApplication:
	OpenApplication(files []File, hint string)
	// QuitApplication:
	QuitApplication()
	// RegisterApplication:
	RegisterApplication(cancellable Cancellable) error
	// ReleaseApplication:
	ReleaseApplication()
	// RunApplication:
	RunApplication(argv []string) int
	// SendNotificationApplication:
	SendNotificationApplication(id string, notification Notification)
	// SetActionGroupApplication:
	SetActionGroupApplication(actionGroup ActionGroup)
	// SetApplicationIDApplication:
	SetApplicationIDApplication(applicationId string)
	// SetDefaultApplication:
	SetDefaultApplication()
	// SetFlagsApplication:
	SetFlagsApplication(flags ApplicationFlags)
	// SetInactivityTimeoutApplication:
	SetInactivityTimeoutApplication(inactivityTimeout uint)
	// SetOptionContextDescriptionApplication:
	SetOptionContextDescriptionApplication(description string)
	// SetOptionContextParameterStringApplication:
	SetOptionContextParameterStringApplication(parameterString string)
	// SetOptionContextSummaryApplication:
	SetOptionContextSummaryApplication(summary string)
	// SetResourceBasePathApplication:
	SetResourceBasePathApplication(resourcePath string)
	// UnbindBusyPropertyApplication:
	UnbindBusyPropertyApplication(object gextras.Objector, property string)
	// UnmarkBusyApplication:
	UnmarkBusyApplication()
	// WithdrawNotificationApplication:
	WithdrawNotificationApplication(id string)
}

// application implements the Application class.
type application struct {
	gextras.Objector
}

// WrapApplication wraps a GObject to the right type. It is
// primarily used internally.
func WrapApplication(obj *externglib.Object) Application {
	return application{
		Objector: obj,
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplication(obj), nil
}

// NewApplication:
func NewApplication(applicationId string, flags ApplicationFlags) Application {
	var _arg1 *C.gchar            // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.GApplication     // in

	_arg1 = (*C.gchar)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.g_application_new(_arg1, _arg2)

	var _application Application // out

	_application = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Application)

	return _application
}

func (a application) ActivateApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_activate(_arg0)
}

func (a application) AddMainOptionApplication(longName string, shortName byte, flags glib.OptionFlags, arg glib.OptionArg, description string, argDescription string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.char         // out
	var _arg2 C.char          // out
	var _arg3 C.GOptionFlags  // out
	var _arg4 C.GOptionArg    // out
	var _arg5 *C.char         // out
	var _arg6 *C.char         // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.char)(C.CString(longName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.char(shortName)
	_arg3 = C.GOptionFlags(flags)
	_arg4 = C.GOptionArg(arg)
	_arg5 = (*C.char)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (*C.char)(C.CString(argDescription))
	defer C.free(unsafe.Pointer(_arg6))

	C.g_application_add_main_option(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

func (a application) AddMainOptionEntriesApplication(entries []glib.OptionEntry) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.GOptionEntry

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	{
		var zero glib.OptionEntry
		entries = append(entries, zero)
	}
	_arg1 = (*C.GOptionEntry)(unsafe.Pointer(&entries[0]))

	C.g_application_add_main_option_entries(_arg0, _arg1)
}

func (a application) AddOptionGroupApplication(group *glib.OptionGroup) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.GOptionGroup // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GOptionGroup)(unsafe.Pointer(group.Native()))

	C.g_application_add_option_group(_arg0, _arg1)
}

func (a application) BindBusyPropertyApplication(object gextras.Objector, property string) {
	var _arg0 *C.GApplication // out
	var _arg1 C.gpointer      // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_application_bind_busy_property(_arg0, _arg1, _arg2)
}

func (a application) ApplicationID() string {
	var _arg0 *C.GApplication // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_application_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a application) DBusConnection() DBusConnection {
	var _arg0 *C.GApplication    // out
	var _cret *C.GDBusConnection // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_dbus_connection(_arg0)

	var _dBusConnection DBusConnection // out

	_dBusConnection = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DBusConnection)

	return _dBusConnection
}

func (a application) DBusObjectPath() string {
	var _arg0 *C.GApplication // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_dbus_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a application) Flags() ApplicationFlags {
	var _arg0 *C.GApplication     // out
	var _cret C.GApplicationFlags // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_flags(_arg0)

	var _applicationFlags ApplicationFlags // out

	_applicationFlags = ApplicationFlags(_cret)

	return _applicationFlags
}

func (a application) InactivityTimeout() uint {
	var _arg0 *C.GApplication // out
	var _cret C.guint         // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_inactivity_timeout(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (a application) IsBusy() bool {
	var _arg0 *C.GApplication // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_is_busy(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a application) IsRegistered() bool {
	var _arg0 *C.GApplication // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_is_registered(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a application) IsRemote() bool {
	var _arg0 *C.GApplication // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_is_remote(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a application) ResourceBasePath() string {
	var _arg0 *C.GApplication // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	_cret = C.g_application_get_resource_base_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a application) HoldApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_hold(_arg0)
}

func (a application) MarkBusyApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_mark_busy(_arg0)
}

func (a application) OpenApplication(files []File, hint string) {
	var _arg0 *C.GApplication // out
	var _arg1 **C.GFile
	var _arg2 C.gint
	var _arg3 *C.gchar // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg2 = C.gint(len(files))
	_arg1 = (**C.GFile)(C.malloc(C.ulong(len(files)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(files))
		for i := range files {
			out[i] = (*C.GFile)(unsafe.Pointer(files[i].Native()))
		}
	}
	_arg3 = (*C.gchar)(C.CString(hint))
	defer C.free(unsafe.Pointer(_arg3))

	C.g_application_open(_arg0, _arg1, _arg2, _arg3)
}

func (a application) QuitApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_quit(_arg0)
}

func (a application) RegisterApplication(cancellable Cancellable) error {
	var _arg0 *C.GApplication // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_application_register(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

func (a application) ReleaseApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_release(_arg0)
}

func (a application) RunApplication(argv []string) int {
	var _arg0 *C.GApplication // out
	var _arg2 **C.char
	var _arg1 C.int
	var _cret C.int // in

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.int(len(argv))
	_arg2 = (**C.char)(C.malloc(C.ulong(len(argv)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(argv))
		for i := range argv {
			out[i] = (*C.char)(C.CString(argv[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.g_application_run(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a application) SendNotificationApplication(id string, notification Notification) {
	var _arg0 *C.GApplication  // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GNotification // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GNotification)(unsafe.Pointer(notification.Native()))

	C.g_application_send_notification(_arg0, _arg1, _arg2)
}

func (a application) SetActionGroupApplication(actionGroup ActionGroup) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.GActionGroup // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GActionGroup)(unsafe.Pointer(actionGroup.Native()))

	C.g_application_set_action_group(_arg0, _arg1)
}

func (a application) SetApplicationIDApplication(applicationId string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(applicationId))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_set_application_id(_arg0, _arg1)
}

func (a application) SetDefaultApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_set_default(_arg0)
}

func (a application) SetFlagsApplication(flags ApplicationFlags) {
	var _arg0 *C.GApplication     // out
	var _arg1 C.GApplicationFlags // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.GApplicationFlags(flags)

	C.g_application_set_flags(_arg0, _arg1)
}

func (a application) SetInactivityTimeoutApplication(inactivityTimeout uint) {
	var _arg0 *C.GApplication // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(inactivityTimeout)

	C.g_application_set_inactivity_timeout(_arg0, _arg1)
}

func (a application) SetOptionContextDescriptionApplication(description string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_set_option_context_description(_arg0, _arg1)
}

func (a application) SetOptionContextParameterStringApplication(parameterString string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(parameterString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_set_option_context_parameter_string(_arg0, _arg1)
}

func (a application) SetOptionContextSummaryApplication(summary string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(summary))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_set_option_context_summary(_arg0, _arg1)
}

func (a application) SetResourceBasePathApplication(resourcePath string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_set_resource_base_path(_arg0, _arg1)
}

func (a application) UnbindBusyPropertyApplication(object gextras.Objector, property string) {
	var _arg0 *C.GApplication // out
	var _arg1 C.gpointer      // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_application_unbind_busy_property(_arg0, _arg1, _arg2)
}

func (a application) UnmarkBusyApplication() {
	var _arg0 *C.GApplication // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))

	C.g_application_unmark_busy(_arg0)
}

func (a application) WithdrawNotificationApplication(id string) {
	var _arg0 *C.GApplication // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GApplication)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.gchar)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_application_withdraw_notification(_arg0, _arg1)
}

func (a application) ActionAdded(actionName string) {
	WrapActionGroup(gextras.InternObject(a)).ActionAdded(actionName)
}

func (a application) ActionEnabledChanged(actionName string, enabled bool) {
	WrapActionGroup(gextras.InternObject(a)).ActionEnabledChanged(actionName, enabled)
}

func (a application) ActionRemoved(actionName string) {
	WrapActionGroup(gextras.InternObject(a)).ActionRemoved(actionName)
}

func (a application) ActionStateChanged(actionName string, state *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ActionStateChanged(actionName, state)
}

func (a application) ActivateAction(actionName string, parameter *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ActivateAction(actionName, parameter)
}

func (a application) ChangeActionState(actionName string, value *glib.Variant) {
	WrapActionGroup(gextras.InternObject(a)).ChangeActionState(actionName, value)
}

func (a application) ActionEnabled(actionName string) bool {
	return WrapActionGroup(gextras.InternObject(a)).ActionEnabled(actionName)
}

func (a application) ActionParameterType(actionName string) *glib.VariantType {
	return WrapActionGroup(gextras.InternObject(a)).ActionParameterType(actionName)
}

func (a application) ActionState(actionName string) *glib.Variant {
	return WrapActionGroup(gextras.InternObject(a)).ActionState(actionName)
}

func (a application) ActionStateHint(actionName string) *glib.Variant {
	return WrapActionGroup(gextras.InternObject(a)).ActionStateHint(actionName)
}

func (a application) ActionStateType(actionName string) *glib.VariantType {
	return WrapActionGroup(gextras.InternObject(a)).ActionStateType(actionName)
}

func (a application) HasAction(actionName string) bool {
	return WrapActionGroup(gextras.InternObject(a)).HasAction(actionName)
}

func (a application) ListActions() []string {
	return WrapActionGroup(gextras.InternObject(a)).ListActions()
}

func (a application) QueryAction(actionName string) (enabled bool, parameterType *glib.VariantType, stateType *glib.VariantType, stateHint *glib.Variant, state *glib.Variant, ok bool) {
	return WrapActionGroup(gextras.InternObject(a)).QueryAction(actionName)
}

func (a application) AddAction(action Action) {
	WrapActionMap(gextras.InternObject(a)).AddAction(action)
}

func (a application) LookupAction(actionName string) Action {
	return WrapActionMap(gextras.InternObject(a)).LookupAction(actionName)
}

func (a application) RemoveAction(actionName string) {
	WrapActionMap(gextras.InternObject(a)).RemoveAction(actionName)
}
