// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.g_settings_bind_flags_get_type()), F: marshalSettingsBindFlags},
		{T: externglib.Type(C.g_settings_get_type()), F: marshalSettings},
	})
}

// SettingsBindFlags flags used when creating a binding. These flags determine
// in which direction the binding works. The default is to synchronize in both
// directions.
type SettingsBindFlags int

const (
	// SettingsBindFlagsDefault: equivalent to
	// `G_SETTINGS_BIND_GET|G_SETTINGS_BIND_SET`
	SettingsBindFlagsDefault SettingsBindFlags = 0b0
	// SettingsBindFlagsGet: update the #GObject property when the setting
	// changes. It is an error to use this flag if the property is not writable.
	SettingsBindFlagsGet SettingsBindFlags = 0b1
	// SettingsBindFlagsSet: update the setting when the #GObject property
	// changes. It is an error to use this flag if the property is not readable.
	SettingsBindFlagsSet SettingsBindFlags = 0b10
	// SettingsBindFlagsNoSensitivity: do not try to bind a "sensitivity"
	// property to the writability of the setting
	SettingsBindFlagsNoSensitivity SettingsBindFlags = 0b100
	// SettingsBindFlagsGetNoChanges: when set in addition to SETTINGS_BIND_GET,
	// set the #GObject property value initially from the setting, but do not
	// listen for changes of the setting
	SettingsBindFlagsGetNoChanges SettingsBindFlags = 0b1000
	// SettingsBindFlagsInvertBoolean: when passed to g_settings_bind(), uses a
	// pair of mapping functions that invert the boolean value when mapping
	// between the setting and the property. The setting and property must both
	// be booleans. You cannot pass this flag to g_settings_bind_with_mapping().
	SettingsBindFlagsInvertBoolean SettingsBindFlags = 0b10000
)

func marshalSettingsBindFlags(p uintptr) (interface{}, error) {
	return SettingsBindFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Settings: the #GSettings class provides a convenient API for storing and
// retrieving application settings.
//
// Reads and writes can be considered to be non-blocking. Reading settings with
// #GSettings is typically extremely fast: on approximately the same order of
// magnitude (but slower than) a Table lookup. Writing settings is also
// extremely fast in terms of time to return to your application, but can be
// extremely expensive for other threads and other processes. Many settings
// backends (including dconf) have lazy initialisation which means in the common
// case of the user using their computer without modifying any settings a lot of
// work can be avoided. For dconf, the D-Bus service doesn't even need to be
// started in this case. For this reason, you should only ever modify #GSettings
// keys in response to explicit user action. Particular care should be paid to
// ensure that modifications are not made during startup -- for example, when
// setting the initial value of preferences widgets. The built-in
// g_settings_bind() functionality is careful not to write settings in response
// to notify signals as a result of modifications that it makes to widgets.
//
// When creating a GSettings instance, you have to specify a schema that
// describes the keys in your settings and their types and default values, as
// well as some other information.
//
// Normally, a schema has a fixed path that determines where the settings are
// stored in the conceptual global tree of settings. However, schemas can also
// be '[relocatable][gsettings-relocatable]', i.e. not equipped with a fixed
// path. This is useful e.g. when the schema describes an 'account', and you
// want to be able to store a arbitrary number of accounts.
//
// Paths must start with and end with a forward slash character ('/') and must
// not contain two sequential slash characters. Paths should be chosen based on
// a domain name associated with the program or library to which the settings
// belong. Examples of paths are "/org/gtk/settings/file-chooser/" and
// "/ca/desrt/dconf-editor/". Paths should not start with "/apps/", "/desktop/"
// or "/system/" as they often did in GConf.
//
// Unlike other configuration systems (like GConf), GSettings does not restrict
// keys to basic types like strings and numbers. GSettings stores values as
// #GVariant, and allows any Type for keys. Key names are restricted to
// lowercase characters, numbers and '-'. Furthermore, the names must begin with
// a lowercase character, must not end with a '-', and must not contain
// consecutive dashes.
//
// Similar to GConf, the default values in GSettings schemas can be localized,
// but the localized values are stored in gettext catalogs and looked up with
// the domain that is specified in the `gettext-domain` attribute of the
// <schemalist> or <schema> elements and the category that is specified in the
// `l10n` attribute of the <default> element. The string which is translated
// includes all text in the <default> element, including any surrounding
// quotation marks.
//
// The `l10n` attribute must be set to `messages` or `time`, and sets the
// [locale category for
// translation](https://www.gnu.org/software/gettext/manual/html_node/Aspects.html#index-locale-categories-1).
// The `messages` category should be used by default; use `time` for
// translatable date or time formats. A translation comment can be added as an
// XML comment immediately above the <default> element — it is recommended to
// add these comments to aid translators understand the meaning and implications
// of the default value. An optional translation `context` attribute can be set
// on the <default> element to disambiguate multiple defaults which use the same
// string.
//
// For example:
//
//    <!-- Translators: A list of words which are not allowed to be typed, in
//         GVariant serialization syntax.
//         See: https://developer.gnome.org/glib/stable/gvariant-text.html -->
//    <default l10n='messages' context='Banned words'>['bad', 'words']</default>
//
// Translations of default values must remain syntactically valid serialized
// #GVariants (e.g. retaining any surrounding quotation marks) or runtime errors
// will occur.
//
// GSettings uses schemas in a compact binary form that is created by the
// [glib-compile-schemas][glib-compile-schemas] utility. The input is a schema
// description in an XML format.
//
// A DTD for the gschema XML format can be found here: gschema.dtd
// (https://git.gnome.org/browse/glib/tree/gio/gschema.dtd)
//
// The [glib-compile-schemas][glib-compile-schemas] tool expects schema files to
// have the extension `.gschema.xml`.
//
// At runtime, schemas are identified by their id (as specified in the id
// attribute of the <schema> element). The convention for schema ids is to use a
// dotted name, similar in style to a D-Bus bus name, e.g.
// "org.gnome.SessionManager". In particular, if the settings are for a specific
// service that owns a D-Bus bus name, the D-Bus bus name and schema id should
// match. For schemas which deal with settings not associated with one named
// application, the id should not use StudlyCaps, e.g.
// "org.gnome.font-rendering".
//
// In addition to #GVariant types, keys can have types that have enumerated
// types. These can be described by a <choice>, <enum> or <flags> element, as
// seen in the [example][schema-enumerated]. The underlying type of such a key
// is string, but you can use g_settings_get_enum(), g_settings_set_enum(),
// g_settings_get_flags(), g_settings_set_flags() access the numeric values
// corresponding to the string value of enum and flags keys.
//
// An example for default value:
//
//    <schemalist>
//      <schema id="org.gtk.Test" path="/org/gtk/Test/" gettext-domain="test">
//
//        <key name="greeting" type="s">
//          <default l10n="messages">"Hello, earthlings"</default>
//          <summary>A greeting</summary>
//          <description>
//            Greeting of the invading martians
//          </description>
//        </key>
//
//        <key name="box" type="(ii)">
//          <default>(20,30)</default>
//        </key>
//
//        <key name="empty-string" type="s">
//          <default>""</default>
//          <summary>Empty strings have to be provided in GVariant form</summary>
//        </key>
//
//      </schema>
//    </schemalist>
//
// An example for ranges, choices and enumerated types:
//
//    <schemalist>
//
//      <enum id="org.gtk.Test.myenum">
//        <value nick="first" value="1"/>
//        <value nick="second" value="2"/>
//      </enum>
//
//      <flags id="org.gtk.Test.myflags">
//        <value nick="flag1" value="1"/>
//        <value nick="flag2" value="2"/>
//        <value nick="flag3" value="4"/>
//      </flags>
//
//      <schema id="org.gtk.Test">
//
//        <key name="key-with-range" type="i">
//          <range min="1" max="100"/>
//          <default>10</default>
//        </key>
//
//        <key name="key-with-choices" type="s">
//          <choices>
//            <choice value='Elisabeth'/>
//            <choice value='Annabeth'/>
//            <choice value='Joe'/>
//          </choices>
//          <aliases>
//            <alias value='Anna' target='Annabeth'/>
//            <alias value='Beth' target='Elisabeth'/>
//          </aliases>
//          <default>'Joe'</default>
//        </key>
//
//        <key name='enumerated-key' enum='org.gtk.Test.myenum'>
//          <default>'first'</default>
//        </key>
//
//        <key name='flags-key' flags='org.gtk.Test.myflags'>
//          <default>["flag1","flag2"]</default>
//        </key>
//      </schema>
//    </schemalist>
//
//
// Vendor overrides
//
// Default values are defined in the schemas that get installed by an
// application. Sometimes, it is necessary for a vendor or distributor to adjust
// these defaults. Since patching the XML source for the schema is inconvenient
// and error-prone, [glib-compile-schemas][glib-compile-schemas] reads so-called
// vendor override' files. These are keyfiles in the same directory as the XML
// schema sources which can override default values. The schema id serves as the
// group name in the key file, and the values are expected in serialized
// GVariant form, as in the following example:
//
//    [org.gtk.Example]
//    key1='string'
//    key2=1.5
//
// glib-compile-schemas expects schema files to have the extension
// `.gschema.override`.
//
//
// Binding
//
// A very convenient feature of GSettings lets you bind #GObject properties
// directly to settings, using g_settings_bind(). Once a GObject property has
// been bound to a setting, changes on either side are automatically propagated
// to the other side. GSettings handles details like mapping between GObject and
// GVariant types, and preventing infinite cycles.
//
// This makes it very easy to hook up a preferences dialog to the underlying
// settings. To make this even more convenient, GSettings looks for a boolean
// property with the name "sensitivity" and automatically binds it to the
// writability of the bound setting. If this 'magic' gets in the way, it can be
// suppressed with the SETTINGS_BIND_NO_SENSITIVITY flag.
//
//
// Relocatable schemas
//
// A relocatable schema is one with no `path` attribute specified on its
// <schema> element. By using g_settings_new_with_path(), a #GSettings object
// can be instantiated for a relocatable schema, assigning a path to the
// instance. Paths passed to g_settings_new_with_path() will typically be
// constructed dynamically from a constant prefix plus some form of instance
// identifier; but they must still be valid GSettings paths. Paths could also be
// constant and used with a globally installed schema originating from a
// dependency library.
//
// For example, a relocatable schema could be used to store geometry information
// for different windows in an application. If the schema ID was
// `org.foo.MyApp.Window`, it could be instantiated for paths
// `/org/foo/MyApp/main/`, `/org/foo/MyApp/document-1/`,
// `/org/foo/MyApp/document-2/`, etc. If any of the paths are well-known they
// can be specified as <child> elements in the parent schema, e.g.:
//
//    <schema id="org.foo.MyApp" path="/org/foo/MyApp/">
//      <child name="main" schema="org.foo.MyApp.Window"/>
//    </schema>
//
//
// Build system integration
//
// GSettings comes with autotools integration to simplify compiling and
// installing schemas. To add GSettings support to an application, add the
// following to your `configure.ac`:
//
//    GLIB_GSETTINGS
//
// In the appropriate `Makefile.am`, use the following snippet to compile and
// install the named schema:
//
//    gsettings_SCHEMAS = org.foo.MyApp.gschema.xml
//    EXTRA_DIST = $(gsettings_SCHEMAS)
//
//    @GSETTINGS_RULES@
//
// No changes are needed to the build system to mark a schema XML file for
// translation. Assuming it sets the `gettext-domain` attribute, a schema may be
// marked for translation by adding it to `POTFILES.in`, assuming gettext 0.19
// is in use (the preferred method for translation):
//
//    data/org.foo.MyApp.gschema.xml
//
// Alternatively, if intltool 0.50.1 is in use:
//
//    [type: gettext/gsettings]data/org.foo.MyApp.gschema.xml
//
// GSettings will use gettext to look up translations for the <summary> and
// <description> elements, and also any <default> elements which have a `l10n`
// attribute set. Translations must not be included in the `.gschema.xml` file
// by the build system, for example by using intltool XML rules with a
// `.gschema.xml.in` template.
//
// If an enumerated type defined in a C header file is to be used in a GSettings
// schema, it can either be defined manually using an <enum> element in the
// schema XML, or it can be extracted automatically from the C header. This
// approach is preferred, as it ensures the two representations are always
// synchronised. To do so, add the following to the relevant `Makefile.am`:
//
//    gsettings_ENUM_NAMESPACE = org.foo.MyApp
//    gsettings_ENUM_FILES = my-app-enums.h my-app-misc.h
//
// `gsettings_ENUM_NAMESPACE` specifies the schema namespace for the enum files,
// which are specified in `gsettings_ENUM_FILES`. This will generate a
// `org.foo.MyApp.enums.xml` file containing the extracted enums, which will be
// automatically included in the schema compilation, install and uninstall
// rules. It should not be committed to version control or included in
// `EXTRA_DIST`.
type Settings interface {
	gextras.Objector

	// ApplySettings:
	ApplySettings()
	// BindSettings:
	BindSettings(key string, object gextras.Objector, property string, flags SettingsBindFlags)
	// BindWritableSettings:
	BindWritableSettings(key string, object gextras.Objector, property string, inverted bool)
	// CreateActionSettings:
	CreateActionSettings(key string) Action
	// DelaySettings:
	DelaySettings()
	// Boolean:
	Boolean(key string) bool
	// Child:
	Child(name string) Settings
	// DefaultValue:
	DefaultValue(key string) *glib.Variant
	// Double:
	Double(key string) float64
	// Enum:
	Enum(key string) int
	// Flags:
	Flags(key string) uint
	// HasUnapplied:
	HasUnapplied() bool
	// Int:
	Int(key string) int
	// Int64:
	Int64(key string) int64
	// Range:
	Range(key string) *glib.Variant
	// String:
	String(key string) string
	// Strv:
	Strv(key string) []string
	// Uint:
	Uint(key string) uint
	// Uint64:
	Uint64(key string) uint64
	// UserValue:
	UserValue(key string) *glib.Variant
	// Value:
	Value(key string) *glib.Variant
	// IsWritableSettings:
	IsWritableSettings(name string) bool
	// ListChildrenSettings:
	ListChildrenSettings() []string
	// ListKeysSettings:
	ListKeysSettings() []string
	// RangeCheckSettings:
	RangeCheckSettings(key string, value *glib.Variant) bool
	// ResetSettings:
	ResetSettings(key string)
	// RevertSettings:
	RevertSettings()
	// SetBooleanSettings:
	SetBooleanSettings(key string, value bool) bool
	// SetDoubleSettings:
	SetDoubleSettings(key string, value float64) bool
	// SetEnumSettings:
	SetEnumSettings(key string, value int) bool
	// SetFlagsSettings:
	SetFlagsSettings(key string, value uint) bool
	// SetIntSettings:
	SetIntSettings(key string, value int) bool
	// SetInt64Settings:
	SetInt64Settings(key string, value int64) bool
	// SetStringSettings:
	SetStringSettings(key string, value string) bool
	// SetStrvSettings:
	SetStrvSettings(key string, value []string) bool
	// SetUintSettings:
	SetUintSettings(key string, value uint) bool
	// SetUint64Settings:
	SetUint64Settings(key string, value uint64) bool
	// SetValueSettings:
	SetValueSettings(key string, value *glib.Variant) bool
}

// settings implements the Settings class.
type settings struct {
	gextras.Objector
}

// WrapSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapSettings(obj *externglib.Object) Settings {
	return settings{
		Objector: obj,
	}
}

func marshalSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSettings(obj), nil
}

// NewSettings:
func NewSettings(schemaId string) Settings {
	var _arg1 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg1 = (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_new(_arg1)

	var _settings Settings // out

	_settings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Settings)

	return _settings
}

// NewSettingsWithPath:
func NewSettingsWithPath(schemaId string, path string) Settings {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg1 = (*C.gchar)(C.CString(schemaId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_settings_new_with_path(_arg1, _arg2)

	var _settings Settings // out

	_settings = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Settings)

	return _settings
}

func (s settings) ApplySettings() {
	var _arg0 *C.GSettings // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	C.g_settings_apply(_arg0)
}

func (s settings) BindSettings(key string, object gextras.Objector, property string, flags SettingsBindFlags) {
	var _arg0 *C.GSettings         // out
	var _arg1 *C.gchar             // out
	var _arg2 C.gpointer           // out
	var _arg3 *C.gchar             // out
	var _arg4 C.GSettingsBindFlags // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.gpointer)(unsafe.Pointer(object.Native()))
	_arg3 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GSettingsBindFlags(flags)

	C.g_settings_bind(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (s settings) BindWritableSettings(key string, object gextras.Objector, property string, inverted bool) {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gpointer   // out
	var _arg3 *C.gchar     // out
	var _arg4 C.gboolean   // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.gpointer)(unsafe.Pointer(object.Native()))
	_arg3 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg3))
	if inverted {
		_arg4 = C.TRUE
	}

	C.g_settings_bind_writable(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (s settings) CreateActionSettings(key string) Action {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GAction   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_create_action(_arg0, _arg1)

	var _action Action // out

	_action = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (s settings) DelaySettings() {
	var _arg0 *C.GSettings // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	C.g_settings_delay(_arg0)
}

func (s settings) Boolean(key string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) Child(name string) Settings {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GSettings // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_child(_arg0, _arg1)

	var _ret Settings // out

	_ret = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Settings)

	return _ret
}

func (s settings) DefaultValue(key string) *glib.Variant {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GVariant  // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_default_value(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_variant, func(v **glib.Variant) {
		C.free(unsafe.Pointer(v))
	})

	return _variant
}

func (s settings) Double(key string) float64 {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gdouble    // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_double(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s settings) Enum(key string) int {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gint       // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_enum(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s settings) Flags(key string) uint {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.guint      // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_flags(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s settings) HasUnapplied() bool {
	var _arg0 *C.GSettings // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	_cret = C.g_settings_get_has_unapplied(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) Int(key string) int {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gint       // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_int(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s settings) Int64(key string) int64 {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gint64     // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_int64(_arg0, _arg1)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

func (s settings) Range(key string) *glib.Variant {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GVariant  // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_range(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_variant, func(v **glib.Variant) {
		C.free(unsafe.Pointer(v))
	})

	return _variant
}

func (s settings) String(key string) string {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (s settings) Strv(key string) []string {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret **C.gchar

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_strv(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (s settings) Uint(key string) uint {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.guint      // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_uint(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s settings) Uint64(key string) uint64 {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.guint64    // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_uint64(_arg0, _arg1)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

func (s settings) UserValue(key string) *glib.Variant {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GVariant  // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_user_value(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_variant, func(v **glib.Variant) {
		C.free(unsafe.Pointer(v))
	})

	return _variant
}

func (s settings) Value(key string) *glib.Variant {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret *C.GVariant  // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_get_value(_arg0, _arg1)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_variant, func(v **glib.Variant) {
		C.free(unsafe.Pointer(v))
	})

	return _variant
}

func (s settings) IsWritableSettings(name string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_settings_is_writable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) ListChildrenSettings() []string {
	var _arg0 *C.GSettings // out
	var _cret **C.gchar

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	_cret = C.g_settings_list_children(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (s settings) ListKeysSettings() []string {
	var _arg0 *C.GSettings // out
	var _cret **C.gchar

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	_cret = C.g_settings_list_keys(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

func (s settings) RangeCheckSettings(key string, value *glib.Variant) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GVariant  // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	_cret = C.g_settings_range_check(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) ResetSettings(key string) {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_settings_reset(_arg0, _arg1)
}

func (s settings) RevertSettings() {
	var _arg0 *C.GSettings // out

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))

	C.g_settings_revert(_arg0)
}

func (s settings) SetBooleanSettings(key string, value bool) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gboolean   // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	_cret = C.g_settings_set_boolean(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetDoubleSettings(key string, value float64) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gdouble    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	_cret = C.g_settings_set_double(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetEnumSettings(key string, value int) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(value)

	_cret = C.g_settings_set_enum(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetFlagsSettings(key string, value uint) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.guint      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(value)

	_cret = C.g_settings_set_flags(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetIntSettings(key string, value int) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint       // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(value)

	_cret = C.g_settings_set_int(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetInt64Settings(key string, value int64) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.gint64     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint64(value)

	_cret = C.g_settings_set_int64(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetStringSettings(key string, value string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_settings_set_string(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetStrvSettings(key string, value []string) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 **C.gchar
	var _cret C.gboolean // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.gchar)(C.malloc(C.ulong(len(value)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice(_arg2, len(value))
		for i := range value {
			out[i] = (*C.gchar)(C.CString(value[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.g_settings_set_strv(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetUintSettings(key string, value uint) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.guint      // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(value)

	_cret = C.g_settings_set_uint(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetUint64Settings(key string, value uint64) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 C.guint64    // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint64(value)

	_cret = C.g_settings_set_uint64(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s settings) SetValueSettings(key string, value *glib.Variant) bool {
	var _arg0 *C.GSettings // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.GVariant  // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GSettings)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	_cret = C.g_settings_set_value(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
