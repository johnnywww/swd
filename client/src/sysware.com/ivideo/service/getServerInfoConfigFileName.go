package service

import (
	"sysware.com/ivideo/model"
)

type GetServerInfoConfigFileName interface {
	GetInfo(*model.ServerInfo) (string, error)
}

func NewGetServerInfoConfigFileName() GetServerInfoConfigFileName {
	return &getServerInfoConfigFileNameImpl{}
}
