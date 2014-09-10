package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type QuitServerProcessInfo interface {
	Quit(*model.ServerInfo) error
}

var quitServerProcessInfo QuitServerProcessInfo

func init() {
	quitServerProcessInfo = nil
	if utils.IsLinuxOS() {
		quitServerProcessInfo = &quitServerProcessInfoLinux{}
	}
}

func NewQuitServerProcessInfo() QuitServerProcessInfo {
	return quitServerProcessInfo
}
