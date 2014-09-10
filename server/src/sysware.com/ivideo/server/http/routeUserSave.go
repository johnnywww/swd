package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type RouteUserSave struct {
}

func (routeUserSave *RouteUserSave) Routes(m *martini.ClassicMartini) {
	m.Post("/user/save", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /user/save")
		SetResponseJsonHeader(w)
		setupInfo := model.SetupInfo{}
		setupHandle := model.NewSetupHandle()
		if !setupHandle.GetSetupInfo(&setupInfo) {
			return model.GetErrorDtoJson("读取用户信息失败")
		}
		r.ParseForm()
		log.WriteLog("user save %v", r.Form)
		ousername := r.Form.Get("ousername")
		username := r.Form.Get("username")
		addFlag := utils.IsEmptyStr(ousername)
		password := r.Form.Get("password")
		if utils.IsEmptyStr(username) {
			return model.GetErrorDtoJson("没有用户名称")
		}
		if utils.IsEmptyStr(password) {
			return model.GetErrorDtoJson("没有密码")
		}
		index := setupInfo.GetUserIndexByUserName(username)
		oindex := setupInfo.GetUserIndexByUserName(ousername)
		var pUserInfo *model.UserInfo = nil
		if addFlag {
			if index > -1 {
				return model.GetErrorDtoJson("用户名称输入重复")
			}
			userInfo := model.UserInfo{UserName: username, Password: password}
			setupInfo.AddUserInfo(userInfo)
			pUserInfo = &userInfo
		} else {
			if oindex < 0 {
				return model.GetErrorDtoJson("没有找到用户")
			} else {
				if index > -1 && oindex != index {
					return model.GetErrorDtoJson("用户名称输入重复")
				}
			}
			setupInfo.Users[oindex].UserName = username
			setupInfo.Users[oindex].Password = password
			pUserInfo = &setupInfo.Users[oindex]
		}
		if !setupHandle.SaveSetupInfo(&setupInfo) {
			log.WriteLog("保存删除用户信息失败")
			return model.GetErrorDtoJson("保存删除用信息失败")
		}
		return model.GetDataDtoJson(model.UserSaveInfo{UserInfo: *pUserInfo, OUserName: ousername, AddFlag: addFlag})
	})

}
