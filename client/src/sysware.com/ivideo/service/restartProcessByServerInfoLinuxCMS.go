package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type restartProcessByServerInfoLinuxCMS struct {
}

type processTimerFuncHandleRestartCMSProc struct {
}

func (processTimerFuncHandleRestartCMSProc *processTimerFuncHandleRestartCMSProc) run(param interface{}) error {
	var cmdLine string = param.(string)
	getProcInfo := &getCMSProcessInfoLinux{}
	_, err := getProcInfo.GetInfo(cmdLine)
	if nil != err {
		return err
	}
	resp, err := utils.CmdOutputNoLn("ls", CMS_UNIX_DOMAIN_SOCKET)
	if nil != err {
		return err
	}
	if resp != CMS_UNIX_DOMAIN_SOCKET {
		return errors.New(fmt.Sprintf("没有找到文件%s", CMS_UNIX_DOMAIN_SOCKET))
	}
	return nil
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
	err = NewExecuteLinuxProc().Exec(cmdLine, webPath)
	if nil != err {
		return err
	}
	if !processHandleTimerFunc(serverInfo.Address, 180, &processTimerFuncHandleRestartCMSProc{}) {
		return errors.New(fmt.Sprintf("无法启动服务器%s程序", serverInfo.ServerName))
	}
	return nil
}
