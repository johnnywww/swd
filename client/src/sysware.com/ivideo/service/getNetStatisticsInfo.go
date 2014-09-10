package service

import (
	"sysware.com/ivideo/utils"
)

type GetNetStatisticsInfo interface {
	GetInfo() ([][]string, error)
}

var getNetStatisticsInfo GetNetStatisticsInfo

func init() {
	getNetStatisticsInfo = nil
	if utils.IsLinuxOS() {
		getNetStatisticsInfo = &GetNetStatisticsInfoLinux{}
	}
}

func NewGetNetStatisticsInfo() GetNetStatisticsInfo {
	return getNetStatisticsInfo
}
