package model

import (
	"encoding/xml"
)

type AddressInfo struct {
	XMLName xml.Name `xml:"addressinfo"`
	IP      string   `xml:"ip"`
	Port    int      `xml:"port"`
}
