// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/florafauna-ai/flora-go"
	"github.com/florafauna-ai/flora-go/internal/testutil"
	"github.com/florafauna-ai/flora-go/option"
)

func TestTechniqueRunNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := flora.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Techniques.Runs.New(
		context.TODO(),
		"art-directors-critique",
		flora.TechniqueRunNewParams{
			Inputs: []flora.TechniqueRunNewParamsInput{{
				ID:    "id",
				Type:  "text",
				Value: "value",
			}},
			Mode:           flora.TechniqueRunNewParamsModeAsync,
			CallbackURL:    flora.String("https://example.com"),
			IdempotencyKey: flora.String("idempotency_key"),
		},
	)
	if err != nil {
		var apierr *flora.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTechniqueRunGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := flora.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Techniques.Runs.Get(
		context.TODO(),
		"run_abc123",
		flora.TechniqueRunGetParams{
			TechniqueID: "art-directors-critique",
		},
	)
	if err != nil {
		var apierr *flora.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
