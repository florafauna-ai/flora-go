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

func TestAssetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.New(context.TODO(), florafaunaai.AssetNewParams{
		Source:      "signed-url",
		WorkspaceID: "ws_abc123",
		ContentType: florafaunaai.String("image/png"),
		FileName:    florafaunaai.String("hero.png"),
		Folder:      florafaunaai.String("campaign-assets"),
	})
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssetGet(t *testing.T) {
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
	_, err := client.Assets.Get(context.TODO(), "asset_abc123")
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.List(context.TODO(), florafaunaai.AssetListParams{
		Cursor:      florafaunaai.String("cursor"),
		Limit:       florafaunaai.Int(1),
		ProjectID:   florafaunaai.String("prj_abc123"),
		Query:       florafaunaai.String("logo"),
		WorkspaceID: florafaunaai.String("ws_abc123"),
	})
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssetCompleteUpload(t *testing.T) {
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
	_, err := client.Assets.CompleteUpload(context.TODO(), "asset_abc123")
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAssetRetryUpload(t *testing.T) {
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
	_, err := client.Assets.RetryUpload(context.TODO(), "asset_abc123")
	if err != nil {
		var apierr *florafaunaai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
