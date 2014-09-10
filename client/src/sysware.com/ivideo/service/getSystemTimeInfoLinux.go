package service

import (
	"sysware.com/ivideo/utils"
)

type GetSystemTimeInfoLinux struct {
}

func (getSystemTimeInfoLinux *GetSystemTimeInfoLinux) GetInfo() (string, error) {
	osErr := getLinuxOSError()
	if nil != osErr {
		return "", osErr
	}
	return utils.CmdOutputNoLn("/bin/date")
}
