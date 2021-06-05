// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <glib.h>
import "C"

// BookmarkFileError: error codes returned by bookmark file parsing.
type BookmarkFileError int

const (
	// BookmarkFileErrorInvalidURI: URI was ill-formed
	BookmarkFileErrorInvalidURI BookmarkFileError = 0
	// BookmarkFileErrorInvalidValue: a requested field was not found
	BookmarkFileErrorInvalidValue BookmarkFileError = 1
	// BookmarkFileErrorAppNotRegistered: a requested application did not
	// register a bookmark
	BookmarkFileErrorAppNotRegistered BookmarkFileError = 2
	// BookmarkFileErrorURINotFound: a requested URI was not found
	BookmarkFileErrorURINotFound BookmarkFileError = 3
	// BookmarkFileErrorRead: document was ill formed
	BookmarkFileErrorRead BookmarkFileError = 4
	// BookmarkFileErrorUnknownEncoding: the text being parsed was in an unknown
	// encoding
	BookmarkFileErrorUnknownEncoding BookmarkFileError = 5
	// BookmarkFileErrorWrite: an error occurred while writing
	BookmarkFileErrorWrite BookmarkFileError = 6
	// BookmarkFileErrorFileNotFound: requested file was not found
	BookmarkFileErrorFileNotFound BookmarkFileError = 7
)

// BookmarkFile: the `GBookmarkFile` structure contains only private data and
// should not be directly accessed.
type BookmarkFile struct {
	native C.GBookmarkFile
}

// WrapBookmarkFile wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBookmarkFile(ptr unsafe.Pointer) *BookmarkFile {
	if ptr == nil {
		return nil
	}

	return (*BookmarkFile)(ptr)
}

func marshalBookmarkFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBookmarkFile(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (b *BookmarkFile) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// AddApplication adds the application with @name and @exec to the list of
// applications that have registered a bookmark for @uri into @bookmark.
//
// Every bookmark inside a File must have at least an application registered.
// Each application must provide a name, a command line useful for launching the
// bookmark, the number of times the bookmark has been registered by the
// application and the last time the application registered this bookmark.
//
// If @name is nil, the name of the application will be the same returned by
// g_get_application_name(); if @exec is nil, the command line will be a
// composition of the program name as returned by g_get_prgname() and the "\u"
// modifier, which will be expanded to the bookmark's URI.
//
// This function will automatically take care of updating the registrations
// count and timestamping in case an application with the same @name had already
// registered a bookmark for @uri inside @bookmark.
//
// If no bookmark for @uri is found, one is created.
func (b *BookmarkFile) AddApplication(uri string, name string, exec string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(exec))
	defer C.free(unsafe.Pointer(arg3))

	C.g_bookmark_file_add_application(arg0, uri, name, exec)
}

// AddGroup adds @group to the list of groups to which the bookmark for @uri
// belongs to.
//
// If no bookmark for @uri is found then it is created.
func (b *BookmarkFile) AddGroup(uri string, group string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(group))
	defer C.free(unsafe.Pointer(arg2))

	C.g_bookmark_file_add_group(arg0, uri, group)
}

// Free frees a File.
func (b *BookmarkFile) Free() {
	var arg0 *C.GBookmarkFile

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	C.g_bookmark_file_free(arg0)
}

// Added gets the time the bookmark for @uri was added to @bookmark
//
// In the event the URI cannot be found, -1 is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Added(uri string) (glong int32, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.time_t
	var goret1 int32
	var goerr error

	cret = C.g_bookmark_file_get_added(arg0, uri, &errout)

	goret1 = C.time_t(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// AddedDateTime gets the time the bookmark for @uri was added to @bookmark
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) AddedDateTime(uri string) (dateTime *DateTime, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDateTime
	var goret1 *DateTime
	var goerr error

	cret = C.g_bookmark_file_get_added_date_time(arg0, uri, &errout)

	goret1 = WrapDateTime(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// AppInfo gets the registration information of @app_name for the bookmark for
// @uri. See g_bookmark_file_set_application_info() for more information about
// the returned data.
//
// The string returned in @app_exec must be freed.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event that no application with name
// @app_name has registered a bookmark for @uri, false is returned and error is
// set to BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED. In the event that unquoting
// the command line fails, an error of the SHELL_ERROR domain is set and false
// is returned.
func (b *BookmarkFile) AppInfo(uri string, name string) (exec string, count uint, stamp int32, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 **C.gchar
	var ret3 string
	var arg4 *C.guint
	var ret4 uint
	var arg5 *C.time_t
	var ret5 int32
	var goerr error

	C.g_bookmark_file_get_app_info(arg0, uri, name, &arg3, &arg4, &arg5, &errout)

	ret3 = C.GoString(arg3)
	defer C.free(unsafe.Pointer(arg3))
	ret4 = *C.guint(arg4)
	ret5 = *C.time_t(arg5)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret3, ret4, ret5, goerr
}

// ApplicationInfo gets the registration information of @app_name for the
// bookmark for @uri. See g_bookmark_file_set_application_info() for more
// information about the returned data.
//
// The string returned in @app_exec must be freed.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event that no application with name
// @app_name has registered a bookmark for @uri, false is returned and error is
// set to BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED. In the event that unquoting
// the command line fails, an error of the SHELL_ERROR domain is set and false
// is returned.
func (b *BookmarkFile) ApplicationInfo(uri string, name string) (exec string, count uint, stamp *DateTime, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var arg2 *C.char
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 **C.char
	var ret3 string
	var arg4 *C.uint
	var ret4 uint
	var arg5 **C.GDateTime
	var ret5 **DateTime
	var goerr error

	C.g_bookmark_file_get_application_info(arg0, uri, name, &arg3, &arg4, &arg5, &errout)

	ret3 = C.GoString(arg3)
	defer C.free(unsafe.Pointer(arg3))
	ret4 = *C.uint(arg4)
	ret5 = WrapDateTime(unsafe.Pointer(arg5))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret3, ret4, ret5, goerr
}

// Applications retrieves the names of the applications that have registered the
// bookmark for @uri.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Applications(uri string) (length uint, utf8s []string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret **C.gchar
	var arg2 *C.gsize
	var goret2 []string
	var goerr error

	cret = C.g_bookmark_file_get_applications(arg0, uri, &arg2, &errout)

	goret2 = make([]string, arg2)
	for i := 0; i < uintptr(arg2); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		goret2[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret2, goret2, goerr
}

// Description retrieves the description of the bookmark for @uri.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Description(uri string) (utf8 string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret1 string
	var goerr error

	cret = C.g_bookmark_file_get_description(arg0, uri, &errout)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Groups retrieves the list of group names of the bookmark for @uri.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
//
// The returned array is nil terminated, so @length may optionally be nil.
func (b *BookmarkFile) Groups(uri string) (length uint, utf8s []string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret **C.gchar
	var arg2 *C.gsize
	var goret2 []string
	var goerr error

	cret = C.g_bookmark_file_get_groups(arg0, uri, &arg2, &errout)

	goret2 = make([]string, arg2)
	for i := 0; i < uintptr(arg2); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		goret2[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret2, goret2, goerr
}

// Icon gets the icon of the bookmark for @uri.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Icon(uri string) (href string, mimeType string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 **C.gchar
	var ret2 string
	var arg3 **C.gchar
	var ret3 string
	var goerr error

	C.g_bookmark_file_get_icon(arg0, uri, &arg2, &arg3, &errout)

	ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	ret3 = C.GoString(arg3)
	defer C.free(unsafe.Pointer(arg3))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret2, ret3, goerr
}

// IsPrivate gets whether the private flag of the bookmark for @uri is set.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event that the private flag cannot
// be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_INVALID_VALUE.
func (b *BookmarkFile) IsPrivate(uri string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var goerr error

	C.g_bookmark_file_get_is_private(arg0, uri, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// MIMEType retrieves the MIME type of the resource pointed by @uri.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event that the MIME type cannot be
// found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_INVALID_VALUE.
func (b *BookmarkFile) MIMEType(uri string) (utf8 string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret1 string
	var goerr error

	cret = C.g_bookmark_file_get_mime_type(arg0, uri, &errout)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Modified gets the time when the bookmark for @uri was last modified.
//
// In the event the URI cannot be found, -1 is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Modified(uri string) (glong int32, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.time_t
	var goret1 int32
	var goerr error

	cret = C.g_bookmark_file_get_modified(arg0, uri, &errout)

	goret1 = C.time_t(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// ModifiedDateTime gets the time when the bookmark for @uri was last modified.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) ModifiedDateTime(uri string) (dateTime *DateTime, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDateTime
	var goret1 *DateTime
	var goerr error

	cret = C.g_bookmark_file_get_modified_date_time(arg0, uri, &errout)

	goret1 = WrapDateTime(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Size gets the number of bookmarks inside @bookmark.
func (b *BookmarkFile) Size() int {
	var arg0 *C.GBookmarkFile

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	var cret C.gint
	var goret1 int

	cret = C.g_bookmark_file_get_size(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// Title returns the title of the bookmark for @uri.
//
// If @uri is nil, the title of @bookmark is returned.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Title(uri string) (utf8 string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar
	var goret1 string
	var goerr error

	cret = C.g_bookmark_file_get_title(arg0, uri, &errout)

	goret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// Uris returns all URIs of the bookmarks in the bookmark file @bookmark. The
// array of returned URIs will be nil-terminated, so @length may optionally be
// nil.
func (b *BookmarkFile) Uris() (length uint, utf8s []string) {
	var arg0 *C.GBookmarkFile

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	var cret **C.gchar
	var arg1 *C.gsize
	var goret2 []string

	cret = C.g_bookmark_file_get_uris(arg0, &arg1)

	goret2 = make([]string, arg1)
	for i := 0; i < uintptr(arg1); i++ {
		src := (*C.gchar)(ptr.Add(unsafe.Pointer(cret), i))
		goret2[i] = C.GoString(src)
		defer C.free(unsafe.Pointer(src))
	}

	return ret1, goret2
}

// Visited gets the time the bookmark for @uri was last visited.
//
// In the event the URI cannot be found, -1 is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) Visited(uri string) (glong int32, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.time_t
	var goret1 int32
	var goerr error

	cret = C.g_bookmark_file_get_visited(arg0, uri, &errout)

	goret1 = C.time_t(cret)
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// VisitedDateTime gets the time the bookmark for @uri was last visited.
//
// In the event the URI cannot be found, nil is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) VisitedDateTime(uri string) (dateTime *DateTime, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDateTime
	var goret1 *DateTime
	var goerr error

	cret = C.g_bookmark_file_get_visited_date_time(arg0, uri, &errout)

	goret1 = WrapDateTime(unsafe.Pointer(cret))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return goret1, goerr
}

// HasApplication checks whether the bookmark for @uri inside @bookmark has been
// registered by application @name.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) HasApplication(uri string, name string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.g_bookmark_file_has_application(arg0, uri, name, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// HasGroup checks whether @group appears in the list of groups to which the
// bookmark for @uri belongs to.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) HasGroup(uri string, group string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(group))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.g_bookmark_file_has_group(arg0, uri, group, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// HasItem looks whether the desktop bookmark has an item with its URI set to
// @uri.
func (b *BookmarkFile) HasItem(uri string) bool {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var goret1 bool

	cret = C.g_bookmark_file_has_item(arg0, uri)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// LoadFromData loads a bookmark file from memory into an empty File structure.
// If the object cannot be created then @error is set to a FileError.
func (b *BookmarkFile) LoadFromData(data []byte) error {
	var arg0 *C.GBookmarkFile

	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	var goerr error

	C.g_bookmark_file_load_from_data(arg0, data, length, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// LoadFromDataDirs: this function looks for a desktop bookmark file named @file
// in the paths returned from g_get_user_data_dir() and
// g_get_system_data_dirs(), loads the file into @bookmark and returns the
// file's full path in @full_path. If the file could not be loaded then @error
// is set to either a Error or FileError.
func (b *BookmarkFile) LoadFromDataDirs(file string) (fullPath string, err error) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(file))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 **C.gchar
	var ret2 string
	var goerr error

	C.g_bookmark_file_load_from_data_dirs(arg0, file, &arg2, &errout)

	ret2 = C.GoString(arg2)
	defer C.free(unsafe.Pointer(arg2))
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret2, goerr
}

// LoadFromFile loads a desktop bookmark file into an empty File structure. If
// the file could not be loaded then @error is set to either a Error or
// FileError.
func (b *BookmarkFile) LoadFromFile(filename string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var goerr error

	C.g_bookmark_file_load_from_file(arg0, filename, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// MoveItem changes the URI of a bookmark item from @old_uri to @new_uri. Any
// existing bookmark for @new_uri will be overwritten. If @new_uri is nil, then
// the bookmark is removed.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND.
func (b *BookmarkFile) MoveItem(oldURI string, newURI string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(oldURI))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(newURI))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.g_bookmark_file_move_item(arg0, oldURI, newURI, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// RemoveApplication removes application registered with @name from the list of
// applications that have registered a bookmark for @uri inside @bookmark.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event that no application with name
// @app_name has registered a bookmark for @uri, false is returned and error is
// set to BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED.
func (b *BookmarkFile) RemoveApplication(uri string, name string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.g_bookmark_file_remove_application(arg0, uri, name, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// RemoveGroup removes @group from the list of groups to which the bookmark for
// @uri belongs to.
//
// In the event the URI cannot be found, false is returned and @error is set to
// BOOKMARK_FILE_ERROR_URI_NOT_FOUND. In the event no group was defined, false
// is returned and @error is set to BOOKMARK_FILE_ERROR_INVALID_VALUE.
func (b *BookmarkFile) RemoveGroup(uri string, group string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(group))
	defer C.free(unsafe.Pointer(arg2))

	var goerr error

	C.g_bookmark_file_remove_group(arg0, uri, group, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// RemoveItem removes the bookmark for @uri from the bookmark file @bookmark.
func (b *BookmarkFile) RemoveItem(uri string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	var goerr error

	C.g_bookmark_file_remove_item(arg0, uri, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// SetAdded sets the time the bookmark for @uri was added into @bookmark.
//
// If no bookmark for @uri is found then it is created.
func (b *BookmarkFile) SetAdded(uri string, added int32) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 C.time_t

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.time_t(added)

	C.g_bookmark_file_set_added(arg0, uri, added)
}

// SetAddedDateTime sets the time the bookmark for @uri was added into
// @bookmark.
//
// If no bookmark for @uri is found then it is created.
func (b *BookmarkFile) SetAddedDateTime(uri string, added *DateTime) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var arg2 *C.GDateTime

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GDateTime)(unsafe.Pointer(added.Native()))

	C.g_bookmark_file_set_added_date_time(arg0, uri, added)
}

// SetAppInfo sets the meta-data of application @name inside the list of
// applications that have registered a bookmark for @uri inside @bookmark.
//
// You should rarely use this function; use g_bookmark_file_add_application()
// and g_bookmark_file_remove_application() instead.
//
// @name can be any UTF-8 encoded string used to identify an application. @exec
// can have one of these two modifiers: "\f", which will be expanded as the
// local file name retrieved from the bookmark's URI; "\u", which will be
// expanded as the bookmark's URI. The expansion is done automatically when
// retrieving the stored command line using the
// g_bookmark_file_get_application_info() function. @count is the number of
// times the application has registered the bookmark; if is < 0, the current
// registration count will be increased by one, if is 0, the application with
// @name will be removed from the list of registered applications. @stamp is the
// Unix time of the last registration; if it is -1, the current time will be
// used.
//
// If you try to remove an application by setting its registration count to
// zero, and no bookmark for @uri is found, false is returned and @error is set
// to BOOKMARK_FILE_ERROR_URI_NOT_FOUND; similarly, in the event that no
// application @name has registered a bookmark for @uri, false is returned and
// error is set to BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED. Otherwise, if no
// bookmark for @uri is found, one is created.
func (b *BookmarkFile) SetAppInfo(uri string, name string, exec string, count int, stamp int32) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gint
	var arg5 C.time_t
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(exec))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gint(count)
	arg5 = C.time_t(stamp)

	var goerr error

	C.g_bookmark_file_set_app_info(arg0, uri, name, exec, count, stamp, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// SetApplicationInfo sets the meta-data of application @name inside the list of
// applications that have registered a bookmark for @uri inside @bookmark.
//
// You should rarely use this function; use g_bookmark_file_add_application()
// and g_bookmark_file_remove_application() instead.
//
// @name can be any UTF-8 encoded string used to identify an application. @exec
// can have one of these two modifiers: "\f", which will be expanded as the
// local file name retrieved from the bookmark's URI; "\u", which will be
// expanded as the bookmark's URI. The expansion is done automatically when
// retrieving the stored command line using the
// g_bookmark_file_get_application_info() function. @count is the number of
// times the application has registered the bookmark; if is < 0, the current
// registration count will be increased by one, if is 0, the application with
// @name will be removed from the list of registered applications. @stamp is the
// Unix time of the last registration.
//
// If you try to remove an application by setting its registration count to
// zero, and no bookmark for @uri is found, false is returned and @error is set
// to BOOKMARK_FILE_ERROR_URI_NOT_FOUND; similarly, in the event that no
// application @name has registered a bookmark for @uri, false is returned and
// error is set to BOOKMARK_FILE_ERROR_APP_NOT_REGISTERED. Otherwise, if no
// bookmark for @uri is found, one is created.
func (b *BookmarkFile) SetApplicationInfo(uri string, name string, exec string, count int, stamp *DateTime) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var arg2 *C.char
	var arg3 *C.char
	var arg4 C.int
	var arg5 *C.GDateTime
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.char)(C.CString(exec))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.int(count)
	arg5 = (*C.GDateTime)(unsafe.Pointer(stamp.Native()))

	var goerr error

	C.g_bookmark_file_set_application_info(arg0, uri, name, exec, count, stamp, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}

// SetDescription sets @description as the description of the bookmark for @uri.
//
// If @uri is nil, the description of @bookmark is set.
//
// If a bookmark for @uri cannot be found then it is created.
func (b *BookmarkFile) SetDescription(uri string, description string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(arg2))

	C.g_bookmark_file_set_description(arg0, uri, description)
}

// SetGroups sets a list of group names for the item with URI @uri. Each
// previously set group name list is removed.
//
// If @uri cannot be found then an item for it is created.
func (b *BookmarkFile) SetGroups(uri string, groups []string) {
	var arg0 *C.GBookmarkFile

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	C.g_bookmark_file_set_groups(arg0, uri, groups, length)
}

// SetIcon sets the icon for the bookmark for @uri. If @href is nil, unsets the
// currently set icon. @href can either be a full URL for the icon file or the
// icon name following the Icon Naming specification.
//
// If no bookmark for @uri is found one is created.
func (b *BookmarkFile) SetIcon(uri string, href string, mimeType string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(href))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg3))

	C.g_bookmark_file_set_icon(arg0, uri, href, mimeType)
}

// SetIsPrivate sets the private flag of the bookmark for @uri.
//
// If a bookmark for @uri cannot be found then it is created.
func (b *BookmarkFile) SetIsPrivate(uri string, isPrivate bool) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 C.gboolean

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	if isPrivate {
		arg2 = C.gboolean(1)
	}

	C.g_bookmark_file_set_is_private(arg0, uri, isPrivate)
}

// SetMIMEType sets @mime_type as the MIME type of the bookmark for @uri.
//
// If a bookmark for @uri cannot be found then it is created.
func (b *BookmarkFile) SetMIMEType(uri string, mimeType string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg2))

	C.g_bookmark_file_set_mime_type(arg0, uri, mimeType)
}

// SetModified sets the last time the bookmark for @uri was last modified.
//
// If no bookmark for @uri is found then it is created.
//
// The "modified" time should only be set when the bookmark's meta-data was
// actually changed. Every function of File that modifies a bookmark also
// changes the modification time, except for
// g_bookmark_file_set_visited_date_time().
func (b *BookmarkFile) SetModified(uri string, modified int32) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 C.time_t

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.time_t(modified)

	C.g_bookmark_file_set_modified(arg0, uri, modified)
}

// SetModifiedDateTime sets the last time the bookmark for @uri was last
// modified.
//
// If no bookmark for @uri is found then it is created.
//
// The "modified" time should only be set when the bookmark's meta-data was
// actually changed. Every function of File that modifies a bookmark also
// changes the modification time, except for
// g_bookmark_file_set_visited_date_time().
func (b *BookmarkFile) SetModifiedDateTime(uri string, modified *DateTime) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var arg2 *C.GDateTime

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GDateTime)(unsafe.Pointer(modified.Native()))

	C.g_bookmark_file_set_modified_date_time(arg0, uri, modified)
}

// SetTitle sets @title as the title of the bookmark for @uri inside the
// bookmark file @bookmark.
//
// If @uri is nil, the title of @bookmark is set.
//
// If a bookmark for @uri cannot be found then it is created.
func (b *BookmarkFile) SetTitle(uri string, title string) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg2))

	C.g_bookmark_file_set_title(arg0, uri, title)
}

// SetVisited sets the time the bookmark for @uri was last visited.
//
// If no bookmark for @uri is found then it is created.
//
// The "visited" time should only be set if the bookmark was launched, either
// using the command line retrieved by g_bookmark_file_get_application_info() or
// by the default application for the bookmark's MIME type, retrieved using
// g_bookmark_file_get_mime_type(). Changing the "visited" time does not affect
// the "modified" time.
func (b *BookmarkFile) SetVisited(uri string, visited int32) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var arg2 C.time_t

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.time_t(visited)

	C.g_bookmark_file_set_visited(arg0, uri, visited)
}

// SetVisitedDateTime sets the time the bookmark for @uri was last visited.
//
// If no bookmark for @uri is found then it is created.
//
// The "visited" time should only be set if the bookmark was launched, either
// using the command line retrieved by g_bookmark_file_get_application_info() or
// by the default application for the bookmark's MIME type, retrieved using
// g_bookmark_file_get_mime_type(). Changing the "visited" time does not affect
// the "modified" time.
func (b *BookmarkFile) SetVisitedDateTime(uri string, visited *DateTime) {
	var arg0 *C.GBookmarkFile
	var arg1 *C.char
	var arg2 *C.GDateTime

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GDateTime)(unsafe.Pointer(visited.Native()))

	C.g_bookmark_file_set_visited_date_time(arg0, uri, visited)
}

// ToData: this function outputs @bookmark as a string.
func (b *BookmarkFile) ToData() (length uint, guint8s []byte, err error) {
	var arg0 *C.GBookmarkFile
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var arg1 *C.gsize
	var goret2 []byte
	var goerr error

	cret = C.g_bookmark_file_to_data(arg0, &arg1, &errout)

	ptr.SetSlice(unsafe.Pointer(&goret2), unsafe.Pointer(cret), int(arg1))
	runtime.SetFinalizer(&goret2, func(v *[]byte) {
		C.free(ptr.Slice(unsafe.Pointer(v)))
	})
	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

	return ret1, goret2, goerr
}

// ToFile: this function outputs @bookmark into a file. The write process is
// guaranteed to be atomic by using g_file_set_contents() internally.
func (b *BookmarkFile) ToFile(filename string) error {
	var arg0 *C.GBookmarkFile
	var arg1 *C.gchar
	var errout *C.GError

	arg0 = (*C.GBookmarkFile)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var goerr error

	C.g_bookmark_file_to_file(arg0, filename, &errout)

	if errout != nil {
		goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))
		C.g_error_free(errout)
	}

}
