// Code generated by girgen. DO NOT EDIT.

package glib

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_io_condition_get_type()), F: marshalIOCondition},
	})
}

// IOCondition: a bitwise combination representing a condition to watch for on
// an event source.
type IOCondition int

const (
	// IOConditionIn: there is data to read.
	IOConditionIn IOCondition = 0b1
	// IOConditionOut: data can be written (without blocking).
	IOConditionOut IOCondition = 0b100
	// IOConditionPri: there is urgent data to read.
	IOConditionPri IOCondition = 0b10
	// IOConditionErr: error condition.
	IOConditionErr IOCondition = 0b1000
	// IOConditionHup: hung up (the connection has been broken, usually for
	// pipes and sockets).
	IOConditionHup IOCondition = 0b10000
	// IOConditionNval: invalid request. The file descriptor is not open.
	IOConditionNval IOCondition = 0b100000
)

func marshalIOCondition(p uintptr) (interface{}, error) {
	return IOCondition(C.g_value_get_bitfield((*C.GValue)(unsafe.Pointer(p)))), nil
}
