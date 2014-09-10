package service

import ()

type GetUnixSocketCommandXMLInfo interface {
	GetInfo(cmdType string, value interface{}, sockeName string) error
}
