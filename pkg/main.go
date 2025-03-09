package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend/tracing"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"go.opentelemetry.io/otel/attribute"
)

// QueryData 处理前端发来的查询请求
func (ds *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()

	for _, q := range req.Queries {
		var query struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal(q.JSON, &query); err != nil {
			return nil, err
		}

		// 模拟内置 API 的逻辑，根据变量值返回数据
		result := simulateAPI(query.Query)

		// 创建数据帧（Data Frame）
		frame := data.NewFrame("response",
			data.NewField("time", nil, []int64{1, 2, 3}),
			data.NewField("value", nil, result),
		)
		response.Responses[q.RefID] = backend.DataResponse{
			Frames: []*data.Frame{frame},
		}
	}

	return response, nil
}

// simulateAPI 模拟内置 API
func simulateAPI(query string) []float64 {
	// 根据用户输入的变量值返回数据
	if query == "test" {
		return []float64{10, 20, 30}
	}
	return []float64{1, 2, 3} // 默认数据
}

// 实现其他必要接口
func (ds *Datasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	log.DefaultLogger.Warn("Health check called")
	return &backend.CheckHealthResult{
		Status:  backend.HealthStatusOk,
		Message: "Data source is working",
	}, nil
}

func NewDatasource(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	log.DefaultLogger.Error("Creating new datasource instance", "url", settings.URL)
	// Uncomment the following to forward all HTTP headers in the requests made by the client
	// (disabled by default since SDK v0.161.0)
	// opts.ForwardHTTPHeaders = true

	// Using httpclient.New without any provided httpclient.Options creates a new HTTP client with a set of
	// default middlewares (httpclient.DefaultMiddlewares) providing additional built-in functionality, such as:
	//	- TracingMiddleware (creates spans for each outgoing HTTP request)
	//	- BasicAuthenticationMiddleware (populates Authorization header if basic authentication been configured via the
	//		DataSourceHttpSettings component from @grafana/ui)
	//	- CustomHeadersMiddleware (populates headers if Custom HTTP Headers been configured via the DataSourceHttpSettings
	//		component from @grafana/ui)
	//	- ContextualMiddleware (custom middlewares per context.Context, see httpclient.WithContextualMiddleware)

	return &Datasource{
		settings: settings,
	}, nil
}

var DatasourceOpts = datasource.ManageOpts{
	TracingOpts: tracing.Opts{
		// Optional custom attributes attached to the tracer's resource.
		// The tracer will already have some SDK and runtime ones pre-populated.
		CustomAttributes: []attribute.KeyValue{
			attribute.String("my_plugin.my_attribute", "custom value"),
		},
	},
}

// Datasource is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type Datasource struct {
	settings backend.DataSourceInstanceSettings
}

func main() {
	if err := datasource.Manage("smark-grafanacalculator-datasource", NewDatasource, DatasourceOpts); err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
