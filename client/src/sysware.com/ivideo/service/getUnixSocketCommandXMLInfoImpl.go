package service

import (
	"encoding/xml"
	"fmt"
	"net"
	"sysware.com/ivideo/common"
)

const (
	CMS_UNIX_DOMAIN_SOCKET = "/tmp/cmsserver.sock"
)

type getUnixSocketCommandXMLInfoImpl struct {
}

func getAbstarctUnixSocketXMLInfo(serverId string, cmdType string, value interface{}) error {
	return (&getUnixSocketCommandXMLInfoImpl{}).GetInfo(cmdType, value, fmt.Sprintf("@%s", serverId))
}

func getCmsUnixSocketXMLInfo(cmdType string, value interface{}) error {
	return (&getUnixSocketCommandXMLInfoImpl{}).GetInfo(cmdType, value, CMS_UNIX_DOMAIN_SOCKET)
}

func (getUnixSocketCommandXMLInfoImpl *getUnixSocketCommandXMLInfoImpl) GetInfo(cmdType string, value interface{}, sockeName string) error {
	conn, err := net.Dial("unix", sockeName)
	if nil != err {
		return err
	}
	defer conn.Close()
	_, err = conn.Write([]byte(NewGetServerRequestXMLInfo().GetXML(cmdType)))
	if nil != err {
		return err
	}
	buf := make([]byte, 3*common.KB_SIZE)
	readByteNum, err := conn.Read(buf[:])
	if nil != err {
		return err
	}
	responseStr := string(buf[0:readByteNum])
	err = xml.Unmarshal([]byte(responseStr), value)
	return err
}
