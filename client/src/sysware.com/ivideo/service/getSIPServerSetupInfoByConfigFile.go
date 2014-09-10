package service

import (
	"errors"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type getSIPServerSetupInfoByConfigFile struct {
}

func (getSIPServerSetupInfoByConfigFile *getSIPServerSetupInfoByConfigFile) GetInfo(cfgFileName string) (interface{}, error) {
	var result interface{} = nil
	if utils.IsEmptyStr(cfgFileName) {
		return nil, errors.New("没有配置文件")
	}
	var sIPServerSetupInfo = model.SIPServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(&sIPServerSetupInfo, cfgFileName)
	if nil != err {
		return nil, err
	}
	result = &sIPServerSetupInfo
	if utils.IsEmptyStr(sIPServerSetupInfo.SIPInfo.Id) {
		return nil, errors.New("配置文件格式不正确")
	}
	return result, nil
}
