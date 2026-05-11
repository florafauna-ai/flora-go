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

func TestRunStartGenerationWithOptionalParams(t *testing.T) {
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
	_, err := client.Runs.StartGeneration(context.TODO(), florafaunaai.RunStartGenerationParams{
		ProjectID:   "prj_abc123",
		Prompt:      "A cinematic product photo of a ceramic mug on a sunlit table",
		Type:        florafaunaai.RunStartGenerationParamsTypeImage,
		WorkspaceID: "ws_abc123",
		Model:       florafaunaai.String("t2i-flux-2-pro"),
		Params: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunStartTechnique(t *testing.T) {
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
	_, err := client.Runs.StartTechnique(context.TODO(), florafaunaai.RunStartTechniqueParams{
		Inputs: map[string]any{
			"foo": "bar",
		},
		TechniqueID: "tech_abcd1234",
		WorkspaceID: "ws_abc123",
	})
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
