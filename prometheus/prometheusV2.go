package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 创建并注册一些 Prometheus 指标
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests to my app.",
		},
	)
	prometheus.MustRegister(counter)

	gauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "myapp_current_users",
			Help: "Current number of users in my app.",
		},
	)
	prometheus.MustRegister(gauge)

	histogram := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "myapp_request_duration_seconds",
			Help: "Request duration in seconds.",
			Buckets: []float64{0.1, 0.2, 0.5, 1, 2, 5},
		},
	)
	prometheus.MustRegister(histogram)

	// 自定义指标
	customMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "myapp_custom_metric",
			Help: "A custom metric for my app.",
		},
	)
	prometheus.MustRegister(customMetric)

	// 模拟请求处理
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		// 记录请求计数
		counter.Inc()

		// 模拟一些业务逻辑
		// ...

		// 设置当前用户数
		gauge.Set(42)

		// 模拟请求耗时
		histogram.Observe(0.5)

		// 设置自定义指标
		customMetric.Set(123.45)

		// 返回响应
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Prometheus!"))
	})

	// 设置 Prometheus 指标 HTTP 路由
	http.Handle("/metrics", promhttp.Handler())

	// 启动 HTTP 服务器
	http.ListenAndServe(":8080", nil)
}
