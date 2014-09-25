package service

import (
	"os"
	"sysware.com/ivideo/utils"
)

type executeLinuxProcStartProcessImpl struct {
}

func (executeLinuxProcStartProcessImpl *executeLinuxProcStartProcessImpl) Exec(cmdLine string, dir string) error {
	procattr := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: []*os.File{nil, nil, nil}}
	_, err := os.StartProcess(cmdLine, []string{utils.GetFileName(cmdLine)}, &procattr)
	if nil != err {
		return err
	}
	return nil
}
