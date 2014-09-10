package service

import (
	"sysware.com/ivideo/model"
)

type getServerProcessInfoLinux struct {
}

func (getServerProcessInfoLinux *getServerProcessInfoLinux) GetInfo(serverName string) (*model.ServerProcessInfo, error) {
	serverInfo, err := NewGetServerInfoByServerName().getInfo(serverName)
	if nil != err {
		return nil, err
	}
	processInfo, err := NewGetProcessInfoByServerInfo().GetInfo(serverInfo)
	if nil != err {
		return nil, err
	}
	return &model.ServerProcessInfo{*serverInfo, *processInfo}, nil
}
