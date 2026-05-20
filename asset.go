// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/apiquery"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/pagination"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
	"github.com/florafauna-ai/flora-go/shared/constant"
)

// Asset upload and retrieval endpoints.
//
// AssetService contains methods and other services that help with interacting with
// the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.options = opts
	return
}

// Creates an asset from an allowlisted source URL or reserves a signed upload URL.
// Mutating public API requests support an optional Idempotency-Key header for
// client retries; duplicate keys within two hours return idempotency_duplicate.
func (r *AssetService) New(ctx context.Context, body AssetNewParams, opts ...option.RequestOption) (res *AssetNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns metadata for one asset when it is accessible to the authenticated public
// API key. Missing and inaccessible assets both return 404.
func (r *AssetService) Get(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s", url.PathEscape(assetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns assets visible to the authenticated public API key. Filter by workspace,
// project canvas, search query, cursor, and limit without exposing raw file bytes
// or internal graph data.
func (r *AssetService) List(ctx context.Context, query AssetListParams, opts ...option.RequestOption) (res *pagination.AssetsCursorPage[AssetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "assets"
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

// Returns assets visible to the authenticated public API key. Filter by workspace,
// project canvas, search query, cursor, and limit without exposing raw file bytes
// or internal graph data.
func (r *AssetService) ListAutoPaging(ctx context.Context, query AssetListParams, opts ...option.RequestOption) *pagination.AssetsCursorPageAutoPager[AssetListResponse] {
	return pagination.NewAssetsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Marks a signed asset upload as complete after the file has been uploaded.
// Mutating public API requests support an optional Idempotency-Key header for
// client retries; duplicate keys within two hours return idempotency_duplicate.
func (r *AssetService) Complete(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetCompleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/complete", url.PathEscape(assetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Creates a fresh signed upload reservation for a failed or expired asset upload.
// Mutating public API requests support an optional Idempotency-Key header for
// client retries; duplicate keys within two hours return idempotency_duplicate.
func (r *AssetService) Retry(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetRetryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if assetID == "" {
		err = errors.New("missing required assetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("assets/%s/retry", url.PathEscape(assetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AssetNewResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Any of "pending_upload", "ready", "failed".
	Status AssetNewResponseStatus `json:"status" api:"required"`
	// Asset source
	//
	// Any of "url", "signed_url".
	UploadedVia AssetNewResponseUploadedVia `json:"uploaded_via" api:"required"`
	// Asset URL
	URL        string             `json:"url" api:"required" format:"uri"`
	Visibility constant.Workspace `json:"visibility" default:"workspace"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Expiration time for the upload URL
	ExpiresAt time.Time              `json:"expires_at" format:"date-time"`
	Upload    AssetNewResponseUpload `json:"upload"`
	// Upload URL (serialized)
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID     respjson.Field
		Status      respjson.Field
		UploadedVia respjson.Field
		URL         respjson.Field
		Visibility  respjson.Field
		WorkspaceID respjson.Field
		ExpiresAt   respjson.Field
		Upload      respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewResponseStatus string

const (
	AssetNewResponseStatusPendingUpload AssetNewResponseStatus = "pending_upload"
	AssetNewResponseStatusReady         AssetNewResponseStatus = "ready"
	AssetNewResponseStatusFailed        AssetNewResponseStatus = "failed"
)

// Asset source
type AssetNewResponseUploadedVia string

const (
	AssetNewResponseUploadedViaURL       AssetNewResponseUploadedVia = "url"
	AssetNewResponseUploadedViaSignedURL AssetNewResponseUploadedVia = "signed_url"
)

type AssetNewResponseUpload struct {
	ContentType constant.MultipartFormData `json:"content_type" default:"multipart/form-data"`
	FileField   constant.File              `json:"file_field" default:"file"`
	// Upload form fields
	FormFields map[string]string `json:"form_fields" api:"required"`
	Method     constant.Post     `json:"method" default:"POST"`
	// Upload URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		FileField   respjson.Field
		FormFields  respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetNewResponseUpload) RawJSON() string { return r.JSON.raw }
func (r *AssetNewResponseUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetGetResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Asset content type
	ContentType string `json:"content_type" api:"required"`
	// Asset creation time (ISO 8601 datetime)
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Asset description
	Description string `json:"description" api:"required"`
	// Failure message when the asset is in failed status
	FailureMessage string `json:"failure_message" api:"required"`
	Height         int64  `json:"height" api:"required"`
	// Asset name
	Name      string `json:"name" api:"required"`
	SizeBytes int64  `json:"size_bytes" api:"required"`
	// Any of "pending_upload", "ready", "failed".
	Status AssetGetResponseStatus `json:"status" api:"required"`
	// Content type provided at upload time
	UploadContentType string `json:"upload_content_type" api:"required"`
	// Asset source
	UploadedVia string `json:"uploaded_via" api:"required"`
	// Asset URL
	URL   string `json:"url" api:"required" format:"uri"`
	Width int64  `json:"width" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID           respjson.Field
		ContentType       respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		FailureMessage    respjson.Field
		Height            respjson.Field
		Name              respjson.Field
		SizeBytes         respjson.Field
		Status            respjson.Field
		UploadContentType respjson.Field
		UploadedVia       respjson.Field
		URL               respjson.Field
		Width             respjson.Field
		WorkspaceID       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetGetResponseStatus string

const (
	AssetGetResponseStatusPendingUpload AssetGetResponseStatus = "pending_upload"
	AssetGetResponseStatusReady         AssetGetResponseStatus = "ready"
	AssetGetResponseStatusFailed        AssetGetResponseStatus = "failed"
)

type AssetListResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Asset content type
	ContentType string `json:"content_type" api:"required"`
	CreatedAt   int64  `json:"created_at" api:"required"`
	// Asset description
	Description string `json:"description" api:"required"`
	Height      int64  `json:"height" api:"required"`
	// Asset name
	Name      string `json:"name" api:"required"`
	SizeBytes int64  `json:"size_bytes" api:"required"`
	// Any of "pending_upload", "ready", "failed".
	Status AssetListResponseStatus `json:"status" api:"required"`
	// Content type provided at upload time
	UploadContentType string `json:"upload_content_type" api:"required"`
	// Asset source
	UploadedVia string `json:"uploaded_via" api:"required"`
	// Asset URL
	URL   string `json:"url" api:"required" format:"uri"`
	Width int64  `json:"width" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Associated node identifier
	NodeID string `json:"node_id" api:"nullable"`
	// Project identifier
	ProjectID string `json:"project_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID           respjson.Field
		ContentType       respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		Height            respjson.Field
		Name              respjson.Field
		SizeBytes         respjson.Field
		Status            respjson.Field
		UploadContentType respjson.Field
		UploadedVia       respjson.Field
		URL               respjson.Field
		Width             respjson.Field
		WorkspaceID       respjson.Field
		NodeID            respjson.Field
		ProjectID         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetListResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetListResponseStatus string

const (
	AssetListResponseStatusPendingUpload AssetListResponseStatus = "pending_upload"
	AssetListResponseStatusReady         AssetListResponseStatus = "ready"
	AssetListResponseStatusFailed        AssetListResponseStatus = "failed"
)

type AssetCompleteResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Any of "pending_upload", "ready", "failed".
	Status AssetCompleteResponseStatus `json:"status" api:"required"`
	// Asset URL
	URL        string             `json:"url" api:"required" format:"uri"`
	Visibility constant.Workspace `json:"visibility" default:"workspace"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Expiration time for the upload URL
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Failure message when the asset is in failed status
	FailureMessage string                      `json:"failure_message" api:"nullable"`
	Upload         AssetCompleteResponseUpload `json:"upload"`
	// Upload URL (serialized)
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID        respjson.Field
		Status         respjson.Field
		URL            respjson.Field
		Visibility     respjson.Field
		WorkspaceID    respjson.Field
		ExpiresAt      respjson.Field
		FailureMessage respjson.Field
		Upload         respjson.Field
		UploadURL      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetCompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetCompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetCompleteResponseStatus string

const (
	AssetCompleteResponseStatusPendingUpload AssetCompleteResponseStatus = "pending_upload"
	AssetCompleteResponseStatusReady         AssetCompleteResponseStatus = "ready"
	AssetCompleteResponseStatusFailed        AssetCompleteResponseStatus = "failed"
)

type AssetCompleteResponseUpload struct {
	ContentType constant.MultipartFormData `json:"content_type" default:"multipart/form-data"`
	FileField   constant.File              `json:"file_field" default:"file"`
	// Upload form fields
	FormFields map[string]string `json:"form_fields" api:"required"`
	Method     constant.Post     `json:"method" default:"POST"`
	// Upload URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		FileField   respjson.Field
		FormFields  respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetCompleteResponseUpload) RawJSON() string { return r.JSON.raw }
func (r *AssetCompleteResponseUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetRetryResponse struct {
	// Asset identifier
	AssetID string `json:"asset_id" api:"required"`
	// Any of "pending_upload", "ready", "failed".
	Status AssetRetryResponseStatus `json:"status" api:"required"`
	// Asset URL
	URL        string             `json:"url" api:"required" format:"uri"`
	Visibility constant.Workspace `json:"visibility" default:"workspace"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Expiration time for the upload URL
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Failure message when the asset is in failed status
	FailureMessage string                   `json:"failure_message" api:"nullable"`
	Upload         AssetRetryResponseUpload `json:"upload"`
	// Upload URL (serialized)
	UploadURL string `json:"upload_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetID        respjson.Field
		Status         respjson.Field
		URL            respjson.Field
		Visibility     respjson.Field
		WorkspaceID    respjson.Field
		ExpiresAt      respjson.Field
		FailureMessage respjson.Field
		Upload         respjson.Field
		UploadURL      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetRetryResponseStatus string

const (
	AssetRetryResponseStatusPendingUpload AssetRetryResponseStatus = "pending_upload"
	AssetRetryResponseStatusReady         AssetRetryResponseStatus = "ready"
	AssetRetryResponseStatusFailed        AssetRetryResponseStatus = "failed"
)

type AssetRetryResponseUpload struct {
	ContentType constant.MultipartFormData `json:"content_type" default:"multipart/form-data"`
	FileField   constant.File              `json:"file_field" default:"file"`
	// Upload form fields
	FormFields map[string]string `json:"form_fields" api:"required"`
	Method     constant.Post     `json:"method" default:"POST"`
	// Upload URL
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		FileField   respjson.Field
		FormFields  respjson.Field
		Method      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetRetryResponseUpload) RawJSON() string { return r.JSON.raw }
func (r *AssetRetryResponseUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewParams struct {
	// Asset source URL or signed-url upload mode
	Source string `json:"source" api:"required"`
	// Workspace identifier
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Asset content type
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// Asset file name
	FileName param.Opt[string] `json:"file_name,omitzero"`
	// Destination folder
	Folder param.Opt[string] `json:"folder,omitzero"`
	paramObj
}

func (r AssetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetListParams struct {
	// Opaque cursor for fetching the next page
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Project identifier
	ProjectID param.Opt[string] `query:"project_id,omitzero" json:"-"`
	// Search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Workspace identifier
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
