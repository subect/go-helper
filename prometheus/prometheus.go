package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	responseTimeSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http_request_duration",
	},
		[]string{"endpoint"},
	)

	errorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "myapp_errors_total",
		Help: "Total number of errors in your application.",
	})

	hourlyRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_hourly_requests_total",
		Help: "Total number of requests per hour in your application.",
	}, []string{"hour"})
)

func ping(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	path := req.URL.Path
	// 模拟业务处理
	ti := rand.Intn(5000)
	time.Sleep(time.Millisecond * time.Duration(ti))
	requestCount++
	//duration := time.Since(start)

	elapsed := (float64)(time.Since(start) / time.Millisecond)
	responseTimeSummary.WithLabelValues(path).Observe(elapsed)
	//tim := math.Round(duration.Seconds())
	//responseTimeSummary.Observe(tim)
	//fmt.Println("接口响应时间:", duration.Seconds())

	if ti < 100 {
		errorCounter.Inc()
	}
	//fmt.Println("接口错误数:", errorCounter)

	// 记录每小时的请求总数
	currentTime := time.Now()
	hour := currentTime.Format("2006-01-02 15")
	hourlyRequestCounter.WithLabelValues(hour).Inc()

	fmt.Fprintf(w, "pong")
}

//CpuCollector ----------------------------------------------
type CpuCollector struct {
	Cpu *prometheus.Desc
}

func NewCpuCollector() *CpuCollector {
	return &CpuCollector{
		Cpu: prometheus.NewDesc(
			"myapp_response_CPU",
			"Response time of my custom service",
			nil, nil,
		),
	}
}

func (c *CpuCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Cpu
}

func (c *CpuCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if v, err := CPUPercent(); err == nil {
		metricValue = v
	} else {
		fmt.Println("cpu err: ", err)
	}
	// 创建并发送指标
	ch <- prometheus.MustNewConstMetric(c.Cpu, prometheus.GaugeValue, metricValue)
}

//--------------------------------------------

//MemoryCollector --------------------------------------------
type MemoryCollector struct {
	Memory *prometheus.Desc
}

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{
		Memory: prometheus.NewDesc(
			"myapp_response_MEMORY",
			"Response time of my custom service",
			nil, nil,
		),
	}
}

func (m *MemoryCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.Memory
}

func (m *MemoryCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if memInfo, err := MEMORYPercent(); err == nil {
		metricValue = memInfo
	} else {
		fmt.Println("memory err: ", err)
	}
	totalMemory, err := TotalSystemMemory()
	if err != nil {
		fmt.Println("TotalSystemMemory err: ", err)
	}
	//fmt.Println("系统总内存: ", totalMemory)
	memoryUsagePercent := (metricValue / totalMemory) * 100.0
	//fmt.Printf("内存占用: %.2f%%\n", memoryUsagePercent)
	// 创建并发送指标
	ch <- prometheus.MustNewConstMetric(m.Memory, prometheus.GaugeValue, memoryUsagePercent)
}

func TotalSystemMemory() (float64, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	// 将内存从字节转换为兆字节
	totalMemoryMB := float64(memory.Total) / (1024 * 1024)

	return totalMemoryMB, nil
}

//--------------------------------------------

//NetCollector --------------------------------------------
type NetCollector struct {
	Net *prometheus.Desc
}

func NewNetCollector() *NetCollector {
	return &NetCollector{
		Net: prometheus.NewDesc(
			"myapp_response_Net",
			"Response time of my custom service",
			nil, nil,
		),
	}
}

func (m *NetCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.Net
}

func (m *NetCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	if memInfo, err := NetPercent(); err == nil {
		metricValue = memInfo
	} else {
		fmt.Println("net err: ", err)
	}
	// 创建并发送指标
	ch <- prometheus.MustNewConstMetric(m.Net, prometheus.GaugeValue, metricValue)
}

//--------------------------------------

//QpcCollector --------------------------------------
var requestCount uint64
var lastCollectionTime time.Time

func calculateQpc() float64 {
	now := time.Now()
	elapsedTime := now.Sub(lastCollectionTime).Seconds()
	qpc := float64(requestCount) / elapsedTime
	lastCollectionTime = now
	requestCount = 0
	return qpc
}

type QpcCollector struct {
	Qpc *prometheus.Desc
}

func NewQpcCollector() *QpcCollector {
	return &QpcCollector{
		Qpc: prometheus.NewDesc(
			"myapp_response_QPC",
			"Queries Per Second for my custom service",
			nil, nil,
		),
	}
}

func (q *QpcCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- q.Qpc
}

func (q *QpcCollector) Collect(ch chan<- prometheus.Metric) {
	// 在这里计算Qpc的值，然后创建并发送指标
	qpcValue := calculateQpc() // 请确保实现calculateQpc函数
	fmt.Println("Qps: ", qpcValue)
	ch <- prometheus.MustNewConstMetric(q.Qpc, prometheus.GaugeValue, qpcValue)
}

//=====================================
func CPUPercent() (float64, error) {
	//fmt.Println(os.Getpid())
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return 0, err
	}
	cpuPercent, err := p.Percent(time.Second * 3) //取样3s内的cpu使用， 返回的是总的cpu使用率;mac上不用除以cpuCounts
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cpu占用: %.2f%%\n", cpuPercent)
	return cpuPercent, nil
}

func MEMORYPercent() (float64, error) {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return 0, err
	}
	memInfo, err := p.MemoryInfo()
	if err != nil {
		log.Fatal(err)
	}
	mem := float64(memInfo.RSS) / 1024 / 1024
	fmt.Printf("内存占用: %.2f MB\n", mem)

	return mem, nil
}

func NetPercent() (float64, error) {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return 0, err
	}
	memInfo, err := p.Connections()
	if err != nil {
		log.Fatal(err)
	}
	mem := float64(len(memInfo))
	fmt.Printf("NET: %v \n", mem)

	return mem, nil
}

//=====================================

func main() {
	// 创建自定义指标收集器并注册
	prometheus.MustRegister(NewCpuCollector())    //CPU
	prometheus.MustRegister(NewMemoryCollector()) //Memory
	prometheus.MustRegister(NewNetCollector())    //Net
	prometheus.MustRegister(NewQpcCollector())    //Qps

	prometheus.MustRegister(responseTimeSummary)  //响应时间
	prometheus.MustRegister(errorCounter)         // 错误总数
	prometheus.MustRegister(hourlyRequestCounter) //每小时的吞吐量

	go handleCh()

	http.HandleFunc("/ping", ping)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func handleCh() {
	ch := make(chan int)
	for {
		select {
		case <-ch:
			// 模拟处理
			time.Sleep(1 * time.Second)
			fmt.Println("ch")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
