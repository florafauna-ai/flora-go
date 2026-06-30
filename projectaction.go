// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

// Project canvas endpoints.
//
// ProjectActionService contains methods and other services that help with
// interacting with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectActionService] method instead.
type ProjectActionService struct {
	options []option.RequestOption
}

// NewProjectActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectActionService(opts ...option.RequestOption) (r ProjectActionService) {
	r = ProjectActionService{}
	r.options = opts
	return
}

// Creates a prebuilt action node on a project canvas using a raw action slug.
func (r *ProjectActionService) New(ctx context.Context, projectID string, body ProjectActionNewParams, opts ...option.RequestOption) (res *ProjectActionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/actions", url.PathEscape(projectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Runs an existing canvas action node through the action execution workflow.
func (r *ProjectActionService) Run(ctx context.Context, nodeID string, body ProjectActionRunParams, opts ...option.RequestOption) (res *ProjectActionRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if nodeID == "" {
		err = errors.New("missing required nodeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/actions/%s/run", url.PathEscape(body.ProjectID), url.PathEscape(nodeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ProjectActionNewResponse struct {
	// Action identifier
	//
	// Any of "color-grade-image-browser", "overlay-image-browser",
	// "draw-image-browser", "crop-image-browser", "scene-3d-image-browser",
	// "blur-image-browser", "change-image-ar-browser", "rotate-image-browser",
	// "color-filter-image-browser", "color-tint-image-browser",
	// "filter-color-image-browser", "duplicate-image-browser",
	// "side-by-side-composite-browser", "add-shape-to-image-browser",
	// "add-text-to-image-browser", "qr-code-generator-browser",
	// "resize-image-browser", "shader-effect-browser", "split-text-browser",
	// "find-and-replace-text-browser", "concat-text-browser", "ken-burns-video",
	// "stitch-videos", "split-video", "extract-video-frames", "color-grade-video",
	// "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID ProjectActionNewResponseActionID `json:"action_id" api:"required"`
	// Project canvas URL
	CanvasURL string `json:"canvas_url" api:"required" format:"uri"`
	// Canvas action node identifier
	NodeID string `json:"node_id" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		CanvasURL   respjson.Field
		NodeID      respjson.Field
		ProjectID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectActionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectActionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action identifier
type ProjectActionNewResponseActionID string

const (
	ProjectActionNewResponseActionIDColorGradeImageBrowser     ProjectActionNewResponseActionID = "color-grade-image-browser"
	ProjectActionNewResponseActionIDOverlayImageBrowser        ProjectActionNewResponseActionID = "overlay-image-browser"
	ProjectActionNewResponseActionIDDrawImageBrowser           ProjectActionNewResponseActionID = "draw-image-browser"
	ProjectActionNewResponseActionIDCropImageBrowser           ProjectActionNewResponseActionID = "crop-image-browser"
	ProjectActionNewResponseActionIDScene3dImageBrowser        ProjectActionNewResponseActionID = "scene-3d-image-browser"
	ProjectActionNewResponseActionIDBlurImageBrowser           ProjectActionNewResponseActionID = "blur-image-browser"
	ProjectActionNewResponseActionIDChangeImageArBrowser       ProjectActionNewResponseActionID = "change-image-ar-browser"
	ProjectActionNewResponseActionIDRotateImageBrowser         ProjectActionNewResponseActionID = "rotate-image-browser"
	ProjectActionNewResponseActionIDColorFilterImageBrowser    ProjectActionNewResponseActionID = "color-filter-image-browser"
	ProjectActionNewResponseActionIDColorTintImageBrowser      ProjectActionNewResponseActionID = "color-tint-image-browser"
	ProjectActionNewResponseActionIDFilterColorImageBrowser    ProjectActionNewResponseActionID = "filter-color-image-browser"
	ProjectActionNewResponseActionIDDuplicateImageBrowser      ProjectActionNewResponseActionID = "duplicate-image-browser"
	ProjectActionNewResponseActionIDSideBySideCompositeBrowser ProjectActionNewResponseActionID = "side-by-side-composite-browser"
	ProjectActionNewResponseActionIDAddShapeToImageBrowser     ProjectActionNewResponseActionID = "add-shape-to-image-browser"
	ProjectActionNewResponseActionIDAddTextToImageBrowser      ProjectActionNewResponseActionID = "add-text-to-image-browser"
	ProjectActionNewResponseActionIDQrCodeGeneratorBrowser     ProjectActionNewResponseActionID = "qr-code-generator-browser"
	ProjectActionNewResponseActionIDResizeImageBrowser         ProjectActionNewResponseActionID = "resize-image-browser"
	ProjectActionNewResponseActionIDShaderEffectBrowser        ProjectActionNewResponseActionID = "shader-effect-browser"
	ProjectActionNewResponseActionIDSplitTextBrowser           ProjectActionNewResponseActionID = "split-text-browser"
	ProjectActionNewResponseActionIDFindAndReplaceTextBrowser  ProjectActionNewResponseActionID = "find-and-replace-text-browser"
	ProjectActionNewResponseActionIDConcatTextBrowser          ProjectActionNewResponseActionID = "concat-text-browser"
	ProjectActionNewResponseActionIDKenBurnsVideo              ProjectActionNewResponseActionID = "ken-burns-video"
	ProjectActionNewResponseActionIDStitchVideos               ProjectActionNewResponseActionID = "stitch-videos"
	ProjectActionNewResponseActionIDSplitVideo                 ProjectActionNewResponseActionID = "split-video"
	ProjectActionNewResponseActionIDExtractVideoFrames         ProjectActionNewResponseActionID = "extract-video-frames"
	ProjectActionNewResponseActionIDColorGradeVideo            ProjectActionNewResponseActionID = "color-grade-video"
	ProjectActionNewResponseActionIDVideoToFrameGrid           ProjectActionNewResponseActionID = "video-to-frame-grid"
	ProjectActionNewResponseActionIDBoomerangVideo             ProjectActionNewResponseActionID = "boomerang-video"
	ProjectActionNewResponseActionIDReverseVideo               ProjectActionNewResponseActionID = "reverse-video"
	ProjectActionNewResponseActionIDVideoToLongExposure        ProjectActionNewResponseActionID = "video-to-long-exposure"
	ProjectActionNewResponseActionIDVideoEffect                ProjectActionNewResponseActionID = "video-effect"
	ProjectActionNewResponseActionIDColorFilterVideo           ProjectActionNewResponseActionID = "color-filter-video"
	ProjectActionNewResponseActionIDSpeedUpVideo               ProjectActionNewResponseActionID = "speed-up-video"
	ProjectActionNewResponseActionIDSlowDownVideo              ProjectActionNewResponseActionID = "slow-down-video"
	ProjectActionNewResponseActionIDDuplicateVideo             ProjectActionNewResponseActionID = "duplicate-video"
	ProjectActionNewResponseActionIDGreenscreenVideo           ProjectActionNewResponseActionID = "greenscreen-video"
	ProjectActionNewResponseActionIDResizeVideo                ProjectActionNewResponseActionID = "resize-video"
	ProjectActionNewResponseActionIDChangeVideoAr              ProjectActionNewResponseActionID = "change-video-ar"
	ProjectActionNewResponseActionIDSplitAudioFromVideo        ProjectActionNewResponseActionID = "split-audio-from-video"
	ProjectActionNewResponseActionIDMergeAudioIntoVideo        ProjectActionNewResponseActionID = "merge-audio-into-video"
)

type ProjectActionRunResponse struct {
	// Cost charged in USD
	ChargedCost      float64 `json:"charged_cost" api:"required"`
	EstimatedSeconds int64   `json:"estimated_seconds" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Run type
	//
	// Any of "generation", "technique", "action".
	Type   ProjectActionRunResponseType   `json:"type" api:"required"`
	Action ProjectActionRunResponseAction `json:"action" api:"nullable"`
	Model  ProjectActionRunResponseModel  `json:"model" api:"nullable"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL string `json:"poll_url" api:"nullable" format:"uri"`
	// Project identifier
	ProjectID string                            `json:"project_id" api:"nullable"`
	Technique ProjectActionRunResponseTechnique `json:"technique" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargedCost      respjson.Field
		EstimatedSeconds respjson.Field
		RunID            respjson.Field
		Type             respjson.Field
		Action           respjson.Field
		Model            respjson.Field
		PollURL          respjson.Field
		ProjectID        respjson.Field
		Technique        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectActionRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectActionRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run type
type ProjectActionRunResponseType string

const (
	ProjectActionRunResponseTypeGeneration ProjectActionRunResponseType = "generation"
	ProjectActionRunResponseTypeTechnique  ProjectActionRunResponseType = "technique"
	ProjectActionRunResponseTypeAction     ProjectActionRunResponseType = "action"
)

type ProjectActionRunResponseAction struct {
	// Action identifier
	//
	// Any of "color-grade-image-browser", "overlay-image-browser",
	// "draw-image-browser", "crop-image-browser", "scene-3d-image-browser",
	// "blur-image-browser", "change-image-ar-browser", "rotate-image-browser",
	// "color-filter-image-browser", "color-tint-image-browser",
	// "filter-color-image-browser", "duplicate-image-browser",
	// "side-by-side-composite-browser", "add-shape-to-image-browser",
	// "add-text-to-image-browser", "qr-code-generator-browser",
	// "resize-image-browser", "shader-effect-browser", "split-text-browser",
	// "find-and-replace-text-browser", "concat-text-browser", "ken-burns-video",
	// "stitch-videos", "split-video", "extract-video-frames", "color-grade-video",
	// "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID string `json:"action_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectActionRunResponseAction) RawJSON() string { return r.JSON.raw }
func (r *ProjectActionRunResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectActionRunResponseModel struct {
	// Model identifier
	ModelID string `json:"model_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectActionRunResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ProjectActionRunResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectActionRunResponseTechnique struct {
	// Technique name
	Name string `json:"name" api:"required"`
	// Technique identifier
	TechniqueID string `json:"technique_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		TechniqueID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectActionRunResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *ProjectActionRunResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectActionNewParams struct {
	// Action identifier
	//
	// Any of "color-grade-image-browser", "overlay-image-browser",
	// "draw-image-browser", "crop-image-browser", "scene-3d-image-browser",
	// "blur-image-browser", "change-image-ar-browser", "rotate-image-browser",
	// "color-filter-image-browser", "color-tint-image-browser",
	// "filter-color-image-browser", "duplicate-image-browser",
	// "side-by-side-composite-browser", "add-shape-to-image-browser",
	// "add-text-to-image-browser", "qr-code-generator-browser",
	// "resize-image-browser", "shader-effect-browser", "split-text-browser",
	// "find-and-replace-text-browser", "concat-text-browser", "ken-burns-video",
	// "stitch-videos", "split-video", "extract-video-frames", "color-grade-video",
	// "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID ProjectActionNewParamsActionID `json:"action_id,omitzero" api:"required"`
	// Action parameters (snake_case keys). The accepted keys depend on action_id; see
	// GET /actions/{actionId} or POST /runs/action for the per-action schema.
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r ProjectActionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectActionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectActionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action identifier
type ProjectActionNewParamsActionID string

const (
	ProjectActionNewParamsActionIDColorGradeImageBrowser     ProjectActionNewParamsActionID = "color-grade-image-browser"
	ProjectActionNewParamsActionIDOverlayImageBrowser        ProjectActionNewParamsActionID = "overlay-image-browser"
	ProjectActionNewParamsActionIDDrawImageBrowser           ProjectActionNewParamsActionID = "draw-image-browser"
	ProjectActionNewParamsActionIDCropImageBrowser           ProjectActionNewParamsActionID = "crop-image-browser"
	ProjectActionNewParamsActionIDScene3dImageBrowser        ProjectActionNewParamsActionID = "scene-3d-image-browser"
	ProjectActionNewParamsActionIDBlurImageBrowser           ProjectActionNewParamsActionID = "blur-image-browser"
	ProjectActionNewParamsActionIDChangeImageArBrowser       ProjectActionNewParamsActionID = "change-image-ar-browser"
	ProjectActionNewParamsActionIDRotateImageBrowser         ProjectActionNewParamsActionID = "rotate-image-browser"
	ProjectActionNewParamsActionIDColorFilterImageBrowser    ProjectActionNewParamsActionID = "color-filter-image-browser"
	ProjectActionNewParamsActionIDColorTintImageBrowser      ProjectActionNewParamsActionID = "color-tint-image-browser"
	ProjectActionNewParamsActionIDFilterColorImageBrowser    ProjectActionNewParamsActionID = "filter-color-image-browser"
	ProjectActionNewParamsActionIDDuplicateImageBrowser      ProjectActionNewParamsActionID = "duplicate-image-browser"
	ProjectActionNewParamsActionIDSideBySideCompositeBrowser ProjectActionNewParamsActionID = "side-by-side-composite-browser"
	ProjectActionNewParamsActionIDAddShapeToImageBrowser     ProjectActionNewParamsActionID = "add-shape-to-image-browser"
	ProjectActionNewParamsActionIDAddTextToImageBrowser      ProjectActionNewParamsActionID = "add-text-to-image-browser"
	ProjectActionNewParamsActionIDQrCodeGeneratorBrowser     ProjectActionNewParamsActionID = "qr-code-generator-browser"
	ProjectActionNewParamsActionIDResizeImageBrowser         ProjectActionNewParamsActionID = "resize-image-browser"
	ProjectActionNewParamsActionIDShaderEffectBrowser        ProjectActionNewParamsActionID = "shader-effect-browser"
	ProjectActionNewParamsActionIDSplitTextBrowser           ProjectActionNewParamsActionID = "split-text-browser"
	ProjectActionNewParamsActionIDFindAndReplaceTextBrowser  ProjectActionNewParamsActionID = "find-and-replace-text-browser"
	ProjectActionNewParamsActionIDConcatTextBrowser          ProjectActionNewParamsActionID = "concat-text-browser"
	ProjectActionNewParamsActionIDKenBurnsVideo              ProjectActionNewParamsActionID = "ken-burns-video"
	ProjectActionNewParamsActionIDStitchVideos               ProjectActionNewParamsActionID = "stitch-videos"
	ProjectActionNewParamsActionIDSplitVideo                 ProjectActionNewParamsActionID = "split-video"
	ProjectActionNewParamsActionIDExtractVideoFrames         ProjectActionNewParamsActionID = "extract-video-frames"
	ProjectActionNewParamsActionIDColorGradeVideo            ProjectActionNewParamsActionID = "color-grade-video"
	ProjectActionNewParamsActionIDVideoToFrameGrid           ProjectActionNewParamsActionID = "video-to-frame-grid"
	ProjectActionNewParamsActionIDBoomerangVideo             ProjectActionNewParamsActionID = "boomerang-video"
	ProjectActionNewParamsActionIDReverseVideo               ProjectActionNewParamsActionID = "reverse-video"
	ProjectActionNewParamsActionIDVideoToLongExposure        ProjectActionNewParamsActionID = "video-to-long-exposure"
	ProjectActionNewParamsActionIDVideoEffect                ProjectActionNewParamsActionID = "video-effect"
	ProjectActionNewParamsActionIDColorFilterVideo           ProjectActionNewParamsActionID = "color-filter-video"
	ProjectActionNewParamsActionIDSpeedUpVideo               ProjectActionNewParamsActionID = "speed-up-video"
	ProjectActionNewParamsActionIDSlowDownVideo              ProjectActionNewParamsActionID = "slow-down-video"
	ProjectActionNewParamsActionIDDuplicateVideo             ProjectActionNewParamsActionID = "duplicate-video"
	ProjectActionNewParamsActionIDGreenscreenVideo           ProjectActionNewParamsActionID = "greenscreen-video"
	ProjectActionNewParamsActionIDResizeVideo                ProjectActionNewParamsActionID = "resize-video"
	ProjectActionNewParamsActionIDChangeVideoAr              ProjectActionNewParamsActionID = "change-video-ar"
	ProjectActionNewParamsActionIDSplitAudioFromVideo        ProjectActionNewParamsActionID = "split-audio-from-video"
	ProjectActionNewParamsActionIDMergeAudioIntoVideo        ProjectActionNewParamsActionID = "merge-audio-into-video"
)

type ProjectActionRunParams struct {
	// Project identifier
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}
