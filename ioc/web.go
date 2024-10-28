package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"net/http"
	"prometheusdemo/handler"
	"prometheusdemo/pkg/ginx"
	"prometheusdemo/pkg/ginx/middlewares/metric"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *handler.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")

	})
	userHdl.RegisterRoutes(server)
	(&handler.ObservabilityHandler{}).RegisterRoutes(server)
	return server
}

func InitMiddlewares() []gin.HandlerFunc {

	ginx.InitCounter(prometheus.CounterOpts{
		Namespace: "test",
		Subsystem: "test",
		Name:      "http_biz_code",
		Help:      "HTTP 的业务错误码",
	})
	return []gin.HandlerFunc{

		(&metric.MiddlewareBuilder{
			Namespace:  "test",
			Subsystem:  "test",
			Name:       "gin_http",
			Help:       "统计 GIN 的 HTTP 接口",
			InstanceID: "my-instance-1",
		}).Build(),
		otelgin.Middleware("test"),
	}
}
