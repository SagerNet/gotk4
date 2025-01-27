// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_tool_palette_drag_targets_get_type()), F: marshalToolPaletteDragTargets},
		{T: externglib.Type(C.gtk_tool_palette_get_type()), F: marshalToolPaletter},
	})
}

// ToolPaletteDragTargets flags used to specify the supported drag targets.
type ToolPaletteDragTargets int

const (
	// ToolPaletteDragItems: support drag of items.
	ToolPaletteDragItems ToolPaletteDragTargets = 0b1
	// ToolPaletteDragGroups: support drag of groups.
	ToolPaletteDragGroups ToolPaletteDragTargets = 0b10
)

func marshalToolPaletteDragTargets(p uintptr) (interface{}, error) {
	return ToolPaletteDragTargets(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for ToolPaletteDragTargets.
func (t ToolPaletteDragTargets) String() string {
	if t == 0 {
		return "ToolPaletteDragTargets(0)"
	}

	var builder strings.Builder
	builder.Grow(42)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case ToolPaletteDragItems:
			builder.WriteString("Items|")
		case ToolPaletteDragGroups:
			builder.WriteString("Groups|")
		default:
			builder.WriteString(fmt.Sprintf("ToolPaletteDragTargets(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t ToolPaletteDragTargets) Has(other ToolPaletteDragTargets) bool {
	return (t & other) == other
}

// ToolPalette allows you to add ToolItems to a palette-like container with
// different categories and drag and drop support.
//
// A ToolPalette is created with a call to gtk_tool_palette_new().
//
// ToolItems cannot be added directly to a ToolPalette - instead they are added
// to a ToolItemGroup which can than be added to a ToolPalette. To add a
// ToolItemGroup to a ToolPalette, use gtk_container_add().
//
//    static void
//    passive_canvas_drag_data_received (GtkWidget        *widget,
//                                       GdkDragContext   *context,
//                                       gint              x,
//                                       gint              y,
//                                       GtkSelectionData *selection,
//                                       guint             info,
//                                       guint             time,
//                                       gpointer          data)
//    {
//      GtkWidget *palette;
//      GtkWidget *item;
//
//      // Get the dragged item
//      palette = gtk_widget_get_ancestor (gtk_drag_get_source_widget (context),
//                                         GTK_TYPE_TOOL_PALETTE);
//      if (palette != NULL)
//        item = gtk_tool_palette_get_drag_item (GTK_TOOL_PALETTE (palette),
//                                               selection);
//
//      // Do something with item
//    }
//
//    GtkWidget *target, palette;
//
//    palette = gtk_tool_palette_new ();
//    target = gtk_drawing_area_new ();
//
//    g_signal_connect (G_OBJECT (target), "drag-data-received",
//                      G_CALLBACK (passive_canvas_drag_data_received), NULL);
//    gtk_tool_palette_add_drag_dest (GTK_TOOL_PALETTE (palette), target,
//                                    GTK_DEST_DEFAULT_ALL,
//                                    GTK_TOOL_PALETTE_DRAG_ITEMS,
//                                    GDK_ACTION_COPY);
//
//
// CSS nodes
//
// GtkToolPalette has a single CSS node named toolpalette.
type ToolPalette struct {
	Container

	Orientable
	Scrollable
	*externglib.Object
}

func WrapToolPalette(obj *externglib.Object) *ToolPalette {
	return &ToolPalette{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
		Scrollable: Scrollable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalToolPaletter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolPalette(obj), nil
}

// NewToolPalette creates a new tool palette.
func NewToolPalette() *ToolPalette {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tool_palette_new()

	var _toolPalette *ToolPalette // out

	_toolPalette = WrapToolPalette(externglib.Take(unsafe.Pointer(_cret)))

	return _toolPalette
}

// AddDragDest sets palette as drag source (see
// gtk_tool_palette_set_drag_source()) and sets widget as a drag destination for
// drags from palette. See gtk_drag_dest_set().
func (palette *ToolPalette) AddDragDest(widget Widgetter, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction) {
	var _arg0 *C.GtkToolPalette           // out
	var _arg1 *C.GtkWidget                // out
	var _arg2 C.GtkDestDefaults           // out
	var _arg3 C.GtkToolPaletteDragTargets // out
	var _arg4 C.GdkDragAction             // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkDestDefaults(flags)
	_arg3 = C.GtkToolPaletteDragTargets(targets)
	_arg4 = C.GdkDragAction(actions)

	C.gtk_tool_palette_add_drag_dest(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(targets)
	runtime.KeepAlive(actions)
}

// DragItem: get the dragged item from the selection. This could be a ToolItem
// or a ToolItemGroup.
func (palette *ToolPalette) DragItem(selection *SelectionData) Widgetter {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkSelectionData // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selection)))

	_cret = C.gtk_tool_palette_get_drag_item(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(selection)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// DropGroup gets the group at position (x, y).
func (palette *ToolPalette) DropGroup(x int, y int) *ToolItemGroup {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 C.gint              // out
	var _arg2 C.gint              // out
	var _cret *C.GtkToolItemGroup // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_palette_get_drop_group(_arg0, _arg1, _arg2)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _toolItemGroup *ToolItemGroup // out

	if _cret != nil {
		_toolItemGroup = WrapToolItemGroup(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItemGroup
}

// DropItem gets the item at position (x, y). See
// gtk_tool_palette_get_drop_group().
func (palette *ToolPalette) DropItem(x int, y int) *ToolItem {
	var _arg0 *C.GtkToolPalette // out
	var _arg1 C.gint            // out
	var _arg2 C.gint            // out
	var _cret *C.GtkToolItem    // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_palette_get_drop_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _toolItem *ToolItem // out

	if _cret != nil {
		_toolItem = WrapToolItem(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItem
}

// Exclusive gets whether group is exclusive or not. See
// gtk_tool_palette_set_exclusive().
func (palette *ToolPalette) Exclusive(group *ToolItemGroup) bool {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_palette_get_exclusive(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expand gets whether group should be given extra space. See
// gtk_tool_palette_set_expand().
func (palette *ToolPalette) Expand(group *ToolItemGroup) bool {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_palette_get_expand(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GroupPosition gets the position of group in palette as index. See
// gtk_tool_palette_set_group_position().
func (palette *ToolPalette) GroupPosition(group *ToolItemGroup) int {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_palette_get_group_position(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// HAdjustment gets the horizontal adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_hadjustment().
func (palette *ToolPalette) HAdjustment() *Adjustment {
	var _arg0 *C.GtkToolPalette // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_hadjustment(_arg0)
	runtime.KeepAlive(palette)

	var _adjustment *Adjustment // out

	_adjustment = WrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// IconSize gets the size of icons in the tool palette. See
// gtk_tool_palette_set_icon_size().
func (palette *ToolPalette) IconSize() int {
	var _arg0 *C.GtkToolPalette // out
	var _cret C.GtkIconSize     // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_icon_size(_arg0)
	runtime.KeepAlive(palette)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Style gets the style (icons, text or both) of items in the tool palette.
func (palette *ToolPalette) Style() ToolbarStyle {
	var _arg0 *C.GtkToolPalette // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_style(_arg0)
	runtime.KeepAlive(palette)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// VAdjustment gets the vertical adjustment of the tool palette.
//
// Deprecated: Use gtk_scrollable_get_vadjustment().
func (palette *ToolPalette) VAdjustment() *Adjustment {
	var _arg0 *C.GtkToolPalette // out
	var _cret *C.GtkAdjustment  // in

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	_cret = C.gtk_tool_palette_get_vadjustment(_arg0)
	runtime.KeepAlive(palette)

	var _adjustment *Adjustment // out

	_adjustment = WrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// SetDragSource sets the tool palette as a drag source. Enables all groups and
// items in the tool palette as drag sources on button 1 and button 3 press with
// copy and move actions. See gtk_drag_source_set().
func (palette *ToolPalette) SetDragSource(targets ToolPaletteDragTargets) {
	var _arg0 *C.GtkToolPalette           // out
	var _arg1 C.GtkToolPaletteDragTargets // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.GtkToolPaletteDragTargets(targets)

	C.gtk_tool_palette_set_drag_source(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(targets)
}

// SetExclusive sets whether the group should be exclusive or not. If an
// exclusive group is expanded all other groups are collapsed.
func (palette *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if exclusive {
		_arg2 = C.TRUE
	}

	C.gtk_tool_palette_set_exclusive(_arg0, _arg1, _arg2)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(exclusive)
}

// SetExpand sets whether the group should be given extra space.
func (palette *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tool_palette_set_expand(_arg0, _arg1, _arg2)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(expand)
}

// SetGroupPosition sets the position of the group as an index of the tool
// palette. If position is 0 the group will become the first child, if position
// is -1 it will become the last child.
func (palette *ToolPalette) SetGroupPosition(group *ToolItemGroup, position int) {
	var _arg0 *C.GtkToolPalette   // out
	var _arg1 *C.GtkToolItemGroup // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_palette_set_group_position(_arg0, _arg1, _arg2)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(group)
	runtime.KeepAlive(position)
}

// SetIconSize sets the size of icons in the tool palette.
func (palette *ToolPalette) SetIconSize(iconSize int) {
	var _arg0 *C.GtkToolPalette // out
	var _arg1 C.GtkIconSize     // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_tool_palette_set_icon_size(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(iconSize)
}

// SetStyle sets the style (text, icons or both) of items in the tool palette.
func (palette *ToolPalette) SetStyle(style ToolbarStyle) {
	var _arg0 *C.GtkToolPalette // out
	var _arg1 C.GtkToolbarStyle // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))
	_arg1 = C.GtkToolbarStyle(style)

	C.gtk_tool_palette_set_style(_arg0, _arg1)
	runtime.KeepAlive(palette)
	runtime.KeepAlive(style)
}

// UnsetIconSize unsets the tool palette icon size set with
// gtk_tool_palette_set_icon_size(), so that user preferences will be used to
// determine the icon size.
func (palette *ToolPalette) UnsetIconSize() {
	var _arg0 *C.GtkToolPalette // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	C.gtk_tool_palette_unset_icon_size(_arg0)
	runtime.KeepAlive(palette)
}

// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(), so
// that user preferences will be used to determine the toolbar style.
func (palette *ToolPalette) UnsetStyle() {
	var _arg0 *C.GtkToolPalette // out

	_arg0 = (*C.GtkToolPalette)(unsafe.Pointer(palette.Native()))

	C.gtk_tool_palette_unset_style(_arg0)
	runtime.KeepAlive(palette)
}

// ToolPaletteGetDragTargetGroup: get the target entry for a dragged
// ToolItemGroup.
func ToolPaletteGetDragTargetGroup() *TargetEntry {
	var _cret *C.GtkTargetEntry // in

	_cret = C.gtk_tool_palette_get_drag_target_group()

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _targetEntry
}

// ToolPaletteGetDragTargetItem gets the target entry for a dragged ToolItem.
func ToolPaletteGetDragTargetItem() *TargetEntry {
	var _cret *C.GtkTargetEntry // in

	_cret = C.gtk_tool_palette_get_drag_target_item()

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _targetEntry
}
