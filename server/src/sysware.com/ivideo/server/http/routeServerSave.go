package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type CheckServerSave interface {
	CheckServerSave(oServerName string, pServerInfo *model.ServerInfo) error
}

type RouteServerSave struct {
	CheckServerSave
}

func (routeServerSave *RouteServerSave) Routes(m *martini.ClassicMartini) {
	m.Post("/server/save", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/save")
		SetResponseJsonHeader(w)
		setupInfo := model.SetupInfo{}
		setupHandle := model.NewSetupHandle()
		if !setupHandle.GetSetupInfo(&setupInfo) {
			return model.GetErrorDtoJson("读取服务器信息失败")
		}
		r.ParseForm()
		log.WriteLog("server save %v", r.Form)
		oservername := r.Form.Get("oservername")
		servername := r.Form.Get("servername")
		addFlag := utils.IsEmptyStr(oservername)
		address := r.Form.Get("address")
		iType, err := strconv.Atoi(r.FormValue("servertype"))
		if nil != err {
			log.WriteLog("error: %v", err)
			iType = common.SERVER_TYPE_CMS
		}
		if utils.IsEmptyStr(servername) {
			return model.GetErrorDtoJson("没有服务器名称")
		}
		if utils.IsEmptyStr(address) {
			return model.GetErrorDtoJson("没有服务器地址")
		}
		index := setupInfo.GetServerIndexByServerName(servername)
		oindex := setupInfo.GetServerIndexByServerName(oservername)
		var pServerInfo *model.ServerInfo = nil
		if addFlag {
			if index > -1 {
				return model.GetErrorDtoJson("服务器名称输入重复")
			}
			serverInfo := model.ServerInfo{ServerName: servername, Address: address, Type: iType}
			setupInfo.AddServerInfo(serverInfo)
			pServerInfo = &serverInfo
		} else {
			if oindex < 0 {
				return model.GetErrorDtoJson("没有找到服务器")
			} else {
				if index > -1 && oindex != index {
					return model.GetErrorDtoJson("服务器名称输入重复")
				}
			}
			setupInfo.Servers[oindex].ServerName = servername
			setupInfo.Servers[oindex].Address = address
			setupInfo.Servers[oindex].Type = iType
			pServerInfo = &setupInfo.Servers[oindex]
		}
		if nil != routeServerSave.CheckServerSave {
			err := routeServerSave.CheckServerSave.CheckServerSave(oservername, pServerInfo)
			if nil != err {
				return model.GetErrorObjDtoJson(err)
			}
		}
		if !setupHandle.SaveSetupInfo(&setupInfo) {
			log.WriteLog("保存删除服务器信息失败")
			return model.GetErrorDtoJson("保存删除服务器信息失败")
		}
		return model.GetDataDtoJson(model.ServerSaveInfo{ServerInfo: *pServerInfo, OServerName: oservername, AddFlag: addFlag})
	})
}
