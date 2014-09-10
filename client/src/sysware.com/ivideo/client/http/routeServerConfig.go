package http

import (
	"github.com/go-martini/martini"
	"sysware.com/ivideo/model"
)

type routeServerConfig struct {
}

func (routeServerConfig *routeServerConfig) Routes(m *martini.ClassicMartini) {
	routeHandles := &model.RouteHandles{Handles: []model.RouteHandle{&routeServerConfigRoot{}, &routeServerConfigSave{}}}
	routeHandles.Routes(m)
}
