package model

import (
	"encoding/xml"
)

type APSServerSetupInfo struct {
	XMLName       xml.Name `xml:"sysware"`
	APSInfo       APSServerAPSInfo
	SipServerInfo APSServerSipServerInfo
	CMSServerInfo CMSServerInfo
}

type APSServerAPSInfo struct {
	XMLName              xml.Name `xml:"aps"`
	Id                   string   `xml:"id"`
	Password             string   `xml:"password"`
	RegisterInterval     int      `xml:"registerinterval"`
	DefaultHeartInterval int      `xml:"defaultheartinterval"`
	AddressInfo          AddressInfo
	LogInfo              APSServerLogInfo
}

type APSServerLogInfo struct {
	XMLName     xml.Name `xml:"loginfo"`
	Logcfgfile  string   `xml:"logcfgfile"`
	SaveCatalog string   `xml:"savecatalog"`
}

type APSServerSipServerInfo struct {
	XMLName     xml.Name `xml:"sipserver"`
	Id          string   `xml:"id"`
	Domain      string   `xml:"domain"`
	AddressInfo AddressInfo
}
