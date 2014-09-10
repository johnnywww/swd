package service

import (
	"strconv"
	"strings"
	"sysware.com/ivideo/utils"
)

type GetLoaderAvgInfoLinux struct {
}

func (getLoaderAvgInfoLinux *GetLoaderAvgInfoLinux) GetInfo() (*LoaderAvgInfo, error) {
	loaderAvgInfo := LoaderAvgInfo{}

	data, err := utils.ReadFileToStringNoLn("/proc/loadavg")
	if nil != err {
		return nil, err
	}

	slice := strings.Split(string(data), " ")
	if loaderAvgInfo.Avg1min, err = strconv.ParseFloat(slice[0], 64); nil != err {
		return nil, err
	}
	if loaderAvgInfo.Avg5min, err = strconv.ParseFloat(slice[1], 64); nil != err {
		return nil, err
	}
	if loaderAvgInfo.Avg15min, err = strconv.ParseFloat(slice[2], 64); nil != err {
		return nil, err
	}

	return &loaderAvgInfo, nil
}
