package service

import (
	"errors"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type getCMSServerSetupInfoByConfigFile struct {
}

func (getCMSServerSetupInfoByConfigFile *getCMSServerSetupInfoByConfigFile) GetInfo(cfgFileName string) (interface{}, error) {
	if utils.IsEmptyStr(cfgFileName) {
		return nil, errors.New("没有配置文件")
	}
	var cMSServerSetupInfo = model.CMSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(&cMSServerSetupInfo, cfgFileName)
	if nil != err {
		return nil, err
	}
	password, err := cMSServerSetupInfo.GetProperty("connection.password")
	if nil != err {
		return nil, errors.New("配置文件格式不正确")
	}
	result := make(map[string]string)
	result["dbPassword"] = password
	result["dbAddress"], _ = cMSServerSetupInfo.GetProperty("connection.url")
	result["dbUserName"], _ = cMSServerSetupInfo.GetProperty("connection.username")
	result["dbDefaultSchema"], _ = cMSServerSetupInfo.GetProperty("default_schema")
	return result, nil
}
