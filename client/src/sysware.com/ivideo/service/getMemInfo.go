package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetMemInfo interface {
	GetInfo() (*model.MemInfo, error)
}

var getMemInfo GetMemInfo

func init() {
	getMemInfo = nil
	if utils.IsLinuxOS() {
		getMemInfo = &GetMemInfoLinux{}
	}
}

func NewGetMemInfo() GetMemInfo {
	return getMemInfo
}
