package service

import (
	"sysware.com/ivideo/utils"
)

type GetListenPortInfo interface {
	GetInfo() ([]int64, error)
}

var getListenPortInfo GetListenPortInfo

func init() {
	getListenPortInfo = nil
	if utils.IsLinuxOS() {
		getListenPortInfo = &GetListenPortInfoLinux{}
	}
}

func NewGetListenPortInfo() GetListenPortInfo {
	return getListenPortInfo
}
