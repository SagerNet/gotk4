// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// BookmarkFileError: error codes returned by bookmark file parsing.
type BookmarkFileError int

const (
	// BookmarkFileErrorInvalidURI: URI was ill-formed
	BookmarkFileErrorInvalidURI BookmarkFileError = iota
	// BookmarkFileErrorInvalidValue: requested field was not found
	BookmarkFileErrorInvalidValue
	// BookmarkFileErrorAppNotRegistered: requested application did not register
	// a bookmark
	BookmarkFileErrorAppNotRegistered
	// BookmarkFileErrorURINotFound: requested URI was not found
	BookmarkFileErrorURINotFound
	// BookmarkFileErrorRead: document was ill formed
	BookmarkFileErrorRead
	// BookmarkFileErrorUnknownEncoding: text being parsed was in an unknown
	// encoding
	BookmarkFileErrorUnknownEncoding
	// BookmarkFileErrorWrite: error occurred while writing
	BookmarkFileErrorWrite
	// BookmarkFileErrorFileNotFound: requested file was not found
	BookmarkFileErrorFileNotFound
)

// String returns the name in string for BookmarkFileError.
func (b BookmarkFileError) String() string {
	switch b {
	case BookmarkFileErrorInvalidURI:
		return "InvalidURI"
	case BookmarkFileErrorInvalidValue:
		return "InvalidValue"
	case BookmarkFileErrorAppNotRegistered:
		return "AppNotRegistered"
	case BookmarkFileErrorURINotFound:
		return "URINotFound"
	case BookmarkFileErrorRead:
		return "Read"
	case BookmarkFileErrorUnknownEncoding:
		return "UnknownEncoding"
	case BookmarkFileErrorWrite:
		return "Write"
	case BookmarkFileErrorFileNotFound:
		return "FileNotFound"
	default:
		return fmt.Sprintf("BookmarkFileError(%d)", b)
	}
}
