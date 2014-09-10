package service

import (
	"errors"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type getAPSServerSetupInfoByConfigFile struct {
}

func (getAPSServerSetupInfoByConfigFile *getAPSServerSetupInfoByConfigFile) GetInfo(cfgFileName string) (interface{}, error) {
	var result interface{} = nil
	if utils.IsEmptyStr(cfgFileName) {
		return nil, errors.New("没有配置文件")
	}
	var aPSServerSetupInfo = model.APSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(&aPSServerSetupInfo, cfgFileName)
	if nil != err {
		return nil, err
	}
	result = &aPSServerSetupInfo
	if utils.IsEmptyStr(aPSServerSetupInfo.APSInfo.Id) {
		return nil, errors.New("配置文件格式不正确")
	}
	return result, nil
}
