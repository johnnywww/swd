package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetProcessInfo interface {
	GetInfo(cmdLine string) (*model.ProcessInfo, error)
}

var getProcessInfo GetProcessInfo

func init() {
	getProcessInfo = nil
	if utils.IsLinuxOS() {
		getProcessInfo = &GetProcessInfoLinux{}
	}
}

func NewGetProcessInfo() GetProcessInfo {
	return getProcessInfo
}
