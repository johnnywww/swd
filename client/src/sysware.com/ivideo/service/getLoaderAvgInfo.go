package service

import (
	"sysware.com/ivideo/utils"
)

type LoaderAvgInfo struct {
	Avg1min  float64
	Avg5min  float64
	Avg15min float64
}

type GetLoaderAvgInfo interface {
	GetInfo() (*LoaderAvgInfo, error)
}

var getLoaderAvgInfo GetLoaderAvgInfo

func init() {
	getLoaderAvgInfo = nil
	if utils.IsLinuxOS() {
		getLoaderAvgInfo = &GetLoaderAvgInfoLinux{}
	}
}

func NewGetLoaderAvgInfo() GetLoaderAvgInfo {
	return getLoaderAvgInfo
}
