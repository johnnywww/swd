package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
)

type getServerDetailSetupInfoCMSImpl struct {
}

func getCMSActiveMQConfigFileName(cmdLine string) string {
	return cmdLine + "/WEB-INF/classes/activemq.properties"
}

func (getServerDetailSetupInfoCMSImpl *getServerDetailSetupInfoCMSImpl) GetInfo(serverInfo *model.ServerInfo) (interface{}, error) {
	if nil == serverInfo {
		return nil, errors.New("没有服务器信息对象")
	}
	if common.SERVER_TYPE_CMS != serverInfo.Type {
		return nil, errors.New("不支持的类型")
	}
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return nil, err
	}

	result, err := (&getCMSServerSetupInfoByConfigFile{}).GetInfo(cfgFileName)
	if nil != err {
		return nil, err
	}
	r1, err := setup.NewPropertiesFileService().GetInfo(getCMSActiveMQConfigFileName(serverInfo.Address))
	if nil != err {
		return nil, err
	}
	(result.(map[string]string))["activemqAddress"] = r1["activemq.address"]
	return result, nil
}
