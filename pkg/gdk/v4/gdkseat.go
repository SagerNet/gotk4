// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeater},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities int

const (
	// SeatCapabilitiesNone: no input capabilities
	SeatCapabilitiesNone SeatCapabilities = 0b0
	// SeatCapabilitiesPointer: the seat has a pointer (e.g. mouse)
	SeatCapabilitiesPointer SeatCapabilities = 0b1
	// SeatCapabilitiesTouch: the seat has touchscreen(s) attached
	SeatCapabilitiesTouch SeatCapabilities = 0b10
	// SeatCapabilitiesTabletStylus: the seat has drawing tablet(s) attached
	SeatCapabilitiesTabletStylus SeatCapabilities = 0b100
	// SeatCapabilitiesKeyboard: the seat has keyboard(s) attached
	SeatCapabilitiesKeyboard SeatCapabilities = 0b1000
	// SeatCapabilitiesTabletPad: the seat has drawing tablet pad(s) attached
	SeatCapabilitiesTabletPad SeatCapabilities = 0b10000
	// SeatCapabilitiesAllPointing: the union of all pointing capabilities
	SeatCapabilitiesAllPointing SeatCapabilities = 0b111
	// SeatCapabilitiesAll: the union of all capabilities
	SeatCapabilitiesAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Seater describes Seat's methods.
type Seater interface {
	gextras.Objector

	Capabilities() SeatCapabilities
	Display() *Display
	Keyboard() *Device
	Pointer() *Device
}

// Seat: the `GdkSeat` object represents a collection of input devices that
// belong to a user.
type Seat struct {
	*externglib.Object
}

var _ Seater = (*Seat)(nil)

func wrapSeater(obj *externglib.Object) Seater {
	return &Seat{
		Object: obj,
	}
}

func marshalSeater(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeater(obj), nil
}

// Capabilities returns the capabilities this `GdkSeat` currently has.
func (seat *Seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = (SeatCapabilities)(_cret)

	return _seatCapabilities
}

// Display returns the `GdkDisplay` this seat belongs to.
func (seat *Seat) Display() *Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_display(_arg0)

	var _display *Display // out

	_display = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Display)

	return _display
}

// Keyboard returns the device that routes keyboard events.
func (seat *Seat) Keyboard() *Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	var _device *Device // out

	_device = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Device)

	return _device
}

// Pointer returns the device that routes pointer events.
func (seat *Seat) Pointer() *Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(seat.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	var _device *Device // out

	_device = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Device)

	return _device
}
