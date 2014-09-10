package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type RouteNotFound struct {
}

func (routeNotFound *RouteNotFound) Routes(m *martini.ClassicMartini) {
	m.NotFound(func(w http.ResponseWriter, r *http.Request, re render.Render) {
		re.HTML(200, "error", nil)
	})
}
