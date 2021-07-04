// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_progress_bar_get_type()), F: marshalProgressBar},
	})
}

// ProgressBar: the ProgressBar is typically used to display the progress of a
// long running operation. It provides a visual clue that processing is
// underway. The GtkProgressBar can be used in two different modes: percentage
// mode and activity mode.
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the GtkProgressBar in percentage mode and the user sees a growing bar
// indicating the percentage of the work that has been completed. In this mode,
// the application is required to call gtk_progress_bar_set_fraction()
// periodically to update the progress bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the ProgressBar in activity mode, which shows activity by a block
// moving back and forth within the progress area. In this mode, the application
// is required to call gtk_progress_bar_pulse() periodically to update the
// progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// ProgressBar. Functions are provided to control the orientation of the bar,
// optional text can be displayed along with the bar, and the step size used in
// activity mode can be set.
//
// CSS nodes
//
//    progressbar[.osd]
//    ├── [text]
//    ╰── trough[.empty][.full]
//        ╰── progress[.pulse]
//
// GtkProgressBar has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
type ProgressBar interface {
	Widget
	Orientable

	Ellipsize() pango.EllipsizeMode

	Fraction() float64

	Inverted() bool

	PulseStep() float64

	ShowText() bool

	Text() string

	PulseProgressBar()

	SetEllipsizeProgressBar(mode pango.EllipsizeMode)

	SetFractionProgressBar(fraction float64)

	SetInvertedProgressBar(inverted bool)

	SetPulseStepProgressBar(fraction float64)

	SetShowTextProgressBar(showText bool)

	SetTextProgressBar(text string)
}

// progressBar implements the ProgressBar class.
type progressBar struct {
	Widget
}

// WrapProgressBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapProgressBar(obj *externglib.Object) ProgressBar {
	return progressBar{
		Widget: WrapWidget(obj),
	}
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProgressBar(obj), nil
}

func NewProgressBar() ProgressBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_progress_bar_new()

	var _progressBar ProgressBar // out

	_progressBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ProgressBar)

	return _progressBar
}

func (p progressBar) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkProgressBar    // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

func (p progressBar) Fraction() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (p progressBar) Inverted() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p progressBar) PulseStep() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gdouble         // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (p progressBar) ShowText() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_show_text(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p progressBar) Text() string {
	var _arg0 *C.GtkProgressBar // out
	var _cret *C.gchar          // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p progressBar) PulseProgressBar() {
	var _arg0 *C.GtkProgressBar // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	C.gtk_progress_bar_pulse(_arg0)
}

func (p progressBar) SetEllipsizeProgressBar(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkProgressBar    // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_progress_bar_set_ellipsize(_arg0, _arg1)
}

func (p progressBar) SetFractionProgressBar(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_fraction(_arg0, _arg1)
}

func (p progressBar) SetInvertedProgressBar(inverted bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_inverted(_arg0, _arg1)
}

func (p progressBar) SetPulseStepProgressBar(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gdouble         // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_pulse_step(_arg0, _arg1)
}

func (p progressBar) SetShowTextProgressBar(showText bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if showText {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_show_text(_arg0, _arg1)
}

func (p progressBar) SetTextProgressBar(text string) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_progress_bar_set_text(_arg0, _arg1)
}

func (b progressBar) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b progressBar) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b progressBar) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b progressBar) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b progressBar) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b progressBar) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b progressBar) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b progressBar) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b progressBar) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b progressBar) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o progressBar) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o progressBar) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}