package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type getServerIdInfoImpl struct {
}

func (getServerIdInfoImpl *getServerIdInfoImpl) getInfo(serverName string) (string, error) {
	result := ""
	serverInfo, vInfo, err := NewGetServerSetupInfo().GetInfo(serverName)
	if nil != err {
		return result, err
	}
	switch serverInfo.Type {
	case common.SERVER_TYPE_APS:
		{
			aPSServerSetupInfo := vInfo.(*model.APSServerSetupInfo)
			result = aPSServerSetupInfo.APSInfo.Id
			if utils.IsEmptyStr(result) {
				return result, errors.New(fmt.Sprintf("没有找到服务器程序%s对应的服务器编号", serverName))
			}
		}
	case common.SERVER_TYPE_SIP:
		{
			sIPServerSetupInfo := vInfo.(*model.SIPServerSetupInfo)
			result = sIPServerSetupInfo.SIPInfo.Id
			if utils.IsEmptyStr(result) {
				return result, errors.New(fmt.Sprintf("没有找到服务器程序%s对应的服务器编号", serverName))
			}
		}
	case common.SERVER_TYPE_MTS:
		{
			mTSServerSetupInfo := vInfo.(*model.MTSServerSetupInfo)
			result = mTSServerSetupInfo.MTSInfo.Id
			if utils.IsEmptyStr(result) {
				return result, errors.New(fmt.Sprintf("没有找到服务器程序%s对应的服务器编号", serverName))
			}
		}
	}
	return result, nil
}
