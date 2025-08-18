package info

import (
	"github.com/shirou/gopsutil/v4/host"
	"runtime"
	"time"
)

var serverStartTime = time.Now()

type GetServerInfoService struct {
}

func NewGetServerInfoService() *GetServerInfoService {
	return &GetServerInfoService{}
}

func (service *GetServerInfoService) GetServerInfo() ServerInfo {
	var serverInfo ServerInfo

	if hostInfo, err := host.Info(); err == nil {
		serverInfo = setServerInfo(hostInfo)
	}
	return serverInfo
}

func setServerInfo(info *host.InfoStat) ServerInfo {
	return ServerInfo{
		OS: OS{
			Name:            info.OS,
			PlatformName:    info.Platform,
			KernelVersion:   info.KernelVersion,
			PlatformVersion: info.PlatformVersion,
		},
		GoVersion: runtime.Version(),
		UpTime:    time.Since(serverStartTime).Round(time.Second).String(),
	}
}
