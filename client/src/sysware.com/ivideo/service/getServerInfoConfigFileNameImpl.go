package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type getServerInfoConfigFileNameImpl struct {
}

func getCMSConfigFileName(cmdLine string) string {
	return cmdLine + "/WEB-INF/classes/hibernate.cfg.xml"
}

func (getServerInfoConfigFileNameImpl *getServerInfoConfigFileNameImpl) GetInfo(serverInfo *model.ServerInfo) (string, error) {
	if nil == serverInfo {
		return "", errors.New("没有服务器对象")
	}
	var configFileName string = ""
	switch serverInfo.Type {
	case common.SERVER_TYPE_APS, common.SERVER_TYPE_MTS, common.SERVER_TYPE_SIP:
		configFileName = utils.GetServerConfigFileName(serverInfo.Address)
	case common.SERVER_TYPE_CMS:
		configFileName = getCMSConfigFileName(serverInfo.Address)
	}
	if utils.IsEmptyStr(configFileName) {
		return "", errors.New("不支持的服务器类型")
	}
	if !utils.IsFileExist(configFileName) {
		return "", errors.New("服务器配置文件不存在")
	}
	return configFileName, nil
}
