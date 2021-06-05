package girgen

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// TypeTree is a structure for a type that is resolved to the lowest level of
// inheritance.
type TypeTree struct {
	Resolved *ResolvedType
	Requires []TypeTree

	// Level sets the maximum recursion level to go. It only applies if set
	// to something more than -1.
	Level int

	res TypeResolver
}

// TypeTree returns a new type tree for resolving.
func (ng *NamespaceGenerator) TypeTree() TypeTree {
	return TypeTree{res: ng, Level: -1}
}

// TypeTree returns a new type tree for resolving.
func (fg *FileGenerator) TypeTree() TypeTree {
	return TypeTree{res: fg, Level: -1}
}

func (tree *TypeTree) reset() {
	// Zero out the fields to prevent dangling pointers.
	for i := range tree.Requires {
		tree.Requires[i] = TypeTree{}
	}

	tree.Resolved = nil
	tree.Requires = tree.Requires[:0]
}

// IsInterface returns true if the current type in the tree is an interface.
func (tree *TypeTree) IsInterface() bool {
	return tree.Resolved.Extern != nil && tree.Resolved.Extern.Result.Interface != nil
}

// Resolve resolves the given toplevel type into the TypeTree, overriding the
// Resolved and Requires fields. True is returned if the tree is successfully
// resolved.
func (tree *TypeTree) Resolve(toplevel string) bool {
	resolved := ResolveTypeName(tree.res, toplevel)
	if resolved == nil {
		tree.reset()
		return false
	}

	return tree.ResolveFromType(resolved)
}

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *TypeTree) ResolveFromType(toplevel *ResolvedType) bool {
	tree.reset()
	if tree.Level == 0 {
		return true
	}

	tree.Resolved = toplevel

	// Edge cases for builtin types.
	if tree.Resolved.Builtin != nil {
		switch {
		case toplevel.IsExternGLib("InitiallyUnowned"):
			return tree.resolveParents(externGLibType("*Object", gir.Type{}, "GObject*"))
		case toplevel.IsExternGLib("Object"):
			return true
		}

		return true
	}

	switch {
	case tree.Resolved.Extern.Result.Class != nil:
		// Resolving the parent type is crucial to make the class working, so if
		// this fails, halt and bail.
		if !tree.resolveName(tree.Resolved.Extern.Result.Class.Parent) {
			tryLogln(tree.res, LogUnknown,
				"can't resolve parent", tree.Resolved.Extern.Result.Class.Parent,
				"for class", tree.Resolved.Extern.Result.Class.Name,
			)
			return false
		}

		for _, impl := range tree.Resolved.Extern.Result.Class.Implements {
			if !tree.resolveName(impl.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve impl", impl.Name,
					"for class", tree.Resolved.Extern.Result.Class.Name,
				)
			}
		}

	case tree.Resolved.Extern.Result.Interface != nil:
		// All interfaces are derived from GObjects, so we override the list if
		// it's empty.
		if len(tree.Resolved.Extern.Result.Interface.Prerequisites) == 0 {
			return tree.resolveParents(
				externGLibType("*Object", gir.Type{}, "GObject*"),
			)
		}

		for _, prereq := range tree.Resolved.Extern.Result.Interface.Prerequisites {
			// Like class parents, interface prerequisites are important.
			if !tree.resolveName(prereq.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve prerequisite", prereq.Name,
					"for interface", tree.Resolved.Extern.Result.Interface.Name,
				)
				return false
			}
		}
	}

	return true
}

func (tree *TypeTree) parentLevel() int {
	if tree.Level <= 0 {
		return tree.Level
	}
	return tree.Level - 1
}

// resolveName resolves and adds the resolved type into the TypeTree.
func (tree *TypeTree) resolveName(name string) bool {
	parent := TypeTree{
		res:   tree.res,
		Level: tree.parentLevel(),
	}

	if !parent.Resolve(name) {
		return false
	}

	tree.Requires = append(tree.Requires, parent)
	return true
}

// resolveParents manually adds the given parents and resolve them to be added
// into the TypeTree.
func (tree *TypeTree) resolveParents(parents ...*ResolvedType) bool {
	for _, parent := range parents {
		parentTree := TypeTree{
			res:   tree.res,
			Level: tree.parentLevel(),
		}

		if !parentTree.ResolveFromType(parent) {
			// This shouldn't happen, unless the parent type made above is
			// invalid.
			return false
		}

		tree.Requires = append(tree.Requires, parentTree)
	}

	return true
}

// PublicChildren returns the list of the toplevel type's children as Go
// exported type names. The namespaces are appropriately prepended if needed.
func (tree *TypeTree) PublicChildren() []string {
	names := make([]string, len(tree.Requires))

	for i, req := range tree.Requires {
		namespace := req.Resolved.NeedsNamespace(tree.res.Namespace())
		names[i] = req.Resolved.PublicType(namespace)
	}

	return names
}

// Wrap creates a wrapper that uses public fields to create code that wraps the
// type tree to the top-level type. The fields are assumed to be public
// (exported) types. Types are assumed to all have valid wrap functions, so no
// nested wraps will actually be done.
//
// Wrapper functions for all types are assumed to follow this format:
//
//    func WrapTypeName(obj *externglib.Object) TypeName
//
func (tree *TypeTree) Wrap(obj string) string {
	p := pen.NewPiece()
	p.Write(tree.Resolved.PublicType(false)).Char('{')
	p.EmptyLine()

	for _, typ := range tree.Requires {
		if typ.Resolved.Builtin != nil {
			// If these cases hit, then the type is an Objector (as deefined by
			// gextras.Objector), so obj satisfies it.
			for _, glibType := range []string{"InitiallyUnowned", "Object"} {
				if typ.Resolved.IsExternGLib(glibType) {
					p.Line("Objector: obj,")
					goto glibContinue
				}
			}

			tryLogln(tree.res, LogUnknown, "builtin wrapping:", spew.Sdump(typ.Resolved))

		glibContinue:
			continue
		}

		// Extern types are generated by us, so the wrapper guarantee is
		// provided.
		namespace := typ.Resolved.NeedsNamespace(tree.res.Namespace())

		p.Linef(
			"%s: %s(%s),",
			typ.Resolved.PublicType(namespace),
			typ.Resolved.WrapName(namespace),
			obj,
		)
	}

	p.Char('}')
	return p.String()
}
