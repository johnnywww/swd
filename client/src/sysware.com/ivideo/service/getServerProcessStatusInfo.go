package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetServerProcessStatusInfo interface {
	GetInfo(serverId string) (map[string][]*model.ServerStatusInfo, error)
}

var getServerProcessStatusInfo GetServerProcessStatusInfo

func init() {
	getServerProcessStatusInfo = nil
	if utils.IsLinuxOS() {
		getServerProcessStatusInfo = &getServerProcessStatusInfoLinux{}
	}
}

func NewGetServerProcessStatusInfo() GetServerProcessStatusInfo {
	return getServerProcessStatusInfo
}
