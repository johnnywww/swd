package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetNetInfo interface {
	GetInfo() ([]model.NetCardInfo, error)
}

var getNetInfo GetNetInfo

func init() {
	getNetInfo = nil
	if utils.IsLinuxOS() {
		getNetInfo = &GetNetInfoLinux{}
	}
}

func NewGetNetInfo() GetNetInfo {
	return getNetInfo
}
