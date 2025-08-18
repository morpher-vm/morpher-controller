package info

type OS struct {
	Name            string
	PlatformName    string
	PlatformVersion string
	KernelVersion   string
}

type ServerInfo struct {
	OS        OS
	GoVersion string
	UpTime    string
}
