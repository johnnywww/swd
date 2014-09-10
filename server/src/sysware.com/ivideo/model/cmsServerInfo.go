package model

import (
	"encoding/xml"
)

type CMSServerInfo struct {
	XMLName xml.Name `xml:"cmsserver"`
	Address string   `xml:"address"`
}
