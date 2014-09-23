package model

import (
	"encoding/xml"
)

type SIPServerSetupInfo struct {
	XMLName       xml.Name `xml:"sysware"`
	SIPInfo       SIPServerSIPInfo
	CMSServerInfo CMSServerInfo
	OptionInfo    SIPServerOptionInfo
	ActiveMQInfo  SIPServerActiveMQInfo
}

type SIPServerSIPInfo struct {
	XMLName              xml.Name `xml:"sipserver"`
	Id                   string   `xml:"id"`
	Password             string   `xml:"password"`
	Domain               string   `xml:"domain"`
	DefaultHeartInterval int      `xml:"defaultheartinterval"`
	AddressInfo          AddressInfo
}

type SIPServerActiveMQInfo struct {
	XMLName     xml.Name `xml:"activeMQ"`
	AddressInfo AddressInfo
	Topic       string `xml:"topic"`
}

type SIPServerOptionInfo struct {
	XMLName      xml.Name `xml:"option"`
	Daemon       int      `xml:"daemon"`
	Havesiptrace int      `xml:"havesiptrace"`
	Sipmode      int      `xml:"sipmode"`
}
