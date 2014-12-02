package model

import (
	"encoding/xml"
)

type MTSServerSetupInfo struct {
	XMLName       xml.Name `xml:"sysware"`
	MTSInfo       MTSServerMTSInfo
	SipServerInfo APSServerSipServerInfo
	CMSServerInfo CMSServerInfo
}

type MTSServerMTSInfo struct {
	XMLName              xml.Name `xml:"mts"`
	Id                   string   `xml:"id"`
	Password             string   `xml:"password"`
	RegisterInterval     int      `xml:"registerinterval"`
	DefaultHeartInterval int      `xml:"defaultheartinterval"`
	AddressInfo          AddressInfo
}
