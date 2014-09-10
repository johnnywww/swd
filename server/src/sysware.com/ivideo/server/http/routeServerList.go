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

func (routeServerList *RouteServerList) Routes(m *martini.ClassicMartini) {

	m.Get("/serverList", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		log.WriteLog("addr: /serverList")
		r.ParseForm()
		startPage, err := strconv.Atoi(r.FormValue("pageNo"))
		if nil != err {
			startPage = 0
		}
		serverPageInfo := model.NewServerPageHandle().GetServerPageInfo(startPage, 10)
		for index, _ := range serverPageInfo.Servers {
			serverPageInfo.Servers[index].Status = utils.GetServerState(serverPageInfo.Servers[index].Address)
		}
		re.HTML(200, "serverList", serverPageInfo)
	})

}
