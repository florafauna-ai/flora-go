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

// Nested technique run endpoints.
//
// TechniqueRunService contains methods and other services that help with
// interacting with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTechniqueRunService] method instead.
type TechniqueRunService struct {
	options []option.RequestOption
}

// NewTechniqueRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTechniqueRunService(opts ...option.RequestOption) (r TechniqueRunService) {
	r = TechniqueRunService{}
	r.options = opts
	return
}

// Starts a run for a specific technique using the backward-compatible nested
// route. Mutating public API requests support an optional Idempotency-Key header
// for client retries; duplicate keys within two hours return
// idempotency_duplicate.
func (r *TechniqueRunService) New(ctx context.Context, techniqueID string, body TechniqueRunNewParams, opts ...option.RequestOption) (res *TechniqueRunNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if techniqueID == "" {
		err = errors.New("missing required techniqueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("techniques/%s/runs", url.PathEscape(techniqueID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists technique run history for the authenticated caller, including pending,
// running, completed, and failed technique runs. Results are newest first and can
// be filtered by workspace_id, project_id, technique_id, and status. Each item
// includes poll_url; use it to poll pending/running technique runs and to fetch
// completed or failed run details and outputs.
func (r *TechniqueRunService) List(ctx context.Context, query TechniqueRunListParams, opts ...option.RequestOption) (res *pagination.TechniqueRunsCursorPage[TechniqueRunListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "technique-runs"
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

// Lists technique run history for the authenticated caller, including pending,
// running, completed, and failed technique runs. Results are newest first and can
// be filtered by workspace_id, project_id, technique_id, and status. Each item
// includes poll_url; use it to poll pending/running technique runs and to fetch
// completed or failed run details and outputs.
func (r *TechniqueRunService) ListAutoPaging(ctx context.Context, query TechniqueRunListParams, opts ...option.RequestOption) *pagination.TechniqueRunsCursorPageAutoPager[TechniqueRunListResponse] {
	return pagination.NewTechniqueRunsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Returns status, progress, outputs, and error details for a technique run when it
// is accessible to the authenticated public API key.
func (r *TechniqueRunService) Get(ctx context.Context, runID string, query TechniqueRunGetParams, opts ...option.RequestOption) (res *TechniqueRunGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.TechniqueID == "" {
		err = errors.New("missing required techniqueId parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return nil, err
	}
	path := fmt.Sprintf("techniques/%s/runs/%s", url.PathEscape(query.TechniqueID), url.PathEscape(runID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TechniqueRunNewResponse struct {
	CreatedAt float64 `json:"created_at" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status TechniqueRunNewResponseStatus `json:"status" api:"required"`
	// Cost charged in USD
	ChargedCost float64 `json:"charged_cost"`
	CompletedAt float64 `json:"completed_at"`
	// Machine-readable run error code
	ErrorCode string `json:"error_code"`
	// Human-readable run error message
	ErrorMessage string                          `json:"error_message"`
	Outputs      []TechniqueRunNewResponseOutput `json:"outputs"`
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
func (r TechniqueRunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunNewResponseStatus string

const (
	TechniqueRunNewResponseStatusPending   TechniqueRunNewResponseStatus = "pending"
	TechniqueRunNewResponseStatusRunning   TechniqueRunNewResponseStatus = "running"
	TechniqueRunNewResponseStatusCompleted TechniqueRunNewResponseStatus = "completed"
	TechniqueRunNewResponseStatusFailed    TechniqueRunNewResponseStatus = "failed"
)

type TechniqueRunNewResponseOutput struct {
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
func (r TechniqueRunNewResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunNewResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunGetResponse struct {
	CreatedAt float64 `json:"created_at" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status TechniqueRunGetResponseStatus `json:"status" api:"required"`
	// Cost charged in USD
	ChargedCost float64 `json:"charged_cost"`
	CompletedAt float64 `json:"completed_at"`
	// Machine-readable run error code
	ErrorCode string `json:"error_code"`
	// Human-readable run error message
	ErrorMessage string                          `json:"error_message"`
	Outputs      []TechniqueRunGetResponseOutput `json:"outputs"`
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
func (r TechniqueRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunGetResponseStatus string

const (
	TechniqueRunGetResponseStatusPending   TechniqueRunGetResponseStatus = "pending"
	TechniqueRunGetResponseStatusRunning   TechniqueRunGetResponseStatus = "running"
	TechniqueRunGetResponseStatusCompleted TechniqueRunGetResponseStatus = "completed"
	TechniqueRunGetResponseStatusFailed    TechniqueRunGetResponseStatus = "failed"
)

type TechniqueRunGetResponseOutput struct {
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
func (r TechniqueRunGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunListResponse struct {
	CreatedAt float64 `json:"created_at" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status    TechniqueRunListResponseStatus    `json:"status" api:"required"`
	Technique TechniqueRunListResponseTechnique `json:"technique" api:"required"`
	// Run identifier
	TechniqueRunID string `json:"technique_run_id" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Cost charged in USD
	ChargedCost float64 `json:"charged_cost"`
	CompletedAt float64 `json:"completed_at"`
	// Machine-readable run error code
	ErrorCode string `json:"error_code"`
	// Human-readable run error message
	ErrorMessage string                           `json:"error_message"`
	Outputs      []TechniqueRunListResponseOutput `json:"outputs"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL   string  `json:"poll_url" format:"uri"`
	StartedAt float64 `json:"started_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		Progress       respjson.Field
		ProjectID      respjson.Field
		RunID          respjson.Field
		Status         respjson.Field
		Technique      respjson.Field
		TechniqueRunID respjson.Field
		WorkspaceID    respjson.Field
		ChargedCost    respjson.Field
		CompletedAt    respjson.Field
		ErrorCode      respjson.Field
		ErrorMessage   respjson.Field
		Outputs        respjson.Field
		PollURL        respjson.Field
		StartedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunListResponseStatus string

const (
	TechniqueRunListResponseStatusPending   TechniqueRunListResponseStatus = "pending"
	TechniqueRunListResponseStatusRunning   TechniqueRunListResponseStatus = "running"
	TechniqueRunListResponseStatusCompleted TechniqueRunListResponseStatus = "completed"
	TechniqueRunListResponseStatusFailed    TechniqueRunListResponseStatus = "failed"
)

type TechniqueRunListResponseTechnique struct {
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
func (r TechniqueRunListResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunListResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunListResponseOutput struct {
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
func (r TechniqueRunListResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunListResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunNewParams struct {
	// Technique inputs
	Inputs []TechniqueRunNewParamsInput `json:"inputs,omitzero" api:"required"`
	// Technique run execution mode
	//
	// Any of "async", "stream".
	Mode TechniqueRunNewParamsMode `json:"mode,omitzero" api:"required"`
	// HTTPS URL that receives a signed POST webhook when the run reaches a terminal
	// state (events run.completed / run.failed). The JSON body is HMAC-SHA256 signed
	// via the Flora-Signature header and delivery is retried up to 3 times with
	// exponential backoff. Must be HTTPS and must not resolve to a private/internal
	// host. See docs/system-overviews/webhooks.md for the payload schema and
	// verification example.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero" format:"uri"`
	// Idempotency key for safely retrying requests
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	paramObj
}

func (r TechniqueRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TechniqueRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TechniqueRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type, Value are required.
type TechniqueRunNewParamsInput struct {
	// Technique input identifier
	ID string `json:"id" api:"required"`
	// Technique input type
	//
	// Any of "text", "imageUrl", "videoUrl".
	Type string `json:"type,omitzero" api:"required"`
	// Technique input value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TechniqueRunNewParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow TechniqueRunNewParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TechniqueRunNewParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TechniqueRunNewParamsInput](
		"type", "text", "imageUrl", "videoUrl",
	)
}

// Technique run execution mode
type TechniqueRunNewParamsMode string

const (
	TechniqueRunNewParamsModeAsync  TechniqueRunNewParamsMode = "async"
	TechniqueRunNewParamsModeStream TechniqueRunNewParamsMode = "stream"
)

type TechniqueRunListParams struct {
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Project identifier
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	// Technique identifier
	TechniqueID param.Opt[string] `query:"technique_id,omitzero" json:"-"`
	// Workspace identifier
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	// Run status filter
	//
	// Any of "pending", "running", "completed", "failed".
	Status TechniqueRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TechniqueRunListParams]'s query parameters as `url.Values`.
func (r TechniqueRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Run status filter
type TechniqueRunListParamsStatus string

const (
	TechniqueRunListParamsStatusPending   TechniqueRunListParamsStatus = "pending"
	TechniqueRunListParamsStatusRunning   TechniqueRunListParamsStatus = "running"
	TechniqueRunListParamsStatusCompleted TechniqueRunListParamsStatus = "completed"
	TechniqueRunListParamsStatusFailed    TechniqueRunListParamsStatus = "failed"
)

type TechniqueRunGetParams struct {
	// Technique identifier or slug
	TechniqueID string `path:"techniqueId" api:"required" json:"-"`
	paramObj
}
