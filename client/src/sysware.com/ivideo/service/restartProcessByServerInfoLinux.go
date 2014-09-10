package service

import (
	"errors"
	"os"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type restartProcessByServerInfoLinux struct {
}

func executeLinuxProc(cmdLine string, dir string) error {
	procattr := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: []*os.File{nil, nil, nil}}
	_, err := os.StartProcess(cmdLine, []string{}, &procattr)
	if nil != err {
		return err
	}
	return nil
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
