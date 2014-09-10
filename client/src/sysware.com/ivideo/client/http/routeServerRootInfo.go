package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/utils"
)

type RouteServerRootInfo struct {
}

func gotoServerInfo(w http.ResponseWriter, r *http.Request, re render.Render) {
	log.WriteLog("addr: /serverinfo")
	if !shttp.CheckSessionUserLogin(w, r) {
		gotoLoginPage(w, r, re)
		return
	}
	r.ParseForm()
	servername := r.FormValue("servername")
	if utils.IsEmptyStr(servername) {
		servername = getFirstServerName()
	}
	retInfos := make(map[string]string)
	retInfos["serverName"] = servername
	if utils.IsEmptyStr(servername) {
		retInfos["errorInfo"] = "没有服务器"
	}
	re.HTML(200, "serverInfo", retInfos)
}

func (routeServerRootInfo *RouteServerRootInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/serverinfo", gotoServerInfo)
}
