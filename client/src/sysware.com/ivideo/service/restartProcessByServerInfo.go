package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type RestartProcessByServerInfo interface {
	Restart(*model.ServerInfo) error
}

var restartProcessByServerInfo RestartProcessByServerInfo

func init() {
	restartProcessByServerInfo = nil
	if utils.IsLinuxOS() {
		restartProcessByServerInfo = &restartProcessByServerInfoLinux{}
	}
}

func NewRestartProcessByServerInfo() RestartProcessByServerInfo {
	return restartProcessByServerInfo
}
