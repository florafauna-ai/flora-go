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

func TestTechniqueRunNew(t *testing.T) {
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
	_, err := client.Techniques.Runs.New(context.TODO(), "tech_def_abc123")
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
			TechniqueID: "tech_def_abc123",
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
