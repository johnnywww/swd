package http

import (
	"github.com/go-martini/martini"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
)

type RouteServer struct {
}

func (routeServer *RouteServer) Routes(m *martini.ClassicMartini) {
	routeHandles := &model.RouteHandles{Handles: []model.RouteHandle{&RouteServerList{}, &RouteServerInfo{}, NewRouteServerSave(), &shttp.RouteServerDelete{}, &RouteServerProcessInfo{}, &routeServerConfig{}}}
	routeHandles.Routes(m)
}
