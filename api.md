# Techniques

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueGetResponse">TechniqueGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueListResponse">TechniqueListResponse</a>

Methods:

- <code title="get /techniques/{techniqueId}">client.Techniques.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, techniqueID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueGetResponse">TechniqueGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /techniques">client.Techniques.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueListParams">TechniqueListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#TechniquesCursorPage">TechniquesCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueListResponse">TechniqueListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunNewResponse">TechniqueRunNewResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunGetResponse">TechniqueRunGetResponse</a>

Methods:

- <code title="post /techniques/{techniqueId}/runs">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, techniqueID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunNewParams">TechniqueRunNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunNewResponse">TechniqueRunNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /techniques/{techniqueId}/runs/{runId}">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunGetParams">TechniqueRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunGetResponse">TechniqueRunGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetNewResponse">AssetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetGetResponse">AssetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetListResponse">AssetListResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetCompleteResponse">AssetCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetRetryResponse">AssetRetryResponse</a>

Methods:

- <code title="post /assets">client.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetNewParams">AssetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetNewResponse">AssetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets/{assetId}">client.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetGetResponse">AssetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assets">client.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetListParams">AssetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#AssetsCursorPage">AssetsCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetListResponse">AssetListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /assets/{assetId}/complete">client.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetCompleteResponse">AssetCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /assets/{assetId}/retry">client.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetService.Retry">Retry</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#AssetRetryResponse">AssetRetryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workspaces

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#WorkspaceListResponse">WorkspaceListResponse</a>

Methods:

- <code title="get /workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#WorkspaceListResponse">WorkspaceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectNewResponse">ProjectNewResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectGetResponse">ProjectGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListResponse">ProjectListResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListNodesResponse">ProjectListNodesResponse</a>

Methods:

- <code title="post /projects">client.Projects.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectNewResponse">ProjectNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectGetResponse">ProjectGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects">client.Projects.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListParams">ProjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#ProjectsCursorPage">ProjectsCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListResponse">ProjectListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/nodes">client.Projects.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectService.ListNodes">ListNodes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListNodesParams">ProjectListNodesParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#CanvasNodesCursorPage">CanvasNodesCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectListNodesResponse">ProjectListNodesResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectAssetAttachAssetResponse">ProjectAssetAttachAssetResponse</a>

Methods:

- <code title="post /projects/{projectId}/assets/{assetId}/attach">client.Projects.Assets.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectAssetService.AttachAsset">AttachAsset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectAssetAttachAssetParams">ProjectAssetAttachAssetParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectAssetAttachAssetResponse">ProjectAssetAttachAssetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListParams">ModelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartGenerationResponse">RunStartGenerationResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartTechniqueResponse">RunStartTechniqueResponse</a>

Methods:

- <code title="post /runs/generation">client.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunService.StartGeneration">StartGeneration</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartGenerationParams">RunStartGenerationParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartGenerationResponse">RunStartGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /runs/technique">client.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunService.StartTechnique">StartTechnique</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartTechniqueParams">RunStartTechniqueParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#RunStartTechniqueResponse">RunStartTechniqueResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Feedback

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#FeedbackRecordResponse">FeedbackRecordResponse</a>

Methods:

- <code title="post /feedback">client.Feedback.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#FeedbackService.Record">Record</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#FeedbackRecordParams">FeedbackRecordParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#FeedbackRecordResponse">FeedbackRecordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
