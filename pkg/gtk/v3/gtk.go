// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_movement_step_get_type()), F: marshalMovementStep},
		{T: externglib.Type(C.gtk_notebook_tab_get_type()), F: marshalNotebookTab},
		{T: externglib.Type(C.gtk_resize_mode_get_type()), F: marshalResizeMode},
		{T: externglib.Type(C.gtk_scroll_step_get_type()), F: marshalScrollStep},
		{T: externglib.Type(C.gtk_debug_flag_get_type()), F: marshalDebugFlag},
		{T: externglib.Type(C.gtk_entry_icon_accessible_get_type()), F: marshalEntryIconAccessibler},
	})
}

type MovementStep int

const (
	// LogicalPositions: move forward or back by graphemes
	MovementStepLogicalPositions MovementStep = iota
	// VisualPositions: move left or right by graphemes
	MovementStepVisualPositions
	// Words: move forward or back by words
	MovementStepWords
	// DisplayLines: move up or down lines (wrapped lines)
	MovementStepDisplayLines
	// DisplayLineEnds: move to either end of a line
	MovementStepDisplayLineEnds
	// Paragraphs: move up or down paragraphs (newline-ended lines)
	MovementStepParagraphs
	// ParagraphEnds: move to either end of a paragraph
	MovementStepParagraphEnds
	// Pages: move by pages
	MovementStepPages
	// BufferEnds: move to ends of the buffer
	MovementStepBufferEnds
	// HorizontalPages: move horizontally by pages
	MovementStepHorizontalPages
)

func marshalMovementStep(p uintptr) (interface{}, error) {
	return MovementStep(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type NotebookTab int

const (
	NotebookTabFirst NotebookTab = iota
	NotebookTabLast
)

func marshalNotebookTab(p uintptr) (interface{}, error) {
	return NotebookTab(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type ResizeMode int

const (
	// Parent pass resize request to the parent
	ResizeModeParent ResizeMode = iota
	// Queue: queue resizes on this widget
	ResizeModeQueue
	// Immediate: resize immediately. Deprecated.
	ResizeModeImmediate
)

func marshalResizeMode(p uintptr) (interface{}, error) {
	return ResizeMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type ScrollStep int

const (
	// Steps: scroll in steps.
	ScrollStepSteps ScrollStep = iota
	// Pages: scroll by pages.
	ScrollStepPages
	// Ends: scroll to ends.
	ScrollStepEnds
	// HorizontalSteps: scroll in horizontal steps.
	ScrollStepHorizontalSteps
	// HorizontalPages: scroll by horizontal pages.
	ScrollStepHorizontalPages
	// HorizontalEnds: scroll to the horizontal ends.
	ScrollStepHorizontalEnds
)

func marshalScrollStep(p uintptr) (interface{}, error) {
	return ScrollStep(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

type DebugFlag int

const (
	DebugFlagMisc         DebugFlag = 0b1
	DebugFlagPlugsocket   DebugFlag = 0b10
	DebugFlagText         DebugFlag = 0b100
	DebugFlagTree         DebugFlag = 0b1000
	DebugFlagUpdates      DebugFlag = 0b10000
	DebugFlagKeybindings  DebugFlag = 0b100000
	DebugFlagMultihead    DebugFlag = 0b1000000
	DebugFlagModules      DebugFlag = 0b10000000
	DebugFlagGeometry     DebugFlag = 0b100000000
	DebugFlagIcontheme    DebugFlag = 0b1000000000
	DebugFlagPrinting     DebugFlag = 0b10000000000
	DebugFlagBuilder      DebugFlag = 0b100000000000
	DebugFlagSizeRequest  DebugFlag = 0b1000000000000
	DebugFlagNoCSSCache   DebugFlag = 0b10000000000000
	DebugFlagBaselines    DebugFlag = 0b100000000000000
	DebugFlagPixelCache   DebugFlag = 0b1000000000000000
	DebugFlagNoPixelCache DebugFlag = 0b10000000000000000
	DebugFlagInteractive  DebugFlag = 0b100000000000000000
	DebugFlagTouchscreen  DebugFlag = 0b1000000000000000000
	DebugFlagActions      DebugFlag = 0b10000000000000000000
	DebugFlagResize       DebugFlag = 0b100000000000000000000
	DebugFlagLayout       DebugFlag = 0b1000000000000000000000
)

func marshalDebugFlag(p uintptr) (interface{}, error) {
	return DebugFlag(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EntryIconAccessibler describes EntryIconAccessible's methods.
type EntryIconAccessibler interface {
	gextras.Objector

	privateEntryIconAccessible()
}

type EntryIconAccessible struct {
	*externglib.Object
	atk.Object
	atk.Action
}

var _ EntryIconAccessibler = (*EntryIconAccessible)(nil)

func wrapEntryIconAccessibler(obj *externglib.Object) EntryIconAccessibler {
	return &EntryIconAccessible{
		Object: obj,
		Object: atk.Object{
			Object: obj,
		},
		Action: atk.Action{
			Object: obj,
		},
	}
}

func marshalEntryIconAccessibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryIconAccessibler(obj), nil
}

func (*EntryIconAccessible) privateEntryIconAccessible() {}
