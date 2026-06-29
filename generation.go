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
	"github.com/florafauna-ai/flora-go/internal/apiquery"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/pagination"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

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

// Starts a model generation using type, prompt, workspace_id, project_id, optional
// model endpoint ID, and optional model parameters. Use
// type=image|video|audio|text and model IDs returned by GET /models or
// list_models. Poll the returned run_id via GET /runs/{runId} for progress and
// outputs. Mutating public API requests support an optional Idempotency-Key header
// for client retries; duplicate keys within two hours return
// idempotency_duplicate.
func (r *GenerationService) New(ctx context.Context, body GenerationNewParams, opts ...option.RequestOption) (res *GenerationNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns status and completed output URLs for a public API run, including action
// runs started through POST /runs/action.
func (r *GenerationService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *GenerationGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("runs/%s", url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists generation history for the authenticated caller, including pending,
// running, completed, and failed generations. Results are newest first and can be
// filtered by workspace_id, project_id, and status. Each item includes poll_url;
// use it to poll pending/running generations and to fetch completed or failed run
// details and outputs.
func (r *GenerationService) List(ctx context.Context, query GenerationListParams, opts ...option.RequestOption) (res *pagination.GenerationsCursorPage[GenerationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "generations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists generation history for the authenticated caller, including pending,
// running, completed, and failed generations. Results are newest first and can be
// filtered by workspace_id, project_id, and status. Each item includes poll_url;
// use it to poll pending/running generations and to fetch completed or failed run
// details and outputs.
func (r *GenerationService) ListAutoPaging(ctx context.Context, query GenerationListParams, opts ...option.RequestOption) *pagination.GenerationsCursorPageAutoPager[GenerationListResponse] {
	return pagination.NewGenerationsCursorPageAutoPager(r.List(ctx, query, opts...))
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
	Type   GenerationNewResponseType   `json:"type" api:"required"`
	Action GenerationNewResponseAction `json:"action" api:"nullable"`
	Model  GenerationNewResponseModel  `json:"model" api:"nullable"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL string `json:"poll_url" api:"nullable" format:"uri"`
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

type GenerationGetResponse struct {
	CreatedAt float64 `json:"created_at" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status GenerationGetResponseStatus `json:"status" api:"required"`
	// Cost charged in USD
	ChargedCost float64 `json:"charged_cost"`
	CompletedAt float64 `json:"completed_at"`
	// Machine-readable run error code
	ErrorCode string `json:"error_code"`
	// Human-readable run error message
	ErrorMessage string                        `json:"error_message"`
	Outputs      []GenerationGetResponseOutput `json:"outputs"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL   string  `json:"poll_url" format:"uri"`
	StartedAt float64 `json:"started_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Progress     respjson.Field
		RunID        respjson.Field
		Status       respjson.Field
		ChargedCost  respjson.Field
		CompletedAt  respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		Outputs      respjson.Field
		PollURL      respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GenerationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationGetResponseStatus string

const (
	GenerationGetResponseStatusPending   GenerationGetResponseStatus = "pending"
	GenerationGetResponseStatusRunning   GenerationGetResponseStatus = "running"
	GenerationGetResponseStatusCompleted GenerationGetResponseStatus = "completed"
	GenerationGetResponseStatusFailed    GenerationGetResponseStatus = "failed"
)

type GenerationGetResponseOutput struct {
	// Run output identifier
	OutputID string `json:"output_id" api:"required"`
	// Run output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Run output URL or text content
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputID    respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *GenerationGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationListResponse struct {
	CreatedAt float64 `json:"created_at" api:"required"`
	// Run identifier
	GenerationID string  `json:"generation_id" api:"required"`
	Progress     float64 `json:"progress" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status GenerationListResponseStatus `json:"status" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Cost charged in USD
	ChargedCost float64 `json:"charged_cost"`
	CompletedAt float64 `json:"completed_at"`
	// Machine-readable run error code
	ErrorCode string `json:"error_code"`
	// Human-readable run error message
	ErrorMessage string                         `json:"error_message"`
	Model        GenerationListResponseModel    `json:"model"`
	Outputs      []GenerationListResponseOutput `json:"outputs"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL   string  `json:"poll_url" format:"uri"`
	StartedAt float64 `json:"started_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		GenerationID respjson.Field
		Progress     respjson.Field
		ProjectID    respjson.Field
		RunID        respjson.Field
		Status       respjson.Field
		WorkspaceID  respjson.Field
		ChargedCost  respjson.Field
		CompletedAt  respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		Model        respjson.Field
		Outputs      respjson.Field
		PollURL      respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationListResponse) RawJSON() string { return r.JSON.raw }
func (r *GenerationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationListResponseStatus string

const (
	GenerationListResponseStatusPending   GenerationListResponseStatus = "pending"
	GenerationListResponseStatusRunning   GenerationListResponseStatus = "running"
	GenerationListResponseStatusCompleted GenerationListResponseStatus = "completed"
	GenerationListResponseStatusFailed    GenerationListResponseStatus = "failed"
)

type GenerationListResponseModel struct {
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
func (r GenerationListResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GenerationListResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationListResponseOutput struct {
	// Run output identifier
	OutputID string `json:"output_id" api:"required"`
	// Run output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Run output URL or text content
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutputID    respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationListResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *GenerationListResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationNewParams struct {
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_.
	ProjectID string `json:"project_id" api:"required"`
	// Generation prompt
	Prompt string `json:"prompt" api:"required"`
	// Generation type. Use "image", "video", "audio", or "text"; do not pass model
	// families such as "t2i" or "i2v".
	//
	// Any of "image", "video", "audio", "text".
	Type GenerationNewParamsType `json:"type,omitzero" api:"required"`
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

func (r GenerationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GenerationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generation type. Use "image", "video", "audio", or "text"; do not pass model
// families such as "t2i" or "i2v".
type GenerationNewParamsType string

const (
	GenerationNewParamsTypeImage GenerationNewParamsType = "image"
	GenerationNewParamsTypeVideo GenerationNewParamsType = "video"
	GenerationNewParamsTypeAudio GenerationNewParamsType = "audio"
	GenerationNewParamsTypeText  GenerationNewParamsType = "text"
)

type GenerationListParams struct {
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Project identifier
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	// Workspace identifier
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	// Run status filter
	//
	// Any of "pending", "running", "completed", "failed".
	Status GenerationListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GenerationListParams]'s query parameters as `url.Values`.
func (r GenerationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Run status filter
type GenerationListParamsStatus string

const (
	GenerationListParamsStatusPending   GenerationListParamsStatus = "pending"
	GenerationListParamsStatusRunning   GenerationListParamsStatus = "running"
	GenerationListParamsStatusCompleted GenerationListParamsStatus = "completed"
	GenerationListParamsStatusFailed    GenerationListParamsStatus = "failed"
)
