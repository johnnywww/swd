package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	shttp "sysware.com/ivideo/server/http"
)

type RouteLogout struct {
}

func logOut(w http.ResponseWriter, r *http.Request, re render.Render) {
	log.WriteLog("addr: /logout")
	userInfo := shttp.GetSessionUserInfo(w, r)
	retInfos := getStrMapRetInfo()
	if nil != userInfo {
		userInfo.Logout()
		retInfos["userName"] = userInfo.UserName
	}
	re.HTML(200, "login", retInfos)
}

func (routeLogout *RouteLogout) Routes(m *martini.ClassicMartini) {
	m.Get("/logout", logOut)
	m.Post("/logout", logOut)
}
