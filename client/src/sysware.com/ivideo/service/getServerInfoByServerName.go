package service

import (
	"sysware.com/ivideo/model"
)

type GetServerInfoByServerName interface {
	getInfo(serverName string) (*model.ServerInfo, error)
}

func NewGetServerInfoByServerName() GetServerInfoByServerName {
	return &getServerInfoByServerNameImpl{}
}
