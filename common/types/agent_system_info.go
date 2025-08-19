package agent

type AgentSystemInfo struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	OS   OSInfo   `json:"os"`
	CPU  CPUInfo  `json:"cpu"`
	RAM  RAMInfo  `json:"ram"`
	Disk DiskInfo `json:"disk"`
}

type OSInfo struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	KernelVersion string `json:"kernel_version"`
}

type CPUInfo struct {
	Arch      string `json:"arch"`
	VCPUs     int    `json:"vcpus"`
	Model     string `json:"model"`
	MHzPerCPU int    `json:"mhz_per_cpu"`
}

type RAMInfo struct {
	TotalMB int `json:"total_mb"`
	UsedMB  int `json:"used_mb"`
	FreeMB  int `json:"free_mb"`
}

type DiskInfo struct {
	Mounts []MountInfo `json:"mounts"`
}

type MountInfo struct {
	Mount       string  `json:"mount"`
	TotalGB     float64 `json:"total_gb"`
	UsedGB      float64 `json:"used_gb"`
	UsedPercent float64 `json:"used_percent"`
}
