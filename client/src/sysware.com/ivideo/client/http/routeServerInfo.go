package http

import (
	"github.com/go-martini/martini"
	"sysware.com/ivideo/model"
)

type RouteServerInfo struct {
}

func (routeServerInfo *RouteServerInfo) Routes(m *martini.ClassicMartini) {
	routeHandles := &model.RouteHandles{Handles: []model.RouteHandle{&RouteServerRootInfo{}, &RouteServerGeneralInfo{}, &RouteServerCpuInfo{}, &RouteServerMemInfo{}, &RouteServerNetCardInfo{}, &RouteServerStatusInfo{}}}
	routeHandles.Routes(m)
}
