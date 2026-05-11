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

// ProjectService contains methods and other services that help with interacting
// with the florafauna-ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	options []option.RequestOption
	Assets  ProjectAssetService
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
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns sanitized visible media nodes on a project canvas. The response omits
// raw graph documents, Liveblocks internals, raw Convex IDs, and unbounded node
// data blobs.
func (r *ProjectService) ListNodes(ctx context.Context, projectID string, query ProjectListNodesParams, opts ...option.RequestOption) (res *ProjectListNodesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/nodes", url.PathEscape(projectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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
	Meta     ProjectListResponseMeta      `json:"meta" api:"required"`
	Projects []ProjectListResponseProject `json:"projects" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta        respjson.Field
		Projects    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseMeta struct {
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
func (r ProjectListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponseProject struct {
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
func (r ProjectListResponseProject) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponseProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListNodesResponse struct {
	// Project canvas URL
	CanvasURL string                         `json:"canvas_url" api:"required" format:"uri"`
	Meta      ProjectListNodesResponseMeta   `json:"meta" api:"required"`
	Nodes     []ProjectListNodesResponseNode `json:"nodes" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanvasURL   respjson.Field
		Meta        respjson.Field
		Nodes       respjson.Field
		ProjectID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectListNodesResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListNodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListNodesResponseMeta struct {
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
func (r ProjectListNodesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ProjectListNodesResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListNodesResponseNode struct {
	// Canvas node identifier
	NodeID string `json:"node_id" api:"required"`
	// Canvas node media type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type" api:"required"`
	// Asset identifier
	AssetID string `json:"asset_id" api:"nullable"`
	Height  int64  `json:"height" api:"nullable"`
	// Canvas node label
	Label string `json:"label" api:"nullable"`
	// Canvas node media URL
	URL   string `json:"url" api:"nullable" format:"uri"`
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
func (r ProjectListNodesResponseNode) RawJSON() string { return r.JSON.raw }
func (r *ProjectListNodesResponseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
