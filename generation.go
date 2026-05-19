// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"net/http"
	"slices"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

// Generation endpoints.
//
// GenerationService contains methods and other services that help with interacting
// with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGenerationService] method instead.
type GenerationService struct {
	options []option.RequestOption
}

// NewGenerationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGenerationService(opts ...option.RequestOption) (r GenerationService) {
	r = GenerationService{}
	r.options = opts
	return
}

// Starts a model generation using a prompt, workspace, project, optional model,
// and optional model parameters. Poll the returned run_id via GET /runs/{runId}
// for progress and outputs. Mutating public API requests support an optional
// Idempotency-Key header for client retries; duplicate keys within two hours
// return idempotency_duplicate.
func (r *GenerationService) New(ctx context.Context, body GenerationNewParams, opts ...option.RequestOption) (res *GenerationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type GenerationNewResponse struct {
	// Cost charged in USD
	ChargedCost      float64 `json:"charged_cost" api:"required"`
	EstimatedSeconds int64   `json:"estimated_seconds" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Run type
	//
	// Any of "generation", "technique", "action".
	Type    GenerationNewResponseType   `json:"type" api:"required"`
	Action  GenerationNewResponseAction `json:"action" api:"nullable"`
	Model   GenerationNewResponseModel  `json:"model" api:"nullable"`
	PollURL string                      `json:"poll_url" api:"nullable" format:"uri"`
	// Project identifier
	ProjectID string                         `json:"project_id" api:"nullable"`
	Technique GenerationNewResponseTechnique `json:"technique" api:"nullable"`
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
func (r GenerationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GenerationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run type
type GenerationNewResponseType string

const (
	GenerationNewResponseTypeGeneration GenerationNewResponseType = "generation"
	GenerationNewResponseTypeTechnique  GenerationNewResponseType = "technique"
	GenerationNewResponseTypeAction     GenerationNewResponseType = "action"
)

type GenerationNewResponseAction struct {
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
func (r GenerationNewResponseAction) RawJSON() string { return r.JSON.raw }
func (r *GenerationNewResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationNewResponseModel struct {
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
func (r GenerationNewResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GenerationNewResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationNewResponseTechnique struct {
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
func (r GenerationNewResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *GenerationNewResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationNewParams struct {
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Generation prompt
	Prompt string `json:"prompt" api:"required"`
	// Generation type
	//
	// Any of "image", "video", "audio", "text".
	Type GenerationNewParamsType `json:"type,omitzero" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Model endpoint ID
	Model param.Opt[string] `json:"model,omitzero"`
	// Model parameters
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r GenerationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GenerationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generation type
type GenerationNewParamsType string

const (
	GenerationNewParamsTypeImage GenerationNewParamsType = "image"
	GenerationNewParamsTypeVideo GenerationNewParamsType = "video"
	GenerationNewParamsTypeAudio GenerationNewParamsType = "audio"
	GenerationNewParamsTypeText  GenerationNewParamsType = "text"
)
