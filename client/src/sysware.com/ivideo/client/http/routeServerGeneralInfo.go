package http

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"os"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
	"sysware.com/ivideo/utils"
)

type RouteServerGeneralInfo struct {
}

func (routeServerGeneralInfo *RouteServerGeneralInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/general", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/general")
		shttp.SetResponseJsonHeader(w)
		result := map[string]string{}
		val, err := service.NewGetSystemTimeInfo().GetInfo()
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		result["time"] = val
		runTimeInfo, err1 := service.NewGetSystemRunTimeInfo().GetInfo()
		if nil != err1 {
			return model.GetErrorObjDtoJson(err1)
		}
		result["runtime"] = fmt.Sprintf("%d天 %d小时 %d分钟", runTimeInfo.Day, runTimeInfo.Hour, runTimeInfo.Min)
		hostname, err3 := os.Hostname()
		if nil != err3 {
			return model.GetErrorObjDtoJson(err3)
		}
		result["hostname"] = hostname
		ver, err2 := utils.CmdOutputNoLn("uname", "-r")
		if nil != err2 {
			return model.GetErrorObjDtoJson(err2)
		}
		result["kernel"] = ver
		return model.GetDataDtoJson(result)
	})
}
