package service

import (
	"testing"
	"time"
)

func Test_GetNetStatisticsInfo(t *testing.T) {
	getInfo := &GetNetStatisticsInfoLinux{}
	time.Sleep(5 * time.Second)
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
