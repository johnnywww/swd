package service

import (
	"sysware.com/ivideo/model"
)

type GetServerDetailSetupInfo interface {
	GetInfo(*model.ServerInfo) (interface{}, error)
}

func NewGetServerDetailSetupInfo() GetServerDetailSetupInfo {
	return &getServerDetailSetupInfoImpl{}
}
