// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.g_file_icon_get_type()), F: marshalFileIcon},
	})
}

// FileIcon specifies an icon by pointing to an image file to be used as icon.
type FileIcon interface {
	gextras.Objector

	// AsIcon casts the class to the Icon interface.
	AsIcon() Icon
	// AsLoadableIcon casts the class to the LoadableIcon interface.
	AsLoadableIcon() LoadableIcon

	// File gets the #GFile associated with the given @icon.
	File() File
}

// fileIcon implements the FileIcon class.
type fileIcon struct {
	gextras.Objector
}

// WrapFileIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileIcon(obj *externglib.Object) FileIcon {
	return fileIcon{
		Objector: obj,
	}
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileIcon(obj), nil
}

// NewFileIcon creates a new icon for a file.
func NewFileIcon(file File) FileIcon {
	var _arg1 *C.GFile // out
	var _cret *C.GIcon // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.g_file_icon_new(_arg1)

	var _fileIcon FileIcon // out

	_fileIcon = WrapFileIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileIcon
}

func (f fileIcon) AsIcon() Icon {
	return WrapIcon(gextras.InternObject(f))
}

func (f fileIcon) AsLoadableIcon() LoadableIcon {
	return WrapLoadableIcon(gextras.InternObject(f))
}

func (i fileIcon) File() File {
	var _arg0 *C.GFileIcon // out
	var _cret *C.GFile     // in

	_arg0 = (*C.GFileIcon)(unsafe.Pointer(i.Native()))

	_cret = C.g_file_icon_get_file(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(File)

	return _file
}
