// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_builder_error_get_type()), F: marshalBuilderError},
		{T: externglib.Type(C.gtk_builder_get_type()), F: marshalBuilderer},
	})
}

// BuilderError: error codes that identify various errors that can occur while
// using Builder.
type BuilderError int

const (
	// BuilderErrorInvalidTypeFunction: type-func attribute didn’t name a
	// function that returns a #GType.
	BuilderErrorInvalidTypeFunction BuilderError = iota
	// BuilderErrorUnhandledTag: input contained a tag that Builder can’t
	// handle.
	BuilderErrorUnhandledTag
	// BuilderErrorMissingAttribute: attribute that is required by Builder was
	// missing.
	BuilderErrorMissingAttribute
	// BuilderErrorInvalidAttribute found an attribute that it doesn’t
	// understand.
	BuilderErrorInvalidAttribute
	// BuilderErrorInvalidTag found a tag that it doesn’t understand.
	BuilderErrorInvalidTag
	// BuilderErrorMissingPropertyValue: required property value was missing.
	BuilderErrorMissingPropertyValue
	// BuilderErrorInvalidValue couldn’t parse some attribute value.
	BuilderErrorInvalidValue
	// BuilderErrorVersionMismatch: input file requires a newer version of GTK+.
	BuilderErrorVersionMismatch
	// BuilderErrorDuplicateID: object id occurred twice.
	BuilderErrorDuplicateID
	// BuilderErrorObjectTypeRefused: specified object type is of the same type
	// or derived from the type of the composite class being extended with
	// builder XML.
	BuilderErrorObjectTypeRefused
	// BuilderErrorTemplateMismatch: wrong type was specified in a composite
	// class’s template XML
	BuilderErrorTemplateMismatch
	// BuilderErrorInvalidProperty: specified property is unknown for the object
	// class.
	BuilderErrorInvalidProperty
	// BuilderErrorInvalidSignal: specified signal is unknown for the object
	// class.
	BuilderErrorInvalidSignal
	// BuilderErrorInvalidID: object id is unknown
	BuilderErrorInvalidID
)

func marshalBuilderError(p uintptr) (interface{}, error) {
	return BuilderError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for BuilderError.
func (b BuilderError) String() string {
	switch b {
	case BuilderErrorInvalidTypeFunction:
		return "InvalidTypeFunction"
	case BuilderErrorUnhandledTag:
		return "UnhandledTag"
	case BuilderErrorMissingAttribute:
		return "MissingAttribute"
	case BuilderErrorInvalidAttribute:
		return "InvalidAttribute"
	case BuilderErrorInvalidTag:
		return "InvalidTag"
	case BuilderErrorMissingPropertyValue:
		return "MissingPropertyValue"
	case BuilderErrorInvalidValue:
		return "InvalidValue"
	case BuilderErrorVersionMismatch:
		return "VersionMismatch"
	case BuilderErrorDuplicateID:
		return "DuplicateID"
	case BuilderErrorObjectTypeRefused:
		return "ObjectTypeRefused"
	case BuilderErrorTemplateMismatch:
		return "TemplateMismatch"
	case BuilderErrorInvalidProperty:
		return "InvalidProperty"
	case BuilderErrorInvalidSignal:
		return "InvalidSignal"
	case BuilderErrorInvalidID:
		return "InvalidID"
	default:
		return fmt.Sprintf("BuilderError(%d)", b)
	}
}

// BuilderOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BuilderOverrider interface {
	// TypeFromName looks up a type by name, using the virtual function that
	// Builder has for that purpose. This is mainly used when implementing the
	// Buildable interface on a type.
	TypeFromName(typeName string) externglib.Type
}

// Builder is an auxiliary object that reads textual descriptions of a user
// interface and instantiates the described objects. To create a GtkBuilder from
// a user interface description, call gtk_builder_new_from_file(),
// gtk_builder_new_from_resource() or gtk_builder_new_from_string().
//
// In the (unusual) case that you want to add user interface descriptions from
// multiple sources to the same GtkBuilder you can call gtk_builder_new() to get
// an empty builder and populate it by (multiple) calls to
// gtk_builder_add_from_file(), gtk_builder_add_from_resource() or
// gtk_builder_add_from_string().
//
// A GtkBuilder holds a reference to all objects that it has constructed and
// drops these references when it is finalized. This finalization can cause the
// destruction of non-widget objects or widgets which are not contained in a
// toplevel window. For toplevel windows constructed by a builder, it is the
// responsibility of the user to call gtk_widget_destroy() to get rid of them
// and all the widgets they contain.
//
// The functions gtk_builder_get_object() and gtk_builder_get_objects() can be
// used to access the widgets in the interface by the names assigned to them
// inside the UI description. Toplevel windows returned by these functions will
// stay around until the user explicitly destroys them with
// gtk_widget_destroy(). Other widgets will either be part of a larger hierarchy
// constructed by the builder (in which case you should not have to worry about
// their lifecycle), or without a parent, in which case they have to be added to
// some container to make use of them. Non-widget objects need to be reffed with
// g_object_ref() to keep them beyond the lifespan of the builder.
//
// The function gtk_builder_connect_signals() and variants thereof can be used
// to connect handlers to the named signals in the description.
//
//
// GtkBuilder UI Definitions
//
// GtkBuilder parses textual descriptions of user interfaces which are specified
// in an XML format which can be roughly described by the RELAX NG schema below.
// We refer to these descriptions as “GtkBuilder UI definitions” or just “UI
// definitions” if the context is clear. Do not confuse GtkBuilder UI
// Definitions with [GtkUIManager UI Definitions][XML-UI], which are more
// limited in scope. It is common to use .ui as the filename extension for files
// containing GtkBuilder UI definitions.
//
// RELAX NG Compact Syntax
// (https://gitlab.gnome.org/GNOME/gtk/-/blob/gtk-3-24/gtk/gtkbuilder.rnc)
//
// The toplevel element is <interface>. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings using
// dgettext() in the domain specified. This can also be done by calling
// gtk_builder_set_translation_domain() on the builder. Objects are described by
// <object> elements, which can contain <property> elements to set properties,
// <signal> elements which connect signals to handlers, and <child> elements,
// which describe child objects (most often widgets inside a container, but also
// e.g. actions in an action group, or columns in a tree model). A <child>
// element contains an <object> element which describes the child object. The
// target toolkit version(s) are described by <requires> elements, the “lib”
// attribute specifies the widget library in question (currently the only
// supported value is “gtk+”) and the “version” attribute specifies the target
// version in the form “<major>.<minor>”. The builder will error out if the
// version requirements are not met.
//
// Typically, the specific kind of object represented by an <object> element is
// specified by the “class” attribute. If the type has not been loaded yet, GTK+
// tries to find the get_type() function from the class name by applying
// heuristics. This works in most cases, but if necessary, it is possible to
// specify the name of the get_type() function explictly with the "type-func"
// attribute. As a special case, GtkBuilder allows to use an object that has
// been constructed by a UIManager in another part of the UI definition by
// specifying the id of the UIManager in the “constructor” attribute and the
// name of the object in the “id” attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with gtk_builder_get_object().
// An id is also necessary to use the object as property value in other parts of
// the UI definition. GTK+ reserves ids starting and ending with ___ (3
// underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the <property>
// element: the “name” attribute specifies the name of the property, and the
// content of the element specifies the value. If the “translatable” attribute
// is set to a true value, GTK+ uses gettext() (or dgettext() if the builder has
// a translation domain set) to find a translation for the value. This happens
// before the value is parsed, so it can be used for properties of any type, but
// it is probably most useful for string properties. It is also possible to
// specify a context to disambiguate short strings, and comments which may help
// the translators.
//
// GtkBuilder can parse textual representations for the most common property
// types: characters, strings, integers, floating-point numbers, booleans
// (strings like “TRUE”, “t”, “yes”, “y”, “1” are interpreted as TRUE, strings
// like “FALSE”, “f”, “no”, “n”, “0” are interpreted as FALSE), enumerations
// (can be specified by their name, nick or integer value), flags (can be
// specified by their name, nick, integer value, optionally combined with “|”,
// e.g. “GTK_VISIBLE|GTK_REALIZED”) and colors (in a format understood by
// gdk_rgba_parse()).
//
// GVariants can be specified in the format understood by g_variant_parse(), and
// pixbufs can be specified as a filename of an image file to load.
//
// Objects can be referred to by their name and by default refer to objects
// declared in the local xml fragment and objects exposed via
// gtk_builder_expose_object(). In general, GtkBuilder allows forward references
// to objects — declared in the local xml; an object doesn’t have to be
// constructed before it can be referred to. The exception to this rule is that
// an object has to be constructed before it can be used as the value of a
// construct-only property.
//
// It is also possible to bind a property value to another object's property
// value using the attributes "bind-source" to specify the source object of the
// binding, "bind-property" to specify the source property and optionally
// "bind-flags" to specify the binding flags. Internally builder implements this
// using GBinding objects. For more information see g_object_bind_property()
//
// Signal handlers are set up with the <signal> element. The “name” attribute
// specifies the name of the signal, and the “handler” attribute specifies the
// function to connect to the signal. By default, GTK+ tries to find the handler
// using g_module_symbol(), but this can be changed by passing a custom
// BuilderConnectFunc to gtk_builder_connect_signals_full(). The remaining
// attributes, “after”, “swapped” and “object”, have the same meaning as the
// corresponding parameters of the g_signal_connect_object() or
// g_signal_connect_data() functions. A “last_modification_time” attribute is
// also allowed, but it does not have a meaning to the builder.
//
// Sometimes it is necessary to refer to widgets which have implicitly been
// constructed by GTK+ as part of a composite widget, to set properties on them
// or to add further children (e.g. the vbox of a Dialog). This can be achieved
// by setting the “internal-child” property of the <child> element to a true
// value. Note that GtkBuilder still requires an <object> element for the
// internal child, even if it has already been constructed.
//
// A number of widgets have different places where a child can be added (e.g.
// tabs vs. page content in notebooks). This can be reflected in a UI definition
// by specifying the “type” attribute on a <child> The possible values for the
// “type” attribute are described in the sections describing the widget-specific
// portions of UI definitions.
//
// A GtkBuilder UI Definition
//
//    <interface>
//      <object class="GtkDialog" id="dialog1">
//        <child internal-child="vbox">
//          <object class="GtkBox" id="vbox1">
//            <property name="border-width">10</property>
//            <child internal-child="action_area">
//              <object class="GtkButtonBox" id="hbuttonbox1">
//                <property name="border-width">20</property>
//                <child>
//                  <object class="GtkButton" id="ok_button">
//                    <property name="label">gtk-ok</property>
//                    <property name="use-stock">TRUE</property>
//                    <signal name="clicked" handler="ok_button_clicked"/>
//                  </object>
//                </child>
//              </object>
//            </child>
//          </object>
//        </child>
//      </object>
//    </interface>
//
// Beyond this general structure, several object classes define their own XML
// DTD fragments for filling in the ANY placeholders in the DTD above. Note that
// a custom element in a <child> element gets parsed by the custom tag handler
// of the parent object, while a custom element in an <object> element gets
// parsed by the custom tag handler of the object.
//
// These XML fragments are explained in the documentation of the respective
// objects.
//
// Additionally, since 3.10 a special <template> tag has been added to the
// format allowing one to define a widget class’s components. See the [GtkWidget
// documentation][composite-templates] for details.
type Builder struct {
	*externglib.Object
}

func WrapBuilder(obj *externglib.Object) *Builder {
	return &Builder{
		Object: obj,
	}
}

func marshalBuilderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuilder(obj), nil
}

// NewBuilder creates a new empty builder object.
//
// This function is only useful if you intend to make multiple calls to
// gtk_builder_add_from_file(), gtk_builder_add_from_resource() or
// gtk_builder_add_from_string() in order to merge multiple UI descriptions into
// a single builder.
//
// Most users will probably want to use gtk_builder_new_from_file(),
// gtk_builder_new_from_resource() or gtk_builder_new_from_string().
func NewBuilder() *Builder {
	var _cret *C.GtkBuilder // in

	_cret = C.gtk_builder_new()

	var _builder *Builder // out

	_builder = WrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromFile builds the [GtkBuilder UI definition][BUILDER-UI] in the
// file filename.
//
// If there is an error opening the file or parsing the description then the
// program will be aborted. You should only ever attempt to parse user interface
// descriptions that are shipped as part of your program.
func NewBuilderFromFile(filename string) *Builder {
	var _arg1 *C.gchar      // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_file(_arg1)
	runtime.KeepAlive(filename)

	var _builder *Builder // out

	_builder = WrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromResource builds the [GtkBuilder UI definition][BUILDER-UI] at
// resource_path.
//
// If there is an error locating the resource or parsing the description, then
// the program will be aborted.
func NewBuilderFromResource(resourcePath string) *Builder {
	var _arg1 *C.gchar      // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _builder *Builder // out

	_builder = WrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromString builds the user interface described by string (in the
// [GtkBuilder UI definition][BUILDER-UI] format).
//
// If string is NULL-terminated, then length should be -1. If length is not -1,
// then it is the length of string.
//
// If there is an error parsing string then the program will be aborted. You
// should not attempt to parse user interface description from untrusted
// sources.
func NewBuilderFromString(_string string, length int) *Builder {
	var _arg1 *C.gchar      // out
	var _arg2 C.gssize      // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.gtk_builder_new_from_string(_arg1, _arg2)
	runtime.KeepAlive(_string)
	runtime.KeepAlive(length)

	var _builder *Builder // out

	_builder = WrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// AddFromFile parses a file containing a [GtkBuilder UI definition][BUILDER-UI]
// and merges it with the current contents of builder.
//
// Most users will probably want to use gtk_builder_new_from_file().
//
// If an error occurs, 0 will be returned and error will be assigned a #GError
// from the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call. You
// should not use this function with untrusted files (ie: files that are not
// part of your application). Broken Builder files can easily crash your
// program, and it’s possible that memory was leaked leading up to the reported
// failure. The only reasonable thing to do when an error is detected is to call
// g_error().
func (builder *Builder) AddFromFile(filename string) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_add_from_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(filename)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of builder.
//
// Most users will probably want to use gtk_builder_new_from_resource().
//
// If an error occurs, 0 will be returned and error will be assigned a #GError
// from the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call. The
// only reasonable thing to do when an error is detected is to call g_error().
func (builder *Builder) AddFromResource(resourcePath string) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_add_from_resource(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(resourcePath)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of builder.
//
// Most users will probably want to use gtk_builder_new_from_string().
//
// Upon errors 0 will be returned and error will be assigned a #GError from the
// K_BUILDER_ERROR, MARKUP_ERROR or VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call. The
// only reasonable thing to do when an error is detected is to call g_error().
func (builder *Builder) AddFromString(buffer string, length uint) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize       // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(length)

	_cret = C.gtk_builder_add_from_string(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(length)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddObjectsFromFile parses a file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges them
// with the current contents of builder.
//
// Upon errors 0 will be returned and error will be assigned a #GError from the
// K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a TreeView that depends on its TreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromFile(filename string, objectIds []string) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _arg2 **C.gchar     // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(objectIds)+1)
			var zero *C.gchar
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_builder_add_objects_from_file(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(objectIds)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddObjectsFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges them
// with the current contents of builder.
//
// Upon errors 0 will be returned and error will be assigned a #GError from the
// K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a TreeView that depends on its TreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromResource(resourcePath string, objectIds []string) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _arg2 **C.gchar     // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(objectIds)+1)
			var zero *C.gchar
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_builder_add_objects_from_resource(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(resourcePath)
	runtime.KeepAlive(objectIds)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// AddObjectsFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges them
// with the current contents of builder.
//
// Upon errors 0 will be returned and error will be assigned a #GError from the
// K_BUILDER_ERROR or MARKUP_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a TreeView that depends on its TreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromString(buffer string, length uint, objectIds []string) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _arg2 C.gsize       // out
	var _arg3 **C.gchar     // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gsize(length)
	{
		_arg3 = (**C.gchar)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(objectIds)+1)
			var zero *C.gchar
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gtk_builder_add_objects_from_string(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(length)
	runtime.KeepAlive(objectIds)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// ConnectSignals: this method is a simpler variation of
// gtk_builder_connect_signals_full(). It uses symbols explicitly added to
// builder with prior calls to gtk_builder_add_callback_symbol(). In the case
// that symbols are not explicitly added; it uses #GModule’s introspective
// features (by opening the module NULL) to look at the application’s symbol
// table. From here it tries to match the signal handler names given in the
// interface description with symbols in the application and connects the
// signals. Note that this function can only be called once, subsequent calls
// will do nothing.
//
// Note that unless gtk_builder_add_callback_symbol() is called for all signal
// callbacks which are referenced by the loaded XML, this function will require
// that #GModule be supported on the platform.
//
// If you rely on #GModule support to lookup callbacks in the symbol table, the
// following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with MODULE_EXPORT, or they will not be put in the symbol table. On Linux and
// Unices, this is not necessary; applications should instead be compiled with
// the -Wl,--export-dynamic CFLAGS, and linked against gmodule-export-2.0.
func (builder *Builder) ConnectSignals(userData cgo.Handle) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 C.gpointer    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (C.gpointer)(unsafe.Pointer(userData))

	C.gtk_builder_connect_signals(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(userData)
}

// ExposeObject: add object to the builder object pool so it can be referenced
// just like any other object built by builder.
func (builder *Builder) ExposeObject(name string, object *externglib.Object) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _arg2 *C.GObject    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GObject)(unsafe.Pointer(object.Native()))

	C.gtk_builder_expose_object(_arg0, _arg1, _arg2)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)
	runtime.KeepAlive(object)
}

// ExtendWithTemplate: main private entry point for building composite container
// components from template XML.
//
// This is exported purely to let gtk-builder-tool validate templates,
// applications have no need to call this function.
func (builder *Builder) ExtendWithTemplate(widget Widgetter, templateType externglib.Type, buffer string, length uint) (uint, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.GType       // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gsize       // out
	var _cret C.guint       // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GType(templateType)
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gsize(length)

	_cret = C.gtk_builder_extend_with_template(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(templateType)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(length)

	var _guint uint  // out
	var _goerr error // out

	_guint = uint(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _guint, _goerr
}

// Application gets the Application associated with the builder.
//
// The Application is used for creating action proxies as requested from XML
// that the builder is loading.
//
// By default, the builder uses the default application: the one from
// g_application_get_default(). If you want to use another application for
// constructing proxies, use gtk_builder_set_application().
func (builder *Builder) Application() *Application {
	var _arg0 *C.GtkBuilder     // out
	var _cret *C.GtkApplication // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_application(_arg0)
	runtime.KeepAlive(builder)

	var _application *Application // out

	if _cret != nil {
		_application = WrapApplication(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _application
}

// GetObject gets the object named name. Note that this function does not
// increment the reference count of the returned object.
func (builder *Builder) GetObject(name string) *externglib.Object {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_object(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)

	var _object *externglib.Object // out

	if _cret != nil {
		_object = externglib.Take(unsafe.Pointer(_cret))
	}

	return _object
}

// Objects gets all objects that have been constructed by builder. Note that
// this function does not increment the reference counts of the returned
// objects.
func (builder *Builder) Objects() []*externglib.Object {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.GSList     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_objects(_arg0)
	runtime.KeepAlive(builder)

	var _sList []*externglib.Object // out

	_sList = make([]*externglib.Object, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GObject)(v)
		var dst *externglib.Object // out
		dst = externglib.Take(unsafe.Pointer(src))
		_sList = append(_sList, dst)
	})

	return _sList
}

// TranslationDomain gets the translation domain of builder.
func (builder *Builder) TranslationDomain() string {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_translation_domain(_arg0)
	runtime.KeepAlive(builder)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TypeFromName looks up a type by name, using the virtual function that Builder
// has for that purpose. This is mainly used when implementing the Buildable
// interface on a type.
func (builder *Builder) TypeFromName(typeName string) externglib.Type {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cret C.GType       // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(typeName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_type_from_name(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typeName)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetApplication sets the application associated with builder.
//
// You only need this function if there is more than one #GApplication in your
// process. application cannot be NULL.
func (builder *Builder) SetApplication(application *Application) {
	var _arg0 *C.GtkBuilder     // out
	var _arg1 *C.GtkApplication // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.GtkApplication)(unsafe.Pointer(application.Native()))

	C.gtk_builder_set_application(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(application)
}

// SetTranslationDomain sets the translation domain of builder. See
// Builder:translation-domain.
func (builder *Builder) SetTranslationDomain(domain string) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	if domain != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_builder_set_translation_domain(_arg0, _arg1)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(domain)
}

// ValueFromStringType: like gtk_builder_value_from_string(), this function
// demarshals a value from a string, but takes a #GType instead of Spec. This
// function calls g_value_init() on the value argument, so it need not be
// initialised beforehand.
//
// Upon errors FALSE will be returned and error will be assigned a #GError from
// the K_BUILDER_ERROR domain.
func (builder *Builder) ValueFromStringType(typ externglib.Type, _string string) (externglib.Value, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 C.GType       // out
	var _arg2 *C.gchar      // out
	var _arg3 C.GValue      // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = C.GType(typ)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_builder_value_from_string_type(_arg0, _arg1, _arg2, &_arg3, &_cerr)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(_string)

	var _value externglib.Value // out
	var _goerr error            // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}
