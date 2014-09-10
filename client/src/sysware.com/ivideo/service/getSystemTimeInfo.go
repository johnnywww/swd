package service

import (
	"sysware.com/ivideo/utils"
)

type GetSystemTimeInfo interface {
	GetInfo() (string, error)
}

var getSystemTimeInfo GetSystemTimeInfo

func init() {
	getSystemTimeInfo = nil
	if utils.IsLinuxOS() {
		getSystemTimeInfo = &GetSystemTimeInfoLinux{}
	}
}

func NewGetSystemTimeInfo() GetSystemTimeInfo {
	return getSystemTimeInfo
}
