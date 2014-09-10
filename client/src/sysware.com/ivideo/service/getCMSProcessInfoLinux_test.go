package service

import (
	"testing"
)

func Test_GetCMSProcessInfo(t *testing.T) {
	getInfo := &getCMSProcessInfoLinux{}
	_, err := getInfo.GetInfo("/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms")
	if nil != err {
		t.Error(err)
	}
}

func Test_GetNotExistCMSProcessInfo(t *testing.T) {
	getInfo := &getCMSProcessInfoLinux{}
	_, err := getInfo.GetInfo("/home/pc01/source/java/apache-tomcat-7.0.54/webapps/cms")
	if nil == err {
		t.Error("程序居然运行了")
	}
}
