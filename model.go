// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/apiquery"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.options = opts
	return
}

// Returns the public model catalog visible to API clients. Use the optional type
// filter to narrow results to image, video, audio, or text models.
func (r *ModelService) List(ctx context.Context, query ModelListParams, opts ...option.RequestOption) (res *ModelListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ModelListResponse struct {
	Models []ModelListResponseModel `json:"models" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Models      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseModel struct {
	Capabilities []string `json:"capabilities" api:"required"`
	// Estimated credits
	EstimatedCredits int64 `json:"estimated_credits" api:"required"`
	// Estimated seconds
	EstimatedSeconds int64 `json:"estimated_seconds" api:"required"`
	// Model identifier
	ModelID string `json:"model_id" api:"required"`
	// Model name
	Name string `json:"name" api:"required"`
	// Supported generation parameters for this model
	Params []ModelListResponseModelParam `json:"params" api:"required"`
	// Model provider
	Provider string `json:"provider" api:"required"`
	// Model type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type" api:"required"`
	Beta bool   `json:"beta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Capabilities     respjson.Field
		EstimatedCredits respjson.Field
		EstimatedSeconds respjson.Field
		ModelID          respjson.Field
		Name             respjson.Field
		Params           respjson.Field
		Provider         respjson.Field
		Type             respjson.Field
		Beta             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseModelParam struct {
	// Parameter name to pass in generation params
	Name string `json:"name" api:"required"`
	// Whether the model requires this parameter
	Required bool `json:"required" api:"required"`
	// Parameter value type
	//
	// Any of "string", "string[]", "bool", "int", "int?", "seed", "float", "dict".
	Type string `json:"type" api:"required"`
	// Default parameter value
	Default any `json:"default"`
	// Parameter help text
	Description string `json:"description"`
	// Human-readable parameter label
	Label string `json:"label"`
	// Maximum numeric value
	Max float64 `json:"max"`
	// Minimum numeric value
	Min float64 `json:"min"`
	// Allowed values for enum-like parameters
	Options []ModelListResponseModelParamOption `json:"options"`
	// Nested numeric properties for object parameters
	Properties map[string]ModelListResponseModelParamProperty `json:"properties"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		Default     respjson.Field
		Description respjson.Field
		Label       respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		Options     respjson.Field
		Properties  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseModelParam) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseModelParamOption struct {
	// Displayed option label
	Label string `json:"label" api:"required"`
	// Option value to pass in generation params
	Value string `json:"value" api:"required"`
	// Option description
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseModelParamOption) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseModelParamOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseModelParamProperty struct {
	Default float64 `json:"default" api:"required"`
	Max     float64 `json:"max" api:"required"`
	Min     float64 `json:"min" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseModelParamProperty) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseModelParamProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListParams struct {
	// Model type
	//
	// Any of "image", "video", "audio", "text".
	Type ModelListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Model type
type ModelListParamsType string

const (
	ModelListParamsTypeImage ModelListParamsType = "image"
	ModelListParamsTypeVideo ModelListParamsType = "video"
	ModelListParamsTypeAudio ModelListParamsType = "audio"
	ModelListParamsTypeText  ModelListParamsType = "text"
)
