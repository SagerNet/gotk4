// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_file_get_type()), F: marshalMediaFiler},
	})
}

const MEDIA_FILE_EXTENSION_POINT_NAME = "gtk-media-file"

// MediaFileOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MediaFileOverrider interface {
	Close()
	Open()
}

// MediaFile: GtkMediaFile implements GtkMediaStream for files.
//
// This provides a simple way to play back video files with GTK.
//
// GTK provides a GIO extension point for GtkMediaFile implementations to allow
// for external implementations using various media frameworks.
//
// GTK itself includes implementations using GStreamer and ffmpeg.
type MediaFile struct {
	MediaStream
}

// MediaFiler describes MediaFile's abstract methods.
type MediaFiler interface {
	externglib.Objector

	// Clear resets the media file to be empty.
	Clear()
	// File returns the file that self is currently playing from.
	File() gio.Filer
	// InputStream returns the stream that self is currently playing from.
	InputStream() gio.InputStreamer
	// SetFile sets the GtkMediaFile to play the given file.
	SetFile(file gio.Filer)
	// SetFilename sets the `GtkMediaFile to play the given file.
	SetFilename(filename string)
	// SetInputStream sets the GtkMediaFile to play the given stream.
	SetInputStream(stream gio.InputStreamer)
	// SetResource sets the `GtkMediaFile to play the given resource.
	SetResource(resourcePath string)
}

var _ MediaFiler = (*MediaFile)(nil)

func WrapMediaFile(obj *externglib.Object) *MediaFile {
	return &MediaFile{
		MediaStream: MediaStream{
			Object: obj,
			Paintable: gdk.Paintable{
				Object: obj,
			},
		},
	}
}

func marshalMediaFiler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaFile(obj), nil
}

// NewMediaFile creates a new empty media file.
func NewMediaFile() *MediaFile {
	var _cret *C.GtkMediaStream // in

	_cret = C.gtk_media_file_new()

	var _mediaFile *MediaFile // out

	_mediaFile = WrapMediaFile(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mediaFile
}

// NewMediaFileForFile creates a new media file to play file.
func NewMediaFileForFile(file gio.Filer) *MediaFile {
	var _arg1 *C.GFile          // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.gtk_media_file_new_for_file(_arg1)
	runtime.KeepAlive(file)

	var _mediaFile *MediaFile // out

	_mediaFile = WrapMediaFile(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mediaFile
}

// NewMediaFileForFilename creates a new media file for the given filename.
//
// This is a utility function that converts the given filename to a GFile and
// calls gtk.MediaFile.NewForFile.
func NewMediaFileForFilename(filename string) *MediaFile {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_filename(_arg1)
	runtime.KeepAlive(filename)

	var _mediaFile *MediaFile // out

	_mediaFile = WrapMediaFile(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mediaFile
}

// NewMediaFileForInputStream creates a new media file to play stream.
//
// If you want the resulting media to be seekable, the stream should implement
// the GSeekable interface.
func NewMediaFileForInputStream(stream gio.InputStreamer) *MediaFile {
	var _arg1 *C.GInputStream   // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.gtk_media_file_new_for_input_stream(_arg1)
	runtime.KeepAlive(stream)

	var _mediaFile *MediaFile // out

	_mediaFile = WrapMediaFile(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mediaFile
}

// NewMediaFileForResource creates a new new media file for the given resource.
//
// This is a utility function that converts the given resource to a GFile and
// calls gtk.MediaFile.NewForFile.
func NewMediaFileForResource(resourcePath string) *MediaFile {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _mediaFile *MediaFile // out

	_mediaFile = WrapMediaFile(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _mediaFile
}

// Clear resets the media file to be empty.
func (self *MediaFile) Clear() {
	var _arg0 *C.GtkMediaFile // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))

	C.gtk_media_file_clear(_arg0)
	runtime.KeepAlive(self)
}

// File returns the file that self is currently playing from.
//
// When self is not playing or not playing from a file, NULL is returned.
func (self *MediaFile) File() gio.Filer {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GFile        // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_media_file_get_file(_arg0)
	runtime.KeepAlive(self)

	var _file gio.Filer // out

	if _cret != nil {
		_file = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Filer)
	}

	return _file
}

// InputStream returns the stream that self is currently playing from.
//
// When self is not playing or not playing from a stream, NULL is returned.
func (self *MediaFile) InputStream() gio.InputStreamer {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GInputStream // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_media_file_get_input_stream(_arg0)
	runtime.KeepAlive(self)

	var _inputStream gio.InputStreamer // out

	if _cret != nil {
		_inputStream = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.InputStreamer)
	}

	return _inputStream
}

// SetFile sets the GtkMediaFile to play the given file.
//
// If any file is still playing, stop playing it.
func (self *MediaFile) SetFile(file gio.Filer) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GFile        // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))
	if file != nil {
		_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))
	}

	C.gtk_media_file_set_file(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetFilename sets the `GtkMediaFile to play the given file.
//
// This is a utility function that converts the given filename to a GFile and
// calls gtk.MediaFile.SetFile().
func (self *MediaFile) SetFilename(filename string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))
	if filename != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_media_file_set_filename(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filename)
}

// SetInputStream sets the GtkMediaFile to play the given stream.
//
// If anything is still playing, stop playing it.
//
// Full control about the stream is assumed for the duration of playback. The
// stream will not be closed.
func (self *MediaFile) SetInputStream(stream gio.InputStreamer) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GInputStream // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))
	if stream != nil {
		_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))
	}

	C.gtk_media_file_set_input_stream(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stream)
}

// SetResource sets the `GtkMediaFile to play the given resource.
//
// This is a utility function that converts the given resource_path to a GFile
// and calls gtk.MediaFile.SetFile().
func (self *MediaFile) SetResource(resourcePath string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(self.Native()))
	if resourcePath != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_media_file_set_resource(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(resourcePath)
}
