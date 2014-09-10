package service

import (
	"net/http"
	"sysware.com/ivideo/model"
)

type ServerConfigSaveHandle interface {
	Save(*http.Request, *model.ServerInfo) error
}

func NewServerConfigSaveHandle() ServerConfigSaveHandle {
	return &serverConfigSaveHandleImpl{}
}
