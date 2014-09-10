package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
	"sysware.com/ivideo/utils"
)

type RouteServerStatusInfo struct {
}

func (routeServerStatusInfo *RouteServerStatusInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/status", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/status")
		shttp.SetResponseJsonHeader(w)
		r.ParseForm()
		serverName := r.FormValue("servername")
		if utils.IsEmptyStr(serverName) {
			return model.GetErrorDtoJson("没有服务器名称")
		}
		info, err := service.NewGetServerStatusInfo().GetInfo(serverName)
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		return model.GetDataDtoJson(info)
	})
}
