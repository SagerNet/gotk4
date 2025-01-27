// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TextBufferDeserializeFunc: function that is called to deserialize rich text
// that has been serialized with gtk_text_buffer_serialize(), and insert it at
// iter.
type TextBufferDeserializeFunc func(registerBuffer *TextBuffer, contentBuffer *TextBuffer, iter *TextIter, data []byte, createTags bool) (ok bool)

//export _gotk4_gtk3_TextBufferDeserializeFunc
func _gotk4_gtk3_TextBufferDeserializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.guint8, arg4 C.gsize, arg5 C.gboolean, arg6 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	var registerBuffer *TextBuffer // out
	var contentBuffer *TextBuffer  // out
	var iter *TextIter             // out
	var data []byte                // out
	var createTags bool            // out

	registerBuffer = WrapTextBuffer(externglib.Take(unsafe.Pointer(arg0)))
	contentBuffer = WrapTextBuffer(externglib.Take(unsafe.Pointer(arg1)))
	iter = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(iter)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_iter_free((*C.GtkTextIter)(intern.C))
		},
	)
	defer C.free(unsafe.Pointer(arg3))
	data = make([]byte, arg4)
	copy(data, unsafe.Slice((*byte)(unsafe.Pointer(arg3)), arg4))
	if arg5 != 0 {
		createTags = true
	}

	fn := v.(TextBufferDeserializeFunc)
	ok := fn(registerBuffer, contentBuffer, iter, data, createTags)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// TextBufferSerializeFunc: function that is called to serialize the content of
// a text buffer. It must return the serialized form of the content.
type TextBufferSerializeFunc func(registerBuffer *TextBuffer, contentBuffer *TextBuffer, start *TextIter, end *TextIter) (length uint, guint8 *byte)

//export _gotk4_gtk3_TextBufferSerializeFunc
func _gotk4_gtk3_TextBufferSerializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.GtkTextIter, arg4 *C.gsize, arg5 C.gpointer) (cret *C.guint8) {
	v := gbox.Get(uintptr(arg5))
	if v == nil {
		panic(`callback not found`)
	}

	var registerBuffer *TextBuffer // out
	var contentBuffer *TextBuffer  // out
	var start *TextIter            // out
	var end *TextIter              // out

	registerBuffer = WrapTextBuffer(externglib.Take(unsafe.Pointer(arg0)))
	contentBuffer = WrapTextBuffer(externglib.Take(unsafe.Pointer(arg1)))
	start = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(start)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_iter_free((*C.GtkTextIter)(intern.C))
		},
	)
	end = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(end)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_iter_free((*C.GtkTextIter)(intern.C))
		},
	)

	fn := v.(TextBufferSerializeFunc)
	length, guint8 := fn(registerBuffer, contentBuffer, start, end)

	*arg4 = C.gsize(length)
	if guint8 != nil {
		cret = (*C.guint8)(unsafe.Pointer(guint8))
	}

	return cret
}
