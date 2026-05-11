// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package florafaunaai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/florafauna-ai-go"
	"github.com/stainless-sdks/florafauna-ai-go/internal/testutil"
	"github.com/stainless-sdks/florafauna-ai-go/option"
)

func TestTechniqueRunGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := florafaunaai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Techniques.Runs.Get(
		context.TODO(),
		"run_abc123",
		florafaunaai.TechniqueRunGetParams{
			TechniqueID: "tech_def_abc123",
		},
	)
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTechniqueRunStartWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := florafaunaai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Techniques.Runs.Start(
		context.TODO(),
		"tech_def_abc123",
		florafaunaai.TechniqueRunStartParams{
			Inputs: []florafaunaai.TechniqueRunStartParamsInput{{
				ID:    "id",
				Type:  "imageUrl",
				Value: "value",
			}},
			Mode:           florafaunaai.TechniqueRunStartParamsModeAsync,
			CallbackURL:    florafaunaai.String("https://example.com"),
			IdempotencyKey: florafaunaai.String("idempotency_key"),
		},
	)
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
