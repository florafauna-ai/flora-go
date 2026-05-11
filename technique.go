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
	"github.com/stainless-sdks/florafauna-ai-go/internal/apiquery"
	"github.com/stainless-sdks/florafauna-ai-go/internal/requestconfig"
	"github.com/stainless-sdks/florafauna-ai-go/option"
	"github.com/stainless-sdks/florafauna-ai-go/packages/param"
	"github.com/stainless-sdks/florafauna-ai-go/packages/respjson"
)

// TechniqueService contains methods and other services that help with interacting
// with the florafauna-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTechniqueService] method instead.
type TechniqueService struct {
	options []option.RequestOption
	Runs    TechniqueRunService
}

// NewTechniqueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTechniqueService(opts ...option.RequestOption) (r TechniqueService) {
	r = TechniqueService{}
	r.options = opts
	r.Runs = NewTechniqueRunService(opts...)
	return
}

// Returns the public definition for one technique, including its input and output
// schema used to start runs.
func (r *TechniqueService) Get(ctx context.Context, techniqueID string, opts ...option.RequestOption) (res *TechniqueGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if techniqueID == "" {
		err = errors.New("missing required techniqueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("techniques/%s", url.PathEscape(techniqueID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns reusable Flora techniques visible to the authenticated public API key.
// Use workspace_id, query, cursor, and limit to filter the catalog.
func (r *TechniqueService) List(ctx context.Context, query TechniqueListParams, opts ...option.RequestOption) (res *TechniqueListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "techniques"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TechniqueGetResponse struct {
	Inputs []TechniqueGetResponseInput `json:"inputs" api:"required"`
	// Technique name
	Name    string                       `json:"name" api:"required"`
	Outputs []TechniqueGetResponseOutput `json:"outputs" api:"required"`
	RunCost float64                      `json:"run_cost" api:"required"`
	// Technique identifier
	TechniqueID string `json:"technique_id" api:"required"`
	// Technique description
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inputs      respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		RunCost     respjson.Field
		TechniqueID respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueGetResponseInput struct {
	// Technique input or output identifier
	ID string `json:"id" api:"required"`
	// Technique input or output display name
	Name string `json:"name" api:"required"`
	// Technique input or output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Technique input or output description
	Description string `json:"description"`
	// Required aspect ratio
	SpecifiedAspectRatio string `json:"specifiedAspectRatio"`
	// Required duration in seconds
	SpecifiedDuration float64 `json:"specifiedDuration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Type                 respjson.Field
		Description          respjson.Field
		SpecifiedAspectRatio respjson.Field
		SpecifiedDuration    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueGetResponseOutput struct {
	// Technique input or output identifier
	ID string `json:"id" api:"required"`
	// Technique input or output display name
	Name string `json:"name" api:"required"`
	// Technique input or output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Technique input or output description
	Description string `json:"description"`
	// Required aspect ratio
	SpecifiedAspectRatio string `json:"specifiedAspectRatio"`
	// Required duration in seconds
	SpecifiedDuration float64 `json:"specifiedDuration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Type                 respjson.Field
		Description          respjson.Field
		SpecifiedAspectRatio respjson.Field
		SpecifiedDuration    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponse struct {
	Meta       TechniqueListResponseMeta        `json:"meta" api:"required"`
	Techniques []TechniqueListResponseTechnique `json:"techniques" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		Techniques  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueListResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseMeta struct {
	// Opaque cursor for fetching the next page
	NextCursor string `json:"next_cursor" api:"required"`
	// Estimated total matching items
	TotalEstimate int64 `json:"total_estimate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor    respjson.Field
		TotalEstimate respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseTechnique struct {
	Inputs []TechniqueListResponseTechniqueInput `json:"inputs" api:"required"`
	// Technique name
	Name    string                                 `json:"name" api:"required"`
	Outputs []TechniqueListResponseTechniqueOutput `json:"outputs" api:"required"`
	RunCost float64                                `json:"run_cost" api:"required"`
	// Technique identifier
	TechniqueID string `json:"technique_id" api:"required"`
	// Technique description
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inputs      respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		RunCost     respjson.Field
		TechniqueID respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueListResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseTechniqueInput struct {
	// Technique input or output identifier
	ID string `json:"id" api:"required"`
	// Technique input or output display name
	Name string `json:"name" api:"required"`
	// Technique input or output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Technique input or output description
	Description string `json:"description"`
	// Required aspect ratio
	SpecifiedAspectRatio string `json:"specifiedAspectRatio"`
	// Required duration in seconds
	SpecifiedDuration float64 `json:"specifiedDuration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Type                 respjson.Field
		Description          respjson.Field
		SpecifiedAspectRatio respjson.Field
		SpecifiedDuration    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueListResponseTechniqueInput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseTechniqueInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseTechniqueOutput struct {
	// Technique input or output identifier
	ID string `json:"id" api:"required"`
	// Technique input or output display name
	Name string `json:"name" api:"required"`
	// Technique input or output media type
	//
	// Any of "imageUrl", "videoUrl", "audioUrl", "text", "documentUrl".
	Type string `json:"type" api:"required"`
	// Technique input or output description
	Description string `json:"description"`
	// Required aspect ratio
	SpecifiedAspectRatio string `json:"specifiedAspectRatio"`
	// Required duration in seconds
	SpecifiedDuration float64 `json:"specifiedDuration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		Type                 respjson.Field
		Description          respjson.Field
		SpecifiedAspectRatio respjson.Field
		SpecifiedDuration    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueListResponseTechniqueOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseTechniqueOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListParams struct {
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Workspace identifier
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TechniqueListParams]'s query parameters as `url.Values`.
func (r TechniqueListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
