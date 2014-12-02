package service

import (
	"errors"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type getMTSServerSetupInfoByConfigFile struct {
}

func (getMTSServerSetupInfoByConfigFile *getMTSServerSetupInfoByConfigFile) GetInfo(cfgFileName string) (interface{}, error) {
	var result interface{} = nil
	if utils.IsEmptyStr(cfgFileName) {
		return nil, errors.New("没有配置文件")
	}
	var mTSServerSetupInfo = model.MTSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(&mTSServerSetupInfo, cfgFileName)
	if nil != err {
		return nil, err
	}
	result = &mTSServerSetupInfo
	if utils.IsEmptyStr(mTSServerSetupInfo.MTSInfo.Id) {
		return nil, errors.New("配置文件格式不正确")
	}
	return result, nil
}
