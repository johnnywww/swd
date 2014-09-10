package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
)

type RouteLogin struct {
}

func gotoLoginPage(w http.ResponseWriter, r *http.Request, re render.Render) {
	userInfo := GetSessionUserInfo(w, r)
	if nil == userInfo {
		log.WriteLog("not session user name")
		re.HTML(200, "login", nil)
	} else {
		retInfos := make(map[string]string)
		retInfos["userName"] = userInfo.UserName
		re.HTML(200, "login", retInfos)
	}
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
			sess := globalSessions.SessionStart(w, r)
			defer sess.SessionRelease(w)
			sess.Set("userInfo", userInfo)
			gotoRootPage(re, userInfo)
		}
	})
}
