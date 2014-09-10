package http

import (
	"github.com/go-martini/martini"
	"sysware.com/ivideo/model"
)

type RouteServer struct {
}

func (routeServer *RouteServer) Routes(m *martini.ClassicMartini) {
	routeHandles := &model.RouteHandles{Handles: []model.RouteHandle{&RouteServerList{}, &RouteServerSave{}, &RouteServerDelete{}}}
	routeHandles.Routes(m)
}
