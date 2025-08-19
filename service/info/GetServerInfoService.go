package info

import (
	"github.com/shirou/gopsutil/v4/host"
	"runtime"
	"time"
)

type GetServerInfoService struct {
	serverStartTime time.Time
}

func NewGetServerInfoService() *GetServerInfoService {
	return &GetServerInfoService{
		serverStartTime: time.Now(),
	}
}

func (service *GetServerInfoService) GetServerInfo() ServerInfo {
	var serverInfo ServerInfo

	if hostInfo, err := host.Info(); err == nil {
		serverInfo = setServerInfo(hostInfo, service.serverStartTime)
	}
	return serverInfo
}

func setServerInfo(info *host.InfoStat, serverStartTime time.Time) ServerInfo {
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
