package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type RouteServerList struct {
}

func (routeServerList *RouteServerList) routeMoreServer(r *http.Request, re render.Render) {
	r.ParseForm()
	startPage, err := strconv.Atoi(r.FormValue("pageNo"))
	if nil != err {
		startPage = 0
	}
	serverPageInfo := model.NewServerPageHandle().GetServerPageInfo(startPage, 10)
	re.HTML(200, "serverList", serverPageInfo)
}

func (routeServerList *RouteServerList) Routes(m *martini.ClassicMartini) {
	m.Get("/serverList", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		log.WriteLog("addr: /serverList")

		serverName := getUniqueServerName()
		if utils.IsEmptyStr(serverName) {
			routeServerList.routeMoreServer(r, re)
		} else {
			gotoServerInfo(w, r, re)
		}

	})

}
