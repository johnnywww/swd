package service

import (
	"errors"
	"sysware.com/ivideo/utils"
)

type executeLinuxProcImpl struct {
}

func (executeLinuxProcImpl *executeLinuxProcImpl) Exec(cmdLine string, dir string) error {
	if !utils.IsLinuxOS() {
		return errors.New("not linux os")
	}
	executeLinuxProc := executeLinuxProcForkChildImpl{}
	return executeLinuxProc.Exec(cmdLine, dir)
}
