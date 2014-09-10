package service

import (
	"encoding/xml"
)

type respQueryStatusXmlInfo struct {
	XMLName xml.Name                  `xml:"Response"`
	CmdType string                    `xml:"CmdType"`
	Sn      int                       `xml:"SN"`
	Items   []respQueryStatusItemInfo `xml:"Items>Item"`
}

type respQueryStatusItemInfo struct {
	Type     string   `xml:"Type"`
	Contents []string `xml:"Content"`
}

type GetServerStatusXMLInfo interface {
	getInfo(serverId string) (*respQueryStatusXmlInfo, error)
}

func NewGetServerStatusXMLInfo() GetServerStatusXMLInfo {
	return &getServerStatusXMLInfoImpl{}
}
