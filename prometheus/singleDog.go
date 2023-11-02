package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type CustomCollectorV1 struct {
	Cpu *prometheus.Desc
}

func NewCustomCollectorV1() *CustomCollectorV1 {
	return &CustomCollectorV1{
		Cpu: prometheus.NewDesc(
			"myapp_response_CPU",
			"Response time of my custom service",
			nil, nil,
		),
	}
}

func (c *CustomCollectorV1) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Cpu
}

func (c *CustomCollectorV1) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if v, err := CPUPercent(); err == nil {
		metricValue = v
	} else {
		fmt.Println("cpu err: ", err)
	}
	// 创建并发送指标
	ch <- prometheus.MustNewConstMetric(c.Cpu, prometheus.GaugeValue, metricValue)
}
