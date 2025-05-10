package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var templateMetrics = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace: "image",
	Subsystem: "kafka",
	Name:      "message",
}, []string{"template"})

func ObserveCreateImage(d time.Duration, template string) {
	templateMetrics.WithLabelValues(template).Observe(d.Seconds())
}

func Listen(host string) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(host, mux)
}
