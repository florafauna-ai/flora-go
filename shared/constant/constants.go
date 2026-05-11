// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/florafauna-ai-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type File string              // Always "file"
type MultipartFormData string // Always "multipart/form-data"
type Post string              // Always "POST"
type Workspace string         // Always "workspace"

func (c File) Default() File                           { return "file" }
func (c MultipartFormData) Default() MultipartFormData { return "multipart/form-data" }
func (c Post) Default() Post                           { return "POST" }
func (c Workspace) Default() Workspace                 { return "workspace" }

func (c File) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c MultipartFormData) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Post) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Workspace) MarshalJSON() ([]byte, error)         { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
