package service

import (
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetAllProcessInfo interface {
	GetInfo() ([]model.ProcessInfo, error)
}

var getAllProcessInfo GetAllProcessInfo

func init() {
	getAllProcessInfo = nil
	if utils.IsLinuxOS() {
		getAllProcessInfo = &GetAllProcessInfoLinux{}
	}
}

func NewGetAllProcessInfo() GetAllProcessInfo {
	return getAllProcessInfo
}
