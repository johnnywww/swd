package model

import (
	"github.com/go-martini/martini"
)

type RouteHandle interface {
	Routes(m *martini.ClassicMartini)
}

type RouteHandles struct {
	Handles []RouteHandle
}

func (routeHandle *RouteHandles) Routes(m *martini.ClassicMartini) {
	for _, routeH := range routeHandle.Handles {
		routeH.Routes(m)
	}
}
