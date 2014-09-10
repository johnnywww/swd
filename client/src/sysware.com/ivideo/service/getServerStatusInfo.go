package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetServerStatusInfo interface {
	GetInfo(serverName string) (map[string][]*model.ServerStatusInfo, error)
}

var getServerStatusInfo GetServerStatusInfo

func init() {
	getServerStatusInfo = nil
	if utils.IsLinuxOS() {
		getServerStatusInfo = &getServerStatusInfoLinux{}
	}
}

func NewGetServerStatusInfo() GetServerStatusInfo {
	return getServerStatusInfo
}
