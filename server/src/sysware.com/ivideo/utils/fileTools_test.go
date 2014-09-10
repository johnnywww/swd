package utils

import (
	"sysware.com/ivideo/common"
	"testing"
)

func Test_GetServerConfigFileName(t *testing.T) {
	if "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/"+common.DEFAULT_CONFIG_FILE_NAME != GetServerConfigFileName("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps") {
		t.Error("配置文件名错误")
	}
}

func Test_IsServerConfigExist(t *testing.T) {
	if !IsServerConfigExist("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps") {
		t.Error("服务器配置文件不存在")
	}
}
