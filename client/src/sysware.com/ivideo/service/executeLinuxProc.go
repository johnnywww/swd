package service

import (
	"sysware.com/ivideo/utils"
)

type ExecuteLinuxProc interface {
	Exec(cmdLine string, dir string) error
}

var executeLinuxProc ExecuteLinuxProc

func init() {
	executeLinuxProc = nil
	if utils.IsLinuxOS() {
		executeLinuxProc = &executeLinuxProcImpl{}
	}
}

func NewExecuteLinuxProc() ExecuteLinuxProc {
	return executeLinuxProc
}
