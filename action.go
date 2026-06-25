// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
	"github.com/florafauna-ai/flora-go/shared/constant"
)

// ActionService contains methods and other services that help with interacting
// with the flora API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	options []option.RequestOption
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.options = opts
	return
}

// Returns metadata for one released prebuilt Flora action. Action identifiers are
// raw slugs such as rotate-image, not action-prefixed IDs.
func (r *ActionService) Get(ctx context.Context, actionID ActionGetParamsActionID, opts ...option.RequestOption) (res *ActionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("actions/%v", actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns released prebuilt Flora actions that can be executed by the public API.
// Action identifiers are raw slugs such as rotate-image, not action-prefixed IDs.
func (r *ActionService) List(ctx context.Context, opts ...option.RequestOption) (res *ActionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Starts a headless prebuilt action run using raw action slugs such as
// rotate-image, workspace_id, project_id, and non-empty inputs. Direct action runs
// do not create or mutate canvas nodes. Mutating public API requests support an
// optional Idempotency-Key header for client retries; duplicate keys within two
// hours return idempotency_duplicate.
func (r *ActionService) Run(ctx context.Context, body ActionRunParams, opts ...option.RequestOption) (res *ActionRunResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "runs/action"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ActionGetResponse struct {
	// Action identifier
	//
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID ActionGetResponseActionID `json:"action_id" api:"required"`
	// Cost per execution in USD
	ChargedCost float64 `json:"charged_cost" api:"required"`
	// Action description
	Description string `json:"description" api:"required"`
	// Action input slots
	Inputs []ActionGetResponseInput `json:"inputs" api:"required"`
	// Action runtime language
	//
	// Any of "javascript", "python".
	Language ActionGetResponseLanguage `json:"language" api:"required"`
	// Action name
	Name string `json:"name" api:"required"`
	// Action output slots
	Outputs []ActionGetResponseOutput `json:"outputs" api:"required"`
	// Action parameters
	Params []ActionGetResponseParam `json:"params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ChargedCost respjson.Field
		Description respjson.Field
		Inputs      respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action identifier
type ActionGetResponseActionID string

const (
	ActionGetResponseActionIDSplitText           ActionGetResponseActionID = "split-text"
	ActionGetResponseActionIDFindAndReplaceText  ActionGetResponseActionID = "find-and-replace-text"
	ActionGetResponseActionIDConcatText          ActionGetResponseActionID = "concat-text"
	ActionGetResponseActionIDKenBurnsVideo       ActionGetResponseActionID = "ken-burns-video"
	ActionGetResponseActionIDColorGradeImage     ActionGetResponseActionID = "color-grade-image"
	ActionGetResponseActionIDChangeImageAr       ActionGetResponseActionID = "change-image-ar"
	ActionGetResponseActionIDRotateImage         ActionGetResponseActionID = "rotate-image"
	ActionGetResponseActionIDFlipImage           ActionGetResponseActionID = "flip-image"
	ActionGetResponseActionIDColorFilterImage    ActionGetResponseActionID = "color-filter-image"
	ActionGetResponseActionIDColorTintImage      ActionGetResponseActionID = "color-tint-image"
	ActionGetResponseActionIDFilterColorImage    ActionGetResponseActionID = "filter-color-image"
	ActionGetResponseActionIDBlurImage           ActionGetResponseActionID = "blur-image"
	ActionGetResponseActionIDDuplicateImage      ActionGetResponseActionID = "duplicate-image"
	ActionGetResponseActionIDSideBySideComposite ActionGetResponseActionID = "side-by-side-composite"
	ActionGetResponseActionIDAddShapeToImage     ActionGetResponseActionID = "add-shape-to-image"
	ActionGetResponseActionIDGenerateShapeImage  ActionGetResponseActionID = "generate-shape-image"
	ActionGetResponseActionIDAddTextToImage      ActionGetResponseActionID = "add-text-to-image"
	ActionGetResponseActionIDGenerateTextImage   ActionGetResponseActionID = "generate-text-image"
	ActionGetResponseActionIDQrCodeGenerator     ActionGetResponseActionID = "qr-code-generator"
	ActionGetResponseActionIDStitchVideos        ActionGetResponseActionID = "stitch-videos"
	ActionGetResponseActionIDSplitVideo          ActionGetResponseActionID = "split-video"
	ActionGetResponseActionIDExtractVideoFrames  ActionGetResponseActionID = "extract-video-frames"
	ActionGetResponseActionIDColorGradeVideo     ActionGetResponseActionID = "color-grade-video"
	ActionGetResponseActionIDVideoToFrameGrid    ActionGetResponseActionID = "video-to-frame-grid"
	ActionGetResponseActionIDBoomerangVideo      ActionGetResponseActionID = "boomerang-video"
	ActionGetResponseActionIDReverseVideo        ActionGetResponseActionID = "reverse-video"
	ActionGetResponseActionIDVideoToLongExposure ActionGetResponseActionID = "video-to-long-exposure"
	ActionGetResponseActionIDVideoEffect         ActionGetResponseActionID = "video-effect"
	ActionGetResponseActionIDColorFilterVideo    ActionGetResponseActionID = "color-filter-video"
	ActionGetResponseActionIDSpeedUpVideo        ActionGetResponseActionID = "speed-up-video"
	ActionGetResponseActionIDSlowDownVideo       ActionGetResponseActionID = "slow-down-video"
	ActionGetResponseActionIDDuplicateVideo      ActionGetResponseActionID = "duplicate-video"
	ActionGetResponseActionIDGreenscreenVideo    ActionGetResponseActionID = "greenscreen-video"
	ActionGetResponseActionIDResizeVideo         ActionGetResponseActionID = "resize-video"
	ActionGetResponseActionIDChangeVideoAr       ActionGetResponseActionID = "change-video-ar"
	ActionGetResponseActionIDSplitAudioFromVideo ActionGetResponseActionID = "split-audio-from-video"
	ActionGetResponseActionIDMergeAudioIntoVideo ActionGetResponseActionID = "merge-audio-into-video"
)

type ActionGetResponseInput struct {
	// Action input or output name
	Name string `json:"name" api:"required"`
	// Action input or output media type
	//
	// Any of "image", "video", "text", "audio".
	Type string `json:"type" api:"required"`
	// Deprecated alias for `multiple: true`. Mirrors `multiple` for back-compat.
	Dynamic bool `json:"dynamic"`
	// Many-connections form. `true` is unbounded shorthand; the object form sets
	// explicit min/max bounds.
	Multiple ActionGetResponseInputMultipleUnion `json:"multiple"`
	// Whether the slot allows zero connections in single-input mode. Ignored when
	// `multiple` is the object form (use `min: 0` there instead).
	Optional bool `json:"optional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Dynamic     respjson.Field
		Multiple    respjson.Field
		Optional    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponseInput) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActionGetResponseInputMultipleUnion contains all possible properties and values
// from [bool], [ActionGetResponseInputMultipleObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfActionGetResponseInputMultipleBoolean]
type ActionGetResponseInputMultipleUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfActionGetResponseInputMultipleBoolean bool `json:",inline"`
	// This field is from variant [ActionGetResponseInputMultipleObject].
	Max int64 `json:"max"`
	// This field is from variant [ActionGetResponseInputMultipleObject].
	Min  int64 `json:"min"`
	JSON struct {
		OfActionGetResponseInputMultipleBoolean respjson.Field
		Max                                     respjson.Field
		Min                                     respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ActionGetResponseInputMultipleUnion) AsActionGetResponseInputMultipleBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActionGetResponseInputMultipleUnion) AsActionGetResponseInputMultipleObject() (v ActionGetResponseInputMultipleObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActionGetResponseInputMultipleUnion) RawJSON() string { return u.JSON.raw }

func (r *ActionGetResponseInputMultipleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionGetResponseInputMultipleBoolean bool

const (
	ActionGetResponseInputMultipleBooleanTrue ActionGetResponseInputMultipleBoolean = true
)

type ActionGetResponseInputMultipleObject struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponseInputMultipleObject) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseInputMultipleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action runtime language
type ActionGetResponseLanguage string

const (
	ActionGetResponseLanguageJavascript ActionGetResponseLanguage = "javascript"
	ActionGetResponseLanguagePython     ActionGetResponseLanguage = "python"
)

type ActionGetResponseOutput struct {
	// Action input or output name
	Name string `json:"name" api:"required"`
	// Action input or output media type
	//
	// Any of "image", "video", "text", "audio".
	Type string `json:"type" api:"required"`
	// Deprecated alias for `multiple: true`. Mirrors `multiple` for back-compat.
	Dynamic bool `json:"dynamic"`
	// Many-connections form. `true` is unbounded shorthand; the object form sets
	// explicit min/max bounds.
	Multiple ActionGetResponseOutputMultipleUnion `json:"multiple"`
	// Whether the slot allows zero connections in single-input mode. Ignored when
	// `multiple` is the object form (use `min: 0` there instead).
	Optional bool `json:"optional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Dynamic     respjson.Field
		Multiple    respjson.Field
		Optional    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponseOutput) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActionGetResponseOutputMultipleUnion contains all possible properties and values
// from [bool], [ActionGetResponseOutputMultipleObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfActionGetResponseOutputMultipleBoolean]
type ActionGetResponseOutputMultipleUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfActionGetResponseOutputMultipleBoolean bool `json:",inline"`
	// This field is from variant [ActionGetResponseOutputMultipleObject].
	Max int64 `json:"max"`
	// This field is from variant [ActionGetResponseOutputMultipleObject].
	Min  int64 `json:"min"`
	JSON struct {
		OfActionGetResponseOutputMultipleBoolean respjson.Field
		Max                                      respjson.Field
		Min                                      respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u ActionGetResponseOutputMultipleUnion) AsActionGetResponseOutputMultipleBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActionGetResponseOutputMultipleUnion) AsActionGetResponseOutputMultipleObject() (v ActionGetResponseOutputMultipleObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActionGetResponseOutputMultipleUnion) RawJSON() string { return u.JSON.raw }

func (r *ActionGetResponseOutputMultipleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionGetResponseOutputMultipleBoolean bool

const (
	ActionGetResponseOutputMultipleBooleanTrue ActionGetResponseOutputMultipleBoolean = true
)

type ActionGetResponseOutputMultipleObject struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponseOutputMultipleObject) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseOutputMultipleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionGetResponseParam struct {
	// Action parameter key
	Key string `json:"key" api:"required"`
	// Action parameter control type
	Type string `json:"type" api:"required"`
	// Allowed parameter values
	AvailableValues []ActionGetResponseParamAvailableValue `json:"available_values"`
	// Default parameter value
	DefaultValue any `json:"default_value"`
	// Action parameter description
	InfoTooltip string `json:"info_tooltip"`
	// Action parameter label
	Label          string  `json:"label"`
	MaxNumberValue float64 `json:"max_number_value"`
	MinNumberValue float64 `json:"min_number_value"`
	NumberStep     float64 `json:"number_step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key             respjson.Field
		Type            respjson.Field
		AvailableValues respjson.Field
		DefaultValue    respjson.Field
		InfoTooltip     respjson.Field
		Label           respjson.Field
		MaxNumberValue  respjson.Field
		MinNumberValue  respjson.Field
		NumberStep      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionGetResponseParam) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionGetResponseParamAvailableValue struct {
	Label       string `json:"label" api:"required"`
	Value       string `json:"value" api:"required"`
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
func (r ActionGetResponseParamAvailableValue) RawJSON() string { return r.JSON.raw }
func (r *ActionGetResponseParamAvailableValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponse struct {
	Actions []ActionListResponseAction `json:"actions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseAction struct {
	// Action identifier
	//
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID string `json:"action_id" api:"required"`
	// Cost per execution in USD
	ChargedCost float64 `json:"charged_cost" api:"required"`
	// Action description
	Description string `json:"description" api:"required"`
	// Action input slots
	Inputs []ActionListResponseActionInput `json:"inputs" api:"required"`
	// Action runtime language
	//
	// Any of "javascript", "python".
	Language string `json:"language" api:"required"`
	// Action name
	Name string `json:"name" api:"required"`
	// Action output slots
	Outputs []ActionListResponseActionOutput `json:"outputs" api:"required"`
	// Action parameters
	Params []ActionListResponseActionParam `json:"params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ChargedCost respjson.Field
		Description respjson.Field
		Inputs      respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseAction) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionInput struct {
	// Action input or output name
	Name string `json:"name" api:"required"`
	// Action input or output media type
	//
	// Any of "image", "video", "text", "audio".
	Type string `json:"type" api:"required"`
	// Deprecated alias for `multiple: true`. Mirrors `multiple` for back-compat.
	Dynamic bool `json:"dynamic"`
	// Many-connections form. `true` is unbounded shorthand; the object form sets
	// explicit min/max bounds.
	Multiple ActionListResponseActionInputMultipleUnion `json:"multiple"`
	// Whether the slot allows zero connections in single-input mode. Ignored when
	// `multiple` is the object form (use `min: 0` there instead).
	Optional bool `json:"optional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Dynamic     respjson.Field
		Multiple    respjson.Field
		Optional    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseActionInput) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActionListResponseActionInputMultipleUnion contains all possible properties and
// values from [bool], [ActionListResponseActionInputMultipleObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfActionListResponseActionInputMultipleBoolean]
type ActionListResponseActionInputMultipleUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfActionListResponseActionInputMultipleBoolean bool `json:",inline"`
	// This field is from variant [ActionListResponseActionInputMultipleObject].
	Max int64 `json:"max"`
	// This field is from variant [ActionListResponseActionInputMultipleObject].
	Min  int64 `json:"min"`
	JSON struct {
		OfActionListResponseActionInputMultipleBoolean respjson.Field
		Max                                            respjson.Field
		Min                                            respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u ActionListResponseActionInputMultipleUnion) AsActionListResponseActionInputMultipleBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActionListResponseActionInputMultipleUnion) AsActionListResponseActionInputMultipleObject() (v ActionListResponseActionInputMultipleObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActionListResponseActionInputMultipleUnion) RawJSON() string { return u.JSON.raw }

func (r *ActionListResponseActionInputMultipleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionInputMultipleBoolean bool

const (
	ActionListResponseActionInputMultipleBooleanTrue ActionListResponseActionInputMultipleBoolean = true
)

type ActionListResponseActionInputMultipleObject struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseActionInputMultipleObject) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionInputMultipleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionOutput struct {
	// Action input or output name
	Name string `json:"name" api:"required"`
	// Action input or output media type
	//
	// Any of "image", "video", "text", "audio".
	Type string `json:"type" api:"required"`
	// Deprecated alias for `multiple: true`. Mirrors `multiple` for back-compat.
	Dynamic bool `json:"dynamic"`
	// Many-connections form. `true` is unbounded shorthand; the object form sets
	// explicit min/max bounds.
	Multiple ActionListResponseActionOutputMultipleUnion `json:"multiple"`
	// Whether the slot allows zero connections in single-input mode. Ignored when
	// `multiple` is the object form (use `min: 0` there instead).
	Optional bool `json:"optional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		Dynamic     respjson.Field
		Multiple    respjson.Field
		Optional    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseActionOutput) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ActionListResponseActionOutputMultipleUnion contains all possible properties and
// values from [bool], [ActionListResponseActionOutputMultipleObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfActionListResponseActionOutputMultipleBoolean]
type ActionListResponseActionOutputMultipleUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfActionListResponseActionOutputMultipleBoolean bool `json:",inline"`
	// This field is from variant [ActionListResponseActionOutputMultipleObject].
	Max int64 `json:"max"`
	// This field is from variant [ActionListResponseActionOutputMultipleObject].
	Min  int64 `json:"min"`
	JSON struct {
		OfActionListResponseActionOutputMultipleBoolean respjson.Field
		Max                                             respjson.Field
		Min                                             respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u ActionListResponseActionOutputMultipleUnion) AsActionListResponseActionOutputMultipleBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ActionListResponseActionOutputMultipleUnion) AsActionListResponseActionOutputMultipleObject() (v ActionListResponseActionOutputMultipleObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ActionListResponseActionOutputMultipleUnion) RawJSON() string { return u.JSON.raw }

func (r *ActionListResponseActionOutputMultipleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionOutputMultipleBoolean bool

const (
	ActionListResponseActionOutputMultipleBooleanTrue ActionListResponseActionOutputMultipleBoolean = true
)

type ActionListResponseActionOutputMultipleObject struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseActionOutputMultipleObject) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionOutputMultipleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionParam struct {
	// Action parameter key
	Key string `json:"key" api:"required"`
	// Action parameter control type
	Type string `json:"type" api:"required"`
	// Allowed parameter values
	AvailableValues []ActionListResponseActionParamAvailableValue `json:"available_values"`
	// Default parameter value
	DefaultValue any `json:"default_value"`
	// Action parameter description
	InfoTooltip string `json:"info_tooltip"`
	// Action parameter label
	Label          string  `json:"label"`
	MaxNumberValue float64 `json:"max_number_value"`
	MinNumberValue float64 `json:"min_number_value"`
	NumberStep     float64 `json:"number_step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key             respjson.Field
		Type            respjson.Field
		AvailableValues respjson.Field
		DefaultValue    respjson.Field
		InfoTooltip     respjson.Field
		Label           respjson.Field
		MaxNumberValue  respjson.Field
		MinNumberValue  respjson.Field
		NumberStep      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionListResponseActionParam) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionListResponseActionParamAvailableValue struct {
	Label       string `json:"label" api:"required"`
	Value       string `json:"value" api:"required"`
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
func (r ActionListResponseActionParamAvailableValue) RawJSON() string { return r.JSON.raw }
func (r *ActionListResponseActionParamAvailableValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRunResponse struct {
	// Cost charged in USD
	ChargedCost      float64 `json:"charged_cost" api:"required"`
	EstimatedSeconds int64   `json:"estimated_seconds" api:"required"`
	// Run identifier
	RunID string `json:"run_id" api:"required"`
	// Run type
	//
	// Any of "generation", "technique", "action".
	Type   ActionRunResponseType   `json:"type" api:"required"`
	Action ActionRunResponseAction `json:"action" api:"nullable"`
	Model  ActionRunResponseModel  `json:"model" api:"nullable"`
	// URL to poll pending/running runs or fetch completed/failed run details.
	PollURL string `json:"poll_url" api:"nullable" format:"uri"`
	// Project identifier
	ProjectID string                     `json:"project_id" api:"nullable"`
	Technique ActionRunResponseTechnique `json:"technique" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargedCost      respjson.Field
		EstimatedSeconds respjson.Field
		RunID            respjson.Field
		Type             respjson.Field
		Action           respjson.Field
		Model            respjson.Field
		PollURL          respjson.Field
		ProjectID        respjson.Field
		Technique        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run type
type ActionRunResponseType string

const (
	ActionRunResponseTypeGeneration ActionRunResponseType = "generation"
	ActionRunResponseTypeTechnique  ActionRunResponseType = "technique"
	ActionRunResponseTypeAction     ActionRunResponseType = "action"
)

type ActionRunResponseAction struct {
	// Action identifier
	//
	// Any of "split-text", "find-and-replace-text", "concat-text", "ken-burns-video",
	// "color-grade-image", "change-image-ar", "rotate-image", "flip-image",
	// "color-filter-image", "color-tint-image", "filter-color-image", "blur-image",
	// "duplicate-image", "side-by-side-composite", "add-shape-to-image",
	// "generate-shape-image", "add-text-to-image", "generate-text-image",
	// "qr-code-generator", "stitch-videos", "split-video", "extract-video-frames",
	// "color-grade-video", "video-to-frame-grid", "boomerang-video", "reverse-video",
	// "video-to-long-exposure", "video-effect", "color-filter-video",
	// "speed-up-video", "slow-down-video", "duplicate-video", "greenscreen-video",
	// "resize-video", "change-video-ar", "split-audio-from-video",
	// "merge-audio-into-video".
	ActionID string `json:"action_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionRunResponseAction) RawJSON() string { return r.JSON.raw }
func (r *ActionRunResponseAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRunResponseModel struct {
	// Model identifier
	ModelID string `json:"model_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionRunResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ActionRunResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRunResponseTechnique struct {
	// Technique name
	Name string `json:"name" api:"required"`
	// Technique identifier
	TechniqueID string `json:"technique_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		TechniqueID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionRunResponseTechnique) RawJSON() string { return r.JSON.raw }
func (r *ActionRunResponseTechnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action identifier
type ActionGetParamsActionID string

const (
	ActionGetParamsActionIDSplitText           ActionGetParamsActionID = "split-text"
	ActionGetParamsActionIDFindAndReplaceText  ActionGetParamsActionID = "find-and-replace-text"
	ActionGetParamsActionIDConcatText          ActionGetParamsActionID = "concat-text"
	ActionGetParamsActionIDKenBurnsVideo       ActionGetParamsActionID = "ken-burns-video"
	ActionGetParamsActionIDColorGradeImage     ActionGetParamsActionID = "color-grade-image"
	ActionGetParamsActionIDChangeImageAr       ActionGetParamsActionID = "change-image-ar"
	ActionGetParamsActionIDRotateImage         ActionGetParamsActionID = "rotate-image"
	ActionGetParamsActionIDFlipImage           ActionGetParamsActionID = "flip-image"
	ActionGetParamsActionIDColorFilterImage    ActionGetParamsActionID = "color-filter-image"
	ActionGetParamsActionIDColorTintImage      ActionGetParamsActionID = "color-tint-image"
	ActionGetParamsActionIDFilterColorImage    ActionGetParamsActionID = "filter-color-image"
	ActionGetParamsActionIDBlurImage           ActionGetParamsActionID = "blur-image"
	ActionGetParamsActionIDDuplicateImage      ActionGetParamsActionID = "duplicate-image"
	ActionGetParamsActionIDSideBySideComposite ActionGetParamsActionID = "side-by-side-composite"
	ActionGetParamsActionIDAddShapeToImage     ActionGetParamsActionID = "add-shape-to-image"
	ActionGetParamsActionIDGenerateShapeImage  ActionGetParamsActionID = "generate-shape-image"
	ActionGetParamsActionIDAddTextToImage      ActionGetParamsActionID = "add-text-to-image"
	ActionGetParamsActionIDGenerateTextImage   ActionGetParamsActionID = "generate-text-image"
	ActionGetParamsActionIDQrCodeGenerator     ActionGetParamsActionID = "qr-code-generator"
	ActionGetParamsActionIDStitchVideos        ActionGetParamsActionID = "stitch-videos"
	ActionGetParamsActionIDSplitVideo          ActionGetParamsActionID = "split-video"
	ActionGetParamsActionIDExtractVideoFrames  ActionGetParamsActionID = "extract-video-frames"
	ActionGetParamsActionIDColorGradeVideo     ActionGetParamsActionID = "color-grade-video"
	ActionGetParamsActionIDVideoToFrameGrid    ActionGetParamsActionID = "video-to-frame-grid"
	ActionGetParamsActionIDBoomerangVideo      ActionGetParamsActionID = "boomerang-video"
	ActionGetParamsActionIDReverseVideo        ActionGetParamsActionID = "reverse-video"
	ActionGetParamsActionIDVideoToLongExposure ActionGetParamsActionID = "video-to-long-exposure"
	ActionGetParamsActionIDVideoEffect         ActionGetParamsActionID = "video-effect"
	ActionGetParamsActionIDColorFilterVideo    ActionGetParamsActionID = "color-filter-video"
	ActionGetParamsActionIDSpeedUpVideo        ActionGetParamsActionID = "speed-up-video"
	ActionGetParamsActionIDSlowDownVideo       ActionGetParamsActionID = "slow-down-video"
	ActionGetParamsActionIDDuplicateVideo      ActionGetParamsActionID = "duplicate-video"
	ActionGetParamsActionIDGreenscreenVideo    ActionGetParamsActionID = "greenscreen-video"
	ActionGetParamsActionIDResizeVideo         ActionGetParamsActionID = "resize-video"
	ActionGetParamsActionIDChangeVideoAr       ActionGetParamsActionID = "change-video-ar"
	ActionGetParamsActionIDSplitAudioFromVideo ActionGetParamsActionID = "split-audio-from-video"
	ActionGetParamsActionIDMergeAudioIntoVideo ActionGetParamsActionID = "merge-audio-into-video"
)

type ActionRunParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfObject *ActionRunParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject2 *ActionRunParamsBodyObject2 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject3 *ActionRunParamsBodyObject3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject4 *ActionRunParamsBodyObject4 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject5 *ActionRunParamsBodyObject5 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject6 *ActionRunParamsBodyObject6 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject7 *ActionRunParamsBodyObject7 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject8 *ActionRunParamsBodyObject8 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject9 *ActionRunParamsBodyObject9 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject10 *ActionRunParamsBodyObject10 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject11 *ActionRunParamsBodyObject11 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject12 *ActionRunParamsBodyObject12 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject13 *ActionRunParamsBodyObject13 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject14 *ActionRunParamsBodyObject14 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject15 *ActionRunParamsBodyObject15 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject16 *ActionRunParamsBodyObject16 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject17 *ActionRunParamsBodyObject17 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject18 *ActionRunParamsBodyObject18 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject19 *ActionRunParamsBodyObject19 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject20 *ActionRunParamsBodyObject20 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject21 *ActionRunParamsBodyObject21 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject22 *ActionRunParamsBodyObject22 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject23 *ActionRunParamsBodyObject23 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject24 *ActionRunParamsBodyObject24 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject25 *ActionRunParamsBodyObject25 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject26 *ActionRunParamsBodyObject26 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject27 *ActionRunParamsBodyObject27 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject28 *ActionRunParamsBodyObject28 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject29 *ActionRunParamsBodyObject29 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject30 *ActionRunParamsBodyObject30 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject31 *ActionRunParamsBodyObject31 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject32 *ActionRunParamsBodyObject32 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject33 *ActionRunParamsBodyObject33 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject34 *ActionRunParamsBodyObject34 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject35 *ActionRunParamsBodyObject35 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject36 *ActionRunParamsBodyObject36 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfActionRunsBodyObject37 *ActionRunParamsBodyObject37 `json:",inline"`

	paramObj
}

func (u ActionRunParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfObject,
		u.OfActionRunsBodyObject2,
		u.OfActionRunsBodyObject3,
		u.OfActionRunsBodyObject4,
		u.OfActionRunsBodyObject5,
		u.OfActionRunsBodyObject6,
		u.OfActionRunsBodyObject7,
		u.OfActionRunsBodyObject8,
		u.OfActionRunsBodyObject9,
		u.OfActionRunsBodyObject10,
		u.OfActionRunsBodyObject11,
		u.OfActionRunsBodyObject12,
		u.OfActionRunsBodyObject13,
		u.OfActionRunsBodyObject14,
		u.OfActionRunsBodyObject15,
		u.OfActionRunsBodyObject16,
		u.OfActionRunsBodyObject17,
		u.OfActionRunsBodyObject18,
		u.OfActionRunsBodyObject19,
		u.OfActionRunsBodyObject20,
		u.OfActionRunsBodyObject21,
		u.OfActionRunsBodyObject22,
		u.OfActionRunsBodyObject23,
		u.OfActionRunsBodyObject24,
		u.OfActionRunsBodyObject25,
		u.OfActionRunsBodyObject26,
		u.OfActionRunsBodyObject27,
		u.OfActionRunsBodyObject28,
		u.OfActionRunsBodyObject29,
		u.OfActionRunsBodyObject30,
		u.OfActionRunsBodyObject31,
		u.OfActionRunsBodyObject32,
		u.OfActionRunsBodyObject33,
		u.OfActionRunsBodyObject34,
		u.OfActionRunsBodyObject35,
		u.OfActionRunsBodyObject36,
		u.OfActionRunsBodyObject37)
}
func (r *ActionRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObjectInput `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObjectParams `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "split-text".
	ActionID constant.SplitText `json:"action_id" default:"split-text"`
	paramObj
}

func (r ActionRunParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObjectInput struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObjectInput) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObjectInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObjectInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObjectInput](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObjectParams struct {
	// Characters per Part
	CharsPerPart param.Opt[float64] `json:"chars_per_part,omitzero"`
	// Lines per Part
	LinesPerPart param.Opt[float64] `json:"lines_per_part,omitzero"`
	// Max Parts
	MaxParts param.Opt[float64] `json:"max_parts,omitzero"`
	// Separator
	Separator param.Opt[string] `json:"separator,omitzero"`
	// Skip Empty Parts
	SkipEmpty param.Opt[bool] `json:"skip_empty,omitzero"`
	// Trim Whitespace
	TrimParts param.Opt[bool] `json:"trim_parts,omitzero"`
	// Split Mode
	//
	// Any of "separator", "paragraph", "lines", "charCount".
	SplitMode string `json:"split_mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObjectParams) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObjectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObjectParams](
		"split_mode", "separator", "paragraph", "lines", "charCount",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject2 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject2Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject2Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "find-and-replace-text".
	ActionID constant.FindAndReplaceText `json:"action_id" default:"find-and-replace-text"`
	paramObj
}

func (r ActionRunParamsBodyObject2) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject2Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject2Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject2Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject2Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject2Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject2Params struct {
	// Case Sensitive
	CaseSensitive param.Opt[bool] `json:"case_sensitive,omitzero"`
	// Find
	Find param.Opt[string] `json:"find,omitzero"`
	// Replace With
	Replace param.Opt[string] `json:"replace,omitzero"`
	// Replace All Occurrences
	ReplaceAll param.Opt[bool] `json:"replace_all,omitzero"`
	// Whole Word Only
	WholeWord param.Opt[bool] `json:"whole_word,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject2Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject2Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject2Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject3 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject3Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject3Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "concat-text".
	ActionID constant.ConcatText `json:"action_id" default:"concat-text"`
	paramObj
}

func (r ActionRunParamsBodyObject3) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject3Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject3Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject3Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject3Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject3Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject3Params struct {
	// Add Input Name Headers
	AddHeaders param.Opt[bool] `json:"add_headers,omitzero"`
	// Prefix
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Separator
	Separator param.Opt[string] `json:"separator,omitzero"`
	// Skip Empty Parts
	SkipEmpty param.Opt[bool] `json:"skip_empty,omitzero"`
	// Suffix
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// Trim Each Part
	TrimParts param.Opt[bool] `json:"trim_parts,omitzero"`
	// Wrap Each Part
	WrapEachPart param.Opt[bool] `json:"wrap_each_part,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject3Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject3Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject3Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject4 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject4Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject4Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "ken-burns-video".
	ActionID constant.KenBurnsVideo `json:"action_id" default:"ken-burns-video"`
	paramObj
}

func (r ActionRunParamsBodyObject4) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject4Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject4Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject4Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject4Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject4Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject4Params struct {
	// Duration (s)
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// Zoom Level
	Zoom param.Opt[float64] `json:"zoom,omitzero"`
	// Easing
	//
	// Any of "linear", "ease-in", "ease-out", "ease-in-out".
	Easing string `json:"easing,omitzero"`
	// FPS
	//
	// Any of "24", "30", "60".
	Fps string `json:"fps,omitzero"`
	// Pan
	//
	// Any of "none", "left", "right", "up", "down".
	PanDirection string `json:"pan_direction,omitzero"`
	// Zoom Direction
	//
	// Any of "none", "in", "out".
	ZoomDirection string `json:"zoom_direction,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject4Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject4Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject4Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject4Params](
		"easing", "linear", "ease-in", "ease-out", "ease-in-out",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject4Params](
		"fps", "24", "30", "60",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject4Params](
		"pan_direction", "none", "left", "right", "up", "down",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject4Params](
		"zoom_direction", "none", "in", "out",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject5 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject5Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject5Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "color-grade-image".
	ActionID constant.ColorGradeImage `json:"action_id" default:"color-grade-image"`
	paramObj
}

func (r ActionRunParamsBodyObject5) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject5
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject5Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject5Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject5Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject5Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject5Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject5Params struct {
	// Brightness
	Brightness param.Opt[float64] `json:"brightness,omitzero"`
	// Contrast
	Contrast param.Opt[float64] `json:"contrast,omitzero"`
	// Saturation
	Saturation param.Opt[float64] `json:"saturation,omitzero"`
	// Warmth
	Warmth param.Opt[float64] `json:"warmth,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject5Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject5Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject5Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject6 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject6Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject6Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "change-image-ar".
	ActionID constant.ChangeImageAr `json:"action_id" default:"change-image-ar"`
	paramObj
}

func (r ActionRunParamsBodyObject6) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject6
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject6) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject6Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject6Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject6Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject6Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject6Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject6Params struct {
	// Blur Amount
	BlurAmount param.Opt[float64] `json:"blur_amount,omitzero"`
	// Pad Color
	PadColor param.Opt[string] `json:"pad_color,omitzero"`
	// Aspect Ratio
	//
	// Any of "1:1", "16:9", "9:16", "4:3", "3:4", "3:2", "2:3", "21:9".
	AspectRatio string `json:"aspect_ratio,omitzero"`
	// Background
	//
	// Any of "solid", "blur".
	BackgroundMode string `json:"background_mode,omitzero"`
	// Fit
	//
	// Any of "crop", "pad".
	Fit string `json:"fit,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject6Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject6Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject6Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject6Params](
		"aspect_ratio", "1:1", "16:9", "9:16", "4:3", "3:4", "3:2", "2:3", "21:9",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject6Params](
		"background_mode", "solid", "blur",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject6Params](
		"fit", "crop", "pad",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject7 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject7Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject7Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "rotate-image".
	ActionID constant.RotateImage `json:"action_id" default:"rotate-image"`
	paramObj
}

func (r ActionRunParamsBodyObject7) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject7
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject7) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject7Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject7Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject7Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject7Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject7Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject7Params struct {
	// Angle (°, clockwise)
	Angle param.Opt[float64] `json:"angle,omitzero"`
	// Background Color
	Background param.Opt[string] `json:"background,omitzero"`
	// Expand Canvas to Fit
	Expand param.Opt[bool] `json:"expand,omitzero"`
	// Transparent Background
	Transparent param.Opt[bool] `json:"transparent,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject7Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject7Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject7Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject8 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject8Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject8Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "flip-image".
	ActionID constant.FlipImage `json:"action_id" default:"flip-image"`
	paramObj
}

func (r ActionRunParamsBodyObject8) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject8
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject8) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject8Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject8Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject8Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject8Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject8Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject8Params struct {
	// Direction
	//
	// Any of "horizontal", "vertical", "both".
	Direction string `json:"direction,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject8Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject8Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject8Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject8Params](
		"direction", "horizontal", "vertical", "both",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject9 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject9Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject9Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "color-filter-image".
	ActionID constant.ColorFilterImage `json:"action_id" default:"color-filter-image"`
	paramObj
}

func (r ActionRunParamsBodyObject9) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject9
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject9) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject9Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject9Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject9Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject9Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject9Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject9Params struct {
	// Threshold
	BwThreshold param.Opt[float64] `json:"bw_threshold,omitzero"`
	// Color Tolerance
	ColorpopTolerance param.Opt[float64] `json:"colorpop_tolerance,omitzero"`
	// Grid Angle (°)
	DotAngle param.Opt[float64] `json:"dot_angle,omitzero"`
	// Background
	DotBg param.Opt[string] `json:"dot_bg,omitzero"`
	// Dot Color
	DotColor param.Opt[string] `json:"dot_color,omitzero"`
	// Dot Size
	DotSize param.Opt[float64] `json:"dot_size,omitzero"`
	// Grain
	Grain param.Opt[float64] `json:"grain,omitzero"`
	// Highlight Color
	HighlightColor param.Opt[string] `json:"highlight_color,omitzero"`
	// Intensity
	Intensity param.Opt[float64] `json:"intensity,omitzero"`
	// Bits per channel
	PosterizeBits param.Opt[float64] `json:"posterize_bits,omitzero"`
	// Warmth
	SepiaWarmth param.Opt[float64] `json:"sepia_warmth,omitzero"`
	// Shadow Color
	ShadowColor param.Opt[string] `json:"shadow_color,omitzero"`
	// Threshold
	SolarizeThreshold param.Opt[float64] `json:"solarize_threshold,omitzero"`
	// Keep Color
	TargetColor param.Opt[string] `json:"target_color,omitzero"`
	// Vignette Softness
	VignetteSoftness param.Opt[float64] `json:"vignette_softness,omitzero"`
	// Vignette Strength
	VignetteStrength param.Opt[float64] `json:"vignette_strength,omitzero"`
	// Filter
	//
	// Any of "grayscale", "sepia", "invert", "bw", "posterize", "solarize", "duotone",
	// "clarendon", "moon", "nashville", "noir", "fade", "vignette", "colorpop",
	// "crossprocess", "halftone".
	Filter string `json:"filter,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject9Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject9Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject9Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject9Params](
		"filter", "grayscale", "sepia", "invert", "bw", "posterize", "solarize", "duotone", "clarendon", "moon", "nashville", "noir", "fade", "vignette", "colorpop", "crossprocess", "halftone",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject10 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject10Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject10Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "color-tint-image".
	ActionID constant.ColorTintImage `json:"action_id" default:"color-tint-image"`
	paramObj
}

func (r ActionRunParamsBodyObject10) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject10
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject10) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject10Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject10Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject10Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject10Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject10Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject10Params struct {
	// Tint Color
	Color param.Opt[string] `json:"color,omitzero"`
	// Intensity
	Intensity param.Opt[float64] `json:"intensity,omitzero"`
	// Blend Mode
	//
	// Any of "multiply", "screen", "overlay", "soft_light", "color".
	BlendMode string `json:"blend_mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject10Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject10Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject10Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject10Params](
		"blend_mode", "multiply", "screen", "overlay", "soft_light", "color",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject11 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject11Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject11Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "filter-color-image".
	ActionID constant.FilterColorImage `json:"action_id" default:"filter-color-image"`
	paramObj
}

func (r ActionRunParamsBodyObject11) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject11
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject11) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject11Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject11Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject11Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject11Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject11Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject11Params struct {
	// Invert Selection
	Invert param.Opt[bool] `json:"invert,omitzero"`
	// Replacement Color
	ReplacementColor param.Opt[string] `json:"replacement_color,omitzero"`
	// Edge Softness
	Softness param.Opt[float64] `json:"softness,omitzero"`
	// Target Color
	TargetColor param.Opt[string] `json:"target_color,omitzero"`
	// Tolerance
	Tolerance param.Opt[float64] `json:"tolerance,omitzero"`
	// Mode
	//
	// Any of "remove", "replace", "keep".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject11Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject11Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject11Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject11Params](
		"mode", "remove", "replace", "keep",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject12 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject12Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject12Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "blur-image".
	ActionID constant.BlurImage `json:"action_id" default:"blur-image"`
	paramObj
}

func (r ActionRunParamsBodyObject12) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject12
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject12) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject12Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject12Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject12Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject12Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject12Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject12Params struct {
	// Motion Angle (°)
	Angle param.Opt[float64] `json:"angle,omitzero"`
	// Edge Preservation
	EdgeThreshold param.Opt[float64] `json:"edge_threshold,omitzero"`
	// Radial Strength
	RadialStrength param.Opt[float64] `json:"radial_strength,omitzero"`
	// Radius
	Radius param.Opt[float64] `json:"radius,omitzero"`
	// Target Color
	TargetColor param.Opt[string] `json:"target_color,omitzero"`
	// Invert Selection (blur non-matching)
	TargetInvert param.Opt[bool] `json:"target_invert,omitzero"`
	// Color Tolerance
	TargetTolerance param.Opt[float64] `json:"target_tolerance,omitzero"`
	// Sharp Band Center
	TiltCenter param.Opt[float64] `json:"tilt_center,omitzero"`
	// Sharp Band Width
	TiltWidth param.Opt[float64] `json:"tilt_width,omitzero"`
	// Blur Type
	//
	// Any of "gaussian", "box", "motion", "radial", "bilateral", "bokeh", "tiltshift",
	// "targetcolor".
	BlurType string `json:"blur_type,omitzero"`
	// Bokeh Shape
	//
	// Any of "circle", "hexagon", "pentagon".
	BokehShape string `json:"bokeh_shape,omitzero"`
	// Radial Mode
	//
	// Any of "zoom", "spin".
	RadialMode string `json:"radial_mode,omitzero"`
	// Band Orientation
	//
	// Any of "horizontal", "vertical".
	TiltOrientation string `json:"tilt_orientation,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject12Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject12Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject12Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject12Params](
		"blur_type", "gaussian", "box", "motion", "radial", "bilateral", "bokeh", "tiltshift", "targetcolor",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject12Params](
		"bokeh_shape", "circle", "hexagon", "pentagon",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject12Params](
		"radial_mode", "zoom", "spin",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject12Params](
		"tilt_orientation", "horizontal", "vertical",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject13 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject13Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject13Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "duplicate-image".
	ActionID constant.DuplicateImage `json:"action_id" default:"duplicate-image"`
	paramObj
}

func (r ActionRunParamsBodyObject13) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject13
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject13) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject13Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject13Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject13Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject13Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject13Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject13Params struct {
	// Copies
	Count param.Opt[float64] `json:"count,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject13Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject13Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject13Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject14 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject14Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject14Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "side-by-side-composite".
	ActionID constant.SideBySideComposite `json:"action_id" default:"side-by-side-composite"`
	paramObj
}

func (r ActionRunParamsBodyObject14) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject14
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject14) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject14Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject14Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject14Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject14Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject14Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject14Params struct {
	// Background
	Background param.Opt[string] `json:"background,omitzero"`
	// Gap (px)
	Gap param.Opt[float64] `json:"gap,omitzero"`
	// Layout
	//
	// Any of "auto", "horizontal-2", "horizontal-3", "vertical-2", "vertical-3",
	// "grid-2x2".
	Layout string `json:"layout,omitzero"`
	// Size Match
	//
	// Any of "match-shortest", "match-largest", "pad-to-largest".
	Normalize string `json:"normalize,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject14Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject14Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject14Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject14Params](
		"layout", "auto", "horizontal-2", "horizontal-3", "vertical-2", "vertical-3", "grid-2x2",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject14Params](
		"normalize", "match-shortest", "match-largest", "pad-to-largest",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject15 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject15Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject15Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "add-shape-to-image".
	ActionID constant.AddShapeToImage `json:"action_id" default:"add-shape-to-image"`
	paramObj
}

func (r ActionRunParamsBodyObject15) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject15
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject15) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject15Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject15Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject15Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject15Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject15Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject15Params struct {
	// Fill Color
	FillColor param.Opt[string] `json:"fill_color,omitzero"`
	// Fill Opacity
	FillOpacity param.Opt[float64] `json:"fill_opacity,omitzero"`
	// Rotation (deg)
	Rotation param.Opt[float64] `json:"rotation,omitzero"`
	// Stroke Color
	StrokeColor param.Opt[string] `json:"stroke_color,omitzero"`
	// Stroke Width (px)
	StrokeWidth param.Opt[float64] `json:"stroke_width,omitzero"`
	// Center
	Center ActionRunParamsBodyObject15ParamsCenter `json:"center,omitzero"`
	// Shape
	//
	// Any of "rectangle", "ellipse", "triangle", "star", "hexagon".
	Shape string `json:"shape,omitzero"`
	// Size
	Size ActionRunParamsBodyObject15ParamsSize `json:"size,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject15Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject15Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject15Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject15Params](
		"shape", "rectangle", "ellipse", "triangle", "star", "hexagon",
	)
}

// Center
//
// The properties X, Y are required.
type ActionRunParamsBodyObject15ParamsCenter struct {
	X float64 `json:"x" api:"required"`
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r ActionRunParamsBodyObject15ParamsCenter) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject15ParamsCenter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject15ParamsCenter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size
//
// The properties X, Y are required.
type ActionRunParamsBodyObject15ParamsSize struct {
	X float64 `json:"x" api:"required"`
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r ActionRunParamsBodyObject15ParamsSize) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject15ParamsSize
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject15ParamsSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject16 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject16Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject16Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "generate-shape-image".
	ActionID constant.GenerateShapeImage `json:"action_id" default:"generate-shape-image"`
	paramObj
}

func (r ActionRunParamsBodyObject16) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject16
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject16) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject16Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject16Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject16Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject16Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject16Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject16Params struct {
	// Background
	Background param.Opt[string] `json:"background,omitzero"`
	// Fill Color
	FillColor param.Opt[string] `json:"fill_color,omitzero"`
	// Fill Opacity
	FillOpacity param.Opt[float64] `json:"fill_opacity,omitzero"`
	// Height (px)
	Height param.Opt[float64] `json:"height,omitzero"`
	// Rotation (deg)
	Rotation param.Opt[float64] `json:"rotation,omitzero"`
	// Stroke Color
	StrokeColor param.Opt[string] `json:"stroke_color,omitzero"`
	// Stroke Width (px)
	StrokeWidth param.Opt[float64] `json:"stroke_width,omitzero"`
	// Width (px)
	Width param.Opt[float64] `json:"width,omitzero"`
	// Center
	Center ActionRunParamsBodyObject16ParamsCenter `json:"center,omitzero"`
	// Shape
	//
	// Any of "rectangle", "ellipse", "triangle", "star", "hexagon".
	Shape string `json:"shape,omitzero"`
	// Size
	Size ActionRunParamsBodyObject16ParamsSize `json:"size,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject16Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject16Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject16Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject16Params](
		"shape", "rectangle", "ellipse", "triangle", "star", "hexagon",
	)
}

// Center
//
// The properties X, Y are required.
type ActionRunParamsBodyObject16ParamsCenter struct {
	X float64 `json:"x" api:"required"`
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r ActionRunParamsBodyObject16ParamsCenter) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject16ParamsCenter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject16ParamsCenter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size
//
// The properties X, Y are required.
type ActionRunParamsBodyObject16ParamsSize struct {
	X float64 `json:"x" api:"required"`
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r ActionRunParamsBodyObject16ParamsSize) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject16ParamsSize
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject16ParamsSize) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject17 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject17Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject17Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "add-text-to-image".
	ActionID constant.AddTextToImage `json:"action_id" default:"add-text-to-image"`
	paramObj
}

func (r ActionRunParamsBodyObject17) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject17
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject17) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject17Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject17Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject17Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject17Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject17Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject17Params struct {
	// Text Color
	Color param.Opt[string] `json:"color,omitzero"`
	// Font Size (px)
	FontSize param.Opt[float64] `json:"font_size,omitzero"`
	// Margin (px)
	Margin param.Opt[float64] `json:"margin,omitzero"`
	// Opacity
	Opacity param.Opt[float64] `json:"opacity,omitzero"`
	// Shadow
	Shadow param.Opt[bool] `json:"shadow,omitzero"`
	// Text
	Text param.Opt[string] `json:"text,omitzero"`
	// Font
	//
	// Any of "sans", "sans-bold", "sans-italic", "serif", "serif-bold", "mono",
	// "mono-bold".
	FontFamily string `json:"font_family,omitzero"`
	// Position
	//
	// Any of "top-left", "top-center", "top-right", "middle-left", "center",
	// "middle-right", "bottom-left", "bottom-center", "bottom-right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject17Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject17Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject17Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject17Params](
		"font_family", "sans", "sans-bold", "sans-italic", "serif", "serif-bold", "mono", "mono-bold",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject17Params](
		"position", "top-left", "top-center", "top-right", "middle-left", "center", "middle-right", "bottom-left", "bottom-center", "bottom-right",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject18 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject18Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject18Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "generate-text-image".
	ActionID constant.GenerateTextImage `json:"action_id" default:"generate-text-image"`
	paramObj
}

func (r ActionRunParamsBodyObject18) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject18
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject18) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject18Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject18Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject18Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject18Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject18Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject18Params struct {
	// Background
	Background param.Opt[string] `json:"background,omitzero"`
	// Text Color
	Color param.Opt[string] `json:"color,omitzero"`
	// Font Size (px)
	FontSize param.Opt[float64] `json:"font_size,omitzero"`
	// Height (px)
	Height param.Opt[float64] `json:"height,omitzero"`
	// Margin (px)
	Margin param.Opt[float64] `json:"margin,omitzero"`
	// Text
	Text param.Opt[string] `json:"text,omitzero"`
	// Width (px)
	Width param.Opt[float64] `json:"width,omitzero"`
	// Font
	//
	// Any of "sans", "sans-bold", "sans-italic", "serif", "serif-bold", "mono",
	// "mono-bold".
	FontFamily string `json:"font_family,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject18Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject18Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject18Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject18Params](
		"font_family", "sans", "sans-bold", "sans-italic", "serif", "serif-bold", "mono", "mono-bold",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject19 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject19Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject19Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "qr-code-generator".
	ActionID constant.QrCodeGenerator `json:"action_id" default:"qr-code-generator"`
	paramObj
}

func (r ActionRunParamsBodyObject19) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject19
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject19) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject19Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject19Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject19Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject19Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject19Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject19Params struct {
	// Background Color
	BgColor param.Opt[string] `json:"bg_color,omitzero"`
	// Quiet-Zone Border (modules)
	Border param.Opt[float64] `json:"border,omitzero"`
	// Foreground Color
	FgColor param.Opt[string] `json:"fg_color,omitzero"`
	// Size (px)
	Size param.Opt[float64] `json:"size,omitzero"`
	// Error Correction
	//
	// Any of "L", "M", "Q", "H".
	ErrorCorrection string `json:"error_correction,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject19Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject19Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject19Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject19Params](
		"error_correction", "L", "M", "Q", "H",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject20 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject20Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject20Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "stitch-videos".
	ActionID constant.StitchVideos `json:"action_id" default:"stitch-videos"`
	paramObj
}

func (r ActionRunParamsBodyObject20) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject20
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject20) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject20Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject20Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject20Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject20Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject20Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject20Params struct {
	// Background Color
	BackgroundColor param.Opt[string] `json:"background_color,omitzero"`
	// Blur Amount
	BlurAmount param.Opt[float64] `json:"blur_amount,omitzero"`
	// Transition Duration (s)
	TransitionDuration param.Opt[float64] `json:"transition_duration,omitzero"`
	// Output Aspect Ratio
	//
	// Any of "auto", "21:9", "16:9", "3:2", "4:3", "5:4", "1:1", "4:5", "3:4", "2:3",
	// "9:16", "9:21".
	AspectRatio string `json:"aspect_ratio,omitzero"`
	// Background Mode
	//
	// Any of "solid", "blur".
	BackgroundMode string `json:"background_mode,omitzero"`
	// Fit Mode
	//
	// Any of "contain", "cover".
	FitMode string `json:"fit_mode,omitzero"`
	// Transition
	//
	// Any of "none", "fade", "wipeleft", "wiperight".
	Transition string `json:"transition,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject20Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject20Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject20Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject20Params](
		"aspect_ratio", "auto", "21:9", "16:9", "3:2", "4:3", "5:4", "1:1", "4:5", "3:4", "2:3", "9:16", "9:21",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject20Params](
		"background_mode", "solid", "blur",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject20Params](
		"fit_mode", "contain", "cover",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject20Params](
		"transition", "none", "fade", "wipeleft", "wiperight",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject21 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject21Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject21Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "split-video".
	ActionID constant.SplitVideo `json:"action_id" default:"split-video"`
	paramObj
}

func (r ActionRunParamsBodyObject21) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject21
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject21) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject21Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject21Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject21Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject21Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject21Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject21Params struct {
	// Scene Sensitivity
	SceneSensitivity param.Opt[float64] `json:"scene_sensitivity,omitzero"`
	// Segment Duration (s)
	SegmentDuration param.Opt[float64] `json:"segment_duration,omitzero"`
	// Number of Segments
	Segments param.Opt[float64] `json:"segments,omitzero"`
	// Strip Audio
	StripAudio param.Opt[bool] `json:"strip_audio,omitzero"`
	// Trim Handles (s)
	TrimHandles param.Opt[float64] `json:"trim_handles,omitzero"`
	// Split Mode
	//
	// Any of "equal", "duration", "scene".
	SplitMode string `json:"split_mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject21Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject21Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject21Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject21Params](
		"split_mode", "equal", "duration", "scene",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject22 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject22Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject22Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "extract-video-frames".
	ActionID constant.ExtractVideoFrames `json:"action_id" default:"extract-video-frames"`
	paramObj
}

func (r ActionRunParamsBodyObject22) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject22
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject22) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject22Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject22Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject22Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject22Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject22Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject22Params struct {
	// Every (s)
	EverySeconds param.Opt[float64] `json:"every_seconds,omitzero"`
	// Frame Count
	FrameCount param.Opt[float64] `json:"frame_count,omitzero"`
	// Scene Sensitivity
	SceneSensitivity param.Opt[float64] `json:"scene_sensitivity,omitzero"`
	// Time (%)
	TimePercent param.Opt[float64] `json:"time_percent,omitzero"`
	// Frames
	//
	// Any of "single", "evenly", "interval", "scene".
	Mode string `json:"mode,omitzero"`
	// Video Range (%)
	Range ActionRunParamsBodyObject22ParamsRange `json:"range,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject22Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject22Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject22Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject22Params](
		"mode", "single", "evenly", "interval", "scene",
	)
}

// Video Range (%)
//
// The properties Max, Min are required.
type ActionRunParamsBodyObject22ParamsRange struct {
	Max float64 `json:"max" api:"required"`
	Min float64 `json:"min" api:"required"`
	paramObj
}

func (r ActionRunParamsBodyObject22ParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject22ParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject22ParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject23 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject23Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject23Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "color-grade-video".
	ActionID constant.ColorGradeVideo `json:"action_id" default:"color-grade-video"`
	paramObj
}

func (r ActionRunParamsBodyObject23) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject23
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject23) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject23Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject23Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject23Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject23Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject23Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject23Params struct {
	// Brightness
	Brightness param.Opt[float64] `json:"brightness,omitzero"`
	// Contrast
	Contrast param.Opt[float64] `json:"contrast,omitzero"`
	// Gamma
	Gamma param.Opt[float64] `json:"gamma,omitzero"`
	// Saturation
	Saturation param.Opt[float64] `json:"saturation,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject23Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject23Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject23Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject24 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject24Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject24Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "video-to-frame-grid".
	ActionID constant.VideoToFrameGrid `json:"action_id" default:"video-to-frame-grid"`
	paramObj
}

func (r ActionRunParamsBodyObject24) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject24
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject24) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject24Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject24Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject24Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject24Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject24Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject24Params struct {
	// Background Color
	Background param.Opt[string] `json:"background,omitzero"`
	// Cell Width (px)
	CellWidth param.Opt[float64] `json:"cell_width,omitzero"`
	// Columns
	Cols param.Opt[float64] `json:"cols,omitzero"`
	// Gap Between Cells (px)
	Gap param.Opt[float64] `json:"gap,omitzero"`
	// Rows
	Rows param.Opt[float64] `json:"rows,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject24Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject24Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject24Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject25 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject25Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject25Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "boomerang-video".
	ActionID constant.BoomerangVideo `json:"action_id" default:"boomerang-video"`
	paramObj
}

func (r ActionRunParamsBodyObject25) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject25
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject25) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject25Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject25Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject25Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject25Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject25Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject25Params struct {
	// Include Audio
	IncludeAudio param.Opt[bool] `json:"include_audio,omitzero"`
	// Speed
	Speed param.Opt[float64] `json:"speed,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject25Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject25Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject25Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject26 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject26Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject26Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "reverse-video".
	ActionID constant.ReverseVideo `json:"action_id" default:"reverse-video"`
	paramObj
}

func (r ActionRunParamsBodyObject26) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject26
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject26) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject26Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject26Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject26Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject26Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject26Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject26Params struct {
	// Audio
	//
	// Any of "strip", "reverse", "keep".
	AudioMode string `json:"audio_mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject26Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject26Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject26Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject26Params](
		"audio_mode", "strip", "reverse", "keep",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject27 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject27Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject27Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "video-to-long-exposure".
	ActionID constant.VideoToLongExposure `json:"action_id" default:"video-to-long-exposure"`
	paramObj
}

func (r ActionRunParamsBodyObject27) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject27
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject27) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject27Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject27Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject27Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject27Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject27Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject27Params struct {
	// Frame Stride (sample every N frames)
	FrameStride param.Opt[float64] `json:"frame_stride,omitzero"`
	// Max Frames to Sample
	MaxFrames param.Opt[float64] `json:"max_frames,omitzero"`
	// Blend Mode
	//
	// Any of "average", "lighten", "darken".
	BlendMode string `json:"blend_mode,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject27Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject27Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject27Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject27Params](
		"blend_mode", "average", "lighten", "darken",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject28 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject28Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject28Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "video-effect".
	ActionID constant.VideoEffect `json:"action_id" default:"video-effect"`
	paramObj
}

func (r ActionRunParamsBodyObject28) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject28
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject28) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject28Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject28Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject28Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject28Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject28Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject28Params struct {
	// Channel Offset (px)
	ChromaticOffset param.Opt[float64] `json:"chromatic_offset,omitzero"`
	// Grain Strength
	GrainStrength param.Opt[float64] `json:"grain_strength,omitzero"`
	// Block Size (px)
	PixelBlockSize param.Opt[float64] `json:"pixel_block_size,omitzero"`
	// Shake Amount (px)
	ShakeAmount param.Opt[float64] `json:"shake_amount,omitzero"`
	// Vignette Falloff
	VignetteAngle param.Opt[float64] `json:"vignette_angle,omitzero"`
	// Effect
	//
	// Any of "vignette", "grain", "pixelate", "shake", "chromatic", "vhs".
	Effect string `json:"effect,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject28Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject28Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject28Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject28Params](
		"effect", "vignette", "grain", "pixelate", "shake", "chromatic", "vhs",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject29 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject29Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject29Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "color-filter-video".
	ActionID constant.ColorFilterVideo `json:"action_id" default:"color-filter-video"`
	paramObj
}

func (r ActionRunParamsBodyObject29) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject29
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject29) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject29Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject29Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject29Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject29Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject29Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject29Params struct {
	// B&W Threshold
	BwThreshold param.Opt[float64] `json:"bw_threshold,omitzero"`
	// Grain
	Grain param.Opt[float64] `json:"grain,omitzero"`
	// Posterize Bits
	PosterizeBits param.Opt[float64] `json:"posterize_bits,omitzero"`
	// Solarize Threshold
	SolarizeThreshold param.Opt[float64] `json:"solarize_threshold,omitzero"`
	// Filter
	//
	// Any of "grayscale", "sepia", "invert", "bw", "posterize", "solarize",
	// "clarendon", "moon", "nashville", "noir", "fade", "crossprocess".
	Filter string `json:"filter,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject29Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject29Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject29Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject29Params](
		"filter", "grayscale", "sepia", "invert", "bw", "posterize", "solarize", "clarendon", "moon", "nashville", "noir", "fade", "crossprocess",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject30 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject30Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject30Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "speed-up-video".
	ActionID constant.SpeedUpVideo `json:"action_id" default:"speed-up-video"`
	paramObj
}

func (r ActionRunParamsBodyObject30) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject30
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject30) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject30Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject30Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject30Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject30Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject30Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject30Params struct {
	// Speed Factor (x)
	Factor param.Opt[float64] `json:"factor,omitzero"`
	// Keep Audio (pitch-preserved)
	KeepAudio param.Opt[bool] `json:"keep_audio,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject30Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject30Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject30Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject31 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject31Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject31Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "slow-down-video".
	ActionID constant.SlowDownVideo `json:"action_id" default:"slow-down-video"`
	paramObj
}

func (r ActionRunParamsBodyObject31) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject31
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject31) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject31Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject31Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject31Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject31Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject31Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject31Params struct {
	// Slow Factor (x)
	Factor param.Opt[float64] `json:"factor,omitzero"`
	// Keep Audio (pitch-preserved)
	KeepAudio param.Opt[bool] `json:"keep_audio,omitzero"`
	// Smoothing
	//
	// Any of "off", "blend", "motion".
	Smoothing string `json:"smoothing,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject31Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject31Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject31Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject31Params](
		"smoothing", "off", "blend", "motion",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject32 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject32Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject32Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "duplicate-video".
	ActionID constant.DuplicateVideo `json:"action_id" default:"duplicate-video"`
	paramObj
}

func (r ActionRunParamsBodyObject32) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject32
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject32) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject32Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject32Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject32Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject32Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject32Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject32Params struct {
	// Copies
	Count param.Opt[float64] `json:"count,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject32Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject32Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject32Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject33 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject33Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject33Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "greenscreen-video".
	ActionID constant.GreenscreenVideo `json:"action_id" default:"greenscreen-video"`
	paramObj
}

func (r ActionRunParamsBodyObject33) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject33
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject33) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject33Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject33Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject33Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject33Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject33Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject33Params struct {
	// Edge Blend (softness)
	Blend param.Opt[float64] `json:"blend,omitzero"`
	// Custom Background Color
	CustomColor param.Opt[string] `json:"custom_color,omitzero"`
	// Similarity (how close to key color counts)
	Similarity param.Opt[float64] `json:"similarity,omitzero"`
	// Spill Suppression
	Spill param.Opt[bool] `json:"spill,omitzero"`
	// Background Preset
	//
	// Any of "green", "blue", "custom".
	ColorPreset string `json:"color_preset,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject33Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject33Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject33Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject33Params](
		"color_preset", "green", "blue", "custom",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject34 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject34Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject34Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "resize-video".
	ActionID constant.ResizeVideo `json:"action_id" default:"resize-video"`
	paramObj
}

func (r ActionRunParamsBodyObject34) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject34
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject34) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject34Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject34Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject34Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject34Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject34Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject34Params struct {
	// Custom Resolution (px)
	CustomResolution param.Opt[float64] `json:"custom_resolution,omitzero"`
	// Resolution
	//
	// Any of "240", "360", "480", "540", "720", "1080", "1440", "2160", "custom".
	Resolution string `json:"resolution,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject34Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject34Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject34Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject34Params](
		"resolution", "240", "360", "480", "540", "720", "1080", "1440", "2160", "custom",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject35 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject35Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject35Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as "change-video-ar".
	ActionID constant.ChangeVideoAr `json:"action_id" default:"change-video-ar"`
	paramObj
}

func (r ActionRunParamsBodyObject35) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject35
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject35) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject35Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject35Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject35Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject35Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject35Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject35Params struct {
	// Blur Amount
	BlurAmount param.Opt[float64] `json:"blur_amount,omitzero"`
	// Pad Color
	PadColor param.Opt[string] `json:"pad_color,omitzero"`
	// Aspect Ratio
	//
	// Any of "1:1", "16:9", "9:16", "4:3", "3:4", "3:2", "2:3", "21:9".
	AspectRatio string `json:"aspect_ratio,omitzero"`
	// Background
	//
	// Any of "solid", "blur".
	BackgroundMode string `json:"background_mode,omitzero"`
	// Fit
	//
	// Any of "crop", "pad".
	Fit string `json:"fit,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject35Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject35Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject35Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject35Params](
		"aspect_ratio", "1:1", "16:9", "9:16", "4:3", "3:4", "3:2", "2:3", "21:9",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject35Params](
		"background_mode", "solid", "blur",
	)
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject35Params](
		"fit", "crop", "pad",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject36 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject36Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject36Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "split-audio-from-video".
	ActionID constant.SplitAudioFromVideo `json:"action_id" default:"split-audio-from-video"`
	paramObj
}

func (r ActionRunParamsBodyObject36) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject36
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject36) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject36Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject36Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject36Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject36Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject36Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject36Params struct {
	// Audio Format
	//
	// Any of "auto", "m4a", "mp3", "wav".
	AudioFormat string `json:"audio_format,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject36Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject36Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject36Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject36Params](
		"audio_format", "auto", "m4a", "mp3", "wav",
	)
}

// The properties ActionID, Inputs, ProjectID, WorkspaceID are required.
type ActionRunParamsBodyObject37 struct {
	// Action inputs. Direct action runs are headless and read inputs from URLs/text
	// rather than canvas edges.
	Inputs []ActionRunParamsBodyObject37Input `json:"inputs,omitzero" api:"required"`
	// Project identifier. Use the public API ID returned by list projects; it must
	// start with prj\_. Used for ownership, billing, and run history. Direct action
	// runs do not mutate the project canvas.
	ProjectID string `json:"project_id" api:"required"`
	// Workspace identifier. Use the public API ID returned by list workspaces; it must
	// start with ws\_.
	WorkspaceID string `json:"workspace_id" api:"required"`
	// Action parameters
	Params ActionRunParamsBodyObject37Params `json:"params,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "merge-audio-into-video".
	ActionID constant.MergeAudioIntoVideo `json:"action_id" default:"merge-audio-into-video"`
	paramObj
}

func (r ActionRunParamsBodyObject37) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject37
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject37) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ActionRunParamsBodyObject37Input struct {
	// Action input type
	//
	// Any of "image", "video", "audio", "text".
	Type string `json:"type,omitzero" api:"required"`
	// Optional input name
	Name param.Opt[string] `json:"name,omitzero"`
	// Input text value
	Text param.Opt[string] `json:"text,omitzero"`
	// Input media URL
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r ActionRunParamsBodyObject37Input) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject37Input
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject37Input) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject37Input](
		"type", "image", "video", "audio", "text",
	)
}

// Action parameters
type ActionRunParamsBodyObject37Params struct {
	// Duration
	//
	// Any of "shortest", "video", "audio".
	Duration string `json:"duration,omitzero"`
	paramObj
}

func (r ActionRunParamsBodyObject37Params) MarshalJSON() (data []byte, err error) {
	type shadow ActionRunParamsBodyObject37Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRunParamsBodyObject37Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ActionRunParamsBodyObject37Params](
		"duration", "shortest", "video", "audio",
	)
}
