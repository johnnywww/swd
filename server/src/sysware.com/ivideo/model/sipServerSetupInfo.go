package model

import (
	"encoding/xml"
)

type SIPServerSetupInfo struct {
	XMLName       xml.Name `xml:"sysware"`
	SIPInfo       SIPServerSIPInfo
	CMSServerInfo CMSServerInfo
	DbInfo        SIPServerDbInfo
	OptionInfo    SIPServerOptionInfo
}

type SIPServerSIPInfo struct {
	XMLName              xml.Name `xml:"sipserver"`
	Id                   string   `xml:"id"`
	Password             string   `xml:"password"`
	Domain               string   `xml:"domain"`
	DefaultHeartInterval int      `xml:"defaultheartinterval"`
	AddressInfo          AddressInfo
}

type SIPServerDbInfo struct {
	XMLName     xml.Name `xml:"db"`
	AddressInfo AddressInfo
	Password    string `xml:"password"`
}

type SIPServerOptionInfo struct {
	XMLName      xml.Name `xml:"option"`
	Daemon       int      `xml:"daemon"`
	Havesiptrace int      `xml:"havesiptrace"`
	Sipmode      int      `xml:"sipmode"`
}
