// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeat},
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
	// SeatCapabilitiesAllPointing: the union of all pointing capabilities
	SeatCapabilitiesAllPointing SeatCapabilities = 0b111
	// SeatCapabilitiesAll: the union of all capabilities
	SeatCapabilitiesAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Seat: the Seat object represents a collection of input devices that belong to
// a user.
type Seat interface {
	gextras.Objector

	// Capabilities:
	Capabilities() SeatCapabilities
	// Display:
	Display() Display
	// Keyboard:
	Keyboard() Device
	// Pointer:
	Pointer() Device
	// UngrabSeat:
	UngrabSeat()
}

// seat implements the Seat class.
type seat struct {
	gextras.Objector
}

// WrapSeat wraps a GObject to the right type. It is
// primarily used internally.
func WrapSeat(obj *externglib.Object) Seat {
	return seat{
		Objector: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeat(obj), nil
}

func (s seat) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = SeatCapabilities(_cret)

	return _seatCapabilities
}

func (s seat) Display() Display {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (s seat) Keyboard() Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (s seat) Pointer() Device {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	var _device Device // out

	_device = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Device)

	return _device
}

func (s seat) UngrabSeat() {
	var _arg0 *C.GdkSeat // out

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	C.gdk_seat_ungrab(_arg0)
}
