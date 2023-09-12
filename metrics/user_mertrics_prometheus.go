package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GetUserCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_user_count",
			Help: "Number of requests to the user",
		},
		[]string{"status"},
	)

	GetUserDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "request_user_duration",
			Help: "Duration of requests to the user",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(GetUserCount)
	prometheus.MustRegister(GetUserDuration)
}
