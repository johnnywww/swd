package service

import (
	"fmt"
	"sysware.com/ivideo/common"
)

var getSeqNo GetSeqNo

func init() {
	getSeqNo = NewGetSeqNo()
}

type getServerRequestXMLInfoImpl struct {
}

func (getServerRequestXMLInfoImpl *getServerRequestXMLInfoImpl) GetXML(cmdType string) string {
	return fmt.Sprintf(common.SERVER_REQUEST_XML, getSeqNo.GetNo(), cmdType)
}
