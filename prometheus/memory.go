package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func TotalSystemMemoryV1() (uint64, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	return memory.Total, nil
}

func main() {
	val, _ := TotalSystemMemoryV1()
	fmt.Println(val)
}
