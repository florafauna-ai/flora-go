// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flora_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/florafauna-ai-go"
	"github.com/stainless-sdks/florafauna-ai-go/internal/testutil"
	"github.com/stainless-sdks/florafauna-ai-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Techniques.List(context.TODO(), flora.TechniqueListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, technique := range page.Techniques {
		t.Logf("%+v\n", technique.TechniqueID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, technique := range page.Techniques {
			t.Logf("%+v\n", technique.TechniqueID)
		}
	}
}
