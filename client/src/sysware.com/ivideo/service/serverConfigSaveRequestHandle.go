package service

import (
	"net/http"
)

type ServerConfigSaveRequestHandle interface {
	Save(*http.Request) error
}

func NewServerConfigSaveRequestHandle() ServerConfigSaveRequestHandle {
	return &serverConfigSaveRequestHandleImpl{}
}
