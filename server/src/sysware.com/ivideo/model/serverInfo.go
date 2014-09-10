package model

import (
	"encoding/xml"
)

type ServerInfo struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	Address    string   `xml:"address"`
	Type       int      `xml:"type"`
	Status     int      `xml:"-"`
}

type ServerSaveInfo struct {
	ServerInfo
	OServerName string
	AddFlag     bool
}
