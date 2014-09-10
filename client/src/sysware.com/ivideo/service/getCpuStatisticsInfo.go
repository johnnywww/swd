package service

import (
	"sysware.com/ivideo/utils"
)

type GetCpuStatisticsInfo interface {
	GetInfo() (map[string]string, error)
}

var getCpuStatisticsInfo GetCpuStatisticsInfo

func init() {
	getCpuStatisticsInfo = nil
	if utils.IsLinuxOS() {
		getCpuStatisticsInfo = &GetCpuStatisticsInfoLinux{}
	}
}

func NewGetCpuStatisticsInfo() GetCpuStatisticsInfo {
	return getCpuStatisticsInfo
}
