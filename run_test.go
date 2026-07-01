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

func TestRunStartGenerationWithOptionalParams(t *testing.T) {
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
	_, err := client.Runs.StartGeneration(context.TODO(), flora.RunStartGenerationParams{
		ProjectID:   "prj_abc123",
		Prompt:      "A cinematic product photo of a ceramic mug on a sunlit table",
		Type:        flora.RunStartGenerationParamsTypeImage,
		WorkspaceID: "ws_abc123",
		CallbackURL: flora.String("https://example.com"),
		Model:       flora.String("t2i-flux-2-pro"),
		Params: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *flora.Error
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
	client := flora.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Runs.StartTechnique(context.TODO(), flora.RunStartTechniqueParams{
		Inputs: map[string]any{
			"foo": "bar",
		},
		TechniqueID: "tech_abcd1234",
		WorkspaceID: "ws_abc123",
	})
	if err != nil {
		var apierr *flora.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
