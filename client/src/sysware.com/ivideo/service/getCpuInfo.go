package service

import (
	"errors"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetCpuInfo interface {
	GetInfo() (*model.CpuInfo, error)
}

var getCpuInfo GetCpuInfo

func init() {
	getCpuInfo = nil
	if utils.IsLinuxOS() {
		getCpuInfo = &GetCpuInfoLinux{}
	}
}

func NewGetCpuInfo() GetCpuInfo {
	return getCpuInfo
}

func getLinuxOSError() error {
	if !utils.IsLinuxOS() {
		return errors.New("not linux os")
	}
	return nil
}
