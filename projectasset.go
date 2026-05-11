// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

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
	"github.com/stainless-sdks/florafauna-ai-go/packages/respjson"
)

// ProjectAssetService contains methods and other services that help with
// interacting with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectAssetService] method instead.
type ProjectAssetService struct {
	options []option.RequestOption
}

// NewProjectAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectAssetService(opts ...option.RequestOption) (r ProjectAssetService) {
	r = ProjectAssetService{}
	r.options = opts
	return
}

// Attaches an existing ready asset to a project canvas as a static media node.
// Mutating public API requests support an optional Idempotency-Key header for
// client retries; duplicate keys within two hours return idempotency_duplicate.
func (r *ProjectAssetService) AttachAsset(ctx context.Context, assetID string, body ProjectAssetAttachAssetParams, opts ...option.RequestOption) (res *ProjectAssetAttachAssetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ProjectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/assets/%s/attach", url.PathEscape(body.ProjectID), url.PathEscape(assetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ProjectAssetAttachAssetResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Project canvas URL
	CanvasURL string `json:"canvas_url" api:"required" format:"uri"`
	// Canvas node identifier
	NodeID string `json:"node_id" api:"required"`
	// Project identifier
	ProjectID string `json:"project_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID     respjson.Field
		CanvasURL   respjson.Field
		NodeID      respjson.Field
		ProjectID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectAssetAttachAssetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectAssetAttachAssetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectAssetAttachAssetParams struct {
	// Project identifier
	ProjectID string `path:"projectId" api:"required" json:"-"`
	paramObj
}
