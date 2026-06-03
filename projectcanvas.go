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
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

// Project management endpoints.
//
// ProjectCanvasService contains methods and other services that help with
// interacting with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectCanvasService] method instead.
type ProjectCanvasService struct {
	options []option.RequestOption
}

// NewProjectCanvasService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectCanvasService(opts ...option.RequestOption) (r ProjectCanvasService) {
	r = ProjectCanvasService{}
	r.options = opts
	return
}

// Returns the current project canvas topology as a Mermaid flowchart using the
// same serializer as the Fauna agent.
func (r *ProjectCanvasService) Get(ctx context.Context, projectID string, opts ...option.RequestOption) (res *ProjectCanvasGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/canvas", url.PathEscape(projectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Applies a Mermaid flowchart patch to the project canvas using the same
// create_workflow path as the Fauna agent. The diagram may add nodes, connect
// nodes, and reference existing canvas nodes by their Mermaid short IDs in edges
// (e.g. `n1 --> out`). This endpoint is add-only: re-declaring an existing node id
// with a label (e.g. `n3["..."]`) creates a NEW node instead of updating the
// existing one, and returns a warning. To attach to an existing node, reference
// its id in an edge without re-declaring its label. Subgraph grouping is not
// applied (nodes inside a `subgraph` are added ungrouped) and returns a warning.
// To place an existing image/video/audio as a static node, set `node_params` —
// which is keyed by Mermaid node id, e.g.
// `{ "img1": { "content_url": "https://…" } }`, NOT a bare `{ content_url }`
// object. `prompt` and `content_url` are mutually exclusive for a node: use
// `prompt` (or a label that doubles as the prompt) for generation, or
// `content_url` for existing media. When using `content_url`, give the node a
// content-free type-only label such as `img1["(Image)"]` so no prompt is inferred
// from the label.
func (r *ProjectCanvasService) Update(ctx context.Context, projectID string, body ProjectCanvasUpdateParams, opts ...option.RequestOption) (res *ProjectCanvasUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return nil, err
	}
	path := fmt.Sprintf("projects/%s/canvas", url.PathEscape(projectID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ProjectCanvasGetResponse struct {
	// Project canvas URL
	CanvasURL string `json:"canvas_url" api:"required" format:"uri"`
	// Mermaid flowchart diagram
	Diagram string `json:"diagram" api:"required"`
	// Project identifier
	ProjectID string                          `json:"project_id" api:"required"`
	Summary   ProjectCanvasGetResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanvasURL   respjson.Field
		Diagram     respjson.Field
		ProjectID   respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectCanvasGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectCanvasGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectCanvasGetResponseSummary struct {
	EdgeCount         int64 `json:"edge_count" api:"required"`
	GroupCount        int64 `json:"group_count" api:"required"`
	IsolatedNodeCount int64 `json:"isolated_node_count" api:"required"`
	NodeCount         int64 `json:"node_count" api:"required"`
	WorkflowCount     int64 `json:"workflow_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EdgeCount         respjson.Field
		GroupCount        respjson.Field
		IsolatedNodeCount respjson.Field
		NodeCount         respjson.Field
		WorkflowCount     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectCanvasGetResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *ProjectCanvasGetResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectCanvasUpdateResponse struct {
	// Project canvas URL
	CanvasURL        string `json:"canvas_url" api:"required" format:"uri"`
	CreatedEdgeCount int64  `json:"created_edge_count" api:"required"`
	CreatedNodeCount int64  `json:"created_node_count" api:"required"`
	// Applied Mermaid flowchart diagram
	Diagram string `json:"diagram" api:"required"`
	// Project identifier
	ProjectID string   `json:"project_id" api:"required"`
	Warnings  []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanvasURL        respjson.Field
		CreatedEdgeCount respjson.Field
		CreatedNodeCount respjson.Field
		Diagram          respjson.Field
		ProjectID        respjson.Field
		Warnings         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectCanvasUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectCanvasUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectCanvasUpdateParams struct {
	// Mermaid flowchart diagram to apply
	Diagram string `json:"diagram" api:"required"`
	// Optional per-node parameters, keyed by Mermaid node id (a Record<nodeId,
	// NodeParams>), e.g. { "img1": { "content_url": "https://…" } }. Pass a map keyed
	// by node id, NOT a bare { content_url } object.
	NodeParams map[string]ProjectCanvasUpdateParamsNodeParam `json:"node_params,omitzero"`
	paramObj
}

func (r ProjectCanvasUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectCanvasUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectCanvasUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectCanvasUpdateParamsNodeParam struct {
	AspectRatio param.Opt[string] `json:"aspect_ratio,omitzero"`
	// HTTPS URL of existing media to place as a static node. Mutually exclusive with
	// prompt: give the node a content-free type-only label such as `(Image)` so no
	// prompt is inferred from the label. Only supported for Image, Video, and Audio
	// nodes.
	ContentURL param.Opt[string] `json:"content_url,omitzero" format:"uri"`
	Model      param.Opt[string] `json:"model,omitzero"`
	// Generation prompt for this node. Mutually exclusive with content_url. If
	// omitted, the node's Mermaid label is used as the prompt.
	Prompt          param.Opt[string] `json:"prompt,omitzero"`
	Resolution      param.Opt[string] `json:"resolution,omitzero"`
	ModelParameters map[string]any    `json:"model_parameters,omitzero"`
	paramObj
}

func (r ProjectCanvasUpdateParamsNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectCanvasUpdateParamsNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectCanvasUpdateParamsNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
