package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
)

type RouteServerDelete struct {
}

func (routeServerDelete *RouteServerDelete) Routes(m *martini.ClassicMartini) {
	m.Post("/server/delete", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/delete")
		SetResponseJsonHeader(w)
		setupInfo := model.SetupInfo{}
		setupHandle := model.NewSetupHandle()
		if !setupHandle.GetSetupInfo(&setupInfo) {
			return model.GetErrorDtoJson("读取服务器信息失败")
		}
		r.ParseForm()
		log.WriteLog("server delete %v", r.Form)
		serverName := r.Form.Get("servername")
		if !setupInfo.DeleteServerByServerName(serverName) {
			log.WriteLog("删除服务器信息失败")
			return model.GetErrorDtoJson("删除服务器信息失败")
		}
		if !setupHandle.SaveSetupInfo(&setupInfo) {
			log.WriteLog("保存删除服务器信息失败")
			return model.GetErrorDtoJson("保存删除服务器信息失败")
		}
		return model.GetDataDtoJson(nil)
	})
}
