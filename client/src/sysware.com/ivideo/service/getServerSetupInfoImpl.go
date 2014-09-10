package service

import (
	"sysware.com/ivideo/model"
)

type getServerSetupInfoImpl struct {
}

func (getServerSetupInfoImpl *getServerSetupInfoImpl) GetInfo(serverName string) (*model.ServerInfo, interface{}, error) {
	var result interface{} = nil
	serverInfo, err := NewGetServerInfoByServerName().getInfo(serverName)
	if nil != err {
		return nil, nil, err
	}
	result, err = NewGetServerDetailSetupInfo().GetInfo(serverInfo)
	return serverInfo, result, err
}
