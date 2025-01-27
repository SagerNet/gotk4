// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_texture_get_type()), F: marshalTexturer},
	})
}

// Texture: GdkTexture is the basic element used to refer to pixel data.
//
// It is primarily meant for pixel data that will not change over multiple
// frames, and will be used for a long time.
//
// There are various ways to create GdkTexture objects from a GdkPixbuf, or a
// Cairo surface, or other pixel data.
//
// The ownership of the pixel data is transferred to the GdkTexture instance;
// you can only make a copy of it, via gdk.Texture.Download().
//
// GdkTexture is an immutable object: That means you cannot change anything
// about it other than increasing the reference count via g_object_ref().
type Texture struct {
	*externglib.Object

	Paintable
}

// Texturer describes Texture's abstract methods.
type Texturer interface {
	externglib.Objector

	// Height returns the height of the texture, in pixels.
	Height() int
	// Width returns the width of texture, in pixels.
	Width() int
	// SaveToPng: store the given texture to the filename as a PNG file.
	SaveToPng(filename string) bool
}

var _ Texturer = (*Texture)(nil)

func WrapTexture(obj *externglib.Object) *Texture {
	return &Texture{
		Object: obj,
		Paintable: Paintable{
			Object: obj,
		},
	}
}

func marshalTexturer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTexture(obj), nil
}

// NewTextureForPixbuf creates a new texture object representing the GdkPixbuf.
func NewTextureForPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Texture {
	var _arg1 *C.GdkPixbuf  // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gdk_texture_new_for_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _texture *Texture // out

	_texture = WrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _texture
}

// NewTextureFromFile creates a new texture by loading an image from a file.
//
// The file format is detected automatically. The supported formats are PNG and
// JPEG, though more formats might be available.
//
// If NULL is returned, then error will be set.
func NewTextureFromFile(file gio.Filer) (*Texture, error) {
	var _arg1 *C.GFile      // out
	var _cret *C.GdkTexture // in
	var _cerr *C.GError     // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.gdk_texture_new_from_file(_arg1, &_cerr)
	runtime.KeepAlive(file)

	var _texture *Texture // out
	var _goerr error      // out

	_texture = WrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _texture, _goerr
}

// NewTextureFromResource creates a new texture by loading an image from a
// resource.
//
// The file format is detected automatically. The supported formats are PNG and
// JPEG, though more formats might be available.
//
// It is a fatal error if resource_path does not specify a valid image resource
// and the program will abort if that happens. If you are unsure about the
// validity of a resource, use gdk.Texture.NewFromFile to load it.
func NewTextureFromResource(resourcePath string) *Texture {
	var _arg1 *C.char       // out
	var _cret *C.GdkTexture // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_new_from_resource(_arg1)
	runtime.KeepAlive(resourcePath)

	var _texture *Texture // out

	_texture = WrapTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _texture
}

// Height returns the height of the texture, in pixels.
func (texture *Texture) Height() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gdk_texture_get_height(_arg0)
	runtime.KeepAlive(texture)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Width returns the width of texture, in pixels.
func (texture *Texture) Width() int {
	var _arg0 *C.GdkTexture // out
	var _cret C.int         // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gdk_texture_get_width(_arg0)
	runtime.KeepAlive(texture)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SaveToPng: store the given texture to the filename as a PNG file.
//
// This is a utility function intended for debugging and testing. If you want
// more control over formats, proper error handling or want to store to a GFile
// or other location, you might want to look into using the gdk-pixbuf library.
func (texture *Texture) SaveToPng(filename string) bool {
	var _arg0 *C.GdkTexture // out
	var _arg1 *C.char       // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_texture_save_to_png(_arg0, _arg1)
	runtime.KeepAlive(texture)
	runtime.KeepAlive(filename)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
