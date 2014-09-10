package service

import (
	"testing"
)

func Test_GetSystemRunTimeInfo(t *testing.T) {
	getInfo := &GetSystemRunTimeInfoLinux{}
	_, err := getInfo.GetInfo()
	if err != nil {
		t.Error(err)
	}
}
