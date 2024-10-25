package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func initPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()
}

func main() {
	initPrometheus()

	r := InitWeb()

	r.Run(":9944")

}
