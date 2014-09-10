package service

import (
	"testing"
	"time"
)

func Test_GetCpuStatisticsInfo(t *testing.T) {
	getInfo := &GetCpuStatisticsInfoLinux{}
	time.Sleep(5 * time.Second)
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
