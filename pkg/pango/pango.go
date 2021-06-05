// Code generated by girgen. DO NOT EDIT.

package pango

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_alignment_get_type()), F: marshalAlignment},
		{T: externglib.Type(C.pango_attr_type_get_type()), F: marshalAttrType},
		{T: externglib.Type(C.pango_bidi_type_get_type()), F: marshalBidiType},
		{T: externglib.Type(C.pango_coverage_level_get_type()), F: marshalCoverageLevel},
		{T: externglib.Type(C.pango_direction_get_type()), F: marshalDirection},
		{T: externglib.Type(C.pango_ellipsize_mode_get_type()), F: marshalEllipsizeMode},
		{T: externglib.Type(C.pango_gravity_get_type()), F: marshalGravity},
		{T: externglib.Type(C.pango_gravity_hint_get_type()), F: marshalGravityHint},
		{T: externglib.Type(C.pango_overline_get_type()), F: marshalOverline},
		{T: externglib.Type(C.pango_render_part_get_type()), F: marshalRenderPart},
		{T: externglib.Type(C.pango_script_get_type()), F: marshalScript},
		{T: externglib.Type(C.pango_stretch_get_type()), F: marshalStretch},
		{T: externglib.Type(C.pango_style_get_type()), F: marshalStyle},
		{T: externglib.Type(C.pango_tab_align_get_type()), F: marshalTabAlign},
		{T: externglib.Type(C.pango_underline_get_type()), F: marshalUnderline},
		{T: externglib.Type(C.pango_variant_get_type()), F: marshalVariant},
		{T: externglib.Type(C.pango_weight_get_type()), F: marshalWeight},
		{T: externglib.Type(C.pango_wrap_mode_get_type()), F: marshalWrapMode},
		{T: externglib.Type(C.pango_font_mask_get_type()), F: marshalFontMask},
		{T: externglib.Type(C.pango_shape_flags_get_type()), F: marshalShapeFlags},
		{T: externglib.Type(C.pango_show_flags_get_type()), F: marshalShowFlags},
		{T: externglib.Type(C.pango_coverage_get_type()), F: marshalCoverage},
	})
}

// Alignment: `PangoAlignment` describes how to align the lines of a
// `PangoLayout` within the available space.
//
// If the `PangoLayout` is set to justify using
// [method@Pango.Layout.set_justify], this only has effect for partial lines.
type Alignment int

const (
	// AlignmentLeft: put all available space on the right
	AlignmentLeft Alignment = 0
	// AlignmentCenter: center the line within the available space
	AlignmentCenter Alignment = 1
	// AlignmentRight: put all available space on the left
	AlignmentRight Alignment = 2
)

func marshalAlignment(p uintptr) (interface{}, error) {
	return Alignment(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AttrType: the `PangoAttrType` distinguishes between different types of
// attributes.
//
// Along with the predefined values, it is possible to allocate additional
// values for custom attributes using [type_func@attr_type_register]. The
// predefined values are given below. The type of structure used to store the
// attribute is listed in parentheses after the description.
type AttrType int

const (
	// AttrTypeInvalid does not happen
	AttrTypeInvalid AttrType = 0
	// AttrTypeLanguage: language ([struct@Pango.AttrLanguage])
	AttrTypeLanguage AttrType = 1
	// AttrTypeFamily: font family name list ([struct@Pango.AttrString])
	AttrTypeFamily AttrType = 2
	// AttrTypeStyle: font slant style ([struct@Pango.AttrInt])
	AttrTypeStyle AttrType = 3
	// AttrTypeWeight: font weight ([struct@Pango.AttrInt])
	AttrTypeWeight AttrType = 4
	// AttrTypeVariant: font variant (normal or small caps)
	// ([struct@Pango.AttrInt])
	AttrTypeVariant AttrType = 5
	// AttrTypeStretch: font stretch ([struct@Pango.AttrInt])
	AttrTypeStretch AttrType = 6
	// AttrTypeSize: font size in points scaled by PANGO_SCALE
	// ([struct@Pango.AttrInt])
	AttrTypeSize AttrType = 7
	// AttrTypeFontDesc: font description ([struct@Pango.AttrFontDesc])
	AttrTypeFontDesc AttrType = 8
	// AttrTypeForeground: foreground color ([struct@Pango.AttrColor])
	AttrTypeForeground AttrType = 9
	// AttrTypeBackground: background color ([struct@Pango.AttrColor])
	AttrTypeBackground AttrType = 10
	// AttrTypeUnderline: whether the text has an underline
	// ([struct@Pango.AttrInt])
	AttrTypeUnderline AttrType = 11
	// AttrTypeStrikethrough: whether the text is struck-through
	// ([struct@Pango.AttrInt])
	AttrTypeStrikethrough AttrType = 12
	// AttrTypeRise: baseline displacement ([struct@Pango.AttrInt])
	AttrTypeRise AttrType = 13
	// AttrTypeShape: shape ([struct@Pango.AttrShape])
	AttrTypeShape AttrType = 14
	// AttrTypeScale: font size scale factor ([struct@Pango.AttrFloat])
	AttrTypeScale AttrType = 15
	// AttrTypeFallback: whether fallback is enabled ([struct@Pango.AttrInt])
	AttrTypeFallback AttrType = 16
	// AttrTypeLetterSpacing: letter spacing ([struct@PangoAttrInt])
	AttrTypeLetterSpacing AttrType = 17
	// AttrTypeUnderlineColor: underline color ([struct@Pango.AttrColor])
	AttrTypeUnderlineColor AttrType = 18
	// AttrTypeStrikethroughColor: strikethrough color
	// ([struct@Pango.AttrColor])
	AttrTypeStrikethroughColor AttrType = 19
	// AttrTypeAbsoluteSize: font size in pixels scaled by PANGO_SCALE
	// ([struct@Pango.AttrInt])
	AttrTypeAbsoluteSize AttrType = 20
	// AttrTypeGravity: base text gravity ([struct@Pango.AttrInt])
	AttrTypeGravity AttrType = 21
	// AttrTypeGravityHint: gravity hint ([struct@Pango.AttrInt])
	AttrTypeGravityHint AttrType = 22
	// AttrTypeFontFeatures: openType font features ([struct@Pango.AttrString]).
	// Since 1.38
	AttrTypeFontFeatures AttrType = 23
	// AttrTypeForegroundAlpha: foreground alpha ([struct@Pango.AttrInt]). Since
	// 1.38
	AttrTypeForegroundAlpha AttrType = 24
	// AttrTypeBackgroundAlpha: background alpha ([struct@Pango.AttrInt]). Since
	// 1.38
	AttrTypeBackgroundAlpha AttrType = 25
	// AttrTypeAllowBreaks: whether breaks are allowed ([struct@Pango.AttrInt]).
	// Since 1.44
	AttrTypeAllowBreaks AttrType = 26
	// AttrTypeShow: how to render invisible characters
	// ([struct@Pango.AttrInt]). Since 1.44
	AttrTypeShow AttrType = 27
	// AttrTypeInsertHyphens: whether to insert hyphens at intra-word line
	// breaks ([struct@Pango.AttrInt]). Since 1.44
	AttrTypeInsertHyphens AttrType = 28
	// AttrTypeOverline: whether the text has an overline
	// ([struct@Pango.AttrInt]). Since 1.46
	AttrTypeOverline AttrType = 29
	// AttrTypeOverlineColor: overline color ([struct@Pango.AttrColor]). Since
	// 1.46
	AttrTypeOverlineColor AttrType = 30
)

func marshalAttrType(p uintptr) (interface{}, error) {
	return AttrType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// BidiType: `PangoBidiType` represents the bidirectional character type of a
// Unicode character as specified by the <ulink
// url="http://www.unicode.org/reports/tr9/">Unicode bidirectional
// algorithm</ulink>.
type BidiType int

const (
	// BidiTypeL: left-to-Right
	BidiTypeL BidiType = 0
	// BidiTypeLre: left-to-Right Embedding
	BidiTypeLre BidiType = 1
	// BidiTypeLro: left-to-Right Override
	BidiTypeLro BidiType = 2
	// BidiTypeR: right-to-Left
	BidiTypeR BidiType = 3
	// BidiTypeAl: right-to-Left Arabic
	BidiTypeAl BidiType = 4
	// BidiTypeRle: right-to-Left Embedding
	BidiTypeRle BidiType = 5
	// BidiTypeRlo: right-to-Left Override
	BidiTypeRlo BidiType = 6
	// BidiTypePdf: pop Directional Format
	BidiTypePdf BidiType = 7
	// BidiTypeEn: european Number
	BidiTypeEn BidiType = 8
	// BidiTypeES: european Number Separator
	BidiTypeES BidiType = 9
	// BidiTypeEt: european Number Terminator
	BidiTypeEt BidiType = 10
	// BidiTypeAn: arabic Number
	BidiTypeAn BidiType = 11
	// BidiTypeCs: common Number Separator
	BidiTypeCs BidiType = 12
	// BidiTypeNsm: nonspacing Mark
	BidiTypeNsm BidiType = 13
	// BidiTypeBn: boundary Neutral
	BidiTypeBn BidiType = 14
	// BidiTypeB: paragraph Separator
	BidiTypeB BidiType = 15
	// BidiTypeS: segment Separator
	BidiTypeS BidiType = 16
	// BidiTypeWs: whitespace
	BidiTypeWs BidiType = 17
	// BidiTypeOn: other Neutrals
	BidiTypeOn BidiType = 18
)

func marshalBidiType(p uintptr) (interface{}, error) {
	return BidiType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// CoverageLevel: `PangoCoverageLevel` is used to indicate how well a font can
// represent a particular Unicode character for a particular script.
//
// Since 1.44, only PANGO_COVERAGE_NONE and PANGO_COVERAGE_EXACT will be
// returned.
type CoverageLevel int

const (
	// CoverageLevelNone: the character is not representable with the font.
	CoverageLevelNone CoverageLevel = 0
	// CoverageLevelFallback: the character is represented in a way that may be
	// comprehensible but is not the correct graphical form. For instance, a
	// Hangul character represented as a a sequence of Jamos, or a Latin
	// transliteration of a Cyrillic word.
	CoverageLevelFallback CoverageLevel = 1
	// CoverageLevelApproximate: the character is represented as basically the
	// correct graphical form, but with a stylistic variant inappropriate for
	// the current script.
	CoverageLevelApproximate CoverageLevel = 2
	// CoverageLevelExact: the character is represented as the correct graphical
	// form.
	CoverageLevelExact CoverageLevel = 3
)

func marshalCoverageLevel(p uintptr) (interface{}, error) {
	return CoverageLevel(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Direction: `PangoDirection` represents a direction in the Unicode
// bidirectional algorithm.
//
// Not every value in this enumeration makes sense for every usage of
// `PangoDirection`; for example, the return value of [func@unichar_direction]
// and [func@find_base_dir] cannot be PANGO_DIRECTION_WEAK_LTR or
// PANGO_DIRECTION_WEAK_RTL, since every character is either neutral or has a
// strong direction; on the other hand PANGO_DIRECTION_NEUTRAL doesn't make
// sense to pass to [func@itemize_with_base_dir].
//
// The PANGO_DIRECTION_TTB_LTR, PANGO_DIRECTION_TTB_RTL values come from an
// earlier interpretation of this enumeration as the writing direction of a
// block of text and are no longer used; See `PangoGravity` for how vertical
// text is handled in Pango.
//
// If you are interested in text direction, you should really use fribidi
// directly. `PangoDirection` is only retained because it is used in some public
// apis.
type Direction int

const (
	// DirectionLTR: a strong left-to-right direction
	DirectionLTR Direction = 0
	// DirectionRTL: a strong right-to-left direction
	DirectionRTL Direction = 1
	// DirectionTtbLTR: deprecated value; treated the same as
	// PANGO_DIRECTION_RTL.
	DirectionTtbLTR Direction = 2
	// DirectionTtbRTL: deprecated value; treated the same as
	// PANGO_DIRECTION_LTR
	DirectionTtbRTL Direction = 3
	// DirectionWeakLTR: a weak left-to-right direction
	DirectionWeakLTR Direction = 4
	// DirectionWeakRTL: a weak right-to-left direction
	DirectionWeakRTL Direction = 5
	// DirectionNeutral: no direction specified
	DirectionNeutral Direction = 6
)

func marshalDirection(p uintptr) (interface{}, error) {
	return Direction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EllipsizeMode: `PangoEllipsizeMode` describes what sort of ellipsization
// should be applied to text.
//
// In the ellipsization process characters are removed from the text in order to
// make it fit to a given width and replaced with an ellipsis.
type EllipsizeMode int

const (
	// EllipsizeModeNone: no ellipsization
	EllipsizeModeNone EllipsizeMode = 0
	// EllipsizeModeStart: omit characters at the start of the text
	EllipsizeModeStart EllipsizeMode = 1
	// EllipsizeModeMiddle: omit characters in the middle of the text
	EllipsizeModeMiddle EllipsizeMode = 2
	// EllipsizeModeEnd: omit characters at the end of the text
	EllipsizeModeEnd EllipsizeMode = 3
)

func marshalEllipsizeMode(p uintptr) (interface{}, error) {
	return EllipsizeMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Gravity: `PangoGravity` represents the orientation of glyphs in a segment of
// text.
//
// This is useful when rendering vertical text layouts. In those situations, the
// layout is rotated using a non-identity [struct@Pango.Matrix], and then glyph
// orientation is controlled using `PangoGravity`.
//
// Not every value in this enumeration makes sense for every usage of
// `PangoGravity`; for example, PANGO_GRAVITY_AUTO only can be passed to
// [method@Pango.Context.set_base_gravity] and can only be returned by
// [method@Pango.Context.get_base_gravity].
//
// See also: [enum@Pango.GravityHint]
type Gravity int

const (
	// GravitySouth glyphs stand upright (default)
	GravitySouth Gravity = 0
	// GravityEast glyphs are rotated 90 degrees clockwise
	GravityEast Gravity = 1
	// GravityNorth glyphs are upside-down
	GravityNorth Gravity = 2
	// GravityWest glyphs are rotated 90 degrees counter-clockwise
	GravityWest Gravity = 3
	// GravityAuto: gravity is resolved from the context matrix
	GravityAuto Gravity = 4
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// GravityHint: `PangoGravityHint` defines how horizontal scripts should behave
// in a vertical context.
//
// That is, English excerpts in a vertical paragraph for example.
//
// See also [enum@Pango.Gravity]
type GravityHint int

const (
	// GravityHintNatural scripts will take their natural gravity based on the
	// base gravity and the script. This is the default.
	GravityHintNatural GravityHint = 0
	// GravityHintStrong always use the base gravity set, regardless of the
	// script.
	GravityHintStrong GravityHint = 1
	// GravityHintLine: for scripts not in their natural direction (eg. Latin in
	// East gravity), choose per-script gravity such that every script respects
	// the line progression. This means, Latin and Arabic will take opposite
	// gravities and both flow top-to-bottom for example.
	GravityHintLine GravityHint = 2
)

func marshalGravityHint(p uintptr) (interface{}, error) {
	return GravityHint(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Overline: the `PangoOverline` enumeration is used to specify whether text
// should be overlined, and if so, the type of line.
type Overline int

const (
	// OverlineNone: no overline should be drawn
	OverlineNone Overline = 0
	// OverlineSingle: draw a single line above the ink extents of the text
	// being underlined.
	OverlineSingle Overline = 1
)

func marshalOverline(p uintptr) (interface{}, error) {
	return Overline(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RenderPart defines different items to render for such purposes as setting
// colors.
type RenderPart int

const (
	// RenderPartForeground: the text itself
	RenderPartForeground RenderPart = 0
	// RenderPartBackground: the area behind the text
	RenderPartBackground RenderPart = 1
	// RenderPartUnderline: underlines
	RenderPartUnderline RenderPart = 2
	// RenderPartStrikethrough: strikethrough lines
	RenderPartStrikethrough RenderPart = 3
	// RenderPartOverline: overlines
	RenderPartOverline RenderPart = 4
)

func marshalRenderPart(p uintptr) (interface{}, error) {
	return RenderPart(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Script: the `PangoScript` enumeration identifies different writing systems.
//
// The values correspond to the names as defined in the Unicode standard. See
// Unicode Standard Annex 24: Script names
// (http://www.unicode.org/reports/tr24/)
//
// Note that this enumeration is deprecated and will not be updated to include
// values in newer versions of the Unicode standard. Applications should use the
// `GUnicodeScript` enumeration instead, whose values are interchangeable with
// `PangoScript`.
type Script int

const (
	// ScriptInvalidCode: a value never returned from pango_script_for_unichar()
	ScriptInvalidCode Script = -1
	// ScriptCommon: a character used by multiple different scripts
	ScriptCommon Script = 0
	// ScriptInherited: a mark glyph that takes its script from the base glyph
	// to which it is attached
	ScriptInherited Script = 1
	// ScriptArabic: arabic
	ScriptArabic Script = 2
	// ScriptArmenian: armenian
	ScriptArmenian Script = 3
	// ScriptBengali: bengali
	ScriptBengali Script = 4
	// ScriptBopomofo: bopomofo
	ScriptBopomofo Script = 5
	// ScriptCherokee: cherokee
	ScriptCherokee Script = 6
	// ScriptCoptic: coptic
	ScriptCoptic Script = 7
	// ScriptCyrillic: cyrillic
	ScriptCyrillic Script = 8
	// ScriptDeseret: deseret
	ScriptDeseret Script = 9
	// ScriptDevanagari: devanagari
	ScriptDevanagari Script = 10
	// ScriptEthiopic: ethiopic
	ScriptEthiopic Script = 11
	// ScriptGeorgian: georgian
	ScriptGeorgian Script = 12
	// ScriptGothic: gothic
	ScriptGothic Script = 13
	// ScriptGreek: greek
	ScriptGreek Script = 14
	// ScriptGujarati: gujarati
	ScriptGujarati Script = 15
	// ScriptGurmukhi: gurmukhi
	ScriptGurmukhi Script = 16
	// ScriptHan: han
	ScriptHan Script = 17
	// ScriptHangul: hangul
	ScriptHangul Script = 18
	// ScriptHebrew: hebrew
	ScriptHebrew Script = 19
	// ScriptHiragana: hiragana
	ScriptHiragana Script = 20
	// ScriptKannada: kannada
	ScriptKannada Script = 21
	// ScriptKatakana: katakana
	ScriptKatakana Script = 22
	// ScriptKhmer: khmer
	ScriptKhmer Script = 23
	// ScriptLao: lao
	ScriptLao Script = 24
	// ScriptLatin: latin
	ScriptLatin Script = 25
	// ScriptMalayalam: malayalam
	ScriptMalayalam Script = 26
	// ScriptMongolian: mongolian
	ScriptMongolian Script = 27
	// ScriptMyanmar: myanmar
	ScriptMyanmar Script = 28
	// ScriptOgham: ogham
	ScriptOgham Script = 29
	// ScriptOldItalic: old Italic
	ScriptOldItalic Script = 30
	// ScriptOriya: oriya
	ScriptOriya Script = 31
	// ScriptRunic: runic
	ScriptRunic Script = 32
	// ScriptSinhala: sinhala
	ScriptSinhala Script = 33
	// ScriptSyriac: syriac
	ScriptSyriac Script = 34
	// ScriptTamil: tamil
	ScriptTamil Script = 35
	// ScriptTelugu: telugu
	ScriptTelugu Script = 36
	// ScriptThaana: thaana
	ScriptThaana Script = 37
	// ScriptThai: thai
	ScriptThai Script = 38
	// ScriptTibetan: tibetan
	ScriptTibetan Script = 39
	// ScriptCanadianAboriginal: canadian Aboriginal
	ScriptCanadianAboriginal Script = 40
	// ScriptYi: yi
	ScriptYi Script = 41
	// ScriptTagalog: tagalog
	ScriptTagalog Script = 42
	// ScriptHanunoo: hanunoo
	ScriptHanunoo Script = 43
	// ScriptBuhid: buhid
	ScriptBuhid Script = 44
	// ScriptTagbanwa: tagbanwa
	ScriptTagbanwa Script = 45
	// ScriptBraille: braille
	ScriptBraille Script = 46
	// ScriptCypriot: cypriot
	ScriptCypriot Script = 47
	// ScriptLimbu: limbu
	ScriptLimbu Script = 48
	// ScriptOsmanya: osmanya
	ScriptOsmanya Script = 49
	// ScriptShavian: shavian
	ScriptShavian Script = 50
	// ScriptLinearB: linear B
	ScriptLinearB Script = 51
	// ScriptTaiLe: tai Le
	ScriptTaiLe Script = 52
	// ScriptUgaritic: ugaritic
	ScriptUgaritic Script = 53
	// ScriptNewTaiLue: new Tai Lue. Since 1.10
	ScriptNewTaiLue Script = 54
	// ScriptBuginese: buginese. Since 1.10
	ScriptBuginese Script = 55
	// ScriptGlagolitic: glagolitic. Since 1.10
	ScriptGlagolitic Script = 56
	// ScriptTifinagh: tifinagh. Since 1.10
	ScriptTifinagh Script = 57
	// ScriptSylotiNagri: syloti Nagri. Since 1.10
	ScriptSylotiNagri Script = 58
	// ScriptOldPersian: old Persian. Since 1.10
	ScriptOldPersian Script = 59
	// ScriptKharoshthi: kharoshthi. Since 1.10
	ScriptKharoshthi Script = 60
	// ScriptUnknown: an unassigned code point. Since 1.14
	ScriptUnknown Script = 61
	// ScriptBalinese: balinese. Since 1.14
	ScriptBalinese Script = 62
	// ScriptCuneiform: cuneiform. Since 1.14
	ScriptCuneiform Script = 63
	// ScriptPhoenician: phoenician. Since 1.14
	ScriptPhoenician Script = 64
	// ScriptPhagsPa: phags-pa. Since 1.14
	ScriptPhagsPa Script = 65
	// ScriptNko: n'Ko. Since 1.14
	ScriptNko Script = 66
	// ScriptKayahLi: kayah Li. Since 1.20.1
	ScriptKayahLi Script = 67
	// ScriptLepcha: lepcha. Since 1.20.1
	ScriptLepcha Script = 68
	// ScriptRejang: rejang. Since 1.20.1
	ScriptRejang Script = 69
	// ScriptSundanese: sundanese. Since 1.20.1
	ScriptSundanese Script = 70
	// ScriptSaurashtra: saurashtra. Since 1.20.1
	ScriptSaurashtra Script = 71
	// ScriptCham: cham. Since 1.20.1
	ScriptCham Script = 72
	// ScriptOlChiki: ol Chiki. Since 1.20.1
	ScriptOlChiki Script = 73
	// ScriptVai: vai. Since 1.20.1
	ScriptVai Script = 74
	// ScriptCarian: carian. Since 1.20.1
	ScriptCarian Script = 75
	// ScriptLycian: lycian. Since 1.20.1
	ScriptLycian Script = 76
	// ScriptLydian: lydian. Since 1.20.1
	ScriptLydian Script = 77
	// ScriptBatak: batak. Since 1.32
	ScriptBatak Script = 78
	// ScriptBrahmi: brahmi. Since 1.32
	ScriptBrahmi Script = 79
	// ScriptMandaic: mandaic. Since 1.32
	ScriptMandaic Script = 80
	// ScriptChakma: chakma. Since: 1.32
	ScriptChakma Script = 81
	// ScriptMeroiticCursive: meroitic Cursive. Since: 1.32
	ScriptMeroiticCursive Script = 82
	// ScriptMeroiticHieroglyphs: meroitic Hieroglyphs. Since: 1.32
	ScriptMeroiticHieroglyphs Script = 83
	// ScriptMiao: miao. Since: 1.32
	ScriptMiao Script = 84
	// ScriptSharada: sharada. Since: 1.32
	ScriptSharada Script = 85
	// ScriptSoraSompeng: sora Sompeng. Since: 1.32
	ScriptSoraSompeng Script = 86
	// ScriptTakri: takri. Since: 1.32
	ScriptTakri Script = 87
	// ScriptBassaVah: bassa. Since: 1.40
	ScriptBassaVah Script = 88
	// ScriptCaucasianAlbanian: caucasian Albanian. Since: 1.40
	ScriptCaucasianAlbanian Script = 89
	// ScriptDuployan: duployan. Since: 1.40
	ScriptDuployan Script = 90
	// ScriptElbasan: elbasan. Since: 1.40
	ScriptElbasan Script = 91
	// ScriptGrantha: grantha. Since: 1.40
	ScriptGrantha Script = 92
	// ScriptKhojki: kjohki. Since: 1.40
	ScriptKhojki Script = 93
	// ScriptKhudawadi: khudawadi, Sindhi. Since: 1.40
	ScriptKhudawadi Script = 94
	// ScriptLinearA: linear A. Since: 1.40
	ScriptLinearA Script = 95
	// ScriptMahajani: mahajani. Since: 1.40
	ScriptMahajani Script = 96
	// ScriptManichaean: manichaean. Since: 1.40
	ScriptManichaean Script = 97
	// ScriptMendeKikakui: mende Kikakui. Since: 1.40
	ScriptMendeKikakui Script = 98
	// ScriptModi: modi. Since: 1.40
	ScriptModi Script = 99
	// ScriptMro: mro. Since: 1.40
	ScriptMro Script = 100
	// ScriptNabataean: nabataean. Since: 1.40
	ScriptNabataean Script = 101
	// ScriptOldNorthArabian: old North Arabian. Since: 1.40
	ScriptOldNorthArabian Script = 102
	// ScriptOldPermic: old Permic. Since: 1.40
	ScriptOldPermic Script = 103
	// ScriptPahawhHmong: pahawh Hmong. Since: 1.40
	ScriptPahawhHmong Script = 104
	// ScriptPalmyrene: palmyrene. Since: 1.40
	ScriptPalmyrene Script = 105
	// ScriptPauCinHau: pau Cin Hau. Since: 1.40
	ScriptPauCinHau Script = 106
	// ScriptPsalterPahlavi: psalter Pahlavi. Since: 1.40
	ScriptPsalterPahlavi Script = 107
	// ScriptSiddham: siddham. Since: 1.40
	ScriptSiddham Script = 108
	// ScriptTirhuta: tirhuta. Since: 1.40
	ScriptTirhuta Script = 109
	// ScriptWarangCiti: warang Citi. Since: 1.40
	ScriptWarangCiti Script = 110
	// ScriptAhom: ahom. Since: 1.40
	ScriptAhom Script = 111
	// ScriptAnatolianHieroglyphs: anatolian Hieroglyphs. Since: 1.40
	ScriptAnatolianHieroglyphs Script = 112
	// ScriptHatran: hatran. Since: 1.40
	ScriptHatran Script = 113
	// ScriptMultani: multani. Since: 1.40
	ScriptMultani Script = 114
	// ScriptOldHungarian: old Hungarian. Since: 1.40
	ScriptOldHungarian Script = 115
	// ScriptSignwriting: signwriting. Since: 1.40
	ScriptSignwriting Script = 116
)

func marshalScript(p uintptr) (interface{}, error) {
	return Script(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Stretch: an enumeration specifying the width of the font relative to other
// designs within a family.
type Stretch int

const (
	// StretchUltraCondensed: ultra condensed width
	StretchUltraCondensed Stretch = 0
	// StretchExtraCondensed: extra condensed width
	StretchExtraCondensed Stretch = 1
	// StretchCondensed: condensed width
	StretchCondensed Stretch = 2
	// StretchSemiCondensed: semi condensed width
	StretchSemiCondensed Stretch = 3
	// StretchNormal: the normal width
	StretchNormal Stretch = 4
	// StretchSemiExpanded: semi expanded width
	StretchSemiExpanded Stretch = 5
	// StretchExpanded: expanded width
	StretchExpanded Stretch = 6
	// StretchExtraExpanded: extra expanded width
	StretchExtraExpanded Stretch = 7
	// StretchUltraExpanded: ultra expanded width
	StretchUltraExpanded Stretch = 8
)

func marshalStretch(p uintptr) (interface{}, error) {
	return Stretch(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Style: an enumeration specifying the various slant styles possible for a
// font.
type Style int

const (
	// StyleNormal: the font is upright.
	StyleNormal Style = 0
	// StyleOblique: the font is slanted, but in a roman style.
	StyleOblique Style = 1
	// StyleItalic: the font is slanted in an italic style.
	StyleItalic Style = 2
)

func marshalStyle(p uintptr) (interface{}, error) {
	return Style(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TabAlign: `PangoTabAlign` specifies where a tab stop appears relative to the
// text.
type TabAlign int

const (
	// TabAlignLeft: the tab stop appears to the left of the text.
	TabAlignLeft TabAlign = 0
)

func marshalTabAlign(p uintptr) (interface{}, error) {
	return TabAlign(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Underline: the `PangoUnderline` enumeration is used to specify whether text
// should be underlined, and if so, the type of underlining.
type Underline int

const (
	// UnderlineNone: no underline should be drawn
	UnderlineNone Underline = 0
	// UnderlineSingle: a single underline should be drawn
	UnderlineSingle Underline = 1
	// UnderlineDouble: a double underline should be drawn
	UnderlineDouble Underline = 2
	// UnderlineLow: a single underline should be drawn at a position beneath
	// the ink extents of the text being underlined. This should be used only
	// for underlining single characters, such as for keyboard accelerators.
	// PANGO_UNDERLINE_SINGLE should be used for extended portions of text.
	UnderlineLow Underline = 3
	// UnderlineError: a wavy underline should be drawn below. This underline is
	// typically used to indicate an error such as a possible mispelling; in
	// some cases a contrasting color may automatically be used. This type of
	// underlining is available since Pango 1.4.
	UnderlineError Underline = 4
	// UnderlineSingleLine: like @PANGO_UNDERLINE_SINGLE, but drawn continuously
	// across multiple runs. This type of underlining is available since Pango
	// 1.46.
	UnderlineSingleLine Underline = 5
	// UnderlineDoubleLine: like @PANGO_UNDERLINE_DOUBLE, but drawn continuously
	// across multiple runs. This type of underlining is available since Pango
	// 1.46.
	UnderlineDoubleLine Underline = 6
	// UnderlineErrorLine: like @PANGO_UNDERLINE_ERROR, but drawn continuously
	// across multiple runs. This type of underlining is available since Pango
	// 1.46.
	UnderlineErrorLine Underline = 7
)

func marshalUnderline(p uintptr) (interface{}, error) {
	return Underline(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Variant: an enumeration specifying capitalization variant of the font.
type Variant int

const (
	// VariantNormal: a normal font.
	VariantNormal Variant = 0
	// VariantSmallCaps: a font with the lower case characters replaced by
	// smaller variants of the capital characters.
	VariantSmallCaps Variant = 1
)

func marshalVariant(p uintptr) (interface{}, error) {
	return Variant(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Weight: an enumeration specifying the weight (boldness) of a font.
//
// This is a numerical value ranging from 100 to 1000, but there are some
// predefined values.
type Weight int

const (
	// WeightThin: the thin weight (= 100; Since: 1.24)
	WeightThin Weight = 100
	// WeightUltralight: the ultralight weight (= 200)
	WeightUltralight Weight = 200
	// WeightLight: the light weight (= 300)
	WeightLight Weight = 300
	// WeightSemilight: the semilight weight (= 350; Since: 1.36.7)
	WeightSemilight Weight = 350
	// WeightBook: the book weight (= 380; Since: 1.24)
	WeightBook Weight = 380
	// WeightNormal: the default weight (= 400)
	WeightNormal Weight = 400
	// WeightMedium: the normal weight (= 500; Since: 1.24)
	WeightMedium Weight = 500
	// WeightSemibold: the semibold weight (= 600)
	WeightSemibold Weight = 600
	// WeightBold: the bold weight (= 700)
	WeightBold Weight = 700
	// WeightUltrabold: the ultrabold weight (= 800)
	WeightUltrabold Weight = 800
	// WeightHeavy: the heavy weight (= 900)
	WeightHeavy Weight = 900
	// WeightUltraheavy: the ultraheavy weight (= 1000; Since: 1.24)
	WeightUltraheavy Weight = 1000
)

func marshalWeight(p uintptr) (interface{}, error) {
	return Weight(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// WrapMode: `PangoWrapMode` describes how to wrap the lines of a `PangoLayout`
// to the desired width.
type WrapMode int

const (
	// WrapModeWord: wrap lines at word boundaries.
	WrapModeWord WrapMode = 0
	// WrapModeChar: wrap lines at character boundaries.
	WrapModeChar WrapMode = 1
	// WrapModeWordChar: wrap lines at word boundaries, but fall back to
	// character boundaries if there is not enough space for a full word.
	WrapModeWordChar WrapMode = 2
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	return WrapMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FontMask: the bits in a `PangoFontMask` correspond to the set fields in a
// `PangoFontDescription`.
type FontMask int

const (
	// FontMaskFamily: the font family is specified.
	FontMaskFamily FontMask = 0b1
	// FontMaskStyle: the font style is specified.
	FontMaskStyle FontMask = 0b10
	// FontMaskVariant: the font variant is specified.
	FontMaskVariant FontMask = 0b100
	// FontMaskWeight: the font weight is specified.
	FontMaskWeight FontMask = 0b1000
	// FontMaskStretch: the font stretch is specified.
	FontMaskStretch FontMask = 0b10000
	// FontMaskSize: the font size is specified.
	FontMaskSize FontMask = 0b100000
	// FontMaskGravity: the font gravity is specified (Since: 1.16.)
	FontMaskGravity FontMask = 0b1000000
	// FontMaskVariations: openType font variations are specified (Since: 1.42)
	FontMaskVariations FontMask = 0b10000000
)

func marshalFontMask(p uintptr) (interface{}, error) {
	return FontMask(C.g_value_get_bitfield((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShapeFlags flags influencing the shaping process.
//
// `PangoShapeFlags` can be passed to pango_shape_with_flags().
type ShapeFlags int

const (
	// ShapeFlagsNone: default value.
	ShapeFlagsNone ShapeFlags = 0b0
	// ShapeFlagsRoundPositions: round glyph positions and widths to whole
	// device units. This option should be set if the target renderer can't do
	// subpixel positioning of glyphs.
	ShapeFlagsRoundPositions ShapeFlags = 0b1
)

func marshalShapeFlags(p uintptr) (interface{}, error) {
	return ShapeFlags(C.g_value_get_bitfield((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShowFlags: these flags affect how Pango treats characters that are normally
// not visible in the output.
type ShowFlags int

const (
	// ShowFlagsNone: no special treatment for invisible characters
	ShowFlagsNone ShowFlags = 0b0
	// ShowFlagsSpaces: render spaces, tabs and newlines visibly
	ShowFlagsSpaces ShowFlags = 0b1
	// ShowFlagsLineBreaks: render line breaks visibly
	ShowFlagsLineBreaks ShowFlags = 0b10
	// ShowFlagsIgnorables: render default-ignorable Unicode characters visibly
	ShowFlagsIgnorables ShowFlags = 0b100
)

func marshalShowFlags(p uintptr) (interface{}, error) {
	return ShowFlags(C.g_value_get_bitfield((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Coverage: a Coverage structure is a map from Unicode characters to
// CoverageLevel values.
//
// It is often necessary in Pango to determine if a particular font can
// represent a particular character, and also how well it can represent that
// character. The Coverage is a data structure that is used to represent that
// information. It is an opaque structure with no public fields.
type Coverage interface {
	gextras.Objector

	// Copy: copy an existing `PangoCoverage`.
	Copy() Coverage
	// Get: determine whether a particular index is covered by @coverage.
	Get(index_ int) CoverageLevel
	// Max: set the coverage for each index in @coverage to be the max (better)
	// value of the current coverage for the index and the coverage for the
	// corresponding index in @other.
	Max(other Coverage)
	// Ref: increase the reference count on the `PangoCoverage` by one.
	Ref() Coverage
	// Set: modify a particular index within @coverage
	Set(index_ int, level CoverageLevel)
	// ToBytes: convert a `PangoCoverage` structure into a flat binary format.
	ToBytes() (bytes []byte, nBytes int)
	// Unref: decrease the reference count on the `PangoCoverage` by one.
	//
	// If the result is zero, free the coverage and all associated memory.
	Unref()
}

// coverage implements the Coverage interface.
type coverage struct {
	gextras.Objector
}

var _ Coverage = (*coverage)(nil)

// WrapCoverage wraps a GObject to the right type. It is
// primarily used internally.
func WrapCoverage(obj *externglib.Object) Coverage {
	return Coverage{
		Objector: obj,
	}
}

func marshalCoverage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCoverage(obj), nil
}

// NewCoverage constructs a class Coverage.
func NewCoverage() Coverage {
	var cret C.PangoCoverage
	var goret1 Coverage

	cret = C.pango_coverage_new()

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Coverage)

	return goret1
}

// Copy: copy an existing `PangoCoverage`.
func (c coverage) Copy() Coverage {
	var arg0 *C.PangoCoverage

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	var cret *C.PangoCoverage
	var goret1 Coverage

	cret = C.pango_coverage_copy(arg0)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Coverage)

	return goret1
}

// Get: determine whether a particular index is covered by @coverage.
func (c coverage) Get(index_ int) CoverageLevel {
	var arg0 *C.PangoCoverage
	var arg1 C.int

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	arg1 = C.int(index_)

	var cret C.PangoCoverageLevel
	var goret1 CoverageLevel

	cret = C.pango_coverage_get(arg0, index_)

	goret1 = CoverageLevel(cret)

	return goret1
}

// Max: set the coverage for each index in @coverage to be the max (better)
// value of the current coverage for the index and the coverage for the
// corresponding index in @other.
func (c coverage) Max(other Coverage) {
	var arg0 *C.PangoCoverage
	var arg1 *C.PangoCoverage

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoCoverage)(unsafe.Pointer(other.Native()))

	C.pango_coverage_max(arg0, other)
}

// Ref: increase the reference count on the `PangoCoverage` by one.
func (c coverage) Ref() Coverage {
	var arg0 *C.PangoCoverage

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	var cret *C.PangoCoverage
	var goret1 Coverage

	cret = C.pango_coverage_ref(arg0)

	goret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Coverage)

	return goret1
}

// Set: modify a particular index within @coverage
func (c coverage) Set(index_ int, level CoverageLevel) {
	var arg0 *C.PangoCoverage
	var arg1 C.int
	var arg2 C.PangoCoverageLevel

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))
	arg1 = C.int(index_)
	arg2 = (C.PangoCoverageLevel)(level)

	C.pango_coverage_set(arg0, index_, level)
}

// ToBytes: convert a `PangoCoverage` structure into a flat binary format.
func (c coverage) ToBytes() (bytes []byte, nBytes int) {
	var arg0 *C.PangoCoverage

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	C.pango_coverage_to_bytes(arg0, &arg1, &arg2)

	return ret1, ret2
}

// Unref: decrease the reference count on the `PangoCoverage` by one.
//
// If the result is zero, free the coverage and all associated memory.
func (c coverage) Unref() {
	var arg0 *C.PangoCoverage

	arg0 = (*C.PangoCoverage)(unsafe.Pointer(c.Native()))

	C.pango_coverage_unref(arg0)
}
