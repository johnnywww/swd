package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/model"
)

type GetProcessInfoLinux struct {
}

func (getProcessInfoLinux *GetProcessInfoLinux) GetInfo(cmdLine string) (*model.ProcessInfo, error) {
	processInfos, err := NewGetAllProcessInfo().GetInfo()
	if nil != err {
		return nil, err
	}
	for _, processInfo := range processInfos {
		if processInfo.RealPath == cmdLine {
			return &processInfo, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("not found process %s", cmdLine))
}
