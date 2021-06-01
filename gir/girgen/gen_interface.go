package girgen

import (
	"github.com/diamondburned/gotk4/gir"
)

// TODO: unexported type implementation
// TODO: methods for implementation

var interfaceTmpl = newGoTemplate(`
	{{ if .Virtuals }}
	// {{ .InterfaceName }}Overrider contains methods that are overridable. This
	// interface is a subset of the interface {{ .InterfaceName }}.
	type {{ .InterfaceName }}Overrider interface {
		{{ range .Virtuals -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}
	{{ end }}

	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{- end }}
		{{ if .Virtuals -}}
		{{ .InterfaceName }}Overrider
		{{- end }}

		{{ range .Methods -}}
		{{ if not ($.IsVirtual .Name) -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end -}}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{ end }}
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	// Wrap{{ .InterfaceName }} wraps a GObject to a type that implements interface
	// {{ .InterfaceName }}. It is primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .TypeTree.Wrap "obj" }}
	}

	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}

	{{ range .Methods }}
	{{ GoDoc .Doc 1 .Name }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type ifaceGenerator struct {
	gir.Interface
	InterfaceName string
	StructName    string

	TypeTree TypeTree
	Virtuals []callableGenerator // for interface
	Methods  []callableGenerator // for object implementation

	Ng *NamespaceGenerator
}

func newIfaceGenerator(ng *NamespaceGenerator) *ifaceGenerator {
	return &ifaceGenerator{
		Ng:       ng,
		TypeTree: *ng.TypeTree(),
	}
}

func (ig *ifaceGenerator) Use(iface gir.Interface) bool {
	ig.Interface = iface
	ig.InterfaceName = PascalToGo(iface.Name)
	ig.StructName = UnexportPascal(ig.InterfaceName)
	ig.updateMethods()

	resolved := TypeFromResult(ig.Ng, gir.TypeFindResult{Interface: &iface})
	if !ig.TypeTree.ResolveFromType(resolved) {
		ig.Ng.logln(logSkip, "interface", ig.InterfaceName, "cannot be type-resolved")
		return false
	}

	return true
}

func (ig *ifaceGenerator) updateMethods() {
	ig.Methods = callableGrow(ig.Methods, len(ig.Interface.Methods))
	ig.Virtuals = callableGrow(ig.Virtuals, len(ig.Interface.VirtualMethods))

	for _, vmethod := range ig.Interface.VirtualMethods {
		cbgen := newCallableGenerator(ig.Ng)
		if !cbgen.Use(vmethod.CallableAttrs) {
			continue
		}

		ig.Virtuals = append(ig.Virtuals, cbgen)
	}

	for _, method := range ig.Interface.Methods {
		cbgen := newCallableGenerator(ig.Ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		ig.Methods = append(ig.Methods, cbgen)
	}

	callableRenameGetters(ig.Methods)
	callableRenameGetters(ig.Virtuals)
}

// IsVirtual returns true if the given method name is a virtual method's.
func (ig *ifaceGenerator) IsVirtual(name string) bool {
	for _, vmethod := range ig.Virtuals {
		if vmethod.Name == name {
			return true
		}
	}

	return false
}

func (ng *NamespaceGenerator) generateIfaces() {
	ig := newIfaceGenerator(ng)

	for _, iface := range ng.current.Namespace.Interfaces {
		if ng.mustIgnore(iface.Name, iface.CType) {
			continue
		}

		if !ig.Use(iface) {
			continue
		}

		if iface.GLibGetType != "" {
			ng.addMarshaler(iface.GLibGetType, ig.InterfaceName)
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(interfaceTmpl, &ig)
	}
}
