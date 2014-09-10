package service

import ()

type GetServerSetupInfoByConfigFile interface {
	GetInfo(cfgFileName string) (interface{}, error)
}
