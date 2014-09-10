package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type getServerDetailSetupInfoImpl struct {
}

func (getServerDetailSetupInfoImpl *getServerDetailSetupInfoImpl) GetInfo(serverInfo *model.ServerInfo) (interface{}, error) {
	if nil == serverInfo {
		return nil, errors.New("没有服务器信息对象")
	}
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return nil, err
	}
	switch serverInfo.Type {
	case common.SERVER_TYPE_APS:
		return (&getAPSServerSetupInfoByConfigFile{}).GetInfo(cfgFileName)
	case common.SERVER_TYPE_SIP:
		return (&getSIPServerSetupInfoByConfigFile{}).GetInfo(cfgFileName)
	case common.SERVER_TYPE_CMS:
		return (&getCMSServerSetupInfoByConfigFile{}).GetInfo(cfgFileName)

	}
	return nil, errors.New("没有对应的对象")
}
