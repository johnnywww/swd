package service

import (
	"encoding/xml"
	"testing"
)

var respServerStatusXML string = `<?xml version="1.0"?>
<Response>
<CmdType>QueryStatus</CmdType>
<SN>11</SN>
<Items>
<Item>
<Type>Basic</Type>
<Content>2014-08-20T02:14:35 APS服务器启动</Content>
</Item>
<Item>
<Type>Register</Type>
<Content>APS服务器连接SipSvr失败</Content>
</Item>
<Item>
<Type>Link</Type>
<Content>APS服务器连接CMS失败</Content>
</Item>
<Item>
<Type>Alarm</Type>
<Content>APS服务器获取0条报警预案</Content>
<Content>APS服务器获取0条报警预案</Content>
</Item>
</Items>
</Response>
`

func Test_GetServerProcessStatusInfo(t *testing.T) {
	getInfo := &getServerProcessStatusInfoLinux{}
	_, err := getInfo.GetInfo("64020000002050000001")
	if nil != err {
		t.Error(err)
	}
}

func Test_GetCMSServerProcessStatusInfo(t *testing.T) {
	getInfo := &getServerProcessStatusInfoLinux{}
	_, err := getInfo.GetInfo("")
	if nil != err {
		t.Error(err)
	}
}

func Test_ParseServerStatusXML(t *testing.T) {
	respQueryStatusXmlInfo := respQueryStatusXmlInfo{}
	err := xml.Unmarshal([]byte(respServerStatusXML), &respQueryStatusXmlInfo)
	if nil != err {
		t.Error(err)
	}
}
