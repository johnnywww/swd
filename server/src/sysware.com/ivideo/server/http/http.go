package http

import (
	"github.com/astaxie/beego/session"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go globalSessions.GC()
}

func GetSessionManager() *session.Manager {
	return globalSessions
}

func StartHttp(routeHandle model.RouteHandle) {
	port := 9500
	setupInfo := &model.SetupInfo{}
	if model.NewSetupHandle().GetSetupInfo(setupInfo) {
		if setupInfo.Port > 0 {
			port = setupInfo.Port
		}
	}
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"pageStartIndex": getPageStartIndex,
			"pageEndIndex":   getPageEndIndex,
			"pageIndexHtml":  getPageIndexHtml,
			"serverType":     getServerType,
		}},
	}))
	routeHandle.Routes(m)
	error := http.ListenAndServe(":"+strconv.Itoa(port), m)
	if nil != error {
		log.WriteLog("start http %d failed %s", port, error.Error())
	}

}

func SetResponseJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func GetSessionUserInfo(w http.ResponseWriter, r *http.Request) *model.UserInfo {
	sess := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	userInfo := sess.Get("userInfo")
	if nil == userInfo {
		return (*model.UserInfo)(nil)
	}
	return userInfo.(*model.UserInfo)
}

func CheckValidUser(form url.Values) (userInfo *model.UserInfo, errorMsg string) {
	userInfo = nil
	username := form.Get("username")
	password := form.Get("password")
	setupInfo := &model.SetupInfo{}
	if !model.NewSetupHandle().GetSetupInfo(setupInfo) {
		errorMsg = "没有配置信息，请找管理员"
		return
	}
	for _, userInfo1 := range setupInfo.Users {
		if userInfo1.Equal(username, password) {
			userInfo = &userInfo1
			break
		}
	}
	if nil == userInfo {
		errorMsg = "用户名或密码错误"
	}
	return userInfo, errorMsg
}

func CheckSessionUserLogin(w http.ResponseWriter, r *http.Request) bool {
	userInfo := GetSessionUserInfo(w, r)
	if nil == userInfo {
		return false
	}
	if !userInfo.IsLogin() {
		return false
	}
	return true
}
