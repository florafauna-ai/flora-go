// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/florafauna-ai-go/internal/apijson"
	"github.com/stainless-sdks/florafauna-ai-go/internal/requestconfig"
	"github.com/stainless-sdks/florafauna-ai-go/option"
	"github.com/stainless-sdks/florafauna-ai-go/packages/param"
	"github.com/stainless-sdks/florafauna-ai-go/packages/respjson"
)

// FeedbackService contains methods and other services that help with interacting
// with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeedbackService] method instead.
type FeedbackService struct {
	options []option.RequestOption
}

// NewFeedbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeedbackService(opts ...option.RequestOption) (r FeedbackService) {
	r = FeedbackService{}
	r.options = opts
	return
}

// Records product feedback from the authenticated user, optionally linked to a
// workspace, project, run, and attempted tools. Mutating public API requests
// support an optional Idempotency-Key header for client retries; duplicate keys
// within two hours return idempotency_duplicate.
func (r *FeedbackService) Record(ctx context.Context, body FeedbackRecordParams, opts ...option.RequestOption) (res *FeedbackRecordResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FeedbackRecordResponse struct {
	// Feedback identifier
	FeedbackID string `json:"feedback_id" api:"required"`
	ReceivedAt int64  `json:"received_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeedbackID  respjson.Field
		ReceivedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackRecordResponse) RawJSON() string { return r.JSON.raw }
func (r *FeedbackRecordResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackRecordParams struct {
	// Detailed description
	Detail string `json:"detail" api:"required"`
	// Feedback kind
	//
	// Any of "feature_request", "bug", "technique_request", "missing_capability".
	Kind FeedbackRecordParamsKind `json:"kind,omitzero" api:"required"`
	// Short summary
	Summary string `json:"summary" api:"required"`
	// Project identifier
	ProjectID param.Opt[string] `json:"project_id,omitzero"`
	// Run identifier
	RunID param.Opt[string] `json:"run_id,omitzero"`
	// Workspace identifier
	WorkspaceID    param.Opt[string] `json:"workspace_id,omitzero"`
	AttemptedTools []string          `json:"attempted_tools,omitzero"`
	paramObj
}

func (r FeedbackRecordParams) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackRecordParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeedbackRecordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feedback kind
type FeedbackRecordParamsKind string

const (
	FeedbackRecordParamsKindFeatureRequest    FeedbackRecordParamsKind = "feature_request"
	FeedbackRecordParamsKindBug               FeedbackRecordParamsKind = "bug"
	FeedbackRecordParamsKindTechniqueRequest  FeedbackRecordParamsKind = "technique_request"
	FeedbackRecordParamsKindMissingCapability FeedbackRecordParamsKind = "missing_capability"
)
