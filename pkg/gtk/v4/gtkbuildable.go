// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buildable_get_type()), F: marshalBuildabler},
	})
}

// BuildableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BuildableOverrider interface {
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	AddChild(builder *Builder, child *externglib.Object, typ string)
	// CustomFinished: similar to gtk_buildable_parser_finished() but is called
	// once for each custom tag handled by the buildable.
	CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagEnd: called at the end of each custom element handled by the
	// buildable.
	CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagStart: called for each unknown element under <child>.
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (BuildableParser, cgo.Handle, bool)
	ID() string
	// InternalChild retrieves the internal child called childname of the
	// buildable object.
	InternalChild(builder *Builder, childname string) *externglib.Object
	ParserFinished(builder *Builder)
	SetBuildableProperty(builder *Builder, name string, value *externglib.Value)
	SetID(id string)
}

// Buildable: GtkBuildable allows objects to extend and customize their
// deserialization from ui files.
//
// The interface includes methods for setting names and properties of objects,
// parsing custom tags and constructing child objects.
//
// The GtkBuildable interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK. The main user of this interface
// is gtk.Builder. There should be very little need for applications to call any
// of these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// GtkBuilder XML format or run any extra routines at deserialization time.
type Buildable struct {
	*externglib.Object
}

// Buildabler describes Buildable's abstract methods.
type Buildabler interface {
	externglib.Objector

	// BuildableID gets the ID of the buildable object.
	BuildableID() string
}

var _ Buildabler = (*Buildable)(nil)

func WrapBuildable(obj *externglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuildable(obj), nil
}

// BuildableID gets the ID of the buildable object.
//
// GtkBuilder sets the name based on the ID attribute of the <object> tag used
// to construct the buildable.
func (buildable *Buildable) BuildableID() string {
	var _arg0 *C.GtkBuildable // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))

	_cret = C.gtk_buildable_get_buildable_id(_arg0)
	runtime.KeepAlive(buildable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// BuildableParser: sub-parser for GtkBuildable implementations.
//
// An instance of this type is always passed by reference.
type BuildableParser struct {
	*buildableParser
}

// buildableParser is the struct that's finalized.
type buildableParser struct {
	native *C.GtkBuildableParser
}
