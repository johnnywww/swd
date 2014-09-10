package service

import ()

type GetServerIdInfo interface {
	getInfo(serverName string) (string, error)
}

func NewGetServerIdInfo() GetServerIdInfo {
	return &getServerIdInfoImpl{}
}
