package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"log"
	"os"
	"time"
)

// CPUPercentV1 获取当前 进程 使用的CPU
func CPUPercentV1() (float64, error) {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return 0, err
	}
	cpuPercent, err := p.Percent(time.Second * 3) //取样3s内的cpu使用， 返回的是总的cpu使用率;mac上不用除以cpuCounts
	if err != nil {
		log.Fatal(err)
	}
	return cpuPercent, nil
}

func main() {
	//getCpu()
	fmt.Println("=====================================")
	//getMemInfo()
	fmt.Println("=====================================")
	//getUser()
	fmt.Println("=====================================")
	//getDisk()
	fmt.Println("=====================================")
	getPid()
}

func getCpu() {
	//打印cpu相关信息
	info, _ := cpu.Info()
	fmt.Println(info)
	for _, ci := range info {
		fmt.Println(ci)
	}

	//打印cpu使用率,每5秒一次，总共9次
	for i := 1; i < 10; i++ {
		time.Sleep(time.Millisecond * 5000)
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Println()
		fmt.Printf("%v, cpu percent: %v", i, percent)
	}
	fmt.Println()
	//显示cpu load值
	avg, _ := load.Avg()
	fmt.Println(avg)
}
func getMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("get memory info fail. err： ", err)
	}
	// 获取总内存大小，单位GB
	memTotal := memInfo.Total / 1024 / 1024 / 1024
	// 获取已用内存大小，单位MB
	memUsed := memInfo.Used / 1024 / 1024
	// 可用内存大小
	memAva := memInfo.Available / 1024 / 1024
	// 内存可用率
	memUsedPercent := memInfo.UsedPercent
	fmt.Printf("总内存: %v GB, 已用内存: %v MB, 可用内存: %v MB, 内存使用率: %.3f %% \n", memTotal, memUsed, memAva, memUsedPercent)
}

func getUser() {
	//显示机器启动时间戳
	bootTime, _ := host.BootTime()
	fmt.Println(bootTime)
	//显示机器信息
	info, _ := host.Info()
	fmt.Println(info)
	//显示终端用户
	users, _ := host.Users()
	for _, user := range users {
		fmt.Println(user.User)
	}
}

func getDisk() {
	//显示磁盘分区信息
	partitions, _ := disk.Partitions(true)
	for _, part := range partitions {
		fmt.Printf("part:%v\n", part.String())
		usage, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used :%v free:%v\n", usage.UsedPercent, usage.Free)
	}

	//显示磁盘分区IO信息
	counters, _ := disk.IOCounters()
	for k, v := range counters {
		fmt.Printf("%v,%v\n", k, v)
	}
}

func getProcess() {
	//显示所有进程名称和PID
	processes, _ := process.Processes()
	for _, process := range processes {
		fmt.Println(process.Pid)
		fmt.Println(process.Name())
	}
}

func getPid() {
	info, _ := process.Pids() //获取当前所有进程的pid
	fmt.Println(info)
}
