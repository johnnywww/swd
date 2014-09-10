package service

import (
	"sysware.com/ivideo/utils"
)

type GetSystemRunTimeInfo interface {
	GetInfo() (*RunTimeInfo, error)
}

type RunTimeInfo struct {
	Day  int64
	Hour int64
	Min  int64
}

var getSystemRunTimeInfo GetSystemRunTimeInfo

func init() {
	getSystemRunTimeInfo = nil
	if utils.IsLinuxOS() {
		getSystemRunTimeInfo = &GetSystemRunTimeInfoLinux{}
	}
}

func NewGetSystemRunTimeInfo() GetSystemRunTimeInfo {
	return getSystemRunTimeInfo
}
