package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type restartProcessByServerInfoLinuxCommon struct {
}

type processTimerFuncHandleRestartProc struct {
}

func (processTimerFuncHandleRestartProc *processTimerFuncHandleRestartProc) run(param interface{}) error {
	var cmdLine string = param.(string)
	_, err := NewGetProcessInfo().GetInfo(cmdLine)
	return err
}

func (restartProcessByServerInfoLinuxCommon *restartProcessByServerInfoLinuxCommon) Restart(serverInfo *model.ServerInfo) error {
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		return errors.New("不支持的类型")
	}
	cmdLine := serverInfo.Address
	err := NewExecuteLinuxProc().Exec(cmdLine, utils.GetFileDir(cmdLine))
	if nil != err {
		return err
	}
	if !processHandleTimerFunc(cmdLine, 15, &processTimerFuncHandleRestartProc{}) {
		return errors.New(fmt.Sprintf("无法启动服务器%s程序", serverInfo.ServerName))
	}
	return nil
}
