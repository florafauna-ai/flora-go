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
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunListResponse">TechniqueRunListResponse</a>

Methods:

- <code title="post /techniques/{techniqueId}/runs">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, techniqueID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunNewParams">TechniqueRunNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunNewResponse">TechniqueRunNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /techniques/{techniqueId}/runs/{runId}">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunGetParams">TechniqueRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunGetResponse">TechniqueRunGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /technique-runs">client.Techniques.Runs.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunListParams">TechniqueRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#TechniqueRunsCursorPage">TechniqueRunsCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#TechniqueRunListResponse">TechniqueRunListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

## Canvas

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasGetResponse">ProjectCanvasGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasUpdateResponse">ProjectCanvasUpdateResponse</a>

Methods:

- <code title="get /projects/{projectId}/canvas">client.Projects.Canvas.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasGetResponse">ProjectCanvasGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /projects/{projectId}/canvas">client.Projects.Canvas.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasUpdateParams">ProjectCanvasUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectCanvasUpdateResponse">ProjectCanvasUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionNewResponse">ProjectActionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionRunResponse">ProjectActionRunResponse</a>

Methods:

- <code title="post /projects/{projectId}/actions">client.Projects.Actions.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionNewParams">ProjectActionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionNewResponse">ProjectActionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /projects/{projectId}/actions/{nodeId}/run">client.Projects.Actions.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, nodeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionRunParams">ProjectActionRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ProjectActionRunResponse">ProjectActionRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionGetResponse">ActionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionListResponse">ActionListResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionRunResponse">ActionRunResponse</a>

Methods:

- <code title="get /actions/{actionId}">client.Actions.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, actionID <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionGetParamsActionID">ActionGetParamsActionID</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionGetResponse">ActionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /actions">client.Actions.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionListResponse">ActionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /runs/action">client.Actions.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionRunParams">ActionRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ActionRunResponse">ActionRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListParams">ModelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Generations

Response Types:

- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationNewResponse">GenerationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationGetResponse">GenerationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationListResponse">GenerationListResponse</a>

Methods:

- <code title="post /generate">client.Generations.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationNewParams">GenerationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationNewResponse">GenerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /runs/{runId}">client.Generations.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, runID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationGetResponse">GenerationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /generations">client.Generations.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationListParams">GenerationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go/packages/pagination#GenerationsCursorPage">GenerationsCursorPage</a>[<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go">flora</a>.<a href="https://pkg.go.dev/github.com/florafauna-ai/flora-go#GenerationListResponse">GenerationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
