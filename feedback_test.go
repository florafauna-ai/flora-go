// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/florafauna-ai-go"
	"github.com/stainless-sdks/florafauna-ai-go/internal/testutil"
	"github.com/stainless-sdks/florafauna-ai-go/option"
)

func TestFeedbackRecordWithOptionalParams(t *testing.T) {
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
	_, err := client.Feedback.Record(context.TODO(), flora.FeedbackRecordParams{
		Detail:         "I want to export all generated campaign images at once.",
		Kind:           flora.FeedbackRecordParamsKindFeatureRequest,
		Summary:        "Need batch export support",
		AttemptedTools: []string{"generate_image"},
		ProjectID:      flora.String("prj_abc123"),
		RunID:          flora.String("run_abc123"),
		WorkspaceID:    flora.String("ws_abc123"),
	})
	if err != nil {
		var apierr *flora.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
