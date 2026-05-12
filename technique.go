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

// TechniqueService contains methods and other services that help with interacting
// with the flora API.
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
func (r *TechniqueService) List(ctx context.Context, query TechniqueListParams, opts ...option.RequestOption) (res *pagination.TechniquesCursorPage[TechniqueListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "techniques"
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

// Returns reusable Flora techniques visible to the authenticated public API key.
// Use workspace_id, query, cursor, and limit to filter the catalog.
func (r *TechniqueService) ListAutoPaging(ctx context.Context, query TechniqueListParams, opts ...option.RequestOption) *pagination.TechniquesCursorPageAutoPager[TechniqueListResponse] {
	return pagination.NewTechniquesCursorPageAutoPager(r.List(ctx, query, opts...))
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
	Inputs []TechniqueListResponseInput `json:"inputs" api:"required"`
	// Technique name
	Name    string                        `json:"name" api:"required"`
	Outputs []TechniqueListResponseOutput `json:"outputs" api:"required"`
	RunCost float64                       `json:"run_cost" api:"required"`
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
func (r TechniqueListResponse) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseInput struct {
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
func (r TechniqueListResponseInput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueListResponseOutput struct {
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
func (r TechniqueListResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *TechniqueListResponseOutput) UnmarshalJSON(data []byte) error {
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
