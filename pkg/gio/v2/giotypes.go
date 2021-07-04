// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_file_attribute_matcher_get_type()), F: marshalFileAttributeMatcher},
		{T: externglib.Type(C.g_resource_get_type()), F: marshalResource},
		{T: externglib.Type(C.g_srv_target_get_type()), F: marshalSrvTarget},
	})
}

// FileAttributeMatcher determines if a string matches a file attribute.
type FileAttributeMatcher C.GFileAttributeMatcher

// WrapFileAttributeMatcher wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeMatcher(ptr unsafe.Pointer) *FileAttributeMatcher {
	return (*FileAttributeMatcher)(ptr)
}

func marshalFileAttributeMatcher(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*FileAttributeMatcher)(unsafe.Pointer(b)), nil
}

// NewFileAttributeMatcher constructs a struct FileAttributeMatcher.
func NewFileAttributeMatcher(attributes string) *FileAttributeMatcher {
	var _arg1 *C.char                  // out
	var _cret *C.GFileAttributeMatcher // in

	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_new(_arg1)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_fileAttributeMatcher, func(v **FileAttributeMatcher) {
		C.free(unsafe.Pointer(v))
	})

	return _fileAttributeMatcher
}

// Native returns the underlying C source pointer.
func (f *FileAttributeMatcher) Native() unsafe.Pointer {
	return unsafe.Pointer(f)
}

// EnumerateNamespace unreferences @matcher. If the reference count falls below
// 1, the @matcher is automatically freed.
func (m *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.char)(C.CString(ns))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_enumerate_namespace(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EnumerateNext unreferences @matcher. If the reference count falls below 1,
// the @matcher is automatically freed.
func (m *FileAttributeMatcher) EnumerateNext() string {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.char                  // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))

	_cret = C.g_file_attribute_matcher_enumerate_next(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Matches unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) Matches(attribute string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchesOnly unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_matches_only(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref unreferences @matcher. If the reference count falls below 1, the @matcher
// is automatically freed.
func (m *FileAttributeMatcher) Ref() *FileAttributeMatcher {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.GFileAttributeMatcher // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))

	_cret = C.g_file_attribute_matcher_ref(_arg0)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_fileAttributeMatcher, func(v **FileAttributeMatcher) {
		C.free(unsafe.Pointer(v))
	})

	return _fileAttributeMatcher
}

// Subtract unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) *FileAttributeMatcher {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.GFileAttributeMatcher // out
	var _cret *C.GFileAttributeMatcher // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(subtract.Native()))

	_cret = C.g_file_attribute_matcher_subtract(_arg0, _arg1)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_fileAttributeMatcher, func(v **FileAttributeMatcher) {
		C.free(unsafe.Pointer(v))
	})

	return _fileAttributeMatcher
}

// String unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) String() string {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.char                  // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))

	_cret = C.g_file_attribute_matcher_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) Unref() {
	var _arg0 *C.GFileAttributeMatcher // out

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m.Native()))

	C.g_file_attribute_matcher_unref(_arg0)
}

// IOExtension is an opaque data structure and can only be accessed using the
// following functions.
type IOExtension C.GIOExtension

// WrapIOExtension wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOExtension(ptr unsafe.Pointer) *IOExtension {
	return (*IOExtension)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IOExtension) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// Name gets a reference to the class for the type that is associated with
// @extension.
func (e *IOExtension) Name() string {
	var _arg0 *C.GIOExtension // out
	var _cret *C.char         // in

	_arg0 = (*C.GIOExtension)(unsafe.Pointer(e.Native()))

	_cret = C.g_io_extension_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Priority gets a reference to the class for the type that is associated with
// @extension.
func (e *IOExtension) Priority() int {
	var _arg0 *C.GIOExtension // out
	var _cret C.gint          // in

	_arg0 = (*C.GIOExtension)(unsafe.Pointer(e.Native()))

	_cret = C.g_io_extension_get_priority(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Type gets a reference to the class for the type that is associated with
// @extension.
func (e *IOExtension) Type() externglib.Type {
	var _arg0 *C.GIOExtension // out
	var _cret C.GType         // in

	_arg0 = (*C.GIOExtension)(unsafe.Pointer(e.Native()))

	_cret = C.g_io_extension_get_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// IOExtensionPoint is an opaque data structure and can only be accessed using
// the following functions.
type IOExtensionPoint C.GIOExtensionPoint

// WrapIOExtensionPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOExtensionPoint(ptr unsafe.Pointer) *IOExtensionPoint {
	return (*IOExtensionPoint)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IOExtensionPoint) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// ExtensionByName sets the required type for @extension_point to @type. All
// implementations must henceforth have this type.
func (e *IOExtensionPoint) ExtensionByName(name string) *IOExtension {
	var _arg0 *C.GIOExtensionPoint // out
	var _arg1 *C.char              // out
	var _cret *C.GIOExtension      // in

	_arg0 = (*C.GIOExtensionPoint)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_io_extension_point_get_extension_by_name(_arg0, _arg1)

	var _ioExtension *IOExtension // out

	_ioExtension = (*IOExtension)(unsafe.Pointer(_cret))

	return _ioExtension
}

// RequiredType sets the required type for @extension_point to @type. All
// implementations must henceforth have this type.
func (e *IOExtensionPoint) RequiredType() externglib.Type {
	var _arg0 *C.GIOExtensionPoint // out
	var _cret C.GType              // in

	_arg0 = (*C.GIOExtensionPoint)(unsafe.Pointer(e.Native()))

	_cret = C.g_io_extension_point_get_required_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetRequiredType sets the required type for @extension_point to @type. All
// implementations must henceforth have this type.
func (e *IOExtensionPoint) SetRequiredType(typ externglib.Type) {
	var _arg0 *C.GIOExtensionPoint // out
	var _arg1 C.GType              // out

	_arg0 = (*C.GIOExtensionPoint)(unsafe.Pointer(e.Native()))
	_arg1 = (C.GType)(typ)

	C.g_io_extension_point_set_required_type(_arg0, _arg1)
}

// IOSchedulerJob: opaque class for defining and scheduling IO jobs.
type IOSchedulerJob C.GIOSchedulerJob

// WrapIOSchedulerJob wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOSchedulerJob(ptr unsafe.Pointer) *IOSchedulerJob {
	return (*IOSchedulerJob)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IOSchedulerJob) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// IOStreamAdapter:
type IOStreamAdapter C.GIOStreamAdapter

// WrapIOStreamAdapter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOStreamAdapter(ptr unsafe.Pointer) *IOStreamAdapter {
	return (*IOStreamAdapter)(ptr)
}

// Native returns the underlying C source pointer.
func (i *IOStreamAdapter) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// InputMessage: structure used for scatter/gather data input when receiving
// multiple messages or packets in one go. You generally pass in an array of
// empty Vectors and the operation will use all the buffers as if they were one
// buffer, and will set @bytes_received to the total number of bytes received
// across all Vectors.
//
// This structure closely mirrors `struct mmsghdr` and `struct msghdr` from the
// POSIX sockets API (see `man 2 recvmmsg`).
//
// If @address is non-nil then it is set to the source address the message was
// received from, and the caller must free it afterwards.
//
// If @control_messages is non-nil then it is set to an array of control
// messages received with the message (if any), and the caller must free it
// afterwards. @num_control_messages is set to the number of elements in this
// array, which may be zero.
//
// Flags relevant to this message will be returned in @flags. For example,
// `MSG_EOR` or `MSG_TRUNC`.
type InputMessage C.GInputMessage

// WrapInputMessage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInputMessage(ptr unsafe.Pointer) *InputMessage {
	return (*InputMessage)(ptr)
}

// Native returns the underlying C source pointer.
func (i *InputMessage) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// InputVector: structure used for scatter/gather data input. You generally pass
// in an array of Vectors and the operation will store the read data starting in
// the first buffer, switching to the next as needed.
type InputVector C.GInputVector

// WrapInputVector wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInputVector(ptr unsafe.Pointer) *InputVector {
	return (*InputVector)(ptr)
}

// Native returns the underlying C source pointer.
func (i *InputVector) Native() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// OutputMessage: structure used for scatter/gather data output when sending
// multiple messages or packets in one go. You generally pass in an array of
// Vectors and the operation will use all the buffers as if they were one
// buffer.
//
// If @address is nil then the message is sent to the default receiver (as
// previously set by g_socket_connect()).
type OutputMessage C.GOutputMessage

// WrapOutputMessage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOutputMessage(ptr unsafe.Pointer) *OutputMessage {
	return (*OutputMessage)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OutputMessage) Native() unsafe.Pointer {
	return unsafe.Pointer(o)
}

// OutputVector: structure used for scatter/gather data output. You generally
// pass in an array of Vectors and the operation will use all the buffers as if
// they were one buffer.
type OutputVector C.GOutputVector

// WrapOutputVector wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOutputVector(ptr unsafe.Pointer) *OutputVector {
	return (*OutputVector)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OutputVector) Native() unsafe.Pointer {
	return unsafe.Pointer(o)
}

// Resource applications and libraries often contain binary or textual data that
// is really part of the application, rather than user data. For instance
// Builder .ui files, splashscreen images, GMenu markup XML, CSS files, icons,
// etc. These are often shipped as files in `$datadir/appname`, or manually
// included as literal strings in the code.
//
// The #GResource API and the [glib-compile-resources][glib-compile-resources]
// program provide a convenient and efficient alternative to this which has some
// nice properties. You maintain the files as normal files, so its easy to edit
// them, but during the build the files are combined into a binary bundle that
// is linked into the executable. This means that loading the resource files are
// efficient (as they are already in memory, shared with other instances) and
// simple (no need to check for things like I/O errors or locate the files in
// the filesystem). It also makes it easier to create relocatable applications.
//
// Resource files can also be marked as compressed. Such files will be included
// in the resource bundle in a compressed form, but will be automatically
// uncompressed when the resource is used. This is very useful e.g. for larger
// text files that are parsed once (or rarely) and then thrown away.
//
// Resource files can also be marked to be preprocessed, by setting the value of
// the `preprocess` attribute to a comma-separated list of preprocessing
// options. The only options currently supported are:
//
// `xml-stripblanks` which will use the xmllint command to strip ignorable
// whitespace from the XML file. For this to work, the `XMLLINT` environment
// variable must be set to the full path to the xmllint executable, or xmllint
// must be in the `PATH`; otherwise the preprocessing step is skipped.
//
// `to-pixdata` (deprecated since gdk-pixbuf 2.32) which will use the
// `gdk-pixbuf-pixdata` command to convert images to the Pixdata format, which
// allows you to create pixbufs directly using the data inside the resource
// file, rather than an (uncompressed) copy of it. For this, the
// `gdk-pixbuf-pixdata` program must be in the `PATH`, or the
// `GDK_PIXBUF_PIXDATA` environment variable must be set to the full path to the
// `gdk-pixbuf-pixdata` executable; otherwise the resource compiler will abort.
// `to-pixdata` has been deprecated since gdk-pixbuf 2.32, as #GResource
// supports embedding modern image formats just as well. Instead of using it,
// embed a PNG or SVG file in your #GResource.
//
// `json-stripblanks` which will use the `json-glib-format` command to strip
// ignorable whitespace from the JSON file. For this to work, the
// `JSON_GLIB_FORMAT` environment variable must be set to the full path to the
// `json-glib-format` executable, or it must be in the `PATH`; otherwise the
// preprocessing step is skipped. In addition, at least version 1.6 of
// `json-glib-format` is required.
//
// Resource files will be exported in the GResource namespace using the
// combination of the given `prefix` and the filename from the `file` element.
// The `alias` attribute can be used to alter the filename to expose them at a
// different location in the resource namespace. Typically, this is used to
// include files from a different source directory without exposing the source
// directory in the resource namespace, as in the example below.
//
// Resource bundles are created by the
// [glib-compile-resources][glib-compile-resources] program which takes an XML
// file that describes the bundle, and a set of files that the XML references.
// These are combined into a binary resource bundle.
//
// An example resource description:
//
//    <?xml version="1.0" encoding="UTF-8"?>
//    <gresources>
//      <gresource prefix="/org/gtk/Example">
//        <file>data/splashscreen.png</file>
//        <file compressed="true">dialog.ui</file>
//        <file preprocess="xml-stripblanks">menumarkup.xml</file>
//        <file alias="example.css">data/example.css</file>
//      </gresource>
//    </gresources>
//
// This will create a resource bundle with the following files:
//
//    /org/gtk/Example/data/splashscreen.png
//    /org/gtk/Example/dialog.ui
//    /org/gtk/Example/menumarkup.xml
//    /org/gtk/Example/example.css
//
// Note that all resources in the process share the same namespace, so use
// Java-style path prefixes (like in the above example) to avoid conflicts.
//
// You can then use [glib-compile-resources][glib-compile-resources] to compile
// the XML to a binary bundle that you can load with g_resource_load(). However,
// its more common to use the --generate-source and --generate-header arguments
// to create a source file and header to link directly into your application.
// This will generate `get_resource()`, `register_resource()` and
// `unregister_resource()` functions, prefixed by the `--c-name` argument passed
// to [glib-compile-resources][glib-compile-resources]. `get_resource()` returns
// the generated #GResource object. The register and unregister functions
// register the resource so its files can be accessed using
// g_resources_lookup_data().
//
// Once a #GResource has been created and registered all the data in it can be
// accessed globally in the process by using API calls like
// g_resources_open_stream() to stream the data or g_resources_lookup_data() to
// get a direct pointer to the data. You can also use URIs like
// "resource:///org/gtk/Example/data/splashscreen.png" with #GFile to access the
// resource data.
//
// Some higher-level APIs, such as Application, will automatically load
// resources from certain well-known paths in the resource namespace as a
// convenience. See the documentation for those APIs for details.
//
// There are two forms of the generated source, the default version uses the
// compiler support for constructor and destructor functions (where available)
// to automatically create and register the #GResource on startup or library
// load time. If you pass `--manual-register`, two functions to
// register/unregister the resource are created instead. This requires an
// explicit initialization call in your application/library, but it works on all
// platforms, even on the minor ones where constructors are not supported.
// (Constructor support is available for at least Win32, Mac OS and Linux.)
//
// Note that resource data can point directly into the data segment of e.g. a
// library, so if you are unloading libraries during runtime you need to be very
// careful with keeping around pointers to data from a resource, as this goes
// away when the library is unloaded. However, in practice this is not generally
// a problem, since most resource accesses are for your own resources, and
// resource data is often used once, during parsing, and then released.
//
// When debugging a program or testing a change to an installed version, it is
// often useful to be able to replace resources in the program or library,
// without recompiling, for debugging or quick hacking and testing purposes.
// Since GLib 2.50, it is possible to use the `G_RESOURCE_OVERLAYS` environment
// variable to selectively overlay resources with replacements from the
// filesystem. It is a G_SEARCHPATH_SEPARATOR-separated list of substitutions to
// perform during resource lookups. It is ignored when running in a setuid
// process.
//
// A substitution has the form
//
//    /org/gtk/libgtk=/home/desrt/gtk-overlay
//
// The part before the `=` is the resource subpath for which the overlay
// applies. The part after is a filesystem path which contains files and
// subdirectories as you would like to be loaded as resources with the
// equivalent names.
//
// In the example above, if an application tried to load a resource with the
// resource path `/org/gtk/libgtk/ui/gtkdialog.ui` then GResource would check
// the filesystem path `/home/desrt/gtk-overlay/ui/gtkdialog.ui`. If a file was
// found there, it would be used instead. This is an overlay, not an outright
// replacement, which means that if a file is not found at that path, the
// built-in version will be used instead. Whiteouts are not currently supported.
//
// Substitutions must start with a slash, and must not contain a trailing slash
// before the '='. The path after the slash should ideally be absolute, but this
// is not strictly required. It is possible to overlay the location of a single
// resource with an individual file.
type Resource C.GResource

// WrapResource wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapResource(ptr unsafe.Pointer) *Resource {
	return (*Resource)(ptr)
}

func marshalResource(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Resource)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Resource) Native() unsafe.Pointer {
	return unsafe.Pointer(r)
}

// EnumerateChildren: atomically decrements the reference count of @resource by
// one. If the reference count drops to 0, all memory allocated by the resource
// is released. This function is MT-safe and may be called from any thread.
func (r *Resource) EnumerateChildren(path string, lookupFlags ResourceLookupFlags) ([]string, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resource_enumerate_children(_arg0, _arg1, _arg2, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.char
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
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

// Info: atomically decrements the reference count of @resource by one. If the
// reference count drops to 0, all memory allocated by the resource is released.
// This function is MT-safe and may be called from any thread.
func (r *Resource) Info(path string, lookupFlags ResourceLookupFlags) (uint, uint32, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _arg3 C.gsize                // in
	var _arg4 C.guint32              // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	C.g_resource_get_info(_arg0, _arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _size uint    // out
	var _flags uint32 // out
	var _goerr error  // out

	_size = uint(_arg3)
	_flags = uint32(_arg4)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _size, _flags, _goerr
}

// OpenStream: atomically decrements the reference count of @resource by one. If
// the reference count drops to 0, all memory allocated by the resource is
// released. This function is MT-safe and may be called from any thread.
func (r *Resource) OpenStream(path string, lookupFlags ResourceLookupFlags) (InputStream, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret *C.GInputStream        // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resource_open_stream(_arg0, _arg1, _arg2, &_cerr)

	var _inputStream InputStream // out
	var _goerr error             // out

	_inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InputStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inputStream, _goerr
}

// Ref: atomically decrements the reference count of @resource by one. If the
// reference count drops to 0, all memory allocated by the resource is released.
// This function is MT-safe and may be called from any thread.
func (r *Resource) Ref() *Resource {
	var _arg0 *C.GResource // out
	var _cret *C.GResource // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r.Native()))

	_cret = C.g_resource_ref(_arg0)

	var _ret *Resource // out

	_ret = (*Resource)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_ret, func(v **Resource) {
		C.free(unsafe.Pointer(v))
	})

	return _ret
}

// Unref: atomically decrements the reference count of @resource by one. If the
// reference count drops to 0, all memory allocated by the resource is released.
// This function is MT-safe and may be called from any thread.
func (r *Resource) Unref() {
	var _arg0 *C.GResource // out

	_arg0 = (*C.GResource)(unsafe.Pointer(r.Native()))

	C.g_resource_unref(_arg0)
}

// SrvTarget: SRV (service) records are used by some network protocols to
// provide service-specific aliasing and load-balancing. For example, XMPP
// (Jabber) uses SRV records to locate the XMPP server for a domain; rather than
// connecting directly to "example.com" or assuming a specific server hostname
// like "xmpp.example.com", an XMPP client would look up the "xmpp-client" SRV
// record for "example.com", and then connect to whatever host was pointed to by
// that record.
//
// You can use g_resolver_lookup_service() or g_resolver_lookup_service_async()
// to find the Targets for a given service. However, if you are simply planning
// to connect to the remote service, you can use Service's Connectable interface
// and not need to worry about Target at all.
type SrvTarget C.GSrvTarget

// WrapSrvTarget wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSrvTarget(ptr unsafe.Pointer) *SrvTarget {
	return (*SrvTarget)(ptr)
}

func marshalSrvTarget(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SrvTarget)(unsafe.Pointer(b)), nil
}

// NewSrvTarget constructs a struct SrvTarget.
func NewSrvTarget(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	var _arg1 *C.gchar      // out
	var _arg2 C.guint16     // out
	var _arg3 C.guint16     // out
	var _arg4 C.guint16     // out
	var _cret *C.GSrvTarget // in

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(port)
	_arg3 = C.guint16(priority)
	_arg4 = C.guint16(weight)

	_cret = C.g_srv_target_new(_arg1, _arg2, _arg3, _arg4)

	var _srvTarget *SrvTarget // out

	_srvTarget = (*SrvTarget)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_srvTarget, func(v **SrvTarget) {
		C.free(unsafe.Pointer(v))
	})

	return _srvTarget
}

// Native returns the underlying C source pointer.
func (s *SrvTarget) Native() unsafe.Pointer {
	return unsafe.Pointer(s)
}

// Copy gets @target's weight. You should not need to look at this; #GResolver
// already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Copy() *SrvTarget {
	var _arg0 *C.GSrvTarget // out
	var _cret *C.GSrvTarget // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	_cret = C.g_srv_target_copy(_arg0)

	var _srvTarget *SrvTarget // out

	_srvTarget = (*SrvTarget)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_srvTarget, func(v **SrvTarget) {
		C.free(unsafe.Pointer(v))
	})

	return _srvTarget
}

// Free gets @target's weight. You should not need to look at this; #GResolver
// already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Free() {
	var _arg0 *C.GSrvTarget // out

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	C.g_srv_target_free(_arg0)
}

// Hostname gets @target's weight. You should not need to look at this;
// #GResolver already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Hostname() string {
	var _arg0 *C.GSrvTarget // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	_cret = C.g_srv_target_get_hostname(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Port gets @target's weight. You should not need to look at this; #GResolver
// already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Port() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	_cret = C.g_srv_target_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Priority gets @target's weight. You should not need to look at this;
// #GResolver already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Priority() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	_cret = C.g_srv_target_get_priority(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Weight gets @target's weight. You should not need to look at this; #GResolver
// already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Weight() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t.Native()))

	_cret = C.g_srv_target_get_weight(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}
