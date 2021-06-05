// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_image_get_type()), F: marshalImage},
	})
}

// Image: the Image widget displays an image. Various kinds of object can be
// displayed as an image; most typically, you would load a Pixbuf ("pixel
// buffer") from a file, and then display that. There’s a convenience function
// to do this, gtk_image_new_from_file(), used as follows:
//
//      static gboolean
//      button_press_callback (GtkWidget      *event_box,
//                             GdkEventButton *event,
//                             gpointer        data)
//      {
//        g_print ("Event box clicked at coordinates f,f\n",
//                 event->x, event->y);
//
//        // Returning TRUE means we handled the event, so the signal
//        // emission should be stopped (don’t call any further callbacks
//        // that may be connected). Return FALSE to continue invoking callbacks.
//        return TRUE;
//      }
//
//      static GtkWidget*
//      create_image (void)
//      {
//        GtkWidget *image;
//        GtkWidget *event_box;
//
//        image = gtk_image_new_from_file ("myfile.png");
//
//        event_box = gtk_event_box_new ();
//
//        gtk_container_add (GTK_CONTAINER (event_box), image);
//
//        g_signal_connect (G_OBJECT (event_box),
//                          "button_press_event",
//                          G_CALLBACK (button_press_callback),
//                          image);
//
//        return image;
//      }
//
// When handling events on the event box, keep in mind that coordinates in the
// image may be different from event box coordinates due to the alignment and
// padding settings on the image (see Misc). The simplest way to solve this is
// to set the alignment to 0.0 (left/top), and set the padding to zero. Then the
// origin of the image will be the same as the origin of the event box.
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. GTK+ comes with a program to avoid this, called
// “gdk-pixbuf-csource”. This library allows you to convert an image into a C
// variable declaration, which can then be loaded into a Pixbuf using
// gdk_pixbuf_new_from_inline().
//
//
// CSS nodes
//
// GtkImage has a single CSS node with the name image. The style classes may
// appear on image CSS nodes: .icon-dropshadow, .lowres-icon.
type Image interface {
	Misc
	Buildable

	// Clear resets the image to be empty.
	Clear()
	// Animation gets the PixbufAnimation being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION
	// (see gtk_image_get_storage_type()). The caller of this function does not
	// own a reference to the returned animation.
	Animation() gdkpixbuf.PixbufAnimation
	// GIcon gets the #GIcon and size being displayed by the Image. The storage
	// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
	// gtk_image_get_storage_type()). The caller of this function does not own a
	// reference to the returned #GIcon.
	GIcon() (gicon gio.Icon, size int)
	// IconName gets the icon name and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
	// (see gtk_image_get_storage_type()). The returned string is owned by the
	// Image and should not be freed.
	IconName() (iconName string, size int)
	// IconSet gets the icon set and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET
	// (see gtk_image_get_storage_type()).
	IconSet() (iconSet *IconSet, size int)
	// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of
	// the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
	// gtk_image_get_storage_type()). The caller of this function does not own a
	// reference to the returned pixbuf.
	Pixbuf() gdkpixbuf.Pixbuf
	// PixelSize gets the pixel size used for named icons.
	PixelSize() int
	// Stock gets the stock icon name and size being displayed by the Image. The
	// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK (see
	// gtk_image_get_storage_type()). The returned string is owned by the Image
	// and should not be freed.
	Stock() (stockID string, size int)
	// StorageType gets the type of representation being used by the Image to
	// store image data. If the Image has no image data, the return value will
	// be GTK_IMAGE_EMPTY.
	StorageType() ImageType
	// SetFromAnimation causes the Image to display the given animation (or
	// display nothing, if you set the animation to nil).
	SetFromAnimation(animation gdkpixbuf.PixbufAnimation)
	// SetFromFile: see gtk_image_new_from_file() for details.
	SetFromFile(filename string)
	// SetFromGIcon: see gtk_image_new_from_gicon() for details.
	SetFromGIcon(icon gio.Icon, size int)
	// SetFromIconName: see gtk_image_new_from_icon_name() for details.
	SetFromIconName(iconName string, size int)
	// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
	SetFromIconSet(iconSet *IconSet, size int)
	// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
	SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf)
	// SetFromResource: see gtk_image_new_from_resource() for details.
	SetFromResource(resourcePath string)
	// SetFromStock: see gtk_image_new_from_stock() for details.
	SetFromStock(stockID string, size int)
	// SetFromSurface: see gtk_image_new_from_surface() for details.
	SetFromSurface(surface *cairo.Surface)
	// SetPixelSize sets the pixel size to use for named icons. If the pixel
	// size is set to a value != -1, it is used instead of the icon size set by
	// gtk_image_set_from_icon_name().
	SetPixelSize(pixelSize int)
}

// image implements the Image interface.
type image struct {
	Misc
	Buildable
}

var _ Image = (*image)(nil)

// WrapImage wraps a GObject to the right type. It is
// primarily used internally.
func WrapImage(obj *externglib.Object) Image {
	return Image{
		Misc:      WrapMisc(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalImage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImage(obj), nil
}

// NewImage constructs a class Image.
func NewImage() Image {
	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromAnimation constructs a class Image.
func NewImageFromAnimation(animation gdkpixbuf.PixbufAnimation) Image {
	var arg1 *C.GdkPixbufAnimation

	arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_animation(animation)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromFile constructs a class Image.
func NewImageFromFile(filename string) Image {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_file(filename)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromGIcon constructs a class Image.
func NewImageFromGIcon(icon gio.Icon, size int) Image {
	var arg1 *C.GIcon
	var arg2 C.GtkIconSize

	arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	arg2 = C.GtkIconSize(size)

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_gicon(icon, size)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromIconName constructs a class Image.
func NewImageFromIconName(iconName string, size int) Image {
	var arg1 *C.gchar
	var arg2 C.GtkIconSize

	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.GtkIconSize(size)

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_icon_name(iconName, size)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromIconSet constructs a class Image.
func NewImageFromIconSet(iconSet *IconSet, size int) Image {
	var arg1 *C.GtkIconSet
	var arg2 C.GtkIconSize

	arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet.Native()))
	arg2 = C.GtkIconSize(size)

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_icon_set(iconSet, size)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromPixbuf constructs a class Image.
func NewImageFromPixbuf(pixbuf gdkpixbuf.Pixbuf) Image {
	var arg1 *C.GdkPixbuf

	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_pixbuf(pixbuf)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromResource constructs a class Image.
func NewImageFromResource(resourcePath string) Image {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_resource(resourcePath)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromStock constructs a class Image.
func NewImageFromStock(stockID string, size int) Image {
	var arg1 *C.gchar
	var arg2 C.GtkIconSize

	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.GtkIconSize(size)

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_stock(stockID, size)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// NewImageFromSurface constructs a class Image.
func NewImageFromSurface(surface *cairo.Surface) Image {
	var arg1 *C.cairo_surface_t

	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	var cret C.GtkImage
	var goret1 Image

	cret = C.gtk_image_new_from_surface(surface)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Image)

	return goret1
}

// Clear resets the image to be empty.
func (i image) Clear() {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	C.gtk_image_clear(arg0)
}

// Animation gets the PixbufAnimation being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ANIMATION
// (see gtk_image_get_storage_type()). The caller of this function does not
// own a reference to the returned animation.
func (i image) Animation() gdkpixbuf.PixbufAnimation {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var cret *C.GdkPixbufAnimation
	var goret1 gdkpixbuf.PixbufAnimation

	cret = C.gtk_image_get_animation(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdkpixbuf.PixbufAnimation)

	return goret1
}

// GIcon gets the #GIcon and size being displayed by the Image. The storage
// type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned #GIcon.
func (i image) GIcon() (gicon gio.Icon, size int) {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var arg1 **C.GIcon
	var ret1 *gio.Icon
	var arg2 *C.GtkIconSize
	var ret2 int

	C.gtk_image_get_gicon(arg0, &arg1, &arg2)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1.Native()))).(*gio.Icon)
	ret2 = *C.GtkIconSize(arg2)

	return ret1, ret2
}

// IconName gets the icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_NAME
// (see gtk_image_get_storage_type()). The returned string is owned by the
// Image and should not be freed.
func (i image) IconName() (iconName string, size int) {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var arg1 **C.gchar
	var ret1 string
	var arg2 *C.GtkIconSize
	var ret2 int

	C.gtk_image_get_icon_name(arg0, &arg1, &arg2)

	ret1 = C.GoString(arg1)
	ret2 = *C.GtkIconSize(arg2)

	return ret1, ret2
}

// IconSet gets the icon set and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_ICON_SET
// (see gtk_image_get_storage_type()).
func (i image) IconSet() (iconSet *IconSet, size int) {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var arg1 **C.GtkIconSet
	var ret1 **IconSet
	var arg2 *C.GtkIconSize
	var ret2 int

	C.gtk_image_get_icon_set(arg0, &arg1, &arg2)

	ret1 = WrapIconSet(unsafe.Pointer(arg1))
	ret2 = *C.GtkIconSize(arg2)

	return ret1, ret2
}

// Pixbuf gets the Pixbuf being displayed by the Image. The storage type of
// the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_image_get_storage_type()). The caller of this function does not own a
// reference to the returned pixbuf.
func (i image) Pixbuf() gdkpixbuf.Pixbuf {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var cret *C.GdkPixbuf
	var goret1 gdkpixbuf.Pixbuf

	cret = C.gtk_image_get_pixbuf(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdkpixbuf.Pixbuf)

	return goret1
}

// PixelSize gets the pixel size used for named icons.
func (i image) PixelSize() int {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gtk_image_get_pixel_size(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// Stock gets the stock icon name and size being displayed by the Image. The
// storage type of the image must be GTK_IMAGE_EMPTY or GTK_IMAGE_STOCK (see
// gtk_image_get_storage_type()). The returned string is owned by the Image
// and should not be freed.
func (i image) Stock() (stockID string, size int) {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var arg1 **C.gchar
	var ret1 string
	var arg2 *C.GtkIconSize
	var ret2 int

	C.gtk_image_get_stock(arg0, &arg1, &arg2)

	ret1 = C.GoString(arg1)
	ret2 = *C.GtkIconSize(arg2)

	return ret1, ret2
}

// StorageType gets the type of representation being used by the Image to
// store image data. If the Image has no image data, the return value will
// be GTK_IMAGE_EMPTY.
func (i image) StorageType() ImageType {
	var arg0 *C.GtkImage

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))

	var cret C.GtkImageType
	var goret1 ImageType

	cret = C.gtk_image_get_storage_type(arg0)

	goret1 = ImageType(cret)

	return goret1
}

// SetFromAnimation causes the Image to display the given animation (or
// display nothing, if you set the animation to nil).
func (i image) SetFromAnimation(animation gdkpixbuf.PixbufAnimation) {
	var arg0 *C.GtkImage
	var arg1 *C.GdkPixbufAnimation

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GdkPixbufAnimation)(unsafe.Pointer(animation.Native()))

	C.gtk_image_set_from_animation(arg0, animation)
}

// SetFromFile: see gtk_image_new_from_file() for details.
func (i image) SetFromFile(filename string) {
	var arg0 *C.GtkImage
	var arg1 *C.gchar

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_image_set_from_file(arg0, filename)
}

// SetFromGIcon: see gtk_image_new_from_gicon() for details.
func (i image) SetFromGIcon(icon gio.Icon, size int) {
	var arg0 *C.GtkImage
	var arg1 *C.GIcon
	var arg2 C.GtkIconSize

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_gicon(arg0, icon, size)
}

// SetFromIconName: see gtk_image_new_from_icon_name() for details.
func (i image) SetFromIconName(iconName string, size int) {
	var arg0 *C.GtkImage
	var arg1 *C.gchar
	var arg2 C.GtkIconSize

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_name(arg0, iconName, size)
}

// SetFromIconSet: see gtk_image_new_from_icon_set() for details.
func (i image) SetFromIconSet(iconSet *IconSet, size int) {
	var arg0 *C.GtkImage
	var arg1 *C.GtkIconSet
	var arg2 C.GtkIconSize

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GtkIconSet)(unsafe.Pointer(iconSet.Native()))
	arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_icon_set(arg0, iconSet, size)
}

// SetFromPixbuf: see gtk_image_new_from_pixbuf() for details.
func (i image) SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkImage
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_image_set_from_pixbuf(arg0, pixbuf)
}

// SetFromResource: see gtk_image_new_from_resource() for details.
func (i image) SetFromResource(resourcePath string) {
	var arg0 *C.GtkImage
	var arg1 *C.gchar

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_image_set_from_resource(arg0, resourcePath)
}

// SetFromStock: see gtk_image_new_from_stock() for details.
func (i image) SetFromStock(stockID string, size int) {
	var arg0 *C.GtkImage
	var arg1 *C.gchar
	var arg2 C.GtkIconSize

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.GtkIconSize(size)

	C.gtk_image_set_from_stock(arg0, stockID, size)
}

// SetFromSurface: see gtk_image_new_from_surface() for details.
func (i image) SetFromSurface(surface *cairo.Surface) {
	var arg0 *C.GtkImage
	var arg1 *C.cairo_surface_t

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))

	C.gtk_image_set_from_surface(arg0, surface)
}

// SetPixelSize sets the pixel size to use for named icons. If the pixel
// size is set to a value != -1, it is used instead of the icon size set by
// gtk_image_set_from_icon_name().
func (i image) SetPixelSize(pixelSize int) {
	var arg0 *C.GtkImage
	var arg1 C.gint

	arg0 = (*C.GtkImage)(unsafe.Pointer(i.Native()))
	arg1 = C.gint(pixelSize)

	C.gtk_image_set_pixel_size(arg0, pixelSize)
}

type ImagePrivate struct {
	native C.GtkImagePrivate
}

// WrapImagePrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapImagePrivate(ptr unsafe.Pointer) *ImagePrivate {
	if ptr == nil {
		return nil
	}

	return (*ImagePrivate)(ptr)
}

func marshalImagePrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapImagePrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *ImagePrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}
