package service

import ()

type GetServerRequestXMLInfo interface {
	GetXML(cmdType string) string
}

func NewGetServerRequestXMLInfo() GetServerRequestXMLInfo {
	return &getServerRequestXMLInfoImpl{}
}
