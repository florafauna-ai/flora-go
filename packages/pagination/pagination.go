// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/florafauna-ai/flora-go/internal/apijson"
	"github.com/florafauna-ai/flora-go/internal/requestconfig"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/florafauna-ai/flora-go/packages/param"
	"github.com/florafauna-ai/flora-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type ProjectsCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectsCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *ProjectsCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectsCursorPage[T any] struct {
	Projects []T                    `json:"projects"`
	Meta     ProjectsCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Projects    respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ProjectsCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ProjectsCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ProjectsCursorPage[T]) GetNextPage() (res *ProjectsCursorPage[T], err error) {
	if len(r.Projects) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ProjectsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ProjectsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ProjectsCursorPageAutoPager[T any] struct {
	page *ProjectsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewProjectsCursorPageAutoPager[T any](page *ProjectsCursorPage[T], err error) *ProjectsCursorPageAutoPager[T] {
	return &ProjectsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ProjectsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Projects) == 0 {
		return false
	}
	if r.idx >= len(r.page.Projects) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Projects) == 0 {
			return false
		}
	}
	r.cur = r.page.Projects[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ProjectsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ProjectsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ProjectsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type TechniquesCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniquesCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *TechniquesCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniquesCursorPage[T any] struct {
	Techniques []T                      `json:"techniques"`
	Meta       TechniquesCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Techniques  respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TechniquesCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TechniquesCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TechniquesCursorPage[T]) GetNextPage() (res *TechniquesCursorPage[T], err error) {
	if len(r.Techniques) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TechniquesCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TechniquesCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TechniquesCursorPageAutoPager[T any] struct {
	page *TechniquesCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTechniquesCursorPageAutoPager[T any](page *TechniquesCursorPage[T], err error) *TechniquesCursorPageAutoPager[T] {
	return &TechniquesCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TechniquesCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Techniques) == 0 {
		return false
	}
	if r.idx >= len(r.page.Techniques) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Techniques) == 0 {
			return false
		}
	}
	r.cur = r.page.Techniques[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TechniquesCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TechniquesCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TechniquesCursorPageAutoPager[T]) Index() int {
	return r.run
}

type GenerationsCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationsCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *GenerationsCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationsCursorPage[T any] struct {
	Generations []T                       `json:"generations"`
	Meta        GenerationsCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Generations respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r GenerationsCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *GenerationsCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *GenerationsCursorPage[T]) GetNextPage() (res *GenerationsCursorPage[T], err error) {
	if len(r.Generations) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *GenerationsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &GenerationsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type GenerationsCursorPageAutoPager[T any] struct {
	page *GenerationsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewGenerationsCursorPageAutoPager[T any](page *GenerationsCursorPage[T], err error) *GenerationsCursorPageAutoPager[T] {
	return &GenerationsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *GenerationsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Generations) == 0 {
		return false
	}
	if r.idx >= len(r.page.Generations) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Generations) == 0 {
			return false
		}
	}
	r.cur = r.page.Generations[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *GenerationsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *GenerationsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *GenerationsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type TechniqueRunsCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TechniqueRunsCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunsCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TechniqueRunsCursorPage[T any] struct {
	TechniqueRuns []T                         `json:"technique_runs"`
	Meta          TechniqueRunsCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TechniqueRuns respjson.Field
		Meta          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TechniqueRunsCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TechniqueRunsCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TechniqueRunsCursorPage[T]) GetNextPage() (res *TechniqueRunsCursorPage[T], err error) {
	if len(r.TechniqueRuns) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *TechniqueRunsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TechniqueRunsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TechniqueRunsCursorPageAutoPager[T any] struct {
	page *TechniqueRunsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTechniqueRunsCursorPageAutoPager[T any](page *TechniqueRunsCursorPage[T], err error) *TechniqueRunsCursorPageAutoPager[T] {
	return &TechniqueRunsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TechniqueRunsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.TechniqueRuns) == 0 {
		return false
	}
	if r.idx >= len(r.page.TechniqueRuns) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.TechniqueRuns) == 0 {
			return false
		}
	}
	r.cur = r.page.TechniqueRuns[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TechniqueRunsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TechniqueRunsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TechniqueRunsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type AssetsCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetsCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *AssetsCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetsCursorPage[T any] struct {
	Assets []T                  `json:"assets"`
	Meta   AssetsCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r AssetsCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *AssetsCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *AssetsCursorPage[T]) GetNextPage() (res *AssetsCursorPage[T], err error) {
	if len(r.Assets) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *AssetsCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &AssetsCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type AssetsCursorPageAutoPager[T any] struct {
	page *AssetsCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewAssetsCursorPageAutoPager[T any](page *AssetsCursorPage[T], err error) *AssetsCursorPageAutoPager[T] {
	return &AssetsCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *AssetsCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Assets) == 0 {
		return false
	}
	if r.idx >= len(r.page.Assets) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Assets) == 0 {
			return false
		}
	}
	r.cur = r.page.Assets[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *AssetsCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *AssetsCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *AssetsCursorPageAutoPager[T]) Index() int {
	return r.run
}

type CanvasNodesCursorPageMeta struct {
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CanvasNodesCursorPageMeta) RawJSON() string { return r.JSON.raw }
func (r *CanvasNodesCursorPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CanvasNodesCursorPage[T any] struct {
	Nodes []T                       `json:"nodes"`
	Meta  CanvasNodesCursorPageMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Nodes       respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CanvasNodesCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CanvasNodesCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CanvasNodesCursorPage[T]) GetNextPage() (res *CanvasNodesCursorPage[T], err error) {
	if len(r.Nodes) == 0 {
		return nil, nil
	}
	next := r.Meta.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CanvasNodesCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CanvasNodesCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CanvasNodesCursorPageAutoPager[T any] struct {
	page *CanvasNodesCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCanvasNodesCursorPageAutoPager[T any](page *CanvasNodesCursorPage[T], err error) *CanvasNodesCursorPageAutoPager[T] {
	return &CanvasNodesCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CanvasNodesCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Nodes) == 0 {
		return false
	}
	if r.idx >= len(r.page.Nodes) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Nodes) == 0 {
			return false
		}
	}
	r.cur = r.page.Nodes[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CanvasNodesCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CanvasNodesCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CanvasNodesCursorPageAutoPager[T]) Index() int {
	return r.run
}
