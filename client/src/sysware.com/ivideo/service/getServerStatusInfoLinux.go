package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type getServerStatusInfoLinux struct {
}

func (getServerStatusInfoLinux *getServerStatusInfoLinux) getCommonInfo(serverName string) (map[string][]*model.ServerStatusInfo, error) {
	result := map[string][]*model.ServerStatusInfo{}
	serverId, err := NewGetServerIdInfo().getInfo(serverName)
	if nil != err {
		return result, err
	}
	if utils.IsEmptyStr(serverId) {
		return result, errors.New(fmt.Sprintf("没有找到服务器程序%s对应的服务器编号", serverName))
	}
	return NewGetServerProcessStatusInfo().GetInfo(serverId)
}

func (getServerStatusInfoLinux *getServerStatusInfoLinux) GetInfo(serverName string) (map[string][]*model.ServerStatusInfo, error) {
	result := map[string][]*model.ServerStatusInfo{}
	serverInfo, err := NewGetServerInfoByServerName().getInfo(serverName)
	if nil != err {
		return result, err
	}
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		return NewGetServerProcessStatusInfo().GetInfo("")
	} else {
		return getServerStatusInfoLinux.getCommonInfo(serverName)
	}
}
