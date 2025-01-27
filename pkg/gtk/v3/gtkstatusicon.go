// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_status_icon_get_type()), F: marshalStatusIconner},
	})
}

// StatusIconOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StatusIconOverrider interface {
	Activate()
	ButtonPressEvent(event *gdk.EventButton) bool
	ButtonReleaseEvent(event *gdk.EventButton) bool
	PopupMenu(button uint, activateTime uint32)
	QueryTooltip(x int, y int, keyboardMode bool, tooltip *Tooltip) bool
	ScrollEvent(event *gdk.EventScroll) bool
	SizeChanged(size int) bool
}

// StatusIcon: “system tray” or notification area is normally used for transient
// icons that indicate some special state. For example, a system tray icon might
// appear to tell the user that they have new mail, or have an incoming instant
// message, or something along those lines. The basic idea is that creating an
// icon in the notification area is less annoying than popping up a dialog.
//
// A StatusIcon object can be used to display an icon in a “system tray”. The
// icon can have a tooltip, and the user can interact with it by activating it
// or popping up a context menu.
//
// It is very important to notice that status icons depend on the existence of a
// notification area being available to the user; you should not use status
// icons as the only way to convey critical information regarding your
// application, as the notification area may not exist on the user's
// environment, or may have been removed. You should always check that a status
// icon has been embedded into a notification area by using
// gtk_status_icon_is_embedded(), and gracefully recover if the function returns
// FALSE.
//
// On X11, the implementation follows the FreeDesktop System Tray Specification
// (http://www.freedesktop.org/wiki/Specifications/systemtray-spec).
// Implementations of the “tray” side of this specification can be found e.g. in
// the GNOME 2 and KDE panel applications.
//
// Note that a GtkStatusIcon is not a widget, but just a #GObject. Making it a
// widget would be impractical, since the system tray on Windows doesn’t allow
// to embed arbitrary widgets.
//
// GtkStatusIcon has been deprecated in 3.14. You should consider using
// notifications or more modern platform-specific APIs instead. GLib provides
// the #GNotification API which works well with Application on multiple
// platforms and environments, and should be the preferred mechanism to notify
// the users of transient status updates. See this HowDoI
// (https://wiki.gnome.org/HowDoI/GNotification) for code examples.
type StatusIcon struct {
	*externglib.Object
}

func WrapStatusIcon(obj *externglib.Object) *StatusIcon {
	return &StatusIcon{
		Object: obj,
	}
}

func marshalStatusIconner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusIcon(obj), nil
}

// NewStatusIcon creates an empty status icon object.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIcon() *StatusIcon {
	var _cret *C.GtkStatusIcon // in

	_cret = C.gtk_status_icon_new()

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// NewStatusIconFromFile creates a status icon displaying the file filename.
//
// The image will be scaled down to fit in the available space in the
// notification area, if necessary.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIconFromFile(filename string) *StatusIcon {
	var _arg1 *C.gchar         // out
	var _cret *C.GtkStatusIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_status_icon_new_from_file(_arg1)
	runtime.KeepAlive(filename)

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// NewStatusIconFromGIcon creates a status icon displaying a #GIcon. If the icon
// is a themed icon, it will be updated when the theme changes.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIconFromGIcon(icon gio.Iconner) *StatusIcon {
	var _arg1 *C.GIcon         // out
	var _cret *C.GtkStatusIcon // in

	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	_cret = C.gtk_status_icon_new_from_gicon(_arg1)
	runtime.KeepAlive(icon)

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// NewStatusIconFromIconName creates a status icon displaying an icon from the
// current icon theme. If the current icon theme is changed, the icon will be
// updated appropriately.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIconFromIconName(iconName string) *StatusIcon {
	var _arg1 *C.gchar         // out
	var _cret *C.GtkStatusIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_status_icon_new_from_icon_name(_arg1)
	runtime.KeepAlive(iconName)

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// NewStatusIconFromPixbuf creates a status icon displaying pixbuf.
//
// The image will be scaled down to fit in the available space in the
// notification area, if necessary.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIconFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *StatusIcon {
	var _arg1 *C.GdkPixbuf     // out
	var _cret *C.GtkStatusIcon // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_status_icon_new_from_pixbuf(_arg1)
	runtime.KeepAlive(pixbuf)

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// NewStatusIconFromStock creates a status icon displaying a stock icon. Sample
// stock icon names are K_STOCK_OPEN, K_STOCK_QUIT. You can register your own
// stock icon names, see gtk_icon_factory_add_default() and
// gtk_icon_factory_add().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications.
func NewStatusIconFromStock(stockId string) *StatusIcon {
	var _arg1 *C.gchar         // out
	var _cret *C.GtkStatusIcon // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_status_icon_new_from_stock(_arg1)
	runtime.KeepAlive(stockId)

	var _statusIcon *StatusIcon // out

	_statusIcon = WrapStatusIcon(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _statusIcon
}

// Geometry obtains information about the location of the status icon on screen.
// This information can be used to e.g. position popups like notification
// bubbles.
//
// See gtk_status_icon_position_menu() for a more convenient alternative for
// positioning menus.
//
// Note that some platforms do not allow GTK+ to provide this information, and
// even on platforms that do allow it, the information is not reliable unless
// the status icon is embedded in a notification area, see
// gtk_status_icon_is_embedded().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as the
// platform is responsible for the presentation of notifications.
func (statusIcon *StatusIcon) Geometry() (*gdk.Screen, gdk.Rectangle, Orientation, bool) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GdkScreen     // in
	var _arg2 C.GdkRectangle   // in
	var _arg3 C.GtkOrientation // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_geometry(_arg0, &_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(statusIcon)

	var _screen *gdk.Screen      // out
	var _area gdk.Rectangle      // out
	var _orientation Orientation // out
	var _ok bool                 // out

	if _arg1 != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_arg1))
			_screen = &gdk.Screen{
				Object: obj,
			}
		}
	}
	_area = *(*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_orientation = Orientation(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _screen, _area, _orientation, _ok
}

// GIcon retrieves the #GIcon being displayed by the StatusIcon. The storage
// type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_status_icon_get_storage_type()). The caller of this function does not own
// a reference to the returned #GIcon.
//
// If this function fails, icon is left unchanged;
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) GIcon() gio.Iconner {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.GIcon         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_gicon(_arg0)
	runtime.KeepAlive(statusIcon)

	var _icon gio.Iconner // out

	if _cret != nil {
		_icon = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Iconner)
	}

	return _icon
}

// HasTooltip returns the current value of the has-tooltip property. See
// StatusIcon:has-tooltip for more information.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) HasTooltip() bool {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_has_tooltip(_arg0)
	runtime.KeepAlive(statusIcon)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the icon being displayed by the StatusIcon. The
// storage type of the status icon must be GTK_IMAGE_EMPTY or
// GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()). The returned
// string is owned by the StatusIcon and should not be freed or modified.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) IconName() string {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_icon_name(_arg0)
	runtime.KeepAlive(statusIcon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Pixbuf gets the Pixbuf being displayed by the StatusIcon. The storage type of
// the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_status_icon_get_storage_type()). The caller of this function does not own
// a reference to the returned pixbuf.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) Pixbuf() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.GdkPixbuf     // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_pixbuf(_arg0)
	runtime.KeepAlive(statusIcon)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// Screen returns the Screen associated with status_icon.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as
// notifications are managed by the platform.
func (statusIcon *StatusIcon) Screen() *gdk.Screen {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.GdkScreen     // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_screen(_arg0)
	runtime.KeepAlive(statusIcon)

	var _screen *gdk.Screen // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_screen = &gdk.Screen{
			Object: obj,
		}
	}

	return _screen
}

// Size gets the size in pixels that is available for the image. Stock icons and
// named icons adapt their size automatically if the size of the notification
// area changes. For other storage types, the size-changed signal can be used to
// react to size changes.
//
// Note that the returned size is only meaningful while the status icon is
// embedded (see gtk_status_icon_is_embedded()).
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as the
// representation of a notification is left to the platform.
func (statusIcon *StatusIcon) Size() int {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.gint           // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_size(_arg0)
	runtime.KeepAlive(statusIcon)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Stock gets the id of the stock icon being displayed by the StatusIcon. The
// storage type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK
// (see gtk_status_icon_get_storage_type()). The returned string is owned by the
// StatusIcon and should not be freed or modified.
//
// Deprecated: Use gtk_status_icon_get_icon_name() instead.
func (statusIcon *StatusIcon) Stock() string {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_stock(_arg0)
	runtime.KeepAlive(statusIcon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// StorageType gets the type of representation being used by the StatusIcon to
// store image data. If the StatusIcon has no image data, the return value will
// be GTK_IMAGE_EMPTY.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, and
// #GNotification only supports #GIcon instances.
func (statusIcon *StatusIcon) StorageType() ImageType {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.GtkImageType   // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_storage_type(_arg0)
	runtime.KeepAlive(statusIcon)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

// Title gets the title of this tray icon. See gtk_status_icon_set_title().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) Title() string {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_title(_arg0)
	runtime.KeepAlive(statusIcon)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TooltipMarkup gets the contents of the tooltip for status_icon.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) TooltipMarkup() string {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_tooltip_markup(_arg0)
	runtime.KeepAlive(statusIcon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// TooltipText gets the contents of the tooltip for status_icon.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) TooltipText() string {
	var _arg0 *C.GtkStatusIcon // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_tooltip_text(_arg0)
	runtime.KeepAlive(statusIcon)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Visible returns whether the status icon is visible or not. Note that being
// visible does not guarantee that the user can actually see the icon, see also
// gtk_status_icon_is_embedded().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) Visible() bool {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_visible(_arg0)
	runtime.KeepAlive(statusIcon)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// X11WindowID: this function is only useful on the X11/freedesktop.org
// platform.
//
// It returns a window ID for the widget in the underlying status icon
// implementation. This is useful for the Galago notification service, which can
// send a window ID in the protocol in order for the server to position
// notification windows pointing to a status icon reliably.
//
// This function is not intended for other use cases which are more likely to be
// met by one of the non-X11 specific methods, such as
// gtk_status_icon_position_menu().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) X11WindowID() uint32 {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.guint32        // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_get_x11_window_id(_arg0)
	runtime.KeepAlive(statusIcon)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// IsEmbedded returns whether the status icon is embedded in a notification
// area.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) IsEmbedded() bool {
	var _arg0 *C.GtkStatusIcon // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))

	_cret = C.gtk_status_icon_is_embedded(_arg0)
	runtime.KeepAlive(statusIcon)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFromFile makes status_icon display the file filename. See
// gtk_status_icon_new_from_file() for details.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; you can use g_notification_set_icon() to associate a #GIcon
// with a notification.
func (statusIcon *StatusIcon) SetFromFile(filename string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_file(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(filename)
}

// SetFromGIcon makes status_icon display the #GIcon. See
// gtk_status_icon_new_from_gicon() for details.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; you can use g_notification_set_icon() to associate a #GIcon
// with a notification.
func (statusIcon *StatusIcon) SetFromGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GIcon         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_status_icon_set_from_gicon(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(icon)
}

// SetFromIconName makes status_icon display the icon named icon_name from the
// current icon theme. See gtk_status_icon_new_from_icon_name() for details.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; you can use g_notification_set_icon() to associate a #GIcon
// with a notification.
func (statusIcon *StatusIcon) SetFromIconName(iconName string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_icon_name(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(iconName)
}

// SetFromPixbuf makes status_icon display pixbuf. See
// gtk_status_icon_new_from_pixbuf() for details.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; you can use g_notification_set_icon() to associate a #GIcon
// with a notification.
func (statusIcon *StatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GdkPixbuf     // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	if pixbuf != nil {
		_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	C.gtk_status_icon_set_from_pixbuf(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(pixbuf)
}

// SetFromStock makes status_icon display the stock icon with the id stock_id.
// See gtk_status_icon_new_from_stock() for details.
//
// Deprecated: Use gtk_status_icon_set_from_icon_name() instead.
func (statusIcon *StatusIcon) SetFromStock(stockId string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_from_stock(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(stockId)
}

// SetHasTooltip sets the has-tooltip property on status_icon to has_tooltip.
// See StatusIcon:has-tooltip for more information.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, but
// notifications can display an arbitrary amount of text using
// g_notification_set_body().
func (statusIcon *StatusIcon) SetHasTooltip(hasTooltip bool) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	if hasTooltip {
		_arg1 = C.TRUE
	}

	C.gtk_status_icon_set_has_tooltip(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(hasTooltip)
}

// SetName sets the name of this tray icon. This should be a string identifying
// this icon. It is may be used for sorting the icons in the tray and will not
// be shown to the user.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as
// notifications are associated with a unique application identifier by
// #GApplication.
func (statusIcon *StatusIcon) SetName(name string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_name(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(name)
}

// SetScreen sets the Screen where status_icon is displayed; if the icon is
// already mapped, it will be unmapped, and then remapped on the new screen.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as GTK
// typically only has one Screen and notifications are managed by the platform.
func (statusIcon *StatusIcon) SetScreen(screen *gdk.Screen) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.GdkScreen     // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_status_icon_set_screen(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(screen)
}

// SetTitle sets the title of this tray icon. This should be a short,
// human-readable, localized string describing the tray icon. It may be used by
// tools like screen readers to render the tray icon.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; you should use g_notification_set_title() and
// g_notification_set_body() to present text inside your notification.
func (statusIcon *StatusIcon) SetTitle(title string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_title(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(title)
}

// SetTooltipMarkup sets markup as the contents of the tooltip, which is marked
// up with the [Pango text markup language][PangoMarkupFormat].
//
// This function will take care of setting StatusIcon:has-tooltip to TRUE and of
// the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-markup property and gtk_tooltip_set_markup().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) SetTooltipMarkup(markup string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	if markup != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_status_icon_set_tooltip_markup(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(markup)
}

// SetTooltipText sets text as the contents of the tooltip.
//
// This function will take care of setting StatusIcon:has-tooltip to TRUE and of
// the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-text property and gtk_tooltip_set_text().
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function.
func (statusIcon *StatusIcon) SetTooltipText(text string) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 *C.gchar         // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_status_icon_set_tooltip_text(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(text)
}

// SetVisible shows or hides a status icon.
//
// Deprecated: Use #GNotification and Application to provide status
// notifications; there is no direct replacement for this function, as
// notifications are managed by the platform.
func (statusIcon *StatusIcon) SetVisible(visible bool) {
	var _arg0 *C.GtkStatusIcon // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkStatusIcon)(unsafe.Pointer(statusIcon.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_status_icon_set_visible(_arg0, _arg1)
	runtime.KeepAlive(statusIcon)
	runtime.KeepAlive(visible)
}
