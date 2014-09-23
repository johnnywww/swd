package service

import (
	"errors"
	"net/http"
	"sysware.com/ivideo/utils"
)

type serverConfigSaveRequestHandleImpl struct {
}

func (serverConfigSaveRequestHandleImpl *serverConfigSaveRequestHandleImpl) Save(r *http.Request) error {
	if nil == r {
		return errors.New("没有请求")
	}

	serverName := r.FormValue("servername")
	if utils.IsEmptyStr(serverName) {
		return errors.New("没有服务器名称")
	}
	serverInfo, err := NewGetServerInfoByServerName().getInfo(serverName)
	if nil != err {
		return err
	}
	err = NewServerConfigSaveHandle().Save(r, serverInfo)
	if nil != err {
		return err
	}
	processInfo, err := NewGetProcessInfoByServerInfo().GetInfo(serverInfo)
	if nil != processInfo {
		err = NewRestartServerProcessInfo().Restart(serverName)
		if nil != err {
			return err
		}
	}
	return nil
}
