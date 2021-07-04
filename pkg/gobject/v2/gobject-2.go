// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_type_plugin_get_type()), F: marshalTypePlugin},
		{T: externglib.Type(C.g_binding_get_type()), F: marshalBinding},
	})
}

// TypePlugin: the GObject type system supports dynamic loading of types. The
// Plugin interface is used to handle the lifecycle of dynamically loaded types.
// It goes as follows:
//
// 1. The type is initially introduced (usually upon loading the module the
// first time, or by your main application that knows what modules introduces
// what types), like this: |[<!-- language="C" --> new_type_id =
// g_type_register_dynamic (parent_type_id, "TypeName", new_type_plugin,
// type_flags); ]| where @new_type_plugin is an implementation of the Plugin
// interface.
//
// 2. The type's implementation is referenced, e.g. through g_type_class_ref()
// or through g_type_create_instance() (this is being called by g_object_new())
// or through one of the above done on a type derived from @new_type_id.
//
// 3. This causes the type system to load the type's implementation by calling
// g_type_plugin_use() and g_type_plugin_complete_type_info() on
// @new_type_plugin.
//
// 4. At some point the type's implementation isn't required anymore, e.g. after
// g_type_class_unref() or g_type_free_instance() (called when the reference
// count of an instance drops to zero).
//
// 5. This causes the type system to throw away the information retrieved from
// g_type_plugin_complete_type_info() and then it calls g_type_plugin_unuse() on
// @new_type_plugin.
//
// 6. Things may repeat from the second step.
//
// So basically, you need to implement a Plugin type that carries a use_count,
// once use_count goes from zero to one, you need to load the implementation to
// successfully handle the upcoming g_type_plugin_complete_type_info() call.
// Later, maybe after succeeding use/unuse calls, once use_count drops to zero,
// you can unload the implementation again. The type system makes sure to call
// g_type_plugin_use() and g_type_plugin_complete_type_info() again when the
// type is needed again.
//
// Module is an implementation of Plugin that already implements most of this
// except for the actual module loading and unloading. It even handles multiple
// registered types per module.
type TypePlugin interface {
	gextras.Objector

	// CompleteInterfaceInfo calls the @use_plugin function from the PluginClass
	// of @plugin. There should be no need to use this function outside of the
	// GObject type system itself.
	CompleteInterfaceInfo(instanceType externglib.Type, interfaceType externglib.Type, info *InterfaceInfo)
	// CompleteTypeInfo calls the @use_plugin function from the PluginClass of
	// @plugin. There should be no need to use this function outside of the
	// GObject type system itself.
	CompleteTypeInfo(gType externglib.Type, info *TypeInfo, valueTable *TypeValueTable)
	// Unuse calls the @use_plugin function from the PluginClass of @plugin.
	// There should be no need to use this function outside of the GObject type
	// system itself.
	Unuse()
	// Use calls the @use_plugin function from the PluginClass of @plugin. There
	// should be no need to use this function outside of the GObject type system
	// itself.
	Use()
}

// typePlugin implements the TypePlugin interface.
type typePlugin struct {
	gextras.Objector
}

var _ TypePlugin = (*typePlugin)(nil)

// WrapTypePlugin wraps a GObject to a type that implements
// interface TypePlugin. It is primarily used internally.
func WrapTypePlugin(obj *externglib.Object) TypePlugin {
	return typePlugin{
		Objector: obj,
	}
}

func marshalTypePlugin(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTypePlugin(obj), nil
}

func (p typePlugin) CompleteInterfaceInfo(instanceType externglib.Type, interfaceType externglib.Type, info *InterfaceInfo) {
	var _arg0 *C.GTypePlugin    // out
	var _arg1 C.GType           // out
	var _arg2 C.GType           // out
	var _arg3 *C.GInterfaceInfo // out

	_arg0 = (*C.GTypePlugin)(unsafe.Pointer(p.Native()))
	_arg1 = (C.GType)(instanceType)
	_arg2 = (C.GType)(interfaceType)
	_arg3 = (*C.GInterfaceInfo)(unsafe.Pointer(info.Native()))

	C.g_type_plugin_complete_interface_info(_arg0, _arg1, _arg2, _arg3)
}

func (p typePlugin) CompleteTypeInfo(gType externglib.Type, info *TypeInfo, valueTable *TypeValueTable) {
	var _arg0 *C.GTypePlugin     // out
	var _arg1 C.GType            // out
	var _arg2 *C.GTypeInfo       // out
	var _arg3 *C.GTypeValueTable // out

	_arg0 = (*C.GTypePlugin)(unsafe.Pointer(p.Native()))
	_arg1 = (C.GType)(gType)
	_arg2 = (*C.GTypeInfo)(unsafe.Pointer(info.Native()))
	_arg3 = (*C.GTypeValueTable)(unsafe.Pointer(valueTable.Native()))

	C.g_type_plugin_complete_type_info(_arg0, _arg1, _arg2, _arg3)
}

func (p typePlugin) Unuse() {
	var _arg0 *C.GTypePlugin // out

	_arg0 = (*C.GTypePlugin)(unsafe.Pointer(p.Native()))

	C.g_type_plugin_unuse(_arg0)
}

func (p typePlugin) Use() {
	var _arg0 *C.GTypePlugin // out

	_arg0 = (*C.GTypePlugin)(unsafe.Pointer(p.Native()))

	C.g_type_plugin_use(_arg0)
}

// Binding is the representation of a binding between a property on a #GObject
// instance (or source) and another property on another #GObject instance (or
// target). Whenever the source property changes, the same value is applied to
// the target property; for instance, the following binding:
//
//    g_object_bind_property_full (adjustment1, "value",
//                                 adjustment2, "value",
//                                 G_BINDING_BIDIRECTIONAL,
//                                 celsius_to_fahrenheit,
//                                 fahrenheit_to_celsius,
//                                 NULL, NULL);
//
// will keep the "value" property of the two adjustments in sync; the
// @celsius_to_fahrenheit function will be called whenever the "value" property
// of @adjustment1 changes and will transform the current value of the property
// before applying it to the "value" property of @adjustment2.
//
// Vice versa, the @fahrenheit_to_celsius function will be called whenever the
// "value" property of @adjustment2 changes, and will transform the current
// value of the property before applying it to the "value" property of
// @adjustment1.
//
// Note that #GBinding does not resolve cycles by itself; a cycle like
//
//    object1:propertyA -> object2:propertyB
//    object2:propertyB -> object3:propertyC
//    object3:propertyC -> object1:propertyA
//
// might lead to an infinite loop. The loop, in this particular case, can be
// avoided if the objects emit the #GObject::notify signal only if the value has
// effectively been changed. A binding is implemented using the #GObject::notify
// signal, so it is susceptible to all the various ways of blocking a signal
// emission, like g_signal_stop_emission() or g_signal_handler_block().
//
// A binding will be severed, and the resources it allocates freed, whenever
// either one of the #GObject instances it refers to are finalized, or when the
// #GBinding instance loses its last reference.
//
// Bindings for languages with garbage collection can use g_binding_unbind() to
// explicitly release a binding between the source and target properties,
// instead of relying on the last reference on the binding, source, and target
// instances to drop.
//
// #GBinding is available since GObject 2.26
type Binding interface {
	gextras.Objector
}

// binding implements the Binding class.
type binding struct {
	gextras.Objector
}

// WrapBinding wraps a GObject to the right type. It is
// primarily used internally.
func WrapBinding(obj *externglib.Object) Binding {
	return binding{
		Objector: obj,
	}
}

func marshalBinding(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBinding(obj), nil
}
