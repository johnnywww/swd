package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
)

type RouteUserDelete struct {
}

func (routeUserDelete *RouteUserDelete) Routes(m *martini.ClassicMartini) {
	m.Post("/user/delete", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /user/delete")
		SetResponseJsonHeader(w)
		setupInfo := model.SetupInfo{}
		setupHandle := model.NewSetupHandle()
		if !setupHandle.GetSetupInfo(&setupInfo) {
			return model.GetErrorDtoJson("读取用户信息失败")
		}
		r.ParseForm()
		log.WriteLog("user delete %v", r.Form)
		userName := r.Form.Get("username")

		if common.USER_NAME_ADMIN == userName {
			log.WriteLog("删除admin失败")
			return model.GetErrorDtoJson("删除管理员信息失败")
		}
		if !setupInfo.DeleteUserByUserName(userName) {
			log.WriteLog("删除用户信息失败")
			return model.GetErrorDtoJson("删除用户信息失败")
		}

		if !setupHandle.SaveSetupInfo(&setupInfo) {
			log.WriteLog("保存删除用户信息失败")
			return model.GetErrorDtoJson("保存删除用户信息失败")
		}
		return model.GetDataDtoJson(nil)
	})
}
