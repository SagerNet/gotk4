// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_emoji_chooser_get_type()), F: marshalEmojiChooserer},
	})
}

// EmojiChooser: GtkEmojiChooser is used by text widgets such as GtkEntry or
// GtkTextView to let users insert Emoji characters.
//
// !An example GtkEmojiChooser (emojichooser.png)
//
// GtkEmojiChooser emits the gtk.EmojiChooser::emoji-picked signal when an Emoji
// is selected.
//
// CSS nodes
//
//    popover
//    ├── box.emoji-searchbar
//    │   ╰── entry.search
//    ╰── box.emoji-toolbar
//        ├── button.image-button.emoji-section
//        ├── ...
//        ╰── button.image-button.emoji-section
//
//
// Every GtkEmojiChooser consists of a main node called popover. The contents of
// the popover are largely implementation defined and supposed to inherit
// general styles. The top searchbar used to search emoji and gets the
// .emoji-searchbar style class itself. The bottom toolbar used to switch
// between different emoji categories consists of buttons with the
// .emoji-section style class and gets the .emoji-toolbar style class itself.
type EmojiChooser struct {
	Popover
}

func WrapEmojiChooser(obj *externglib.Object) *EmojiChooser {
	return &EmojiChooser{
		Popover: Popover{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
			NativeSurface: NativeSurface{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
					Object: obj,
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalEmojiChooserer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEmojiChooser(obj), nil
}

// NewEmojiChooser creates a new GtkEmojiChooser.
func NewEmojiChooser() *EmojiChooser {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_emoji_chooser_new()

	var _emojiChooser *EmojiChooser // out

	_emojiChooser = WrapEmojiChooser(externglib.Take(unsafe.Pointer(_cret)))

	return _emojiChooser
}

func (*EmojiChooser) privateEmojiChooser() {}
