package metrics

import (
	"fmt"
	"golang-api-template/internal/application/config"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics interface {
	Listen() error
}

type metrics struct {
	config config.MetricsConfig
}

func NewMetrics(config config.MetricsConfig) Metrics {
	return &metrics{config: config}
}

func (m *metrics) Listen() error {

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	address := fmt.Sprintf("%s:%s", m.config.IpAddress, m.config.Port)

	return http.ListenAndServe(address, mux)
}
