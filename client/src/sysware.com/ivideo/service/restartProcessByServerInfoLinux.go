package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type restartProcessByServerInfoLinux struct {
}

func (restartProcessByServerInfoLinux *restartProcessByServerInfoLinux) Restart(serverInfo *model.ServerInfo) error {
	if nil == serverInfo {
		return errors.New("没有服务器对象")
	}
	var restartProcessByServerInfo RestartProcessByServerInfo = nil
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		restartProcessByServerInfo = &restartProcessByServerInfoLinuxCMS{}
	} else {
		restartProcessByServerInfo = &restartProcessByServerInfoLinuxCommon{}
	}
	return restartProcessByServerInfo.Restart(serverInfo)
}
