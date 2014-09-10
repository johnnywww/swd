package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type quitServerProcessInfoLinuxCommon struct {
}

type processTimerFuncHandleQuitProc struct {
}

func (processTimerFuncHandleQuitProc *processTimerFuncHandleQuitProc) run(cmdLine string) error {
	_, err := NewGetProcessInfo().GetInfo(cmdLine)
	if nil == err {
		return errors.New("找到了进程")
	} else {
		return nil
	}
}

func (quitServerProcessInfoLinuxCommon *quitServerProcessInfoLinuxCommon) Quit(serverInfo *model.ServerInfo) error {
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		return errors.New("不支持的类型")
	}
	_, err := NewGetProcessInfo().GetInfo(serverInfo.Address)
	if nil == err {
		serverId, err := NewGetServerIdInfo().getInfo(serverInfo.ServerName)
		if nil != err {
			return err
		}
		if utils.IsEmptyStr(serverId) {
			return errors.New(fmt.Sprintf("没有找到服务器程序%s对应的服务器编号", serverInfo.ServerName))
		}
		respQuitXmlInfo := respQuitXmlInfo{}
		err = getAbstarctUnixSocketXMLInfo(serverId, common.SERVER_COMMAND_QUIT, &respQuitXmlInfo)
		if nil != err {
			return err
		}
		if common.SERVER_COMMAND_QUIT != respQuitXmlInfo.CmdType {
			return errors.New("返回的命令不为退出命令")
		}
		if !processHandleTimerFunc(serverInfo.Address, 15, &processTimerFuncHandleQuitProc{}) {
			return errors.New(fmt.Sprintf("无法停止服务器%s程序", serverInfo.ServerName))
		}
	}
	return nil
}
