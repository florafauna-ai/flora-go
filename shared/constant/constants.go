// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/florafauna-ai/flora-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type AddShapeToImageBrowser string     // Always "add-shape-to-image-browser"
type AddTextToImageBrowser string      // Always "add-text-to-image-browser"
type BlurImageBrowser string           // Always "blur-image-browser"
type BoomerangVideo string             // Always "boomerang-video"
type ChangeImageArBrowser string       // Always "change-image-ar-browser"
type ChangeVideoAr string              // Always "change-video-ar"
type ColorFilterImageBrowser string    // Always "color-filter-image-browser"
type ColorFilterVideo string           // Always "color-filter-video"
type ColorGradeImageBrowser string     // Always "color-grade-image-browser"
type ColorGradeVideo string            // Always "color-grade-video"
type ColorTintImageBrowser string      // Always "color-tint-image-browser"
type ConcatTextBrowser string          // Always "concat-text-browser"
type CropImageBrowser string           // Always "crop-image-browser"
type DrawImageBrowser string           // Always "draw-image-browser"
type DuplicateImageBrowser string      // Always "duplicate-image-browser"
type DuplicateVideo string             // Always "duplicate-video"
type ExtractVideoFrames string         // Always "extract-video-frames"
type File string                       // Always "file"
type FilterColorImageBrowser string    // Always "filter-color-image-browser"
type FindAndReplaceTextBrowser string  // Always "find-and-replace-text-browser"
type GreenscreenVideo string           // Always "greenscreen-video"
type KenBurnsVideo string              // Always "ken-burns-video"
type MergeAudioIntoVideo string        // Always "merge-audio-into-video"
type MultipartFormData string          // Always "multipart/form-data"
type OverlayImageBrowser string        // Always "overlay-image-browser"
type Post string                       // Always "POST"
type QrCodeGeneratorBrowser string     // Always "qr-code-generator-browser"
type ResizeImageBrowser string         // Always "resize-image-browser"
type ResizeVideo string                // Always "resize-video"
type ReverseVideo string               // Always "reverse-video"
type RotateImageBrowser string         // Always "rotate-image-browser"
type Scene3dImageBrowser string        // Always "scene-3d-image-browser"
type ShaderEffectBrowser string        // Always "shader-effect-browser"
type SideBySideCompositeBrowser string // Always "side-by-side-composite-browser"
type SlowDownVideo string              // Always "slow-down-video"
type SpeedUpVideo string               // Always "speed-up-video"
type SplitAudioFromVideo string        // Always "split-audio-from-video"
type SplitTextBrowser string           // Always "split-text-browser"
type SplitVideo string                 // Always "split-video"
type StitchVideos string               // Always "stitch-videos"
type VideoEffect string                // Always "video-effect"
type VideoToFrameGrid string           // Always "video-to-frame-grid"
type VideoToLongExposure string        // Always "video-to-long-exposure"
type Workspace string                  // Always "workspace"

func (c AddShapeToImageBrowser) Default() AddShapeToImageBrowser { return "add-shape-to-image-browser" }
func (c AddTextToImageBrowser) Default() AddTextToImageBrowser   { return "add-text-to-image-browser" }
func (c BlurImageBrowser) Default() BlurImageBrowser             { return "blur-image-browser" }
func (c BoomerangVideo) Default() BoomerangVideo                 { return "boomerang-video" }
func (c ChangeImageArBrowser) Default() ChangeImageArBrowser     { return "change-image-ar-browser" }
func (c ChangeVideoAr) Default() ChangeVideoAr                   { return "change-video-ar" }
func (c ColorFilterImageBrowser) Default() ColorFilterImageBrowser {
	return "color-filter-image-browser"
}
func (c ColorFilterVideo) Default() ColorFilterVideo             { return "color-filter-video" }
func (c ColorGradeImageBrowser) Default() ColorGradeImageBrowser { return "color-grade-image-browser" }
func (c ColorGradeVideo) Default() ColorGradeVideo               { return "color-grade-video" }
func (c ColorTintImageBrowser) Default() ColorTintImageBrowser   { return "color-tint-image-browser" }
func (c ConcatTextBrowser) Default() ConcatTextBrowser           { return "concat-text-browser" }
func (c CropImageBrowser) Default() CropImageBrowser             { return "crop-image-browser" }
func (c DrawImageBrowser) Default() DrawImageBrowser             { return "draw-image-browser" }
func (c DuplicateImageBrowser) Default() DuplicateImageBrowser   { return "duplicate-image-browser" }
func (c DuplicateVideo) Default() DuplicateVideo                 { return "duplicate-video" }
func (c ExtractVideoFrames) Default() ExtractVideoFrames         { return "extract-video-frames" }
func (c File) Default() File                                     { return "file" }
func (c FilterColorImageBrowser) Default() FilterColorImageBrowser {
	return "filter-color-image-browser"
}
func (c FindAndReplaceTextBrowser) Default() FindAndReplaceTextBrowser {
	return "find-and-replace-text-browser"
}
func (c GreenscreenVideo) Default() GreenscreenVideo             { return "greenscreen-video" }
func (c KenBurnsVideo) Default() KenBurnsVideo                   { return "ken-burns-video" }
func (c MergeAudioIntoVideo) Default() MergeAudioIntoVideo       { return "merge-audio-into-video" }
func (c MultipartFormData) Default() MultipartFormData           { return "multipart/form-data" }
func (c OverlayImageBrowser) Default() OverlayImageBrowser       { return "overlay-image-browser" }
func (c Post) Default() Post                                     { return "POST" }
func (c QrCodeGeneratorBrowser) Default() QrCodeGeneratorBrowser { return "qr-code-generator-browser" }
func (c ResizeImageBrowser) Default() ResizeImageBrowser         { return "resize-image-browser" }
func (c ResizeVideo) Default() ResizeVideo                       { return "resize-video" }
func (c ReverseVideo) Default() ReverseVideo                     { return "reverse-video" }
func (c RotateImageBrowser) Default() RotateImageBrowser         { return "rotate-image-browser" }
func (c Scene3dImageBrowser) Default() Scene3dImageBrowser       { return "scene-3d-image-browser" }
func (c ShaderEffectBrowser) Default() ShaderEffectBrowser       { return "shader-effect-browser" }
func (c SideBySideCompositeBrowser) Default() SideBySideCompositeBrowser {
	return "side-by-side-composite-browser"
}
func (c SlowDownVideo) Default() SlowDownVideo             { return "slow-down-video" }
func (c SpeedUpVideo) Default() SpeedUpVideo               { return "speed-up-video" }
func (c SplitAudioFromVideo) Default() SplitAudioFromVideo { return "split-audio-from-video" }
func (c SplitTextBrowser) Default() SplitTextBrowser       { return "split-text-browser" }
func (c SplitVideo) Default() SplitVideo                   { return "split-video" }
func (c StitchVideos) Default() StitchVideos               { return "stitch-videos" }
func (c VideoEffect) Default() VideoEffect                 { return "video-effect" }
func (c VideoToFrameGrid) Default() VideoToFrameGrid       { return "video-to-frame-grid" }
func (c VideoToLongExposure) Default() VideoToLongExposure { return "video-to-long-exposure" }
func (c Workspace) Default() Workspace                     { return "workspace" }

func (c AddShapeToImageBrowser) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c AddTextToImageBrowser) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c BlurImageBrowser) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BoomerangVideo) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ChangeImageArBrowser) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ChangeVideoAr) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ColorFilterImageBrowser) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ColorFilterVideo) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ColorGradeImageBrowser) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ColorGradeVideo) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ColorTintImageBrowser) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ConcatTextBrowser) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c CropImageBrowser) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c DrawImageBrowser) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c DuplicateImageBrowser) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c DuplicateVideo) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ExtractVideoFrames) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FilterColorImageBrowser) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c FindAndReplaceTextBrowser) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c GreenscreenVideo) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c KenBurnsVideo) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c MergeAudioIntoVideo) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c MultipartFormData) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c OverlayImageBrowser) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Post) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c QrCodeGeneratorBrowser) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ResizeImageBrowser) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ResizeVideo) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ReverseVideo) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c RotateImageBrowser) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Scene3dImageBrowser) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ShaderEffectBrowser) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c SideBySideCompositeBrowser) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c SlowDownVideo) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c SpeedUpVideo) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c SplitAudioFromVideo) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c SplitTextBrowser) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c SplitVideo) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c StitchVideos) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c VideoEffect) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c VideoToFrameGrid) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c VideoToLongExposure) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Workspace) MarshalJSON() ([]byte, error)                  { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
