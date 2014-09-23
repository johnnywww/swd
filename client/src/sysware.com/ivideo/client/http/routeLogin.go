package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	shttp "sysware.com/ivideo/server/http"
)

type RouteLogin struct {
}

func gotoLoginPage(w http.ResponseWriter, r *http.Request, re render.Render) {
	userInfo := shttp.GetSessionUserInfo(w, r)
	retInfos := getStrMapRetInfo()
	if nil == userInfo {
		log.WriteLog("not session user name")
	} else {
		retInfos["userName"] = userInfo.UserName
	}
	re.HTML(200, "login", retInfos)
}

func (routeLogin *RouteLogin) Routes(m *martini.ClassicMartini) {

	m.Get("/login", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		gotoLoginPage(w, r, re)
	})

	m.Post("/login", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		r.ParseForm()
		userInfo := checkLoginUser(r.Form, re, "login")
		if nil != userInfo {
			userInfo.SetLoginState()
			shttp.SetSessionUserInfo(w, r, userInfo)
			gotoRootPage(re, userInfo)
		}
	})
}
