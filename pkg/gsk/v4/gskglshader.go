// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/graphene"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_gl_shader_get_type()), F: marshalGLShader},
		{T: externglib.Type(C.gsk_shader_args_builder_get_type()), F: marshalShaderArgsBuilder},
	})
}

// GLShader: an object representing a GL shader program.
type GLShader interface {
	gextras.Objector

	// Compile tries to compile the @shader for the given @renderer, and reports
	// false with an error if there is a problem. You should use this function
	// before relying on the shader for rendering and use a fallback with a
	// simpler shader or without shaders if it fails.
	//
	// Note that this will modify the rendering state (for example change the
	// current GL context) and requires the renderer to be set up. This means
	// that the widget has to be realized. Commonly you want to call this from
	// the realize signal of a widget, or during widget snapshot.
	Compile(renderer Renderer) error
	// FindUniformByName looks for a uniform by the name @name, and returns the
	// index of the uniform, or -1 if it was not found.
	FindUniformByName(name string) int
	// ArgBool gets the value of the uniform @idx in the @args block. The
	// uniform must be of bool type.
	ArgBool(args *glib.Bytes, idx int) bool
	// ArgFloat gets the value of the uniform @idx in the @args block. The
	// uniform must be of float type.
	ArgFloat(args *glib.Bytes, idx int) float32
	// ArgInt gets the value of the uniform @idx in the @args block. The uniform
	// must be of int type.
	ArgInt(args *glib.Bytes, idx int) int32
	// ArgUint gets the value of the uniform @idx in the @args block. The
	// uniform must be of uint type.
	ArgUint(args *glib.Bytes, idx int) uint32
	// ArgVec2 gets the value of the uniform @idx in the @args block. The
	// uniform must be of vec2 type.
	ArgVec2(args *glib.Bytes, idx int, outValue *graphene.Vec2)
	// ArgVec3 gets the value of the uniform @idx in the @args block. The
	// uniform must be of vec3 type.
	ArgVec3(args *glib.Bytes, idx int, outValue *graphene.Vec3)
	// ArgVec4 gets the value of the uniform @idx in the @args block. The
	// uniform must be of vec4 type.
	ArgVec4(args *glib.Bytes, idx int, outValue *graphene.Vec4)
	// ArgsSize: get the size of the data block used to specify arguments for
	// this shader.
	ArgsSize() uint
	// NTextures returns the number of textures that the shader requires.
	//
	// This can be used to check that the a passed shader works in your usecase.
	// It is determined by looking at the highest u_textureN value that the
	// shader defines.
	NTextures() int
	// NUniforms: get the number of declared uniforms for this shader.
	NUniforms() int
	// Resource gets the resource path for the GLSL sourcecode being used to
	// render this shader.
	Resource() string
	// Source gets the GLSL sourcecode being used to render this shader.
	Source() *glib.Bytes
	// UniformName: get the name of the declared uniform for this shader at
	// index @idx.
	UniformName(idx int) string
	// UniformOffset: get the offset into the data block where data for this
	// uniforms is stored.
	UniformOffset(idx int) int
	// UniformType: get the type of the declared uniform for this shader at
	// index @idx.
	UniformType(idx int) GLUniformType
}

// glShader implements the GLShader interface.
type glShader struct {
	gextras.Objector
}

var _ GLShader = (*glShader)(nil)

// WrapGLShader wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLShader(obj *externglib.Object) GLShader {
	return GLShader{
		Objector: obj,
	}
}

func marshalGLShader(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLShader(obj), nil
}

// NewGLShaderFromBytes constructs a class GLShader.
func NewGLShaderFromBytes(sourcecode *glib.Bytes) GLShader {
	var arg1 *C.GBytes

	arg1 = (*C.GBytes)(unsafe.Pointer(sourcecode.Native()))

	var cret C.GskGLShader
	var ret1 GLShader

	cret = C.gsk_gl_shader_new_from_bytes(sourcecode)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GLShader)

	return ret1
}

// NewGLShaderFromResource constructs a class GLShader.
func NewGLShaderFromResource(resourcePath string) GLShader {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GskGLShader
	var ret1 GLShader

	cret = C.gsk_gl_shader_new_from_resource(resourcePath)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GLShader)

	return ret1
}

// Compile tries to compile the @shader for the given @renderer, and reports
// false with an error if there is a problem. You should use this function
// before relying on the shader for rendering and use a fallback with a
// simpler shader or without shaders if it fails.
//
// Note that this will modify the rendering state (for example change the
// current GL context) and requires the renderer to be set up. This means
// that the widget has to be realized. Commonly you want to call this from
// the realize signal of a widget, or during widget snapshot.
func (s glShader) Compile(renderer Renderer) error {
	var arg0 *C.GskGLShader
	var arg1 *C.GskRenderer

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GskRenderer)(unsafe.Pointer(renderer.Native()))

	var errout *C.GError
	var goerr error

	C.gsk_gl_shader_compile(arg0, renderer, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))

	return goerr
}

// FindUniformByName looks for a uniform by the name @name, and returns the
// index of the uniform, or -1 if it was not found.
func (s glShader) FindUniformByName(name string) int {
	var arg0 *C.GskGLShader
	var arg1 *C.char

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.int
	var ret1 int

	cret = C.gsk_gl_shader_find_uniform_by_name(arg0, name)

	ret1 = C.int(cret)

	return ret1
}

// ArgBool gets the value of the uniform @idx in the @args block. The
// uniform must be of bool type.
func (s glShader) ArgBool(args *glib.Bytes, idx int) bool {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)

	var cret C.gboolean
	var ret1 bool

	cret = C.gsk_gl_shader_get_arg_bool(arg0, args, idx)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ArgFloat gets the value of the uniform @idx in the @args block. The
// uniform must be of float type.
func (s glShader) ArgFloat(args *glib.Bytes, idx int) float32 {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)

	var cret C.float
	var ret1 float32

	cret = C.gsk_gl_shader_get_arg_float(arg0, args, idx)

	ret1 = C.float(cret)

	return ret1
}

// ArgInt gets the value of the uniform @idx in the @args block. The uniform
// must be of int type.
func (s glShader) ArgInt(args *glib.Bytes, idx int) int32 {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)

	var cret C.gint32
	var ret1 int32

	cret = C.gsk_gl_shader_get_arg_int(arg0, args, idx)

	ret1 = C.gint32(cret)

	return ret1
}

// ArgUint gets the value of the uniform @idx in the @args block. The
// uniform must be of uint type.
func (s glShader) ArgUint(args *glib.Bytes, idx int) uint32 {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)

	var cret C.guint32
	var ret1 uint32

	cret = C.gsk_gl_shader_get_arg_uint(arg0, args, idx)

	ret1 = C.guint32(cret)

	return ret1
}

// ArgVec2 gets the value of the uniform @idx in the @args block. The
// uniform must be of vec2 type.
func (s glShader) ArgVec2(args *glib.Bytes, idx int, outValue *graphene.Vec2) {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int
	var arg3 *C.graphene_vec2_t

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)
	arg3 = (*C.graphene_vec2_t)(unsafe.Pointer(outValue.Native()))

	C.gsk_gl_shader_get_arg_vec2(arg0, args, idx, outValue)
}

// ArgVec3 gets the value of the uniform @idx in the @args block. The
// uniform must be of vec3 type.
func (s glShader) ArgVec3(args *glib.Bytes, idx int, outValue *graphene.Vec3) {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int
	var arg3 *C.graphene_vec3_t

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)
	arg3 = (*C.graphene_vec3_t)(unsafe.Pointer(outValue.Native()))

	C.gsk_gl_shader_get_arg_vec3(arg0, args, idx, outValue)
}

// ArgVec4 gets the value of the uniform @idx in the @args block. The
// uniform must be of vec4 type.
func (s glShader) ArgVec4(args *glib.Bytes, idx int, outValue *graphene.Vec4) {
	var arg0 *C.GskGLShader
	var arg1 *C.GBytes
	var arg2 C.int
	var arg3 *C.graphene_vec4_t

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GBytes)(unsafe.Pointer(args.Native()))
	arg2 = C.int(idx)
	arg3 = (*C.graphene_vec4_t)(unsafe.Pointer(outValue.Native()))

	C.gsk_gl_shader_get_arg_vec4(arg0, args, idx, outValue)
}

// ArgsSize: get the size of the data block used to specify arguments for
// this shader.
func (s glShader) ArgsSize() uint {
	var arg0 *C.GskGLShader

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))

	var cret C.gsize
	var ret1 uint

	cret = C.gsk_gl_shader_get_args_size(arg0)

	ret1 = C.gsize(cret)

	return ret1
}

// NTextures returns the number of textures that the shader requires.
//
// This can be used to check that the a passed shader works in your usecase.
// It is determined by looking at the highest u_textureN value that the
// shader defines.
func (s glShader) NTextures() int {
	var arg0 *C.GskGLShader

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))

	var cret C.int
	var ret1 int

	cret = C.gsk_gl_shader_get_n_textures(arg0)

	ret1 = C.int(cret)

	return ret1
}

// NUniforms: get the number of declared uniforms for this shader.
func (s glShader) NUniforms() int {
	var arg0 *C.GskGLShader

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))

	var cret C.int
	var ret1 int

	cret = C.gsk_gl_shader_get_n_uniforms(arg0)

	ret1 = C.int(cret)

	return ret1
}

// Resource gets the resource path for the GLSL sourcecode being used to
// render this shader.
func (s glShader) Resource() string {
	var arg0 *C.GskGLShader

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gsk_gl_shader_get_resource(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Source gets the GLSL sourcecode being used to render this shader.
func (s glShader) Source() *glib.Bytes {
	var arg0 *C.GskGLShader

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))

	var cret *C.GBytes
	var ret1 *glib.Bytes

	cret = C.gsk_gl_shader_get_source(arg0)

	ret1 = glib.WrapBytes(unsafe.Pointer(cret))

	return ret1
}

// UniformName: get the name of the declared uniform for this shader at
// index @idx.
func (s glShader) UniformName(idx int) string {
	var arg0 *C.GskGLShader
	var arg1 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = C.int(idx)

	var cret *C.char
	var ret1 string

	cret = C.gsk_gl_shader_get_uniform_name(arg0, idx)

	ret1 = C.GoString(cret)

	return ret1
}

// UniformOffset: get the offset into the data block where data for this
// uniforms is stored.
func (s glShader) UniformOffset(idx int) int {
	var arg0 *C.GskGLShader
	var arg1 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = C.int(idx)

	var cret C.int
	var ret1 int

	cret = C.gsk_gl_shader_get_uniform_offset(arg0, idx)

	ret1 = C.int(cret)

	return ret1
}

// UniformType: get the type of the declared uniform for this shader at
// index @idx.
func (s glShader) UniformType(idx int) GLUniformType {
	var arg0 *C.GskGLShader
	var arg1 C.int

	arg0 = (*C.GskGLShader)(unsafe.Pointer(s.Native()))
	arg1 = C.int(idx)

	var cret C.GskGLUniformType
	var ret1 GLUniformType

	cret = C.gsk_gl_shader_get_uniform_type(arg0, idx)

	ret1 = GLUniformType(cret)

	return ret1
}

// ShaderArgsBuilder: an object to build the uniforms data for a GLShader.
type ShaderArgsBuilder struct {
	native C.GskShaderArgsBuilder
}

// WrapShaderArgsBuilder wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapShaderArgsBuilder(ptr unsafe.Pointer) *ShaderArgsBuilder {
	if ptr == nil {
		return nil
	}

	return (*ShaderArgsBuilder)(ptr)
}

func marshalShaderArgsBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapShaderArgsBuilder(unsafe.Pointer(b)), nil
}

// NewShaderArgsBuilder constructs a struct ShaderArgsBuilder.
func NewShaderArgsBuilder(shader GLShader, initialValues *glib.Bytes) *ShaderArgsBuilder {
	var arg1 *C.GskGLShader
	var arg2 *C.GBytes

	arg1 = (*C.GskGLShader)(unsafe.Pointer(shader.Native()))
	arg2 = (*C.GBytes)(unsafe.Pointer(initialValues.Native()))

	var cret *C.GskShaderArgsBuilder
	var ret1 *ShaderArgsBuilder

	cret = C.gsk_shader_args_builder_new(shader, initialValues)

	ret1 = WrapShaderArgsBuilder(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ShaderArgsBuilder) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Native returns the underlying C source pointer.
func (s *ShaderArgsBuilder) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// FreeToArgs creates a new #GBytes args from the current state of the given
// @builder, and frees the @builder instance. Any uniforms of the shader that
// have not been explicitly set on the @builder are zero-initialized.
func (b *ShaderArgsBuilder) FreeToArgs() *glib.Bytes {
	var arg0 *C.GskShaderArgsBuilder

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GBytes
	var ret1 *glib.Bytes

	cret = C.gsk_shader_args_builder_free_to_args(arg0)

	ret1 = glib.WrapBytes(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.Bytes) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Ref increases the reference count of a ShaderArgsBuilder by one.
func (b *ShaderArgsBuilder) Ref() *ShaderArgsBuilder {
	var arg0 *C.GskShaderArgsBuilder

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GskShaderArgsBuilder
	var ret1 *ShaderArgsBuilder

	cret = C.gsk_shader_args_builder_ref(arg0)

	ret1 = WrapShaderArgsBuilder(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *ShaderArgsBuilder) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// SetBool sets the value of the uniform @idx. The uniform must be of bool type.
func (b *ShaderArgsBuilder) SetBool(idx int, value bool) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 C.gboolean

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	if value {
		arg2 = C.gboolean(1)
	}

	C.gsk_shader_args_builder_set_bool(arg0, idx, value)
}

// SetFloat sets the value of the uniform @idx. The uniform must be of float
// type.
func (b *ShaderArgsBuilder) SetFloat(idx int, value float32) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 C.float

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = C.float(value)

	C.gsk_shader_args_builder_set_float(arg0, idx, value)
}

// SetInt sets the value of the uniform @idx. The uniform must be of int type.
func (b *ShaderArgsBuilder) SetInt(idx int, value int32) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 C.gint32

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = C.gint32(value)

	C.gsk_shader_args_builder_set_int(arg0, idx, value)
}

// SetUint sets the value of the uniform @idx. The uniform must be of uint type.
func (b *ShaderArgsBuilder) SetUint(idx int, value uint32) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 C.guint32

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = C.guint32(value)

	C.gsk_shader_args_builder_set_uint(arg0, idx, value)
}

// SetVec2 sets the value of the uniform @idx. The uniform must be of vec2 type.
func (b *ShaderArgsBuilder) SetVec2(idx int, value *graphene.Vec2) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 *C.graphene_vec2_t

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = (*C.graphene_vec2_t)(unsafe.Pointer(value.Native()))

	C.gsk_shader_args_builder_set_vec2(arg0, idx, value)
}

// SetVec3 sets the value of the uniform @idx. The uniform must be of vec3 type.
func (b *ShaderArgsBuilder) SetVec3(idx int, value *graphene.Vec3) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 *C.graphene_vec3_t

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = (*C.graphene_vec3_t)(unsafe.Pointer(value.Native()))

	C.gsk_shader_args_builder_set_vec3(arg0, idx, value)
}

// SetVec4 sets the value of the uniform @idx. The uniform must be of vec4 type.
func (b *ShaderArgsBuilder) SetVec4(idx int, value *graphene.Vec4) {
	var arg0 *C.GskShaderArgsBuilder
	var arg1 C.int
	var arg2 *C.graphene_vec4_t

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.int(idx)
	arg2 = (*C.graphene_vec4_t)(unsafe.Pointer(value.Native()))

	C.gsk_shader_args_builder_set_vec4(arg0, idx, value)
}

// ToArgs creates a new #GBytes args from the current state of the given
// @builder. Any uniforms of the shader that have not been explicitly set on the
// @builder are zero-initialized.
//
// The given ShaderArgsBuilder is reset once this function returns; you cannot
// call this function multiple times on the same @builder instance.
//
// This function is intended primarily for bindings. C code should use
// gsk_shader_args_builder_free_to_args().
func (b *ShaderArgsBuilder) ToArgs() *glib.Bytes {
	var arg0 *C.GskShaderArgsBuilder

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GBytes
	var ret1 *glib.Bytes

	cret = C.gsk_shader_args_builder_to_args(arg0)

	ret1 = glib.WrapBytes(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *glib.Bytes) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Unref decreases the reference count of a ShaderArgBuilder by one. If the
// resulting reference count is zero, frees the builder.
func (b *ShaderArgsBuilder) Unref() {
	var arg0 *C.GskShaderArgsBuilder

	arg0 = (*C.GskShaderArgsBuilder)(unsafe.Pointer(b.Native()))

	C.gsk_shader_args_builder_unref(arg0)
}