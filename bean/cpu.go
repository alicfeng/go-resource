package bean

type cpu struct {
	Number  int       `json:"number"`
	CPUInfo []cpuInfo `json:"cpu_info"`
}

type cpuInfo struct {
	Name  string  `json:"name"`
	Total uint64  `json:"total"`
	Busy  uint64  `json:"busy"`
	Free  uint64  `json:"free"`
	Usage float64 `json:"usage"`
}
