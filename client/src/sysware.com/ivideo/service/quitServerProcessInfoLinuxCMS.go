package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type quitServerProcessInfoLinuxCMS struct {
}

type processTimerFuncHandleQuitCMSProc struct {
}

func (processTimerFuncHandleQuitCMSProc *processTimerFuncHandleQuitCMSProc) run(param interface{}) error {
	var serverInfo *model.ServerInfo = param.(*model.ServerInfo)
	_, err := NewGetProcessInfoByServerInfo().GetInfo(serverInfo)
	if nil != err {
		return nil
	}
	return errors.New("没有中心管理服务器信息")
}

func (quitServerProcessInfoLinuxCMS *quitServerProcessInfoLinuxCMS) Quit(serverInfo *model.ServerInfo) error {
	if common.SERVER_TYPE_CMS != serverInfo.Type {
		return errors.New("不支持的类型")
	}
	respQuitXmlInfo := respQuitXmlInfo{}
	err := getCmsUnixSocketXMLInfo(common.SERVER_COMMAND_QUIT, &respQuitXmlInfo)
	if nil != err {
		return err
	}
	if common.SERVER_COMMAND_QUIT != respQuitXmlInfo.CmdType {
		return errors.New("返回的命令不为退出命令")
	}
	webPath, err := getCMSWebPath(serverInfo.Address)
	if nil != err {
		return err
	}
	cmdLine := webPath + "/bin/shutdown.sh"
	err = NewExecuteLinuxProc().Exec(cmdLine, webPath)
	if nil != err {
		return err
	}
	if !processHandleTimerFunc(serverInfo, 30, &processTimerFuncHandleQuitCMSProc{}) {
		return errors.New(fmt.Sprintf("无法停止中心管理服务器%s程序", serverInfo.ServerName))
	}
	return nil
}
