package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// C to Go type conversions.

// CGoConverter returns Go code that is the conversion from the given C value
// type to its respective Go value type. An empty string returned is invalid.
//
// The given argPrefix is used to get the nth parameter by concatenating the
// prefix with the index number. This is used for length parameters.
func (ng *NamespaceGenerator) CGoConverter(conv TypeConversionToGo) string {
	switch {
	case conv.Type.Array != nil:
		return ng.cgoArrayConverter(conv)
	case conv.Type.Type != nil:
		return ng.cgoTypeConverter(conv)
	}

	// Ignore VarArgs.
	return ""
}

func (ng *NamespaceGenerator) cgoArrayConverter(conv TypeConversionToGo) string {
	array := conv.Type.Array

	if array.Type == nil {
		ng.gen.logln(logSkip, "nested array", array.CType)
		return ""
	}

	innerResolved := ng.ResolveType(*array.AnyType.Type)
	if innerResolved == nil {
		return ""
	}
	innerType := ng.PublicType(innerResolved)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "${target}[i]" variables.
	innerTypeConv := conv
	innerTypeConv.TypeConversion = conv.inner("src", conv.Target+"[i]")

	innerConv := ng.CGoConverter(innerTypeConv)
	if innerConv == "" {
		return ""
	}

	b := pen.NewBlock()

	switch {
	case array.FixedSize > 0:
		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast a fixed-size array.
		if p := girPrimitiveGo(array.Type.Name); p != "" {
			return conv.callf("[%d]%s", array.FixedSize, p)
		}

		// TODO: nested array support
		b.Linef("cArray := ([%d]%s)(%s)", array.FixedSize, array.Type.CType, conv.Value)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := cArray[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		ng.addImport("unsafe")

		lengthArg := conv.ArgAt(*array.Length)

		// If we're owning the new data, then we will directly use the backing
		// array, but we can only do this if the underlying type is a primitive,
		// since those have equivalent Go representations. Any other types will
		// have to be copied or otherwise converted somehow.
		//
		// TODO: record conversion should handle ownership: if
		// transfer-ownership is none, then the native pointer should probably
		// not be freed.
		if conv.isTransferring() && innerResolved.IsPrimitive() {
			ng.addImport("runtime")
			ng.addImport("reflect")

			b.Linef(goSliceFromPtr(conv.Target, conv.Value, lengthArg))

			// Ensure that Go's GC doesn't touch the pointer within the duration
			// of the function.
			// See: https://golang.org/misc/cgo/gmp/gmp.go?s=3086:3757#L87
			b.Linef("runtime.SetFinalizer(&%s, func() {", conv.Value)
			b.Linef("  C.free(unsafe.Pointer(%s))", conv.Value)
			b.Linef("})")
			b.Linef("defer runtime.KeepAlive(%s)", conv.Value)
			return b.String()
		}

		b.Linef("%s = make([]%s, %s)", conv.Target, innerType, lengthArg)
		b.Linef("for i := 0; i < uintptr(%s); i++ {", lengthArg)
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Name == "GLib.Array": // treat as Go array
		ng.addImport("unsafe")

		b.Linef("var length uintptr")
		b.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&length))", conv.Value)
		b.Linef("%s = make([]%s, length)", conv.Target, innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		// TODO: nested array support
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.IsZeroTerminated():
		ng.addImport("unsafe")

		// Scan for the length.
		b.Linef("var length uint")
		b.Linef("for p := unsafe.Pointer(%s); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {", conv.Value)
		b.Linef("  length++")
		b.Linef("}")

		b.EmptyLine()

		// Preallocate the slice.
		b.Linef("%s = make([]%s, length)", conv.Target, innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + i))", innerCGoType, conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")

	default:
		ng.logln(logSkip, conv.ParentName+":", "weird array type to Go")
		return ""
	}

	return b.String()
}

// goSliceFromPtr crafts a typ slice from the given ptr as the backing array
// with the given len, then set it into target. typ should be innerType. A
// temporary variable named sliceHeader is made.
func goSliceFromPtr(target, ptr, len string) string {
	return pen.NewPiece().
		Linef("sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&%s))", target).
		Linef("sliceHeader.Data = uintptr(unsafe.Pointer(%s))", ptr).
		Linef("sliceHeader.Len = %s", len).
		Linef("sliceHeader.Cap = %s", len).
		String()
}

func (ng *NamespaceGenerator) cgoTypeConverter(conv TypeConversionToGo) string {
	typ := conv.Type.Type

	for _, unsupported := range unsupportedTypes {
		if unsupported == typ.Name {
			return ""
		}
	}

	// Resolve primitive types.
	if prim, ok := girToBuiltin[typ.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef("%s = C.GoString(%s)", conv.Target, conv.Value)

			// Only free this if C is transferring ownership to us.
			if conv.isTransferring() {
				ng.addImport("unsafe")
				p.Linef("C.free(unsafe.Pointer(%s))", conv.Value)
			}

			return p.String()

		case "bool":
			return fmt.Sprintf("%s = C.BOOL(%s) != 0", conv.Target, conv.Value)

		default:
			return directCall(conv.Value, conv.Target, prim)
		}
	}

	// Resolve special-case GLib types.
	switch ng.ensureNamespace(typ.Name) {
	case "gpointer":
		ng.addImport("github.com/diamondburned/gotk4/internal/box")

		castTail := conv.BoxCast
		if castTail != "" {
			castTail = fmt.Sprintf(".(%s)", castTail)
		}

		return fmt.Sprintf("%s = box.Get(uintptr(%s))%s", conv.Target, conv.Value, castTail)

	case "GObject.Object", "GObject.InitiallyUnowned":
		ng.addImport("unsafe")
		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		return cgoTakeObject(conv.TypeConversion, "gextras.Objector")

	case "GObject.Type", "GType":
		ng.addGLibImport()
		return conv.call("externglib.Type")

	case "GObject.Value":
		ng.addGLibImport()
		ng.addImport("unsafe")

		p := pen.NewPiece()
		p.Linef("%s = externglib.ValueFromNative(unsafe.Pointer(%s))", conv.Target, conv.Value)

		// Set this to be freed if we have the ownership now.
		if conv.isTransferring() {
			p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", conv.Target)
			p.Linef("  C.g_value_unset((*C.GValue)(v.GValue))")
			p.Linef("})")
		}

		return p.String()
	}

	// Pretend that ignored types don't exist.
	if ng.mustIgnore(typ.Name, typ.CType) {
		return ""
	}

	result := ng.gen.Repos.FindType(ng.current, typ.Name)
	if result == nil {
		// Probably already warned.
		return ""
	}

	resolved := typeFromResult(ng.gen, *typ, result)

	// goName contains the pointer.
	goName := ng.PublicType(resolved)

	// Resolve alias.
	if result.Alias != nil {
		rootType := ng.ResolveType(result.Alias.Type)
		if rootType == nil {
			ng.logln(logUnknown, "from alias", result.Alias.Name, "has", result.Alias.Type.Name)
			return ""
		}

		rootConv := conv
		rootConv.Type = gir.AnyType{Type: &result.Alias.Type}
		rootConv.Target = "tmp"

		b := pen.NewBlock()
		b.Linef("var tmp %s", ng.PublicType(rootType))
		b.Linef(ng.CGoConverter(rootConv))
		b.Linef("%s = %s(tmp)", conv.Target, goName)
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return conv.call(goName)
	}

	if result.Class != nil || result.Record != nil || result.Interface != nil {
		switch {
		case result.Class != nil, result.Interface != nil:
			ng.addImport("unsafe")
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return cgoTakeObject(conv.TypeConversion, goName)

		case result.Record != nil:
			// We should only use the concrete wrapper for the record, since the
			// returned type is concretely known here.
			wrapName := resolved.WrapName(resolved.NeedsNamespace(ng.current))

			b := pen.NewBlock()
			b.Linef("%s = %s(unsafe.Pointer(%s))", conv.Target, wrapName, conv.Value)

			// If ownership is being transferred to us on the Go side, then we
			// should free.
			if conv.isTransferring() {
				ng.addImport("runtime")

				arg := conv.Target
				typ := goName

				if resolved.Ptr < 1 {
					arg = "&" + arg
					typ = "*" + typ
				}

				b.Linef("runtime.SetFinalizer(%s, func(v %s) {", arg, typ)
				b.Linef("  C.free(unsafe.Pointer(v.Native()))")
				b.Linef("})")
			}

			return b.String()
		}
	}

	// Callbacks returned don't seem to have an output closure, so we can't get
	// our closure here.
	if result.Callback != nil {
		return ""
	}

	// TODO: function
	// TODO: union
	// TODO: callback

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.

	return ""
}

// cgoTakeObject generates a glib.Take or glib.AssumeOwnership.
func cgoTakeObject(conv TypeConversion, ifaceType string) string {
	var gobjectFunction string
	if conv.isTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	return fmt.Sprintf(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s.Native()))).(%s)",
		conv.Target, gobjectFunction, conv.Value, ifaceType,
	)
}
