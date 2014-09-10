package http

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
	"sysware.com/ivideo/utils"
)

type RouteServerProcessInfo struct {
}

func (routeServerProcessInfo *RouteServerProcessInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/process", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/process")
		shttp.SetResponseJsonHeader(w)
		r.ParseForm()
		serverName := r.FormValue("servername")
		if utils.IsEmptyStr(serverName) {
			return model.GetErrorDtoJson("没有服务器名称")
		}
		serverProcessInfo, err := service.NewGetServerProcessInfo().GetInfo(serverName)
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		result := map[string]string{}
		result["cpuRate"] = fmt.Sprintf("%.1f%%", serverProcessInfo.CPURate)
		result["memRate"] = fmt.Sprintf("%.1f%%", serverProcessInfo.MemRate)
		result["realMem"] = fmt.Sprintf("%s", utils.DisplaySizeStr(float64(serverProcessInfo.MemRealSize)))
		result["virtualMem"] = fmt.Sprintf("%s", utils.DisplaySizeStr(float64(serverProcessInfo.MemVmSize)))
		result["state"] = serverProcessInfo.State
		return model.GetDataDtoJson(result)
	})
}
