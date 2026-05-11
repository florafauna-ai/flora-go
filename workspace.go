// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/florafauna-ai-go/internal/apijson"
	"github.com/stainless-sdks/florafauna-ai-go/internal/requestconfig"
	"github.com/stainless-sdks/florafauna-ai-go/option"
	"github.com/stainless-sdks/florafauna-ai-go/packages/respjson"
)

// WorkspaceService contains methods and other services that help with interacting
// with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	options []option.RequestOption
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r WorkspaceService) {
	r = WorkspaceService{}
	r.options = opts
	return
}

// Returns the workspaces available to the authenticated public API key, including
// each workspace's public ID, name, creation timestamp, and caller role.
func (r *WorkspaceService) List(ctx context.Context, opts ...option.RequestOption) (res *WorkspaceListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WorkspaceListResponse struct {
	Workspaces []WorkspaceListResponseWorkspace `json:"workspaces" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workspaces  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceListResponseWorkspace struct {
	CreatedAt int64 `json:"created_at" api:"required"`
	// Workspace name
	Name string `json:"name" api:"required"`
	// Workspace role
	Role string `json:"role" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceListResponseWorkspace) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceListResponseWorkspace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
