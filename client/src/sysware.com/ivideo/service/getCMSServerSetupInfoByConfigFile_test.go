package service

import (
	"testing"
)

func Test_GetCMSServerSetupInfoByConfigFile(t *testing.T) {
	getInfo := &getCMSServerSetupInfoByConfigFile{}
	_, err := getInfo.GetInfo("/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms/WEB-INF/classes/hibernate.cfg.xml")
	if nil != err {
		t.Error(err)
	}
}
