package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
)

type RouteRoot struct {
}

func gotoRootPage(re render.Render, userInfo *model.UserInfo) {
	retInfos := getStrMapRetInfo()
	retInfos["userName"] = userInfo.UserName
	re.HTML(200, "index", retInfos)
}

func (routeRoot *RouteRoot) Routes(m *martini.ClassicMartini) {
	m.Get("/", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		log.WriteLog("addr: /")
		if !shttp.CheckSessionUserLogin(w, r) {
			gotoLoginPage(w, r, re)
		} else {
			gotoRootPage(re, shttp.GetSessionUserInfo(w, r))
		}
	})
}
