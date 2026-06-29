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

// Top-level run creation endpoints.
//
// RunService contains methods and other services that help with interacting with
// the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunService] method instead.
type RunService struct {
	options []option.RequestOption
}

// NewRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunService(opts ...option.RequestOption) (r RunService) {
	r = RunService{}
	r.options = opts
	return
}

// Starts a model generation run in a project canvas using type, prompt,
// workspace_id, project_id, optional model endpoint ID, and optional model
// parameters. Use type=image|video|audio|text and model IDs returned by GET
// /models or list_models. Mutating public API requests support an optional
// Idempotency-Key header for client retries; duplicate keys within two hours
// return idempotency_duplicate.
//
// Deprecated: deprecated
func (r *RunService) StartGeneration(ctx context.Context, body RunStartGenerationParams, opts ...option.RequestOption) (res *RunStartGenerationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "runs/generation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts a technique run through the normalized top-level run resource using
// technique*id, workspace_id, and inputs. technique_id must use the tech* public
// API ID returned by list techniques. Mutating public API requests support an
// optional Idempotency-Key header for client retries; duplicate keys within two
// hours return idempotency_duplicate.
func (r *RunService) StartTechnique(ctx context.Context, body RunStartTechniqueParams, opts ...option.RequestOption) (res *RunStartTechniqueResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "runs/technique"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type RunStartGenerationResponse struct {
	// Cost charged in USD
	ChargedCost      float64 `json:"charged_cost" api:"required"`
	EstimatedSeconds int64   `json:"estimated_seconds" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Run type
	//
	// Any of "generation", "technique", "action".
	Type   RunStartGenerationResponseType   `json:"type" api:"required"`
	Action RunStartGenerationResponseAction `json:"action" api:"nullable"`
	Model  RunStartGenerationResponseModel  `json:"model" api:"nullable"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL string `json:"poll_url" api:"nullable" format:"uri"`
	// Project identifier
	ProjectID string                              `json:"project_id" api:"nullable"`
	Technique RunStartGenerationResponseTechnique `json:"technique" api:"nullable"`
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
func (r RunStartGenerationResponse) RawJSON() string { return r.JSON.raw }
func (r *RunStartGenerationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run type
type RunStartGenerationResponseType string

const (
	RunStartGenerationResponseTypeGeneration RunStartGenerationResponseType = "generation"
	RunStartGenerationResponseTypeTechnique  RunStartGenerationResponseType = "technique"
	RunStartGenerationResponseTypeAction     RunStartGenerationResponseType = "action"
)

type RunStartGenerationResponseAction struct {
	// Action identifier
	//
	// Any of "split-text", "find-and-replace-text", "concat-text",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "ken-burns-video", "stitch-videos", "split-video",
	// "extract-video-frames", "color-grade-video", "video-to-frame-grid",
	// "boomerang-video", "reverse-video", "video-to-long-exposure", "video-effect",
	// "color-filter-video", "speed-up-video", "slow-down-video", "duplicate-video",
	// "greenscreen-video", "resize-video", "change-video-ar",
	// "split-audio-from-video", "merge-audio-into-video".
	ActionID string `json:"action_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunStartGenerationResponseAction) RawJSON() string { return r.JSON.raw }
func (r *RunStartGenerationResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartGenerationResponseModel struct {
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
func (r RunStartGenerationResponseModel) RawJSON() string { return r.JSON.raw }
func (r *RunStartGenerationResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartGenerationResponseTechnique struct {
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
func (r RunStartGenerationResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *RunStartGenerationResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartTechniqueResponse struct {
	// Cost charged in USD
	ChargedCost      float64 `json:"charged_cost" api:"required"`
	EstimatedSeconds int64   `json:"estimated_seconds" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Run type
	//
	// Any of "generation", "technique", "action".
	Type   RunStartTechniqueResponseType   `json:"type" api:"required"`
	Action RunStartTechniqueResponseAction `json:"action" api:"nullable"`
	Model  RunStartTechniqueResponseModel  `json:"model" api:"nullable"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL string `json:"poll_url" api:"nullable" format:"uri"`
	// Project identifier
	ProjectID string                             `json:"project_id" api:"nullable"`
	Technique RunStartTechniqueResponseTechnique `json:"technique" api:"nullable"`
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
func (r RunStartTechniqueResponse) RawJSON() string { return r.JSON.raw }
func (r *RunStartTechniqueResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run type
type RunStartTechniqueResponseType string

const (
	RunStartTechniqueResponseTypeGeneration RunStartTechniqueResponseType = "generation"
	RunStartTechniqueResponseTypeTechnique  RunStartTechniqueResponseType = "technique"
	RunStartTechniqueResponseTypeAction     RunStartTechniqueResponseType = "action"
)

type RunStartTechniqueResponseAction struct {
	// Action identifier
	//
	// Any of "split-text", "find-and-replace-text", "concat-text",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "ken-burns-video", "stitch-videos", "split-video",
	// "extract-video-frames", "color-grade-video", "video-to-frame-grid",
	// "boomerang-video", "reverse-video", "video-to-long-exposure", "video-effect",
	// "color-filter-video", "speed-up-video", "slow-down-video", "duplicate-video",
	// "greenscreen-video", "resize-video", "change-video-ar",
	// "split-audio-from-video", "merge-audio-into-video".
	ActionID string `json:"action_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunStartTechniqueResponseAction) RawJSON() string { return r.JSON.raw }
func (r *RunStartTechniqueResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartTechniqueResponseModel struct {
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
func (r RunStartTechniqueResponseModel) RawJSON() string { return r.JSON.raw }
func (r *RunStartTechniqueResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartTechniqueResponseTechnique struct {
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
func (r RunStartTechniqueResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *RunStartTechniqueResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunStartGenerationParams struct {
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_.
	ProjectID string `json:"project_id" api:"required"`
	// Generation prompt
	Prompt string `json:"prompt" api:"required"`
	// Generation type. Use "image", "video", "audio", or "text"; do not pass model
	// families such as "t2i" or "i2v".
	//
	// Any of "image", "video", "audio", "text".
	Type RunStartGenerationParamsType `json:"type,omitzero" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Model endpoint ID, not a display name. Use list_models (or GET /models) to find
	// accessible endpoint IDs for the requested type.
	Model param.Opt[string] `json:"model,omitzero"`
	// Model parameters
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r RunStartGenerationParams) MarshalJSON() (data []byte, err error) {
	type shadow RunStartGenerationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RunStartGenerationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generation type. Use "image", "video", "audio", or "text"; do not pass model
// families such as "t2i" or "i2v".
type RunStartGenerationParamsType string

const (
	RunStartGenerationParamsTypeImage RunStartGenerationParamsType = "image"
	RunStartGenerationParamsTypeVideo RunStartGenerationParamsType = "video"
	RunStartGenerationParamsTypeAudio RunStartGenerationParamsType = "audio"
	RunStartGenerationParamsTypeText  RunStartGenerationParamsType = "text"
)

type RunStartTechniqueParams struct {
	// Technique inputs
	Inputs map[string]any `json:"inputs,omitzero" api:"required"`
	// Technique identifier. Use the public API ID returned by list techniques; it must
	// start with tech\_.
	TechniqueID string `json:"technique_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	paramObj
}

func (r RunStartTechniqueParams) MarshalJSON() (data []byte, err error) {
	type shadow RunStartTechniqueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RunStartTechniqueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
