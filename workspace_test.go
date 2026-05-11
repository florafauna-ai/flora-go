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

func TestWorkspaceList(t *testing.T) {
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
	_, err := client.Workspaces.List(context.TODO())
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
