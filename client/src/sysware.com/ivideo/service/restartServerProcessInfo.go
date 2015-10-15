package service

import (
	"sysware.com/ivideo/utils"
)

type RestartServerProcessInfo interface {
	Restart(serverName string, oServerId string) error
}

var restartServerProcessInfo RestartServerProcessInfo

func init() {
	restartServerProcessInfo = nil
	if utils.IsLinuxOS() {
		restartServerProcessInfo = &restartServerProcessInfoLinux{}
	}
}

func NewRestartServerProcessInfo() RestartServerProcessInfo {
	return restartServerProcessInfo
}
