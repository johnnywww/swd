package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetProcessInfoByServerInfo interface {
	GetInfo(serverInfo *model.ServerInfo) (*model.ProcessInfo, error)
}

var getProcessInfoByServerInfo GetProcessInfoByServerInfo

func init() {
	getProcessInfoByServerInfo = nil
	if utils.IsLinuxOS() {
		getProcessInfoByServerInfo = &getProcessInfoByServerInfoLinux{}
	}
}

func NewGetProcessInfoByServerInfo() GetProcessInfoByServerInfo {
	return getProcessInfoByServerInfo
}
