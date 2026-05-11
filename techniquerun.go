// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package florafaunaai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/florafauna-ai-go/internal/apijson"
	"github.com/stainless-sdks/florafauna-ai-go/internal/requestconfig"
	"github.com/stainless-sdks/florafauna-ai-go/option"
	"github.com/stainless-sdks/florafauna-ai-go/packages/param"
	"github.com/stainless-sdks/florafauna-ai-go/packages/respjson"
)

// TechniqueRunService contains methods and other services that help with
// interacting with the florafauna-ai API.
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

// Starts a run for a specific technique using the backward-compatible nested
// route. Mutating public API requests support an optional Idempotency-Key header
// for client retries; duplicate keys within two hours return
// idempotency_duplicate.
func (r *TechniqueRunService) Start(ctx context.Context, techniqueID string, body TechniqueRunStartParams, opts ...option.RequestOption) (res *TechniqueRunStartResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if techniqueID == "" {
		err = errors.New("missing required techniqueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("techniques/%s/runs", url.PathEscape(techniqueID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TechniqueRunGetResponse struct {
	CreatedAt float64 `json:"createdAt" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Run identifier
	RunID string `json:"runId" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status      TechniqueRunGetResponseStatus `json:"status" api:"required"`
	ChargedCost float64                       `json:"chargedCost"`
	CompletedAt float64                       `json:"completedAt"`
	// Machine-readable run error code
	ErrorCode string `json:"errorCode"`
	// Human-readable run error message
	ErrorMessage string                          `json:"errorMessage"`
	Outputs      []TechniqueRunGetResponseOutput `json:"outputs"`
	PollURL      string                          `json:"pollUrl" format:"uri"`
	StartedAt    float64                         `json:"startedAt"`
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
	OutputID string `json:"outputId" api:"required"`
	// Run output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Run output URL
	URL string `json:"url" api:"required" format:"uri"`
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

type TechniqueRunStartResponse struct {
	CreatedAt float64 `json:"createdAt" api:"required"`
	Progress  float64 `json:"progress" api:"required"`
	// Run identifier
	RunID string `json:"runId" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status      TechniqueRunStartResponseStatus `json:"status" api:"required"`
	ChargedCost float64                         `json:"chargedCost"`
	CompletedAt float64                         `json:"completedAt"`
	// Machine-readable run error code
	ErrorCode string `json:"errorCode"`
	// Human-readable run error message
	ErrorMessage string                            `json:"errorMessage"`
	Outputs      []TechniqueRunStartResponseOutput `json:"outputs"`
	PollURL      string                            `json:"pollUrl" format:"uri"`
	StartedAt    float64                           `json:"startedAt"`
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
func (r TechniqueRunStartResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunStartResponseStatus string

const (
	TechniqueRunStartResponseStatusPending   TechniqueRunStartResponseStatus = "pending"
	TechniqueRunStartResponseStatusRunning   TechniqueRunStartResponseStatus = "running"
	TechniqueRunStartResponseStatusCompleted TechniqueRunStartResponseStatus = "completed"
	TechniqueRunStartResponseStatusFailed    TechniqueRunStartResponseStatus = "failed"
)

type TechniqueRunStartResponseOutput struct {
	// Run output identifier
	OutputID string `json:"outputId" api:"required"`
	// Run output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Run output URL
	URL string `json:"url" api:"required" format:"uri"`
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
func (r TechniqueRunStartResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunStartResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunGetParams struct {
	// Technique identifier or slug
	TechniqueID string `path:"techniqueId" api:"required" json:"-"`
	paramObj
}

type TechniqueRunStartParams struct {
	Inputs []TechniqueRunStartParamsInput `json:"inputs,omitzero" api:"required"`
	// Any of "async", "stream".
	Mode        TechniqueRunStartParamsMode `json:"mode,omitzero" api:"required"`
	CallbackURL param.Opt[string]           `json:"callback_url,omitzero" format:"uri"`
	// Idempotency key for safely retrying requests
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	paramObj
}

func (r TechniqueRunStartParams) MarshalJSON() (data []byte, err error) {
	type shadow TechniqueRunStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TechniqueRunStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type, Value are required.
type TechniqueRunStartParamsInput struct {
	// Technique input identifier
	ID string `json:"id" api:"required"`
	// Technique input type
	//
	// Any of "imageUrl", "videoUrl", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Technique input value
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TechniqueRunStartParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow TechniqueRunStartParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TechniqueRunStartParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TechniqueRunStartParamsInput](
		"type", "imageUrl", "videoUrl", "text",
	)
}

type TechniqueRunStartParamsMode string

const (
	TechniqueRunStartParamsModeAsync  TechniqueRunStartParamsMode = "async"
	TechniqueRunStartParamsModeStream TechniqueRunStartParamsMode = "stream"
)
