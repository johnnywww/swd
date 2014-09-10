package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

func gotoSetupPage(re render.Render) {
	setupInfo := model.SetupInfo{}
	model.NewSetupHandle().GetSetupInfo(&setupInfo)
	re.HTML(200, "setup", setupInfo)
}

type RouteServerSetup struct {
}

func (routeSetup *RouteServerSetup) Routes(m *martini.ClassicMartini) {

	m.Get("/setup", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		log.WriteLog("addr: /setup")
		userInfo := GetSessionUserInfo(w, r)
		if nil != userInfo && utils.IsAdminUser(userInfo.UserName) {
			gotoSetupPage(re)
		} else {
			retInfos := make(map[string]string)
			retInfos["userName"] = common.USER_NAME_ADMIN
			re.HTML(200, "setupLogin", retInfos)
		}
	})

	m.Get("/setupLogin", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		log.WriteLog("addr: /setupLogin")
		retInfos := make(map[string]string)
		retInfos["userName"] = common.USER_NAME_ADMIN
		re.HTML(200, "setupLogin", retInfos)
	})

	m.Post("/setupLogin", func(w http.ResponseWriter, r *http.Request, re render.Render) {
		r.ParseForm()
		username := r.Form.Get("username")
		if !utils.IsAdminUser(username) {
			retInfos := make(map[string]string)
			retInfos["userName"] = username
			retInfos["password"] = "******"
			retInfos["errorInfo"] = "不是管理员账号"
			re.HTML(200, "setupLogin", retInfos)
			return
		}
		userInfo := checkLoginUser(r.Form, re, "setupLogin")
		if nil != userInfo {
			gotoSetupPage(re)
		}
	})

}
