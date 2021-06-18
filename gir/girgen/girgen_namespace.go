package girgen

import (
	"bytes"
	"fmt"
	"go/format"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
	"github.com/pkg/errors"
)

// NamespaceGenerator is a generator for a specific namespace.
type NamespaceGenerator struct {
	pen     *pen.PaperBuffer
	gen     *Generator
	current *gir.NamespaceFindResult
	pkgPath string
	pkgName string

	marshalers []string
	SideEffects
}

// NewNamespaceGenerator creates a new NamespaceGenerator.
func NewNamespaceGenerator(g *Generator, res *gir.NamespaceFindResult) *NamespaceGenerator {
	return &NamespaceGenerator{
		pen:     pen.NewPaperBufferSize(5 * 1024 * 1024), // 5MB
		gen:     g,
		current: res,
		pkgPath: g.ModPath(res.Namespace),
		pkgName: gir.GoNamespace(res.Namespace),
	}
}

// Namespace returns the generator's namespace that includes the repository it's
// in.
func (ng *NamespaceGenerator) Namespace() *gir.NamespaceFindResult {
	return ng.current
}

// PackageName returns the current namespace's package name.
func (ng *NamespaceGenerator) PackageName() string { return ng.pkgName }

func (ng *NamespaceGenerator) Logln(level LogLevel, v ...interface{}) {
	v = append(v, nil)
	copy(v[1:], v) // shift rightwards once
	v[0] = ng.logPrefix() + ":"

	ng.gen.Logln(level, v...)
}

func (ng *NamespaceGenerator) logPrefix() string {
	return fmt.Sprintf("package %s/v%s", ng.pkgName, gir.MajorVersion(ng.current.Namespace.Version))
}

// TypeTree returns a new type tree for resolving.
func (ng *NamespaceGenerator) TypeTree() TypeTree {
	return NewTypeTree(ng)
}

// WrapClass generates the class value that wraps the given object.
func (ng *NamespaceGenerator) WrapClass(typ *ResolvedType, obj string) (WrapOutput, bool) {
	tree := NewTypeTree(ng)
	if !tree.ResolveFromType(typ) {
		return WrapOutput{}, false
	}
	return tree.Wrap(obj), true
}

// Generate generates the current namespace. It returns a filesystem consisting
// of only files. For correctness, the caller should WalkDir at root.
func (ng *NamespaceGenerator) Generate() ([]byte, error) {
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateCallbacks()
	ng.generateFuncs()
	ng.generateIfaces()
	ng.generateClasses()
	ng.generateRecords()

	if len(ng.marshalers) > 0 {
		ng.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")
	}

	if ng.CallbackDelete {
		ng.addImportInternal("box")
		ng.addCallbackHeader("extern void callbackDelete(gpointer);")
	}

	var out bytes.Buffer
	// Preallocate 10KB + existing buffers.
	out.Grow(10*1024 + ng.pen.Len())

	pen := pen.NewPen(&out)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.EmptyLine()

	pen.Words("package", ng.pkgName)
	pen.EmptyLine()

	if len(ng.Imports) > 0 {
		builtin := make([]string, 0, len(ng.Imports))
		externs := make([]string, 0, len(ng.Imports))

		for path, alias := range ng.Imports {
			// Skip importing the current package.
			if path == ng.pkgPath {
				continue
			}

			if !strings.Contains(path, "/") {
				builtin = append(builtin, makeImport(path, alias))
			} else {
				externs = append(externs, makeImport(path, alias))
			}
		}

		sort.Strings(builtin)
		sort.Strings(externs)

		pen.Words("import (")

		for _, str := range builtin {
			pen.Words(str)
		}
		if len(builtin) > 0 && len(externs) > 0 {
			pen.EmptyLine()
		}
		for _, str := range externs {
			pen.Words(str)
		}

		pen.Line(")")
		pen.EmptyLine()
		pen.EmptyLine()
	}

	pen.Words(append([]string{"// #cgo pkg-config:"}, ng.pkgconfig()...)...)
	pen.Words("// #cgo CFLAGS: -Wno-deprecated-declarations")

	for _, cIncl := range ng.cIncludes() {
		pen.Linef("// #include <%s>", cIncl)
	}

	if len(ng.Callbacks) > 0 {
		pen.Words("//")
		for _, callback := range ng.callbackHeaders() {
			pen.Words("//", callback)
		}
	}

	pen.Words(`import "C"`)
	pen.EmptyLine()

	if len(ng.marshalers) > 0 {
		pen.Words("func init() {")
		pen.Words("  externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{")

		for _, marshaler := range ng.marshalers {
			pen.Words("      " + marshaler)
		}

		pen.Words("  })")
		pen.Words("}")
		pen.EmptyLine()
	}

	// Only write the Go definition in one file.
	if ng.CallbackDelete {
		pen.Words("//export callbackDelete")
		pen.Words("func callbackDelete(ptr C.gpointer) {")
		pen.Words("  box.Delete(box.Callback, uintptr(ptr))")
		pen.Words("}")
		pen.EmptyLine()
	}

	pen.WriteString(ng.pen.String())

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		ng.Logln(LogError, "failed to fmt pkg", ng.pkgName)
		return out.Bytes(), errors.Wrap(err, "fmt "+ng.logPrefix())
	}

	return formatted, nil
}

func makeImport(importPath, alias string) string {
	pathBase := path.Base(importPath)

	// Check if the base is a version part.
	if strings.HasPrefix(pathBase, "v") {
		_, err := strconv.Atoi(strings.TrimPrefix(pathBase, "v"))
		if err == nil {
			// Valid version part. Trim it.
			pathBase = path.Base(path.Dir(importPath))
		}
	}

	if alias == "" || alias == pathBase {
		return strconv.Quote(importPath)
	}

	// Only use the import alias if it's provided and does not match the base
	// name of the import path for idiomaticity.
	return alias + " " + strconv.Quote(importPath)
}

func (ng *NamespaceGenerator) cIncludes() []string {
	includes := make([]string, 0,
		len(ng.CIncludes)+len(ng.current.Repository.CIncludes))

	for cIncl := range ng.CIncludes {
		includes = append(includes, cIncl)
	}
	for _, cIncl := range ng.current.Repository.CIncludes {
		includes = append(includes, cIncl.Name)
	}

	sort.Strings(includes)
	return includes
}

// callbackHeaders returns the sorted C callback headers.
func (ng *NamespaceGenerator) callbackHeaders() []string {
	headers := make([]string, 0, len(ng.Callbacks))
	for callback := range ng.Callbacks {
		headers = append(headers, callback)
	}

	sort.Strings(headers)
	return headers
}

// mustIgnore checks the generator's filters to see if the given girType in this
// namespace should be ignored.
func (ng *NamespaceGenerator) mustIgnore(girName, cType *string) (ignore bool) {
	girType := ensureNamespace(ng.Namespace(), *girName)
	hadNamespace := girType == *girName

	names := FilterTypeName{girType, *cType}

	for _, filter := range ng.gen.Filters {
		// Filter returns keep=false.
		if !filter.Filter(ng, &names) {
			ng.Logln(LogDebug, "ignoring", girType)
			return true
		}
	}

	if hadNamespace {
		*girName = names.GIRType
	} else {
		*girName = names.Name()
	}

	*cType = names.CType

	return false
}

// mustIgnoreC is similar to mustIgnore but only works on C types.
func (ng *NamespaceGenerator) mustIgnoreC(cType string) (ignore bool) {
	nul := "\x00"
	return ng.mustIgnore(&nul, &cType)
}

// fullGIR returns the full GIR type name if it doesn't contain a namespace.
func (ng *NamespaceGenerator) fullGIR(girType string) string {
	// Skip builtin types.
	_, isBuiltin := girToBuiltin[girType]
	if isBuiltin {
		return girType
	}

	if !strings.Contains(girType, ".") {
		return ng.current.Namespace.Name + "." + girType
	}
	return girType
}

// pkgconfig returns the current repository's pkg-config names.
func (ng *NamespaceGenerator) pkgconfig() []string {
	foundRoot := false
	pkgs := make([]string, 0, len(ng.current.Repository.Packages)+1)

	for _, pkg := range ng.current.Repository.Packages {
		if pkg.Name == ng.current.Repository.Pkg {
			foundRoot = true
		}

		pkgs = append(pkgs, pkg.Name)
	}

	if !foundRoot {
		pkgs = append(pkgs, ng.current.Repository.Pkg)
	}

	return pkgs
}

// addMarshaler adds the type marshaler into the init header. It also adds
// imports.
func (ng *NamespaceGenerator) addMarshaler(glibGetType, goName string) {
	ng.marshalers = append(ng.marshalers, fmt.Sprintf(
		`{T: externglib.Type(C.%s()), F: marshal%s},`, glibGetType, goName,
	))
	// Need this for g_value functions inside marshal.
	ng.needsGLibObject()
	// Need this for the pointer cast.
	ng.addImport("unsafe")
}

// applyConvertedFxs applies all side effects of the given list of type converted
// results.
func (ng *NamespaceGenerator) applyConvertedFxs(results []ValueConverted) {
	for _, result := range results {
		result.ApplySideEffects(&ng.SideEffects)
	}
}

// SideEffects describes the side effects of the conversion, such as
// importing new things or modifying the Cgo preamble.
type SideEffects struct {
	Imports        map[string]string
	CIncludes      map[string]struct{}
	Packages       map[string]struct{} // for pkg-config
	Callbacks      map[string]struct{}
	CallbackDelete bool
}

const internalImportPath = "github.com/diamondburned/gotk4/core"

func (sides *SideEffects) addImportInternal(internal string) {
	sides.addImport(internalImportPath + "/" + internal)
}

func (sides *SideEffects) addImport(path string) {
	sides.addImportAlias(path, "")
}

func (sides *SideEffects) addImportAlias(path, alias string) {
	if sides.Imports == nil {
		sides.Imports = map[string]string{}
	}

	sides.Imports[path] = alias
}

// needsExternGLib adds the external gotk3/glib import.
func (sides *SideEffects) needsExternGLib() {
	sides.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")
}

func (sides *SideEffects) importResolved(resolved *ResolvedType) {
	if resolved == nil {
		return
	}
	if resolved.Import.Path != "" {
		sides.addImportAlias(resolved.Import.Path, resolved.Import.Package)
	}
	if resolved.IsCallback() {
		sides.addCallbackHeader(CallbackCHeader(resolved.Extern.Result.Callback))
	}
}

func (sides *SideEffects) addCallback(callback *gir.Callback) {
	sides.addCallbackHeader(CallbackCHeader(callback))
}

func (sides *SideEffects) addCallbackHeader(header string) {
	if sides.Callbacks == nil {
		sides.Callbacks = map[string]struct{}{}
	}

	sides.Callbacks[header] = struct{}{}
}

// addPackage adds a pkg-config package.
func (sides *SideEffects) addPackage(pkg string) {
	if sides.Packages == nil {
		sides.Packages = map[string]struct{}{}
	}

	sides.Packages[pkg] = struct{}{}
}

// includeC adds a C header file into the cgo preamble.
func (sides *SideEffects) includeC(include string) {
	if sides.CIncludes == nil {
		sides.CIncludes = map[string]struct{}{}
	}

	sides.CIncludes[include] = struct{}{}
}

// needsCbool adds the C stdbool.h include.
func (sides *SideEffects) needsCbool() {
	sides.includeC("stdbool.h")
}

// needsGLibObject adds the glib-object.h include and the glib-2.0 package.
func (sides *SideEffects) needsGLibObject() {
	// Need this for g_value_get_boxed.
	sides.includeC("glib-object.h")
	// Need this for the above header.
	sides.addPackage("glib-2.0")
}

// ApplySideEffects applies the side effects of the conversion. The caller is
// responsible for calling this.
func (sides *SideEffects) ApplySideEffects(dst *SideEffects) {
	if sides.CallbackDelete {
		dst.CallbackDelete = true
	}
	for path, alias := range sides.Imports {
		dst.addImportAlias(path, alias)
	}
	for callback := range sides.Callbacks {
		dst.addCallbackHeader(callback)
	}
	for cIncl := range sides.CIncludes {
		dst.includeC(cIncl)
	}
	for pkg := range sides.Packages {
		dst.addPackage(pkg)
	}
}
