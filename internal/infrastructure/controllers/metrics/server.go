package metrics

import (
	"fmt"
	"golang-api-template/internal/application/config"
	"golang-api-template/pkg/advancedlog"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type Metrics interface {
	Listen() error
}

type metrics struct {
	config config.MetricsConfig

	log *logrus.Entry
}

func NewMetrics(config config.MetricsConfig, log *logrus.Entry) Metrics {
	return &metrics{config: config, log: log}
}

func (m *metrics) Listen() error {
	logF := advancedlog.FunctionLog(m.log)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	address := fmt.Sprintf("%s:%s", m.config.IpAddress, m.config.Port)

	logF.Infof("Start Listen metrics to address: %s", address)
	return http.ListenAndServe(address, mux)
}
