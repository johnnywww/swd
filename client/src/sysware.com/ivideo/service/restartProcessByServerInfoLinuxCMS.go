package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type restartProcessByServerInfoLinuxCMS struct {
}

type processTimerFuncHandleRestartCMSProc struct {
}

func (processTimerFuncHandleRestartCMSProc *processTimerFuncHandleRestartCMSProc) run(cmdLine string) error {
	getProcInfo := &getCMSProcessInfoLinux{}
	_, err := getProcInfo.GetInfo(cmdLine)
	return err
}

func (restartProcessByServerInfoLinuxCMS *restartProcessByServerInfoLinuxCMS) Restart(serverInfo *model.ServerInfo) error {
	if common.SERVER_TYPE_CMS != serverInfo.Type {
		return errors.New("不支持的类型")
	}
	webPath, err := getCMSWebPath(serverInfo.Address)
	if nil != err {
		return err
	}
	cmdLine := webPath + "/bin/startup.sh"
	err = restartLinuxProc(cmdLine, webPath)
	if nil != err {
		return err
	}
	if !processHandleTimerFunc(serverInfo.Address, 15, &processTimerFuncHandleRestartCMSProc{}) {
		return errors.New(fmt.Sprintf("无法启动服务器%s程序", serverInfo.ServerName))
	}
	return nil
}
