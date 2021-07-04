// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_window_get_type()), F: marshalShortcutsWindow},
	})
}

// ShortcutsWindow: a GtkShortcutsWindow shows brief information about the
// keyboard shortcuts and gestures of an application. The shortcuts can be
// grouped, and you can have multiple sections in this window, corresponding to
// the major modes of your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with GtkBuilder, by
// populating a ShortcutsWindow with one or more ShortcutsSection objects, which
// contain ShortcutsGroups that in turn contain objects of class
// ShortcutsShortcut.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a ShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a ShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow interface {
	Window
}

// shortcutsWindow implements the ShortcutsWindow class.
type shortcutsWindow struct {
	Window
}

// WrapShortcutsWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutsWindow(obj *externglib.Object) ShortcutsWindow {
	return shortcutsWindow{
		Window: WrapWindow(obj),
	}
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutsWindow(obj), nil
}

func (b shortcutsWindow) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b shortcutsWindow) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b shortcutsWindow) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b shortcutsWindow) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b shortcutsWindow) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b shortcutsWindow) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b shortcutsWindow) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
