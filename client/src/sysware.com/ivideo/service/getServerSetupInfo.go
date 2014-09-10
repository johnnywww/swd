package service

import (
	"sysware.com/ivideo/model"
)

type GetServerSetupInfo interface {
	GetInfo(serverName string) (*model.ServerInfo, interface{}, error)
}

func NewGetServerSetupInfo() GetServerSetupInfo {
	return &getServerSetupInfoImpl{}
}
