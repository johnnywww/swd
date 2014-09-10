package service

import (
	"errors"
	"fmt"
	"strings"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type getCMSProcessInfoLinux struct {
}

func getCMSWebPath(cmdLine string) (string, error) {
	result := utils.GetFileDir(utils.GetFileDir(cmdLine))
	if utils.IsEmptyStr(result) {
		return result, errors.New("不是正确的中心管理服务器路径")
	}
	return result, nil
}

func (getCMSProcessInfoLinux *getCMSProcessInfoLinux) GetInfo(cmdLine string) (*model.ProcessInfo, error) {
	webPath, err := getCMSWebPath(cmdLine)
	if nil != err {
		return nil, err
	}
	processInfos, err := NewGetAllProcessInfo().GetInfo()
	if nil != err {
		return nil, err
	}
	for _, processInfo := range processInfos {
		if strings.Contains(processInfo.Cmdline, webPath) {
			return &processInfo, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("not found process %s", webPath))
}
