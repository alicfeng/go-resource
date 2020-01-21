package bean

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/mem"
)

type resource struct {
	System system `json:"system"`
	GPU    gpu    `json:"gpu"`
	CPU    cpu    `json:"cpu"`
	Memory memory `json:"mem"`
	Bios   bios   `json:"bios"`
}

var (
	Resource resource
)

func init() {
	Resource.System = initSystem()
	Resource.CPU = initCPU()
	Resource.Memory = initMemory()
	Resource.Bios = initBios()
}

/**
初始化系统信息
*/
func initSystem() (System system) {
	System.Arch = runtime.GOARCH
	System.OS = runtime.GOOS
	hostname, err := os.Hostname()
	if err == nil {
		System.Hostname = hostname
	}

	return System
}

/**
初始化CPU信息
*/
func initCPU() (CPU cpu) {
	CPU.Number = runtime.NumCPU()
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return CPU
	}

	lines := strings.Split(string(contents), "\n")[0:CPU.Number]
	for _, line := range lines {
		var CPUInfo cpuInfo
		fields := strings.Fields(line)
		numFields := len(fields)
		for index := 1; index < numFields; index++ {
			value, err := strconv.ParseUint(fields[index], 10, 64)
			if err != nil {
				fmt.Println("Error: ", index, fields[index], err)
			}
			CPUInfo.Total += value // tally up all the numbers to get total ticks
			if index == 4 { // idle is the 5th field in the cpu line
				CPUInfo.Free = value
			}
		}

		CPUInfo.Name = fields[0]
		CPUInfo.Busy = CPUInfo.Total - CPUInfo.Free
		CPUInfo.Usage = (float64(CPUInfo.Busy)) / float64(CPUInfo.Total)

		CPU.CPUInfo = append(CPU.CPUInfo, CPUInfo)
	}
	return CPU
}

/**
初始化内存信息
*/
func initMemory() (Memory memory) {
	memInfo, _ := mem.VirtualMemory()
	Memory.Total = memInfo.Total
	Memory.Used = memInfo.Used
	Memory.Free = memInfo.Available
	return Memory
}

func initBios() (Bios bios) {
	biosInfo, err := ghw.BIOS()
	if err == nil {
		Bios.Vendor = biosInfo.Vendor
		Bios.Version = biosInfo.Version
		Bios.Date = biosInfo.Date
	}
	return Bios
}
