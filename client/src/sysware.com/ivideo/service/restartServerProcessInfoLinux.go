package service

import (
	"sysware.com/ivideo/common"
	"time"
)

type restartServerProcessInfoLinux struct {
}

type processTimerFuncHandle interface {
	run(cmdLine string) error
}

func processHandleTimerFunc(cmdLine string, intervalSec int, processTimerFuncHandle processTimerFuncHandle) bool {
	quitChan := make(chan int)
	go func() {
		for {
			err := processTimerFuncHandle.run(cmdLine)
			if nil == err {
				quitChan <- common.RET_CODE_SUCCESS
			}
			time.Sleep(1 * time.Second)
			intervalSec = intervalSec - 1
			if 0 == intervalSec {
				quitChan <- common.RET_CODE_TIMEOUT
			}
		}
	}()
	result := <-quitChan
	return common.RET_CODE_SUCCESS == result
}

func (restartServerProcessInfoLinux *restartServerProcessInfoLinux) Restart(serverName string) error {
	serverInfo, err := NewGetServerInfoByServerName().getInfo(serverName)
	if nil != err {
		return err
	}
	err = NewQuitServerProcessInfo().Quit(serverInfo)
	if nil != err {
		return err
	}
	return NewRestartProcessByServerInfo().Restart(serverInfo)
}
