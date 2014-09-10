package http

import (
	"github.com/go-martini/martini"
	"sysware.com/ivideo/model"
)

type RouteUser struct {
}

func (routeUser *RouteUser) Routes(m *martini.ClassicMartini) {
	routeHandles := &model.RouteHandles{Handles: []model.RouteHandle{&RouteUserSave{}, &RouteUserDelete{}}}
	routeHandles.Routes(m)

}
