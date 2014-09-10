package model

import (
	"sysware.com/ivideo/service/setup"
	"testing"
)

func Test_CMSServerSetupInfoGetPassword(t *testing.T) {
	setupInfo := &CMSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms/WEB-INF/classes/hibernate.cfg.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	value, err := setupInfo.GetProperty("connection.password")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}

	if "sysware" != value {
		t.Error("错误密码")
	}
}

func Test_CMSServerSetupInfoSetPassword(t *testing.T) {
	setupInfo := &CMSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms/WEB-INF/classes/hibernate.cfg.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	err = setupInfo.SetProperty("connection.password", "sysware1")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	v1, err := setupInfo.GetProperty("connection.password")
	if nil != err {
		t.Error(err)
	}
	if v1 != "sysware1" {
		t.Error("设置密码失败")
	}
}

func Test_CMSServerSetupInfoAddNewPassword(t *testing.T) {
	setupInfo := &CMSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms/WEB-INF/classes/hibernate.cfg.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	err = setupInfo.SetProperty("connection.password1", "sysware1")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	v1, err := setupInfo.GetProperty("connection.password1")
	if nil != err {
		t.Error(err)
	}
	if v1 != "sysware1" {
		t.Error("设置密码失败")
	}
}
