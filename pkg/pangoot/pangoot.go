// Code generated by girgen. DO NOT EDIT.

package pangoot

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoot
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango-ot.h>
//
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		// Enums
		// Skipped TableType.

		// Records
		{T: externglib.Type(C.pango_ot_buffer_get_type()), F: marshalBuffer},
		// Skipped FeatureMap.
		// Skipped Glyph.
		{T: externglib.Type(C.pango_ot_ruleset_description_get_type()), F: marshalRulesetDescription},

		// Classes
		{T: externglib.Type(C.pango_ot_info_get_type()), F: marshalInfo},
		{T: externglib.Type(C.pango_ot_ruleset_get_type()), F: marshalRuleset},
	})
}

// Tag: the OTTag typedef is used to represent TrueType and OpenType four letter
// tags inside Pango. Use PANGO_OT_TAG_MAKE() or PANGO_OT_TAG_MAKE_FROM_STRING()
// macros to create PangoOTTags manually.
type Tag uint32

// TableType: the PangoOTTableType enumeration values are used to identify the
// various OpenType tables in the pango_ot_info_… functions.
type TableType int

const (
	// TableTypeGsub: the GSUB table.
	TableTypeGsub TableType = 0
	// TableTypeGpos: the GPOS table.
	TableTypeGpos TableType = 1
)

// TagFromLanguage finds the OpenType language-system tag best describing
// @language.
func TagFromLanguage(language *pango.Language) Tag {
	var arg0 *C.PangoLanguage
	arg0 = (*C.C.PangoLanguage)(language.Native())

	ret := C.pango_ot_tag_from_language(arg0)

	var ret0 Tag
	{
		var tmp uint32
		tmp = uint32(ret)
		ret0 = Tag(tmp)
	}

	return ret0
}

// TagFromScript finds the OpenType script tag corresponding to @script.
//
// The PANGO_SCRIPT_COMMON, PANGO_SCRIPT_INHERITED, and PANGO_SCRIPT_UNKNOWN
// scripts are mapped to the OpenType 'DFLT' script tag that is also defined as
// PANGO_OT_TAG_DEFAULT_SCRIPT.
//
// Note that multiple Script values may map to the same OpenType script tag. In
// particular, PANGO_SCRIPT_HIRAGANA and PANGO_SCRIPT_KATAKANA both map to the
// OT tag 'kana'.
func TagFromScript(script pango.Script) Tag {
	var arg0 C.PangoScript
	arg0 = (C.C.PangoScript)(script)

	ret := C.pango_ot_tag_from_script(arg0)

	var ret0 Tag
	{
		var tmp uint32
		tmp = uint32(ret)
		ret0 = Tag(tmp)
	}

	return ret0
}

// TagToLanguage finds a Language corresponding to @language_tag.
func TagToLanguage(languageTag Tag) *pango.Language {
	ret := C.pango_ot_tag_to_language()

	var ret0 *pango.Language
	ret0 = pango.WrapLanguage(ret)

	return ret0
}

// TagToScript finds the Script corresponding to @script_tag.
//
// The 'DFLT' script tag is mapped to PANGO_SCRIPT_COMMON.
//
// Note that an OpenType script tag may correspond to multiple Script values. In
// such cases, the Script value with the smallest value is returned. In
// particular, PANGO_SCRIPT_HIRAGANA and PANGO_SCRIPT_KATAKANA both map to the
// OT tag 'kana'. This function will return PANGO_SCRIPT_HIRAGANA for 'kana'.
func TagToScript(scriptTag Tag) pango.Script {
	ret := C.pango_ot_tag_to_script()

	var ret0 pango.Script
	ret0 = pango.Script(ret)

	return ret0
}

type Buffer struct {
	native *C.PangoOTBuffer
}

// WrapBuffer wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBuffer(ptr unsafe.Pointer) *Buffer {
	p := (*C.PangoOTBuffer)(ptr)
	v := Buffer{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Buffer).free)

	return &v
}

func marshalBuffer(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapBuffer(unsafe.Pointer(b))
}

func (b *Buffer) free() {
	C.free(b.Native())
}

// Native returns the underlying source pointer.
func (b *Buffer) Native() unsafe.Pointer {
	return unsafe.Pointer(b.native)
}

// Native returns the pointer to *C.PangoOTBuffer. The caller is expected to
// cast.
func (b *Buffer) Native() unsafe.Pointer {
	return unsafe.Pointer(b.native)
}

func NewBuffer(font pangofc.Font) *Buffer

// FeatureMap: the OTFeatureMap typedef is used to represent an OpenType feature
// with the property bit associated with it. The feature tag is represented as a
// char array instead of a OTTag for convenience.
type FeatureMap struct {
	// FeatureName: feature tag in represented as four-letter ASCII string.
	FeatureName [5]byte
	// PropertyBit: the property bit to use for this feature. See
	// pango_ot_ruleset_add_feature() for details.
	PropertyBit uint32

	native *C.PangoOTFeatureMap
}

// WrapFeatureMap wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFeatureMap(ptr unsafe.Pointer) *FeatureMap {
	p := (*C.PangoOTFeatureMap)(ptr)
	v := FeatureMap{native: p}

	v.FeatureName = [5]byte(p.feature_name)
	v.PropertyBit = uint32(p.property_bit)

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*FeatureMap).free)

	return &v
}

func marshalFeatureMap(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFeatureMap(unsafe.Pointer(b))
}

func (f *FeatureMap) free() {
	C.free(f.Native())
}

// Native returns the underlying source pointer.
func (f *FeatureMap) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

// Native returns the pointer to *C.PangoOTFeatureMap. The caller is expected to
// cast.
func (f *FeatureMap) Native() unsafe.Pointer {
	return unsafe.Pointer(f.native)
}

// Glyph: the OTGlyph structure represents a single glyph together with
// information used for OpenType layout processing of the glyph. It contains the
// following fields.
type Glyph struct {
	// Glyph: the glyph itself.
	Glyph uint32
	// Properties: the properties value, identifying which features should be
	// applied on this glyph. See pango_ot_ruleset_add_feature().
	Properties uint
	// Cluster: the cluster that this glyph belongs to.
	Cluster uint
	// Component: a component value, set by the OpenType layout engine.
	Component uint16
	// LigID: a ligature index value, set by the OpenType layout engine.
	LigID uint16
	// Internal: for Pango internal use
	Internal uint

	native *C.PangoOTGlyph
}

// WrapGlyph wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapGlyph(ptr unsafe.Pointer) *Glyph {
	p := (*C.PangoOTGlyph)(ptr)
	v := Glyph{native: p}

	v.Glyph = uint32(p.glyph)
	v.Properties = uint(p.properties)
	v.Cluster = uint(p.cluster)
	v.Component = uint16(p.component)
	v.LigID = uint16(p.ligID)
	v.Internal = uint(p.internal)

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Glyph).free)

	return &v
}

func marshalGlyph(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapGlyph(unsafe.Pointer(b))
}

func (g *Glyph) free() {
	C.free(g.Native())
}

// Native returns the underlying source pointer.
func (g *Glyph) Native() unsafe.Pointer {
	return unsafe.Pointer(g.native)
}

// Native returns the pointer to *C.PangoOTGlyph. The caller is expected to
// cast.
func (g *Glyph) Native() unsafe.Pointer {
	return unsafe.Pointer(g.native)
}

// RulesetDescription: the OTRuleset structure holds all the information needed
// to build a complete OTRuleset from an OpenType font. The main use of this
// struct is to act as the key for a per-font hash of rulesets. The user
// populates a ruleset description and gets the ruleset using
// pango_ot_ruleset_get_for_description() or create a new one using
// pango_ot_ruleset_new_from_description().
type RulesetDescription struct {
	// Script: a Script.
	Script pango.Script
	// Language: a Language.
	Language *pango.Language
	// StaticGsubFeatures: static map of GSUB features, or nil.
	StaticGsubFeatures *FeatureMap
	// NStaticGsubFeatures: length of @static_gsub_features, or 0.
	NStaticGsubFeatures uint
	// StaticGposFeatures: static map of GPOS features, or nil.
	StaticGposFeatures *FeatureMap
	// NStaticGposFeatures: length of @static_gpos_features, or 0.
	NStaticGposFeatures uint
	// OtherFeatures: map of extra features to add to both GSUB and GPOS, or
	// nil. Unlike the static maps, this pointer need not live beyond the life
	// of function calls taking this struct.
	OtherFeatures *FeatureMap
	// NOtherFeatures: length of @other_features, or 0.
	NOtherFeatures uint

	native *C.PangoOTRulesetDescription
}

// WrapRulesetDescription wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRulesetDescription(ptr unsafe.Pointer) *RulesetDescription {
	p := (*C.PangoOTRulesetDescription)(ptr)
	v := RulesetDescription{native: p}

	v.Script = pango.Script(p.script)
	v.Language = pango.WrapLanguage(p.language)
	v.StaticGsubFeatures = WrapFeatureMap(p.static_gsub_features)
	v.NStaticGsubFeatures = uint(p.n_static_gsub_features)
	v.StaticGposFeatures = WrapFeatureMap(p.static_gpos_features)
	v.NStaticGposFeatures = uint(p.n_static_gpos_features)
	v.OtherFeatures = WrapFeatureMap(p.other_features)
	v.NOtherFeatures = uint(p.n_other_features)

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*RulesetDescription).free)

	return &v
}

func marshalRulesetDescription(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRulesetDescription(unsafe.Pointer(b))
}

func (r *RulesetDescription) free() {
	C.free(r.Native())
}

// Native returns the underlying source pointer.
func (r *RulesetDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

// Native returns the pointer to *C.PangoOTRulesetDescription. The caller is expected to
// cast.
func (r *RulesetDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(r.native)
}

type Info interface {
	gextras.Objector

	// FindFeature finds the index of a feature. If the feature is not found,
	// sets @feature_index to PANGO_OT_NO_FEATURE, which is safe to pass to
	// pango_ot_ruleset_add_feature() and similar functions.
	//
	// In the future, this may set @feature_index to an special value that if
	// used in pango_ot_ruleset_add_feature() will ask Pango to synthesize the
	// requested feature based on Unicode properties and data. However, this
	// function will still return false in those cases. So, users may want to
	// ignore the return value of this function in certain cases.
	FindFeature(tableType TableType, featureTag Tag, scriptIndex uint, languageIndex uint) (featureIndex uint, ok bool)
	// FindLanguage finds the index of a language and its required feature
	// index. If the language is not found, sets @language_index to
	// PANGO_OT_DEFAULT_LANGUAGE and the required feature of the default
	// language system is returned in required_feature_index. For best
	// compatibility with some fonts, also searches the language system tag
	// 'dflt' before falling back to the default language system, but that is
	// transparent to the user. The user can simply ignore the return value of
	// this function to automatically fall back to the default language system.
	FindLanguage(tableType TableType, scriptIndex uint, languageTag Tag) (languageIndex uint, requiredFeatureIndex uint, ok bool)
	// FindScript finds the index of a script. If not found, tries to find the
	// 'DFLT' and then 'dflt' scripts and return the index of that in
	// @script_index. If none of those is found either, PANGO_OT_NO_SCRIPT is
	// placed in @script_index.
	//
	// All other functions taking an input script_index parameter know how to
	// handle PANGO_OT_NO_SCRIPT, so one can ignore the return value of this
	// function completely and proceed, to enjoy the automatic fallback to the
	// 'DFLT'/'dflt' script.
	FindScript(tableType TableType, scriptTag Tag) (scriptIndex uint, ok bool)
	// ListFeatures obtains the list of features for the given language of the
	// given script.
	ListFeatures(tableType TableType, tag Tag, scriptIndex uint, languageIndex uint) *Tag
	// ListLanguages obtains the list of available languages for a given script.
	ListLanguages(tableType TableType, scriptIndex uint, languageTag Tag) *Tag
	// ListScripts obtains the list of available scripts.
	ListScripts(tableType TableType) *Tag
}

type info struct {
	*externglib.Object
}

// WrapInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapInfo(obj *externglib.Object) Info {
	return info{*externglib.Object{obj}}
}

func marshalInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInfo(obj), nil
}

func (i info) FindFeature(tableType TableType, featureTag Tag, scriptIndex uint, languageIndex uint) (featureIndex uint, ok bool)

func (i info) FindLanguage(tableType TableType, scriptIndex uint, languageTag Tag) (languageIndex uint, requiredFeatureIndex uint, ok bool)

func (i info) FindScript(tableType TableType, scriptTag Tag) (scriptIndex uint, ok bool)

func (i info) ListFeatures(tableType TableType, tag Tag, scriptIndex uint, languageIndex uint) *Tag

func (i info) ListLanguages(tableType TableType, scriptIndex uint, languageTag Tag) *Tag

func (i info) ListScripts(tableType TableType) *Tag

// Ruleset: the OTRuleset structure holds a set of features selected from the
// tables in an OpenType font. (A feature is an operation such as adjusting
// glyph positioning that should be applied to a text feature such as a certain
// type of accent.) A OTRuleset is created with pango_ot_ruleset_new(), features
// are added to it with pango_ot_ruleset_add_feature(), then it is applied to a
// GlyphString with pango_ot_ruleset_shape().
type Ruleset interface {
	gextras.Objector

	// AddFeature adds a feature to the ruleset.
	AddFeature(tableType TableType, featureIndex uint, propertyBit uint32)
	// FeatureCount gets the number of GSUB and GPOS features in the ruleset.
	FeatureCount() (nGsubFeatures uint, nGposFeatures uint, guint uint)
	// MaybeAddFeature: this is a convenience function that first tries to find
	// the feature using pango_ot_info_find_feature() and the ruleset script and
	// language passed to pango_ot_ruleset_new_for(), and if the feature is
	// found, adds it to the ruleset.
	//
	// If @ruleset was not created using pango_ot_ruleset_new_for(), this
	// function does nothing.
	MaybeAddFeature(tableType TableType, featureTag Tag, propertyBit uint32) bool
	// MaybeAddFeatures: this is a convenience function that for each feature in
	// the feature map array @features converts the feature name to a OTTag
	// feature tag using PANGO_OT_TAG_MAKE() and calls
	// pango_ot_ruleset_maybe_add_feature() on it.
	MaybeAddFeatures(tableType TableType, features *FeatureMap, nFeatures uint) uint
	// Position performs the OpenType GPOS positioning on @buffer using the
	// features in @ruleset
	Position(buffer *Buffer)
	// Substitute performs the OpenType GSUB substitution on @buffer using the
	// features in @ruleset
	Substitute(buffer *Buffer)
}

type ruleset struct {
	*externglib.Object
}

// WrapRuleset wraps a GObject to the right type. It is
// primarily used internally.
func WrapRuleset(obj *externglib.Object) Ruleset {
	return ruleset{*externglib.Object{obj}}
}

func marshalRuleset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRuleset(obj), nil
}

func NewRuleset(info Info) Ruleset

func NewRulesetFor(info Info, script pango.Script, language *pango.Language) Ruleset

func NewRulesetFromDescription(info Info, desc *RulesetDescription) Ruleset

func (r ruleset) AddFeature(tableType TableType, featureIndex uint, propertyBit uint32)

func (r ruleset) FeatureCount() (nGsubFeatures uint, nGposFeatures uint, guint uint)

func (r ruleset) MaybeAddFeature(tableType TableType, featureTag Tag, propertyBit uint32) bool

func (r ruleset) MaybeAddFeatures(tableType TableType, features *FeatureMap, nFeatures uint) uint

func (r ruleset) Position(buffer *Buffer)

func (r ruleset) Substitute(buffer *Buffer)
