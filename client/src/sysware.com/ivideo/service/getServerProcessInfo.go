package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetServerProcessInfo interface {
	GetInfo(serverName string) (*model.ServerProcessInfo, error)
}

var getServerProcessInfo GetServerProcessInfo

func init() {
	getServerProcessInfo = nil
	if utils.IsLinuxOS() {
		getServerProcessInfo = &getServerProcessInfoLinux{}
	}
}

func NewGetServerProcessInfo() GetServerProcessInfo {
	return getServerProcessInfo
}
