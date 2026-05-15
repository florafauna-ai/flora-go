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

// ProjectService contains methods and other services that help with interacting
// with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	options []option.RequestOption
	// Project canvas endpoints.
	Assets ProjectAssetService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.options = opts
	r.Assets = NewProjectAssetService(opts...)
	return
}

// Creates a new Flora project in the requested workspace. Mutating public API
// requests support an optional Idempotency-Key header for client retries;
// duplicate keys within two hours return idempotency_duplicate.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns metadata for a single project when it is accessible to the authenticated
// public API key. Missing and inaccessible projects both return 404.
func (r *ProjectService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s", url.PathEscape(projectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns projects in the requested workspace that are accessible to the
// authenticated public API key, ordered by recent activity.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.ProjectsCursorPage[ProjectListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "projects"
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

// Returns projects in the requested workspace that are accessible to the
// authenticated public API key, ordered by recent activity.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.ProjectsCursorPageAutoPager[ProjectListResponse] {
	return pagination.NewProjectsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Returns sanitized visible media nodes on a project canvas. The response omits
// raw graph documents, Liveblocks internals, raw Convex IDs, and unbounded node
// data blobs.
func (r *ProjectService) ListNodes(ctx context.Context, projectID string, query ProjectListNodesParams, opts ...option.RequestOption) (res *pagination.CanvasNodesCursorPage[ProjectListNodesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/nodes", url.PathEscape(projectID))
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

// Returns sanitized visible media nodes on a project canvas. The response omits
// raw graph documents, Liveblocks internals, raw Convex IDs, and unbounded node
// data blobs.
func (r *ProjectService) ListNodesAutoPaging(ctx context.Context, projectID string, query ProjectListNodesParams, opts ...option.RequestOption) *pagination.CanvasNodesCursorPageAutoPager[ProjectListNodesResponse] {
	return pagination.NewCanvasNodesCursorPageAutoPager(r.ListNodes(ctx, projectID, query, opts...))
}

type ProjectNewResponse struct {
	CreatedAt    float64 `json:"created_at" api:"required"`
	LastModified float64 `json:"last_modified" api:"required"`
	// Project name
	Name string `json:"name" api:"required"`
	// Project origin
	Origin string `json:"origin" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		LastModified respjson.Field
		Name         respjson.Field
		Origin       respjson.Field
		ProjectID    respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectGetResponse struct {
	CreatedAt    float64 `json:"created_at" api:"required"`
	LastModified float64 `json:"last_modified" api:"required"`
	// Project name
	Name string `json:"name" api:"required"`
	// Project origin
	Origin string `json:"origin" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		LastModified respjson.Field
		Name         respjson.Field
		Origin       respjson.Field
		ProjectID    respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponse struct {
	CreatedAt    float64 `json:"created_at" api:"required"`
	LastModified float64 `json:"last_modified" api:"required"`
	// Project name
	Name string `json:"name" api:"required"`
	// Project origin
	Origin string `json:"origin" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		LastModified respjson.Field
		Name         respjson.Field
		Origin       respjson.Field
		ProjectID    respjson.Field
		WorkspaceID  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListNodesResponse struct {
	// Canvas node identifier
	NodeID string `json:"node_id" api:"required"`
	// Canvas node media type
	//
	// Any of "image", "video", "audio", "text".
	Type ProjectListNodesResponseType `json:"type" api:"required"`
	// Asset identifier
	AssetID string `json:"asset_id" api:"nullable"`
	Height  int64  `json:"height" api:"nullable"`
	// Canvas node label
	Label string `json:"label" api:"nullable"`
	// Canvas node output URL or text content
	URL   string `json:"url" api:"nullable"`
	Width int64  `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NodeID      respjson.Field
		Type        respjson.Field
		AssetID     respjson.Field
		Height      respjson.Field
		Label       respjson.Field
		URL         respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListNodesResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListNodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canvas node media type
type ProjectListNodesResponseType string

const (
	ProjectListNodesResponseTypeImage ProjectListNodesResponseType = "image"
	ProjectListNodesResponseTypeVideo ProjectListNodesResponseType = "video"
	ProjectListNodesResponseTypeAudio ProjectListNodesResponseType = "audio"
	ProjectListNodesResponseTypeText  ProjectListNodesResponseType = "text"
)

type ProjectNewParams struct {
	// Project name
	Name string `json:"name" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
	// Workspace identifier
	WorkspaceID string `query:"workspace_id" api:"required" json:"-"`
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ProjectListNodesParams struct {
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListNodesParams]'s query parameters as `url.Values`.
func (r ProjectListNodesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
