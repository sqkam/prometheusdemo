package ginx

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Vector *prometheus.CounterVec

func InitCounter(opt prometheus.CounterOpts) {
	Vector = prometheus.NewCounterVec(opt,
		[]string{"code"})
	prometheus.MustRegister(Vector)
	// 考虑使用 code, method, 命中路由，HTTP 状态码
}
