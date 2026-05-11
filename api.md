# Workspaces

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#WorkspaceListResponse">WorkspaceListResponse</a>

Methods:

- <code title="get /workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#WorkspaceListResponse">WorkspaceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListResponse">ProjectListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListNodesResponse">ProjectListNodesResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListParams">ProjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/nodes">client.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectService.ListNodes">ListNodes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListNodesParams">ProjectListNodesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectListNodesResponse">ProjectListNodesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectAssetAttachResponse">ProjectAssetAttachResponse</a>

Methods:

- <code title="post /projects/{projectId}/assets/{assetId}/attach">client.Projects.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectAssetService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectAssetAttachParams">ProjectAssetAttachParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ProjectAssetAttachResponse">ProjectAssetAttachResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ModelListParams">ModelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Techniques

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueGetResponse">TechniqueGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueListResponse">TechniqueListResponse</a>

Methods:

- <code title="get /techniques/{techniqueId}">client.Techniques.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, techniqueID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueGetResponse">TechniqueGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /techniques">client.Techniques.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueListParams">TechniqueListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueListResponse">TechniqueListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunGetResponse">TechniqueRunGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunStartResponse">TechniqueRunStartResponse</a>

Methods:

- <code title="get /techniques/{techniqueId}/runs/{runId}">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunGetParams">TechniqueRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunGetResponse">TechniqueRunGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /techniques/{techniqueId}/runs">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, techniqueID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunStartParams">TechniqueRunStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#TechniqueRunStartResponse">TechniqueRunStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetNewResponse">AssetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetGetResponse">AssetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetListResponse">AssetListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetCompleteUploadResponse">AssetCompleteUploadResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetRetryUploadResponse">AssetRetryUploadResponse</a>

Methods:

- <code title="post /assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetNewParams">AssetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetNewResponse">AssetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetGetResponse">AssetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetListParams">AssetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetListResponse">AssetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /assets/{assetId}/complete">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetService.CompleteUpload">CompleteUpload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetCompleteUploadResponse">AssetCompleteUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /assets/{assetId}/retry">client.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetService.RetryUpload">RetryUpload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#AssetRetryUploadResponse">AssetRetryUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartGenerationResponse">RunStartGenerationResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartTechniqueResponse">RunStartTechniqueResponse</a>

Methods:

- <code title="post /runs/generation">client.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunService.StartGeneration">StartGeneration</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartGenerationParams">RunStartGenerationParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartGenerationResponse">RunStartGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /runs/technique">client.Runs.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunService.StartTechnique">StartTechnique</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartTechniqueParams">RunStartTechniqueParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#RunStartTechniqueResponse">RunStartTechniqueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Feedback

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#FeedbackRecordResponse">FeedbackRecordResponse</a>

Methods:

- <code title="post /feedback">client.Feedback.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#FeedbackService.Record">Record</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#FeedbackRecordParams">FeedbackRecordParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go">florafaunaai</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/florafauna-ai-go#FeedbackRecordResponse">FeedbackRecordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
