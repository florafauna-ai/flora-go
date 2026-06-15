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

type AddShapeToImage string     // Always "add-shape-to-image"
type AddTextToImage string      // Always "add-text-to-image"
type BlurImage string           // Always "blur-image"
type BoomerangVideo string      // Always "boomerang-video"
type ChangeImageAr string       // Always "change-image-ar"
type ChangeVideoAr string       // Always "change-video-ar"
type ColorFilterImage string    // Always "color-filter-image"
type ColorFilterVideo string    // Always "color-filter-video"
type ColorGradeImage string     // Always "color-grade-image"
type ColorGradeVideo string     // Always "color-grade-video"
type ColorTintImage string      // Always "color-tint-image"
type ConcatText string          // Always "concat-text"
type DuplicateImage string      // Always "duplicate-image"
type DuplicateVideo string      // Always "duplicate-video"
type ExtractVideoFrames string  // Always "extract-video-frames"
type File string                // Always "file"
type FilterColorImage string    // Always "filter-color-image"
type FindAndReplaceText string  // Always "find-and-replace-text"
type FlipImage string           // Always "flip-image"
type GenerateShapeImage string  // Always "generate-shape-image"
type GenerateTextImage string   // Always "generate-text-image"
type GreenscreenVideo string    // Always "greenscreen-video"
type KenBurnsVideo string       // Always "ken-burns-video"
type MergeAudioIntoVideo string // Always "merge-audio-into-video"
type MultipartFormData string   // Always "multipart/form-data"
type Post string                // Always "POST"
type QrCodeGenerator string     // Always "qr-code-generator"
type ResizeVideo string         // Always "resize-video"
type ReverseVideo string        // Always "reverse-video"
type RotateImage string         // Always "rotate-image"
type SideBySideComposite string // Always "side-by-side-composite"
type SlowDownVideo string       // Always "slow-down-video"
type SpeedUpVideo string        // Always "speed-up-video"
type SplitAudioFromVideo string // Always "split-audio-from-video"
type SplitText string           // Always "split-text"
type SplitVideo string          // Always "split-video"
type StitchVideos string        // Always "stitch-videos"
type VideoEffect string         // Always "video-effect"
type VideoToFrameGrid string    // Always "video-to-frame-grid"
type VideoToLongExposure string // Always "video-to-long-exposure"
type Workspace string           // Always "workspace"

func (c AddShapeToImage) Default() AddShapeToImage         { return "add-shape-to-image" }
func (c AddTextToImage) Default() AddTextToImage           { return "add-text-to-image" }
func (c BlurImage) Default() BlurImage                     { return "blur-image" }
func (c BoomerangVideo) Default() BoomerangVideo           { return "boomerang-video" }
func (c ChangeImageAr) Default() ChangeImageAr             { return "change-image-ar" }
func (c ChangeVideoAr) Default() ChangeVideoAr             { return "change-video-ar" }
func (c ColorFilterImage) Default() ColorFilterImage       { return "color-filter-image" }
func (c ColorFilterVideo) Default() ColorFilterVideo       { return "color-filter-video" }
func (c ColorGradeImage) Default() ColorGradeImage         { return "color-grade-image" }
func (c ColorGradeVideo) Default() ColorGradeVideo         { return "color-grade-video" }
func (c ColorTintImage) Default() ColorTintImage           { return "color-tint-image" }
func (c ConcatText) Default() ConcatText                   { return "concat-text" }
func (c DuplicateImage) Default() DuplicateImage           { return "duplicate-image" }
func (c DuplicateVideo) Default() DuplicateVideo           { return "duplicate-video" }
func (c ExtractVideoFrames) Default() ExtractVideoFrames   { return "extract-video-frames" }
func (c File) Default() File                               { return "file" }
func (c FilterColorImage) Default() FilterColorImage       { return "filter-color-image" }
func (c FindAndReplaceText) Default() FindAndReplaceText   { return "find-and-replace-text" }
func (c FlipImage) Default() FlipImage                     { return "flip-image" }
func (c GenerateShapeImage) Default() GenerateShapeImage   { return "generate-shape-image" }
func (c GenerateTextImage) Default() GenerateTextImage     { return "generate-text-image" }
func (c GreenscreenVideo) Default() GreenscreenVideo       { return "greenscreen-video" }
func (c KenBurnsVideo) Default() KenBurnsVideo             { return "ken-burns-video" }
func (c MergeAudioIntoVideo) Default() MergeAudioIntoVideo { return "merge-audio-into-video" }
func (c MultipartFormData) Default() MultipartFormData     { return "multipart/form-data" }
func (c Post) Default() Post                               { return "POST" }
func (c QrCodeGenerator) Default() QrCodeGenerator         { return "qr-code-generator" }
func (c ResizeVideo) Default() ResizeVideo                 { return "resize-video" }
func (c ReverseVideo) Default() ReverseVideo               { return "reverse-video" }
func (c RotateImage) Default() RotateImage                 { return "rotate-image" }
func (c SideBySideComposite) Default() SideBySideComposite { return "side-by-side-composite" }
func (c SlowDownVideo) Default() SlowDownVideo             { return "slow-down-video" }
func (c SpeedUpVideo) Default() SpeedUpVideo               { return "speed-up-video" }
func (c SplitAudioFromVideo) Default() SplitAudioFromVideo { return "split-audio-from-video" }
func (c SplitText) Default() SplitText                     { return "split-text" }
func (c SplitVideo) Default() SplitVideo                   { return "split-video" }
func (c StitchVideos) Default() StitchVideos               { return "stitch-videos" }
func (c VideoEffect) Default() VideoEffect                 { return "video-effect" }
func (c VideoToFrameGrid) Default() VideoToFrameGrid       { return "video-to-frame-grid" }
func (c VideoToLongExposure) Default() VideoToLongExposure { return "video-to-long-exposure" }
func (c Workspace) Default() Workspace                     { return "workspace" }

func (c AddShapeToImage) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c AddTextToImage) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c BlurImage) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c BoomerangVideo) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ChangeImageAr) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ChangeVideoAr) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ColorFilterImage) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ColorFilterVideo) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ColorGradeImage) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ColorGradeVideo) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ColorTintImage) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ConcatText) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c DuplicateImage) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c DuplicateVideo) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ExtractVideoFrames) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c FilterColorImage) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c FindAndReplaceText) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c FlipImage) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c GenerateShapeImage) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c GenerateTextImage) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c GreenscreenVideo) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c KenBurnsVideo) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c MergeAudioIntoVideo) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c MultipartFormData) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Post) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c QrCodeGenerator) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ResizeVideo) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ReverseVideo) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c RotateImage) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SideBySideComposite) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c SlowDownVideo) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c SpeedUpVideo) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c SplitAudioFromVideo) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c SplitText) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c SplitVideo) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c StitchVideos) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c VideoEffect) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c VideoToFrameGrid) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c VideoToLongExposure) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Workspace) MarshalJSON() ([]byte, error)           { return marshalString(c) }

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
