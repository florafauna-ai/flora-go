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
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
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
	ProjectActionNewResponseActionIDSplitText           ProjectActionNewResponseActionID = "split-text"
	ProjectActionNewResponseActionIDFindAndReplaceText  ProjectActionNewResponseActionID = "find-and-replace-text"
	ProjectActionNewResponseActionIDConcatText          ProjectActionNewResponseActionID = "concat-text"
	ProjectActionNewResponseActionIDKenBurnsVideo       ProjectActionNewResponseActionID = "ken-burns-video"
	ProjectActionNewResponseActionIDColorGradeImage     ProjectActionNewResponseActionID = "color-grade-image"
	ProjectActionNewResponseActionIDChangeImageAr       ProjectActionNewResponseActionID = "change-image-ar"
	ProjectActionNewResponseActionIDRotateImage         ProjectActionNewResponseActionID = "rotate-image"
	ProjectActionNewResponseActionIDFlipImage           ProjectActionNewResponseActionID = "flip-image"
	ProjectActionNewResponseActionIDColorFilterImage    ProjectActionNewResponseActionID = "color-filter-image"
	ProjectActionNewResponseActionIDColorTintImage      ProjectActionNewResponseActionID = "color-tint-image"
	ProjectActionNewResponseActionIDFilterColorImage    ProjectActionNewResponseActionID = "filter-color-image"
	ProjectActionNewResponseActionIDBlurImage           ProjectActionNewResponseActionID = "blur-image"
	ProjectActionNewResponseActionIDDuplicateImage      ProjectActionNewResponseActionID = "duplicate-image"
	ProjectActionNewResponseActionIDSideBySideComposite ProjectActionNewResponseActionID = "side-by-side-composite"
	ProjectActionNewResponseActionIDAddShapeToImage     ProjectActionNewResponseActionID = "add-shape-to-image"
	ProjectActionNewResponseActionIDGenerateShapeImage  ProjectActionNewResponseActionID = "generate-shape-image"
	ProjectActionNewResponseActionIDAddTextToImage      ProjectActionNewResponseActionID = "add-text-to-image"
	ProjectActionNewResponseActionIDGenerateTextImage   ProjectActionNewResponseActionID = "generate-text-image"
	ProjectActionNewResponseActionIDQrCodeGenerator     ProjectActionNewResponseActionID = "qr-code-generator"
	ProjectActionNewResponseActionIDStitchVideos        ProjectActionNewResponseActionID = "stitch-videos"
	ProjectActionNewResponseActionIDSplitVideo          ProjectActionNewResponseActionID = "split-video"
	ProjectActionNewResponseActionIDExtractVideoFrames  ProjectActionNewResponseActionID = "extract-video-frames"
	ProjectActionNewResponseActionIDColorGradeVideo     ProjectActionNewResponseActionID = "color-grade-video"
	ProjectActionNewResponseActionIDVideoToFrameGrid    ProjectActionNewResponseActionID = "video-to-frame-grid"
	ProjectActionNewResponseActionIDBoomerangVideo      ProjectActionNewResponseActionID = "boomerang-video"
	ProjectActionNewResponseActionIDReverseVideo        ProjectActionNewResponseActionID = "reverse-video"
	ProjectActionNewResponseActionIDVideoToLongExposure ProjectActionNewResponseActionID = "video-to-long-exposure"
	ProjectActionNewResponseActionIDVideoEffect         ProjectActionNewResponseActionID = "video-effect"
	ProjectActionNewResponseActionIDColorFilterVideo    ProjectActionNewResponseActionID = "color-filter-video"
	ProjectActionNewResponseActionIDSpeedUpVideo        ProjectActionNewResponseActionID = "speed-up-video"
	ProjectActionNewResponseActionIDSlowDownVideo       ProjectActionNewResponseActionID = "slow-down-video"
	ProjectActionNewResponseActionIDDuplicateVideo      ProjectActionNewResponseActionID = "duplicate-video"
	ProjectActionNewResponseActionIDGreenscreenVideo    ProjectActionNewResponseActionID = "greenscreen-video"
	ProjectActionNewResponseActionIDResizeVideo         ProjectActionNewResponseActionID = "resize-video"
	ProjectActionNewResponseActionIDChangeVideoAr       ProjectActionNewResponseActionID = "change-video-ar"
	ProjectActionNewResponseActionIDSplitAudioFromVideo ProjectActionNewResponseActionID = "split-audio-from-video"
	ProjectActionNewResponseActionIDMergeAudioIntoVideo ProjectActionNewResponseActionID = "merge-audio-into-video"
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
	Type    ProjectActionRunResponseType   `json:"type" api:"required"`
	Action  ProjectActionRunResponseAction `json:"action" api:"nullable"`
	Model   ProjectActionRunResponseModel  `json:"model" api:"nullable"`
	PollURL string                         `json:"poll_url" api:"nullable" format:"uri"`
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
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
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
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID ProjectActionNewParamsActionID `json:"action_id,omitzero" api:"required"`
	// Action parameters
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
	ProjectActionNewParamsActionIDSplitText           ProjectActionNewParamsActionID = "split-text"
	ProjectActionNewParamsActionIDFindAndReplaceText  ProjectActionNewParamsActionID = "find-and-replace-text"
	ProjectActionNewParamsActionIDConcatText          ProjectActionNewParamsActionID = "concat-text"
	ProjectActionNewParamsActionIDKenBurnsVideo       ProjectActionNewParamsActionID = "ken-burns-video"
	ProjectActionNewParamsActionIDColorGradeImage     ProjectActionNewParamsActionID = "color-grade-image"
	ProjectActionNewParamsActionIDChangeImageAr       ProjectActionNewParamsActionID = "change-image-ar"
	ProjectActionNewParamsActionIDRotateImage         ProjectActionNewParamsActionID = "rotate-image"
	ProjectActionNewParamsActionIDFlipImage           ProjectActionNewParamsActionID = "flip-image"
	ProjectActionNewParamsActionIDColorFilterImage    ProjectActionNewParamsActionID = "color-filter-image"
	ProjectActionNewParamsActionIDColorTintImage      ProjectActionNewParamsActionID = "color-tint-image"
	ProjectActionNewParamsActionIDFilterColorImage    ProjectActionNewParamsActionID = "filter-color-image"
	ProjectActionNewParamsActionIDBlurImage           ProjectActionNewParamsActionID = "blur-image"
	ProjectActionNewParamsActionIDDuplicateImage      ProjectActionNewParamsActionID = "duplicate-image"
	ProjectActionNewParamsActionIDSideBySideComposite ProjectActionNewParamsActionID = "side-by-side-composite"
	ProjectActionNewParamsActionIDAddShapeToImage     ProjectActionNewParamsActionID = "add-shape-to-image"
	ProjectActionNewParamsActionIDGenerateShapeImage  ProjectActionNewParamsActionID = "generate-shape-image"
	ProjectActionNewParamsActionIDAddTextToImage      ProjectActionNewParamsActionID = "add-text-to-image"
	ProjectActionNewParamsActionIDGenerateTextImage   ProjectActionNewParamsActionID = "generate-text-image"
	ProjectActionNewParamsActionIDQrCodeGenerator     ProjectActionNewParamsActionID = "qr-code-generator"
	ProjectActionNewParamsActionIDStitchVideos        ProjectActionNewParamsActionID = "stitch-videos"
	ProjectActionNewParamsActionIDSplitVideo          ProjectActionNewParamsActionID = "split-video"
	ProjectActionNewParamsActionIDExtractVideoFrames  ProjectActionNewParamsActionID = "extract-video-frames"
	ProjectActionNewParamsActionIDColorGradeVideo     ProjectActionNewParamsActionID = "color-grade-video"
	ProjectActionNewParamsActionIDVideoToFrameGrid    ProjectActionNewParamsActionID = "video-to-frame-grid"
	ProjectActionNewParamsActionIDBoomerangVideo      ProjectActionNewParamsActionID = "boomerang-video"
	ProjectActionNewParamsActionIDReverseVideo        ProjectActionNewParamsActionID = "reverse-video"
	ProjectActionNewParamsActionIDVideoToLongExposure ProjectActionNewParamsActionID = "video-to-long-exposure"
	ProjectActionNewParamsActionIDVideoEffect         ProjectActionNewParamsActionID = "video-effect"
	ProjectActionNewParamsActionIDColorFilterVideo    ProjectActionNewParamsActionID = "color-filter-video"
	ProjectActionNewParamsActionIDSpeedUpVideo        ProjectActionNewParamsActionID = "speed-up-video"
	ProjectActionNewParamsActionIDSlowDownVideo       ProjectActionNewParamsActionID = "slow-down-video"
	ProjectActionNewParamsActionIDDuplicateVideo      ProjectActionNewParamsActionID = "duplicate-video"
	ProjectActionNewParamsActionIDGreenscreenVideo    ProjectActionNewParamsActionID = "greenscreen-video"
	ProjectActionNewParamsActionIDResizeVideo         ProjectActionNewParamsActionID = "resize-video"
	ProjectActionNewParamsActionIDChangeVideoAr       ProjectActionNewParamsActionID = "change-video-ar"
	ProjectActionNewParamsActionIDSplitAudioFromVideo ProjectActionNewParamsActionID = "split-audio-from-video"
	ProjectActionNewParamsActionIDMergeAudioIntoVideo ProjectActionNewParamsActionID = "merge-audio-into-video"
)

type ProjectActionRunParams struct {
	// Project identifier
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}
