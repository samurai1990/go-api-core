package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func HandleHealth(c *gin.Context) (*server, error) {
	cpuNumber, _ := cpu.Counts(true)
	cpuInfo, _ := cpu.Percent(0, true)
	CpuPercent := make([]int, len(cpuInfo))
	for i, info := range cpuInfo {
		CpuPercent[i] = int(info)
	}
	memInfo, _ := mem.VirtualMemory()

	data := struct {
		Description   string `json:"description"`
		Version       string `json:"version"`
		CpuNumber     int    `json:"cpu_number"`
		CpuPercent    []int  `json:"cpu_percent"`
		MemoryPercent int    `json:"memory_percent"`
	}{
		Description:   "health of api service",
		Version:       "1.0.0",
		CpuNumber:     cpuNumber,
		CpuPercent:    CpuPercent,
		MemoryPercent: int(memInfo.UsedPercent),
	}
	serv := &server{
		ctx:  c,
		data: data,
	}
	return serv, nil
}
